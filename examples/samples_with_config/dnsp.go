//go:build exclude

// dnsp.go script dumps and print all zones (including cache) from the dns server.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/ndr"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/dnsp"
	"github.com/oiweiwei/go-msrpc/msrpc/dnsp/dnsserver/v5"
	"github.com/oiweiwei/go-msrpc/msrpc/dnsp/record"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

var (
	cfg        = config.New()
	zoneName   string
	nodeName   string
	serverName string
)

var P = fmt.Println

func init() {

	config_flag.BindFlags(cfg, flag.CommandLine)

	flag.CommandLine.StringVar(&zoneName, "zone", "", "zone name")
	flag.CommandLine.StringVar(&nodeName, "node", "@", "node name")
	flag.CommandLine.StringVar(&serverName, "server-name", "", "server name")
}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, cfg.ServerAddr(), cfg.DialOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "dial", err)
		return
	}

	defer cc.Close(ctx)

	cli, err := dnsserver.NewDNSServerClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_dnsserver_client", err)
		return
	}

	if zoneName != "" {
		P("-----------------------")
		P("ZONE:", zoneName)
		P("-----------------------")
		EnumRR(ctx, cli, zoneName, nodeName)
		return
	}

	resp, err := cli.ComplexOperation2(ctx, &dnsserver.ComplexOperation2Request{
		ClientVersion: 0x0,
		ServerName:    serverName,
		Operation:     "EnumZones",
		TypeIn:        uint32(dnsp.TypeIDDword),
		DataIn:        &dnsp.Union{Value: &dnsp.Union_Dword{Dword: uint32(dnsp.ZoneRequestFilterAll)}},
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "enum_zones", err)
		return
	}

	for _, zone := range resp.DataOut.GetValue().(*dnsp.ZoneListW2K).ZoneArray {

		P("-----------------------")
		P("ZONE:", zone.ZoneName)
		P("-----------------------")

		EnumRR(ctx, cli, zone.ZoneName, "@")

	}
}

type ZoneNode struct {
	Zone string
	Node string
}

func EnumRR(ctx context.Context, cli dnsserver.DNSServerClient, zone string, n string) {

	next := []*ZoneNode{}

	rr, err := cli.EnumRecords2(ctx, &dnsserver.EnumRecords2Request{
		ServerName: serverName,
		Zone:       zone,
		NodeName:   n,
		RecordType: 0x00FF,
		SelectFlag: 0x0000001F,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, n, zone, err)
		return
	}

	if n == "@" {
		n = ""
	} else {
		n = "." + n
	}

	var nodes record.NodeList

	if err := ndr.Unmarshal(rr.Buffer, &nodes, ndr.Opaque); err != nil {
		fmt.Fprintln(os.Stderr, "umarshal_nodes_list", err)
		return
	}

	for _, node := range nodes.DNSNodes {

		if len(node.DNSRecords) == 0 {
			if node.ChildCount > 0 {
				fmt.Printf("%s\t?\tNODE\t?\t(child: %d)\n", NodeName(node, n, zone), node.ChildCount)
				next = append(next, &ZoneNode{zone, string(node.DNSNodeName.DNSName) + n})
			}
		}

		for _, rr := range node.RRs() {
			if rr.Header().Name == "" {
				rr.Header().Name = NodeName(node, n, zone)
			}
			P(rr)
		}
	}

	for _, nxt := range next {
		P("-----------------------")
		P("ZONE CHILD:", nxt.Node+"."+strings.TrimPrefix(nxt.Zone, "."))
		P("-----------------------")
		EnumRR(ctx, cli, nxt.Zone, nxt.Node)
	}
}

func NodeName(node *record.Node, p, zone string) string {
	n := string(node.DNSNodeName.DNSName)
	if n == "" {
		return NameJoin(p, zone)
	}
	return n
}

func NameJoin(n ...string) string {

	ret := ""

	for _, nn := range n {
		ret += "." + strings.Trim(nn, ".")
	}

	return strings.Trim(ret, ".") + "."
}

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

	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/msrpc/dnsp"
	"github.com/oiweiwei/go-msrpc/msrpc/dnsp/dnsserver/v5"
	"github.com/oiweiwei/go-msrpc/msrpc/dnsp/record"
	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

var (
	zoneName   string
	nodeName   string
	serverName string
)

func init() {
	flag.StringVar(&zoneName, "zone", "", "zone name")
	flag.StringVar(&nodeName, "node", "@", "node name")
	flag.StringVar(&serverName, "server", os.Getenv("SERVER_NAME"), "server name")
	flag.Parse()
}

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	// add mechanism.
	gssapi.AddMechanism(ssp.KRB5)
}

func main() {

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, os.Getenv("SERVER"), epm.EndpointMapper(ctx, os.Getenv("SERVER")))
	if err != nil {
		fmt.Fprintln(os.Stderr, "dcerpc_dial", err)
		return
	}

	defer cc.Close(ctx)

	cli, err := dnsserver.NewDNSServerClient(ctx, cc, dcerpc.WithSign(), dcerpc.WithEndpoint("ncacn_ip_tcp:"), dcerpc.WithTargetName(os.Getenv("TARGET")))
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_dnsserver_client", err)
		return
	}

	if zoneName != "" {
		fmt.Println("-----------------------")
		fmt.Println("ZONE:", zoneName)
		fmt.Println("-----------------------")
		EnumRR(ctx, cli, zoneName, nodeName)
		return
	}

	resp, err := cli.ComplexOperation2(ctx, &dnsserver.ComplexOperation2Request{
		ClientVersion: 0x0,
		ServerName:    serverName,
		Operation:     "EnumZones",
		TypeIn:        uint32(dnsp.TypeIDDword),
		DataIn:        &dnsp.Union{Value: &dnsp.Union_Dword{Dword: 0x00003FFF}},
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "enum_zones", err)
		return
	}

	for _, zone := range resp.DataOut.GetValue().(*dnsp.ZoneListW2K).ZoneArray {

		fmt.Println("-----------------------")
		fmt.Println("ZONE:", zone.ZoneName)
		fmt.Println("-----------------------")

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
			fmt.Println(rr)
		}
	}

	for _, nxt := range next {
		fmt.Println("-----------------------")
		fmt.Println("ZONE CHILD:", nxt.Node+"."+strings.TrimPrefix(nxt.Zone, "."))
		fmt.Println("-----------------------")
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

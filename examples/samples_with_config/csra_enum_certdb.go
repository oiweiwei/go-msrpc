//go:build ignore

// csra_enum_certdb.go performs enumeration of certificate database tables using CSRA interface.
// Usage:
//
// go run csra_enum_certdb.go username%password@ncacn_ip_tcp:dc01.msad.local[privacy] --count=10 --table=req --authority=msad-DC01-CA-1
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/midl/uuid"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom"
	csra_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/csra/client"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/csra/icertadmind/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/csra/icertadmind2/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iactivation/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iobjectexporter/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"

	"github.com/oiweiwei/go-msrpc/msrpc/erref/hresult"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/hresult"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/ntstatus"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"

	. "github.com/oiweiwei/go-msrpc/examples/common"
)

var (
	cfg       = config.New()
	count     int
	authority string
	table     string
)

func init() {
	config_flag.BindFlags(cfg, flag.CommandLine)
	flag.IntVar(&count, "count", 100, "Number of entries to fetch per page")
	flag.StringVar(&authority, "authority", "msad-DC01-CA-1", "CA Authority name")
	flag.StringVar(&table, "table", "request", "table to enumerate (request, extension, attribute, crl)")
}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var tableID uint32

	switch table := strings.ToLower(table); {
	case strings.HasPrefix(table, "req"):
		tableID = 0x00000000
	case strings.HasPrefix(table, "ext"):
		tableID = 0x00003000
	case strings.HasPrefix(table, "attr"):
		tableID = 0x00004000
	case strings.HasPrefix(table, "crl"):
		tableID = 0x00005000
	default:
		fmt.Fprintln(os.Stderr, "unknown table:", table)
		return
	}

	ctx := gssapi.NewSecurityContext(context.Background())

	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dcom/2be2642e-67a1-4690-883b-642b505ddb1d

	// ObjectExporter uses well-known endpoint 135.
	cc, err := dcerpc.Dial(ctx, cfg.ServerAddr(), cfg.DialOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "dial_well_known_endpoint", err)
		return
	}

	defer cc.Close(ctx)

	// new object exporter client.
	cli, err := iobjectexporter.NewObjectExporterClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_object_exporter", err)
		return
	}

	// server-alive to determine the bindings.
	srv, err := cli.ServerAlive2(ctx, &iobjectexporter.ServerAlive2Request{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "server_alive2", err)
		return
	}

	// new activation-client.
	iact, err := iactivation.NewActivationClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_activation_client", err)
		return
	}

	act, err := iact.RemoteActivation(ctx, &iactivation.RemoteActivationRequest{
		ORPCThis: &dcom.ORPCThis{Version: srv.COMVersion},
		ClassID:  dtyp.GUIDFromUUID(uuid.MustParse("d99e6e73-fc88-11d0-b498-00a0c90312f3")), // CertAdminD2
		IIDs:     []*dcom.IID{icertadmind2.CertAdminD2IID},
		// for TCP/IP it must be []uint16{7} / for named pipes: []uint16{15}.
		RequestedProtocolSequences: []uint16{7, 15},
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "remote_activation", err)
		return
	}

	if act.HResult != 0 {
		fmt.Fprintln(os.Stderr, hresult.FromCode(uint32(act.HResult)))
		return
	}

	// dial WMI using OXID bindings. (use ncacn_tcp_ip).
	conn, err := dcerpc.Dial(ctx, cfg.ServerAddr(), append(cfg.DialOptions(ctx), act.OXIDBindings.EndpointsByProtocol("ncacn_ip_tcp")...)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "dial_wmi_endpoint", err)
		return
	}

	defer conn.Close(ctx)

	// establish context that will be shared between NewLevel1LoginClient and
	// NewServicesClient.
	ctx = gssapi.NewSecurityContext(ctx)

	csraCli, err := csra_client.NewClient(ctx, conn, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	csraCli = csraCli.IPID(ctx, act.InterfaceData[0].IPID())

	_, err = csraCli.CertAdminD2().Ping2(ctx, &icertadmind2.Ping2Request{
		This: &dcom.ORPCThis{Version: srv.COMVersion},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	columns, err := csraCli.CertAdminD2().EnumViewColumnTable(ctx, &icertadmind2.EnumViewColumnTableRequest{
		This:        &dcom.ORPCThis{Version: srv.COMVersion},
		Table:       tableID,
		Authority:   authority,
		ColumnCount: 100,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	index := make(map[uint32]string)

	dbColumns, err := columns.ColumnInfo.Columns(int(columns.ColumnOutCount))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var colsOut []uint32

	for i := range dbColumns {
		index[dbColumns[i].Index] = dbColumns[i].Name
		colsOut = append(colsOut, dbColumns[i].Index)
	}

	ret := []map[string]any{}

	view, err := csraCli.CertAdminD2().CertAdminD().OpenView(ctx, &icertadmind.OpenViewRequest{
		This:       &dcom.ORPCThis{Version: srv.COMVersion},
		Authority:  authority,
		ColumnsOut: colsOut,
		ID:         1,
		Count:      uint32(count),
	})

	defer func() {
		_, err := csraCli.CertAdminD2().CertAdminD().CloseView(ctx, &icertadmind.CloseViewRequest{
			This:      &dcom.ORPCThis{Version: srv.COMVersion},
			Authority: authority,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "close_view", err)
		}
	}()

	resultRows := view.ResultRows

	if err != nil && !errors.Is(err, hresult.ErrorArithmeticOverflow) {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	rows, _, err := resultRows.Rows()
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse rows", err)
		return
	}

	for i := range rows {
		o := make(map[string]any)
		for _, col := range rows[i].Columns {
			o["@RowID"] = rows[i].ID
			if val, _ := col.Value(); val != nil {
				o[index[col.Index]] = val
			}
		}
		ret = append(ret, o)
	}

	fmt.Println(J(ret))

}

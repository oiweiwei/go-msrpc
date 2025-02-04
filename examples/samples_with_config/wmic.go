//go:build ignore

// wmic.go performs the fast wmi queries using optimized wco smart enumeration.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iactivation/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iobjectexporter/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmio"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/ienumwbemclassobject/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemfetchsmartenum/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemlevel1login/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemservices/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemwcosmartenum/v0"

	"github.com/oiweiwei/go-msrpc/msrpc/erref/hresult"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/wmi"

	. "github.com/oiweiwei/go-msrpc/examples/common"
)

var (
	cfg      = config.New().UseGlobalCredentials()
	query    string
	resource string
	proto    bool
	forward  bool
	debug    bool
	timeout  time.Duration
	limit    int
	page     int
)

func init() {

	config_flag.BindFlags(cfg, flag.CommandLine)

	flag.StringVar(&query, "query", "SELECT * FROM Win32_Process", "wmic query")
	flag.StringVar(&resource, "resource", "//./root/cimv2", "wmi resource")
	flag.BoolVar(&proto, "proto", false, "return prototype")
	flag.IntVar(&limit, "limit", 0, "limit")
	flag.IntVar(&page, "page", 100, "page")
	flag.BoolVar(&forward, "forward-only", false, "forward-only")
}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, mech := range cfg.Mechanisms() {
		gssapi.AddMechanism(mech)
	}

	for _, cred := range cfg.Credentials() {
		gssapi.AddCredential(cred)
	}

	startTime := time.Now()

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

	// activate the WMI interface.
	act, err := iact.RemoteActivation(ctx, &iactivation.RemoteActivationRequest{
		ORPCThis: &dcom.ORPCThis{Version: srv.COMVersion},
		ClassID:  wmi.Level1LoginClassID.GUID(),
		IIDs:     []*dcom.IID{iwbemlevel1login.Level1LoginIID},
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
	wcc, err := dcerpc.Dial(ctx, cfg.ServerAddr(), append(cfg.DialOptions(ctx), act.OXIDBindings.EndpointsByProtocol("ncacn_ip_tcp")...)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "dial_wmi_endpoint", err)
		return
	}

	defer wcc.Close(ctx)

	// establish context that will be shared between NewLevel1LoginClient and
	// NewServicesClient.
	ctx = gssapi.NewSecurityContext(ctx)

	l1login, err := iwbemlevel1login.NewLevel1LoginClient(ctx, wcc, append(cfg.ClientOptions(ctx), dcom.WithIPID(act.InterfaceData[0].IPID()))...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_level1_login_client", err)
		return
	}

	// login to WMI.
	login, err := l1login.NTLMLogin(ctx, &iwbemlevel1login.NTLMLoginRequest{
		This:            &dcom.ORPCThis{Version: srv.COMVersion},
		NetworkResource: resource,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// start services client.
	svcs, err := iwbemservices.NewServicesClient(ctx, wcc, dcom.WithIPID(login.Namespace.InterfacePointer().IPID()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_services_client", err)
		return
	}

	flags := wmi.QueryFlagType(0)

	if proto {
		flags |= wmi.QueryFlagTypePrototype
	}

	if forward {
		flags |= wmi.QueryFlagType(wmi.GenericFlagTypeForwardOnly)
	}

	now := time.Now()

	enum, err := svcs.ExecQuery(ctx, &iwbemservices.ExecQueryRequest{
		This:          &dcom.ORPCThis{Version: srv.COMVersion},
		QueryLanguage: &oaut.String{Data: "WQL"},
		Query:         &oaut.String{Data: query},
		Flags:         int32(flags),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "wmi_exec_query", query, err)
		return
	}

	if !forward {
		enums, err := ienumwbemclassobject.NewEnumClassObjectClient(ctx, wcc, dcom.WithIPID(enum.Enum.InterfacePointer().IPID()))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		_, err = enums.Reset(ctx, &ienumwbemclassobject.ResetRequest{
			This: &dcom.ORPCThis{Version: srv.COMVersion},
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "wmi_enum_reset", err)
			return
		}
	}

	irem, err := iremunknown2.NewRemoteUnknown2Client(ctx, wcc,
		dcom.WithIPID(act.RemoteUnknown))
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_remote_unknown2_client", err)
		return
	}

	qif, err := irem.RemoteQueryInterface2(ctx, &iremunknown2.RemoteQueryInterface2Request{
		This: &dcom.ORPCThis{Version: srv.COMVersion},
		IPID: enum.Enum.InterfacePointer().IPID().GUID(),
		IIDs: []*dcom.IID{iwbemfetchsmartenum.FetchSmartEnumIID},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "query_wco_smartenum_fetcher_interface", err)
		return
	}

	smart, err := iwbemfetchsmartenum.NewFetchSmartEnumClient(ctx, wcc, dcom.WithIPID(qif.Interface[0].IPID()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_fetch_smart_enum_client", err)
		return
	}

	smartenum, err := smart.GetSmartEnum(ctx, &iwbemfetchsmartenum.GetSmartEnumRequest{
		This: &dcom.ORPCThis{Version: srv.COMVersion},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "fetch_smart_enum", err)
		return
	}

	wco, err := iwbemwcosmartenum.NewWCOSmartEnumClient(ctx, wcc, dcom.WithIPID(smartenum.SmartEnum.InterfacePointer().IPID()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_wco_smart_enum_client", err)
		return
	}

	if limit > 0 && limit < page {
		page = limit
	}

	// classes should store the class definitions across the calls.
	var classes = make(map[string]*wmio.Class)

	count, buffer := 0, 0

	for i := 0; limit == 0 || i < limit; i += page {

		ret, err := wco.Next(ctx, &iwbemwcosmartenum.NextRequest{
			This:    &dcom.ORPCThis{Version: srv.COMVersion},
			Timeout: -1,
			Count:   uint32(page),
		})
		if err != nil {
			if wmi.Status(ret.Return) != wmi.StatusFalse {
				fmt.Fprintln(os.Stderr, "smart_enum_next", err)
				return
			}
		}

		if len(ret.Buffer) == 0 {
			break
		}

		oa, err := wmi.UnmarshalObjectArrayWithClasses(ret.Buffer, classes)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unmarshal_object_array_with_class", err)
			return
		}

		for _, po := range oa.Objects {
			if po.Object.Class != nil {
				fmt.Println(J(po.Object.Properties()))
			} else {
				fmt.Println(J(po.Object.Values()))
			}
		}

		count += len(oa.Objects)
		buffer += len(ret.Buffer)
	}

	fmt.Fprintln(os.Stderr, "buffer fetched:", buffer/1024, "KiB")
	fmt.Fprintln(os.Stderr, "records fetched:", count)
	fmt.Fprintln(os.Stderr, "query execution time:", time.Now().Sub(now))
	fmt.Fprintln(os.Stderr, "script execution time:", time.Now().Sub(startTime))
	PrintMemUsage()

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Fprintf(os.Stderr, "alloc / total / sys / numgc: %v MiB / %v MiB / %v MiB / %v\n",
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys),
		m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

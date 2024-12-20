//go:build ignore

// wmiexec.go is a sample for executing arbitrary WMI class method on the remote machine, examples:

// (1) enumerate sub keys on the remote machine, where 2147483650 is HKEY_LOCAL_MACHINE
// (see https://learn.microsoft.com/en-us/previous-versions/windows/desktop/regprov/enumkey-method-in-class-stdregprov)
//
//		go run wmiexec.go \
//	         --debug \
//				--class StdRegProv \
//				--method EnumKey \
//				--args '{"hDefKey": 2147483650, "sSubKeyName": "SYSTEM\\CurrentControlSet\\Services"}'
//
// (2) cerate a process on the remote machine:
//
//	go run wmiexec.go --class Win32Process --method Create --args '{"CommandLine": "calc.exe", "CurrentDirectory": "C:\\"}'
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iactivation/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iobjectexporter/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmio/query"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmio"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemlevel1login/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemservices/v0"

	"github.com/oiweiwei/go-msrpc/msrpc/erref/hresult"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/wmi"

	. "github.com/oiweiwei/go-msrpc/examples/common"
)

var (
	cfg      = config.New().UseGlobalCredentials()
	class    string
	method   string
	arg      string
	resource string
	args     wmio.Values
)

func init() {

	config_flag.BindFlags(cfg, flag.CommandLine)

	flag.StringVar(&class, "class", "Win32_Process", "class name")
	flag.StringVar(&method, "method", "Create", "method name")
	flag.StringVar(&arg, "args", `{"CommandLine":"calc.exe","CurrentDirectory":"C:\\"}`, "method args")
	flag.StringVar(&resource, "resource", "//./root/cimv2", "resource name, ie root/default, or root/cimv2")

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

	if err := json.Unmarshal([]byte(arg), &args); err != nil {
		fmt.Fprintln(os.Stderr, "parse_args", err)
		return
	}

	log := zerolog.New(os.Stderr)

	if !cfg.Debug {
		log = zerolog.New(io.Discard)
	}

	log.Info().Str("class", class).Str("method", method).Str("args", fmt.Sprintf("%+v", args)).Msg("execute")

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

	if cfg.Debug {
		fmt.Println("-----------------------")
		fmt.Println("OBJECT RESOLVER STRING BINDINGS")
		fmt.Println("-----------------------")

		fmt.Println(J(srv.ObjectResolverBindings.GetStringBindings()))

		fmt.Println("-----------------------")
		fmt.Println("OBJECT RESOLVER SECURITY BINDINGS")
		fmt.Println("-----------------------")

		fmt.Println(J(srv.ObjectResolverBindings.GetSecurityBindings()))
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

	if cfg.Debug {
		fmt.Println("-----------------------")
		fmt.Println("REMOTE ACTIVATION")
		fmt.Println("-----------------------")

		fmt.Println(J(act))

		for i, ifd := range act.InterfaceData {
			fmt.Println("-----------------------")
			fmt.Println(i, "REMOTE ACTIVATION INTERFACE")
			fmt.Println("-----------------------")
			fmt.Println(J(ifd.GetObjectReference()))
		}
	}

	std := act.InterfaceData[0].GetStandardObjectReference().Std

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

	l1login, err := iwbemlevel1login.NewLevel1LoginClient(ctx, wcc, append(cfg.ClientOptions(ctx), dcom.WithIPID(std.IPID))...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
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

	if cfg.Debug {
		fmt.Println(J(login))
	}

	ns := login.Namespace

	// start services client.
	svcs, err := iwbemservices.NewServicesClient(ctx, wcc, dcom.WithIPID(ns.InterfacePointer().IPID()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	builder := query.NewBuilder(ctx, svcs, srv.COMVersion)

	// spawn the Win32_Process class.

	objWithValues := builder.Spawn(class).Method(method).Values(args, wmio.JSONValueToType)

	if cfg.Debug {
		obj, err := objWithValues.Object()
		if err != nil {
			fmt.Fprintln(os.Stderr, "svcs_exec_method", err)
			return
		}
		fmt.Println(J(obj))
	}

	// use simple query builder to execute the Create static method of the Win32_Process.
	out, err := objWithValues.Exec().Object()
	if err != nil {
		fmt.Fprintln(os.Stderr, "svcs_exec_method", err)
		return
	}

	fmt.Println(J(out.Values()))
}

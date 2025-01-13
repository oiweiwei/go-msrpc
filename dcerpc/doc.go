// Package dcerpc implements the DCE/RPC (C706) client with MS-RPCE extensions.
//
// # Manual
//
// The package is used by the generated stub/client code, so any examples
// make sense only within the scope of using generated clients.
//
// One of the key concepts in DCE/RPC is interface. Interface is just set of endopints
// serving certain logical purpose that are handled by server.
//
// From client perspective, interface is represented by presentation context.
// Presentation context is simply a Interface ID (Syntax UUID) and Interface Version
// (Syntax Version) and client-generated unique integer identifier.
// Syntax UUID and Syntax Version are also called as Abstract Sytnax.
//
// When we call NewXXXClient on the connection, generated stub code provides the
// abstract syntax as an option, and then performs the Bind request, which
// creates new Presentation Context (Abstract Syntax + unique identifier) and
// associates it with particular endpoint to which we are connecting.
//
// Endpoint can be either the named pipe name on the IPC$ share, or the TCP/IP
// port.
//
// So, Presentation Context and Endpoint make up a client connection.
//
// Let's begin with simpliest program that connects to EPM and lists all the
// endpoints registered on the server:
//
//	import (
//		"fmt"
//
//		"github.com/oiweiwei/go-msrpc/dcerpc"
//
//		"github.com/oiweiwei/go-msrpc/epm/epm/v3"
//	)
//
//	func main() {
//
//		// create a new connection to the server my-server.com.
//		// note, that no actual connection will be created, since we don't
//		// know at this point, to which port and protocol we should connect.
//		conn, err := dcerpc.Dial(context.TODO(), "my-server.com")
//		if err != nil {
//			// exit.
//		}
//
//		// close the connection (and all opened transports).
//		defer conn.Close(context.TODO())
//
//		// establish connection to the epm client, port 135.
//		// internally, all stubs add WithAbstractSyntax option when establishing the connection.
//		cli, err := epm.NewClient(ctx, conn, dcerpc.WithInsecure(), dcerpc.WithEndpoint(":135"))
//		if err != nil {
//			// exit.
//		}
//
//		// execute lookup request remote procedure call.
//		resp, err := cli.Lookup(context.TODO(), &epm.LookupRequest{MaxEntries: 500})
//		if err != nil {
//			// exit.
//		}
//
//		// print results.
//		for i, entries := range resp.Entries {
//			fmt.Printf("[%d] %s \n", i, entries.Tower.Binding().StringBinding)
//		}
//	}
//
// # Security
//
// The example above doesn't use any security, which is often not true for the real world
// applications.
//
// DCE/RPC provides multiple ways to configure the security level:
//
// * Insecure: no security.
//
// * Packet (not used).
//
// * Connect: establish context but do not provide integrity/confidentiality services.
//
// * Sign: integrity.
//
// * Seal: confidentiality.
//
// In order to establish the security context, first that we need to do is provision the
// SSP packages and credentials. This can be done in two ways. First, provision via global settings.
//
//	import (
//		"github.com/oiweiwei/go-msrpc/ssp"
//		"github.com/oiweiwei/go-msrpc/ssp/credential"
//		"github.com/oiweiwei/go-msrpc/ssp/gssapi"
//	)
//
//	func init() {
//		// This way we add the credentials from the password.
//		gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
//		// This way we provision the security package.
//		gssapi.AddMechanism(ssp.NTLM)
//	}
//
//	func main() {
//
//		// ...
//
//		// create new security context.
//		ctx = gssapi.NewSecurityContext(ctx)
//
//		// We use credentials to establish the NTLM security for Singing the requests/responses.
//		cli, err := epm.NewClient(ctx, conn, dcerpc.WithSign(), dcerpc.WithEndpoint(":135"))
//
//		// ...
//	}
//
// Second, we can provision packages directly when creating the connection.
//
//	cli, err := epm.NewClient(ctx, conn,
//		dcerpc.WithCredential(creds),
//		// only top-level mechanism is used, when this mechanism is SPNEGO, it will
//		// combine all the mechanisms and will perform the negotiation with server.
//		dcerpc.WithMechanism(ssp.SPNEGO),
//		dcerpc.WithMechanism(ssp.KRB5),
//		dcerpc.WithMechanism(ssp.NTLM),
//		dcerpc.WithSeal(),
//		dcerpc.WithEndpoint(":135"))
//
// Or by combination of both approaches (specify mechanisms globally, and credentials locally and so on).
//
//	ctx := gssapi.NewSecurityContext(ctx, ssp.NTLM)
//
//	cli, err := epm.NewClient(ctx, conn, dcerpc.WithCredential(creds), ...)
//
// Security context can be altered by client:
//
//	cli.AlterContext(ctx, dcerpc.WithMechanism(ssp.KRB5), dcerpc.WithSeal())
//
// ## Security Context Use-Cases
//
// As you may have noticed, the security can be configured on many levels, so
// to summarize following is a set of use-cases and proposed approach:
//
// ### Global Configuration
//
// Use init function to establish the mechanism / credentials once and for all connections:
//
//	func init() {
//		gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
//		gssapi.AddMechanism(ssp.NTLM)
//	}
//
// Note, that you cannot add mechanism to the global configuration more than once, as it is not allowed
// by GSSAPI package. If you want to have a freedom of choice for the configuration, use dcerpc.WithMechanism
// per-client options, or NewSecurityContext's gssapi.Option for Reusable Configuration scenario.
//
// ### Reusable Configuration
//
// When you're creating multiple clients over the same endpoint (same named pipe, or TCP port)
// it may be desired to not establish a new security context every time you instantinate the client.
//
// To achieve this, you can define credentials / SSP mechanisms per security context:
//
//	ctx := gssapi.NewSecurityContext(ctx, ssp.NTLM, yourCreds)
//
//	cli1, err := epm.NewEpmClient(
//		ctx, // this context contains NTLM SSP and yourCreds
//		dcerpc.WithSign(), // MUST specify for the initial context establishment.
//		dcerpc.WithEndpoint(":135"))
//
//	cli2, err := iobjectexporter.NewObjectExporterClient(
//		ctx, // this contains already established context.
//		// dcerpc.WithSign(): DON'T specify it, as you are going to use same TCP connection / and context.
//		dcerpc.WithEndpoint(":135"))
//
// As an effect, both cli1 and cli2 will use same security context identifier.
//
// ## Acquire Security Context Attributes
//
// After establishing the security context, you can acquire security attributes from the
// client using connection context:
//
//	epmCli, err := epm.NewClient(ctx, conn, dcerpc.WithSeal(), dcerpc.WithEndpoint(":135"))
//	if err != nil {
//		// handle error
//	}
//
//	key, ok := gssapi.GetAttribute(epmCli.Conn().Context(), gssapi.AttributeSessionKey)
//	if ok {
//		fmt.Printf("Session Key: %x\n", key)
//	}
//
// # Per-Client Configuration
//
// When you wish for each client to have different security context / credentials / mechanism
// you should use dcerpc.WithMechanism and dcerpc.WithCredential:
//
//	cli, err := epm.NewClient(ctx, conn,
//		dcerpc.WithMechanism(ssp.NTLM),
//		dcerpc.WithCredential(creds),
//		dcerpc.WithSeal(),
//		dcerpc.WithEndpoint(":135"))
//
// ## Kerberos
//
// Kerberos uses several environment variables, KRB5_CONFIG to specify the path
// to the kerberos 5 config file, and KRB5_CCACHE for credentials cache path.
//
// The following kerberos configuration should be sufficient to connect to
// the MSAD KDC:
//
//	[realms]
//	CONTOSO.NET = {
//		kdc = win2019-0-1.contoso.net
//		admin_server = win2019-0-1.contoso.net
//	}
//
//	[libdefaults]
//	default_realm = CONTOSO.NET
//	default_tkt_enctypes = rc4-hmac # or aes128-cts-hmac-sha1-96 or aes256-cts-hmac-sha1-96
//	default_tgs_enctypes = rc4-hmac # or aes128-cts-hmac-sha1-96 or aes256-cts-hmac-sha1-96
//
// To acquire the credentials cache you must perform kinit:
//
//	$ KRB5_CONFIG=path/to/config.conf kinit <Username> -c /path/to/output.ccache
//	> Password for <Username>@<Domain>:
//
// The cache can be used via environment variable:
//
//	$ export KRB5_CCACHE=/path/to/output.ccache ./your-dcerpc-prog
//
// Note that kerberos requires valid service principal name, like "host/my-server.com".
//
// # Verification
//
// MS-RPCE provides a feature to include unprotected header parts into the request payload
// and hence add integrity check for them when header signing is not supported.
//
// Use WithVerifyBitMask, WithVerifyHeader2, WithVerifyPresentation options in order
// to include verification trailer to the request.
//
//	cli, err := epm.NewClient(ctx, conn,
//		// ... other opts ...
//		dcerpc.WithVerifyBitMask(true), // required verification
//		dcerpc.WithVerifyPresentation(false), // optional verification
//		dcerpc.WithEndpoint(":135"))
//
// # String Binding
//
// String binding is a special syntax used by DCE/RPC (MS-RPCE) to describe the ways to
// locate the interface.
//
// The string binding syntax is following:
//
//	[ 'ObjectUUID' '@' ] ProtocolSequence ':' NetworkAddress '[' Endpoint ']'
//
// For example:
//
//	"ncacn_ip_tcp:[135]" // TCP/IP on Port 135.
//	"ncacn_np:WIN2019[winreg]" // Named Pipe "winreg" over SMB IPC$ share, WIN2019 is a NetBIOS name.
//
// # Endpont Mapping
//
// Endpoint mapper is not only the service of its own, but also a core component
// in DCE/RPC architecture. The endpoint mapper interface (EPM) is used to determine
// the endpoints (port or named pipes) that are used by interfaces.
//
// Some endpoints are so-called well-known endpoint, that is, they are always the same
// for any system, for example, EPM itself has a well-known endpoint ncacn_ip_tcp:[135], and
// ncacn_np:[epmapper].
//
// You can use the well_known endpoint mapper like following:
//
//	import "github.com/oiweiwei/go-msrpc/well_known"
//
//	// most of the well-known endpoints are SMB endpoints.
//	// (note that SMB will perform significantly slower).
//	conn, err := dcerpc.Dial(ctx, "my-server.com", well_known.EndpointMapper())
//
// But most of the endpoints are dynamic, like DNSServer, DHCPServer, Eventlog TCP/IP endpoints.
// To locate them Endpoint mapper must be used:
//
//	import "github.com/oiweiwei/go-msrpc/epm/epm/v3"
//
//	// specify EPMv3 endpoint mapper, it will use well_known.EndpointMapper()
//	// internally to determine it's own address, by you can also force
//	// it to use SMB by replacing "my-server.com" with "ncacn_np:my-server.com[epmapper]"
//	// or passing dcerpc.WithEndpoint("ncacn_np:[epmapper]") as an option.
//	conn, err := dcerpc.Dial(ctx, "my-server.com", epm.EndpointMapper(ctx, "my-server.com", dcerpc.WithSeal()))
//
// # Error Handling
//
// The "github.com/oiweiwei/go-msrpc/msrpc/erref" package contains the error handlers for
// the list of known errors. If you will import any of the packages inside it, any error
// in `Return` will be automatically matched and converted to human-readable strings.
//
//	// importing win32 errors will convert the error codes into human-readable errors.
//	// go run examples/rrp.go # with erorr
//	// key enumerate: query_info: dcerpc: invoke: /winreg/v1/BaseRegQueryInfoKey: response: decode packet: win32: RPC_X_BAD_STUB_DATA (0x000006f7): The stub received bad data.
//	// go run examples/rrp.go # without error
//	// key enumerate: query_info: dcerpc: invoke: /winreg/v1/BaseRegQueryInfoKey: response: decode packet: error: code: 0x000006f7
//	import _ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
//
// # SMB Performance
//
// Note that using SMB may slow-down the performance, since every request write and response read
// are dedicated SMB Write and SMB Read calls. To improve the performance, SMB transaction support
// must be implemented.
//
// Also note, that the only working security scenario for SMB is WithInsecure and WithSeal, where
// second will also slow-down the performace.
//
// # Examples
//
// See github.com/oiweiwei/go-msrpc/examples for more examples.
package dcerpc

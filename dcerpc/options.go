package dcerpc

import (
	"context"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/rs/zerolog"
)

// Option represents the common interface for DCE/RPC options.
type Option interface {
	is_rpcOption()
}

// CallOption represents the DCE/RPC Call option.
type CallOption interface {
	is_rpcCallOption()
}

// The option set.
type option struct {
	// The security context.
	Security *Security
	// The security context options.
	SecurityOptions []gssapi.ContextOption
	// The flag that indicates whether this is a fresh
	// security context.
	IsNewSecurity bool
	// The association group.
	Group *Group
	// The presentation context list.
	Presentations []*Presentation
	// The list of abstract syntaxes, note, this option
	// is exclusive with Presentations option.
	AbstractSyntaxes []*SyntaxID
	// The transfer syntaxes option. (One possible value is
	// NDR20).
	TransferSyntaxes []*SyntaxID
	// The verification trailer commands.
	Verify []*VerificationCommand
	// The debug logger.
	Logger zerolog.Logger
	// The binding string.
	Bindings []string
}

// HasObjectUUID function returns `true` if set of options
// contains the ObjectUUID option.
func HasObjectUUID(opts any) (*uuid.UUID, bool) {

	switch opts := opts.(type) {
	case []CallOption:
		for i := range opts {
			if opt, ok := (any)(opts[i]).(ObjectUUIDOption); ok {
				return opt(), true
			}
		}
	case []Option:
		for i := range opts {
			if opt, ok := (any)(opts[i]).(ObjectUUIDOption); ok {
				return opt(), true
			}
		}
	}
	return nil, false
}

// SecurityContextOption defines the security options before establishing
// the security context.
type SecurityContextOption func(o *option)

// Option implmentation.
func (SecurityContextOption) is_rpcOption() {}

// WithMechanism option specifies the allowed mechanism for the
// security context.
//
// Use this method to trigger the new security context establishment
// for the client with mechanism provided.
//
//	import "github.com/oiweiwei/go-msrpc/ssp"
//
//	cli, err := winreg.NewWinregClient(ctx, conn, dcerpc.WithSeal(), dcerpc.WithMechanism(ssp.NTLM))
func WithMechanism(m gssapi.Mechanism) SecurityContextOption {
	return SecurityContextOption(func(o *option) {
		o.SecurityOptions = append(o.SecurityOptions, m)
	})
}

// WithCredentials option specifies the credentials set for the
// security context.
//
// Use this method to trigger the new security context establishement
// for the client with credentials provided.
//
//	import "github.com/oiweiwei/go-msrpc/ssp/credential"
//
//	creds := credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
//
//	cli, err := winreg.NewWinregClient(ctx, conn, dcerpc.WithSeal(), dcerpc.WithCredentials(creds))
func WithCredentials(creds gssapi.Credential) SecurityContextOption {
	return SecurityContextOption(func(o *option) {
		o.SecurityOptions = append(o.SecurityOptions, creds)
	})
}

// NoBindOption option indicates that no bind must be performed
// for this connection.
type NoBindOption struct{ Conn Conn }

// Option implementation.
func (NoBindOption) is_rpcOption() {}

// HasNoBind function returns `true` if set of options contains the
// NoBind option.
func HasNoBind(opts any) (Conn, bool) {
	switch opts := opts.(type) {
	case []Option:
		for i := range opts {
			if opt, ok := (any)(opts[i]).(NoBindOption); ok {
				return opt.Conn, true
			}
		}
	}
	return nil, false
}

// WithNoBind option is used to indicate, that no bind must be
// performed to the connection.
func WithNoBind(cc Conn) NoBindOption {
	return NoBindOption{Conn: cc}
}

// The ObjectUUID option.
type ObjectUUIDOption func() *uuid.UUID

// Option interface implementation.
func (ObjectUUIDOption) is_rpcOption() {}

// CallOption interface implementation.
func (ObjectUUIDOption) is_rpcCallOption() {}

// WithObjectUUID option specifies the object UUID for the RPC call.
// For more explicit scenarios, use InvokeObject RPC.
func WithObjectUUID(u *uuid.UUID) ObjectUUIDOption {
	if u == nil {
		u = &uuid.UUID{}
	}
	return ObjectUUIDOption(func() *uuid.UUID {
		return u
	})
}

// BindOption represents the DCE/RPC binding option.
type BindOption func(*option)

func (BindOption) is_rpcOption() {}

// withPresentation option specifies the existing presentation context
// to be used for possible security context negotiation.
func withPresentation(p *Presentation) BindOption {
	return BindOption(func(opt *option) {
		if p != nil && p.id != 0 {
			opt.Presentations = append([]*Presentation{p}, opt.Presentations...)
		}
	})
}

// WithGroup option specifies the association group for the
// connection or is used to initialize the association group id.
func WithGroup(g *Group) BindOption {
	return BindOption(func(opt *option) {
		if g != nil {
			opt.Group = g
		}
	})
}

// WithAbstractSyntax option specifies the abstract syntax for
// the DCE/RPC connection.
func WithAbstractSyntax(abstractSyntax *SyntaxID) BindOption {
	return BindOption(func(opt *option) {
		if abstractSyntax != nil {
			for _, tr := range opt.AbstractSyntaxes {
				if *tr == *abstractSyntax {
					return
				}
			}
			opt.AbstractSyntaxes = append(opt.AbstractSyntaxes, abstractSyntax)
		}
	})
}

// WithNDR20 option specifies the (default) NDR encoding.
func WithNDR20() BindOption {
	return BindOption(func(opt *option) {
		for _, tr := range opt.TransferSyntaxes {
			if *tr == *TransferNDRSyntaxV2_0 {
				return
			}
		}
		opt.TransferSyntaxes = append(opt.TransferSyntaxes, TransferNDRSyntaxV2_0)
	})
}

// WithNDR64 option specifies the NDR64 encoding. (NYI)
func WithNDR64() BindOption {
	return BindOption(func(opt *option) {
		for _, tr := range opt.TransferSyntaxes {
			if *tr == *TransferNDR64SyntaxV1_0 {
				return
			}
		}
		opt.TransferSyntaxes = append(opt.TransferSyntaxes, TransferNDRSyntaxV2_0)
	})
}

// WithVerifyBitMask option includes the bit-mask verification to the
// request verification trailer.
func WithVerifyBitMask(must bool) BindOption {
	return BindOption(func(opt *option) {
		opt.Verify = append(opt.Verify, &VerificationCommand{
			Command:  VerifyBitMask(0),
			Required: must,
		})
	})
}

// WithVerifyHeader2 option includes the header2 verification to the
// request verification trailer.
func WithVerifyHeader2(must bool) BindOption {
	return BindOption(func(opt *option) {
		opt.Verify = append(opt.Verify, &VerificationCommand{
			Command:  &VerifyHeader2{},
			Required: must,
		})
	})
}

// WithVerifyPresentation option includes the presentation verification
// to the request verification trailer.
func WithVerifyPresenetation(must bool) BindOption {
	return BindOption(func(opt *option) {
		opt.Verify = append(opt.Verify, &VerificationCommand{
			Command:  &VerifyPresentation{},
			Required: must,
		})
	})
}

// SecurityOption represents the function on security context.
type SecurityOption func(*Security)

func (SecurityOption) is_rpcOption() {}

func (SecurityOption) is_secOption() {}

// NoHeaderSign option disables the header signing.
func NoHeaderSign() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.RequestHeaderSign = 0
	})
}

// The server can impersonate the client's security context while
// acting on behalf of the client. The server can access local
// resources as the client. If the server is local, it can access
// network resources as the client. If the server is remote, it can
// access only resources that are on the same computer as the server.
func Impersonate() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Impersonation = ImpersonationLevelImpersonate
	})
}

// The most powerful impersonation level. When this level is selected,
// the server (whether local or remote) can impersonate the client's
// security context while acting on behalf of the client. During
// impersonation, the client's credentials (both local and network) can
// be passed to any number of computers.
func Delegate() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Impersonation = ImpersonationLevelDelegate
	})
}

// The server can obtain the client's identity, and the server
// can impersonate the client to do ACL checks.
func Identify() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Impersonation = ImpersonationLevelIdentify
	})
}

// The client is anonymous to the server. The server process
// can impersonate the client, but the impersonation token does
// not contain any information about the client. This level is
// only supported over the local interprocess communication transport.
// All other transports silently promote this level to identify.
func Anonymize() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Impersonation = ImpersonationLevelAnonymous
	})
}

// WithInsecure option specifies the plain-text connection over RPC.
func WithInsecure() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Level = AuthLevelNone
		ctx.Type = AuthTypeNone
	})
}

// WithSign option specifies the connection with packet integrity check.
func WithSign() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Level = AuthLevelPktIntegrity
	})
}

// WithSeal option specifies the connection with packet stub encryption.
func WithSeal() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Level = AuthLevelPktPrivacy
	})
}

// WithSecurityConfig option specifies the mechanism configuration.
//
// Example:
//
//	cli, err := winreg.NewWinregClient(ctx, conn, dcerpc.WithSecurityConfig(&krb5.Config{
//		KRB5ConfPath:       "/tmp/my-krb.conf",
//		CCachePath: 	    "/tmp/ccache",
//		RequireHostAddress: true,
//	}))
func WithSecurityConfig(cfg gssapi.MechanismConfig) SecurityOption {
	return SecurityOption(func(ctx *Security) {
		// select the auth type.
		ctx.Type = MechanismToAuthType(cfg.Type())
		// append the gssapi options.
		ctx.opts = append(ctx.opts, gssapi.WithMechanismType(cfg.Type()), gssapi.WithMechanismConfig(cfg))
	})
}

// WithConnect option sepcifies the connect authentication level.
func WithConnect() SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Level = AuthLevelConnect
	})
}

// WithSecurityLevel option specifies the authenticate security level.
//
// Use this option to specify the security level for the connection.
func WithSecurityLevel(lvl AuthLevel) SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Level = lvl
	})
}

// WithSecurityProvider option specifies the exact security provider.
//
// Use this option to specify the security provider. Note, that security
// provider must be provisioned either globally, or using the GSSAPI
// security context option `WithMechanism`.
//
//	cli, err := winreg.NewWinregClient(ctx, conn, dcerpc.WithSecurityProvider(dcerpc.AuthTypeKerberos), dcerpc.WithSeal())
func WithSecurtyProvider(typ AuthType) SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.Type = typ
	})
}

// WithTargetName option specifies the target name.
//
// Use this option to set the SSP target name:
//
//	cli, err := winreg.NewWinregClient(ctx, conn, dcerpc.WithSeal(), dcerpc.WithTargetName("host/contoso.svc.net"))
func WithTargetName(n string) SecurityOption {
	return SecurityOption(func(ctx *Security) {
		ctx.TargetName = n
	})
}

// WithLogger option sets the debug logger.
//
// Specify this option to turn on the debug logging for the DCE/RPC connection:
//
//	conn, err := dcerpc.Dial(ctx, "contoso.net", dcerpc.WithLogger(zerolog.New(os.Stderr)))
func WithLogger(logger zerolog.Logger) BindOption {
	return BindOption(func(o *option) {
		o.Logger = logger
	})
}

// WithEndpoint option specifies the string binding for the
// connection.
//
// You can specify exact binding string that should be used
// to establish the connection to the server, for example:
//
//	conn, err := dcerpc.Dial(ctx, "contoso.net", dcerpc.WithEndpoint("ncacn_ip_tcp:[48621]")
//
// This code should create a server connection, that will bind any transport
// to the TCP/IP port 48621.
//
// Alternatively, you can specify this option during NewXClient call, for example:
//
//	cli, err := winreg.NewWinregClient(ctx, conn, dcerpc.WithSeal(), dcerpc.WithEndpoint("ncacn_np:[winreg]"))
//
// This code should establish a new transport over a named pipe to the windows registry
// server.
func WithEndpoint(s string) BindOption {
	return BindOption(func(o *option) {
		o.Bindings = append(o.Bindings, s)
	})
}

func ParseOptions(ctx context.Context, opts ...Option) (*option, error) {

	option := &option{}

	for i := range opts {
		// first collect all security context options.
		if o, ok := (opts[i]).(SecurityContextOption); ok {
			o(option)
		}
	}

	for i := range opts {
		switch o := (opts[i]).(type) {
		case SecurityOption:
			if option.Security == nil {
				option.Security, option.IsNewSecurity = NewSecurity(ctx, option.SecurityOptions...), true
			}
			o(option.Security)
		case BindOption:
			o(option)
		}
	}

	if option.Security == nil && len(option.SecurityOptions) > 0 {
		option.Security, option.IsNewSecurity = NewSecurity(ctx, option.SecurityOptions...), true
	}

	if len(option.TransferSyntaxes) == 0 {
		option.TransferSyntaxes = []*SyntaxID{TransferNDRSyntaxV2_0}
	}

	if option.Security == nil {
		cc, ok := gssapi.GetAttribute(ctx, gssapi.AttributeRPCContext)
		if !ok {
			return option, ErrNoSecurityContext
		}
		option.Security = cc.(*Security)
	}

	if len(option.Presentations) == 0 {

		if len(option.AbstractSyntaxes) == 0 {
			return option, ErrNoPresentationContext
		}

		for i := range option.AbstractSyntaxes {
			option.Presentations = append(option.Presentations, NewPresentation(option.AbstractSyntaxes[i]))
		}
	}

	return option, nil
}

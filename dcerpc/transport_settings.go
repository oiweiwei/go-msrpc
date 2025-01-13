package dcerpc

import (
	"context"
	"net"
	"time"

	"github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/ndr"
)

// The Endpoint Mapper interface maps the given syntax identifier
// and returns the list of bindings of following format:
type EndpointMapper interface {
	Map(context.Context, *Binding) ([]StringBinding, error)
}

// The transport settings.
type Transport struct {
	// The receive buffer size.
	MaxRecvFrag int
	// The transmit buffer size.
	MaxXmitFrag int
	// The association group identifier.
	GroupID int
	// The hostname.
	HostName string
	// Clients that implement MS-RPCE extensions SHOULD ignore the
	// secondary endpoint address.
	SecondaryAddr string
	// The data representation for the connection.
	DataRepresentation ndr.DataRepresentation
	// Once concurrent multiplexing on a connection is negotiated,
	// a client is allowed to send another request on a connection
	// before it receives a response on a previous request
	Multiplexing bool
	// The server supports keeping the connection open after an
	// orphaned PDU is received.
	KeepConnOpenOnOrphaned bool
	// If this flag is set to `true` it is allowed for a client
	// to use more than one security context per connection.
	SecurityContextMultiplexing bool
	// The number of security contexts established for this
	// connection.
	SecurityContextCount int
	// The number of allowed outstanding calls.
	MultiplexingOutstandingCalls int
	// The network interaction timeout.
	Timeout time.Duration
	// The programmable operation deadline for deecoding and
	// encryption/decryption.
	Deadline time.Duration
	// Network dialer.
	Dialer Dialer
	// SMB port.
	SMBPort int
	// SMB dialer.
	SMBDialer any
	// Endpoint Mapper.
	EndpointMapper EndpointMapper
	// Preferred protocol sequence.
	StringBinding StringBinding
	// If set to `true`, new connection will be established
	// for every new client with matching binding.
	NoReuseTransport bool
}

// The transport connection option.
type ConnectOption func(*Transport)

func (ConnectOption) is_rpcOption() {}

// WithBufferSize options sets the buffer input/output size.
func WithFragmentSize(sz int) ConnectOption {
	return func(o *Transport) {
		if sz >= MinimumXmitSize { // minimum size.
			o.MaxXmitFrag, o.MaxRecvFrag = sz, sz
		}
	}
}

func WithNewTransport() ConnectOption {
	return func(o *Transport) { o.NoReuseTransport = true }
}

// WithGroupID option sets the association group identifier.
func WithGroupID(id int) ConnectOption {
	return func(o *Transport) { o.GroupID = id }
}

// Dialer interface is used to dial a TCP connection.
type Dialer interface {
	// DialContext dials a TCP connection.
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

// WithDialer option sets the network dialer dialer for TCP and SMB connections.
func WithDialer(dialer Dialer) ConnectOption {
	return func(o *Transport) { o.Dialer = dialer }
}

// WithTimeout option sets the networking timeout.
func WithTimeout(timeout time.Duration) ConnectOption {
	return func(o *Transport) { o.Timeout = timeout }
}

// WithSMBPort function sets the SMB communication port.
func WithSMBPort(port int) ConnectOption {
	return func(o *Transport) { o.SMBPort = port }
}

// WithSMBDialer function sets the SMB dialer.
func WithSMBDialer(dialer any) ConnectOption {
	return func(o *Transport) { o.SMBDialer = dialer }
}

// WithEndpointMapper option sets the endpoint mapper to find the endpoint
// (port or named pipe) for the selected abstract syntax.
//
// There are two interface implementation, first is EPM service:
//
//	import "github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
//
//	conn, err := dcerpc.Dial(ctx, "contoso.net", epm.EndpointMapper(ctx, "contoso.net", dcerpc.WithSign()))
//
// Second is well-known endpoints:
//
//	import "github.com/oiweiwei/go-msrpc/msrpc/well_known"
//
//	conn, err := dcerpc.Dial(ctx, "contoso.net", well_known.EndpointMapper())
func WithEndpointMapper(m EndpointMapper) ConnectOption {
	return func(o *Transport) { o.EndpointMapper = m }
}

// NewTransport function returns the default transport configuration.
func NewTransport() Transport {
	return Transport{
		MaxRecvFrag:                  DefaultXmitSize,
		MaxXmitFrag:                  DefaultXmitSize,
		DataRepresentation:           ndr.DefaultDataRepresentation,
		MultiplexingOutstandingCalls: 64,
		Timeout:                      10 * time.Second,
		Deadline:                     3 * time.Second,
		SMBPort:                      445,
	}
}

func (s Transport) MarshalZerologObject(e *zerolog.Event) {
	e.Int("max_xmit_frag", s.MaxXmitFrag)
	e.Int("max_recv_frag", s.MaxRecvFrag)
	e.Int("group_id", s.GroupID)
	e.Str("secondary_addr", s.SecondaryAddr)
	e.Bool("multiplexing", s.Multiplexing)
	e.Bool("security_context_multiplexing", s.SecurityContextMultiplexing)
	e.Bool("keep_conn_on_orphaned", s.KeepConnOpenOnOrphaned)
	e.Dur("network_timeout", s.Timeout)
	e.Dur("programmable_deadline", s.Deadline)

}

// IsSecurityMultiplexed function returns `true` if security multiplexing is enabled
// and there are more than one security context captured.
func (s Transport) IsSecurityMultiplexed() bool {
	return s.SecurityContextMultiplexing && s.SecurityContextCount > 1
}

// BufferSize function returns the biggest number between max_recv_frag and
// max_xmit_frag.
func (s Transport) FragmentSize() int {
	if s.MaxRecvFrag >= s.MaxXmitFrag {
		return s.MaxRecvFrag
	}
	return s.MaxXmitFrag
}

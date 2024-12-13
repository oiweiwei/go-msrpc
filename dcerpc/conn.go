package dcerpc

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/smb2"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/rs/zerolog"
)

// BufferedConn is a raw connection wrapper to optimize the short
// read header operations.
type BufferedConn struct {
	// The wrapped raw connection.
	RawConn
	// cur for current position, total is buffer start position.
	cur, total []byte
}

func (conn *BufferedConn) Resized(sz int) *BufferedConn {
	if sz > len(conn.total) {
		conn.total = make([]byte, sz)
	}
	conn.cur, conn.total = nil, conn.total[:sz]
	return conn
}

// NewBufferedConn function returns the new buffered connection.
func NewBufferedConn(cc RawConn, sz int) RawConn {
	return &BufferedConn{RawConn: cc, cur: nil, total: make([]byte, sz)}
}

// Read function reads the data into buffer `b` and optionally fetches
// the next data block.
func (conn *BufferedConn) Read(b []byte) (int, error) {

	want := len(b)

	if len(b) <= len(conn.cur) {
		// copy current to bytes and advance conn.cur to b lenght.
		conn.cur = conn.cur[copy(b, conn.cur):]
		return want, nil
	}

	if len(conn.cur) > 0 {
		// copy all bytes to b.
		b = b[copy(b, conn.cur):]
	}

	for len(b) > 0 {
		// reset current buffer.
		conn.cur = conn.total

		// read available data.
		n, err := conn.RawConn.Read(conn.cur)
		if err != nil {
			return n, err
		}

		// limit current buffer to read chunk.
		conn.cur = conn.cur[:n]
		// copy data to buffer.
		n = copy(b, conn.cur)
		// advance the buffer.
		conn.cur, b = conn.cur[n:], b[n:]
	}

	return want, nil
}

// conn represents the client's set of transports
// per binding.
type conn struct {
	mu sync.Mutex
	// The server address.
	serverAddr string
	// The association group identifier.
	group *Group
	// The transport set settings.
	settings *Transport
	// The set of transports per binding-string.
	transports map[string][]*transport
	// The logger.
	logger zerolog.Logger
	// options.
	opts []Option
}

// Dial function creates a new transport set that should be used for
// the binding procedures performed by clients.
func Dial(ctx context.Context, addr string, opts ...Option) (Conn, error) {
	settings := NewTransport()
	// apply connect options.
	for _, o := range opts {
		o, ok := o.(ConnectOption)
		if !ok {
			continue
		}
		o(&settings)
	}
	// establish new association group.
	group := &Group{}
	group.SetID(settings.GroupID)
	// new transport set.
	tr := &conn{
		group:      group,
		serverAddr: addr,
		settings:   &settings,
		transports: make(map[string][]*transport),
		opts:       append(opts, WithGroup(group)),
	}

	if net.ParseIP(addr) != nil {
		// this is an ip address.
		return tr, nil
	}

	if !strings.Contains(addr, ":") {
		tr.settings.HostName = addr
		// this is a fqdn.
		ips, err := net.LookupIP(addr)
		if err != nil {
			return nil, fmt.Errorf("dial: lookup server address: %w", err)
		}
		for _, ip := range ips {
			if ip.To4() != nil {
				tr.serverAddr = ip.String()
				return tr, nil
			}
		}
		return nil, fmt.Errorf("dial: lookup server address: no ipv4 address")
	}

	// address is the string binding (ie, ncacn_ip_tcp:127.0.0.1,
	// or, ncacn_np:127.0.0.1[\PIPE\pipe_name]).
	binding, err := ParseStringBinding(addr)
	if err != nil {
		return nil, fmt.Errorf("dial: %s: %w", addr, err)
	}

	// set the string binding.
	tr.settings.StringBinding = *binding
	// set the server address.
	if tr.serverAddr = binding.NetworkAddress; tr.serverAddr == "" {
		tr.serverAddr = binding.ComputerName
	}

	// return the transport set.
	return tr, nil
}

// Alter Context.
func (t *conn) AlterContext(_ context.Context, _ ...Option) error {
	return fmt.Errorf("alter context: the transport is not binded")
}

// Invoke.
func (t *conn) Invoke(_ context.Context, _ Operation, _ ...CallOption) error {
	return fmt.Errorf("invoke: connection is not binded")
}

// Invoke Object.
func (t *conn) InvokeObject(_ context.Context, _ *uuid.UUID, _ Operation, _ ...CallOption) error {
	return fmt.Errorf("invoke_object: connection is not binded")
}

// Close.
func (t *conn) Close(ctx context.Context) error {

	t.mu.Lock()
	defer t.mu.Unlock()

	for _, transports := range t.transports {
		for i := range transports {
			if err := transports[i].shutdown(ctx); err != nil {
				t.logger.Err(err).Msg("shutdown transport")
			}
		}
	}

	t.transports = make(map[string][]*transport)
	return nil
}

// RegisterServer.
func (t *conn) RegisterServer(_ ServerHandle, _ ...Option) {
	// TODO: NYI
}

func (t *conn) Bind(ctx context.Context, opts ...Option) (Conn, error) {

	t.mu.Lock()
	defer t.mu.Unlock()

	opts = append(t.opts, opts...)

	// parse selected option.
	o, err := ParseOptions(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("bind: parse options: %w", err)
	}

	t.logger = o.Logger

	var bindings []StringBinding

	if len(o.Bindings) > 0 {
		// first check the bindings provided by WithEndpoint parameter.
		for i := range o.Bindings {
			binding, err := ParseStringBinding(o.Bindings[i])
			if err != nil {
				return nil, fmt.Errorf("bind: parse string binding: %w", err)
			}

			if !binding.Complete() {
				// try to complete the binding with endpoint mapper.
				if t.settings.EndpointMapper != nil {
					bs, err := t.settings.EndpointMapper.Map(ctx, &Binding{
						SyntaxID:      *o.AbstractSyntaxes[0],
						StringBinding: *binding,
					})
					if err != nil {
						return nil, fmt.Errorf("bind: endpoint mapper: %w", err)
					}
					if len(bs) > 0 {
						binding = &bs[0]
					}
				}
			}
			bindings = append(bindings, *binding)
		}
	} else if t.settings.StringBinding.Complete() {
		// check the binding encoded into the server address.
		bindings = []StringBinding{t.settings.StringBinding}
	} else {
		// use endpoint mapper to retrieve the bindings.
		if t.settings.EndpointMapper != nil {
			// figure out the string binding from the endpoint mapper.
			if bindings, err = t.settings.EndpointMapper.Map(ctx, &Binding{
				SyntaxID:      *o.AbstractSyntaxes[0],
				StringBinding: t.settings.StringBinding,
			}); err != nil {
				return nil, fmt.Errorf("bind: endpoint mapper: %w", err)
			}
		}
	}

	target := o.TargetBinding()
	if target == "" {
		target = t.serverAddr
	}

	for i, j := 0, 0; i < len(bindings); i++ {
		// bubble-sort round to prioritize matching target names.
		if bindings[i].MatchTarget(target) {
			bindings[i], bindings[j] = bindings[j], bindings[i]
			j++
		}
	}

	t.logger.Debug().Str("target_name", target).Interface("bindings", bindings).Msgf("found %d bindings", len(bindings))

	var selected []*transport

	if !t.settings.NoReuseTransport {
		for i := range bindings {
			// find already established transport.
			if selected = t.transports[bindings[i].String()]; len(selected) > 0 {
				// check if connections are alive.
				active := false
				for j := range selected {
					if selected[j].err == nil {
						active = true
						break
					}
				}
				// found alive connection, breaking.
				if active {
					t.logger.Debug().Msgf("found established transport for binding %s", bindings[i])
					break
				}

			}
		}
	}

	if len(selected) == 0 {
		t.logger.Debug().Msgf("no established transport was found")
		for i := range bindings {
			t.logger.Debug().Msgf("etablishing new transport for binding %s", bindings[i])
			if selected, err = t.dial(ctx, bindings[i]); err != nil {
				t.logger.Error().Err(err).Msgf("bind: dial %s error", bindings[i])
				continue
			}
			t.logger.Debug().Msgf("new transport for binding %s has been successfully established", bindings[i])
			t.transports[bindings[i].String()] = append(t.transports[bindings[i].String()], selected...)
			break
		}
	}

	var conn Conn

	for i := range selected {
		// skip dead connections.
		if selected[i].err != nil {
			continue
		}
		t.logger.Debug().Msgf("binding the selected transport")
		if conn, err = selected[i].Bind(ctx, opts...); err != nil {
			t.logger.Err(err).Msgf("selected transport error")
			continue
		}
		t.logger.Debug().Msgf("selected transport has been successfully binded")
		return conn, nil
	}

	if err != nil {
		return nil, fmt.Errorf("bind: could not bind the selected transport: %w", err)
	}

	return nil, fmt.Errorf("bind: could not find matching binding")
}

func (t *conn) dial(ctx context.Context, binding StringBinding) ([]*transport, error) {

	conn, err := t.dialConn(ctx, binding)
	if err != nil {
		return nil, err
	}

	settings := *t.settings

	return []*transport{{
		id:       rand.Int(),
		cc:       NewBufferedConn(conn, t.settings.MaxRecvFrag),
		settings: &settings,
		tx:       make([]byte, t.settings.MaxXmitFrag),
		rx:       make([]byte, t.settings.MaxRecvFrag),
		txQ:      make(chan *call),
		rxQ:      make(chan *call, 64),
		logger:   t.logger,
		conn:     t,
	}}, nil
}

func (t *conn) dialConn(ctx context.Context, binding StringBinding) (RawConn, error) {

	switch binding.ProtocolSequence {
	case ProtocolSequenceIPTCP:

		addr := net.JoinHostPort(binding.NetworkAddress, binding.Endpoint)

		if binding.NetworkAddress == "" || binding.NetworkAddress == "0.0.0.0" {
			addr = net.JoinHostPort(t.serverAddr, binding.Endpoint)
		}

		t.logger.Debug().Msgf("dialing tcp %s", addr)

		conn, err := net.DialTimeout("tcp", addr, t.settings.Timeout)
		if err != nil {
			return nil, fmt.Errorf("ncacn_ip_tcp: %w", err)
		}

		t.logger.Debug().Msgf("dialing tcp %s done", addr)

		return conn, nil

	case ProtocolSequenceNamedPipe:

		t.logger.Debug().Msgf("dialing smb named pipe %s:%d:%s\\%s",
			t.serverAddr, t.settings.SMBPort, binding.ShareName(), binding.NamedPipe())

		dialer, err := smb2.CompatDialer(t.settings.SMBDialer)
		if err != nil {
			return nil, fmt.Errorf("ncacn_np: %w", err)
		}

		if dialer == nil {
			o := ParseSecurityOptions(ctx, t.opts...).Security
			if o != nil {
				dialer = smb2.NewDialer(
					smb2.WithSecurity(
						// set target name by default if available.
						gssapi.WithTargetName(o.TargetName),
					),
				)
			} else {
				dialer = smb2.NewDialer(smb2.WithSecurity())
			}
		}

		pipe := &smb2.NamedPipe{
			Logger:    t.logger,
			Address:   t.serverAddr,
			Port:      t.settings.SMBPort,
			Timeout:   t.settings.Timeout,
			Dialer:    dialer,
			ShareName: binding.ShareName(),
			Name:      binding.NamedPipe(),
		}

		if err := pipe.Connect(ctx); err != nil {
			return nil, fmt.Errorf("ncacn_np: %w", err)
		}

		t.logger.Debug().Msgf("dialing smb named pipe done")

		return pipe, nil
	}

	return nil, fmt.Errorf("ncacn: %s: not supported", binding.String())
}

func (c *conn) closeTransport(ctx context.Context, tr *transport) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	for k, transports := range c.transports {
		for i := 0; i < len(transports); i++ {
			if transports[i].id == tr.id {
				c.transports[k] = append(transports[:i], transports[i+1:]...)
				return nil
			}
		}
	}

	return fmt.Errorf("transport not found")
}

package smb2

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/go-smb2.fork"
	"github.com/rs/zerolog"
)

type Dialect uint16

const (
	// SMB2.0.2 dialects.
	SMB202 Dialect = 0x0202
	// SMB2.1 dialects.
	SMB210 Dialect = 0x0210
	// SMB3.0 dialects.
	SMB300 Dialect = 0x0300
	// SMB3.0.2 dialects.
	SMB302 Dialect = 0x0302
	// SMB3.1.1 dialects.
	SMB311 Dialect = 0x0311
)

// DialerOption is a function that configures a SMB2/3 Dialer.
type DialerOption func(*smb2.Dialer)

// WithDialect returns a DialerOption that configures a SMB2/3 Dialer to
// use the specified dialect.
func WithDialect(dialect Dialect) DialerOption {
	return func(d *smb2.Dialer) {
		d.Negotiator.SpecifiedDialect = uint16(dialect)
	}
}

// WithSign returns a DialerOption that configures a SMB2/3 Dialer to
// require message signing.
func WithSign() DialerOption {
	return func(d *smb2.Dialer) {
		d.Negotiator.RequireMessageSigning = true
	}
}

// WithSeal returns a DialerOption that configures a SMB2/3 Dialer to
// use message encryption whenever possible.
func WithSeal() DialerOption {
	return func(d *smb2.Dialer) {
		d.Negotiator.UseMessageEncryption = true
	}
}

// WithSecurity returns a DialerOption that configures a SMB2/3 Dialer to
// use the specified GSSAPI security context.
func WithSecurity(opts ...gssapi.ContextOption) DialerOption {
	return func(d *smb2.Dialer) {
		d.Initiator = SecurtiyContextInitiator(opts...)
	}
}

func NewDialer(opts ...DialerOption) *smb2.Dialer {
	d := &smb2.Dialer{}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

type NamedPipe struct {
	*smb2.File
	Logger          zerolog.Logger
	Address         string
	Port            int
	Timeout         time.Duration
	Dialer          *smb2.Dialer
	NetworkDialFunc func(ctx context.Context, network, address string) (net.Conn, error)
	ShareName       string
	Name            string
	Share           *smb2.Share
}

func (pipe *NamedPipe) Dialect() Dialect {
	return Dialect(pipe.Dialer.Negotiator.SpecifiedDialect)
}

func (pipe *NamedPipe) SessionKey() []byte {
	return pipe.Dialer.Initiator.SessionKey()
}

func (pipe *NamedPipe) ApplicationKey() []byte {
	return pipe.Share.ApplicationKey()
}

const ErrNotActive = "An instance of a named pipe cannot be found in the listening state"

func (pipe *NamedPipe) dial(ctx context.Context, addr string) (net.Conn, error) {
	if pipe.NetworkDialFunc != nil {
		return pipe.NetworkDialFunc(ctx, "tcp", addr)
	}
	return net.DialTimeout("tcp", addr, pipe.Timeout)
}

func (pipe *NamedPipe) Connect(ctx context.Context) error {

	addr := net.JoinHostPort(pipe.Address, strconv.Itoa(pipe.Port))

	conn, err := pipe.dial(ctx, addr)
	if err != nil {
		return fmt.Errorf("dial smb server: %s: %w", addr, err)
	}

	session, err := pipe.Dialer.DialContext(ctx, conn)
	if err != nil {
		return fmt.Errorf("open smb session: %w", err)
	}

	pipe.Share, err = session.Mount(pipe.ShareName)
	if err != nil {
		return fmt.Errorf("mount share: %w", err)
	}

	for {
		pipe.File, err = pipe.Share.OpenFile(pipe.Name, os.O_RDWR, 0666)
		if err != nil {
			if strings.Contains(err.Error(), ErrNotActive) {
				pipe.Logger.Err(err).Msgf("open share file %s", pipe.Name)
				continue
			}
			return fmt.Errorf("open file: %w", err)
		}
		break
	}

	return err
}

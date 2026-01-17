package dcerpc

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/ndr"
)

var (
	// Transfer Syntax NDR UUID.
	TransferNDR = uuid.New(0x8A885D04, 0x1CEB, 0x11C9, 0x9F, 0xE8, [6]byte{0x08, 0x00, 0x2B, 0x10, 0x48, 0x60})
	// Transfer Syntax NDR v2.0.
	TransferNDRSyntaxV2_0 = &SyntaxID{TransferNDR, 2, 0}

	// Transfer Syntax NDR64 UUID.
	TransferNDR64 = uuid.New(0x71710533, 0xBEBA, 0x4937, 0x83, 0x19, [6]uint8{0xB5, 0xDB, 0xEF, 0x9C, 0xCC, 0x36})
	// Transfer Syntax NDR64 v1.0.
	TransferNDR64SyntaxV1_0 = &SyntaxID{TransferNDR64, 1, 0}

	// Bind-time feature negotiation flags.
	BindFlags = KeepConnOpenOnOrphaned | SecurityContextMultiplexing
	// Bind-time feature negotiation UUID.
	BindFeature = uuid.New(0x6CB71C2C, 0x9812, 0x4540, uint8(BindFlags), 0, [6]byte{})
	// Bind-time feature negotiation Syntax.
	BindFeatureSyntaxV1_0 = &SyntaxID{BindFeature, 1, 0}
)

// NewBindFeatureSyntaxV1_0 function returns the bind-feature negotiation flags.
func NewBindFeatureSyntaxV1_0(flags ProviderReason) *SyntaxID {
	return &SyntaxID{uuid.New(0x6CB71C2C, 0x9812, 0x4540, uint8(flags&BindFlags), 0, [6]byte{}), 1, 0}
}

// The default transmit/receive size.
const DefaultXmitSize = 4096

// The MinimumXmitSize.
const MinimumXmitSize = 4096

// The DCE/RPC Operation.
type Operation = ndr.Operation

// The Raw connection.
type RawConn interface {
	io.ReadWriteCloser
}

// The DCE/RPC Connection.
type Conn interface {
	Bind(context.Context, ...Option) (Conn, error)
	// Alter Context.
	AlterContext(context.Context, ...Option) error
	// Context.
	Context() context.Context
	// Invoke.
	Invoke(context.Context, Operation, ...CallOption) error
	// Invoke Object.
	InvokeObject(context.Context, *uuid.UUID, Operation, ...CallOption) error
	// Close.
	Close(context.Context) error
	// RegisterServer.
	RegisterServer(ServerHandle, ...Option)
	// Error.
	Error(context.Context, any) error
}

var (
	// Server terminated the connection.
	ErrShutdown = errors.New("server terminated")
	// UUID mismatch.
	ErrObjectUUIDMismatch = errors.New("object uuid mismatch")
	// Security Context is empty.
	ErrNoSecurityContext = errors.New("security context is empty")
	// Presentation Context is empty.
	ErrNoPresentationContext = errors.New("presentation context is empty")
)

// ServerHandle is a function that must accept the incoming reader and operation
// number and return the operation response.
type ServerHandle func(context.Context, int, ndr.Reader) (Operation, error)

var ErrNotImplemented = errors.New("not implemented")

// Group is used to hold the association group identifier.
// The association can be passed through the options using multiple
// connections, thus sharing the same group id between them.
type Group struct {
	mu sync.Mutex
	id int
}

// GroupID function returns the association group
// (or 0 if association is null).
func (a *Group) GroupID() int {
	if a != nil {
		a.mu.Lock()
		defer a.mu.Unlock()
		return a.id
	}
	return 0
}

// SetID function sets the association group id (when present)
// and returns that identifier.
func (a *Group) SetID(id int) int {
	if a != nil {
		a.mu.Lock()
		defer a.mu.Unlock()
		if id != 0 {
			a.id = id
		}
		return a.id
	}
	return id
}

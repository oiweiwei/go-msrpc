package iobjectexporter

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
)

// IObjectExporter server interface.
type ObjectExporterServer interface {

	// The ResolveOxid method returns the bindings and Remote Unknown IPID for an object
	// exporter.
	ResolveOXID(context.Context, *ResolveOXIDRequest) (*ResolveOXIDResponse, error)

	// The SimplePing method performs a ping of a previously allocated ping set to maintain
	// the reference counts on the objects referred to by the set.
	SimplePing(context.Context, *SimplePingRequest) (*SimplePingResponse, error)

	// The ComplexPing (Opnum 2) method is invoked to create or modify a ping set, to ping
	// a ping set, or to perform a combination of these operations in one invocation.
	ComplexPing(context.Context, *ComplexPingRequest) (*ComplexPingResponse, error)

	// The ServerAlive (Opnum 3) method is used by clients to test the aliveness of the
	// server using a given RPC protocol. If it returns without an error, the server is
	// assumed to be reachable.
	ServerAlive(context.Context, *ServerAliveRequest) (*ServerAliveResponse, error)

	// The ResolveOxid2 method returns the bindings and Remote Unknown IPID for an object
	// exporter, and the COMVERSION of the object server. This method was introduced with
	// version 5.2 of the DCOM Remote Protocol.
	ResolveOxid2(context.Context, *ResolveOxid2Request) (*ResolveOxid2Response, error)

	// The ServerAlive2 (Opnum 5) method was introduced with version 5.6 of the DCOM Remote
	// Protocol. This method extends the ServerAlive method. It returns string and security
	// bindings for the object resolver, which allows the client to choose the most appropriate,
	// mutually compatible settings.
	ServerAlive2(context.Context, *ServerAlive2Request) (*ServerAlive2Response, error)
}

func RegisterObjectExporterServer(conn dcerpc.Conn, o ObjectExporterServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewObjectExporterServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ObjectExporterSyntaxV0_0))...)
}

func NewObjectExporterServerHandle(o ObjectExporterServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ObjectExporterServerHandle(ctx, o, opNum, r)
	}
}

func ObjectExporterServerHandle(ctx context.Context, o ObjectExporterServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // ResolveOxid
		in := &ResolveOXIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResolveOXID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // SimplePing
		in := &SimplePingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SimplePing(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // ComplexPing
		in := &ComplexPingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ComplexPing(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // ServerAlive
		in := &ServerAliveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerAlive(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // ResolveOxid2
		in := &ResolveOxid2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResolveOxid2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // ServerAlive2
		in := &ServerAlive2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerAlive2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}

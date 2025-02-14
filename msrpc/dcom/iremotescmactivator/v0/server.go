package iremotescmactivator

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

// IRemoteSCMActivator server interface.
type RemoteSCMActivatorServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// The RemoteGetClassObject (Opnum 3) method is used by clients to create an object
	// reference for the class factory object.
	RemoteGetClassObject(context.Context, *RemoteGetClassObjectRequest) (*RemoteGetClassObjectResponse, error)

	// The RemoteCreateInstance (Opnum 4) method is used by clients to create an object
	// reference for the actual object.
	RemoteCreateInstance(context.Context, *RemoteCreateInstanceRequest) (*RemoteCreateInstanceResponse, error)
}

func RegisterRemoteSCMActivatorServer(conn dcerpc.Conn, o RemoteSCMActivatorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteSCMActivatorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteSCMActivatorSyntaxV0_0))...)
}

func NewRemoteSCMActivatorServerHandle(o RemoteSCMActivatorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteSCMActivatorServerHandle(ctx, o, opNum, r)
	}
}

func RemoteSCMActivatorServerHandle(ctx context.Context, o RemoteSCMActivatorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // Opnum2NotUsedOnWire
		// Opnum2NotUsedOnWire
		return nil, nil
	case 3: // RemoteGetClassObject
		op := &xxx_RemoteGetClassObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteGetClassObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteGetClassObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RemoteCreateInstance
		op := &xxx_RemoteCreateInstanceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteCreateInstanceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteCreateInstance(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteSCMActivator
type UnimplementedRemoteSCMActivatorServer struct {
}

func (UnimplementedRemoteSCMActivatorServer) RemoteGetClassObject(context.Context, *RemoteGetClassObjectRequest) (*RemoteGetClassObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteSCMActivatorServer) RemoteCreateInstance(context.Context, *RemoteCreateInstanceRequest) (*RemoteCreateInstanceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteSCMActivatorServer = (*UnimplementedRemoteSCMActivatorServer)(nil)

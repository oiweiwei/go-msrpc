package imsmqapplication

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IMSMQApplication server interface.
type ImsmqApplicationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// MachineIdOfMachineName operation.
	MachineIDOfMachineName(context.Context, *MachineIDOfMachineNameRequest) (*MachineIDOfMachineNameResponse, error)
}

func RegisterImsmqApplicationServer(conn dcerpc.Conn, o ImsmqApplicationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqApplicationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqApplicationSyntaxV0_0))...)
}

func NewImsmqApplicationServerHandle(o ImsmqApplicationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqApplicationServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqApplicationServerHandle(ctx context.Context, o ImsmqApplicationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // MachineIdOfMachineName
		op := &xxx_MachineIDOfMachineNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MachineIDOfMachineNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MachineIDOfMachineName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQApplication
type UnimplementedImsmqApplicationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqApplicationServer) MachineIDOfMachineName(context.Context, *MachineIDOfMachineNameRequest) (*MachineIDOfMachineNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqApplicationServer = (*UnimplementedImsmqApplicationServer)(nil)

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
type ApplicationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// MachineIdOfMachineName operation.
	MachineIDOfMachineName(context.Context, *MachineIDOfMachineNameRequest) (*MachineIDOfMachineNameResponse, error)
}

func RegisterApplicationServer(conn dcerpc.Conn, o ApplicationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewApplicationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ApplicationSyntaxV0_0))...)
}

func NewApplicationServerHandle(o ApplicationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ApplicationServerHandle(ctx, o, opNum, r)
	}
}

func ApplicationServerHandle(ctx context.Context, o ApplicationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedApplicationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedApplicationServer) MachineIDOfMachineName(context.Context, *MachineIDOfMachineNameRequest) (*MachineIDOfMachineNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ApplicationServer = (*UnimplementedApplicationServer)(nil)

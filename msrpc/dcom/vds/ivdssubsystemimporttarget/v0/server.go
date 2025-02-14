package ivdssubsystemimporttarget

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IVdsSubSystemImportTarget server interface.
type SubSystemImportTargetServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetImportTarget method retrieves the name of the import target to associate with
	// the LUNs being imported on the subsystem.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetImportTarget(context.Context, *GetImportTargetRequest) (*GetImportTargetResponse, error)

	// The SetImportTarget method sets the name of the import target to associate with the
	// LUNs being imported on the subsystem.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetImportTarget(context.Context, *SetImportTargetRequest) (*SetImportTargetResponse, error)
}

func RegisterSubSystemImportTargetServer(conn dcerpc.Conn, o SubSystemImportTargetServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSubSystemImportTargetServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SubSystemImportTargetSyntaxV0_0))...)
}

func NewSubSystemImportTargetServerHandle(o SubSystemImportTargetServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SubSystemImportTargetServerHandle(ctx, o, opNum, r)
	}
}

func SubSystemImportTargetServerHandle(ctx context.Context, o SubSystemImportTargetServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetImportTarget
		op := &xxx_GetImportTargetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImportTargetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImportTarget(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SetImportTarget
		op := &xxx_SetImportTargetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetImportTargetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetImportTarget(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsSubSystemImportTarget
type UnimplementedSubSystemImportTargetServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedSubSystemImportTargetServer) GetImportTarget(context.Context, *GetImportTargetRequest) (*GetImportTargetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSubSystemImportTargetServer) SetImportTarget(context.Context, *SetImportTargetRequest) (*SetImportTargetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SubSystemImportTargetServer = (*UnimplementedSubSystemImportTargetServer)(nil)

package ivdsdisk

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// IVdsDisk server interface.
type DiskServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// GetPack operation.
	GetPack(context.Context, *GetPackRequest) (*GetPackResponse, error)

	// The GetIdentificationData method retrieves information that uniquely identifies a
	// disk.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetIdentificationData(context.Context, *GetIdentificationDataRequest) (*GetIdentificationDataResponse, error)

	// QueryExtents operation.
	QueryExtents(context.Context, *QueryExtentsRequest) (*QueryExtentsResponse, error)

	// The ConvertStyle method converts a disk's partitioning format.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	ConvertStyle(context.Context, *ConvertStyleRequest) (*ConvertStyleResponse, error)

	// SetFlags operation.
	SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error)

	// ClearFlags operation.
	ClearFlags(context.Context, *ClearFlagsRequest) (*ClearFlagsResponse, error)
}

func RegisterDiskServer(conn dcerpc.Conn, o DiskServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDiskServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DiskSyntaxV0_0))...)
}

func NewDiskServerHandle(o DiskServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DiskServerHandle(ctx, o, opNum, r)
	}
}

func DiskServerHandle(ctx context.Context, o DiskServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProperties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetPack
		op := &xxx_GetPackOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPackRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPack(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetIdentificationData
		op := &xxx_GetIdentificationDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIdentificationDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIdentificationData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // QueryExtents
		op := &xxx_QueryExtentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryExtentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryExtents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // ConvertStyle
		op := &xxx_ConvertStyleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConvertStyleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConvertStyle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SetFlags
		op := &xxx_SetFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ClearFlags
		op := &xxx_ClearFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClearFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsDisk
type UnimplementedDiskServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDiskServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskServer) GetPack(context.Context, *GetPackRequest) (*GetPackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskServer) GetIdentificationData(context.Context, *GetIdentificationDataRequest) (*GetIdentificationDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskServer) QueryExtents(context.Context, *QueryExtentsRequest) (*QueryExtentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskServer) ConvertStyle(context.Context, *ConvertStyleRequest) (*ConvertStyleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskServer) SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskServer) ClearFlags(context.Context, *ClearFlagsRequest) (*ClearFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DiskServer = (*UnimplementedDiskServer)(nil)

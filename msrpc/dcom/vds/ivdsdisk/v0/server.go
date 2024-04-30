package ivdsdisk

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
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetPack
		in := &GetPackRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPack(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetIdentificationData
		in := &GetIdentificationDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIdentificationData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // QueryExtents
		in := &QueryExtentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryExtents(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // ConvertStyle
		in := &ConvertStyleRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ConvertStyle(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SetFlags
		in := &SetFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // ClearFlags
		in := &ClearFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}

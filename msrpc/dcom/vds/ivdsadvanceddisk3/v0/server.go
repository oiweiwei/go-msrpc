package ivdsadvanceddisk3

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

// IVdsAdvancedDisk3 server interface.
type AdvancedDisk3Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The GetUniqueId method<104> retrieves the device path that the operating system uses
	// to identify the disk.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetUniqueID(context.Context, *GetUniqueIDRequest) (*GetUniqueIDResponse, error)
}

func RegisterAdvancedDisk3Server(conn dcerpc.Conn, o AdvancedDisk3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAdvancedDisk3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AdvancedDisk3SyntaxV0_0))...)
}

func NewAdvancedDisk3ServerHandle(o AdvancedDisk3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AdvancedDisk3ServerHandle(ctx, o, opNum, r)
	}
}

func AdvancedDisk3ServerHandle(ctx context.Context, o AdvancedDisk3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // GetUniqueId
		op := &xxx_GetUniqueIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUniqueIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUniqueID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsAdvancedDisk3
type UnimplementedAdvancedDisk3Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAdvancedDisk3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAdvancedDisk3Server) GetUniqueID(context.Context, *GetUniqueIDRequest) (*GetUniqueIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AdvancedDisk3Server = (*UnimplementedAdvancedDisk3Server)(nil)

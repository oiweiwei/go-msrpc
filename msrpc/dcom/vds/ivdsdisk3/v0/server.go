package ivdsdisk3

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

// IVdsDisk3 server interface.
type Disk3Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties2 operation.
	GetProperties2(context.Context, *GetProperties2Request) (*GetProperties2Response, error)

	// The QueryFreeExtents method enumerates a disk's free extents.Returns all free extents
	// on the disk and aligns them to the alignment value supplied in the ulAlign parameter.
	// If there is no alignment value supplied, QueryFreeExtents aligns the free extents
	// based on the default alignment values.
	//
	// Return Values: QueryFreeExtents MUST return zero to indicate success, or an implementation-specific,
	// nonzero error code to indicate failure.
	//
	// Free extents are not returned for CD/DVD, or super floppy devices.
	//
	// If the disk has no partition format (it is not formatted as either MBR or GPT), then
	// this method MUST return VDS_E_DISK_NOT_INITIALIZED.
	QueryFreeExtents(context.Context, *QueryFreeExtentsRequest) (*QueryFreeExtentsResponse, error)
}

func RegisterDisk3Server(conn dcerpc.Conn, o Disk3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDisk3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Disk3SyntaxV0_0))...)
}

func NewDisk3ServerHandle(o Disk3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Disk3ServerHandle(ctx, o, opNum, r)
	}
}

func Disk3ServerHandle(ctx context.Context, o Disk3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProperties2
		op := &xxx_GetProperties2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProperties2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // QueryFreeExtents
		op := &xxx_QueryFreeExtentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFreeExtentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFreeExtents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsDisk3
type UnimplementedDisk3Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDisk3Server) GetProperties2(context.Context, *GetProperties2Request) (*GetProperties2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDisk3Server) QueryFreeExtents(context.Context, *QueryFreeExtentsRequest) (*QueryFreeExtentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Disk3Server = (*UnimplementedDisk3Server)(nil)

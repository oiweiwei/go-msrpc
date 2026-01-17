package ivdsdiskpartitionmf

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

// IVdsDiskPartitionMF server interface.
type DiskPartitionMFServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetPartitionFileSystemProperties method returns property details about the file
	// system on a disk partition at a specified byte offset. This method is only supported
	// on OEM, ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	GetPartitionFileSystemProperties(context.Context, *GetPartitionFileSystemPropertiesRequest) (*GetPartitionFileSystemPropertiesResponse, error)

	// The GetPartitionFileSystemTypeName method retrieves the name of the file system on
	// a disk partition at a specified byte offset. This method is only supported on OEM,
	// ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetPartitionFileSystemTypeName(context.Context, *GetPartitionFileSystemTypeNameRequest) (*GetPartitionFileSystemTypeNameResponse, error)

	// The QueryPartitionFileSystemFormatSupport method retrieves the properties of the
	// file systems that support formatting a disk partition at a specified byte offset.
	// This method is only supported on OEM, ESP, recovery, and unknown partitions.
	//
	// Note  This method is not valid on CD/DVD or super floppy devices. These devices do
	// not support partition tables.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.<107>
	//
	// ERROR_SUCCESS (0x00000000)
	QueryPartitionFileSystemFormatSupport(context.Context, *QueryPartitionFileSystemFormatSupportRequest) (*QueryPartitionFileSystemFormatSupportResponse, error)

	// The FormatPartitionEx method formats an existing OEM, ESP, or unknown partition.
	// This method is only supported on OEM, ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	FormatPartitionEx(context.Context, *FormatPartitionExRequest) (*FormatPartitionExResponse, error)
}

func RegisterDiskPartitionMFServer(conn dcerpc.Conn, o DiskPartitionMFServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDiskPartitionMFServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DiskPartitionMFSyntaxV0_0))...)
}

func NewDiskPartitionMFServerHandle(o DiskPartitionMFServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DiskPartitionMFServerHandle(ctx, o, opNum, r)
	}
}

func DiskPartitionMFServerHandle(ctx context.Context, o DiskPartitionMFServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetPartitionFileSystemProperties
		op := &xxx_GetPartitionFileSystemPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPartitionFileSystemPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPartitionFileSystemProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetPartitionFileSystemTypeName
		op := &xxx_GetPartitionFileSystemTypeNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPartitionFileSystemTypeNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPartitionFileSystemTypeName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // QueryPartitionFileSystemFormatSupport
		op := &xxx_QueryPartitionFileSystemFormatSupportOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryPartitionFileSystemFormatSupportRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryPartitionFileSystemFormatSupport(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // FormatPartitionEx
		op := &xxx_FormatPartitionExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FormatPartitionExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FormatPartitionEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsDiskPartitionMF
type UnimplementedDiskPartitionMFServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDiskPartitionMFServer) GetPartitionFileSystemProperties(context.Context, *GetPartitionFileSystemPropertiesRequest) (*GetPartitionFileSystemPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskPartitionMFServer) GetPartitionFileSystemTypeName(context.Context, *GetPartitionFileSystemTypeNameRequest) (*GetPartitionFileSystemTypeNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskPartitionMFServer) QueryPartitionFileSystemFormatSupport(context.Context, *QueryPartitionFileSystemFormatSupportRequest) (*QueryPartitionFileSystemFormatSupportResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDiskPartitionMFServer) FormatPartitionEx(context.Context, *FormatPartitionExRequest) (*FormatPartitionExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DiskPartitionMFServer = (*UnimplementedDiskPartitionMFServer)(nil)

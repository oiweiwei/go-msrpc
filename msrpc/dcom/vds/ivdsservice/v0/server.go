package ivdsservice

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

// IVdsService server interface.
type ServiceServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IsServiceReady method determines whether a service is finished initializing.
	// Until the service initialization completes, an application SHOULD NOT call any method
	// other than GetProperties.<71>
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	IsServiceReady(context.Context, *IsServiceReadyRequest) (*IsServiceReadyResponse, error)

	// The WaitForServiceReady method waits for VDS initialization to complete and returns
	// the status of the VDS initialization in the HRESULT.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero to indicate success or the error code
	// VDS_E_INITIALIZED_FAILED if the service-ready state is "failed".
	WaitForServiceReady(context.Context, *WaitForServiceReadyRequest) (*WaitForServiceReadyResponse, error)

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The QueryProviders method enumerates the providers of the server.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryProviders(context.Context, *QueryProvidersRequest) (*QueryProvidersResponse, error)

	// Opnum07NotUsedOnWire operation.
	// Opnum07NotUsedOnWire

	// The QueryUnallocatedDisks method enumerates the unallocated disks on the server.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryUnallocatedDisks(context.Context, *QueryUnallocatedDisksRequest) (*QueryUnallocatedDisksResponse, error)

	// The GetObject method retrieves an IUnknown pointer to a specified object.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)

	// The QueryDriveLetters method enumerates the drive letters of the server.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryDriveLetters(context.Context, *QueryDriveLettersRequest) (*QueryDriveLettersResponse, error)

	// The QueryFileSystemTypes method returns property details for all file systems that
	// are known to VDS.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryFileSystemTypes(context.Context, *QueryFileSystemTypesRequest) (*QueryFileSystemTypesResponse, error)

	// The Reenumerate method discovers newly added and newly removed disks and returns
	// the status of the operation in the HRESULT.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Reenumerate(context.Context, *ReenumerateRequest) (*ReenumerateResponse, error)

	// The Refresh method refreshes the ownership and layout of disks on the server.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)

	// The CleanupObsoleteMountPoints method removes any mount points that point to volumes
	// that no longer exist.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	CleanupObsoleteMountPoints(context.Context, *CleanupObsoleteMountPointsRequest) (*CleanupObsoleteMountPointsResponse, error)

	// The Advise method registers a notification callback with the server. Clients pass
	// the callback object to the server to receive notifications.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Advise(context.Context, *AdviseRequest) (*AdviseResponse, error)

	// The Unadvise method unregisters a client from being notified by the server of changes
	// to storage objects.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Unadvise(context.Context, *UnadviseRequest) (*UnadviseResponse, error)

	// The Reboot method restarts the computer on which the server is running.<73>
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Reboot(context.Context, *RebootRequest) (*RebootResponse, error)

	// SetFlags operation.
	SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error)

	// ClearFlags operation.
	ClearFlags(context.Context, *ClearFlagsRequest) (*ClearFlagsResponse, error)
}

func RegisterServiceServer(conn dcerpc.Conn, o ServiceServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceSyntaxV0_0))...)
}

func NewServiceServerHandle(o ServiceServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceServerHandle(ctx, o, opNum, r)
	}
}

func ServiceServerHandle(ctx context.Context, o ServiceServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // IsServiceReady
		op := &xxx_IsServiceReadyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsServiceReadyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsServiceReady(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // WaitForServiceReady
		op := &xxx_WaitForServiceReadyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WaitForServiceReadyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WaitForServiceReady(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetProperties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // QueryProviders
		op := &xxx_QueryProvidersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryProvidersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryProviders(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Opnum07NotUsedOnWire
		// Opnum07NotUsedOnWire
		return nil, nil
	case 8: // QueryUnallocatedDisks
		op := &xxx_QueryUnallocatedDisksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryUnallocatedDisksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryUnallocatedDisks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // GetObject
		op := &xxx_GetObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // QueryDriveLetters
		op := &xxx_QueryDriveLettersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDriveLettersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDriveLetters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // QueryFileSystemTypes
		op := &xxx_QueryFileSystemTypesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFileSystemTypesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFileSystemTypes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Reenumerate
		op := &xxx_ReenumerateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReenumerateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reenumerate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Refresh
		op := &xxx_RefreshOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Refresh(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // CleanupObsoleteMountPoints
		op := &xxx_CleanupObsoleteMountPointsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CleanupObsoleteMountPointsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CleanupObsoleteMountPoints(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Advise
		op := &xxx_AdviseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AdviseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Advise(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Unadvise
		op := &xxx_UnadviseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnadviseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Unadvise(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Reboot
		op := &xxx_RebootOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RebootRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reboot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SetFlags
		op := &xxx_SetFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // ClearFlags
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

// Unimplemented IVdsService
type UnimplementedServiceServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServiceServer) IsServiceReady(context.Context, *IsServiceReadyRequest) (*IsServiceReadyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) WaitForServiceReady(context.Context, *WaitForServiceReadyRequest) (*WaitForServiceReadyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) QueryProviders(context.Context, *QueryProvidersRequest) (*QueryProvidersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) QueryUnallocatedDisks(context.Context, *QueryUnallocatedDisksRequest) (*QueryUnallocatedDisksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) QueryDriveLetters(context.Context, *QueryDriveLettersRequest) (*QueryDriveLettersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) QueryFileSystemTypes(context.Context, *QueryFileSystemTypesRequest) (*QueryFileSystemTypesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) Reenumerate(context.Context, *ReenumerateRequest) (*ReenumerateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) CleanupObsoleteMountPoints(context.Context, *CleanupObsoleteMountPointsRequest) (*CleanupObsoleteMountPointsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) Advise(context.Context, *AdviseRequest) (*AdviseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) Unadvise(context.Context, *UnadviseRequest) (*UnadviseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) Reboot(context.Context, *RebootRequest) (*RebootResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceServer) ClearFlags(context.Context, *ClearFlagsRequest) (*ClearFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServiceServer = (*UnimplementedServiceServer)(nil)

package ivdsservice

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
		in := &IsServiceReadyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsServiceReady(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // WaitForServiceReady
		in := &WaitForServiceReadyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WaitForServiceReady(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetProperties
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // QueryProviders
		in := &QueryProvidersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryProviders(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // Opnum07NotUsedOnWire
		// Opnum07NotUsedOnWire
		return nil, nil
	case 8: // QueryUnallocatedDisks
		in := &QueryUnallocatedDisksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryUnallocatedDisks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetObject
		in := &GetObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // QueryDriveLetters
		in := &QueryDriveLettersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryDriveLetters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // QueryFileSystemTypes
		in := &QueryFileSystemTypesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryFileSystemTypes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Reenumerate
		in := &ReenumerateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Reenumerate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Refresh
		in := &RefreshRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Refresh(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // CleanupObsoleteMountPoints
		in := &CleanupObsoleteMountPointsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CleanupObsoleteMountPoints(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Advise
		in := &AdviseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Advise(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Unadvise
		in := &UnadviseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Unadvise(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Reboot
		in := &RebootRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Reboot(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // SetFlags
		in := &SetFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // ClearFlags
		in := &ClearFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}

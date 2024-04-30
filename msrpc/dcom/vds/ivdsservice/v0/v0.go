package ivdsservice

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsService interface identifier 0818a8ef-9ba9-40d8-a6f9-e22833cc771e
	ServiceIID = &dcom.IID{Data1: 0x0818a8ef, Data2: 0x9ba9, Data3: 0x40d8, Data4: []byte{0xa6, 0xf9, 0xe2, 0x28, 0x33, 0xcc, 0x77, 0x1e}}
	// Syntax UUID
	ServiceSyntaxUUID = &uuid.UUID{TimeLow: 0x818a8ef, TimeMid: 0x9ba9, TimeHiAndVersion: 0x40d8, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xf9, Node: [6]uint8{0xe2, 0x28, 0x33, 0xcc, 0x77, 0x1e}}
	// Syntax ID
	ServiceSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServiceSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsService interface.
type ServiceClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

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
	IsServiceReady(context.Context, *IsServiceReadyRequest, ...dcerpc.CallOption) (*IsServiceReadyResponse, error)

	// The WaitForServiceReady method waits for VDS initialization to complete and returns
	// the status of the VDS initialization in the HRESULT.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero to indicate success or the error code
	// VDS_E_INITIALIZED_FAILED if the service-ready state is "failed".
	WaitForServiceReady(context.Context, *WaitForServiceReadyRequest, ...dcerpc.CallOption) (*WaitForServiceReadyResponse, error)

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The QueryProviders method enumerates the providers of the server.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryProviders(context.Context, *QueryProvidersRequest, ...dcerpc.CallOption) (*QueryProvidersResponse, error)

	// Opnum07NotUsedOnWire operation.
	// Opnum07NotUsedOnWire

	// The QueryUnallocatedDisks method enumerates the unallocated disks on the server.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryUnallocatedDisks(context.Context, *QueryUnallocatedDisksRequest, ...dcerpc.CallOption) (*QueryUnallocatedDisksResponse, error)

	// The GetObject method retrieves an IUnknown pointer to a specified object.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetObject(context.Context, *GetObjectRequest, ...dcerpc.CallOption) (*GetObjectResponse, error)

	// The QueryDriveLetters method enumerates the drive letters of the server.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryDriveLetters(context.Context, *QueryDriveLettersRequest, ...dcerpc.CallOption) (*QueryDriveLettersResponse, error)

	// The QueryFileSystemTypes method returns property details for all file systems that
	// are known to VDS.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryFileSystemTypes(context.Context, *QueryFileSystemTypesRequest, ...dcerpc.CallOption) (*QueryFileSystemTypesResponse, error)

	// The Reenumerate method discovers newly added and newly removed disks and returns
	// the status of the operation in the HRESULT.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Reenumerate(context.Context, *ReenumerateRequest, ...dcerpc.CallOption) (*ReenumerateResponse, error)

	// The Refresh method refreshes the ownership and layout of disks on the server.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Refresh(context.Context, *RefreshRequest, ...dcerpc.CallOption) (*RefreshResponse, error)

	// The CleanupObsoleteMountPoints method removes any mount points that point to volumes
	// that no longer exist.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	CleanupObsoleteMountPoints(context.Context, *CleanupObsoleteMountPointsRequest, ...dcerpc.CallOption) (*CleanupObsoleteMountPointsResponse, error)

	// The Advise method registers a notification callback with the server. Clients pass
	// the callback object to the server to receive notifications.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Advise(context.Context, *AdviseRequest, ...dcerpc.CallOption) (*AdviseResponse, error)

	// The Unadvise method unregisters a client from being notified by the server of changes
	// to storage objects.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Unadvise(context.Context, *UnadviseRequest, ...dcerpc.CallOption) (*UnadviseResponse, error)

	// The Reboot method restarts the computer on which the server is running.<73>
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Reboot(context.Context, *RebootRequest, ...dcerpc.CallOption) (*RebootResponse, error)

	// SetFlags operation.
	SetFlags(context.Context, *SetFlagsRequest, ...dcerpc.CallOption) (*SetFlagsResponse, error)

	// ClearFlags operation.
	ClearFlags(context.Context, *ClearFlagsRequest, ...dcerpc.CallOption) (*ClearFlagsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServiceClient
}

type xxx_DefaultServiceClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServiceClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServiceClient) IsServiceReady(ctx context.Context, in *IsServiceReadyRequest, opts ...dcerpc.CallOption) (*IsServiceReadyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsServiceReadyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) WaitForServiceReady(ctx context.Context, in *WaitForServiceReadyRequest, opts ...dcerpc.CallOption) (*WaitForServiceReadyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WaitForServiceReadyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) QueryProviders(ctx context.Context, in *QueryProvidersRequest, opts ...dcerpc.CallOption) (*QueryProvidersResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryProvidersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) QueryUnallocatedDisks(ctx context.Context, in *QueryUnallocatedDisksRequest, opts ...dcerpc.CallOption) (*QueryUnallocatedDisksResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryUnallocatedDisksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...dcerpc.CallOption) (*GetObjectResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) QueryDriveLetters(ctx context.Context, in *QueryDriveLettersRequest, opts ...dcerpc.CallOption) (*QueryDriveLettersResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryDriveLettersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) QueryFileSystemTypes(ctx context.Context, in *QueryFileSystemTypesRequest, opts ...dcerpc.CallOption) (*QueryFileSystemTypesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryFileSystemTypesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) Reenumerate(ctx context.Context, in *ReenumerateRequest, opts ...dcerpc.CallOption) (*ReenumerateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReenumerateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...dcerpc.CallOption) (*RefreshResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RefreshResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) CleanupObsoleteMountPoints(ctx context.Context, in *CleanupObsoleteMountPointsRequest, opts ...dcerpc.CallOption) (*CleanupObsoleteMountPointsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CleanupObsoleteMountPointsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) Advise(ctx context.Context, in *AdviseRequest, opts ...dcerpc.CallOption) (*AdviseResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AdviseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) Unadvise(ctx context.Context, in *UnadviseRequest, opts ...dcerpc.CallOption) (*UnadviseResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnadviseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) Reboot(ctx context.Context, in *RebootRequest, opts ...dcerpc.CallOption) (*RebootResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RebootResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) SetFlags(ctx context.Context, in *SetFlagsRequest, opts ...dcerpc.CallOption) (*SetFlagsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) ClearFlags(ctx context.Context, in *ClearFlagsRequest, opts ...dcerpc.CallOption) (*ClearFlagsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClearFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServiceClient) IPID(ctx context.Context, ipid *dcom.IPID) ServiceClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServiceClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewServiceClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServiceClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServiceSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultServiceClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_IsServiceReadyOperation structure represents the IsServiceReady operation
type xxx_IsServiceReadyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsServiceReadyOperation) OpNum() int { return 3 }

func (o *xxx_IsServiceReadyOperation) OpName() string { return "/IVdsService/v0/IsServiceReady" }

func (o *xxx_IsServiceReadyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServiceReadyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServiceReadyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServiceReadyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServiceReadyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsServiceReadyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsServiceReadyRequest structure represents the IsServiceReady operation request
type IsServiceReadyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsServiceReadyRequest) xxx_ToOp(ctx context.Context) *xxx_IsServiceReadyOperation {
	if o == nil {
		return &xxx_IsServiceReadyOperation{}
	}
	return &xxx_IsServiceReadyOperation{
		This: o.This,
	}
}

func (o *IsServiceReadyRequest) xxx_FromOp(ctx context.Context, op *xxx_IsServiceReadyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsServiceReadyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsServiceReadyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsServiceReadyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsServiceReadyResponse structure represents the IsServiceReady operation response
type IsServiceReadyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IsServiceReady return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsServiceReadyResponse) xxx_ToOp(ctx context.Context) *xxx_IsServiceReadyOperation {
	if o == nil {
		return &xxx_IsServiceReadyOperation{}
	}
	return &xxx_IsServiceReadyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *IsServiceReadyResponse) xxx_FromOp(ctx context.Context, op *xxx_IsServiceReadyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsServiceReadyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsServiceReadyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsServiceReadyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitForServiceReadyOperation structure represents the WaitForServiceReady operation
type xxx_WaitForServiceReadyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WaitForServiceReadyOperation) OpNum() int { return 4 }

func (o *xxx_WaitForServiceReadyOperation) OpName() string {
	return "/IVdsService/v0/WaitForServiceReady"
}

func (o *xxx_WaitForServiceReadyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForServiceReadyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForServiceReadyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForServiceReadyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForServiceReadyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForServiceReadyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WaitForServiceReadyRequest structure represents the WaitForServiceReady operation request
type WaitForServiceReadyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *WaitForServiceReadyRequest) xxx_ToOp(ctx context.Context) *xxx_WaitForServiceReadyOperation {
	if o == nil {
		return &xxx_WaitForServiceReadyOperation{}
	}
	return &xxx_WaitForServiceReadyOperation{
		This: o.This,
	}
}

func (o *WaitForServiceReadyRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitForServiceReadyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *WaitForServiceReadyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WaitForServiceReadyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForServiceReadyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WaitForServiceReadyResponse structure represents the WaitForServiceReady operation response
type WaitForServiceReadyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The WaitForServiceReady return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitForServiceReadyResponse) xxx_ToOp(ctx context.Context) *xxx_WaitForServiceReadyOperation {
	if o == nil {
		return &xxx_WaitForServiceReadyOperation{}
	}
	return &xxx_WaitForServiceReadyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *WaitForServiceReadyResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitForServiceReadyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *WaitForServiceReadyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WaitForServiceReadyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForServiceReadyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This            *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ServiceProperty *vds.ServiceProperty `idl:"name:pServiceProp" json:"service_property"`
	Return          int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 5 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IVdsService/v0/GetProperties" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pServiceProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_SERVICE_PROP}(struct))
	{
		if o.ServiceProperty != nil {
			if err := o.ServiceProperty.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ServiceProperty{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pServiceProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_SERVICE_PROP}(struct))
	{
		if o.ServiceProperty == nil {
			o.ServiceProperty = &vds.ServiceProperty{}
		}
		if err := o.ServiceProperty.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPropertiesRequest structure represents the GetProperties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesOperation {
	if o == nil {
		return &xxx_GetPropertiesOperation{}
	}
	return &xxx_GetPropertiesOperation{
		This: o.This,
	}
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the GetProperties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ServiceProperty *vds.ServiceProperty `idl:"name:pServiceProp" json:"service_property"`
	// Return: The GetProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesOperation {
	if o == nil {
		return &xxx_GetPropertiesOperation{}
	}
	return &xxx_GetPropertiesOperation{
		That:            o.That,
		ServiceProperty: o.ServiceProperty,
		Return:          o.Return,
	}
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServiceProperty = op.ServiceProperty
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryProvidersOperation structure represents the QueryProviders operation
type xxx_QueryProvidersOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Masks  uint32          `idl:"name:masks" json:"masks"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryProvidersOperation) OpNum() int { return 6 }

func (o *xxx_QueryProvidersOperation) OpName() string { return "/IVdsService/v0/QueryProviders" }

func (o *xxx_QueryProvidersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProvidersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// masks {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Masks); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProvidersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// masks {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Masks); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProvidersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProvidersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.EnumObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProvidersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &vds.EnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**vds.EnumObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryProvidersRequest structure represents the QueryProviders operation request
type QueryProvidersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// masks: The combination of any values, by using a bitwise OR operator, that the VDS_QUERY_PROVIDER_FLAG
	// enumeration defines. The values that are set in the mask specify the types of providers
	// to return.
	Masks uint32 `idl:"name:masks" json:"masks"`
}

func (o *QueryProvidersRequest) xxx_ToOp(ctx context.Context) *xxx_QueryProvidersOperation {
	if o == nil {
		return &xxx_QueryProvidersOperation{}
	}
	return &xxx_QueryProvidersOperation{
		This:  o.This,
		Masks: o.Masks,
	}
}

func (o *QueryProvidersRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryProvidersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Masks = op.Masks
}
func (o *QueryProvidersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryProvidersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProvidersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryProvidersResponse structure represents the QueryProviders operation response
type QueryProvidersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if successfully completed,
	// receives the IEnumVdsObject interface of the object that contains an enumeration
	// of provider objects on the server. Callers MUST release the interface that is received
	// when they are done with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryProviders return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryProvidersResponse) xxx_ToOp(ctx context.Context) *xxx_QueryProvidersOperation {
	if o == nil {
		return &xxx_QueryProvidersOperation{}
	}
	return &xxx_QueryProvidersOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *QueryProvidersResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryProvidersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryProvidersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryProvidersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProvidersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryUnallocatedDisksOperation structure represents the QueryUnallocatedDisks operation
type xxx_QueryUnallocatedDisksOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryUnallocatedDisksOperation) OpNum() int { return 8 }

func (o *xxx_QueryUnallocatedDisksOperation) OpName() string {
	return "/IVdsService/v0/QueryUnallocatedDisks"
}

func (o *xxx_QueryUnallocatedDisksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUnallocatedDisksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUnallocatedDisksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUnallocatedDisksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUnallocatedDisksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.EnumObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUnallocatedDisksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &vds.EnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**vds.EnumObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryUnallocatedDisksRequest structure represents the QueryUnallocatedDisks operation request
type QueryUnallocatedDisksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryUnallocatedDisksRequest) xxx_ToOp(ctx context.Context) *xxx_QueryUnallocatedDisksOperation {
	if o == nil {
		return &xxx_QueryUnallocatedDisksOperation{}
	}
	return &xxx_QueryUnallocatedDisksOperation{
		This: o.This,
	}
}

func (o *QueryUnallocatedDisksRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryUnallocatedDisksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryUnallocatedDisksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryUnallocatedDisksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryUnallocatedDisksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryUnallocatedDisksResponse structure represents the QueryUnallocatedDisks operation response
type QueryUnallocatedDisksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if the operation is successfully
	// completed, receives the IEnumVdsObject interface of the object that contains an enumeration
	// of disk objects that correspond to unallocated disks on the server. Callers MUST
	// release the interface that is received when they are done with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryUnallocatedDisks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryUnallocatedDisksResponse) xxx_ToOp(ctx context.Context) *xxx_QueryUnallocatedDisksOperation {
	if o == nil {
		return &xxx_QueryUnallocatedDisksOperation{}
	}
	return &xxx_QueryUnallocatedDisksOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *QueryUnallocatedDisksResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryUnallocatedDisksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryUnallocatedDisksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryUnallocatedDisksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryUnallocatedDisksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectOperation structure represents the GetObject operation
type xxx_GetObjectOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID      *vds.ObjectID  `idl:"name:ObjectId" json:"object_id"`
	Type          vds.ObjectType `idl:"name:type" json:"type"`
	ObjectUnknown *dcom.Unknown  `idl:"name:ppObjectUnk" json:"object_unknown"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectOperation) OpNum() int { return 9 }

func (o *xxx_GetObjectOperation) OpName() string { return "/IVdsService/v0/GetObject" }

func (o *xxx_GetObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ObjectId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// type {in} (1:{alias=VDS_OBJECT_TYPE}(enum))
	{
		if err := w.WriteData(uint16(o.Type)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ObjectId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &vds.ObjectID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// type {in} (1:{alias=VDS_OBJECT_TYPE}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Type)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppObjectUnk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.ObjectUnknown != nil {
			_ptr_ppObjectUnk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectUnknown != nil {
					if err := o.ObjectUnknown.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectUnknown, _ptr_ppObjectUnk); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppObjectUnk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppObjectUnk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectUnknown == nil {
				o.ObjectUnknown = &dcom.Unknown{}
			}
			if err := o.ObjectUnknown.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppObjectUnk := func(ptr interface{}) { o.ObjectUnknown = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.ObjectUnknown, _s_ppObjectUnk, _ptr_ppObjectUnk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetObjectRequest structure represents the GetObject operation request
type GetObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ObjectId: The GUID of the desired object.
	ObjectID *vds.ObjectID `idl:"name:ObjectId" json:"object_id"`
	// type: The object type that VDS_OBJECT_TYPE enumerates. All object types are valid
	// except VDS_OT_UNKNOWN, VDS_OT_PROVIDER, VDS_OT_ASYNC, and VDS_OT_ENUM.
	Type vds.ObjectType `idl:"name:type" json:"type"`
}

func (o *GetObjectRequest) xxx_ToOp(ctx context.Context) *xxx_GetObjectOperation {
	if o == nil {
		return &xxx_GetObjectOperation{}
	}
	return &xxx_GetObjectOperation{
		This:     o.This,
		ObjectID: o.ObjectID,
		Type:     o.Type,
	}
}

func (o *GetObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
}
func (o *GetObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectResponse structure represents the GetObject operation response
type GetObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppObjectUnk: A pointer to an IUnknown interface that, if the operation is successfully
	// completed, receives an IUnknown interface of the object. Callers MUST release the
	// interface that is received when they are done with it.
	ObjectUnknown *dcom.Unknown `idl:"name:ppObjectUnk" json:"object_unknown"`
	// Return: The GetObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectResponse) xxx_ToOp(ctx context.Context) *xxx_GetObjectOperation {
	if o == nil {
		return &xxx_GetObjectOperation{}
	}
	return &xxx_GetObjectOperation{
		That:          o.That,
		ObjectUnknown: o.ObjectUnknown,
		Return:        o.Return,
	}
}

func (o *GetObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ObjectUnknown = op.ObjectUnknown
	o.Return = op.Return
}
func (o *GetObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDriveLettersOperation structure represents the QueryDriveLetters operation
type xxx_QueryDriveLettersOperation struct {
	This                     *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat             `idl:"name:That" json:"that"`
	FirstLetter              uint16                     `idl:"name:wcFirstLetter" json:"first_letter"`
	Count                    uint32                     `idl:"name:count" json:"count"`
	DriveLetterPropertyArray []*vds.DriveLetterProperty `idl:"name:pDriveLetterPropArray;size_is:(count)" json:"drive_letter_property_array"`
	Return                   int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDriveLettersOperation) OpNum() int { return 10 }

func (o *xxx_QueryDriveLettersOperation) OpName() string { return "/IVdsService/v0/QueryDriveLetters" }

func (o *xxx_QueryDriveLettersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDriveLettersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// wcFirstLetter {in} (1:{alias=WCHAR}(wchar))
	{
		if err := w.WriteData(o.FirstLetter); err != nil {
			return err
		}
	}
	// count {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDriveLettersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wcFirstLetter {in} (1:{alias=WCHAR}(wchar))
	{
		if err := w.ReadData(&o.FirstLetter); err != nil {
			return err
		}
	}
	// count {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDriveLettersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDriveLettersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pDriveLetterPropArray {out} (1:{pointer=ref}*(1))(2:{alias=VDS_DRIVE_LETTER_PROP}[dim:0,size_is=count](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DriveLetterPropertyArray {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DriveLetterPropertyArray[i1] != nil {
				if err := o.DriveLetterPropertyArray[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&vds.DriveLetterProperty{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DriveLetterPropertyArray); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&vds.DriveLetterProperty{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDriveLettersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pDriveLetterPropArray {out} (1:{pointer=ref}*(1))(2:{alias=VDS_DRIVE_LETTER_PROP}[dim:0,size_is=count](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DriveLetterPropertyArray", sizeInfo[0])
		}
		o.DriveLetterPropertyArray = make([]*vds.DriveLetterProperty, sizeInfo[0])
		for i1 := range o.DriveLetterPropertyArray {
			i1 := i1
			if o.DriveLetterPropertyArray[i1] == nil {
				o.DriveLetterPropertyArray[i1] = &vds.DriveLetterProperty{}
			}
			if err := o.DriveLetterPropertyArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryDriveLettersRequest structure represents the QueryDriveLetters operation request
type QueryDriveLettersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// wcFirstLetter: The first drive letter to query as a single uppercase or lowercase
	// alphabetical (A-Z) Unicode character.
	FirstLetter uint16 `idl:"name:wcFirstLetter" json:"first_letter"`
	// count: The total number of drive letters to retrieve, beginning with the letter that
	// wcFirstLetter specifies. This MUST also be the number of elements in the pDriveLetterPropArray.
	// It MUST NOT exceed the total number of drive letters between the letter in wcFirstLetter
	// and the last possible drive letter (Z), inclusive.
	Count uint32 `idl:"name:count" json:"count"`
}

func (o *QueryDriveLettersRequest) xxx_ToOp(ctx context.Context) *xxx_QueryDriveLettersOperation {
	if o == nil {
		return &xxx_QueryDriveLettersOperation{}
	}
	return &xxx_QueryDriveLettersOperation{
		This:        o.This,
		FirstLetter: o.FirstLetter,
		Count:       o.Count,
	}
}

func (o *QueryDriveLettersRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryDriveLettersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FirstLetter = op.FirstLetter
	o.Count = op.Count
}
func (o *QueryDriveLettersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryDriveLettersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDriveLettersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDriveLettersResponse structure represents the QueryDriveLetters operation response
type QueryDriveLettersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pDriveLetterPropArray: An array of VDS_DRIVE_LETTER_PROP structures that, if the
	// operation is successfully completed, receives the array of drive letter properties.
	DriveLetterPropertyArray []*vds.DriveLetterProperty `idl:"name:pDriveLetterPropArray;size_is:(count)" json:"drive_letter_property_array"`
	// Return: The QueryDriveLetters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDriveLettersResponse) xxx_ToOp(ctx context.Context) *xxx_QueryDriveLettersOperation {
	if o == nil {
		return &xxx_QueryDriveLettersOperation{}
	}
	return &xxx_QueryDriveLettersOperation{
		That:                     o.That,
		DriveLetterPropertyArray: o.DriveLetterPropertyArray,
		Return:                   o.Return,
	}
}

func (o *QueryDriveLettersResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryDriveLettersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DriveLetterPropertyArray = op.DriveLetterPropertyArray
	o.Return = op.Return
}
func (o *QueryDriveLettersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryDriveLettersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDriveLettersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryFileSystemTypesOperation structure represents the QueryFileSystemTypes operation
type xxx_QueryFileSystemTypesOperation struct {
	This                     *dcom.ORPCThis                `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat                `idl:"name:That" json:"that"`
	FileSystemTypeProperties []*vds.FileSystemTypeProperty `idl:"name:ppFileSystemTypeProps;size_is:(, plNumberOfFileSystems)" json:"file_system_type_properties"`
	NumberOfFileSystems      int32                         `idl:"name:plNumberOfFileSystems" json:"number_of_file_systems"`
	Return                   int32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryFileSystemTypesOperation) OpNum() int { return 11 }

func (o *xxx_QueryFileSystemTypesOperation) OpName() string {
	return "/IVdsService/v0/QueryFileSystemTypes"
}

func (o *xxx_QueryFileSystemTypesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileSystemTypesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileSystemTypesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileSystemTypesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.FileSystemTypeProperties != nil && o.NumberOfFileSystems == 0 {
		o.NumberOfFileSystems = int32(len(o.FileSystemTypeProperties))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileSystemTypesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppFileSystemTypeProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_FILE_SYSTEM_TYPE_PROP}[dim:0,size_is=plNumberOfFileSystems](struct))
	{
		if o.FileSystemTypeProperties != nil || o.NumberOfFileSystems > 0 {
			_ptr_ppFileSystemTypeProps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfFileSystems)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.FileSystemTypeProperties {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.FileSystemTypeProperties[i1] != nil {
						if err := o.FileSystemTypeProperties[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.FileSystemTypeProperty{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.FileSystemTypeProperties); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.FileSystemTypeProperty{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemTypeProperties, _ptr_ppFileSystemTypeProps); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfFileSystems {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfFileSystems); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileSystemTypesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppFileSystemTypeProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_FILE_SYSTEM_TYPE_PROP}[dim:0,size_is=plNumberOfFileSystems](struct))
	{
		_ptr_ppFileSystemTypeProps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.FileSystemTypeProperties", sizeInfo[0])
			}
			o.FileSystemTypeProperties = make([]*vds.FileSystemTypeProperty, sizeInfo[0])
			for i1 := range o.FileSystemTypeProperties {
				i1 := i1
				if o.FileSystemTypeProperties[i1] == nil {
					o.FileSystemTypeProperties[i1] = &vds.FileSystemTypeProperty{}
				}
				if err := o.FileSystemTypeProperties[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppFileSystemTypeProps := func(ptr interface{}) { o.FileSystemTypeProperties = *ptr.(*[]*vds.FileSystemTypeProperty) }
		if err := w.ReadPointer(&o.FileSystemTypeProperties, _s_ppFileSystemTypeProps, _ptr_ppFileSystemTypeProps); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfFileSystems {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfFileSystems); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryFileSystemTypesRequest structure represents the QueryFileSystemTypes operation request
type QueryFileSystemTypesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryFileSystemTypesRequest) xxx_ToOp(ctx context.Context) *xxx_QueryFileSystemTypesOperation {
	if o == nil {
		return &xxx_QueryFileSystemTypesOperation{}
	}
	return &xxx_QueryFileSystemTypesOperation{
		This: o.This,
	}
}

func (o *QueryFileSystemTypesRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryFileSystemTypesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryFileSystemTypesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryFileSystemTypesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFileSystemTypesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryFileSystemTypesResponse structure represents the QueryFileSystemTypes operation response
type QueryFileSystemTypesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppFileSystemTypeProps: A pointer to an array of VDS_FILE_SYSTEM_TYPE_PROP structures
	// that, if the operation is successfully completed, receives the array of file system
	// type properties.
	FileSystemTypeProperties []*vds.FileSystemTypeProperty `idl:"name:ppFileSystemTypeProps;size_is:(, plNumberOfFileSystems)" json:"file_system_type_properties"`
	// plNumberOfFileSystems: A pointer to a variable that, if the operation is successfully
	// completed, receives the total number of elements returned in ppFileSystemTypeProps.
	NumberOfFileSystems int32 `idl:"name:plNumberOfFileSystems" json:"number_of_file_systems"`
	// Return: The QueryFileSystemTypes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryFileSystemTypesResponse) xxx_ToOp(ctx context.Context) *xxx_QueryFileSystemTypesOperation {
	if o == nil {
		return &xxx_QueryFileSystemTypesOperation{}
	}
	return &xxx_QueryFileSystemTypesOperation{
		That:                     o.That,
		FileSystemTypeProperties: o.FileSystemTypeProperties,
		NumberOfFileSystems:      o.NumberOfFileSystems,
		Return:                   o.Return,
	}
}

func (o *QueryFileSystemTypesResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryFileSystemTypesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemTypeProperties = op.FileSystemTypeProperties
	o.NumberOfFileSystems = op.NumberOfFileSystems
	o.Return = op.Return
}
func (o *QueryFileSystemTypesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryFileSystemTypesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFileSystemTypesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReenumerateOperation structure represents the Reenumerate operation
type xxx_ReenumerateOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReenumerateOperation) OpNum() int { return 12 }

func (o *xxx_ReenumerateOperation) OpName() string { return "/IVdsService/v0/Reenumerate" }

func (o *xxx_ReenumerateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReenumerateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReenumerateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReenumerateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReenumerateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReenumerateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReenumerateRequest structure represents the Reenumerate operation request
type ReenumerateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ReenumerateRequest) xxx_ToOp(ctx context.Context) *xxx_ReenumerateOperation {
	if o == nil {
		return &xxx_ReenumerateOperation{}
	}
	return &xxx_ReenumerateOperation{
		This: o.This,
	}
}

func (o *ReenumerateRequest) xxx_FromOp(ctx context.Context, op *xxx_ReenumerateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ReenumerateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReenumerateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReenumerateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReenumerateResponse structure represents the Reenumerate operation response
type ReenumerateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reenumerate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReenumerateResponse) xxx_ToOp(ctx context.Context) *xxx_ReenumerateOperation {
	if o == nil {
		return &xxx_ReenumerateOperation{}
	}
	return &xxx_ReenumerateOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ReenumerateResponse) xxx_FromOp(ctx context.Context, op *xxx_ReenumerateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ReenumerateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReenumerateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReenumerateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RefreshOperation structure represents the Refresh operation
type xxx_RefreshOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshOperation) OpNum() int { return 13 }

func (o *xxx_RefreshOperation) OpName() string { return "/IVdsService/v0/Refresh" }

func (o *xxx_RefreshOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RefreshRequest structure represents the Refresh operation request
type RefreshRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RefreshRequest) xxx_ToOp(ctx context.Context) *xxx_RefreshOperation {
	if o == nil {
		return &xxx_RefreshOperation{}
	}
	return &xxx_RefreshOperation{
		This: o.This,
	}
}

func (o *RefreshRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RefreshRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RefreshRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshResponse structure represents the Refresh operation response
type RefreshResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Refresh return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshResponse) xxx_ToOp(ctx context.Context) *xxx_RefreshOperation {
	if o == nil {
		return &xxx_RefreshOperation{}
	}
	return &xxx_RefreshOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RefreshResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RefreshResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CleanupObsoleteMountPointsOperation structure represents the CleanupObsoleteMountPoints operation
type xxx_CleanupObsoleteMountPointsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CleanupObsoleteMountPointsOperation) OpNum() int { return 14 }

func (o *xxx_CleanupObsoleteMountPointsOperation) OpName() string {
	return "/IVdsService/v0/CleanupObsoleteMountPoints"
}

func (o *xxx_CleanupObsoleteMountPointsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupObsoleteMountPointsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupObsoleteMountPointsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupObsoleteMountPointsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupObsoleteMountPointsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupObsoleteMountPointsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CleanupObsoleteMountPointsRequest structure represents the CleanupObsoleteMountPoints operation request
type CleanupObsoleteMountPointsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CleanupObsoleteMountPointsRequest) xxx_ToOp(ctx context.Context) *xxx_CleanupObsoleteMountPointsOperation {
	if o == nil {
		return &xxx_CleanupObsoleteMountPointsOperation{}
	}
	return &xxx_CleanupObsoleteMountPointsOperation{
		This: o.This,
	}
}

func (o *CleanupObsoleteMountPointsRequest) xxx_FromOp(ctx context.Context, op *xxx_CleanupObsoleteMountPointsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CleanupObsoleteMountPointsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CleanupObsoleteMountPointsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupObsoleteMountPointsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CleanupObsoleteMountPointsResponse structure represents the CleanupObsoleteMountPoints operation response
type CleanupObsoleteMountPointsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CleanupObsoleteMountPoints return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CleanupObsoleteMountPointsResponse) xxx_ToOp(ctx context.Context) *xxx_CleanupObsoleteMountPointsOperation {
	if o == nil {
		return &xxx_CleanupObsoleteMountPointsOperation{}
	}
	return &xxx_CleanupObsoleteMountPointsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CleanupObsoleteMountPointsResponse) xxx_FromOp(ctx context.Context, op *xxx_CleanupObsoleteMountPointsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CleanupObsoleteMountPointsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CleanupObsoleteMountPointsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupObsoleteMountPointsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AdviseOperation structure represents the Advise operation
type xxx_AdviseOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Sink   *vds.AdviseSink `idl:"name:pSink" json:"sink"`
	Cookie uint32          `idl:"name:pdwCookie" json:"cookie"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_AdviseOperation) OpNum() int { return 15 }

func (o *xxx_AdviseOperation) OpName() string { return "/IVdsService/v0/Advise" }

func (o *xxx_AdviseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pSink {in} (1:{pointer=ref}*(1))(2:{alias=IVdsAdviseSink}(interface))
	{
		if o.Sink != nil {
			_ptr_pSink := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Sink != nil {
					if err := o.Sink.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.AdviseSink{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Sink, _ptr_pSink); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pSink {in} (1:{pointer=ref}*(1))(2:{alias=IVdsAdviseSink}(interface))
	{
		_ptr_pSink := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Sink == nil {
				o.Sink = &vds.AdviseSink{}
			}
			if err := o.Sink.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSink := func(ptr interface{}) { o.Sink = *ptr.(**vds.AdviseSink) }
		if err := w.ReadPointer(&o.Sink, _s_pSink, _ptr_pSink); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwCookie {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cookie); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwCookie {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cookie); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AdviseRequest structure represents the Advise operation request
type AdviseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pSink: A pointer to an IVdsAdviseSink interface of the callback object to register
	// with the server for notification of object changes.
	Sink *vds.AdviseSink `idl:"name:pSink" json:"sink"`
}

func (o *AdviseRequest) xxx_ToOp(ctx context.Context) *xxx_AdviseOperation {
	if o == nil {
		return &xxx_AdviseOperation{}
	}
	return &xxx_AdviseOperation{
		This: o.This,
		Sink: o.Sink,
	}
}

func (o *AdviseRequest) xxx_FromOp(ctx context.Context, op *xxx_AdviseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Sink = op.Sink
}
func (o *AdviseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AdviseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AdviseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AdviseResponse structure represents the Advise operation response
type AdviseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwCookie: A pointer to a variable that, if the operation is successfully completed,
	// receives a unique cookie value that the client can later use to unregister the callback
	// object from receiving notification changes from the service. For information about
	// how to register callback objects, see section 3.3.1.1.
	Cookie uint32 `idl:"name:pdwCookie" json:"cookie"`
	// Return: The Advise return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AdviseResponse) xxx_ToOp(ctx context.Context) *xxx_AdviseOperation {
	if o == nil {
		return &xxx_AdviseOperation{}
	}
	return &xxx_AdviseOperation{
		That:   o.That,
		Cookie: o.Cookie,
		Return: o.Return,
	}
}

func (o *AdviseResponse) xxx_FromOp(ctx context.Context, op *xxx_AdviseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Cookie = op.Cookie
	o.Return = op.Return
}
func (o *AdviseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AdviseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AdviseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnadviseOperation structure represents the Unadvise operation
type xxx_UnadviseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Cookie uint32         `idl:"name:dwCookie" json:"cookie"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_UnadviseOperation) OpNum() int { return 16 }

func (o *xxx_UnadviseOperation) OpName() string { return "/IVdsService/v0/Unadvise" }

func (o *xxx_UnadviseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwCookie {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cookie); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwCookie {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cookie); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UnadviseRequest structure represents the Unadvise operation request
type UnadviseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwCookie: The cookie value generated when the IVdsAdviseSink interface was registered.
	Cookie uint32 `idl:"name:dwCookie" json:"cookie"`
}

func (o *UnadviseRequest) xxx_ToOp(ctx context.Context) *xxx_UnadviseOperation {
	if o == nil {
		return &xxx_UnadviseOperation{}
	}
	return &xxx_UnadviseOperation{
		This:   o.This,
		Cookie: o.Cookie,
	}
}

func (o *UnadviseRequest) xxx_FromOp(ctx context.Context, op *xxx_UnadviseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Cookie = op.Cookie
}
func (o *UnadviseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UnadviseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnadviseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnadviseResponse structure represents the Unadvise operation response
type UnadviseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Unadvise return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UnadviseResponse) xxx_ToOp(ctx context.Context) *xxx_UnadviseOperation {
	if o == nil {
		return &xxx_UnadviseOperation{}
	}
	return &xxx_UnadviseOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *UnadviseResponse) xxx_FromOp(ctx context.Context, op *xxx_UnadviseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *UnadviseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UnadviseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnadviseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RebootOperation structure represents the Reboot operation
type xxx_RebootOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RebootOperation) OpNum() int { return 17 }

func (o *xxx_RebootOperation) OpName() string { return "/IVdsService/v0/Reboot" }

func (o *xxx_RebootOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RebootRequest structure represents the Reboot operation request
type RebootRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RebootRequest) xxx_ToOp(ctx context.Context) *xxx_RebootOperation {
	if o == nil {
		return &xxx_RebootOperation{}
	}
	return &xxx_RebootOperation{
		This: o.This,
	}
}

func (o *RebootRequest) xxx_FromOp(ctx context.Context, op *xxx_RebootOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RebootRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RebootRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RebootOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RebootResponse structure represents the Reboot operation response
type RebootResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reboot return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RebootResponse) xxx_ToOp(ctx context.Context) *xxx_RebootOperation {
	if o == nil {
		return &xxx_RebootOperation{}
	}
	return &xxx_RebootOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RebootResponse) xxx_FromOp(ctx context.Context, op *xxx_RebootOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RebootResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RebootResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RebootOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFlagsOperation structure represents the SetFlags operation
type xxx_SetFlagsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags  uint32         `idl:"name:ulFlags" json:"flags"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFlagsOperation) OpNum() int { return 18 }

func (o *xxx_SetFlagsOperation) OpName() string { return "/IVdsService/v0/SetFlags" }

func (o *xxx_SetFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetFlagsRequest structure represents the SetFlags operation request
type SetFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flags uint32         `idl:"name:ulFlags" json:"flags"`
}

func (o *SetFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_SetFlagsOperation {
	if o == nil {
		return &xxx_SetFlagsOperation{}
	}
	return &xxx_SetFlagsOperation{
		This:  o.This,
		Flags: o.Flags,
	}
}

func (o *SetFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *SetFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFlagsResponse structure represents the SetFlags operation response
type SetFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_SetFlagsOperation {
	if o == nil {
		return &xxx_SetFlagsOperation{}
	}
	return &xxx_SetFlagsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearFlagsOperation structure represents the ClearFlags operation
type xxx_ClearFlagsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags  uint32         `idl:"name:ulFlags" json:"flags"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearFlagsOperation) OpNum() int { return 19 }

func (o *xxx_ClearFlagsOperation) OpName() string { return "/IVdsService/v0/ClearFlags" }

func (o *xxx_ClearFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClearFlagsRequest structure represents the ClearFlags operation request
type ClearFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flags uint32         `idl:"name:ulFlags" json:"flags"`
}

func (o *ClearFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_ClearFlagsOperation {
	if o == nil {
		return &xxx_ClearFlagsOperation{}
	}
	return &xxx_ClearFlagsOperation{
		This:  o.This,
		Flags: o.Flags,
	}
}

func (o *ClearFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *ClearFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClearFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearFlagsResponse structure represents the ClearFlags operation response
type ClearFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClearFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_ClearFlagsOperation {
	if o == nil {
		return &xxx_ClearFlagsOperation{}
	}
	return &xxx_ClearFlagsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ClearFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClearFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

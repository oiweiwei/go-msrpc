package intmslibrarycontrol1

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
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsLibraryControl1 interface identifier 4e934f30-341a-11d1-8fb1-00a024cb6019
	LibraryControl1IID = &dcom.IID{Data1: 0x4e934f30, Data2: 0x341a, Data3: 0x11d1, Data4: []byte{0x8f, 0xb1, 0x00, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax UUID
	LibraryControl1SyntaxUUID = &uuid.UUID{TimeLow: 0x4e934f30, TimeMid: 0x341a, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xb1, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax ID
	LibraryControl1SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: LibraryControl1SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsLibraryControl1 interface.
type LibraryControl1Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	EjectNTMSMedia(context.Context, *EjectNTMSMediaRequest, ...dcerpc.CallOption) (*EjectNTMSMediaResponse, error)

	InjectNTMSMedia(context.Context, *InjectNTMSMediaRequest, ...dcerpc.CallOption) (*InjectNTMSMediaResponse, error)

	AccessNTMSLibraryDoor(context.Context, *AccessNTMSLibraryDoorRequest, ...dcerpc.CallOption) (*AccessNTMSLibraryDoorResponse, error)

	CleanNTMSDrive(context.Context, *CleanNTMSDriveRequest, ...dcerpc.CallOption) (*CleanNTMSDriveResponse, error)

	DismountNTMSDrive(context.Context, *DismountNTMSDriveRequest, ...dcerpc.CallOption) (*DismountNTMSDriveResponse, error)

	InventoryNTMSLibrary(context.Context, *InventoryNTMSLibraryRequest, ...dcerpc.CallOption) (*InventoryNTMSLibraryResponse, error)

	// INtmsLibraryControl1_LocalOnlyOpnum09 operation.
	LibraryControl1LocalOnlyOpnum09(context.Context, *LibraryControl1LocalOnlyOpnum09Request, ...dcerpc.CallOption) (*LibraryControl1LocalOnlyOpnum09Response, error)

	CancelNTMSLibraryRequest(context.Context, *CancelNTMSLibraryRequestRequest, ...dcerpc.CallOption) (*CancelNTMSLibraryRequestResponse, error)

	ReserveNTMSCleanerSlot(context.Context, *ReserveNTMSCleanerSlotRequest, ...dcerpc.CallOption) (*ReserveNTMSCleanerSlotResponse, error)

	ReleaseNTMSCleanerSlot(context.Context, *ReleaseNTMSCleanerSlotRequest, ...dcerpc.CallOption) (*ReleaseNTMSCleanerSlotResponse, error)

	InjectNTMSCleaner(context.Context, *InjectNTMSCleanerRequest, ...dcerpc.CallOption) (*InjectNTMSCleanerResponse, error)

	EjectNTMSCleaner(context.Context, *EjectNTMSCleanerRequest, ...dcerpc.CallOption) (*EjectNTMSCleanerResponse, error)

	DeleteNTMSLibrary(context.Context, *DeleteNTMSLibraryRequest, ...dcerpc.CallOption) (*DeleteNTMSLibraryResponse, error)

	DeleteNTMSDrive(context.Context, *DeleteNTMSDriveRequest, ...dcerpc.CallOption) (*DeleteNTMSDriveResponse, error)

	GetNTMSRequestOrder(context.Context, *GetNTMSRequestOrderRequest, ...dcerpc.CallOption) (*GetNTMSRequestOrderResponse, error)

	SetNTMSRequestOrder(context.Context, *SetNTMSRequestOrderRequest, ...dcerpc.CallOption) (*SetNTMSRequestOrderResponse, error)

	DeleteNTMSRequests(context.Context, *DeleteNTMSRequestsRequest, ...dcerpc.CallOption) (*DeleteNTMSRequestsResponse, error)

	BeginNTMSDeviceChangeDetection(context.Context, *BeginNTMSDeviceChangeDetectionRequest, ...dcerpc.CallOption) (*BeginNTMSDeviceChangeDetectionResponse, error)

	SetNTMSDeviceChangeDetection(context.Context, *SetNTMSDeviceChangeDetectionRequest, ...dcerpc.CallOption) (*SetNTMSDeviceChangeDetectionResponse, error)

	EndNTMSDeviceChangeDetection(context.Context, *EndNTMSDeviceChangeDetectionRequest, ...dcerpc.CallOption) (*EndNTMSDeviceChangeDetectionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) LibraryControl1Client
}

type xxx_DefaultLibraryControl1Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultLibraryControl1Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultLibraryControl1Client) EjectNTMSMedia(ctx context.Context, in *EjectNTMSMediaRequest, opts ...dcerpc.CallOption) (*EjectNTMSMediaResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EjectNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) InjectNTMSMedia(ctx context.Context, in *InjectNTMSMediaRequest, opts ...dcerpc.CallOption) (*InjectNTMSMediaResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &InjectNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) AccessNTMSLibraryDoor(ctx context.Context, in *AccessNTMSLibraryDoorRequest, opts ...dcerpc.CallOption) (*AccessNTMSLibraryDoorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &AccessNTMSLibraryDoorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) CleanNTMSDrive(ctx context.Context, in *CleanNTMSDriveRequest, opts ...dcerpc.CallOption) (*CleanNTMSDriveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CleanNTMSDriveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) DismountNTMSDrive(ctx context.Context, in *DismountNTMSDriveRequest, opts ...dcerpc.CallOption) (*DismountNTMSDriveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DismountNTMSDriveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) InventoryNTMSLibrary(ctx context.Context, in *InventoryNTMSLibraryRequest, opts ...dcerpc.CallOption) (*InventoryNTMSLibraryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &InventoryNTMSLibraryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) LibraryControl1LocalOnlyOpnum09(ctx context.Context, in *LibraryControl1LocalOnlyOpnum09Request, opts ...dcerpc.CallOption) (*LibraryControl1LocalOnlyOpnum09Response, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &LibraryControl1LocalOnlyOpnum09Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) CancelNTMSLibraryRequest(ctx context.Context, in *CancelNTMSLibraryRequestRequest, opts ...dcerpc.CallOption) (*CancelNTMSLibraryRequestResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CancelNTMSLibraryRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) ReserveNTMSCleanerSlot(ctx context.Context, in *ReserveNTMSCleanerSlotRequest, opts ...dcerpc.CallOption) (*ReserveNTMSCleanerSlotResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &ReserveNTMSCleanerSlotResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) ReleaseNTMSCleanerSlot(ctx context.Context, in *ReleaseNTMSCleanerSlotRequest, opts ...dcerpc.CallOption) (*ReleaseNTMSCleanerSlotResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &ReleaseNTMSCleanerSlotResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) InjectNTMSCleaner(ctx context.Context, in *InjectNTMSCleanerRequest, opts ...dcerpc.CallOption) (*InjectNTMSCleanerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &InjectNTMSCleanerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) EjectNTMSCleaner(ctx context.Context, in *EjectNTMSCleanerRequest, opts ...dcerpc.CallOption) (*EjectNTMSCleanerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EjectNTMSCleanerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) DeleteNTMSLibrary(ctx context.Context, in *DeleteNTMSLibraryRequest, opts ...dcerpc.CallOption) (*DeleteNTMSLibraryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteNTMSLibraryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) DeleteNTMSDrive(ctx context.Context, in *DeleteNTMSDriveRequest, opts ...dcerpc.CallOption) (*DeleteNTMSDriveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteNTMSDriveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) GetNTMSRequestOrder(ctx context.Context, in *GetNTMSRequestOrderRequest, opts ...dcerpc.CallOption) (*GetNTMSRequestOrderResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetNTMSRequestOrderResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) SetNTMSRequestOrder(ctx context.Context, in *SetNTMSRequestOrderRequest, opts ...dcerpc.CallOption) (*SetNTMSRequestOrderResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetNTMSRequestOrderResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) DeleteNTMSRequests(ctx context.Context, in *DeleteNTMSRequestsRequest, opts ...dcerpc.CallOption) (*DeleteNTMSRequestsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteNTMSRequestsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) BeginNTMSDeviceChangeDetection(ctx context.Context, in *BeginNTMSDeviceChangeDetectionRequest, opts ...dcerpc.CallOption) (*BeginNTMSDeviceChangeDetectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &BeginNTMSDeviceChangeDetectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) SetNTMSDeviceChangeDetection(ctx context.Context, in *SetNTMSDeviceChangeDetectionRequest, opts ...dcerpc.CallOption) (*SetNTMSDeviceChangeDetectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetNTMSDeviceChangeDetectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) EndNTMSDeviceChangeDetection(ctx context.Context, in *EndNTMSDeviceChangeDetectionRequest, opts ...dcerpc.CallOption) (*EndNTMSDeviceChangeDetectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EndNTMSDeviceChangeDetectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLibraryControl1Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLibraryControl1Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultLibraryControl1Client) IPID(ctx context.Context, ipid *dcom.IPID) LibraryControl1Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultLibraryControl1Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewLibraryControl1Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LibraryControl1Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LibraryControl1SyntaxV0_0))...)
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
	return &xxx_DefaultLibraryControl1Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_EjectNTMSMediaOperation structure represents the EjectNtmsMedia operation
type xxx_EjectNTMSMediaOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID        *dtyp.GUID     `idl:"name:lpMediaId;pointer:unique" json:"media_id"`
	EjectOperation *dtyp.GUID     `idl:"name:lpEjectOperation" json:"eject_operation"`
	Action         uint32         `idl:"name:dwAction" json:"action"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EjectNTMSMediaOperation) OpNum() int { return 0 }

func (o *xxx_EjectNTMSMediaOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/EjectNtmsMedia"
}

func (o *xxx_EjectNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			_ptr_lpMediaId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MediaID != nil {
					if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MediaID, _ptr_lpMediaId); err != nil {
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
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation != nil {
			if err := o.EjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpMediaId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MediaID == nil {
				o.MediaID = &dtyp.GUID{}
			}
			if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpMediaId := func(ptr interface{}) { o.MediaID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.MediaID, _s_lpMediaId, _ptr_lpMediaId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation == nil {
			o.EjectOperation = &dtyp.GUID{}
		}
		if err := o.EjectOperation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation != nil {
			if err := o.EjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EjectNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation == nil {
			o.EjectOperation = &dtyp.GUID{}
		}
		if err := o.EjectOperation.UnmarshalNDR(ctx, w); err != nil {
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

// EjectNTMSMediaRequest structure represents the EjectNtmsMedia operation request
type EjectNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	MediaID        *dtyp.GUID     `idl:"name:lpMediaId;pointer:unique" json:"media_id"`
	EjectOperation *dtyp.GUID     `idl:"name:lpEjectOperation" json:"eject_operation"`
	Action         uint32         `idl:"name:dwAction" json:"action"`
}

func (o *EjectNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_EjectNTMSMediaOperation) *xxx_EjectNTMSMediaOperation {
	if op == nil {
		op = &xxx_EjectNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	op.EjectOperation = o.EjectOperation
	op.Action = o.Action
	return op
}

func (o *EjectNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_EjectNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
	o.EjectOperation = op.EjectOperation
	o.Action = op.Action
}
func (o *EjectNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EjectNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EjectNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EjectNTMSMediaResponse structure represents the EjectNtmsMedia operation response
type EjectNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	EjectOperation *dtyp.GUID     `idl:"name:lpEjectOperation" json:"eject_operation"`
	// Return: The EjectNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EjectNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_EjectNTMSMediaOperation) *xxx_EjectNTMSMediaOperation {
	if op == nil {
		op = &xxx_EjectNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.EjectOperation = o.EjectOperation
	op.Return = o.Return
	return op
}

func (o *EjectNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_EjectNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EjectOperation = op.EjectOperation
	o.Return = op.Return
}
func (o *EjectNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EjectNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EjectNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InjectNTMSMediaOperation structure represents the InjectNtmsMedia operation
type xxx_InjectNTMSMediaOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	LibraryID       *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
	InjectOperation *dtyp.GUID     `idl:"name:lpInjectOperation" json:"inject_operation"`
	Action          uint32         `idl:"name:dwAction" json:"action"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InjectNTMSMediaOperation) OpNum() int { return 1 }

func (o *xxx_InjectNTMSMediaOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/InjectNtmsMedia"
}

func (o *xxx_InjectNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID != nil {
			if err := o.LibraryID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation != nil {
			if err := o.InjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID == nil {
			o.LibraryID = &dtyp.GUID{}
		}
		if err := o.LibraryID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation == nil {
			o.InjectOperation = &dtyp.GUID{}
		}
		if err := o.InjectOperation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation != nil {
			if err := o.InjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_InjectNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation == nil {
			o.InjectOperation = &dtyp.GUID{}
		}
		if err := o.InjectOperation.UnmarshalNDR(ctx, w); err != nil {
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

// InjectNTMSMediaRequest structure represents the InjectNtmsMedia operation request
type InjectNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	LibraryID       *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
	InjectOperation *dtyp.GUID     `idl:"name:lpInjectOperation" json:"inject_operation"`
	Action          uint32         `idl:"name:dwAction" json:"action"`
}

func (o *InjectNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_InjectNTMSMediaOperation) *xxx_InjectNTMSMediaOperation {
	if op == nil {
		op = &xxx_InjectNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LibraryID = o.LibraryID
	op.InjectOperation = o.InjectOperation
	op.Action = o.Action
	return op
}

func (o *InjectNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_InjectNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LibraryID = op.LibraryID
	o.InjectOperation = op.InjectOperation
	o.Action = op.Action
}
func (o *InjectNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InjectNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InjectNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InjectNTMSMediaResponse structure represents the InjectNtmsMedia operation response
type InjectNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	InjectOperation *dtyp.GUID     `idl:"name:lpInjectOperation" json:"inject_operation"`
	// Return: The InjectNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InjectNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_InjectNTMSMediaOperation) *xxx_InjectNTMSMediaOperation {
	if op == nil {
		op = &xxx_InjectNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.InjectOperation = o.InjectOperation
	op.Return = o.Return
	return op
}

func (o *InjectNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_InjectNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InjectOperation = op.InjectOperation
	o.Return = op.Return
}
func (o *InjectNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InjectNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InjectNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AccessNTMSLibraryDoorOperation structure represents the AccessNtmsLibraryDoor operation
type xxx_AccessNTMSLibraryDoorOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	LibraryID *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
	Action    uint32         `idl:"name:dwAction" json:"action"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AccessNTMSLibraryDoorOperation) OpNum() int { return 2 }

func (o *xxx_AccessNTMSLibraryDoorOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/AccessNtmsLibraryDoor"
}

func (o *xxx_AccessNTMSLibraryDoorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessNTMSLibraryDoorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID != nil {
			if err := o.LibraryID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessNTMSLibraryDoorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID == nil {
			o.LibraryID = &dtyp.GUID{}
		}
		if err := o.LibraryID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessNTMSLibraryDoorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessNTMSLibraryDoorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AccessNTMSLibraryDoorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AccessNTMSLibraryDoorRequest structure represents the AccessNtmsLibraryDoor operation request
type AccessNTMSLibraryDoorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	LibraryID *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
	Action    uint32         `idl:"name:dwAction" json:"action"`
}

func (o *AccessNTMSLibraryDoorRequest) xxx_ToOp(ctx context.Context, op *xxx_AccessNTMSLibraryDoorOperation) *xxx_AccessNTMSLibraryDoorOperation {
	if op == nil {
		op = &xxx_AccessNTMSLibraryDoorOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LibraryID = o.LibraryID
	op.Action = o.Action
	return op
}

func (o *AccessNTMSLibraryDoorRequest) xxx_FromOp(ctx context.Context, op *xxx_AccessNTMSLibraryDoorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LibraryID = op.LibraryID
	o.Action = op.Action
}
func (o *AccessNTMSLibraryDoorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AccessNTMSLibraryDoorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AccessNTMSLibraryDoorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AccessNTMSLibraryDoorResponse structure represents the AccessNtmsLibraryDoor operation response
type AccessNTMSLibraryDoorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AccessNtmsLibraryDoor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AccessNTMSLibraryDoorResponse) xxx_ToOp(ctx context.Context, op *xxx_AccessNTMSLibraryDoorOperation) *xxx_AccessNTMSLibraryDoorOperation {
	if op == nil {
		op = &xxx_AccessNTMSLibraryDoorOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AccessNTMSLibraryDoorResponse) xxx_FromOp(ctx context.Context, op *xxx_AccessNTMSLibraryDoorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AccessNTMSLibraryDoorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AccessNTMSLibraryDoorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AccessNTMSLibraryDoorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CleanNTMSDriveOperation structure represents the CleanNtmsDrive operation
type xxx_CleanNTMSDriveOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	DriveID *dtyp.GUID     `idl:"name:lpDriveId" json:"drive_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CleanNTMSDriveOperation) OpNum() int { return 3 }

func (o *xxx_CleanNTMSDriveOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/CleanNtmsDrive"
}

func (o *xxx_CleanNTMSDriveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanNTMSDriveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpDriveId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.DriveID != nil {
			if err := o.DriveID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CleanNTMSDriveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpDriveId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.DriveID == nil {
			o.DriveID = &dtyp.GUID{}
		}
		if err := o.DriveID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanNTMSDriveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanNTMSDriveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CleanNTMSDriveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CleanNTMSDriveRequest structure represents the CleanNtmsDrive operation request
type CleanNTMSDriveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	DriveID *dtyp.GUID     `idl:"name:lpDriveId" json:"drive_id"`
}

func (o *CleanNTMSDriveRequest) xxx_ToOp(ctx context.Context, op *xxx_CleanNTMSDriveOperation) *xxx_CleanNTMSDriveOperation {
	if op == nil {
		op = &xxx_CleanNTMSDriveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DriveID = o.DriveID
	return op
}

func (o *CleanNTMSDriveRequest) xxx_FromOp(ctx context.Context, op *xxx_CleanNTMSDriveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DriveID = op.DriveID
}
func (o *CleanNTMSDriveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CleanNTMSDriveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanNTMSDriveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CleanNTMSDriveResponse structure represents the CleanNtmsDrive operation response
type CleanNTMSDriveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CleanNtmsDrive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CleanNTMSDriveResponse) xxx_ToOp(ctx context.Context, op *xxx_CleanNTMSDriveOperation) *xxx_CleanNTMSDriveOperation {
	if op == nil {
		op = &xxx_CleanNTMSDriveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CleanNTMSDriveResponse) xxx_FromOp(ctx context.Context, op *xxx_CleanNTMSDriveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CleanNTMSDriveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CleanNTMSDriveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanNTMSDriveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DismountNTMSDriveOperation structure represents the DismountNtmsDrive operation
type xxx_DismountNTMSDriveOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	DriveID *dtyp.GUID     `idl:"name:lpDriveId" json:"drive_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DismountNTMSDriveOperation) OpNum() int { return 4 }

func (o *xxx_DismountNTMSDriveOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/DismountNtmsDrive"
}

func (o *xxx_DismountNTMSDriveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountNTMSDriveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpDriveId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.DriveID != nil {
			if err := o.DriveID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DismountNTMSDriveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpDriveId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.DriveID == nil {
			o.DriveID = &dtyp.GUID{}
		}
		if err := o.DriveID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountNTMSDriveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountNTMSDriveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DismountNTMSDriveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DismountNTMSDriveRequest structure represents the DismountNtmsDrive operation request
type DismountNTMSDriveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	DriveID *dtyp.GUID     `idl:"name:lpDriveId" json:"drive_id"`
}

func (o *DismountNTMSDriveRequest) xxx_ToOp(ctx context.Context, op *xxx_DismountNTMSDriveOperation) *xxx_DismountNTMSDriveOperation {
	if op == nil {
		op = &xxx_DismountNTMSDriveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DriveID = o.DriveID
	return op
}

func (o *DismountNTMSDriveRequest) xxx_FromOp(ctx context.Context, op *xxx_DismountNTMSDriveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DriveID = op.DriveID
}
func (o *DismountNTMSDriveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DismountNTMSDriveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DismountNTMSDriveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DismountNTMSDriveResponse structure represents the DismountNtmsDrive operation response
type DismountNTMSDriveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DismountNtmsDrive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DismountNTMSDriveResponse) xxx_ToOp(ctx context.Context, op *xxx_DismountNTMSDriveOperation) *xxx_DismountNTMSDriveOperation {
	if op == nil {
		op = &xxx_DismountNTMSDriveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DismountNTMSDriveResponse) xxx_FromOp(ctx context.Context, op *xxx_DismountNTMSDriveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DismountNTMSDriveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DismountNTMSDriveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DismountNTMSDriveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InventoryNTMSLibraryOperation structure represents the InventoryNtmsLibrary operation
type xxx_InventoryNTMSLibraryOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	LibraryID *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
	Action    uint32         `idl:"name:dwAction" json:"action"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InventoryNTMSLibraryOperation) OpNum() int { return 5 }

func (o *xxx_InventoryNTMSLibraryOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/InventoryNtmsLibrary"
}

func (o *xxx_InventoryNTMSLibraryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InventoryNTMSLibraryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID != nil {
			if err := o.LibraryID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InventoryNTMSLibraryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID == nil {
			o.LibraryID = &dtyp.GUID{}
		}
		if err := o.LibraryID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InventoryNTMSLibraryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InventoryNTMSLibraryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_InventoryNTMSLibraryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// InventoryNTMSLibraryRequest structure represents the InventoryNtmsLibrary operation request
type InventoryNTMSLibraryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	LibraryID *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
	Action    uint32         `idl:"name:dwAction" json:"action"`
}

func (o *InventoryNTMSLibraryRequest) xxx_ToOp(ctx context.Context, op *xxx_InventoryNTMSLibraryOperation) *xxx_InventoryNTMSLibraryOperation {
	if op == nil {
		op = &xxx_InventoryNTMSLibraryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LibraryID = o.LibraryID
	op.Action = o.Action
	return op
}

func (o *InventoryNTMSLibraryRequest) xxx_FromOp(ctx context.Context, op *xxx_InventoryNTMSLibraryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LibraryID = op.LibraryID
	o.Action = op.Action
}
func (o *InventoryNTMSLibraryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InventoryNTMSLibraryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InventoryNTMSLibraryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InventoryNTMSLibraryResponse structure represents the InventoryNtmsLibrary operation response
type InventoryNTMSLibraryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The InventoryNtmsLibrary return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InventoryNTMSLibraryResponse) xxx_ToOp(ctx context.Context, op *xxx_InventoryNTMSLibraryOperation) *xxx_InventoryNTMSLibraryOperation {
	if op == nil {
		op = &xxx_InventoryNTMSLibraryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *InventoryNTMSLibraryResponse) xxx_FromOp(ctx context.Context, op *xxx_InventoryNTMSLibraryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *InventoryNTMSLibraryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InventoryNTMSLibraryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InventoryNTMSLibraryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LibraryControl1LocalOnlyOpnum09Operation structure represents the INtmsLibraryControl1_LocalOnlyOpnum09 operation
type xxx_LibraryControl1LocalOnlyOpnum09Operation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) OpNum() int { return 6 }

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) OpName() string {
	return "/INtmsLibraryControl1/v0/INtmsLibraryControl1_LocalOnlyOpnum09"
}

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_LibraryControl1LocalOnlyOpnum09Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// LibraryControl1LocalOnlyOpnum09Request structure represents the INtmsLibraryControl1_LocalOnlyOpnum09 operation request
type LibraryControl1LocalOnlyOpnum09Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *LibraryControl1LocalOnlyOpnum09Request) xxx_ToOp(ctx context.Context, op *xxx_LibraryControl1LocalOnlyOpnum09Operation) *xxx_LibraryControl1LocalOnlyOpnum09Operation {
	if op == nil {
		op = &xxx_LibraryControl1LocalOnlyOpnum09Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *LibraryControl1LocalOnlyOpnum09Request) xxx_FromOp(ctx context.Context, op *xxx_LibraryControl1LocalOnlyOpnum09Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *LibraryControl1LocalOnlyOpnum09Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LibraryControl1LocalOnlyOpnum09Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LibraryControl1LocalOnlyOpnum09Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LibraryControl1LocalOnlyOpnum09Response structure represents the INtmsLibraryControl1_LocalOnlyOpnum09 operation response
type LibraryControl1LocalOnlyOpnum09Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The INtmsLibraryControl1_LocalOnlyOpnum09 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LibraryControl1LocalOnlyOpnum09Response) xxx_ToOp(ctx context.Context, op *xxx_LibraryControl1LocalOnlyOpnum09Operation) *xxx_LibraryControl1LocalOnlyOpnum09Operation {
	if op == nil {
		op = &xxx_LibraryControl1LocalOnlyOpnum09Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *LibraryControl1LocalOnlyOpnum09Response) xxx_FromOp(ctx context.Context, op *xxx_LibraryControl1LocalOnlyOpnum09Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *LibraryControl1LocalOnlyOpnum09Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LibraryControl1LocalOnlyOpnum09Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LibraryControl1LocalOnlyOpnum09Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelNTMSLibraryRequestOperation structure represents the CancelNtmsLibraryRequest operation
type xxx_CancelNTMSLibraryRequestOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelNTMSLibraryRequestOperation) OpNum() int { return 7 }

func (o *xxx_CancelNTMSLibraryRequestOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/CancelNtmsLibraryRequest"
}

func (o *xxx_CancelNTMSLibraryRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelNTMSLibraryRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CancelNTMSLibraryRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelNTMSLibraryRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelNTMSLibraryRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelNTMSLibraryRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CancelNTMSLibraryRequestRequest structure represents the CancelNtmsLibraryRequest operation request
type CancelNTMSLibraryRequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
}

func (o *CancelNTMSLibraryRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelNTMSLibraryRequestOperation) *xxx_CancelNTMSLibraryRequestOperation {
	if op == nil {
		op = &xxx_CancelNTMSLibraryRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestID = o.RequestID
	return op
}

func (o *CancelNTMSLibraryRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelNTMSLibraryRequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestID = op.RequestID
}
func (o *CancelNTMSLibraryRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelNTMSLibraryRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelNTMSLibraryRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelNTMSLibraryRequestResponse structure represents the CancelNtmsLibraryRequest operation response
type CancelNTMSLibraryRequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CancelNtmsLibraryRequest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelNTMSLibraryRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelNTMSLibraryRequestOperation) *xxx_CancelNTMSLibraryRequestOperation {
	if op == nil {
		op = &xxx_CancelNTMSLibraryRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CancelNTMSLibraryRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelNTMSLibraryRequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelNTMSLibraryRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelNTMSLibraryRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelNTMSLibraryRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReserveNTMSCleanerSlotOperation structure represents the ReserveNtmsCleanerSlot operation
type xxx_ReserveNTMSCleanerSlotOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Library *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
	Slot    *dtyp.GUID     `idl:"name:lpSlot" json:"slot"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReserveNTMSCleanerSlotOperation) OpNum() int { return 8 }

func (o *xxx_ReserveNTMSCleanerSlotOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/ReserveNtmsCleanerSlot"
}

func (o *xxx_ReserveNTMSCleanerSlotOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReserveNTMSCleanerSlotOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library != nil {
			if err := o.Library.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSlot {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Slot != nil {
			if err := o.Slot.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ReserveNTMSCleanerSlotOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library == nil {
			o.Library = &dtyp.GUID{}
		}
		if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSlot {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.Slot == nil {
			o.Slot = &dtyp.GUID{}
		}
		if err := o.Slot.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReserveNTMSCleanerSlotOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReserveNTMSCleanerSlotOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReserveNTMSCleanerSlotOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ReserveNTMSCleanerSlotRequest structure represents the ReserveNtmsCleanerSlot operation request
type ReserveNTMSCleanerSlotRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Library *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
	Slot    *dtyp.GUID     `idl:"name:lpSlot" json:"slot"`
}

func (o *ReserveNTMSCleanerSlotRequest) xxx_ToOp(ctx context.Context, op *xxx_ReserveNTMSCleanerSlotOperation) *xxx_ReserveNTMSCleanerSlotOperation {
	if op == nil {
		op = &xxx_ReserveNTMSCleanerSlotOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Library = o.Library
	op.Slot = o.Slot
	return op
}

func (o *ReserveNTMSCleanerSlotRequest) xxx_FromOp(ctx context.Context, op *xxx_ReserveNTMSCleanerSlotOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Library = op.Library
	o.Slot = op.Slot
}
func (o *ReserveNTMSCleanerSlotRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReserveNTMSCleanerSlotRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReserveNTMSCleanerSlotOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReserveNTMSCleanerSlotResponse structure represents the ReserveNtmsCleanerSlot operation response
type ReserveNTMSCleanerSlotResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReserveNtmsCleanerSlot return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReserveNTMSCleanerSlotResponse) xxx_ToOp(ctx context.Context, op *xxx_ReserveNTMSCleanerSlotOperation) *xxx_ReserveNTMSCleanerSlotOperation {
	if op == nil {
		op = &xxx_ReserveNTMSCleanerSlotOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ReserveNTMSCleanerSlotResponse) xxx_FromOp(ctx context.Context, op *xxx_ReserveNTMSCleanerSlotOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ReserveNTMSCleanerSlotResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReserveNTMSCleanerSlotResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReserveNTMSCleanerSlotOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReleaseNTMSCleanerSlotOperation structure represents the ReleaseNtmsCleanerSlot operation
type xxx_ReleaseNTMSCleanerSlotOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Library *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReleaseNTMSCleanerSlotOperation) OpNum() int { return 9 }

func (o *xxx_ReleaseNTMSCleanerSlotOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/ReleaseNtmsCleanerSlot"
}

func (o *xxx_ReleaseNTMSCleanerSlotOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReleaseNTMSCleanerSlotOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library != nil {
			if err := o.Library.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ReleaseNTMSCleanerSlotOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library == nil {
			o.Library = &dtyp.GUID{}
		}
		if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReleaseNTMSCleanerSlotOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReleaseNTMSCleanerSlotOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReleaseNTMSCleanerSlotOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ReleaseNTMSCleanerSlotRequest structure represents the ReleaseNtmsCleanerSlot operation request
type ReleaseNTMSCleanerSlotRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Library *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
}

func (o *ReleaseNTMSCleanerSlotRequest) xxx_ToOp(ctx context.Context, op *xxx_ReleaseNTMSCleanerSlotOperation) *xxx_ReleaseNTMSCleanerSlotOperation {
	if op == nil {
		op = &xxx_ReleaseNTMSCleanerSlotOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Library = o.Library
	return op
}

func (o *ReleaseNTMSCleanerSlotRequest) xxx_FromOp(ctx context.Context, op *xxx_ReleaseNTMSCleanerSlotOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Library = op.Library
}
func (o *ReleaseNTMSCleanerSlotRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReleaseNTMSCleanerSlotRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReleaseNTMSCleanerSlotOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReleaseNTMSCleanerSlotResponse structure represents the ReleaseNtmsCleanerSlot operation response
type ReleaseNTMSCleanerSlotResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReleaseNtmsCleanerSlot return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReleaseNTMSCleanerSlotResponse) xxx_ToOp(ctx context.Context, op *xxx_ReleaseNTMSCleanerSlotOperation) *xxx_ReleaseNTMSCleanerSlotOperation {
	if op == nil {
		op = &xxx_ReleaseNTMSCleanerSlotOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ReleaseNTMSCleanerSlotResponse) xxx_FromOp(ctx context.Context, op *xxx_ReleaseNTMSCleanerSlotOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ReleaseNTMSCleanerSlotResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReleaseNTMSCleanerSlotResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReleaseNTMSCleanerSlotOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InjectNTMSCleanerOperation structure represents the InjectNtmsCleaner operation
type xxx_InjectNTMSCleanerOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	Library            *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
	InjectOperation    *dtyp.GUID     `idl:"name:lpInjectOperation" json:"inject_operation"`
	NumberOfCleansLeft uint32         `idl:"name:dwNumberOfCleansLeft" json:"number_of_cleans_left"`
	Action             uint32         `idl:"name:dwAction" json:"action"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InjectNTMSCleanerOperation) OpNum() int { return 10 }

func (o *xxx_InjectNTMSCleanerOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/InjectNtmsCleaner"
}

func (o *xxx_InjectNTMSCleanerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSCleanerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library != nil {
			if err := o.Library.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation != nil {
			if err := o.InjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwNumberOfCleansLeft {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfCleansLeft); err != nil {
			return err
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSCleanerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library == nil {
			o.Library = &dtyp.GUID{}
		}
		if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation == nil {
			o.InjectOperation = &dtyp.GUID{}
		}
		if err := o.InjectOperation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwNumberOfCleansLeft {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfCleansLeft); err != nil {
			return err
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSCleanerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InjectNTMSCleanerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation != nil {
			if err := o.InjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_InjectNTMSCleanerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpInjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.InjectOperation == nil {
			o.InjectOperation = &dtyp.GUID{}
		}
		if err := o.InjectOperation.UnmarshalNDR(ctx, w); err != nil {
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

// InjectNTMSCleanerRequest structure represents the InjectNtmsCleaner operation request
type InjectNTMSCleanerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	Library            *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
	InjectOperation    *dtyp.GUID     `idl:"name:lpInjectOperation" json:"inject_operation"`
	NumberOfCleansLeft uint32         `idl:"name:dwNumberOfCleansLeft" json:"number_of_cleans_left"`
	Action             uint32         `idl:"name:dwAction" json:"action"`
}

func (o *InjectNTMSCleanerRequest) xxx_ToOp(ctx context.Context, op *xxx_InjectNTMSCleanerOperation) *xxx_InjectNTMSCleanerOperation {
	if op == nil {
		op = &xxx_InjectNTMSCleanerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Library = o.Library
	op.InjectOperation = o.InjectOperation
	op.NumberOfCleansLeft = o.NumberOfCleansLeft
	op.Action = o.Action
	return op
}

func (o *InjectNTMSCleanerRequest) xxx_FromOp(ctx context.Context, op *xxx_InjectNTMSCleanerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Library = op.Library
	o.InjectOperation = op.InjectOperation
	o.NumberOfCleansLeft = op.NumberOfCleansLeft
	o.Action = op.Action
}
func (o *InjectNTMSCleanerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InjectNTMSCleanerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InjectNTMSCleanerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InjectNTMSCleanerResponse structure represents the InjectNtmsCleaner operation response
type InjectNTMSCleanerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	InjectOperation *dtyp.GUID     `idl:"name:lpInjectOperation" json:"inject_operation"`
	// Return: The InjectNtmsCleaner return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InjectNTMSCleanerResponse) xxx_ToOp(ctx context.Context, op *xxx_InjectNTMSCleanerOperation) *xxx_InjectNTMSCleanerOperation {
	if op == nil {
		op = &xxx_InjectNTMSCleanerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.InjectOperation = o.InjectOperation
	op.Return = o.Return
	return op
}

func (o *InjectNTMSCleanerResponse) xxx_FromOp(ctx context.Context, op *xxx_InjectNTMSCleanerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InjectOperation = op.InjectOperation
	o.Return = op.Return
}
func (o *InjectNTMSCleanerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InjectNTMSCleanerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InjectNTMSCleanerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EjectNTMSCleanerOperation structure represents the EjectNtmsCleaner operation
type xxx_EjectNTMSCleanerOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	Library        *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
	EjectOperation *dtyp.GUID     `idl:"name:lpEjectOperation" json:"eject_operation"`
	Action         uint32         `idl:"name:dwAction" json:"action"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EjectNTMSCleanerOperation) OpNum() int { return 11 }

func (o *xxx_EjectNTMSCleanerOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/EjectNtmsCleaner"
}

func (o *xxx_EjectNTMSCleanerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSCleanerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library != nil {
			if err := o.Library.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation != nil {
			if err := o.EjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSCleanerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibrary {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.Library == nil {
			o.Library = &dtyp.GUID{}
		}
		if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation == nil {
			o.EjectOperation = &dtyp.GUID{}
		}
		if err := o.EjectOperation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSCleanerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EjectNTMSCleanerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation != nil {
			if err := o.EjectOperation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EjectNTMSCleanerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpEjectOperation {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.EjectOperation == nil {
			o.EjectOperation = &dtyp.GUID{}
		}
		if err := o.EjectOperation.UnmarshalNDR(ctx, w); err != nil {
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

// EjectNTMSCleanerRequest structure represents the EjectNtmsCleaner operation request
type EjectNTMSCleanerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	Library        *dtyp.GUID     `idl:"name:lpLibrary" json:"library"`
	EjectOperation *dtyp.GUID     `idl:"name:lpEjectOperation" json:"eject_operation"`
	Action         uint32         `idl:"name:dwAction" json:"action"`
}

func (o *EjectNTMSCleanerRequest) xxx_ToOp(ctx context.Context, op *xxx_EjectNTMSCleanerOperation) *xxx_EjectNTMSCleanerOperation {
	if op == nil {
		op = &xxx_EjectNTMSCleanerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Library = o.Library
	op.EjectOperation = o.EjectOperation
	op.Action = o.Action
	return op
}

func (o *EjectNTMSCleanerRequest) xxx_FromOp(ctx context.Context, op *xxx_EjectNTMSCleanerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Library = op.Library
	o.EjectOperation = op.EjectOperation
	o.Action = op.Action
}
func (o *EjectNTMSCleanerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EjectNTMSCleanerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EjectNTMSCleanerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EjectNTMSCleanerResponse structure represents the EjectNtmsCleaner operation response
type EjectNTMSCleanerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	EjectOperation *dtyp.GUID     `idl:"name:lpEjectOperation" json:"eject_operation"`
	// Return: The EjectNtmsCleaner return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EjectNTMSCleanerResponse) xxx_ToOp(ctx context.Context, op *xxx_EjectNTMSCleanerOperation) *xxx_EjectNTMSCleanerOperation {
	if op == nil {
		op = &xxx_EjectNTMSCleanerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.EjectOperation = o.EjectOperation
	op.Return = o.Return
	return op
}

func (o *EjectNTMSCleanerResponse) xxx_FromOp(ctx context.Context, op *xxx_EjectNTMSCleanerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EjectOperation = op.EjectOperation
	o.Return = op.Return
}
func (o *EjectNTMSCleanerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EjectNTMSCleanerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EjectNTMSCleanerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteNTMSLibraryOperation structure represents the DeleteNtmsLibrary operation
type xxx_DeleteNTMSLibraryOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	LibraryID *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteNTMSLibraryOperation) OpNum() int { return 12 }

func (o *xxx_DeleteNTMSLibraryOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/DeleteNtmsLibrary"
}

func (o *xxx_DeleteNTMSLibraryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSLibraryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID != nil {
			if err := o.LibraryID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSLibraryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpLibraryId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibraryID == nil {
			o.LibraryID = &dtyp.GUID{}
		}
		if err := o.LibraryID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSLibraryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSLibraryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteNTMSLibraryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteNTMSLibraryRequest structure represents the DeleteNtmsLibrary operation request
type DeleteNTMSLibraryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	LibraryID *dtyp.GUID     `idl:"name:lpLibraryId" json:"library_id"`
}

func (o *DeleteNTMSLibraryRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSLibraryOperation) *xxx_DeleteNTMSLibraryOperation {
	if op == nil {
		op = &xxx_DeleteNTMSLibraryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LibraryID = o.LibraryID
	return op
}

func (o *DeleteNTMSLibraryRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSLibraryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LibraryID = op.LibraryID
}
func (o *DeleteNTMSLibraryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteNTMSLibraryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSLibraryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteNTMSLibraryResponse structure represents the DeleteNtmsLibrary operation response
type DeleteNTMSLibraryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteNtmsLibrary return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteNTMSLibraryResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSLibraryOperation) *xxx_DeleteNTMSLibraryOperation {
	if op == nil {
		op = &xxx_DeleteNTMSLibraryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteNTMSLibraryResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSLibraryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteNTMSLibraryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteNTMSLibraryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSLibraryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteNTMSDriveOperation structure represents the DeleteNtmsDrive operation
type xxx_DeleteNTMSDriveOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	DriveID *dtyp.GUID     `idl:"name:lpDriveId" json:"drive_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteNTMSDriveOperation) OpNum() int { return 13 }

func (o *xxx_DeleteNTMSDriveOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/DeleteNtmsDrive"
}

func (o *xxx_DeleteNTMSDriveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSDriveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpDriveId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.DriveID != nil {
			if err := o.DriveID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSDriveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpDriveId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.DriveID == nil {
			o.DriveID = &dtyp.GUID{}
		}
		if err := o.DriveID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSDriveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSDriveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteNTMSDriveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteNTMSDriveRequest structure represents the DeleteNtmsDrive operation request
type DeleteNTMSDriveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	DriveID *dtyp.GUID     `idl:"name:lpDriveId" json:"drive_id"`
}

func (o *DeleteNTMSDriveRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSDriveOperation) *xxx_DeleteNTMSDriveOperation {
	if op == nil {
		op = &xxx_DeleteNTMSDriveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DriveID = o.DriveID
	return op
}

func (o *DeleteNTMSDriveRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSDriveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DriveID = op.DriveID
}
func (o *DeleteNTMSDriveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteNTMSDriveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSDriveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteNTMSDriveResponse structure represents the DeleteNtmsDrive operation response
type DeleteNTMSDriveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteNtmsDrive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteNTMSDriveResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSDriveOperation) *xxx_DeleteNTMSDriveOperation {
	if op == nil {
		op = &xxx_DeleteNTMSDriveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteNTMSDriveResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSDriveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteNTMSDriveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteNTMSDriveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSDriveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSRequestOrderOperation structure represents the GetNtmsRequestOrder operation
type xxx_GetNTMSRequestOrderOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestID   *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	OrderNumber uint32         `idl:"name:lpdwOrderNumber" json:"order_number"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSRequestOrderOperation) OpNum() int { return 14 }

func (o *xxx_GetNTMSRequestOrderOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/GetNtmsRequestOrder"
}

func (o *xxx_GetNTMSRequestOrderOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSRequestOrderOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNTMSRequestOrderOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSRequestOrderOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSRequestOrderOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpdwOrderNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OrderNumber); err != nil {
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

func (o *xxx_GetNTMSRequestOrderOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpdwOrderNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OrderNumber); err != nil {
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

// GetNTMSRequestOrderRequest structure represents the GetNtmsRequestOrder operation request
type GetNTMSRequestOrderRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
}

func (o *GetNTMSRequestOrderRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSRequestOrderOperation) *xxx_GetNTMSRequestOrderOperation {
	if op == nil {
		op = &xxx_GetNTMSRequestOrderOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestID = o.RequestID
	return op
}

func (o *GetNTMSRequestOrderRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSRequestOrderOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestID = op.RequestID
}
func (o *GetNTMSRequestOrderRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSRequestOrderRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSRequestOrderOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSRequestOrderResponse structure represents the GetNtmsRequestOrder operation response
type GetNTMSRequestOrderResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	OrderNumber uint32         `idl:"name:lpdwOrderNumber" json:"order_number"`
	// Return: The GetNtmsRequestOrder return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSRequestOrderResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSRequestOrderOperation) *xxx_GetNTMSRequestOrderOperation {
	if op == nil {
		op = &xxx_GetNTMSRequestOrderOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.OrderNumber = o.OrderNumber
	op.Return = o.Return
	return op
}

func (o *GetNTMSRequestOrderResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSRequestOrderOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.OrderNumber = op.OrderNumber
	o.Return = op.Return
}
func (o *GetNTMSRequestOrderResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSRequestOrderResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSRequestOrderOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSRequestOrderOperation structure represents the SetNtmsRequestOrder operation
type xxx_SetNTMSRequestOrderOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestID   *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	OrderNumber uint32         `idl:"name:dwOrderNumber" json:"order_number"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSRequestOrderOperation) OpNum() int { return 15 }

func (o *xxx_SetNTMSRequestOrderOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/SetNtmsRequestOrder"
}

func (o *xxx_SetNTMSRequestOrderOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSRequestOrderOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwOrderNumber {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OrderNumber); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSRequestOrderOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwOrderNumber {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OrderNumber); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSRequestOrderOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSRequestOrderOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSRequestOrderOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSRequestOrderRequest structure represents the SetNtmsRequestOrder operation request
type SetNTMSRequestOrderRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	RequestID   *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	OrderNumber uint32         `idl:"name:dwOrderNumber" json:"order_number"`
}

func (o *SetNTMSRequestOrderRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSRequestOrderOperation) *xxx_SetNTMSRequestOrderOperation {
	if op == nil {
		op = &xxx_SetNTMSRequestOrderOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestID = o.RequestID
	op.OrderNumber = o.OrderNumber
	return op
}

func (o *SetNTMSRequestOrderRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSRequestOrderOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestID = op.RequestID
	o.OrderNumber = op.OrderNumber
}
func (o *SetNTMSRequestOrderRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSRequestOrderRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSRequestOrderOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSRequestOrderResponse structure represents the SetNtmsRequestOrder operation response
type SetNTMSRequestOrderResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsRequestOrder return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSRequestOrderResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSRequestOrderOperation) *xxx_SetNTMSRequestOrderOperation {
	if op == nil {
		op = &xxx_SetNTMSRequestOrderOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSRequestOrderResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSRequestOrderOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSRequestOrderResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSRequestOrderResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSRequestOrderOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteNTMSRequestsOperation structure represents the DeleteNtmsRequests operation
type xxx_DeleteNTMSRequestsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestID []*dtyp.GUID   `idl:"name:lpRequestId;size_is:(dwCount)" json:"request_id"`
	Type      uint32         `idl:"name:dwType" json:"type"`
	Count     uint32         `idl:"name:dwCount" json:"count"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteNTMSRequestsOperation) OpNum() int { return 16 }

func (o *xxx_DeleteNTMSRequestsOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/DeleteNtmsRequests"
}

func (o *xxx_DeleteNTMSRequestsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RequestID != nil && o.Count == 0 {
		o.Count = uint32(len(o.RequestID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSRequestsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.RequestID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.RequestID[i1] != nil {
				if err := o.RequestID[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.RequestID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSRequestsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.RequestID", sizeInfo[0])
		}
		o.RequestID = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.RequestID {
			i1 := i1
			if o.RequestID[i1] == nil {
				o.RequestID[i1] = &dtyp.GUID{}
			}
			if err := o.RequestID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSRequestsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSRequestsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteNTMSRequestsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteNTMSRequestsRequest structure represents the DeleteNtmsRequests operation request
type DeleteNTMSRequestsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	RequestID []*dtyp.GUID   `idl:"name:lpRequestId;size_is:(dwCount)" json:"request_id"`
	Type      uint32         `idl:"name:dwType" json:"type"`
	Count     uint32         `idl:"name:dwCount" json:"count"`
}

func (o *DeleteNTMSRequestsRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSRequestsOperation) *xxx_DeleteNTMSRequestsOperation {
	if op == nil {
		op = &xxx_DeleteNTMSRequestsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestID = o.RequestID
	op.Type = o.Type
	op.Count = o.Count
	return op
}

func (o *DeleteNTMSRequestsRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSRequestsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestID = op.RequestID
	o.Type = op.Type
	o.Count = op.Count
}
func (o *DeleteNTMSRequestsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteNTMSRequestsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSRequestsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteNTMSRequestsResponse structure represents the DeleteNtmsRequests operation response
type DeleteNTMSRequestsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteNtmsRequests return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteNTMSRequestsResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSRequestsOperation) *xxx_DeleteNTMSRequestsOperation {
	if op == nil {
		op = &xxx_DeleteNTMSRequestsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteNTMSRequestsResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSRequestsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteNTMSRequestsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteNTMSRequestsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSRequestsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BeginNTMSDeviceChangeDetectionOperation structure represents the BeginNtmsDeviceChangeDetection operation
type xxx_BeginNTMSDeviceChangeDetectionOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	DetectHandle uint64         `idl:"name:lpDetectHandle" json:"detect_handle"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) OpNum() int { return 17 }

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/BeginNtmsDeviceChangeDetection"
}

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpDetectHandle {out} (1:{pointer=ref}*(1))(2:{alias=NTMS_HANDLE, names=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.DetectHandle)); err != nil {
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

func (o *xxx_BeginNTMSDeviceChangeDetectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpDetectHandle {out} (1:{pointer=ref}*(1))(2:{alias=NTMS_HANDLE, names=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.DetectHandle)); err != nil {
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

// BeginNTMSDeviceChangeDetectionRequest structure represents the BeginNtmsDeviceChangeDetection operation request
type BeginNTMSDeviceChangeDetectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BeginNTMSDeviceChangeDetectionRequest) xxx_ToOp(ctx context.Context, op *xxx_BeginNTMSDeviceChangeDetectionOperation) *xxx_BeginNTMSDeviceChangeDetectionOperation {
	if op == nil {
		op = &xxx_BeginNTMSDeviceChangeDetectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *BeginNTMSDeviceChangeDetectionRequest) xxx_FromOp(ctx context.Context, op *xxx_BeginNTMSDeviceChangeDetectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BeginNTMSDeviceChangeDetectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BeginNTMSDeviceChangeDetectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginNTMSDeviceChangeDetectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BeginNTMSDeviceChangeDetectionResponse structure represents the BeginNtmsDeviceChangeDetection operation response
type BeginNTMSDeviceChangeDetectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	DetectHandle uint64         `idl:"name:lpDetectHandle" json:"detect_handle"`
	// Return: The BeginNtmsDeviceChangeDetection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BeginNTMSDeviceChangeDetectionResponse) xxx_ToOp(ctx context.Context, op *xxx_BeginNTMSDeviceChangeDetectionOperation) *xxx_BeginNTMSDeviceChangeDetectionOperation {
	if op == nil {
		op = &xxx_BeginNTMSDeviceChangeDetectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DetectHandle = o.DetectHandle
	op.Return = o.Return
	return op
}

func (o *BeginNTMSDeviceChangeDetectionResponse) xxx_FromOp(ctx context.Context, op *xxx_BeginNTMSDeviceChangeDetectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DetectHandle = op.DetectHandle
	o.Return = op.Return
}
func (o *BeginNTMSDeviceChangeDetectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BeginNTMSDeviceChangeDetectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginNTMSDeviceChangeDetectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSDeviceChangeDetectionOperation structure represents the SetNtmsDeviceChangeDetection operation
type xxx_SetNTMSDeviceChangeDetectionOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	DetectHandle uint64         `idl:"name:DetectHandle" json:"detect_handle"`
	ObjectID     []*dtyp.GUID   `idl:"name:lpObjectId;size_is:(dwCount)" json:"object_id"`
	Type         uint32         `idl:"name:dwType" json:"type"`
	Count        uint32         `idl:"name:dwCount" json:"count"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) OpNum() int { return 18 }

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/SetNtmsDeviceChangeDetection"
}

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ObjectID != nil && o.Count == 0 {
		o.Count = uint32(len(o.ObjectID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// DetectHandle {in} (1:{alias=NTMS_HANDLE, names=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.DetectHandle)); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ObjectID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ObjectID[i1] != nil {
				if err := o.ObjectID[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ObjectID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// DetectHandle {in} (1:{alias=NTMS_HANDLE, names=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.DetectHandle)); err != nil {
			return err
		}
	}
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ObjectID", sizeInfo[0])
		}
		o.ObjectID = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.ObjectID {
			i1 := i1
			if o.ObjectID[i1] == nil {
				o.ObjectID[i1] = &dtyp.GUID{}
			}
			if err := o.ObjectID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSDeviceChangeDetectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSDeviceChangeDetectionRequest structure represents the SetNtmsDeviceChangeDetection operation request
type SetNTMSDeviceChangeDetectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	DetectHandle uint64         `idl:"name:DetectHandle" json:"detect_handle"`
	ObjectID     []*dtyp.GUID   `idl:"name:lpObjectId;size_is:(dwCount)" json:"object_id"`
	Type         uint32         `idl:"name:dwType" json:"type"`
	Count        uint32         `idl:"name:dwCount" json:"count"`
}

func (o *SetNTMSDeviceChangeDetectionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSDeviceChangeDetectionOperation) *xxx_SetNTMSDeviceChangeDetectionOperation {
	if op == nil {
		op = &xxx_SetNTMSDeviceChangeDetectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DetectHandle = o.DetectHandle
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.Count = o.Count
	return op
}

func (o *SetNTMSDeviceChangeDetectionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSDeviceChangeDetectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DetectHandle = op.DetectHandle
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.Count = op.Count
}
func (o *SetNTMSDeviceChangeDetectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSDeviceChangeDetectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSDeviceChangeDetectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSDeviceChangeDetectionResponse structure represents the SetNtmsDeviceChangeDetection operation response
type SetNTMSDeviceChangeDetectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsDeviceChangeDetection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSDeviceChangeDetectionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSDeviceChangeDetectionOperation) *xxx_SetNTMSDeviceChangeDetectionOperation {
	if op == nil {
		op = &xxx_SetNTMSDeviceChangeDetectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSDeviceChangeDetectionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSDeviceChangeDetectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSDeviceChangeDetectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSDeviceChangeDetectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSDeviceChangeDetectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EndNTMSDeviceChangeDetectionOperation structure represents the EndNtmsDeviceChangeDetection operation
type xxx_EndNTMSDeviceChangeDetectionOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	DetectHandle uint64         `idl:"name:DetectHandle" json:"detect_handle"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) OpNum() int { return 19 }

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) OpName() string {
	return "/INtmsLibraryControl1/v0/EndNtmsDeviceChangeDetection"
}

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// DetectHandle {in} (1:{alias=NTMS_HANDLE, names=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.DetectHandle)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// DetectHandle {in} (1:{alias=NTMS_HANDLE, names=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.DetectHandle)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EndNTMSDeviceChangeDetectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EndNTMSDeviceChangeDetectionRequest structure represents the EndNtmsDeviceChangeDetection operation request
type EndNTMSDeviceChangeDetectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	DetectHandle uint64         `idl:"name:DetectHandle" json:"detect_handle"`
}

func (o *EndNTMSDeviceChangeDetectionRequest) xxx_ToOp(ctx context.Context, op *xxx_EndNTMSDeviceChangeDetectionOperation) *xxx_EndNTMSDeviceChangeDetectionOperation {
	if op == nil {
		op = &xxx_EndNTMSDeviceChangeDetectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DetectHandle = o.DetectHandle
	return op
}

func (o *EndNTMSDeviceChangeDetectionRequest) xxx_FromOp(ctx context.Context, op *xxx_EndNTMSDeviceChangeDetectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DetectHandle = op.DetectHandle
}
func (o *EndNTMSDeviceChangeDetectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EndNTMSDeviceChangeDetectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndNTMSDeviceChangeDetectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EndNTMSDeviceChangeDetectionResponse structure represents the EndNtmsDeviceChangeDetection operation response
type EndNTMSDeviceChangeDetectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EndNtmsDeviceChangeDetection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EndNTMSDeviceChangeDetectionResponse) xxx_ToOp(ctx context.Context, op *xxx_EndNTMSDeviceChangeDetectionOperation) *xxx_EndNTMSDeviceChangeDetectionOperation {
	if op == nil {
		op = &xxx_EndNTMSDeviceChangeDetectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *EndNTMSDeviceChangeDetectionResponse) xxx_FromOp(ctx context.Context, op *xxx_EndNTMSDeviceChangeDetectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EndNTMSDeviceChangeDetectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EndNTMSDeviceChangeDetectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndNTMSDeviceChangeDetectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

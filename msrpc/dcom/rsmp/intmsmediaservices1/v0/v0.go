package intmsmediaservices1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	rsmp "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
	_ = rsmp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsMediaServices1 interface identifier d02e4be0-3419-11d1-8fb1-00a024cb6019
	MediaServices1IID = &dcom.IID{Data1: 0xd02e4be0, Data2: 0x3419, Data3: 0x11d1, Data4: []byte{0x8f, 0xb1, 0x00, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax UUID
	MediaServices1SyntaxUUID = &uuid.UUID{TimeLow: 0xd02e4be0, TimeMid: 0x3419, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xb1, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax ID
	MediaServices1SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: MediaServices1SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsMediaServices1 interface.
type MediaServices1Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The MountNtmsMedia method mounts one or more pieces of media.
	MountNTMSMedia(context.Context, *MountNTMSMediaRequest, ...dcerpc.CallOption) (*MountNTMSMediaResponse, error)

	// The DismountNtmsMedia method queues a command to move a medium in a drive to its
	// storage.
	DismountNTMSMedia(context.Context, *DismountNTMSMediaRequest, ...dcerpc.CallOption) (*DismountNTMSMediaResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// The AllocateNtmsMedia method allocates a piece of available media.
	AllocateNTMSMedia(context.Context, *AllocateNTMSMediaRequest, ...dcerpc.CallOption) (*AllocateNTMSMediaResponse, error)

	// The DeallocateNtmsMedia method deallocates the side that is associated with a piece
	// of logical media.
	DeallocateNTMSMedia(context.Context, *DeallocateNTMSMediaRequest, ...dcerpc.CallOption) (*DeallocateNTMSMediaResponse, error)

	// The SwapNtmsMedia method swaps the position of two media sides.
	SwapNTMSMedia(context.Context, *SwapNTMSMediaRequest, ...dcerpc.CallOption) (*SwapNTMSMediaResponse, error)

	// The DecommissionNtmsMedia method moves media from available state to decommissioned
	// state. Media that are decommissioned by the DecommissionNtmsMedia method are recognized
	// by the server, but decommissioned media does not contain any data and is never again
	// used. Only media that are in an available state can be decommissioned.
	DecommissionNTMSMedia(context.Context, *DecommissionNTMSMediaRequest, ...dcerpc.CallOption) (*DecommissionNTMSMediaResponse, error)

	// The SetNtmsMediaComplete method marks a piece of logical media as complete.
	SetNTMSMediaComplete(context.Context, *SetNTMSMediaCompleteRequest, ...dcerpc.CallOption) (*SetNTMSMediaCompleteResponse, error)

	// The DeleteNtmsMedia method deletes a physical piece of offline media by removing
	// all references to it.
	DeleteNTMSMedia(context.Context, *DeleteNTMSMediaRequest, ...dcerpc.CallOption) (*DeleteNTMSMediaResponse, error)

	// The CreateNtmsMediaPoolA method creates a new application media pool, with strings
	// encoded using ASCII.
	CreateNTMSMediaPoolA(context.Context, *CreateNTMSMediaPoolARequest, ...dcerpc.CallOption) (*CreateNTMSMediaPoolAResponse, error)

	// The CreateNtmsMediaPoolW method creates a new application media pool whose name is
	// composed of a sequence of Unicode characters.
	CreateNTMSMediaPoolW(context.Context, *CreateNTMSMediaPoolWRequest, ...dcerpc.CallOption) (*CreateNTMSMediaPoolWResponse, error)

	// The GetNtmsMediaPoolNameA method retrieves the full name hierarchy of a media pool,
	// with null-terminated strings encoded using ASCII.
	GetNTMSMediaPoolNameA(context.Context, *GetNTMSMediaPoolNameARequest, ...dcerpc.CallOption) (*GetNTMSMediaPoolNameAResponse, error)

	// The GetNtmsMediaPoolNameW method retrieves the full name hierarchy of a media pool,
	// with strings encoded using Unicode.
	GetNTMSMediaPoolNameW(context.Context, *GetNTMSMediaPoolNameWRequest, ...dcerpc.CallOption) (*GetNTMSMediaPoolNameWResponse, error)

	// The MoveToNtmsMediaPool method moves a medium from its current media pool to another
	// media pool.
	MoveToNTMSMediaPool(context.Context, *MoveToNTMSMediaPoolRequest, ...dcerpc.CallOption) (*MoveToNTMSMediaPoolResponse, error)

	// The DeleteNtmsMediaPool method deletes an application media pool.
	DeleteNTMSMediaPool(context.Context, *DeleteNTMSMediaPoolRequest, ...dcerpc.CallOption) (*DeleteNTMSMediaPoolResponse, error)

	// The AddNtmsMediaType method MUST add a media type to a library if there is not currently
	// a relation in the library. The method MUST create the system media pools (FREE, IMPORT,
	// and UNRECOGNIZED) if they do not exist.
	AddNTMSMediaType(context.Context, *AddNTMSMediaTypeRequest, ...dcerpc.CallOption) (*AddNTMSMediaTypeResponse, error)

	// The DeleteNtmsMediaType method deletes a media type from a library.
	DeleteNTMSMediaType(context.Context, *DeleteNTMSMediaTypeRequest, ...dcerpc.CallOption) (*DeleteNTMSMediaTypeResponse, error)

	// The ChangeNtmsMediaType method moves a physical media identifier to a new media pool
	// and sets the media type of the medium to that of the pool.
	ChangeNTMSMediaType(context.Context, *ChangeNTMSMediaTypeRequest, ...dcerpc.CallOption) (*ChangeNTMSMediaTypeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) MediaServices1Client
}

type xxx_DefaultMediaServices1Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultMediaServices1Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultMediaServices1Client) MountNTMSMedia(ctx context.Context, in *MountNTMSMediaRequest, opts ...dcerpc.CallOption) (*MountNTMSMediaResponse, error) {
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
	out := &MountNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) DismountNTMSMedia(ctx context.Context, in *DismountNTMSMediaRequest, opts ...dcerpc.CallOption) (*DismountNTMSMediaResponse, error) {
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
	out := &DismountNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) AllocateNTMSMedia(ctx context.Context, in *AllocateNTMSMediaRequest, opts ...dcerpc.CallOption) (*AllocateNTMSMediaResponse, error) {
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
	out := &AllocateNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) DeallocateNTMSMedia(ctx context.Context, in *DeallocateNTMSMediaRequest, opts ...dcerpc.CallOption) (*DeallocateNTMSMediaResponse, error) {
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
	out := &DeallocateNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) SwapNTMSMedia(ctx context.Context, in *SwapNTMSMediaRequest, opts ...dcerpc.CallOption) (*SwapNTMSMediaResponse, error) {
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
	out := &SwapNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) DecommissionNTMSMedia(ctx context.Context, in *DecommissionNTMSMediaRequest, opts ...dcerpc.CallOption) (*DecommissionNTMSMediaResponse, error) {
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
	out := &DecommissionNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) SetNTMSMediaComplete(ctx context.Context, in *SetNTMSMediaCompleteRequest, opts ...dcerpc.CallOption) (*SetNTMSMediaCompleteResponse, error) {
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
	out := &SetNTMSMediaCompleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) DeleteNTMSMedia(ctx context.Context, in *DeleteNTMSMediaRequest, opts ...dcerpc.CallOption) (*DeleteNTMSMediaResponse, error) {
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
	out := &DeleteNTMSMediaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) CreateNTMSMediaPoolA(ctx context.Context, in *CreateNTMSMediaPoolARequest, opts ...dcerpc.CallOption) (*CreateNTMSMediaPoolAResponse, error) {
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
	out := &CreateNTMSMediaPoolAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) CreateNTMSMediaPoolW(ctx context.Context, in *CreateNTMSMediaPoolWRequest, opts ...dcerpc.CallOption) (*CreateNTMSMediaPoolWResponse, error) {
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
	out := &CreateNTMSMediaPoolWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) GetNTMSMediaPoolNameA(ctx context.Context, in *GetNTMSMediaPoolNameARequest, opts ...dcerpc.CallOption) (*GetNTMSMediaPoolNameAResponse, error) {
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
	out := &GetNTMSMediaPoolNameAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) GetNTMSMediaPoolNameW(ctx context.Context, in *GetNTMSMediaPoolNameWRequest, opts ...dcerpc.CallOption) (*GetNTMSMediaPoolNameWResponse, error) {
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
	out := &GetNTMSMediaPoolNameWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) MoveToNTMSMediaPool(ctx context.Context, in *MoveToNTMSMediaPoolRequest, opts ...dcerpc.CallOption) (*MoveToNTMSMediaPoolResponse, error) {
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
	out := &MoveToNTMSMediaPoolResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) DeleteNTMSMediaPool(ctx context.Context, in *DeleteNTMSMediaPoolRequest, opts ...dcerpc.CallOption) (*DeleteNTMSMediaPoolResponse, error) {
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
	out := &DeleteNTMSMediaPoolResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) AddNTMSMediaType(ctx context.Context, in *AddNTMSMediaTypeRequest, opts ...dcerpc.CallOption) (*AddNTMSMediaTypeResponse, error) {
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
	out := &AddNTMSMediaTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) DeleteNTMSMediaType(ctx context.Context, in *DeleteNTMSMediaTypeRequest, opts ...dcerpc.CallOption) (*DeleteNTMSMediaTypeResponse, error) {
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
	out := &DeleteNTMSMediaTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) ChangeNTMSMediaType(ctx context.Context, in *ChangeNTMSMediaTypeRequest, opts ...dcerpc.CallOption) (*ChangeNTMSMediaTypeResponse, error) {
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
	out := &ChangeNTMSMediaTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMediaServices1Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMediaServices1Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultMediaServices1Client) IPID(ctx context.Context, ipid *dcom.IPID) MediaServices1Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultMediaServices1Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewMediaServices1Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (MediaServices1Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(MediaServices1SyntaxV0_0))...)
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
	return &xxx_DefaultMediaServices1Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_MountNTMSMediaOperation structure represents the MountNtmsMedia operation
type xxx_MountNTMSMediaOperation struct {
	This             *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat         `idl:"name:That" json:"that"`
	MediaID          []*dtyp.GUID           `idl:"name:lpMediaId;size_is:(dwCount)" json:"media_id"`
	DriveID          []*dtyp.GUID           `idl:"name:lpDriveId;size_is:(dwCount)" json:"drive_id"`
	Count            uint32                 `idl:"name:dwCount" json:"count"`
	Options          uint32                 `idl:"name:dwOptions" json:"options"`
	Priority         int32                  `idl:"name:dwPriority" json:"priority"`
	Timeout          uint32                 `idl:"name:dwTimeout" json:"timeout"`
	MountInformation *rsmp.MountInformation `idl:"name:lpMountInformation" json:"mount_information"`
	Return           int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_MountNTMSMediaOperation) OpNum() int { return 0 }

func (o *xxx_MountNTMSMediaOperation) OpName() string {
	return "/INtmsMediaServices1/v0/MountNtmsMedia"
}

func (o *xxx_MountNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.MediaID != nil && o.Count == 0 {
		o.Count = uint32(len(o.MediaID))
	}
	if o.DriveID != nil && o.Count == 0 {
		o.Count = uint32(len(o.DriveID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MountNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.MediaID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.MediaID[i1] != nil {
				if err := o.MediaID[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.MediaID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpDriveId {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DriveID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DriveID[i1] != nil {
				if err := o.DriveID[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DriveID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	// dwPriority {in} (1:(int32))
	{
		if err := w.WriteData(o.Priority); err != nil {
			return err
		}
	}
	// dwTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// lpMountInformation {in, out} (1:{alias=LPNTMS_MOUNT_INFORMATION}*(1))(2:{alias=NTMS_MOUNT_INFORMATION}(struct))
	{
		if o.MountInformation != nil {
			if err := o.MountInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.MountInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MountNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.MediaID", sizeInfo[0])
		}
		o.MediaID = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.MediaID {
			i1 := i1
			if o.MediaID[i1] == nil {
				o.MediaID[i1] = &dtyp.GUID{}
			}
			if err := o.MediaID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpDriveId {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.DriveID", sizeInfo[0])
		}
		o.DriveID = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.DriveID {
			i1 := i1
			if o.DriveID[i1] == nil {
				o.DriveID[i1] = &dtyp.GUID{}
			}
			if err := o.DriveID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	// dwPriority {in} (1:(int32))
	{
		if err := w.ReadData(&o.Priority); err != nil {
			return err
		}
	}
	// dwTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// lpMountInformation {in, out} (1:{alias=LPNTMS_MOUNT_INFORMATION,pointer=ref}*(1))(2:{alias=NTMS_MOUNT_INFORMATION}(struct))
	{
		if o.MountInformation == nil {
			o.MountInformation = &rsmp.MountInformation{}
		}
		if err := o.MountInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MountNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MountNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpDriveId {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DriveID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DriveID[i1] != nil {
				if err := o.DriveID[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DriveID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpMountInformation {in, out} (1:{alias=LPNTMS_MOUNT_INFORMATION}*(1))(2:{alias=NTMS_MOUNT_INFORMATION}(struct))
	{
		if o.MountInformation != nil {
			if err := o.MountInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.MountInformation{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MountNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpDriveId {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.DriveID", sizeInfo[0])
		}
		o.DriveID = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.DriveID {
			i1 := i1
			if o.DriveID[i1] == nil {
				o.DriveID[i1] = &dtyp.GUID{}
			}
			if err := o.DriveID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpMountInformation {in, out} (1:{alias=LPNTMS_MOUNT_INFORMATION,pointer=ref}*(1))(2:{alias=NTMS_MOUNT_INFORMATION}(struct))
	{
		if o.MountInformation == nil {
			o.MountInformation = &rsmp.MountInformation{}
		}
		if err := o.MountInformation.UnmarshalNDR(ctx, w); err != nil {
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

// MountNTMSMediaRequest structure represents the MountNtmsMedia operation request
type MountNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: An array of logical media identifiers or media side identifiers. Each
	// entry in the array MUST be unique.
	MediaID []*dtyp.GUID `idl:"name:lpMediaId;size_is:(dwCount)" json:"media_id"`
	// lpDriveId: An array of drive identifiers that correspond to the media listed in the
	// lpMediaId parameter. This array MUST either specify a list of drives into which media
	// will be mounted or receive the list of drives into which media will be mounted when
	// the operation completes.
	DriveID []*dtyp.GUID `idl:"name:lpDriveId;size_is:(dwCount)" json:"drive_id"`
	// dwCount: The number of elements in the lpMediaId and lpDriveId arrays.
	Count uint32 `idl:"name:dwCount" json:"count"`
	// dwOptions: A bitmap of mount options from the NtmsMountOptions (section 2.2.3.3)
	// enumeration.
	Options uint32 `idl:"name:dwOptions" json:"options"`
	// dwPriority: A value from the NtmsMountPriority (section 2.2.3.4) enumeration specifying
	// the priority of the mount request.
	Priority int32 `idl:"name:dwPriority" json:"priority"`
	// dwTimeout: The maximum time, in milliseconds, allowed for mounting of the specified
	// media. To wait as long as the mount takes, this parameter MUST be set to 0xFFFFFFFF.
	// If dwOptions is specified as NTMS_MOUNT_NOWAIT, ignore this value.
	//
	//	+------------+--------------------------------------------------------------------+
	//	|            |                                                                    |
	//	|   VALUE    |                              MEANING                               |
	//	|            |                                                                    |
	//	+------------+--------------------------------------------------------------------+
	//	+------------+--------------------------------------------------------------------+
	//	| 0xFFFFFFFF | Use this value to wait as long as required for the mount to occur. |
	//	+------------+--------------------------------------------------------------------+
	Timeout uint32 `idl:"name:dwTimeout" json:"timeout"`
	// lpMountInformation: This parameter is currently unused. It MUST be sent as NULL and
	// ignored on receipt.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                         | The call was successful.                                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED          | Access to an object was denied.                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY      | A allocation failure occurred during processing.                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007000F ERROR_INVALID_DRIVE          | The drive identifier is not valid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070013 ERROR_WRITE_PROTECT          | The media state is set to NTMS_PARTSTATE_COMPLETE, from the NtmsPartitionState   |
	//	|                                         | enumeration, and the NTMS_MOUNT_WRITE value was specified.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER      | A parameter is not valid.                                                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700AA ERROR_BUSY                   | The media or drives are busy.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800705B4 ERROR_TIMEOUT                | The time-out event expired before the medium was available.                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA          | The media identifier is not valid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CD ERROR_INVALID_LIBRARY        | The library identifier is not valid.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CF ERROR_DRIVE_MEDIA_MISMATCH   | The specified media and drive are not in the same library.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D0 ERROR_MEDIA_OFFLINE          | The specified media is offline and cannot be allocated.                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D5 ERROR_RESOURCE_DISABLED      | A resource required for this operation is disabled.                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE       | The database query or update failed.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL          | The database is full.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710E0 ERROR_REQUEST_REFUSED        | The request is refused as a user canceled the request through the user           |
	//	|                                         | interface.                                                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139E ERROR_RESOURCE_NOT_AVAILABLE | The specified resource is not available.                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE          | An unexpected state was encountered.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704C7 ERROR_CANCELLED              | The request was cancelled.                                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	MountInformation *rsmp.MountInformation `idl:"name:lpMountInformation" json:"mount_information"`
}

func (o *MountNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_MountNTMSMediaOperation) *xxx_MountNTMSMediaOperation {
	if op == nil {
		op = &xxx_MountNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	op.DriveID = o.DriveID
	op.Count = o.Count
	op.Options = o.Options
	op.Priority = o.Priority
	op.Timeout = o.Timeout
	op.MountInformation = o.MountInformation
	return op
}

func (o *MountNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_MountNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
	o.DriveID = op.DriveID
	o.Count = op.Count
	o.Options = op.Options
	o.Priority = op.Priority
	o.Timeout = op.Timeout
	o.MountInformation = op.MountInformation
}
func (o *MountNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MountNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MountNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MountNTMSMediaResponse structure represents the MountNtmsMedia operation response
type MountNTMSMediaResponse struct {
	// XXX: dwCount is an implicit input depedency for output parameters
	Count uint32 `idl:"name:dwCount" json:"count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpDriveId: An array of drive identifiers that correspond to the media listed in the
	// lpMediaId parameter. This array MUST either specify a list of drives into which media
	// will be mounted or receive the list of drives into which media will be mounted when
	// the operation completes.
	DriveID []*dtyp.GUID `idl:"name:lpDriveId;size_is:(dwCount)" json:"drive_id"`
	// lpMountInformation: This parameter is currently unused. It MUST be sent as NULL and
	// ignored on receipt.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                         | The call was successful.                                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED          | Access to an object was denied.                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY      | A allocation failure occurred during processing.                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007000F ERROR_INVALID_DRIVE          | The drive identifier is not valid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070013 ERROR_WRITE_PROTECT          | The media state is set to NTMS_PARTSTATE_COMPLETE, from the NtmsPartitionState   |
	//	|                                         | enumeration, and the NTMS_MOUNT_WRITE value was specified.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER      | A parameter is not valid.                                                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700AA ERROR_BUSY                   | The media or drives are busy.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800705B4 ERROR_TIMEOUT                | The time-out event expired before the medium was available.                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA          | The media identifier is not valid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CD ERROR_INVALID_LIBRARY        | The library identifier is not valid.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CF ERROR_DRIVE_MEDIA_MISMATCH   | The specified media and drive are not in the same library.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D0 ERROR_MEDIA_OFFLINE          | The specified media is offline and cannot be allocated.                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D5 ERROR_RESOURCE_DISABLED      | A resource required for this operation is disabled.                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE       | The database query or update failed.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL          | The database is full.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710E0 ERROR_REQUEST_REFUSED        | The request is refused as a user canceled the request through the user           |
	//	|                                         | interface.                                                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139E ERROR_RESOURCE_NOT_AVAILABLE | The specified resource is not available.                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE          | An unexpected state was encountered.                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704C7 ERROR_CANCELLED              | The request was cancelled.                                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	MountInformation *rsmp.MountInformation `idl:"name:lpMountInformation" json:"mount_information"`
	// Return: The MountNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MountNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_MountNTMSMediaOperation) *xxx_MountNTMSMediaOperation {
	if op == nil {
		op = &xxx_MountNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Count == uint32(0) {
		op.Count = o.Count
	}

	op.That = o.That
	op.DriveID = o.DriveID
	op.MountInformation = o.MountInformation
	op.Return = o.Return
	return op
}

func (o *MountNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_MountNTMSMediaOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Count = op.Count

	o.That = op.That
	o.DriveID = op.DriveID
	o.MountInformation = op.MountInformation
	o.Return = op.Return
}
func (o *MountNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MountNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MountNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DismountNTMSMediaOperation structure represents the DismountNtmsMedia operation
type xxx_DismountNTMSMediaOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID []*dtyp.GUID   `idl:"name:lpMediaId;size_is:(dwCount)" json:"media_id"`
	Count   uint32         `idl:"name:dwCount" json:"count"`
	Options uint32         `idl:"name:dwOptions" json:"options"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DismountNTMSMediaOperation) OpNum() int { return 1 }

func (o *xxx_DismountNTMSMediaOperation) OpName() string {
	return "/INtmsMediaServices1/v0/DismountNtmsMedia"
}

func (o *xxx_DismountNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.MediaID != nil && o.Count == 0 {
		o.Count = uint32(len(o.MediaID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.MediaID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.MediaID[i1] != nil {
				if err := o.MediaID[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.MediaID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.MediaID", sizeInfo[0])
		}
		o.MediaID = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.MediaID {
			i1 := i1
			if o.MediaID[i1] == nil {
				o.MediaID[i1] = &dtyp.GUID{}
			}
			if err := o.MediaID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DismountNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DismountNTMSMediaRequest structure represents the DismountNtmsMedia operation request
type DismountNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: An array of logical media or media side identifiers.
	MediaID []*dtyp.GUID `idl:"name:lpMediaId;size_is:(dwCount)" json:"media_id"`
	// dwCount: The number of elements in the lpMediaId array.
	Count uint32 `idl:"name:dwCount" json:"count"`
	// dwOptions: One of the options from the NtmsDismountOptions (section 2.2.1.9) numeration.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED        | NTMS_USE_ACCESS to the media pool or library that contains the media is denied;  |
	//	|                                       | other security errors are possible, but indicate a security subsystem error.     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER    | A parameter is missing.                                                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE     | The database is inaccessible or damaged.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL        | The database is full.                                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DF ERROR_DEVICE_NOT_AVAILABLE | One or more resources required to perform the dismount are not currently         |
	//	|                                       | available.                                                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CD ERROR_INVALID_LIBRARY      | The library that contains the drives or media is not valid.                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA        | A medium is not valid, or lpMediaId contains duplicate identifiers.              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | An unexpected media or device state occurred during dismount.                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D0 ERROR_MEDIA_OFFLINE        | The specified media is offline.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY    | A memory allocation failure occurred during processing.                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800705B4 ERROR_TIMEOUT              | The time-out event expired while the application attempted to acquire one or     |
	//	|                                       | more resources.                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
}

func (o *DismountNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_DismountNTMSMediaOperation) *xxx_DismountNTMSMediaOperation {
	if op == nil {
		op = &xxx_DismountNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	op.Count = o.Count
	op.Options = o.Options
	return op
}

func (o *DismountNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_DismountNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
	o.Count = op.Count
	o.Options = op.Options
}
func (o *DismountNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DismountNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DismountNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DismountNTMSMediaResponse structure represents the DismountNtmsMedia operation response
type DismountNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DismountNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DismountNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_DismountNTMSMediaOperation) *xxx_DismountNTMSMediaOperation {
	if op == nil {
		op = &xxx_DismountNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DismountNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_DismountNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DismountNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DismountNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DismountNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AllocateNTMSMediaOperation structure represents the AllocateNtmsMedia operation
type xxx_AllocateNTMSMediaOperation struct {
	This                *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat              `idl:"name:That" json:"that"`
	MediaPool           *dtyp.GUID                  `idl:"name:lpMediaPool" json:"media_pool"`
	Partition           *dtyp.GUID                  `idl:"name:lpPartition;pointer:unique" json:"partition"`
	MediaID             *dtyp.GUID                  `idl:"name:lpMediaId" json:"media_id"`
	Options             uint32                      `idl:"name:dwOptions" json:"options"`
	Timeout             uint32                      `idl:"name:dwTimeout" json:"timeout"`
	AllocateInformation *rsmp.AllocationInformation `idl:"name:lpAllocateInformation" json:"allocate_information"`
	Return              int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_AllocateNTMSMediaOperation) OpNum() int { return 3 }

func (o *xxx_AllocateNTMSMediaOperation) OpName() string {
	return "/INtmsMediaServices1/v0/AllocateNtmsMedia"
}

func (o *xxx_AllocateNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AllocateNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaPool {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaPool != nil {
			if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpPartition {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Partition != nil {
			_ptr_lpPartition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Partition != nil {
					if err := o.Partition.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Partition, _ptr_lpPartition); err != nil {
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
	// lpMediaId {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	// dwTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// lpAllocateInformation {in, out} (1:{alias=LPNTMS_ALLOCATION_INFORMATION}*(1))(2:{alias=NTMS_ALLOCATION_INFORMATION}(struct))
	{
		if o.AllocateInformation != nil {
			if err := o.AllocateInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.AllocationInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AllocateNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaPool {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaPool == nil {
			o.MediaPool = &dtyp.GUID{}
		}
		if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpPartition {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpPartition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Partition == nil {
				o.Partition = &dtyp.GUID{}
			}
			if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpPartition := func(ptr interface{}) { o.Partition = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Partition, _s_lpPartition, _ptr_lpPartition); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpMediaId {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	// dwTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// lpAllocateInformation {in, out} (1:{alias=LPNTMS_ALLOCATION_INFORMATION,pointer=ref}*(1))(2:{alias=NTMS_ALLOCATION_INFORMATION}(struct))
	{
		if o.AllocateInformation == nil {
			o.AllocateInformation = &rsmp.AllocationInformation{}
		}
		if err := o.AllocateInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AllocateNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AllocateNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in, out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpAllocateInformation {in, out} (1:{alias=LPNTMS_ALLOCATION_INFORMATION}*(1))(2:{alias=NTMS_ALLOCATION_INFORMATION}(struct))
	{
		if o.AllocateInformation != nil {
			if err := o.AllocateInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rsmp.AllocationInformation{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_AllocateNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in, out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpAllocateInformation {in, out} (1:{alias=LPNTMS_ALLOCATION_INFORMATION,pointer=ref}*(1))(2:{alias=NTMS_ALLOCATION_INFORMATION}(struct))
	{
		if o.AllocateInformation == nil {
			o.AllocateInformation = &rsmp.AllocationInformation{}
		}
		if err := o.AllocateInformation.UnmarshalNDR(ctx, w); err != nil {
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

// AllocateNTMSMediaRequest structure represents the AllocateNtmsMedia operation request
type AllocateNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaPool: A pointer to the identifier of the media pool from which the media is
	// allocated.
	MediaPool *dtyp.GUID `idl:"name:lpMediaPool" json:"media_pool"`
	// lpPartition: A pointer to the partition identifier of the side that MUST be used
	// for a logical media identifier. This feature MUST be used to allocate a particular
	// side or to import media.
	Partition *dtyp.GUID `idl:"name:lpPartition;pointer:unique" json:"partition"`
	// lpMediaId: A pointer to the identifier of the allocated medium.
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
	// dwOptions: A bitmap of allocation options from the NtmsAllocateOptions (section 2.2.3.1)
	// enumeration.
	Options uint32 `idl:"name:dwOptions" json:"options"`
	// dwTimeout: The maximum time, in milliseconds, allowed to allocate the specified media.
	// If this parameter is -1, the function MUST NOT time out. If this parameter is 0,
	// it MUST NOT wait for media.
	Timeout uint32 `idl:"name:dwTimeout" json:"timeout"`
	// lpAllocateInformation: A pointer to an NTMS_ALLOCATION_INFORMATION (section 2.2.3.6)
	// structure that MUST be filled with the source media pool from which the medium was
	// taken. A NULL pointer MUST be passed if this information is not needed.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED        | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                       | but indicate a security subsystem error.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY    | An allocation error occurred during processing.                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER    | The media or media pool identifiers are missing.                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704C7 ERROR_CANCELLED            | The operator canceled the request for new media.                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800705B4 ERROR_TIMEOUT              | The time-out event expired before the medium was available.                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA        | The partition identifier or logical media identifier was invalid when combined   |
	//	|                                       | with the NTMS_ALLOCATE_NEXT flag.                                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL   | The media pool identifier is invalid.                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D0 ERROR_MEDIA_OFFLINE        | The specified media are offline and cannot be allocated.                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D4 ERROR_MEDIA_UNAVAILABLE    | No media have been allocated in the specified time-out.                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE     | The database is inaccessible or damaged.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL        | The database is full.                                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DF ERROR_DEVICE_NOT_AVAILABLE | An intermediate resource is not available.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	AllocateInformation *rsmp.AllocationInformation `idl:"name:lpAllocateInformation" json:"allocate_information"`
}

func (o *AllocateNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_AllocateNTMSMediaOperation) *xxx_AllocateNTMSMediaOperation {
	if op == nil {
		op = &xxx_AllocateNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaPool = o.MediaPool
	op.Partition = o.Partition
	op.MediaID = o.MediaID
	op.Options = o.Options
	op.Timeout = o.Timeout
	op.AllocateInformation = o.AllocateInformation
	return op
}

func (o *AllocateNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_AllocateNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaPool = op.MediaPool
	o.Partition = op.Partition
	o.MediaID = op.MediaID
	o.Options = op.Options
	o.Timeout = op.Timeout
	o.AllocateInformation = op.AllocateInformation
}
func (o *AllocateNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AllocateNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AllocateNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AllocateNTMSMediaResponse structure represents the AllocateNtmsMedia operation response
type AllocateNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpMediaId: A pointer to the identifier of the allocated medium.
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
	// lpAllocateInformation: A pointer to an NTMS_ALLOCATION_INFORMATION (section 2.2.3.6)
	// structure that MUST be filled with the source media pool from which the medium was
	// taken. A NULL pointer MUST be passed if this information is not needed.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED        | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                       | but indicate a security subsystem error.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY    | An allocation error occurred during processing.                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER    | The media or media pool identifiers are missing.                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704C7 ERROR_CANCELLED            | The operator canceled the request for new media.                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800705B4 ERROR_TIMEOUT              | The time-out event expired before the medium was available.                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA        | The partition identifier or logical media identifier was invalid when combined   |
	//	|                                       | with the NTMS_ALLOCATE_NEXT flag.                                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL   | The media pool identifier is invalid.                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D0 ERROR_MEDIA_OFFLINE        | The specified media are offline and cannot be allocated.                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D4 ERROR_MEDIA_UNAVAILABLE    | No media have been allocated in the specified time-out.                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE     | The database is inaccessible or damaged.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL        | The database is full.                                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DF ERROR_DEVICE_NOT_AVAILABLE | An intermediate resource is not available.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	AllocateInformation *rsmp.AllocationInformation `idl:"name:lpAllocateInformation" json:"allocate_information"`
	// Return: The AllocateNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AllocateNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_AllocateNTMSMediaOperation) *xxx_AllocateNTMSMediaOperation {
	if op == nil {
		op = &xxx_AllocateNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MediaID = o.MediaID
	op.AllocateInformation = o.AllocateInformation
	op.Return = o.Return
	return op
}

func (o *AllocateNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_AllocateNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MediaID = op.MediaID
	o.AllocateInformation = op.AllocateInformation
	o.Return = op.Return
}
func (o *AllocateNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AllocateNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AllocateNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeallocateNTMSMediaOperation structure represents the DeallocateNtmsMedia operation
type xxx_DeallocateNTMSMediaOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID *dtyp.GUID     `idl:"name:lpMediaId" json:"media_id"`
	Options uint32         `idl:"name:dwOptions" json:"options"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeallocateNTMSMediaOperation) OpNum() int { return 4 }

func (o *xxx_DeallocateNTMSMediaOperation) OpName() string {
	return "/INtmsMediaServices1/v0/DeallocateNtmsMedia"
}

func (o *xxx_DeallocateNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeallocateNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeallocateNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeallocateNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeallocateNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeallocateNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeallocateNTMSMediaRequest structure represents the DeallocateNtmsMedia operation request
type DeallocateNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: A pointer to the identifier of the logical media.
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
	// dwOptions: This parameter is unused. It MUST be sent as 0 and MUST be ignored on
	// receipt.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                    | but indicate a security subsystem error.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation error occurred during processing.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The media or media pool identifiers are missing.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database is inaccessible or damaged.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full.                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
}

func (o *DeallocateNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_DeallocateNTMSMediaOperation) *xxx_DeallocateNTMSMediaOperation {
	if op == nil {
		op = &xxx_DeallocateNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	op.Options = o.Options
	return op
}

func (o *DeallocateNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_DeallocateNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
	o.Options = op.Options
}
func (o *DeallocateNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeallocateNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeallocateNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeallocateNTMSMediaResponse structure represents the DeallocateNtmsMedia operation response
type DeallocateNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeallocateNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeallocateNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_DeallocateNTMSMediaOperation) *xxx_DeallocateNTMSMediaOperation {
	if op == nil {
		op = &xxx_DeallocateNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeallocateNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_DeallocateNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeallocateNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeallocateNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeallocateNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SwapNTMSMediaOperation structure represents the SwapNtmsMedia operation
type xxx_SwapNTMSMediaOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID1 *dtyp.GUID     `idl:"name:lpMediaId1" json:"media_id1"`
	MediaID2 *dtyp.GUID     `idl:"name:lpMediaId2" json:"media_id2"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SwapNTMSMediaOperation) OpNum() int { return 5 }

func (o *xxx_SwapNTMSMediaOperation) OpName() string { return "/INtmsMediaServices1/v0/SwapNtmsMedia" }

func (o *xxx_SwapNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SwapNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId1 {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID1 != nil {
			if err := o.MediaID1.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpMediaId2 {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID2 != nil {
			if err := o.MediaID2.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SwapNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId1 {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID1 == nil {
			o.MediaID1 = &dtyp.GUID{}
		}
		if err := o.MediaID1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpMediaId2 {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID2 == nil {
			o.MediaID2 = &dtyp.GUID{}
		}
		if err := o.MediaID2.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SwapNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SwapNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SwapNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SwapNTMSMediaRequest structure represents the SwapNtmsMedia operation request
type SwapNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId1: A pointer to the identifier of a logical medium.
	MediaID1 *dtyp.GUID `idl:"name:lpMediaId1" json:"media_id1"`
	// lpMediaId2:  A pointer to the identifier of a logical medium.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                     | but indicate a security subsystem error.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation error occurred during processing.                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007000B ERROR_BAD_FORMAT         | No media label library recognizes the media label.                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | At least one media identifier is missing.                                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | A media identifier is invalid.                                                   |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | A media pool of the logical media is invalid.                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL      | The database is full.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	MediaID2 *dtyp.GUID `idl:"name:lpMediaId2" json:"media_id2"`
}

func (o *SwapNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_SwapNTMSMediaOperation) *xxx_SwapNTMSMediaOperation {
	if op == nil {
		op = &xxx_SwapNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID1 = o.MediaID1
	op.MediaID2 = o.MediaID2
	return op
}

func (o *SwapNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_SwapNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID1 = op.MediaID1
	o.MediaID2 = op.MediaID2
}
func (o *SwapNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SwapNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SwapNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SwapNTMSMediaResponse structure represents the SwapNtmsMedia operation response
type SwapNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SwapNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SwapNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_SwapNTMSMediaOperation) *xxx_SwapNTMSMediaOperation {
	if op == nil {
		op = &xxx_SwapNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SwapNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_SwapNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SwapNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SwapNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SwapNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DecommissionNTMSMediaOperation structure represents the DecommissionNtmsMedia operation
type xxx_DecommissionNTMSMediaOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID *dtyp.GUID     `idl:"name:lpMediaId" json:"media_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DecommissionNTMSMediaOperation) OpNum() int { return 6 }

func (o *xxx_DecommissionNTMSMediaOperation) OpName() string {
	return "/INtmsMediaServices1/v0/DecommissionNtmsMedia"
}

func (o *xxx_DecommissionNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecommissionNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_DecommissionNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecommissionNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecommissionNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DecommissionNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DecommissionNTMSMediaRequest structure represents the DecommissionNtmsMedia operation request
type DecommissionNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: A pointer to the medium identifier of the partition to be decommissioned.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | Access to the object is denied; other security errors are possible, but indicate |
	//	|                                     | a security subsystem error.                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation error occurred during processing.                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | The parameter is not valid.                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | The media identifier is not valid.                                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | The media pool identifier is not valid.                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database query or update failed.                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE      | An unexpected state was encountered; might be disabled or offline.               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
}

func (o *DecommissionNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_DecommissionNTMSMediaOperation) *xxx_DecommissionNTMSMediaOperation {
	if op == nil {
		op = &xxx_DecommissionNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	return op
}

func (o *DecommissionNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_DecommissionNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
}
func (o *DecommissionNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DecommissionNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DecommissionNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DecommissionNTMSMediaResponse structure represents the DecommissionNtmsMedia operation response
type DecommissionNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DecommissionNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DecommissionNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_DecommissionNTMSMediaOperation) *xxx_DecommissionNTMSMediaOperation {
	if op == nil {
		op = &xxx_DecommissionNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DecommissionNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_DecommissionNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DecommissionNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DecommissionNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DecommissionNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNTMSMediaCompleteOperation structure represents the SetNtmsMediaComplete operation
type xxx_SetNTMSMediaCompleteOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID *dtyp.GUID     `idl:"name:lpMediaId" json:"media_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNTMSMediaCompleteOperation) OpNum() int { return 7 }

func (o *xxx_SetNTMSMediaCompleteOperation) OpName() string {
	return "/INtmsMediaServices1/v0/SetNtmsMediaComplete"
}

func (o *xxx_SetNTMSMediaCompleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSMediaCompleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetNTMSMediaCompleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSMediaCompleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNTMSMediaCompleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNTMSMediaCompleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNTMSMediaCompleteRequest structure represents the SetNtmsMediaComplete operation request
type SetNTMSMediaCompleteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: A pointer to the identifier of the logical medium.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                     | but indicate a security subsystem error.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation error occurred during processing.                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | The media identifier is missing.                                                 |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | The media identifier is invalid.                                                 |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | The media pool of the media is invalid.                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL      | The database is full.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE      | The medium is not in the allocated state, or is currently mounted.               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
}

func (o *SetNTMSMediaCompleteRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSMediaCompleteOperation) *xxx_SetNTMSMediaCompleteOperation {
	if op == nil {
		op = &xxx_SetNTMSMediaCompleteOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	return op
}

func (o *SetNTMSMediaCompleteRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSMediaCompleteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
}
func (o *SetNTMSMediaCompleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNTMSMediaCompleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSMediaCompleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNTMSMediaCompleteResponse structure represents the SetNtmsMediaComplete operation response
type SetNTMSMediaCompleteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetNtmsMediaComplete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNTMSMediaCompleteResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNTMSMediaCompleteOperation) *xxx_SetNTMSMediaCompleteOperation {
	if op == nil {
		op = &xxx_SetNTMSMediaCompleteOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNTMSMediaCompleteResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNTMSMediaCompleteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNTMSMediaCompleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNTMSMediaCompleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNTMSMediaCompleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteNTMSMediaOperation structure represents the DeleteNtmsMedia operation
type xxx_DeleteNTMSMediaOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID *dtyp.GUID     `idl:"name:lpMediaId" json:"media_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteNTMSMediaOperation) OpNum() int { return 8 }

func (o *xxx_DeleteNTMSMediaOperation) OpName() string {
	return "/INtmsMediaServices1/v0/DeleteNtmsMedia"
}

func (o *xxx_DeleteNTMSMediaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_DeleteNTMSMediaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteNTMSMediaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteNTMSMediaRequest structure represents the DeleteNtmsMedia operation request
type DeleteNTMSMediaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: A pointer to the identifier of a physical medium.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible  |
	//	|                                     | but indicate a security subsystem error.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation error occurred during processing.                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | The media identifier is missing.                                                 |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | The media identifier is invalid.                                                 |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | The media pool of the media is invalid.                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL      | The database is full.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE      | The media is not offline.                                                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
}

func (o *DeleteNTMSMediaRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSMediaOperation) *xxx_DeleteNTMSMediaOperation {
	if op == nil {
		op = &xxx_DeleteNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	return op
}

func (o *DeleteNTMSMediaRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
}
func (o *DeleteNTMSMediaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteNTMSMediaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSMediaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteNTMSMediaResponse structure represents the DeleteNtmsMedia operation response
type DeleteNTMSMediaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteNtmsMedia return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteNTMSMediaResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSMediaOperation) *xxx_DeleteNTMSMediaOperation {
	if op == nil {
		op = &xxx_DeleteNTMSMediaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteNTMSMediaResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSMediaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteNTMSMediaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteNTMSMediaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSMediaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNTMSMediaPoolAOperation structure represents the CreateNtmsMediaPoolA operation
type xxx_CreateNTMSMediaPoolAOperation struct {
	This               *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat               `idl:"name:That" json:"that"`
	PoolName           string                       `idl:"name:lpPoolName;string" json:"pool_name"`
	MediaType          *dtyp.GUID                   `idl:"name:lpMediaType;pointer:unique" json:"media_type"`
	Options            uint32                       `idl:"name:dwOptions" json:"options"`
	SecurityAttributes *rsmp.SecurityAttributesNTMS `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
	PoolID             *dtyp.GUID                   `idl:"name:lpPoolId" json:"pool_id"`
	Return             int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNTMSMediaPoolAOperation) OpNum() int { return 9 }

func (o *xxx_CreateNTMSMediaPoolAOperation) OpName() string {
	return "/INtmsMediaServices1/v0/CreateNtmsMediaPoolA"
}

func (o *xxx_CreateNTMSMediaPoolAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaPoolAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.PoolName); err != nil {
			return err
		}
	}
	// lpMediaType {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaType != nil {
			_ptr_lpMediaType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MediaType != nil {
					if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MediaType, _ptr_lpMediaType); err != nil {
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
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=LPSECURITY_ATTRIBUTES_NTMS}*(1))(2:{alias=SECURITY_ATTRIBUTES_NTMS}(struct))
	{
		if o.SecurityAttributes != nil {
			_ptr_lpSecurityAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecurityAttributes != nil {
					if err := o.SecurityAttributes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsmp.SecurityAttributesNTMS{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
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

func (o *xxx_CreateNTMSMediaPoolAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.PoolName); err != nil {
			return err
		}
	}
	// lpMediaType {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpMediaType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MediaType == nil {
				o.MediaType = &dtyp.GUID{}
			}
			if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpMediaType := func(ptr interface{}) { o.MediaType = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.MediaType, _s_lpMediaType, _ptr_lpMediaType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=LPSECURITY_ATTRIBUTES_NTMS}*(1))(2:{alias=SECURITY_ATTRIBUTES_NTMS}(struct))
	{
		_ptr_lpSecurityAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecurityAttributes == nil {
				o.SecurityAttributes = &rsmp.SecurityAttributesNTMS{}
			}
			if err := o.SecurityAttributes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpSecurityAttributes := func(ptr interface{}) { o.SecurityAttributes = *ptr.(**rsmp.SecurityAttributesNTMS) }
		if err := w.ReadPointer(&o.SecurityAttributes, _s_lpSecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaPoolAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaPoolAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolId {out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CreateNTMSMediaPoolAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolId {out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
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

// CreateNTMSMediaPoolARequest structure represents the CreateNtmsMediaPoolA operation request
type CreateNTMSMediaPoolARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpPoolName: A null-terminated sequence of ASCII characters that constitute the name
	// of the new media pool; MUST be unique among all the media pool present in the server.
	PoolName string `idl:"name:lpPoolName;string" json:"pool_name"`
	// lpMediaType: Pointer to a unique identifier for the type of media in this media pool.
	// INtmsObjectManagement1::EnumerateNtmsObject produces a list of available media types
	// and their attributes. Use of a NULL pointer creates a media pool that contains only
	// other media pools.
	MediaType *dtyp.GUID `idl:"name:lpMediaType;pointer:unique" json:"media_type"`
	// dwOptions: A value from the NtmsCreateOptions (section 2.2.3.2) enumeration that
	// specifies the type of creation to undertake.
	Options uint32 `idl:"name:dwOptions" json:"options"`
	// lpSecurityAttributes: A pointer to an optional SECURITY_ATTRIBUTES_NTMS structure
	// that is used to restrict access to the pool.
	SecurityAttributes *rsmp.SecurityAttributesNTMS `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
}

func (o *CreateNTMSMediaPoolARequest) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolAOperation) *xxx_CreateNTMSMediaPoolAOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaPoolAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PoolName = o.PoolName
	op.MediaType = o.MediaType
	op.Options = o.Options
	op.SecurityAttributes = o.SecurityAttributes
	return op
}

func (o *CreateNTMSMediaPoolARequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PoolName = op.PoolName
	o.MediaType = op.MediaType
	o.Options = op.Options
	o.SecurityAttributes = op.SecurityAttributes
}
func (o *CreateNTMSMediaPoolARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateNTMSMediaPoolARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaPoolAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNTMSMediaPoolAResponse structure represents the CreateNtmsMediaPoolA operation response
type CreateNTMSMediaPoolAResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpPoolId: A pointer to the identifier of the new media pool.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible  |
	//	|                                    | but indicate a security subsystem error.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The media pool name or identifier is missing.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007B ERROR_INVALID_NAME      | The pool name syntax is invalid or the name is too long.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700B7 ERROR_ALREADY_EXISTS    | A new media pool could not be created because one already exists with this name. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA     | The selected media type is not valid.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | Unable to open existing media pool.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database is inaccessible or damaged.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full; other security errors are possible but indicate a security |
	//	|                                    | subsystem error.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	PoolID *dtyp.GUID `idl:"name:lpPoolId" json:"pool_id"`
	// Return: The CreateNtmsMediaPoolA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateNTMSMediaPoolAResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolAOperation) *xxx_CreateNTMSMediaPoolAOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaPoolAOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PoolID = o.PoolID
	op.Return = o.Return
	return op
}

func (o *CreateNTMSMediaPoolAResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolAOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PoolID = op.PoolID
	o.Return = op.Return
}
func (o *CreateNTMSMediaPoolAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateNTMSMediaPoolAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaPoolAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNTMSMediaPoolWOperation structure represents the CreateNtmsMediaPoolW operation
type xxx_CreateNTMSMediaPoolWOperation struct {
	This               *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat               `idl:"name:That" json:"that"`
	PoolName           string                       `idl:"name:lpPoolName;string" json:"pool_name"`
	MediaType          *dtyp.GUID                   `idl:"name:lpMediaType;pointer:unique" json:"media_type"`
	Options            uint32                       `idl:"name:dwOptions" json:"options"`
	SecurityAttributes *rsmp.SecurityAttributesNTMS `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
	PoolID             *dtyp.GUID                   `idl:"name:lpPoolId" json:"pool_id"`
	Return             int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNTMSMediaPoolWOperation) OpNum() int { return 10 }

func (o *xxx_CreateNTMSMediaPoolWOperation) OpName() string {
	return "/INtmsMediaServices1/v0/CreateNtmsMediaPoolW"
}

func (o *xxx_CreateNTMSMediaPoolWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaPoolWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PoolName); err != nil {
			return err
		}
	}
	// lpMediaType {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaType != nil {
			_ptr_lpMediaType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MediaType != nil {
					if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MediaType, _ptr_lpMediaType); err != nil {
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
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=LPSECURITY_ATTRIBUTES_NTMS}*(1))(2:{alias=SECURITY_ATTRIBUTES_NTMS}(struct))
	{
		if o.SecurityAttributes != nil {
			_ptr_lpSecurityAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecurityAttributes != nil {
					if err := o.SecurityAttributes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsmp.SecurityAttributesNTMS{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
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

func (o *xxx_CreateNTMSMediaPoolWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PoolName); err != nil {
			return err
		}
	}
	// lpMediaType {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpMediaType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MediaType == nil {
				o.MediaType = &dtyp.GUID{}
			}
			if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpMediaType := func(ptr interface{}) { o.MediaType = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.MediaType, _s_lpMediaType, _ptr_lpMediaType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=LPSECURITY_ATTRIBUTES_NTMS}*(1))(2:{alias=SECURITY_ATTRIBUTES_NTMS}(struct))
	{
		_ptr_lpSecurityAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecurityAttributes == nil {
				o.SecurityAttributes = &rsmp.SecurityAttributesNTMS{}
			}
			if err := o.SecurityAttributes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpSecurityAttributes := func(ptr interface{}) { o.SecurityAttributes = *ptr.(**rsmp.SecurityAttributesNTMS) }
		if err := w.ReadPointer(&o.SecurityAttributes, _s_lpSecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaPoolWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNTMSMediaPoolWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolId {out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CreateNTMSMediaPoolWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolId {out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
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

// CreateNTMSMediaPoolWRequest structure represents the CreateNtmsMediaPoolW operation request
type CreateNTMSMediaPoolWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpPoolName: A null-terminated sequence of Unicode characters, encoded using UTF-16,
	// that constitute the name of the new media pool; MUST be unique among all the media
	// pools present in the server.
	PoolName string `idl:"name:lpPoolName;string" json:"pool_name"`
	// lpMediaType: Pointer to a unique identifier for the type of media in this media pool.
	// INtmsObjectManagement1::EnumerateNtmsObject produces a list of available media types
	// and their attributes. Use of a NULL pointer creates a media pool that contains only
	// other media pools.
	MediaType *dtyp.GUID `idl:"name:lpMediaType;pointer:unique" json:"media_type"`
	// dwOptions: A value from the NtmsCreateOptions (section 2.2.3.2) enumeration specifying
	// the type of creation to undertake.
	Options uint32 `idl:"name:dwOptions" json:"options"`
	// lpSecurityAttributes: A pointer to an optional SECURITY_ATTRIBUTES_NTMS (section
	// 2.2.3.5) structure used to restrict access to the pool.
	SecurityAttributes *rsmp.SecurityAttributesNTMS `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
}

func (o *CreateNTMSMediaPoolWRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolWOperation) *xxx_CreateNTMSMediaPoolWOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaPoolWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PoolName = o.PoolName
	op.MediaType = o.MediaType
	op.Options = o.Options
	op.SecurityAttributes = o.SecurityAttributes
	return op
}

func (o *CreateNTMSMediaPoolWRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PoolName = op.PoolName
	o.MediaType = op.MediaType
	o.Options = op.Options
	o.SecurityAttributes = op.SecurityAttributes
}
func (o *CreateNTMSMediaPoolWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateNTMSMediaPoolWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaPoolWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNTMSMediaPoolWResponse structure represents the CreateNtmsMediaPoolW operation response
type CreateNTMSMediaPoolWResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpPoolId: A pointer to the identifier of the new media pool.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                    | but indicate a security subsystem error.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The media pool name or identifier is missing.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007B ERROR_INVALID_NAME      | The pool name syntax is invalid or the name is too long.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700B7 ERROR_ALREADY_EXISTS    | A new media pool could not be created because one already exists with this name. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA     | The selected media type is not valid.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | Unable to open an existing media pool.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database is inaccessible or damaged.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full; other security errors are possible, but indicate a         |
	//	|                                    | security subsystem error.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	PoolID *dtyp.GUID `idl:"name:lpPoolId" json:"pool_id"`
	// Return: The CreateNtmsMediaPoolW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateNTMSMediaPoolWResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolWOperation) *xxx_CreateNTMSMediaPoolWOperation {
	if op == nil {
		op = &xxx_CreateNTMSMediaPoolWOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PoolID = o.PoolID
	op.Return = o.Return
	return op
}

func (o *CreateNTMSMediaPoolWResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNTMSMediaPoolWOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PoolID = op.PoolID
	o.Return = op.Return
}
func (o *CreateNTMSMediaPoolWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateNTMSMediaPoolWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNTMSMediaPoolWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSMediaPoolNameAOperation structure represents the GetNtmsMediaPoolNameA operation
type xxx_GetNTMSMediaPoolNameAOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	PoolID         *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	BufferName     []byte         `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSizeBuf)" json:"buffer_name"`
	NameSizeBuffer uint32         `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
	NameSize       uint32         `idl:"name:lpdwNameSize" json:"name_size"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSMediaPoolNameAOperation) OpNum() int { return 11 }

func (o *xxx_GetNTMSMediaPoolNameAOperation) OpName() string {
	return "/INtmsMediaServices1/v0/GetNtmsMediaPoolNameA"
}

func (o *xxx_GetNTMSMediaPoolNameAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSizeBuf](uchar))
	{
		dimSize1 := uint64(o.NameSizeBuffer)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.NameSizeBuffer)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.BufferName {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.BufferName[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.BufferName); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
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

func (o *xxx_GetNTMSMediaPoolNameAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSizeBuf](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.BufferName", sizeInfo[0])
		}
		o.BufferName = make([]byte, sizeInfo[0])
		for i1 := range o.BufferName {
			i1 := i1
			if err := w.ReadData(&o.BufferName[i1]); err != nil {
				return err
			}
		}
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
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

// GetNTMSMediaPoolNameARequest structure represents the GetNtmsMediaPoolNameA operation request
type GetNTMSMediaPoolNameARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpPoolId: A pointer to the identifier of the media pool from which to retrieve the
	// name.
	PoolID *dtyp.GUID `idl:"name:lpPoolId" json:"pool_id"`
	// lpdwNameSizeBuf:  A pointer to the size, in bytes, of lpBufName.
	NameSizeBuffer uint32 `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
}

func (o *GetNTMSMediaPoolNameARequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameAOperation) *xxx_GetNTMSMediaPoolNameAOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PoolID = o.PoolID
	op.NameSizeBuffer = o.NameSizeBuffer
	return op
}

func (o *GetNTMSMediaPoolNameARequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PoolID = op.PoolID
	o.NameSizeBuffer = op.NameSizeBuffer
}
func (o *GetNTMSMediaPoolNameARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSMediaPoolNameARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSMediaPoolNameAResponse structure represents the GetNtmsMediaPoolNameA operation response
type GetNTMSMediaPoolNameAResponse struct {
	// XXX: lpdwNameSizeBuf is an implicit input depedency for output parameters
	NameSizeBuffer uint32 `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpBufName: A null-terminated buffer that contains the name of the media pool.
	BufferName []byte `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSizeBuf)" json:"buffer_name"`
	// lpdwNameSize:  A pointer to the length of the string in lpBufName, including the
	// terminating null character.
	//
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                |
	//	|              VALUE/CODE              |                                  DESCRIPTION                                   |
	//	|                                      |                                                                                |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                       |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access to an object was denied.                                                |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY   | An allocation error occurred during processing.                                |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER   | A parameter is missing or invalid.                                             |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The buffer is not large enough. The required size is returned in lpdwNameSize. |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL  | The media pool identifier is missing or invalid.                               |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	NameSize uint32 `idl:"name:lpdwNameSize" json:"name_size"`
	// Return: The GetNtmsMediaPoolNameA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSMediaPoolNameAResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameAOperation) *xxx_GetNTMSMediaPoolNameAOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameAOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NameSizeBuffer == uint32(0) {
		op.NameSizeBuffer = o.NameSizeBuffer
	}

	op.That = o.That
	op.BufferName = o.BufferName
	op.NameSize = o.NameSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSMediaPoolNameAResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameAOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NameSizeBuffer = op.NameSizeBuffer

	o.That = op.That
	o.BufferName = op.BufferName
	o.NameSize = op.NameSize
	o.Return = op.Return
}
func (o *GetNTMSMediaPoolNameAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSMediaPoolNameAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSMediaPoolNameWOperation structure represents the GetNtmsMediaPoolNameW operation
type xxx_GetNTMSMediaPoolNameWOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	PoolID         *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	BufferName     string         `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSizeBuf)" json:"buffer_name"`
	NameSizeBuffer uint32         `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
	NameSize       uint32         `idl:"name:lpdwNameSize" json:"name_size"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSMediaPoolNameWOperation) OpNum() int { return 12 }

func (o *xxx_GetNTMSMediaPoolNameWOperation) OpName() string {
	return "/INtmsMediaServices1/v0/GetNtmsMediaPoolNameW"
}

func (o *xxx_GetNTMSMediaPoolNameWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSizeBuf,string](wchar))
	{
		dimSize1 := uint64(o.NameSizeBuffer)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.NameSizeBuffer)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_BufferName_buf := utf16.Encode([]rune(o.BufferName))
		if uint64(len(_BufferName_buf)) > sizeInfo[0] {
			_BufferName_buf = _BufferName_buf[:sizeInfo[0]]
		}
		for i1 := range _BufferName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_BufferName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_BufferName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
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

func (o *xxx_GetNTMSMediaPoolNameWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSizeBuf,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _BufferName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BufferName_buf", sizeInfo[0])
		}
		_BufferName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _BufferName_buf {
			i1 := i1
			if err := w.ReadData(&_BufferName_buf[i1]); err != nil {
				return err
			}
		}
		o.BufferName = strings.TrimRight(string(utf16.Decode(_BufferName_buf)), ndr.ZeroString)
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
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

// GetNTMSMediaPoolNameWRequest structure represents the GetNtmsMediaPoolNameW operation request
type GetNTMSMediaPoolNameWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpPoolId: A pointer to the identifier of the media pool for which to retrieve the
	// name.
	PoolID *dtyp.GUID `idl:"name:lpPoolId" json:"pool_id"`
	// lpdwNameSizeBuf: A pointer to the size, in bytes, of the client buffer that is allocated
	// to store lpBufName.
	NameSizeBuffer uint32 `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
}

func (o *GetNTMSMediaPoolNameWRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWOperation) *xxx_GetNTMSMediaPoolNameWOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PoolID = o.PoolID
	op.NameSizeBuffer = o.NameSizeBuffer
	return op
}

func (o *GetNTMSMediaPoolNameWRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PoolID = op.PoolID
	o.NameSizeBuffer = op.NameSizeBuffer
}
func (o *GetNTMSMediaPoolNameWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSMediaPoolNameWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSMediaPoolNameWResponse structure represents the GetNtmsMediaPoolNameW operation response
type GetNTMSMediaPoolNameWResponse struct {
	// XXX: lpdwNameSizeBuf is an implicit input depedency for output parameters
	NameSizeBuffer uint32 `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpBufName: A null-terminated buffer that contains the name of the media pool.
	BufferName string `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSizeBuf)" json:"buffer_name"`
	// lpdwNameSize: A pointer to the length of the string in lpBufName, including the terminating
	// null character.
	//
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                |
	//	|              VALUE/CODE              |                                  DESCRIPTION                                   |
	//	|                                      |                                                                                |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                       |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access to an object was denied.                                                |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY   | An allocation error occurred during processing.                                |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER   | A parameter is missing or invalid.                                             |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The buffer is not large enough. The required size is returned in lpdwNameSize. |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL  | The media pool identifier is missing or invalid.                               |
	//	+--------------------------------------+--------------------------------------------------------------------------------+
	NameSize uint32 `idl:"name:lpdwNameSize" json:"name_size"`
	// Return: The GetNtmsMediaPoolNameW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSMediaPoolNameWResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWOperation) *xxx_GetNTMSMediaPoolNameWOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameWOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NameSizeBuffer == uint32(0) {
		op.NameSizeBuffer = o.NameSizeBuffer
	}

	op.That = o.That
	op.BufferName = o.BufferName
	op.NameSize = o.NameSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSMediaPoolNameWResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NameSizeBuffer = op.NameSizeBuffer

	o.That = op.That
	o.BufferName = op.BufferName
	o.NameSize = op.NameSize
	o.Return = op.Return
}
func (o *GetNTMSMediaPoolNameWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSMediaPoolNameWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveToNTMSMediaPoolOperation structure represents the MoveToNtmsMediaPool operation
type xxx_MoveToNTMSMediaPoolOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID *dtyp.GUID     `idl:"name:lpMediaId" json:"media_id"`
	PoolID  *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveToNTMSMediaPoolOperation) OpNum() int { return 13 }

func (o *xxx_MoveToNTMSMediaPoolOperation) OpName() string {
	return "/INtmsMediaServices1/v0/MoveToNtmsMediaPool"
}

func (o *xxx_MoveToNTMSMediaPoolOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveToNTMSMediaPoolOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpPoolId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MoveToNTMSMediaPoolOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpPoolId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveToNTMSMediaPoolOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveToNTMSMediaPoolOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MoveToNTMSMediaPoolOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MoveToNTMSMediaPoolRequest structure represents the MoveToNtmsMediaPool operation request
type MoveToNTMSMediaPoolRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: A pointer to the identifier of a physical medium.
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
	// lpPoolId:  A pointer to the identifier of a media pool to which the medium is to
	// be moved.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                     | but indicate a security subsystem error.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation failure occurred during processing.                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | The parameter is missing or invalid.                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700AA ERROR_BUSY               | A side of the media is in use or currently unavailable.                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | The source media or implied source media pool is invalid.                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | Either the destination media pool is invalid, or media in the unrecognized or    |
	//	|                                     | import pool can be moved only to the free pool.                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL      | The database is full.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DB ERROR_MEDIA_INCOMPATIBLE | The source media type differs from the media type of the destination pool.       |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	PoolID *dtyp.GUID `idl:"name:lpPoolId" json:"pool_id"`
}

func (o *MoveToNTMSMediaPoolRequest) xxx_ToOp(ctx context.Context, op *xxx_MoveToNTMSMediaPoolOperation) *xxx_MoveToNTMSMediaPoolOperation {
	if op == nil {
		op = &xxx_MoveToNTMSMediaPoolOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	op.PoolID = o.PoolID
	return op
}

func (o *MoveToNTMSMediaPoolRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveToNTMSMediaPoolOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
	o.PoolID = op.PoolID
}
func (o *MoveToNTMSMediaPoolRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MoveToNTMSMediaPoolRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveToNTMSMediaPoolOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveToNTMSMediaPoolResponse structure represents the MoveToNtmsMediaPool operation response
type MoveToNTMSMediaPoolResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MoveToNtmsMediaPool return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MoveToNTMSMediaPoolResponse) xxx_ToOp(ctx context.Context, op *xxx_MoveToNTMSMediaPoolOperation) *xxx_MoveToNTMSMediaPoolOperation {
	if op == nil {
		op = &xxx_MoveToNTMSMediaPoolOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *MoveToNTMSMediaPoolResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveToNTMSMediaPoolOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *MoveToNTMSMediaPoolResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MoveToNTMSMediaPoolResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveToNTMSMediaPoolOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteNTMSMediaPoolOperation structure represents the DeleteNtmsMediaPool operation
type xxx_DeleteNTMSMediaPoolOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	PoolID *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteNTMSMediaPoolOperation) OpNum() int { return 14 }

func (o *xxx_DeleteNTMSMediaPoolOperation) OpName() string {
	return "/INtmsMediaServices1/v0/DeleteNtmsMediaPool"
}

func (o *xxx_DeleteNTMSMediaPoolOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaPoolOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_DeleteNTMSMediaPoolOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaPoolOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaPoolOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteNTMSMediaPoolOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteNTMSMediaPoolRequest structure represents the DeleteNtmsMediaPool operation request
type DeleteNTMSMediaPoolRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpPoolId: A pointer to the identifier of a media pool.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_CONTROL_ACCESS to the media pool is denied (for more information, see       |
	//	|                                     | [MSDN-SetNtmsObjectSecurity]); other security errors are possible, but indicate  |
	//	|                                     | a security subsystem error.                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation failure occurred during processing.                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | The media pool identifier is missing.                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | Unable to open the media pool or delete the free, import, or unrecognized media  |
	//	|                                     | pools.                                                                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D3 ERROR_NOT_EMPTY          | The media pool must be empty to be deleted.                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// The DeleteNtmsMediaPool method deletes the specified application media pool. Only
	// empty media pools can be deleted with the DeleteNtmsMediaPool method. Free, unrecognized,
	// and import media pools are managed by RSM and cannot be deleted with DeleteNtmsMediaPool.
	PoolID *dtyp.GUID `idl:"name:lpPoolId" json:"pool_id"`
}

func (o *DeleteNTMSMediaPoolRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSMediaPoolOperation) *xxx_DeleteNTMSMediaPoolOperation {
	if op == nil {
		op = &xxx_DeleteNTMSMediaPoolOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PoolID = o.PoolID
	return op
}

func (o *DeleteNTMSMediaPoolRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSMediaPoolOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PoolID = op.PoolID
}
func (o *DeleteNTMSMediaPoolRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteNTMSMediaPoolRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSMediaPoolOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteNTMSMediaPoolResponse structure represents the DeleteNtmsMediaPool operation response
type DeleteNTMSMediaPoolResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteNtmsMediaPool return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteNTMSMediaPoolResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSMediaPoolOperation) *xxx_DeleteNTMSMediaPoolOperation {
	if op == nil {
		op = &xxx_DeleteNTMSMediaPoolOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteNTMSMediaPoolResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSMediaPoolOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteNTMSMediaPoolResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteNTMSMediaPoolResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSMediaPoolOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNTMSMediaTypeOperation structure represents the AddNtmsMediaType operation
type xxx_AddNTMSMediaTypeOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaTypeID *dtyp.GUID     `idl:"name:lpMediaTypeId" json:"media_type_id"`
	LibID       *dtyp.GUID     `idl:"name:lpLibId" json:"lib_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNTMSMediaTypeOperation) OpNum() int { return 15 }

func (o *xxx_AddNTMSMediaTypeOperation) OpName() string {
	return "/INtmsMediaServices1/v0/AddNtmsMediaType"
}

func (o *xxx_AddNTMSMediaTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNTMSMediaTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaTypeId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaTypeID != nil {
			if err := o.MediaTypeID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpLibId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibID != nil {
			if err := o.LibID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_AddNTMSMediaTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaTypeId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaTypeID == nil {
			o.MediaTypeID = &dtyp.GUID{}
		}
		if err := o.MediaTypeID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpLibId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibID == nil {
			o.LibID = &dtyp.GUID{}
		}
		if err := o.LibID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNTMSMediaTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNTMSMediaTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddNTMSMediaTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddNTMSMediaTypeRequest structure represents the AddNtmsMediaType operation request
type AddNTMSMediaTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaTypeId: A pointer to the identifier of a media type to add to the library.
	MediaTypeID *dtyp.GUID `idl:"name:lpMediaTypeId" json:"media_type_id"`
	// lpLibId:  A pointer to the identifier of the library to which the media type is
	// to be added.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible, |
	//	|                                    | but indicate a security subsystem error.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing.                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The media type or library identifiers are missing.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CD ERROR_INVALID_LIBRARY   | The library identifier is invalid.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database is inaccessible or damaged.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full.                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the specified media type is not in the library object's list of already supported
	// media types, the AddNtmsMediaType method adds the it to the specified library. If
	// the specified media type is already in the library object's list of already supported
	// media types, the AddNtmsMediaType method will not add the it to the specified library
	// also it will not return error because of this. AddNtmsMediaType then creates the
	// system media pools if they do not exist.
	//
	// If the specified media type is not in the library object's list of already supported
	// media types, the AddNtmsMediaType method adds it to the specified library. If the
	// specified media type is already in the library object's list of supported media types,
	// the specified media type is not added to the library object's list. In both instances,
	// AddNtmsMediaType creates the system media pools if they do not exist.
	LibID *dtyp.GUID `idl:"name:lpLibId" json:"lib_id"`
}

func (o *AddNTMSMediaTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_AddNTMSMediaTypeOperation) *xxx_AddNTMSMediaTypeOperation {
	if op == nil {
		op = &xxx_AddNTMSMediaTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaTypeID = o.MediaTypeID
	op.LibID = o.LibID
	return op
}

func (o *AddNTMSMediaTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNTMSMediaTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaTypeID = op.MediaTypeID
	o.LibID = op.LibID
}
func (o *AddNTMSMediaTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddNTMSMediaTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNTMSMediaTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNTMSMediaTypeResponse structure represents the AddNtmsMediaType operation response
type AddNTMSMediaTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddNtmsMediaType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddNTMSMediaTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_AddNTMSMediaTypeOperation) *xxx_AddNTMSMediaTypeOperation {
	if op == nil {
		op = &xxx_AddNTMSMediaTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AddNTMSMediaTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNTMSMediaTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddNTMSMediaTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddNTMSMediaTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNTMSMediaTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteNTMSMediaTypeOperation structure represents the DeleteNtmsMediaType operation
type xxx_DeleteNTMSMediaTypeOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaTypeID *dtyp.GUID     `idl:"name:lpMediaTypeId" json:"media_type_id"`
	LibID       *dtyp.GUID     `idl:"name:lpLibId" json:"lib_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteNTMSMediaTypeOperation) OpNum() int { return 16 }

func (o *xxx_DeleteNTMSMediaTypeOperation) OpName() string {
	return "/INtmsMediaServices1/v0/DeleteNtmsMediaType"
}

func (o *xxx_DeleteNTMSMediaTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaTypeId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaTypeID != nil {
			if err := o.MediaTypeID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpLibId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibID != nil {
			if err := o.LibID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_DeleteNTMSMediaTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaTypeId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaTypeID == nil {
			o.MediaTypeID = &dtyp.GUID{}
		}
		if err := o.MediaTypeID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpLibId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.LibID == nil {
			o.LibID = &dtyp.GUID{}
		}
		if err := o.LibID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteNTMSMediaTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteNTMSMediaTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteNTMSMediaTypeRequest structure represents the DeleteNtmsMediaType operation request
type DeleteNTMSMediaTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaTypeId: A pointer to the identifier of a media type to delete from the library.
	MediaTypeID *dtyp.GUID `idl:"name:lpMediaTypeId" json:"media_type_id"`
	// lpLibId:  A pointer to the identifier of the library from which to delete the media
	// type.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | NTMS_MODIFY_ACCESS to the library is denied; other security errors are possible  |
	//	|                                    | but indicate a security subsystem error.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | An allocation failure occurred during processing.                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | The media type or library identifiers are missing.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710CD ERROR_INVALID_LIBRARY   | The library identifier is invalid.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D3 ERROR_NOT_EMPTY         | The media pool must be empty to be deleted.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database is inaccessible or damaged.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL     | The database is full.                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	LibID *dtyp.GUID `idl:"name:lpLibId" json:"lib_id"`
}

func (o *DeleteNTMSMediaTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSMediaTypeOperation) *xxx_DeleteNTMSMediaTypeOperation {
	if op == nil {
		op = &xxx_DeleteNTMSMediaTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaTypeID = o.MediaTypeID
	op.LibID = o.LibID
	return op
}

func (o *DeleteNTMSMediaTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSMediaTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaTypeID = op.MediaTypeID
	o.LibID = op.LibID
}
func (o *DeleteNTMSMediaTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteNTMSMediaTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSMediaTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteNTMSMediaTypeResponse structure represents the DeleteNtmsMediaType operation response
type DeleteNTMSMediaTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteNtmsMediaType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteNTMSMediaTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteNTMSMediaTypeOperation) *xxx_DeleteNTMSMediaTypeOperation {
	if op == nil {
		op = &xxx_DeleteNTMSMediaTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteNTMSMediaTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteNTMSMediaTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteNTMSMediaTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteNTMSMediaTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteNTMSMediaTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ChangeNTMSMediaTypeOperation structure represents the ChangeNtmsMediaType operation
type xxx_ChangeNTMSMediaTypeOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MediaID *dtyp.GUID     `idl:"name:lpMediaId" json:"media_id"`
	PoolID  *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangeNTMSMediaTypeOperation) OpNum() int { return 17 }

func (o *xxx_ChangeNTMSMediaTypeOperation) OpName() string {
	return "/INtmsMediaServices1/v0/ChangeNtmsMediaType"
}

func (o *xxx_ChangeNTMSMediaTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNTMSMediaTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID != nil {
			if err := o.MediaID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpPoolId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ChangeNTMSMediaTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpMediaId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.MediaID == nil {
			o.MediaID = &dtyp.GUID{}
		}
		if err := o.MediaID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpPoolId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNTMSMediaTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeNTMSMediaTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ChangeNTMSMediaTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ChangeNTMSMediaTypeRequest structure represents the ChangeNtmsMediaType operation request
type ChangeNTMSMediaTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpMediaId: A pointer to the identifier of the physical media to be moved.
	MediaID *dtyp.GUID `idl:"name:lpMediaId" json:"media_id"`
	// lpPoolId:  A pointer to the identifier of the media pool to which the media will
	// be allocated.
	//
	//	+-------------------------------------+--------------------------------------------------------------+
	//	|               RETURN                |                                                              |
	//	|             VALUE/CODE              |                         DESCRIPTION                          |
	//	|                                     |                                                              |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                     |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | NTMS_MODIFY_ACCESS to the media pool of the media is denied. |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | An allocation failure occurred during processing.            |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER  | The media pool or media identifiers are missing.             |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x800710CC ERROR_INVALID_MEDIA      | The media identifier is not valid.                           |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x800710CE ERROR_INVALID_MEDIA_POOL | The identifier of the media pool is invalid.                 |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE   | The database is inaccessible or damaged.                     |
	//	+-------------------------------------+--------------------------------------------------------------+
	//	| 0x800710DA ERROR_DATABASE_FULL      | The database is full.                                        |
	//	+-------------------------------------+--------------------------------------------------------------+
	PoolID *dtyp.GUID `idl:"name:lpPoolId" json:"pool_id"`
}

func (o *ChangeNTMSMediaTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_ChangeNTMSMediaTypeOperation) *xxx_ChangeNTMSMediaTypeOperation {
	if op == nil {
		op = &xxx_ChangeNTMSMediaTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MediaID = o.MediaID
	op.PoolID = o.PoolID
	return op
}

func (o *ChangeNTMSMediaTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangeNTMSMediaTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MediaID = op.MediaID
	o.PoolID = op.PoolID
}
func (o *ChangeNTMSMediaTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ChangeNTMSMediaTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeNTMSMediaTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangeNTMSMediaTypeResponse structure represents the ChangeNtmsMediaType operation response
type ChangeNTMSMediaTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ChangeNtmsMediaType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ChangeNTMSMediaTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_ChangeNTMSMediaTypeOperation) *xxx_ChangeNTMSMediaTypeOperation {
	if op == nil {
		op = &xxx_ChangeNTMSMediaTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ChangeNTMSMediaTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangeNTMSMediaTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ChangeNTMSMediaTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ChangeNTMSMediaTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeNTMSMediaTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

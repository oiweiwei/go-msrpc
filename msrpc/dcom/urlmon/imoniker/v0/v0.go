package imoniker

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	urlmon "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon"
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
	_ = urlmon.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/urlmon"
)

var (
	// IMoniker interface identifier 0000000f-0000-0000-c000-000000000046
	MonikerIID = &dcom.IID{Data1: 0x0000000f, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	MonikerSyntaxUUID = &uuid.UUID{TimeLow: 0xf, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	MonikerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: MonikerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMoniker interface.
type MonikerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetClassID operation.
	GetClassID(context.Context, *GetClassIDRequest, ...dcerpc.CallOption) (*GetClassIDResponse, error)

	// IsDirty operation.
	IsDirty(context.Context, *IsDirtyRequest, ...dcerpc.CallOption) (*IsDirtyResponse, error)

	// Load operation.
	Load(context.Context, *LoadRequest, ...dcerpc.CallOption) (*LoadResponse, error)

	// Save operation.
	Save(context.Context, *SaveRequest, ...dcerpc.CallOption) (*SaveResponse, error)

	// GetSizeMax operation.
	GetSizeMax(context.Context, *GetSizeMaxRequest, ...dcerpc.CallOption) (*GetSizeMaxResponse, error)

	// BindToObject operation.
	BindToObject(context.Context, *BindToObjectRequest, ...dcerpc.CallOption) (*BindToObjectResponse, error)

	// BindToStorage operation.
	BindToStorage(context.Context, *BindToStorageRequest, ...dcerpc.CallOption) (*BindToStorageResponse, error)

	// Reduce operation.
	Reduce(context.Context, *ReduceRequest, ...dcerpc.CallOption) (*ReduceResponse, error)

	// ComposeWith operation.
	ComposeWith(context.Context, *ComposeWithRequest, ...dcerpc.CallOption) (*ComposeWithResponse, error)

	// Enum operation.
	Enum(context.Context, *EnumRequest, ...dcerpc.CallOption) (*EnumResponse, error)

	// IsEqual operation.
	IsEqual(context.Context, *IsEqualRequest, ...dcerpc.CallOption) (*IsEqualResponse, error)

	// Hash operation.
	Hash(context.Context, *HashRequest, ...dcerpc.CallOption) (*HashResponse, error)

	// IsRunning operation.
	IsRunning(context.Context, *IsRunningRequest, ...dcerpc.CallOption) (*IsRunningResponse, error)

	// GetTimeOfLastChange operation.
	GetTimeOfLastChange(context.Context, *GetTimeOfLastChangeRequest, ...dcerpc.CallOption) (*GetTimeOfLastChangeResponse, error)

	// Inverse operation.
	Inverse(context.Context, *InverseRequest, ...dcerpc.CallOption) (*InverseResponse, error)

	// CommonPrefixWith operation.
	CommonPrefixWith(context.Context, *CommonPrefixWithRequest, ...dcerpc.CallOption) (*CommonPrefixWithResponse, error)

	// RelativePathTo operation.
	RelativePathTo(context.Context, *RelativePathToRequest, ...dcerpc.CallOption) (*RelativePathToResponse, error)

	// GetDisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest, ...dcerpc.CallOption) (*GetDisplayNameResponse, error)

	// ParseDisplayName operation.
	ParseDisplayName(context.Context, *ParseDisplayNameRequest, ...dcerpc.CallOption) (*ParseDisplayNameResponse, error)

	// IsSystemMoniker operation.
	IsSystemMoniker(context.Context, *IsSystemMonikerRequest, ...dcerpc.CallOption) (*IsSystemMonikerResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) MonikerClient
}

type xxx_DefaultMonikerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultMonikerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultMonikerClient) GetClassID(ctx context.Context, in *GetClassIDRequest, opts ...dcerpc.CallOption) (*GetClassIDResponse, error) {
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
	out := &GetClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) IsDirty(ctx context.Context, in *IsDirtyRequest, opts ...dcerpc.CallOption) (*IsDirtyResponse, error) {
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
	out := &IsDirtyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) Load(ctx context.Context, in *LoadRequest, opts ...dcerpc.CallOption) (*LoadResponse, error) {
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
	out := &LoadResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) Save(ctx context.Context, in *SaveRequest, opts ...dcerpc.CallOption) (*SaveResponse, error) {
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
	out := &SaveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) GetSizeMax(ctx context.Context, in *GetSizeMaxRequest, opts ...dcerpc.CallOption) (*GetSizeMaxResponse, error) {
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
	out := &GetSizeMaxResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) BindToObject(ctx context.Context, in *BindToObjectRequest, opts ...dcerpc.CallOption) (*BindToObjectResponse, error) {
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
	out := &BindToObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) BindToStorage(ctx context.Context, in *BindToStorageRequest, opts ...dcerpc.CallOption) (*BindToStorageResponse, error) {
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
	out := &BindToStorageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) Reduce(ctx context.Context, in *ReduceRequest, opts ...dcerpc.CallOption) (*ReduceResponse, error) {
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
	out := &ReduceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) ComposeWith(ctx context.Context, in *ComposeWithRequest, opts ...dcerpc.CallOption) (*ComposeWithResponse, error) {
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
	out := &ComposeWithResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) Enum(ctx context.Context, in *EnumRequest, opts ...dcerpc.CallOption) (*EnumResponse, error) {
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
	out := &EnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) IsEqual(ctx context.Context, in *IsEqualRequest, opts ...dcerpc.CallOption) (*IsEqualResponse, error) {
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
	out := &IsEqualResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) Hash(ctx context.Context, in *HashRequest, opts ...dcerpc.CallOption) (*HashResponse, error) {
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
	out := &HashResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) IsRunning(ctx context.Context, in *IsRunningRequest, opts ...dcerpc.CallOption) (*IsRunningResponse, error) {
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
	out := &IsRunningResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) GetTimeOfLastChange(ctx context.Context, in *GetTimeOfLastChangeRequest, opts ...dcerpc.CallOption) (*GetTimeOfLastChangeResponse, error) {
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
	out := &GetTimeOfLastChangeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) Inverse(ctx context.Context, in *InverseRequest, opts ...dcerpc.CallOption) (*InverseResponse, error) {
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
	out := &InverseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) CommonPrefixWith(ctx context.Context, in *CommonPrefixWithRequest, opts ...dcerpc.CallOption) (*CommonPrefixWithResponse, error) {
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
	out := &CommonPrefixWithResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) RelativePathTo(ctx context.Context, in *RelativePathToRequest, opts ...dcerpc.CallOption) (*RelativePathToResponse, error) {
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
	out := &RelativePathToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) GetDisplayName(ctx context.Context, in *GetDisplayNameRequest, opts ...dcerpc.CallOption) (*GetDisplayNameResponse, error) {
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
	out := &GetDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) ParseDisplayName(ctx context.Context, in *ParseDisplayNameRequest, opts ...dcerpc.CallOption) (*ParseDisplayNameResponse, error) {
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
	out := &ParseDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) IsSystemMoniker(ctx context.Context, in *IsSystemMonikerRequest, opts ...dcerpc.CallOption) (*IsSystemMonikerResponse, error) {
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
	out := &IsSystemMonikerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMonikerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMonikerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultMonikerClient) IPID(ctx context.Context, ipid *dcom.IPID) MonikerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultMonikerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewMonikerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (MonikerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(MonikerSyntaxV0_0))...)
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
	return &xxx_DefaultMonikerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetClassIDOperation structure represents the GetClassID operation
type xxx_GetClassIDOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClassID *dtyp.GUID     `idl:"name:pClassID" json:"class_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassIDOperation) OpNum() int { return 3 }

func (o *xxx_GetClassIDOperation) OpName() string { return "/IMoniker/v0/GetClassID" }

func (o *xxx_GetClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pClassID {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClassID != nil {
			if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pClassID {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClassID == nil {
			o.ClassID = &dtyp.GUID{}
		}
		if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
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

// GetClassIDRequest structure represents the GetClassID operation request
type GetClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClassIDOperation) *xxx_GetClassIDOperation {
	if op == nil {
		op = &xxx_GetClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassIDResponse structure represents the GetClassID operation response
type GetClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClassID *dtyp.GUID     `idl:"name:pClassID" json:"class_id"`
	// Return: The GetClassID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClassIDOperation) *xxx_GetClassIDOperation {
	if op == nil {
		op = &xxx_GetClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ClassID = o.ClassID
	op.Return = o.Return
	return op
}

func (o *GetClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClassID = op.ClassID
	o.Return = op.Return
}
func (o *GetClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsDirtyOperation structure represents the IsDirty operation
type xxx_IsDirtyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsDirtyOperation) OpNum() int { return 4 }

func (o *xxx_IsDirtyOperation) OpName() string { return "/IMoniker/v0/IsDirty" }

func (o *xxx_IsDirtyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsDirtyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsDirtyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsDirtyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsDirtyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsDirtyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IsDirtyRequest structure represents the IsDirty operation request
type IsDirtyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsDirtyRequest) xxx_ToOp(ctx context.Context, op *xxx_IsDirtyOperation) *xxx_IsDirtyOperation {
	if op == nil {
		op = &xxx_IsDirtyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *IsDirtyRequest) xxx_FromOp(ctx context.Context, op *xxx_IsDirtyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsDirtyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsDirtyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsDirtyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsDirtyResponse structure represents the IsDirty operation response
type IsDirtyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IsDirty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsDirtyResponse) xxx_ToOp(ctx context.Context, op *xxx_IsDirtyOperation) *xxx_IsDirtyOperation {
	if op == nil {
		op = &xxx_IsDirtyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *IsDirtyResponse) xxx_FromOp(ctx context.Context, op *xxx_IsDirtyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsDirtyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsDirtyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsDirtyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LoadOperation structure represents the Load operation
type xxx_LoadOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Stream *urlmon.Stream `idl:"name:pStm" json:"stream"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_LoadOperation) OpNum() int { return 5 }

func (o *xxx_LoadOperation) OpName() string { return "/IMoniker/v0/Load" }

func (o *xxx_LoadOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pStm {in} (1:{pointer=ref}*(1))(2:{alias=IStream}(interface))
	{
		if o.Stream != nil {
			_ptr_pStm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Stream != nil {
					if err := o.Stream.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Stream{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Stream, _ptr_pStm); err != nil {
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

func (o *xxx_LoadOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pStm {in} (1:{pointer=ref}*(1))(2:{alias=IStream}(interface))
	{
		_ptr_pStm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Stream == nil {
				o.Stream = &urlmon.Stream{}
			}
			if err := o.Stream.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pStm := func(ptr interface{}) { o.Stream = *ptr.(**urlmon.Stream) }
		if err := w.ReadPointer(&o.Stream, _s_pStm, _ptr_pStm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_LoadOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// LoadRequest structure represents the Load operation request
type LoadRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Stream *urlmon.Stream `idl:"name:pStm" json:"stream"`
}

func (o *LoadRequest) xxx_ToOp(ctx context.Context, op *xxx_LoadOperation) *xxx_LoadOperation {
	if op == nil {
		op = &xxx_LoadOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Stream = o.Stream
	return op
}

func (o *LoadRequest) xxx_FromOp(ctx context.Context, op *xxx_LoadOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Stream = op.Stream
}
func (o *LoadRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LoadRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LoadOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LoadResponse structure represents the Load operation response
type LoadResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Load return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LoadResponse) xxx_ToOp(ctx context.Context, op *xxx_LoadOperation) *xxx_LoadOperation {
	if op == nil {
		op = &xxx_LoadOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *LoadResponse) xxx_FromOp(ctx context.Context, op *xxx_LoadOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *LoadResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LoadResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LoadOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SaveOperation structure represents the Save operation
type xxx_SaveOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Stream     *urlmon.Stream `idl:"name:pStm" json:"stream"`
	ClearDirty bool           `idl:"name:fClearDirty" json:"clear_dirty"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SaveOperation) OpNum() int { return 6 }

func (o *xxx_SaveOperation) OpName() string { return "/IMoniker/v0/Save" }

func (o *xxx_SaveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pStm {in} (1:{pointer=ref}*(1))(2:{alias=IStream}(interface))
	{
		if o.Stream != nil {
			_ptr_pStm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Stream != nil {
					if err := o.Stream.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Stream{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Stream, _ptr_pStm); err != nil {
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
	// fClearDirty {in} (1:{alias=BOOL}(int32))
	{
		if !o.ClearDirty {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SaveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pStm {in} (1:{pointer=ref}*(1))(2:{alias=IStream}(interface))
	{
		_ptr_pStm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Stream == nil {
				o.Stream = &urlmon.Stream{}
			}
			if err := o.Stream.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pStm := func(ptr interface{}) { o.Stream = *ptr.(**urlmon.Stream) }
		if err := w.ReadPointer(&o.Stream, _s_pStm, _ptr_pStm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fClearDirty {in} (1:{alias=BOOL}(int32))
	{
		var _bClearDirty int32
		if err := w.ReadData(&_bClearDirty); err != nil {
			return err
		}
		o.ClearDirty = _bClearDirty != 0
	}
	return nil
}

func (o *xxx_SaveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SaveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SaveRequest structure represents the Save operation request
type SaveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	Stream     *urlmon.Stream `idl:"name:pStm" json:"stream"`
	ClearDirty bool           `idl:"name:fClearDirty" json:"clear_dirty"`
}

func (o *SaveRequest) xxx_ToOp(ctx context.Context, op *xxx_SaveOperation) *xxx_SaveOperation {
	if op == nil {
		op = &xxx_SaveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Stream = o.Stream
	op.ClearDirty = o.ClearDirty
	return op
}

func (o *SaveRequest) xxx_FromOp(ctx context.Context, op *xxx_SaveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Stream = op.Stream
	o.ClearDirty = op.ClearDirty
}
func (o *SaveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SaveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SaveResponse structure represents the Save operation response
type SaveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Save return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SaveResponse) xxx_ToOp(ctx context.Context, op *xxx_SaveOperation) *xxx_SaveOperation {
	if op == nil {
		op = &xxx_SaveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SaveResponse) xxx_FromOp(ctx context.Context, op *xxx_SaveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SaveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SaveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSizeMaxOperation structure represents the GetSizeMax operation
type xxx_GetSizeMaxOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Length *dtyp.UlargeInteger `idl:"name:pcbSize" json:"length"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSizeMaxOperation) OpNum() int { return 7 }

func (o *xxx_GetSizeMaxOperation) OpName() string { return "/IMoniker/v0/GetSizeMax" }

func (o *xxx_GetSizeMaxOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSizeMaxOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSizeMaxOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSizeMaxOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSizeMaxOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcbSize {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length != nil {
			if err := o.Length.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetSizeMaxOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcbSize {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length == nil {
			o.Length = &dtyp.UlargeInteger{}
		}
		if err := o.Length.UnmarshalNDR(ctx, w); err != nil {
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

// GetSizeMaxRequest structure represents the GetSizeMax operation request
type GetSizeMaxRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSizeMaxRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSizeMaxOperation) *xxx_GetSizeMaxOperation {
	if op == nil {
		op = &xxx_GetSizeMaxOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSizeMaxRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSizeMaxOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSizeMaxRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSizeMaxRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSizeMaxOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSizeMaxResponse structure represents the GetSizeMax operation response
type GetSizeMaxResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Length *dtyp.UlargeInteger `idl:"name:pcbSize" json:"length"`
	// Return: The GetSizeMax return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSizeMaxResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSizeMaxOperation) *xxx_GetSizeMaxOperation {
	if op == nil {
		op = &xxx_GetSizeMaxOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Length = o.Length
	op.Return = o.Return
	return op
}

func (o *GetSizeMaxResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSizeMaxOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Length = op.Length
	o.Return = op.Return
}
func (o *GetSizeMaxResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSizeMaxResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSizeMaxOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BindToObjectOperation structure represents the BindToObject operation
type xxx_BindToObjectOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	IIDResult   *dcom.IID           `idl:"name:riidResult" json:"iid_result"`
	Result      []byte              `idl:"name:ppvResult" json:"result"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_BindToObjectOperation) OpNum() int { return 8 }

func (o *xxx_BindToObjectOperation) OpName() string { return "/IMoniker/v0/BindToObject" }

func (o *xxx_BindToObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindToObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_pmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_pmkToLeft); err != nil {
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
	// riidResult {in} (1:{alias=REFIID}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IIDResult != nil {
			if err := o.IIDResult.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BindToObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_pmkToLeft, _ptr_pmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// riidResult {in} (1:{alias=REFIID,pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IIDResult == nil {
			o.IIDResult = &dcom.IID{}
		}
		if err := o.IIDResult.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindToObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindToObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppvResult {out, iid_is={riidResult}} (1:{pointer=ref}*(2)*(1)(void))
	{
		if o.Result != nil {
			_ptr_ppvResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				// FIXME unknown type ppvResult
				return nil
			})
			if err := w.WritePointer(&o.Result, _ptr_ppvResult); err != nil {
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

func (o *xxx_BindToObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppvResult {out, iid_is={riidResult}} (1:{pointer=ref}*(2)*(1)(void))
	{
		_ptr_ppvResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			// FIXME: unknown type ppvResult
			return nil
		})
		_s_ppvResult := func(ptr interface{}) { o.Result = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Result, _s_ppvResult, _ptr_ppvResult); err != nil {
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

// BindToObjectRequest structure represents the BindToObject operation request
type BindToObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	IIDResult   *dcom.IID           `idl:"name:riidResult" json:"iid_result"`
}

func (o *BindToObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_BindToObjectOperation) *xxx_BindToObjectOperation {
	if op == nil {
		op = &xxx_BindToObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindContext = o.BindContext
	op.ToLeft = o.ToLeft
	op.IIDResult = o.IIDResult
	return op
}

func (o *BindToObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_BindToObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindContext = op.BindContext
	o.ToLeft = op.ToLeft
	o.IIDResult = op.IIDResult
}
func (o *BindToObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BindToObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindToObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BindToObjectResponse structure represents the BindToObject operation response
type BindToObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Result []byte         `idl:"name:ppvResult" json:"result"`
	// Return: The BindToObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BindToObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_BindToObjectOperation) *xxx_BindToObjectOperation {
	if op == nil {
		op = &xxx_BindToObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *BindToObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_BindToObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Result = op.Result
	o.Return = op.Return
}
func (o *BindToObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BindToObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindToObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BindToStorageOperation structure represents the BindToStorage operation
type xxx_BindToStorageOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	IID         *dcom.IID           `idl:"name:riid" json:"iid"`
	Object      []byte              `idl:"name:ppvObj" json:"object"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_BindToStorageOperation) OpNum() int { return 9 }

func (o *xxx_BindToStorageOperation) OpName() string { return "/IMoniker/v0/BindToStorage" }

func (o *xxx_BindToStorageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindToStorageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_pmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_pmkToLeft); err != nil {
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
	// riid {in} (1:{alias=REFIID}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID != nil {
			if err := o.IID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BindToStorageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_pmkToLeft, _ptr_pmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// riid {in} (1:{alias=REFIID,pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID == nil {
			o.IID = &dcom.IID{}
		}
		if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindToStorageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindToStorageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppvObj {out, iid_is={riid}} (1:{pointer=ref}*(2)*(1)(void))
	{
		if o.Object != nil {
			_ptr_ppvObj := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				// FIXME unknown type ppvObj
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_ppvObj); err != nil {
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

func (o *xxx_BindToStorageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppvObj {out, iid_is={riid}} (1:{pointer=ref}*(2)*(1)(void))
	{
		_ptr_ppvObj := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			// FIXME: unknown type ppvObj
			return nil
		})
		_s_ppvObj := func(ptr interface{}) { o.Object = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Object, _s_ppvObj, _ptr_ppvObj); err != nil {
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

// BindToStorageRequest structure represents the BindToStorage operation request
type BindToStorageRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	IID         *dcom.IID           `idl:"name:riid" json:"iid"`
}

func (o *BindToStorageRequest) xxx_ToOp(ctx context.Context, op *xxx_BindToStorageOperation) *xxx_BindToStorageOperation {
	if op == nil {
		op = &xxx_BindToStorageOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindContext = o.BindContext
	op.ToLeft = o.ToLeft
	op.IID = o.IID
	return op
}

func (o *BindToStorageRequest) xxx_FromOp(ctx context.Context, op *xxx_BindToStorageOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindContext = op.BindContext
	o.ToLeft = op.ToLeft
	o.IID = op.IID
}
func (o *BindToStorageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BindToStorageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindToStorageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BindToStorageResponse structure represents the BindToStorage operation response
type BindToStorageResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Object []byte         `idl:"name:ppvObj" json:"object"`
	// Return: The BindToStorage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BindToStorageResponse) xxx_ToOp(ctx context.Context, op *xxx_BindToStorageOperation) *xxx_BindToStorageOperation {
	if op == nil {
		op = &xxx_BindToStorageOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Object = o.Object
	op.Return = o.Return
	return op
}

func (o *BindToStorageResponse) xxx_FromOp(ctx context.Context, op *xxx_BindToStorageOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Object = op.Object
	o.Return = op.Return
}
func (o *BindToStorageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BindToStorageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindToStorageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReduceOperation structure represents the Reduce operation
type xxx_ReduceOperation struct {
	This         *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindContext  *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ReduceHowFar uint32              `idl:"name:dwReduceHowFar" json:"reduce_how_far"`
	ToLeft       *urlmon.Moniker     `idl:"name:ppmkToLeft" json:"to_left"`
	Reduced      *urlmon.Moniker     `idl:"name:ppmkReduced" json:"reduced"`
	Return       int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_ReduceOperation) OpNum() int { return 10 }

func (o *xxx_ReduceOperation) OpName() string { return "/IMoniker/v0/Reduce" }

func (o *xxx_ReduceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReduceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// dwReduceHowFar {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReduceHowFar); err != nil {
			return err
		}
	}
	// ppmkToLeft {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_ppmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_ppmkToLeft); err != nil {
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

func (o *xxx_ReduceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwReduceHowFar {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReduceHowFar); err != nil {
			return err
		}
	}
	// ppmkToLeft {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_ppmkToLeft, _ptr_ppmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReduceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReduceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmkToLeft {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_ppmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_ppmkToLeft); err != nil {
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
	// ppmkReduced {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Reduced != nil {
			_ptr_ppmkReduced := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Reduced != nil {
					if err := o.Reduced.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Reduced, _ptr_ppmkReduced); err != nil {
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

func (o *xxx_ReduceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmkToLeft {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_ppmkToLeft, _ptr_ppmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppmkReduced {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmkReduced := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Reduced == nil {
				o.Reduced = &urlmon.Moniker{}
			}
			if err := o.Reduced.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmkReduced := func(ptr interface{}) { o.Reduced = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Reduced, _s_ppmkReduced, _ptr_ppmkReduced); err != nil {
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

// ReduceRequest structure represents the Reduce operation request
type ReduceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindContext  *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ReduceHowFar uint32              `idl:"name:dwReduceHowFar" json:"reduce_how_far"`
	ToLeft       *urlmon.Moniker     `idl:"name:ppmkToLeft" json:"to_left"`
}

func (o *ReduceRequest) xxx_ToOp(ctx context.Context, op *xxx_ReduceOperation) *xxx_ReduceOperation {
	if op == nil {
		op = &xxx_ReduceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindContext = o.BindContext
	op.ReduceHowFar = o.ReduceHowFar
	op.ToLeft = o.ToLeft
	return op
}

func (o *ReduceRequest) xxx_FromOp(ctx context.Context, op *xxx_ReduceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindContext = op.BindContext
	o.ReduceHowFar = op.ReduceHowFar
	o.ToLeft = op.ToLeft
}
func (o *ReduceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReduceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReduceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReduceResponse structure represents the Reduce operation response
type ReduceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ToLeft  *urlmon.Moniker `idl:"name:ppmkToLeft" json:"to_left"`
	Reduced *urlmon.Moniker `idl:"name:ppmkReduced" json:"reduced"`
	// Return: The Reduce return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReduceResponse) xxx_ToOp(ctx context.Context, op *xxx_ReduceOperation) *xxx_ReduceOperation {
	if op == nil {
		op = &xxx_ReduceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ToLeft = o.ToLeft
	op.Reduced = o.Reduced
	op.Return = o.Return
	return op
}

func (o *ReduceResponse) xxx_FromOp(ctx context.Context, op *xxx_ReduceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ToLeft = op.ToLeft
	o.Reduced = op.Reduced
	o.Return = op.Return
}
func (o *ReduceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReduceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReduceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ComposeWithOperation structure represents the ComposeWith operation
type xxx_ComposeWithOperation struct {
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Right                   *urlmon.Moniker `idl:"name:pmkRight" json:"right"`
	OnlyInterfaceNotGeneric bool            `idl:"name:fOnlyIfNotGeneric" json:"only_interface_not_generic"`
	Composite               *urlmon.Moniker `idl:"name:ppmkComposite" json:"composite"`
	Return                  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ComposeWithOperation) OpNum() int { return 11 }

func (o *xxx_ComposeWithOperation) OpName() string { return "/IMoniker/v0/ComposeWith" }

func (o *xxx_ComposeWithOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComposeWithOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pmkRight {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Right != nil {
			_ptr_pmkRight := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Right != nil {
					if err := o.Right.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Right, _ptr_pmkRight); err != nil {
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
	// fOnlyIfNotGeneric {in} (1:{alias=BOOL}(int32))
	{
		if !o.OnlyInterfaceNotGeneric {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ComposeWithOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pmkRight {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkRight := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Right == nil {
				o.Right = &urlmon.Moniker{}
			}
			if err := o.Right.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkRight := func(ptr interface{}) { o.Right = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Right, _s_pmkRight, _ptr_pmkRight); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fOnlyIfNotGeneric {in} (1:{alias=BOOL}(int32))
	{
		var _bOnlyInterfaceNotGeneric int32
		if err := w.ReadData(&_bOnlyInterfaceNotGeneric); err != nil {
			return err
		}
		o.OnlyInterfaceNotGeneric = _bOnlyInterfaceNotGeneric != 0
	}
	return nil
}

func (o *xxx_ComposeWithOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComposeWithOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmkComposite {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Composite != nil {
			_ptr_ppmkComposite := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Composite != nil {
					if err := o.Composite.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Composite, _ptr_ppmkComposite); err != nil {
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

func (o *xxx_ComposeWithOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmkComposite {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmkComposite := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Composite == nil {
				o.Composite = &urlmon.Moniker{}
			}
			if err := o.Composite.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmkComposite := func(ptr interface{}) { o.Composite = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Composite, _s_ppmkComposite, _ptr_ppmkComposite); err != nil {
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

// ComposeWithRequest structure represents the ComposeWith operation request
type ComposeWithRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Right                   *urlmon.Moniker `idl:"name:pmkRight" json:"right"`
	OnlyInterfaceNotGeneric bool            `idl:"name:fOnlyIfNotGeneric" json:"only_interface_not_generic"`
}

func (o *ComposeWithRequest) xxx_ToOp(ctx context.Context, op *xxx_ComposeWithOperation) *xxx_ComposeWithOperation {
	if op == nil {
		op = &xxx_ComposeWithOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Right = o.Right
	op.OnlyInterfaceNotGeneric = o.OnlyInterfaceNotGeneric
	return op
}

func (o *ComposeWithRequest) xxx_FromOp(ctx context.Context, op *xxx_ComposeWithOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Right = op.Right
	o.OnlyInterfaceNotGeneric = op.OnlyInterfaceNotGeneric
}
func (o *ComposeWithRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ComposeWithRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComposeWithOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ComposeWithResponse structure represents the ComposeWith operation response
type ComposeWithResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Composite *urlmon.Moniker `idl:"name:ppmkComposite" json:"composite"`
	// Return: The ComposeWith return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ComposeWithResponse) xxx_ToOp(ctx context.Context, op *xxx_ComposeWithOperation) *xxx_ComposeWithOperation {
	if op == nil {
		op = &xxx_ComposeWithOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Composite = o.Composite
	op.Return = o.Return
	return op
}

func (o *ComposeWithResponse) xxx_FromOp(ctx context.Context, op *xxx_ComposeWithOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Composite = op.Composite
	o.Return = op.Return
}
func (o *ComposeWithResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ComposeWithResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComposeWithOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumOperation structure represents the Enum operation
type xxx_EnumOperation struct {
	This    *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Forward bool                `idl:"name:fForward" json:"forward"`
	Moniker *urlmon.EnumMoniker `idl:"name:ppenumMoniker" json:"moniker"`
	Return  int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumOperation) OpNum() int { return 12 }

func (o *xxx_EnumOperation) OpName() string { return "/IMoniker/v0/Enum" }

func (o *xxx_EnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fForward {in} (1:{alias=BOOL}(int32))
	{
		if !o.Forward {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fForward {in} (1:{alias=BOOL}(int32))
	{
		var _bForward int32
		if err := w.ReadData(&_bForward); err != nil {
			return err
		}
		o.Forward = _bForward != 0
	}
	return nil
}

func (o *xxx_EnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppenumMoniker {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumMoniker}(interface))
	{
		if o.Moniker != nil {
			_ptr_ppenumMoniker := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Moniker != nil {
					if err := o.Moniker.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.EnumMoniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Moniker, _ptr_ppenumMoniker); err != nil {
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

func (o *xxx_EnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppenumMoniker {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumMoniker}(interface))
	{
		_ptr_ppenumMoniker := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Moniker == nil {
				o.Moniker = &urlmon.EnumMoniker{}
			}
			if err := o.Moniker.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppenumMoniker := func(ptr interface{}) { o.Moniker = *ptr.(**urlmon.EnumMoniker) }
		if err := w.ReadPointer(&o.Moniker, _s_ppenumMoniker, _ptr_ppenumMoniker); err != nil {
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

// EnumRequest structure represents the Enum operation request
type EnumRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Forward bool           `idl:"name:fForward" json:"forward"`
}

func (o *EnumRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumOperation) *xxx_EnumOperation {
	if op == nil {
		op = &xxx_EnumOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Forward = o.Forward
	return op
}

func (o *EnumRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Forward = op.Forward
}
func (o *EnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumResponse structure represents the Enum operation response
type EnumResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Moniker *urlmon.EnumMoniker `idl:"name:ppenumMoniker" json:"moniker"`
	// Return: The Enum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumOperation) *xxx_EnumOperation {
	if op == nil {
		op = &xxx_EnumOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Moniker = o.Moniker
	op.Return = o.Return
	return op
}

func (o *EnumResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Moniker = op.Moniker
	o.Return = op.Return
}
func (o *EnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsEqualOperation structure represents the IsEqual operation
type xxx_IsEqualOperation struct {
	This         *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	OtherMoniker *urlmon.Moniker `idl:"name:pmkOtherMoniker" json:"other_moniker"`
	Return       int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_IsEqualOperation) OpNum() int { return 13 }

func (o *xxx_IsEqualOperation) OpName() string { return "/IMoniker/v0/IsEqual" }

func (o *xxx_IsEqualOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsEqualOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pmkOtherMoniker {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.OtherMoniker != nil {
			_ptr_pmkOtherMoniker := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OtherMoniker != nil {
					if err := o.OtherMoniker.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OtherMoniker, _ptr_pmkOtherMoniker); err != nil {
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

func (o *xxx_IsEqualOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pmkOtherMoniker {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkOtherMoniker := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OtherMoniker == nil {
				o.OtherMoniker = &urlmon.Moniker{}
			}
			if err := o.OtherMoniker.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkOtherMoniker := func(ptr interface{}) { o.OtherMoniker = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.OtherMoniker, _s_pmkOtherMoniker, _ptr_pmkOtherMoniker); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsEqualOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsEqualOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsEqualOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IsEqualRequest structure represents the IsEqual operation request
type IsEqualRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis  `idl:"name:This" json:"this"`
	OtherMoniker *urlmon.Moniker `idl:"name:pmkOtherMoniker" json:"other_moniker"`
}

func (o *IsEqualRequest) xxx_ToOp(ctx context.Context, op *xxx_IsEqualOperation) *xxx_IsEqualOperation {
	if op == nil {
		op = &xxx_IsEqualOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OtherMoniker = o.OtherMoniker
	return op
}

func (o *IsEqualRequest) xxx_FromOp(ctx context.Context, op *xxx_IsEqualOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OtherMoniker = op.OtherMoniker
}
func (o *IsEqualRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsEqualRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsEqualOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsEqualResponse structure represents the IsEqual operation response
type IsEqualResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IsEqual return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsEqualResponse) xxx_ToOp(ctx context.Context, op *xxx_IsEqualOperation) *xxx_IsEqualOperation {
	if op == nil {
		op = &xxx_IsEqualOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *IsEqualResponse) xxx_FromOp(ctx context.Context, op *xxx_IsEqualOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsEqualResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsEqualResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsEqualOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_HashOperation structure represents the Hash operation
type xxx_HashOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Hash   uint32         `idl:"name:pdwHash" json:"hash"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_HashOperation) OpNum() int { return 14 }

func (o *xxx_HashOperation) OpName() string { return "/IMoniker/v0/Hash" }

func (o *xxx_HashOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HashOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_HashOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_HashOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HashOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwHash {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Hash); err != nil {
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

func (o *xxx_HashOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwHash {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Hash); err != nil {
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

// HashRequest structure represents the Hash operation request
type HashRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *HashRequest) xxx_ToOp(ctx context.Context, op *xxx_HashOperation) *xxx_HashOperation {
	if op == nil {
		op = &xxx_HashOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *HashRequest) xxx_FromOp(ctx context.Context, op *xxx_HashOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *HashRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *HashRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HashOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// HashResponse structure represents the Hash operation response
type HashResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Hash uint32         `idl:"name:pdwHash" json:"hash"`
	// Return: The Hash return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *HashResponse) xxx_ToOp(ctx context.Context, op *xxx_HashOperation) *xxx_HashOperation {
	if op == nil {
		op = &xxx_HashOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Hash = o.Hash
	op.Return = o.Return
	return op
}

func (o *HashResponse) xxx_FromOp(ctx context.Context, op *xxx_HashOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Hash = op.Hash
	o.Return = op.Return
}
func (o *HashResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *HashResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HashOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsRunningOperation structure represents the IsRunning operation
type xxx_IsRunningOperation struct {
	This         *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindContext  *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft       *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	NewlyRunning *urlmon.Moniker     `idl:"name:pmkNewlyRunning" json:"newly_running"`
	Return       int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_IsRunningOperation) OpNum() int { return 15 }

func (o *xxx_IsRunningOperation) OpName() string { return "/IMoniker/v0/IsRunning" }

func (o *xxx_IsRunningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsRunningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_pmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_pmkToLeft); err != nil {
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
	// pmkNewlyRunning {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.NewlyRunning != nil {
			_ptr_pmkNewlyRunning := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewlyRunning != nil {
					if err := o.NewlyRunning.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewlyRunning, _ptr_pmkNewlyRunning); err != nil {
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

func (o *xxx_IsRunningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_pmkToLeft, _ptr_pmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkNewlyRunning {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkNewlyRunning := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewlyRunning == nil {
				o.NewlyRunning = &urlmon.Moniker{}
			}
			if err := o.NewlyRunning.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkNewlyRunning := func(ptr interface{}) { o.NewlyRunning = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.NewlyRunning, _s_pmkNewlyRunning, _ptr_pmkNewlyRunning); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsRunningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsRunningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsRunningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IsRunningRequest structure represents the IsRunning operation request
type IsRunningRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindContext  *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft       *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	NewlyRunning *urlmon.Moniker     `idl:"name:pmkNewlyRunning" json:"newly_running"`
}

func (o *IsRunningRequest) xxx_ToOp(ctx context.Context, op *xxx_IsRunningOperation) *xxx_IsRunningOperation {
	if op == nil {
		op = &xxx_IsRunningOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindContext = o.BindContext
	op.ToLeft = o.ToLeft
	op.NewlyRunning = o.NewlyRunning
	return op
}

func (o *IsRunningRequest) xxx_FromOp(ctx context.Context, op *xxx_IsRunningOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindContext = op.BindContext
	o.ToLeft = op.ToLeft
	o.NewlyRunning = op.NewlyRunning
}
func (o *IsRunningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsRunningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsRunningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsRunningResponse structure represents the IsRunning operation response
type IsRunningResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IsRunning return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsRunningResponse) xxx_ToOp(ctx context.Context, op *xxx_IsRunningOperation) *xxx_IsRunningOperation {
	if op == nil {
		op = &xxx_IsRunningOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *IsRunningResponse) xxx_FromOp(ctx context.Context, op *xxx_IsRunningOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsRunningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsRunningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsRunningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTimeOfLastChangeOperation structure represents the GetTimeOfLastChange operation
type xxx_GetTimeOfLastChangeOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	FileTime    *dtyp.Filetime      `idl:"name:pFileTime" json:"file_time"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTimeOfLastChangeOperation) OpNum() int { return 16 }

func (o *xxx_GetTimeOfLastChangeOperation) OpName() string { return "/IMoniker/v0/GetTimeOfLastChange" }

func (o *xxx_GetTimeOfLastChangeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTimeOfLastChangeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_pmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_pmkToLeft); err != nil {
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

func (o *xxx_GetTimeOfLastChangeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_pmkToLeft, _ptr_pmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTimeOfLastChangeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTimeOfLastChangeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pFileTime {out} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.FileTime != nil {
			if err := o.FileTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetTimeOfLastChangeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pFileTime {out} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.FileTime == nil {
			o.FileTime = &dtyp.Filetime{}
		}
		if err := o.FileTime.UnmarshalNDR(ctx, w); err != nil {
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

// GetTimeOfLastChangeRequest structure represents the GetTimeOfLastChange operation request
type GetTimeOfLastChangeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
}

func (o *GetTimeOfLastChangeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTimeOfLastChangeOperation) *xxx_GetTimeOfLastChangeOperation {
	if op == nil {
		op = &xxx_GetTimeOfLastChangeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindContext = o.BindContext
	op.ToLeft = o.ToLeft
	return op
}

func (o *GetTimeOfLastChangeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTimeOfLastChangeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindContext = op.BindContext
	o.ToLeft = op.ToLeft
}
func (o *GetTimeOfLastChangeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTimeOfLastChangeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTimeOfLastChangeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTimeOfLastChangeResponse structure represents the GetTimeOfLastChange operation response
type GetTimeOfLastChangeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileTime *dtyp.Filetime `idl:"name:pFileTime" json:"file_time"`
	// Return: The GetTimeOfLastChange return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTimeOfLastChangeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTimeOfLastChangeOperation) *xxx_GetTimeOfLastChangeOperation {
	if op == nil {
		op = &xxx_GetTimeOfLastChangeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileTime = o.FileTime
	op.Return = o.Return
	return op
}

func (o *GetTimeOfLastChangeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTimeOfLastChangeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileTime = op.FileTime
	o.Return = op.Return
}
func (o *GetTimeOfLastChangeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTimeOfLastChangeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTimeOfLastChangeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InverseOperation structure represents the Inverse operation
type xxx_InverseOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Moniker *urlmon.Moniker `idl:"name:ppmk" json:"moniker"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_InverseOperation) OpNum() int { return 17 }

func (o *xxx_InverseOperation) OpName() string { return "/IMoniker/v0/Inverse" }

func (o *xxx_InverseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InverseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_InverseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_InverseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InverseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Moniker != nil {
			_ptr_ppmk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Moniker != nil {
					if err := o.Moniker.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Moniker, _ptr_ppmk); err != nil {
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

func (o *xxx_InverseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Moniker == nil {
				o.Moniker = &urlmon.Moniker{}
			}
			if err := o.Moniker.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmk := func(ptr interface{}) { o.Moniker = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Moniker, _s_ppmk, _ptr_ppmk); err != nil {
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

// InverseRequest structure represents the Inverse operation request
type InverseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *InverseRequest) xxx_ToOp(ctx context.Context, op *xxx_InverseOperation) *xxx_InverseOperation {
	if op == nil {
		op = &xxx_InverseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *InverseRequest) xxx_FromOp(ctx context.Context, op *xxx_InverseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *InverseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InverseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InverseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InverseResponse structure represents the Inverse operation response
type InverseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Moniker *urlmon.Moniker `idl:"name:ppmk" json:"moniker"`
	// Return: The Inverse return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InverseResponse) xxx_ToOp(ctx context.Context, op *xxx_InverseOperation) *xxx_InverseOperation {
	if op == nil {
		op = &xxx_InverseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Moniker = o.Moniker
	op.Return = o.Return
	return op
}

func (o *InverseResponse) xxx_FromOp(ctx context.Context, op *xxx_InverseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Moniker = op.Moniker
	o.Return = op.Return
}
func (o *InverseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InverseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InverseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CommonPrefixWithOperation structure represents the CommonPrefixWith operation
type xxx_CommonPrefixWithOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Other  *urlmon.Moniker `idl:"name:pmkOther" json:"other"`
	Prefix *urlmon.Moniker `idl:"name:ppmkPrefix" json:"prefix"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CommonPrefixWithOperation) OpNum() int { return 18 }

func (o *xxx_CommonPrefixWithOperation) OpName() string { return "/IMoniker/v0/CommonPrefixWith" }

func (o *xxx_CommonPrefixWithOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommonPrefixWithOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pmkOther {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Other != nil {
			_ptr_pmkOther := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Other != nil {
					if err := o.Other.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Other, _ptr_pmkOther); err != nil {
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

func (o *xxx_CommonPrefixWithOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pmkOther {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkOther := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Other == nil {
				o.Other = &urlmon.Moniker{}
			}
			if err := o.Other.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkOther := func(ptr interface{}) { o.Other = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Other, _s_pmkOther, _ptr_pmkOther); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommonPrefixWithOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommonPrefixWithOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmkPrefix {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Prefix != nil {
			_ptr_ppmkPrefix := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Prefix != nil {
					if err := o.Prefix.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Prefix, _ptr_ppmkPrefix); err != nil {
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

func (o *xxx_CommonPrefixWithOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmkPrefix {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmkPrefix := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Prefix == nil {
				o.Prefix = &urlmon.Moniker{}
			}
			if err := o.Prefix.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmkPrefix := func(ptr interface{}) { o.Prefix = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Prefix, _s_ppmkPrefix, _ptr_ppmkPrefix); err != nil {
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

// CommonPrefixWithRequest structure represents the CommonPrefixWith operation request
type CommonPrefixWithRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Other *urlmon.Moniker `idl:"name:pmkOther" json:"other"`
}

func (o *CommonPrefixWithRequest) xxx_ToOp(ctx context.Context, op *xxx_CommonPrefixWithOperation) *xxx_CommonPrefixWithOperation {
	if op == nil {
		op = &xxx_CommonPrefixWithOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Other = o.Other
	return op
}

func (o *CommonPrefixWithRequest) xxx_FromOp(ctx context.Context, op *xxx_CommonPrefixWithOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Other = op.Other
}
func (o *CommonPrefixWithRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CommonPrefixWithRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommonPrefixWithOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommonPrefixWithResponse structure represents the CommonPrefixWith operation response
type CommonPrefixWithResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Prefix *urlmon.Moniker `idl:"name:ppmkPrefix" json:"prefix"`
	// Return: The CommonPrefixWith return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommonPrefixWithResponse) xxx_ToOp(ctx context.Context, op *xxx_CommonPrefixWithOperation) *xxx_CommonPrefixWithOperation {
	if op == nil {
		op = &xxx_CommonPrefixWithOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Prefix = o.Prefix
	op.Return = o.Return
	return op
}

func (o *CommonPrefixWithResponse) xxx_FromOp(ctx context.Context, op *xxx_CommonPrefixWithOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Prefix = op.Prefix
	o.Return = op.Return
}
func (o *CommonPrefixWithResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CommonPrefixWithResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommonPrefixWithOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RelativePathToOperation structure represents the RelativePathTo operation
type xxx_RelativePathToOperation struct {
	This         *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Other        *urlmon.Moniker `idl:"name:pmkOther" json:"other"`
	RelationPath *urlmon.Moniker `idl:"name:ppmkRelPath" json:"relation_path"`
	Return       int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_RelativePathToOperation) OpNum() int { return 19 }

func (o *xxx_RelativePathToOperation) OpName() string { return "/IMoniker/v0/RelativePathTo" }

func (o *xxx_RelativePathToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RelativePathToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pmkOther {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Other != nil {
			_ptr_pmkOther := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Other != nil {
					if err := o.Other.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Other, _ptr_pmkOther); err != nil {
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

func (o *xxx_RelativePathToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pmkOther {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkOther := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Other == nil {
				o.Other = &urlmon.Moniker{}
			}
			if err := o.Other.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkOther := func(ptr interface{}) { o.Other = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Other, _s_pmkOther, _ptr_pmkOther); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RelativePathToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RelativePathToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmkRelPath {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.RelationPath != nil {
			_ptr_ppmkRelPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelationPath != nil {
					if err := o.RelationPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelationPath, _ptr_ppmkRelPath); err != nil {
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

func (o *xxx_RelativePathToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmkRelPath {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmkRelPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelationPath == nil {
				o.RelationPath = &urlmon.Moniker{}
			}
			if err := o.RelationPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmkRelPath := func(ptr interface{}) { o.RelationPath = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.RelationPath, _s_ppmkRelPath, _ptr_ppmkRelPath); err != nil {
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

// RelativePathToRequest structure represents the RelativePathTo operation request
type RelativePathToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Other *urlmon.Moniker `idl:"name:pmkOther" json:"other"`
}

func (o *RelativePathToRequest) xxx_ToOp(ctx context.Context, op *xxx_RelativePathToOperation) *xxx_RelativePathToOperation {
	if op == nil {
		op = &xxx_RelativePathToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Other = o.Other
	return op
}

func (o *RelativePathToRequest) xxx_FromOp(ctx context.Context, op *xxx_RelativePathToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Other = op.Other
}
func (o *RelativePathToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RelativePathToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RelativePathToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RelativePathToResponse structure represents the RelativePathTo operation response
type RelativePathToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	RelationPath *urlmon.Moniker `idl:"name:ppmkRelPath" json:"relation_path"`
	// Return: The RelativePathTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RelativePathToResponse) xxx_ToOp(ctx context.Context, op *xxx_RelativePathToOperation) *xxx_RelativePathToOperation {
	if op == nil {
		op = &xxx_RelativePathToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RelationPath = o.RelationPath
	op.Return = o.Return
	return op
}

func (o *RelativePathToResponse) xxx_FromOp(ctx context.Context, op *xxx_RelativePathToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RelationPath = op.RelationPath
	o.Return = op.Return
}
func (o *RelativePathToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RelativePathToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RelativePathToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDisplayNameOperation structure represents the GetDisplayName operation
type xxx_GetDisplayNameOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	DisplayName string              `idl:"name:ppszDisplayName" json:"display_name"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDisplayNameOperation) OpNum() int { return 20 }

func (o *xxx_GetDisplayNameOperation) OpName() string { return "/IMoniker/v0/GetDisplayName" }

func (o *xxx_GetDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_pmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_pmkToLeft); err != nil {
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

func (o *xxx_GetDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_pmkToLeft, _ptr_pmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppszDisplayName {out} (1:{pointer=ref}*(2))(2:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DisplayName != "" {
			_ptr_ppszDisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DisplayName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DisplayName, _ptr_ppszDisplayName); err != nil {
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

func (o *xxx_GetDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppszDisplayName {out} (1:{pointer=ref}*(2))(2:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppszDisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DisplayName); err != nil {
				return err
			}
			return nil
		})
		_s_ppszDisplayName := func(ptr interface{}) { o.DisplayName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DisplayName, _s_ppszDisplayName, _ptr_ppszDisplayName); err != nil {
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

// GetDisplayNameRequest structure represents the GetDisplayName operation request
type GetDisplayNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
}

func (o *GetDisplayNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDisplayNameOperation) *xxx_GetDisplayNameOperation {
	if op == nil {
		op = &xxx_GetDisplayNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindContext = o.BindContext
	op.ToLeft = o.ToLeft
	return op
}

func (o *GetDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindContext = op.BindContext
	o.ToLeft = op.ToLeft
}
func (o *GetDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDisplayNameResponse structure represents the GetDisplayName operation response
type GetDisplayNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	DisplayName string         `idl:"name:ppszDisplayName" json:"display_name"`
	// Return: The GetDisplayName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDisplayNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDisplayNameOperation) *xxx_GetDisplayNameOperation {
	if op == nil {
		op = &xxx_GetDisplayNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DisplayName = o.DisplayName
	op.Return = o.Return
	return op
}

func (o *GetDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DisplayName = op.DisplayName
	o.Return = op.Return
}
func (o *GetDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ParseDisplayNameOperation structure represents the ParseDisplayName operation
type xxx_ParseDisplayNameOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	DisplayName string              `idl:"name:pszDisplayName" json:"display_name"`
	Eaten       uint32              `idl:"name:pchEaten" json:"eaten"`
	Out         *urlmon.Moniker     `idl:"name:ppmkOut" json:"out"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_ParseDisplayNameOperation) OpNum() int { return 21 }

func (o *xxx_ParseDisplayNameOperation) OpName() string { return "/IMoniker/v0/ParseDisplayName" }

func (o *xxx_ParseDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ParseDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ToLeft != nil {
			_ptr_pmkToLeft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ToLeft != nil {
					if err := o.ToLeft.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ToLeft, _ptr_pmkToLeft); err != nil {
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
	// pszDisplayName {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DisplayName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ParseDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkToLeft {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkToLeft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ToLeft == nil {
				o.ToLeft = &urlmon.Moniker{}
			}
			if err := o.ToLeft.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkToLeft := func(ptr interface{}) { o.ToLeft = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ToLeft, _s_pmkToLeft, _ptr_pmkToLeft); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszDisplayName {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DisplayName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ParseDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ParseDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pchEaten {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Eaten); err != nil {
			return err
		}
	}
	// ppmkOut {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Out != nil {
			_ptr_ppmkOut := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Out != nil {
					if err := o.Out.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Out, _ptr_ppmkOut); err != nil {
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

func (o *xxx_ParseDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pchEaten {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Eaten); err != nil {
			return err
		}
	}
	// ppmkOut {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppmkOut := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Out == nil {
				o.Out = &urlmon.Moniker{}
			}
			if err := o.Out.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmkOut := func(ptr interface{}) { o.Out = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Out, _s_ppmkOut, _ptr_ppmkOut); err != nil {
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

// ParseDisplayNameRequest structure represents the ParseDisplayName operation request
type ParseDisplayNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	ToLeft      *urlmon.Moniker     `idl:"name:pmkToLeft" json:"to_left"`
	DisplayName string              `idl:"name:pszDisplayName" json:"display_name"`
}

func (o *ParseDisplayNameRequest) xxx_ToOp(ctx context.Context, op *xxx_ParseDisplayNameOperation) *xxx_ParseDisplayNameOperation {
	if op == nil {
		op = &xxx_ParseDisplayNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindContext = o.BindContext
	op.ToLeft = o.ToLeft
	op.DisplayName = o.DisplayName
	return op
}

func (o *ParseDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_ParseDisplayNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindContext = op.BindContext
	o.ToLeft = op.ToLeft
	o.DisplayName = op.DisplayName
}
func (o *ParseDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ParseDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ParseDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ParseDisplayNameResponse structure represents the ParseDisplayName operation response
type ParseDisplayNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Eaten uint32          `idl:"name:pchEaten" json:"eaten"`
	Out   *urlmon.Moniker `idl:"name:ppmkOut" json:"out"`
	// Return: The ParseDisplayName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ParseDisplayNameResponse) xxx_ToOp(ctx context.Context, op *xxx_ParseDisplayNameOperation) *xxx_ParseDisplayNameOperation {
	if op == nil {
		op = &xxx_ParseDisplayNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Eaten = o.Eaten
	op.Out = o.Out
	op.Return = o.Return
	return op
}

func (o *ParseDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_ParseDisplayNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Eaten = op.Eaten
	o.Out = op.Out
	o.Return = op.Return
}
func (o *ParseDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ParseDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ParseDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsSystemMonikerOperation structure represents the IsSystemMoniker operation
type xxx_IsSystemMonikerOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MKSYS  uint32         `idl:"name:pdwMksys" json:"mksys"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsSystemMonikerOperation) OpNum() int { return 22 }

func (o *xxx_IsSystemMonikerOperation) OpName() string { return "/IMoniker/v0/IsSystemMoniker" }

func (o *xxx_IsSystemMonikerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsSystemMonikerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsSystemMonikerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsSystemMonikerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsSystemMonikerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwMksys {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MKSYS); err != nil {
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

func (o *xxx_IsSystemMonikerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwMksys {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MKSYS); err != nil {
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

// IsSystemMonikerRequest structure represents the IsSystemMoniker operation request
type IsSystemMonikerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsSystemMonikerRequest) xxx_ToOp(ctx context.Context, op *xxx_IsSystemMonikerOperation) *xxx_IsSystemMonikerOperation {
	if op == nil {
		op = &xxx_IsSystemMonikerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *IsSystemMonikerRequest) xxx_FromOp(ctx context.Context, op *xxx_IsSystemMonikerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsSystemMonikerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsSystemMonikerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsSystemMonikerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsSystemMonikerResponse structure represents the IsSystemMoniker operation response
type IsSystemMonikerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	MKSYS uint32         `idl:"name:pdwMksys" json:"mksys"`
	// Return: The IsSystemMoniker return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsSystemMonikerResponse) xxx_ToOp(ctx context.Context, op *xxx_IsSystemMonikerOperation) *xxx_IsSystemMonikerOperation {
	if op == nil {
		op = &xxx_IsSystemMonikerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MKSYS = o.MKSYS
	op.Return = o.Return
	return op
}

func (o *IsSystemMonikerResponse) xxx_FromOp(ctx context.Context, op *xxx_IsSystemMonikerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MKSYS = op.MKSYS
	o.Return = op.Return
}
func (o *IsSystemMonikerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsSystemMonikerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsSystemMonikerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

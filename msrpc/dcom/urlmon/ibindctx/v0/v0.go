package ibindctx

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
	urlmon "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon"
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
	_ = urlmon.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/urlmon"
)

var (
	// IBindCtx interface identifier 0000000e-0000-0000-c000-000000000046
	BindContextIID = &dcom.IID{Data1: 0x0000000e, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	BindContextSyntaxUUID = &uuid.UUID{TimeLow: 0xe, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	BindContextSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: BindContextSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IBindCtx interface.
type BindContextClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// RegisterObjectBound operation.
	RegisterObjectBound(context.Context, *RegisterObjectBoundRequest, ...dcerpc.CallOption) (*RegisterObjectBoundResponse, error)

	// RevokeObjectBound operation.
	RevokeObjectBound(context.Context, *RevokeObjectBoundRequest, ...dcerpc.CallOption) (*RevokeObjectBoundResponse, error)

	// ReleaseBoundObjects operation.
	ReleaseBoundObjects(context.Context, *ReleaseBoundObjectsRequest, ...dcerpc.CallOption) (*ReleaseBoundObjectsResponse, error)

	// SetBindOptions operation.
	SetBindOptions(context.Context, *SetBindOptionsRequest, ...dcerpc.CallOption) (*SetBindOptionsResponse, error)

	// GetBindOptions operation.
	GetBindOptions(context.Context, *GetBindOptionsRequest, ...dcerpc.CallOption) (*GetBindOptionsResponse, error)

	// GetRunningObjectTable operation.
	GetRunningObjectTable(context.Context, *GetRunningObjectTableRequest, ...dcerpc.CallOption) (*GetRunningObjectTableResponse, error)

	// RegisterObjectParam operation.
	RegisterObjectParam(context.Context, *RegisterObjectParamRequest, ...dcerpc.CallOption) (*RegisterObjectParamResponse, error)

	// GetObjectParam operation.
	GetObjectParam(context.Context, *GetObjectParamRequest, ...dcerpc.CallOption) (*GetObjectParamResponse, error)

	// EnumObjectParam operation.
	EnumObjectParam(context.Context, *EnumObjectParamRequest, ...dcerpc.CallOption) (*EnumObjectParamResponse, error)

	// RevokeObjectParam operation.
	RevokeObjectParam(context.Context, *RevokeObjectParamRequest, ...dcerpc.CallOption) (*RevokeObjectParamResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) BindContextClient
}

type xxx_DefaultBindContextClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultBindContextClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultBindContextClient) RegisterObjectBound(ctx context.Context, in *RegisterObjectBoundRequest, opts ...dcerpc.CallOption) (*RegisterObjectBoundResponse, error) {
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
	out := &RegisterObjectBoundResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) RevokeObjectBound(ctx context.Context, in *RevokeObjectBoundRequest, opts ...dcerpc.CallOption) (*RevokeObjectBoundResponse, error) {
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
	out := &RevokeObjectBoundResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) ReleaseBoundObjects(ctx context.Context, in *ReleaseBoundObjectsRequest, opts ...dcerpc.CallOption) (*ReleaseBoundObjectsResponse, error) {
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
	out := &ReleaseBoundObjectsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) SetBindOptions(ctx context.Context, in *SetBindOptionsRequest, opts ...dcerpc.CallOption) (*SetBindOptionsResponse, error) {
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
	out := &SetBindOptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) GetBindOptions(ctx context.Context, in *GetBindOptionsRequest, opts ...dcerpc.CallOption) (*GetBindOptionsResponse, error) {
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
	out := &GetBindOptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) GetRunningObjectTable(ctx context.Context, in *GetRunningObjectTableRequest, opts ...dcerpc.CallOption) (*GetRunningObjectTableResponse, error) {
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
	out := &GetRunningObjectTableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) RegisterObjectParam(ctx context.Context, in *RegisterObjectParamRequest, opts ...dcerpc.CallOption) (*RegisterObjectParamResponse, error) {
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
	out := &RegisterObjectParamResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) GetObjectParam(ctx context.Context, in *GetObjectParamRequest, opts ...dcerpc.CallOption) (*GetObjectParamResponse, error) {
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
	out := &GetObjectParamResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) EnumObjectParam(ctx context.Context, in *EnumObjectParamRequest, opts ...dcerpc.CallOption) (*EnumObjectParamResponse, error) {
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
	out := &EnumObjectParamResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) RevokeObjectParam(ctx context.Context, in *RevokeObjectParamRequest, opts ...dcerpc.CallOption) (*RevokeObjectParamResponse, error) {
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
	out := &RevokeObjectParamResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBindContextClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultBindContextClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultBindContextClient) IPID(ctx context.Context, ipid *dcom.IPID) BindContextClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultBindContextClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewBindContextClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (BindContextClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(BindContextSyntaxV0_0))...)
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
	return &xxx_DefaultBindContextClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RegisterObjectBoundOperation structure represents the RegisterObjectBound operation
type xxx_RegisterObjectBoundOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Unknown *dcom.Unknown  `idl:"name:punk" json:"unknown"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterObjectBoundOperation) OpNum() int { return 3 }

func (o *xxx_RegisterObjectBoundOperation) OpName() string { return "/IBindCtx/v0/RegisterObjectBound" }

func (o *xxx_RegisterObjectBoundOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterObjectBoundOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// punk {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Unknown != nil {
			_ptr_punk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Unknown != nil {
					if err := o.Unknown.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Unknown, _ptr_punk); err != nil {
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

func (o *xxx_RegisterObjectBoundOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// punk {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_punk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Unknown == nil {
				o.Unknown = &dcom.Unknown{}
			}
			if err := o.Unknown.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_punk := func(ptr interface{}) { o.Unknown = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Unknown, _s_punk, _ptr_punk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterObjectBoundOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterObjectBoundOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RegisterObjectBoundOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RegisterObjectBoundRequest structure represents the RegisterObjectBound operation request
type RegisterObjectBoundRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Unknown *dcom.Unknown  `idl:"name:punk" json:"unknown"`
}

func (o *RegisterObjectBoundRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterObjectBoundOperation) *xxx_RegisterObjectBoundOperation {
	if op == nil {
		op = &xxx_RegisterObjectBoundOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Unknown = o.Unknown
	return op
}

func (o *RegisterObjectBoundRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterObjectBoundOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Unknown = op.Unknown
}
func (o *RegisterObjectBoundRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterObjectBoundRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterObjectBoundOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterObjectBoundResponse structure represents the RegisterObjectBound operation response
type RegisterObjectBoundResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RegisterObjectBound return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterObjectBoundResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterObjectBoundOperation) *xxx_RegisterObjectBoundOperation {
	if op == nil {
		op = &xxx_RegisterObjectBoundOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RegisterObjectBoundResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterObjectBoundOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RegisterObjectBoundResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterObjectBoundResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterObjectBoundOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RevokeObjectBoundOperation structure represents the RevokeObjectBound operation
type xxx_RevokeObjectBoundOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Unknown *dcom.Unknown  `idl:"name:punk" json:"unknown"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RevokeObjectBoundOperation) OpNum() int { return 4 }

func (o *xxx_RevokeObjectBoundOperation) OpName() string { return "/IBindCtx/v0/RevokeObjectBound" }

func (o *xxx_RevokeObjectBoundOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeObjectBoundOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// punk {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Unknown != nil {
			_ptr_punk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Unknown != nil {
					if err := o.Unknown.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Unknown, _ptr_punk); err != nil {
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

func (o *xxx_RevokeObjectBoundOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// punk {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_punk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Unknown == nil {
				o.Unknown = &dcom.Unknown{}
			}
			if err := o.Unknown.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_punk := func(ptr interface{}) { o.Unknown = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Unknown, _s_punk, _ptr_punk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeObjectBoundOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeObjectBoundOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RevokeObjectBoundOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RevokeObjectBoundRequest structure represents the RevokeObjectBound operation request
type RevokeObjectBoundRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Unknown *dcom.Unknown  `idl:"name:punk" json:"unknown"`
}

func (o *RevokeObjectBoundRequest) xxx_ToOp(ctx context.Context, op *xxx_RevokeObjectBoundOperation) *xxx_RevokeObjectBoundOperation {
	if op == nil {
		op = &xxx_RevokeObjectBoundOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Unknown = o.Unknown
	return op
}

func (o *RevokeObjectBoundRequest) xxx_FromOp(ctx context.Context, op *xxx_RevokeObjectBoundOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Unknown = op.Unknown
}
func (o *RevokeObjectBoundRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RevokeObjectBoundRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeObjectBoundOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RevokeObjectBoundResponse structure represents the RevokeObjectBound operation response
type RevokeObjectBoundResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RevokeObjectBound return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RevokeObjectBoundResponse) xxx_ToOp(ctx context.Context, op *xxx_RevokeObjectBoundOperation) *xxx_RevokeObjectBoundOperation {
	if op == nil {
		op = &xxx_RevokeObjectBoundOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RevokeObjectBoundResponse) xxx_FromOp(ctx context.Context, op *xxx_RevokeObjectBoundOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RevokeObjectBoundResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RevokeObjectBoundResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeObjectBoundOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReleaseBoundObjectsOperation structure represents the ReleaseBoundObjects operation
type xxx_ReleaseBoundObjectsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReleaseBoundObjectsOperation) OpNum() int { return 5 }

func (o *xxx_ReleaseBoundObjectsOperation) OpName() string { return "/IBindCtx/v0/ReleaseBoundObjects" }

func (o *xxx_ReleaseBoundObjectsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReleaseBoundObjectsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReleaseBoundObjectsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ReleaseBoundObjectsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReleaseBoundObjectsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReleaseBoundObjectsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ReleaseBoundObjectsRequest structure represents the ReleaseBoundObjects operation request
type ReleaseBoundObjectsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ReleaseBoundObjectsRequest) xxx_ToOp(ctx context.Context, op *xxx_ReleaseBoundObjectsOperation) *xxx_ReleaseBoundObjectsOperation {
	if op == nil {
		op = &xxx_ReleaseBoundObjectsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ReleaseBoundObjectsRequest) xxx_FromOp(ctx context.Context, op *xxx_ReleaseBoundObjectsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ReleaseBoundObjectsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReleaseBoundObjectsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReleaseBoundObjectsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReleaseBoundObjectsResponse structure represents the ReleaseBoundObjects operation response
type ReleaseBoundObjectsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReleaseBoundObjects return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReleaseBoundObjectsResponse) xxx_ToOp(ctx context.Context, op *xxx_ReleaseBoundObjectsOperation) *xxx_ReleaseBoundObjectsOperation {
	if op == nil {
		op = &xxx_ReleaseBoundObjectsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ReleaseBoundObjectsResponse) xxx_FromOp(ctx context.Context, op *xxx_ReleaseBoundObjectsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ReleaseBoundObjectsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReleaseBoundObjectsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReleaseBoundObjectsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetBindOptionsOperation structure represents the SetBindOptions operation
type xxx_SetBindOptionsOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindOptions *urlmon.BindOptions `idl:"name:pbindopts" json:"bind_options"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_SetBindOptionsOperation) OpNum() int { return 6 }

func (o *xxx_SetBindOptionsOperation) OpName() string { return "/IBindCtx/v0/SetBindOptions" }

func (o *xxx_SetBindOptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBindOptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbindopts {in} (1:{pointer=ref}*(1))(2:{alias=BIND_OPTS}(struct))
	{
		if o.BindOptions != nil {
			if err := o.BindOptions.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&urlmon.BindOptions{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetBindOptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbindopts {in} (1:{pointer=ref}*(1))(2:{alias=BIND_OPTS}(struct))
	{
		if o.BindOptions == nil {
			o.BindOptions = &urlmon.BindOptions{}
		}
		if err := o.BindOptions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBindOptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBindOptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetBindOptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetBindOptionsRequest structure represents the SetBindOptions operation request
type SetBindOptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindOptions *urlmon.BindOptions `idl:"name:pbindopts" json:"bind_options"`
}

func (o *SetBindOptionsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetBindOptionsOperation) *xxx_SetBindOptionsOperation {
	if op == nil {
		op = &xxx_SetBindOptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindOptions = o.BindOptions
	return op
}

func (o *SetBindOptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetBindOptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindOptions = op.BindOptions
}
func (o *SetBindOptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetBindOptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBindOptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetBindOptionsResponse structure represents the SetBindOptions operation response
type SetBindOptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetBindOptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetBindOptionsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetBindOptionsOperation) *xxx_SetBindOptionsOperation {
	if op == nil {
		op = &xxx_SetBindOptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetBindOptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetBindOptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetBindOptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetBindOptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBindOptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBindOptionsOperation structure represents the GetBindOptions operation
type xxx_GetBindOptionsOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindOptions *urlmon.BindOptions `idl:"name:pbindopts" json:"bind_options"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBindOptionsOperation) OpNum() int { return 7 }

func (o *xxx_GetBindOptionsOperation) OpName() string { return "/IBindCtx/v0/GetBindOptions" }

func (o *xxx_GetBindOptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBindOptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pbindopts {in, out} (1:{pointer=ref}*(1))(2:{alias=BIND_OPTS}(struct))
	{
		if o.BindOptions != nil {
			if err := o.BindOptions.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&urlmon.BindOptions{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetBindOptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pbindopts {in, out} (1:{pointer=ref}*(1))(2:{alias=BIND_OPTS}(struct))
	{
		if o.BindOptions == nil {
			o.BindOptions = &urlmon.BindOptions{}
		}
		if err := o.BindOptions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBindOptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBindOptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbindopts {in, out} (1:{pointer=ref}*(1))(2:{alias=BIND_OPTS}(struct))
	{
		if o.BindOptions != nil {
			if err := o.BindOptions.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&urlmon.BindOptions{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetBindOptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbindopts {in, out} (1:{pointer=ref}*(1))(2:{alias=BIND_OPTS}(struct))
	{
		if o.BindOptions == nil {
			o.BindOptions = &urlmon.BindOptions{}
		}
		if err := o.BindOptions.UnmarshalNDR(ctx, w); err != nil {
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

// GetBindOptionsRequest structure represents the GetBindOptions operation request
type GetBindOptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	BindOptions *urlmon.BindOptions `idl:"name:pbindopts" json:"bind_options"`
}

func (o *GetBindOptionsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBindOptionsOperation) *xxx_GetBindOptionsOperation {
	if op == nil {
		op = &xxx_GetBindOptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BindOptions = o.BindOptions
	return op
}

func (o *GetBindOptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBindOptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BindOptions = op.BindOptions
}
func (o *GetBindOptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBindOptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBindOptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBindOptionsResponse structure represents the GetBindOptions operation response
type GetBindOptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	BindOptions *urlmon.BindOptions `idl:"name:pbindopts" json:"bind_options"`
	// Return: The GetBindOptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBindOptionsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBindOptionsOperation) *xxx_GetBindOptionsOperation {
	if op == nil {
		op = &xxx_GetBindOptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BindOptions = o.BindOptions
	op.Return = o.Return
	return op
}

func (o *GetBindOptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBindOptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BindOptions = op.BindOptions
	o.Return = op.Return
}
func (o *GetBindOptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBindOptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBindOptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRunningObjectTableOperation structure represents the GetRunningObjectTable operation
type xxx_GetRunningObjectTableOperation struct {
	This               *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat             `idl:"name:That" json:"that"`
	RunningObjectTable *urlmon.RunningObjectTable `idl:"name:pprot" json:"running_object_table"`
	Return             int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRunningObjectTableOperation) OpNum() int { return 8 }

func (o *xxx_GetRunningObjectTableOperation) OpName() string {
	return "/IBindCtx/v0/GetRunningObjectTable"
}

func (o *xxx_GetRunningObjectTableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRunningObjectTableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRunningObjectTableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRunningObjectTableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRunningObjectTableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pprot {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IRunningObjectTable}(interface))
	{
		if o.RunningObjectTable != nil {
			_ptr_pprot := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RunningObjectTable != nil {
					if err := o.RunningObjectTable.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.RunningObjectTable{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RunningObjectTable, _ptr_pprot); err != nil {
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

func (o *xxx_GetRunningObjectTableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pprot {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IRunningObjectTable}(interface))
	{
		_ptr_pprot := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RunningObjectTable == nil {
				o.RunningObjectTable = &urlmon.RunningObjectTable{}
			}
			if err := o.RunningObjectTable.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pprot := func(ptr interface{}) { o.RunningObjectTable = *ptr.(**urlmon.RunningObjectTable) }
		if err := w.ReadPointer(&o.RunningObjectTable, _s_pprot, _ptr_pprot); err != nil {
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

// GetRunningObjectTableRequest structure represents the GetRunningObjectTable operation request
type GetRunningObjectTableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRunningObjectTableRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRunningObjectTableOperation) *xxx_GetRunningObjectTableOperation {
	if op == nil {
		op = &xxx_GetRunningObjectTableOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetRunningObjectTableRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRunningObjectTableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRunningObjectTableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRunningObjectTableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRunningObjectTableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRunningObjectTableResponse structure represents the GetRunningObjectTable operation response
type GetRunningObjectTableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat             `idl:"name:That" json:"that"`
	RunningObjectTable *urlmon.RunningObjectTable `idl:"name:pprot" json:"running_object_table"`
	// Return: The GetRunningObjectTable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRunningObjectTableResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRunningObjectTableOperation) *xxx_GetRunningObjectTableOperation {
	if op == nil {
		op = &xxx_GetRunningObjectTableOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RunningObjectTable = o.RunningObjectTable
	op.Return = o.Return
	return op
}

func (o *GetRunningObjectTableResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRunningObjectTableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RunningObjectTable = op.RunningObjectTable
	o.Return = op.Return
}
func (o *GetRunningObjectTableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRunningObjectTableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRunningObjectTableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterObjectParamOperation structure represents the RegisterObjectParam operation
type xxx_RegisterObjectParamOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Key     string         `idl:"name:pszKey" json:"key"`
	Unknown *dcom.Unknown  `idl:"name:punk" json:"unknown"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterObjectParamOperation) OpNum() int { return 9 }

func (o *xxx_RegisterObjectParamOperation) OpName() string { return "/IBindCtx/v0/RegisterObjectParam" }

func (o *xxx_RegisterObjectParamOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterObjectParamOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszKey {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Key); err != nil {
			return err
		}
	}
	// punk {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Unknown != nil {
			_ptr_punk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Unknown != nil {
					if err := o.Unknown.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Unknown, _ptr_punk); err != nil {
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

func (o *xxx_RegisterObjectParamOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszKey {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Key); err != nil {
			return err
		}
	}
	// punk {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_punk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Unknown == nil {
				o.Unknown = &dcom.Unknown{}
			}
			if err := o.Unknown.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_punk := func(ptr interface{}) { o.Unknown = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Unknown, _s_punk, _ptr_punk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterObjectParamOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterObjectParamOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RegisterObjectParamOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RegisterObjectParamRequest structure represents the RegisterObjectParam operation request
type RegisterObjectParamRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Key     string         `idl:"name:pszKey" json:"key"`
	Unknown *dcom.Unknown  `idl:"name:punk" json:"unknown"`
}

func (o *RegisterObjectParamRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterObjectParamOperation) *xxx_RegisterObjectParamOperation {
	if op == nil {
		op = &xxx_RegisterObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Key = o.Key
	op.Unknown = o.Unknown
	return op
}

func (o *RegisterObjectParamRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterObjectParamOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Key = op.Key
	o.Unknown = op.Unknown
}
func (o *RegisterObjectParamRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterObjectParamRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterObjectParamOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterObjectParamResponse structure represents the RegisterObjectParam operation response
type RegisterObjectParamResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RegisterObjectParam return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterObjectParamResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterObjectParamOperation) *xxx_RegisterObjectParamOperation {
	if op == nil {
		op = &xxx_RegisterObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RegisterObjectParamResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterObjectParamOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RegisterObjectParamResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterObjectParamResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterObjectParamOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectParamOperation structure represents the GetObjectParam operation
type xxx_GetObjectParamOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Key     string         `idl:"name:pszKey" json:"key"`
	Unknown *dcom.Unknown  `idl:"name:ppunk" json:"unknown"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectParamOperation) OpNum() int { return 10 }

func (o *xxx_GetObjectParamOperation) OpName() string { return "/IBindCtx/v0/GetObjectParam" }

func (o *xxx_GetObjectParamOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectParamOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszKey {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectParamOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszKey {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectParamOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectParamOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppunk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Unknown != nil {
			_ptr_ppunk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Unknown != nil {
					if err := o.Unknown.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Unknown, _ptr_ppunk); err != nil {
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

func (o *xxx_GetObjectParamOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppunk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppunk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Unknown == nil {
				o.Unknown = &dcom.Unknown{}
			}
			if err := o.Unknown.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppunk := func(ptr interface{}) { o.Unknown = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Unknown, _s_ppunk, _ptr_ppunk); err != nil {
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

// GetObjectParamRequest structure represents the GetObjectParam operation request
type GetObjectParamRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Key  string         `idl:"name:pszKey" json:"key"`
}

func (o *GetObjectParamRequest) xxx_ToOp(ctx context.Context, op *xxx_GetObjectParamOperation) *xxx_GetObjectParamOperation {
	if op == nil {
		op = &xxx_GetObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Key = o.Key
	return op
}

func (o *GetObjectParamRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectParamOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Key = op.Key
}
func (o *GetObjectParamRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectParamRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectParamOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectParamResponse structure represents the GetObjectParam operation response
type GetObjectParamResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Unknown *dcom.Unknown  `idl:"name:ppunk" json:"unknown"`
	// Return: The GetObjectParam return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectParamResponse) xxx_ToOp(ctx context.Context, op *xxx_GetObjectParamOperation) *xxx_GetObjectParamOperation {
	if op == nil {
		op = &xxx_GetObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Unknown = o.Unknown
	op.Return = o.Return
	return op
}

func (o *GetObjectParamResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectParamOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Unknown = op.Unknown
	o.Return = op.Return
}
func (o *GetObjectParamResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectParamResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectParamOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumObjectParamOperation structure represents the EnumObjectParam operation
type xxx_EnumObjectParamOperation struct {
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Enum   *urlmon.EnumString `idl:"name:ppenum" json:"enum"`
	Return int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumObjectParamOperation) OpNum() int { return 11 }

func (o *xxx_EnumObjectParamOperation) OpName() string { return "/IBindCtx/v0/EnumObjectParam" }

func (o *xxx_EnumObjectParamOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumObjectParamOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumObjectParamOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_EnumObjectParamOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumObjectParamOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppenum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumString}(interface))
	{
		if o.Enum != nil {
			_ptr_ppenum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.EnumString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppenum); err != nil {
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

func (o *xxx_EnumObjectParamOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppenum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumString}(interface))
	{
		_ptr_ppenum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &urlmon.EnumString{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppenum := func(ptr interface{}) { o.Enum = *ptr.(**urlmon.EnumString) }
		if err := w.ReadPointer(&o.Enum, _s_ppenum, _ptr_ppenum); err != nil {
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

// EnumObjectParamRequest structure represents the EnumObjectParam operation request
type EnumObjectParamRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EnumObjectParamRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumObjectParamOperation) *xxx_EnumObjectParamOperation {
	if op == nil {
		op = &xxx_EnumObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EnumObjectParamRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumObjectParamOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EnumObjectParamRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumObjectParamRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumObjectParamOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumObjectParamResponse structure represents the EnumObjectParam operation response
type EnumObjectParamResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Enum *urlmon.EnumString `idl:"name:ppenum" json:"enum"`
	// Return: The EnumObjectParam return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumObjectParamResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumObjectParamOperation) *xxx_EnumObjectParamOperation {
	if op == nil {
		op = &xxx_EnumObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *EnumObjectParamResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumObjectParamOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *EnumObjectParamResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumObjectParamResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumObjectParamOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RevokeObjectParamOperation structure represents the RevokeObjectParam operation
type xxx_RevokeObjectParamOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Key    string         `idl:"name:pszKey" json:"key"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RevokeObjectParamOperation) OpNum() int { return 12 }

func (o *xxx_RevokeObjectParamOperation) OpName() string { return "/IBindCtx/v0/RevokeObjectParam" }

func (o *xxx_RevokeObjectParamOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeObjectParamOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszKey {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeObjectParamOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszKey {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeObjectParamOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeObjectParamOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RevokeObjectParamOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RevokeObjectParamRequest structure represents the RevokeObjectParam operation request
type RevokeObjectParamRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Key  string         `idl:"name:pszKey" json:"key"`
}

func (o *RevokeObjectParamRequest) xxx_ToOp(ctx context.Context, op *xxx_RevokeObjectParamOperation) *xxx_RevokeObjectParamOperation {
	if op == nil {
		op = &xxx_RevokeObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Key = o.Key
	return op
}

func (o *RevokeObjectParamRequest) xxx_FromOp(ctx context.Context, op *xxx_RevokeObjectParamOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Key = op.Key
}
func (o *RevokeObjectParamRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RevokeObjectParamRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeObjectParamOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RevokeObjectParamResponse structure represents the RevokeObjectParam operation response
type RevokeObjectParamResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RevokeObjectParam return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RevokeObjectParamResponse) xxx_ToOp(ctx context.Context, op *xxx_RevokeObjectParamOperation) *xxx_RevokeObjectParamOperation {
	if op == nil {
		op = &xxx_RevokeObjectParamOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RevokeObjectParamResponse) xxx_FromOp(ctx context.Context, op *xxx_RevokeObjectParamOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RevokeObjectParamResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RevokeObjectParamResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeObjectParamOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

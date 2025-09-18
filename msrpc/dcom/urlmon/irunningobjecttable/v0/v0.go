package irunningobjecttable

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
	_ = urlmon.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/urlmon"
)

var (
	// IRunningObjectTable interface identifier 00000010-0000-0000-c000-000000000046
	RunningObjectTableIID = &dcom.IID{Data1: 0x00000010, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	RunningObjectTableSyntaxUUID = &uuid.UUID{TimeLow: 0x10, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	RunningObjectTableSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RunningObjectTableSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRunningObjectTable interface.
type RunningObjectTableClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Register operation.
	Register(context.Context, *RegisterRequest, ...dcerpc.CallOption) (*RegisterResponse, error)

	// Revoke operation.
	Revoke(context.Context, *RevokeRequest, ...dcerpc.CallOption) (*RevokeResponse, error)

	// IsRunning operation.
	IsRunning(context.Context, *IsRunningRequest, ...dcerpc.CallOption) (*IsRunningResponse, error)

	// GetObject operation.
	GetObject(context.Context, *GetObjectRequest, ...dcerpc.CallOption) (*GetObjectResponse, error)

	// NoteChangeTime operation.
	NoteChangeTime(context.Context, *NoteChangeTimeRequest, ...dcerpc.CallOption) (*NoteChangeTimeResponse, error)

	// GetTimeOfLastChange operation.
	GetTimeOfLastChange(context.Context, *GetTimeOfLastChangeRequest, ...dcerpc.CallOption) (*GetTimeOfLastChangeResponse, error)

	// EnumRunning operation.
	EnumRunning(context.Context, *EnumRunningRequest, ...dcerpc.CallOption) (*EnumRunningResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RunningObjectTableClient
}

type xxx_DefaultRunningObjectTableClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRunningObjectTableClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRunningObjectTableClient) Register(ctx context.Context, in *RegisterRequest, opts ...dcerpc.CallOption) (*RegisterResponse, error) {
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
	out := &RegisterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRunningObjectTableClient) Revoke(ctx context.Context, in *RevokeRequest, opts ...dcerpc.CallOption) (*RevokeResponse, error) {
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
	out := &RevokeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRunningObjectTableClient) IsRunning(ctx context.Context, in *IsRunningRequest, opts ...dcerpc.CallOption) (*IsRunningResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRunningObjectTableClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...dcerpc.CallOption) (*GetObjectResponse, error) {
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
	out := &GetObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRunningObjectTableClient) NoteChangeTime(ctx context.Context, in *NoteChangeTimeRequest, opts ...dcerpc.CallOption) (*NoteChangeTimeResponse, error) {
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
	out := &NoteChangeTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRunningObjectTableClient) GetTimeOfLastChange(ctx context.Context, in *GetTimeOfLastChangeRequest, opts ...dcerpc.CallOption) (*GetTimeOfLastChangeResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRunningObjectTableClient) EnumRunning(ctx context.Context, in *EnumRunningRequest, opts ...dcerpc.CallOption) (*EnumRunningResponse, error) {
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
	out := &EnumRunningResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRunningObjectTableClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRunningObjectTableClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRunningObjectTableClient) IPID(ctx context.Context, ipid *dcom.IPID) RunningObjectTableClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRunningObjectTableClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRunningObjectTableClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RunningObjectTableClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RunningObjectTableSyntaxV0_0))...)
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
	return &xxx_DefaultRunningObjectTableClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RegisterOperation structure represents the Register operation
type xxx_RegisterOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Flags      uint32          `idl:"name:grfFlags" json:"flags"`
	Object     *dcom.Unknown   `idl:"name:punkObject" json:"object"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
	Register   uint32          `idl:"name:pdwRegister" json:"register"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterOperation) OpNum() int { return 3 }

func (o *xxx_RegisterOperation) OpName() string { return "/IRunningObjectTable/v0/Register" }

func (o *xxx_RegisterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// grfFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// punkObject {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Object != nil {
			_ptr_punkObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_punkObject); err != nil {
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
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ObjectName != nil {
			_ptr_pmkObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectName != nil {
					if err := o.ObjectName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectName, _ptr_pmkObjectName); err != nil {
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

func (o *xxx_RegisterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// grfFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// punkObject {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_punkObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &dcom.Unknown{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_punkObject := func(ptr interface{}) { o.Object = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Object, _s_punkObject, _ptr_punkObject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectName == nil {
				o.ObjectName = &urlmon.Moniker{}
			}
			if err := o.ObjectName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ObjectName, _s_pmkObjectName, _ptr_pmkObjectName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwRegister {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Register); err != nil {
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

func (o *xxx_RegisterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwRegister {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Register); err != nil {
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

// RegisterRequest structure represents the Register operation request
type RegisterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Flags      uint32          `idl:"name:grfFlags" json:"flags"`
	Object     *dcom.Unknown   `idl:"name:punkObject" json:"object"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
}

func (o *RegisterRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterOperation) *xxx_RegisterOperation {
	if op == nil {
		op = &xxx_RegisterOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Flags = o.Flags
	op.Object = o.Object
	op.ObjectName = o.ObjectName
	return op
}

func (o *RegisterRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.Object = op.Object
	o.ObjectName = op.ObjectName
}
func (o *RegisterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterResponse structure represents the Register operation response
type RegisterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Register uint32         `idl:"name:pdwRegister" json:"register"`
	// Return: The Register return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterOperation) *xxx_RegisterOperation {
	if op == nil {
		op = &xxx_RegisterOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Register = o.Register
	op.Return = o.Return
	return op
}

func (o *RegisterResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Register = op.Register
	o.Return = op.Return
}
func (o *RegisterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RevokeOperation structure represents the Revoke operation
type xxx_RevokeOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Register uint32         `idl:"name:dwRegister" json:"register"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RevokeOperation) OpNum() int { return 4 }

func (o *xxx_RevokeOperation) OpName() string { return "/IRunningObjectTable/v0/Revoke" }

func (o *xxx_RevokeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Register); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Register); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RevokeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RevokeRequest structure represents the Revoke operation request
type RevokeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Register uint32         `idl:"name:dwRegister" json:"register"`
}

func (o *RevokeRequest) xxx_ToOp(ctx context.Context, op *xxx_RevokeOperation) *xxx_RevokeOperation {
	if op == nil {
		op = &xxx_RevokeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Register = o.Register
	return op
}

func (o *RevokeRequest) xxx_FromOp(ctx context.Context, op *xxx_RevokeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Register = op.Register
}
func (o *RevokeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RevokeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RevokeResponse structure represents the Revoke operation response
type RevokeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Revoke return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RevokeResponse) xxx_ToOp(ctx context.Context, op *xxx_RevokeOperation) *xxx_RevokeOperation {
	if op == nil {
		op = &xxx_RevokeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RevokeResponse) xxx_FromOp(ctx context.Context, op *xxx_RevokeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RevokeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RevokeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsRunningOperation structure represents the IsRunning operation
type xxx_IsRunningOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_IsRunningOperation) OpNum() int { return 5 }

func (o *xxx_IsRunningOperation) OpName() string { return "/IRunningObjectTable/v0/IsRunning" }

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
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ObjectName != nil {
			_ptr_pmkObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectName != nil {
					if err := o.ObjectName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectName, _ptr_pmkObjectName); err != nil {
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
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectName == nil {
				o.ObjectName = &urlmon.Moniker{}
			}
			if err := o.ObjectName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ObjectName, _s_pmkObjectName, _ptr_pmkObjectName); err != nil {
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
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
}

func (o *IsRunningRequest) xxx_ToOp(ctx context.Context, op *xxx_IsRunningOperation) *xxx_IsRunningOperation {
	if op == nil {
		op = &xxx_IsRunningOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectName = o.ObjectName
	return op
}

func (o *IsRunningRequest) xxx_FromOp(ctx context.Context, op *xxx_IsRunningOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectName = op.ObjectName
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

// xxx_GetObjectOperation structure represents the GetObject operation
type xxx_GetObjectOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
	Object     *dcom.Unknown   `idl:"name:ppunkObject" json:"object"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectOperation) OpNum() int { return 6 }

func (o *xxx_GetObjectOperation) OpName() string { return "/IRunningObjectTable/v0/GetObject" }

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
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ObjectName != nil {
			_ptr_pmkObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectName != nil {
					if err := o.ObjectName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectName, _ptr_pmkObjectName); err != nil {
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
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectName == nil {
				o.ObjectName = &urlmon.Moniker{}
			}
			if err := o.ObjectName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ObjectName, _s_pmkObjectName, _ptr_pmkObjectName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
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
	// ppunkObject {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Object != nil {
			_ptr_ppunkObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_ppunkObject); err != nil {
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
	// ppunkObject {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppunkObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &dcom.Unknown{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppunkObject := func(ptr interface{}) { o.Object = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Object, _s_ppunkObject, _ptr_ppunkObject); err != nil {
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
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
}

func (o *GetObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_GetObjectOperation) *xxx_GetObjectOperation {
	if op == nil {
		op = &xxx_GetObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectName = o.ObjectName
	return op
}

func (o *GetObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectName = op.ObjectName
}
func (o *GetObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Object *dcom.Unknown  `idl:"name:ppunkObject" json:"object"`
	// Return: The GetObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_GetObjectOperation) *xxx_GetObjectOperation {
	if op == nil {
		op = &xxx_GetObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Object = o.Object
	op.Return = o.Return
	return op
}

func (o *GetObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Object = op.Object
	o.Return = op.Return
}
func (o *GetObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NoteChangeTimeOperation structure represents the NoteChangeTime operation
type xxx_NoteChangeTimeOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Register uint32         `idl:"name:dwRegister" json:"register"`
	Filetime *dtyp.Filetime `idl:"name:pfiletime" json:"filetime"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_NoteChangeTimeOperation) OpNum() int { return 7 }

func (o *xxx_NoteChangeTimeOperation) OpName() string {
	return "/IRunningObjectTable/v0/NoteChangeTime"
}

func (o *xxx_NoteChangeTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NoteChangeTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Register); err != nil {
			return err
		}
	}
	// pfiletime {in} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.Filetime != nil {
			if err := o.Filetime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_NoteChangeTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Register); err != nil {
			return err
		}
	}
	// pfiletime {in} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.Filetime == nil {
			o.Filetime = &dtyp.Filetime{}
		}
		if err := o.Filetime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NoteChangeTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NoteChangeTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_NoteChangeTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// NoteChangeTimeRequest structure represents the NoteChangeTime operation request
type NoteChangeTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Register uint32         `idl:"name:dwRegister" json:"register"`
	Filetime *dtyp.Filetime `idl:"name:pfiletime" json:"filetime"`
}

func (o *NoteChangeTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_NoteChangeTimeOperation) *xxx_NoteChangeTimeOperation {
	if op == nil {
		op = &xxx_NoteChangeTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Register = o.Register
	op.Filetime = o.Filetime
	return op
}

func (o *NoteChangeTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_NoteChangeTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Register = op.Register
	o.Filetime = op.Filetime
}
func (o *NoteChangeTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NoteChangeTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NoteChangeTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NoteChangeTimeResponse structure represents the NoteChangeTime operation response
type NoteChangeTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The NoteChangeTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NoteChangeTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_NoteChangeTimeOperation) *xxx_NoteChangeTimeOperation {
	if op == nil {
		op = &xxx_NoteChangeTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *NoteChangeTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_NoteChangeTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *NoteChangeTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NoteChangeTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NoteChangeTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTimeOfLastChangeOperation structure represents the GetTimeOfLastChange operation
type xxx_GetTimeOfLastChangeOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
	Filetime   *dtyp.Filetime  `idl:"name:pfiletime" json:"filetime"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTimeOfLastChangeOperation) OpNum() int { return 8 }

func (o *xxx_GetTimeOfLastChangeOperation) OpName() string {
	return "/IRunningObjectTable/v0/GetTimeOfLastChange"
}

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
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.ObjectName != nil {
			_ptr_pmkObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectName != nil {
					if err := o.ObjectName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectName, _ptr_pmkObjectName); err != nil {
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
	// pmkObjectName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pmkObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectName == nil {
				o.ObjectName = &urlmon.Moniker{}
			}
			if err := o.ObjectName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pmkObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.ObjectName, _s_pmkObjectName, _ptr_pmkObjectName); err != nil {
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
	// pfiletime {out} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.Filetime != nil {
			if err := o.Filetime.MarshalNDR(ctx, w); err != nil {
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
	// pfiletime {out} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.Filetime == nil {
			o.Filetime = &dtyp.Filetime{}
		}
		if err := o.Filetime.UnmarshalNDR(ctx, w); err != nil {
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
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ObjectName *urlmon.Moniker `idl:"name:pmkObjectName" json:"object_name"`
}

func (o *GetTimeOfLastChangeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTimeOfLastChangeOperation) *xxx_GetTimeOfLastChangeOperation {
	if op == nil {
		op = &xxx_GetTimeOfLastChangeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectName = o.ObjectName
	return op
}

func (o *GetTimeOfLastChangeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTimeOfLastChangeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectName = op.ObjectName
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
	Filetime *dtyp.Filetime `idl:"name:pfiletime" json:"filetime"`
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
	op.Filetime = o.Filetime
	op.Return = o.Return
	return op
}

func (o *GetTimeOfLastChangeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTimeOfLastChangeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Filetime = op.Filetime
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

// xxx_EnumRunningOperation structure represents the EnumRunning operation
type xxx_EnumRunningOperation struct {
	This    *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Moniker *urlmon.EnumMoniker `idl:"name:ppenumMoniker" json:"moniker"`
	Return  int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumRunningOperation) OpNum() int { return 9 }

func (o *xxx_EnumRunningOperation) OpName() string { return "/IRunningObjectTable/v0/EnumRunning" }

func (o *xxx_EnumRunningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRunningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumRunningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_EnumRunningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRunningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumRunningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EnumRunningRequest structure represents the EnumRunning operation request
type EnumRunningRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EnumRunningRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumRunningOperation) *xxx_EnumRunningOperation {
	if op == nil {
		op = &xxx_EnumRunningOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EnumRunningRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumRunningOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EnumRunningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumRunningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRunningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumRunningResponse structure represents the EnumRunning operation response
type EnumRunningResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Moniker *urlmon.EnumMoniker `idl:"name:ppenumMoniker" json:"moniker"`
	// Return: The EnumRunning return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumRunningResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumRunningOperation) *xxx_EnumRunningOperation {
	if op == nil {
		op = &xxx_EnumRunningOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Moniker = o.Moniker
	op.Return = o.Return
	return op
}

func (o *EnumRunningResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumRunningOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Moniker = op.Moniker
	o.Return = op.Return
}
func (o *EnumRunningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumRunningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRunningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

package ienumwbemclassobject

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
	wmi "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
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
	_ = wmi.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IEnumWbemClassObject interface identifier 027947e1-d731-11ce-a357-000000000001
	EnumClassObjectIID = &dcom.IID{Data1: 0x027947e1, Data2: 0xd731, Data3: 0x11ce, Data4: []byte{0xa3, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}}
	// Syntax UUID
	EnumClassObjectSyntaxUUID = &uuid.UUID{TimeLow: 0x27947e1, TimeMid: 0xd731, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x57, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x1}}
	// Syntax ID
	EnumClassObjectSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EnumClassObjectSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEnumWbemClassObject interface.
type EnumClassObjectClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// When the IEnumWbemClassObject::Reset method is invoked, the server MUST reset the
	// enumeration sequence to the beginning of the collection of CIM objects.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method. If the IEnumWbemClassObject::Reset
	// method is invoked on an enumerator that does not support reset capability, the server
	// MUST return WBEM_E_INVALID_OPERATION.
	Reset(context.Context, *ResetRequest, ...dcerpc.CallOption) (*ResetResponse, error)

	// Next operation.
	Next(context.Context, *NextRequest, ...dcerpc.CallOption) (*NextResponse, error)

	// The IEnumWbemClassObject::NextAsync method is the asynchronous version of the IEnumWbemClassObject::Next
	// method. It provides controlled asynchronous retrieval of CIM objects to a sink. The
	// server MUST asynchronously get one or more CIM objects, starting at the current position
	// in an enumeration, and MUST move the current position by the number of CIM objects.
	// When IEnumWbemClassObject is created, the current position MUST be set on the first
	// CIM object of the collection. The order of the CIM objects that are stored in the
	// enumerator is arbitrary.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	NextAsync(context.Context, *NextAsyncRequest, ...dcerpc.CallOption) (*NextAsyncResponse, error)

	// The IEnumWbemClassObject::Clone method makes a logical copy of the entire enumerator.
	// The cloned enumerator MUST have the same current position as the source enumerator.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	Clone(context.Context, *CloneRequest, ...dcerpc.CallOption) (*CloneResponse, error)

	// When the IEnumWbemClassObject::Skip method is invoked, the server MUST move the current
	// position in an enumeration ahead by a specified number of CIM objects.
	//
	// The IEnumWbemClassObject::Skip method opnum equals 7.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	Skip(context.Context, *SkipRequest, ...dcerpc.CallOption) (*SkipResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EnumClassObjectClient
}

type xxx_DefaultEnumClassObjectClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEnumClassObjectClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultEnumClassObjectClient) Reset(ctx context.Context, in *ResetRequest, opts ...dcerpc.CallOption) (*ResetResponse, error) {
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
	out := &ResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEnumClassObjectClient) Next(ctx context.Context, in *NextRequest, opts ...dcerpc.CallOption) (*NextResponse, error) {
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
	out := &NextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEnumClassObjectClient) NextAsync(ctx context.Context, in *NextAsyncRequest, opts ...dcerpc.CallOption) (*NextAsyncResponse, error) {
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
	out := &NextAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEnumClassObjectClient) Clone(ctx context.Context, in *CloneRequest, opts ...dcerpc.CallOption) (*CloneResponse, error) {
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
	out := &CloneResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEnumClassObjectClient) Skip(ctx context.Context, in *SkipRequest, opts ...dcerpc.CallOption) (*SkipResponse, error) {
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
	out := &SkipResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEnumClassObjectClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEnumClassObjectClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEnumClassObjectClient) IPID(ctx context.Context, ipid *dcom.IPID) EnumClassObjectClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEnumClassObjectClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewEnumClassObjectClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EnumClassObjectClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EnumClassObjectSyntaxV0_0))...)
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
	return &xxx_DefaultEnumClassObjectClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ResetOperation structure represents the Reset operation
type xxx_ResetOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResetOperation) OpNum() int { return 3 }

func (o *xxx_ResetOperation) OpName() string { return "/IEnumWbemClassObject/v0/Reset" }

func (o *xxx_ResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResetRequest structure represents the Reset operation request
type ResetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResetRequest) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ResetRequest) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResetResponse structure represents the Reset operation response
type ResetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reset return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResetResponse) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ResetResponse) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NextOperation structure represents the Next operation
type xxx_NextOperation struct {
	This     *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Timeout  int32              `idl:"name:lTimeout" json:"timeout"`
	Count    uint32             `idl:"name:uCount" json:"count"`
	Objects  []*wmi.ClassObject `idl:"name:apObjects;size_is:(uCount);length_is:(puReturned)" json:"objects"`
	Returned uint32             `idl:"name:puReturned" json:"returned"`
	Return   int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_NextOperation) OpNum() int { return 4 }

func (o *xxx_NextOperation) OpName() string { return "/IEnumWbemClassObject/v0/Next" }

func (o *xxx_NextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// uCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// uCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Objects != nil && o.Returned == 0 {
		o.Returned = uint32(len(o.Objects))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// apObjects {out} (1:{pointer=ref}*(1)[dim:0,size_is=uCount,length_is=puReturned]*(1))(2:{alias=IWbemClassObject}(interface))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.Returned)
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
		for i1 := range o.Objects {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Objects[i1] != nil {
				_ptr_apObjects := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Objects[i1] != nil {
						if err := o.Objects[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Objects[i1], _ptr_apObjects); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Objects); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// puReturned {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Returned); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// apObjects {out} (1:{pointer=ref}*(1)[dim:0,size_is=uCount,length_is=puReturned]*(1))(2:{alias=IWbemClassObject}(interface))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Objects", sizeInfo[0])
		}
		o.Objects = make([]*wmi.ClassObject, sizeInfo[0])
		for i1 := range o.Objects {
			i1 := i1
			_ptr_apObjects := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Objects[i1] == nil {
					o.Objects[i1] = &wmi.ClassObject{}
				}
				if err := o.Objects[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_apObjects := func(ptr interface{}) { o.Objects[i1] = *ptr.(**wmi.ClassObject) }
			if err := w.ReadPointer(&o.Objects[i1], _s_apObjects, _ptr_apObjects); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// puReturned {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Returned); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NextRequest structure represents the Next operation request
type NextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Timeout int32          `idl:"name:lTimeout" json:"timeout"`
	Count   uint32         `idl:"name:uCount" json:"count"`
}

func (o *NextRequest) xxx_ToOp(ctx context.Context, op *xxx_NextOperation) *xxx_NextOperation {
	if op == nil {
		op = &xxx_NextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Timeout = o.Timeout
	op.Count = o.Count
	return op
}

func (o *NextRequest) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Timeout = op.Timeout
	o.Count = op.Count
}
func (o *NextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NextResponse structure represents the Next operation response
type NextResponse struct {
	// XXX: uCount is an implicit input depedency for output parameters
	Count uint32 `idl:"name:uCount" json:"count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Objects  []*wmi.ClassObject `idl:"name:apObjects;size_is:(uCount);length_is:(puReturned)" json:"objects"`
	Returned uint32             `idl:"name:puReturned" json:"returned"`
	// Return: The Next return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NextResponse) xxx_ToOp(ctx context.Context, op *xxx_NextOperation) *xxx_NextOperation {
	if op == nil {
		op = &xxx_NextOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Count == uint32(0) {
		op.Count = o.Count
	}

	op.That = o.That
	op.Objects = o.Objects
	op.Returned = o.Returned
	op.Return = o.Return
	return op
}

func (o *NextResponse) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Count = op.Count

	o.That = op.That
	o.Objects = op.Objects
	o.Returned = op.Returned
	o.Return = op.Return
}
func (o *NextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NextAsyncOperation structure represents the NextAsync operation
type xxx_NextAsyncOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Count  uint32          `idl:"name:uCount" json:"count"`
	Sink   *wmi.ObjectSink `idl:"name:pSink" json:"sink"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_NextAsyncOperation) OpNum() int { return 5 }

func (o *xxx_NextAsyncOperation) OpName() string { return "/IEnumWbemClassObject/v0/NextAsync" }

func (o *xxx_NextAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// uCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// pSink {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		if o.Sink != nil {
			_ptr_pSink := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Sink != nil {
					if err := o.Sink.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ObjectSink{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_NextAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// uCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// pSink {in} (1:{pointer=ref}*(1))(2:{alias=IWbemObjectSink}(interface))
	{
		_ptr_pSink := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Sink == nil {
				o.Sink = &wmi.ObjectSink{}
			}
			if err := o.Sink.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSink := func(ptr interface{}) { o.Sink = *ptr.(**wmi.ObjectSink) }
		if err := w.ReadPointer(&o.Sink, _s_pSink, _ptr_pSink); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NextAsyncRequest structure represents the NextAsync operation request
type NextAsyncRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// uCount: MUST be the number of CIM objects being requested.
	Count uint32 `idl:"name:uCount" json:"count"`
	// pSink: MUST be a pointer to the IWbemObjectSink interface, which MUST represent the
	// sink to receive the CIM object. As each batch of CIM objects is requested, they MUST
	// be delivered to the IWbemObjectSink::Indicate method to which pSink points (as specified
	// in section 3.1.4.2.1) and MUST be followed by a final call to the IWbemObjectSink::SetStatus
	// method to which pSink points, as specified in section 3.1.4.2.2. This parameter MUST
	// NOT be NULL. In error cases, indicated by the HRESULT return value, the supplied
	// IWbemObjectSink interface pointer MUST NOT be used by the server.
	Sink *wmi.ObjectSink `idl:"name:pSink" json:"sink"`
}

func (o *NextAsyncRequest) xxx_ToOp(ctx context.Context, op *xxx_NextAsyncOperation) *xxx_NextAsyncOperation {
	if op == nil {
		op = &xxx_NextAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Count = o.Count
	op.Sink = o.Sink
	return op
}

func (o *NextAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_NextAsyncOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Count = op.Count
	o.Sink = op.Sink
}
func (o *NextAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NextAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NextAsyncResponse structure represents the NextAsync operation response
type NextAsyncResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The NextAsync return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NextAsyncResponse) xxx_ToOp(ctx context.Context, op *xxx_NextAsyncOperation) *xxx_NextAsyncOperation {
	if op == nil {
		op = &xxx_NextAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *NextAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_NextAsyncOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *NextAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NextAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloneOperation structure represents the Clone operation
type xxx_CloneOperation struct {
	This   *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Enum   *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	Return int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CloneOperation) OpNum() int { return 6 }

func (o *xxx_CloneOperation) OpName() string { return "/IEnumWbemClassObject/v0/Clone" }

func (o *xxx_CloneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CloneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.EnumClassObject{}).MarshalNDR(ctx, w); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumWbemClassObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &wmi.EnumClassObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**wmi.EnumClassObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloneRequest structure represents the Clone operation request
type CloneRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CloneRequest) xxx_ToOp(ctx context.Context, op *xxx_CloneOperation) *xxx_CloneOperation {
	if op == nil {
		op = &xxx_CloneOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CloneRequest) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CloneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloneResponse structure represents the Clone operation response
type CloneResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: Upon return, MUST contain a pointer to an IEnumWbemClassObject interface
	// CIM object that is a logical copy of the entire enumerator that made the Clone method
	// call, retaining the current position in an enumeration. This parameter MUST NOT be
	// NULL. When returned by the server, this parameter can be NULL if a failure occurred
	// or if there are no results.
	Enum *wmi.EnumClassObject `idl:"name:ppEnum" json:"enum"`
	// Return: The Clone return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloneResponse) xxx_ToOp(ctx context.Context, op *xxx_CloneOperation) *xxx_CloneOperation {
	if op == nil {
		op = &xxx_CloneOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *CloneResponse) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *CloneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SkipOperation structure represents the Skip operation
type xxx_SkipOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Timeout int32          `idl:"name:lTimeout" json:"timeout"`
	Count   uint32         `idl:"name:nCount" json:"count"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SkipOperation) OpNum() int { return 7 }

func (o *xxx_SkipOperation) OpName() string { return "/IEnumWbemClassObject/v0/Skip" }

func (o *xxx_SkipOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// nCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// nCount {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SkipRequest structure represents the Skip operation request
type SkipRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lTimeout: MUST be the maximum amount of time, in milliseconds, that the call to Skip
	// allows to pass before it times out. If the constant WBEM_INFINITE (0xFFFFFFFF) is
	// used, the Skip method call waits until the operation succeeds.
	Timeout int32 `idl:"name:lTimeout" json:"timeout"`
	// nCount: MUST be the number of CIM objects to skip in the enumeration. If this parameter
	// is greater than the number of CIM objects that remain to enumerate, the call MUST
	// skip to the end of the enumeration, and WBEM_S_FALSE MUST be the returned value for
	// the method.
	Count uint32 `idl:"name:nCount" json:"count"`
}

func (o *SkipRequest) xxx_ToOp(ctx context.Context, op *xxx_SkipOperation) *xxx_SkipOperation {
	if op == nil {
		op = &xxx_SkipOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Timeout = o.Timeout
	op.Count = o.Count
	return op
}

func (o *SkipRequest) xxx_FromOp(ctx context.Context, op *xxx_SkipOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Timeout = op.Timeout
	o.Count = op.Count
}
func (o *SkipRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SkipRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SkipOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SkipResponse structure represents the Skip operation response
type SkipResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Skip return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SkipResponse) xxx_ToOp(ctx context.Context, op *xxx_SkipOperation) *xxx_SkipOperation {
	if op == nil {
		op = &xxx_SkipOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SkipResponse) xxx_FromOp(ctx context.Context, op *xxx_SkipOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SkipResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SkipResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SkipOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

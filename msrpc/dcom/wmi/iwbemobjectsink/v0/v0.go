package iwbemobjectsink

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
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemObjectSink interface identifier 7c857801-7381-11cf-884d-00aa004b2e24
	ObjectSinkIID = &dcom.IID{Data1: 0x7c857801, Data2: 0x7381, Data3: 0x11cf, Data4: []byte{0x88, 0x4d, 0x00, 0xaa, 0x00, 0x4b, 0x2e, 0x24}}
	// Syntax UUID
	ObjectSinkSyntaxUUID = &uuid.UUID{TimeLow: 0x7c857801, TimeMid: 0x7381, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x4d, Node: [6]uint8{0x0, 0xaa, 0x0, 0x4b, 0x2e, 0x24}}
	// Syntax ID
	ObjectSinkSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ObjectSinkSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemObjectSink interface.
type ObjectSinkClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// When the IWbemObjectSink::Indicate method is called, the apObjArray parameter MUST
	// contain additional result objects as an array of an IWbemClassObject, sent by the
	// client to the server. The IWbemObjectSink::Indicate method has the following syntax,
	// expressed in Microsoft Interface Definition Language (MIDL).
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call.
	//
	// When the IWbemObjectSink::Indicate method is called for the first time, the server
	// that implements the ObjectArray structure MUST return WBEM_S_NEW_STYLE if the execution
	// of the method succeeds. If a server does not implement the ObjectArray structure,
	// it MUST return WBEM_S_NO_ERROR for successful execution of the method.
	//
	// If the server implements the ObjectArray structure and WBEM_S_NEW_STYLE is returned
	// during the first call to the IWbemObjectSink::Indicate method, the server MUST support
	// subsequent calls to the IWbemObjectSink::Indicate method by using both the DCOM Remote
	// Protocol marshaling and the ObjectArray structure as specified in section 2.2.14.
	Indicate(context.Context, *IndicateRequest, ...dcerpc.CallOption) (*IndicateResponse, error)

	// When the IWbemObjectSink::SetStatus method is called, the parameter MUST contain
	// the result of the operation or the progress status information.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	SetStatus(context.Context, *SetStatusRequest, ...dcerpc.CallOption) (*SetStatusResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ObjectSinkClient
}

type xxx_DefaultObjectSinkClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultObjectSinkClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultObjectSinkClient) Indicate(ctx context.Context, in *IndicateRequest, opts ...dcerpc.CallOption) (*IndicateResponse, error) {
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
	out := &IndicateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectSinkClient) SetStatus(ctx context.Context, in *SetStatusRequest, opts ...dcerpc.CallOption) (*SetStatusResponse, error) {
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
	out := &SetStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectSinkClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultObjectSinkClient) IPID(ctx context.Context, ipid *dcom.IPID) ObjectSinkClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultObjectSinkClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewObjectSinkClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ObjectSinkClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ObjectSinkSyntaxV0_0))...)
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
	return &xxx_DefaultObjectSinkClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_IndicateOperation structure represents the Indicate operation
type xxx_IndicateOperation struct {
	This        *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat     `idl:"name:That" json:"that"`
	ObjectCount int32              `idl:"name:lObjectCount" json:"object_count"`
	ObjectArray []*wmi.ClassObject `idl:"name:apObjArray;size_is:(lObjectCount)" json:"object_array"`
	Return      int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_IndicateOperation) OpNum() int { return 3 }

func (o *xxx_IndicateOperation) OpName() string { return "/IWbemObjectSink/v0/Indicate" }

func (o *xxx_IndicateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ObjectArray != nil && o.ObjectCount == 0 {
		o.ObjectCount = int32(len(o.ObjectArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IndicateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lObjectCount {in} (1:(int32))
	{
		if err := w.WriteData(o.ObjectCount); err != nil {
			return err
		}
	}
	// apObjArray {in} (1:{pointer=ref}*(1)[dim:0,size_is=lObjectCount]*(1))(2:{alias=IWbemClassObject}(interface))
	{
		dimSize1 := uint64(o.ObjectCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ObjectArray {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ObjectArray[i1] != nil {
				_ptr_apObjArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.ObjectArray[i1] != nil {
						if err := o.ObjectArray[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.ObjectArray[i1], _ptr_apObjArray); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ObjectArray); uint64(i1) < sizeInfo[0]; i1++ {
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

func (o *xxx_IndicateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lObjectCount {in} (1:(int32))
	{
		if err := w.ReadData(&o.ObjectCount); err != nil {
			return err
		}
	}
	// apObjArray {in} (1:{pointer=ref}*(1)[dim:0,size_is=lObjectCount]*(1))(2:{alias=IWbemClassObject}(interface))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ObjectArray", sizeInfo[0])
		}
		o.ObjectArray = make([]*wmi.ClassObject, sizeInfo[0])
		for i1 := range o.ObjectArray {
			i1 := i1
			_ptr_apObjArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ObjectArray[i1] == nil {
					o.ObjectArray[i1] = &wmi.ClassObject{}
				}
				if err := o.ObjectArray[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_apObjArray := func(ptr interface{}) { o.ObjectArray[i1] = *ptr.(**wmi.ClassObject) }
			if err := w.ReadPointer(&o.ObjectArray[i1], _s_apObjArray, _ptr_apObjArray); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IndicateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IndicateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IndicateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IndicateRequest structure represents the Indicate operation request
type IndicateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lObjectCount: MUST be the number of CIM objects in the array of pointers in the ppObjArray
	// parameter.
	ObjectCount int32 `idl:"name:lObjectCount" json:"object_count"`
	// apObjArray: MUST contain an array of result objects sent by the caller.
	ObjectArray []*wmi.ClassObject `idl:"name:apObjArray;size_is:(lObjectCount)" json:"object_array"`
}

func (o *IndicateRequest) xxx_ToOp(ctx context.Context) *xxx_IndicateOperation {
	if o == nil {
		return &xxx_IndicateOperation{}
	}
	return &xxx_IndicateOperation{
		This:        o.This,
		ObjectCount: o.ObjectCount,
		ObjectArray: o.ObjectArray,
	}
}

func (o *IndicateRequest) xxx_FromOp(ctx context.Context, op *xxx_IndicateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectCount = op.ObjectCount
	o.ObjectArray = op.ObjectArray
}
func (o *IndicateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IndicateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IndicateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IndicateResponse structure represents the Indicate operation response
type IndicateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Indicate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IndicateResponse) xxx_ToOp(ctx context.Context) *xxx_IndicateOperation {
	if o == nil {
		return &xxx_IndicateOperation{}
	}
	return &xxx_IndicateOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *IndicateResponse) xxx_FromOp(ctx context.Context, op *xxx_IndicateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IndicateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IndicateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IndicateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetStatusOperation structure represents the SetStatus operation
type xxx_SetStatusOperation struct {
	This        *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Flags       int32            `idl:"name:lFlags" json:"flags"`
	HResult     int32            `idl:"name:hResult" json:"hresult"`
	Param       *oaut.String     `idl:"name:strParam" json:"param"`
	ObjectParam *wmi.ClassObject `idl:"name:pObjParam" json:"object_param"`
	Return      int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetStatusOperation) OpNum() int { return 4 }

func (o *xxx_SetStatusOperation) OpName() string { return "/IWbemObjectSink/v0/SetStatus" }

func (o *xxx_SetStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// hResult {in} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.HResult); err != nil {
			return err
		}
	}
	// strParam {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Param != nil {
			_ptr_strParam := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Param != nil {
					if err := o.Param.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Param, _ptr_strParam); err != nil {
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
	// pObjParam {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.ObjectParam != nil {
			_ptr_pObjParam := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectParam != nil {
					if err := o.ObjectParam.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectParam, _ptr_pObjParam); err != nil {
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

func (o *xxx_SetStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// hResult {in} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.HResult); err != nil {
			return err
		}
	}
	// strParam {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strParam := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Param == nil {
				o.Param = &oaut.String{}
			}
			if err := o.Param.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strParam := func(ptr interface{}) { o.Param = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Param, _s_strParam, _ptr_strParam); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pObjParam {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pObjParam := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectParam == nil {
				o.ObjectParam = &wmi.ClassObject{}
			}
			if err := o.ObjectParam.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pObjParam := func(ptr interface{}) { o.ObjectParam = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.ObjectParam, _s_pObjParam, _ptr_pObjParam); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetStatusRequest structure represents the SetStatus operation request
type SetStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lFlags: Flags that give information about the operation status. The flags MUST be
	// interpreted as specified in the following table.
	//
	// Note  The flags are not bit flags and cannot be combined.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|                                 |                                                             |
	//	|              VALUE              |                           MEANING                           |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| WBEM_STATUS_COMPLETE 0x00000000 | Indicates the end of the asynchronous operation.            |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| WBEM_STATUS_PROGRESS 0x00000002 | Indicates the progress state of the asynchronous operation. |
	//	+---------------------------------+-------------------------------------------------------------+
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// hResult: The HRESULT value of the asynchronous operation or notification. This hResult
	// MUST be the same HRESULT that the WMI client gets from the matching synchronous operation
	// when the WMI client makes an asynchronous request to the WMI server.
	HResult int32 `idl:"name:hResult" json:"hresult"`
	// strParam: If the parameter is NULL, the server MUST ignore the parameter. If the
	// parameter is not NULL, it MUST represent the operational result of the asynchronous
	// operation. The string MUST be the same as the string that is returned from the IWbemCallResult::GetResultString
	// (Opnum 4) method when the operation is executed synchronously.
	Param *oaut.String `idl:"name:strParam" json:"param"`
	// pObjParam: If the parameter is NULL, the server MUST ignore the parameter. If the
	// parameter is not NULL, the object MUST contain additional error information for the
	// asynchronous operation failure.
	ObjectParam *wmi.ClassObject `idl:"name:pObjParam" json:"object_param"`
}

func (o *SetStatusRequest) xxx_ToOp(ctx context.Context) *xxx_SetStatusOperation {
	if o == nil {
		return &xxx_SetStatusOperation{}
	}
	return &xxx_SetStatusOperation{
		This:        o.This,
		Flags:       o.Flags,
		HResult:     o.HResult,
		Param:       o.Param,
		ObjectParam: o.ObjectParam,
	}
}

func (o *SetStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_SetStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.HResult = op.HResult
	o.Param = op.Param
	o.ObjectParam = op.ObjectParam
}
func (o *SetStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetStatusResponse structure represents the SetStatus operation response
type SetStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetStatusResponse) xxx_ToOp(ctx context.Context) *xxx_SetStatusOperation {
	if o == nil {
		return &xxx_SetStatusOperation{}
	}
	return &xxx_SetStatusOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_SetStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

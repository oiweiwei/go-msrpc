package ivdssubsystemimporttarget

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
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsSubSystemImportTarget interface identifier 83bfb87f-43fb-4903-baa6-127f01029eec
	SubSystemImportTargetIID = &dcom.IID{Data1: 0x83bfb87f, Data2: 0x43fb, Data3: 0x4903, Data4: []byte{0xba, 0xa6, 0x12, 0x7f, 0x01, 0x02, 0x9e, 0xec}}
	// Syntax UUID
	SubSystemImportTargetSyntaxUUID = &uuid.UUID{TimeLow: 0x83bfb87f, TimeMid: 0x43fb, TimeHiAndVersion: 0x4903, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0xa6, Node: [6]uint8{0x12, 0x7f, 0x1, 0x2, 0x9e, 0xec}}
	// Syntax ID
	SubSystemImportTargetSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: SubSystemImportTargetSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsSubSystemImportTarget interface.
type SubSystemImportTargetClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetImportTarget method retrieves the name of the import target to associate with
	// the LUNs being imported on the subsystem.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetImportTarget(context.Context, *GetImportTargetRequest, ...dcerpc.CallOption) (*GetImportTargetResponse, error)

	// The SetImportTarget method sets the name of the import target to associate with the
	// LUNs being imported on the subsystem.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetImportTarget(context.Context, *SetImportTargetRequest, ...dcerpc.CallOption) (*SetImportTargetResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) SubSystemImportTargetClient
}

type xxx_DefaultSubSystemImportTargetClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultSubSystemImportTargetClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultSubSystemImportTargetClient) GetImportTarget(ctx context.Context, in *GetImportTargetRequest, opts ...dcerpc.CallOption) (*GetImportTargetResponse, error) {
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
	out := &GetImportTargetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSubSystemImportTargetClient) SetImportTarget(ctx context.Context, in *SetImportTargetRequest, opts ...dcerpc.CallOption) (*SetImportTargetResponse, error) {
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
	out := &SetImportTargetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSubSystemImportTargetClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSubSystemImportTargetClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultSubSystemImportTargetClient) IPID(ctx context.Context, ipid *dcom.IPID) SubSystemImportTargetClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultSubSystemImportTargetClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewSubSystemImportTargetClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (SubSystemImportTargetClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(SubSystemImportTargetSyntaxV0_0))...)
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
	return &xxx_DefaultSubSystemImportTargetClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetImportTargetOperation structure represents the GetImportTarget operation
type xxx_GetImportTargetOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ISCSIName string         `idl:"name:ppwszIscsiName;string" json:"iscsi_name"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetImportTargetOperation) OpNum() int { return 3 }

func (o *xxx_GetImportTargetOperation) OpName() string {
	return "/IVdsSubSystemImportTarget/v0/GetImportTarget"
}

func (o *xxx_GetImportTargetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImportTargetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetImportTargetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetImportTargetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImportTargetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppwszIscsiName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ISCSIName != "" {
			_ptr_ppwszIscsiName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ISCSIName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ISCSIName, _ptr_ppwszIscsiName); err != nil {
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

func (o *xxx_GetImportTargetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppwszIscsiName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppwszIscsiName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ISCSIName); err != nil {
				return err
			}
			return nil
		})
		_s_ppwszIscsiName := func(ptr interface{}) { o.ISCSIName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ISCSIName, _s_ppwszIscsiName, _ptr_ppwszIscsiName); err != nil {
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

// GetImportTargetRequest structure represents the GetImportTarget operation request
type GetImportTargetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetImportTargetRequest) xxx_ToOp(ctx context.Context, op *xxx_GetImportTargetOperation) *xxx_GetImportTargetOperation {
	if op == nil {
		op = &xxx_GetImportTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetImportTargetRequest) xxx_FromOp(ctx context.Context, op *xxx_GetImportTargetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetImportTargetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetImportTargetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImportTargetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetImportTargetResponse structure represents the GetImportTarget operation response
type GetImportTargetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszIscsiName: A pointer to a string that contains the name of the import target
	// of the subsystem. Callers MUST free the memory that is allocated for the string when
	// they are finished with it.
	ISCSIName string `idl:"name:ppwszIscsiName;string" json:"iscsi_name"`
	// Return: The GetImportTarget return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetImportTargetResponse) xxx_ToOp(ctx context.Context, op *xxx_GetImportTargetOperation) *xxx_GetImportTargetOperation {
	if op == nil {
		op = &xxx_GetImportTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ISCSIName = o.ISCSIName
	op.Return = o.Return
	return op
}

func (o *GetImportTargetResponse) xxx_FromOp(ctx context.Context, op *xxx_GetImportTargetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ISCSIName = op.ISCSIName
	o.Return = op.Return
}
func (o *GetImportTargetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetImportTargetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImportTargetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetImportTargetOperation structure represents the SetImportTarget operation
type xxx_SetImportTargetOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ISCSIName string         `idl:"name:pwszIscsiName;string;pointer:unique" json:"iscsi_name"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetImportTargetOperation) OpNum() int { return 4 }

func (o *xxx_SetImportTargetOperation) OpName() string {
	return "/IVdsSubSystemImportTarget/v0/SetImportTarget"
}

func (o *xxx_SetImportTargetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetImportTargetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszIscsiName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ISCSIName != "" {
			_ptr_pwszIscsiName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ISCSIName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ISCSIName, _ptr_pwszIscsiName); err != nil {
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

func (o *xxx_SetImportTargetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszIscsiName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszIscsiName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ISCSIName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszIscsiName := func(ptr interface{}) { o.ISCSIName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ISCSIName, _s_pwszIscsiName, _ptr_pwszIscsiName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetImportTargetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetImportTargetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetImportTargetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetImportTargetRequest structure represents the SetImportTarget operation request
type SetImportTargetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszIscsiName: A string that contains the name of the import target of the subsystem.
	ISCSIName string `idl:"name:pwszIscsiName;string;pointer:unique" json:"iscsi_name"`
}

func (o *SetImportTargetRequest) xxx_ToOp(ctx context.Context, op *xxx_SetImportTargetOperation) *xxx_SetImportTargetOperation {
	if op == nil {
		op = &xxx_SetImportTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ISCSIName = o.ISCSIName
	return op
}

func (o *SetImportTargetRequest) xxx_FromOp(ctx context.Context, op *xxx_SetImportTargetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ISCSIName = op.ISCSIName
}
func (o *SetImportTargetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetImportTargetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetImportTargetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetImportTargetResponse structure represents the SetImportTarget operation response
type SetImportTargetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetImportTarget return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetImportTargetResponse) xxx_ToOp(ctx context.Context, op *xxx_SetImportTargetOperation) *xxx_SetImportTargetOperation {
	if op == nil {
		op = &xxx_SetImportTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetImportTargetResponse) xxx_FromOp(ctx context.Context, op *xxx_SetImportTargetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetImportTargetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetImportTargetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetImportTargetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

package iwindowsdriverupdate4

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
	iwindowsdriverupdate3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate3/v0"
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
	_ = iwindowsdriverupdate3.GoPackage
	_ = uamg.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IWindowsDriverUpdate4 interface identifier 004c6a2b-0c19-4c69-9f5c-a269b2560db9
	WindowsDriverUpdate4IID = &dcom.IID{Data1: 0x004c6a2b, Data2: 0x0c19, Data3: 0x4c69, Data4: []byte{0x9f, 0x5c, 0xa2, 0x69, 0xb2, 0x56, 0x0d, 0xb9}}
	// Syntax UUID
	WindowsDriverUpdate4SyntaxUUID = &uuid.UUID{TimeLow: 0x4c6a2b, TimeMid: 0xc19, TimeHiAndVersion: 0x4c69, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x5c, Node: [6]uint8{0xa2, 0x69, 0xb2, 0x56, 0xd, 0xb9}}
	// Syntax ID
	WindowsDriverUpdate4SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: WindowsDriverUpdate4SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWindowsDriverUpdate4 interface.
type WindowsDriverUpdate4Client interface {

	// IWindowsDriverUpdate3 retrieval method.
	WindowsDriverUpdate3() iwindowsdriverupdate3.WindowsDriverUpdate3Client

	// The IWindowsDriverUpdate4::WindowsDriverUpdateEntries (opnum 66) method retrieves
	// a collection of driver update entries associated with this driver.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the WindowsDriverUpdateEntries ADM element.
	GetWindowsDriverUpdateEntries(context.Context, *GetWindowsDriverUpdateEntriesRequest, ...dcerpc.CallOption) (*GetWindowsDriverUpdateEntriesResponse, error)

	// The IUpdate4::PerUser (opnum 58) method retrieves whether the update is per user.
	//
	// The IWindowsDriverUpdate4::PerUser (opnum 67) method retrieves whether the update
	// is per user.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the PerUser ADM element.
	GetPerUser(context.Context, *GetPerUserRequest, ...dcerpc.CallOption) (*GetPerUserResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) WindowsDriverUpdate4Client
}

type xxx_DefaultWindowsDriverUpdate4Client struct {
	iwindowsdriverupdate3.WindowsDriverUpdate3Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultWindowsDriverUpdate4Client) WindowsDriverUpdate3() iwindowsdriverupdate3.WindowsDriverUpdate3Client {
	return o.WindowsDriverUpdate3Client
}

func (o *xxx_DefaultWindowsDriverUpdate4Client) GetWindowsDriverUpdateEntries(ctx context.Context, in *GetWindowsDriverUpdateEntriesRequest, opts ...dcerpc.CallOption) (*GetWindowsDriverUpdateEntriesResponse, error) {
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
	out := &GetWindowsDriverUpdateEntriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWindowsDriverUpdate4Client) GetPerUser(ctx context.Context, in *GetPerUserRequest, opts ...dcerpc.CallOption) (*GetPerUserResponse, error) {
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
	out := &GetPerUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWindowsDriverUpdate4Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWindowsDriverUpdate4Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultWindowsDriverUpdate4Client) IPID(ctx context.Context, ipid *dcom.IPID) WindowsDriverUpdate4Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultWindowsDriverUpdate4Client{
		WindowsDriverUpdate3Client: o.WindowsDriverUpdate3Client.IPID(ctx, ipid),
		cc:                         o.cc,
		ipid:                       ipid,
	}
}

func NewWindowsDriverUpdate4Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WindowsDriverUpdate4Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate4SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iwindowsdriverupdate3.NewWindowsDriverUpdate3Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultWindowsDriverUpdate4Client{
		WindowsDriverUpdate3Client: base,
		cc:                         cc,
		ipid:                       ipid,
	}, nil
}

// xxx_GetWindowsDriverUpdateEntriesOperation structure represents the WindowsDriverUpdateEntries operation
type xxx_GetWindowsDriverUpdateEntriesOperation struct {
	This        *dcom.ORPCThis                           `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat                           `idl:"name:That" json:"that"`
	ReturnValue *uamg.WindowsDriverUpdateEntryCollection `idl:"name:retval" json:"return_value"`
	Return      int32                                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) OpNum() int { return 65 }

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) OpName() string {
	return "/IWindowsDriverUpdate4/v0/WindowsDriverUpdateEntries"
}

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IWindowsDriverUpdateEntryCollection}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.WindowsDriverUpdateEntryCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_GetWindowsDriverUpdateEntriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IWindowsDriverUpdateEntryCollection}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.WindowsDriverUpdateEntryCollection{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.WindowsDriverUpdateEntryCollection) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// GetWindowsDriverUpdateEntriesRequest structure represents the WindowsDriverUpdateEntries operation request
type GetWindowsDriverUpdateEntriesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetWindowsDriverUpdateEntriesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetWindowsDriverUpdateEntriesOperation) *xxx_GetWindowsDriverUpdateEntriesOperation {
	if op == nil {
		op = &xxx_GetWindowsDriverUpdateEntriesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetWindowsDriverUpdateEntriesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetWindowsDriverUpdateEntriesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetWindowsDriverUpdateEntriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetWindowsDriverUpdateEntriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetWindowsDriverUpdateEntriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetWindowsDriverUpdateEntriesResponse structure represents the WindowsDriverUpdateEntries operation response
type GetWindowsDriverUpdateEntriesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An IWindowsDriverUpdateEntryCollection instance containing IWindowsDriverUpdateEntry
	// instances for each device for which this driver can be installed.
	ReturnValue *uamg.WindowsDriverUpdateEntryCollection `idl:"name:retval" json:"return_value"`
	// Return: The WindowsDriverUpdateEntries return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetWindowsDriverUpdateEntriesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetWindowsDriverUpdateEntriesOperation) *xxx_GetWindowsDriverUpdateEntriesOperation {
	if op == nil {
		op = &xxx_GetWindowsDriverUpdateEntriesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetWindowsDriverUpdateEntriesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetWindowsDriverUpdateEntriesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetWindowsDriverUpdateEntriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetWindowsDriverUpdateEntriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetWindowsDriverUpdateEntriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPerUserOperation structure represents the PerUser operation
type xxx_GetPerUserOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPerUserOperation) OpNum() int { return 66 }

func (o *xxx_GetPerUserOperation) OpName() string { return "/IWindowsDriverUpdate4/v0/PerUser" }

func (o *xxx_GetPerUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPerUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPerUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
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

func (o *xxx_GetPerUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
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

// GetPerUserRequest structure represents the PerUser operation request
type GetPerUserRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPerUserRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPerUserOperation) *xxx_GetPerUserOperation {
	if op == nil {
		op = &xxx_GetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPerUserRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPerUserOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPerUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPerUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPerUserResponse structure represents the PerUser operation response
type GetPerUserResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: MUST be set either to VARIANT_TRUE if this update is installed per user rather
	// than per machine or to VARIANT_FALSE if not.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
	// Return: The PerUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPerUserResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPerUserOperation) *xxx_GetPerUserOperation {
	if op == nil {
		op = &xxx_GetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetPerUserResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPerUserOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetPerUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPerUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

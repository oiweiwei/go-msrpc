package iupdatesession2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iupdatesession "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesession/v0"
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
	_ = iupdatesession.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateSession2 interface identifier 91caf7b0-eb23-49ed-9937-c52d817f46f7
	UpdateSession2IID = &dcom.IID{Data1: 0x91caf7b0, Data2: 0xeb23, Data3: 0x49ed, Data4: []byte{0x99, 0x37, 0xc5, 0x2d, 0x81, 0x7f, 0x46, 0xf7}}
	// Syntax UUID
	UpdateSession2SyntaxUUID = &uuid.UUID{TimeLow: 0x91caf7b0, TimeMid: 0xeb23, TimeHiAndVersion: 0x49ed, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x37, Node: [6]uint8{0xc5, 0x2d, 0x81, 0x7f, 0x46, 0xf7}}
	// Syntax ID
	UpdateSession2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateSession2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateSession2 interface.
type UpdateSession2Client interface {

	// IUpdateSession retrieval method.
	UpdateSession() iupdatesession.UpdateSessionClient

	// The IUpdateSession2::UserLocale (opnum 14) method gets the language for update results.
	//
	// The IUpdateSession2::UserLocale (opnum 15) method sets the language for update results.
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
	// This method SHOULD return the value of the UserLocale ADM element.
	GetUserLocale(context.Context, *GetUserLocaleRequest, ...dcerpc.CallOption) (*GetUserLocaleResponse, error)

	// The IUpdateSession2::UserLocale (opnum 14) method gets the language for update results.
	//
	// The IUpdateSession2::UserLocale (opnum 15) method sets the language for update results.
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
	// This method SHOULD return the value of the UserLocale ADM element.
	SetUserLocale(context.Context, *SetUserLocaleRequest, ...dcerpc.CallOption) (*SetUserLocaleResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateSession2Client
}

type xxx_DefaultUpdateSession2Client struct {
	iupdatesession.UpdateSessionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateSession2Client) UpdateSession() iupdatesession.UpdateSessionClient {
	return o.UpdateSessionClient
}

func (o *xxx_DefaultUpdateSession2Client) GetUserLocale(ctx context.Context, in *GetUserLocaleRequest, opts ...dcerpc.CallOption) (*GetUserLocaleResponse, error) {
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
	out := &GetUserLocaleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSession2Client) SetUserLocale(ctx context.Context, in *SetUserLocaleRequest, opts ...dcerpc.CallOption) (*SetUserLocaleResponse, error) {
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
	out := &SetUserLocaleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSession2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateSession2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateSession2Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateSession2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateSession2Client{
		UpdateSessionClient: o.UpdateSessionClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}

func NewUpdateSession2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateSession2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateSession2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdatesession.NewUpdateSessionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateSession2Client{
		UpdateSessionClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_GetUserLocaleOperation structure represents the UserLocale operation
type xxx_GetUserLocaleOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue uint32         `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUserLocaleOperation) OpNum() int { return 15 }

func (o *xxx_GetUserLocaleOperation) OpName() string { return "/IUpdateSession2/v0/UserLocale" }

func (o *xxx_GetUserLocaleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserLocaleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetUserLocaleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetUserLocaleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserLocaleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=LCID, names=DWORD}(uint32))
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

func (o *xxx_GetUserLocaleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=LCID, names=DWORD}(uint32))
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

// GetUserLocaleRequest structure represents the UserLocale operation request
type GetUserLocaleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetUserLocaleRequest) xxx_ToOp(ctx context.Context, op *xxx_GetUserLocaleOperation) *xxx_GetUserLocaleOperation {
	if op == nil {
		op = &xxx_GetUserLocaleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetUserLocaleRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserLocaleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetUserLocaleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetUserLocaleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserLocaleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUserLocaleResponse structure represents the UserLocale operation response
type GetUserLocaleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An LCID that contains the ID of the locale for results of operations in this
	// session. If no locale is set by using IUpdateSession2::UserLocale (opnum 15) (section
	// 3.16.5.2), this MUST contain the current user's default locale as obtained from the
	// system.<17>
	ReturnValue uint32 `idl:"name:retval" json:"return_value"`
	// Return: The UserLocale return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUserLocaleResponse) xxx_ToOp(ctx context.Context, op *xxx_GetUserLocaleOperation) *xxx_GetUserLocaleOperation {
	if op == nil {
		op = &xxx_GetUserLocaleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetUserLocaleResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserLocaleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetUserLocaleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetUserLocaleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserLocaleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetUserLocaleOperation structure represents the UserLocale operation
type xxx_SetUserLocaleOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	LocaleID uint32         `idl:"name:lcid" json:"locale_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetUserLocaleOperation) OpNum() int { return 16 }

func (o *xxx_SetUserLocaleOperation) OpName() string { return "/IUpdateSession2/v0/UserLocale" }

func (o *xxx_SetUserLocaleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserLocaleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserLocaleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserLocaleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserLocaleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetUserLocaleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetUserLocaleRequest structure represents the UserLocale operation request
type SetUserLocaleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lcid: An LCID that contains the language identifier for results of operations in
	// this session.
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
}

func (o *SetUserLocaleRequest) xxx_ToOp(ctx context.Context, op *xxx_SetUserLocaleOperation) *xxx_SetUserLocaleOperation {
	if op == nil {
		op = &xxx_SetUserLocaleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LocaleID = o.LocaleID
	return op
}

func (o *SetUserLocaleRequest) xxx_FromOp(ctx context.Context, op *xxx_SetUserLocaleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LocaleID = op.LocaleID
}
func (o *SetUserLocaleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetUserLocaleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserLocaleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetUserLocaleResponse structure represents the UserLocale operation response
type SetUserLocaleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The UserLocale return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetUserLocaleResponse) xxx_ToOp(ctx context.Context, op *xxx_SetUserLocaleOperation) *xxx_SetUserLocaleOperation {
	if op == nil {
		op = &xxx_SetUserLocaleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetUserLocaleResponse) xxx_FromOp(ctx context.Context, op *xxx_SetUserLocaleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetUserLocaleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetUserLocaleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserLocaleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

package imanagetelnetsessions

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/tsrap"
)

var (
	// IManageTelnetSessions interface identifier 034634fd-ba3f-11d1-856a-00a0c944138c
	ManageTelnetSessionsIID = &dcom.IID{Data1: 0x034634fd, Data2: 0xba3f, Data3: 0x11d1, Data4: []byte{0x85, 0x6a, 0x00, 0xa0, 0xc9, 0x44, 0x13, 0x8c}}
	// Syntax UUID
	ManageTelnetSessionsSyntaxUUID = &uuid.UUID{TimeLow: 0x34634fd, TimeMid: 0xba3f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x6a, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x44, 0x13, 0x8c}}
	// Syntax ID
	ManageTelnetSessionsSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ManageTelnetSessionsSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IManageTelnetSessions interface.
type ManageTelnetSessionsClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The GetTelnetSessions method is used to query the telnet server for information about
	// all active telnet sessions.
	//
	// Return Values: The server MUST return zero if the method is successful. The server
	// MUST return 0x01 if processing fails and set output parameters to NULL. These are
	// in addition to the values that can be returned by the underlying [MS-DCOM] implementation.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// DCOM protocol [MS-DCOM].
	GetTelnetSessions(context.Context, *GetTelnetSessionsRequest, ...dcerpc.CallOption) (*GetTelnetSessionsResponse, error)

	// The TerminateSession method terminates a telnet session.
	//
	// Return Values: The server MUST return zero if the method is successful. The server
	// MUST return 0x01 if processing fails. These are in addition to the values that can
	// be returned by the underlying [MS-DCOM] implementation.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// DCOM protocol [MS-DCOM].
	TerminateSession(context.Context, *TerminateSessionRequest, ...dcerpc.CallOption) (*TerminateSessionResponse, error)

	// The SendMsgToASession method directs the telnet server to send a text message to
	// the telnet client that initiated the session.
	//
	// Return Values: The server MUST return zero if the method is successful. The server
	// MUST return 0x01 if processing fails. These are in addition to the values that can
	// be returned by the underlying [MS-DCOM] implementation.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// DCOM protocol [MS-DCOM].
	SendMessageToASession(context.Context, *SendMessageToASessionRequest, ...dcerpc.CallOption) (*SendMessageToASessionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ManageTelnetSessionsClient
}

type xxx_DefaultManageTelnetSessionsClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultManageTelnetSessionsClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultManageTelnetSessionsClient) GetTelnetSessions(ctx context.Context, in *GetTelnetSessionsRequest, opts ...dcerpc.CallOption) (*GetTelnetSessionsResponse, error) {
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
	out := &GetTelnetSessionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManageTelnetSessionsClient) TerminateSession(ctx context.Context, in *TerminateSessionRequest, opts ...dcerpc.CallOption) (*TerminateSessionResponse, error) {
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
	out := &TerminateSessionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManageTelnetSessionsClient) SendMessageToASession(ctx context.Context, in *SendMessageToASessionRequest, opts ...dcerpc.CallOption) (*SendMessageToASessionResponse, error) {
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
	out := &SendMessageToASessionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManageTelnetSessionsClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultManageTelnetSessionsClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultManageTelnetSessionsClient) IPID(ctx context.Context, ipid *dcom.IPID) ManageTelnetSessionsClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultManageTelnetSessionsClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewManageTelnetSessionsClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ManageTelnetSessionsClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ManageTelnetSessionsSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultManageTelnetSessionsClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetTelnetSessionsOperation structure represents the GetTelnetSessions operation
type xxx_GetTelnetSessionsOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	SessionData *oaut.String   `idl:"name:pszSessionData" json:"session_data"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTelnetSessionsOperation) OpNum() int { return 7 }

func (o *xxx_GetTelnetSessionsOperation) OpName() string {
	return "/IManageTelnetSessions/v0/GetTelnetSessions"
}

func (o *xxx_GetTelnetSessionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTelnetSessionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTelnetSessionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTelnetSessionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTelnetSessionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pszSessionData {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SessionData != nil {
			_ptr_pszSessionData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SessionData != nil {
					if err := o.SessionData.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SessionData, _ptr_pszSessionData); err != nil {
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

func (o *xxx_GetTelnetSessionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pszSessionData {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pszSessionData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SessionData == nil {
				o.SessionData = &oaut.String{}
			}
			if err := o.SessionData.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pszSessionData := func(ptr interface{}) { o.SessionData = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SessionData, _s_pszSessionData, _ptr_pszSessionData); err != nil {
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

// GetTelnetSessionsRequest structure represents the GetTelnetSessions operation request
type GetTelnetSessionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTelnetSessionsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTelnetSessionsOperation) *xxx_GetTelnetSessionsOperation {
	if op == nil {
		op = &xxx_GetTelnetSessionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTelnetSessionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTelnetSessionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTelnetSessionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTelnetSessionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTelnetSessionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTelnetSessionsResponse structure represents the GetTelnetSessions operation response
type GetTelnetSessionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pszSessionData: A string pointer to PSZSESSIONDATA string that contains information
	// about telnet sessions in the server.<1>
	SessionData *oaut.String `idl:"name:pszSessionData" json:"session_data"`
	// Return: The GetTelnetSessions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTelnetSessionsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTelnetSessionsOperation) *xxx_GetTelnetSessionsOperation {
	if op == nil {
		op = &xxx_GetTelnetSessionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SessionData = o.SessionData
	op.Return = o.Return
	return op
}

func (o *GetTelnetSessionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTelnetSessionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SessionData = op.SessionData
	o.Return = op.Return
}
func (o *GetTelnetSessionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTelnetSessionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTelnetSessionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TerminateSessionOperation structure represents the TerminateSession operation
type xxx_TerminateSessionOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	UniqueID uint32         `idl:"name:dwUniqueId" json:"unique_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_TerminateSessionOperation) OpNum() int { return 8 }

func (o *xxx_TerminateSessionOperation) OpName() string {
	return "/IManageTelnetSessions/v0/TerminateSession"
}

func (o *xxx_TerminateSessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateSessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwUniqueId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UniqueID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateSessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwUniqueId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UniqueID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateSessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateSessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_TerminateSessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// TerminateSessionRequest structure represents the TerminateSession operation request
type TerminateSessionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwUniqueId: The ID of the session. The ID of a session can be obtained by calling
	// the GetTelnetSessions method or can be user provided. The server MUST ensure that
	// at any given point in time only one telnet session exists with a particular ID. Refer
	// to Section 3.1.1 for an abstract data model that the server can maintain.
	UniqueID uint32 `idl:"name:dwUniqueId" json:"unique_id"`
}

func (o *TerminateSessionRequest) xxx_ToOp(ctx context.Context, op *xxx_TerminateSessionOperation) *xxx_TerminateSessionOperation {
	if op == nil {
		op = &xxx_TerminateSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.UniqueID = o.UniqueID
	return op
}

func (o *TerminateSessionRequest) xxx_FromOp(ctx context.Context, op *xxx_TerminateSessionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.UniqueID = op.UniqueID
}
func (o *TerminateSessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *TerminateSessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TerminateSessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TerminateSessionResponse structure represents the TerminateSession operation response
type TerminateSessionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The TerminateSession return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *TerminateSessionResponse) xxx_ToOp(ctx context.Context, op *xxx_TerminateSessionOperation) *xxx_TerminateSessionOperation {
	if op == nil {
		op = &xxx_TerminateSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *TerminateSessionResponse) xxx_FromOp(ctx context.Context, op *xxx_TerminateSessionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *TerminateSessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *TerminateSessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TerminateSessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendMessageToASessionOperation structure represents the SendMsgToASession operation
type xxx_SendMessageToASessionOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	UniqueID uint32         `idl:"name:dwUniqueId" json:"unique_id"`
	Message  *oaut.String   `idl:"name:szMsg" json:"message"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SendMessageToASessionOperation) OpNum() int { return 9 }

func (o *xxx_SendMessageToASessionOperation) OpName() string {
	return "/IManageTelnetSessions/v0/SendMsgToASession"
}

func (o *xxx_SendMessageToASessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageToASessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwUniqueId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UniqueID); err != nil {
			return err
		}
	}
	// szMsg {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Message != nil {
			_ptr_szMsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_szMsg); err != nil {
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

func (o *xxx_SendMessageToASessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwUniqueId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UniqueID); err != nil {
			return err
		}
	}
	// szMsg {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_szMsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &oaut.String{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_szMsg := func(ptr interface{}) { o.Message = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Message, _s_szMsg, _ptr_szMsg); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageToASessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageToASessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SendMessageToASessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SendMessageToASessionRequest structure represents the SendMsgToASession operation request
type SendMessageToASessionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwUniqueId: The ID of the session. The ID of a session can be obtained using the
	// GetTelnetSessions method or can be user provided. The server MUST ensure that at
	// any given point in time only one telnet session exists with a particular ID. Refer
	// to Section 3.1.1 for an abstract data model that the server can maintain.
	UniqueID uint32 `idl:"name:dwUniqueId" json:"unique_id"`
	// szMsg: The string text that has to be sent.
	Message *oaut.String `idl:"name:szMsg" json:"message"`
}

func (o *SendMessageToASessionRequest) xxx_ToOp(ctx context.Context, op *xxx_SendMessageToASessionOperation) *xxx_SendMessageToASessionOperation {
	if op == nil {
		op = &xxx_SendMessageToASessionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.UniqueID = o.UniqueID
	op.Message = o.Message
	return op
}

func (o *SendMessageToASessionRequest) xxx_FromOp(ctx context.Context, op *xxx_SendMessageToASessionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.UniqueID = op.UniqueID
	o.Message = op.Message
}
func (o *SendMessageToASessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendMessageToASessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendMessageToASessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendMessageToASessionResponse structure represents the SendMsgToASession operation response
type SendMessageToASessionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SendMsgToASession return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SendMessageToASessionResponse) xxx_ToOp(ctx context.Context, op *xxx_SendMessageToASessionOperation) *xxx_SendMessageToASessionOperation {
	if op == nil {
		op = &xxx_SendMessageToASessionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SendMessageToASessionResponse) xxx_FromOp(ctx context.Context, op *xxx_SendMessageToASessionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SendMessageToASessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendMessageToASessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendMessageToASessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

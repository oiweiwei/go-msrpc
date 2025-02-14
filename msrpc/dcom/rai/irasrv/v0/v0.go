package irasrv

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
	GoPackage = "dcom/rai"
)

var (
	// IRASrv interface identifier f120a684-b926-447f-9df4-c966cb785648
	RemoteAssistanceServerIID = &dcom.IID{Data1: 0xf120a684, Data2: 0xb926, Data3: 0x447f, Data4: []byte{0x9d, 0xf4, 0xc9, 0x66, 0xcb, 0x78, 0x56, 0x48}}
	// Syntax UUID
	RemoteAssistanceServerSyntaxUUID = &uuid.UUID{TimeLow: 0xf120a684, TimeMid: 0xb926, TimeHiAndVersion: 0x447f, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xf4, Node: [6]uint8{0xc9, 0x66, 0xcb, 0x78, 0x56, 0x48}}
	// Syntax ID
	RemoteAssistanceServerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteAssistanceServerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRASrv interface.
type RemoteAssistanceServerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The GetNoviceUserInfo method is received by the server in an RPC_REQUEST packet.
	// The method is received in the terminal services session as specified by the Client.
	// In response, the server returns the Remote Assistance Connection String 2 for the
	// specified terminal services session.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure.
	//
	//	+------------------------------+------------------------------------------------------------------------------+
	//	|            RETURN            |                                                                              |
	//	|          VALUE/CODE          |                                 DESCRIPTION                                  |
	//	|                              |                                                                              |
	//	+------------------------------+------------------------------------------------------------------------------+
	//	+------------------------------+------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK              | The call was successful.                                                     |
	//	+------------------------------+------------------------------------------------------------------------------+
	//	| 0x80004003 E_POINTER         | The method failed due to an invalid pointer for szName.                      |
	//	+------------------------------+------------------------------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY     | Out of memory.                                                               |
	//	+------------------------------+------------------------------------------------------------------------------+
	//	| 0x80070052 ERROR_CANNOT_MAKE | An instance of Remote Assistance is already running on the novice's machine. |
	//	+------------------------------+------------------------------------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	GetNoviceUserInfo(context.Context, *GetNoviceUserInfoRequest, ...dcerpc.CallOption) (*GetNoviceUserInfoResponse, error)

	// The GetSessionInfo method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns the terminal services session information for the various
	// sessions on the computer. The terminal services session information is returned as
	// a SAFEARRAY of BSTRs. Each BSTR contains the DomainName, UserName and SessionID in
	// the format DomainName\UserName:SessionID.
	//
	// This method also returns the count of the total number of sessions.
	//
	// This method does not return Idle and Disconnected terminal services sessions. Any
	// null values returned in the SAFEARRAY can be ignored.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure.
	//
	//	+--------------------------+----------------------------------------------------+
	//	|          RETURN          |                                                    |
	//	|        VALUE/CODE        |                    DESCRIPTION                     |
	//	|                          |                                                    |
	//	+--------------------------+----------------------------------------------------+
	//	+--------------------------+----------------------------------------------------+
	//	| 0x00000000 S_OK          | The call was successful.                           |
	//	+--------------------------+----------------------------------------------------+
	//	| 0x80004003 E_POINTER     | The method failed due to an invalid pointer.       |
	//	+--------------------------+----------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY | The method was unable to allocate required memory. |
	//	+--------------------------+----------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	GetSessionInfo(context.Context, *GetSessionInfoRequest, ...dcerpc.CallOption) (*GetSessionInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteAssistanceServerClient
}

type xxx_DefaultRemoteAssistanceServerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteAssistanceServerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultRemoteAssistanceServerClient) GetNoviceUserInfo(ctx context.Context, in *GetNoviceUserInfoRequest, opts ...dcerpc.CallOption) (*GetNoviceUserInfoResponse, error) {
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
	out := &GetNoviceUserInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteAssistanceServerClient) GetSessionInfo(ctx context.Context, in *GetSessionInfoRequest, opts ...dcerpc.CallOption) (*GetSessionInfoResponse, error) {
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
	out := &GetSessionInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteAssistanceServerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteAssistanceServerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteAssistanceServerClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteAssistanceServerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteAssistanceServerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewRemoteAssistanceServerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteAssistanceServerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteAssistanceServerSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteAssistanceServerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetNoviceUserInfoOperation structure represents the GetNoviceUserInfo operation
type xxx_GetNoviceUserInfoOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   string         `idl:"name:szName" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNoviceUserInfoOperation) OpNum() int { return 7 }

func (o *xxx_GetNoviceUserInfoOperation) OpName() string { return "/IRASrv/v0/GetNoviceUserInfo" }

func (o *xxx_GetNoviceUserInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNoviceUserInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szName {in, out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string](wchar))
	{
		if o.Name != "" {
			_ptr_szName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.Name); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_szName); err != nil {
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

func (o *xxx_GetNoviceUserInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szName {in, out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		_ptr_szName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.Name); err != nil {
				return err
			}
			return nil
		})
		_s_szName := func(ptr interface{}) { o.Name = *ptr.(*string) }
		if err := w.ReadPointer(&o.Name, _s_szName, _ptr_szName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNoviceUserInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNoviceUserInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// szName {in, out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string](wchar))
	{
		if o.Name != "" {
			_ptr_szName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.Name); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_szName); err != nil {
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

func (o *xxx_GetNoviceUserInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// szName {in, out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		_ptr_szName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.Name); err != nil {
				return err
			}
			return nil
		})
		_s_szName := func(ptr interface{}) { o.Name = *ptr.(*string) }
		if err := w.ReadPointer(&o.Name, _s_szName, _ptr_szName); err != nil {
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

// GetNoviceUserInfoRequest structure represents the GetNoviceUserInfo operation request
type GetNoviceUserInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szName: A pointer to a NULL-terminated Unicode string that contains the Remote Assistance
	// Connection String 2 for the specified terminal services session.
	Name string `idl:"name:szName" json:"name"`
}

func (o *GetNoviceUserInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNoviceUserInfoOperation) *xxx_GetNoviceUserInfoOperation {
	if op == nil {
		op = &xxx_GetNoviceUserInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Name = op.Name
	return op
}

func (o *GetNoviceUserInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNoviceUserInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *GetNoviceUserInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNoviceUserInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNoviceUserInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNoviceUserInfoResponse structure represents the GetNoviceUserInfo operation response
type GetNoviceUserInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// szName: A pointer to a NULL-terminated Unicode string that contains the Remote Assistance
	// Connection String 2 for the specified terminal services session.
	Name string `idl:"name:szName" json:"name"`
	// Return: The GetNoviceUserInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNoviceUserInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNoviceUserInfoOperation) *xxx_GetNoviceUserInfoOperation {
	if op == nil {
		op = &xxx_GetNoviceUserInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
	return op
}

func (o *GetNoviceUserInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNoviceUserInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetNoviceUserInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNoviceUserInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNoviceUserInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSessionInfoOperation structure represents the GetSessionInfo operation
type xxx_GetSessionInfoOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	UserNames *oaut.SafeArray `idl:"name:UserNames" json:"user_names"`
	Count     int32           `idl:"name:Count" json:"count"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSessionInfoOperation) OpNum() int { return 8 }

func (o *xxx_GetSessionInfoOperation) OpName() string { return "/IRASrv/v0/GetSessionInfo" }

func (o *xxx_GetSessionInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// UserNames {in, out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.UserNames != nil {
			_ptr_UserNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserNames != nil {
					if err := o.UserNames.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserNames, _ptr_UserNames); err != nil {
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
	// Count {in, out} (1:{pointer=ref}*(1))(2:{alias=INT}(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// UserNames {in, out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_UserNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserNames == nil {
				o.UserNames = &oaut.SafeArray{}
			}
			if err := o.UserNames.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserNames := func(ptr interface{}) { o.UserNames = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.UserNames, _s_UserNames, _ptr_UserNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Count {in, out} (1:{pointer=ref}*(1))(2:{alias=INT}(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// UserNames {in, out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.UserNames != nil {
			_ptr_UserNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserNames != nil {
					if err := o.UserNames.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserNames, _ptr_UserNames); err != nil {
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
	// Count {in, out} (1:{pointer=ref}*(1))(2:{alias=INT}(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
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

func (o *xxx_GetSessionInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// UserNames {in, out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_UserNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserNames == nil {
				o.UserNames = &oaut.SafeArray{}
			}
			if err := o.UserNames.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserNames := func(ptr interface{}) { o.UserNames = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.UserNames, _s_UserNames, _ptr_UserNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Count {in, out} (1:{pointer=ref}*(1))(2:{alias=INT}(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
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

// GetSessionInfoRequest structure represents the GetSessionInfo operation request
type GetSessionInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// UserNames: A pointer to a SAFEARRAY, as specified in [MS-OAUT], of BSTRs containing
	// the terminal services session information. Each BSTR element in the array contains
	// the DomainName, UserName, and SessionID in the format DomainName\UserName:SessionID.
	// This is returned to the expert.
	UserNames *oaut.SafeArray `idl:"name:UserNames" json:"user_names"`
	// Count: A pointer to an INT that returns the number of terminal services sessions
	// on the novice.
	Count int32 `idl:"name:Count" json:"count"`
}

func (o *GetSessionInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSessionInfoOperation) *xxx_GetSessionInfoOperation {
	if op == nil {
		op = &xxx_GetSessionInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.UserNames = op.UserNames
	o.Count = op.Count
	return op
}

func (o *GetSessionInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSessionInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.UserNames = op.UserNames
	o.Count = op.Count
}
func (o *GetSessionInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSessionInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSessionInfoResponse structure represents the GetSessionInfo operation response
type GetSessionInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// UserNames: A pointer to a SAFEARRAY, as specified in [MS-OAUT], of BSTRs containing
	// the terminal services session information. Each BSTR element in the array contains
	// the DomainName, UserName, and SessionID in the format DomainName\UserName:SessionID.
	// This is returned to the expert.
	UserNames *oaut.SafeArray `idl:"name:UserNames" json:"user_names"`
	// Count: A pointer to an INT that returns the number of terminal services sessions
	// on the novice.
	Count int32 `idl:"name:Count" json:"count"`
	// Return: The GetSessionInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSessionInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSessionInfoOperation) *xxx_GetSessionInfoOperation {
	if op == nil {
		op = &xxx_GetSessionInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.UserNames = op.UserNames
	o.Count = op.Count
	o.Return = op.Return
	return op
}

func (o *GetSessionInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSessionInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UserNames = op.UserNames
	o.Count = op.Count
	o.Return = op.Return
}
func (o *GetSessionInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSessionInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

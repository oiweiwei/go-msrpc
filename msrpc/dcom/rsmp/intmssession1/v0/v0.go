package intmssession1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsSession1 interface identifier 8da03f40-3419-11d1-8fb1-00a024cb6019
	Session1IID = &dcom.IID{Data1: 0x8da03f40, Data2: 0x3419, Data3: 0x11d1, Data4: []byte{0x8f, 0xb1, 0x00, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax UUID
	Session1SyntaxUUID = &uuid.UUID{TimeLow: 0x8da03f40, TimeMid: 0x3419, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xb1, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	// Syntax ID
	Session1SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Session1SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsSession1 interface.
type Session1Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The OpenNtmsServerSessionW method sets up a session with the server. All input parameters
	// for this method are optional.
	OpenNTMSServerSessionW(context.Context, *OpenNTMSServerSessionWRequest, ...dcerpc.CallOption) (*OpenNTMSServerSessionWResponse, error)

	// The OpenNtmsServerSessionA method sets up a session with the server. All input parameters
	// for this method are optional.
	OpenNTMSServerSessionA(context.Context, *OpenNTMSServerSessionARequest, ...dcerpc.CallOption) (*OpenNTMSServerSessionAResponse, error)

	// The CloseNtmsSession method closes a session.
	//
	// This method has no parameters.
	//
	//	+-------------------------------------+------------------------------------------------+
	//	|               RETURN                |                                                |
	//	|             VALUE/CODE              |                  DESCRIPTION                   |
	//	|                                     |                                                |
	//	+-------------------------------------+------------------------------------------------+
	//	+-------------------------------------+------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                       |
	//	+-------------------------------------+------------------------------------------------+
	//	| 0x800704B1 ERROR_CONNECTION_UNAVAIL | The connection to the server is not available. |
	//	+-------------------------------------+------------------------------------------------+
	//
	// This method will close the session. This method MUST be called when the client has
	// dereferenced the session. Note, however, that the server might not yet have dereferenced
	// the session, because one or more interfaces might still be in process. The server,
	// therefore, does not tear down the session, but rather cleans up the session and releases
	// the open session reference, allowing the session to be ended when the last reference
	// is removed.
	//
	// If the client has an outstanding synchronous request, the requests MUST be unwound
	// and canceled.
	CloseNTMSSession(context.Context, *CloseNTMSSessionRequest, ...dcerpc.CallOption) (*CloseNTMSSessionResponse, error)

	// The SubmitNtmsOperatorRequestW method submits an operator request, with strings encoded
	// using Unicode.
	SubmitNTMSOperatorRequestW(context.Context, *SubmitNTMSOperatorRequestWRequest, ...dcerpc.CallOption) (*SubmitNTMSOperatorRequestWResponse, error)

	// The SubmitNtmsOperatorRequestA method submits an operator request, with strings encoded
	// using ASCII.
	SubmitNTMSOperatorRequestA(context.Context, *SubmitNTMSOperatorRequestARequest, ...dcerpc.CallOption) (*SubmitNTMSOperatorRequestAResponse, error)

	// The WaitForNtmsOperatorRequest method waits for an operator request.
	WaitForNTMSOperatorRequest(context.Context, *WaitForNTMSOperatorRequestRequest, ...dcerpc.CallOption) (*WaitForNTMSOperatorRequestResponse, error)

	// The CancelNtmsOperatorRequest method cancels an operator request.
	CancelNTMSOperatorRequest(context.Context, *CancelNTMSOperatorRequestRequest, ...dcerpc.CallOption) (*CancelNTMSOperatorRequestResponse, error)

	// The SatisfyNtmsOperatorRequest method completes an operator request.
	SatisfyNTMSOperatorRequest(context.Context, *SatisfyNTMSOperatorRequestRequest, ...dcerpc.CallOption) (*SatisfyNTMSOperatorRequestResponse, error)

	// The ImportNtmsDatabase method imports the state of the server at the next server
	// restart.
	//
	// This method has no parameters.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                   | The call was successful.                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED    | Access to object was denied. For more information regarding object, see section  |
	//	|                                   | 3.2.5.2.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED    | Unable to connect to the server.                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE | The database query or update failed.                                             |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// The ImportNtmsDatabase method directs the server to import the content of the export
	// directory into the database upon the next restart. When the server restarts, it MUST
	// import the already saved database to its current database if the import is set.
	ImportNTMSDatabase(context.Context, *ImportNTMSDatabaseRequest, ...dcerpc.CallOption) (*ImportNTMSDatabaseResponse, error)

	// The ExportNtmsDatabase method exports the state of the server.
	//
	// This method has no parameters.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to object was denied. For more information regarding object, see section  |
	//	|                                    | 3.2.5.2.                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070020 ERROR_SHARING_VIOLATION | One of the files that needs to be written to is open.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED     | Unable to connect to the server.                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database query or update failed.                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	ExportNTMSDatabase(context.Context, *ExportNTMSDatabaseRequest, ...dcerpc.CallOption) (*ExportNTMSDatabaseResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// The AddNotification method registers a client to receive change notifications for
	// a type of object.
	AddNotification(context.Context, *AddNotificationRequest, ...dcerpc.CallOption) (*AddNotificationResponse, error)

	// The RemoveNotification method unregisters a client from receiving change notifications
	// for a type of object.
	RemoveNotification(context.Context, *RemoveNotificationRequest, ...dcerpc.CallOption) (*RemoveNotificationResponse, error)

	// The DispatchNotification method sends a notification to all registered sinks.
	DispatchNotification(context.Context, *DispatchNotificationRequest, ...dcerpc.CallOption) (*DispatchNotificationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Session1Client
}

type xxx_DefaultSession1Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultSession1Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultSession1Client) OpenNTMSServerSessionW(ctx context.Context, in *OpenNTMSServerSessionWRequest, opts ...dcerpc.CallOption) (*OpenNTMSServerSessionWResponse, error) {
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
	out := &OpenNTMSServerSessionWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) OpenNTMSServerSessionA(ctx context.Context, in *OpenNTMSServerSessionARequest, opts ...dcerpc.CallOption) (*OpenNTMSServerSessionAResponse, error) {
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
	out := &OpenNTMSServerSessionAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) CloseNTMSSession(ctx context.Context, in *CloseNTMSSessionRequest, opts ...dcerpc.CallOption) (*CloseNTMSSessionResponse, error) {
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
	out := &CloseNTMSSessionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) SubmitNTMSOperatorRequestW(ctx context.Context, in *SubmitNTMSOperatorRequestWRequest, opts ...dcerpc.CallOption) (*SubmitNTMSOperatorRequestWResponse, error) {
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
	out := &SubmitNTMSOperatorRequestWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) SubmitNTMSOperatorRequestA(ctx context.Context, in *SubmitNTMSOperatorRequestARequest, opts ...dcerpc.CallOption) (*SubmitNTMSOperatorRequestAResponse, error) {
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
	out := &SubmitNTMSOperatorRequestAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) WaitForNTMSOperatorRequest(ctx context.Context, in *WaitForNTMSOperatorRequestRequest, opts ...dcerpc.CallOption) (*WaitForNTMSOperatorRequestResponse, error) {
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
	out := &WaitForNTMSOperatorRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) CancelNTMSOperatorRequest(ctx context.Context, in *CancelNTMSOperatorRequestRequest, opts ...dcerpc.CallOption) (*CancelNTMSOperatorRequestResponse, error) {
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
	out := &CancelNTMSOperatorRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) SatisfyNTMSOperatorRequest(ctx context.Context, in *SatisfyNTMSOperatorRequestRequest, opts ...dcerpc.CallOption) (*SatisfyNTMSOperatorRequestResponse, error) {
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
	out := &SatisfyNTMSOperatorRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) ImportNTMSDatabase(ctx context.Context, in *ImportNTMSDatabaseRequest, opts ...dcerpc.CallOption) (*ImportNTMSDatabaseResponse, error) {
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
	out := &ImportNTMSDatabaseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) ExportNTMSDatabase(ctx context.Context, in *ExportNTMSDatabaseRequest, opts ...dcerpc.CallOption) (*ExportNTMSDatabaseResponse, error) {
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
	out := &ExportNTMSDatabaseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) AddNotification(ctx context.Context, in *AddNotificationRequest, opts ...dcerpc.CallOption) (*AddNotificationResponse, error) {
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
	out := &AddNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) RemoveNotification(ctx context.Context, in *RemoveNotificationRequest, opts ...dcerpc.CallOption) (*RemoveNotificationResponse, error) {
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
	out := &RemoveNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) DispatchNotification(ctx context.Context, in *DispatchNotificationRequest, opts ...dcerpc.CallOption) (*DispatchNotificationResponse, error) {
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
	out := &DispatchNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSession1Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSession1Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultSession1Client) IPID(ctx context.Context, ipid *dcom.IPID) Session1Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultSession1Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewSession1Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Session1Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Session1SyntaxV0_0))...)
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
	return &xxx_DefaultSession1Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_OpenNTMSServerSessionWOperation structure represents the OpenNtmsServerSessionW operation
type xxx_OpenNTMSServerSessionWOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Server      string         `idl:"name:lpServer;string;pointer:unique" json:"server"`
	Application string         `idl:"name:lpApplication;string;pointer:unique" json:"application"`
	ClientName  string         `idl:"name:lpClientName;string" json:"client_name"`
	UserName    string         `idl:"name:lpUserName;string" json:"user_name"`
	Options     uint32         `idl:"name:dwOptions" json:"options"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenNTMSServerSessionWOperation) OpNum() int { return 0 }

func (o *xxx_OpenNTMSServerSessionWOperation) OpName() string {
	return "/INtmsSession1/v0/OpenNtmsServerSessionW"
}

func (o *xxx_OpenNTMSServerSessionWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpServer {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_lpServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_lpServer); err != nil {
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
	// lpApplication {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Application != "" {
			_ptr_lpApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Application); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Application, _ptr_lpApplication); err != nil {
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
	// lpClientName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
			return err
		}
	}
	// lpUserName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpServer {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_lpServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_lpServer, _ptr_lpServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpApplication {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Application); err != nil {
				return err
			}
			return nil
		})
		_s_lpApplication := func(ptr interface{}) { o.Application = *ptr.(*string) }
		if err := w.ReadPointer(&o.Application, _s_lpApplication, _ptr_lpApplication); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpClientName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
	}
	// lpUserName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OpenNTMSServerSessionWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OpenNTMSServerSessionWRequest structure represents the OpenNtmsServerSessionW operation request
type OpenNTMSServerSessionWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpServer: The null-terminated Domain Name System (DNS) or Windows Internet Naming
	// Service (WINS) name of the computer with which to set up the session. If this parameter
	// is NULL, the current computer name MUST be used.
	Server string `idl:"name:lpServer;string;pointer:unique" json:"server"`
	// lpApplication: The null-terminated unique character string that identifies the application.
	// This name identifies resources and operator requests, and is optional.<76>
	Application string `idl:"name:lpApplication;string;pointer:unique" json:"application"`
	// lpClientName: The null-terminated DNS or WINS name of the computer sending the request.
	ClientName string `idl:"name:lpClientName;string" json:"client_name"`
	// lpUserName: A null-terminated sequence of Unicode characters specifying the name
	// of the interactive user sending the request.
	UserName string `idl:"name:lpUserName;string" json:"user_name"`
	// dwOptions: This parameter is unused. It MUST be sent as 0 and MUST be ignored on
	// receipt.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED        | Access to an object was denied.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070015 ERROR_NOT_READY            | The service has not started. The application is to wait and retry its request.   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704BA ERROR_INVALID_COMPUTERNAME | The supplied computer name format is invalid.                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070719 RPC_S_NO_INTERFACES        | The server is using a version of RSM that is older than the version that is used |
	//	|                                       | by the client.                                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
}

func (o *OpenNTMSServerSessionWRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenNTMSServerSessionWOperation) *xxx_OpenNTMSServerSessionWOperation {
	if op == nil {
		op = &xxx_OpenNTMSServerSessionWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Server = o.Server
	op.Application = o.Application
	op.ClientName = o.ClientName
	op.UserName = o.UserName
	op.Options = o.Options
	return op
}

func (o *OpenNTMSServerSessionWRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenNTMSServerSessionWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Server = op.Server
	o.Application = op.Application
	o.ClientName = op.ClientName
	o.UserName = op.UserName
	o.Options = op.Options
}
func (o *OpenNTMSServerSessionWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenNTMSServerSessionWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNTMSServerSessionWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenNTMSServerSessionWResponse structure represents the OpenNtmsServerSessionW operation response
type OpenNTMSServerSessionWResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OpenNtmsServerSessionW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenNTMSServerSessionWResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenNTMSServerSessionWOperation) *xxx_OpenNTMSServerSessionWOperation {
	if op == nil {
		op = &xxx_OpenNTMSServerSessionWOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OpenNTMSServerSessionWResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenNTMSServerSessionWOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OpenNTMSServerSessionWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenNTMSServerSessionWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNTMSServerSessionWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenNTMSServerSessionAOperation structure represents the OpenNtmsServerSessionA operation
type xxx_OpenNTMSServerSessionAOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Server      string         `idl:"name:lpServer;pointer:unique" json:"server"`
	Application string         `idl:"name:lpApplication;pointer:unique" json:"application"`
	ClientName  uint8          `idl:"name:lpClientName" json:"client_name"`
	UserName    uint8          `idl:"name:lpUserName" json:"user_name"`
	Options     uint32         `idl:"name:dwOptions" json:"options"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenNTMSServerSessionAOperation) OpNum() int { return 1 }

func (o *xxx_OpenNTMSServerSessionAOperation) OpName() string {
	return "/INtmsSession1/v0/OpenNtmsServerSessionA"
}

func (o *xxx_OpenNTMSServerSessionAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpServer {in} (1:{pointer=unique}*(1)[dim:0,string](char))
	{
		if o.Server != "" {
			_ptr_lpServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_lpServer); err != nil {
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
	// lpApplication {in} (1:{pointer=unique}*(1)[dim:0,string](char))
	{
		if o.Application != "" {
			_ptr_lpApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharString(ctx, w, o.Application); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Application, _ptr_lpApplication); err != nil {
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
	// lpClientName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.WriteData(o.ClientName); err != nil {
			return err
		}
	}
	// lpUserName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.WriteData(o.UserName); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpServer {in} (1:{pointer=unique}*(1)[dim:0,string](char))
	{
		_ptr_lpServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_lpServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_lpServer, _ptr_lpServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpApplication {in} (1:{pointer=unique}*(1)[dim:0,string](char))
	{
		_ptr_lpApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharString(ctx, w, &o.Application); err != nil {
				return err
			}
			return nil
		})
		_s_lpApplication := func(ptr interface{}) { o.Application = *ptr.(*string) }
		if err := w.ReadPointer(&o.Application, _s_lpApplication, _ptr_lpApplication); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpClientName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.ReadData(&o.ClientName); err != nil {
			return err
		}
	}
	// lpUserName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.ReadData(&o.UserName); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNTMSServerSessionAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OpenNTMSServerSessionAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OpenNTMSServerSessionARequest structure represents the OpenNtmsServerSessionA operation request
type OpenNTMSServerSessionARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpServer: The null-terminated DNS or WINS name of the computer with which to set
	// up the session. If this parameter is NULL, the current computer name MUST be used.
	Server string `idl:"name:lpServer;pointer:unique" json:"server"`
	// lpApplication: The null-terminated unique character string that identifies the application.
	// This name identifies resources and requests made by the operator. It is optional
	// and can be NULL.
	Application string `idl:"name:lpApplication;pointer:unique" json:"application"`
	// lpClientName: The null-terminated DNS or WINS name of the computer sending the request.
	ClientName uint8 `idl:"name:lpClientName" json:"client_name"`
	// lpUserName: A null-terminated sequence of characters specifying the name of the interactive
	// user sending the request.
	UserName uint8 `idl:"name:lpUserName" json:"user_name"`
	// dwOptions: This parameter is unused. It MUST be sent as 0 and MUST be ignored on
	// receipt.
	//
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                |
	//	|              VALUE/CODE               |                                  DESCRIPTION                                   |
	//	|                                       |                                                                                |
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                       |
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED        | Access to an object was denied.                                                |
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070015 ERROR_NOT_READY            | The service has not started. The application is to wait and retry its request. |
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x800704BA ERROR_INVALID_COMPUTERNAME | The format of the supplied computer name is invalid.                           |
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070719 RPC_S_NO_INTERFACES        | The server is using an older version of RSM than that of the client.           |
	//	+---------------------------------------+--------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
}

func (o *OpenNTMSServerSessionARequest) xxx_ToOp(ctx context.Context, op *xxx_OpenNTMSServerSessionAOperation) *xxx_OpenNTMSServerSessionAOperation {
	if op == nil {
		op = &xxx_OpenNTMSServerSessionAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Server = o.Server
	op.Application = o.Application
	op.ClientName = o.ClientName
	op.UserName = o.UserName
	op.Options = o.Options
	return op
}

func (o *OpenNTMSServerSessionARequest) xxx_FromOp(ctx context.Context, op *xxx_OpenNTMSServerSessionAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Server = op.Server
	o.Application = op.Application
	o.ClientName = op.ClientName
	o.UserName = op.UserName
	o.Options = op.Options
}
func (o *OpenNTMSServerSessionARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenNTMSServerSessionARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNTMSServerSessionAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenNTMSServerSessionAResponse structure represents the OpenNtmsServerSessionA operation response
type OpenNTMSServerSessionAResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OpenNtmsServerSessionA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenNTMSServerSessionAResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenNTMSServerSessionAOperation) *xxx_OpenNTMSServerSessionAOperation {
	if op == nil {
		op = &xxx_OpenNTMSServerSessionAOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OpenNTMSServerSessionAResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenNTMSServerSessionAOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OpenNTMSServerSessionAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenNTMSServerSessionAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNTMSServerSessionAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseNTMSSessionOperation structure represents the CloseNtmsSession operation
type xxx_CloseNTMSSessionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseNTMSSessionOperation) OpNum() int { return 2 }

func (o *xxx_CloseNTMSSessionOperation) OpName() string { return "/INtmsSession1/v0/CloseNtmsSession" }

func (o *xxx_CloseNTMSSessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNTMSSessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseNTMSSessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CloseNTMSSessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNTMSSessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseNTMSSessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CloseNTMSSessionRequest structure represents the CloseNtmsSession operation request
type CloseNTMSSessionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CloseNTMSSessionRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseNTMSSessionOperation) *xxx_CloseNTMSSessionOperation {
	if op == nil {
		op = &xxx_CloseNTMSSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CloseNTMSSessionRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseNTMSSessionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CloseNTMSSessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseNTMSSessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNTMSSessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseNTMSSessionResponse structure represents the CloseNtmsSession operation response
type CloseNTMSSessionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CloseNtmsSession return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseNTMSSessionResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseNTMSSessionOperation) *xxx_CloseNTMSSessionOperation {
	if op == nil {
		op = &xxx_CloseNTMSSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CloseNTMSSessionResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseNTMSSessionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CloseNTMSSessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseNTMSSessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNTMSSessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SubmitNTMSOperatorRequestWOperation structure represents the SubmitNtmsOperatorRequestW operation
type xxx_SubmitNTMSOperatorRequestWOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Request   uint32         `idl:"name:dwRequest" json:"request"`
	Message   string         `idl:"name:lpMessage;string;pointer:unique" json:"message"`
	Arg1ID    *dtyp.GUID     `idl:"name:lpArg1Id;pointer:unique" json:"arg1_id"`
	Arg2ID    *dtyp.GUID     `idl:"name:lpArg2Id;pointer:unique" json:"arg2_id"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SubmitNTMSOperatorRequestWOperation) OpNum() int { return 3 }

func (o *xxx_SubmitNTMSOperatorRequestWOperation) OpName() string {
	return "/INtmsSession1/v0/SubmitNtmsOperatorRequestW"
}

func (o *xxx_SubmitNTMSOperatorRequestWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitNTMSOperatorRequestWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwRequest {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Request); err != nil {
			return err
		}
	}
	// lpMessage {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Message != "" {
			_ptr_lpMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Message); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_lpMessage); err != nil {
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
	// lpArg1Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Arg1ID != nil {
			_ptr_lpArg1Id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Arg1ID != nil {
					if err := o.Arg1ID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Arg1ID, _ptr_lpArg1Id); err != nil {
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
	// lpArg2Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Arg2ID != nil {
			_ptr_lpArg2Id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Arg2ID != nil {
					if err := o.Arg2ID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Arg2ID, _ptr_lpArg2Id); err != nil {
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

func (o *xxx_SubmitNTMSOperatorRequestWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwRequest {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Request); err != nil {
			return err
		}
	}
	// lpMessage {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Message); err != nil {
				return err
			}
			return nil
		})
		_s_lpMessage := func(ptr interface{}) { o.Message = *ptr.(*string) }
		if err := w.ReadPointer(&o.Message, _s_lpMessage, _ptr_lpMessage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpArg1Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpArg1Id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Arg1ID == nil {
				o.Arg1ID = &dtyp.GUID{}
			}
			if err := o.Arg1ID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpArg1Id := func(ptr interface{}) { o.Arg1ID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Arg1ID, _s_lpArg1Id, _ptr_lpArg1Id); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpArg2Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpArg2Id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Arg2ID == nil {
				o.Arg2ID = &dtyp.GUID{}
			}
			if err := o.Arg2ID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpArg2Id := func(ptr interface{}) { o.Arg2ID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Arg2ID, _s_lpArg2Id, _ptr_lpArg2Id); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitNTMSOperatorRequestWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitNTMSOperatorRequestWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SubmitNTMSOperatorRequestWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
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

// SubmitNTMSOperatorRequestWRequest structure represents the SubmitNtmsOperatorRequestW operation request
type SubmitNTMSOperatorRequestWRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwRequest: A value from the NtmsOpreqCommand (section 2.2.1.7) enumeration specifying
	// the type of operation requested. If dwRequest is set to NTMS_OPREQ_UNKNOWN, the server
	// MUST return ERROR_INVALID_PARAMETER and take no action.
	Request uint32 `idl:"name:dwRequest" json:"request"`
	// lpMessage: An optional null-terminated message string to send to the user.
	Message string `idl:"name:lpMessage;string;pointer:unique" json:"message"`
	// lpArg1Id: This parameter MUST be set according to the value of dwRequest.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|              VALUE OF               |                                                                                  |
	//	|              DWREQUEST              |                                     LPARG1ID                                     |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_NEWMEDIA 0x00000001      | MUST be set to a pointer to the identifier of the media pool from which to       |
	//	|                                     | retrieve the new media.                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_CLEANER 0x00000002       | Must be set to the pointer to the identifier of the device (library or slot).    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_DEVICESERVICE 0x00000003 | MUST be set to a pointer to the identifier of the device that needs service.     |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MOVEMEDIA 0x00000004     | MUST be set to a pointer to the identifier of the piece of physical media to     |
	//	|                                     | move.                                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MESSAGE 0x00000005       | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Arg1ID *dtyp.GUID `idl:"name:lpArg1Id;pointer:unique" json:"arg1_id"`
	// lpArg2Id: This parameter MUST be set according to the value of dwRequest.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|              VALUE OF               |                                                                                  |
	//	|              DWREQUEST              |                                     LPARG2ID                                     |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_NEWMEDIA 0x00000001      | MUST be set to a pointer to the identifier of the media library to which the new |
	//	|                                     | media MUST be added.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_CLEANER 0x00000002       | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_DEVICESERVICE 0x00000003 | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MOVEMEDIA 0x00000004     | MUST be set to a pointer to the identifier of the target library to which the    |
	//	|                                     | media MUST be moved.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MESSAGE 0x00000005       | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Arg2ID *dtyp.GUID `idl:"name:lpArg2Id;pointer:unique" json:"arg2_id"`
}

func (o *SubmitNTMSOperatorRequestWRequest) xxx_ToOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestWOperation) *xxx_SubmitNTMSOperatorRequestWOperation {
	if op == nil {
		op = &xxx_SubmitNTMSOperatorRequestWOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Request = o.Request
	op.Message = o.Message
	op.Arg1ID = o.Arg1ID
	op.Arg2ID = o.Arg2ID
	return op
}

func (o *SubmitNTMSOperatorRequestWRequest) xxx_FromOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestWOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Request = op.Request
	o.Message = op.Message
	o.Arg1ID = op.Arg1ID
	o.Arg2ID = op.Arg2ID
}
func (o *SubmitNTMSOperatorRequestWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SubmitNTMSOperatorRequestWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitNTMSOperatorRequestWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SubmitNTMSOperatorRequestWResponse structure represents the SubmitNtmsOperatorRequestW operation response
type SubmitNTMSOperatorRequestWResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpRequestId: A pointer to the identifier of the created operator request.
	//
	//	+------------------------------------+--------------------------------------------------+
	//	|               RETURN               |                                                  |
	//	|             VALUE/CODE             |                   DESCRIPTION                    |
	//	|                                    |                                                  |
	//	+------------------------------------+--------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                         |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to one or more objects is denied.         |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | An invalid parameter was found.                  |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED     | Unable to connect to the service.                |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | Unable to find the source or destination object. |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database query or update failed.             |
	//	+------------------------------------+--------------------------------------------------+
	//
	// The SubmitNtmsOperatorRequestW method submits an operator request, and returns the
	// status of the request (Satisfied or Canceled), or times out (if the operator does
	// not act upon the request). Operator requests are used to request media, to request
	// that the specified medium be moved from one library to another, or to request the
	// server device service.
	RequestID *dtyp.GUID `idl:"name:lpRequestId" json:"request_id"`
	// Return: The SubmitNtmsOperatorRequestW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SubmitNTMSOperatorRequestWResponse) xxx_ToOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestWOperation) *xxx_SubmitNTMSOperatorRequestWOperation {
	if op == nil {
		op = &xxx_SubmitNTMSOperatorRequestWOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RequestID = o.RequestID
	op.Return = o.Return
	return op
}

func (o *SubmitNTMSOperatorRequestWResponse) xxx_FromOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestWOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RequestID = op.RequestID
	o.Return = op.Return
}
func (o *SubmitNTMSOperatorRequestWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SubmitNTMSOperatorRequestWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitNTMSOperatorRequestWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SubmitNTMSOperatorRequestAOperation structure represents the SubmitNtmsOperatorRequestA operation
type xxx_SubmitNTMSOperatorRequestAOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Request   uint32         `idl:"name:dwRequest" json:"request"`
	Message   string         `idl:"name:lpMessage;pointer:unique" json:"message"`
	Arg1ID    *dtyp.GUID     `idl:"name:lpArg1Id;pointer:unique" json:"arg1_id"`
	Arg2ID    *dtyp.GUID     `idl:"name:lpArg2Id;pointer:unique" json:"arg2_id"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SubmitNTMSOperatorRequestAOperation) OpNum() int { return 4 }

func (o *xxx_SubmitNTMSOperatorRequestAOperation) OpName() string {
	return "/INtmsSession1/v0/SubmitNtmsOperatorRequestA"
}

func (o *xxx_SubmitNTMSOperatorRequestAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitNTMSOperatorRequestAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwRequest {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Request); err != nil {
			return err
		}
	}
	// lpMessage {in} (1:{pointer=unique}*(1)[dim:0,string](char))
	{
		if o.Message != "" {
			_ptr_lpMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharString(ctx, w, o.Message); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_lpMessage); err != nil {
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
	// lpArg1Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Arg1ID != nil {
			_ptr_lpArg1Id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Arg1ID != nil {
					if err := o.Arg1ID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Arg1ID, _ptr_lpArg1Id); err != nil {
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
	// lpArg2Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.Arg2ID != nil {
			_ptr_lpArg2Id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Arg2ID != nil {
					if err := o.Arg2ID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Arg2ID, _ptr_lpArg2Id); err != nil {
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

func (o *xxx_SubmitNTMSOperatorRequestAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwRequest {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Request); err != nil {
			return err
		}
	}
	// lpMessage {in} (1:{pointer=unique}*(1)[dim:0,string](char))
	{
		_ptr_lpMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharString(ctx, w, &o.Message); err != nil {
				return err
			}
			return nil
		})
		_s_lpMessage := func(ptr interface{}) { o.Message = *ptr.(*string) }
		if err := w.ReadPointer(&o.Message, _s_lpMessage, _ptr_lpMessage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpArg1Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpArg1Id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Arg1ID == nil {
				o.Arg1ID = &dtyp.GUID{}
			}
			if err := o.Arg1ID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpArg1Id := func(ptr interface{}) { o.Arg1ID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Arg1ID, _s_lpArg1Id, _ptr_lpArg1Id); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpArg2Id {in} (1:{pointer=unique, alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_lpArg2Id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Arg2ID == nil {
				o.Arg2ID = &dtyp.GUID{}
			}
			if err := o.Arg2ID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpArg2Id := func(ptr interface{}) { o.Arg2ID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.Arg2ID, _s_lpArg2Id, _ptr_lpArg2Id); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitNTMSOperatorRequestAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitNTMSOperatorRequestAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {out} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SubmitNTMSOperatorRequestAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {out} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
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

// SubmitNTMSOperatorRequestARequest structure represents the SubmitNtmsOperatorRequestA operation request
type SubmitNTMSOperatorRequestARequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwRequest: A value from the NtmsOpreqCommand (section 2.2.1.7) enumeration specifying
	// the type of operation requested. If dwRequest is set to NTMS_OPREQ_UNKNOWN, the server
	// MUST return ERROR_INVALID_PARAMETER and take no action.
	Request uint32 `idl:"name:dwRequest" json:"request"`
	// lpMessage: An optional null-terminated message string to send to the user.
	Message string `idl:"name:lpMessage;pointer:unique" json:"message"`
	// lpArg1Id: This parameter MUST be set according to the value of dwRequest.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|              VALUE OF               |                                                                                  |
	//	|              DWREQUEST              |                                     LPARG1ID                                     |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_CLEANER 0x00000002       | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_DEVICESERVICE 0x00000003 | MUST be set to a pointer to the identifier of the device that needs service.     |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MESSAGE 0x00000005       | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MOVEMEDIA 0x00000004     | MUST be set to a pointer to the identifier of the piece of physical media to     |
	//	|                                     | move.                                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_NEWMEDIA 0x00000001      | MUST be set to a pointer to the identifier of the media pool from which to       |
	//	|                                     | retrieve the new media.                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Arg1ID *dtyp.GUID `idl:"name:lpArg1Id;pointer:unique" json:"arg1_id"`
	// lpArg2Id: This parameter MUST be set according to the value of dwRequest.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|              VALUE OF               |                                                                                  |
	//	|              DWREQUEST              |                                     LPARG2ID                                     |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_CLEANER 0x00000002       | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_DEVICESERVICE 0x00000003 | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MESSAGE 0x00000005       | MUST be set to NULL.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_MOVEMEDIA 0x00000004     | MUST be set to a pointer to the identifier of the target library to which the    |
	//	|                                     | media MUST be moved.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQ_NEWMEDIA 0x00000001      | MUST be set to a pointer to the identifier of the media library to which the new |
	//	|                                     | media MUST be added.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Arg2ID *dtyp.GUID `idl:"name:lpArg2Id;pointer:unique" json:"arg2_id"`
}

func (o *SubmitNTMSOperatorRequestARequest) xxx_ToOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestAOperation) *xxx_SubmitNTMSOperatorRequestAOperation {
	if op == nil {
		op = &xxx_SubmitNTMSOperatorRequestAOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Request = o.Request
	op.Message = o.Message
	op.Arg1ID = o.Arg1ID
	op.Arg2ID = o.Arg2ID
	return op
}

func (o *SubmitNTMSOperatorRequestARequest) xxx_FromOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestAOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Request = op.Request
	o.Message = op.Message
	o.Arg1ID = op.Arg1ID
	o.Arg2ID = op.Arg2ID
}
func (o *SubmitNTMSOperatorRequestARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SubmitNTMSOperatorRequestARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitNTMSOperatorRequestAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SubmitNTMSOperatorRequestAResponse structure represents the SubmitNtmsOperatorRequestA operation response
type SubmitNTMSOperatorRequestAResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpRequestId: A pointer to the identifier of the created operator request.
	//
	//	+------------------------------------+--------------------------------------------------+
	//	|               RETURN               |                                                  |
	//	|             VALUE/CODE             |                   DESCRIPTION                    |
	//	|                                    |                                                  |
	//	+------------------------------------+--------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                         |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to one or more objects is denied.         |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | An invalid parameter was found.                  |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED     | Unable to connect to the service.                |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | Unable to find the source or destination object. |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE  | The database query or update failed.             |
	//	+------------------------------------+--------------------------------------------------+
	//
	// The SubmitNtmsOperatorRequestA method submits an operator request, and returns the
	// status of the request (Satisfied or Canceled), or times out (if the operator does
	// not act upon the request). Operator requests are used to request media, to request
	// that the specified medium be moved from one library to another, or to request the
	// server device service.
	RequestID *dtyp.GUID `idl:"name:lpRequestId" json:"request_id"`
	// Return: The SubmitNtmsOperatorRequestA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SubmitNTMSOperatorRequestAResponse) xxx_ToOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestAOperation) *xxx_SubmitNTMSOperatorRequestAOperation {
	if op == nil {
		op = &xxx_SubmitNTMSOperatorRequestAOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RequestID = o.RequestID
	op.Return = o.Return
	return op
}

func (o *SubmitNTMSOperatorRequestAResponse) xxx_FromOp(ctx context.Context, op *xxx_SubmitNTMSOperatorRequestAOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RequestID = op.RequestID
	o.Return = op.Return
}
func (o *SubmitNTMSOperatorRequestAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SubmitNTMSOperatorRequestAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitNTMSOperatorRequestAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitForNTMSOperatorRequestOperation structure represents the WaitForNtmsOperatorRequest operation
type xxx_WaitForNTMSOperatorRequestOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	Timeout   uint32         `idl:"name:dwTimeout" json:"timeout"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WaitForNTMSOperatorRequestOperation) OpNum() int { return 5 }

func (o *xxx_WaitForNTMSOperatorRequestOperation) OpName() string {
	return "/INtmsSession1/v0/WaitForNtmsOperatorRequest"
}

func (o *xxx_WaitForNTMSOperatorRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForNTMSOperatorRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForNTMSOperatorRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForNTMSOperatorRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForNTMSOperatorRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_WaitForNTMSOperatorRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// WaitForNTMSOperatorRequestRequest structure represents the WaitForNtmsOperatorRequest operation request
type WaitForNTMSOperatorRequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpRequestId: A pointer to the identifier of the request for which to wait.
	RequestID *dtyp.GUID `idl:"name:lpRequestId" json:"request_id"`
	// dwTimeout: The number of milliseconds to wait. To check for an operator request,
	// specify a time-out value of 0. If a value of 0xFFFFFFFF is specified, this method
	// does not time out.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | Access to one or more objects is denied.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | An invalid parameter was found.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704C7 ERROR_CANCELLED         | The operator request was canceled by an administrator.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800705B4 ERROR_TIMEOUT           | The time specified in the dwTimeout parameter elapsed before completion of the   |
	//	|                                    | operator request.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED     | Unable to connect to the service.                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | Unable to find the operator request object. Object requests are flushed from the |
	//	|                                    | database.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Timeout uint32 `idl:"name:dwTimeout" json:"timeout"`
}

func (o *WaitForNTMSOperatorRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_WaitForNTMSOperatorRequestOperation) *xxx_WaitForNTMSOperatorRequestOperation {
	if op == nil {
		op = &xxx_WaitForNTMSOperatorRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestID = o.RequestID
	op.Timeout = o.Timeout
	return op
}

func (o *WaitForNTMSOperatorRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitForNTMSOperatorRequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestID = op.RequestID
	o.Timeout = op.Timeout
}
func (o *WaitForNTMSOperatorRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WaitForNTMSOperatorRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForNTMSOperatorRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WaitForNTMSOperatorRequestResponse structure represents the WaitForNtmsOperatorRequest operation response
type WaitForNTMSOperatorRequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The WaitForNtmsOperatorRequest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitForNTMSOperatorRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_WaitForNTMSOperatorRequestOperation) *xxx_WaitForNTMSOperatorRequestOperation {
	if op == nil {
		op = &xxx_WaitForNTMSOperatorRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *WaitForNTMSOperatorRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitForNTMSOperatorRequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *WaitForNTMSOperatorRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WaitForNTMSOperatorRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForNTMSOperatorRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelNTMSOperatorRequestOperation structure represents the CancelNtmsOperatorRequest operation
type xxx_CancelNTMSOperatorRequestOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelNTMSOperatorRequestOperation) OpNum() int { return 6 }

func (o *xxx_CancelNTMSOperatorRequestOperation) OpName() string {
	return "/INtmsSession1/v0/CancelNtmsOperatorRequest"
}

func (o *xxx_CancelNTMSOperatorRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelNTMSOperatorRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CancelNTMSOperatorRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelNTMSOperatorRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelNTMSOperatorRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelNTMSOperatorRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CancelNTMSOperatorRequestRequest structure represents the CancelNtmsOperatorRequest operation request
type CancelNTMSOperatorRequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpRequestId: A pointer to the identifier of the request to cancel.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | The user who tried to execute this method does not have administrator            |
	//	|                                    | privileges. Only a server administrator can cancel operator requests.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is missing.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The operator request object identifier was not found. This error occurs if the   |
	//	|                                    | request is completed prior to cancellation of the operation, or when an invalid  |
	//	|                                    | request identifier is supplied.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE     | The request has already been completed or cancelled.                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	RequestID *dtyp.GUID `idl:"name:lpRequestId" json:"request_id"`
}

func (o *CancelNTMSOperatorRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelNTMSOperatorRequestOperation) *xxx_CancelNTMSOperatorRequestOperation {
	if op == nil {
		op = &xxx_CancelNTMSOperatorRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestID = o.RequestID
	return op
}

func (o *CancelNTMSOperatorRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelNTMSOperatorRequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestID = op.RequestID
}
func (o *CancelNTMSOperatorRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelNTMSOperatorRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelNTMSOperatorRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelNTMSOperatorRequestResponse structure represents the CancelNtmsOperatorRequest operation response
type CancelNTMSOperatorRequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CancelNtmsOperatorRequest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelNTMSOperatorRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelNTMSOperatorRequestOperation) *xxx_CancelNTMSOperatorRequestOperation {
	if op == nil {
		op = &xxx_CancelNTMSOperatorRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CancelNTMSOperatorRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelNTMSOperatorRequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelNTMSOperatorRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelNTMSOperatorRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelNTMSOperatorRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SatisfyNTMSOperatorRequestOperation structure represents the SatisfyNtmsOperatorRequest operation
type xxx_SatisfyNTMSOperatorRequestOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RequestID *dtyp.GUID     `idl:"name:lpRequestId" json:"request_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SatisfyNTMSOperatorRequestOperation) OpNum() int { return 7 }

func (o *xxx_SatisfyNTMSOperatorRequestOperation) OpName() string {
	return "/INtmsSession1/v0/SatisfyNtmsOperatorRequest"
}

func (o *xxx_SatisfyNTMSOperatorRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SatisfyNTMSOperatorRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID != nil {
			if err := o.RequestID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SatisfyNTMSOperatorRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpRequestId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.RequestID == nil {
			o.RequestID = &dtyp.GUID{}
		}
		if err := o.RequestID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SatisfyNTMSOperatorRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SatisfyNTMSOperatorRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SatisfyNTMSOperatorRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SatisfyNTMSOperatorRequestRequest structure represents the SatisfyNtmsOperatorRequest operation request
type SatisfyNTMSOperatorRequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpRequestId: A pointer to the identifier of the request to complete.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED     | The user who tried to execute this method does not have administrator            |
	//	|                                    | privileges. Only a server administrator can satisfy operator requests.           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is missing.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND  | The operator request object identifier was not found. This error occurs if the   |
	//	|                                    | request is completed before the operation has been satisfied, or when an invalid |
	//	|                                    | request identifier is supplied.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	RequestID *dtyp.GUID `idl:"name:lpRequestId" json:"request_id"`
}

func (o *SatisfyNTMSOperatorRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_SatisfyNTMSOperatorRequestOperation) *xxx_SatisfyNTMSOperatorRequestOperation {
	if op == nil {
		op = &xxx_SatisfyNTMSOperatorRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RequestID = o.RequestID
	return op
}

func (o *SatisfyNTMSOperatorRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_SatisfyNTMSOperatorRequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RequestID = op.RequestID
}
func (o *SatisfyNTMSOperatorRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SatisfyNTMSOperatorRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SatisfyNTMSOperatorRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SatisfyNTMSOperatorRequestResponse structure represents the SatisfyNtmsOperatorRequest operation response
type SatisfyNTMSOperatorRequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SatisfyNtmsOperatorRequest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SatisfyNTMSOperatorRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_SatisfyNTMSOperatorRequestOperation) *xxx_SatisfyNTMSOperatorRequestOperation {
	if op == nil {
		op = &xxx_SatisfyNTMSOperatorRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SatisfyNTMSOperatorRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_SatisfyNTMSOperatorRequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SatisfyNTMSOperatorRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SatisfyNTMSOperatorRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SatisfyNTMSOperatorRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportNTMSDatabaseOperation structure represents the ImportNtmsDatabase operation
type xxx_ImportNTMSDatabaseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportNTMSDatabaseOperation) OpNum() int { return 8 }

func (o *xxx_ImportNTMSDatabaseOperation) OpName() string {
	return "/INtmsSession1/v0/ImportNtmsDatabase"
}

func (o *xxx_ImportNTMSDatabaseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportNTMSDatabaseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ImportNTMSDatabaseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ImportNTMSDatabaseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportNTMSDatabaseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ImportNTMSDatabaseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ImportNTMSDatabaseRequest structure represents the ImportNtmsDatabase operation request
type ImportNTMSDatabaseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ImportNTMSDatabaseRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportNTMSDatabaseOperation) *xxx_ImportNTMSDatabaseOperation {
	if op == nil {
		op = &xxx_ImportNTMSDatabaseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ImportNTMSDatabaseRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportNTMSDatabaseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ImportNTMSDatabaseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportNTMSDatabaseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportNTMSDatabaseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportNTMSDatabaseResponse structure represents the ImportNtmsDatabase operation response
type ImportNTMSDatabaseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ImportNtmsDatabase return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportNTMSDatabaseResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportNTMSDatabaseOperation) *xxx_ImportNTMSDatabaseOperation {
	if op == nil {
		op = &xxx_ImportNTMSDatabaseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ImportNTMSDatabaseResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportNTMSDatabaseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ImportNTMSDatabaseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportNTMSDatabaseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportNTMSDatabaseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportNTMSDatabaseOperation structure represents the ExportNtmsDatabase operation
type xxx_ExportNTMSDatabaseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportNTMSDatabaseOperation) OpNum() int { return 9 }

func (o *xxx_ExportNTMSDatabaseOperation) OpName() string {
	return "/INtmsSession1/v0/ExportNtmsDatabase"
}

func (o *xxx_ExportNTMSDatabaseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportNTMSDatabaseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExportNTMSDatabaseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ExportNTMSDatabaseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportNTMSDatabaseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExportNTMSDatabaseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ExportNTMSDatabaseRequest structure represents the ExportNtmsDatabase operation request
type ExportNTMSDatabaseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ExportNTMSDatabaseRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportNTMSDatabaseOperation) *xxx_ExportNTMSDatabaseOperation {
	if op == nil {
		op = &xxx_ExportNTMSDatabaseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ExportNTMSDatabaseRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportNTMSDatabaseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ExportNTMSDatabaseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportNTMSDatabaseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportNTMSDatabaseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportNTMSDatabaseResponse structure represents the ExportNtmsDatabase operation response
type ExportNTMSDatabaseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExportNtmsDatabase return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportNTMSDatabaseResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportNTMSDatabaseOperation) *xxx_ExportNTMSDatabaseOperation {
	if op == nil {
		op = &xxx_ExportNTMSDatabaseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ExportNTMSDatabaseResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportNTMSDatabaseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ExportNTMSDatabaseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportNTMSDatabaseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportNTMSDatabaseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotificationOperation structure represents the AddNotification operation
type xxx_AddNotificationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Type   uint32         `idl:"name:dwType" json:"type"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotificationOperation) OpNum() int { return 11 }

func (o *xxx_AddNotificationOperation) OpName() string { return "/INtmsSession1/v0/AddNotification" }

func (o *xxx_AddNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddNotificationRequest structure represents the AddNotification operation request
type AddNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of object for which the client wants to receive change notifications.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	Type uint32 `idl:"name:dwType" json:"type"`
}

func (o *AddNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_AddNotificationOperation) *xxx_AddNotificationOperation {
	if op == nil {
		op = &xxx_AddNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	return op
}

func (o *AddNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
}
func (o *AddNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotificationResponse structure represents the AddNotification operation response
type AddNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_AddNotificationOperation) *xxx_AddNotificationOperation {
	if op == nil {
		op = &xxx_AddNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AddNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveNotificationOperation structure represents the RemoveNotification operation
type xxx_RemoveNotificationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Type   uint32         `idl:"name:dwType" json:"type"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveNotificationOperation) OpNum() int { return 12 }

func (o *xxx_RemoveNotificationOperation) OpName() string {
	return "/INtmsSession1/v0/RemoveNotification"
}

func (o *xxx_RemoveNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RemoveNotificationRequest structure represents the RemoveNotification operation request
type RemoveNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of object for which the client no longer wants to receive change notifications.
	//
	//	+--------------------------------+---------------------------------+
	//	|             RETURN             |                                 |
	//	|           VALUE/CODE           |           DESCRIPTION           |
	//	|                                |                                 |
	//	+--------------------------------+---------------------------------+
	//	+--------------------------------+---------------------------------+
	//	| 0x00000000 S_OK                | The call was successful.        |
	//	+--------------------------------+---------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED | Access to an object was denied. |
	//	+--------------------------------+---------------------------------+
	Type uint32 `idl:"name:dwType" json:"type"`
}

func (o *RemoveNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveNotificationOperation) *xxx_RemoveNotificationOperation {
	if op == nil {
		op = &xxx_RemoveNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	return op
}

func (o *RemoveNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
}
func (o *RemoveNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveNotificationResponse structure represents the RemoveNotification operation response
type RemoveNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemoveNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveNotificationOperation) *xxx_RemoveNotificationOperation {
	if op == nil {
		op = &xxx_RemoveNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RemoveNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoveNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DispatchNotificationOperation structure represents the DispatchNotification operation
type xxx_DispatchNotificationOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Type      uint32         `idl:"name:dwType" json:"type"`
	Operation uint32         `idl:"name:dwOperation" json:"operation"`
	ID        *dtyp.GUID     `idl:"name:lpIdentifier" json:"id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DispatchNotificationOperation) OpNum() int { return 13 }

func (o *xxx_DispatchNotificationOperation) OpName() string {
	return "/INtmsSession1/v0/DispatchNotification"
}

func (o *xxx_DispatchNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DispatchNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Operation); err != nil {
			return err
		}
	}
	// lpIdentifier {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ID != nil {
			if err := o.ID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DispatchNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Operation); err != nil {
			return err
		}
	}
	// lpIdentifier {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ID == nil {
			o.ID = &dtyp.GUID{}
		}
		if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DispatchNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DispatchNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DispatchNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DispatchNotificationRequest structure represents the DispatchNotification operation request
type DispatchNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwType: A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of object to which the notification refers.
	Type uint32 `idl:"name:dwType" json:"type"`
	// dwOperation: A value from the NtmsNotificationOperations (section 2.2.1.8) enumeration
	// specifying the type of operation to which the notification refers.
	Operation uint32 `idl:"name:dwOperation" json:"operation"`
	// lpIdentifier:  A pointer to the identifier of the event for which notification is
	// being sent.
	//
	//	+--------------------------------+---------------------------------+
	//	|             RETURN             |                                 |
	//	|           VALUE/CODE           |           DESCRIPTION           |
	//	|                                |                                 |
	//	+--------------------------------+---------------------------------+
	//	+--------------------------------+---------------------------------+
	//	| 0x00000000 S_OK                | The call was successful.        |
	//	+--------------------------------+---------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED | Access to an object was denied. |
	//	+--------------------------------+---------------------------------+
	ID *dtyp.GUID `idl:"name:lpIdentifier" json:"id"`
}

func (o *DispatchNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_DispatchNotificationOperation) *xxx_DispatchNotificationOperation {
	if op == nil {
		op = &xxx_DispatchNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	op.Operation = o.Operation
	op.ID = o.ID
	return op
}

func (o *DispatchNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_DispatchNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.Operation = op.Operation
	o.ID = op.ID
}
func (o *DispatchNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DispatchNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DispatchNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DispatchNotificationResponse structure represents the DispatchNotification operation response
type DispatchNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DispatchNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DispatchNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_DispatchNotificationOperation) *xxx_DispatchNotificationOperation {
	if op == nil {
		op = &xxx_DispatchNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DispatchNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_DispatchNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DispatchNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DispatchNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DispatchNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

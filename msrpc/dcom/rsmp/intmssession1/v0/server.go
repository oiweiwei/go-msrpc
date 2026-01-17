package intmssession1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = iunknown.GoPackage
)

// INtmsSession1 server interface.
type Session1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The OpenNtmsServerSessionW method sets up a session with the server. All input parameters
	// for this method are optional.
	OpenNTMSServerSessionW(context.Context, *OpenNTMSServerSessionWRequest) (*OpenNTMSServerSessionWResponse, error)

	// The OpenNtmsServerSessionA method sets up a session with the server. All input parameters
	// for this method are optional.
	OpenNTMSServerSessionA(context.Context, *OpenNTMSServerSessionARequest) (*OpenNTMSServerSessionAResponse, error)

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
	CloseNTMSSession(context.Context, *CloseNTMSSessionRequest) (*CloseNTMSSessionResponse, error)

	// The SubmitNtmsOperatorRequestW method submits an operator request, with strings encoded
	// using Unicode.
	SubmitNTMSOperatorRequestW(context.Context, *SubmitNTMSOperatorRequestWRequest) (*SubmitNTMSOperatorRequestWResponse, error)

	// The SubmitNtmsOperatorRequestA method submits an operator request, with strings encoded
	// using ASCII.
	SubmitNTMSOperatorRequestA(context.Context, *SubmitNTMSOperatorRequestARequest) (*SubmitNTMSOperatorRequestAResponse, error)

	// The WaitForNtmsOperatorRequest method waits for an operator request.
	WaitForNTMSOperatorRequest(context.Context, *WaitForNTMSOperatorRequestRequest) (*WaitForNTMSOperatorRequestResponse, error)

	// The CancelNtmsOperatorRequest method cancels an operator request.
	CancelNTMSOperatorRequest(context.Context, *CancelNTMSOperatorRequestRequest) (*CancelNTMSOperatorRequestResponse, error)

	// The SatisfyNtmsOperatorRequest method completes an operator request.
	SatisfyNTMSOperatorRequest(context.Context, *SatisfyNTMSOperatorRequestRequest) (*SatisfyNTMSOperatorRequestResponse, error)

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
	ImportNTMSDatabase(context.Context, *ImportNTMSDatabaseRequest) (*ImportNTMSDatabaseResponse, error)

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
	ExportNTMSDatabase(context.Context, *ExportNTMSDatabaseRequest) (*ExportNTMSDatabaseResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// The AddNotification method registers a client to receive change notifications for
	// a type of object.
	AddNotification(context.Context, *AddNotificationRequest) (*AddNotificationResponse, error)

	// The RemoveNotification method unregisters a client from receiving change notifications
	// for a type of object.
	RemoveNotification(context.Context, *RemoveNotificationRequest) (*RemoveNotificationResponse, error)

	// The DispatchNotification method sends a notification to all registered sinks.
	DispatchNotification(context.Context, *DispatchNotificationRequest) (*DispatchNotificationResponse, error)
}

func RegisterSession1Server(conn dcerpc.Conn, o Session1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSession1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Session1SyntaxV0_0))...)
}

func NewSession1ServerHandle(o Session1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Session1ServerHandle(ctx, o, opNum, r)
	}
}

func Session1ServerHandle(ctx context.Context, o Session1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // OpenNtmsServerSessionW
		op := &xxx_OpenNTMSServerSessionWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNTMSServerSessionWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNTMSServerSessionW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // OpenNtmsServerSessionA
		op := &xxx_OpenNTMSServerSessionAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNTMSServerSessionARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNTMSServerSessionA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // CloseNtmsSession
		op := &xxx_CloseNTMSSessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseNTMSSessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseNTMSSession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // SubmitNtmsOperatorRequestW
		op := &xxx_SubmitNTMSOperatorRequestWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SubmitNTMSOperatorRequestWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SubmitNTMSOperatorRequestW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SubmitNtmsOperatorRequestA
		op := &xxx_SubmitNTMSOperatorRequestAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SubmitNTMSOperatorRequestARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SubmitNTMSOperatorRequestA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // WaitForNtmsOperatorRequest
		op := &xxx_WaitForNTMSOperatorRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WaitForNTMSOperatorRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WaitForNTMSOperatorRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // CancelNtmsOperatorRequest
		op := &xxx_CancelNTMSOperatorRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelNTMSOperatorRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CancelNTMSOperatorRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // SatisfyNtmsOperatorRequest
		op := &xxx_SatisfyNTMSOperatorRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SatisfyNTMSOperatorRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SatisfyNTMSOperatorRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ImportNtmsDatabase
		op := &xxx_ImportNTMSDatabaseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportNTMSDatabaseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportNTMSDatabase(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ExportNtmsDatabase
		op := &xxx_ExportNTMSDatabaseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportNTMSDatabaseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportNTMSDatabase(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Opnum13NotUsedOnWire
		// Opnum13NotUsedOnWire
		return nil, nil
	case 11: // AddNotification
		op := &xxx_AddNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RemoveNotification
		op := &xxx_RemoveNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DispatchNotification
		op := &xxx_DispatchNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DispatchNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DispatchNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsSession1
type UnimplementedSession1Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedSession1Server) OpenNTMSServerSessionW(context.Context, *OpenNTMSServerSessionWRequest) (*OpenNTMSServerSessionWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) OpenNTMSServerSessionA(context.Context, *OpenNTMSServerSessionARequest) (*OpenNTMSServerSessionAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) CloseNTMSSession(context.Context, *CloseNTMSSessionRequest) (*CloseNTMSSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) SubmitNTMSOperatorRequestW(context.Context, *SubmitNTMSOperatorRequestWRequest) (*SubmitNTMSOperatorRequestWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) SubmitNTMSOperatorRequestA(context.Context, *SubmitNTMSOperatorRequestARequest) (*SubmitNTMSOperatorRequestAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) WaitForNTMSOperatorRequest(context.Context, *WaitForNTMSOperatorRequestRequest) (*WaitForNTMSOperatorRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) CancelNTMSOperatorRequest(context.Context, *CancelNTMSOperatorRequestRequest) (*CancelNTMSOperatorRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) SatisfyNTMSOperatorRequest(context.Context, *SatisfyNTMSOperatorRequestRequest) (*SatisfyNTMSOperatorRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) ImportNTMSDatabase(context.Context, *ImportNTMSDatabaseRequest) (*ImportNTMSDatabaseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) ExportNTMSDatabase(context.Context, *ExportNTMSDatabaseRequest) (*ExportNTMSDatabaseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) AddNotification(context.Context, *AddNotificationRequest) (*AddNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) RemoveNotification(context.Context, *RemoveNotificationRequest) (*RemoveNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSession1Server) DispatchNotification(context.Context, *DispatchNotificationRequest) (*DispatchNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Session1Server = (*UnimplementedSession1Server)(nil)

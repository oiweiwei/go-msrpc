package irasrv

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = idispatch.GoPackage
)

// IRASrv server interface.
type RemoteAssistanceServerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

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
	GetNoviceUserInfo(context.Context, *GetNoviceUserInfoRequest) (*GetNoviceUserInfoResponse, error)

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
	GetSessionInfo(context.Context, *GetSessionInfoRequest) (*GetSessionInfoResponse, error)
}

func RegisterRemoteAssistanceServerServer(conn dcerpc.Conn, o RemoteAssistanceServerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteAssistanceServerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteAssistanceServerSyntaxV0_0))...)
}

func NewRemoteAssistanceServerServerHandle(o RemoteAssistanceServerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteAssistanceServerServerHandle(ctx, o, opNum, r)
	}
}

func RemoteAssistanceServerServerHandle(ctx context.Context, o RemoteAssistanceServerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetNoviceUserInfo
		op := &xxx_GetNoviceUserInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNoviceUserInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNoviceUserInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetSessionInfo
		op := &xxx_GetSessionInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRASrv
type UnimplementedRemoteAssistanceServerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedRemoteAssistanceServerServer) GetNoviceUserInfo(context.Context, *GetNoviceUserInfoRequest) (*GetNoviceUserInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteAssistanceServerServer) GetSessionInfo(context.Context, *GetSessionInfoRequest) (*GetSessionInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteAssistanceServerServer = (*UnimplementedRemoteAssistanceServerServer)(nil)

package msgsvc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// msgsvc server interface.
type MsgsvcServer interface {

	// The NetrMessageNameAdd (Opnum 0) interface is used to configure the message server
	// to listen for messages sent to an additional NetBIOS name.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	|             RETURN             |                                                                         |
	//	|           VALUE/CODE           |                               DESCRIPTION                               |
	//	|                                |                                                                         |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The operation completed successfully.                                   |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                       |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME  | The file name, directory name, or volume label syntax is incorrect.     |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000859 NERR_NetworkError   | A general network error occurred.                                       |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x0000085C NERR_InternalError  | An internal error occurred.                                             |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x000008E4 NERR_AlreadyExists  | This message alias already exists locally.                              |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x000008E5 NERR_TooManyNames   | The maximum number of added message aliases has been exceeded.          |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x000008F9 NERR_DuplicateName  | The name specified is already in use as a message alias on the network. |
	//	+--------------------------------+-------------------------------------------------------------------------+
	MessageNameAdd(context.Context, *MessageNameAddRequest) (*MessageNameAddResponse, error)

	// The NetrMessageNameEnum (Opnum 1) interface is used to enumerate the NetBIOS names
	// for which the message server is currently listening for messages.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+------------------------------------+---------------------------------------+
	//	|               RETURN               |                                       |
	//	|             VALUE/CODE             |              DESCRIPTION              |
	//	|                                    |                                       |
	//	+------------------------------------+---------------------------------------+
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000000 NERR_Success            | The operation completed successfully. |
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                     |
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.           |
	//	+------------------------------------+---------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct. |
	//	+------------------------------------+---------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall        | The API return buffer is too small.   |
	//	+------------------------------------+---------------------------------------+
	MessageNameEnum(context.Context, *MessageNameEnumRequest) (*MessageNameEnumResponse, error)

	// The NetrMessageNameGetInfo (Opnum 2) interface is used to retrieve information from
	// the message server on a NetBIOS name for which the message server is currently listening
	// for messages.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN               |                                                                     |
	//	|             VALUE/CODE             |                             DESCRIPTION                             |
	//	|                                    |                                                                     |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The operation completed successfully.                               |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                   |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.            |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME      | The file name, directory name, or volume label syntax is incorrect. |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                               |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x000008ED NERR_NotLocalName       | The name is not on the local computer.                              |
	//	+------------------------------------+---------------------------------------------------------------------+
	MessageNameGetInfo(context.Context, *MessageNameGetInfoRequest) (*MessageNameGetInfoResponse, error)

	// The NetrMessageNameDel (Opnum 3) interface is used to configure the message server
	// to stop listening for messages for a particular NetBIOS name.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+---------------------------------+---------------------------------------------------------------------+
	//	|             RETURN              |                                                                     |
	//	|           VALUE/CODE            |                             DESCRIPTION                             |
	//	|                                 |                                                                     |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success         | The operation completed successfully.                               |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED  | Access is denied.                                                   |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME   | The file name, directory name, or volume label syntax is incorrect. |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008E6 NERR_DelComputerName | The computer name could not be deleted.                             |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008EB NERR_NameInUse       | The message alias is currently in use. Try again later.             |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008ED NERR_NotLocalName    | The name is not on the local computer.                              |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008FB NERR_IncompleteDel   | The message alias was not successfully deleted from all networks.   |
	//	+---------------------------------+---------------------------------------------------------------------+
	MessageNameDelete(context.Context, *MessageNameDeleteRequest) (*MessageNameDeleteResponse, error)
}

func RegisterMsgsvcServer(conn dcerpc.Conn, o MsgsvcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMsgsvcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MsgsvcSyntaxV1_0))...)
}

func NewMsgsvcServerHandle(o MsgsvcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MsgsvcServerHandle(ctx, o, opNum, r)
	}
}

func MsgsvcServerHandle(ctx context.Context, o MsgsvcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NetrMessageNameAdd
		op := &xxx_MessageNameAddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MessageNameAddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MessageNameAdd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // NetrMessageNameEnum
		op := &xxx_MessageNameEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MessageNameEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MessageNameEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // NetrMessageNameGetInfo
		op := &xxx_MessageNameGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MessageNameGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MessageNameGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // NetrMessageNameDel
		op := &xxx_MessageNameDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MessageNameDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MessageNameDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented msgsvc
type UnimplementedMsgsvcServer struct {
}

func (UnimplementedMsgsvcServer) MessageNameAdd(context.Context, *MessageNameAddRequest) (*MessageNameAddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMsgsvcServer) MessageNameEnum(context.Context, *MessageNameEnumRequest) (*MessageNameEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMsgsvcServer) MessageNameGetInfo(context.Context, *MessageNameGetInfoRequest) (*MessageNameGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMsgsvcServer) MessageNameDelete(context.Context, *MessageNameDeleteRequest) (*MessageNameDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ MsgsvcServer = (*UnimplementedMsgsvcServer)(nil)

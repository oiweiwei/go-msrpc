package winsi2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

// winsi2 server interface.
type Winsi2Server interface {

	// The R_WinsTombstoneDbRecs method tombstones records whose version numbers fall within
	// a range of version numbers and are owned by a server with a specified address.
	//
	// Return Values: A 32 bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation completed
	// successfully. Any other return value is a Win32 error code as specified in [MS-ERREF].
	// The following Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller does not have sufficient permissions. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	TombstoneDBRecords(context.Context, *TombstoneDBRecordsRequest) (*TombstoneDBRecordsResponse, error)

	// The R_ WinsCheckAccess method retrieves the level of access the client is granted.<12>
	//
	// Return Values: A 32-bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation completed
	// successfully. Any other return value is a Win32 error code as specified in [MS-ERREF].
	// The following Win32 error codes can be returned:
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CheckAccess(context.Context, *CheckAccessRequest) (*CheckAccessResponse, error)
}

func RegisterWinsi2Server(conn dcerpc.Conn, o Winsi2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWinsi2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Winsi2SyntaxV1_0))...)
}

func NewWinsi2ServerHandle(o Winsi2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Winsi2ServerHandle(ctx, o, opNum, r)
	}
}

func Winsi2ServerHandle(ctx context.Context, o Winsi2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_WinsTombstoneDbRecs
		in := &TombstoneDBRecordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.TombstoneDBRecords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // R_WinsCheckAccess
		in := &CheckAccessRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CheckAccess(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}

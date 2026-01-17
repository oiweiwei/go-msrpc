package browser

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

// browser server interface.
type BrowserServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// The I_BrowserrQueryOtherDomains method is received by the server in an RPC_REQUEST
	// packet. The client SHOULD NOT send this RPC request to a server that is not a primary
	// domain controller (PDC) acting as the domain master browser server.
	//
	// If this server is not a primary domain controller it MAY fail the request.<3>
	//
	// If the server is a primary domain controller, the server MUST update OtherDomains
	// as specified in [MS-WKST] section 3.2.6.1, WkstaQueryOtherDomains Event. The server
	// MUST construct a SERVER_ENUM structure as specified in 2.2.3.2, containing a SERVER_INFO_100
	// structure as specified in [MS-DTYP] section 2.3.11 for each name in OtherDomains,
	// and return this to the caller.
	//
	// Return Values: The method returns NERR_Success on success; otherwise, it returns
	// a nonzero error code, as specified in either Win32 Error Codes. The most common error
	// codes are listed in the following table.<5>
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The operation completed successfully.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | This value MUST be returned when the server could not allocate enough memory to  |
	//	|                                    | complete this operation.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | This value MUST be returned when a parameter is incorrect. For example, this     |
	//	|                                    | value is returned when the InfoStruct parameter is NULL or the Level100 member   |
	//	|                                    | in the structure pointed to by the InfoStruct parameter is NULL.                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | This value MUST be returned when the Level member is not 100.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The error ERROR_MORE_DATA indicates that not all available entries were          |
	//	|                                    | returned. Some more entries exist which were not returned in the response.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	QueryOtherDomains(context.Context, *QueryOtherDomainsRequest) (*QueryOtherDomainsResponse, error)

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire
}

func RegisterBrowserServer(conn dcerpc.Conn, o BrowserServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewBrowserServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(BrowserSyntaxV0_0))...)
}

func NewBrowserServerHandle(o BrowserServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return BrowserServerHandle(ctx, o, opNum, r)
	}
}

func BrowserServerHandle(ctx context.Context, o BrowserServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // I_BrowserrQueryOtherDomains
		op := &xxx_QueryOtherDomainsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryOtherDomainsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryOtherDomains(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented browser
type UnimplementedBrowserServer struct {
}

func (UnimplementedBrowserServer) QueryOtherDomains(context.Context, *QueryOtherDomainsRequest) (*QueryOtherDomainsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ BrowserServer = (*UnimplementedBrowserServer)(nil)

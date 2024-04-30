package iwamadmin2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iwamadmin "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/iwamadmin/v0"
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
	_ = iwamadmin.GoPackage
)

// IWamAdmin2 server interface.
type WAMAdmin2Server interface {

	// IWamAdmin base class.
	iwamadmin.WAMAdminServer

	// The AppCreate2 method creates a new application at the specified metabase path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// The opnum field value for this method is 9.
	AppCreate2(context.Context, *AppCreate2Request) (*AppCreate2Response, error)
}

func RegisterWAMAdmin2Server(conn dcerpc.Conn, o WAMAdmin2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWAMAdmin2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WAMAdmin2SyntaxV0_0))...)
}

func NewWAMAdmin2ServerHandle(o WAMAdmin2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WAMAdmin2ServerHandle(ctx, o, opNum, r)
	}
}

func WAMAdmin2ServerHandle(ctx context.Context, o WAMAdmin2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 9 {
		// IWamAdmin base method.
		return iwamadmin.WAMAdminServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 9: // AppCreate2
		in := &AppCreate2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AppCreate2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}

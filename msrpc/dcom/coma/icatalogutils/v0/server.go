package icatalogutils

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = iunknown.GoPackage
)

// ICatalogUtils server interface.
type CatalogUtilsServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to verify that a user account and password are
	// valid.
	//
	// Return Values: This method MUST return a value in the table below on success, or
	// a failure result, as specified in [MS-ERREF] section 2.1, on failure. All failure
	// results MUST be treated identically.
	//
	//	+--------------------+-------------------------------------------------------------+
	//	|       RETURN       |                                                             |
	//	|     VALUE/CODE     |                         DESCRIPTION                         |
	//	|                    |                                                             |
	//	+--------------------+-------------------------------------------------------------+
	//	+--------------------+-------------------------------------------------------------+
	//	| 0x00000000 S_OK    | The user account and password are valid.                    |
	//	+--------------------+-------------------------------------------------------------+
	//	| 0x00000001 S_FALSE | The user account was not found or the password was invalid. |
	//	+--------------------+-------------------------------------------------------------+
	ValidateUser(context.Context, *ValidateUserRequest) (*ValidateUserResponse, error)

	// This method is called by a COMA client to synchronize with the server after performing
	// a write operation.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	WaitForEndWrites(context.Context, *WaitForEndWritesRequest) (*WaitForEndWritesResponse, error)

	// This method is called by a client to get information about the event classes associated
	// with an IID that are configured in the Global Partition.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	GetEventClassesForIID(context.Context, *GetEventClassesForIIDRequest) (*GetEventClassesForIIDResponse, error)
}

func RegisterCatalogUtilsServer(conn dcerpc.Conn, o CatalogUtilsServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCatalogUtilsServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CatalogUtilsSyntaxV0_0))...)
}

func NewCatalogUtilsServerHandle(o CatalogUtilsServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CatalogUtilsServerHandle(ctx, o, opNum, r)
	}
}

func CatalogUtilsServerHandle(ctx context.Context, o CatalogUtilsServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ValidateUser
		in := &ValidateUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ValidateUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // WaitForEndWrites
		in := &WaitForEndWritesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WaitForEndWrites(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetEventClassesForIID
		in := &GetEventClassesForIIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassesForIID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}

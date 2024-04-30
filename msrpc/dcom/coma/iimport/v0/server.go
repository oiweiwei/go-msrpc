package iimport

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

// IImport server interface.
type ImportServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to import one or more conglomerations from an installer
	// package file. Importing a conglomeration from an installer package file conceptually
	// consists of installation of modules, including the registration of the components
	// in those modules, and creating a conglomeration and component configurations equivalent
	// to the conglomeration and the component configurations that were exported to create
	// the installer package file. As a side effect, this method returns implementation-specific
	// detailed results of the registration operation for informational purposes.
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
	ImportFromFile(context.Context, *ImportFromFileRequest) (*ImportFromFileResponse, error)

	// This method is called by a client to retrieve information about an installer package
	// file.
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success, and a failure
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
	QueryFile(context.Context, *QueryFileRequest) (*QueryFileResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire
}

func RegisterImportServer(conn dcerpc.Conn, o ImportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImportSyntaxV0_0))...)
}

func NewImportServerHandle(o ImportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImportServerHandle(ctx, o, opNum, r)
	}
}

func ImportServerHandle(ctx context.Context, o ImportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ImportFromFile
		in := &ImportFromFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ImportFromFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // QueryFile
		in := &QueryFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

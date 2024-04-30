package iexport

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

// IExport server interface.
type ExportServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to export a conglomeration to an installer package
	// file.
	//
	// pConglomerationIdentifier: The conglomeration identifier of a conglomeration on the
	// server.
	//
	// pwszInstallerPackage: A path in UNC that is to be used as the location for the server
	// to create an installer package file.
	//
	// pwszReserved:  MUST be an empty (zero-length) string.
	//
	// dwFlags:  MUST be a combination of zero or more of the following flags.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|             FLAG             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_OVERWRITE 0x00000001 | The server SHOULD mark the installer package file with a directive that existing |
	//	|                              | files be overwritten on import (section 3.1.4.12.1).                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_WITHUSERS 0x00000010 | The server SHOULD include user account information in the installer package      |
	//	|                              | file.                                                                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_PROXY 0x00000020     | The server SHOULD mark the exported conglomeration as a proxy conglomeration by  |
	//	|                              | setting the IsProxyApp property to TRUE (0x00000001).                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| fEXPORT_CATVER300 0x00000080 | The server SHOULD only include configuration that is defined in catalog version  |
	//	|                              | 3.00.                                                                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
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
	ExportConglomeration(context.Context, *ExportConglomerationRequest) (*ExportConglomerationResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire
}

func RegisterExportServer(conn dcerpc.Conn, o ExportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewExportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ExportSyntaxV0_0))...)
}

func NewExportServerHandle(o ExportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ExportServerHandle(ctx, o, opNum, r)
	}
}

func ExportServerHandle(ctx context.Context, o ExportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ExportConglomeration
		in := &ExportConglomerationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExportConglomeration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

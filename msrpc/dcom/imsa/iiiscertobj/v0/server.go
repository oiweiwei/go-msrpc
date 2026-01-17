package iiiscertobj

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IIISCertObj server interface.
type IISCertObjectServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Opnum7NotUsedOnWire operation.
	// SetOpnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// SetOpnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// SetOpnum9NotUsedOnWire

	// The InstanceName method sets the web server instance to be used by subsequent method
	// calls.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [[MS-ERREF].
	//
	//	+----------------------------------+------------------------------------+
	//	|              RETURN              |                                    |
	//	|            VALUE/CODE            |            DESCRIPTION             |
	//	|                                  |                                    |
	//	+----------------------------------+------------------------------------+
	//	+----------------------------------+------------------------------------+
	//	| 0x00000000 S_OK                  | The call was successful.           |
	//	+----------------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One or more arguments are invalid. |
	//	+----------------------------------+------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG | The string is too long.            |
	//	+----------------------------------+------------------------------------+
	//
	// The opnum field value for this method is 10.
	SetInstanceName(context.Context, *SetInstanceNameRequest) (*SetInstanceNameResponse, error)

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// The IsInstalledRemote method determines if a certificate is associated with the specified
	// InstanceName.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.           |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// The opnum field value for this method is 12.
	IsInstalledRemote(context.Context, *IsInstalledRemoteRequest) (*IsInstalledRemoteResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// The IsExportableRemote method determines whether the server certificate associated
	// with InstanceName can be exported.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.           |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// The opnum field value for this method is 14.
	IsExportableRemote(context.Context, *IsExportableRemoteRequest) (*IsExportableRemoteResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// The GetCertInfoRemote method retrieves properties from a certificate associated with
	// the specified InstanceName.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------+------------------------------------------------+
	//	|         RETURN          |                                                |
	//	|       VALUE/CODE        |                  DESCRIPTION                   |
	//	|                         |                                                |
	//	+-------------------------+------------------------------------------------+
	//	+-------------------------+------------------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.                       |
	//	+-------------------------+------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid.             |
	//	+-------------------------+------------------------------------------------+
	//	| 0x00000001 S_FALSE      | The call was successful. No data was returned. |
	//	+-------------------------+------------------------------------------------+
	//
	// The opnum field value for this method is 16.
	GetCertInfoRemote(context.Context, *GetCertInfoRemoteRequest) (*GetCertInfoRemoteResponse, error)

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Opnum18NotUsedOnWire operation.
	// Opnum18NotUsedOnWire

	// Opnum19NotUsedOnWire operation.
	// Opnum19NotUsedOnWire

	// Opnum20NotUsedOnWire operation.
	// Opnum20NotUsedOnWire

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// The ImportFromBlob method imports a previously exported certificate blob on the target
	// machine.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------------+----------------------------------------+
	//	|              RETURN              |                                        |
	//	|            VALUE/CODE            |              DESCRIPTION               |
	//	|                                  |                                        |
	//	+----------------------------------+----------------------------------------+
	//	+----------------------------------+----------------------------------------+
	//	| 0x00000000 S_OK                  | The call was successful.               |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One or more arguments are invalid.     |
	//	+----------------------------------+----------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG | The string is too long.                |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80092005 CRYPT_E_EXISTS        | The object or property already exists. |
	//	+----------------------------------+----------------------------------------+
	//
	// The opnum field value for this method is 22.
	ImportFromBlob(context.Context, *ImportFromBlobRequest) (*ImportFromBlobResponse, error)

	// The ImportFromBlobGetHash method imports a previously exported certificate blob on
	// the target machine. In addition to data returned by method ImportFromBlob, this method
	// returns certificate hash and certificate hash buffer size in client-provided parameters
	// pcbCertHashSize and pCertHash. The server MUST allocate memory for the hash buffer
	// and assign this memory block to pCertHash. Size of required buffer is assigned to
	// pcbCertHashSize. If client will pass pCertHash equal to NULL, hash data will not
	// be returned.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------------+----------------------------------------+
	//	|              RETURN              |                                        |
	//	|            VALUE/CODE            |              DESCRIPTION               |
	//	|                                  |                                        |
	//	+----------------------------------+----------------------------------------+
	//	+----------------------------------+----------------------------------------+
	//	| 0x00000000 S_OK                  | The call was successful.               |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One or more arguments are invalid.     |
	//	+----------------------------------+----------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG | The string is too long.                |
	//	+----------------------------------+----------------------------------------+
	//	| 0x80092005 CRYPT_E_EXISTS        | The object or property already exists. |
	//	+----------------------------------+----------------------------------------+
	//
	// The opnum field value for this method is 23.
	ImportFromBlobGetHash(context.Context, *ImportFromBlobGetHashRequest) (*ImportFromBlobGetHashResponse, error)

	// Opnum24NotUsedOnWire operation.
	// Opnum24NotUsedOnWire

	// The ExportToBlob method exports the certificate referenced at InstanceName to a memory
	// buffer.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+------------------------------------+-------------------------------------------------------+
	//	|               RETURN               |                                                       |
	//	|             VALUE/CODE             |                      DESCRIPTION                      |
	//	|                                    |                                                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                              |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG            | One or more arguments are invalid.                    |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x000006cf RPC_S_STRING_TOO_LONG   | The string is too long.                               |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x800CC801 MD_ERROR_DATA_NOT_FOUND | The specified metadata was not found.                 |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80092004 CRYPT_E_NOT_FOUND       | Cannot find object or property.                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80090349 SEC_E_CERT_WRONG_USAGE  | The certificate is not valid for the requested usage. |
	//	+------------------------------------+-------------------------------------------------------+
	//
	// The opnum field value for this method is 25.
	ExportToBlob(context.Context, *ExportToBlobRequest) (*ExportToBlobResponse, error)
}

func RegisterIISCertObjectServer(conn dcerpc.Conn, o IISCertObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIISCertObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IISCertObjectSyntaxV0_0))...)
}

func NewIISCertObjectServerHandle(o IISCertObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IISCertObjectServerHandle(ctx, o, opNum, r)
	}
}

func IISCertObjectServerHandle(ctx context.Context, o IISCertObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // InstanceName
		op := &xxx_SetInstanceNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInstanceNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInstanceName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	case 12: // IsInstalledRemote
		op := &xxx_IsInstalledRemoteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsInstalledRemoteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsInstalledRemote(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Opnum13NotUsedOnWire
		// Opnum13NotUsedOnWire
		return nil, nil
	case 14: // IsExportableRemote
		op := &xxx_IsExportableRemoteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsExportableRemoteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsExportableRemote(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // GetCertInfoRemote
		op := &xxx_GetCertInfoRemoteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCertInfoRemoteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCertInfoRemote(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 18: // Opnum18NotUsedOnWire
		// Opnum18NotUsedOnWire
		return nil, nil
	case 19: // Opnum19NotUsedOnWire
		// Opnum19NotUsedOnWire
		return nil, nil
	case 20: // Opnum20NotUsedOnWire
		// Opnum20NotUsedOnWire
		return nil, nil
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // ImportFromBlob
		op := &xxx_ImportFromBlobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportFromBlobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportFromBlob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // ImportFromBlobGetHash
		op := &xxx_ImportFromBlobGetHashOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportFromBlobGetHashRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportFromBlobGetHash(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // Opnum24NotUsedOnWire
		// Opnum24NotUsedOnWire
		return nil, nil
	case 25: // ExportToBlob
		op := &xxx_ExportToBlobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExportToBlobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExportToBlob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IIISCertObj
type UnimplementedIISCertObjectServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedIISCertObjectServer) SetInstanceName(context.Context, *SetInstanceNameRequest) (*SetInstanceNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISCertObjectServer) IsInstalledRemote(context.Context, *IsInstalledRemoteRequest) (*IsInstalledRemoteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISCertObjectServer) IsExportableRemote(context.Context, *IsExportableRemoteRequest) (*IsExportableRemoteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISCertObjectServer) GetCertInfoRemote(context.Context, *GetCertInfoRemoteRequest) (*GetCertInfoRemoteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISCertObjectServer) ImportFromBlob(context.Context, *ImportFromBlobRequest) (*ImportFromBlobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISCertObjectServer) ImportFromBlobGetHash(context.Context, *ImportFromBlobGetHashRequest) (*ImportFromBlobGetHashResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISCertObjectServer) ExportToBlob(context.Context, *ExportToBlobRequest) (*ExportToBlobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IISCertObjectServer = (*UnimplementedIISCertObjectServer)(nil)

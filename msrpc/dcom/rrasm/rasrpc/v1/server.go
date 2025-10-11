package rasrpc

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

// rasrpc server interface.
type RasrpcServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// The RasRpcDeleteEntry method deletes a specific RRAS Entry from an RRAS Phonebook
	// path.
	//
	// Return Values: Specifies the return status as explained in section 2.2.1.2.218 for
	// retcode field.
	//
	// The return value can be one of the following error codes. All other error values
	// MUST be treated the same by the RRASM client.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| ERROR_SUCCESS 0x00000000 | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// Exceptions Thrown: This method throws an exception with the exception code RPC_S_ACCESS_DENIED
	// (0x00000005) if the client is not an administrator on the RRASM server, with access
	// permission to perform the operation.<334>
	//
	// The opnum field value for this method is 5.
	//
	// * Validate as specified in section 3.3.4 ( 83a083ac-cd43-44bf-a301-7c88e64a32fe )
	// whether this method was called by a client which is an administrator of the RRASM
	// server.
	//
	// * If lpszEntry is NULL, return an error other than one of the errors specified in
	// the preceding table.
	//
	// * If lpszEntry is not present in *PhonebookEntryNameList* , return an error other
	// than one of the errors specified in the preceding table.
	//
	// * Call the abstract interface *Invoke RASRPC* method, specifying the operation and
	// the parameters necessary to enable the RRASM server to perform the required management
	// task.
	//
	// * Return any error result that the RRASM server returns as a part of the processing.
	// Otherwise return ERROR_SUCCESS (0x00000000).
	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// The RasRpcGetUserPreferences method retrieves the configuration information. The
	// configuration information consists of the callback information associated with the
	// various ports, and the number of the last successful callback done by the RRAS. This
	// configuration information is set by RasRpcSetUserPreferences (section 3.3.4.3).
	//
	// Return Values: Specifies the return status as explained in section 2.2.1.2.218 for
	// the retcode field.
	//
	// The return value can be one of the following error codes. All other error values
	// MUST be treated the same by the RRASM client.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| ERROR_SUCCESS 0x00000000 | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// Exceptions Thrown: This method throws an exception with the exception code RPC_S_ACCESS_DENIED
	// (0x00000005) if the client is not an administrator on the RRASM server, with access
	// permission to perform the operation.<335>
	//
	// The Opnum field value for this method is 9.
	//
	// * Validate as specified in section 3.3.4 ( 83a083ac-cd43-44bf-a301-7c88e64a32fe )
	// whether this method was called by a client that is an administrator of the RRASM
	// server. <336> ( 3bb906f0-b077-47ab-ad11-d8d807afde26#Appendix_A_336 )
	//
	// * Call the abstract interface *Invoke RASRPC* method, specifying the operation and
	// the parameters necessary to enable RRAS server to perform the required management
	// task.
	//
	// * Populate the pUser structure with the configuration information returned by the
	// RRAS server and returning ERROR_SUCCESS (0x00000000).
	GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*GetUserPreferencesResponse, error)

	// The RasRpcSetUserPreferences method sets the configuration information. The configuration
	// information consists of the callback information associated with the various ports,
	// and the number of the last successful callback done by the RRAS.
	//
	// Return Values: Specifies the return status as explained in section 2.2.1.2.218 for
	// the retcode field.
	//
	// The return value can be one of the error codes that follow. All other error values
	// MUST be treated the same by the RRASM client.
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
	// Exceptions Thrown: This method throws an exception with the exception code RPC_S_ACCESS_DENIED
	// (0x00000005) if the client is not an administrator on the RRASM server, with access
	// permission to perform the operation.<337>
	//
	// The opnum field value for this method is 10.
	//
	// * Validate as specified in section 3.3.4 ( 83a083ac-cd43-44bf-a301-7c88e64a32fe )
	// whether this method was called by a client that is an administrator of the RRASM
	// server.
	//
	// * Call the abstract interface *Invoke RASRPC* method, specifying the operation and
	// the parameters necessary to enable RRAS server to perform the required management
	// task.
	//
	// * Provide the configuration information as specified by the pUser structure to the
	// RRAS server for further processing and returning ERROR_SUCCESS (0x00000000).
	SetUserPreferences(context.Context, *SetUserPreferencesRequest) (*SetUserPreferencesResponse, error)

	// The RasRpcGetSystemDirectory method retrieves the path of the system directory.
	//
	// Return Values: Specifies the return status as explained in section 2.2.1.2.218 for
	// retcode field.
	//
	// The return value can be one of the following error codes. All other error values
	// MUST be treated the same by the RRASM client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|      RETURN      |                                                                                  |
	//	|      VALUE       |                                   DESCRIPTION                                    |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000       | The actual processing to retrieve the system directory on the remote server has  |
	//	|                  | failed.                                                                          |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| Any other values | Indicate the length of the string in Unicode characters copied to the buffer.    |
	//	+------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: This method throws an exception with the exception code RPC_S_ACCESS_DENIED
	// (0x00000005) if the client is not an administrator on the RRASM server, with access
	// permission to perform the operation.<338>
	//
	// The Opnum field value for this method is 11.
	//
	// * Validate as specified in section 3.3.4 ( 83a083ac-cd43-44bf-a301-7c88e64a32fe )
	// whether this method was called by a client that is an administrator of the RRASM
	// server. <339> ( 3bb906f0-b077-47ab-ad11-d8d807afde26#Appendix_A_339 )
	//
	// * If uSize is less than 260, return an error other than one of the errors specified
	// in the preceding table.
	//
	// * Call the abstract interface *Invoke RASRPC* method specifying the operation and
	// the parameters to enable RRAS server to perform the required management task.
	//
	// * If all validations are successful, return the processing information result for
	// the RRAS server and populate the lpBuffer with the system directory path returned
	// by the RRAS server. Return the length of the string in Unicode characters populated
	// to the lpBuffer.
	GetSystemDirectory(context.Context, *GetSystemDirectoryRequest) (*GetSystemDirectoryResponse, error)

	// The RasRpcSubmitRequest method retrieves or sets the configuration data on RRAS server.
	//
	// Return Values: Specifies the return status as explained in section 2.2.1.2.218 for
	// retcode field.
	//
	// Exceptions Thrown: This method throws an exception with the exception code RPC_S_ACCESS_DENIED
	// (0x00000005) if the client is not an administrator on the RRASM server, with access
	// permission to perform the operation.<340>
	//
	// Validations which SHOULD be done by the RRASM for all ReqTypes are:
	//
	// * Return ERROR_SUCCESS (0x00000000) if one of the following conditions are met without
	// any further processing of the call:
	//
	// * *dwcbBufSize* is less than the sum of size of *RequestBuffer* and 5000, i.e. if
	// the condition ( *dwcbBufSize* < size of *RequestBuffer* + 5000) is TRUE.
	//
	// * *pReqBuffer* is NULL
	//
	// * *pReqBuffer.RB_ReqType* is less than zero (0) or greater than maximum ReqTypes
	// <341> ( 3bb906f0-b077-47ab-ad11-d8d807afde26#Appendix_A_341 )
	//
	// * Validate as specified in section 3.3.4 ( 83a083ac-cd43-44bf-a301-7c88e64a32fe )
	// whether this method was called by a client that is an administrator ( fc2dfae9-0d04-4e1d-97c9-c51c2dc06c3b#gt_eb335f5b-619b-46ed-9138-06e9910108c3
	// ) of the RRASM server ( fc2dfae9-0d04-4e1d-97c9-c51c2dc06c3b#gt_f91f3fc7-d1a1-4a07-8c92-d19eb0c9acb0
	// ).
	//
	// Specific RRASM behavior for each ReqTypes value follows.
	SubmitRequest(context.Context, *SubmitRequestRequest) (*SubmitRequestResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// The RasRpcGetInstalledProtocolsEx method retrieves the protocol information on the
	// RRAS server. The list of protocols is defined in the following return value section.
	//
	// Return Values: Specifies the return status as explained in section 2.2.1.2.218 for
	// the retcode field.
	//
	// The return value can be one of the error codes that follow. All other error values
	// MUST be treated the same by the RRASM client.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN          |                                                                                  |
	//	|          VALUE           |                                   DESCRIPTION                                    |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000               | There is no protocol installed on the RRAS server or there is some error when    |
	//	|                          | RRAS server retrieves the information.                                           |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 to 0x0000000F | Specifies the protocols enabled on the RRAS server. This value SHOULD be a       |
	//	|                          | combination of one or more of the following flags: NP_Nbf (0x00000001): NetBEUI  |
	//	|                          | protocol is enabled.<344> NP_Ipx (0x00000002): IPX protocol is enabled.<345>     |
	//	|                          | NP_Ip (0x00000004): TCP/IPv4 protocol is enabled. NP_Ipv6 (0x00000008): TCP/IPv6 |
	//	|                          | protocol is enabled.<346>                                                        |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: This method throws an exception with the exception code RPC_S_ACCESS_DENIED
	// (0x00000005) if the client is not an administrator on the RRASM server, with access
	// permission to perform the operation.<347>
	//
	// The opnum field value for this method is 14.
	//
	// * Validate as specified in section 3.3.4 ( 83a083ac-cd43-44bf-a301-7c88e64a32fe )
	// whether this method was called by a client that is an administrator of the RRASM
	// server.
	//
	// * Call the abstract interface *Invoke RASRPC* method specifying the operation and
	// the parameters to enable RRAS server to perform the required management task.
	//
	// * If all validation is successful, return the installed protocol information as provided
	// by the RRAS server.
	GetInstalledProtocolsEx(context.Context, *GetInstalledProtocolsExRequest) (*GetInstalledProtocolsExResponse, error)

	// The RasRpcGetVersion method retrieves the Rasrpc server interface version.
	//
	// Return Values: Specifies the return status as explained in section 2.2.1.2.218 for
	// the retcode field
	//
	// The return value can be one of the error codes that follow. All other error values
	// MUST be treated the same by the RRASM client.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| ERROR_SUCCESS 0x00000000 | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// Exceptions Thrown: This method throws an exception with the exception code RPC_S_ACCESS_DENIED
	// (0x00000005) if the client is not an administrator on the RRASM server, with access
	// permission to perform the operation.<348>
	//
	// The opnum field value for this method is 15.
	//
	// * Validate as specified in section 3.3.4 ( 83a083ac-cd43-44bf-a301-7c88e64a32fe )
	// whether this method was called by a client that is an administrator of the RRASM
	// server. <349> ( 3bb906f0-b077-47ab-ad11-d8d807afde26#Appendix_A_349 )
	//
	// * Call the abstract interface *Invoke RASRPC method* specifying the operation and
	// the parameters to enable the RRAS server to perform the required management task.
	//
	// * Set the value pointed by pdwVersion to the version of RRAS server. <350> ( 3bb906f0-b077-47ab-ad11-d8d807afde26#Appendix_A_350
	// )
	//
	// * If there is no error, the server MUST return ERROR_SUCCESS (0x00000000).
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire
}

func RegisterRasrpcServer(conn dcerpc.Conn, o RasrpcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRasrpcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RasrpcSyntaxV1_0))...)
}

func NewRasrpcServerHandle(o RasrpcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RasrpcServerHandle(ctx, o, opNum, r)
	}
}

func RasrpcServerHandle(ctx context.Context, o RasrpcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // Opnum2NotUsedOnWire
		// Opnum2NotUsedOnWire
		return nil, nil
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // RasRpcDeleteEntry
		op := &xxx_DeleteEntryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteEntryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteEntry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // RasRpcGetUserPreferences
		op := &xxx_GetUserPreferencesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserPreferencesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserPreferences(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RasRpcSetUserPreferences
		op := &xxx_SetUserPreferencesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetUserPreferencesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetUserPreferences(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RasRpcGetSystemDirectory
		op := &xxx_GetSystemDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSystemDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSystemDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RasRpcSubmitRequest
		op := &xxx_SubmitRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SubmitRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SubmitRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Opnum13NotUsedOnWire
		// Opnum13NotUsedOnWire
		return nil, nil
	case 14: // RasRpcGetInstalledProtocolsEx
		op := &xxx_GetInstalledProtocolsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInstalledProtocolsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInstalledProtocolsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RasRpcGetVersion
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented rasrpc
type UnimplementedRasrpcServer struct {
}

func (UnimplementedRasrpcServer) DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*GetUserPreferencesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) SetUserPreferences(context.Context, *SetUserPreferencesRequest) (*SetUserPreferencesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetSystemDirectory(context.Context, *GetSystemDirectoryRequest) (*GetSystemDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) SubmitRequest(context.Context, *SubmitRequestRequest) (*SubmitRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetInstalledProtocolsEx(context.Context, *GetInstalledProtocolsExRequest) (*GetInstalledProtocolsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRasrpcServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RasrpcServer = (*UnimplementedRasrpcServer)(nil)

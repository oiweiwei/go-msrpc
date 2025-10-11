package rasrpc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	rrasm "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm"
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
	_ = rrasm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// Syntax UUID
	RasrpcSyntaxUUID = &uuid.UUID{TimeLow: 0x20610036, TimeMid: 0xfa22, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x23, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x11, 0xe5, 0xdf}}
	// Syntax ID
	RasrpcSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: RasrpcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// rasrpc interface.
type RasrpcClient interface {

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
	DeleteEntry(context.Context, *DeleteEntryRequest, ...dcerpc.CallOption) (*DeleteEntryResponse, error)

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
	GetUserPreferences(context.Context, *GetUserPreferencesRequest, ...dcerpc.CallOption) (*GetUserPreferencesResponse, error)

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
	SetUserPreferences(context.Context, *SetUserPreferencesRequest, ...dcerpc.CallOption) (*SetUserPreferencesResponse, error)

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
	GetSystemDirectory(context.Context, *GetSystemDirectoryRequest, ...dcerpc.CallOption) (*GetSystemDirectoryResponse, error)

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
	SubmitRequest(context.Context, *SubmitRequestRequest, ...dcerpc.CallOption) (*SubmitRequestResponse, error)

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
	GetInstalledProtocolsEx(context.Context, *GetInstalledProtocolsExRequest, ...dcerpc.CallOption) (*GetInstalledProtocolsExResponse, error)

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
	GetVersion(context.Context, *GetVersionRequest, ...dcerpc.CallOption) (*GetVersionResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultRasrpcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRasrpcClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...dcerpc.CallOption) (*DeleteEntryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteEntryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) GetUserPreferences(ctx context.Context, in *GetUserPreferencesRequest, opts ...dcerpc.CallOption) (*GetUserPreferencesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetUserPreferencesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) SetUserPreferences(ctx context.Context, in *SetUserPreferencesRequest, opts ...dcerpc.CallOption) (*SetUserPreferencesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetUserPreferencesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) GetSystemDirectory(ctx context.Context, in *GetSystemDirectoryRequest, opts ...dcerpc.CallOption) (*GetSystemDirectoryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSystemDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) SubmitRequest(ctx context.Context, in *SubmitRequestRequest, opts ...dcerpc.CallOption) (*SubmitRequestResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SubmitRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) GetInstalledProtocolsEx(ctx context.Context, in *GetInstalledProtocolsExRequest, opts ...dcerpc.CallOption) (*GetInstalledProtocolsExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInstalledProtocolsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...dcerpc.CallOption) (*GetVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRasrpcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewRasrpcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RasrpcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RasrpcSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRasrpcClient{cc: cc}, nil
}

// xxx_DeleteEntryOperation structure represents the RasRpcDeleteEntry operation
type xxx_DeleteEntryOperation struct {
	Phonebook string `idl:"name:lpszPhonebook;string" json:"phonebook"`
	Entry     string `idl:"name:lpszEntry;string" json:"entry"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteEntryOperation) OpNum() int { return 5 }

func (o *xxx_DeleteEntryOperation) OpName() string { return "/rasrpc/v1/RasRpcDeleteEntry" }

func (o *xxx_DeleteEntryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteEntryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszPhonebook {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Phonebook); err != nil {
			return err
		}
	}
	// lpszEntry {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Entry); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteEntryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszPhonebook {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phonebook); err != nil {
			return err
		}
	}
	// lpszEntry {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Entry); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteEntryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteEntryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteEntryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteEntryRequest structure represents the RasRpcDeleteEntry operation request
type DeleteEntryRequest struct {
	// lpszPhonebook: A null-terminated Unicode string specifying the RRAS Phonebook path
	// as specified in section 2.2.2.
	Phonebook string `idl:"name:lpszPhonebook;string" json:"phonebook"`
	// lpszEntry: A null-terminated Unicode string specifying the RRAS Entry name as specified
	// in section 2.2.2.1 to be deleted.
	Entry string `idl:"name:lpszEntry;string" json:"entry"`
}

func (o *DeleteEntryRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteEntryOperation) *xxx_DeleteEntryOperation {
	if op == nil {
		op = &xxx_DeleteEntryOperation{}
	}
	if o == nil {
		return op
	}
	op.Phonebook = o.Phonebook
	op.Entry = o.Entry
	return op
}

func (o *DeleteEntryRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteEntryOperation) {
	if o == nil {
		return
	}
	o.Phonebook = op.Phonebook
	o.Entry = op.Entry
}
func (o *DeleteEntryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteEntryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteEntryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteEntryResponse structure represents the RasRpcDeleteEntry operation response
type DeleteEntryResponse struct {
	// Return: The RasRpcDeleteEntry return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteEntryResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteEntryOperation) *xxx_DeleteEntryOperation {
	if op == nil {
		op = &xxx_DeleteEntryOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *DeleteEntryResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteEntryOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteEntryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteEntryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteEntryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUserPreferencesOperation structure represents the RasRpcGetUserPreferences operation
type xxx_GetUserPreferencesOperation struct {
	User   *rrasm.User `idl:"name:pUser" json:"user"`
	Mode   uint32      `idl:"name:dwMode" json:"mode"`
	Return uint32      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUserPreferencesOperation) OpNum() int { return 9 }

func (o *xxx_GetUserPreferencesOperation) OpName() string {
	return "/rasrpc/v1/RasRpcGetUserPreferences"
}

func (o *xxx_GetUserPreferencesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserPreferencesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User != nil {
			if err := o.User.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.User{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserPreferencesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER,pointer=ref}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User == nil {
			o.User = &rrasm.User{}
		}
		if err := o.User.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserPreferencesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserPreferencesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User != nil {
			if err := o.User.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.User{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserPreferencesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER,pointer=ref}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User == nil {
			o.User = &rrasm.User{}
		}
		if err := o.User.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetUserPreferencesRequest structure represents the RasRpcGetUserPreferences operation request
type GetUserPreferencesRequest struct {
	// pUser: Pointer to the RASRPC_PBUSER (section 2.2.1.2.229) structure which on successful
	// return contains the configuration information on the RRAS server.
	User *rrasm.User `idl:"name:pUser" json:"user"`
	// dwMode: This MUST be set to 2.
	Mode uint32 `idl:"name:dwMode" json:"mode"`
}

func (o *GetUserPreferencesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetUserPreferencesOperation) *xxx_GetUserPreferencesOperation {
	if op == nil {
		op = &xxx_GetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.User = o.User
	op.Mode = o.Mode
	return op
}

func (o *GetUserPreferencesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.User = op.User
	o.Mode = op.Mode
}
func (o *GetUserPreferencesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetUserPreferencesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUserPreferencesResponse structure represents the RasRpcGetUserPreferences operation response
type GetUserPreferencesResponse struct {
	// pUser: Pointer to the RASRPC_PBUSER (section 2.2.1.2.229) structure which on successful
	// return contains the configuration information on the RRAS server.
	User *rrasm.User `idl:"name:pUser" json:"user"`
	// Return: The RasRpcGetUserPreferences return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetUserPreferencesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetUserPreferencesOperation) *xxx_GetUserPreferencesOperation {
	if op == nil {
		op = &xxx_GetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.User = o.User
	op.Return = o.Return
	return op
}

func (o *GetUserPreferencesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.User = op.User
	o.Return = op.Return
}
func (o *GetUserPreferencesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetUserPreferencesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetUserPreferencesOperation structure represents the RasRpcSetUserPreferences operation
type xxx_SetUserPreferencesOperation struct {
	User   *rrasm.User `idl:"name:pUser" json:"user"`
	Mode   uint32      `idl:"name:dwMode" json:"mode"`
	Return uint32      `idl:"name:Return" json:"return"`
}

func (o *xxx_SetUserPreferencesOperation) OpNum() int { return 10 }

func (o *xxx_SetUserPreferencesOperation) OpName() string {
	return "/rasrpc/v1/RasRpcSetUserPreferences"
}

func (o *xxx_SetUserPreferencesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserPreferencesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pUser {in} (1:{alias=LPRASRPC_PBUSER}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User != nil {
			if err := o.User.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.User{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserPreferencesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pUser {in} (1:{alias=LPRASRPC_PBUSER,pointer=ref}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User == nil {
			o.User = &rrasm.User{}
		}
		if err := o.User.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserPreferencesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserPreferencesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserPreferencesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetUserPreferencesRequest structure represents the RasRpcSetUserPreferences operation request
type SetUserPreferencesRequest struct {
	// pUser: Pointer to the RASRPC_PBUSER (section 2.2.1.2.229) structure which on successful
	// return contains the configuration information on the RRAS server.
	User *rrasm.User `idl:"name:pUser" json:"user"`
	// dwMode: This MUST be set to 2.
	Mode uint32 `idl:"name:dwMode" json:"mode"`
}

func (o *SetUserPreferencesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetUserPreferencesOperation) *xxx_SetUserPreferencesOperation {
	if op == nil {
		op = &xxx_SetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.User = o.User
	op.Mode = o.Mode
	return op
}

func (o *SetUserPreferencesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.User = op.User
	o.Mode = op.Mode
}
func (o *SetUserPreferencesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetUserPreferencesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetUserPreferencesResponse structure represents the RasRpcSetUserPreferences operation response
type SetUserPreferencesResponse struct {
	// Return: The RasRpcSetUserPreferences return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetUserPreferencesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetUserPreferencesOperation) *xxx_SetUserPreferencesOperation {
	if op == nil {
		op = &xxx_SetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetUserPreferencesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetUserPreferencesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetUserPreferencesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSystemDirectoryOperation structure represents the RasRpcGetSystemDirectory operation
type xxx_GetSystemDirectoryOperation struct {
	Buffer string `idl:"name:lpBuffer;size_is:(uSize);string" json:"buffer"`
	Size   uint32 `idl:"name:uSize" json:"size"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSystemDirectoryOperation) OpNum() int { return 11 }

func (o *xxx_GetSystemDirectoryOperation) OpName() string {
	return "/rasrpc/v1/RasRpcGetSystemDirectory"
}

func (o *xxx_GetSystemDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != "" && o.Size == 0 {
		o.Size = uint32(ndr.UTF16NLen(o.Buffer))
	}
	if o.Size > uint32(260) {
		return fmt.Errorf("Size is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpBuffer {in, out} (1:{string, alias=LPWSTR}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
		dimSize1 := uint64(o.Size)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.Buffer)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0]-1 {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]-1]
		}
		if o.Buffer != ndr.ZeroString {
			_Buffer_buf = append(_Buffer_buf, uint16(0))
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// uSize {in} (1:{range=(0,260), alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpBuffer {in, out} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	// uSize {in} (1:{range=(0,260), alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpBuffer {in, out} (1:{string, alias=LPWSTR}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
		dimSize1 := uint64(o.Size)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.Buffer)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0]-1 {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]-1]
		}
		if o.Buffer != ndr.ZeroString {
			_Buffer_buf = append(_Buffer_buf, uint16(0))
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpBuffer {in, out} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	// Return {out} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSystemDirectoryRequest structure represents the RasRpcGetSystemDirectory operation request
type GetSystemDirectoryRequest struct {
	// lpBuffer: A null-terminated Unicode string that is populated with the path of the
	// system directory. The length of the string MUST be equal to uSize.
	Buffer string `idl:"name:lpBuffer;size_is:(uSize);string" json:"buffer"`
	// uSize: Specifies the size of the lpBuffer in Unicode characters. This value MUST
	// be equal to 260.
	Size uint32 `idl:"name:uSize" json:"size"`
}

func (o *GetSystemDirectoryRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSystemDirectoryOperation) *xxx_GetSystemDirectoryOperation {
	if op == nil {
		op = &xxx_GetSystemDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.Buffer = o.Buffer
	op.Size = o.Size
	return op
}

func (o *GetSystemDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSystemDirectoryOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Size = op.Size
}
func (o *GetSystemDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSystemDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSystemDirectoryResponse structure represents the RasRpcGetSystemDirectory operation response
type GetSystemDirectoryResponse struct {
	// XXX: uSize is an implicit input depedency for output parameters
	Size uint32 `idl:"name:uSize" json:"size"`

	// lpBuffer: A null-terminated Unicode string that is populated with the path of the
	// system directory. The length of the string MUST be equal to uSize.
	Buffer string `idl:"name:lpBuffer;size_is:(uSize);string" json:"buffer"`
	// Return: The RasRpcGetSystemDirectory return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSystemDirectoryResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSystemDirectoryOperation) *xxx_GetSystemDirectoryOperation {
	if op == nil {
		op = &xxx_GetSystemDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Size == uint32(0) {
		op.Size = o.Size
	}

	op.Buffer = o.Buffer
	op.Return = o.Return
	return op
}

func (o *GetSystemDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSystemDirectoryOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Size = op.Size

	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *GetSystemDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSystemDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SubmitRequestOperation structure represents the RasRpcSubmitRequest operation
type xxx_SubmitRequestOperation struct {
	RequestBuffer []byte `idl:"name:pReqBuffer;size_is:(dwcbBufSize);pointer:unique" json:"request_buffer"`
	BufferSize    uint32 `idl:"name:dwcbBufSize" json:"buffer_size"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SubmitRequestOperation) OpNum() int { return 12 }

func (o *xxx_SubmitRequestOperation) OpName() string { return "/rasrpc/v1/RasRpcSubmitRequest" }

func (o *xxx_SubmitRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RequestBuffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.RequestBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		if o.RequestBuffer != nil || o.BufferSize > 0 {
			_ptr_pReqBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RequestBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.RequestBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.RequestBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RequestBuffer, _ptr_pReqBuffer); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwcbBufSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		_ptr_pReqBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RequestBuffer", sizeInfo[0])
			}
			o.RequestBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.RequestBuffer {
				i1 := i1
				if err := w.ReadData(&o.RequestBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pReqBuffer := func(ptr interface{}) { o.RequestBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RequestBuffer, _s_pReqBuffer, _ptr_pReqBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwcbBufSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		if o.RequestBuffer != nil || o.BufferSize > 0 {
			_ptr_pReqBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RequestBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.RequestBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.RequestBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RequestBuffer, _ptr_pReqBuffer); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		_ptr_pReqBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RequestBuffer", sizeInfo[0])
			}
			o.RequestBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.RequestBuffer {
				i1 := i1
				if err := w.ReadData(&o.RequestBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pReqBuffer := func(ptr interface{}) { o.RequestBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RequestBuffer, _s_pReqBuffer, _ptr_pReqBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SubmitRequestRequest structure represents the RasRpcSubmitRequest operation request
type SubmitRequestRequest struct {
	// pReqBuffer: A pointer to a buffer of size dwcbBufSize. The buffer MUST be large enough
	// to hold the RequestBuffer structure (section 2.2.1.2.217) and RequestBuffer.RB_Buffer
	// data. RequestBuffer.RB_Reqtype specifies the request type which will be processed
	// by the server and RequestBuffer.RB_Buffer specifies the structure specific to RB_Reqtype
	// to be processed. The structure that MUST be used for each ReqTypes value is explained
	// in section 2.2.1.2.217. The client MUST NOT send the ReqType other than those defined
	// in ReqTypes (section 2.2.1.1.18). RequestBuffer.RB_PCBIndex MUST be set to the unique
	// port identifier whose information is sought for ReqTypes REQTYPE_GETINFO and REQTYPE_GETDEVCONFIG.
	// For other valid ReqTypes, RequestBuffer.RB_PCBIndex MUST be set to zero (0).
	RequestBuffer []byte `idl:"name:pReqBuffer;size_is:(dwcbBufSize);pointer:unique" json:"request_buffer"`
	// dwcbBufSize: Size in byte of pReqBuffer.
	BufferSize uint32 `idl:"name:dwcbBufSize" json:"buffer_size"`
}

func (o *SubmitRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_SubmitRequestOperation) *xxx_SubmitRequestOperation {
	if op == nil {
		op = &xxx_SubmitRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.RequestBuffer = o.RequestBuffer
	op.BufferSize = o.BufferSize
	return op
}

func (o *SubmitRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_SubmitRequestOperation) {
	if o == nil {
		return
	}
	o.RequestBuffer = op.RequestBuffer
	o.BufferSize = op.BufferSize
}
func (o *SubmitRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SubmitRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SubmitRequestResponse structure represents the RasRpcSubmitRequest operation response
type SubmitRequestResponse struct {
	// XXX: dwcbBufSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:dwcbBufSize" json:"buffer_size"`

	// pReqBuffer: A pointer to a buffer of size dwcbBufSize. The buffer MUST be large enough
	// to hold the RequestBuffer structure (section 2.2.1.2.217) and RequestBuffer.RB_Buffer
	// data. RequestBuffer.RB_Reqtype specifies the request type which will be processed
	// by the server and RequestBuffer.RB_Buffer specifies the structure specific to RB_Reqtype
	// to be processed. The structure that MUST be used for each ReqTypes value is explained
	// in section 2.2.1.2.217. The client MUST NOT send the ReqType other than those defined
	// in ReqTypes (section 2.2.1.1.18). RequestBuffer.RB_PCBIndex MUST be set to the unique
	// port identifier whose information is sought for ReqTypes REQTYPE_GETINFO and REQTYPE_GETDEVCONFIG.
	// For other valid ReqTypes, RequestBuffer.RB_PCBIndex MUST be set to zero (0).
	RequestBuffer []byte `idl:"name:pReqBuffer;size_is:(dwcbBufSize);pointer:unique" json:"request_buffer"`
	// Return: The RasRpcSubmitRequest return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SubmitRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_SubmitRequestOperation) *xxx_SubmitRequestOperation {
	if op == nil {
		op = &xxx_SubmitRequestOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.RequestBuffer = o.RequestBuffer
	op.Return = o.Return
	return op
}

func (o *SubmitRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_SubmitRequestOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.RequestBuffer = op.RequestBuffer
	o.Return = op.Return
}
func (o *SubmitRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SubmitRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInstalledProtocolsExOperation structure represents the RasRpcGetInstalledProtocolsEx operation
type xxx_GetInstalledProtocolsExOperation struct {
	Router    bool   `idl:"name:fRouter" json:"router"`
	RASCli    bool   `idl:"name:fRasCli" json:"ras_cli"`
	RASServer bool   `idl:"name:fRasSrv" json:"ras_server"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInstalledProtocolsExOperation) OpNum() int { return 14 }

func (o *xxx_GetInstalledProtocolsExOperation) OpName() string {
	return "/rasrpc/v1/RasRpcGetInstalledProtocolsEx"
}

func (o *xxx_GetInstalledProtocolsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledProtocolsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fRouter {in} (1:{alias=BOOL}(int32))
	{
		if !o.Router {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// fRasCli {in} (1:{alias=BOOL}(int32))
	{
		if !o.RASCli {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// fRasSrv {in} (1:{alias=BOOL}(int32))
	{
		if !o.RASServer {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetInstalledProtocolsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fRouter {in} (1:{alias=BOOL}(int32))
	{
		var _bRouter int32
		if err := w.ReadData(&_bRouter); err != nil {
			return err
		}
		o.Router = _bRouter != 0
	}
	// fRasCli {in} (1:{alias=BOOL}(int32))
	{
		var _bRASCli int32
		if err := w.ReadData(&_bRASCli); err != nil {
			return err
		}
		o.RASCli = _bRASCli != 0
	}
	// fRasSrv {in} (1:{alias=BOOL}(int32))
	{
		var _bRASServer int32
		if err := w.ReadData(&_bRASServer); err != nil {
			return err
		}
		o.RASServer = _bRASServer != 0
	}
	return nil
}

func (o *xxx_GetInstalledProtocolsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledProtocolsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstalledProtocolsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInstalledProtocolsExRequest structure represents the RasRpcGetInstalledProtocolsEx operation request
type GetInstalledProtocolsExRequest struct {
	// fRouter: If set to TRUE, protocols enabled for Demand Dial are retrieved. If set
	// to FALSE, protocols enabled for Demand Dial are not retrieved.
	Router bool `idl:"name:fRouter" json:"router"`
	// fRasCli: This flag is not used and MUST be set to FALSE.
	RASCli bool `idl:"name:fRasCli" json:"ras_cli"`
	// fRasSrv: If set to TRUE, retrieves the protocol enabled for RRAS incoming connections.
	// If set to FALSE, protocol for RRAS incoming connections are not retrieved.
	RASServer bool `idl:"name:fRasSrv" json:"ras_server"`
}

func (o *GetInstalledProtocolsExRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInstalledProtocolsExOperation) *xxx_GetInstalledProtocolsExOperation {
	if op == nil {
		op = &xxx_GetInstalledProtocolsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Router = o.Router
	op.RASCli = o.RASCli
	op.RASServer = o.RASServer
	return op
}

func (o *GetInstalledProtocolsExRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInstalledProtocolsExOperation) {
	if o == nil {
		return
	}
	o.Router = op.Router
	o.RASCli = op.RASCli
	o.RASServer = op.RASServer
}
func (o *GetInstalledProtocolsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInstalledProtocolsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstalledProtocolsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInstalledProtocolsExResponse structure represents the RasRpcGetInstalledProtocolsEx operation response
type GetInstalledProtocolsExResponse struct {
	// Return: The RasRpcGetInstalledProtocolsEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetInstalledProtocolsExResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInstalledProtocolsExOperation) *xxx_GetInstalledProtocolsExOperation {
	if op == nil {
		op = &xxx_GetInstalledProtocolsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *GetInstalledProtocolsExResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInstalledProtocolsExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *GetInstalledProtocolsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInstalledProtocolsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstalledProtocolsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionOperation structure represents the RasRpcGetVersion operation
type xxx_GetVersionOperation struct {
	Version uint32 `idl:"name:pdwVersion;pointer:ref" json:"version"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionOperation) OpNum() int { return 15 }

func (o *xxx_GetVersionOperation) OpName() string { return "/rasrpc/v1/RasRpcGetVersion" }

func (o *xxx_GetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetVersionRequest structure represents the RasRpcGetVersion operation request
type GetVersionRequest struct {
	// pdwVersion: This is a pointer to type DWORD which, after a successful function call,
	// specifies the version of the Rasrpc interface.
	Version uint32 `idl:"name:pdwVersion;pointer:ref" json:"version"`
}

func (o *GetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	return op
}

func (o *GetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
}
func (o *GetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionResponse structure represents the RasRpcGetVersion operation response
type GetVersionResponse struct {
	// pdwVersion: This is a pointer to type DWORD which, after a successful function call,
	// specifies the version of the Rasrpc interface.
	Version uint32 `idl:"name:pdwVersion;pointer:ref" json:"version"`
	// Return: The RasRpcGetVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	op.Return = o.Return
	return op
}

func (o *GetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.Return = op.Return
}
func (o *GetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

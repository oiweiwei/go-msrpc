package efsrpc

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

// efsrpc server interface.
type EfsrpcServer interface {

	// The EfsRpcOpenFileRaw method is used to open an encrypted object on the server for
	// backup or restore. It allocates resources that MUST be released by calling the EfsRpcCloseRaw
	// method.<42>
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	//
	// First, the server SHOULD perform any additional access checks prescribed by the implementation.
	// If any of these checks fail, it MUST return a nonzero value.
	//
	// EFSRPC servers SHOULD return an error unless at least one of the following conditions
	// is true:
	//
	// * The calling user has a private key ( 230807ac-20be-494f-86e3-4c8ac23ea584#gt_6fca10f4-e829-42ab-ad40-1566585060ca
	// ) that grants the user authorized access to the file.
	//
	// * The *CREATE_FOR_IMPORT* flag is set, and the user has restore rights on the server.
	//
	// * The *CREATE_FOR_IMPORT* flag is not set, and the user has backup rights on the
	// server.
	//
	// If the CREATE_FOR_IMPORT flag is set, the server MUST attempt to create an object
	// with the given name and prepare it for writing data received in future EfsRpcWriteFileRaw
	// calls. The server MUST return a nonzero value if this fails.
	//
	// If the CREATE_FOR_IMPORT flag is not set, the server MUST attempt to locate the object
	// requested and prepare it for reading data to be sent through future EfsRpcReadFileRaw
	// calls. The server MUST return a nonzero value if it fails.
	//
	// If the server supports the CREATE_FOR_DIR flag, and this flag is set:
	//
	// * If the *CREATE_FOR_IMPORT* flag is not set:
	//
	// * If the data object referred to by FileName exists on the server and is not a container
	// for other objects, the server SHOULD return a nonzero value.
	//
	// * Otherwise, the server SHOULD ignore the *CREATE_FOR_DIR* flag.
	//
	// * If the *CREATE_FOR_IMPORT* flag is set, the server MUST attempt to create a container
	// with the given name and prepare it for writing data received in future EfsRpcWriteFileRaw
	// calls. The server MUST return a nonzero value if this fails.
	//
	// If the server supports the OVERWRITE_HIDDEN flag, and this flag is set:
	//
	// * If the *CREATE_FOR_IMPORT* flag is not set, the server SHOULD ignore this flag.
	//
	// * If the *CREATE_FOR_IMPORT* flag is set, the server SHOULD overwrite an existing
	// object even if the object is "hidden". The meaning of "hidden" is specific to the
	// implementation of the data store, and this meaning does not affect protocol behavior.
	OpenFileRaw(context.Context, *OpenFileRawRequest) (*OpenFileRawResponse, error)

	// The method EfsRpcReadFileRaw is used by a client to obtain marshaled data for an
	// encrypted object from the server.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client.
	//
	// If called with a context handle that has not been obtained by calling the EfsRpcOpenFileRaw
	// method without the CREATE_FOR_IMPORT flag set, the server SHOULD throw an RPC exception.
	ReadFileRaw(context.Context, *ReadFileRawRequest) (*ReadFileRawResponse, error)

	// The method EfsRpcWriteFileRaw is used to create an encrypted object on the server
	// from the marshaled data provided by the client.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client.
	//
	// If called with a context handle that has not been obtained by calling EfsRpcOpenFileRaw
	// with the CREATE_FOR_IMPORT flag set, the server MUST abort the operation. In this
	// case, it SHOULD throw an RPC exception.
	WriteFileRaw(context.Context, *WriteFileRawRequest) (*WriteFileRawResponse, error)

	// The EfsRpcCloseRaw method is called to release any resources allocated by the EfsRpcOpenFileRaw
	// method, or by subsequent calls to the EfsRpcReadFileRaw or EfsRpcWriteFileRaw methods.
	//
	// Return Values: This method has no return values.
	CloseRaw(context.Context, *CloseRawRequest) (*CloseRawResponse, error)

	// EfsRpcEncryptFileSrv operation.
	EncryptFileServer(context.Context, *EncryptFileServerRequest) (*EncryptFileServerResponse, error)

	// The EfsRpcDecryptFileSrv method is used to convert an existing encrypted object to
	// the unencrypted state in the server's data store.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	//
	// If no object exists on the server with the specified name, the server MUST return
	// a nonzero value. If the object exists and is not encrypted, the server MUST return
	// success.
	//
	// Otherwise, the server performs the following actions to convert the object in its
	// data store to an unencrypted state:
	//
	// * If the data object referred to by FileName is a container for other objects, the
	// server MUST clear the attribute on the container that instructs the data store to
	// encrypt any new objects created in that container. The server MAY decrypt encrypted
	// objects that were already in the container before this message was received.
	//
	// * Otherwise, the server SHOULD:
	//
	// * Check that the calling user has access to a private key that will decrypt the file;
	// if the user does not have access, return a nonzero value.
	//
	// * Decrypt the object and discard its EFSRPC Metadata ( 3166cf4a-e085-47be-94c6-b69bddf274ff
	// ).
	//
	// * Return 0 to indicate success.
	DecryptFileServer(context.Context, *DecryptFileServerRequest) (*DecryptFileServerResponse, error)

	// EfsRpcQueryUsersOnFile operation.
	QueryUsersOnFile(context.Context, *QueryUsersOnFileRequest) (*QueryUsersOnFileResponse, error)

	// The EfsRpcQueryRecoveryAgents method is used to query the EFSRPC Metadata of an encrypted
	// object for the X.509 certificates of the data recovery agents whose private keys
	// can be used to decrypt the object.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	QueryRecoveryAgents(context.Context, *QueryRecoveryAgentsRequest) (*QueryRecoveryAgentsResponse, error)

	// The EfsRpcRemoveUsersFromFile method is used to revoke a user's access to an encrypted
	// object. This method revokes the ability of the private key corresponding to a given
	// X.509 certificate to decrypt the object.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	//
	// If none of the preceding errors occur, the server MUST remove the parts of the object's
	// EFSRPC Metadata that refer to the user certificates listed in the Users structure.
	RemoveUsersFromFile(context.Context, *RemoveUsersFromFileRequest) (*RemoveUsersFromFileResponse, error)

	// The EfsRpcAddUsersToFile method is used to grant the possessors of the private keys
	// corresponding to certain X.509 certificates the ability to decrypt the object.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	//
	// If no object exists on the server with the specified name, or if the object exists
	// and is not encrypted, the server MUST return a nonzero value. Otherwise, the server
	// MUST modify the object's EFSRPC Metadata such that all the user certificates listed
	// in the Users structure have the ability to decrypt the object.
	AddUsersToFile(context.Context, *AddUsersToFileRequest) (*AddUsersToFileResponse, error)

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Return Values: The EFSRPC server SHOULD return a nonzero value. However, the server
	// MAY<45> process this as described in section 3.1.4.2.13.
	NotSupported(context.Context, *NotSupportedRequest) (*NotSupportedResponse, error)

	// The EfsRpcFileKeyInfo method is used to query and modify information about the keys
	// used to encrypt a given object.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	//
	// If no object exists on the server with the specified name the server MUST return
	// a nonzero value.
	//
	// If the InfoClass parameter is not equal to CHECK_ENCRYPTION_STATUS and the object
	// with the specified name is not encrypted, the server MUST return a nonzero value.
	//
	// If the value in the InfoClass parameter is unsupported by the server, the server
	// MUST return a nonzero value.
	//
	// If the value in the InfoClass parameter is equal to BASIC_KEY_INFO, the server SHOULD
	// read the EFSRPC Metadata of the object referred to by the FileName argument and return
	// information about its FEK in an EFS_KEY_INFO structure within the KeyInfo argument.
	//
	// If the value in the InfoClass parameter is equal to UPDATE_KEY_USED, the implementation
	// supports this value, and the FileName parameter does not satisfy the implementation-specific
	// requirements for this operation<47>, the server MUST return a nonzero value.
	//
	// If the value in the InfoClass parameter is equal to UPDATE_KEY_USED, the implementation
	// supports this value, and the FileName parameter does satisfy all implementation-specific
	// requirements, the server MUST update the EFSRPC Metadata of all the data objects
	// referred by FileName in an implementation-specific way<48>, and return a newline-separated
	// list of EFSRPC Identifiers thus updated in the KeyInfo parameter.
	//
	// If the value in the InfoClass parameter is equal to CHECK_ENCRYPTION_STATUS, the
	// server MUST return an EFS_ENCRYPTION_STATUS_INFO structure in the KeyInfo parameter,
	// which provides a hint to the client what error code would be returned if encryption
	// was attempted on this object without any further user interaction or higher-level
	// events.
	//
	// If the value in the InfoClass parameter is equal to CHECK_DECRYPTION_STATUS, the
	// server SHOULD return ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION ([MS-ERREF] section
	// 2.2). The server MAY, instead, return an EFS_DECRYPTION_STATUS_INFO structure in
	// the KeyInfo parameter, which provides a hint to the client what error code would
	// be returned if decryption were attempted on this object without any further user
	// interaction or higher-level events.
	FileKeyInfo(context.Context, *FileKeyInfoRequest) (*FileKeyInfoResponse, error)

	// The EfsRpcDuplicateEncryptionInfoFile method is used to duplicate the EFSRPC Metadata
	// of one encrypted object and attach it to another encrypted object. This is typically
	// done when copying objects to maintain the same set of keys and users for the copy
	// as for the original.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.<50>
	//
	// If no object exists on the server with the name specified in the SrcFileName parameter,
	// or if it exists and is not encrypted, the server MUST return a nonzero value.
	//
	// * If an object exists with the name specified in the DestFileName parameter, the
	// server MUST return a nonzero value.
	//
	// * If no object exists with the name specified in the DestFileName parameter, the
	// server MUST create a new object with this name and duplicate the EFSRPC Metadata
	// from the SrcFileName parameter into it. If the object specified in SrcFileName is
	// a container for other objects, the server MUST create the object as a container for
	// objects, and it MUST encrypt any objects that are subsequently placed in the container
	// after this message has been processed. Otherwise, the server MUST create the object
	// as a non-container encrypted data object.
	//
	// If an encrypted object exists with the name specified in the SrcFileName and dwCreationDisposition
	// parameters is not equal to CREATE_NEW, then:
	//
	// * If an object already exists with the name specified in the DestFileName parameter,
	// the server MUST check whether the object referred to by SrcFileName is of the same
	// type (either simple object or container for other objects); if the object is not
	// of the same type, the server MUST return a nonzero value. In addition, if the object
	// referred to by DestFileName is a container for other objects, and it is not already
	// encrypted, the server MUST return a nonzero value. Otherwise, the server SHOULD overwrite
	// the object, clear its existing attributes, create a new object in its place with
	// the attributes specified, and duplicate the EFSRPC Metadata from the SrcFileName
	// parameter into it.
	//
	// * If no object exists with the name specified in the DestFileName parameter, the
	// server MUST create a new object with this name and duplicate the EFSRPC Metadata
	// from the SrcFileName parameter into it. If the object specified in SrcFileName is
	// a container for other objects, the server MUST create the object as a container for
	// objects, and it MUST encrypt any objects that are subsequently placed in the container
	// after this message has been processed. Otherwise, the server MUST create the object
	// as a non-container encrypted data object.
	//
	// In duplicating the EFSRPC Metadata from the SrcFileName parameter to the DestFileName
	// parameter, the server MAY<51> change the metadata. However, upon successful completion,
	// the set of users and DRAs with access to the DestFileName parameter MUST be the same
	// set of users who had access to the SrcFileName parameter at the outset.
	DuplicateEncryptionInfoFile(context.Context, *DuplicateEncryptionInfoFileRequest) (*DuplicateEncryptionInfoFileResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	// The EfsRpcAddUsersToFileEx method is used to grant the possessors of the private
	// keys corresponding to certain X.509 certificates the ability to decrypt the object.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	//
	// If no object exists on the server with the specified name, or if it exists and is
	// not encrypted, the server MUST return a nonzero value.
	//
	// If the EFSRPC_ADDUSERFLAG_REPLACE_DDF flag is set in the dwFlags parameter, and the
	// EncryptionCertificates parameter contains more than one certificate, the server MUST
	// return a nonzero value.
	//
	// If the EFSRPC_ADDUSERFLAG_REPLACE_DDF flag is set in the dwFlags parameter, and the
	// calling user does not have the ability to decrypt the object, the server MUST return
	// a nonzero value.
	//
	// If the EFSRPC_ADDUSERFLAG_ADD_POLICY_KEYTYPE flag is specified in the dwFlags parameter,
	// then for each certificate specified in the EncryptionCertificates parameter, the
	// server MUST check whether the private key for the certificate is stored on a smart
	// card. If the key is stored, the server MUST return a nonzero value; otherwise, the
	// server MUST ignore this flag.
	//
	// If the EFSRPC_ADDUSERFLAG_REPLACE_DDF flag is set in the dwFlags parameter, and the
	// calling user has the ability to decrypt the object, then the certificate in the EncryptionCertificates
	// parameter is to be given access to the object, replacing one of the calling user's
	// user certificates through which he currently has access.
	AddUsersToFileEx(context.Context, *AddUsersToFileExRequest) (*AddUsersToFileExResponse, error)

	// Return Values: The server SHOULD return a nonzero value.<54>
	FileKeyInfoEx(context.Context, *FileKeyInfoExRequest) (*FileKeyInfoExResponse, error)

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Return Values: The server SHOULD return a nonzero value.<57>
	GetEncryptedFileMetadata(context.Context, *GetEncryptedFileMetadataRequest) (*GetEncryptedFileMetadataResponse, error)

	// Return Values: The server SHOULD return a nonzero value.<62>
	SetEncryptedFileMetadata(context.Context, *SetEncryptedFileMetadataRequest) (*SetEncryptedFileMetadataResponse, error)

	// The EfsRpcFlushEfsCache method causes EFS to flush the logical cache that holds all
	// the sensitive information required to perform EFSRPC operations for the calling user.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	FlushEFSCache(context.Context, *FlushEFSCacheRequest) (*FlushEFSCacheResponse, error)

	// EfsRpcEncryptFileExSrv operation.
	EncryptFileExServer(context.Context, *EncryptFileExServerRequest) (*EncryptFileExServerResponse, error)

	// EfsRpcQueryProtectors operation.
	QueryProtectors(context.Context, *QueryProtectorsRequest) (*QueryProtectorsResponse, error)

	// Opnum23NotUsedOnWire operation.
	// Opnum23NotUsedOnWire

	// Opnum24NotUsedOnWire operation.
	// Opnum24NotUsedOnWire

	// Opnum25NotUsedOnWire operation.
	// Opnum25NotUsedOnWire

	// Opnum26NotUsedOnWire operation.
	// Opnum26NotUsedOnWire

	// Opnum27NotUsedOnWire operation.
	// Opnum27NotUsedOnWire

	// Opnum28NotUsedOnWire operation.
	// Opnum28NotUsedOnWire

	// Opnum29NotUsedOnWire operation.
	// Opnum29NotUsedOnWire

	// Opnum30NotUsedOnWire operation.
	// Opnum30NotUsedOnWire

	// Opnum31NotUsedOnWire operation.
	// Opnum31NotUsedOnWire

	// Opnum32NotUsedOnWire operation.
	// Opnum32NotUsedOnWire

	// Opnum33NotUsedOnWire operation.
	// Opnum33NotUsedOnWire

	// Opnum34NotUsedOnWire operation.
	// Opnum34NotUsedOnWire

	// Opnum35NotUsedOnWire operation.
	// Opnum35NotUsedOnWire

	// Opnum36NotUsedOnWire operation.
	// Opnum36NotUsedOnWire

	// Opnum37NotUsedOnWire operation.
	// Opnum37NotUsedOnWire

	// Opnum38NotUsedOnWire operation.
	// Opnum38NotUsedOnWire

	// Opnum39NotUsedOnWire operation.
	// Opnum39NotUsedOnWire

	// Opnum40NotUsedOnWire operation.
	// Opnum40NotUsedOnWire

	// Opnum41NotUsedOnWire operation.
	// Opnum41NotUsedOnWire

	// Opnum42NotUsedOnWire operation.
	// Opnum42NotUsedOnWire

	// Opnum43NotUsedOnWire operation.
	// Opnum43NotUsedOnWire

	// Opnum44NotUsedOnWire operation.
	// Opnum44NotUsedOnWire
}

func RegisterEfsrpcServer(conn dcerpc.Conn, o EfsrpcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEfsrpcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EfsrpcSyntaxV1_0))...)
}

func NewEfsrpcServerHandle(o EfsrpcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EfsrpcServerHandle(ctx, o, opNum, r)
	}
}

func EfsrpcServerHandle(ctx context.Context, o EfsrpcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // EfsRpcOpenFileRaw
		op := &xxx_OpenFileRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenFileRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenFileRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // EfsRpcReadFileRaw
		op := &xxx_ReadFileRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadFileRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadFileRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // EfsRpcWriteFileRaw
		op := &xxx_WriteFileRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteFileRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteFileRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // EfsRpcCloseRaw
		op := &xxx_CloseRawOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseRawRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseRaw(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // EfsRpcEncryptFileSrv
		op := &xxx_EncryptFileServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EncryptFileServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EncryptFileServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // EfsRpcDecryptFileSrv
		op := &xxx_DecryptFileServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DecryptFileServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DecryptFileServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // EfsRpcQueryUsersOnFile
		op := &xxx_QueryUsersOnFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryUsersOnFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryUsersOnFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // EfsRpcQueryRecoveryAgents
		op := &xxx_QueryRecoveryAgentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryRecoveryAgentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryRecoveryAgents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // EfsRpcRemoveUsersFromFile
		op := &xxx_RemoveUsersFromFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveUsersFromFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveUsersFromFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // EfsRpcAddUsersToFile
		op := &xxx_AddUsersToFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddUsersToFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddUsersToFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // EfsRpcNotSupported
		op := &xxx_NotSupportedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NotSupportedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NotSupported(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // EfsRpcFileKeyInfo
		op := &xxx_FileKeyInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FileKeyInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FileKeyInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // EfsRpcDuplicateEncryptionInfoFile
		op := &xxx_DuplicateEncryptionInfoFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DuplicateEncryptionInfoFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DuplicateEncryptionInfoFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	case 15: // EfsRpcAddUsersToFileEx
		op := &xxx_AddUsersToFileExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddUsersToFileExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddUsersToFileEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // EfsRpcFileKeyInfoEx
		op := &xxx_FileKeyInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FileKeyInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FileKeyInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 18: // EfsRpcGetEncryptedFileMetadata
		op := &xxx_GetEncryptedFileMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEncryptedFileMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEncryptedFileMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // EfsRpcSetEncryptedFileMetadata
		op := &xxx_SetEncryptedFileMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEncryptedFileMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEncryptedFileMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // EfsRpcFlushEfsCache
		op := &xxx_FlushEFSCacheOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FlushEFSCacheRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FlushEFSCache(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // EfsRpcEncryptFileExSrv
		op := &xxx_EncryptFileExServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EncryptFileExServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EncryptFileExServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // EfsRpcQueryProtectors
		op := &xxx_QueryProtectorsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryProtectorsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryProtectors(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // Opnum23NotUsedOnWire
		// Opnum23NotUsedOnWire
		return nil, nil
	case 24: // Opnum24NotUsedOnWire
		// Opnum24NotUsedOnWire
		return nil, nil
	case 25: // Opnum25NotUsedOnWire
		// Opnum25NotUsedOnWire
		return nil, nil
	case 26: // Opnum26NotUsedOnWire
		// Opnum26NotUsedOnWire
		return nil, nil
	case 27: // Opnum27NotUsedOnWire
		// Opnum27NotUsedOnWire
		return nil, nil
	case 28: // Opnum28NotUsedOnWire
		// Opnum28NotUsedOnWire
		return nil, nil
	case 29: // Opnum29NotUsedOnWire
		// Opnum29NotUsedOnWire
		return nil, nil
	case 30: // Opnum30NotUsedOnWire
		// Opnum30NotUsedOnWire
		return nil, nil
	case 31: // Opnum31NotUsedOnWire
		// Opnum31NotUsedOnWire
		return nil, nil
	case 32: // Opnum32NotUsedOnWire
		// Opnum32NotUsedOnWire
		return nil, nil
	case 33: // Opnum33NotUsedOnWire
		// Opnum33NotUsedOnWire
		return nil, nil
	case 34: // Opnum34NotUsedOnWire
		// Opnum34NotUsedOnWire
		return nil, nil
	case 35: // Opnum35NotUsedOnWire
		// Opnum35NotUsedOnWire
		return nil, nil
	case 36: // Opnum36NotUsedOnWire
		// Opnum36NotUsedOnWire
		return nil, nil
	case 37: // Opnum37NotUsedOnWire
		// Opnum37NotUsedOnWire
		return nil, nil
	case 38: // Opnum38NotUsedOnWire
		// Opnum38NotUsedOnWire
		return nil, nil
	case 39: // Opnum39NotUsedOnWire
		// Opnum39NotUsedOnWire
		return nil, nil
	case 40: // Opnum40NotUsedOnWire
		// Opnum40NotUsedOnWire
		return nil, nil
	case 41: // Opnum41NotUsedOnWire
		// Opnum41NotUsedOnWire
		return nil, nil
	case 42: // Opnum42NotUsedOnWire
		// Opnum42NotUsedOnWire
		return nil, nil
	case 43: // Opnum43NotUsedOnWire
		// Opnum43NotUsedOnWire
		return nil, nil
	case 44: // Opnum44NotUsedOnWire
		// Opnum44NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented efsrpc
type UnimplementedEfsrpcServer struct {
}

func (UnimplementedEfsrpcServer) OpenFileRaw(context.Context, *OpenFileRawRequest) (*OpenFileRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) ReadFileRaw(context.Context, *ReadFileRawRequest) (*ReadFileRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) WriteFileRaw(context.Context, *WriteFileRawRequest) (*WriteFileRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) CloseRaw(context.Context, *CloseRawRequest) (*CloseRawResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) EncryptFileServer(context.Context, *EncryptFileServerRequest) (*EncryptFileServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) DecryptFileServer(context.Context, *DecryptFileServerRequest) (*DecryptFileServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) QueryUsersOnFile(context.Context, *QueryUsersOnFileRequest) (*QueryUsersOnFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) QueryRecoveryAgents(context.Context, *QueryRecoveryAgentsRequest) (*QueryRecoveryAgentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) RemoveUsersFromFile(context.Context, *RemoveUsersFromFileRequest) (*RemoveUsersFromFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) AddUsersToFile(context.Context, *AddUsersToFileRequest) (*AddUsersToFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) NotSupported(context.Context, *NotSupportedRequest) (*NotSupportedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) FileKeyInfo(context.Context, *FileKeyInfoRequest) (*FileKeyInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) DuplicateEncryptionInfoFile(context.Context, *DuplicateEncryptionInfoFileRequest) (*DuplicateEncryptionInfoFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) AddUsersToFileEx(context.Context, *AddUsersToFileExRequest) (*AddUsersToFileExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) FileKeyInfoEx(context.Context, *FileKeyInfoExRequest) (*FileKeyInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) GetEncryptedFileMetadata(context.Context, *GetEncryptedFileMetadataRequest) (*GetEncryptedFileMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) SetEncryptedFileMetadata(context.Context, *SetEncryptedFileMetadataRequest) (*SetEncryptedFileMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) FlushEFSCache(context.Context, *FlushEFSCacheRequest) (*FlushEFSCacheResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) EncryptFileExServer(context.Context, *EncryptFileExServerRequest) (*EncryptFileExServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEfsrpcServer) QueryProtectors(context.Context, *QueryProtectorsRequest) (*QueryProtectorsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EfsrpcServer = (*UnimplementedEfsrpcServer)(nil)

package efsrpc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "efsr"
)

var (
	// Syntax UUID
	EfsrpcSyntaxUUID = &uuid.UUID{TimeLow: 0xc681d488, TimeMid: 0xd850, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x52, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0xf, 0x7e}}
	// Syntax ID
	EfsrpcSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: EfsrpcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// efsrpc interface.
type EfsrpcClient interface {

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
	OpenFileRaw(context.Context, *OpenFileRawRequest, ...dcerpc.CallOption) (*OpenFileRawResponse, error)

	// The method EfsRpcReadFileRaw is used by a client to obtain marshaled data for an
	// encrypted object from the server.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client.
	//
	// If called with a context handle that has not been obtained by calling the EfsRpcOpenFileRaw
	// method without the CREATE_FOR_IMPORT flag set, the server SHOULD throw an RPC exception.
	ReadFileRaw(context.Context, *ReadFileRawRequest, ...dcerpc.CallOption) (*ReadFileRawResponse, error)

	// The method EfsRpcWriteFileRaw is used to create an encrypted object on the server
	// from the marshaled data provided by the client.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client.
	//
	// If called with a context handle that has not been obtained by calling EfsRpcOpenFileRaw
	// with the CREATE_FOR_IMPORT flag set, the server MUST abort the operation. In this
	// case, it SHOULD throw an RPC exception.
	WriteFileRaw(context.Context, *WriteFileRawRequest, ...dcerpc.CallOption) (*WriteFileRawResponse, error)

	// The EfsRpcCloseRaw method is called to release any resources allocated by the EfsRpcOpenFileRaw
	// method, or by subsequent calls to the EfsRpcReadFileRaw or EfsRpcWriteFileRaw methods.
	//
	// Return Values: This method has no return values.
	CloseRaw(context.Context, *CloseRawRequest, ...dcerpc.CallOption) (*CloseRawResponse, error)

	// EfsRpcEncryptFileSrv operation.
	EncryptFileServer(context.Context, *EncryptFileServerRequest, ...dcerpc.CallOption) (*EncryptFileServerResponse, error)

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
	DecryptFileServer(context.Context, *DecryptFileServerRequest, ...dcerpc.CallOption) (*DecryptFileServerResponse, error)

	// EfsRpcQueryUsersOnFile operation.
	QueryUsersOnFile(context.Context, *QueryUsersOnFileRequest, ...dcerpc.CallOption) (*QueryUsersOnFileResponse, error)

	// The EfsRpcQueryRecoveryAgents method is used to query the EFSRPC Metadata of an encrypted
	// object for the X.509 certificates of the data recovery agents whose private keys
	// can be used to decrypt the object.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	QueryRecoveryAgents(context.Context, *QueryRecoveryAgentsRequest, ...dcerpc.CallOption) (*QueryRecoveryAgentsResponse, error)

	// The EfsRpcRemoveUsersFromFile method is used to revoke a user's access to an encrypted
	// object. This method revokes the ability of the private key corresponding to a given
	// X.509 certificate to decrypt the object.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	//
	// If none of the preceding errors occur, the server MUST remove the parts of the object's
	// EFSRPC Metadata that refer to the user certificates listed in the Users structure.
	RemoveUsersFromFile(context.Context, *RemoveUsersFromFileRequest, ...dcerpc.CallOption) (*RemoveUsersFromFileResponse, error)

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
	AddUsersToFile(context.Context, *AddUsersToFileRequest, ...dcerpc.CallOption) (*AddUsersToFileResponse, error)

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Return Values: The EFSRPC server SHOULD return a nonzero value. However, the server
	// MAY<45> process this as described in section 3.1.4.2.13.
	NotSupported(context.Context, *NotSupportedRequest, ...dcerpc.CallOption) (*NotSupportedResponse, error)

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
	FileKeyInfo(context.Context, *FileKeyInfoRequest, ...dcerpc.CallOption) (*FileKeyInfoResponse, error)

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
	DuplicateEncryptionInfoFile(context.Context, *DuplicateEncryptionInfoFileRequest, ...dcerpc.CallOption) (*DuplicateEncryptionInfoFileResponse, error)

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
	AddUsersToFileEx(context.Context, *AddUsersToFileExRequest, ...dcerpc.CallOption) (*AddUsersToFileExResponse, error)

	// Return Values: The server SHOULD return a nonzero value.<54>
	FileKeyInfoEx(context.Context, *FileKeyInfoExRequest, ...dcerpc.CallOption) (*FileKeyInfoExResponse, error)

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Return Values: The server SHOULD return a nonzero value.<57>
	GetEncryptedFileMetadata(context.Context, *GetEncryptedFileMetadataRequest, ...dcerpc.CallOption) (*GetEncryptedFileMetadataResponse, error)

	// Return Values: The server SHOULD return a nonzero value.<62>
	SetEncryptedFileMetadata(context.Context, *SetEncryptedFileMetadataRequest, ...dcerpc.CallOption) (*SetEncryptedFileMetadataResponse, error)

	// The EfsRpcFlushEfsCache method causes EFS to flush the logical cache that holds all
	// the sensitive information required to perform EFSRPC operations for the calling user.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client. The server MUST return a nonzero value if processing fails.
	FlushEFSCache(context.Context, *FlushEFSCacheRequest, ...dcerpc.CallOption) (*FlushEFSCacheResponse, error)

	// EfsRpcEncryptFileExSrv operation.
	EncryptFileExServer(context.Context, *EncryptFileExServerRequest, ...dcerpc.CallOption) (*EncryptFileExServerResponse, error)

	// EfsRpcQueryProtectors operation.
	QueryProtectors(context.Context, *QueryProtectorsRequest, ...dcerpc.CallOption) (*QueryProtectorsResponse, error)

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

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Context structure represents PEXIMPORT_CONTEXT_HANDLE RPC structure.
type Context dcetypes.ContextHandle

func (o *Context) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Context) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Context) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Context) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ExportImportPipe type represents EFS_EXIM_PIPE RPC pipe.
type ExportImportPipe interface {
	Pipe() <-chan []byte
	Send(v []byte)
	Append(v []byte)
	Recv() ([]byte, bool)
	Close() error
}

// Pipe provides a channel-based interface to send data to the RPC server.
// Note: pipe must be closed after use to avoid resource leaks.
type xxx_ExportImportPipe struct {
	p    chan []byte
	recv [][]byte
}

// Newxxx_ExportImportPipe creates a new ExportImportPipe instance.
func NewExportImportPipe() ExportImportPipe {
	return &xxx_ExportImportPipe{
		p: make(chan []byte),
	}
}

func (o *xxx_ExportImportPipe) Pipe() <-chan []byte {
	return o.p
}

// Send sends data to the RPC server via the pipe.
func (o *xxx_ExportImportPipe) Send(v []byte) {
	if o != nil && o.p != nil {
		o.p <- v
	}
}

// Append appends data to the internal buffer of the pipe.
func (o *xxx_ExportImportPipe) Append(v []byte) {
	if o != nil {
		o.recv = append(o.recv, v)
	}
}

// Recv receives data from the internal buffer of the pipe.
func (o *xxx_ExportImportPipe) Recv() ([]byte, bool) {
	if o != nil {
		if len(o.recv) > 0 {
			v := o.recv[0]
			o.recv = o.recv[1:]
			return v, true
		}
	}
	return nil, false
}
func (o *xxx_ExportImportPipe) Close() error {
	if o != nil && o.p != nil {
		close(o.p)
	}
	return nil
}

// Blob structure represents EFS_RPC_BLOB RPC structure.
//
// The EFS_RPC_BLOB type is used to represent a generic binary large object (BLOB) (that
// is, an opaque data type).
type Blob struct {
	// cbData:  The length, in bytes, of the data object in the bData field.
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
	// bData:  The contents of the data object.<25>
	Data []byte `idl:"name:bData;size_is:(cbData)" json:"data"`
}

func (o *Blob) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.DataLength > uint32(266240) {
		return fmt.Errorf("DataLength is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Blob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_bData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_bData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Blob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_bData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_bData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_bData, _ptr_bData); err != nil {
		return err
	}
	return nil
}

// CompatibilityInfo structure represents EFS_COMPATIBILITY_INFO RPC structure.
//
// The EFS_COMPATIBILITY_INFO type is used to represent information about the compatibility
// restrictions of an encrypted file.
type CompatibilityInfo struct {
	// EfsVersion:  The EfsVersion associated with the EFSRPC Metadata. Valid values for
	// the EfsVersion field are described in sections 2.2.2.1, 2.2.2.2, and 2.2.2.3.<29>
	Version uint32 `idl:"name:EfsVersion" json:"version"`
}

func (o *CompatibilityInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CompatibilityInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	return nil
}
func (o *CompatibilityInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	return nil
}

// HashBlob structure represents EFS_HASH_BLOB RPC structure.
//
// The EFS_HASH_BLOB type is used to represent an X.509 certificate hash.
type HashBlob struct {
	// cbData:  The number of bytes in the bData buffer.
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
	// bData:  The SHA-1 hash of an X.509 certificate. For more information on SHA-1, see
	// [FIPS180-4].<22>
	Data []byte `idl:"name:bData;size_is:(cbData)" json:"data"`
}

func (o *HashBlob) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.DataLength > uint32(100) {
		return fmt.Errorf("DataLength is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *HashBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_bData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_bData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *HashBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_bData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_bData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_bData, _ptr_bData); err != nil {
		return err
	}
	return nil
}

// EncryptionCertificateHash structure represents ENCRYPTION_CERTIFICATE_HASH RPC structure.
//
// The ENCRYPTION_CERTIFICATE_HASH type is used to represent a single certificate hash.
// For more information on certificates, see [X509].
type EncryptionCertificateHash struct {
	// cbTotalLength:  The length, in bytes, of the structure.
	TotalLength uint32 `idl:"name:cbTotalLength" json:"total_length"`
	// UserSid:  The SID of the user who owns the certificate. This is intended only as
	// a hint. It MAY be set to zero if no such hint is available. The structure of an RPC
	// SID is specified in [MS-DTYP], section 2.4.2.3.
	UserSID *dtyp.SID `idl:"name:UserSid" json:"user_sid"`
	// Hash:  A pointer to an EFS_HASH_BLOB (2.2.7) structure.
	Hash *HashBlob `idl:"name:Hash" json:"hash"`
	// lpDisplayInformation:  A string that contains the subject or principal name of the
	// account the certification is assigned to. The subject name and the principal name
	// can be the same. This is only intended as a hint for display purposes, and is implementation-dependent.
	// This field MAY be set to NULL if no such information is available.
	DisplayInformation string `idl:"name:lpDisplayInformation;string" json:"display_information"`
}

func (o *EncryptionCertificateHash) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionCertificateHash) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalLength); err != nil {
		return err
	}
	if o.UserSID != nil {
		_ptr_UserSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UserSID != nil {
				if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UserSID, _ptr_UserSid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Hash != nil {
		_ptr_Hash := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Hash != nil {
				if err := o.Hash.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&HashBlob{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Hash, _ptr_Hash); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DisplayInformation != "" {
		_ptr_lpDisplayInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DisplayInformation); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DisplayInformation, _ptr_lpDisplayInformation); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptionCertificateHash) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalLength); err != nil {
		return err
	}
	_ptr_UserSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UserSID == nil {
			o.UserSID = &dtyp.SID{}
		}
		if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UserSid := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.UserSID, _s_UserSid, _ptr_UserSid); err != nil {
		return err
	}
	_ptr_Hash := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Hash == nil {
			o.Hash = &HashBlob{}
		}
		if err := o.Hash.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Hash := func(ptr interface{}) { o.Hash = *ptr.(**HashBlob) }
	if err := w.ReadPointer(&o.Hash, _s_Hash, _ptr_Hash); err != nil {
		return err
	}
	_ptr_lpDisplayInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DisplayInformation); err != nil {
			return err
		}
		return nil
	})
	_s_lpDisplayInformation := func(ptr interface{}) { o.DisplayInformation = *ptr.(*string) }
	if err := w.ReadPointer(&o.DisplayInformation, _s_lpDisplayInformation, _ptr_lpDisplayInformation); err != nil {
		return err
	}
	return nil
}

// EncryptionCertificateHashList structure represents ENCRYPTION_CERTIFICATE_HASH_LIST RPC structure.
//
// The ENCRYPTION_CERTIFICATE_HASH_LIST type is used to represent a set of certificate
// hashes.
type EncryptionCertificateHashList struct {
	// nCert_Hash:  The number of certificate hashes in the list.
	CertHash uint32 `idl:"name:nCert_Hash" json:"cert_hash"`
	// Users:  A pointer to an array of pointers to ENCRYPTION_CERTIFICATE_HASH (2.2.10)
	// structures. This array is of size nCert_Hash.<24>
	Users []*EncryptionCertificateHash `idl:"name:Users;size_is:(nCert_Hash, )" json:"users"`
}

func (o *EncryptionCertificateHashList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Users != nil && o.CertHash == 0 {
		o.CertHash = uint32(len(o.Users))
	}
	if o.CertHash > uint32(500) {
		return fmt.Errorf("CertHash is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionCertificateHashList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.CertHash); err != nil {
		return err
	}
	if o.Users != nil || o.CertHash > 0 {
		_ptr_Users := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CertHash)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Users {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Users[i1] != nil {
					_ptr_Users := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Users[i1] != nil {
							if err := o.Users[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&EncryptionCertificateHash{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Users[i1], _ptr_Users); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Users); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Users, _ptr_Users); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptionCertificateHashList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.CertHash); err != nil {
		return err
	}
	_ptr_Users := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CertHash > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CertHash)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Users", sizeInfo[0])
		}
		o.Users = make([]*EncryptionCertificateHash, sizeInfo[0])
		for i1 := range o.Users {
			i1 := i1
			_ptr_Users := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Users[i1] == nil {
					o.Users[i1] = &EncryptionCertificateHash{}
				}
				if err := o.Users[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Users := func(ptr interface{}) { o.Users[i1] = *ptr.(**EncryptionCertificateHash) }
			if err := w.ReadPointer(&o.Users[i1], _s_Users, _ptr_Users); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Users := func(ptr interface{}) { o.Users = *ptr.(*[]*EncryptionCertificateHash) }
	if err := w.ReadPointer(&o.Users, _s_Users, _ptr_Users); err != nil {
		return err
	}
	return nil
}

// CertificateBlob structure represents EFS_CERTIFICATE_BLOB RPC structure.
//
// The EFS_CERTIFICATE_BLOB type is used to represent the encoded contents of an X.509
// certificate.
type CertificateBlob struct {
	// dwCertEncodingType:  The certificate encoding type. This MUST be set to one of the
	// following values. If set to any other value, the certificate is considered invalid
	// and behavior is undefined.
	//
	//	+------------+----------------------------------------+
	//	|            |                                        |
	//	|   VALUE    |                MEANING                 |
	//	|            |                                        |
	//	+------------+----------------------------------------+
	//	+------------+----------------------------------------+
	//	| 0x00000001 | Certificate uses X.509 ASN.1 encoding. |
	//	+------------+----------------------------------------+
	//	| 0x00000002 | Certificate uses X.509 NDR encoding.   |
	//	+------------+----------------------------------------+
	CertEncodingType uint32 `idl:"name:dwCertEncodingType" json:"cert_encoding_type"`
	// cbData:  The number of bytes in the bData buffer.
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
	// bData:  An encoded X.509 certificate. Its format is specified by the dwCertEncodingType
	// member. For more information on ASN encoding, see [X690]. NDR encoding is specified
	// in [C706].<21>
	Data []byte `idl:"name:bData;size_is:(cbData)" json:"data"`
}

func (o *CertificateBlob) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.DataLength > uint32(32768) {
		return fmt.Errorf("DataLength is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CertificateBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.CertEncodingType); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_bData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_bData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertificateBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.CertEncodingType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_bData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_bData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_bData, _ptr_bData); err != nil {
		return err
	}
	return nil
}

// EncryptionCertificate structure represents ENCRYPTION_CERTIFICATE RPC structure.
type EncryptionCertificate struct {
	TotalLength uint32           `idl:"name:cbTotalLength" json:"total_length"`
	UserSID     *dtyp.SID        `idl:"name:UserSid" json:"user_sid"`
	CertBlob    *CertificateBlob `idl:"name:CertBlob" json:"cert_blob"`
}

func (o *EncryptionCertificate) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionCertificate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalLength); err != nil {
		return err
	}
	if o.UserSID != nil {
		_ptr_UserSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UserSID != nil {
				if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UserSID, _ptr_UserSid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.CertBlob != nil {
		_ptr_CertBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CertBlob != nil {
				if err := o.CertBlob.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CertificateBlob{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertBlob, _ptr_CertBlob); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptionCertificate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalLength); err != nil {
		return err
	}
	_ptr_UserSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UserSID == nil {
			o.UserSID = &dtyp.SID{}
		}
		if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UserSid := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.UserSID, _s_UserSid, _ptr_UserSid); err != nil {
		return err
	}
	_ptr_CertBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CertBlob == nil {
			o.CertBlob = &CertificateBlob{}
		}
		if err := o.CertBlob.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_CertBlob := func(ptr interface{}) { o.CertBlob = *ptr.(**CertificateBlob) }
	if err := w.ReadPointer(&o.CertBlob, _s_CertBlob, _ptr_CertBlob); err != nil {
		return err
	}
	return nil
}

// EncryptionCertificateList structure represents ENCRYPTION_CERTIFICATE_LIST RPC structure.
type EncryptionCertificateList struct {
	UsersCount uint32                   `idl:"name:nUsers" json:"users_count"`
	Users      []*EncryptionCertificate `idl:"name:Users;size_is:(nUsers, )" json:"users"`
}

func (o *EncryptionCertificateList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Users != nil && o.UsersCount == 0 {
		o.UsersCount = uint32(len(o.Users))
	}
	if o.UsersCount > uint32(500) {
		return fmt.Errorf("UsersCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionCertificateList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.UsersCount); err != nil {
		return err
	}
	if o.Users != nil || o.UsersCount > 0 {
		_ptr_Users := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.UsersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Users {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Users[i1] != nil {
					_ptr_Users := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Users[i1] != nil {
							if err := o.Users[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&EncryptionCertificate{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Users[i1], _ptr_Users); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Users); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Users, _ptr_Users); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptionCertificateList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.UsersCount); err != nil {
		return err
	}
	_ptr_Users := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.UsersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.UsersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Users", sizeInfo[0])
		}
		o.Users = make([]*EncryptionCertificate, sizeInfo[0])
		for i1 := range o.Users {
			i1 := i1
			_ptr_Users := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Users[i1] == nil {
					o.Users[i1] = &EncryptionCertificate{}
				}
				if err := o.Users[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Users := func(ptr interface{}) { o.Users[i1] = *ptr.(**EncryptionCertificate) }
			if err := w.ReadPointer(&o.Users[i1], _s_Users, _ptr_Users); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Users := func(ptr interface{}) { o.Users = *ptr.(*[]*EncryptionCertificate) }
	if err := w.ReadPointer(&o.Users, _s_Users, _ptr_Users); err != nil {
		return err
	}
	return nil
}

// EncryptedFileMetadataSignature structure represents ENCRYPTED_FILE_METADATA_SIGNATURE RPC structure.
//
// The ENCRYPTED_FILE_METADATA_SIGNATURE structure is used by the client to prove to
// the server that it possesses a private key that is authorized to decrypt a given
// object.
type EncryptedFileMetadataSignature struct {
	// dwEfsAccessType:  The operation being performed. It MUST be set to one of the following
	// values.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EFS_METADATA_ADD_USER 0x00000001     | One or more additional user certificates are being granted access to the object. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EFS_METADATA_REMOVE_USER 0x00000002  | One or more user certificates are having their access to the object revoked.     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EFS_METADATA_REPLACE_USER 0x00000004 | One or more user certificates with access to the object are being replaced.      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EFS_METADATA_GENERAL_OP 0x00000008   | A change is being made to the metadata that is not fully described by exactly    |
	//	|                                      | one of the previous options.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	EFSAccessType uint32 `idl:"name:dwEfsAccessType" json:"efs_access_type"`
	// CertificatesAdded:   The X.509 certificates whose corresponding private keys are
	// to be granted or denied the ability to decrypt the object.
	CertificatesAdded *EncryptionCertificateHashList `idl:"name:CertificatesAdded" json:"certificates_added"`
	// EncryptionCertificate:   The X.509 certificates whose corresponding private key the
	// caller claims to possess.
	EncryptionCertificate *EncryptionCertificate `idl:"name:EncryptionCertificate" json:"encryption_certificate"`
	// EfsStreamSignature:  The signature obtained by signing the SHA-1 hash of the new
	// EFSRPC Metadata with the private RSA key corresponding to EncryptionCertificate.
	StreamSignature *Blob `idl:"name:EfsStreamSignature" json:"stream_signature"`
}

func (o *EncryptedFileMetadataSignature) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptedFileMetadataSignature) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EFSAccessType); err != nil {
		return err
	}
	if o.CertificatesAdded != nil {
		_ptr_CertificatesAdded := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CertificatesAdded != nil {
				if err := o.CertificatesAdded.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EncryptionCertificateHashList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertificatesAdded, _ptr_CertificatesAdded); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EncryptionCertificate != nil {
		_ptr_EncryptionCertificate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.EncryptionCertificate != nil {
				if err := o.EncryptionCertificate.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&EncryptionCertificate{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EncryptionCertificate, _ptr_EncryptionCertificate); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.StreamSignature != nil {
		_ptr_EfsStreamSignature := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.StreamSignature != nil {
				if err := o.StreamSignature.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.StreamSignature, _ptr_EfsStreamSignature); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedFileMetadataSignature) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EFSAccessType); err != nil {
		return err
	}
	_ptr_CertificatesAdded := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CertificatesAdded == nil {
			o.CertificatesAdded = &EncryptionCertificateHashList{}
		}
		if err := o.CertificatesAdded.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_CertificatesAdded := func(ptr interface{}) { o.CertificatesAdded = *ptr.(**EncryptionCertificateHashList) }
	if err := w.ReadPointer(&o.CertificatesAdded, _s_CertificatesAdded, _ptr_CertificatesAdded); err != nil {
		return err
	}
	_ptr_EncryptionCertificate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.EncryptionCertificate == nil {
			o.EncryptionCertificate = &EncryptionCertificate{}
		}
		if err := o.EncryptionCertificate.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_EncryptionCertificate := func(ptr interface{}) { o.EncryptionCertificate = *ptr.(**EncryptionCertificate) }
	if err := w.ReadPointer(&o.EncryptionCertificate, _s_EncryptionCertificate, _ptr_EncryptionCertificate); err != nil {
		return err
	}
	_ptr_EfsStreamSignature := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.StreamSignature == nil {
			o.StreamSignature = &Blob{}
		}
		if err := o.StreamSignature.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_EfsStreamSignature := func(ptr interface{}) { o.StreamSignature = *ptr.(**Blob) }
	if err := w.ReadPointer(&o.StreamSignature, _s_EfsStreamSignature, _ptr_EfsStreamSignature); err != nil {
		return err
	}
	return nil
}

// KeyInfo structure represents EFS_KEY_INFO RPC structure.
//
// The EFS_KEY_INFO type is used to represent information about a key of a symmetric
// cryptosystem.
type KeyInfo struct {
	// dwVersion:  The version of this data structure. It MUST be equal to 0x00000001.
	Version uint32 `idl:"name:dwVersion" json:"version"`
	// Entropy:  The actual number of bits of entropy or true randomness in the key. This
	// value, divided by 8, MUST be less than or equal to the value of the KeyLength member.
	Entropy uint32 `idl:"name:Entropy" json:"entropy"`
	// Algorithm:  The cryptographic algorithm with which the key is intended to be used.
	Algorithm uint32 `idl:"name:Algorithm" json:"algorithm"`
	// KeyLength:  The total length, in bytes, of the key. This value, multiplied by 8,
	// MUST be greater than or equal to the value of the Entropy member. Valid combinations
	// of Entropy, Algorithm, and KeyLength are specified in section 2.2.13.
	KeyLength uint32 `idl:"name:KeyLength" json:"key_length"`
}

func (o *KeyInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *KeyInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Entropy); err != nil {
		return err
	}
	if err := w.WriteData(o.Algorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyLength); err != nil {
		return err
	}
	return nil
}
func (o *KeyInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entropy); err != nil {
		return err
	}
	if err := w.ReadData(&o.Algorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyLength); err != nil {
		return err
	}
	return nil
}

// DecryptionStatusInfo structure represents EFS_DECRYPTION_STATUS_INFO RPC structure.
//
// The EFS_DECRYPTION_STATUS_INFO type is used to represent the predicted outcome if
// an attempt were made to read the plaintext of an encrypted object.
type DecryptionStatusInfo struct {
	// dwDecryptionError:  The error code returned if decryption were attempted. If the
	// operation were to succeed, this value MUST be zero. Otherwise it MUST be set to a
	// nonzero value.
	DecryptionError uint32 `idl:"name:dwDecryptionError" json:"decryption_error"`
	// dwHashOffset:  The offset of the appended certificate hash in bytes from the start
	// of this structure.
	HashOffset uint32 `idl:"name:dwHashOffset" json:"hash_offset"`
	// cbHash:  The length in bytes of the appended certificate hash.
	//
	// If dwDecryptionError is nonzero, the preceding fields are followed by the hash of
	// a certificate whose corresponding private key is required for the decryption to succeed.
	HashLength uint32 `idl:"name:cbHash" json:"hash_length"`
}

func (o *DecryptionStatusInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DecryptionStatusInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DecryptionError); err != nil {
		return err
	}
	if err := w.WriteData(o.HashOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.HashLength); err != nil {
		return err
	}
	return nil
}
func (o *DecryptionStatusInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DecryptionError); err != nil {
		return err
	}
	if err := w.ReadData(&o.HashOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.HashLength); err != nil {
		return err
	}
	return nil
}

// EncryptionStatusInfo structure represents EFS_ENCRYPTION_STATUS_INFO RPC structure.
//
// The EFS_ENCRYPTION_STATUS_INFO structure is used to represent the predicted outcome
// if an attempt were made to convert an unencrypted object to an encrypted state.
type EncryptionStatusInfo struct {
	// bHasCurrentKey:  A Boolean value signifying whether an appropriate key was found
	// that could be used for encryption.
	HasCurrentKey bool `idl:"name:bHasCurrentKey" json:"has_current_key"`
	// dwEncryptionError:  The error code returned if encryption were attempted. If the
	// operation were to succeed, this value MUST be zero. Otherwise, it MUST be set to
	// a nonzero value.
	EncryptionError uint32 `idl:"name:dwEncryptionError" json:"encryption_error"`
}

func (o *EncryptionStatusInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionStatusInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.HasCurrentKey {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EncryptionError); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionStatusInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bHasCurrentKey int32
	if err := w.ReadData(&_bHasCurrentKey); err != nil {
		return err
	}
	o.HasCurrentKey = _bHasCurrentKey != 0
	if err := w.ReadData(&o.EncryptionError); err != nil {
		return err
	}
	return nil
}

// EncryptionProtector structure represents ENCRYPTION_PROTECTOR RPC structure.
type EncryptionProtector struct {
	TotalLength         uint32    `idl:"name:cbTotalLength" json:"total_length"`
	UserSID             *dtyp.SID `idl:"name:UserSid" json:"user_sid"`
	ProtectorDescriptor string    `idl:"name:lpProtectorDescriptor;string" json:"protector_descriptor"`
}

func (o *EncryptionProtector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionProtector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalLength); err != nil {
		return err
	}
	if o.UserSID != nil {
		_ptr_UserSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UserSID != nil {
				if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UserSID, _ptr_UserSid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProtectorDescriptor != "" {
		_ptr_lpProtectorDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProtectorDescriptor); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProtectorDescriptor, _ptr_lpProtectorDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptionProtector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalLength); err != nil {
		return err
	}
	_ptr_UserSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UserSID == nil {
			o.UserSID = &dtyp.SID{}
		}
		if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UserSid := func(ptr interface{}) { o.UserSID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.UserSID, _s_UserSid, _ptr_UserSid); err != nil {
		return err
	}
	_ptr_lpProtectorDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProtectorDescriptor); err != nil {
			return err
		}
		return nil
	})
	_s_lpProtectorDescriptor := func(ptr interface{}) { o.ProtectorDescriptor = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProtectorDescriptor, _s_lpProtectorDescriptor, _ptr_lpProtectorDescriptor); err != nil {
		return err
	}
	return nil
}

// EncryptionProtectorList structure represents ENCRYPTION_PROTECTOR_LIST RPC structure.
type EncryptionProtectorList struct {
	ProtectorsCount uint32                 `idl:"name:nProtectors" json:"protectors_count"`
	Protectors      []*EncryptionProtector `idl:"name:pProtectors;size_is:(nProtectors)" json:"protectors"`
}

func (o *EncryptionProtectorList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Protectors != nil && o.ProtectorsCount == 0 {
		o.ProtectorsCount = uint32(len(o.Protectors))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EncryptionProtectorList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ProtectorsCount); err != nil {
		return err
	}
	if o.Protectors != nil || o.ProtectorsCount > 0 {
		_ptr_pProtectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ProtectorsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Protectors {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Protectors[i1] != nil {
					_ptr_pProtectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Protectors[i1] != nil {
							if err := o.Protectors[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&EncryptionProtector{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Protectors[i1], _ptr_pProtectors); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Protectors); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Protectors, _ptr_pProtectors); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptionProtectorList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtectorsCount); err != nil {
		return err
	}
	_ptr_pProtectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ProtectorsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ProtectorsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Protectors", sizeInfo[0])
		}
		o.Protectors = make([]*EncryptionProtector, sizeInfo[0])
		for i1 := range o.Protectors {
			i1 := i1
			_ptr_pProtectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Protectors[i1] == nil {
					o.Protectors[i1] = &EncryptionProtector{}
				}
				if err := o.Protectors[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_pProtectors := func(ptr interface{}) { o.Protectors[i1] = *ptr.(**EncryptionProtector) }
			if err := w.ReadPointer(&o.Protectors[i1], _s_pProtectors, _ptr_pProtectors); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pProtectors := func(ptr interface{}) { o.Protectors = *ptr.(*[]*EncryptionProtector) }
	if err := w.ReadPointer(&o.Protectors, _s_pProtectors, _ptr_pProtectors); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultEfsrpcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultEfsrpcClient) OpenFileRaw(ctx context.Context, in *OpenFileRawRequest, opts ...dcerpc.CallOption) (*OpenFileRawResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenFileRawResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) ReadFileRaw(ctx context.Context, in *ReadFileRawRequest, opts ...dcerpc.CallOption) (*ReadFileRawResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReadFileRawResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) WriteFileRaw(ctx context.Context, in *WriteFileRawRequest, opts ...dcerpc.CallOption) (*WriteFileRawResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WriteFileRawResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) CloseRaw(ctx context.Context, in *CloseRawRequest, opts ...dcerpc.CallOption) (*CloseRawResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseRawResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) EncryptFileServer(ctx context.Context, in *EncryptFileServerRequest, opts ...dcerpc.CallOption) (*EncryptFileServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EncryptFileServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) DecryptFileServer(ctx context.Context, in *DecryptFileServerRequest, opts ...dcerpc.CallOption) (*DecryptFileServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DecryptFileServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) QueryUsersOnFile(ctx context.Context, in *QueryUsersOnFileRequest, opts ...dcerpc.CallOption) (*QueryUsersOnFileResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryUsersOnFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) QueryRecoveryAgents(ctx context.Context, in *QueryRecoveryAgentsRequest, opts ...dcerpc.CallOption) (*QueryRecoveryAgentsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryRecoveryAgentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) RemoveUsersFromFile(ctx context.Context, in *RemoveUsersFromFileRequest, opts ...dcerpc.CallOption) (*RemoveUsersFromFileResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveUsersFromFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) AddUsersToFile(ctx context.Context, in *AddUsersToFileRequest, opts ...dcerpc.CallOption) (*AddUsersToFileResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddUsersToFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) NotSupported(ctx context.Context, in *NotSupportedRequest, opts ...dcerpc.CallOption) (*NotSupportedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NotSupportedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) FileKeyInfo(ctx context.Context, in *FileKeyInfoRequest, opts ...dcerpc.CallOption) (*FileKeyInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FileKeyInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) DuplicateEncryptionInfoFile(ctx context.Context, in *DuplicateEncryptionInfoFileRequest, opts ...dcerpc.CallOption) (*DuplicateEncryptionInfoFileResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DuplicateEncryptionInfoFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) AddUsersToFileEx(ctx context.Context, in *AddUsersToFileExRequest, opts ...dcerpc.CallOption) (*AddUsersToFileExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddUsersToFileExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) FileKeyInfoEx(ctx context.Context, in *FileKeyInfoExRequest, opts ...dcerpc.CallOption) (*FileKeyInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FileKeyInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) GetEncryptedFileMetadata(ctx context.Context, in *GetEncryptedFileMetadataRequest, opts ...dcerpc.CallOption) (*GetEncryptedFileMetadataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEncryptedFileMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) SetEncryptedFileMetadata(ctx context.Context, in *SetEncryptedFileMetadataRequest, opts ...dcerpc.CallOption) (*SetEncryptedFileMetadataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetEncryptedFileMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) FlushEFSCache(ctx context.Context, in *FlushEFSCacheRequest, opts ...dcerpc.CallOption) (*FlushEFSCacheResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FlushEFSCacheResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) EncryptFileExServer(ctx context.Context, in *EncryptFileExServerRequest, opts ...dcerpc.CallOption) (*EncryptFileExServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EncryptFileExServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) QueryProtectors(ctx context.Context, in *QueryProtectorsRequest, opts ...dcerpc.CallOption) (*QueryProtectorsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryProtectorsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEfsrpcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEfsrpcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewEfsrpcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EfsrpcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EfsrpcSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultEfsrpcClient{cc: cc}, nil
}

// xxx_OpenFileRawOperation structure represents the EfsRpcOpenFileRaw operation
type xxx_OpenFileRawOperation struct {
	Context  *Context `idl:"name:hContext" json:"context"`
	FileName string   `idl:"name:FileName;string" json:"file_name"`
	Flags    int32    `idl:"name:Flags" json:"flags"`
	Return   int32    `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenFileRawOperation) OpNum() int { return 0 }

func (o *xxx_OpenFileRawOperation) OpName() string { return "/efsrpc/v1/EfsRpcOpenFileRaw" }

func (o *xxx_OpenFileRawOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenFileRawOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// Flags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenFileRawOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// Flags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenFileRawOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenFileRawOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// hContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenFileRawOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// hContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenFileRawRequest structure represents the EfsRpcOpenFileRaw operation request
type OpenFileRawRequest struct {
	// FileName: An EFSRPC identifier, as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// Flags: This MUST be set to some combination of the following values. All servers
	// and clients MUST support the CREATE_FOR_IMPORT flag. Servers that implement a hierarchical
	// encrypted store, such as the NTFS file system, SHOULD also support the CREATE_FOR_DIR
	// flag. Servers SHOULD support the OVERWRITE_HIDDEN flag, and MAY interpret it in implementation-specific
	// ways. A client MUST ensure that all the flags it does not support are set to zero.
	// A server MUST ignore all flags it does not support. Flag values are specified in
	// the following table.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| CREATE_FOR_IMPORT 0x00000001          | Open the object for writing (that is, restore). If this flag is not set, open    |
	//	|                                       | the object for reading (that is, backup).                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| CREATE_FOR_DIR 0x00000002             | This flag is only intended for use in conjunction with the CREATE_FOR_IMPORT     |
	//	|                                       | flag. It indicates that the object being restored is a container for other       |
	//	|                                       | objects.<43>                                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| OVERWRITE_HIDDEN 0x00000004           | This flag is only intended for use in conjunction with the CREATE_FOR_IMPORT     |
	//	|                                       | flag. This flag indicates a request from the client for the server to overwrite  |
	//	|                                       | an existing object even if the existing object is "hidden". The meaning of       |
	//	|                                       | "hidden" is specific to the implementation of the data store, and this meaning   |
	//	|                                       | does not affect protocol behavior.                                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| EFS_DROP_ALTERNATE_STREAMS 0x00000010 | This flag indicates that content from any alternate data streams, if present and |
	//	|                                       | implemented by the storage system, will be ignored.                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:Flags" json:"flags"`
}

func (o *OpenFileRawRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenFileRawOperation) *xxx_OpenFileRawOperation {
	if op == nil {
		op = &xxx_OpenFileRawOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	op.Flags = o.Flags
	return op
}

func (o *OpenFileRawRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenFileRawOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.Flags = op.Flags
}
func (o *OpenFileRawRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenFileRawRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenFileRawOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenFileRawResponse structure represents the EfsRpcOpenFileRaw operation response
type OpenFileRawResponse struct {
	// hContext: An implementation-specific context handle that is used in subsequent calls
	// by the client to the EfsRpcReadFileRaw method, EfsRpcWriteFileRaw method, or EfsRpcCloseRaw
	// method.
	Context *Context `idl:"name:hContext" json:"context"`
	// Return: The EfsRpcOpenFileRaw return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenFileRawResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenFileRawOperation) *xxx_OpenFileRawOperation {
	if op == nil {
		op = &xxx_OpenFileRawOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *OpenFileRawResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenFileRawOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
}
func (o *OpenFileRawResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenFileRawResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenFileRawOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReadFileRawOperation structure represents the EfsRpcReadFileRaw operation
type xxx_ReadFileRawOperation struct {
	Context *Context         `idl:"name:hContext" json:"context"`
	OutPipe ExportImportPipe `idl:"name:EfsOutPipe" json:"out_pipe"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_ReadFileRawOperation) OpNum() int { return 1 }

func (o *xxx_ReadFileRawOperation) OpName() string { return "/efsrpc/v1/EfsRpcReadFileRaw" }

func (o *xxx_ReadFileRawOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadFileRawOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hContext {in} (1:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ReadFileRawOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hContext {in} (1:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadFileRawOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadFileRawOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EfsOutPipe {out} (1:{pointer=ref}*(1))(2:{alias=EFS_EXIM_PIPE}(pipe)(uchar))
	{
		// marshal pipe EfsOutPipe
		if o.OutPipe != nil {
			for _chunk := range o.OutPipe.Pipe() {
				dimSize1 := uint64(len(_chunk))
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range _chunk {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_chunk[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_chunk); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
			}
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadFileRawOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EfsOutPipe {out} (1:{pointer=ref}*(1))(2:{alias=EFS_EXIM_PIPE}(pipe)(uchar))
	{
		if o.OutPipe == nil {
			o.OutPipe = NewExportImportPipe()
		}
		for {
			var _chunk []byte
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _chunk", sizeInfo[0])
			}
			_chunk = make([]byte, sizeInfo[0])
			for i1 := range _chunk {
				i1 := i1
				if err := w.ReadData(&_chunk[i1]); err != nil {
					return err
				}
			}
			if len(_chunk) == 0 /* end */ {
				break
			}
			o.OutPipe.Append(_chunk)
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReadFileRawRequest structure represents the EfsRpcReadFileRaw operation request
type ReadFileRawRequest struct {
	// XXX: EfsOutPipe is an implicit input depedency for output parameters
	OutPipe ExportImportPipe `idl:"name:EfsOutPipe" json:"out_pipe"`

	// hContext: A context handle returned by the EfsRpcOpenFileRaw method, which MUST have
	// been called without the CREATE_FOR_IMPORT flag.
	Context *Context `idl:"name:hContext" json:"context"`
}

func (o *ReadFileRawRequest) xxx_ToOp(ctx context.Context, op *xxx_ReadFileRawOperation) *xxx_ReadFileRawOperation {
	if op == nil {
		op = &xxx_ReadFileRawOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.OutPipe == nil {
		op.OutPipe = o.OutPipe
	}

	op.Context = o.Context
	return op
}

func (o *ReadFileRawRequest) xxx_FromOp(ctx context.Context, op *xxx_ReadFileRawOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.OutPipe = op.OutPipe

	o.Context = op.Context
}
func (o *ReadFileRawRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReadFileRawRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadFileRawOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReadFileRawResponse structure represents the EfsRpcReadFileRaw operation response
type ReadFileRawResponse struct {
	// EfsOutPipe: A pipe structure. The push procedure of this pipe will be called with
	// the marshaled data. The structure of this marshaled data is specified in section
	// 2.2.3.
	OutPipe ExportImportPipe `idl:"name:EfsOutPipe" json:"out_pipe"`
	// Return: The EfsRpcReadFileRaw return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReadFileRawResponse) xxx_ToOp(ctx context.Context, op *xxx_ReadFileRawOperation) *xxx_ReadFileRawOperation {
	if op == nil {
		op = &xxx_ReadFileRawOperation{}
	}
	if o == nil {
		return op
	}
	op.OutPipe = o.OutPipe
	op.Return = o.Return
	return op
}

func (o *ReadFileRawResponse) xxx_FromOp(ctx context.Context, op *xxx_ReadFileRawOperation) {
	if o == nil {
		return
	}
	o.OutPipe = op.OutPipe
	o.Return = op.Return
}
func (o *ReadFileRawResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReadFileRawResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadFileRawOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WriteFileRawOperation structure represents the EfsRpcWriteFileRaw operation
type xxx_WriteFileRawOperation struct {
	Context *Context         `idl:"name:hContext" json:"context"`
	InPipe  ExportImportPipe `idl:"name:EfsInPipe" json:"in_pipe"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_WriteFileRawOperation) OpNum() int { return 2 }

func (o *xxx_WriteFileRawOperation) OpName() string { return "/efsrpc/v1/EfsRpcWriteFileRaw" }

func (o *xxx_WriteFileRawOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteFileRawOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hContext {in} (1:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EfsInPipe {in} (1:{pointer=ref}*(1))(2:{alias=EFS_EXIM_PIPE}(pipe)(uchar))
	{
		// marshal pipe EfsInPipe
		if o.InPipe != nil {
			for _chunk := range o.InPipe.Pipe() {
				dimSize1 := uint64(len(_chunk))
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range _chunk {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_chunk[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_chunk); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
			}
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteFileRawOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hContext {in} (1:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EfsInPipe {in} (1:{pointer=ref}*(1))(2:{alias=EFS_EXIM_PIPE}(pipe)(uchar))
	{
		if o.InPipe == nil {
			o.InPipe = NewExportImportPipe()
		}
		for {
			var _chunk []byte
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _chunk", sizeInfo[0])
			}
			_chunk = make([]byte, sizeInfo[0])
			for i1 := range _chunk {
				i1 := i1
				if err := w.ReadData(&_chunk[i1]); err != nil {
					return err
				}
			}
			if len(_chunk) == 0 /* end */ {
				break
			}
			o.InPipe.Append(_chunk)
		}
	}
	return nil
}

func (o *xxx_WriteFileRawOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteFileRawOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteFileRawOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WriteFileRawRequest structure represents the EfsRpcWriteFileRaw operation request
type WriteFileRawRequest struct {
	// hContext: A context handle returned by the EfsRpcOpenFileRaw method, which MUST have
	// been called with the CREATE_FOR_IMPORT flag.
	Context *Context `idl:"name:hContext" json:"context"`
	// EfsInPipe: A pipe structure. The pull procedure of this pipe is expected to provide
	// the marshaled data. The structure of this marshaled data is specified in section
	// 2.2.3.
	InPipe ExportImportPipe `idl:"name:EfsInPipe" json:"in_pipe"`
}

func (o *WriteFileRawRequest) xxx_ToOp(ctx context.Context, op *xxx_WriteFileRawOperation) *xxx_WriteFileRawOperation {
	if op == nil {
		op = &xxx_WriteFileRawOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.InPipe = o.InPipe
	return op
}

func (o *WriteFileRawRequest) xxx_FromOp(ctx context.Context, op *xxx_WriteFileRawOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.InPipe = op.InPipe
}
func (o *WriteFileRawRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WriteFileRawRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteFileRawOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WriteFileRawResponse structure represents the EfsRpcWriteFileRaw operation response
type WriteFileRawResponse struct {
	// Return: The EfsRpcWriteFileRaw return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WriteFileRawResponse) xxx_ToOp(ctx context.Context, op *xxx_WriteFileRawOperation) *xxx_WriteFileRawOperation {
	if op == nil {
		op = &xxx_WriteFileRawOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *WriteFileRawResponse) xxx_FromOp(ctx context.Context, op *xxx_WriteFileRawOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *WriteFileRawResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WriteFileRawResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteFileRawOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseRawOperation structure represents the EfsRpcCloseRaw operation
type xxx_CloseRawOperation struct {
	Context *Context `idl:"name:hContext" json:"context"`
}

func (o *xxx_CloseRawOperation) OpNum() int { return 3 }

func (o *xxx_CloseRawOperation) OpName() string { return "/efsrpc/v1/EfsRpcCloseRaw" }

func (o *xxx_CloseRawOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseRawOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseRawOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseRawOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseRawOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// hContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseRawOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// hContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PEXIMPORT_CONTEXT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CloseRawRequest structure represents the EfsRpcCloseRaw operation request
type CloseRawRequest struct {
	Context *Context `idl:"name:hContext" json:"context"`
}

func (o *CloseRawRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseRawOperation) *xxx_CloseRawOperation {
	if op == nil {
		op = &xxx_CloseRawOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CloseRawRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseRawOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *CloseRawRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseRawRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseRawOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseRawResponse structure represents the EfsRpcCloseRaw operation response
type CloseRawResponse struct {
	Context *Context `idl:"name:hContext" json:"context"`
}

func (o *CloseRawResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseRawOperation) *xxx_CloseRawOperation {
	if op == nil {
		op = &xxx_CloseRawOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CloseRawResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseRawOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *CloseRawResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseRawResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseRawOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EncryptFileServerOperation structure represents the EfsRpcEncryptFileSrv operation
type xxx_EncryptFileServerOperation struct {
	FileName string `idl:"name:FileName;string" json:"file_name"`
	Return   int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_EncryptFileServerOperation) OpNum() int { return 4 }

func (o *xxx_EncryptFileServerOperation) OpName() string { return "/efsrpc/v1/EfsRpcEncryptFileSrv" }

func (o *xxx_EncryptFileServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EncryptFileServerRequest structure represents the EfsRpcEncryptFileSrv operation request
type EncryptFileServerRequest struct {
	FileName string `idl:"name:FileName;string" json:"file_name"`
}

func (o *EncryptFileServerRequest) xxx_ToOp(ctx context.Context, op *xxx_EncryptFileServerOperation) *xxx_EncryptFileServerOperation {
	if op == nil {
		op = &xxx_EncryptFileServerOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	return op
}

func (o *EncryptFileServerRequest) xxx_FromOp(ctx context.Context, op *xxx_EncryptFileServerOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
}
func (o *EncryptFileServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EncryptFileServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EncryptFileServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EncryptFileServerResponse structure represents the EfsRpcEncryptFileSrv operation response
type EncryptFileServerResponse struct {
	// Return: The EfsRpcEncryptFileSrv return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EncryptFileServerResponse) xxx_ToOp(ctx context.Context, op *xxx_EncryptFileServerOperation) *xxx_EncryptFileServerOperation {
	if op == nil {
		op = &xxx_EncryptFileServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EncryptFileServerResponse) xxx_FromOp(ctx context.Context, op *xxx_EncryptFileServerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EncryptFileServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EncryptFileServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EncryptFileServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DecryptFileServerOperation structure represents the EfsRpcDecryptFileSrv operation
type xxx_DecryptFileServerOperation struct {
	FileName string `idl:"name:FileName;string" json:"file_name"`
	OpenFlag uint32 `idl:"name:OpenFlag" json:"open_flag"`
	Return   int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_DecryptFileServerOperation) OpNum() int { return 5 }

func (o *xxx_DecryptFileServerOperation) OpName() string { return "/efsrpc/v1/EfsRpcDecryptFileSrv" }

func (o *xxx_DecryptFileServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecryptFileServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// OpenFlag {in} (1:(uint32))
	{
		if err := w.WriteData(o.OpenFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecryptFileServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// OpenFlag {in} (1:(uint32))
	{
		if err := w.ReadData(&o.OpenFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecryptFileServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecryptFileServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DecryptFileServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DecryptFileServerRequest structure represents the EfsRpcDecryptFileSrv operation request
type DecryptFileServerRequest struct {
	// FileName: An EFSRPC identifier as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// OpenFlag: This parameter is unused and MUST be ignored by the server. It MUST be
	// set to zero by the client.
	OpenFlag uint32 `idl:"name:OpenFlag" json:"open_flag"`
}

func (o *DecryptFileServerRequest) xxx_ToOp(ctx context.Context, op *xxx_DecryptFileServerOperation) *xxx_DecryptFileServerOperation {
	if op == nil {
		op = &xxx_DecryptFileServerOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	op.OpenFlag = o.OpenFlag
	return op
}

func (o *DecryptFileServerRequest) xxx_FromOp(ctx context.Context, op *xxx_DecryptFileServerOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.OpenFlag = op.OpenFlag
}
func (o *DecryptFileServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DecryptFileServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DecryptFileServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DecryptFileServerResponse structure represents the EfsRpcDecryptFileSrv operation response
type DecryptFileServerResponse struct {
	// Return: The EfsRpcDecryptFileSrv return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DecryptFileServerResponse) xxx_ToOp(ctx context.Context, op *xxx_DecryptFileServerOperation) *xxx_DecryptFileServerOperation {
	if op == nil {
		op = &xxx_DecryptFileServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *DecryptFileServerResponse) xxx_FromOp(ctx context.Context, op *xxx_DecryptFileServerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DecryptFileServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DecryptFileServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DecryptFileServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryUsersOnFileOperation structure represents the EfsRpcQueryUsersOnFile operation
type xxx_QueryUsersOnFileOperation struct {
	FileName string                         `idl:"name:FileName;string" json:"file_name"`
	Users    *EncryptionCertificateHashList `idl:"name:Users" json:"users"`
	Return   uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryUsersOnFileOperation) OpNum() int { return 6 }

func (o *xxx_QueryUsersOnFileOperation) OpName() string { return "/efsrpc/v1/EfsRpcQueryUsersOnFile" }

func (o *xxx_QueryUsersOnFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUsersOnFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUsersOnFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUsersOnFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryUsersOnFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Users {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ENCRYPTION_CERTIFICATE_HASH_LIST}(struct))
	{
		if o.Users != nil {
			_ptr_Users := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Users != nil {
					if err := o.Users.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptionCertificateHashList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Users, _ptr_Users); err != nil {
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

func (o *xxx_QueryUsersOnFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Users {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ENCRYPTION_CERTIFICATE_HASH_LIST}(struct))
	{
		_ptr_Users := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Users == nil {
				o.Users = &EncryptionCertificateHashList{}
			}
			if err := o.Users.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Users := func(ptr interface{}) { o.Users = *ptr.(**EncryptionCertificateHashList) }
		if err := w.ReadPointer(&o.Users, _s_Users, _ptr_Users); err != nil {
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

// QueryUsersOnFileRequest structure represents the EfsRpcQueryUsersOnFile operation request
type QueryUsersOnFileRequest struct {
	FileName string `idl:"name:FileName;string" json:"file_name"`
}

func (o *QueryUsersOnFileRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryUsersOnFileOperation) *xxx_QueryUsersOnFileOperation {
	if op == nil {
		op = &xxx_QueryUsersOnFileOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	return op
}

func (o *QueryUsersOnFileRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryUsersOnFileOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
}
func (o *QueryUsersOnFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryUsersOnFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryUsersOnFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryUsersOnFileResponse structure represents the EfsRpcQueryUsersOnFile operation response
type QueryUsersOnFileResponse struct {
	Users *EncryptionCertificateHashList `idl:"name:Users" json:"users"`
	// Return: The EfsRpcQueryUsersOnFile return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryUsersOnFileResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryUsersOnFileOperation) *xxx_QueryUsersOnFileOperation {
	if op == nil {
		op = &xxx_QueryUsersOnFileOperation{}
	}
	if o == nil {
		return op
	}
	op.Users = o.Users
	op.Return = o.Return
	return op
}

func (o *QueryUsersOnFileResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryUsersOnFileOperation) {
	if o == nil {
		return
	}
	o.Users = op.Users
	o.Return = op.Return
}
func (o *QueryUsersOnFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryUsersOnFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryUsersOnFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryRecoveryAgentsOperation structure represents the EfsRpcQueryRecoveryAgents operation
type xxx_QueryRecoveryAgentsOperation struct {
	FileName       string                         `idl:"name:FileName;string" json:"file_name"`
	RecoveryAgents *EncryptionCertificateHashList `idl:"name:RecoveryAgents" json:"recovery_agents"`
	Return         uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryRecoveryAgentsOperation) OpNum() int { return 7 }

func (o *xxx_QueryRecoveryAgentsOperation) OpName() string {
	return "/efsrpc/v1/EfsRpcQueryRecoveryAgents"
}

func (o *xxx_QueryRecoveryAgentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRecoveryAgentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRecoveryAgentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRecoveryAgentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRecoveryAgentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RecoveryAgents {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ENCRYPTION_CERTIFICATE_HASH_LIST}(struct))
	{
		if o.RecoveryAgents != nil {
			_ptr_RecoveryAgents := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecoveryAgents != nil {
					if err := o.RecoveryAgents.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptionCertificateHashList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecoveryAgents, _ptr_RecoveryAgents); err != nil {
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

func (o *xxx_QueryRecoveryAgentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RecoveryAgents {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ENCRYPTION_CERTIFICATE_HASH_LIST}(struct))
	{
		_ptr_RecoveryAgents := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecoveryAgents == nil {
				o.RecoveryAgents = &EncryptionCertificateHashList{}
			}
			if err := o.RecoveryAgents.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RecoveryAgents := func(ptr interface{}) { o.RecoveryAgents = *ptr.(**EncryptionCertificateHashList) }
		if err := w.ReadPointer(&o.RecoveryAgents, _s_RecoveryAgents, _ptr_RecoveryAgents); err != nil {
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

// QueryRecoveryAgentsRequest structure represents the EfsRpcQueryRecoveryAgents operation request
type QueryRecoveryAgentsRequest struct {
	// FileName: An EFSRPC identifier as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
}

func (o *QueryRecoveryAgentsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryRecoveryAgentsOperation) *xxx_QueryRecoveryAgentsOperation {
	if op == nil {
		op = &xxx_QueryRecoveryAgentsOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	return op
}

func (o *QueryRecoveryAgentsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryRecoveryAgentsOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
}
func (o *QueryRecoveryAgentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryRecoveryAgentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryRecoveryAgentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryRecoveryAgentsResponse structure represents the EfsRpcQueryRecoveryAgents operation response
type QueryRecoveryAgentsResponse struct {
	// RecoveryAgents: A list of certificate hashes, represented by an ENCRYPTION_CERTIFICATE_HASH_LIST
	// structure.
	RecoveryAgents *EncryptionCertificateHashList `idl:"name:RecoveryAgents" json:"recovery_agents"`
	// Return: The EfsRpcQueryRecoveryAgents return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryRecoveryAgentsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryRecoveryAgentsOperation) *xxx_QueryRecoveryAgentsOperation {
	if op == nil {
		op = &xxx_QueryRecoveryAgentsOperation{}
	}
	if o == nil {
		return op
	}
	op.RecoveryAgents = o.RecoveryAgents
	op.Return = o.Return
	return op
}

func (o *QueryRecoveryAgentsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryRecoveryAgentsOperation) {
	if o == nil {
		return
	}
	o.RecoveryAgents = op.RecoveryAgents
	o.Return = op.Return
}
func (o *QueryRecoveryAgentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryRecoveryAgentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryRecoveryAgentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveUsersFromFileOperation structure represents the EfsRpcRemoveUsersFromFile operation
type xxx_RemoveUsersFromFileOperation struct {
	FileName string                         `idl:"name:FileName;string" json:"file_name"`
	Users    *EncryptionCertificateHashList `idl:"name:Users" json:"users"`
	Return   uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveUsersFromFileOperation) OpNum() int { return 8 }

func (o *xxx_RemoveUsersFromFileOperation) OpName() string {
	return "/efsrpc/v1/EfsRpcRemoveUsersFromFile"
}

func (o *xxx_RemoveUsersFromFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveUsersFromFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// Users {in} (1:{pointer=ref}*(1))(2:{alias=ENCRYPTION_CERTIFICATE_HASH_LIST}(struct))
	{
		if o.Users != nil {
			if err := o.Users.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EncryptionCertificateHashList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveUsersFromFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// Users {in} (1:{pointer=ref}*(1))(2:{alias=ENCRYPTION_CERTIFICATE_HASH_LIST}(struct))
	{
		if o.Users == nil {
			o.Users = &EncryptionCertificateHashList{}
		}
		if err := o.Users.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveUsersFromFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveUsersFromFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveUsersFromFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveUsersFromFileRequest structure represents the EfsRpcRemoveUsersFromFile operation request
type RemoveUsersFromFileRequest struct {
	// FileName: An EFSRPC identifier as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// Users: A list of certificate hashes, represented by an ENCRYPTION_CERTIFICATE_HASH_LIST
	// structure, whose access is to be removed.
	Users *EncryptionCertificateHashList `idl:"name:Users" json:"users"`
}

func (o *RemoveUsersFromFileRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveUsersFromFileOperation) *xxx_RemoveUsersFromFileOperation {
	if op == nil {
		op = &xxx_RemoveUsersFromFileOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	op.Users = o.Users
	return op
}

func (o *RemoveUsersFromFileRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveUsersFromFileOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.Users = op.Users
}
func (o *RemoveUsersFromFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveUsersFromFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveUsersFromFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveUsersFromFileResponse structure represents the EfsRpcRemoveUsersFromFile operation response
type RemoveUsersFromFileResponse struct {
	// Return: The EfsRpcRemoveUsersFromFile return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveUsersFromFileResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveUsersFromFileOperation) *xxx_RemoveUsersFromFileOperation {
	if op == nil {
		op = &xxx_RemoveUsersFromFileOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RemoveUsersFromFileResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveUsersFromFileOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveUsersFromFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveUsersFromFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveUsersFromFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddUsersToFileOperation structure represents the EfsRpcAddUsersToFile operation
type xxx_AddUsersToFileOperation struct {
	FileName               string                     `idl:"name:FileName;string" json:"file_name"`
	EncryptionCertificates *EncryptionCertificateList `idl:"name:EncryptionCertificates" json:"encryption_certificates"`
	Return                 uint32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_AddUsersToFileOperation) OpNum() int { return 9 }

func (o *xxx_AddUsersToFileOperation) OpName() string { return "/efsrpc/v1/EfsRpcAddUsersToFile" }

func (o *xxx_AddUsersToFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// EncryptionCertificates {in} (1:{pointer=ref}*(1))(2:{alias=ENCRYPTION_CERTIFICATE_LIST}(struct))
	{
		if o.EncryptionCertificates != nil {
			if err := o.EncryptionCertificates.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EncryptionCertificateList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// EncryptionCertificates {in} (1:{pointer=ref}*(1))(2:{alias=ENCRYPTION_CERTIFICATE_LIST}(struct))
	{
		if o.EncryptionCertificates == nil {
			o.EncryptionCertificates = &EncryptionCertificateList{}
		}
		if err := o.EncryptionCertificates.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddUsersToFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddUsersToFileRequest structure represents the EfsRpcAddUsersToFile operation request
type AddUsersToFileRequest struct {
	// FileName: An EFSRPC nonzero name, as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// EncryptionCertificates: A list of certificates, represented by an ENCRYPTION_CERTIFICATE_LIST
	// structure, which are to be given access to the object.
	EncryptionCertificates *EncryptionCertificateList `idl:"name:EncryptionCertificates" json:"encryption_certificates"`
}

func (o *AddUsersToFileRequest) xxx_ToOp(ctx context.Context, op *xxx_AddUsersToFileOperation) *xxx_AddUsersToFileOperation {
	if op == nil {
		op = &xxx_AddUsersToFileOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	op.EncryptionCertificates = o.EncryptionCertificates
	return op
}

func (o *AddUsersToFileRequest) xxx_FromOp(ctx context.Context, op *xxx_AddUsersToFileOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.EncryptionCertificates = op.EncryptionCertificates
}
func (o *AddUsersToFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddUsersToFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddUsersToFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddUsersToFileResponse structure represents the EfsRpcAddUsersToFile operation response
type AddUsersToFileResponse struct {
	// Return: The EfsRpcAddUsersToFile return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddUsersToFileResponse) xxx_ToOp(ctx context.Context, op *xxx_AddUsersToFileOperation) *xxx_AddUsersToFileOperation {
	if op == nil {
		op = &xxx_AddUsersToFileOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AddUsersToFileResponse) xxx_FromOp(ctx context.Context, op *xxx_AddUsersToFileOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddUsersToFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddUsersToFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddUsersToFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NotSupportedOperation structure represents the EfsRpcNotSupported operation
type xxx_NotSupportedOperation struct {
	_      string `idl:"name:Reserved1;string"`
	_      string `idl:"name:Reserved2;string"`
	_      uint32 `idl:"name:dwReserved1"`
	_      uint32 `idl:"name:dwReserved2"`
	_      *Blob  `idl:"name:Reserved;pointer:unique"`
	_      bool   `idl:"name:bReserved"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_NotSupportedOperation) OpNum() int { return 11 }

func (o *xxx_NotSupportedOperation) OpName() string { return "/efsrpc/v1/EfsRpcNotSupported" }

func (o *xxx_NotSupportedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NotSupportedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Reserved1 {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		// reserved Reserved1
		if err := ndr.WriteUTF16NString(ctx, w, ""); err != nil {
			return err
		}
	}
	// Reserved2 {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		// reserved Reserved2
		if err := ndr.WriteUTF16NString(ctx, w, ""); err != nil {
			return err
		}
	}
	// dwReserved1 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved1
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// dwReserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved2
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		// reserved Reserved
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bReserved {in} (1:{alias=BOOL}(int32))
	{
		// reserved bReserved
		if err := w.WriteData(false); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NotSupportedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Reserved1 {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		// reserved Reserved1
		var _Reserved1 string
		if err := ndr.ReadUTF16NString(ctx, w, &_Reserved1); err != nil {
			return err
		}
	}
	// Reserved2 {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		// reserved Reserved2
		var _Reserved2 string
		if err := ndr.ReadUTF16NString(ctx, w, &_Reserved2); err != nil {
			return err
		}
	}
	// dwReserved1 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved1
		var _dwReserved1 uint32
		if err := w.ReadData(&_dwReserved1); err != nil {
			return err
		}
	}
	// dwReserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved2
		var _dwReserved2 uint32
		if err := w.ReadData(&_dwReserved2); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		// reserved Reserved
		var _Reserved *Blob
		_ptr_Reserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _Reserved == nil {
				_Reserved = &Blob{}
			}
			if err := _Reserved.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Reserved := func(ptr interface{}) { _Reserved = *ptr.(**Blob) }
		if err := w.ReadPointer(&_Reserved, _s_Reserved, _ptr_Reserved); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bReserved {in} (1:{alias=BOOL}(int32))
	{
		// reserved bReserved
		var _bReserved bool
		var _b_bReserved int32
		if err := w.ReadData(&_b_bReserved); err != nil {
			return err
		}
		_bReserved = _b_bReserved != 0
		_ = _bReserved
	}
	return nil
}

func (o *xxx_NotSupportedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NotSupportedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_NotSupportedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NotSupportedRequest structure represents the EfsRpcNotSupported operation request
type NotSupportedRequest struct {
}

func (o *NotSupportedRequest) xxx_ToOp(ctx context.Context, op *xxx_NotSupportedOperation) *xxx_NotSupportedOperation {
	if op == nil {
		op = &xxx_NotSupportedOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *NotSupportedRequest) xxx_FromOp(ctx context.Context, op *xxx_NotSupportedOperation) {
	if o == nil {
		return
	}
}
func (o *NotSupportedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NotSupportedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NotSupportedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NotSupportedResponse structure represents the EfsRpcNotSupported operation response
type NotSupportedResponse struct {
	// Return: The EfsRpcNotSupported return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NotSupportedResponse) xxx_ToOp(ctx context.Context, op *xxx_NotSupportedOperation) *xxx_NotSupportedOperation {
	if op == nil {
		op = &xxx_NotSupportedOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *NotSupportedResponse) xxx_FromOp(ctx context.Context, op *xxx_NotSupportedOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *NotSupportedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NotSupportedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NotSupportedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FileKeyInfoOperation structure represents the EfsRpcFileKeyInfo operation
type xxx_FileKeyInfoOperation struct {
	FileName  string `idl:"name:FileName;string" json:"file_name"`
	InfoClass uint32 `idl:"name:InfoClass" json:"info_class"`
	KeyInfo   *Blob  `idl:"name:KeyInfo" json:"key_info"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FileKeyInfoOperation) OpNum() int { return 12 }

func (o *xxx_FileKeyInfoOperation) OpName() string { return "/efsrpc/v1/EfsRpcFileKeyInfo" }

func (o *xxx_FileKeyInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// InfoClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InfoClass); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// InfoClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InfoClass); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// KeyInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		if o.KeyInfo != nil {
			_ptr_KeyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.KeyInfo != nil {
					if err := o.KeyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.KeyInfo, _ptr_KeyInfo); err != nil {
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

func (o *xxx_FileKeyInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// KeyInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		_ptr_KeyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.KeyInfo == nil {
				o.KeyInfo = &Blob{}
			}
			if err := o.KeyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_KeyInfo := func(ptr interface{}) { o.KeyInfo = *ptr.(**Blob) }
		if err := w.ReadPointer(&o.KeyInfo, _s_KeyInfo, _ptr_KeyInfo); err != nil {
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

// FileKeyInfoRequest structure represents the EfsRpcFileKeyInfo operation request
type FileKeyInfoRequest struct {
	// FileName: An EFSRPC identifier, as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// InfoClass: One of the values in the following table. With the exception of UPDATE_KEY_USED
	// (0x00000100), a server SHOULD support all of these values. A server MAY choose to
	// support UPDATE_KEY_USED.<46>
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| BASIC_KEY_INFO 0x00000001           | Request information about the keys used to encrypt the object's contents.        |
	//	|                                     | On success, the server will return the information in an EFS_KEY_INFO            |
	//	|                                     | (2.2.14)structure in the KeyInfo parameter.                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| CHECK_COMPATIBILITY_INFO 0x00000002 | Requests the EfsVersion for the encrypted file. On success, the server will      |
	//	|                                     | return the information in an EFS_COMPATIBILITY_INFO structure in the KeyInfo     |
	//	|                                     | parameter.                                                                       |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| UPDATE_KEY_USED 0x00000100          | Update the user certificates used to give a specific user access to an object.   |
	//	|                                     | The server will populate the KeyInfo parameter with a zero-terminated, wide      |
	//	|                                     | character Unicode string that contains a newline-separated list of names of      |
	//	|                                     | objects successfully updated.                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| CHECK_DECRYPTION_STATUS 0x00000200  | Request a hint from the server as to whether the given object could be           |
	//	|                                     | successfully decrypted without further user intervention or higher-level         |
	//	|                                     | events. The server will return this information in an EFS_DECRYPTION_STATUS_INFO |
	//	|                                     | structure in the KeyInfo parameter.                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| CHECK_ENCRYPTION_STATUS 0x00000400  | Request a hint from the server as to whether the given object could be           |
	//	|                                     | successfully encrypted without further user intervention or higher-level         |
	//	|                                     | events. The server will return this information in an EFS_ENCRYPTION_STATUS_INFO |
	//	|                                     | structure in the KeyInfo parameter.                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	InfoClass uint32 `idl:"name:InfoClass" json:"info_class"`
}

func (o *FileKeyInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_FileKeyInfoOperation) *xxx_FileKeyInfoOperation {
	if op == nil {
		op = &xxx_FileKeyInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	op.InfoClass = o.InfoClass
	return op
}

func (o *FileKeyInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_FileKeyInfoOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.InfoClass = op.InfoClass
}
func (o *FileKeyInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FileKeyInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FileKeyInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FileKeyInfoResponse structure represents the EfsRpcFileKeyInfo operation response
type FileKeyInfoResponse struct {
	// KeyInfo: Returned by the server, as previously specified.
	KeyInfo *Blob `idl:"name:KeyInfo" json:"key_info"`
	// Return: The EfsRpcFileKeyInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FileKeyInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_FileKeyInfoOperation) *xxx_FileKeyInfoOperation {
	if op == nil {
		op = &xxx_FileKeyInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.KeyInfo = o.KeyInfo
	op.Return = o.Return
	return op
}

func (o *FileKeyInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_FileKeyInfoOperation) {
	if o == nil {
		return
	}
	o.KeyInfo = op.KeyInfo
	o.Return = op.Return
}
func (o *FileKeyInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FileKeyInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FileKeyInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DuplicateEncryptionInfoFileOperation structure represents the EfsRpcDuplicateEncryptionInfoFile operation
type xxx_DuplicateEncryptionInfoFileOperation struct {
	SourceFileName             string `idl:"name:SrcFileName;string" json:"source_file_name"`
	DestinationFileName        string `idl:"name:DestFileName;string" json:"destination_file_name"`
	CreationDisposition        uint32 `idl:"name:dwCreationDisposition" json:"creation_disposition"`
	Attributes                 uint32 `idl:"name:dwAttributes" json:"attributes"`
	RelativeSecurityDescriptor *Blob  `idl:"name:RelativeSD;pointer:unique" json:"relative_security_descriptor"`
	InheritHandle              bool   `idl:"name:bInheritHandle" json:"inherit_handle"`
	Return                     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DuplicateEncryptionInfoFileOperation) OpNum() int { return 13 }

func (o *xxx_DuplicateEncryptionInfoFileOperation) OpName() string {
	return "/efsrpc/v1/EfsRpcDuplicateEncryptionInfoFile"
}

func (o *xxx_DuplicateEncryptionInfoFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DuplicateEncryptionInfoFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SrcFileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SourceFileName); err != nil {
			return err
		}
	}
	// DestFileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DestinationFileName); err != nil {
			return err
		}
	}
	// dwCreationDisposition {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreationDisposition); err != nil {
			return err
		}
	}
	// dwAttributes {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Attributes); err != nil {
			return err
		}
	}
	// RelativeSD {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		if o.RelativeSecurityDescriptor != nil {
			_ptr_RelativeSD := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RelativeSecurityDescriptor != nil {
					if err := o.RelativeSecurityDescriptor.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RelativeSecurityDescriptor, _ptr_RelativeSD); err != nil {
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
	// bInheritHandle {in} (1:{alias=BOOL}(int32))
	{
		if !o.InheritHandle {
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

func (o *xxx_DuplicateEncryptionInfoFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SrcFileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SourceFileName); err != nil {
			return err
		}
	}
	// DestFileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationFileName); err != nil {
			return err
		}
	}
	// dwCreationDisposition {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreationDisposition); err != nil {
			return err
		}
	}
	// dwAttributes {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Attributes); err != nil {
			return err
		}
	}
	// RelativeSD {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		_ptr_RelativeSD := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RelativeSecurityDescriptor == nil {
				o.RelativeSecurityDescriptor = &Blob{}
			}
			if err := o.RelativeSecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_RelativeSD := func(ptr interface{}) { o.RelativeSecurityDescriptor = *ptr.(**Blob) }
		if err := w.ReadPointer(&o.RelativeSecurityDescriptor, _s_RelativeSD, _ptr_RelativeSD); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bInheritHandle {in} (1:{alias=BOOL}(int32))
	{
		var _bInheritHandle int32
		if err := w.ReadData(&_bInheritHandle); err != nil {
			return err
		}
		o.InheritHandle = _bInheritHandle != 0
	}
	return nil
}

func (o *xxx_DuplicateEncryptionInfoFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DuplicateEncryptionInfoFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DuplicateEncryptionInfoFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DuplicateEncryptionInfoFileRequest structure represents the EfsRpcDuplicateEncryptionInfoFile operation request
type DuplicateEncryptionInfoFileRequest struct {
	// SrcFileName:  An EFSRPC identifier, as specified in section 2.2.1.
	SourceFileName string `idl:"name:SrcFileName;string" json:"source_file_name"`
	// DestFileName: An EFSRPC identifier, as specified in section 2.2.1.
	DestinationFileName string `idl:"name:DestFileName;string" json:"destination_file_name"`
	// dwCreationDisposition: This parameter specifies what action the server is advised
	// to take if the object referred to by DestFileName does not already exist. It MUST
	// be one of the following values.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|                          |                                                                                  |
	//	|          VALUE           |                                     MEANING                                      |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| CREATE_NEW 0x00000001    | Do not overwrite the data object referred to by DestFileName if it already       |
	//	|                          | exists.                                                                          |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| CREATE_ALWAYS 0x00000002 | Overwrite the data object referred to by DestFileName if it already exists.      |
	//	+--------------------------+----------------------------------------------------------------------------------+
	CreationDisposition uint32 `idl:"name:dwCreationDisposition" json:"creation_disposition"`
	// dwAttributes: Desired attributes for the target object. Clients SHOULD set this parameter
	// to the bitwise OR of zero or more of the following values. Servers SHOULD support
	// all of these values. These values can be interpreted by the underlying server data
	// store, and they do not affect protocol behavior.
	//
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	//	|                                               |                                                                             |
	//	|                     VALUE                     |                                   MEANING                                   |
	//	|                                               |                                                                             |
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	//	| FILE_ATTRIBUTE_HIDDEN 0x00000002              | The file is hidden (not displayed in normal folder listings).               |
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	//	| FILE_ATTRIBUTE_ARCHIVE 0x00000020             | This attribute is used by applications to mark files for backup or removal. |
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	//	| FILE_ATTRIBUTE_TEMPORARY 0x00000100           | The file is being used for temporary storage.                               |
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	//	| FILE_ATTRIBUTE_NOT_CONTENT_INDEXED 0x00002000 | The file's contents are not to be indexed by the content indexing service.  |
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	//	| FILE_ATTRIBUTE_NORMAL 0x00000080              | No other attributes are to be set.                                          |
	//	+-----------------------------------------------+-----------------------------------------------------------------------------+
	Attributes uint32 `idl:"name:dwAttributes" json:"attributes"`
	// RelativeSD: Relative security descriptor for the target object. The format of this
	// is implementation-dependent.<49>
	RelativeSecurityDescriptor *Blob `idl:"name:RelativeSD;pointer:unique" json:"relative_security_descriptor"`
	// bInheritHandle: This parameter SHOULD be set to FALSE by the client and SHOULD be
	// ignored by the server.
	InheritHandle bool `idl:"name:bInheritHandle" json:"inherit_handle"`
}

func (o *DuplicateEncryptionInfoFileRequest) xxx_ToOp(ctx context.Context, op *xxx_DuplicateEncryptionInfoFileOperation) *xxx_DuplicateEncryptionInfoFileOperation {
	if op == nil {
		op = &xxx_DuplicateEncryptionInfoFileOperation{}
	}
	if o == nil {
		return op
	}
	op.SourceFileName = o.SourceFileName
	op.DestinationFileName = o.DestinationFileName
	op.CreationDisposition = o.CreationDisposition
	op.Attributes = o.Attributes
	op.RelativeSecurityDescriptor = o.RelativeSecurityDescriptor
	op.InheritHandle = o.InheritHandle
	return op
}

func (o *DuplicateEncryptionInfoFileRequest) xxx_FromOp(ctx context.Context, op *xxx_DuplicateEncryptionInfoFileOperation) {
	if o == nil {
		return
	}
	o.SourceFileName = op.SourceFileName
	o.DestinationFileName = op.DestinationFileName
	o.CreationDisposition = op.CreationDisposition
	o.Attributes = op.Attributes
	o.RelativeSecurityDescriptor = op.RelativeSecurityDescriptor
	o.InheritHandle = op.InheritHandle
}
func (o *DuplicateEncryptionInfoFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DuplicateEncryptionInfoFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DuplicateEncryptionInfoFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DuplicateEncryptionInfoFileResponse structure represents the EfsRpcDuplicateEncryptionInfoFile operation response
type DuplicateEncryptionInfoFileResponse struct {
	// Return: The EfsRpcDuplicateEncryptionInfoFile return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DuplicateEncryptionInfoFileResponse) xxx_ToOp(ctx context.Context, op *xxx_DuplicateEncryptionInfoFileOperation) *xxx_DuplicateEncryptionInfoFileOperation {
	if op == nil {
		op = &xxx_DuplicateEncryptionInfoFileOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *DuplicateEncryptionInfoFileResponse) xxx_FromOp(ctx context.Context, op *xxx_DuplicateEncryptionInfoFileOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DuplicateEncryptionInfoFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DuplicateEncryptionInfoFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DuplicateEncryptionInfoFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddUsersToFileExOperation structure represents the EfsRpcAddUsersToFileEx operation
type xxx_AddUsersToFileExOperation struct {
	Flags                  uint32                     `idl:"name:dwFlags" json:"flags"`
	_                      *Blob                      `idl:"name:Reserved;pointer:unique"`
	FileName               string                     `idl:"name:FileName;string" json:"file_name"`
	EncryptionCertificates *EncryptionCertificateList `idl:"name:EncryptionCertificates" json:"encryption_certificates"`
	Return                 uint32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_AddUsersToFileExOperation) OpNum() int { return 15 }

func (o *xxx_AddUsersToFileExOperation) OpName() string { return "/efsrpc/v1/EfsRpcAddUsersToFileEx" }

func (o *xxx_AddUsersToFileExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		// reserved Reserved
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// EncryptionCertificates {in} (1:{pointer=ref}*(1))(2:{alias=ENCRYPTION_CERTIFICATE_LIST}(struct))
	{
		if o.EncryptionCertificates != nil {
			if err := o.EncryptionCertificates.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EncryptionCertificateList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		// reserved Reserved
		var _Reserved *Blob
		_ptr_Reserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _Reserved == nil {
				_Reserved = &Blob{}
			}
			if err := _Reserved.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Reserved := func(ptr interface{}) { _Reserved = *ptr.(**Blob) }
		if err := w.ReadPointer(&_Reserved, _s_Reserved, _ptr_Reserved); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// EncryptionCertificates {in} (1:{pointer=ref}*(1))(2:{alias=ENCRYPTION_CERTIFICATE_LIST}(struct))
	{
		if o.EncryptionCertificates == nil {
			o.EncryptionCertificates = &EncryptionCertificateList{}
		}
		if err := o.EncryptionCertificates.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddUsersToFileExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddUsersToFileExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddUsersToFileExRequest structure represents the EfsRpcAddUsersToFileEx operation request
type AddUsersToFileExRequest struct {
	// dwFlags: This MUST be set to a bitwise OR of 0 or more of the following flags. The
	// descriptions of the flags are specified in the following table. If the EFSRPC_ADDUSERFLAG_REPLACE_DDF
	// flag is used, then the EncryptionCertificates parameter MUST contain exactly one
	// certificate.
	//
	//	+---------------------------------------+------------+
	//	|                                       |            |
	//	|                 NAME                  |   VALUE    |
	//	|                                       |            |
	//	+---------------------------------------+------------+
	//	+---------------------------------------+------------+
	//	| EFSRPC_ADDUSERFLAG_ADD_POLICY_KEYTYPE | 0x00000002 |
	//	+---------------------------------------+------------+
	//	| EFSRPC_ADDUSERFLAG_REPLACE_DDF        | 0x00000004 |
	//	+---------------------------------------+------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// FileName: An EFSRPC identifier, as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// EncryptionCertificates: A list of certificates, represented by an ENCRYPTION_CERTIFICATE_LIST
	// structure, which are to be given access to the object.
	EncryptionCertificates *EncryptionCertificateList `idl:"name:EncryptionCertificates" json:"encryption_certificates"`
}

func (o *AddUsersToFileExRequest) xxx_ToOp(ctx context.Context, op *xxx_AddUsersToFileExOperation) *xxx_AddUsersToFileExOperation {
	if op == nil {
		op = &xxx_AddUsersToFileExOperation{}
	}
	if o == nil {
		return op
	}
	op.Flags = o.Flags
	op.FileName = o.FileName
	op.EncryptionCertificates = o.EncryptionCertificates
	return op
}

func (o *AddUsersToFileExRequest) xxx_FromOp(ctx context.Context, op *xxx_AddUsersToFileExOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.FileName = op.FileName
	o.EncryptionCertificates = op.EncryptionCertificates
}
func (o *AddUsersToFileExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddUsersToFileExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddUsersToFileExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddUsersToFileExResponse structure represents the EfsRpcAddUsersToFileEx operation response
type AddUsersToFileExResponse struct {
	// Return: The EfsRpcAddUsersToFileEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddUsersToFileExResponse) xxx_ToOp(ctx context.Context, op *xxx_AddUsersToFileExOperation) *xxx_AddUsersToFileExOperation {
	if op == nil {
		op = &xxx_AddUsersToFileExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AddUsersToFileExResponse) xxx_FromOp(ctx context.Context, op *xxx_AddUsersToFileExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddUsersToFileExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddUsersToFileExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddUsersToFileExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FileKeyInfoExOperation structure represents the EfsRpcFileKeyInfoEx operation
type xxx_FileKeyInfoExOperation struct {
	FileKeyInfoFlags uint32 `idl:"name:dwFileKeyInfoFlags" json:"file_key_info_flags"`
	_                *Blob  `idl:"name:Reserved;pointer:unique"`
	FileName         string `idl:"name:FileName;string" json:"file_name"`
	InfoClass        uint32 `idl:"name:InfoClass" json:"info_class"`
	KeyInfo          *Blob  `idl:"name:KeyInfo" json:"key_info"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FileKeyInfoExOperation) OpNum() int { return 16 }

func (o *xxx_FileKeyInfoExOperation) OpName() string { return "/efsrpc/v1/EfsRpcFileKeyInfoEx" }

func (o *xxx_FileKeyInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwFileKeyInfoFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FileKeyInfoFlags); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		// reserved Reserved
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// InfoClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InfoClass); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwFileKeyInfoFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FileKeyInfoFlags); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		// reserved Reserved
		var _Reserved *Blob
		_ptr_Reserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _Reserved == nil {
				_Reserved = &Blob{}
			}
			if err := _Reserved.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Reserved := func(ptr interface{}) { _Reserved = *ptr.(**Blob) }
		if err := w.ReadPointer(&_Reserved, _s_Reserved, _ptr_Reserved); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// InfoClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InfoClass); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileKeyInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// KeyInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		if o.KeyInfo != nil {
			_ptr_KeyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.KeyInfo != nil {
					if err := o.KeyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.KeyInfo, _ptr_KeyInfo); err != nil {
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

func (o *xxx_FileKeyInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// KeyInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		_ptr_KeyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.KeyInfo == nil {
				o.KeyInfo = &Blob{}
			}
			if err := o.KeyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_KeyInfo := func(ptr interface{}) { o.KeyInfo = *ptr.(**Blob) }
		if err := w.ReadPointer(&o.KeyInfo, _s_KeyInfo, _ptr_KeyInfo); err != nil {
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

// FileKeyInfoExRequest structure represents the EfsRpcFileKeyInfoEx operation request
type FileKeyInfoExRequest struct {
	// dwFileKeyInfoFlags: This parameter is reserved. It MUST be set to zero by the client
	// and ignored by the server.
	FileKeyInfoFlags uint32 `idl:"name:dwFileKeyInfoFlags" json:"file_key_info_flags"`
	// FileName: An EFSRPC identifier, as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// InfoClass: One of the values specified for the InfoClass parameter of the EfsRpcFileKeyInfo
	// method.
	InfoClass uint32 `idl:"name:InfoClass" json:"info_class"`
}

func (o *FileKeyInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_FileKeyInfoExOperation) *xxx_FileKeyInfoExOperation {
	if op == nil {
		op = &xxx_FileKeyInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.FileKeyInfoFlags = o.FileKeyInfoFlags
	op.FileName = o.FileName
	op.InfoClass = o.InfoClass
	return op
}

func (o *FileKeyInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_FileKeyInfoExOperation) {
	if o == nil {
		return
	}
	o.FileKeyInfoFlags = op.FileKeyInfoFlags
	o.FileName = op.FileName
	o.InfoClass = op.InfoClass
}
func (o *FileKeyInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FileKeyInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FileKeyInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FileKeyInfoExResponse structure represents the EfsRpcFileKeyInfoEx operation response
type FileKeyInfoExResponse struct {
	// KeyInfo: The server SHOULD ignore this parameter.<53>
	KeyInfo *Blob `idl:"name:KeyInfo" json:"key_info"`
	// Return: The EfsRpcFileKeyInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FileKeyInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_FileKeyInfoExOperation) *xxx_FileKeyInfoExOperation {
	if op == nil {
		op = &xxx_FileKeyInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.KeyInfo = o.KeyInfo
	op.Return = o.Return
	return op
}

func (o *FileKeyInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_FileKeyInfoExOperation) {
	if o == nil {
		return
	}
	o.KeyInfo = op.KeyInfo
	o.Return = op.Return
}
func (o *FileKeyInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FileKeyInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FileKeyInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEncryptedFileMetadataOperation structure represents the EfsRpcGetEncryptedFileMetadata operation
type xxx_GetEncryptedFileMetadataOperation struct {
	FileName   string `idl:"name:FileName;string;pointer:ref" json:"file_name"`
	StreamBlob *Blob  `idl:"name:EfsStreamBlob;pointer:ref" json:"stream_blob"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEncryptedFileMetadataOperation) OpNum() int { return 18 }

func (o *xxx_GetEncryptedFileMetadataOperation) OpName() string {
	return "/efsrpc/v1/EfsRpcGetEncryptedFileMetadata"
}

func (o *xxx_GetEncryptedFileMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncryptedFileMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncryptedFileMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncryptedFileMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEncryptedFileMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EfsStreamBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		if o.StreamBlob != nil {
			_ptr_EfsStreamBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StreamBlob != nil {
					if err := o.StreamBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StreamBlob, _ptr_EfsStreamBlob); err != nil {
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

func (o *xxx_GetEncryptedFileMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EfsStreamBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		_ptr_EfsStreamBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StreamBlob == nil {
				o.StreamBlob = &Blob{}
			}
			if err := o.StreamBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EfsStreamBlob := func(ptr interface{}) { o.StreamBlob = *ptr.(**Blob) }
		if err := w.ReadPointer(&o.StreamBlob, _s_EfsStreamBlob, _ptr_EfsStreamBlob); err != nil {
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

// GetEncryptedFileMetadataRequest structure represents the EfsRpcGetEncryptedFileMetadata operation request
type GetEncryptedFileMetadataRequest struct {
	// FileName: An EFSRPC identifier, as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string;pointer:ref" json:"file_name"`
}

func (o *GetEncryptedFileMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEncryptedFileMetadataOperation) *xxx_GetEncryptedFileMetadataOperation {
	if op == nil {
		op = &xxx_GetEncryptedFileMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	return op
}

func (o *GetEncryptedFileMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEncryptedFileMetadataOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
}
func (o *GetEncryptedFileMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEncryptedFileMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEncryptedFileMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEncryptedFileMetadataResponse structure represents the EfsRpcGetEncryptedFileMetadata operation response
type GetEncryptedFileMetadataResponse struct {
	// EfsStreamBlob: The server SHOULD ignore this parameter.<56>
	StreamBlob *Blob `idl:"name:EfsStreamBlob;pointer:ref" json:"stream_blob"`
	// Return: The EfsRpcGetEncryptedFileMetadata return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetEncryptedFileMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEncryptedFileMetadataOperation) *xxx_GetEncryptedFileMetadataOperation {
	if op == nil {
		op = &xxx_GetEncryptedFileMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.StreamBlob = o.StreamBlob
	op.Return = o.Return
	return op
}

func (o *GetEncryptedFileMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEncryptedFileMetadataOperation) {
	if o == nil {
		return
	}
	o.StreamBlob = op.StreamBlob
	o.Return = op.Return
}
func (o *GetEncryptedFileMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEncryptedFileMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEncryptedFileMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEncryptedFileMetadataOperation structure represents the EfsRpcSetEncryptedFileMetadata operation
type xxx_SetEncryptedFileMetadataOperation struct {
	FileName         string                          `idl:"name:FileName;string;pointer:ref" json:"file_name"`
	OldEFSStreamBlob *Blob                           `idl:"name:OldEfsStreamBlob;pointer:unique" json:"old_efs_stream_blob"`
	NewEFSStreamBlob *Blob                           `idl:"name:NewEfsStreamBlob;pointer:ref" json:"new_efs_stream_blob"`
	NewEFSSignature  *EncryptedFileMetadataSignature `idl:"name:NewEfsSignature;pointer:unique" json:"new_efs_signature"`
	Return           uint32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEncryptedFileMetadataOperation) OpNum() int { return 19 }

func (o *xxx_SetEncryptedFileMetadataOperation) OpName() string {
	return "/efsrpc/v1/EfsRpcSetEncryptedFileMetadata"
}

func (o *xxx_SetEncryptedFileMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEncryptedFileMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// OldEfsStreamBlob {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		if o.OldEFSStreamBlob != nil {
			_ptr_OldEfsStreamBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldEFSStreamBlob != nil {
					if err := o.OldEFSStreamBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldEFSStreamBlob, _ptr_OldEfsStreamBlob); err != nil {
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
	// NewEfsStreamBlob {in} (1:{pointer=ref}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		if o.NewEFSStreamBlob != nil {
			if err := o.NewEFSStreamBlob.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// NewEfsSignature {in} (1:{pointer=unique}*(1))(2:{alias=ENCRYPTED_FILE_METADATA_SIGNATURE}(struct))
	{
		if o.NewEFSSignature != nil {
			_ptr_NewEfsSignature := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewEFSSignature != nil {
					if err := o.NewEFSSignature.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedFileMetadataSignature{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewEFSSignature, _ptr_NewEfsSignature); err != nil {
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
	return nil
}

func (o *xxx_SetEncryptedFileMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// OldEfsStreamBlob {in} (1:{pointer=unique}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		_ptr_OldEfsStreamBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldEFSStreamBlob == nil {
				o.OldEFSStreamBlob = &Blob{}
			}
			if err := o.OldEFSStreamBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldEfsStreamBlob := func(ptr interface{}) { o.OldEFSStreamBlob = *ptr.(**Blob) }
		if err := w.ReadPointer(&o.OldEFSStreamBlob, _s_OldEfsStreamBlob, _ptr_OldEfsStreamBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NewEfsStreamBlob {in} (1:{pointer=ref}*(1))(2:{alias=EFS_RPC_BLOB}(struct))
	{
		if o.NewEFSStreamBlob == nil {
			o.NewEFSStreamBlob = &Blob{}
		}
		if err := o.NewEFSStreamBlob.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NewEfsSignature {in} (1:{pointer=unique}*(1))(2:{alias=ENCRYPTED_FILE_METADATA_SIGNATURE}(struct))
	{
		_ptr_NewEfsSignature := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewEFSSignature == nil {
				o.NewEFSSignature = &EncryptedFileMetadataSignature{}
			}
			if err := o.NewEFSSignature.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewEfsSignature := func(ptr interface{}) { o.NewEFSSignature = *ptr.(**EncryptedFileMetadataSignature) }
		if err := w.ReadPointer(&o.NewEFSSignature, _s_NewEfsSignature, _ptr_NewEfsSignature); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEncryptedFileMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEncryptedFileMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEncryptedFileMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetEncryptedFileMetadataRequest structure represents the EfsRpcSetEncryptedFileMetadata operation request
type SetEncryptedFileMetadataRequest struct {
	// FileName: An EFSRPC identifier as specified in section 2.2.1.
	FileName string `idl:"name:FileName;string;pointer:ref" json:"file_name"`
	// OldEfsStreamBlob: This parameter SHOULD be set to NULL by the client and ignored
	// by the server.<59>
	OldEFSStreamBlob *Blob `idl:"name:OldEfsStreamBlob;pointer:unique" json:"old_efs_stream_blob"`
	// NewEfsStreamBlob: This parameter SHOULD be set to a zero-length EFS_RPC_BLOB by the
	// client and ignored by the server.<60>
	NewEFSStreamBlob *Blob `idl:"name:NewEfsStreamBlob;pointer:ref" json:"new_efs_stream_blob"`
	// NewEfsSignature:  This parameter SHOULD be set to NULL by the client and ignored
	// by the server.<61>
	NewEFSSignature *EncryptedFileMetadataSignature `idl:"name:NewEfsSignature;pointer:unique" json:"new_efs_signature"`
}

func (o *SetEncryptedFileMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEncryptedFileMetadataOperation) *xxx_SetEncryptedFileMetadataOperation {
	if op == nil {
		op = &xxx_SetEncryptedFileMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	op.OldEFSStreamBlob = o.OldEFSStreamBlob
	op.NewEFSStreamBlob = o.NewEFSStreamBlob
	op.NewEFSSignature = o.NewEFSSignature
	return op
}

func (o *SetEncryptedFileMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEncryptedFileMetadataOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.OldEFSStreamBlob = op.OldEFSStreamBlob
	o.NewEFSStreamBlob = op.NewEFSStreamBlob
	o.NewEFSSignature = op.NewEFSSignature
}
func (o *SetEncryptedFileMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEncryptedFileMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEncryptedFileMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEncryptedFileMetadataResponse structure represents the EfsRpcSetEncryptedFileMetadata operation response
type SetEncryptedFileMetadataResponse struct {
	// Return: The EfsRpcSetEncryptedFileMetadata return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetEncryptedFileMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEncryptedFileMetadataOperation) *xxx_SetEncryptedFileMetadataOperation {
	if op == nil {
		op = &xxx_SetEncryptedFileMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetEncryptedFileMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEncryptedFileMetadataOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetEncryptedFileMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEncryptedFileMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEncryptedFileMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FlushEFSCacheOperation structure represents the EfsRpcFlushEfsCache operation
type xxx_FlushEFSCacheOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FlushEFSCacheOperation) OpNum() int { return 20 }

func (o *xxx_FlushEFSCacheOperation) OpName() string { return "/efsrpc/v1/EfsRpcFlushEfsCache" }

func (o *xxx_FlushEFSCacheOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushEFSCacheOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_FlushEFSCacheOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_FlushEFSCacheOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushEFSCacheOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FlushEFSCacheOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FlushEFSCacheRequest structure represents the EfsRpcFlushEfsCache operation request
type FlushEFSCacheRequest struct {
}

func (o *FlushEFSCacheRequest) xxx_ToOp(ctx context.Context, op *xxx_FlushEFSCacheOperation) *xxx_FlushEFSCacheOperation {
	if op == nil {
		op = &xxx_FlushEFSCacheOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *FlushEFSCacheRequest) xxx_FromOp(ctx context.Context, op *xxx_FlushEFSCacheOperation) {
	if o == nil {
		return
	}
}
func (o *FlushEFSCacheRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FlushEFSCacheRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushEFSCacheOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FlushEFSCacheResponse structure represents the EfsRpcFlushEfsCache operation response
type FlushEFSCacheResponse struct {
	// Return: The EfsRpcFlushEfsCache return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FlushEFSCacheResponse) xxx_ToOp(ctx context.Context, op *xxx_FlushEFSCacheOperation) *xxx_FlushEFSCacheOperation {
	if op == nil {
		op = &xxx_FlushEFSCacheOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FlushEFSCacheResponse) xxx_FromOp(ctx context.Context, op *xxx_FlushEFSCacheOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FlushEFSCacheResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FlushEFSCacheResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushEFSCacheOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EncryptFileExServerOperation structure represents the EfsRpcEncryptFileExSrv operation
type xxx_EncryptFileExServerOperation struct {
	FileName            string `idl:"name:FileName;string" json:"file_name"`
	ProtectorDescriptor string `idl:"name:ProtectorDescriptor;string;pointer:unique" json:"protector_descriptor"`
	Flags               uint32 `idl:"name:Flags" json:"flags"`
	Return              int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_EncryptFileExServerOperation) OpNum() int { return 21 }

func (o *xxx_EncryptFileExServerOperation) OpName() string {
	return "/efsrpc/v1/EfsRpcEncryptFileExSrv"
}

func (o *xxx_EncryptFileExServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileExServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// ProtectorDescriptor {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ProtectorDescriptor != "" {
			_ptr_ProtectorDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ProtectorDescriptor); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ProtectorDescriptor, _ptr_ProtectorDescriptor); err != nil {
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
	// Flags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileExServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// ProtectorDescriptor {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ProtectorDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ProtectorDescriptor); err != nil {
				return err
			}
			return nil
		})
		_s_ProtectorDescriptor := func(ptr interface{}) { o.ProtectorDescriptor = *ptr.(*string) }
		if err := w.ReadPointer(&o.ProtectorDescriptor, _s_ProtectorDescriptor, _ptr_ProtectorDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileExServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileExServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EncryptFileExServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EncryptFileExServerRequest structure represents the EfsRpcEncryptFileExSrv operation request
type EncryptFileExServerRequest struct {
	FileName            string `idl:"name:FileName;string" json:"file_name"`
	ProtectorDescriptor string `idl:"name:ProtectorDescriptor;string;pointer:unique" json:"protector_descriptor"`
	Flags               uint32 `idl:"name:Flags" json:"flags"`
}

func (o *EncryptFileExServerRequest) xxx_ToOp(ctx context.Context, op *xxx_EncryptFileExServerOperation) *xxx_EncryptFileExServerOperation {
	if op == nil {
		op = &xxx_EncryptFileExServerOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	op.ProtectorDescriptor = o.ProtectorDescriptor
	op.Flags = o.Flags
	return op
}

func (o *EncryptFileExServerRequest) xxx_FromOp(ctx context.Context, op *xxx_EncryptFileExServerOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.ProtectorDescriptor = op.ProtectorDescriptor
	o.Flags = op.Flags
}
func (o *EncryptFileExServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EncryptFileExServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EncryptFileExServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EncryptFileExServerResponse structure represents the EfsRpcEncryptFileExSrv operation response
type EncryptFileExServerResponse struct {
	// Return: The EfsRpcEncryptFileExSrv return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EncryptFileExServerResponse) xxx_ToOp(ctx context.Context, op *xxx_EncryptFileExServerOperation) *xxx_EncryptFileExServerOperation {
	if op == nil {
		op = &xxx_EncryptFileExServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EncryptFileExServerResponse) xxx_FromOp(ctx context.Context, op *xxx_EncryptFileExServerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EncryptFileExServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EncryptFileExServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EncryptFileExServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryProtectorsOperation structure represents the EfsRpcQueryProtectors operation
type xxx_QueryProtectorsOperation struct {
	FileName      string                   `idl:"name:FileName;string" json:"file_name"`
	ProtectorList *EncryptionProtectorList `idl:"name:ppProtectorList" json:"protector_list"`
	Return        uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryProtectorsOperation) OpNum() int { return 22 }

func (o *xxx_QueryProtectorsOperation) OpName() string { return "/efsrpc/v1/EfsRpcQueryProtectors" }

func (o *xxx_QueryProtectorsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProtectorsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProtectorsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProtectorsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProtectorsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppProtectorList {out} (1:{pointer=ref}*(3)*(2))(2:{alias=PENCRYPTION_PROTECTOR_LIST}*(1))(3:{alias=ENCRYPTION_PROTECTOR_LIST}(struct))
	{
		if o.ProtectorList != nil {
			_ptr_ppProtectorList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProtectorList != nil {
					_ptr_ppProtectorList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.ProtectorList != nil {
							if err := o.ProtectorList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&EncryptionProtectorList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.ProtectorList, _ptr_ppProtectorList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProtectorList, _ptr_ppProtectorList); err != nil {
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

func (o *xxx_QueryProtectorsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppProtectorList {out} (1:{pointer=ref}*(3)*(2))(2:{alias=PENCRYPTION_PROTECTOR_LIST,pointer=ref}*(1))(3:{alias=ENCRYPTION_PROTECTOR_LIST}(struct))
	{
		_ptr_ppProtectorList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppProtectorList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ProtectorList == nil {
					o.ProtectorList = &EncryptionProtectorList{}
				}
				if err := o.ProtectorList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppProtectorList := func(ptr interface{}) { o.ProtectorList = *ptr.(**EncryptionProtectorList) }
			if err := w.ReadPointer(&o.ProtectorList, _s_ppProtectorList, _ptr_ppProtectorList); err != nil {
				return err
			}
			return nil
		})
		_s_ppProtectorList := func(ptr interface{}) { o.ProtectorList = *ptr.(**EncryptionProtectorList) }
		if err := w.ReadPointer(&o.ProtectorList, _s_ppProtectorList, _ptr_ppProtectorList); err != nil {
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

// QueryProtectorsRequest structure represents the EfsRpcQueryProtectors operation request
type QueryProtectorsRequest struct {
	FileName string `idl:"name:FileName;string" json:"file_name"`
}

func (o *QueryProtectorsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryProtectorsOperation) *xxx_QueryProtectorsOperation {
	if op == nil {
		op = &xxx_QueryProtectorsOperation{}
	}
	if o == nil {
		return op
	}
	op.FileName = o.FileName
	return op
}

func (o *QueryProtectorsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryProtectorsOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
}
func (o *QueryProtectorsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryProtectorsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProtectorsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryProtectorsResponse structure represents the EfsRpcQueryProtectors operation response
type QueryProtectorsResponse struct {
	ProtectorList *EncryptionProtectorList `idl:"name:ppProtectorList" json:"protector_list"`
	// Return: The EfsRpcQueryProtectors return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryProtectorsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryProtectorsOperation) *xxx_QueryProtectorsOperation {
	if op == nil {
		op = &xxx_QueryProtectorsOperation{}
	}
	if o == nil {
		return op
	}
	op.ProtectorList = o.ProtectorList
	op.Return = o.Return
	return op
}

func (o *QueryProtectorsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryProtectorsOperation) {
	if o == nil {
		return
	}
	o.ProtectorList = op.ProtectorList
	o.Return = op.Return
}
func (o *QueryProtectorsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryProtectorsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProtectorsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

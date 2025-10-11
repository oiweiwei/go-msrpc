// The efsr package implements the EFSR client protocol.
//
// # Introduction
//
// The Encrypting File System Remote (EFSRPC) Protocol is used for performing maintenance
// and management operations on encrypted data that is stored remotely and accessed
// over a network. It is used in Windows to manage files that reside on remote file
// servers and are encrypted using the Encrypting File System (EFS).
//
// # Overview
//
// The Encrypting File System Remote Protocol (hereafter referred to as EFSRPC) is a
// Remote Procedure Call (RPC) interface that is used to manage data objects stored
// in an encrypted form. The objective of encrypting data in this fashion is to enforce
// access control policies and to provide confidentiality from unauthorized users.
//
// EFSRPC is implemented in Windows to provide remote management for files encrypted
// by the Encrypting File System (EFS). EFS is the ability of the New Technology File
// System (NTFS) file system to encrypt files on disk in a manner that is transparent
// to the user. For more information on EFS, see [MSFT-EFS]. For more information about
// NTFS, see [MSFT-NTFS].
//
// EFSRPC does not address how data is encrypted, how the encrypted data is stored,
// or how it is accessed for routine operations such as reading, writing, creating,
// and deleting. All these actions are specific to the server implementation. On Windows,
// NTFS provides the storage mechanism (the file is the unit of storage) and the Server
// Message Block (SMB) Protocol provides remote access to such files. For more information
// about SMB, see [MS-SMB] and [MS-SMB2].
//
// EFSRPC models the underlying data encryption architecture using two basic constructs:
//
// * A set of data objects, each of which is encrypted independently and can be managed
// independently.
//
// * A set of access control subjects, each of which is represented by a key pair generated
// by a public key ( 230807ac-20be-494f-86e3-4c8ac23ea584#gt_4cf96ca0-e3a9-4165-8d1a-a21b1397007a
// ) cryptographic algorithm. The public key of this key pair is embedded in a certificate
// ( 230807ac-20be-494f-86e3-4c8ac23ea584#gt_7a0f4b71-23ba-434f-b781-28053ed64879 )
// and can be widely distributed in that form. The corresponding private key ( 230807ac-20be-494f-86e3-4c8ac23ea584#gt_6fca10f4-e829-42ab-ad40-1566585060ca
// ) is held solely by the user or users who represent that subject. Thus, a given access
// control subject can correspond to one or more users, and a given user can possess
// the private keys for zero or more access control subjects. Access control subjects
// are further divided into two types:
//
// * Unprivileged user subjects, which are used for routine data access by ordinary
// users of the system. For convenience, this specification refers to such subjects
// as user certificate.
//
// * Data Recovery Agents (DRAs) ( 230807ac-20be-494f-86e3-4c8ac23ea584#gt_c2d1bb54-31a8-4918-a163-39a7851c347a
// ) , which are controlled by system administrators. The storage system ensures that
// all active DRAs for the system are automatically authorized to access all encrypted
// objects on the system. If an unprivileged user loses the private key, an administrator
// can use a DRA's private key to recover the contents of encrypted objects.
//
// EFSRPC also assumes that each encrypted object is associated with some security-related
// metadata, which contains information required for authorized users and DRAs to access
// the plaintext of the object. This specification refers to this security-related metadata
// as the EFSRPC Metadata.
//
// EFSRPC does not specify how data is encrypted, stored, or accessed. It is possible
// to build a compliant EFSRPC implementation that uses a mechanism, such as access
// control lists (ACLs), instead of encryption to control access to data objects. For
// the purposes of this specification, the term encrypted is used to indicate that a
// data object and its metadata can be successfully manipulated through the EFSRPC methods,
// with the exception of the EfsRpcEncryptFileSrv method, which converts data objects
// from an unencrypted state to an encrypted state.
//
// Within the preceding model, EFSRPC provides various categories of management routines.
// The syntax of the individual methods and rules for how these methods are processed
// on the server are specified in section 3.1.4.2. The categories of management routines
// that EFSRPC provides are as follows:
//
// * Requesting the server to convert objects from encrypted state to unencrypted state
// and vice versa.
//
// * EfsRpcEncryptFileSrv (section 3.1.4.2.5)
//
// * EfsRpcDecryptFileSrv (section 3.1.4.2.6) ( 043715de-caee-402a-a61b-921743337e78
// )
//
// * Creating, querying, and manipulating the EFSRPC Metadata. Clients use the following
// methods to query and change which user certificates can be used to decrypt an encrypted
// object. The set of user certificates with access to an object needs to be changed
// when the set of users with access to the object changes or when a user with access
// to the object changes the user certificate. The following methods can also be used
// to copy the access rights from one object to another; the EfsRpcDuplicateEncryptionInfoFile
// ( b39ec3e2-d3f0-4934-925e-74032365f9d2 ) method is particularly well-suited for this
// purpose. Methods:
//
// * EfsRpcQueryUsersOnFile (section 3.1.4.2.7) ( a058dc6c-bb7e-491c-9143-a5cb1f7e7cea
// )
//
// * EfsRpcQueryRecoveryAgents (section 3.1.4.2.8) ( cf759c00-1b90-4c33-9ace-f51c20149cea
// )
//
// * EfsRpcRemoveUsersFromFile (section 3.1.4.2.9) ( 28609dad-5fa5-4af9-9382-18d40e3e9dec
// )
//
// * EfsRpcAddUsersToFile (section 3.1.4.2.10) ( afd56d24-3732-4477-b5cf-44cc33848d85
// )
//
// * EfsRpcFileKeyInfo (section 3.1.4.2.12) ( 6813bfa8-1538-4c5f-982a-ad58caff3c1c )
//
// * EfsRpcDuplicateEncryptionInfoFile (section 3.1.4.2.13)
//
// * EfsRpcAddUsersToFileEx (section 3.1.4.2.14) ( d36df703-edc9-4482-87b7-d05c7783d65e
// )
//
// * EfsRpcFileKeyInfoEx (section 3.1.4.2.15) ( d0da10ab-3139-4d67-a66c-ea6eb497118d
// )
//
// * EfsRpcGetEncryptedFileMetadata (section 3.1.4.2.16) ( 5cb4e26d-5632-4048-8b2d-f661f8e8ee6b
// )
//
// * EfsRpcSetEncryptedFileMetadata (section 3.1.4.2.17) ( 38a5a7d3-4d47-4075-b830-b657805f2e94
// )
//
// * Performing backup of encrypted objects in ciphertext form along with their EFSRPC
// Metadata, and restoring encrypted objects from such backups. Depending on the implementation
// of these methods, the backups that are created can expose the implementation-specific
// EFSRPC Metadata format to the client. The Windows implementation of these methods
// exposes the Windows EFSRPC Metadata format; however, Windows applications do not
// manipulate this information. The following methods are suitable for secure content
// archival or transferring encrypted data securely between servers of the same implementation
// because they do not require decrypting the data. Methods:
//
// * EfsRpcOpenFileRaw (section 3.1.4.2.1) ( ccc4fb75-1c86-41d7-bbc4-b278ec13bfb8 )
//
// * EfsRpcReadFileRaw (section 3.1.4.2.2) ( a3a6d95f-ebd5-4a08-9d19-6b6c1b7d41f6 )
//
// * EfsRpcWriteFileRaw (section 3.1.4.2.3) ( b4fa10e8-31bb-4dd4-87af-1c4a3dedf417 )
//
// * EfsRpcCloseRaw (section 3.1.4.2.4) ( c5da8aa7-fe38-4940-8d40-2cebbd887cba )
//
// * Controlling the server's encryption subsystem. Methods:
//
// * EfsRpcFlushEfsCache (section 3.1.4.2.18) ( a7e37233-b684-4c3f-905b-81c2e04482f1
// )
//
// Most of the EFSRPC routines are stateless and can be called in any order. When one
// of these routines is called, the message exchange is as follows.
package efsr

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

var (
	// import guard
	GoPackage = "efsr"
)

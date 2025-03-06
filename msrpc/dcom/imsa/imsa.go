// The imsa package implements the IMSA client protocol.
//
// # Introduction
//
// The Internet Information Services (IIS) IMSAdminBaseW Remote Protocol defines interfaces
// that provide Unicode-compliant methods for remotely accessing and administering the
// IIS metabase associated with an application that manages IIS configuration, such
// as the IIS snap-in for Microsoft Management Console (MMC).
//
// # Overview
//
// The Internet Information Services (IIS) IMSAdminBaseW Remote Protocol is a client/server
// protocol that is used for remotely managing a hierarchical configuration data store
// (metabase). The layout and specifics of such a store are specified in section 3.1.1.
//
// The Internet Information Services (IIS) IMSAdminBaseW Remote Protocol also provides
// DCOM interfaces to manage server entities, such as web applications and public key
// certificates, which can be defined or referenced in the metabase data store.
//
// A remote metabase management session begins with the client initiating the connection
// request to the server. If the server grants the request, the connection is established.
// The client can then make multiple requests to read or modify the metabase on the
// server by using the same session until the session is terminated.
//
// A typical remote metabase management session involves the client connecting to the
// server and requesting to open a metabase node on the server. If the server accepts
// the request, it responds with an RPC context handle that refers to the node. The
// client uses this RPC context handle to operate on that node. This involves sending
// another request to the server specifying the type of operation to perform and any
// specific parameters that are associated with that operation. If the server accepts
// this request, it attempts to change the state of the node based on the request and
// responds to the client with the result of the operation. When the client is finished
// operating on the server nodes, it terminates the protocol by sending a request to
// close the RPC context handle.
package imsa

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

// MetadataMasterRootHandle represents the METADATA_MASTER_ROOT_HANDLE RPC constant
var MetadataMasterRootHandle = 0

// AdmindataMaxNameLength represents the ADMINDATA_MAX_NAME_LEN RPC constant
var AdmindataMaxNameLength = 256

// MDBackupMaxLength represents the MD_BACKUP_MAX_LEN RPC constant
var MDBackupMaxLength = 100

// IISCryptoBlob structure represents IIS_CRYPTO_BLOB RPC structure.
//
// The IIS_CRYPTO_BLOB message defines a block of data, possibly encrypted, that is
// transferred between client and server. It is used to transfer public keys, hash information,
// and encrypted and cleartext data.
type IISCryptoBlob struct {
	// BlobSignature:  The structure signature for this binary large object (BLOB).
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                          |                                                                                  |
	//	|                  VALUE                   |                                     MEANING                                      |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| SESSION_KEY_BLOB_SIGNATURE 0x624b6349    | The BlobData member contains the session key used to encrypt sensitive data      |
	//	|                                          | exchanged between client and server. See SESSION_KEY_BLOB (section 2.2.2.2) for  |
	//	|                                          | more information about the BlobData layout.                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| PUBLIC_KEY_BLOB_SIGNATURE 0x62506349     | The BlobData member contains the public key for a particular IIS encryption      |
	//	|                                          | behavior. See PUBLIC_KEY_BLOB (section 2.2.2.1) for more information about the   |
	//	|                                          | BlobData layout.                                                                 |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| ENCRYPTED_DATA_BLOB_SIGNATURE 0x62446349 | The BlobData member contains encrypted data. See ENCRYPTED_DATA_BLOB (section    |
	//	|                                          | 2.2.2.5) for more information about the BlobData layout.                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| HASH_BLOB_SIGNATURE 0x62486349           | The BlobData member contains a hash. See HASH_BLOB (section 2.2.2.3) for more    |
	//	|                                          | information about the BlobData layout.                                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLEARTEXT_DATA_BLOB_SIGNATURE 0x62436349 | The BlobData member contains cleartext data. See CLEARTEXT DATA_BLOB (section    |
	//	|                                          | 2.2.2.4) for more information about the BlobData layout.                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	BlobSignature uint32 `idl:"name:BlobSignature" json:"blob_signature"`
	// BlobDataLength:  The size, in bytes, of BlobData.
	BlobDataLength uint32 `idl:"name:BlobDataLength" json:"blob_data_length"`
	// BlobData:  A block of bytes that can be interpreted based on BlobSignature.
	BlobData []byte `idl:"name:BlobData;size_is:(BlobDataLength)" json:"blob_data"`
}

func (o *IISCryptoBlob) xxx_PreparePayload(ctx context.Context) error {
	if o.BlobData != nil && o.BlobDataLength == 0 {
		o.BlobDataLength = uint32(len(o.BlobData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *IISCryptoBlob) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.BlobDataLength)
	return []uint64{
		dimSize1,
	}
}
func (o *IISCryptoBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.BlobSignature); err != nil {
		return err
	}
	if err := w.WriteData(o.BlobDataLength); err != nil {
		return err
	}
	for i1 := range o.BlobData {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.BlobData[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.BlobData); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *IISCryptoBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlobSignature); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlobDataLength); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.BlobDataLength > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.BlobDataLength)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.BlobData", sizeInfo[0])
	}
	o.BlobData = make([]byte, sizeInfo[0])
	for i1 := range o.BlobData {
		i1 := i1
		if err := w.ReadData(&o.BlobData[i1]); err != nil {
			return err
		}
	}
	return nil
}

// MetadataRecord structure represents METADATA_RECORD RPC structure.
//
// The METADATA_RECORD structure defines information about a metabase entry.
type MetadataRecord struct {
	// dwMDIdentifier:  An unsigned integer value that uniquely identifies the metabase
	// entry.
	ID uint32 `idl:"name:dwMDIdentifier" json:"id"`
	// dwMDAttributes:  An unsigned integer value containing bit flags that specify how
	// to get or set data from the metabase. This member MUST have a valid combination of
	// the following flags set.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_INHERIT 0x00000001       | In Get methods: Returns inheritable data. In Set methods: The data can be        |
	//	|                                   | inherited.                                                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_INSERT_PATH 0x00000040   | For a string data item. In Get methods: Replaces all occurrences of              |
	//	|                                   | "<%INSERT_PATH%>" with the path of the data item relative to the handle. In      |
	//	|                                   | Set methods: Indicate that the string contains the Unicode character substring   |
	//	|                                   | "<%INSERT_PATH%>".                                                               |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_ISINHERITED 0x00000020   | In Get methods: Marks data items that were inherited. In Set methods: Not valid. |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_NO_ATTRIBUTES 0x00000000 | In Get methods: Not applicable. Data is returned regardless of this flag         |
	//	|                                   | setting. In Set methods: The data does not have any attributes.                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_PARTIAL_PATH 0x00000002  | In Get methods: Returns any inherited data even if the entire path is not        |
	//	|                                   | present. This flag is valid only if METADATA_INHERIT is also set. In Set         |
	//	|                                   | methods: Not valid.                                                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_SECURE 0x00000004        | In Get methods: Not valid. In Set methods: Stores and transports the data in a   |
	//	|                                   | secure fashion, as specified in 3.1.4.1.                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_VOLATILE 0x00000010      | In Get methods: Not valid. In Set methods: Does not save the data in long-term   |
	//	|                                   | storage.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	Attributes uint32 `idl:"name:dwMDAttributes" json:"attributes"`
	// dwMDUserType:  An integer value that specifies the user type of the data. The dwMDUserType
	// member MUST be set to one of the following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ASP_MD_UT_APP 0x00000065    | The entry contains information specific to ASP application configuration.        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IIS_MD_UT_FILE 0x00000002   | The entry contains information about a file, such as access permissions or logon |
	//	|                             | methods.                                                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IIS_MD_UT_SERVER 0x00000001 | The entry contains information specific to the server, such as ports in use and  |
	//	|                             | IP addresses.                                                                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IIS_MD_UT_WAM 0x00000064    | The entry contains information specific to WAM.                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	UserType uint32 `idl:"name:dwMDUserType" json:"user_type"`
	// dwMDDataType:  An unsigned integer value that identifies the type of data in the
	// metabase entry. The dwMDDataType member MUST be set to one of the following values.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000      | Specifies all data, regardless of type.                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BINARY_METADATA 0x00000003   | Specifies binary data.                                                           |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DWORD_METADATA 0x00000001    | Specifies all DWORD (unsigned 32-bit integer) data.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| EXPANDSZ_METADATA 0x00000004 | Specifies all data that consists of a string that includes the terminating null  |
	//	|                              | character and which contains environment variables that are not expanded.        |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MULTISZ_METADATA 0x00000005  | Specifies all data represented as an array of strings, where each string         |
	//	|                              | includes the terminating null character, and the array itself is terminated by   |
	//	|                              | two terminating null characters.                                                 |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| STRING_METADATA 0x00000002   | Specifies all data consisting of an ASCII string that includes the terminating   |
	//	|                              | null character.                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DataType uint32 `idl:"name:dwMDDataType" json:"data_type"`
	// dwMDDataLen:  An unsigned integer value that specifies the length of the data in
	// bytes. If the data is a string, this value includes the terminating null character.
	// For lists of strings, this includes an additional terminating null character after
	// the final string (double terminating null characters).
	//
	// For example, the length of a string list containing two strings would be as follows.
	//
	// (wcslen(stringA) + 1) * sizeof(WCHAR) + (wcslen(stringB) + 1) * sizeof(WCHAR) + 1
	// * sizeof(WCHAR)
	DataLength uint32 `idl:"name:dwMDDataLen" json:"data_length"`
	// pbMDData:  When setting data in the metabase, this member contains a pointer to a
	// buffer that holds the data. When getting data from the metabase, this member contains
	// a pointer to a buffer that will receive the data.
	Data []byte `idl:"name:pbMDData;size_is:(dwMDDataLen);pointer:unique" json:"data"`
	// dwMDDataTag:  A reserved member that is currently unused.
	DataTag uint32 `idl:"name:dwMDDataTag" json:"data_tag"`
}

func (o *MetadataRecord) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MetadataRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if err := w.WriteData(o.UserType); err != nil {
		return err
	}
	if err := w.WriteData(o.DataType); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_pbMDData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WritePointer(&o.Data, _ptr_pbMDData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DataTag); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *MetadataRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_pbMDData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
	_s_pbMDData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_pbMDData, _ptr_pbMDData); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataTag); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MetadataGetAllRecord structure represents METADATA_GETALL_RECORD RPC structure.
//
// The METADATA_GETALL_RECORD structure defines an analogous structure to METADATA_RECORD
// but is used only to return data from a call to the R_GetAllData method. Data retrieval
// specifications are provided in R_GetAllData method parameters, not in this structure
// (as is the case with METADATA_RECORD). The R_GetAllData method returns the data from
// multiple entries as an array of METADATA_GETALL_RECORD structures.
type MetadataGetAllRecord struct {
	// dwMDIdentifier:  An unsigned integer value that uniquely identifies the metabase
	// entry.
	ID uint32 `idl:"name:dwMDIdentifier" json:"id"`
	// dwMDAttributes:  An unsigned integer value containing bit flags that specify how
	// to set or get data from the metabase. This member MUST be set to a valid combination
	// of the following values.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_INHERIT 0x00000001       | In Get methods: Return the inheritable data. In Set methods: The data can be     |
	//	|                                   | inherited.                                                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_INSERT_PATH 0x00000040   | For a string data item. In Get methods: Replace all occurrences of               |
	//	|                                   | "<%INSERT_PATH%>" with the path of the data item relative to the handle. In      |
	//	|                                   | Set methods: Indicate that the string contains the Unicode character substring   |
	//	|                                   | "<%INSERT_PATH%>".                                                               |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_ISINHERITED 0x00000020   | In Get methods: Mark the data items that were inherited. In Set methods: Not     |
	//	|                                   | valid.                                                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_NO_ATTRIBUTES 0x00000000 | In Get methods: Not applicable. Data is returned regardless of this flag         |
	//	|                                   | setting. In Set methods: The data does not have any attributes.                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_PARTIAL_PATH 0x00000002  | In Get methods: Return any inherited data even if the entire path is not         |
	//	|                                   | present. This flag is valid only if METADATA_INHERIT is also set. In Set         |
	//	|                                   | methods: Not valid.                                                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_SECURE 0x00000004        | In Get methods: Not valid. In Set methods: The server and client transport and   |
	//	|                                   | store the data in a secure fashion, as specified in 3.1.4.1.1.                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| METADATA_VOLATILE 0x00000010      | In Get methods: Not valid. In Set methods: Do not save the data in long-term     |
	//	|                                   | storage.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	Attributes uint32 `idl:"name:dwMDAttributes" json:"attributes"`
	// dwMDUserType:  An unsigned integer value that specifies the user type of the data.
	// The dwMDUserType member MUST be set to one of the following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ASP_MD_UT_APP 0x00000065    | The entry contains information specific to ASP application configuration.        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IIS_MD_UT_FILE 0x00000002   | The entry contains information about a file, such as access permissions or logon |
	//	|                             | methods.                                                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IIS_MD_UT_SERVER 0x00000001 | The entry contains information specific to the server, such as ports in use and  |
	//	|                             | IP addresses.                                                                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IIS_MD_UT_WAM 0x00000064    | The entry contains information specific to web application management.           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	UserType uint32 `idl:"name:dwMDUserType" json:"user_type"`
	// dwMDDataType:  An integer value that identifies the type of data in the metabase
	// entry. The dwMDDataType member MUST be set to one of the following values.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000      | Specifies all data, regardless of type.                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BINARY_METADATA 0x00000003   | Specifies binary data in any form.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DWORD_METADATA 0x00000001    | Specifies all DWORD (unsigned 32-bit integer) data.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| EXPANDSZ_METADATA 0x00000004 | Specifies all data that consists of a string that includes the terminating null  |
	//	|                              | character, and which contains environment variables that are not expanded.       |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MULTISZ_METADATA 0x00000005  | Specifies all data represented as an array of strings, where each string         |
	//	|                              | includes the terminating null character, and the array itself is terminated by   |
	//	|                              | two terminating null characters.                                                 |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| STRING_METADATA 0x00000002   | Specifies all data consisting of an ASCII string that includes the terminating   |
	//	|                              | null character.                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DataType uint32 `idl:"name:dwMDDataType" json:"data_type"`
	// dwMDDataLen:   An unsigned integer value that specifies the length, in bytes, of
	// the data. If the data is a string, this value includes the ending null character.
	// For lists of strings, this includes an additional terminating null character after
	// the final string (double terminating null characters).
	//
	// For example, the length of a string list containing two strings would be as follows.
	//
	// (wcslen(stringA) + 1) * sizeof(WCHAR) + (wcslen(stringB) + 1) * sizeof(WCHAR) + 1
	// * sizeof(WCHAR)
	DataLength uint32 `idl:"name:dwMDDataLen" json:"data_length"`
	// dwMDDataOffset:  If the data was returned by value, this member contains the byte
	// offset of the data in the buffer specified by the pbMDBuffer parameter of the R_GetAllData
	// method. All out-of-process executions will return data by value. The array of records,
	// excluding the data, is returned in the first part of the buffer. The data associated
	// with the records is returned in the buffer after the array of records, and dwMDDataOffset
	// is the offset to the beginning of the data associated with each record in the array.
	DataOffset uint32 `idl:"name:dwMDDataOffset" json:"data_offset"`
	// dwMDDataTag:  A reserved member that is currently unused.
	DataTag uint32 `idl:"name:dwMDDataTag" json:"data_tag"`
}

func (o *MetadataGetAllRecord) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MetadataGetAllRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if err := w.WriteData(o.UserType); err != nil {
		return err
	}
	if err := w.WriteData(o.DataType); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DataOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.DataTag); err != nil {
		return err
	}
	return nil
}
func (o *MetadataGetAllRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataTag); err != nil {
		return err
	}
	return nil
}

// MetadataHandleInfo structure represents METADATA_HANDLE_INFO RPC structure.
//
// The METADATA_HANDLE_INFO structure defines information about a handle to a metabase
// entry.
type MetadataHandleInfo struct {
	// dwMDPermissions:   An unsigned integer value containing the permissions with which
	// the handle was opened. This member MUST have a valid combination of the following
	// flags set.
	//
	//	+--------------------------------------+--------------------------------------+
	//	|                                      |                                      |
	//	|                VALUE                 |               MEANING                |
	//	|                                      |                                      |
	//	+--------------------------------------+--------------------------------------+
	//	+--------------------------------------+--------------------------------------+
	//	| METADATA_PERMISSION_READ 0x00000001  | The handle can read nodes and data.  |
	//	+--------------------------------------+--------------------------------------+
	//	| METADATA_PERMISSION_WRITE 0x00000002 | The handle can write nodes and data. |
	//	+--------------------------------------+--------------------------------------+
	Permissions uint32 `idl:"name:dwMDPermissions" json:"permissions"`
	// dwMDSystemChangeNumber:  An unsigned integer value containing the system change number
	// when the handle was opened. The system change number is a 32-bit unsigned integer
	// value that is incremented when a change is made to the metabase. See GetSystemChangeNumber
	// (section 3.1.4.21) for a specification of the system change number.
	SystemChangeNumber uint32 `idl:"name:dwMDSystemChangeNumber" json:"system_change_number"`
}

func (o *MetadataHandleInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MetadataHandleInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Permissions); err != nil {
		return err
	}
	if err := w.WriteData(o.SystemChangeNumber); err != nil {
		return err
	}
	return nil
}
func (o *MetadataHandleInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Permissions); err != nil {
		return err
	}
	if err := w.ReadData(&o.SystemChangeNumber); err != nil {
		return err
	}
	return nil
}

// WAMAdmin2 structure represents IWamAdmin2 RPC structure.
type WAMAdmin2 dcom.InterfacePointer

func (o *WAMAdmin2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *WAMAdmin2) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *WAMAdmin2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *WAMAdmin2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *WAMAdmin2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// IMSAdminBaseW structure represents IMSAdminBaseW RPC structure.
type IMSAdminBaseW dcom.InterfacePointer

func (o *IMSAdminBaseW) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *IMSAdminBaseW) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *IMSAdminBaseW) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IMSAdminBaseW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *IMSAdminBaseW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// IISCertObject structure represents IIISCertObj RPC structure.
type IISCertObject dcom.InterfacePointer

func (o *IISCertObject) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *IISCertObject) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *IISCertObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IISCertObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *IISCertObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// IISApplicationAdmin structure represents IIISApplicationAdmin RPC structure.
type IISApplicationAdmin dcom.InterfacePointer

func (o *IISApplicationAdmin) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *IISApplicationAdmin) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *IISApplicationAdmin) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IISApplicationAdmin) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *IISApplicationAdmin) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// IMSAdminBase2W structure represents IMSAdminBase2W RPC structure.
type IMSAdminBase2W dcom.InterfacePointer

func (o *IMSAdminBase2W) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *IMSAdminBase2W) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *IMSAdminBase2W) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IMSAdminBase2W) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *IMSAdminBase2W) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// IMSAdminBase3W structure represents IMSAdminBase3W RPC structure.
type IMSAdminBase3W dcom.InterfacePointer

func (o *IMSAdminBase3W) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *IMSAdminBase3W) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *IMSAdminBase3W) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IMSAdminBase3W) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *IMSAdminBase3W) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

// WAMAdmin structure represents IWamAdmin RPC structure.
type WAMAdmin dcom.InterfacePointer

func (o *WAMAdmin) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *WAMAdmin) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *WAMAdmin) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *WAMAdmin) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
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
}
func (o *WAMAdmin) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
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
}

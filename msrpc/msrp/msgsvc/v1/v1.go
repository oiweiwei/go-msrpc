package msgsvc

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
	GoPackage = "msrp"
)

var (
	// Syntax UUID
	MsgsvcSyntaxUUID = &uuid.UUID{TimeLow: 0x17fdd703, TimeMid: 0x1827, TimeHiAndVersion: 0x4e34, ClockSeqHiAndReserved: 0x79, ClockSeqLow: 0xd4, Node: [6]uint8{0x24, 0xa5, 0x5c, 0x53, 0xbb, 0x37}}
	// Syntax ID
	MsgsvcSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: MsgsvcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// msgsvc interface.
type MsgsvcClient interface {

	// The NetrMessageNameAdd (Opnum 0) interface is used to configure the message server
	// to listen for messages sent to an additional NetBIOS name.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	|             RETURN             |                                                                         |
	//	|           VALUE/CODE           |                               DESCRIPTION                               |
	//	|                                |                                                                         |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The operation completed successfully.                                   |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                       |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME  | The file name, directory name, or volume label syntax is incorrect.     |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000859 NERR_NetworkError   | A general network error occurred.                                       |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x0000085C NERR_InternalError  | An internal error occurred.                                             |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x000008E4 NERR_AlreadyExists  | This message alias already exists locally.                              |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x000008E5 NERR_TooManyNames   | The maximum number of added message aliases has been exceeded.          |
	//	+--------------------------------+-------------------------------------------------------------------------+
	//	| 0x000008F9 NERR_DuplicateName  | The name specified is already in use as a message alias on the network. |
	//	+--------------------------------+-------------------------------------------------------------------------+
	MessageNameAdd(context.Context, *MessageNameAddRequest, ...dcerpc.CallOption) (*MessageNameAddResponse, error)

	// The NetrMessageNameEnum (Opnum 1) interface is used to enumerate the NetBIOS names
	// for which the message server is currently listening for messages.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+------------------------------------+---------------------------------------+
	//	|               RETURN               |                                       |
	//	|             VALUE/CODE             |              DESCRIPTION              |
	//	|                                    |                                       |
	//	+------------------------------------+---------------------------------------+
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000000 NERR_Success            | The operation completed successfully. |
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                     |
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.           |
	//	+------------------------------------+---------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct. |
	//	+------------------------------------+---------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall        | The API return buffer is too small.   |
	//	+------------------------------------+---------------------------------------+
	MessageNameEnum(context.Context, *MessageNameEnumRequest, ...dcerpc.CallOption) (*MessageNameEnumResponse, error)

	// The NetrMessageNameGetInfo (Opnum 2) interface is used to retrieve information from
	// the message server on a NetBIOS name for which the message server is currently listening
	// for messages.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN               |                                                                     |
	//	|             VALUE/CODE             |                             DESCRIPTION                             |
	//	|                                    |                                                                     |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The operation completed successfully.                               |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                   |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.            |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME      | The file name, directory name, or volume label syntax is incorrect. |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                               |
	//	+------------------------------------+---------------------------------------------------------------------+
	//	| 0x000008ED NERR_NotLocalName       | The name is not on the local computer.                              |
	//	+------------------------------------+---------------------------------------------------------------------+
	MessageNameGetInfo(context.Context, *MessageNameGetInfoRequest, ...dcerpc.CallOption) (*MessageNameGetInfoResponse, error)

	// The NetrMessageNameDel (Opnum 3) interface is used to configure the message server
	// to stop listening for messages for a particular NetBIOS name.
	//
	// Return Values: A NET_API_STATUS value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+---------------------------------+---------------------------------------------------------------------+
	//	|             RETURN              |                                                                     |
	//	|           VALUE/CODE            |                             DESCRIPTION                             |
	//	|                                 |                                                                     |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success         | The operation completed successfully.                               |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED  | Access is denied.                                                   |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME   | The file name, directory name, or volume label syntax is incorrect. |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008E6 NERR_DelComputerName | The computer name could not be deleted.                             |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008EB NERR_NameInUse       | The message alias is currently in use. Try again later.             |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008ED NERR_NotLocalName    | The name is not on the local computer.                              |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x000008FB NERR_IncompleteDel   | The message alias was not successfully deleted from all networks.   |
	//	+---------------------------------+---------------------------------------------------------------------+
	MessageNameDelete(context.Context, *MessageNameDeleteRequest, ...dcerpc.CallOption) (*MessageNameDeleteResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// MessageInfo0 structure represents MSG_INFO_0 RPC structure.
//
// MSG_INFO_0 is a data structure that contains a string that specifies the recipient
// name to which a message is to be sent.
type MessageInfo0 struct {
	// msgi0_name:   Pointer to a buffer that receives the name string in UTF-16 format.
	// There are two sources for this parameter:
	//
	// *
	//
	// It is the UTF-16 formatted Name parameter passed in NetrMessageNameGetInfo ( e5845829-414e-43c3-a956-e9e8dd4f6fd3
	// ) (section 3.1.4.3) that has been verified to have an equivalent entry in the message
	// table in section 3.1.1 ( 964e52a9-ddc0-4473-9199-200a417c50ed ) according to the
	// following algorithm.
	//
	//  Function CompareName (passed in Unicode name)
	//    Convert the name to uppercase as specified in [MS-UCODEREF] section 3.1.5.3
	//    Convert the Unicode string to an OEM string as specified in [MS-UCODEREF] section
	// 3.1.5.1.1
	//    If OEM string is < 15 bytes
	//      Pad to 15 bytes with spaces
	//    If OEM string is > 15 bytes
	//      Truncate to 15 bytes
	//    For each table entry
	//      Compare the resulting 15 bytes for equality with the first 15 bytes of the
	// NetBIOS name in the entry
	//      If equal
	//        Return True
	//  End CompareName
	//
	// *
	//
	// It is returned in the InfoStruct parameter of NetrMessageNameEnum ( 14600397-2f79-4704-a902-e4101bbbc0c5
	// ) (section 3.1.4.2) in which it was retrieved from the message table in section 3.1.1,
	// the NetBIOS suffix ( 1bb1da94-fc7e-484c-ad2a-4e0d257a50ce#gt_e0056bac-be0a-488d-86b2-eec9c6bfb947
	// ) and any trailing spaces removed, and the remaining characters converted to UTF-16.
	Name string `idl:"name:msgi0_name;string" json:"name"`
}

func (o *MessageInfo0) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_msgi0_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_msgi0_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_msgi0_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_msgi0_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_msgi0_name, _ptr_msgi0_name); err != nil {
		return err
	}
	return nil
}

// MessageInfo1 structure represents MSG_INFO_1 RPC structure.
//
// MSG_INFO_1 is a data structure that contains a string that specifies the recipient
// name to which a message is to be sent.
type MessageInfo1 struct {
	// msgi1_name:   Pointer to a buffer that receives the name string in UTF-16 format.
	// There are two sources for this parameter:
	//
	// *
	//
	// It is the UTF-16 formatted Name parameter passed in NetrMessageNameGetInfo ( e5845829-414e-43c3-a956-e9e8dd4f6fd3
	// ) (section 3.1.4.3) that has been verified to have an equivalent entry in the message
	// table in section 3.1.1 ( 964e52a9-ddc0-4473-9199-200a417c50ed ) according to the
	// following algorithm.
	//
	//  Function CompareName (passed in Unicode name)
	//    Convert the name to uppercase as specified in [MS-UCODEREF] section 3.1.5.3
	//    Convert the Unicode string to an OEM string as specified in [MS-UCODEREF] section
	// 3.1.5.1.1
	//    If OEM string is < 15 bytes
	//      Pad to 15 bytes with spaces
	//    If OEM string is > 15 bytes
	//      Truncate to 15 bytes
	//    For each table entry
	//      Compare the resulting 15 bytes for equality with the first 15 bytes of the
	// NetBIOS name in the entry
	//      If equal
	//        Return True
	//  End CompareName
	//
	// *
	//
	// It is returned in the InfoStruct parameter of NetrMessageNameEnum ( 14600397-2f79-4704-a902-e4101bbbc0c5
	// ) (section 3.1.4.2) in which it was retrieved from the message table in section 3.1.1,
	// the NetBIOS suffix ( 1bb1da94-fc7e-484c-ad2a-4e0d257a50ce#gt_e0056bac-be0a-488d-86b2-eec9c6bfb947
	// ) and any trailing spaces removed, and the remaining characters converted to UTF-16.
	Name string `idl:"name:msgi1_name;string" json:"name"`
	// msgi1_forward_flag:  MUST be set to zero when sent and ignored on receipt.
	ForwardFlag uint32 `idl:"name:msgi1_forward_flag" json:"forward_flag"`
	// msgi1_forward:  MUST be NULL and ignored on receipt.
	Forward string `idl:"name:msgi1_forward;string" json:"forward"`
}

func (o *MessageInfo1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_msgi1_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_msgi1_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ForwardFlag); err != nil {
		return err
	}
	if o.Forward != "" {
		_ptr_msgi1_forward := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Forward); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Forward, _ptr_msgi1_forward); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_msgi1_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_msgi1_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_msgi1_name, _ptr_msgi1_name); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardFlag); err != nil {
		return err
	}
	_ptr_msgi1_forward := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Forward); err != nil {
			return err
		}
		return nil
	})
	_s_msgi1_forward := func(ptr interface{}) { o.Forward = *ptr.(*string) }
	if err := w.ReadPointer(&o.Forward, _s_msgi1_forward, _ptr_msgi1_forward); err != nil {
		return err
	}
	return nil
}

// MessageInfo0Container structure represents MSG_INFO_0_CONTAINER RPC structure.
//
// MSG_INFO_0_CONTAINER is a container structure that holds one or more MSG_INFO_0 structures.
type MessageInfo0Container struct {
	// EntriesRead:  A 32-bit value that MUST denote the number of entries in Buffer.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  Pointer to a buffer that MUST contain one or more MSG_INFO_0 structures.
	Buffer []*MessageInfo0 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *MessageInfo0Container) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo0Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&MessageInfo0{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&MessageInfo0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo0Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*MessageInfo0, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &MessageInfo0{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*MessageInfo0) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// MessageInfo1Container structure represents MSG_INFO_1_CONTAINER RPC structure.
//
// MSG_INFO_1_CONTAINER is a container structure that holds one or more MSG_INFO_1 structures.
type MessageInfo1Container struct {
	// EntriesRead:  A 32-bit value that MUST denote the number of entries in Buffer.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  A pointer to a variable-size buffer that MUST contain one or more MSG_INFO_1
	// structures.
	Buffer []*MessageInfo1 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *MessageInfo1Container) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo1Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&MessageInfo1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&MessageInfo1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo1Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*MessageInfo1, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &MessageInfo1{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*MessageInfo1) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// MessageEnum structure represents MSG_ENUM_STRUCT RPC structure.
//
// MSG_ENUM_STRUCT is a container structure holding either one MSG_INFO_0_CONTAINER
// container or one MSG_INFO_1_CONTAINER container. The structure also has a member
// to indicate what type of container it contains.
type MessageEnum struct {
	// Level:   A 32-bit enumerated number that MUST denote the type of structure contained
	// in MsgInfo. It must be either 0 or 1.
	Level uint32 `idl:"name:Level" json:"level"`
	// MsgInfo:  A pointer to a buffer that MUST contain a union that consists of either
	// an MSG_INFO_0_CONTAINER structure or an MSG_INFO_1_CONTAINER structure.
	MessageInfo *MessageEnum_MessageInfo `idl:"name:MsgInfo;switch_is:Level" json:"message_info"`
}

func (o *MessageEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swMessageInfo := uint32(o.Level)
	if o.MessageInfo != nil {
		if err := o.MessageInfo.MarshalUnionNDR(ctx, w, _swMessageInfo); err != nil {
			return err
		}
	} else {
		if err := (&MessageEnum_MessageInfo{}).MarshalUnionNDR(ctx, w, _swMessageInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.MessageInfo == nil {
		o.MessageInfo = &MessageEnum_MessageInfo{}
	}
	_swMessageInfo := uint32(o.Level)
	if err := o.MessageInfo.UnmarshalUnionNDR(ctx, w, _swMessageInfo); err != nil {
		return err
	}
	return nil
}

// MessageEnum_MessageInfo structure represents MSG_ENUM_STRUCT union anonymous member.
//
// MSG_ENUM_STRUCT is a container structure holding either one MSG_INFO_0_CONTAINER
// container or one MSG_INFO_1_CONTAINER container. The structure also has a member
// to indicate what type of container it contains.
type MessageEnum_MessageInfo struct {
	// Types that are assignable to Value
	//
	// *MessageInfo_Level0
	// *MessageInfo_Level1
	Value is_MessageEnum_MessageInfo `json:"value"`
}

func (o *MessageEnum_MessageInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MessageInfo_Level0:
		if value != nil {
			return value.Level0
		}
	case *MessageInfo_Level1:
		if value != nil {
			return value.Level1
		}
	}
	return nil
}

type is_MessageEnum_MessageInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MessageEnum_MessageInfo()
}

func (o *MessageEnum_MessageInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *MessageInfo_Level0:
		return uint32(0)
	case *MessageInfo_Level1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *MessageEnum_MessageInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*MessageInfo_Level0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageInfo_Level0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*MessageInfo_Level1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageInfo_Level1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *MessageEnum_MessageInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &MessageInfo_Level0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &MessageInfo_Level1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// MessageInfo_Level0 structure represents MessageEnum_MessageInfo RPC union arm.
//
// It has following labels: 0
type MessageInfo_Level0 struct {
	// Level0:  If Level is 0, MsgInfo MUST contain an MSG_INFO_0_CONTAINER named Level0.
	Level0 *MessageInfo0Container `idl:"name:Level0" json:"level0"`
}

func (*MessageInfo_Level0) is_MessageEnum_MessageInfo() {}

func (o *MessageInfo_Level0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level0 != nil {
		_ptr_Level0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level0 != nil {
				if err := o.Level0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&MessageInfo0Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level0, _ptr_Level0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo_Level0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level0 == nil {
			o.Level0 = &MessageInfo0Container{}
		}
		if err := o.Level0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level0 := func(ptr interface{}) { o.Level0 = *ptr.(**MessageInfo0Container) }
	if err := w.ReadPointer(&o.Level0, _s_Level0, _ptr_Level0); err != nil {
		return err
	}
	return nil
}

// MessageInfo_Level1 structure represents MessageEnum_MessageInfo RPC union arm.
//
// It has following labels: 1
type MessageInfo_Level1 struct {
	// Level1:  If Level is 1, MsgInfo MUST contain an MSG_INFO_1_CONTAINER named Level1.
	Level1 *MessageInfo1Container `idl:"name:Level1" json:"level1"`
}

func (*MessageInfo_Level1) is_MessageEnum_MessageInfo() {}

func (o *MessageInfo_Level1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level1 != nil {
		_ptr_Level1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level1 != nil {
				if err := o.Level1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&MessageInfo1Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level1, _ptr_Level1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo_Level1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level1 == nil {
			o.Level1 = &MessageInfo1Container{}
		}
		if err := o.Level1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level1 := func(ptr interface{}) { o.Level1 = *ptr.(**MessageInfo1Container) }
	if err := w.ReadPointer(&o.Level1, _s_Level1, _ptr_Level1); err != nil {
		return err
	}
	return nil
}

// MessageInfo structure represents MSG_INFO RPC union.
//
// MSG_INFO is a data structure that contains either an MSG_INFO_0 or an MSG_INFO_1
// structure.
type MessageInfo struct {
	// Types that are assignable to Value
	//
	// *MessageInfo_0
	// *MessageInfo_1
	Value is_MessageInfo `json:"value"`
}

func (o *MessageInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MessageInfo_0:
		if value != nil {
			return value.MessageInfo0
		}
	case *MessageInfo_1:
		if value != nil {
			return value.MessageInfo1
		}
	}
	return nil
}

type is_MessageInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MessageInfo()
}

func (o *MessageInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *MessageInfo_0:
		return uint32(0)
	case *MessageInfo_1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *MessageInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*MessageInfo_0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageInfo_0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*MessageInfo_1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageInfo_1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *MessageInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &MessageInfo_0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &MessageInfo_1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// MessageInfo_0 structure represents MSG_INFO RPC union arm.
//
// It has following labels: 0
type MessageInfo_0 struct {
	// MsgInfo0:  A pointer to a variable-size buffer that MUST contain an MSG_INFO_0 data
	// structure.
	MessageInfo0 *MessageInfo0 `idl:"name:MsgInfo0" json:"message_info0"`
}

func (*MessageInfo_0) is_MessageInfo() {}

func (o *MessageInfo_0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MessageInfo0 != nil {
		_ptr_MsgInfo0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.MessageInfo0 != nil {
				if err := o.MessageInfo0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&MessageInfo0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MessageInfo0, _ptr_MsgInfo0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo_0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_MsgInfo0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.MessageInfo0 == nil {
			o.MessageInfo0 = &MessageInfo0{}
		}
		if err := o.MessageInfo0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_MsgInfo0 := func(ptr interface{}) { o.MessageInfo0 = *ptr.(**MessageInfo0) }
	if err := w.ReadPointer(&o.MessageInfo0, _s_MsgInfo0, _ptr_MsgInfo0); err != nil {
		return err
	}
	return nil
}

// MessageInfo_1 structure represents MSG_INFO RPC union arm.
//
// It has following labels: 1
type MessageInfo_1 struct {
	// MsgInfo1:  A pointer to a variable-size buffer that MUST contain an MSG_INFO_1 data
	// structure.
	MessageInfo1 *MessageInfo1 `idl:"name:MsgInfo1" json:"message_info1"`
}

func (*MessageInfo_1) is_MessageInfo() {}

func (o *MessageInfo_1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MessageInfo1 != nil {
		_ptr_MsgInfo1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.MessageInfo1 != nil {
				if err := o.MessageInfo1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&MessageInfo1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MessageInfo1, _ptr_MsgInfo1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageInfo_1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_MsgInfo1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.MessageInfo1 == nil {
			o.MessageInfo1 = &MessageInfo1{}
		}
		if err := o.MessageInfo1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_MsgInfo1 := func(ptr interface{}) { o.MessageInfo1 = *ptr.(**MessageInfo1) }
	if err := w.ReadPointer(&o.MessageInfo1, _s_MsgInfo1, _ptr_MsgInfo1); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultMsgsvcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultMsgsvcClient) MessageNameAdd(ctx context.Context, in *MessageNameAddRequest, opts ...dcerpc.CallOption) (*MessageNameAddResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MessageNameAddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMsgsvcClient) MessageNameEnum(ctx context.Context, in *MessageNameEnumRequest, opts ...dcerpc.CallOption) (*MessageNameEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MessageNameEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMsgsvcClient) MessageNameGetInfo(ctx context.Context, in *MessageNameGetInfoRequest, opts ...dcerpc.CallOption) (*MessageNameGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MessageNameGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMsgsvcClient) MessageNameDelete(ctx context.Context, in *MessageNameDeleteRequest, opts ...dcerpc.CallOption) (*MessageNameDeleteResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MessageNameDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMsgsvcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMsgsvcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewMsgsvcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (MsgsvcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(MsgsvcSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultMsgsvcClient{cc: cc}, nil
}

// xxx_MessageNameAddOperation structure represents the NetrMessageNameAdd operation
type xxx_MessageNameAddOperation struct {
	ServerName  string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	MessageName string `idl:"name:MsgName;string" json:"message_name"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_MessageNameAddOperation) OpNum() int { return 0 }

func (o *xxx_MessageNameAddOperation) OpName() string { return "/msgsvc/v1/NetrMessageNameAdd" }

func (o *xxx_MessageNameAddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameAddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// MsgName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.MessageName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameAddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MsgName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.MessageName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameAddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameAddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameAddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MessageNameAddRequest structure represents the NetrMessageNameAdd operation request
type MessageNameAddRequest struct {
	// ServerName: A pointer to a null-terminated string that MUST denote the NetBIOS name
	// (as specified in [RFC1001] section 5.2) or the fully qualified domain name (FQDN)
	// of the remote computer on which the function is to execute. There are no other constraints
	// on the format of this string. The message server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// MsgName: A null-terminated Unicode string that MUST denote the recipient name to
	// add. The name is not guaranteed to be unique or reachable by this method. The string
	// MUST be represented using Unicode UTF-16.
	MessageName string `idl:"name:MsgName;string" json:"message_name"`
}

func (o *MessageNameAddRequest) xxx_ToOp(ctx context.Context) *xxx_MessageNameAddOperation {
	if o == nil {
		return &xxx_MessageNameAddOperation{}
	}
	return &xxx_MessageNameAddOperation{
		ServerName:  o.ServerName,
		MessageName: o.MessageName,
	}
}

func (o *MessageNameAddRequest) xxx_FromOp(ctx context.Context, op *xxx_MessageNameAddOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.MessageName = op.MessageName
}
func (o *MessageNameAddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MessageNameAddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameAddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MessageNameAddResponse structure represents the NetrMessageNameAdd operation response
type MessageNameAddResponse struct {
	// Return: The NetrMessageNameAdd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MessageNameAddResponse) xxx_ToOp(ctx context.Context) *xxx_MessageNameAddOperation {
	if o == nil {
		return &xxx_MessageNameAddOperation{}
	}
	return &xxx_MessageNameAddOperation{
		Return: o.Return,
	}
}

func (o *MessageNameAddResponse) xxx_FromOp(ctx context.Context, op *xxx_MessageNameAddOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MessageNameAddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MessageNameAddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameAddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MessageNameEnumOperation structure represents the NetrMessageNameEnum operation
type xxx_MessageNameEnumOperation struct {
	ServerName    string       `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Info          *MessageEnum `idl:"name:InfoStruct" json:"info"`
	PrefMaxLength uint32       `idl:"name:PrefMaxLen" json:"pref_max_length"`
	TotalEntries  uint32       `idl:"name:TotalEntries" json:"total_entries"`
	Resume        uint32       `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	Return        uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_MessageNameEnumOperation) OpNum() int { return 1 }

func (o *xxx_MessageNameEnumOperation) OpName() string { return "/msgsvc/v1/NetrMessageNameEnum" }

func (o *xxx_MessageNameEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// InfoStruct {in, out} (1:{alias=LPMSG_ENUM_STRUCT}*(1))(2:{alias=MSG_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// PrefMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PrefMaxLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// InfoStruct {in, out} (1:{alias=LPMSG_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=MSG_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &MessageEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PrefMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PrefMaxLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {in, out} (1:{alias=LPMSG_ENUM_STRUCT}*(1))(2:{alias=MSG_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MessageEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {in, out} (1:{alias=LPMSG_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=MSG_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &MessageEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MessageNameEnumRequest structure represents the NetrMessageNameEnum operation request
type MessageNameEnumRequest struct {
	// ServerName: A pointer to a null-terminated string that MUST denote the NetBIOS name
	// (as specified in [RFC1001] section 5.2) or the fully qualified domain name (FQDN)
	// of the remote computer on which the function is to execute. There are no other constraints
	// on the format of this string. The message server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// InfoStruct: A pointer to a buffer that receives a variable-length data structure
	// of type MSG_ENUM_STRUCT. The buffer MUST be allocated, and the pointer MUST be assigned
	// by the message server. On return, the structure MUST contain the list of names for
	// which the message server is listening for messages.
	Info *MessageEnum `idl:"name:InfoStruct" json:"info"`
	// PrefMaxLen: A 32-bit number that MUST denote the maximum number of bytes the message
	// server allocates for the buffer. If PrefMaxLen is set to 0xFFFFFFFF, the message
	// server MUST always allocate a buffer that can hold all of the information available
	// in a single MSG_ENUM_STRUCT.
	PrefMaxLength uint32 `idl:"name:PrefMaxLen" json:"pref_max_length"`
	// ResumeHandle: A pointer to a 32-bit number that MUST contain the ordinal value of
	// the name, in the message server's internal list, on which to start enumeration. This
	// MAY be null.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
}

func (o *MessageNameEnumRequest) xxx_ToOp(ctx context.Context) *xxx_MessageNameEnumOperation {
	if o == nil {
		return &xxx_MessageNameEnumOperation{}
	}
	return &xxx_MessageNameEnumOperation{
		ServerName:    o.ServerName,
		Info:          o.Info,
		PrefMaxLength: o.PrefMaxLength,
		Resume:        o.Resume,
	}
}

func (o *MessageNameEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_MessageNameEnumOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Info = op.Info
	o.PrefMaxLength = op.PrefMaxLength
	o.Resume = op.Resume
}
func (o *MessageNameEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MessageNameEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MessageNameEnumResponse structure represents the NetrMessageNameEnum operation response
type MessageNameEnumResponse struct {
	// InfoStruct: A pointer to a buffer that receives a variable-length data structure
	// of type MSG_ENUM_STRUCT. The buffer MUST be allocated, and the pointer MUST be assigned
	// by the message server. On return, the structure MUST contain the list of names for
	// which the message server is listening for messages.
	Info *MessageEnum `idl:"name:InfoStruct" json:"info"`
	// TotalEntries: A pointer to a 32-bit number that, on return, MUST contain the total
	// number of entries in InfoStruct.
	TotalEntries uint32 `idl:"name:TotalEntries" json:"total_entries"`
	// ResumeHandle: A pointer to a 32-bit number that MUST contain the ordinal value of
	// the name, in the message server's internal list, on which to start enumeration. This
	// MAY be null.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	// Return: The NetrMessageNameEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MessageNameEnumResponse) xxx_ToOp(ctx context.Context) *xxx_MessageNameEnumOperation {
	if o == nil {
		return &xxx_MessageNameEnumOperation{}
	}
	return &xxx_MessageNameEnumOperation{
		Info:         o.Info,
		TotalEntries: o.TotalEntries,
		Resume:       o.Resume,
		Return:       o.Return,
	}
}

func (o *MessageNameEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_MessageNameEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *MessageNameEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MessageNameEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MessageNameGetInfoOperation structure represents the NetrMessageNameGetInfo operation
type xxx_MessageNameGetInfoOperation struct {
	ServerName  string       `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	MessageName string       `idl:"name:MsgName;string" json:"message_name"`
	Level       uint32       `idl:"name:Level" json:"level"`
	Info        *MessageInfo `idl:"name:InfoStruct;switch_is:Level" json:"info"`
	Return      uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_MessageNameGetInfoOperation) OpNum() int { return 2 }

func (o *xxx_MessageNameGetInfoOperation) OpName() string { return "/msgsvc/v1/NetrMessageNameGetInfo" }

func (o *xxx_MessageNameGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// MsgName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.MessageName); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MsgName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.MessageName); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {out} (1:{switch_type={alias=DWORD}(uint32), alias=LPMSG_INFO}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=MSG_INFO}(union))
	{
		_swInfo := uint32(o.Level)
		if o.Info != nil {
			if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		} else {
			if err := (&MessageInfo{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {out} (1:{switch_type={alias=DWORD}(uint32), alias=LPMSG_INFO,pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=MSG_INFO}(union))
	{
		if o.Info == nil {
			o.Info = &MessageInfo{}
		}
		_swInfo := uint32(o.Level)
		if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MessageNameGetInfoRequest structure represents the NetrMessageNameGetInfo operation request
type MessageNameGetInfoRequest struct {
	// ServerName: A pointer to a null-terminated string that MUST denote the NetBIOS name
	// (as specified in [RFC1001] section 5.2) or the fully qualified domain name (FQDN)
	// of the remote computer on which the function is to execute. There are no other constraints
	// on the format of this string. The message server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// MsgName:  A null-terminated Unicode UTF-16 string. It MUST denote the recipient
	// name for which to get information. The name is not guaranteed to exist.
	MessageName string `idl:"name:MsgName;string" json:"message_name"`
	// Level:  A 32-bit number that MUST be either 0 or 1. It represents the type of structure
	// contained in the InfoStruct MSG_INFO structure. If Level is 0, InfoStruct MUST contain
	// an MSG_INFO_0 data structure. If Level is 1, InfoStruct MUST contain an MSG_INFO_1
	// data structure.
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *MessageNameGetInfoRequest) xxx_ToOp(ctx context.Context) *xxx_MessageNameGetInfoOperation {
	if o == nil {
		return &xxx_MessageNameGetInfoOperation{}
	}
	return &xxx_MessageNameGetInfoOperation{
		ServerName:  o.ServerName,
		MessageName: o.MessageName,
		Level:       o.Level,
	}
}

func (o *MessageNameGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_MessageNameGetInfoOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.MessageName = op.MessageName
	o.Level = op.Level
}
func (o *MessageNameGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MessageNameGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MessageNameGetInfoResponse structure represents the NetrMessageNameGetInfo operation response
type MessageNameGetInfoResponse struct {
	// InfoStruct:  A pointer to a structure of type MSG_INFO.
	Info *MessageInfo `idl:"name:InfoStruct;switch_is:Level" json:"info"`
	// Return: The NetrMessageNameGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MessageNameGetInfoResponse) xxx_ToOp(ctx context.Context) *xxx_MessageNameGetInfoOperation {
	if o == nil {
		return &xxx_MessageNameGetInfoOperation{}
	}
	return &xxx_MessageNameGetInfoOperation{
		Info:   o.Info,
		Return: o.Return,
	}
}

func (o *MessageNameGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_MessageNameGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MessageNameGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MessageNameGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MessageNameDeleteOperation structure represents the NetrMessageNameDel operation
type xxx_MessageNameDeleteOperation struct {
	ServerName  string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	MessageName string `idl:"name:MsgName;string" json:"message_name"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_MessageNameDeleteOperation) OpNum() int { return 3 }

func (o *xxx_MessageNameDeleteOperation) OpName() string { return "/msgsvc/v1/NetrMessageNameDel" }

func (o *xxx_MessageNameDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// MsgName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.MessageName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=MSGSVC_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MsgName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.MessageName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageNameDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MessageNameDeleteRequest structure represents the NetrMessageNameDel operation request
type MessageNameDeleteRequest struct {
	// ServerName: A pointer to a null-terminated string that MUST denote the NetBIOS name
	// (as specified in [RFC1001] section 5.2) or the fully qualified domain name (FQDN)
	// of the remote computer on which the function is to execute. There are no other constraints
	// on the format of this string. The message server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// MsgName:  A null-terminated Unicode UTF-16 string that MUST denote the recipient
	// name to delete. It is limited in length to 16 characters.<10>
	MessageName string `idl:"name:MsgName;string" json:"message_name"`
}

func (o *MessageNameDeleteRequest) xxx_ToOp(ctx context.Context) *xxx_MessageNameDeleteOperation {
	if o == nil {
		return &xxx_MessageNameDeleteOperation{}
	}
	return &xxx_MessageNameDeleteOperation{
		ServerName:  o.ServerName,
		MessageName: o.MessageName,
	}
}

func (o *MessageNameDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_MessageNameDeleteOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.MessageName = op.MessageName
}
func (o *MessageNameDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MessageNameDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MessageNameDeleteResponse structure represents the NetrMessageNameDel operation response
type MessageNameDeleteResponse struct {
	// Return: The NetrMessageNameDel return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MessageNameDeleteResponse) xxx_ToOp(ctx context.Context) *xxx_MessageNameDeleteOperation {
	if o == nil {
		return &xxx_MessageNameDeleteOperation{}
	}
	return &xxx_MessageNameDeleteOperation{
		Return: o.Return,
	}
}

func (o *MessageNameDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_MessageNameDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MessageNameDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MessageNameDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageNameDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

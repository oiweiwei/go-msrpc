// The swn package implements the SWN client protocol.
//
// # Introduction
//
// The Service Witness Protocol is a remote procedure call (RPC)-based protocol that
// is used to promptly notify a client of resource changes that have occurred on a highly
// available server.
//
// # Overview
//
// In highly available systems, there can be many instances of a service (for instance
// an SMB3 file service [MS-SMB2]) running on a server or group of servers. These service
// instances are accessed by clients through network DNS names and associated IP addresses.
//
// The Service Witness Protocol enables a client application (for instance, an SMB3
// client) to receive prompt and explicit notifications about the failure or recovery
// of a network name and associated services, rather than relying on slower detection
// mechanisms such as time-outs and keep-alives.
//
// The Service Witness Protocol is an independent protocol which is used alongside other
// protocols, as illustrated by the following figure.
package swn

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
	GoPackage = "swn"
)

// Shared structure represents PCONTEXT_HANDLE_SHARED RPC structure.
type Shared dcetypes.ContextHandle

func (o *Shared) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Shared) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Shared) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Shared) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ResponseAsyncNotify structure represents RESP_ASYNC_NOTIFY RPC structure.
//
// The RESP_ASYNC_NOTIFY structure contains the resource change type.
type ResponseAsyncNotify struct {
	// MessageType:  Specifies the notification type. This field MUST contain one of the
	// following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | RESOURCE_CHANGE_NOTIFICATION                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | CLIENT_MOVE_NOTIFICATION                                                         |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     3 | SHARE_MOVE_NOTIFICATION This value is applicable only for the server             |
	//	|       | implementing version 2.                                                          |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     4 | IP_CHANGE_NOTIFICATION This value is applicable only for the server implementing |
	//	|       | version 2.                                                                       |
	//	+-------+----------------------------------------------------------------------------------+
	MessageType uint32 `idl:"name:MessageType" json:"message_type"`
	// Length:  Specifies the size of the MessageBuffer field, in bytes.
	Length uint32 `idl:"name:Length" json:"length"`
	// NumberOfMessages:  Total number of notifications in the MessageBuffer field.
	NumberOfMessages uint32 `idl:"name:NumberOfMessages" json:"number_of_messages"`
	// MessageBuffer:  Contains an array of notification information structures whose type
	// is determined by the MessageType field.
	MessageBuffer []byte `idl:"name:MessageBuffer;size_is:(Length);pointer:unique" json:"message_buffer"`
}

func (o *ResponseAsyncNotify) xxx_PreparePayload(ctx context.Context) error {
	if o.MessageBuffer != nil && o.Length == 0 {
		o.Length = uint32(len(o.MessageBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ResponseAsyncNotify) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MessageType); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfMessages); err != nil {
		return err
	}
	if o.MessageBuffer != nil || o.Length > 0 {
		_ptr_MessageBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.MessageBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.MessageBuffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.MessageBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MessageBuffer, _ptr_MessageBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ResponseAsyncNotify) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessageType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfMessages); err != nil {
		return err
	}
	_ptr_MessageBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.MessageBuffer", sizeInfo[0])
		}
		o.MessageBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.MessageBuffer {
			i1 := i1
			if err := w.ReadData(&o.MessageBuffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_MessageBuffer := func(ptr interface{}) { o.MessageBuffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.MessageBuffer, _s_MessageBuffer, _ptr_MessageBuffer); err != nil {
		return err
	}
	return nil
}

// InterfaceInfo structure represents WITNESS_INTERFACE_INFO RPC structure.
//
// The WITNESS_INTERFACE_INFO structure specifies the IP addresses of the interface.
type InterfaceInfo struct {
	// InterfaceGroupName:  The null-terminated string that specifies a name of the interface
	// group.
	InterfaceGroupName []uint16 `idl:"name:InterfaceGroupName" json:"interface_group_name"`
	// Version:  The current version of the Witness Service running on the server.
	Version uint32 `idl:"name:Version" json:"version"`
	// State:  The current state of the interface. This field MUST contain one of the following
	// values:
	//
	//	+--------------------+----------------------------------------+
	//	|                    |                                        |
	//	|       VALUE        |                MEANING                 |
	//	|                    |                                        |
	//	+--------------------+----------------------------------------+
	//	+--------------------+----------------------------------------+
	//	| UNKNOWN 0x0000     | The state of the interface is unknown. |
	//	+--------------------+----------------------------------------+
	//	| AVAILABLE 0x0001   | The interface is available.            |
	//	+--------------------+----------------------------------------+
	//	| UNAVAILABLE 0x00FF | The interface is unavailable.          |
	//	+--------------------+----------------------------------------+
	State uint16 `idl:"name:State" json:"state"`
	// IPV4:  The IPv4 address of the interface, if the IPv4 flag is set in Flags field.
	IPv4 uint32 `idl:"name:IPV4" json:"ipv4"`
	// IPV6:  The IPv6 address of the interface, if the IPv6 flag is set in Flags field.
	IPv6 []uint16 `idl:"name:IPV6" json:"ipv6"`
	// Flags: The Flags field specifies information about the interface. This field MUST
	// be set to combination of zero or more of the following values:
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| IPv4 0x00000001              | If set, the IPV4 field contains a valid address.                                 |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| IPv6 0x00000002              | If set, the IPV6 field contains a valid address.                                 |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| INTERFACE_WITNESS 0x00000004 | If set, the interface is available for witness registration. If not set, the     |
	//	|                              | interface MUST NOT be used for witness registration.                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *InterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.InterfaceGroupName {
		i1 := i1
		if uint64(i1) >= 260 {
			break
		}
		if err := w.WriteData(o.InterfaceGroupName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceGroupName); uint64(i1) < 260; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4); err != nil {
		return err
	}
	for i1 := range o.IPv6 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.IPv6[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.IPv6); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.InterfaceGroupName = make([]uint16, 260)
	for i1 := range o.InterfaceGroupName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceGroupName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4); err != nil {
		return err
	}
	o.IPv6 = make([]uint16, 8)
	for i1 := range o.IPv6 {
		i1 := i1
		if err := w.ReadData(&o.IPv6[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// InterfaceList structure represents WITNESS_INTERFACE_LIST RPC structure.
//
// The WITNESS_INTERFACE_LIST structure specifies the list of interfaces available for
// witness registration.
type InterfaceList struct {
	// NumberOfInterfaces:  The number of WITNESS_INTERFACE_INFO structures in InterfaceInfo.
	NumberOfInterfaces uint32 `idl:"name:NumberOfInterfaces" json:"number_of_interfaces"`
	// InterfaceInfo:  Contains an array of WITNESS_INTERFACE_INFO structures, as specified
	// in section 2.2.2.5.
	InterfaceInfo []*InterfaceInfo `idl:"name:InterfaceInfo;size_is:(NumberOfInterfaces);pointer:unique" json:"interface_info"`
}

func (o *InterfaceList) xxx_PreparePayload(ctx context.Context) error {
	if o.InterfaceInfo != nil && o.NumberOfInterfaces == 0 {
		o.NumberOfInterfaces = uint32(len(o.InterfaceInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfInterfaces); err != nil {
		return err
	}
	if o.InterfaceInfo != nil || o.NumberOfInterfaces > 0 {
		_ptr_InterfaceInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfInterfaces)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.InterfaceInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.InterfaceInfo[i1] != nil {
					if err := o.InterfaceInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&InterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.InterfaceInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&InterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InterfaceInfo, _ptr_InterfaceInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfInterfaces); err != nil {
		return err
	}
	_ptr_InterfaceInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfInterfaces > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfInterfaces)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceInfo", sizeInfo[0])
		}
		o.InterfaceInfo = make([]*InterfaceInfo, sizeInfo[0])
		for i1 := range o.InterfaceInfo {
			i1 := i1
			if o.InterfaceInfo[i1] == nil {
				o.InterfaceInfo[i1] = &InterfaceInfo{}
			}
			if err := o.InterfaceInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_InterfaceInfo := func(ptr interface{}) { o.InterfaceInfo = *ptr.(*[]*InterfaceInfo) }
	if err := w.ReadPointer(&o.InterfaceInfo, _s_InterfaceInfo, _ptr_InterfaceInfo); err != nil {
		return err
	}
	return nil
}

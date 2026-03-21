package termsrvenumeration

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	tsts "github.com/oiweiwei/go-msrpc/msrpc/tsts"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
	_ = tsts.GoPackage
)

var (
	// import guard
	GoPackage = "tsts"
)

var (
	// Syntax UUID
	TerminateServerEnumerationSyntaxUUID = &uuid.UUID{TimeLow: 0x88143fd0, TimeMid: 0xc28d, TimeHiAndVersion: 0x4b2b, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xef, Node: [6]uint8{0x8d, 0x88, 0x2f, 0x6a, 0x93, 0x90}}
	// Syntax ID
	TerminateServerEnumerationSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: TerminateServerEnumerationSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// TermSrvEnumeration interface.
type TerminateServerEnumerationClient interface {

	// The RpcOpenEnum method returns a handle of the type ENUM_HANDLE, which can be used
	// to query information about the sessions that are currently running on a terminal
	// server. No special permissions are required to call this method. However, only sessions
	// to which the caller has WINSTATION_QUERY permission are enumerated.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	OpenEnum(context.Context, *OpenEnumRequest, ...dcerpc.CallOption) (*OpenEnumResponse, error)

	// The RpcCloseEnum method closes the enumeration object returned by RpcOpenEnum. This
	// method MUST be called after RpcOpenEnum. No special permissions are required to call
	// this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	CloseEnum(context.Context, *CloseEnumRequest, ...dcerpc.CallOption) (*CloseEnumResponse, error)

	// The RpcFilterByState method adds a filter to the session enumeration result, running
	// on a terminal server, based on the state of the sessions. This method MUST be called
	// after RpcOpenEnum and before RpcGetEnumResult or RpcGetEnumResultEx. No special permissions
	// are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	FilterByState(context.Context, *FilterByStateRequest, ...dcerpc.CallOption) (*FilterByStateResponse, error)

	// The RpcFilterByCallersName method adds a filter to the session enumeration result,
	// running on a terminal server, based on the caller name. The enumeration will return
	// only sessions belonging to the user making this call. This method MUST be called
	// after RpcOpenEnum and before RpcGetEnumResult or RpcGetEnumResultEx. No special permissions
	// are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	FilterByCallersName(context.Context, *FilterByCallersNameRequest, ...dcerpc.CallOption) (*FilterByCallersNameResponse, error)

	// The RpcEnumAddFilter method adds another filter to the current enumeration. This
	// method MUST be called after RpcOpenEnum and before RpcGetEnumResult or RpcGetEnumResultEx.
	// No special permissions are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	EnumAddFilter(context.Context, *EnumAddFilterRequest, ...dcerpc.CallOption) (*EnumAddFilterResponse, error)

	// The RpcGetEnumResult method returns a pointer of the type PSESSIONENUM which points
	// to the list of sessions currently running on the terminal server after applying the
	// specified filter. This method MUST be called after RpcOpenEnum. No special permissions
	// are required to call this method. However, only sessions for which the caller has
	// WINSTATION_QUERY permission are enumerated. The method checks whether the caller
	// has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access Request
	// mask, and skips the sessions for which the caller does not have the permission.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	GetEnumResult(context.Context, *GetEnumResultRequest, ...dcerpc.CallOption) (*GetEnumResultResponse, error)

	// The RpcFilterBySessionType method adds a filter to the session enumeration result,
	// running on a terminal server, based on the type of the session. This method MUST
	// be called after RpcOpenEnum and before RpcGetEnumResult or RpcGetEnumResultEx. No
	// special permissions are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	FilterBySessionType(context.Context, *FilterBySessionTypeRequest, ...dcerpc.CallOption) (*FilterBySessionTypeResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The RpcGetSessionIds method returns a list of the IDs associated with the sessions
	// running on a terminal server that satisfy the specified filter. No special permissions
	// are required to call this method. However, only sessions for which the caller has
	// WINSTATION_QUERY permission are enumerated. The method checks whether the caller
	// has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access Request
	// mask, and skips sessions for which the caller does not have the permission.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	GetSessionIDs(context.Context, *GetSessionIDsRequest, ...dcerpc.CallOption) (*GetSessionIDsResponse, error)

	// The RpcGetEnumResultEx method returns a pointer of the type PSESSIONENUM_EX, which
	// points to the list of sessions currently running on the terminal server after applying
	// the specified filter. This method MUST be called after RpcOpenEnum. No special permissions
	// are required to call this method. However, only sessions for which the caller has
	// WINSTATION_QUERY permission are enumerated. The method checks whether the caller
	// has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access Request
	// mask, and skips the sessions for which the caller does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	GetEnumResultEx(context.Context, *GetEnumResultExRequest, ...dcerpc.CallOption) (*GetEnumResultExResponse, error)

	// The RpcGetAllSessions method returns a pointer of type PEXECENVDATA, which points
	// to the list of sessions currently running on the terminal server and the sessions
	// running on various virtual machines hosted by the server. No special permissions
	// are required to call this method. However, only sessions for which the caller has
	// WINSTATION_QUERY permission are enumerated. The method checks whether the caller
	// has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access Request
	// mask, and skips the sessions for which the caller does not have the permission.<155>
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	GetAllSessions(context.Context, *GetAllSessionsRequest, ...dcerpc.CallOption) (*GetAllSessionsResponse, error)

	// The RpcGetAllSessionsEx method returns a pointer of type PEXECENVDATAEX, which points
	// to the list of sessions currently running on the terminal server and the sessions
	// running on various virtual machines hosted by the server. No special permissions
	// are required to call this method. However, only sessions for which the caller has
	// WINSTATION_QUERY permission are enumerated. The method checks whether the caller
	// has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access Request
	// mask, and skips the sessions for which the caller does not have the permission.<157>
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	GetAllSessionsEx(context.Context, *GetAllSessionsExRequest, ...dcerpc.CallOption) (*GetAllSessionsExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// EnumLevel1 represents the ENUM_LEVEL1 RPC constant
var EnumLevel1 = 1

// EnumLevel2 represents the ENUM_LEVEL2 RPC constant
var EnumLevel2 = 2

// EnumLevel3 represents the ENUM_LEVEL3 RPC constant
var EnumLevel3 = 3

// CurrentEnumLevel represents the CURRENT_ENUM_LEVEL RPC constant
var CurrentEnumLevel = 2

// UnifiedEnumLevel1 represents the UNIFIED_ENUM_LEVEL1 RPC constant
var UnifiedEnumLevel1 = 1

// UnifiedEnumLevel2 represents the UNIFIED_ENUM_LEVEL2 RPC constant
var UnifiedEnumLevel2 = 2

// CurrentUnifiedEnumLevel represents the CURRENT_UNIFIED_ENUM_LEVEL RPC constant
var CurrentUnifiedEnumLevel = 2

// Enum structure represents ENUM_HANDLE RPC structure.
type Enum dcetypes.ContextHandle

func (o *Enum) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Enum) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Enum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Enum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// SessionEnumLevel1 structure represents SESSIONENUM_LEVEL1 RPC structure.
//
// The SESSIONENUM_LEVEL1 structure contains basic information about sessions running
// on a computer.
type SessionEnumLevel1 struct {
	// SessionId:  An identifier assigned by the operating system to the session contained
	// in this structure.
	SessionID int32 `idl:"name:SessionId" json:"session_id"`
	// State:  The state of the session, as specified in section 3.3.4.1.8.
	State int32 `idl:"name:State" json:"state"`
	// Name:  A string that contains the name of the session assigned by Terminal Services
	// followed by the terminating NULL character.
	Name []uint16 `idl:"name:Name" json:"name"`
}

func (o *SessionEnumLevel1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SessionEnumLevel1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *SessionEnumLevel1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	o.Name = make([]uint16, 33)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// SessionEnumLevel2 structure represents SESSIONENUM_LEVEL2 RPC structure.
//
// The SESSIONENUM_LEVEL2 structure contains information about sessions running on a
// computer that is more detailed than the information contained in SESSIONENUM_LEVEL1.
type SessionEnumLevel2 struct {
	// SessionId:  An identifier assigned by the operating system to the session contained
	// in this structure.
	SessionID int32 `idl:"name:SessionId" json:"session_id"`
	// State:  The state of the session, as specified in section 3.3.4.1.8.
	State int32 `idl:"name:State" json:"state"`
	// Name:  A string that contains the name of the session followed by the terminating
	// NULL character.
	Name []uint16 `idl:"name:Name" json:"name"`
	// Source:  The parameter is always set to zero.
	Source uint32 `idl:"name:Source" json:"source"`
	// bFullDesktop:  The parameter is always set to TRUE.
	FullDesktop bool `idl:"name:bFullDesktop" json:"full_desktop"`
	// SessionType:  Describes the type of the session.<47>
	SessionType *dtyp.GUID `idl:"name:SessionType" json:"session_type"`
}

func (o *SessionEnumLevel2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SessionEnumLevel2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Source); err != nil {
		return err
	}
	if !o.FullDesktop {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.SessionType != nil {
		if err := o.SessionType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionEnumLevel2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	o.Name = make([]uint16, 33)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Source); err != nil {
		return err
	}
	var _bFullDesktop int32
	if err := w.ReadData(&_bFullDesktop); err != nil {
		return err
	}
	o.FullDesktop = _bFullDesktop != 0
	if o.SessionType == nil {
		o.SessionType = &dtyp.GUID{}
	}
	if err := o.SessionType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SessionEnumLevel3 structure represents SESSIONENUM_LEVEL3 RPC structure.
//
// The SESSIONENUM_LEVEL3 structure contains information about sessions running on a
// computer that is more detailed than the information contained in SESSIONENUM_LEVEL1
// and SESSIONENUM_LEVEL2.
type SessionEnumLevel3 struct {
	// SessionId:  An identifier assigned by the operating system to the session contained
	// in this structure.
	SessionID int32 `idl:"name:SessionId" json:"session_id"`
	// State:  The state of the session, as specified in section 3.3.4.1.8.
	State int32 `idl:"name:State" json:"state"`
	// Name:  A string that contains the name of the session followed by the terminating
	// NULL character.
	Name []uint16 `idl:"name:Name" json:"name"`
	// Source:  The parameter is always set to zero.
	//
	// bFullDesktop:  The parameter is always set to TRUE.
	//
	// SessionType:  The parameter is always set to zero.
	//
	// ProtoDataSize:  Size of data, in bytes, contained in the pProtocolData member.
	//
	// pProtocolData:  Data about the protocol status between the terminal server client
	// and server. This data will be of type PROTOCOLSTATUSEX.
	Source        uint32     `idl:"name:Source" json:"source"`
	FullDesktop   bool       `idl:"name:bFullDesktop" json:"full_desktop"`
	SessionType   *dtyp.GUID `idl:"name:SessionType" json:"session_type"`
	ProtoDataSize uint32     `idl:"name:ProtoDataSize" json:"proto_data_size"`
	ProtocolData  []byte     `idl:"name:pProtocolData;size_is:(ProtoDataSize)" json:"protocol_data"`
}

func (o *SessionEnumLevel3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ProtocolData != nil && o.ProtoDataSize == 0 {
		o.ProtoDataSize = uint32(len(o.ProtocolData))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SessionEnumLevel3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Source); err != nil {
		return err
	}
	if !o.FullDesktop {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.SessionType != nil {
		if err := o.SessionType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProtoDataSize); err != nil {
		return err
	}
	if o.ProtocolData != nil || o.ProtoDataSize > 0 {
		_ptr_pProtocolData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ProtoDataSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ProtocolData {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ProtocolData[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ProtocolData); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ProtocolData, _ptr_pProtocolData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionEnumLevel3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	o.Name = make([]uint16, 33)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Source); err != nil {
		return err
	}
	var _bFullDesktop int32
	if err := w.ReadData(&_bFullDesktop); err != nil {
		return err
	}
	o.FullDesktop = _bFullDesktop != 0
	if o.SessionType == nil {
		o.SessionType = &dtyp.GUID{}
	}
	if err := o.SessionType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtoDataSize); err != nil {
		return err
	}
	_ptr_pProtocolData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ProtoDataSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ProtoDataSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ProtocolData", sizeInfo[0])
		}
		o.ProtocolData = make([]byte, sizeInfo[0])
		for i1 := range o.ProtocolData {
			i1 := i1
			if err := w.ReadData(&o.ProtocolData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pProtocolData := func(ptr interface{}) { o.ProtocolData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ProtocolData, _s_pProtocolData, _ptr_pProtocolData); err != nil {
		return err
	}
	return nil
}

// SessionInfo structure represents SessionInfo RPC union.
type SessionInfo struct {
	// Types that are assignable to Value
	//
	// *SessionInfo_SessionEnumLevel1
	// *SessionInfo_SessionEnumLevel2
	Value is_SessionInfo `json:"value"`
}

func (o *SessionInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SessionInfo_SessionEnumLevel1:
		if value != nil {
			return value.SessionEnumLevel1
		}
	case *SessionInfo_SessionEnumLevel2:
		if value != nil {
			return value.SessionEnumLevel2
		}
	}
	return nil
}

type is_SessionInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SessionInfo()
}

func (o *SessionInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SessionInfo_SessionEnumLevel1:
		return uint32(1)
	case *SessionInfo_SessionEnumLevel2:
		return uint32(2)
	}
	return uint32(0)
}

func (o *SessionInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*SessionInfo_SessionEnumLevel1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SessionInfo_SessionEnumLevel1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*SessionInfo_SessionEnumLevel2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SessionInfo_SessionEnumLevel2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SessionInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &SessionInfo_SessionEnumLevel1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &SessionInfo_SessionEnumLevel2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SessionInfo_SessionEnumLevel1 structure represents SessionInfo RPC union arm.
//
// It has following labels: 1
type SessionInfo_SessionEnumLevel1 struct {
	SessionEnumLevel1 *SessionEnumLevel1 `idl:"name:SessionEnum_Level1" json:"session_enum_level1"`
}

func (*SessionInfo_SessionEnumLevel1) is_SessionInfo() {}

func (o *SessionInfo_SessionEnumLevel1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SessionEnumLevel1 != nil {
		if err := o.SessionEnumLevel1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SessionEnumLevel1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionInfo_SessionEnumLevel1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SessionEnumLevel1 == nil {
		o.SessionEnumLevel1 = &SessionEnumLevel1{}
	}
	if err := o.SessionEnumLevel1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SessionInfo_SessionEnumLevel2 structure represents SessionInfo RPC union arm.
//
// It has following labels: 2
type SessionInfo_SessionEnumLevel2 struct {
	SessionEnumLevel2 *SessionEnumLevel2 `idl:"name:SessionEnum_Level2" json:"session_enum_level2"`
}

func (*SessionInfo_SessionEnumLevel2) is_SessionInfo() {}

func (o *SessionInfo_SessionEnumLevel2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SessionEnumLevel2 != nil {
		if err := o.SessionEnumLevel2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SessionEnumLevel2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionInfo_SessionEnumLevel2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SessionEnumLevel2 == nil {
		o.SessionEnumLevel2 = &SessionEnumLevel2{}
	}
	if err := o.SessionEnumLevel2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SessionEnum structure represents SESSIONENUM RPC structure.
//
// PSESSIONENUM is a pointer to a structure containing information about the sessions
// running on the terminal server. It is returned by RpcGetEnumResult.
type SessionEnum struct {
	// Level:  The level of information contained in the Data member; the valid values are
	// 1 and 2.
	Level uint32 `idl:"name:Level" json:"level"`
	// Data:  Contains information at a specified level of detail about sessions running
	// on a computer.
	Data *SessionInfo `idl:"name:Data;switch_is:Level" json:"data"`
}

func (o *SessionEnum) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SessionEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swData := uint32(o.Level)
	if o.Data != nil {
		if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	} else {
		if err := (&SessionInfo{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.Data == nil {
		o.Data = &SessionInfo{}
	}
	_swData := uint32(o.Level)
	if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
		return err
	}
	return nil
}

// SessionInfoEx structure represents SessionInfo_Ex RPC union.
type SessionInfoEx struct {
	// Types that are assignable to Value
	//
	// *SessionInfoEx_SessionEnumLevel1
	// *SessionInfoEx_SessionEnumLevel2
	// *SessionInfoEx_SessionEnumLevel3
	Value is_SessionInfoEx `json:"value"`
}

func (o *SessionInfoEx) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SessionInfoEx_SessionEnumLevel1:
		if value != nil {
			return value.SessionEnumLevel1
		}
	case *SessionInfoEx_SessionEnumLevel2:
		if value != nil {
			return value.SessionEnumLevel2
		}
	case *SessionInfoEx_SessionEnumLevel3:
		if value != nil {
			return value.SessionEnumLevel3
		}
	}
	return nil
}

type is_SessionInfoEx interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SessionInfoEx()
}

func (o *SessionInfoEx) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SessionInfoEx_SessionEnumLevel1:
		return uint32(1)
	case *SessionInfoEx_SessionEnumLevel2:
		return uint32(2)
	case *SessionInfoEx_SessionEnumLevel3:
		return uint32(3)
	}
	return uint32(0)
}

func (o *SessionInfoEx) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*SessionInfoEx_SessionEnumLevel1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SessionInfoEx_SessionEnumLevel1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*SessionInfoEx_SessionEnumLevel2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SessionInfoEx_SessionEnumLevel2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*SessionInfoEx_SessionEnumLevel3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SessionInfoEx_SessionEnumLevel3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SessionInfoEx) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &SessionInfoEx_SessionEnumLevel1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &SessionInfoEx_SessionEnumLevel2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &SessionInfoEx_SessionEnumLevel3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SessionInfoEx_SessionEnumLevel1 structure represents SessionInfo_Ex RPC union arm.
//
// It has following labels: 1
type SessionInfoEx_SessionEnumLevel1 struct {
	SessionEnumLevel1 *SessionEnumLevel1 `idl:"name:SessionEnum_Level1" json:"session_enum_level1"`
}

func (*SessionInfoEx_SessionEnumLevel1) is_SessionInfoEx() {}

func (o *SessionInfoEx_SessionEnumLevel1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SessionEnumLevel1 != nil {
		if err := o.SessionEnumLevel1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SessionEnumLevel1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionInfoEx_SessionEnumLevel1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SessionEnumLevel1 == nil {
		o.SessionEnumLevel1 = &SessionEnumLevel1{}
	}
	if err := o.SessionEnumLevel1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SessionInfoEx_SessionEnumLevel2 structure represents SessionInfo_Ex RPC union arm.
//
// It has following labels: 2
type SessionInfoEx_SessionEnumLevel2 struct {
	SessionEnumLevel2 *SessionEnumLevel2 `idl:"name:SessionEnum_Level2" json:"session_enum_level2"`
}

func (*SessionInfoEx_SessionEnumLevel2) is_SessionInfoEx() {}

func (o *SessionInfoEx_SessionEnumLevel2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SessionEnumLevel2 != nil {
		if err := o.SessionEnumLevel2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SessionEnumLevel2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionInfoEx_SessionEnumLevel2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SessionEnumLevel2 == nil {
		o.SessionEnumLevel2 = &SessionEnumLevel2{}
	}
	if err := o.SessionEnumLevel2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SessionInfoEx_SessionEnumLevel3 structure represents SessionInfo_Ex RPC union arm.
//
// It has following labels: 3
type SessionInfoEx_SessionEnumLevel3 struct {
	SessionEnumLevel3 *SessionEnumLevel3 `idl:"name:SessionEnum_Level3" json:"session_enum_level3"`
}

func (*SessionInfoEx_SessionEnumLevel3) is_SessionInfoEx() {}

func (o *SessionInfoEx_SessionEnumLevel3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SessionEnumLevel3 != nil {
		if err := o.SessionEnumLevel3.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SessionEnumLevel3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionInfoEx_SessionEnumLevel3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SessionEnumLevel3 == nil {
		o.SessionEnumLevel3 = &SessionEnumLevel3{}
	}
	if err := o.SessionEnumLevel3.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SessionEnumEx structure represents SESSIONENUM_EX RPC structure.
//
// The PSESSIONENUM_EX is a pointer to a structure containing information about the
// sessions running on the terminal server. It is returned by RpcGetEnumResultEx.
type SessionEnumEx struct {
	// Level:  The level of information contained in Data; the valid values are 1, 2, and
	// 3.
	//
	//	+-------+----------------------------------------------------------------+
	//	|       |                                                                |
	//	| VALUE |                            MEANING                             |
	//	|       |                                                                |
	//	+-------+----------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------+
	//	|     1 | The union SessionInfo_Ex has the SessionEnum_Level1 structure. |
	//	+-------+----------------------------------------------------------------+
	//	|     2 | The union SessionInfo_Ex has the SessionEnum_Level2 structure. |
	//	+-------+----------------------------------------------------------------+
	//	|     3 | The union SessionInfo_Ex has the SessionEnum_Level3 structure. |
	//	+-------+----------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// Data:  Contains information at a specified level of detail about sessions running
	// on a computer. This parameter is of type SessionInfo_Ex. If Level is set to 1, the
	// union SessionInfo_Ex has the SessionEnum_Level1 structure. If Level is set to 2,
	// the union SessionInfo_Ex has the SessionEnum_Level2 structure. If Level is set to
	// 3, the union SessionInfo_Ex has the SessionEnum_Level3 structure.
	Data *SessionInfoEx `idl:"name:Data;switch_is:Level" json:"data"`
}

func (o *SessionEnumEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SessionEnumEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swData := uint32(o.Level)
	if o.Data != nil {
		if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	} else {
		if err := (&SessionInfoEx{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	}
	return nil
}
func (o *SessionEnumEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.Data == nil {
		o.Data = &SessionInfoEx{}
	}
	_swData := uint32(o.Level)
	if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
		return err
	}
	return nil
}

// ExecEnvDataLevel1 structure represents EXECENVDATA_LEVEL1 RPC structure.
//
// The EXECENVDATA_LEVEL1 structure contains basic information about sessions running
// on a computer.
type ExecEnvDataLevel1 struct {
	// ExecEnvId:  An identifier assigned to the session contained in this structure by
	// the component that aggregates the sessions on the server and sessions within virtual
	// machines hosted on the server.<52>
	ExecEnvID int32 `idl:"name:ExecEnvId" json:"exec_env_id"`
	// State:  The state of the session, as specified in section 3.3.4.1.8.
	State int32 `idl:"name:State" json:"state"`
	// SessionName:  A string that contains the name of the session assigned by Terminal
	// Services followed by the terminating NULL character.
	SessionName []uint16 `idl:"name:SessionName" json:"session_name"`
}

func (o *ExecEnvDataLevel1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ExecEnvDataLevel1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ExecEnvID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	for i1 := range o.SessionName {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.SessionName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SessionName); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *ExecEnvDataLevel1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExecEnvID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	o.SessionName = make([]uint16, 33)
	for i1 := range o.SessionName {
		i1 := i1
		if err := w.ReadData(&o.SessionName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// ExecEnvDataLevel2 structure represents EXECENVDATA_LEVEL2 RPC structure.
//
// The EXECENVDATA_LEVEL2 structure contains information about sessions running on a
// computer that is more detailed than the information contained in EXECENVDATA_LEVEL1.
type ExecEnvDataLevel2 struct {
	// ExecEnvId:  An identifier assigned to the session contained in this structure by
	// the component that aggregates the sessions on the server and sessions within virtual
	// machines hosted on the server.<53>
	ExecEnvID int32 `idl:"name:ExecEnvId" json:"exec_env_id"`
	// State:  The state of the session, as specified in section 3.3.4.1.8.
	State int32 `idl:"name:State" json:"state"`
	// SessionName:  A string that contains the name of the session followed by the terminating
	// NULL character.
	SessionName []uint16 `idl:"name:SessionName" json:"session_name"`
	// AbsSessionId:  An identifier assigned by the operating system running in the virtual
	// machine to the session contained in this structure. If the session contained in this
	// structure is not running under the virtual machine, the value of AbsSessionId is
	// same as ExecEnvId.
	AbsSessionID int32 `idl:"name:AbsSessionId" json:"abs_session_id"`
	// HostName:  A string that contains the name of the machine that hosts the session
	// contained in this structure followed by the terminating NULL character.
	HostName []uint16 `idl:"name:HostName" json:"host_name"`
	// UserName:  A string that contains the name of the user logged onto the session followed
	// by the terminating NULL character.
	UserName []uint16 `idl:"name:UserName" json:"user_name"`
	// DomainName:  A string that contains the domain of the user logged onto the session
	// followed by the terminating NULL character.
	DomainName []uint16 `idl:"name:DomainName" json:"domain_name"`
	// FarmName:  A string that contains the farm name associated with the session followed
	// by the terminating NULL character.
	FarmName []uint16 `idl:"name:FarmName" json:"farm_name"`
}

func (o *ExecEnvDataLevel2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ExecEnvDataLevel2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ExecEnvID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	for i1 := range o.SessionName {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.SessionName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SessionName); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AbsSessionID); err != nil {
		return err
	}
	for i1 := range o.HostName {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.HostName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.HostName); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.UserName {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.UserName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UserName); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DomainName {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.DomainName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DomainName); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.FarmName {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.FarmName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.FarmName); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *ExecEnvDataLevel2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExecEnvID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	o.SessionName = make([]uint16, 33)
	for i1 := range o.SessionName {
		i1 := i1
		if err := w.ReadData(&o.SessionName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AbsSessionID); err != nil {
		return err
	}
	o.HostName = make([]uint16, 33)
	for i1 := range o.HostName {
		i1 := i1
		if err := w.ReadData(&o.HostName[i1]); err != nil {
			return err
		}
	}
	o.UserName = make([]uint16, 33)
	for i1 := range o.UserName {
		i1 := i1
		if err := w.ReadData(&o.UserName[i1]); err != nil {
			return err
		}
	}
	o.DomainName = make([]uint16, 33)
	for i1 := range o.DomainName {
		i1 := i1
		if err := w.ReadData(&o.DomainName[i1]); err != nil {
			return err
		}
	}
	o.FarmName = make([]uint16, 33)
	for i1 := range o.FarmName {
		i1 := i1
		if err := w.ReadData(&o.FarmName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// ExecEnvEnum structure represents ExecEnvEnum RPC union.
type ExecEnvEnum struct {
	// Types that are assignable to Value
	//
	// *ExecEnvEnum_Level1
	// *ExecEnvEnum_Level2
	Value is_ExecEnvEnum `json:"value"`
}

func (o *ExecEnvEnum) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ExecEnvEnum_Level1:
		if value != nil {
			return value.ExecEnvEnumLevel1
		}
	case *ExecEnvEnum_Level2:
		if value != nil {
			return value.ExecEnvEnumLevel2
		}
	}
	return nil
}

type is_ExecEnvEnum interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ExecEnvEnum()
}

func (o *ExecEnvEnum) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ExecEnvEnum_Level1:
		return uint32(1)
	case *ExecEnvEnum_Level2:
		return uint32(2)
	}
	return uint32(0)
}

func (o *ExecEnvEnum) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*ExecEnvEnum_Level1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExecEnvEnum_Level1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*ExecEnvEnum_Level2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExecEnvEnum_Level2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ExecEnvEnum) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &ExecEnvEnum_Level1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &ExecEnvEnum_Level2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ExecEnvEnum_Level1 structure represents ExecEnvEnum RPC union arm.
//
// It has following labels: 1
type ExecEnvEnum_Level1 struct {
	ExecEnvEnumLevel1 *ExecEnvDataLevel1 `idl:"name:ExecEnvEnum_Level1" json:"exec_env_enum_level1"`
}

func (*ExecEnvEnum_Level1) is_ExecEnvEnum() {}

func (o *ExecEnvEnum_Level1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ExecEnvEnumLevel1 != nil {
		if err := o.ExecEnvEnumLevel1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ExecEnvDataLevel1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExecEnvEnum_Level1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ExecEnvEnumLevel1 == nil {
		o.ExecEnvEnumLevel1 = &ExecEnvDataLevel1{}
	}
	if err := o.ExecEnvEnumLevel1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ExecEnvEnum_Level2 structure represents ExecEnvEnum RPC union arm.
//
// It has following labels: 2
type ExecEnvEnum_Level2 struct {
	ExecEnvEnumLevel2 *ExecEnvDataLevel2 `idl:"name:ExecEnvEnum_Level2" json:"exec_env_enum_level2"`
}

func (*ExecEnvEnum_Level2) is_ExecEnvEnum() {}

func (o *ExecEnvEnum_Level2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ExecEnvEnumLevel2 != nil {
		if err := o.ExecEnvEnumLevel2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ExecEnvDataLevel2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExecEnvEnum_Level2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ExecEnvEnumLevel2 == nil {
		o.ExecEnvEnumLevel2 = &ExecEnvDataLevel2{}
	}
	if err := o.ExecEnvEnumLevel2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ExecEnvData structure represents EXECENVDATA RPC structure.
//
// PEXECENVDATA is a pointer to a structure containing information about the sessions
// running on the terminal server and the sessions running on virtual machines hosted
// on the server.<48> It is returned by RpcGetAllSessions.
type ExecEnvData struct {
	// Level:  The level of information contained in Data; the valid values are 1 and 2.
	//
	//	+-------+-------------------------------------------------------------+
	//	|       |                                                             |
	//	| VALUE |                           MEANING                           |
	//	|       |                                                             |
	//	+-------+-------------------------------------------------------------+
	//	+-------+-------------------------------------------------------------+
	//	|     1 | The union ExecEnvData has the EXECENVDATA_LEVEL1 structure. |
	//	+-------+-------------------------------------------------------------+
	//	|     2 | The union ExecEnvData has the EXECENVDATA_LEVEL2 structure. |
	//	+-------+-------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// Data:  Contains information at a specified level of detail about sessions running
	// on a computer. This is of type ExecEnvData.
	Data *ExecEnvEnum `idl:"name:Data;switch_is:Level" json:"data"`
}

func (o *ExecEnvData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ExecEnvData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swData := uint32(o.Level)
	if o.Data != nil {
		if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	} else {
		if err := (&ExecEnvEnum{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExecEnvData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.Data == nil {
		o.Data = &ExecEnvEnum{}
	}
	_swData := uint32(o.Level)
	if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
		return err
	}
	return nil
}

// ExecEnvDataExLevel1 structure represents EXECENVDATAEX_LEVEL1 RPC structure.
//
// The EXECENVDATAEX_LEVEL1 structure contains information about sessions running on
// a computer.
type ExecEnvDataExLevel1 struct {
	// ExecEnvId:  An identifier assigned to the session contained in this structure by
	// the component that aggregates the sessions on the server and sessions within virtual
	// machines hosted on the server.
	ExecEnvID int32 `idl:"name:ExecEnvId" json:"exec_env_id"`
	// State:  The state of the session, as specified in section 3.3.4.1.8.
	State int32 `idl:"name:State" json:"state"`
	// AbsSessionId:  An identifier assigned by the operating system running in the virtual
	// machine to the session contained in this structure. If the session contained in this
	// structure is not running under the virtual machine, the value of AbsSessionId is
	// the same as the value of the ExecEnvId member.
	AbsSessionID int32 `idl:"name:AbsSessionId" json:"abs_session_id"`
	// pszSessionName:  A string that contains the name of the session followed by the terminating
	// NULL character.
	SessionName string `idl:"name:pszSessionName;max_is:(256);string" json:"session_name"`
	// pszHostName:  A string that contains the name of the machine that hosts the session
	// contained in this structure, followed by the terminating NULL character.
	HostName string `idl:"name:pszHostName;max_is:(256);string" json:"host_name"`
	// pszUserName:  A string that contains the name of the user logged onto the session
	// followed by the terminating NULL character.
	UserName string `idl:"name:pszUserName;max_is:(256);string" json:"user_name"`
	// pszDomainName:  A string that contains the domain of the user logged onto the session
	// followed by the terminating NULL character.
	DomainName string `idl:"name:pszDomainName;max_is:(256);string" json:"domain_name"`
	// pszFarmName:  A string that contains the farm name associated with the session followed
	// by the terminating NULL character.
	FarmName string `idl:"name:pszFarmName;max_is:(256);string" json:"farm_name"`
}

func (o *ExecEnvDataExLevel1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	// cannot evaluate expression 257
	// cannot evaluate expression 257
	// cannot evaluate expression 257
	// cannot evaluate expression 257
	// cannot evaluate expression 257
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ExecEnvDataExLevel1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ExecEnvID); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.AbsSessionID); err != nil {
		return err
	}
	if o.SessionName != "" || 257 > 0 {
		_ptr_pszSessionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(257)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.SessionName)
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
			_SessionName_buf := utf16.Encode([]rune(o.SessionName))
			if uint64(len(_SessionName_buf)) > sizeInfo[0]-1 {
				_SessionName_buf = _SessionName_buf[:sizeInfo[0]-1]
			}
			if o.SessionName != ndr.ZeroString {
				_SessionName_buf = append(_SessionName_buf, uint16(0))
			}
			for i1 := range _SessionName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_SessionName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_SessionName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SessionName, _ptr_pszSessionName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.HostName != "" || 257 > 0 {
		_ptr_pszHostName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(257)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.HostName)
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
			_HostName_buf := utf16.Encode([]rune(o.HostName))
			if uint64(len(_HostName_buf)) > sizeInfo[0]-1 {
				_HostName_buf = _HostName_buf[:sizeInfo[0]-1]
			}
			if o.HostName != ndr.ZeroString {
				_HostName_buf = append(_HostName_buf, uint16(0))
			}
			for i1 := range _HostName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_HostName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_HostName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.HostName, _ptr_pszHostName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.UserName != "" || 257 > 0 {
		_ptr_pszUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(257)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.UserName)
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
			_UserName_buf := utf16.Encode([]rune(o.UserName))
			if uint64(len(_UserName_buf)) > sizeInfo[0]-1 {
				_UserName_buf = _UserName_buf[:sizeInfo[0]-1]
			}
			if o.UserName != ndr.ZeroString {
				_UserName_buf = append(_UserName_buf, uint16(0))
			}
			for i1 := range _UserName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_UserName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_UserName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UserName, _ptr_pszUserName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainName != "" || 257 > 0 {
		_ptr_pszDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(257)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.DomainName)
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
			_DomainName_buf := utf16.Encode([]rune(o.DomainName))
			if uint64(len(_DomainName_buf)) > sizeInfo[0]-1 {
				_DomainName_buf = _DomainName_buf[:sizeInfo[0]-1]
			}
			if o.DomainName != ndr.ZeroString {
				_DomainName_buf = append(_DomainName_buf, uint16(0))
			}
			for i1 := range _DomainName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_DomainName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_DomainName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainName, _ptr_pszDomainName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.FarmName != "" || 257 > 0 {
		_ptr_pszFarmName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(257)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.FarmName)
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
			_FarmName_buf := utf16.Encode([]rune(o.FarmName))
			if uint64(len(_FarmName_buf)) > sizeInfo[0]-1 {
				_FarmName_buf = _FarmName_buf[:sizeInfo[0]-1]
			}
			if o.FarmName != ndr.ZeroString {
				_FarmName_buf = append(_FarmName_buf, uint16(0))
			}
			for i1 := range _FarmName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_FarmName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_FarmName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.FarmName, _ptr_pszFarmName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExecEnvDataExLevel1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExecEnvID); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.AbsSessionID); err != nil {
		return err
	}
	_ptr_pszSessionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _SessionName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _SessionName_buf", sizeInfo[0])
		}
		_SessionName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _SessionName_buf {
			i1 := i1
			if err := w.ReadData(&_SessionName_buf[i1]); err != nil {
				return err
			}
		}
		o.SessionName = strings.TrimRight(string(utf16.Decode(_SessionName_buf)), ndr.ZeroString)
		return nil
	})
	_s_pszSessionName := func(ptr interface{}) { o.SessionName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SessionName, _s_pszSessionName, _ptr_pszSessionName); err != nil {
		return err
	}
	_ptr_pszHostName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _HostName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _HostName_buf", sizeInfo[0])
		}
		_HostName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _HostName_buf {
			i1 := i1
			if err := w.ReadData(&_HostName_buf[i1]); err != nil {
				return err
			}
		}
		o.HostName = strings.TrimRight(string(utf16.Decode(_HostName_buf)), ndr.ZeroString)
		return nil
	})
	_s_pszHostName := func(ptr interface{}) { o.HostName = *ptr.(*string) }
	if err := w.ReadPointer(&o.HostName, _s_pszHostName, _ptr_pszHostName); err != nil {
		return err
	}
	_ptr_pszUserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _UserName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _UserName_buf", sizeInfo[0])
		}
		_UserName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _UserName_buf {
			i1 := i1
			if err := w.ReadData(&_UserName_buf[i1]); err != nil {
				return err
			}
		}
		o.UserName = strings.TrimRight(string(utf16.Decode(_UserName_buf)), ndr.ZeroString)
		return nil
	})
	_s_pszUserName := func(ptr interface{}) { o.UserName = *ptr.(*string) }
	if err := w.ReadPointer(&o.UserName, _s_pszUserName, _ptr_pszUserName); err != nil {
		return err
	}
	_ptr_pszDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _DomainName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _DomainName_buf", sizeInfo[0])
		}
		_DomainName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _DomainName_buf {
			i1 := i1
			if err := w.ReadData(&_DomainName_buf[i1]); err != nil {
				return err
			}
		}
		o.DomainName = strings.TrimRight(string(utf16.Decode(_DomainName_buf)), ndr.ZeroString)
		return nil
	})
	_s_pszDomainName := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainName, _s_pszDomainName, _ptr_pszDomainName); err != nil {
		return err
	}
	_ptr_pszFarmName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _FarmName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _FarmName_buf", sizeInfo[0])
		}
		_FarmName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _FarmName_buf {
			i1 := i1
			if err := w.ReadData(&_FarmName_buf[i1]); err != nil {
				return err
			}
		}
		o.FarmName = strings.TrimRight(string(utf16.Decode(_FarmName_buf)), ndr.ZeroString)
		return nil
	})
	_s_pszFarmName := func(ptr interface{}) { o.FarmName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FarmName, _s_pszFarmName, _ptr_pszFarmName); err != nil {
		return err
	}
	return nil
}

// ExecEnvEnumEx structure represents ExecEnvEnumEx RPC union.
type ExecEnvEnumEx struct {
	// Types that are assignable to Value
	//
	// *ExecEnvEnumEx_Level1
	Value is_ExecEnvEnumEx `json:"value"`
}

func (o *ExecEnvEnumEx) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ExecEnvEnumEx_Level1:
		if value != nil {
			return value.ExecEnvEnumExLevel1
		}
	}
	return nil
}

type is_ExecEnvEnumEx interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ExecEnvEnumEx()
}

func (o *ExecEnvEnumEx) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ExecEnvEnumEx_Level1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *ExecEnvEnumEx) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*ExecEnvEnumEx_Level1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExecEnvEnumEx_Level1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ExecEnvEnumEx) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &ExecEnvEnumEx_Level1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ExecEnvEnumEx_Level1 structure represents ExecEnvEnumEx RPC union arm.
//
// It has following labels: 1
type ExecEnvEnumEx_Level1 struct {
	ExecEnvEnumExLevel1 *ExecEnvDataExLevel1 `idl:"name:ExecEnvEnumEx_Level1" json:"exec_env_enum_ex_level1"`
}

func (*ExecEnvEnumEx_Level1) is_ExecEnvEnumEx() {}

func (o *ExecEnvEnumEx_Level1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ExecEnvEnumExLevel1 != nil {
		if err := o.ExecEnvEnumExLevel1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ExecEnvDataExLevel1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExecEnvEnumEx_Level1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ExecEnvEnumExLevel1 == nil {
		o.ExecEnvEnumExLevel1 = &ExecEnvDataExLevel1{}
	}
	if err := o.ExecEnvEnumExLevel1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ExecEnvDataEx structure represents EXECENVDATAEX RPC structure.
//
// PEXECENVDATAEX is a pointer to a structure containing information about the sessions
// running on the terminal server and the sessions running on virtual machines hosted
// on the server.<54> It is returned by RpcGetAllSessionsEx.
type ExecEnvDataEx struct {
	// Level:  The level of information contained in the Data member; the only valid value
	// is 1.
	//
	//	+-------+-----------------------------------------------------------------+
	//	|       |                                                                 |
	//	| VALUE |                             MEANING                             |
	//	|       |                                                                 |
	//	+-------+-----------------------------------------------------------------+
	//	+-------+-----------------------------------------------------------------+
	//	|     1 | The union ExecEnvDataEx has the EXECENVDATAEX_LEVEL1 structure. |
	//	+-------+-----------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// Data:  Contains information at a specified level of detail about sessions running
	// on a computer. This is of type ExecEnvDataEx.
	Data *ExecEnvEnumEx `idl:"name:Data;switch_is:Level" json:"data"`
}

func (o *ExecEnvDataEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ExecEnvDataEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swData := uint32(o.Level)
	if o.Data != nil {
		if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	} else {
		if err := (&ExecEnvEnumEx{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExecEnvDataEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.Data == nil {
		o.Data = &ExecEnvEnumEx{}
	}
	_swData := uint32(o.Level)
	if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultTerminateServerEnumerationClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTerminateServerEnumerationClient) OpenEnum(ctx context.Context, in *OpenEnumRequest, opts ...dcerpc.CallOption) (*OpenEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) CloseEnum(ctx context.Context, in *CloseEnumRequest, opts ...dcerpc.CallOption) (*CloseEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) FilterByState(ctx context.Context, in *FilterByStateRequest, opts ...dcerpc.CallOption) (*FilterByStateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FilterByStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) FilterByCallersName(ctx context.Context, in *FilterByCallersNameRequest, opts ...dcerpc.CallOption) (*FilterByCallersNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FilterByCallersNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) EnumAddFilter(ctx context.Context, in *EnumAddFilterRequest, opts ...dcerpc.CallOption) (*EnumAddFilterResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumAddFilterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) GetEnumResult(ctx context.Context, in *GetEnumResultRequest, opts ...dcerpc.CallOption) (*GetEnumResultResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEnumResultResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) FilterBySessionType(ctx context.Context, in *FilterBySessionTypeRequest, opts ...dcerpc.CallOption) (*FilterBySessionTypeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FilterBySessionTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) GetSessionIDs(ctx context.Context, in *GetSessionIDsRequest, opts ...dcerpc.CallOption) (*GetSessionIDsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSessionIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) GetEnumResultEx(ctx context.Context, in *GetEnumResultExRequest, opts ...dcerpc.CallOption) (*GetEnumResultExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEnumResultExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) GetAllSessions(ctx context.Context, in *GetAllSessionsRequest, opts ...dcerpc.CallOption) (*GetAllSessionsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetAllSessionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) GetAllSessionsEx(ctx context.Context, in *GetAllSessionsExRequest, opts ...dcerpc.CallOption) (*GetAllSessionsExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetAllSessionsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerEnumerationClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTerminateServerEnumerationClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewTerminateServerEnumerationClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TerminateServerEnumerationClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TerminateServerEnumerationSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTerminateServerEnumerationClient{cc: cc}, nil
}

// xxx_OpenEnumOperation structure represents the RpcOpenEnum operation
type xxx_OpenEnumOperation struct {
	Enum   *Enum `idl:"name:phEnum" json:"enum"`
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenEnumOperation) OpNum() int { return 0 }

func (o *xxx_OpenEnumOperation) OpName() string { return "/TermSrvEnumeration/v1/RpcOpenEnum" }

func (o *xxx_OpenEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_OpenEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_OpenEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phEnum {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phEnum {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenEnumRequest structure represents the RpcOpenEnum operation request
type OpenEnumRequest struct {
}

func (o *OpenEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenEnumOperation) *xxx_OpenEnumOperation {
	if op == nil {
		op = &xxx_OpenEnumOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *OpenEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenEnumOperation) {
	if o == nil {
		return
	}
}
func (o *OpenEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenEnumResponse structure represents the RpcOpenEnum operation response
type OpenEnumResponse struct {
	// phEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:phEnum" json:"enum"`
	// Return: The RpcOpenEnum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenEnumOperation) *xxx_OpenEnumOperation {
	if op == nil {
		op = &xxx_OpenEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *OpenEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenEnumOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *OpenEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseEnumOperation structure represents the RpcCloseEnum operation
type xxx_CloseEnumOperation struct {
	Enum   *Enum `idl:"name:phEnum" json:"enum"`
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseEnumOperation) OpNum() int { return 1 }

func (o *xxx_CloseEnumOperation) OpName() string { return "/TermSrvEnumeration/v1/RpcCloseEnum" }

func (o *xxx_CloseEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phEnum {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phEnum {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phEnum {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phEnum {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseEnumRequest structure represents the RpcCloseEnum operation request
type CloseEnumRequest struct {
	// phEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:phEnum" json:"enum"`
}

func (o *CloseEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseEnumOperation) *xxx_CloseEnumOperation {
	if op == nil {
		op = &xxx_CloseEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	return op
}

func (o *CloseEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseEnumOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
}
func (o *CloseEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseEnumResponse structure represents the RpcCloseEnum operation response
type CloseEnumResponse struct {
	// phEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:phEnum" json:"enum"`
	// Return: The RpcCloseEnum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseEnumOperation) *xxx_CloseEnumOperation {
	if op == nil {
		op = &xxx_CloseEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *CloseEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseEnumOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *CloseEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FilterByStateOperation structure represents the RpcFilterByState operation
type xxx_FilterByStateOperation struct {
	Enum   *Enum `idl:"name:hEnum" json:"enum"`
	State  int32 `idl:"name:State" json:"state"`
	Invert bool  `idl:"name:bInvert" json:"invert"`
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FilterByStateOperation) OpNum() int { return 2 }

func (o *xxx_FilterByStateOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcFilterByState"
}

func (o *xxx_FilterByStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterByStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// State {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// bInvert {in} (1:{alias=BOOL}(int32))
	{
		if !o.Invert {
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

func (o *xxx_FilterByStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// State {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// bInvert {in} (1:{alias=BOOL}(int32))
	{
		var _bInvert int32
		if err := w.ReadData(&_bInvert); err != nil {
			return err
		}
		o.Invert = _bInvert != 0
	}
	return nil
}

func (o *xxx_FilterByStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterByStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterByStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FilterByStateRequest structure represents the RpcFilterByState operation request
type FilterByStateRequest struct {
	// hEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:hEnum" json:"enum"`
	// State: The session state, as specified in section 3.3.4.1.8, to be used to filter
	// out the enumeration result. Only the sessions with the specified state will be returned.
	State int32 `idl:"name:State" json:"state"`
	// bInvert:  Set to TRUE to imply that the result of the comparison during enumeration
	// will be inverted, FALSE otherwise.
	Invert bool `idl:"name:bInvert" json:"invert"`
}

func (o *FilterByStateRequest) xxx_ToOp(ctx context.Context, op *xxx_FilterByStateOperation) *xxx_FilterByStateOperation {
	if op == nil {
		op = &xxx_FilterByStateOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.State = o.State
	op.Invert = o.Invert
	return op
}

func (o *FilterByStateRequest) xxx_FromOp(ctx context.Context, op *xxx_FilterByStateOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.State = op.State
	o.Invert = op.Invert
}
func (o *FilterByStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FilterByStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FilterByStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FilterByStateResponse structure represents the RpcFilterByState operation response
type FilterByStateResponse struct {
	// Return: The RpcFilterByState return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FilterByStateResponse) xxx_ToOp(ctx context.Context, op *xxx_FilterByStateOperation) *xxx_FilterByStateOperation {
	if op == nil {
		op = &xxx_FilterByStateOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FilterByStateResponse) xxx_FromOp(ctx context.Context, op *xxx_FilterByStateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FilterByStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FilterByStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FilterByStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FilterByCallersNameOperation structure represents the RpcFilterByCallersName operation
type xxx_FilterByCallersNameOperation struct {
	Enum   *Enum `idl:"name:hEnum" json:"enum"`
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FilterByCallersNameOperation) OpNum() int { return 3 }

func (o *xxx_FilterByCallersNameOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcFilterByCallersName"
}

func (o *xxx_FilterByCallersNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterByCallersNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_FilterByCallersNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterByCallersNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterByCallersNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterByCallersNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FilterByCallersNameRequest structure represents the RpcFilterByCallersName operation request
type FilterByCallersNameRequest struct {
	// hEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:hEnum" json:"enum"`
}

func (o *FilterByCallersNameRequest) xxx_ToOp(ctx context.Context, op *xxx_FilterByCallersNameOperation) *xxx_FilterByCallersNameOperation {
	if op == nil {
		op = &xxx_FilterByCallersNameOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	return op
}

func (o *FilterByCallersNameRequest) xxx_FromOp(ctx context.Context, op *xxx_FilterByCallersNameOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
}
func (o *FilterByCallersNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FilterByCallersNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FilterByCallersNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FilterByCallersNameResponse structure represents the RpcFilterByCallersName operation response
type FilterByCallersNameResponse struct {
	// Return: The RpcFilterByCallersName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FilterByCallersNameResponse) xxx_ToOp(ctx context.Context, op *xxx_FilterByCallersNameOperation) *xxx_FilterByCallersNameOperation {
	if op == nil {
		op = &xxx_FilterByCallersNameOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FilterByCallersNameResponse) xxx_FromOp(ctx context.Context, op *xxx_FilterByCallersNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FilterByCallersNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FilterByCallersNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FilterByCallersNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumAddFilterOperation structure represents the RpcEnumAddFilter operation
type xxx_EnumAddFilterOperation struct {
	Enum    *Enum `idl:"name:hEnum" json:"enum"`
	SubEnum *Enum `idl:"name:hSubEnum" json:"sub_enum"`
	Return  int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumAddFilterOperation) OpNum() int { return 4 }

func (o *xxx_EnumAddFilterOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcEnumAddFilter"
}

func (o *xxx_EnumAddFilterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAddFilterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hSubEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SubEnum != nil {
			if err := o.SubEnum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnumAddFilterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hSubEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SubEnum == nil {
			o.SubEnum = &Enum{}
		}
		if err := o.SubEnum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAddFilterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAddFilterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAddFilterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumAddFilterRequest structure represents the RpcEnumAddFilter operation request
type EnumAddFilterRequest struct {
	// hEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:hEnum" json:"enum"`
	// hSubEnum:  The handle to another enumeration whose filter will be applied to the
	// current enumeration. The other enumeration MUST be created using RpcOpenEnum and
	// any combination of RpcFilterByState, RpcFilterByCallersName, and RpcFilterBySessionType.
	SubEnum *Enum `idl:"name:hSubEnum" json:"sub_enum"`
}

func (o *EnumAddFilterRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumAddFilterOperation) *xxx_EnumAddFilterOperation {
	if op == nil {
		op = &xxx_EnumAddFilterOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.SubEnum = o.SubEnum
	return op
}

func (o *EnumAddFilterRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumAddFilterOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.SubEnum = op.SubEnum
}
func (o *EnumAddFilterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumAddFilterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAddFilterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumAddFilterResponse structure represents the RpcEnumAddFilter operation response
type EnumAddFilterResponse struct {
	// Return: The RpcEnumAddFilter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumAddFilterResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumAddFilterOperation) *xxx_EnumAddFilterOperation {
	if op == nil {
		op = &xxx_EnumAddFilterOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EnumAddFilterResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumAddFilterOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EnumAddFilterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumAddFilterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAddFilterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEnumResultOperation structure represents the RpcGetEnumResult operation
type xxx_GetEnumResultOperation struct {
	Enum              *Enum          `idl:"name:hEnum" json:"enum"`
	SessionEnumResult []*SessionEnum `idl:"name:ppSessionEnumResult;size_is:(, pEntries)" json:"session_enum_result"`
	Level             uint32         `idl:"name:Level" json:"level"`
	Entries           uint32         `idl:"name:pEntries" json:"entries"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEnumResultOperation) OpNum() int { return 5 }

func (o *xxx_GetEnumResultOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcGetEnumResult"
}

func (o *xxx_GetEnumResultOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnumResultOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetEnumResultOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetEnumResultOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SessionEnumResult != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SessionEnumResult))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnumResultOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppSessionEnumResult {out} (1:{pointer=ref}*(2))(2:{alias=PSESSIONENUM}*(1))(3:{alias=SESSIONENUM}[dim:0,size_is=pEntries](struct))
	{
		if o.SessionEnumResult != nil || o.Entries > 0 {
			_ptr_ppSessionEnumResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Entries)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SessionEnumResult {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.SessionEnumResult[i1] != nil {
						if err := o.SessionEnumResult[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&SessionEnum{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.SessionEnumResult); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&SessionEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SessionEnumResult, _ptr_ppSessionEnumResult); err != nil {
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
	// pEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Entries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnumResultOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppSessionEnumResult {out} (1:{pointer=ref}*(2))(2:{alias=PSESSIONENUM,pointer=ref}*(1))(3:{alias=SESSIONENUM}[dim:0,size_is=pEntries](struct))
	{
		_ptr_ppSessionEnumResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SessionEnumResult", sizeInfo[0])
			}
			o.SessionEnumResult = make([]*SessionEnum, sizeInfo[0])
			for i1 := range o.SessionEnumResult {
				i1 := i1
				if o.SessionEnumResult[i1] == nil {
					o.SessionEnumResult[i1] = &SessionEnum{}
				}
				if err := o.SessionEnumResult[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSessionEnumResult := func(ptr interface{}) { o.SessionEnumResult = *ptr.(*[]*SessionEnum) }
		if err := w.ReadPointer(&o.SessionEnumResult, _s_ppSessionEnumResult, _ptr_ppSessionEnumResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Entries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetEnumResultRequest structure represents the RpcGetEnumResult operation request
type GetEnumResultRequest struct {
	// hEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:hEnum" json:"enum"`
	// Level: The level of information requested. This field MUST be set to the supported
	// values of 1 or 2. If the server does not support the level requested, it will set
	// the level to the maximum supported level and return information accordingly.<153>
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *GetEnumResultRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEnumResultOperation) *xxx_GetEnumResultOperation {
	if op == nil {
		op = &xxx_GetEnumResultOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.Level = o.Level
	return op
}

func (o *GetEnumResultRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEnumResultOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.Level = op.Level
}
func (o *GetEnumResultRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEnumResultRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnumResultOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEnumResultResponse structure represents the RpcGetEnumResult operation response
type GetEnumResultResponse struct {
	// ppSessionEnumResult: A pointer of the type PSESSIONENUM which points to the list
	// of sessions currently running on the terminal server.
	SessionEnumResult []*SessionEnum `idl:"name:ppSessionEnumResult;size_is:(, pEntries)" json:"session_enum_result"`
	// pEntries:  The number of sessions currently running on the terminal server and returned
	// in the ppSessionEnumResult structure.
	Entries uint32 `idl:"name:pEntries" json:"entries"`
	// Return: The RpcGetEnumResult return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEnumResultResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEnumResultOperation) *xxx_GetEnumResultOperation {
	if op == nil {
		op = &xxx_GetEnumResultOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionEnumResult = o.SessionEnumResult
	op.Entries = o.Entries
	op.Return = o.Return
	return op
}

func (o *GetEnumResultResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEnumResultOperation) {
	if o == nil {
		return
	}
	o.SessionEnumResult = op.SessionEnumResult
	o.Entries = op.Entries
	o.Return = op.Return
}
func (o *GetEnumResultResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEnumResultResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnumResultOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FilterBySessionTypeOperation structure represents the RpcFilterBySessionType operation
type xxx_FilterBySessionTypeOperation struct {
	Enum        *Enum      `idl:"name:hEnum" json:"enum"`
	SessionType *dtyp.GUID `idl:"name:pSessionType" json:"session_type"`
	Return      int32      `idl:"name:Return" json:"return"`
}

func (o *xxx_FilterBySessionTypeOperation) OpNum() int { return 6 }

func (o *xxx_FilterBySessionTypeOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcFilterBySessionType"
}

func (o *xxx_FilterBySessionTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterBySessionTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pSessionType {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.SessionType != nil {
			if err := o.SessionType.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_FilterBySessionTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pSessionType {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.SessionType == nil {
			o.SessionType = &dtyp.GUID{}
		}
		if err := o.SessionType.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterBySessionTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterBySessionTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FilterBySessionTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FilterBySessionTypeRequest structure represents the RpcFilterBySessionType operation request
type FilterBySessionTypeRequest struct {
	// hEnum:  The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:hEnum" json:"enum"`
	// pSessionType: The session GUID to be used to filter out the enumeration result. Only
	// the session with the specified GUID will be returned.
	SessionType *dtyp.GUID `idl:"name:pSessionType" json:"session_type"`
}

func (o *FilterBySessionTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_FilterBySessionTypeOperation) *xxx_FilterBySessionTypeOperation {
	if op == nil {
		op = &xxx_FilterBySessionTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.SessionType = o.SessionType
	return op
}

func (o *FilterBySessionTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_FilterBySessionTypeOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.SessionType = op.SessionType
}
func (o *FilterBySessionTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FilterBySessionTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FilterBySessionTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FilterBySessionTypeResponse structure represents the RpcFilterBySessionType operation response
type FilterBySessionTypeResponse struct {
	// Return: The RpcFilterBySessionType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FilterBySessionTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_FilterBySessionTypeOperation) *xxx_FilterBySessionTypeOperation {
	if op == nil {
		op = &xxx_FilterBySessionTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FilterBySessionTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_FilterBySessionTypeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FilterBySessionTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FilterBySessionTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FilterBySessionTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSessionIDsOperation structure represents the RpcGetSessionIds operation
type xxx_GetSessionIDsOperation struct {
	Filter          tsts.SessionFilter `idl:"name:Filter" json:"filter"`
	MaxEntries      uint32             `idl:"name:MaxEntries" json:"max_entries"`
	SessionIDs      []int32            `idl:"name:pSessionIds;size_is:(, pcSessionIds)" json:"session_ids"`
	SessionIDsCount uint32             `idl:"name:pcSessionIds" json:"session_ids_count"`
	Return          int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSessionIDsOperation) OpNum() int { return 8 }

func (o *xxx_GetSessionIDsOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcGetSessionIds"
}

func (o *xxx_GetSessionIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.MaxEntries > uint32(65535) {
		return fmt.Errorf("MaxEntries is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Filter {in} (1:{alias=SESSION_FILTER}(enum))
	{
		if err := w.WriteEnum(uint16(o.Filter)); err != nil {
			return err
		}
	}
	// MaxEntries {in} (1:{range=(0,65535), alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MaxEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Filter {in} (1:{alias=SESSION_FILTER}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Filter)); err != nil {
			return err
		}
	}
	// MaxEntries {in} (1:{range=(0,65535), alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MaxEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SessionIDs != nil && o.SessionIDsCount == 0 {
		o.SessionIDsCount = uint32(len(o.SessionIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pSessionIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcSessionIds](int32))
	{
		if o.SessionIDs != nil || o.SessionIDsCount > 0 {
			_ptr_pSessionIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.SessionIDsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SessionIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.SessionIDs[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.SessionIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(int32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SessionIDs, _ptr_pSessionIds); err != nil {
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
	// pcSessionIds {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionIDsCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pSessionIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcSessionIds](int32))
	{
		_ptr_pSessionIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SessionIDs", sizeInfo[0])
			}
			o.SessionIDs = make([]int32, sizeInfo[0])
			for i1 := range o.SessionIDs {
				i1 := i1
				if err := w.ReadData(&o.SessionIDs[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pSessionIds := func(ptr interface{}) { o.SessionIDs = *ptr.(*[]int32) }
		if err := w.ReadPointer(&o.SessionIDs, _s_pSessionIds, _ptr_pSessionIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcSessionIds {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionIDsCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSessionIDsRequest structure represents the RpcGetSessionIds operation request
type GetSessionIDsRequest struct {
	// Filter:  The filter to apply to the sessions running on the terminal server. This
	// is of type SESSION_FILTER.
	Filter tsts.SessionFilter `idl:"name:Filter" json:"filter"`
	// MaxEntries:  The maximum number of session IDs to return. This value MUST be less
	// than or equal to the number of entries in the pSessionIds array.
	MaxEntries uint32 `idl:"name:MaxEntries" json:"max_entries"`
}

func (o *GetSessionIDsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSessionIDsOperation) *xxx_GetSessionIDsOperation {
	if op == nil {
		op = &xxx_GetSessionIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.Filter = o.Filter
	op.MaxEntries = o.MaxEntries
	return op
}

func (o *GetSessionIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSessionIDsOperation) {
	if o == nil {
		return
	}
	o.Filter = op.Filter
	o.MaxEntries = op.MaxEntries
}
func (o *GetSessionIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSessionIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSessionIDsResponse structure represents the RpcGetSessionIds operation response
type GetSessionIDsResponse struct {
	// pSessionIds:  An array to contain the list of session IDs running on the terminal
	// server and filtered as specified by Filter.
	SessionIDs []int32 `idl:"name:pSessionIds;size_is:(, pcSessionIds)" json:"session_ids"`
	// pcSessionIds:  The number of session IDs returned.
	SessionIDsCount uint32 `idl:"name:pcSessionIds" json:"session_ids_count"`
	// Return: The RpcGetSessionIds return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSessionIDsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSessionIDsOperation) *xxx_GetSessionIDsOperation {
	if op == nil {
		op = &xxx_GetSessionIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionIDs = o.SessionIDs
	op.SessionIDsCount = o.SessionIDsCount
	op.Return = o.Return
	return op
}

func (o *GetSessionIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSessionIDsOperation) {
	if o == nil {
		return
	}
	o.SessionIDs = op.SessionIDs
	o.SessionIDsCount = op.SessionIDsCount
	o.Return = op.Return
}
func (o *GetSessionIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSessionIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEnumResultExOperation structure represents the RpcGetEnumResultEx operation
type xxx_GetEnumResultExOperation struct {
	Enum              *Enum            `idl:"name:hEnum" json:"enum"`
	SessionEnumResult []*SessionEnumEx `idl:"name:ppSessionEnumResult;size_is:(, pEntries)" json:"session_enum_result"`
	Level             uint32           `idl:"name:Level" json:"level"`
	Entries           uint32           `idl:"name:pEntries" json:"entries"`
	Return            int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEnumResultExOperation) OpNum() int { return 9 }

func (o *xxx_GetEnumResultExOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcGetEnumResultEx"
}

func (o *xxx_GetEnumResultExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnumResultExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum != nil {
			if err := o.Enum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Enum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetEnumResultExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hEnum {in} (1:{context_handle, alias=ENUM_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Enum == nil {
			o.Enum = &Enum{}
		}
		if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetEnumResultExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SessionEnumResult != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SessionEnumResult))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnumResultExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppSessionEnumResult {out} (1:{pointer=ref}*(2))(2:{alias=PSESSIONENUM_EX}*(1))(3:{alias=SESSIONENUM_EX}[dim:0,size_is=pEntries](struct))
	{
		if o.SessionEnumResult != nil || o.Entries > 0 {
			_ptr_ppSessionEnumResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Entries)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SessionEnumResult {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.SessionEnumResult[i1] != nil {
						if err := o.SessionEnumResult[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&SessionEnumEx{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.SessionEnumResult); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&SessionEnumEx{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SessionEnumResult, _ptr_ppSessionEnumResult); err != nil {
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
	// pEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Entries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnumResultExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppSessionEnumResult {out} (1:{pointer=ref}*(2))(2:{alias=PSESSIONENUM_EX,pointer=ref}*(1))(3:{alias=SESSIONENUM_EX}[dim:0,size_is=pEntries](struct))
	{
		_ptr_ppSessionEnumResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SessionEnumResult", sizeInfo[0])
			}
			o.SessionEnumResult = make([]*SessionEnumEx, sizeInfo[0])
			for i1 := range o.SessionEnumResult {
				i1 := i1
				if o.SessionEnumResult[i1] == nil {
					o.SessionEnumResult[i1] = &SessionEnumEx{}
				}
				if err := o.SessionEnumResult[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSessionEnumResult := func(ptr interface{}) { o.SessionEnumResult = *ptr.(*[]*SessionEnumEx) }
		if err := w.ReadPointer(&o.SessionEnumResult, _s_ppSessionEnumResult, _ptr_ppSessionEnumResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Entries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetEnumResultExRequest structure represents the RpcGetEnumResultEx operation request
type GetEnumResultExRequest struct {
	// hEnum: The handle to the session enumeration object. This is of type ENUM_HANDLE.
	Enum *Enum `idl:"name:hEnum" json:"enum"`
	// Level: The level of information requested. This field MUST be set to the supported
	// values of 1, 2, or 3. If the server does not support the level requested, it will
	// set the level to the maximum supported level and return information accordingly.<154>
	//
	//	+-------+----------------------------------------------------------------------+
	//	|       |                                                                      |
	//	| VALUE |                               MEANING                                |
	//	|       |                                                                      |
	//	+-------+----------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------+
	//	|     1 | The union SessionInfo_Ex will have the SessionEnum_Level1 structure. |
	//	+-------+----------------------------------------------------------------------+
	//	|     2 | The union SessionInfo_Ex will have the SessionEnum_Level2 structure. |
	//	+-------+----------------------------------------------------------------------+
	//	|     3 | The union SessionInfo_Ex will have the SessionEnum_Level3 structure. |
	//	+-------+----------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *GetEnumResultExRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEnumResultExOperation) *xxx_GetEnumResultExOperation {
	if op == nil {
		op = &xxx_GetEnumResultExOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.Level = o.Level
	return op
}

func (o *GetEnumResultExRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEnumResultExOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.Level = op.Level
}
func (o *GetEnumResultExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEnumResultExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnumResultExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEnumResultExResponse structure represents the RpcGetEnumResultEx operation response
type GetEnumResultExResponse struct {
	// ppSessionEnumResult: A pointer of the type PSESSIONENUM_EX which points to the list
	// of sessions currently running on the terminal server.
	SessionEnumResult []*SessionEnumEx `idl:"name:ppSessionEnumResult;size_is:(, pEntries)" json:"session_enum_result"`
	// pEntries: The number of sessions currently running on the terminal server and returned
	// in the ppSessionEnumResult structure.
	Entries uint32 `idl:"name:pEntries" json:"entries"`
	// Return: The RpcGetEnumResultEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEnumResultExResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEnumResultExOperation) *xxx_GetEnumResultExOperation {
	if op == nil {
		op = &xxx_GetEnumResultExOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionEnumResult = o.SessionEnumResult
	op.Entries = o.Entries
	op.Return = o.Return
	return op
}

func (o *GetEnumResultExResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEnumResultExOperation) {
	if o == nil {
		return
	}
	o.SessionEnumResult = op.SessionEnumResult
	o.Entries = op.Entries
	o.Return = op.Return
}
func (o *GetEnumResultExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEnumResultExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnumResultExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllSessionsOperation structure represents the RpcGetAllSessions operation
type xxx_GetAllSessionsOperation struct {
	Level        uint32         `idl:"name:pLevel" json:"level"`
	SessionData  []*ExecEnvData `idl:"name:ppSessionData;size_is:(, pcEntries)" json:"session_data"`
	EntriesCount uint32         `idl:"name:pcEntries" json:"entries_count"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllSessionsOperation) OpNum() int { return 10 }

func (o *xxx_GetAllSessionsOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcGetAllSessions"
}

func (o *xxx_GetAllSessionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pLevel {in, out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pLevel {in, out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SessionData != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.SessionData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pLevel {in, out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// ppSessionData {out} (1:{pointer=ref}*(2))(2:{alias=PEXECENVDATA}*(1))(3:{alias=EXECENVDATA}[dim:0,size_is=pcEntries](struct))
	{
		if o.SessionData != nil || o.EntriesCount > 0 {
			_ptr_ppSessionData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.EntriesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SessionData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.SessionData[i1] != nil {
						if err := o.SessionData[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&ExecEnvData{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.SessionData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&ExecEnvData{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SessionData, _ptr_ppSessionData); err != nil {
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
	// pcEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.EntriesCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pLevel {in, out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// ppSessionData {out} (1:{pointer=ref}*(2))(2:{alias=PEXECENVDATA,pointer=ref}*(1))(3:{alias=EXECENVDATA}[dim:0,size_is=pcEntries](struct))
	{
		_ptr_ppSessionData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SessionData", sizeInfo[0])
			}
			o.SessionData = make([]*ExecEnvData, sizeInfo[0])
			for i1 := range o.SessionData {
				i1 := i1
				if o.SessionData[i1] == nil {
					o.SessionData[i1] = &ExecEnvData{}
				}
				if err := o.SessionData[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSessionData := func(ptr interface{}) { o.SessionData = *ptr.(*[]*ExecEnvData) }
		if err := w.ReadPointer(&o.SessionData, _s_ppSessionData, _ptr_ppSessionData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.EntriesCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetAllSessionsRequest structure represents the RpcGetAllSessions operation request
type GetAllSessionsRequest struct {
	// pLevel: The level of information requested. This field MUST be set to the supported
	// values of 1 or 2. If the server does not support the level requested, it will set
	// the level to the maximum supported level and return an implementation-specific nonzero
	// value.
	//
	//	+-------+-------------------------------------------------------------------+
	//	|       |                                                                   |
	//	| VALUE |                              MEANING                              |
	//	|       |                                                                   |
	//	+-------+-------------------------------------------------------------------+
	//	+-------+-------------------------------------------------------------------+
	//	|     1 | The union ExecEnvData will have the EXECENVDATA_LEVEL1 structure. |
	//	+-------+-------------------------------------------------------------------+
	//	|     2 | The union ExecEnvData will have the EXECENVDATA_LEVEL2 structure. |
	//	+-------+-------------------------------------------------------------------+
	Level uint32 `idl:"name:pLevel" json:"level"`
}

func (o *GetAllSessionsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAllSessionsOperation) *xxx_GetAllSessionsOperation {
	if op == nil {
		op = &xxx_GetAllSessionsOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	return op
}

func (o *GetAllSessionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllSessionsOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
}
func (o *GetAllSessionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAllSessionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllSessionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllSessionsResponse structure represents the RpcGetAllSessions operation response
type GetAllSessionsResponse struct {
	// pLevel: The level of information requested. This field MUST be set to the supported
	// values of 1 or 2. If the server does not support the level requested, it will set
	// the level to the maximum supported level and return an implementation-specific nonzero
	// value.
	//
	//	+-------+-------------------------------------------------------------------+
	//	|       |                                                                   |
	//	| VALUE |                              MEANING                              |
	//	|       |                                                                   |
	//	+-------+-------------------------------------------------------------------+
	//	+-------+-------------------------------------------------------------------+
	//	|     1 | The union ExecEnvData will have the EXECENVDATA_LEVEL1 structure. |
	//	+-------+-------------------------------------------------------------------+
	//	|     2 | The union ExecEnvData will have the EXECENVDATA_LEVEL2 structure. |
	//	+-------+-------------------------------------------------------------------+
	Level uint32 `idl:"name:pLevel" json:"level"`
	// ppSessionData: A pointer of type PEXECENVDATA (section 2.2.2.6), which points to
	// the list of sessions currently running on the terminal server and the session running
	// on virtual machines hosted by the server.<156>
	SessionData []*ExecEnvData `idl:"name:ppSessionData;size_is:(, pcEntries)" json:"session_data"`
	// pcEntries: The number of sessions currently running on the terminal server and returned
	// in the ppSessionData structure.
	EntriesCount uint32 `idl:"name:pcEntries" json:"entries_count"`
	// Return: The RpcGetAllSessions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllSessionsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAllSessionsOperation) *xxx_GetAllSessionsOperation {
	if op == nil {
		op = &xxx_GetAllSessionsOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.SessionData = o.SessionData
	op.EntriesCount = o.EntriesCount
	op.Return = o.Return
	return op
}

func (o *GetAllSessionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllSessionsOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.SessionData = op.SessionData
	o.EntriesCount = op.EntriesCount
	o.Return = op.Return
}
func (o *GetAllSessionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAllSessionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllSessionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllSessionsExOperation structure represents the RpcGetAllSessionsEx operation
type xxx_GetAllSessionsExOperation struct {
	Level        uint32           `idl:"name:Level" json:"level"`
	SessionData  []*ExecEnvDataEx `idl:"name:ppSessionData;size_is:(, pcEntries)" json:"session_data"`
	EntriesCount uint32           `idl:"name:pcEntries" json:"entries_count"`
	Return       int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllSessionsExOperation) OpNum() int { return 11 }

func (o *xxx_GetAllSessionsExOperation) OpName() string {
	return "/TermSrvEnumeration/v1/RpcGetAllSessionsEx"
}

func (o *xxx_GetAllSessionsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Level {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Level {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SessionData != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.SessionData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppSessionData {out} (1:{pointer=ref}*(2))(2:{alias=PEXECENVDATAEX}*(1))(3:{alias=EXECENVDATAEX}[dim:0,size_is=pcEntries](struct))
	{
		if o.SessionData != nil || o.EntriesCount > 0 {
			_ptr_ppSessionData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.EntriesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SessionData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.SessionData[i1] != nil {
						if err := o.SessionData[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&ExecEnvDataEx{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.SessionData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&ExecEnvDataEx{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SessionData, _ptr_ppSessionData); err != nil {
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
	// pcEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.EntriesCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllSessionsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppSessionData {out} (1:{pointer=ref}*(2))(2:{alias=PEXECENVDATAEX,pointer=ref}*(1))(3:{alias=EXECENVDATAEX}[dim:0,size_is=pcEntries](struct))
	{
		_ptr_ppSessionData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SessionData", sizeInfo[0])
			}
			o.SessionData = make([]*ExecEnvDataEx, sizeInfo[0])
			for i1 := range o.SessionData {
				i1 := i1
				if o.SessionData[i1] == nil {
					o.SessionData[i1] = &ExecEnvDataEx{}
				}
				if err := o.SessionData[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSessionData := func(ptr interface{}) { o.SessionData = *ptr.(*[]*ExecEnvDataEx) }
		if err := w.ReadPointer(&o.SessionData, _s_ppSessionData, _ptr_ppSessionData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.EntriesCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetAllSessionsExRequest structure represents the RpcGetAllSessionsEx operation request
type GetAllSessionsExRequest struct {
	// Level: The level of information requested. This field MUST be set to 1. If the server
	// does not support the level requested, it will set the level to the maximum supported
	// level and return an implementation-specific nonzero value.
	//
	//	+-------+-----------------------------------------------------------------+
	//	|       |                                                                 |
	//	| VALUE |                             MEANING                             |
	//	|       |                                                                 |
	//	+-------+-----------------------------------------------------------------+
	//	+-------+-----------------------------------------------------------------+
	//	|     1 | The union ExecEnvDataEx has the EXECENVDATAEX_LEVEL1 structure. |
	//	+-------+-----------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *GetAllSessionsExRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAllSessionsExOperation) *xxx_GetAllSessionsExOperation {
	if op == nil {
		op = &xxx_GetAllSessionsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	return op
}

func (o *GetAllSessionsExRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllSessionsExOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
}
func (o *GetAllSessionsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAllSessionsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllSessionsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllSessionsExResponse structure represents the RpcGetAllSessionsEx operation response
type GetAllSessionsExResponse struct {
	// ppSessionData: A pointer of type PEXECENVDATAEX (section 2.2.2.7), which points to
	// the list of sessions currently running on the terminal server and the sessions running
	// on virtual machines hosted by the server.
	SessionData []*ExecEnvDataEx `idl:"name:ppSessionData;size_is:(, pcEntries)" json:"session_data"`
	// pcEntries: The number of sessions currently running on the terminal server and returned
	// in the ppSessionData structure.
	EntriesCount uint32 `idl:"name:pcEntries" json:"entries_count"`
	// Return: The RpcGetAllSessionsEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllSessionsExResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAllSessionsExOperation) *xxx_GetAllSessionsExOperation {
	if op == nil {
		op = &xxx_GetAllSessionsExOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionData = o.SessionData
	op.EntriesCount = o.EntriesCount
	op.Return = o.Return
	return op
}

func (o *GetAllSessionsExResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllSessionsExOperation) {
	if o == nil {
		return
	}
	o.SessionData = op.SessionData
	o.EntriesCount = op.EntriesCount
	o.Return = op.Return
}
func (o *GetAllSessionsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAllSessionsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllSessionsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

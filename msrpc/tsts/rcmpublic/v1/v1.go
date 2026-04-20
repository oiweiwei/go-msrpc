package rcmpublic

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

var (
	// import guard
	GoPackage = "tsts"
)

var (
	// Syntax UUID
	RcmPublicSyntaxUUID = &uuid.UUID{TimeLow: 0xbde95fdf, TimeMid: 0xeee0, TimeHiAndVersion: 0x45de, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x12, Node: [6]uint8{0xe5, 0xa6, 0x1c, 0xd0, 0xd4, 0xfe}}
	// Syntax ID
	RcmPublicSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: RcmPublicSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// RCMPublic interface.
type RcmPublicClient interface {

	// The RpcGetClientData method returns information about the client that is connected
	// to a particular session running on a terminal server. The caller must have WINSTATION_QUERY
	// permission. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
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
	GetClientData(context.Context, *GetClientDataRequest, ...dcerpc.CallOption) (*GetClientDataResponse, error)

	// The RpcGetConfigData method returns the configuration data associated with the user
	// connected to a particular session running on a terminal server. The caller MUST have
	// WINSTATION_QUERY permission. The method checks whether the caller has WINSTATION_QUERY
	// permission (section 3.1.1) by setting it as the Access Request mask, and fails if
	// the caller does not have the permission.<158>
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
	GetConfigData(context.Context, *GetConfigDataRequest, ...dcerpc.CallOption) (*GetConfigDataResponse, error)

	// The RpcGetProtocolStatus method retrieves information about the status of the protocol
	// used to connect to a particular session running on a terminal server. The caller
	// MUST have WINSTATION_QUERY permission for the session. The method checks whether
	// the caller has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access
	// Request mask, and fails if the caller does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	GetProtocolStatus(context.Context, *GetProtocolStatusRequest, ...dcerpc.CallOption) (*GetProtocolStatusResponse, error)

	// The RpcGetLastInputTime method returns the time the last user input was received
	// by the associated protocol for the specified sessions running on a terminal server.
	// The caller MUST have WINSTATION_QUERY permission for the session. The method checks
	// whether the caller has WINSTATION_QUERY permission (section 3.1.1) by setting it
	// as the Access Request mask, and fails if the caller does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	GetLastInputTime(context.Context, *GetLastInputTimeRequest, ...dcerpc.CallOption) (*GetLastInputTimeResponse, error)

	// The RpcGetRemoteAddress method retrieves the IP address of the client computer connected
	// to the session on the terminal server. The caller MUST have WINSTATION_QUERY permission
	// for the session. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
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
	GetRemoteAddress(context.Context, *GetRemoteAddressRequest, ...dcerpc.CallOption) (*GetRemoteAddressResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The RpcGetAllListeners method returns a list of all Terminal Services listeners running
	// on a terminal server. No special permissions are required to call this method. However,
	// only listeners for which the caller has WINSTATION_QUERY permission are enumerated.
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
	GetAllListeners(context.Context, *GetAllListenersRequest, ...dcerpc.CallOption) (*GetAllListenersResponse, error)

	// The RpcGetSessionProtocolLastInputTime method returns the protocol status and the
	// time the last input was received by the protocol associated with a specific session
	// running on a terminal server. The caller MUST have WINSTATION_QUERY permission for
	// the session. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	GetSessionProtocolLastInputTime(context.Context, *GetSessionProtocolLastInputTimeRequest, ...dcerpc.CallOption) (*GetSessionProtocolLastInputTimeResponse, error)

	// The RpcGetUserCertificates method returns a client X509 certificate if the client
	// used the certificate to connect to a session running on a terminal server. For more
	// information, see [X509]. The caller MUST have WINSTATION_QUERY permission for the
	// session. The method checks whether the caller has WINSTATION_QUERY permission (section
	// 3.1.1) by setting it as the Access Request mask, and fails if the caller does not
	// have the permission.
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
	GetUserCertificates(context.Context, *GetUserCertificatesRequest, ...dcerpc.CallOption) (*GetUserCertificatesResponse, error)

	// The RpcQuerySessionData method returns information about a particular session running
	// on a terminal server. The caller MUST have WINSTATION_QUERY permission to the session
	// being queried. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	QuerySessionData(context.Context, *QuerySessionDataRequest, ...dcerpc.CallOption) (*QuerySessionDataResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// LstEnumLevel1 represents the LST_ENUM_LEVEL1 RPC constant
var LstEnumLevel1 = 1

// CurrentLstEnumLevel represents the CURRENT_LST_ENUM_LEVEL RPC constant
var CurrentLstEnumLevel = 1

// RcmRemoteaddress structure represents RCM_REMOTEADDRESS RPC structure.
//
// The RCM_REMOTEADDRESS structure defines a remote address.
type RcmRemoteaddress struct {
	// sin_family:  Specifies the type of IP address. Valid values are 2 for IPv4 addresses,
	// and 23 for IPv6 addresses.
	Family           uint16                             `idl:"name:sin_family" json:"family"`
	RcmRemoteaddress *RcmRemoteaddress_RcmRemoteaddress `idl:"name:RcmRemoteaddress" json:"rcm_remoteaddress"`
}

func (o *RcmRemoteaddress) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RcmRemoteaddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Family); err != nil {
		return err
	}
	if o.RcmRemoteaddress != nil {
		if err := o.RcmRemoteaddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RcmRemoteaddress_RcmRemoteaddress{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RcmRemoteaddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Family); err != nil {
		return err
	}
	if o.RcmRemoteaddress == nil {
		o.RcmRemoteaddress = &RcmRemoteaddress_RcmRemoteaddress{}
	}
	if err := o.RcmRemoteaddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RcmRemoteaddress_RcmRemoteaddress structure represents RCM_REMOTEADDRESS union anonymous member.
//
// The RCM_REMOTEADDRESS structure defines a remote address.
type RcmRemoteaddress_RcmRemoteaddress struct {
	Family uint16
	// Types that are assignable to Value
	//
	// *RcmRemoteaddress_IPv4
	// *RcmRemoteaddress_IPv6
	Value is_RcmRemoteaddress_RcmRemoteaddress `json:"value"`
}

func (o *RcmRemoteaddress_RcmRemoteaddress) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RcmRemoteaddress_IPv4:
		if value != nil {
			return value
		}
	case *RcmRemoteaddress_IPv6:
		if value != nil {
			return value
		}
	}
	return nil
}

type is_RcmRemoteaddress_RcmRemoteaddress interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RcmRemoteaddress_RcmRemoteaddress()
}

func (o *RcmRemoteaddress_RcmRemoteaddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Family)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	switch o.Family {
	case uint16(2):
		_o, _ := o.Value.(*RcmRemoteaddress_IPv4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RcmRemoteaddress_IPv4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(23):
		_o, _ := o.Value.(*RcmRemoteaddress_IPv6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RcmRemoteaddress_IPv6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Family)
	}
	return nil
}

func (o *RcmRemoteaddress_RcmRemoteaddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Family)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	switch o.Family {
	case uint16(2):
		o.Value = &RcmRemoteaddress_IPv4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(23):
		o.Value = &RcmRemoteaddress_IPv6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Family)
	}
	return nil
}

// RcmRemoteaddress_IPv4 structure represents RcmRemoteaddress_RcmRemoteaddress RPC union arm.
//
// It has following labels: 2
type RcmRemoteaddress_IPv4 struct {
	// sin_port:  Specifies a TCP or UDP port number.
	Port uint16 `idl:"name:sin_port" json:"port"`
	// in_addr:  Indicates the IP address.
	InAddr uint32 `idl:"name:in_addr" json:"in_addr"`
	// sin_zero:  An array filled with zeros.
	Zero []byte `idl:"name:sin_zero" json:"zero"`
}

func (*RcmRemoteaddress_IPv4) is_RcmRemoteaddress_RcmRemoteaddress() {}

func (o *RcmRemoteaddress_IPv4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if err := w.WriteData(o.InAddr); err != nil {
		return err
	}
	for i1 := range o.Zero {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Zero[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Zero); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *RcmRemoteaddress_IPv4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	if err := w.ReadData(&o.InAddr); err != nil {
		return err
	}
	o.Zero = make([]byte, 8)
	for i1 := range o.Zero {
		i1 := i1
		if err := w.ReadData(&o.Zero[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RcmRemoteaddress_IPv6 structure represents RcmRemoteaddress_RcmRemoteaddress RPC union arm.
//
// It has following labels: 23
type RcmRemoteaddress_IPv6 struct {
	// sin6_port:  Specifies a TCP or UDP port number.
	Port uint16 `idl:"name:sin6_port" json:"port"`
	// sin6_flowinfo:  IPv6 flow information.
	FlowInfo uint32 `idl:"name:sin6_flowinfo" json:"flow_info"`
	// sin6_addr:  Indicates the IP address.
	Addr []uint16 `idl:"name:sin6_addr" json:"addr"`
	// sin6_scope_id:  Set of interfaces for a scope. For more information about these interfaces,
	// see [MSDN-SOCKADDR_IN6].
	ScopeID uint32 `idl:"name:sin6_scope_id" json:"scope_id"`
}

func (*RcmRemoteaddress_IPv6) is_RcmRemoteaddress_RcmRemoteaddress() {}

func (o *RcmRemoteaddress_IPv6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if err := w.WriteData(o.FlowInfo); err != nil {
		return err
	}
	for i1 := range o.Addr {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Addr[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Addr); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ScopeID); err != nil {
		return err
	}
	return nil
}
func (o *RcmRemoteaddress_IPv6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	if err := w.ReadData(&o.FlowInfo); err != nil {
		return err
	}
	o.Addr = make([]uint16, 8)
	for i1 := range o.Addr {
		i1 := i1
		if err := w.ReadData(&o.Addr[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.ScopeID); err != nil {
		return err
	}
	return nil
}

// ListenerenumLevel1 structure represents LISTENERENUM_LEVEL1 RPC structure.
//
// LISTENERENUM_LEVEL1 is a structure containing information of level 1 detail about
// a Terminal Services listener.
type ListenerenumLevel1 struct {
	// Id:  The identifier associated with the listener.
	ID int32 `idl:"name:Id" json:"id"`
	// bListening:   Set to TRUE if the listener is listening for incoming connections;
	// FALSE otherwise.
	Listening bool `idl:"name:bListening" json:"listening"`
	// Name:  A string that contains the name of the listener followed by the terminating
	// NULL character.
	Name []uint16 `idl:"name:Name" json:"name"`
}

func (o *ListenerenumLevel1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ListenerenumLevel1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if !o.Listening {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
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
func (o *ListenerenumLevel1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	var _bListening int32
	if err := w.ReadData(&_bListening); err != nil {
		return err
	}
	o.Listening = _bListening != 0
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

// ListenerInfo structure represents ListenerInfo RPC union.
type ListenerInfo struct {
	// Types that are assignable to Value
	//
	// *ListenerInfo_ListenerEnumLevel1
	Value is_ListenerInfo `json:"value"`
}

func (o *ListenerInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ListenerInfo_ListenerEnumLevel1:
		if value != nil {
			return value.ListenerEnumLevel1
		}
	}
	return nil
}

type is_ListenerInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ListenerInfo()
}

func (o *ListenerInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ListenerInfo_ListenerEnumLevel1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *ListenerInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*ListenerInfo_ListenerEnumLevel1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ListenerInfo_ListenerEnumLevel1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *ListenerInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &ListenerInfo_ListenerEnumLevel1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// ListenerInfo_ListenerEnumLevel1 structure represents ListenerInfo RPC union arm.
//
// It has following labels: 1
type ListenerInfo_ListenerEnumLevel1 struct {
	ListenerEnumLevel1 *ListenerenumLevel1 `idl:"name:ListenerEnum_Level1" json:"listener_enum_level1"`
}

func (*ListenerInfo_ListenerEnumLevel1) is_ListenerInfo() {}

func (o *ListenerInfo_ListenerEnumLevel1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ListenerEnumLevel1 != nil {
		if err := o.ListenerEnumLevel1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ListenerenumLevel1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ListenerInfo_ListenerEnumLevel1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ListenerEnumLevel1 == nil {
		o.ListenerEnumLevel1 = &ListenerenumLevel1{}
	}
	if err := o.ListenerEnumLevel1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Listenerenum structure represents LISTENERENUM RPC structure.
//
// PLISTENERENUM contains information about one terminal server listener and the level
// of detail of the information provided.
type Listenerenum struct {
	// Level:  The level of detail provided about the listener. The only supported value
	// is 1.
	Level uint32 `idl:"name:Level" json:"level"`
	// Data:  Information about the listener. This is of the type ListenerInfo.
	Data *ListenerInfo `idl:"name:Data;switch_is:Level" json:"data"`
}

func (o *Listenerenum) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Listenerenum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&ListenerInfo{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
	}
	return nil
}
func (o *Listenerenum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.Data == nil {
		o.Data = &ListenerInfo{}
	}
	_swData := uint32(o.Level)
	if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
		return err
	}
	return nil
}

// ProtocolstatusInfoType type represents PROTOCOLSTATUS_INFO_TYPE RPC enumeration.
//
// The PROTOCOLSTATUS_INFO_TYPE enumeration specifies the protocol status information
// requested for a particular session running on a terminal server.
type ProtocolstatusInfoType uint16

var (
	// PROTOCOLSTATUS_INFO_BASIC:  Returns basic information about the protocol status
	// in a PROTOCOLSTATUS structure.
	ProtocolstatusInfoTypeBasic ProtocolstatusInfoType = 0
	// PROTOCOLSTATUS_INFO_EXTENDED:  Returns extended information about the protocol status.
	// Extended information is returned in a PROTOCOLSTATUSEX structure.
	ProtocolstatusInfoTypeExtended ProtocolstatusInfoType = 1
)

func (o ProtocolstatusInfoType) String() string {
	switch o {
	case ProtocolstatusInfoTypeBasic:
		return "ProtocolstatusInfoTypeBasic"
	case ProtocolstatusInfoTypeExtended:
		return "ProtocolstatusInfoTypeExtended"
	}
	return "Invalid"
}

// QuerySessionDataType type represents QUERY_SESSION_DATA_TYPE RPC enumeration.
//
// The QUERY_SESSION_DATA_TYPE enumeration specifies the type of session information
// that can be requested for a particular session running on a terminal server.
type QuerySessionDataType uint16

var (
	// QUERY_SESSION_DATA_MODULE:  Retrieves data about protocol-specific binaries loaded
	// for the given Terminal Services session. The type of the data is PBYTE.
	QuerySessionDataTypeModule QuerySessionDataType = 0
	// QUERY_SESSION_DATA_WDCONFIG:  Retrieves protocol driver configuration data for the
	// session. The data returned is of type WDCONFIG.
	QuerySessionDataTypeWdconfig QuerySessionDataType = 1
	// QUERY_SESSION_DATA_VIRTUALDATA:  Retrieves data about virtual channels for the given
	// Terminal Services session. The data returned is of type PBYTE.
	QuerySessionDataTypeVirtualdata QuerySessionDataType = 2
	// QUERY_SESSION_DATA_LICENSE:  Retrieves data about the licensing policies associated
	// with a given Terminal Services session.<43><44>
	QuerySessionDataTypeLicense QuerySessionDataType = 3
	// QUERY_SESSION_DATA_DEVICEID:  Retrieves the device ID of the client connected to
	// a given Terminal Services session. The data returned is of type PBYTE.<45>
	QuerySessionDataTypeDeviceid QuerySessionDataType = 4
	// QUERY_SESSION_DATA_LICENSE_VALIDATION:  Retrieves the data required to validate
	// the license associated with a given Terminal Services session. The data returned
	// is of type WINSTATIONVALIDATIONINFORMATION.<46>
	QuerySessionDataTypeLicenseValidation QuerySessionDataType = 5
)

func (o QuerySessionDataType) String() string {
	switch o {
	case QuerySessionDataTypeModule:
		return "QuerySessionDataTypeModule"
	case QuerySessionDataTypeWdconfig:
		return "QuerySessionDataTypeWdconfig"
	case QuerySessionDataTypeVirtualdata:
		return "QuerySessionDataTypeVirtualdata"
	case QuerySessionDataTypeLicense:
		return "QuerySessionDataTypeLicense"
	case QuerySessionDataTypeDeviceid:
		return "QuerySessionDataTypeDeviceid"
	case QuerySessionDataTypeLicenseValidation:
		return "QuerySessionDataTypeLicenseValidation"
	}
	return "Invalid"
}

type xxx_DefaultRcmPublicClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRcmPublicClient) GetClientData(ctx context.Context, in *GetClientDataRequest, opts ...dcerpc.CallOption) (*GetClientDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClientDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) GetConfigData(ctx context.Context, in *GetConfigDataRequest, opts ...dcerpc.CallOption) (*GetConfigDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetConfigDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) GetProtocolStatus(ctx context.Context, in *GetProtocolStatusRequest, opts ...dcerpc.CallOption) (*GetProtocolStatusResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetProtocolStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) GetLastInputTime(ctx context.Context, in *GetLastInputTimeRequest, opts ...dcerpc.CallOption) (*GetLastInputTimeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetLastInputTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) GetRemoteAddress(ctx context.Context, in *GetRemoteAddressRequest, opts ...dcerpc.CallOption) (*GetRemoteAddressResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetRemoteAddressResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) GetAllListeners(ctx context.Context, in *GetAllListenersRequest, opts ...dcerpc.CallOption) (*GetAllListenersResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetAllListenersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) GetSessionProtocolLastInputTime(ctx context.Context, in *GetSessionProtocolLastInputTimeRequest, opts ...dcerpc.CallOption) (*GetSessionProtocolLastInputTimeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSessionProtocolLastInputTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) GetUserCertificates(ctx context.Context, in *GetUserCertificatesRequest, opts ...dcerpc.CallOption) (*GetUserCertificatesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetUserCertificatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) QuerySessionData(ctx context.Context, in *QuerySessionDataRequest, opts ...dcerpc.CallOption) (*QuerySessionDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QuerySessionDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmPublicClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRcmPublicClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewRcmPublicClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RcmPublicClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RcmPublicSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRcmPublicClient{cc: cc}, nil
}

// xxx_GetClientDataOperation structure represents the RpcGetClientData operation
type xxx_GetClientDataOperation struct {
	SessionID           uint32 `idl:"name:SessionId" json:"session_id"`
	Buffer              []byte `idl:"name:ppBuff;size_is:(, pOutBuffByteLen)" json:"buffer"`
	OutBufferByteLength uint32 `idl:"name:pOutBuffByteLen" json:"out_buffer_byte_length"`
	Return              int32  `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetClientData operation.
func (o *xxx_GetClientDataOperation) OpNum() int { return 0 }

// OpName returns the operation name of RpcGetClientData operation.
func (o *xxx_GetClientDataOperation) OpName() string { return "/RCMPublic/v1/RpcGetClientData" }

func (o *xxx_GetClientDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.OutBufferByteLength == 0 {
		o.OutBufferByteLength = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppBuff {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutBuffByteLen](uchar))
	{
		if o.Buffer != nil || o.OutBufferByteLength > 0 {
			_ptr_ppBuff := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.OutBufferByteLength)
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
					if err := w.WriteData(o.Buffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_ppBuff); err != nil {
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
	// pOutBuffByteLen {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.OutBufferByteLength); err != nil {
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

func (o *xxx_GetClientDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppBuff {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutBuffByteLen](uchar))
	{
		_ptr_ppBuff := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
			}
			o.Buffer = make([]byte, sizeInfo[0])
			for i1 := range o.Buffer {
				i1 := i1
				if err := w.ReadData(&o.Buffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppBuff := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_ppBuff, _ptr_ppBuff); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pOutBuffByteLen {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.OutBufferByteLength); err != nil {
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

// GetClientDataRequest structure represents the RpcGetClientData operation request
type GetClientDataRequest struct {
	// SessionId:  The ID of the session from which client data is to be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
}

func (o *GetClientDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClientDataOperation) *xxx_GetClientDataOperation {
	if op == nil {
		op = &xxx_GetClientDataOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	return op
}

func (o *GetClientDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClientDataOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
}
func (o *GetClientDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClientDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetClientDataRequest build a response structure from the given request structure.
func (o *GetClientDataRequest) MakeResponse() *GetClientDataResponse {
	return &GetClientDataResponse{}
}

// OpNum returns the operation number of RpcGetClientData operation.
func (o *GetClientDataRequest) OpNum() int { return 0 }

// OpName returns the operation name of RpcGetClientData operation.
func (o *GetClientDataRequest) OpName() string { return "/RCMPublic/v1/RpcGetClientData" }

// GetClientDataResponse structure represents the RpcGetClientData operation response
type GetClientDataResponse struct {
	// ppBuff:  The buffer that contains the client data. This data is of type PWINSTATIONCLIENT,
	// specified in section 2.2.2.19. The buffer is allocated by RpcGetClientData.
	Buffer []byte `idl:"name:ppBuff;size_is:(, pOutBuffByteLen)" json:"buffer"`
	// pOutBuffByteLen:  The number of bytes of client data that is returned.
	OutBufferByteLength uint32 `idl:"name:pOutBuffByteLen" json:"out_buffer_byte_length"`
	// Return: The RpcGetClientData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClientDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClientDataOperation) *xxx_GetClientDataOperation {
	if op == nil {
		op = &xxx_GetClientDataOperation{}
	}
	if o == nil {
		return op
	}
	op.Buffer = o.Buffer
	op.OutBufferByteLength = o.OutBufferByteLength
	op.Return = o.Return
	return op
}

func (o *GetClientDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClientDataOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.OutBufferByteLength = op.OutBufferByteLength
	o.Return = op.Return
}
func (o *GetClientDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClientDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetConfigDataOperation structure represents the RpcGetConfigData operation
type xxx_GetConfigDataOperation struct {
	SessionID           uint32 `idl:"name:SessionId" json:"session_id"`
	Buffer              []byte `idl:"name:ppBuff;size_is:(, pOutBuffByteLen)" json:"buffer"`
	OutBufferByteLength uint32 `idl:"name:pOutBuffByteLen" json:"out_buffer_byte_length"`
	Return              int32  `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetConfigData operation.
func (o *xxx_GetConfigDataOperation) OpNum() int { return 1 }

// OpName returns the operation name of RpcGetConfigData operation.
func (o *xxx_GetConfigDataOperation) OpName() string { return "/RCMPublic/v1/RpcGetConfigData" }

func (o *xxx_GetConfigDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.OutBufferByteLength == 0 {
		o.OutBufferByteLength = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppBuff {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutBuffByteLen](uchar))
	{
		if o.Buffer != nil || o.OutBufferByteLength > 0 {
			_ptr_ppBuff := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.OutBufferByteLength)
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
					if err := w.WriteData(o.Buffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_ppBuff); err != nil {
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
	// pOutBuffByteLen {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.OutBufferByteLength); err != nil {
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

func (o *xxx_GetConfigDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppBuff {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutBuffByteLen](uchar))
	{
		_ptr_ppBuff := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
			}
			o.Buffer = make([]byte, sizeInfo[0])
			for i1 := range o.Buffer {
				i1 := i1
				if err := w.ReadData(&o.Buffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppBuff := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_ppBuff, _ptr_ppBuff); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pOutBuffByteLen {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.OutBufferByteLength); err != nil {
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

// GetConfigDataRequest structure represents the RpcGetConfigData operation request
type GetConfigDataRequest struct {
	// SessionId:  The ID of the session for which the client configuration data is to
	// be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
}

func (o *GetConfigDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConfigDataOperation) *xxx_GetConfigDataOperation {
	if op == nil {
		op = &xxx_GetConfigDataOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	return op
}

func (o *GetConfigDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConfigDataOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
}
func (o *GetConfigDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConfigDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetConfigDataRequest build a response structure from the given request structure.
func (o *GetConfigDataRequest) MakeResponse() *GetConfigDataResponse {
	return &GetConfigDataResponse{}
}

// OpNum returns the operation number of RpcGetConfigData operation.
func (o *GetConfigDataRequest) OpNum() int { return 1 }

// OpName returns the operation name of RpcGetConfigData operation.
func (o *GetConfigDataRequest) OpName() string { return "/RCMPublic/v1/RpcGetConfigData" }

// GetConfigDataResponse structure represents the RpcGetConfigData operation response
type GetConfigDataResponse struct {
	// ppBuff:  The buffer that will contain the client configuration data. This will be
	// allocated by RpcGetConfigData. This data is of type PWINSTATIONCONFIG.
	Buffer []byte `idl:"name:ppBuff;size_is:(, pOutBuffByteLen)" json:"buffer"`
	// pOutBuffByteLen:  The number of bytes of client configuration data that is returned.
	OutBufferByteLength uint32 `idl:"name:pOutBuffByteLen" json:"out_buffer_byte_length"`
	// Return: The RpcGetConfigData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConfigDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConfigDataOperation) *xxx_GetConfigDataOperation {
	if op == nil {
		op = &xxx_GetConfigDataOperation{}
	}
	if o == nil {
		return op
	}
	op.Buffer = o.Buffer
	op.OutBufferByteLength = o.OutBufferByteLength
	op.Return = o.Return
	return op
}

func (o *GetConfigDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConfigDataOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.OutBufferByteLength = op.OutBufferByteLength
	o.Return = op.Return
}
func (o *GetConfigDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConfigDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetProtocolStatusOperation structure represents the RpcGetProtocolStatus operation
type xxx_GetProtocolStatusOperation struct {
	SessionID         uint32                 `idl:"name:SessionId" json:"session_id"`
	InfoType          ProtocolstatusInfoType `idl:"name:InfoType" json:"info_type"`
	ProtoStatus       []byte                 `idl:"name:ppProtoStatus;size_is:(, pcbProtoStatus)" json:"proto_status"`
	ProtoStatusLength uint32                 `idl:"name:pcbProtoStatus" json:"proto_status_length"`
	Return            int32                  `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetProtocolStatus operation.
func (o *xxx_GetProtocolStatusOperation) OpNum() int { return 2 }

// OpName returns the operation name of RpcGetProtocolStatus operation.
func (o *xxx_GetProtocolStatusOperation) OpName() string { return "/RCMPublic/v1/RpcGetProtocolStatus" }

func (o *xxx_GetProtocolStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProtocolStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	// InfoType {in} (1:{alias=PROTOCOLSTATUS_INFO_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.InfoType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProtocolStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	// InfoType {in} (1:{alias=PROTOCOLSTATUS_INFO_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InfoType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProtocolStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ProtoStatus != nil && o.ProtoStatusLength == 0 {
		o.ProtoStatusLength = uint32(len(o.ProtoStatus))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProtocolStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppProtoStatus {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbProtoStatus](uchar))
	{
		if o.ProtoStatus != nil || o.ProtoStatusLength > 0 {
			_ptr_ppProtoStatus := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ProtoStatusLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ProtoStatus {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ProtoStatus[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ProtoStatus); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProtoStatus, _ptr_ppProtoStatus); err != nil {
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
	// pcbProtoStatus {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ProtoStatusLength); err != nil {
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

func (o *xxx_GetProtocolStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppProtoStatus {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbProtoStatus](uchar))
	{
		_ptr_ppProtoStatus := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ProtoStatus", sizeInfo[0])
			}
			o.ProtoStatus = make([]byte, sizeInfo[0])
			for i1 := range o.ProtoStatus {
				i1 := i1
				if err := w.ReadData(&o.ProtoStatus[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppProtoStatus := func(ptr interface{}) { o.ProtoStatus = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ProtoStatus, _s_ppProtoStatus, _ptr_ppProtoStatus); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbProtoStatus {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ProtoStatusLength); err != nil {
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

// GetProtocolStatusRequest structure represents the RpcGetProtocolStatus operation request
type GetProtocolStatusRequest struct {
	// SessionId:  The ID of the session for which protocol status is to be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
	// InfoType:  Specifies what type of information to gather. This is of the type PROTOCOLSTATUS_INFO_TYPE.
	InfoType ProtocolstatusInfoType `idl:"name:InfoType" json:"info_type"`
}

func (o *GetProtocolStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_GetProtocolStatusOperation) *xxx_GetProtocolStatusOperation {
	if op == nil {
		op = &xxx_GetProtocolStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	op.InfoType = o.InfoType
	return op
}

func (o *GetProtocolStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetProtocolStatusOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
	o.InfoType = op.InfoType
}
func (o *GetProtocolStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetProtocolStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProtocolStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetProtocolStatusRequest build a response structure from the given request structure.
func (o *GetProtocolStatusRequest) MakeResponse() *GetProtocolStatusResponse {
	return &GetProtocolStatusResponse{}
}

// OpNum returns the operation number of RpcGetProtocolStatus operation.
func (o *GetProtocolStatusRequest) OpNum() int { return 2 }

// OpName returns the operation name of RpcGetProtocolStatus operation.
func (o *GetProtocolStatusRequest) OpName() string { return "/RCMPublic/v1/RpcGetProtocolStatus" }

// GetProtocolStatusResponse structure represents the RpcGetProtocolStatus operation response
type GetProtocolStatusResponse struct {
	// ppProtoStatus:  The buffer that will contain protocol status data. This data is
	// of the type PROTOCOLSTATUS.
	ProtoStatus []byte `idl:"name:ppProtoStatus;size_is:(, pcbProtoStatus)" json:"proto_status"`
	// pcbProtoStatus:  The number of bytes of protocol data that is returned.
	ProtoStatusLength uint32 `idl:"name:pcbProtoStatus" json:"proto_status_length"`
	// Return: The RpcGetProtocolStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProtocolStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_GetProtocolStatusOperation) *xxx_GetProtocolStatusOperation {
	if op == nil {
		op = &xxx_GetProtocolStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.ProtoStatus = o.ProtoStatus
	op.ProtoStatusLength = o.ProtoStatusLength
	op.Return = o.Return
	return op
}

func (o *GetProtocolStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetProtocolStatusOperation) {
	if o == nil {
		return
	}
	o.ProtoStatus = op.ProtoStatus
	o.ProtoStatusLength = op.ProtoStatusLength
	o.Return = op.Return
}
func (o *GetProtocolStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetProtocolStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProtocolStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLastInputTimeOperation structure represents the RpcGetLastInputTime operation
type xxx_GetLastInputTimeOperation struct {
	SessionID     uint32 `idl:"name:SessionId" json:"session_id"`
	LastInputTime int64  `idl:"name:pLastInputTime" json:"last_input_time"`
	Return        int32  `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetLastInputTime operation.
func (o *xxx_GetLastInputTimeOperation) OpNum() int { return 3 }

// OpName returns the operation name of RpcGetLastInputTime operation.
func (o *xxx_GetLastInputTimeOperation) OpName() string { return "/RCMPublic/v1/RpcGetLastInputTime" }

func (o *xxx_GetLastInputTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastInputTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastInputTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastInputTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastInputTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pLastInputTime {out} (1:{pointer=ref}*(1)(int64))
	{
		if err := w.WriteData(o.LastInputTime); err != nil {
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

func (o *xxx_GetLastInputTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pLastInputTime {out} (1:{pointer=ref}*(1)(int64))
	{
		if err := w.ReadData(&o.LastInputTime); err != nil {
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

// GetLastInputTimeRequest structure represents the RpcGetLastInputTime operation request
type GetLastInputTimeRequest struct {
	// SessionId:  The ID of the session for which the last user input time is to be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
}

func (o *GetLastInputTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastInputTimeOperation) *xxx_GetLastInputTimeOperation {
	if op == nil {
		op = &xxx_GetLastInputTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	return op
}

func (o *GetLastInputTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastInputTimeOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
}
func (o *GetLastInputTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastInputTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastInputTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetLastInputTimeRequest build a response structure from the given request structure.
func (o *GetLastInputTimeRequest) MakeResponse() *GetLastInputTimeResponse {
	return &GetLastInputTimeResponse{}
}

// OpNum returns the operation number of RpcGetLastInputTime operation.
func (o *GetLastInputTimeRequest) OpNum() int { return 3 }

// OpName returns the operation name of RpcGetLastInputTime operation.
func (o *GetLastInputTimeRequest) OpName() string { return "/RCMPublic/v1/RpcGetLastInputTime" }

// GetLastInputTimeResponse structure represents the RpcGetLastInputTime operation response
type GetLastInputTimeResponse struct {
	// pLastInputTime:  The time when the last user input was received by the server. This
	// is a 64-bit value representing the number of 100-nanosecond intervals since January
	// 1, 1601 (UTC).
	LastInputTime int64 `idl:"name:pLastInputTime" json:"last_input_time"`
	// Return: The RpcGetLastInputTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastInputTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastInputTimeOperation) *xxx_GetLastInputTimeOperation {
	if op == nil {
		op = &xxx_GetLastInputTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.LastInputTime = o.LastInputTime
	op.Return = o.Return
	return op
}

func (o *GetLastInputTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastInputTimeOperation) {
	if o == nil {
		return
	}
	o.LastInputTime = op.LastInputTime
	o.Return = op.Return
}
func (o *GetLastInputTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastInputTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastInputTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRemoteAddressOperation structure represents the RpcGetRemoteAddress operation
type xxx_GetRemoteAddressOperation struct {
	SessionID     uint32            `idl:"name:SessionId" json:"session_id"`
	RemoteAddress *RcmRemoteaddress `idl:"name:pRemoteAddress" json:"remote_address"`
	Return        int32             `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetRemoteAddress operation.
func (o *xxx_GetRemoteAddressOperation) OpNum() int { return 4 }

// OpName returns the operation name of RpcGetRemoteAddress operation.
func (o *xxx_GetRemoteAddressOperation) OpName() string { return "/RCMPublic/v1/RpcGetRemoteAddress" }

func (o *xxx_GetRemoteAddressOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteAddressOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteAddressOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteAddressOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteAddressOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pRemoteAddress {out} (1:{alias=PRCM_REMOTEADDRESS}*(1))(2:{alias=RCM_REMOTEADDRESS}(struct))
	{
		if o.RemoteAddress != nil {
			if err := o.RemoteAddress.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RcmRemoteaddress{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetRemoteAddressOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pRemoteAddress {out} (1:{alias=PRCM_REMOTEADDRESS,pointer=ref}*(1))(2:{alias=RCM_REMOTEADDRESS}(struct))
	{
		if o.RemoteAddress == nil {
			o.RemoteAddress = &RcmRemoteaddress{}
		}
		if err := o.RemoteAddress.UnmarshalNDR(ctx, w); err != nil {
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

// GetRemoteAddressRequest structure represents the RpcGetRemoteAddress operation request
type GetRemoteAddressRequest struct {
	// SessionId:  The ID of the session for which client data is to be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
}

func (o *GetRemoteAddressRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRemoteAddressOperation) *xxx_GetRemoteAddressOperation {
	if op == nil {
		op = &xxx_GetRemoteAddressOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	return op
}

func (o *GetRemoteAddressRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteAddressOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
}
func (o *GetRemoteAddressRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRemoteAddressRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteAddressOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetRemoteAddressRequest build a response structure from the given request structure.
func (o *GetRemoteAddressRequest) MakeResponse() *GetRemoteAddressResponse {
	return &GetRemoteAddressResponse{}
}

// OpNum returns the operation number of RpcGetRemoteAddress operation.
func (o *GetRemoteAddressRequest) OpNum() int { return 4 }

// OpName returns the operation name of RpcGetRemoteAddress operation.
func (o *GetRemoteAddressRequest) OpName() string { return "/RCMPublic/v1/RpcGetRemoteAddress" }

// GetRemoteAddressResponse structure represents the RpcGetRemoteAddress operation response
type GetRemoteAddressResponse struct {
	// pRemoteAddress: The address of the client computer that is connected to the session.
	// This is of the type PRCM_REMOTEADDRESS.
	RemoteAddress *RcmRemoteaddress `idl:"name:pRemoteAddress" json:"remote_address"`
	// Return: The RpcGetRemoteAddress return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRemoteAddressResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRemoteAddressOperation) *xxx_GetRemoteAddressOperation {
	if op == nil {
		op = &xxx_GetRemoteAddressOperation{}
	}
	if o == nil {
		return op
	}
	op.RemoteAddress = o.RemoteAddress
	op.Return = o.Return
	return op
}

func (o *GetRemoteAddressResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteAddressOperation) {
	if o == nil {
		return
	}
	o.RemoteAddress = op.RemoteAddress
	o.Return = op.Return
}
func (o *GetRemoteAddressResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRemoteAddressResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteAddressOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllListenersOperation structure represents the RpcGetAllListeners operation
type xxx_GetAllListenersOperation struct {
	Listeners    []*Listenerenum `idl:"name:ppListeners;size_is:(, pNumListeners)" json:"listeners"`
	Level        uint32          `idl:"name:Level" json:"level"`
	NumListeners uint32          `idl:"name:pNumListeners" json:"num_listeners"`
	Return       int32           `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetAllListeners operation.
func (o *xxx_GetAllListenersOperation) OpNum() int { return 8 }

// OpName returns the operation name of RpcGetAllListeners operation.
func (o *xxx_GetAllListenersOperation) OpName() string { return "/RCMPublic/v1/RpcGetAllListeners" }

func (o *xxx_GetAllListenersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllListenersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllListenersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllListenersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Listeners != nil && o.NumListeners == 0 {
		o.NumListeners = uint32(len(o.Listeners))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllListenersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppListeners {out} (1:{pointer=ref}*(2))(2:{alias=PLISTENERENUM}*(1))(3:{alias=LISTENERENUM}[dim:0,size_is=pNumListeners](struct))
	{
		if o.Listeners != nil || o.NumListeners > 0 {
			_ptr_ppListeners := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumListeners)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Listeners {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Listeners[i1] != nil {
						if err := o.Listeners[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&Listenerenum{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Listeners); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&Listenerenum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Listeners, _ptr_ppListeners); err != nil {
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
	// pNumListeners {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.NumListeners); err != nil {
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

func (o *xxx_GetAllListenersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppListeners {out} (1:{pointer=ref}*(2))(2:{alias=PLISTENERENUM,pointer=ref}*(1))(3:{alias=LISTENERENUM}[dim:0,size_is=pNumListeners](struct))
	{
		_ptr_ppListeners := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Listeners", sizeInfo[0])
			}
			o.Listeners = make([]*Listenerenum, sizeInfo[0])
			for i1 := range o.Listeners {
				i1 := i1
				if o.Listeners[i1] == nil {
					o.Listeners[i1] = &Listenerenum{}
				}
				if err := o.Listeners[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppListeners := func(ptr interface{}) { o.Listeners = *ptr.(*[]*Listenerenum) }
		if err := w.ReadPointer(&o.Listeners, _s_ppListeners, _ptr_ppListeners); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pNumListeners {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.NumListeners); err != nil {
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

// GetAllListenersRequest structure represents the RpcGetAllListeners operation request
type GetAllListenersRequest struct {
	// Level:  The level of information that is requested for the listeners. The only supported
	// value is 1.
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *GetAllListenersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAllListenersOperation) *xxx_GetAllListenersOperation {
	if op == nil {
		op = &xxx_GetAllListenersOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	return op
}

func (o *GetAllListenersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllListenersOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
}
func (o *GetAllListenersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAllListenersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllListenersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetAllListenersRequest build a response structure from the given request structure.
func (o *GetAllListenersRequest) MakeResponse() *GetAllListenersResponse {
	return &GetAllListenersResponse{}
}

// OpNum returns the operation number of RpcGetAllListeners operation.
func (o *GetAllListenersRequest) OpNum() int { return 8 }

// OpName returns the operation name of RpcGetAllListeners operation.
func (o *GetAllListenersRequest) OpName() string { return "/RCMPublic/v1/RpcGetAllListeners" }

// GetAllListenersResponse structure represents the RpcGetAllListeners operation response
type GetAllListenersResponse struct {
	// ppListeners:  The list of Terminal Services listeners running on the terminal server.
	// This is an array of type PLISTENERENUM.
	Listeners []*Listenerenum `idl:"name:ppListeners;size_is:(, pNumListeners)" json:"listeners"`
	// pNumListeners: The number of listeners returned.
	NumListeners uint32 `idl:"name:pNumListeners" json:"num_listeners"`
	// Return: The RpcGetAllListeners return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllListenersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAllListenersOperation) *xxx_GetAllListenersOperation {
	if op == nil {
		op = &xxx_GetAllListenersOperation{}
	}
	if o == nil {
		return op
	}
	op.Listeners = o.Listeners
	op.NumListeners = o.NumListeners
	op.Return = o.Return
	return op
}

func (o *GetAllListenersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllListenersOperation) {
	if o == nil {
		return
	}
	o.Listeners = op.Listeners
	o.NumListeners = op.NumListeners
	o.Return = op.Return
}
func (o *GetAllListenersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAllListenersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllListenersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSessionProtocolLastInputTimeOperation structure represents the RpcGetSessionProtocolLastInputTime operation
type xxx_GetSessionProtocolLastInputTimeOperation struct {
	SessionID         uint32                 `idl:"name:SessionId" json:"session_id"`
	InfoType          ProtocolstatusInfoType `idl:"name:InfoType" json:"info_type"`
	ProtoStatus       []byte                 `idl:"name:ppProtoStatus;size_is:(, pcbProtoStatus)" json:"proto_status"`
	ProtoStatusLength uint32                 `idl:"name:pcbProtoStatus" json:"proto_status_length"`
	LastInputTime     int64                  `idl:"name:pLastInputTime" json:"last_input_time"`
	Return            int32                  `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetSessionProtocolLastInputTime operation.
func (o *xxx_GetSessionProtocolLastInputTimeOperation) OpNum() int { return 9 }

// OpName returns the operation name of RpcGetSessionProtocolLastInputTime operation.
func (o *xxx_GetSessionProtocolLastInputTimeOperation) OpName() string {
	return "/RCMPublic/v1/RpcGetSessionProtocolLastInputTime"
}

func (o *xxx_GetSessionProtocolLastInputTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionProtocolLastInputTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	// InfoType {in} (1:{alias=PROTOCOLSTATUS_INFO_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.InfoType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionProtocolLastInputTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	// InfoType {in} (1:{alias=PROTOCOLSTATUS_INFO_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InfoType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionProtocolLastInputTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ProtoStatus != nil && o.ProtoStatusLength == 0 {
		o.ProtoStatusLength = uint32(len(o.ProtoStatus))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSessionProtocolLastInputTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppProtoStatus {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbProtoStatus](uchar))
	{
		if o.ProtoStatus != nil || o.ProtoStatusLength > 0 {
			_ptr_ppProtoStatus := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ProtoStatusLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ProtoStatus {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ProtoStatus[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ProtoStatus); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProtoStatus, _ptr_ppProtoStatus); err != nil {
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
	// pcbProtoStatus {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ProtoStatusLength); err != nil {
			return err
		}
	}
	// pLastInputTime {out} (1:{pointer=ref}*(1)(int64))
	{
		if err := w.WriteData(o.LastInputTime); err != nil {
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

func (o *xxx_GetSessionProtocolLastInputTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppProtoStatus {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbProtoStatus](uchar))
	{
		_ptr_ppProtoStatus := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ProtoStatus", sizeInfo[0])
			}
			o.ProtoStatus = make([]byte, sizeInfo[0])
			for i1 := range o.ProtoStatus {
				i1 := i1
				if err := w.ReadData(&o.ProtoStatus[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppProtoStatus := func(ptr interface{}) { o.ProtoStatus = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ProtoStatus, _s_ppProtoStatus, _ptr_ppProtoStatus); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbProtoStatus {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ProtoStatusLength); err != nil {
			return err
		}
	}
	// pLastInputTime {out} (1:{pointer=ref}*(1)(int64))
	{
		if err := w.ReadData(&o.LastInputTime); err != nil {
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

// GetSessionProtocolLastInputTimeRequest structure represents the RpcGetSessionProtocolLastInputTime operation request
type GetSessionProtocolLastInputTimeRequest struct {
	// SessionId:  The ID of the session from which information is to be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
	// InfoType:  Specifies what type of information to gather. This is of type PROTOCOLSTATUS_INFO_TYPE.
	InfoType ProtocolstatusInfoType `idl:"name:InfoType" json:"info_type"`
}

func (o *GetSessionProtocolLastInputTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSessionProtocolLastInputTimeOperation) *xxx_GetSessionProtocolLastInputTimeOperation {
	if op == nil {
		op = &xxx_GetSessionProtocolLastInputTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	op.InfoType = o.InfoType
	return op
}

func (o *GetSessionProtocolLastInputTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSessionProtocolLastInputTimeOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
	o.InfoType = op.InfoType
}
func (o *GetSessionProtocolLastInputTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSessionProtocolLastInputTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionProtocolLastInputTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetSessionProtocolLastInputTimeRequest build a response structure from the given request structure.
func (o *GetSessionProtocolLastInputTimeRequest) MakeResponse() *GetSessionProtocolLastInputTimeResponse {
	return &GetSessionProtocolLastInputTimeResponse{}
}

// OpNum returns the operation number of RpcGetSessionProtocolLastInputTime operation.
func (o *GetSessionProtocolLastInputTimeRequest) OpNum() int { return 9 }

// OpName returns the operation name of RpcGetSessionProtocolLastInputTime operation.
func (o *GetSessionProtocolLastInputTimeRequest) OpName() string {
	return "/RCMPublic/v1/RpcGetSessionProtocolLastInputTime"
}

// GetSessionProtocolLastInputTimeResponse structure represents the RpcGetSessionProtocolLastInputTime operation response
type GetSessionProtocolLastInputTimeResponse struct {
	// ppProtoStatus:  The buffer that contains protocol status data. This data is of type
	// PROTOCOLSTATUS, specified in section 2.2.2.20.1.
	ProtoStatus []byte `idl:"name:ppProtoStatus;size_is:(, pcbProtoStatus)" json:"proto_status"`
	// pcbProtoStatus:  The number of bytes of protocol data returned.
	ProtoStatusLength uint32 `idl:"name:pcbProtoStatus" json:"proto_status_length"`
	// pLastInputTime:  The time the last input was received by the server.
	//
	// Time is measured as the number of 100-nanosecond intervals since January 1, 1601
	// (UTC).
	LastInputTime int64 `idl:"name:pLastInputTime" json:"last_input_time"`
	// Return: The RpcGetSessionProtocolLastInputTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSessionProtocolLastInputTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSessionProtocolLastInputTimeOperation) *xxx_GetSessionProtocolLastInputTimeOperation {
	if op == nil {
		op = &xxx_GetSessionProtocolLastInputTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.ProtoStatus = o.ProtoStatus
	op.ProtoStatusLength = o.ProtoStatusLength
	op.LastInputTime = o.LastInputTime
	op.Return = o.Return
	return op
}

func (o *GetSessionProtocolLastInputTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSessionProtocolLastInputTimeOperation) {
	if o == nil {
		return
	}
	o.ProtoStatus = op.ProtoStatus
	o.ProtoStatusLength = op.ProtoStatusLength
	o.LastInputTime = op.LastInputTime
	o.Return = op.Return
}
func (o *GetSessionProtocolLastInputTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSessionProtocolLastInputTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSessionProtocolLastInputTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUserCertificatesOperation structure represents the RpcGetUserCertificates operation
type xxx_GetUserCertificatesOperation struct {
	SessionID   uint32 `idl:"name:SessionId" json:"session_id"`
	CertsCount  uint32 `idl:"name:pcCerts" json:"certs_count"`
	Certs       []byte `idl:"name:ppbCerts;size_is:(, pcbCerts)" json:"certs"`
	CertsLength uint32 `idl:"name:pcbCerts" json:"certs_length"`
	Return      int32  `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcGetUserCertificates operation.
func (o *xxx_GetUserCertificatesOperation) OpNum() int { return 10 }

// OpName returns the operation name of RpcGetUserCertificates operation.
func (o *xxx_GetUserCertificatesOperation) OpName() string {
	return "/RCMPublic/v1/RpcGetUserCertificates"
}

func (o *xxx_GetUserCertificatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserCertificatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserCertificatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserCertificatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Certs != nil && o.CertsLength == 0 {
		o.CertsLength = uint32(len(o.Certs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserCertificatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcCerts {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.CertsCount); err != nil {
			return err
		}
	}
	// ppbCerts {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbCerts](uint8))
	{
		if o.Certs != nil || o.CertsLength > 0 {
			_ptr_ppbCerts := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.CertsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Certs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Certs[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Certs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Certs, _ptr_ppbCerts); err != nil {
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
	// pcbCerts {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.CertsLength); err != nil {
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

func (o *xxx_GetUserCertificatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcCerts {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.CertsCount); err != nil {
			return err
		}
	}
	// ppbCerts {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbCerts](uint8))
	{
		_ptr_ppbCerts := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Certs", sizeInfo[0])
			}
			o.Certs = make([]byte, sizeInfo[0])
			for i1 := range o.Certs {
				i1 := i1
				if err := w.ReadData(&o.Certs[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppbCerts := func(ptr interface{}) { o.Certs = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Certs, _s_ppbCerts, _ptr_ppbCerts); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbCerts {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.CertsLength); err != nil {
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

// GetUserCertificatesRequest structure represents the RpcGetUserCertificates operation request
type GetUserCertificatesRequest struct {
	// SessionId: The ID of the session for which the certificate is to be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
}

func (o *GetUserCertificatesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetUserCertificatesOperation) *xxx_GetUserCertificatesOperation {
	if op == nil {
		op = &xxx_GetUserCertificatesOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	return op
}

func (o *GetUserCertificatesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserCertificatesOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
}
func (o *GetUserCertificatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetUserCertificatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserCertificatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeGetUserCertificatesRequest build a response structure from the given request structure.
func (o *GetUserCertificatesRequest) MakeResponse() *GetUserCertificatesResponse {
	return &GetUserCertificatesResponse{}
}

// OpNum returns the operation number of RpcGetUserCertificates operation.
func (o *GetUserCertificatesRequest) OpNum() int { return 10 }

// OpName returns the operation name of RpcGetUserCertificates operation.
func (o *GetUserCertificatesRequest) OpName() string { return "/RCMPublic/v1/RpcGetUserCertificates" }

// GetUserCertificatesResponse structure represents the RpcGetUserCertificates operation response
type GetUserCertificatesResponse struct {
	// pcCerts: The number of client certificates returned.
	CertsCount uint32 `idl:"name:pcCerts" json:"certs_count"`
	// ppbCerts: Certificate data.
	Certs []byte `idl:"name:ppbCerts;size_is:(, pcbCerts)" json:"certs"`
	// pcbCerts: The size, in bytes, of ppbCerts.
	CertsLength uint32 `idl:"name:pcbCerts" json:"certs_length"`
	// Return: The RpcGetUserCertificates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUserCertificatesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetUserCertificatesOperation) *xxx_GetUserCertificatesOperation {
	if op == nil {
		op = &xxx_GetUserCertificatesOperation{}
	}
	if o == nil {
		return op
	}
	op.CertsCount = o.CertsCount
	op.Certs = o.Certs
	op.CertsLength = o.CertsLength
	op.Return = o.Return
	return op
}

func (o *GetUserCertificatesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserCertificatesOperation) {
	if o == nil {
		return
	}
	o.CertsCount = op.CertsCount
	o.Certs = op.Certs
	o.CertsLength = op.CertsLength
	o.Return = op.Return
}
func (o *GetUserCertificatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetUserCertificatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserCertificatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QuerySessionDataOperation structure represents the RpcQuerySessionData operation
type xxx_QuerySessionDataOperation struct {
	SessionID           uint32               `idl:"name:SessionId" json:"session_id"`
	Type                QuerySessionDataType `idl:"name:type" json:"type"`
	InputData           []byte               `idl:"name:pbInputData;size_is:(cbInputData);pointer:unique" json:"input_data"`
	InputDataLength     uint32               `idl:"name:cbInputData" json:"input_data_length"`
	SessionData         []byte               `idl:"name:pbSessionData;size_is:(cbSessionData);length_is:(pcbReturnLength);pointer:ref" json:"session_data"`
	SessionDataLength   uint32               `idl:"name:cbSessionData" json:"session_data_length"`
	ReturnLength        uint32               `idl:"name:pcbReturnLength;pointer:ref" json:"return_length"`
	RequireBufferLength uint32               `idl:"name:pcbRequireBufferSize;pointer:ref" json:"require_buffer_length"`
	Return              int32                `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcQuerySessionData operation.
func (o *xxx_QuerySessionDataOperation) OpNum() int { return 11 }

// OpName returns the operation name of RpcQuerySessionData operation.
func (o *xxx_QuerySessionDataOperation) OpName() string { return "/RCMPublic/v1/RpcQuerySessionData" }

func (o *xxx_QuerySessionDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InputData != nil && o.InputDataLength == 0 {
		o.InputDataLength = uint32(len(o.InputData))
	}
	if o.InputDataLength > uint32(8192) {
		return fmt.Errorf("InputDataLength is out of range")
	}
	if o.SessionDataLength > uint32(8192) {
		return fmt.Errorf("SessionDataLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySessionDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	// type {in} (1:{alias=QUERY_SESSION_DATA_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.Type)); err != nil {
			return err
		}
	}
	// pbInputData {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbInputData](uint8))
	{
		if o.InputData != nil || o.InputDataLength > 0 {
			_ptr_pbInputData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InputDataLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InputData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InputData[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InputData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InputData, _ptr_pbInputData); err != nil {
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
	// cbInputData {in} (1:{range=(0,8192), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InputDataLength); err != nil {
			return err
		}
	}
	// cbSessionData {in} (1:{range=(0,8192), alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionDataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySessionDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	// type {in} (1:{alias=QUERY_SESSION_DATA_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
			return err
		}
	}
	// pbInputData {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbInputData](uint8))
	{
		_ptr_pbInputData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InputData", sizeInfo[0])
			}
			o.InputData = make([]byte, sizeInfo[0])
			for i1 := range o.InputData {
				i1 := i1
				if err := w.ReadData(&o.InputData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbInputData := func(ptr interface{}) { o.InputData = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InputData, _s_pbInputData, _ptr_pbInputData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbInputData {in} (1:{range=(0,8192), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InputDataLength); err != nil {
			return err
		}
	}
	// cbSessionData {in} (1:{range=(0,8192), alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionDataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySessionDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SessionData != nil && o.ReturnLength == 0 {
		o.ReturnLength = uint32(len(o.SessionData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySessionDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pbSessionData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbSessionData,length_is=pcbReturnLength](uint8))
	{
		dimSize1 := uint64(o.SessionDataLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ReturnLength)
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
		for i1 := range o.SessionData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.SessionData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.SessionData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbReturnLength {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ReturnLength); err != nil {
			return err
		}
	}
	// pcbRequireBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.RequireBufferLength); err != nil {
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

func (o *xxx_QuerySessionDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pbSessionData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbSessionData,length_is=pcbReturnLength](uint8))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SessionData", sizeInfo[0])
		}
		o.SessionData = make([]byte, sizeInfo[0])
		for i1 := range o.SessionData {
			i1 := i1
			if err := w.ReadData(&o.SessionData[i1]); err != nil {
				return err
			}
		}
	}
	// pcbReturnLength {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ReturnLength); err != nil {
			return err
		}
	}
	// pcbRequireBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.RequireBufferLength); err != nil {
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

// QuerySessionDataRequest structure represents the RpcQuerySessionData operation request
type QuerySessionDataRequest struct {
	// SessionId:  The ID of the session for which data is to be retrieved.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
	// type:  The type of data to retrieve about the session. This is of type QUERY_SESSION_DATA_TYPE.
	Type QuerySessionDataType `idl:"name:type" json:"type"`
	// pbInputData:  Input data. This is a string specifying the name of the virtual channel
	// and is required only when querying virtual channel information.
	InputData []byte `idl:"name:pbInputData;size_is:(cbInputData);pointer:unique" json:"input_data"`
	// cbInputData:  The size, in bytes, of input data.
	InputDataLength uint32 `idl:"name:cbInputData" json:"input_data_length"`
	// cbSessionData:  The size, in bytes, of pbSessionData.
	SessionDataLength uint32 `idl:"name:cbSessionData" json:"session_data_length"`
}

func (o *QuerySessionDataRequest) xxx_ToOp(ctx context.Context, op *xxx_QuerySessionDataOperation) *xxx_QuerySessionDataOperation {
	if op == nil {
		op = &xxx_QuerySessionDataOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	op.Type = o.Type
	op.InputData = o.InputData
	op.InputDataLength = o.InputDataLength
	op.SessionDataLength = o.SessionDataLength
	return op
}

func (o *QuerySessionDataRequest) xxx_FromOp(ctx context.Context, op *xxx_QuerySessionDataOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
	o.Type = op.Type
	o.InputData = op.InputData
	o.InputDataLength = op.InputDataLength
	o.SessionDataLength = op.SessionDataLength
}
func (o *QuerySessionDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QuerySessionDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySessionDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeQuerySessionDataRequest build a response structure from the given request structure.
func (o *QuerySessionDataRequest) MakeResponse() *QuerySessionDataResponse {
	if o == nil {
		return &QuerySessionDataResponse{}
	}
	return &QuerySessionDataResponse{
		SessionDataLength: o.SessionDataLength,
	}
}

// OpNum returns the operation number of RpcQuerySessionData operation.
func (o *QuerySessionDataRequest) OpNum() int { return 11 }

// OpName returns the operation name of RpcQuerySessionData operation.
func (o *QuerySessionDataRequest) OpName() string { return "/RCMPublic/v1/RpcQuerySessionData" }

// QuerySessionDataResponse structure represents the RpcQuerySessionData operation response
type QuerySessionDataResponse struct {
	// XXX: cbSessionData is an implicit input depedency for output parameters
	SessionDataLength uint32 `idl:"name:cbSessionData" json:"session_data_length"`

	// pbSessionData:  The output data containing the requested information. The data returned
	// is of type WDCONFIG if the type specified is QUERY_SESSION_DATA_WDCONFIG. It is of
	// type WINSTATIONVALIDATIONINFORMATION if the type specified is QUERY_SESSION_DATA_LICENSE_VALIDATION.
	// For other types, it is protocol-specific.
	SessionData []byte `idl:"name:pbSessionData;size_is:(cbSessionData);length_is:(pcbReturnLength);pointer:ref" json:"session_data"`
	// pcbReturnLength:  The length of the returned data, in bytes.
	ReturnLength uint32 `idl:"name:pcbReturnLength;pointer:ref" json:"return_length"`
	// pcbRequireBufferSize:  The buffer size, in bytes, required by the returned data.
	RequireBufferLength uint32 `idl:"name:pcbRequireBufferSize;pointer:ref" json:"require_buffer_length"`
	// Return: The RpcQuerySessionData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QuerySessionDataResponse) xxx_ToOp(ctx context.Context, op *xxx_QuerySessionDataOperation) *xxx_QuerySessionDataOperation {
	if op == nil {
		op = &xxx_QuerySessionDataOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.SessionDataLength == uint32(0) {
		op.SessionDataLength = o.SessionDataLength
	}

	op.SessionData = o.SessionData
	op.ReturnLength = o.ReturnLength
	op.RequireBufferLength = o.RequireBufferLength
	op.Return = o.Return
	return op
}

func (o *QuerySessionDataResponse) xxx_FromOp(ctx context.Context, op *xxx_QuerySessionDataOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.SessionDataLength = op.SessionDataLength

	o.SessionData = op.SessionData
	o.ReturnLength = op.ReturnLength
	o.RequireBufferLength = op.RequireBufferLength
	o.Return = op.Return
}
func (o *QuerySessionDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QuerySessionDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySessionDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

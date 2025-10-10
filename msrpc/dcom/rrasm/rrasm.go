// The rrasm package implements the RRASM client protocol.
//
// # Introduction
//
// # Overview
package rrasm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dtyp.GoPackage
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

// MaxSSTPHashSize represents the MAX_SSTP_HASH_SIZE RPC constant
var MaxSSTPHashSize = 32

// MaxGroupLength represents the MAX_GROUP_LEN RPC constant
var MaxGroupLength = 64

// MaxAdaptorNameLength represents the MAX_ADAPTOR_NAME_LEN RPC constant
var MaxAdaptorNameLength = 48

// MaxEntryName represents the RASRPC_MaxEntryName RPC constant
var MaxEntryName = 256

// MaxPortName represents the RASRPC_MaxPortName RPC constant
var MaxPortName = 16

// MaxDeviceName represents the RASRPC_MaxDeviceName RPC constant
var MaxDeviceName = 128

// MaxPhoneNumber represents the RASRPC_MaxPhoneNumber RPC constant
var MaxPhoneNumber = 128

// MaxPath represents the RASRPC_MAX_PATH RPC constant
var MaxPath = 260

// RouterInterfaceType type represents ROUTER_INTERFACE_TYPE RPC enumeration.
type RouterInterfaceType uint16

var (
	RouterInterfaceTypeClient     RouterInterfaceType = 0
	RouterInterfaceTypeHomeRouter RouterInterfaceType = 1
	RouterInterfaceTypeFullRouter RouterInterfaceType = 2
	RouterInterfaceTypeDedicated  RouterInterfaceType = 3
	RouterInterfaceTypeInternal   RouterInterfaceType = 4
	RouterInterfaceTypeLoopback   RouterInterfaceType = 5
	RouterInterfaceTypeTunnel1    RouterInterfaceType = 6
	RouterInterfaceTypeDialout    RouterInterfaceType = 7
)

func (o RouterInterfaceType) String() string {
	switch o {
	case RouterInterfaceTypeClient:
		return "RouterInterfaceTypeClient"
	case RouterInterfaceTypeHomeRouter:
		return "RouterInterfaceTypeHomeRouter"
	case RouterInterfaceTypeFullRouter:
		return "RouterInterfaceTypeFullRouter"
	case RouterInterfaceTypeDedicated:
		return "RouterInterfaceTypeDedicated"
	case RouterInterfaceTypeInternal:
		return "RouterInterfaceTypeInternal"
	case RouterInterfaceTypeLoopback:
		return "RouterInterfaceTypeLoopback"
	case RouterInterfaceTypeTunnel1:
		return "RouterInterfaceTypeTunnel1"
	case RouterInterfaceTypeDialout:
		return "RouterInterfaceTypeDialout"
	}
	return "Invalid"
}

// RouterConnectionState type represents ROUTER_CONNECTION_STATE RPC enumeration.
type RouterConnectionState uint16

var (
	RouterConnectionStateInterfaceStateUnreachable  RouterConnectionState = 0
	RouterConnectionStateInterfaceStateDisconnected RouterConnectionState = 1
	RouterConnectionStateInterfaceStateConnecting   RouterConnectionState = 2
	RouterConnectionStateInterfaceStateConnected    RouterConnectionState = 3
)

func (o RouterConnectionState) String() string {
	switch o {
	case RouterConnectionStateInterfaceStateUnreachable:
		return "RouterConnectionStateInterfaceStateUnreachable"
	case RouterConnectionStateInterfaceStateDisconnected:
		return "RouterConnectionStateInterfaceStateDisconnected"
	case RouterConnectionStateInterfaceStateConnecting:
		return "RouterConnectionStateInterfaceStateConnecting"
	case RouterConnectionStateInterfaceStateConnected:
		return "RouterConnectionStateInterfaceStateConnected"
	}
	return "Invalid"
}

// RASQuarantineState type represents RAS_QUARANTINE_STATE RPC enumeration.
type RASQuarantineState uint16

var (
	RASQuarantineStateNormal     RASQuarantineState = 0
	RASQuarantineStateQuarantine RASQuarantineState = 1
	RASQuarantineStateProbation  RASQuarantineState = 2
	RASQuarantineStateUnknown    RASQuarantineState = 3
)

func (o RASQuarantineState) String() string {
	switch o {
	case RASQuarantineStateNormal:
		return "RASQuarantineStateNormal"
	case RASQuarantineStateQuarantine:
		return "RASQuarantineStateQuarantine"
	case RASQuarantineStateProbation:
		return "RASQuarantineStateProbation"
	case RASQuarantineStateUnknown:
		return "RASQuarantineStateUnknown"
	}
	return "Invalid"
}

// RASPortCondition type represents RAS_PORT_CONDITION RPC enumeration.
type RASPortCondition uint16

var (
	RASPortConditionNonOperational RASPortCondition = 0
	RASPortConditionDisconnected   RASPortCondition = 1
	RASPortConditionCallingBack    RASPortCondition = 2
	RASPortConditionListening      RASPortCondition = 3
	RASPortConditionAuthenticating RASPortCondition = 4
	RASPortConditionAuthenticated  RASPortCondition = 5
	RASPortConditionInitializing   RASPortCondition = 6
)

func (o RASPortCondition) String() string {
	switch o {
	case RASPortConditionNonOperational:
		return "RASPortConditionNonOperational"
	case RASPortConditionDisconnected:
		return "RASPortConditionDisconnected"
	case RASPortConditionCallingBack:
		return "RASPortConditionCallingBack"
	case RASPortConditionListening:
		return "RASPortConditionListening"
	case RASPortConditionAuthenticating:
		return "RASPortConditionAuthenticating"
	case RASPortConditionAuthenticated:
		return "RASPortConditionAuthenticated"
	case RASPortConditionInitializing:
		return "RASPortConditionInitializing"
	}
	return "Invalid"
}

// RASHardwareCondition type represents RAS_HARDWARE_CONDITION RPC enumeration.
type RASHardwareCondition uint16

var (
	RASHardwareConditionOperational RASHardwareCondition = 0
	RASHardwareConditionFailure     RASHardwareCondition = 1
)

func (o RASHardwareCondition) String() string {
	switch o {
	case RASHardwareConditionOperational:
		return "RASHardwareConditionOperational"
	case RASHardwareConditionFailure:
		return "RASHardwareConditionFailure"
	}
	return "Invalid"
}

// ForwardAction type represents FORWARD_ACTION RPC enumeration.
type ForwardAction uint16

var (
	ForwardActionForward ForwardAction = 0
	ForwardActionDrop    ForwardAction = 1
)

func (o ForwardAction) String() string {
	switch o {
	case ForwardActionForward:
		return "ForwardActionForward"
	case ForwardActionDrop:
		return "ForwardActionDrop"
	}
	return "Invalid"
}

// MIBIPForwardType type represents MIB_IPFORWARD_TYPE RPC enumeration.
type MIBIPForwardType uint16

var (
	MIBIPForwardTypeRouteTypeOther    MIBIPForwardType = 1
	MIBIPForwardTypeRouteTypeInvalid  MIBIPForwardType = 2
	MIBIPForwardTypeRouteTypeDirect   MIBIPForwardType = 3
	MIBIPForwardTypeRouteTypeIndirect MIBIPForwardType = 4
)

func (o MIBIPForwardType) String() string {
	switch o {
	case MIBIPForwardTypeRouteTypeOther:
		return "MIBIPForwardTypeRouteTypeOther"
	case MIBIPForwardTypeRouteTypeInvalid:
		return "MIBIPForwardTypeRouteTypeInvalid"
	case MIBIPForwardTypeRouteTypeDirect:
		return "MIBIPForwardTypeRouteTypeDirect"
	case MIBIPForwardTypeRouteTypeIndirect:
		return "MIBIPForwardTypeRouteTypeIndirect"
	}
	return "Invalid"
}

// MIBIPForwardProto type represents MIB_IPFORWARD_PROTO RPC enumeration.
type MIBIPForwardProto uint16

var (
	MIBIPForwardProtoOther          MIBIPForwardProto = 1
	MIBIPForwardProtoLocal          MIBIPForwardProto = 2
	MIBIPForwardProtoNETMGMT        MIBIPForwardProto = 3
	MIBIPForwardProtoICMP           MIBIPForwardProto = 4
	MIBIPForwardProtoEGP            MIBIPForwardProto = 5
	MIBIPForwardProtoGGP            MIBIPForwardProto = 6
	MIBIPForwardProtoHello          MIBIPForwardProto = 7
	MIBIPForwardProtoRIP            MIBIPForwardProto = 8
	MIBIPForwardProtoISIS           MIBIPForwardProto = 9
	MIBIPForwardProtoESIS           MIBIPForwardProto = 10
	MIBIPForwardProtoCisco          MIBIPForwardProto = 11
	MIBIPForwardProtoBbn            MIBIPForwardProto = 12
	MIBIPForwardProtoOSPF           MIBIPForwardProto = 13
	MIBIPForwardProtoBGP            MIBIPForwardProto = 14
	MIBIPForwardProtoNTAutostatic   MIBIPForwardProto = 10002
	MIBIPForwardProtoNTStatic       MIBIPForwardProto = 10006
	MIBIPForwardProtoNTStaticNonDOD MIBIPForwardProto = 10007
)

func (o MIBIPForwardProto) String() string {
	switch o {
	case MIBIPForwardProtoOther:
		return "MIBIPForwardProtoOther"
	case MIBIPForwardProtoLocal:
		return "MIBIPForwardProtoLocal"
	case MIBIPForwardProtoNETMGMT:
		return "MIBIPForwardProtoNETMGMT"
	case MIBIPForwardProtoICMP:
		return "MIBIPForwardProtoICMP"
	case MIBIPForwardProtoEGP:
		return "MIBIPForwardProtoEGP"
	case MIBIPForwardProtoGGP:
		return "MIBIPForwardProtoGGP"
	case MIBIPForwardProtoHello:
		return "MIBIPForwardProtoHello"
	case MIBIPForwardProtoRIP:
		return "MIBIPForwardProtoRIP"
	case MIBIPForwardProtoISIS:
		return "MIBIPForwardProtoISIS"
	case MIBIPForwardProtoESIS:
		return "MIBIPForwardProtoESIS"
	case MIBIPForwardProtoCisco:
		return "MIBIPForwardProtoCisco"
	case MIBIPForwardProtoBbn:
		return "MIBIPForwardProtoBbn"
	case MIBIPForwardProtoOSPF:
		return "MIBIPForwardProtoOSPF"
	case MIBIPForwardProtoBGP:
		return "MIBIPForwardProtoBGP"
	case MIBIPForwardProtoNTAutostatic:
		return "MIBIPForwardProtoNTAutostatic"
	case MIBIPForwardProtoNTStatic:
		return "MIBIPForwardProtoNTStatic"
	case MIBIPForwardProtoNTStaticNonDOD:
		return "MIBIPForwardProtoNTStaticNonDOD"
	}
	return "Invalid"
}

// MIBIPStatsForwarding type represents MIB_IPSTATS_FORWARDING RPC enumeration.
type MIBIPStatsForwarding uint16

var (
	MIBIPStatsForwardingIsForwarding    MIBIPStatsForwarding = 1
	MIBIPStatsForwardingIsNotForwarding MIBIPStatsForwarding = 2
)

func (o MIBIPStatsForwarding) String() string {
	switch o {
	case MIBIPStatsForwardingIsForwarding:
		return "MIBIPStatsForwardingIsForwarding"
	case MIBIPStatsForwardingIsNotForwarding:
		return "MIBIPStatsForwardingIsNotForwarding"
	}
	return "Invalid"
}

// MIBTCPState type represents MIB_TCP_STATE RPC enumeration.
type MIBTCPState uint16

var (
	MIBTCPStateClosed    MIBTCPState = 1
	MIBTCPStateListen    MIBTCPState = 2
	MIBTCPStateSynSent   MIBTCPState = 3
	MIBTCPStateSynRcvd   MIBTCPState = 4
	MIBTCPStateEstab     MIBTCPState = 5
	MIBTCPStateInWait1   MIBTCPState = 6
	MIBTCPStateInWait2   MIBTCPState = 7
	MIBTCPStateCloseWait MIBTCPState = 8
	MIBTCPStateClosing   MIBTCPState = 9
	MIBTCPStateLastAck   MIBTCPState = 10
	MIBTCPStateTimeWait  MIBTCPState = 11
	MIBTCPStateDeleteTcb MIBTCPState = 12
)

func (o MIBTCPState) String() string {
	switch o {
	case MIBTCPStateClosed:
		return "MIBTCPStateClosed"
	case MIBTCPStateListen:
		return "MIBTCPStateListen"
	case MIBTCPStateSynSent:
		return "MIBTCPStateSynSent"
	case MIBTCPStateSynRcvd:
		return "MIBTCPStateSynRcvd"
	case MIBTCPStateEstab:
		return "MIBTCPStateEstab"
	case MIBTCPStateInWait1:
		return "MIBTCPStateInWait1"
	case MIBTCPStateInWait2:
		return "MIBTCPStateInWait2"
	case MIBTCPStateCloseWait:
		return "MIBTCPStateCloseWait"
	case MIBTCPStateClosing:
		return "MIBTCPStateClosing"
	case MIBTCPStateLastAck:
		return "MIBTCPStateLastAck"
	case MIBTCPStateTimeWait:
		return "MIBTCPStateTimeWait"
	case MIBTCPStateDeleteTcb:
		return "MIBTCPStateDeleteTcb"
	}
	return "Invalid"
}

// TCPRTOAlgorithm type represents TCP_RTO_ALGORITHM RPC enumeration.
type TCPRTOAlgorithm uint16

var (
	TCPRTOAlgorithmMIBTCPRTOOther    TCPRTOAlgorithm = 1
	TCPRTOAlgorithmMIBTCPRTOConstant TCPRTOAlgorithm = 2
	TCPRTOAlgorithmMIBTCPRTORSRE     TCPRTOAlgorithm = 3
	TCPRTOAlgorithmMIBTCPRTOVANJ     TCPRTOAlgorithm = 4
)

func (o TCPRTOAlgorithm) String() string {
	switch o {
	case TCPRTOAlgorithmMIBTCPRTOOther:
		return "TCPRTOAlgorithmMIBTCPRTOOther"
	case TCPRTOAlgorithmMIBTCPRTOConstant:
		return "TCPRTOAlgorithmMIBTCPRTOConstant"
	case TCPRTOAlgorithmMIBTCPRTORSRE:
		return "TCPRTOAlgorithmMIBTCPRTORSRE"
	case TCPRTOAlgorithmMIBTCPRTOVANJ:
		return "TCPRTOAlgorithmMIBTCPRTOVANJ"
	}
	return "Invalid"
}

// IN6Addr structure represents IN6_ADDR RPC structure.
type IN6Addr struct {
	Union *IN6Addr_Union `idl:"name:u" json:"union"`
}

func (o *IN6Addr) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IN6Addr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	// FIXME unknown type u
	return nil
}
func (o *IN6Addr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	// FIXME: unknown type u
	return nil
}

type IN6Addr_Union struct {
	Byte []byte   `idl:"name:Byte" json:"byte"`
	Word []uint16 `idl:"name:Word" json:"word"`
}

// InformationContainer structure represents DIM_INFORMATION_CONTAINER RPC structure.
type InformationContainer struct {
	BufferSize uint32 `idl:"name:dwBufferSize" json:"buffer_size"`
	Buffer     []byte `idl:"name:pBuffer;size_is:(dwBufferSize)" json:"buffer"`
}

func (o *InformationContainer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InformationContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferSize); err != nil {
		return err
	}
	if o.Buffer != nil || o.BufferSize > 0 {
		_ptr_pBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BufferSize)
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
		if err := w.WritePointer(&o.Buffer, _ptr_pBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InformationContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferSize); err != nil {
		return err
	}
	_ptr_pBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BufferSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BufferSize)
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
	_s_pBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_pBuffer, _ptr_pBuffer); err != nil {
		return err
	}
	return nil
}

// ObjectHeaderIDL structure represents MPRAPI_OBJECT_HEADER_IDL RPC structure.
type ObjectHeaderIDL struct {
	Revision uint8  `idl:"name:revision" json:"revision"`
	Type     uint8  `idl:"name:type" json:"type"`
	Size     uint16 `idl:"name:size" json:"size"`
}

func (o *ObjectHeaderIDL) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectHeaderIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	return nil
}
func (o *ObjectHeaderIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	return nil
}

// PPPProjectionInfo1 structure represents PPP_PROJECTION_INFO_1 RPC structure.
type PPPProjectionInfo1 struct {
	IPv4NegotiationError         uint32   `idl:"name:dwIPv4NegotiationError" json:"ipv4_negotiation_error"`
	Address                      []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress                []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
	IPv4Options                  uint32   `idl:"name:dwIPv4Options" json:"ipv4_options"`
	IPv4RemoteOptions            uint32   `idl:"name:dwIPv4RemoteOptions" json:"ipv4_remote_options"`
	IPv4SubInterfaceIndex        uint64   `idl:"name:IPv4SubInterfaceIndex" json:"ipv4_sub_interface_index"`
	IPv6NegotiationError         uint32   `idl:"name:dwIPv6NegotiationError" json:"ipv6_negotiation_error"`
	InterfaceID                  []byte   `idl:"name:bInterfaceIdentifier" json:"interface_id"`
	RemoteInterfaceID            []byte   `idl:"name:bRemoteInterfaceIdentifier" json:"remote_interface_id"`
	Prefix                       []byte   `idl:"name:bPrefix" json:"prefix"`
	PrefixLength                 uint32   `idl:"name:dwPrefixLength" json:"prefix_length"`
	IPv6SubInterfaceIndex        uint64   `idl:"name:IPv6SubInterfaceIndex" json:"ipv6_sub_interface_index"`
	LCPError                     uint32   `idl:"name:dwLcpError" json:"lcp_error"`
	AuthenticationProtocol       uint32   `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	AuthenticationData           uint32   `idl:"name:dwAuthenticationData" json:"authentication_data"`
	RemoteAuthenticationProtocol uint32   `idl:"name:dwRemoteAuthenticationProtocol" json:"remote_authentication_protocol"`
	RemoteAuthenticationData     uint32   `idl:"name:dwRemoteAuthenticationData" json:"remote_authentication_data"`
	LCPTerminateReason           uint32   `idl:"name:dwLcpTerminateReason" json:"lcp_terminate_reason"`
	LCPRemoteTerminateReason     uint32   `idl:"name:dwLcpRemoteTerminateReason" json:"lcp_remote_terminate_reason"`
	LCPOptions                   uint32   `idl:"name:dwLcpOptions" json:"lcp_options"`
	LCPRemoteOptions             uint32   `idl:"name:dwLcpRemoteOptions" json:"lcp_remote_options"`
	EAPTypeID                    uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
	RemoteEAPTypeID              uint32   `idl:"name:dwRemoteEapTypeId" json:"remote_eap_type_id"`
	CCPError                     uint32   `idl:"name:dwCcpError" json:"ccp_error"`
	CompressionAlgorithm         uint32   `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	CCPOptions                   uint32   `idl:"name:dwCcpOptions" json:"ccp_options"`
	RemoteCompressionAlgorithm   uint32   `idl:"name:dwRemoteCompressionAlgorithm" json:"remote_compression_algorithm"`
	CCPRemoteOptions             uint32   `idl:"name:dwCcpRemoteOptions" json:"ccp_remote_options"`
}

func (o *PPPProjectionInfo1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPProjectionInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4NegotiationError); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteAddress {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteAddress); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPv4Options); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4RemoteOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6NegotiationError); err != nil {
		return err
	}
	for i1 := range o.InterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteInterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Prefix {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Prefix[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Prefix); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrefixLength); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPError); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationData); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAuthenticationData); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.EAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteEAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.CCPError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CCPOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CCPRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *PPPProjectionInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4NegotiationError); err != nil {
		return err
	}
	o.Address = make([]uint16, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	o.RemoteAddress = make([]uint16, 16)
	for i1 := range o.RemoteAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.IPv4Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4RemoteOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6NegotiationError); err != nil {
		return err
	}
	o.InterfaceID = make([]byte, 8)
	for i1 := range o.InterfaceID {
		i1 := i1
		if err := w.ReadData(&o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	o.RemoteInterfaceID = make([]byte, 8)
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if err := w.ReadData(&o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	o.Prefix = make([]byte, 8)
	for i1 := range o.Prefix {
		i1 := i1
		if err := w.ReadData(&o.Prefix[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PrefixLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPError); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationData); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAuthenticationData); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPRemoteOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.EAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteEAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CCPError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CCPOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CCPRemoteOptions); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// PPPProjectionInfo2 structure represents PPP_PROJECTION_INFO_2 RPC structure.
type PPPProjectionInfo2 struct {
	IPv4NegotiationError         uint32   `idl:"name:dwIPv4NegotiationError" json:"ipv4_negotiation_error"`
	Address                      []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress                []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
	IPv4Options                  uint32   `idl:"name:dwIPv4Options" json:"ipv4_options"`
	IPv4RemoteOptions            uint32   `idl:"name:dwIPv4RemoteOptions" json:"ipv4_remote_options"`
	IPv4SubInterfaceIndex        uint64   `idl:"name:IPv4SubInterfaceIndex" json:"ipv4_sub_interface_index"`
	IPv6NegotiationError         uint32   `idl:"name:dwIPv6NegotiationError" json:"ipv6_negotiation_error"`
	InterfaceID                  []byte   `idl:"name:bInterfaceIdentifier" json:"interface_id"`
	RemoteInterfaceID            []byte   `idl:"name:bRemoteInterfaceIdentifier" json:"remote_interface_id"`
	Prefix                       []byte   `idl:"name:bPrefix" json:"prefix"`
	PrefixLength                 uint32   `idl:"name:dwPrefixLength" json:"prefix_length"`
	IPv6SubInterfaceIndex        uint64   `idl:"name:IPv6SubInterfaceIndex" json:"ipv6_sub_interface_index"`
	LCPError                     uint32   `idl:"name:dwLcpError" json:"lcp_error"`
	AuthenticationProtocol       uint32   `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	AuthenticationData           uint32   `idl:"name:dwAuthenticationData" json:"authentication_data"`
	RemoteAuthenticationProtocol uint32   `idl:"name:dwRemoteAuthenticationProtocol" json:"remote_authentication_protocol"`
	RemoteAuthenticationData     uint32   `idl:"name:dwRemoteAuthenticationData" json:"remote_authentication_data"`
	LCPTerminateReason           uint32   `idl:"name:dwLcpTerminateReason" json:"lcp_terminate_reason"`
	LCPRemoteTerminateReason     uint32   `idl:"name:dwLcpRemoteTerminateReason" json:"lcp_remote_terminate_reason"`
	LCPOptions                   uint32   `idl:"name:dwLcpOptions" json:"lcp_options"`
	LCPRemoteOptions             uint32   `idl:"name:dwLcpRemoteOptions" json:"lcp_remote_options"`
	EAPTypeID                    uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
	EmbeddedEAPTypeID            uint32   `idl:"name:dwEmbeddedEAPTypeId" json:"embedded_eap_type_id"`
	RemoteEAPTypeID              uint32   `idl:"name:dwRemoteEapTypeId" json:"remote_eap_type_id"`
	CCPError                     uint32   `idl:"name:dwCcpError" json:"ccp_error"`
	CompressionAlgorithm         uint32   `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	CCPOptions                   uint32   `idl:"name:dwCcpOptions" json:"ccp_options"`
	RemoteCompressionAlgorithm   uint32   `idl:"name:dwRemoteCompressionAlgorithm" json:"remote_compression_algorithm"`
	CCPRemoteOptions             uint32   `idl:"name:dwCcpRemoteOptions" json:"ccp_remote_options"`
}

func (o *PPPProjectionInfo2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPProjectionInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4NegotiationError); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteAddress {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteAddress); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPv4Options); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4RemoteOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6NegotiationError); err != nil {
		return err
	}
	for i1 := range o.InterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteInterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Prefix {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Prefix[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Prefix); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrefixLength); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPError); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationData); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAuthenticationData); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.LCPRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.EAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.EmbeddedEAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteEAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.CCPError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CCPOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CCPRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *PPPProjectionInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4NegotiationError); err != nil {
		return err
	}
	o.Address = make([]uint16, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	o.RemoteAddress = make([]uint16, 16)
	for i1 := range o.RemoteAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.IPv4Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4RemoteOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6NegotiationError); err != nil {
		return err
	}
	o.InterfaceID = make([]byte, 8)
	for i1 := range o.InterfaceID {
		i1 := i1
		if err := w.ReadData(&o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	o.RemoteInterfaceID = make([]byte, 8)
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if err := w.ReadData(&o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	o.Prefix = make([]byte, 8)
	for i1 := range o.Prefix {
		i1 := i1
		if err := w.ReadData(&o.Prefix[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PrefixLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPError); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationData); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAuthenticationData); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCPRemoteOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.EAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.EmbeddedEAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteEAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CCPError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CCPOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CCPRemoteOptions); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// IKEv2ProjectionInfo1 structure represents IKEV2_PROJECTION_INFO_1 RPC structure.
type IKEv2ProjectionInfo1 struct {
	IPv4NegotiationError   uint32   `idl:"name:dwIPv4NegotiationError" json:"ipv4_negotiation_error"`
	Address                []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress          []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
	IPv4SubInterfaceIndex  uint64   `idl:"name:IPv4SubInterfaceIndex" json:"ipv4_sub_interface_index"`
	IPv6NegotiationError   uint32   `idl:"name:dwIPv6NegotiationError" json:"ipv6_negotiation_error"`
	InterfaceID            []byte   `idl:"name:bInterfaceIdentifier" json:"interface_id"`
	RemoteInterfaceID      []byte   `idl:"name:bRemoteInterfaceIdentifier" json:"remote_interface_id"`
	Prefix                 []byte   `idl:"name:bPrefix" json:"prefix"`
	PrefixLength           uint32   `idl:"name:dwPrefixLength" json:"prefix_length"`
	IPv6SubInterfaceIndex  uint64   `idl:"name:IPv6SubInterfaceIndex" json:"ipv6_sub_interface_index"`
	Options                uint32   `idl:"name:dwOptions" json:"options"`
	AuthenticationProtocol uint32   `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	EAPTypeID              uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
	CompressionAlgorithm   uint32   `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	EncryptionMethod       uint32   `idl:"name:dwEncryptionMethod" json:"encryption_method"`
}

func (o *IKEv2ProjectionInfo1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2ProjectionInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4NegotiationError); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteAddress {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteAddress); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6NegotiationError); err != nil {
		return err
	}
	for i1 := range o.InterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteInterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Prefix {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Prefix[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Prefix); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrefixLength); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.EAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2ProjectionInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4NegotiationError); err != nil {
		return err
	}
	o.Address = make([]uint16, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	o.RemoteAddress = make([]uint16, 16)
	for i1 := range o.RemoteAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6NegotiationError); err != nil {
		return err
	}
	o.InterfaceID = make([]byte, 8)
	for i1 := range o.InterfaceID {
		i1 := i1
		if err := w.ReadData(&o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	o.RemoteInterfaceID = make([]byte, 8)
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if err := w.ReadData(&o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	o.Prefix = make([]byte, 8)
	for i1 := range o.Prefix {
		i1 := i1
		if err := w.ReadData(&o.Prefix[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PrefixLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.EAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// IKEv2ProjectionInfo2 structure represents IKEV2_PROJECTION_INFO_2 RPC structure.
type IKEv2ProjectionInfo2 struct {
	IPv4NegotiationError   uint32   `idl:"name:dwIPv4NegotiationError" json:"ipv4_negotiation_error"`
	Address                []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress          []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
	IPv4SubInterfaceIndex  uint64   `idl:"name:IPv4SubInterfaceIndex" json:"ipv4_sub_interface_index"`
	IPv6NegotiationError   uint32   `idl:"name:dwIPv6NegotiationError" json:"ipv6_negotiation_error"`
	InterfaceID            []byte   `idl:"name:bInterfaceIdentifier" json:"interface_id"`
	RemoteInterfaceID      []byte   `idl:"name:bRemoteInterfaceIdentifier" json:"remote_interface_id"`
	Prefix                 []byte   `idl:"name:bPrefix" json:"prefix"`
	PrefixLength           uint32   `idl:"name:dwPrefixLength" json:"prefix_length"`
	IPv6SubInterfaceIndex  uint64   `idl:"name:IPv6SubInterfaceIndex" json:"ipv6_sub_interface_index"`
	Options                uint32   `idl:"name:dwOptions" json:"options"`
	AuthenticationProtocol uint32   `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	EAPTypeID              uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
	EmbeddedEAPTypeID      uint32   `idl:"name:dwEmbeddedEAPTypeId" json:"embedded_eap_type_id"`
	CompressionAlgorithm   uint32   `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	EncryptionMethod       uint32   `idl:"name:dwEncryptionMethod" json:"encryption_method"`
}

func (o *IKEv2ProjectionInfo2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2ProjectionInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv4NegotiationError); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteAddress {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteAddress); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6NegotiationError); err != nil {
		return err
	}
	for i1 := range o.InterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteInterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Prefix {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Prefix[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Prefix); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrefixLength); err != nil {
		return err
	}
	if err := w.WriteData(o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.EAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.EmbeddedEAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2ProjectionInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv4NegotiationError); err != nil {
		return err
	}
	o.Address = make([]uint16, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	o.RemoteAddress = make([]uint16, 16)
	for i1 := range o.RemoteAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.IPv4SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6NegotiationError); err != nil {
		return err
	}
	o.InterfaceID = make([]byte, 8)
	for i1 := range o.InterfaceID {
		i1 := i1
		if err := w.ReadData(&o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	o.RemoteInterfaceID = make([]byte, 8)
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if err := w.ReadData(&o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	o.Prefix = make([]byte, 8)
	for i1 := range o.Prefix {
		i1 := i1
		if err := w.ReadData(&o.Prefix[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PrefixLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6SubInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.EAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.EmbeddedEAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ProjectionInfoIDL1 structure represents PROJECTION_INFO_IDL_1 RPC union.
type ProjectionInfoIDL1 struct {
	ProjectionInfoType uint8
	// Types that are assignable to Value
	//
	// *ProjectionInfoIDL1_PPPProjectionInfo
	// *ProjectionInfoIDL1_IKEv2ProjectionInfo
	Value is_ProjectionInfoIDL1 `json:"value"`
}

func (o *ProjectionInfoIDL1) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ProjectionInfoIDL1_PPPProjectionInfo:
		if value != nil {
			return value.PPPProjectionInfo
		}
	case *ProjectionInfoIDL1_IKEv2ProjectionInfo:
		if value != nil {
			return value.IKEv2ProjectionInfo
		}
	}
	return nil
}

type is_ProjectionInfoIDL1 interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ProjectionInfoIDL1()
}

func (o *ProjectionInfoIDL1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint8(o.ProjectionInfoType)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch o.ProjectionInfoType {
	case uint8(1):
		_o, _ := o.Value.(*ProjectionInfoIDL1_PPPProjectionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProjectionInfoIDL1_PPPProjectionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*ProjectionInfoIDL1_IKEv2ProjectionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProjectionInfoIDL1_IKEv2ProjectionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.ProjectionInfoType)
	}
	return nil
}

func (o *ProjectionInfoIDL1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint8)(&o.ProjectionInfoType)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch o.ProjectionInfoType {
	case uint8(1):
		o.Value = &ProjectionInfoIDL1_PPPProjectionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &ProjectionInfoIDL1_IKEv2ProjectionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.ProjectionInfoType)
	}
	return nil
}

// ProjectionInfoIDL1_PPPProjectionInfo structure represents PROJECTION_INFO_IDL_1 RPC union arm.
//
// It has following labels: 1
type ProjectionInfoIDL1_PPPProjectionInfo struct {
	PPPProjectionInfo *PPPProjectionInfo1 `idl:"name:PppProjectionInfo" json:"ppp_projection_info"`
}

func (*ProjectionInfoIDL1_PPPProjectionInfo) is_ProjectionInfoIDL1() {}

func (o *ProjectionInfoIDL1_PPPProjectionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PPPProjectionInfo != nil {
		if err := o.PPPProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPProjectionInfo1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProjectionInfoIDL1_PPPProjectionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PPPProjectionInfo == nil {
		o.PPPProjectionInfo = &PPPProjectionInfo1{}
	}
	if err := o.PPPProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ProjectionInfoIDL1_IKEv2ProjectionInfo structure represents PROJECTION_INFO_IDL_1 RPC union arm.
//
// It has following labels: 2
type ProjectionInfoIDL1_IKEv2ProjectionInfo struct {
	IKEv2ProjectionInfo *IKEv2ProjectionInfo1 `idl:"name:Ikev2ProjectionInfo" json:"ikev2_projection_info"`
}

func (*ProjectionInfoIDL1_IKEv2ProjectionInfo) is_ProjectionInfoIDL1() {}

func (o *ProjectionInfoIDL1_IKEv2ProjectionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IKEv2ProjectionInfo != nil {
		if err := o.IKEv2ProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2ProjectionInfo1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProjectionInfoIDL1_IKEv2ProjectionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IKEv2ProjectionInfo == nil {
		o.IKEv2ProjectionInfo = &IKEv2ProjectionInfo1{}
	}
	if err := o.IKEv2ProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ProjectionInfoIDL2 structure represents PROJECTION_INFO_IDL_2 RPC union.
type ProjectionInfoIDL2 struct {
	ProjectionInfoType uint8
	// Types that are assignable to Value
	//
	// *ProjectionInfoIDL2_PPPProjectionInfo
	// *ProjectionInfoIDL2_IKEv2ProjectionInfo
	Value is_ProjectionInfoIDL2 `json:"value"`
}

func (o *ProjectionInfoIDL2) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ProjectionInfoIDL2_PPPProjectionInfo:
		if value != nil {
			return value.PPPProjectionInfo
		}
	case *ProjectionInfoIDL2_IKEv2ProjectionInfo:
		if value != nil {
			return value.IKEv2ProjectionInfo
		}
	}
	return nil
}

type is_ProjectionInfoIDL2 interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ProjectionInfoIDL2()
}

func (o *ProjectionInfoIDL2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint8(o.ProjectionInfoType)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch o.ProjectionInfoType {
	case uint8(1):
		_o, _ := o.Value.(*ProjectionInfoIDL2_PPPProjectionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProjectionInfoIDL2_PPPProjectionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*ProjectionInfoIDL2_IKEv2ProjectionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProjectionInfoIDL2_IKEv2ProjectionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.ProjectionInfoType)
	}
	return nil
}

func (o *ProjectionInfoIDL2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint8)(&o.ProjectionInfoType)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch o.ProjectionInfoType {
	case uint8(1):
		o.Value = &ProjectionInfoIDL2_PPPProjectionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &ProjectionInfoIDL2_IKEv2ProjectionInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.ProjectionInfoType)
	}
	return nil
}

// ProjectionInfoIDL2_PPPProjectionInfo structure represents PROJECTION_INFO_IDL_2 RPC union arm.
//
// It has following labels: 1
type ProjectionInfoIDL2_PPPProjectionInfo struct {
	PPPProjectionInfo *PPPProjectionInfo2 `idl:"name:PppProjectionInfo" json:"ppp_projection_info"`
}

func (*ProjectionInfoIDL2_PPPProjectionInfo) is_ProjectionInfoIDL2() {}

func (o *ProjectionInfoIDL2_PPPProjectionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PPPProjectionInfo != nil {
		if err := o.PPPProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPProjectionInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProjectionInfoIDL2_PPPProjectionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PPPProjectionInfo == nil {
		o.PPPProjectionInfo = &PPPProjectionInfo2{}
	}
	if err := o.PPPProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ProjectionInfoIDL2_IKEv2ProjectionInfo structure represents PROJECTION_INFO_IDL_2 RPC union arm.
//
// It has following labels: 2
type ProjectionInfoIDL2_IKEv2ProjectionInfo struct {
	IKEv2ProjectionInfo *IKEv2ProjectionInfo2 `idl:"name:Ikev2ProjectionInfo" json:"ikev2_projection_info"`
}

func (*ProjectionInfoIDL2_IKEv2ProjectionInfo) is_ProjectionInfoIDL2() {}

func (o *ProjectionInfoIDL2_IKEv2ProjectionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IKEv2ProjectionInfo != nil {
		if err := o.IKEv2ProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2ProjectionInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProjectionInfoIDL2_IKEv2ProjectionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IKEv2ProjectionInfo == nil {
		o.IKEv2ProjectionInfo = &IKEv2ProjectionInfo2{}
	}
	if err := o.IKEv2ProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASConnectionEx1IDL structure represents RAS_CONNECTION_EX_1_IDL RPC structure.
type RASConnectionEx1IDL struct {
	Header                *ObjectHeaderIDL    `idl:"name:Header" json:"header"`
	ConnectDuration       uint32              `idl:"name:dwConnectDuration" json:"connect_duration"`
	InterfaceType         RouterInterfaceType `idl:"name:dwInterfaceType" json:"interface_type"`
	ConnectionFlags       uint32              `idl:"name:dwConnectionFlags" json:"connection_flags"`
	InterfaceName         []uint16            `idl:"name:wszInterfaceName" json:"interface_name"`
	UserName              []uint16            `idl:"name:wszUserName" json:"user_name"`
	LogonDomain           []uint16            `idl:"name:wszLogonDomain" json:"logon_domain"`
	RemoteComputer        []uint16            `idl:"name:wszRemoteComputer" json:"remote_computer"`
	GUID                  *dtyp.GUID          `idl:"name:guid" json:"guid"`
	RASQuarantineState    RASQuarantineState  `idl:"name:rasQuarState" json:"ras_quarantine_state"`
	ProbationTime         *dtyp.Filetime      `idl:"name:probationTime" json:"probation_time"`
	BytesXmited           uint32              `idl:"name:dwBytesXmited" json:"bytes_xmited"`
	BytesRcved            uint32              `idl:"name:dwBytesRcved" json:"bytes_rcved"`
	FramesXmited          uint32              `idl:"name:dwFramesXmited" json:"frames_xmited"`
	FramesRcved           uint32              `idl:"name:dwFramesRcved" json:"frames_rcved"`
	CRCError              uint32              `idl:"name:dwCrcErr" json:"crc_error"`
	TimeoutError          uint32              `idl:"name:dwTimeoutErr" json:"timeout_error"`
	AlignmentError        uint32              `idl:"name:dwAlignmentErr" json:"alignment_error"`
	HardwareOverrunError  uint32              `idl:"name:dwHardwareOverrunErr" json:"hardware_overrun_error"`
	FramingError          uint32              `idl:"name:dwFramingErr" json:"framing_error"`
	BufferOverrunError    uint32              `idl:"name:dwBufferOverrunErr" json:"buffer_overrun_error"`
	CompressionRatioIn    uint32              `idl:"name:dwCompressionRatioIn" json:"compression_ratio_in"`
	CompressionRatioOut   uint32              `idl:"name:dwCompressionRatioOut" json:"compression_ratio_out"`
	SwitchOversLength     uint32              `idl:"name:dwNumSwitchOvers" json:"switch_overs_length"`
	RemoteEndpointAddress []uint16            `idl:"name:wszRemoteEndpointAddress" json:"remote_endpoint_address"`
	LocalEndpointAddress  []uint16            `idl:"name:wszLocalEndpointAddress" json:"local_endpoint_address"`
	ProjectionInfo        *ProjectionInfoIDL1 `idl:"name:ProjectionInfo" json:"projection_info"`
	Connection            uint32              `idl:"name:hConnection" json:"connection"`
	Interface             uint32              `idl:"name:hInterface" json:"interface"`
}

func (o *RASConnectionEx1IDL) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASConnectionEx1IDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ConnectDuration); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectionFlags); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.UserName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.UserName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UserName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.LogonDomain {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LogonDomain[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LogonDomain); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteComputer {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.RemoteComputer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteComputer); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.RASQuarantineState)); err != nil {
		return err
	}
	if o.ProbationTime != nil {
		if err := o.ProbationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BytesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.CRCError); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutError); err != nil {
		return err
	}
	if err := w.WriteData(o.AlignmentError); err != nil {
		return err
	}
	if err := w.WriteData(o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.FramingError); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioOut); err != nil {
		return err
	}
	if err := w.WriteData(o.SwitchOversLength); err != nil {
		return err
	}
	for i1 := range o.RemoteEndpointAddress {
		i1 := i1
		if uint64(i1) >= 65 {
			break
		}
		if err := w.WriteData(o.RemoteEndpointAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteEndpointAddress); uint64(i1) < 65; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.LocalEndpointAddress {
		i1 := i1
		if uint64(i1) >= 65 {
			break
		}
		if err := w.WriteData(o.LocalEndpointAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalEndpointAddress); uint64(i1) < 65; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.ProjectionInfo != nil {
		if err := o.ProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ProjectionInfoIDL1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *RASConnectionEx1IDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectDuration); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectionFlags); err != nil {
		return err
	}
	o.InterfaceName = make([]uint16, 257)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	o.UserName = make([]uint16, 257)
	for i1 := range o.UserName {
		i1 := i1
		if err := w.ReadData(&o.UserName[i1]); err != nil {
			return err
		}
	}
	o.LogonDomain = make([]uint16, 16)
	for i1 := range o.LogonDomain {
		i1 := i1
		if err := w.ReadData(&o.LogonDomain[i1]); err != nil {
			return err
		}
	}
	o.RemoteComputer = make([]uint16, 17)
	for i1 := range o.RemoteComputer {
		i1 := i1
		if err := w.ReadData(&o.RemoteComputer[i1]); err != nil {
			return err
		}
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RASQuarantineState)); err != nil {
		return err
	}
	if o.ProbationTime == nil {
		o.ProbationTime = &dtyp.Filetime{}
	}
	if err := o.ProbationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.CRCError); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutError); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlignmentError); err != nil {
		return err
	}
	if err := w.ReadData(&o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramingError); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioOut); err != nil {
		return err
	}
	if err := w.ReadData(&o.SwitchOversLength); err != nil {
		return err
	}
	o.RemoteEndpointAddress = make([]uint16, 65)
	for i1 := range o.RemoteEndpointAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteEndpointAddress[i1]); err != nil {
			return err
		}
	}
	o.LocalEndpointAddress = make([]uint16, 65)
	for i1 := range o.LocalEndpointAddress {
		i1 := i1
		if err := w.ReadData(&o.LocalEndpointAddress[i1]); err != nil {
			return err
		}
	}
	if o.ProjectionInfo == nil {
		o.ProjectionInfo = &ProjectionInfoIDL1{}
	}
	if err := o.ProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// RASConnectionExIDL structure represents RAS_CONNECTION_EX_IDL RPC union.
type RASConnectionExIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *RASConnectionExIDL_RASConnection1
	Value is_RASConnectionExIDL `json:"value"`
}

func (o *RASConnectionExIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RASConnectionExIDL_RASConnection1:
		if value != nil {
			return value.RASConnection1
		}
	}
	return nil
}

type is_RASConnectionExIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RASConnectionExIDL()
}

func (o *RASConnectionExIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint8(o.Revision)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		_o, _ := o.Value.(*RASConnectionExIDL_RASConnection1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RASConnectionExIDL_RASConnection1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *RASConnectionExIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint8)(&o.Revision)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		o.Value = &RASConnectionExIDL_RASConnection1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// RASConnectionExIDL_RASConnection1 structure represents RAS_CONNECTION_EX_IDL RPC union arm.
//
// It has following labels: 1
type RASConnectionExIDL_RASConnection1 struct {
	RASConnection1 *RASConnectionEx1IDL `idl:"name:RasConnection1" json:"ras_connection1"`
}

func (*RASConnectionExIDL_RASConnection1) is_RASConnectionExIDL() {}

func (o *RASConnectionExIDL_RASConnection1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RASConnection1 != nil {
		if err := o.RASConnection1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASConnectionEx1IDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RASConnectionExIDL_RASConnection1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.RASConnection1 == nil {
		o.RASConnection1 = &RASConnectionEx1IDL{}
	}
	if err := o.RASConnection1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASConnection4IDL structure represents RAS_CONNECTION_4_IDL RPC structure.
type RASConnection4IDL struct {
	ConnectDuration       uint32              `idl:"name:dwConnectDuration" json:"connect_duration"`
	InterfaceType         RouterInterfaceType `idl:"name:dwInterfaceType" json:"interface_type"`
	ConnectionFlags       uint32              `idl:"name:dwConnectionFlags" json:"connection_flags"`
	InterfaceName         []uint16            `idl:"name:wszInterfaceName" json:"interface_name"`
	UserName              []uint16            `idl:"name:wszUserName" json:"user_name"`
	LogonDomain           []uint16            `idl:"name:wszLogonDomain" json:"logon_domain"`
	RemoteComputer        []uint16            `idl:"name:wszRemoteComputer" json:"remote_computer"`
	GUID                  *dtyp.GUID          `idl:"name:guid" json:"guid"`
	RASQuarantineState    RASQuarantineState  `idl:"name:rasQuarState" json:"ras_quarantine_state"`
	ProbationTime         *dtyp.Filetime      `idl:"name:probationTime" json:"probation_time"`
	ConnectionStartTime   *dtyp.Filetime      `idl:"name:connectionStartTime" json:"connection_start_time"`
	BytesXmited           uint32              `idl:"name:dwBytesXmited" json:"bytes_xmited"`
	BytesRcved            uint32              `idl:"name:dwBytesRcved" json:"bytes_rcved"`
	FramesXmited          uint32              `idl:"name:dwFramesXmited" json:"frames_xmited"`
	FramesRcved           uint32              `idl:"name:dwFramesRcved" json:"frames_rcved"`
	CRCError              uint32              `idl:"name:dwCrcErr" json:"crc_error"`
	TimeoutError          uint32              `idl:"name:dwTimeoutErr" json:"timeout_error"`
	AlignmentError        uint32              `idl:"name:dwAlignmentErr" json:"alignment_error"`
	HardwareOverrunError  uint32              `idl:"name:dwHardwareOverrunErr" json:"hardware_overrun_error"`
	FramingError          uint32              `idl:"name:dwFramingErr" json:"framing_error"`
	BufferOverrunError    uint32              `idl:"name:dwBufferOverrunErr" json:"buffer_overrun_error"`
	CompressionRatioIn    uint32              `idl:"name:dwCompressionRatioIn" json:"compression_ratio_in"`
	CompressionRatioOut   uint32              `idl:"name:dwCompressionRatioOut" json:"compression_ratio_out"`
	SwitchOversLength     uint32              `idl:"name:dwNumSwitchOvers" json:"switch_overs_length"`
	RemoteEndpointAddress []uint16            `idl:"name:wszRemoteEndpointAddress" json:"remote_endpoint_address"`
	LocalEndpointAddress  []uint16            `idl:"name:wszLocalEndpointAddress" json:"local_endpoint_address"`
	ProjectionInfo        *ProjectionInfoIDL2 `idl:"name:ProjectionInfo" json:"projection_info"`
	Connection            uint32              `idl:"name:hConnection" json:"connection"`
	Interface             uint32              `idl:"name:hInterface" json:"interface"`
	DeviceType            uint32              `idl:"name:dwDeviceType" json:"device_type"`
}

func (o *RASConnection4IDL) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASConnection4IDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectDuration); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectionFlags); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.UserName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.UserName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UserName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.LogonDomain {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LogonDomain[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LogonDomain); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteComputer {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.RemoteComputer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteComputer); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.RASQuarantineState)); err != nil {
		return err
	}
	if o.ProbationTime != nil {
		if err := o.ProbationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ConnectionStartTime != nil {
		if err := o.ConnectionStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BytesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.CRCError); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutError); err != nil {
		return err
	}
	if err := w.WriteData(o.AlignmentError); err != nil {
		return err
	}
	if err := w.WriteData(o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.FramingError); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioOut); err != nil {
		return err
	}
	if err := w.WriteData(o.SwitchOversLength); err != nil {
		return err
	}
	for i1 := range o.RemoteEndpointAddress {
		i1 := i1
		if uint64(i1) >= 65 {
			break
		}
		if err := w.WriteData(o.RemoteEndpointAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteEndpointAddress); uint64(i1) < 65; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.LocalEndpointAddress {
		i1 := i1
		if uint64(i1) >= 65 {
			break
		}
		if err := w.WriteData(o.LocalEndpointAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalEndpointAddress); uint64(i1) < 65; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.ProjectionInfo != nil {
		if err := o.ProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ProjectionInfoIDL2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *RASConnection4IDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectDuration); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectionFlags); err != nil {
		return err
	}
	o.InterfaceName = make([]uint16, 257)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	o.UserName = make([]uint16, 257)
	for i1 := range o.UserName {
		i1 := i1
		if err := w.ReadData(&o.UserName[i1]); err != nil {
			return err
		}
	}
	o.LogonDomain = make([]uint16, 16)
	for i1 := range o.LogonDomain {
		i1 := i1
		if err := w.ReadData(&o.LogonDomain[i1]); err != nil {
			return err
		}
	}
	o.RemoteComputer = make([]uint16, 17)
	for i1 := range o.RemoteComputer {
		i1 := i1
		if err := w.ReadData(&o.RemoteComputer[i1]); err != nil {
			return err
		}
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RASQuarantineState)); err != nil {
		return err
	}
	if o.ProbationTime == nil {
		o.ProbationTime = &dtyp.Filetime{}
	}
	if err := o.ProbationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ConnectionStartTime == nil {
		o.ConnectionStartTime = &dtyp.Filetime{}
	}
	if err := o.ConnectionStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.CRCError); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutError); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlignmentError); err != nil {
		return err
	}
	if err := w.ReadData(&o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramingError); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioOut); err != nil {
		return err
	}
	if err := w.ReadData(&o.SwitchOversLength); err != nil {
		return err
	}
	o.RemoteEndpointAddress = make([]uint16, 65)
	for i1 := range o.RemoteEndpointAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteEndpointAddress[i1]); err != nil {
			return err
		}
	}
	o.LocalEndpointAddress = make([]uint16, 65)
	for i1 := range o.LocalEndpointAddress {
		i1 := i1
		if err := w.ReadData(&o.LocalEndpointAddress[i1]); err != nil {
			return err
		}
	}
	if o.ProjectionInfo == nil {
		o.ProjectionInfo = &ProjectionInfoIDL2{}
	}
	if err := o.ProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// CertBlob1 structure represents CERT_BLOB_1 RPC structure.
type CertBlob1 struct {
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
	Data       []byte `idl:"name:pbData;size_is:(cbData)" json:"data"`
}

func (o *CertBlob1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CertBlob1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		_ptr_pbData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WritePointer(&o.Data, _ptr_pbData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertBlob1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_pbData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
	_s_pbData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_pbData, _ptr_pbData); err != nil {
		return err
	}
	return nil
}

// CertEKU1 structure represents CERT_EKU_1 RPC structure.
type CertEKU1 struct {
	Size     uint32 `idl:"name:dwSize" json:"size"`
	IsEKUOID bool   `idl:"name:IsEKUOID" json:"is_ekuoid"`
	EKU      string `idl:"name:pwszEKU;size_is:(dwSize)" json:"eku"`
}

func (o *CertEKU1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.EKU != "" && o.Size == 0 {
		o.Size = uint32(ndr.UTF16Len(o.EKU))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CertEKU1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if !o.IsEKUOID {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.EKU != "" || o.Size > 0 {
		_ptr_pwszEKU := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_EKU_buf := utf16.Encode([]rune(o.EKU))
			if uint64(len(_EKU_buf)) > sizeInfo[0] {
				_EKU_buf = _EKU_buf[:sizeInfo[0]]
			}
			for i1 := range _EKU_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_EKU_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_EKU_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EKU, _ptr_pwszEKU); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertEKU1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	var _bIsEKUOID int32
	if err := w.ReadData(&_bIsEKUOID); err != nil {
		return err
	}
	o.IsEKUOID = _bIsEKUOID != 0
	_ptr_pwszEKU := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		var _EKU_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _EKU_buf", sizeInfo[0])
		}
		_EKU_buf = make([]uint16, sizeInfo[0])
		for i1 := range _EKU_buf {
			i1 := i1
			if err := w.ReadData(&_EKU_buf[i1]); err != nil {
				return err
			}
		}
		o.EKU = strings.TrimRight(string(utf16.Decode(_EKU_buf)), ndr.ZeroString)
		return nil
	})
	_s_pwszEKU := func(ptr interface{}) { o.EKU = *ptr.(*string) }
	if err := w.ReadPointer(&o.EKU, _s_pwszEKU, _ptr_pwszEKU); err != nil {
		return err
	}
	return nil
}

// IKEv2TunnelConfigParams1 structure represents IKEV2_TUNNEL_CONFIG_PARAMS_1 RPC structure.
type IKEv2TunnelConfigParams1 struct {
	IdleTimeout                uint32       `idl:"name:dwIdleTimeout" json:"idle_timeout"`
	NetworkBlackoutTime        uint32       `idl:"name:dwNetworkBlackoutTime" json:"network_blackout_time"`
	SALifeTime                 uint32       `idl:"name:dwSaLifeTime" json:"sa_life_time"`
	SADataSizeForRenegotiation uint32       `idl:"name:dwSaDataSizeForRenegotiation" json:"sa_data_size_for_renegotiation"`
	ConfigOptions              uint32       `idl:"name:dwConfigOptions" json:"config_options"`
	TotalCertificates          uint32       `idl:"name:dwTotalCertificates" json:"total_certificates"`
	CertificateNames           []*CertBlob1 `idl:"name:certificateNames;size_is:(dwTotalCertificates)" json:"certificate_names"`
}

func (o *IKEv2TunnelConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.CertificateNames != nil && o.TotalCertificates == 0 {
		o.TotalCertificates = uint32(len(o.CertificateNames))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2TunnelConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.NetworkBlackoutTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SALifeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	if err := w.WriteData(o.ConfigOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalCertificates); err != nil {
		return err
	}
	if o.CertificateNames != nil || o.TotalCertificates > 0 {
		_ptr_certificateNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.TotalCertificates)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.CertificateNames {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.CertificateNames[i1] != nil {
					if err := o.CertificateNames[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.CertificateNames); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertificateNames, _ptr_certificateNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IKEv2TunnelConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetworkBlackoutTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SALifeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConfigOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalCertificates); err != nil {
		return err
	}
	_ptr_certificateNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.TotalCertificates > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.TotalCertificates)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CertificateNames", sizeInfo[0])
		}
		o.CertificateNames = make([]*CertBlob1, sizeInfo[0])
		for i1 := range o.CertificateNames {
			i1 := i1
			if o.CertificateNames[i1] == nil {
				o.CertificateNames[i1] = &CertBlob1{}
			}
			if err := o.CertificateNames[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_certificateNames := func(ptr interface{}) { o.CertificateNames = *ptr.(*[]*CertBlob1) }
	if err := w.ReadPointer(&o.CertificateNames, _s_certificateNames, _ptr_certificateNames); err != nil {
		return err
	}
	return nil
}

// RouterCustomIKEv2Policy0 structure represents ROUTER_CUSTOM_IKEv2_POLICY_0 RPC structure.
type RouterCustomIKEv2Policy0 struct {
	IntegrityMethod         uint32 `idl:"name:dwIntegrityMethod" json:"integrity_method"`
	EncryptionMethod        uint32 `idl:"name:dwEncryptionMethod" json:"encryption_method"`
	CipherTransformConstant uint32 `idl:"name:dwCipherTransformConstant" json:"cipher_transform_constant"`
	AuthTransformConstant   uint32 `idl:"name:dwAuthTransformConstant" json:"auth_transform_constant"`
	PFSGroup                uint32 `idl:"name:dwPfsGroup" json:"pfs_group"`
	DHGroup                 uint32 `idl:"name:dwDhGroup" json:"dh_group"`
}

func (o *RouterCustomIKEv2Policy0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterCustomIKEv2Policy0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IntegrityMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.CipherTransformConstant); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthTransformConstant); err != nil {
		return err
	}
	if err := w.WriteData(o.PFSGroup); err != nil {
		return err
	}
	if err := w.WriteData(o.DHGroup); err != nil {
		return err
	}
	return nil
}
func (o *RouterCustomIKEv2Policy0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IntegrityMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.CipherTransformConstant); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthTransformConstant); err != nil {
		return err
	}
	if err := w.ReadData(&o.PFSGroup); err != nil {
		return err
	}
	if err := w.ReadData(&o.DHGroup); err != nil {
		return err
	}
	return nil
}

// RouterCustomL2TPPolicy0 structure represents ROUTER_CUSTOM_L2TP_POLICY_0 RPC structure.
type RouterCustomL2TPPolicy0 struct {
	IntegrityMethod         uint32 `idl:"name:dwIntegrityMethod" json:"integrity_method"`
	EncryptionMethod        uint32 `idl:"name:dwEncryptionMethod" json:"encryption_method"`
	CipherTransformConstant uint32 `idl:"name:dwCipherTransformConstant" json:"cipher_transform_constant"`
	AuthTransformConstant   uint32 `idl:"name:dwAuthTransformConstant" json:"auth_transform_constant"`
	PFSGroup                uint32 `idl:"name:dwPfsGroup" json:"pfs_group"`
	DHGroup                 uint32 `idl:"name:dwDhGroup" json:"dh_group"`
}

func (o *RouterCustomL2TPPolicy0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterCustomL2TPPolicy0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IntegrityMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.CipherTransformConstant); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthTransformConstant); err != nil {
		return err
	}
	if err := w.WriteData(o.PFSGroup); err != nil {
		return err
	}
	if err := w.WriteData(o.DHGroup); err != nil {
		return err
	}
	return nil
}
func (o *RouterCustomL2TPPolicy0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IntegrityMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.CipherTransformConstant); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthTransformConstant); err != nil {
		return err
	}
	if err := w.ReadData(&o.PFSGroup); err != nil {
		return err
	}
	if err := w.ReadData(&o.DHGroup); err != nil {
		return err
	}
	return nil
}

// RouterIKEv2InterfaceCustomConfig0 structure represents ROUTER_IKEv2_IF_CUSTOM_CONFIG_0 RPC structure.
type RouterIKEv2InterfaceCustomConfig0 struct {
	SALifeTime      uint32                    `idl:"name:dwSaLifeTime" json:"sa_life_time"`
	SADataSize      uint32                    `idl:"name:dwSaDataSize" json:"sa_data_size"`
	CertificateName *CertBlob1                `idl:"name:certificateName" json:"certificate_name"`
	CustomPolicy    *RouterCustomIKEv2Policy0 `idl:"name:customPolicy" json:"custom_policy"`
}

func (o *RouterIKEv2InterfaceCustomConfig0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterIKEv2InterfaceCustomConfig0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SALifeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SADataSize); err != nil {
		return err
	}
	if o.CertificateName != nil {
		if err := o.CertificateName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CustomPolicy != nil {
		_ptr_customPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CustomPolicy != nil {
				if err := o.CustomPolicy.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RouterCustomIKEv2Policy0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CustomPolicy, _ptr_customPolicy); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RouterIKEv2InterfaceCustomConfig0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SALifeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SADataSize); err != nil {
		return err
	}
	if o.CertificateName == nil {
		o.CertificateName = &CertBlob1{}
	}
	if err := o.CertificateName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_customPolicy := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CustomPolicy == nil {
			o.CustomPolicy = &RouterCustomIKEv2Policy0{}
		}
		if err := o.CustomPolicy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_customPolicy := func(ptr interface{}) { o.CustomPolicy = *ptr.(**RouterCustomIKEv2Policy0) }
	if err := w.ReadPointer(&o.CustomPolicy, _s_customPolicy, _ptr_customPolicy); err != nil {
		return err
	}
	return nil
}

// RouterIKEv2InterfaceCustomConfig1 structure represents ROUTER_IKEv2_IF_CUSTOM_CONFIG_1 RPC structure.
type RouterIKEv2InterfaceCustomConfig1 struct {
	SALifeTime      uint32                    `idl:"name:dwSaLifeTime" json:"sa_life_time"`
	SADataSize      uint32                    `idl:"name:dwSaDataSize" json:"sa_data_size"`
	CertificateName *CertBlob1                `idl:"name:certificateName" json:"certificate_name"`
	CustomPolicy    *RouterCustomIKEv2Policy0 `idl:"name:customPolicy" json:"custom_policy"`
	CertificateHash *CertBlob1                `idl:"name:certificateHash" json:"certificate_hash"`
}

func (o *RouterIKEv2InterfaceCustomConfig1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterIKEv2InterfaceCustomConfig1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SALifeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SADataSize); err != nil {
		return err
	}
	if o.CertificateName != nil {
		if err := o.CertificateName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CustomPolicy != nil {
		_ptr_customPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CustomPolicy != nil {
				if err := o.CustomPolicy.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RouterCustomIKEv2Policy0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CustomPolicy, _ptr_customPolicy); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.CertificateHash != nil {
		if err := o.CertificateHash.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RouterIKEv2InterfaceCustomConfig1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SALifeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SADataSize); err != nil {
		return err
	}
	if o.CertificateName == nil {
		o.CertificateName = &CertBlob1{}
	}
	if err := o.CertificateName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_customPolicy := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CustomPolicy == nil {
			o.CustomPolicy = &RouterCustomIKEv2Policy0{}
		}
		if err := o.CustomPolicy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_customPolicy := func(ptr interface{}) { o.CustomPolicy = *ptr.(**RouterCustomIKEv2Policy0) }
	if err := w.ReadPointer(&o.CustomPolicy, _s_customPolicy, _ptr_customPolicy); err != nil {
		return err
	}
	if o.CertificateHash == nil {
		o.CertificateHash = &CertBlob1{}
	}
	if err := o.CertificateHash.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// InterfaceCustominfoex0 structure represents MPR_IF_CUSTOMINFOEX_0 RPC structure.
type InterfaceCustominfoex0 struct {
	Header            *ObjectHeaderIDL                   `idl:"name:Header" json:"header"`
	Flags             uint32                             `idl:"name:dwFlags" json:"flags"`
	CustomIKEv2Config *RouterIKEv2InterfaceCustomConfig0 `idl:"name:customIkev2Config" json:"custom_ikev2_config"`
}

func (o *InterfaceCustominfoex0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceCustominfoex0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.CustomIKEv2Config != nil {
		if err := o.CustomIKEv2Config.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RouterIKEv2InterfaceCustomConfig0{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceCustominfoex0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.CustomIKEv2Config == nil {
		o.CustomIKEv2Config = &RouterIKEv2InterfaceCustomConfig0{}
	}
	if err := o.CustomIKEv2Config.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// InterfaceCustominfoex1 structure represents MPR_IF_CUSTOMINFOEX_1 RPC structure.
type InterfaceCustominfoex1 struct {
	Header            *ObjectHeaderIDL                   `idl:"name:Header" json:"header"`
	Flags             uint32                             `idl:"name:dwFlags" json:"flags"`
	CustomIKEv2Config *RouterIKEv2InterfaceCustomConfig1 `idl:"name:customIkev2Config" json:"custom_ikev2_config"`
}

func (o *InterfaceCustominfoex1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceCustominfoex1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.CustomIKEv2Config != nil {
		if err := o.CustomIKEv2Config.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RouterIKEv2InterfaceCustomConfig1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceCustominfoex1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.CustomIKEv2Config == nil {
		o.CustomIKEv2Config = &RouterIKEv2InterfaceCustomConfig1{}
	}
	if err := o.CustomIKEv2Config.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// InterfaceCustominfoexIDL structure represents MPR_IF_CUSTOMINFOEX_IDL RPC union.
type InterfaceCustominfoexIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *InterfaceCustominfoexIDL_InterfaceConfigObj1
	// *InterfaceCustominfoexIDL_InterfaceConfigObj2
	Value is_InterfaceCustominfoexIDL `json:"value"`
}

func (o *InterfaceCustominfoexIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *InterfaceCustominfoexIDL_InterfaceConfigObj1:
		if value != nil {
			return value.InterfaceConfigObj1
		}
	case *InterfaceCustominfoexIDL_InterfaceConfigObj2:
		if value != nil {
			return value.InterfaceConfigObj2
		}
	}
	return nil
}

type is_InterfaceCustominfoexIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_InterfaceCustominfoexIDL()
}

func (o *InterfaceCustominfoexIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint8(o.Revision)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		_o, _ := o.Value.(*InterfaceCustominfoexIDL_InterfaceConfigObj1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InterfaceCustominfoexIDL_InterfaceConfigObj1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*InterfaceCustominfoexIDL_InterfaceConfigObj2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InterfaceCustominfoexIDL_InterfaceConfigObj2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *InterfaceCustominfoexIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint8)(&o.Revision)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		o.Value = &InterfaceCustominfoexIDL_InterfaceConfigObj1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &InterfaceCustominfoexIDL_InterfaceConfigObj2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// InterfaceCustominfoexIDL_InterfaceConfigObj1 structure represents MPR_IF_CUSTOMINFOEX_IDL RPC union arm.
//
// It has following labels: 1
type InterfaceCustominfoexIDL_InterfaceConfigObj1 struct {
	InterfaceConfigObj1 *InterfaceCustominfoex0 `idl:"name:IfConfigObj1" json:"interface_config_obj1"`
}

func (*InterfaceCustominfoexIDL_InterfaceConfigObj1) is_InterfaceCustominfoexIDL() {}

func (o *InterfaceCustominfoexIDL_InterfaceConfigObj1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.InterfaceConfigObj1 != nil {
		if err := o.InterfaceConfigObj1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceCustominfoex0{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceCustominfoexIDL_InterfaceConfigObj1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.InterfaceConfigObj1 == nil {
		o.InterfaceConfigObj1 = &InterfaceCustominfoex0{}
	}
	if err := o.InterfaceConfigObj1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// InterfaceCustominfoexIDL_InterfaceConfigObj2 structure represents MPR_IF_CUSTOMINFOEX_IDL RPC union arm.
//
// It has following labels: 2
type InterfaceCustominfoexIDL_InterfaceConfigObj2 struct {
	InterfaceConfigObj2 *InterfaceCustominfoex1 `idl:"name:IfConfigObj2" json:"interface_config_obj2"`
}

func (*InterfaceCustominfoexIDL_InterfaceConfigObj2) is_InterfaceCustominfoexIDL() {}

func (o *InterfaceCustominfoexIDL_InterfaceConfigObj2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.InterfaceConfigObj2 != nil {
		if err := o.InterfaceConfigObj2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceCustominfoex1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceCustominfoexIDL_InterfaceConfigObj2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.InterfaceConfigObj2 == nil {
		o.InterfaceConfigObj2 = &InterfaceCustominfoex1{}
	}
	if err := o.InterfaceConfigObj2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IKEv2TunnelConfigParams2 structure represents IKEV2_TUNNEL_CONFIG_PARAMS_2 RPC structure.
type IKEv2TunnelConfigParams2 struct {
	IdleTimeout                uint32                    `idl:"name:dwIdleTimeout" json:"idle_timeout"`
	NetworkBlackoutTime        uint32                    `idl:"name:dwNetworkBlackoutTime" json:"network_blackout_time"`
	SALifeTime                 uint32                    `idl:"name:dwSaLifeTime" json:"sa_life_time"`
	SADataSizeForRenegotiation uint32                    `idl:"name:dwSaDataSizeForRenegotiation" json:"sa_data_size_for_renegotiation"`
	ConfigOptions              uint32                    `idl:"name:dwConfigOptions" json:"config_options"`
	TotalCertificates          uint32                    `idl:"name:dwTotalCertificates" json:"total_certificates"`
	CertificateNames           []*CertBlob1              `idl:"name:certificateNames;size_is:(dwTotalCertificates)" json:"certificate_names"`
	MachineCertificateName     *CertBlob1                `idl:"name:machineCertificateName" json:"machine_certificate_name"`
	EncryptionType             uint32                    `idl:"name:dwEncryptionType" json:"encryption_type"`
	CustomPolicy               *RouterCustomIKEv2Policy0 `idl:"name:customPolicy" json:"custom_policy"`
}

func (o *IKEv2TunnelConfigParams2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.CertificateNames != nil && o.TotalCertificates == 0 {
		o.TotalCertificates = uint32(len(o.CertificateNames))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2TunnelConfigParams2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.NetworkBlackoutTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SALifeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	if err := w.WriteData(o.ConfigOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalCertificates); err != nil {
		return err
	}
	if o.CertificateNames != nil || o.TotalCertificates > 0 {
		_ptr_certificateNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.TotalCertificates)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.CertificateNames {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.CertificateNames[i1] != nil {
					if err := o.CertificateNames[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.CertificateNames); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertificateNames, _ptr_certificateNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.MachineCertificateName != nil {
		if err := o.MachineCertificateName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EncryptionType); err != nil {
		return err
	}
	if o.CustomPolicy != nil {
		_ptr_customPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CustomPolicy != nil {
				if err := o.CustomPolicy.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RouterCustomIKEv2Policy0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CustomPolicy, _ptr_customPolicy); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IKEv2TunnelConfigParams2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetworkBlackoutTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SALifeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConfigOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalCertificates); err != nil {
		return err
	}
	_ptr_certificateNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.TotalCertificates > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.TotalCertificates)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CertificateNames", sizeInfo[0])
		}
		o.CertificateNames = make([]*CertBlob1, sizeInfo[0])
		for i1 := range o.CertificateNames {
			i1 := i1
			if o.CertificateNames[i1] == nil {
				o.CertificateNames[i1] = &CertBlob1{}
			}
			if err := o.CertificateNames[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_certificateNames := func(ptr interface{}) { o.CertificateNames = *ptr.(*[]*CertBlob1) }
	if err := w.ReadPointer(&o.CertificateNames, _s_certificateNames, _ptr_certificateNames); err != nil {
		return err
	}
	if o.MachineCertificateName == nil {
		o.MachineCertificateName = &CertBlob1{}
	}
	if err := o.MachineCertificateName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionType); err != nil {
		return err
	}
	_ptr_customPolicy := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CustomPolicy == nil {
			o.CustomPolicy = &RouterCustomIKEv2Policy0{}
		}
		if err := o.CustomPolicy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_customPolicy := func(ptr interface{}) { o.CustomPolicy = *ptr.(**RouterCustomIKEv2Policy0) }
	if err := w.ReadPointer(&o.CustomPolicy, _s_customPolicy, _ptr_customPolicy); err != nil {
		return err
	}
	return nil
}

// IKEv2TunnelConfigParams3 structure represents IKEV2_TUNNEL_CONFIG_PARAMS_3 RPC structure.
type IKEv2TunnelConfigParams3 struct {
	IdleTimeout                uint32                    `idl:"name:dwIdleTimeout" json:"idle_timeout"`
	NetworkBlackoutTime        uint32                    `idl:"name:dwNetworkBlackoutTime" json:"network_blackout_time"`
	SALifeTime                 uint32                    `idl:"name:dwSaLifeTime" json:"sa_life_time"`
	SADataSizeForRenegotiation uint32                    `idl:"name:dwSaDataSizeForRenegotiation" json:"sa_data_size_for_renegotiation"`
	ConfigOptions              uint32                    `idl:"name:dwConfigOptions" json:"config_options"`
	TotalCertificates          uint32                    `idl:"name:dwTotalCertificates" json:"total_certificates"`
	CertificateNames           []*CertBlob1              `idl:"name:certificateNames;size_is:(dwTotalCertificates)" json:"certificate_names"`
	MachineCertificateName     *CertBlob1                `idl:"name:machineCertificateName" json:"machine_certificate_name"`
	EncryptionType             uint32                    `idl:"name:dwEncryptionType" json:"encryption_type"`
	CustomPolicy               *RouterCustomIKEv2Policy0 `idl:"name:customPolicy" json:"custom_policy"`
	TotalEkus                  uint32                    `idl:"name:dwTotalEkus" json:"total_ekus"`
	CertificateEkUs            []*CertEKU1               `idl:"name:certificateEKUs;size_is:(dwTotalEkus)" json:"certificate_ek_us"`
	MachineCertificateHash     *CertBlob1                `idl:"name:machineCertificateHash" json:"machine_certificate_hash"`
}

func (o *IKEv2TunnelConfigParams3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.CertificateNames != nil && o.TotalCertificates == 0 {
		o.TotalCertificates = uint32(len(o.CertificateNames))
	}
	if o.CertificateEkUs != nil && o.TotalEkus == 0 {
		o.TotalEkus = uint32(len(o.CertificateEkUs))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2TunnelConfigParams3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.NetworkBlackoutTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SALifeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	if err := w.WriteData(o.ConfigOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalCertificates); err != nil {
		return err
	}
	if o.CertificateNames != nil || o.TotalCertificates > 0 {
		_ptr_certificateNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.TotalCertificates)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.CertificateNames {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.CertificateNames[i1] != nil {
					if err := o.CertificateNames[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.CertificateNames); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertificateNames, _ptr_certificateNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.MachineCertificateName != nil {
		if err := o.MachineCertificateName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EncryptionType); err != nil {
		return err
	}
	if o.CustomPolicy != nil {
		_ptr_customPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CustomPolicy != nil {
				if err := o.CustomPolicy.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RouterCustomIKEv2Policy0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CustomPolicy, _ptr_customPolicy); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TotalEkus); err != nil {
		return err
	}
	if o.CertificateEkUs != nil || o.TotalEkus > 0 {
		_ptr_certificateEKUs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.TotalEkus)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.CertificateEkUs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.CertificateEkUs[i1] != nil {
					if err := o.CertificateEkUs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CertEKU1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.CertificateEkUs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&CertEKU1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertificateEkUs, _ptr_certificateEKUs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.MachineCertificateHash != nil {
		if err := o.MachineCertificateHash.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IKEv2TunnelConfigParams3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetworkBlackoutTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SALifeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConfigOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalCertificates); err != nil {
		return err
	}
	_ptr_certificateNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.TotalCertificates > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.TotalCertificates)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CertificateNames", sizeInfo[0])
		}
		o.CertificateNames = make([]*CertBlob1, sizeInfo[0])
		for i1 := range o.CertificateNames {
			i1 := i1
			if o.CertificateNames[i1] == nil {
				o.CertificateNames[i1] = &CertBlob1{}
			}
			if err := o.CertificateNames[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_certificateNames := func(ptr interface{}) { o.CertificateNames = *ptr.(*[]*CertBlob1) }
	if err := w.ReadPointer(&o.CertificateNames, _s_certificateNames, _ptr_certificateNames); err != nil {
		return err
	}
	if o.MachineCertificateName == nil {
		o.MachineCertificateName = &CertBlob1{}
	}
	if err := o.MachineCertificateName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionType); err != nil {
		return err
	}
	_ptr_customPolicy := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CustomPolicy == nil {
			o.CustomPolicy = &RouterCustomIKEv2Policy0{}
		}
		if err := o.CustomPolicy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_customPolicy := func(ptr interface{}) { o.CustomPolicy = *ptr.(**RouterCustomIKEv2Policy0) }
	if err := w.ReadPointer(&o.CustomPolicy, _s_customPolicy, _ptr_customPolicy); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalEkus); err != nil {
		return err
	}
	_ptr_certificateEKUs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.TotalEkus > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.TotalEkus)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CertificateEkUs", sizeInfo[0])
		}
		o.CertificateEkUs = make([]*CertEKU1, sizeInfo[0])
		for i1 := range o.CertificateEkUs {
			i1 := i1
			if o.CertificateEkUs[i1] == nil {
				o.CertificateEkUs[i1] = &CertEKU1{}
			}
			if err := o.CertificateEkUs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_certificateEKUs := func(ptr interface{}) { o.CertificateEkUs = *ptr.(*[]*CertEKU1) }
	if err := w.ReadPointer(&o.CertificateEkUs, _s_certificateEKUs, _ptr_certificateEKUs); err != nil {
		return err
	}
	if o.MachineCertificateHash == nil {
		o.MachineCertificateHash = &CertBlob1{}
	}
	if err := o.MachineCertificateHash.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// L2TPTunnelConfigParams1 structure represents L2TP_TUNNEL_CONFIG_PARAMS_1 RPC structure.
type L2TPTunnelConfigParams1 struct {
	IdleTimeout                uint32                   `idl:"name:dwIdleTimeout" json:"idle_timeout"`
	EncryptionType             uint32                   `idl:"name:dwEncryptionType" json:"encryption_type"`
	SALifeTime                 uint32                   `idl:"name:dwSaLifeTime" json:"sa_life_time"`
	SADataSizeForRenegotiation uint32                   `idl:"name:dwSaDataSizeForRenegotiation" json:"sa_data_size_for_renegotiation"`
	CustomPolicy               *RouterCustomL2TPPolicy0 `idl:"name:customPolicy" json:"custom_policy"`
}

func (o *L2TPTunnelConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *L2TPTunnelConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionType); err != nil {
		return err
	}
	if err := w.WriteData(o.SALifeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	if o.CustomPolicy != nil {
		_ptr_customPolicy := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CustomPolicy != nil {
				if err := o.CustomPolicy.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RouterCustomL2TPPolicy0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CustomPolicy, _ptr_customPolicy); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *L2TPTunnelConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionType); err != nil {
		return err
	}
	if err := w.ReadData(&o.SALifeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SADataSizeForRenegotiation); err != nil {
		return err
	}
	_ptr_customPolicy := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CustomPolicy == nil {
			o.CustomPolicy = &RouterCustomL2TPPolicy0{}
		}
		if err := o.CustomPolicy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_customPolicy := func(ptr interface{}) { o.CustomPolicy = *ptr.(**RouterCustomL2TPPolicy0) }
	if err := w.ReadPointer(&o.CustomPolicy, _s_customPolicy, _ptr_customPolicy); err != nil {
		return err
	}
	return nil
}

// IKEv2ConfigParams1 structure represents IKEV2_CONFIG_PARAMS_1 RPC structure.
type IKEv2ConfigParams1 struct {
	PortsLength            uint32                    `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags              uint32                    `idl:"name:dwPortFlags" json:"port_flags"`
	TunnelConfigParamFlags uint32                    `idl:"name:dwTunnelConfigParamFlags" json:"tunnel_config_param_flags"`
	TunnelConfigParams     *IKEv2TunnelConfigParams1 `idl:"name:TunnelConfigParams" json:"tunnel_config_params"`
}

func (o *IKEv2ConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2ConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams != nil {
		if err := o.TunnelConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2TunnelConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IKEv2ConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams == nil {
		o.TunnelConfigParams = &IKEv2TunnelConfigParams1{}
	}
	if err := o.TunnelConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IKEv2ConfigParams2 structure represents IKEV2_CONFIG_PARAMS_2 RPC structure.
type IKEv2ConfigParams2 struct {
	PortsLength            uint32                    `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags              uint32                    `idl:"name:dwPortFlags" json:"port_flags"`
	TunnelConfigParamFlags uint32                    `idl:"name:dwTunnelConfigParamFlags" json:"tunnel_config_param_flags"`
	TunnelConfigParams     *IKEv2TunnelConfigParams2 `idl:"name:TunnelConfigParams" json:"tunnel_config_params"`
}

func (o *IKEv2ConfigParams2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2ConfigParams2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams != nil {
		if err := o.TunnelConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2TunnelConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IKEv2ConfigParams2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams == nil {
		o.TunnelConfigParams = &IKEv2TunnelConfigParams2{}
	}
	if err := o.TunnelConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IKEv2ConfigParams3 structure represents IKEV2_CONFIG_PARAMS_3 RPC structure.
type IKEv2ConfigParams3 struct {
	PortsLength            uint32                    `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags              uint32                    `idl:"name:dwPortFlags" json:"port_flags"`
	TunnelConfigParamFlags uint32                    `idl:"name:dwTunnelConfigParamFlags" json:"tunnel_config_param_flags"`
	TunnelConfigParams     *IKEv2TunnelConfigParams3 `idl:"name:TunnelConfigParams" json:"tunnel_config_params"`
}

func (o *IKEv2ConfigParams3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IKEv2ConfigParams3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams != nil {
		if err := o.TunnelConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2TunnelConfigParams3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IKEv2ConfigParams3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams == nil {
		o.TunnelConfigParams = &IKEv2TunnelConfigParams3{}
	}
	if err := o.TunnelConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PPTPConfigParams1 structure represents PPTP_CONFIG_PARAMS_1 RPC structure.
type PPTPConfigParams1 struct {
	PortsLength uint32 `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags   uint32 `idl:"name:dwPortFlags" json:"port_flags"`
}

func (o *PPTPConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPTPConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PortFlags); err != nil {
		return err
	}
	return nil
}
func (o *PPTPConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortFlags); err != nil {
		return err
	}
	return nil
}

// L2TPConfigParams1 structure represents L2TP_CONFIG_PARAMS_1 RPC structure.
type L2TPConfigParams1 struct {
	PortsLength uint32 `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags   uint32 `idl:"name:dwPortFlags" json:"port_flags"`
}

func (o *L2TPConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *L2TPConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PortFlags); err != nil {
		return err
	}
	return nil
}
func (o *L2TPConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortFlags); err != nil {
		return err
	}
	return nil
}

// L2TPConfigParams2 structure represents L2TP_CONFIG_PARAMS_2 RPC structure.
type L2TPConfigParams2 struct {
	PortsLength            uint32                   `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags              uint32                   `idl:"name:dwPortFlags" json:"port_flags"`
	TunnelConfigParamFlags uint32                   `idl:"name:dwTunnelConfigParamFlags" json:"tunnel_config_param_flags"`
	TunnelConfigParams     *L2TPTunnelConfigParams1 `idl:"name:TunnelConfigParams" json:"tunnel_config_params"`
}

func (o *L2TPConfigParams2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *L2TPConfigParams2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams != nil {
		if err := o.TunnelConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&L2TPTunnelConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *L2TPConfigParams2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.TunnelConfigParamFlags); err != nil {
		return err
	}
	if o.TunnelConfigParams == nil {
		o.TunnelConfigParams = &L2TPTunnelConfigParams1{}
	}
	if err := o.TunnelConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SSTPCertInfo1 structure represents SSTP_CERT_INFO_1 RPC structure.
type SSTPCertInfo1 struct {
	IsDefault bool       `idl:"name:isDefault" json:"is_default"`
	CertBlob  *CertBlob1 `idl:"name:certBlob" json:"cert_blob"`
}

func (o *SSTPCertInfo1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SSTPCertInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if !o.IsDefault {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.CertBlob != nil {
		if err := o.CertBlob.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertBlob1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SSTPCertInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	var _bIsDefault int32
	if err := w.ReadData(&_bIsDefault); err != nil {
		return err
	}
	o.IsDefault = _bIsDefault != 0
	if o.CertBlob == nil {
		o.CertBlob = &CertBlob1{}
	}
	if err := o.CertBlob.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SSTPConfigParams1 structure represents SSTP_CONFIG_PARAMS_1 RPC structure.
type SSTPConfigParams1 struct {
	PortsLength     uint32         `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags       uint32         `idl:"name:dwPortFlags" json:"port_flags"`
	IsUseHTTPS      bool           `idl:"name:isUseHttps" json:"is_use_https"`
	CertAlgorithm   uint32         `idl:"name:certAlgorithm" json:"cert_algorithm"`
	SSTPCertDetails *SSTPCertInfo1 `idl:"name:sstpCertDetails" json:"sstp_cert_details"`
}

func (o *SSTPConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SSTPConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PortFlags); err != nil {
		return err
	}
	if !o.IsUseHTTPS {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CertAlgorithm); err != nil {
		return err
	}
	if o.SSTPCertDetails != nil {
		if err := o.SSTPCertDetails.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SSTPCertInfo1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SSTPConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortFlags); err != nil {
		return err
	}
	var _bIsUseHTTPS int32
	if err := w.ReadData(&_bIsUseHTTPS); err != nil {
		return err
	}
	o.IsUseHTTPS = _bIsUseHTTPS != 0
	if err := w.ReadData(&o.CertAlgorithm); err != nil {
		return err
	}
	if o.SSTPCertDetails == nil {
		o.SSTPCertDetails = &SSTPCertInfo1{}
	}
	if err := o.SSTPCertDetails.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TunnelConfigParams1 structure represents MPRAPI_TUNNEL_CONFIG_PARAMS_1 RPC structure.
type TunnelConfigParams1 struct {
	IKEConfigParams  *IKEv2ConfigParams1 `idl:"name:IkeConfigParams" json:"ike_config_params"`
	PPTPConfigParams *PPTPConfigParams1  `idl:"name:PptpConfigParams" json:"pptp_config_params"`
	L2TPConfigParams *L2TPConfigParams1  `idl:"name:L2tpConfigParams" json:"l2tp_config_params"`
	SSTPConfigParams *SSTPConfigParams1  `idl:"name:SstpConfigParams" json:"sstp_config_params"`
}

func (o *TunnelConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TunnelConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.IKEConfigParams != nil {
		if err := o.IKEConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2ConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PPTPConfigParams != nil {
		if err := o.PPTPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPTPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.L2TPConfigParams != nil {
		if err := o.L2TPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&L2TPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SSTPConfigParams != nil {
		if err := o.SSTPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SSTPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TunnelConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.IKEConfigParams == nil {
		o.IKEConfigParams = &IKEv2ConfigParams1{}
	}
	if err := o.IKEConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PPTPConfigParams == nil {
		o.PPTPConfigParams = &PPTPConfigParams1{}
	}
	if err := o.PPTPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.L2TPConfigParams == nil {
		o.L2TPConfigParams = &L2TPConfigParams1{}
	}
	if err := o.L2TPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SSTPConfigParams == nil {
		o.SSTPConfigParams = &SSTPConfigParams1{}
	}
	if err := o.SSTPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TunnelConfigParams2 structure represents MPRAPI_TUNNEL_CONFIG_PARAMS_2 RPC structure.
type TunnelConfigParams2 struct {
	IKEConfigParams  *IKEv2ConfigParams2 `idl:"name:IkeConfigParams" json:"ike_config_params"`
	PPTPConfigParams *PPTPConfigParams1  `idl:"name:PptpConfigParams" json:"pptp_config_params"`
	L2TPConfigParams *L2TPConfigParams1  `idl:"name:L2tpConfigParams" json:"l2tp_config_params"`
	SSTPConfigParams *SSTPConfigParams1  `idl:"name:SstpConfigParams" json:"sstp_config_params"`
}

func (o *TunnelConfigParams2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TunnelConfigParams2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.IKEConfigParams != nil {
		if err := o.IKEConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2ConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PPTPConfigParams != nil {
		if err := o.PPTPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPTPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.L2TPConfigParams != nil {
		if err := o.L2TPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&L2TPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SSTPConfigParams != nil {
		if err := o.SSTPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SSTPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TunnelConfigParams2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.IKEConfigParams == nil {
		o.IKEConfigParams = &IKEv2ConfigParams2{}
	}
	if err := o.IKEConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PPTPConfigParams == nil {
		o.PPTPConfigParams = &PPTPConfigParams1{}
	}
	if err := o.PPTPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.L2TPConfigParams == nil {
		o.L2TPConfigParams = &L2TPConfigParams1{}
	}
	if err := o.L2TPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SSTPConfigParams == nil {
		o.SSTPConfigParams = &SSTPConfigParams1{}
	}
	if err := o.SSTPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TunnelConfigParams3 structure represents MPRAPI_TUNNEL_CONFIG_PARAMS_3 RPC structure.
type TunnelConfigParams3 struct {
	IKEConfigParams  *IKEv2ConfigParams3 `idl:"name:IkeConfigParams" json:"ike_config_params"`
	PPTPConfigParams *PPTPConfigParams1  `idl:"name:PptpConfigParams" json:"pptp_config_params"`
	L2TPConfigParams *L2TPConfigParams2  `idl:"name:L2tpConfigParams" json:"l2tp_config_params"`
	SSTPConfigParams *SSTPConfigParams1  `idl:"name:SstpConfigParams" json:"sstp_config_params"`
}

func (o *TunnelConfigParams3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TunnelConfigParams3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.IKEConfigParams != nil {
		if err := o.IKEConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IKEv2ConfigParams3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PPTPConfigParams != nil {
		if err := o.PPTPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPTPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.L2TPConfigParams != nil {
		if err := o.L2TPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&L2TPConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SSTPConfigParams != nil {
		if err := o.SSTPConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SSTPConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TunnelConfigParams3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.IKEConfigParams == nil {
		o.IKEConfigParams = &IKEv2ConfigParams3{}
	}
	if err := o.IKEConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PPTPConfigParams == nil {
		o.PPTPConfigParams = &PPTPConfigParams1{}
	}
	if err := o.PPTPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.L2TPConfigParams == nil {
		o.L2TPConfigParams = &L2TPConfigParams2{}
	}
	if err := o.L2TPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SSTPConfigParams == nil {
		o.SSTPConfigParams = &SSTPConfigParams1{}
	}
	if err := o.SSTPConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerEx1 structure represents MPR_SERVER_EX_1 RPC structure.
type ServerEx1 struct {
	Header       *ObjectHeaderIDL     `idl:"name:Header" json:"header"`
	LANOnlyMode  bool                 `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime       uint32               `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts   uint32               `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse   uint32               `idl:"name:dwPortsInUse" json:"ports_in_use"`
	_            uint32               `idl:"name:Reserved"`
	ConfigParams *TunnelConfigParams1 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *ServerEx1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerEx1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.LANOnlyMode {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalPorts); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsInUse); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ConfigParams != nil {
		if err := o.ConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TunnelConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerEx1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bLANOnlyMode int32
	if err := w.ReadData(&_bLANOnlyMode); err != nil {
		return err
	}
	o.LANOnlyMode = _bLANOnlyMode != 0
	if err := w.ReadData(&o.UpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalPorts); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsInUse); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &TunnelConfigParams1{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerEx2 structure represents MPR_SERVER_EX_2 RPC structure.
type ServerEx2 struct {
	Header       *ObjectHeaderIDL     `idl:"name:Header" json:"header"`
	LANOnlyMode  bool                 `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime       uint32               `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts   uint32               `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse   uint32               `idl:"name:dwPortsInUse" json:"ports_in_use"`
	_            uint32               `idl:"name:Reserved"`
	ConfigParams *TunnelConfigParams2 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *ServerEx2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerEx2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.LANOnlyMode {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalPorts); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsInUse); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ConfigParams != nil {
		if err := o.ConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TunnelConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerEx2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bLANOnlyMode int32
	if err := w.ReadData(&_bLANOnlyMode); err != nil {
		return err
	}
	o.LANOnlyMode = _bLANOnlyMode != 0
	if err := w.ReadData(&o.UpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalPorts); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsInUse); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &TunnelConfigParams2{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerEx3 structure represents MPR_SERVER_EX_3 RPC structure.
type ServerEx3 struct {
	Header       *ObjectHeaderIDL     `idl:"name:Header" json:"header"`
	LANOnlyMode  bool                 `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime       uint32               `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts   uint32               `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse   uint32               `idl:"name:dwPortsInUse" json:"ports_in_use"`
	_            uint32               `idl:"name:Reserved"`
	ConfigParams *TunnelConfigParams3 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *ServerEx3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerEx3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.LANOnlyMode {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalPorts); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsInUse); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ConfigParams != nil {
		if err := o.ConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TunnelConfigParams3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerEx3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bLANOnlyMode int32
	if err := w.ReadData(&_bLANOnlyMode); err != nil {
		return err
	}
	o.LANOnlyMode = _bLANOnlyMode != 0
	if err := w.ReadData(&o.UpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalPorts); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsInUse); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &TunnelConfigParams3{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerExIDL structure represents MPR_SERVER_EX_IDL RPC union.
type ServerExIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *ServerExIDL_ServerConfig1
	// *ServerExIDL_ServerConfig2
	// *ServerExIDL_ServerConfig3
	Value is_ServerExIDL `json:"value"`
}

func (o *ServerExIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ServerExIDL_ServerConfig1:
		if value != nil {
			return value.ServerConfig1
		}
	case *ServerExIDL_ServerConfig2:
		if value != nil {
			return value.ServerConfig2
		}
	case *ServerExIDL_ServerConfig3:
		if value != nil {
			return value.ServerConfig3
		}
	}
	return nil
}

type is_ServerExIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ServerExIDL()
}

func (o *ServerExIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint8(o.Revision)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		_o, _ := o.Value.(*ServerExIDL_ServerConfig1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerExIDL_ServerConfig1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*ServerExIDL_ServerConfig2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerExIDL_ServerConfig2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(3):
		_o, _ := o.Value.(*ServerExIDL_ServerConfig3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerExIDL_ServerConfig3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *ServerExIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint8)(&o.Revision)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		o.Value = &ServerExIDL_ServerConfig1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &ServerExIDL_ServerConfig2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(3):
		o.Value = &ServerExIDL_ServerConfig3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// ServerExIDL_ServerConfig1 structure represents MPR_SERVER_EX_IDL RPC union arm.
//
// It has following labels: 1
type ServerExIDL_ServerConfig1 struct {
	ServerConfig1 *ServerEx1 `idl:"name:ServerConfig1" json:"server_config1"`
}

func (*ServerExIDL_ServerConfig1) is_ServerExIDL() {}

func (o *ServerExIDL_ServerConfig1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerConfig1 != nil {
		if err := o.ServerConfig1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ServerEx1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerExIDL_ServerConfig1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerConfig1 == nil {
		o.ServerConfig1 = &ServerEx1{}
	}
	if err := o.ServerConfig1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerExIDL_ServerConfig2 structure represents MPR_SERVER_EX_IDL RPC union arm.
//
// It has following labels: 2
type ServerExIDL_ServerConfig2 struct {
	ServerConfig2 *ServerEx2 `idl:"name:ServerConfig2" json:"server_config2"`
}

func (*ServerExIDL_ServerConfig2) is_ServerExIDL() {}

func (o *ServerExIDL_ServerConfig2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerConfig2 != nil {
		if err := o.ServerConfig2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ServerEx2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerExIDL_ServerConfig2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerConfig2 == nil {
		o.ServerConfig2 = &ServerEx2{}
	}
	if err := o.ServerConfig2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerExIDL_ServerConfig3 structure represents MPR_SERVER_EX_IDL RPC union arm.
//
// It has following labels: 3
type ServerExIDL_ServerConfig3 struct {
	ServerConfig3 *ServerEx3 `idl:"name:ServerConfig3" json:"server_config3"`
}

func (*ServerExIDL_ServerConfig3) is_ServerExIDL() {}

func (o *ServerExIDL_ServerConfig3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerConfig3 != nil {
		if err := o.ServerConfig3.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ServerEx3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerExIDL_ServerConfig3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerConfig3 == nil {
		o.ServerConfig3 = &ServerEx3{}
	}
	if err := o.ServerConfig3.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerSetConfigEx1 structure represents MPR_SERVER_SET_CONFIG_EX_1 RPC structure.
type ServerSetConfigEx1 struct {
	Header                *ObjectHeaderIDL     `idl:"name:Header" json:"header"`
	SetConfigForProtocols uint32               `idl:"name:setConfigForProtocols" json:"set_config_for_protocols"`
	ConfigParams          *TunnelConfigParams1 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *ServerSetConfigEx1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerSetConfigEx1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams != nil {
		if err := o.ConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TunnelConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSetConfigEx1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &TunnelConfigParams1{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerSetConfigEx2 structure represents MPR_SERVER_SET_CONFIG_EX_2 RPC structure.
type ServerSetConfigEx2 struct {
	Header                *ObjectHeaderIDL     `idl:"name:Header" json:"header"`
	SetConfigForProtocols uint32               `idl:"name:setConfigForProtocols" json:"set_config_for_protocols"`
	ConfigParams          *TunnelConfigParams2 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *ServerSetConfigEx2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerSetConfigEx2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams != nil {
		if err := o.ConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TunnelConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSetConfigEx2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &TunnelConfigParams2{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerSetConfigEx3 structure represents MPR_SERVER_SET_CONFIG_EX_3 RPC structure.
type ServerSetConfigEx3 struct {
	Header                *ObjectHeaderIDL     `idl:"name:Header" json:"header"`
	SetConfigForProtocols uint32               `idl:"name:setConfigForProtocols" json:"set_config_for_protocols"`
	ConfigParams          *TunnelConfigParams3 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *ServerSetConfigEx3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerSetConfigEx3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams != nil {
		if err := o.ConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TunnelConfigParams3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSetConfigEx3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &TunnelConfigParams3{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerSetConfigExIDL structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union.
type ServerSetConfigExIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *ServerSetConfigExIDL_ServerSetConfig1
	// *ServerSetConfigExIDL_ServerSetConfig2
	// *ServerSetConfigExIDL_ServerSetConfig3
	Value is_ServerSetConfigExIDL `json:"value"`
}

func (o *ServerSetConfigExIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ServerSetConfigExIDL_ServerSetConfig1:
		if value != nil {
			return value.ServerSetConfig1
		}
	case *ServerSetConfigExIDL_ServerSetConfig2:
		if value != nil {
			return value.ServerSetConfig2
		}
	case *ServerSetConfigExIDL_ServerSetConfig3:
		if value != nil {
			return value.ServerSetConfig3
		}
	}
	return nil
}

type is_ServerSetConfigExIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ServerSetConfigExIDL()
}

func (o *ServerSetConfigExIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint8(o.Revision)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		_o, _ := o.Value.(*ServerSetConfigExIDL_ServerSetConfig1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerSetConfigExIDL_ServerSetConfig1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*ServerSetConfigExIDL_ServerSetConfig2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerSetConfigExIDL_ServerSetConfig2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(3):
		_o, _ := o.Value.(*ServerSetConfigExIDL_ServerSetConfig3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerSetConfigExIDL_ServerSetConfig3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *ServerSetConfigExIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint8)(&o.Revision)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		o.Value = &ServerSetConfigExIDL_ServerSetConfig1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &ServerSetConfigExIDL_ServerSetConfig2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(3):
		o.Value = &ServerSetConfigExIDL_ServerSetConfig3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// ServerSetConfigExIDL_ServerSetConfig1 structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union arm.
//
// It has following labels: 1
type ServerSetConfigExIDL_ServerSetConfig1 struct {
	ServerSetConfig1 *ServerSetConfigEx1 `idl:"name:ServerSetConfig1" json:"server_set_config1"`
}

func (*ServerSetConfigExIDL_ServerSetConfig1) is_ServerSetConfigExIDL() {}

func (o *ServerSetConfigExIDL_ServerSetConfig1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerSetConfig1 != nil {
		if err := o.ServerSetConfig1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ServerSetConfigEx1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSetConfigExIDL_ServerSetConfig1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerSetConfig1 == nil {
		o.ServerSetConfig1 = &ServerSetConfigEx1{}
	}
	if err := o.ServerSetConfig1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerSetConfigExIDL_ServerSetConfig2 structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union arm.
//
// It has following labels: 2
type ServerSetConfigExIDL_ServerSetConfig2 struct {
	ServerSetConfig2 *ServerSetConfigEx2 `idl:"name:ServerSetConfig2" json:"server_set_config2"`
}

func (*ServerSetConfigExIDL_ServerSetConfig2) is_ServerSetConfigExIDL() {}

func (o *ServerSetConfigExIDL_ServerSetConfig2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerSetConfig2 != nil {
		if err := o.ServerSetConfig2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ServerSetConfigEx2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSetConfigExIDL_ServerSetConfig2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerSetConfig2 == nil {
		o.ServerSetConfig2 = &ServerSetConfigEx2{}
	}
	if err := o.ServerSetConfig2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ServerSetConfigExIDL_ServerSetConfig3 structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union arm.
//
// It has following labels: 3
type ServerSetConfigExIDL_ServerSetConfig3 struct {
	ServerSetConfig3 *ServerSetConfigEx3 `idl:"name:ServerSetConfig3" json:"server_set_config3"`
}

func (*ServerSetConfigExIDL_ServerSetConfig3) is_ServerSetConfigExIDL() {}

func (o *ServerSetConfigExIDL_ServerSetConfig3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerSetConfig3 != nil {
		if err := o.ServerSetConfig3.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ServerSetConfigEx3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSetConfigExIDL_ServerSetConfig3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerSetConfig3 == nil {
		o.ServerSetConfig3 = &ServerSetConfigEx3{}
	}
	if err := o.ServerSetConfig3.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASUpdateConnection1IDL structure represents RAS_UPDATE_CONNECTION_1_IDL RPC structure.
type RASUpdateConnection1IDL struct {
	Header                *ObjectHeaderIDL `idl:"name:Header" json:"header"`
	InterfaceIndex        uint32           `idl:"name:dwIfIndex" json:"interface_index"`
	RemoteEndpointAddress []uint16         `idl:"name:wszRemoteEndpointAddress" json:"remote_endpoint_address"`
}

func (o *RASUpdateConnection1IDL) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASUpdateConnection1IDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	for i1 := range o.RemoteEndpointAddress {
		i1 := i1
		if uint64(i1) >= 65 {
			break
		}
		if err := w.WriteData(o.RemoteEndpointAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteEndpointAddress); uint64(i1) < 65; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *RASUpdateConnection1IDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &ObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	o.RemoteEndpointAddress = make([]uint16, 65)
	for i1 := range o.RemoteEndpointAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteEndpointAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// RASUpdateConnectionIDL structure represents RAS_UPDATE_CONNECTION_IDL RPC union.
type RASUpdateConnectionIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *RASUpdateConnectionIDL_UpdateConnection1
	Value is_RASUpdateConnectionIDL `json:"value"`
}

func (o *RASUpdateConnectionIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RASUpdateConnectionIDL_UpdateConnection1:
		if value != nil {
			return value.UpdateConnection1
		}
	}
	return nil
}

type is_RASUpdateConnectionIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RASUpdateConnectionIDL()
}

func (o *RASUpdateConnectionIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint8(o.Revision)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		_o, _ := o.Value.(*RASUpdateConnectionIDL_UpdateConnection1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RASUpdateConnectionIDL_UpdateConnection1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *RASUpdateConnectionIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint8)(&o.Revision)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	switch o.Revision {
	case uint8(1):
		o.Value = &RASUpdateConnectionIDL_UpdateConnection1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// RASUpdateConnectionIDL_UpdateConnection1 structure represents RAS_UPDATE_CONNECTION_IDL RPC union arm.
//
// It has following labels: 1
type RASUpdateConnectionIDL_UpdateConnection1 struct {
	UpdateConnection1 *RASUpdateConnection1IDL `idl:"name:UpdateConnection1" json:"update_connection1"`
}

func (*RASUpdateConnectionIDL_UpdateConnection1) is_RASUpdateConnectionIDL() {}

func (o *RASUpdateConnectionIDL_UpdateConnection1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UpdateConnection1 != nil {
		if err := o.UpdateConnection1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASUpdateConnection1IDL{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RASUpdateConnectionIDL_UpdateConnection1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UpdateConnection1 == nil {
		o.UpdateConnection1 = &RASUpdateConnection1IDL{}
	}
	if err := o.UpdateConnection1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// InterfaceContainer structure represents DIM_INTERFACE_CONTAINER RPC structure.
type InterfaceContainer struct {
	GetInterfaceInfo  uint32 `idl:"name:fGetInterfaceInfo" json:"get_interface_info"`
	InterfaceInfoSize uint32 `idl:"name:dwInterfaceInfoSize" json:"interface_info_size"`
	InterfaceInfo     []byte `idl:"name:pInterfaceInfo;size_is:(dwInterfaceInfoSize)" json:"interface_info"`
	GetGlobalInfo     uint32 `idl:"name:fGetGlobalInfo" json:"get_global_info"`
	GlobalInfoSize    uint32 `idl:"name:dwGlobalInfoSize" json:"global_info_size"`
	GlobalInfo        []byte `idl:"name:pGlobalInfo;size_is:(dwGlobalInfoSize)" json:"global_info"`
}

func (o *InterfaceContainer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.InterfaceInfo != nil && o.InterfaceInfoSize == 0 {
		o.InterfaceInfoSize = uint32(len(o.InterfaceInfo))
	}
	if o.GlobalInfo != nil && o.GlobalInfoSize == 0 {
		o.GlobalInfoSize = uint32(len(o.GlobalInfo))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.GetInterfaceInfo); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceInfoSize); err != nil {
		return err
	}
	if o.InterfaceInfo != nil || o.InterfaceInfoSize > 0 {
		_ptr_pInterfaceInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfaceInfoSize)
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
				if err := w.WriteData(o.InterfaceInfo[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.InterfaceInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InterfaceInfo, _ptr_pInterfaceInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GetGlobalInfo); err != nil {
		return err
	}
	if err := w.WriteData(o.GlobalInfoSize); err != nil {
		return err
	}
	if o.GlobalInfo != nil || o.GlobalInfoSize > 0 {
		_ptr_pGlobalInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.GlobalInfoSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.GlobalInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.GlobalInfo[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.GlobalInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GlobalInfo, _ptr_pGlobalInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.GetInterfaceInfo); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceInfoSize); err != nil {
		return err
	}
	_ptr_pInterfaceInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfaceInfoSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfaceInfoSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceInfo", sizeInfo[0])
		}
		o.InterfaceInfo = make([]byte, sizeInfo[0])
		for i1 := range o.InterfaceInfo {
			i1 := i1
			if err := w.ReadData(&o.InterfaceInfo[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pInterfaceInfo := func(ptr interface{}) { o.InterfaceInfo = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.InterfaceInfo, _s_pInterfaceInfo, _ptr_pInterfaceInfo); err != nil {
		return err
	}
	if err := w.ReadData(&o.GetGlobalInfo); err != nil {
		return err
	}
	if err := w.ReadData(&o.GlobalInfoSize); err != nil {
		return err
	}
	_ptr_pGlobalInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.GlobalInfoSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.GlobalInfoSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.GlobalInfo", sizeInfo[0])
		}
		o.GlobalInfo = make([]byte, sizeInfo[0])
		for i1 := range o.GlobalInfo {
			i1 := i1
			if err := w.ReadData(&o.GlobalInfo[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pGlobalInfo := func(ptr interface{}) { o.GlobalInfo = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.GlobalInfo, _s_pGlobalInfo, _ptr_pGlobalInfo); err != nil {
		return err
	}
	return nil
}

// RouterTOCEntry structure represents RTR_TOC_ENTRY RPC structure.
type RouterTOCEntry struct {
	InfoType uint32 `idl:"name:InfoType" json:"info_type"`
	InfoSize uint32 `idl:"name:InfoSize" json:"info_size"`
	Count    uint32 `idl:"name:Count" json:"count"`
	Offset   uint32 `idl:"name:Offset" json:"offset"`
}

func (o *RouterTOCEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterTOCEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InfoType); err != nil {
		return err
	}
	if err := w.WriteData(o.InfoSize); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	return nil
}
func (o *RouterTOCEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InfoType); err != nil {
		return err
	}
	if err := w.ReadData(&o.InfoSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	return nil
}

// RouterInfoBlockHeader structure represents RTR_INFO_BLOCK_HEADER RPC structure.
type RouterInfoBlockHeader struct {
	Version         uint32            `idl:"name:Version" json:"version"`
	Size            uint32            `idl:"name:Size" json:"size"`
	TOCEntriesCount uint32            `idl:"name:TocEntriesCount" json:"toc_entries_count"`
	TOCEntry        []*RouterTOCEntry `idl:"name:TocEntry" json:"toc_entry"`
}

func (o *RouterInfoBlockHeader) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterInfoBlockHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.TOCEntriesCount); err != nil {
		return err
	}
	for i1 := range o.TOCEntry {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.TOCEntry[i1] != nil {
			if err := o.TOCEntry[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RouterTOCEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.TOCEntry); uint64(i1) < 1; i1++ {
		if err := (&RouterTOCEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RouterInfoBlockHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.TOCEntriesCount); err != nil {
		return err
	}
	o.TOCEntry = make([]*RouterTOCEntry, 1)
	for i1 := range o.TOCEntry {
		i1 := i1
		if o.TOCEntry[i1] == nil {
			o.TOCEntry[i1] = &RouterTOCEntry{}
		}
		if err := o.TOCEntry[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// FilterInfo structure represents FILTER_INFO RPC structure.
type FilterInfo struct {
	SourceAddr      uint32 `idl:"name:dwSrcAddr" json:"source_addr"`
	SourceMask      uint32 `idl:"name:dwSrcMask" json:"source_mask"`
	DestinationAddr uint32 `idl:"name:dwDstAddr" json:"destination_addr"`
	DestinationMask uint32 `idl:"name:dwDstMask" json:"destination_mask"`
	Protocol        uint32 `idl:"name:dwProtocol" json:"protocol"`
	LateBound       uint32 `idl:"name:fLateBound" json:"late_bound"`
	SourcePort      uint16 `idl:"name:wSrcPort" json:"source_port"`
	DestinationPort uint16 `idl:"name:wDstPort" json:"destination_port"`
}

func (o *FilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceMask); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationMask); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.LateBound); err != nil {
		return err
	}
	if err := w.WriteData(o.SourcePort); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationPort); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *FilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.LateBound); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourcePort); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationPort); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// FilterDescriptor structure represents FILTER_DESCRIPTOR RPC structure.
type FilterDescriptor struct {
	Version       uint32        `idl:"name:dwVersion" json:"version"`
	FiltersLength uint32        `idl:"name:dwNumFilters" json:"filters_length"`
	DefaultAction ForwardAction `idl:"name:faDefaultAction" json:"default_action"`
	Filter        []*FilterInfo `idl:"name:fiFilter" json:"filter"`
}

func (o *FilterDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FilterDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.FiltersLength); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.DefaultAction)); err != nil {
		return err
	}
	for i1 := range o.Filter {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Filter[i1] != nil {
			if err := o.Filter[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FilterInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Filter); uint64(i1) < 1; i1++ {
		if err := (&FilterInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.FiltersLength); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.DefaultAction)); err != nil {
		return err
	}
	o.Filter = make([]*FilterInfo, 1)
	for i1 := range o.Filter {
		i1 := i1
		if o.Filter[i1] == nil {
			o.Filter[i1] = &FilterInfo{}
		}
		if err := o.Filter[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// FilterInfoV6 structure represents FILTER_INFO_V6 RPC structure.
type FilterInfoV6 struct {
	IPv6SourceAddr          []byte `idl:"name:ipv6SrcAddr" json:"ipv6_source_addr"`
	SourcePrefixLength      uint32 `idl:"name:dwSrcPrefixLength" json:"source_prefix_length"`
	IPv6DestinationAddr     []byte `idl:"name:ipv6DstAddr" json:"ipv6_destination_addr"`
	DestinationPrefixLength uint32 `idl:"name:dwDstPrefixLength" json:"destination_prefix_length"`
	Protocol                uint32 `idl:"name:dwProtocol" json:"protocol"`
	LateBound               uint32 `idl:"name:fLateBound" json:"late_bound"`
	SourcePort              uint16 `idl:"name:wSrcPort" json:"source_port"`
	DestinationPort         uint16 `idl:"name:wDstPort" json:"destination_port"`
}

func (o *FilterInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FilterInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.IPv6SourceAddr {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.IPv6SourceAddr[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.IPv6SourceAddr); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SourcePrefixLength); err != nil {
		return err
	}
	for i1 := range o.IPv6DestinationAddr {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.IPv6DestinationAddr[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.IPv6DestinationAddr); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DestinationPrefixLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.LateBound); err != nil {
		return err
	}
	if err := w.WriteData(o.SourcePort); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationPort); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *FilterInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.IPv6SourceAddr = make([]byte, 16)
	for i1 := range o.IPv6SourceAddr {
		i1 := i1
		if err := w.ReadData(&o.IPv6SourceAddr[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.SourcePrefixLength); err != nil {
		return err
	}
	o.IPv6DestinationAddr = make([]byte, 16)
	for i1 := range o.IPv6DestinationAddr {
		i1 := i1
		if err := w.ReadData(&o.IPv6DestinationAddr[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.DestinationPrefixLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.LateBound); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourcePort); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationPort); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// FilterDescriptorV6 structure represents FILTER_DESCRIPTOR_V6 RPC structure.
type FilterDescriptorV6 struct {
	Version       uint32          `idl:"name:dwVersion" json:"version"`
	FiltersLength uint32          `idl:"name:dwNumFilters" json:"filters_length"`
	DefaultAction ForwardAction   `idl:"name:faDefaultAction" json:"default_action"`
	Filter        []*FilterInfoV6 `idl:"name:fiFilter" json:"filter"`
}

func (o *FilterDescriptorV6) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FilterDescriptorV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.FiltersLength); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.DefaultAction)); err != nil {
		return err
	}
	for i1 := range o.Filter {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Filter[i1] != nil {
			if err := o.Filter[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FilterInfoV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Filter); uint64(i1) < 1; i1++ {
		if err := (&FilterInfoV6{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterDescriptorV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.FiltersLength); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.DefaultAction)); err != nil {
		return err
	}
	o.Filter = make([]*FilterInfoV6, 1)
	for i1 := range o.Filter {
		i1 := i1
		if o.Filter[i1] == nil {
			o.Filter[i1] = &FilterInfoV6{}
		}
		if err := o.Filter[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// GlobalInfo structure represents GLOBAL_INFO RPC structure.
type GlobalInfo struct {
	FilteringOn  bool   `idl:"name:bFilteringOn" json:"filtering_on"`
	LoggingLevel uint32 `idl:"name:dwLoggingLevel" json:"logging_level"`
}

func (o *GlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *GlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.FilteringOn {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	return nil
}
func (o *GlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bFilteringOn int32
	if err := w.ReadData(&_bFilteringOn); err != nil {
		return err
	}
	o.FilteringOn = _bFilteringOn != 0
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	return nil
}

// InterfaceRouteInfo structure represents INTERFACE_ROUTE_INFO RPC structure.
type InterfaceRouteInfo struct {
	Field1               *InterfaceRouteInfo_Field1 `idl:"name:" json:""`
	RTInfoInterfaceIndex uint32                     `idl:"name:dwRtInfoIfIndex" json:"rt_info_interface_index"`
	RTInfoType           uint32                     `idl:"name:dwRtInfoType" json:"rt_info_type"`
	RTInfoProto          uint32                     `idl:"name:dwRtInfoProto" json:"rt_info_proto"`
	RTInfoPreference     uint32                     `idl:"name:dwRtInfoPreference" json:"rt_info_preference"`
	RTInfoViewSet        uint32                     `idl:"name:dwRtInfoViewSet" json:"rt_info_view_set"`
	V4                   bool                       `idl:"name:bV4" json:"v4"`
}

func (o *InterfaceRouteInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceRouteInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// FIXME unknown type
	if err := w.WriteData(o.RTInfoInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoType); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoProto); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoPreference); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoViewSet); err != nil {
		return err
	}
	if !o.V4 {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceRouteInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// FIXME: unknown type
	if err := w.ReadData(&o.RTInfoInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoType); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoProto); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoPreference); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoViewSet); err != nil {
		return err
	}
	var _bV4 int32
	if err := w.ReadData(&_bV4); err != nil {
		return err
	}
	o.V4 = _bV4 != 0
	return nil
}

type InterfaceRouteInfo_Field1 struct {
	Field1 *InterfaceRouteInfo_Field1_Field1 `idl:"name:" json:""`
	Field2 *InterfaceRouteInfo_Field1_Field2 `idl:"name:" json:""`
}

// InterfaceRouteInfo_Field1_Field1 structure represents INTERFACE_ROUTE_INFO structure anonymous member.
type InterfaceRouteInfo_Field1_Field1 struct {
	RTInfoDestination uint32 `idl:"name:dwRtInfoDest" json:"rt_info_destination"`
	RTInfoMask        uint32 `idl:"name:dwRtInfoMask" json:"rt_info_mask"`
	RTInfoPolicy      uint32 `idl:"name:dwRtInfoPolicy" json:"rt_info_policy"`
	RTInfoNextHop     uint32 `idl:"name:dwRtInfoNextHop" json:"rt_info_next_hop"`
	RTInfoAge         uint32 `idl:"name:dwRtInfoAge" json:"rt_info_age"`
	RTInfoNextHopAs   uint32 `idl:"name:dwRtInfoNextHopAS" json:"rt_info_next_hop_as"`
	RTInfoMetric1     uint32 `idl:"name:dwRtInfoMetric1" json:"rt_info_metric1"`
	RTInfoMetric2     uint32 `idl:"name:dwRtInfoMetric2" json:"rt_info_metric2"`
	RTInfoMetric3     uint32 `idl:"name:dwRtInfoMetric3" json:"rt_info_metric3"`
}

func (o *InterfaceRouteInfo_Field1_Field1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceRouteInfo_Field1_Field1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoDestination); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoMask); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoPolicy); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoNextHop); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoAge); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoNextHopAs); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoMetric1); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoMetric2); err != nil {
		return err
	}
	if err := w.WriteData(o.RTInfoMetric3); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceRouteInfo_Field1_Field1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoDestination); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoPolicy); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoNextHop); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoAge); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoNextHopAs); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoMetric1); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoMetric2); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTInfoMetric3); err != nil {
		return err
	}
	return nil
}

// InterfaceRouteInfo_Field1_Field2 structure represents INTERFACE_ROUTE_INFO structure anonymous member.
type InterfaceRouteInfo_Field1_Field2 struct {
	DestinationPrefix       *IN6Addr `idl:"name:DestinationPrefix" json:"destination_prefix"`
	DestinationPrefixLength uint32   `idl:"name:DestPrefixLength" json:"destination_prefix_length"`
	NextHopAddress          *IN6Addr `idl:"name:NextHopAddress" json:"next_hop_address"`
	ValidLifeTime           uint32   `idl:"name:ValidLifeTime" json:"valid_life_time"`
	Flags                   uint32   `idl:"name:Flags" json:"flags"`
	Metric                  uint32   `idl:"name:Metric" json:"metric"`
}

func (o *InterfaceRouteInfo_Field1_Field2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceRouteInfo_Field1_Field2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.DestinationPrefix != nil {
		if err := o.DestinationPrefix.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IN6Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DestinationPrefixLength); err != nil {
		return err
	}
	if o.NextHopAddress != nil {
		if err := o.NextHopAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IN6Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ValidLifeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Metric); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceRouteInfo_Field1_Field2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.DestinationPrefix == nil {
		o.DestinationPrefix = &IN6Addr{}
	}
	if err := o.DestinationPrefix.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationPrefixLength); err != nil {
		return err
	}
	if o.NextHopAddress == nil {
		o.NextHopAddress = &IN6Addr{}
	}
	if err := o.NextHopAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValidLifeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Metric); err != nil {
		return err
	}
	return nil
}

// ProtocolMetric structure represents PROTOCOL_METRIC RPC structure.
type ProtocolMetric struct {
	ProtocolID uint32 `idl:"name:dwProtocolId" json:"protocol_id"`
	Metric     uint32 `idl:"name:dwMetric" json:"metric"`
}

func (o *ProtocolMetric) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ProtocolMetric) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ProtocolID); err != nil {
		return err
	}
	if err := w.WriteData(o.Metric); err != nil {
		return err
	}
	return nil
}
func (o *ProtocolMetric) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtocolID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Metric); err != nil {
		return err
	}
	return nil
}

// PriorityInfo structure represents PRIORITY_INFO RPC structure.
type PriorityInfo struct {
	ProtocolsLength uint32            `idl:"name:dwNumProtocols" json:"protocols_length"`
	ProtocolMetric  []*ProtocolMetric `idl:"name:ppmProtocolMetric" json:"protocol_metric"`
}

func (o *PriorityInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PriorityInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ProtocolsLength); err != nil {
		return err
	}
	for i1 := range o.ProtocolMetric {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.ProtocolMetric[i1] != nil {
			if err := o.ProtocolMetric[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProtocolMetric{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ProtocolMetric); uint64(i1) < 1; i1++ {
		if err := (&ProtocolMetric{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PriorityInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtocolsLength); err != nil {
		return err
	}
	o.ProtocolMetric = make([]*ProtocolMetric, 1)
	for i1 := range o.ProtocolMetric {
		i1 := i1
		if o.ProtocolMetric[i1] == nil {
			o.ProtocolMetric[i1] = &ProtocolMetric{}
		}
		if err := o.ProtocolMetric[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// ProtocolMetricEx structure represents PROTOCOL_METRIC_EX RPC structure.
type ProtocolMetricEx struct {
	ProtocolID    uint32 `idl:"name:dwProtocolId" json:"protocol_id"`
	SubProtocolID uint32 `idl:"name:dwSubProtocolId" json:"sub_protocol_id"`
	Metric        uint32 `idl:"name:dwMetric" json:"metric"`
}

func (o *ProtocolMetricEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ProtocolMetricEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ProtocolID); err != nil {
		return err
	}
	if err := w.WriteData(o.SubProtocolID); err != nil {
		return err
	}
	if err := w.WriteData(o.Metric); err != nil {
		return err
	}
	return nil
}
func (o *ProtocolMetricEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtocolID); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubProtocolID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Metric); err != nil {
		return err
	}
	return nil
}

// PriorityInfoEx structure represents PRIORITY_INFO_EX RPC structure.
type PriorityInfoEx struct {
	ProtocolsLength uint32              `idl:"name:dwNumProtocols" json:"protocols_length"`
	ProtocolMetric  []*ProtocolMetricEx `idl:"name:ppmProtocolMetric" json:"protocol_metric"`
}

func (o *PriorityInfoEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PriorityInfoEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ProtocolsLength); err != nil {
		return err
	}
	for i1 := range o.ProtocolMetric {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.ProtocolMetric[i1] != nil {
			if err := o.ProtocolMetric[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProtocolMetricEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ProtocolMetric); uint64(i1) < 1; i1++ {
		if err := (&ProtocolMetricEx{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PriorityInfoEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtocolsLength); err != nil {
		return err
	}
	o.ProtocolMetric = make([]*ProtocolMetricEx, 1)
	for i1 := range o.ProtocolMetric {
		i1 := i1
		if o.ProtocolMetric[i1] == nil {
			o.ProtocolMetric[i1] = &ProtocolMetricEx{}
		}
		if err := o.ProtocolMetric[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// RouterDiscInfo structure represents RTR_DISC_INFO RPC structure.
type RouterDiscInfo struct {
	MaxAdvertiseInterval uint16 `idl:"name:wMaxAdvtInterval" json:"max_advertise_interval"`
	MinAdvertiseInterval uint16 `idl:"name:wMinAdvtInterval" json:"min_advertise_interval"`
	AdvertiseLifetime    uint16 `idl:"name:wAdvtLifetime" json:"advertise_lifetime"`
	Advertise            bool   `idl:"name:bAdvertise" json:"advertise"`
	PrefLevel            int32  `idl:"name:lPrefLevel" json:"pref_level"`
}

func (o *RouterDiscInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterDiscInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxAdvertiseInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.MinAdvertiseInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AdvertiseLifetime); err != nil {
		return err
	}
	if !o.Advertise {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrefLevel); err != nil {
		return err
	}
	return nil
}
func (o *RouterDiscInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxAdvertiseInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinAdvertiseInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdvertiseLifetime); err != nil {
		return err
	}
	var _bAdvertise int32
	if err := w.ReadData(&_bAdvertise); err != nil {
		return err
	}
	o.Advertise = _bAdvertise != 0
	if err := w.ReadData(&o.PrefLevel); err != nil {
		return err
	}
	return nil
}

// MulticastHBeatInfo structure represents MCAST_HBEAT_INFO RPC structure.
type MulticastHBeatInfo struct {
	Group        []uint16 `idl:"name:pwszGroup" json:"group"`
	Active       bool     `idl:"name:bActive" json:"active"`
	DeadInterval uint32   `idl:"name:ulDeadInterval" json:"dead_interval"`
	ByProtocol   uint8    `idl:"name:byProtocol" json:"by_protocol"`
	Port         uint16   `idl:"name:wPort" json:"port"`
}

func (o *MulticastHBeatInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MulticastHBeatInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.Group {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.Group[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Group); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if !o.Active {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DeadInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.ByProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MulticastHBeatInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.Group = make([]uint16, 64)
	for i1 := range o.Group {
		i1 := i1
		if err := w.ReadData(&o.Group[i1]); err != nil {
			return err
		}
	}
	var _bActive int32
	if err := w.ReadData(&_bActive); err != nil {
		return err
	}
	o.Active = _bActive != 0
	if err := w.ReadData(&o.DeadInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.ByProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MIBMulticastLimitRow structure represents MIB_MCAST_LIMIT_ROW RPC structure.
type MIBMulticastLimitRow struct {
	TTL       uint32 `idl:"name:dwTtl" json:"ttl"`
	RateLimit uint32 `idl:"name:dwRateLimit" json:"rate_limit"`
}

func (o *MIBMulticastLimitRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBMulticastLimitRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TTL); err != nil {
		return err
	}
	if err := w.WriteData(o.RateLimit); err != nil {
		return err
	}
	return nil
}
func (o *MIBMulticastLimitRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.RateLimit); err != nil {
		return err
	}
	return nil
}

// IPInIPConfigInfo structure represents IPINIP_CONFIG_INFO RPC structure.
type IPInIPConfigInfo struct {
	RemoteAddress uint32 `idl:"name:dwRemoteAddress" json:"remote_address"`
	LocalAddress  uint32 `idl:"name:dwLocalAddress" json:"local_address"`
	ByTTL         uint8  `idl:"name:byTtl" json:"by_ttl"`
}

func (o *IPInIPConfigInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPInIPConfigInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.ByTTL); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPInIPConfigInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.ByTTL); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// InterfaceStatusInfo structure represents INTERFACE_STATUS_INFO RPC structure.
type InterfaceStatusInfo struct {
	AdminStatus uint32 `idl:"name:dwAdminStatus" json:"admin_status"`
}

func (o *InterfaceStatusInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceStatusInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminStatus); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceStatusInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminStatus); err != nil {
		return err
	}
	return nil
}

// MIBEntryContainer structure represents DIM_MIB_ENTRY_CONTAINER RPC structure.
type MIBEntryContainer struct {
	MIBInEntrySize  uint32 `idl:"name:dwMibInEntrySize" json:"mib_in_entry_size"`
	MIBInEntry      []byte `idl:"name:pMibInEntry;size_is:(dwMibInEntrySize)" json:"mib_in_entry"`
	MIBOutEntrySize uint32 `idl:"name:dwMibOutEntrySize" json:"mib_out_entry_size"`
	MIBOutEntry     []byte `idl:"name:pMibOutEntry;size_is:(dwMibOutEntrySize)" json:"mib_out_entry"`
}

func (o *MIBEntryContainer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.MIBInEntry != nil && o.MIBInEntrySize == 0 {
		o.MIBInEntrySize = uint32(len(o.MIBInEntry))
	}
	if o.MIBOutEntry != nil && o.MIBOutEntrySize == 0 {
		o.MIBOutEntrySize = uint32(len(o.MIBOutEntry))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBEntryContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MIBInEntrySize); err != nil {
		return err
	}
	if o.MIBInEntry != nil || o.MIBInEntrySize > 0 {
		_ptr_pMibInEntry := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MIBInEntrySize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.MIBInEntry {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.MIBInEntry[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.MIBInEntry); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MIBInEntry, _ptr_pMibInEntry); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MIBOutEntrySize); err != nil {
		return err
	}
	if o.MIBOutEntry != nil || o.MIBOutEntrySize > 0 {
		_ptr_pMibOutEntry := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MIBOutEntrySize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.MIBOutEntry {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.MIBOutEntry[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.MIBOutEntry); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MIBOutEntry, _ptr_pMibOutEntry); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBEntryContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MIBInEntrySize); err != nil {
		return err
	}
	_ptr_pMibInEntry := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MIBInEntrySize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MIBInEntrySize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.MIBInEntry", sizeInfo[0])
		}
		o.MIBInEntry = make([]byte, sizeInfo[0])
		for i1 := range o.MIBInEntry {
			i1 := i1
			if err := w.ReadData(&o.MIBInEntry[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMibInEntry := func(ptr interface{}) { o.MIBInEntry = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.MIBInEntry, _s_pMibInEntry, _ptr_pMibInEntry); err != nil {
		return err
	}
	if err := w.ReadData(&o.MIBOutEntrySize); err != nil {
		return err
	}
	_ptr_pMibOutEntry := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MIBOutEntrySize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MIBOutEntrySize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.MIBOutEntry", sizeInfo[0])
		}
		o.MIBOutEntry = make([]byte, sizeInfo[0])
		for i1 := range o.MIBOutEntry {
			i1 := i1
			if err := w.ReadData(&o.MIBOutEntry[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMibOutEntry := func(ptr interface{}) { o.MIBOutEntry = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.MIBOutEntry, _s_pMibOutEntry, _ptr_pMibOutEntry); err != nil {
		return err
	}
	return nil
}

// MIBIPForwardRow structure represents MIB_IPFORWARDROW RPC structure.
type MIBIPForwardRow struct {
	ForwardDestination    uint32            `idl:"name:dwForwardDest" json:"forward_destination"`
	ForwardMask           uint32            `idl:"name:dwForwardMask" json:"forward_mask"`
	ForwardPolicy         uint32            `idl:"name:dwForwardPolicy" json:"forward_policy"`
	ForwardNextHop        uint32            `idl:"name:dwForwardNextHop" json:"forward_next_hop"`
	ForwardInterfaceIndex uint32            `idl:"name:dwForwardIfIndex" json:"forward_interface_index"`
	ForwardType           MIBIPForwardType  `idl:"name:ForwardType" json:"forward_type"`
	ForwardProto          MIBIPForwardProto `idl:"name:ForwardProto" json:"forward_proto"`
	ForwardAge            uint32            `idl:"name:dwForwardAge" json:"forward_age"`
	ForwardNextHopAs      uint32            `idl:"name:dwForwardNextHopAS" json:"forward_next_hop_as"`
	ForwardMetric1        uint32            `idl:"name:dwForwardMetric1" json:"forward_metric1"`
	ForwardMetric2        uint32            `idl:"name:dwForwardMetric2" json:"forward_metric2"`
	ForwardMetric3        uint32            `idl:"name:dwForwardMetric3" json:"forward_metric3"`
	ForwardMetric4        uint32            `idl:"name:dwForwardMetric4" json:"forward_metric4"`
	ForwardMetric5        uint32            `idl:"name:dwForwardMetric5" json:"forward_metric5"`
}

func (o *MIBIPForwardRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPForwardRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardDestination); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardMask); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardPolicy); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardNextHop); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ForwardType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ForwardProto)); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardAge); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardNextHopAs); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardMetric1); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardMetric2); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardMetric3); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardMetric4); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardMetric5); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPForwardRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardDestination); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardPolicy); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardNextHop); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ForwardType)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ForwardProto)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardAge); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardNextHopAs); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardMetric1); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardMetric2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardMetric3); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardMetric4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardMetric5); err != nil {
		return err
	}
	return nil
}

// MIBIPDestinationRow structure represents MIB_IPDESTROW RPC structure.
type MIBIPDestinationRow struct {
	ForwardRow        *MIBIPForwardRow `idl:"name:ForwardRow" json:"forward_row"`
	ForwardPreference uint32           `idl:"name:dwForwardPreference" json:"forward_preference"`
	ForwardViewSet    uint32           `idl:"name:dwForwardViewSet" json:"forward_view_set"`
}

func (o *MIBIPDestinationRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPDestinationRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ForwardRow != nil {
		if err := o.ForwardRow.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MIBIPForwardRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ForwardPreference); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwardViewSet); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPDestinationRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ForwardRow == nil {
		o.ForwardRow = &MIBIPForwardRow{}
	}
	if err := o.ForwardRow.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardPreference); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwardViewSet); err != nil {
		return err
	}
	return nil
}

// MIBIPDestinationTable structure represents MIB_IPDESTTABLE RPC structure.
type MIBIPDestinationTable struct {
	EntriesLength uint32                 `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPDestinationRow `idl:"name:table" json:"table"`
}

func (o *MIBIPDestinationTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPDestinationTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPDestinationRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPDestinationRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIPDestinationTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPDestinationRow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPDestinationRow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBRouteState structure represents MIB_ROUTESTATE RPC structure.
type MIBRouteState struct {
	RoutesSetToStack bool `idl:"name:bRoutesSetToStack" json:"routes_set_to_stack"`
}

func (o *MIBRouteState) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBRouteState) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.RoutesSetToStack {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBRouteState) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bRoutesSetToStack int32
	if err := w.ReadData(&_bRoutesSetToStack); err != nil {
		return err
	}
	o.RoutesSetToStack = _bRoutesSetToStack != 0
	return nil
}

// MIBBestInterface structure represents MIB_BEST_IF RPC structure.
type MIBBestInterface struct {
	DestinationAddr uint32 `idl:"name:dwDestAddr" json:"destination_addr"`
	InterfaceIndex  uint32 `idl:"name:dwIfIndex" json:"interface_index"`
}

func (o *MIBBestInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBBestInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *MIBBestInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}

// MIBBoundaryrow structure represents MIB_BOUNDARYROW RPC structure.
type MIBBoundaryrow struct {
	GroupAddress uint32 `idl:"name:dwGroupAddress" json:"group_address"`
	GroupMask    uint32 `idl:"name:dwGroupMask" json:"group_mask"`
}

func (o *MIBBoundaryrow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBBoundaryrow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupMask); err != nil {
		return err
	}
	return nil
}
func (o *MIBBoundaryrow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupMask); err != nil {
		return err
	}
	return nil
}

// MIBICMPStats structure represents MIBICMPSTATS RPC structure.
type MIBICMPStats struct {
	Msgs                uint32 `idl:"name:dwMsgs" json:"msgs"`
	Errors              uint32 `idl:"name:dwErrors" json:"errors"`
	DestinationUnreachs uint32 `idl:"name:dwDestUnreachs" json:"destination_unreachs"`
	TimeExcds           uint32 `idl:"name:dwTimeExcds" json:"time_excds"`
	ParameterProbs      uint32 `idl:"name:dwParmProbs" json:"parameter_probs"`
	SourceQuenchs       uint32 `idl:"name:dwSrcQuenchs" json:"source_quenchs"`
	Redirects           uint32 `idl:"name:dwRedirects" json:"redirects"`
	Echos               uint32 `idl:"name:dwEchos" json:"echos"`
	EchoReps            uint32 `idl:"name:dwEchoReps" json:"echo_reps"`
	Timestamps          uint32 `idl:"name:dwTimestamps" json:"timestamps"`
	TimestampReps       uint32 `idl:"name:dwTimestampReps" json:"timestamp_reps"`
	AddrMasks           uint32 `idl:"name:dwAddrMasks" json:"addr_masks"`
	AddrMaskReps        uint32 `idl:"name:dwAddrMaskReps" json:"addr_mask_reps"`
}

func (o *MIBICMPStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBICMPStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Msgs); err != nil {
		return err
	}
	if err := w.WriteData(o.Errors); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationUnreachs); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeExcds); err != nil {
		return err
	}
	if err := w.WriteData(o.ParameterProbs); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceQuenchs); err != nil {
		return err
	}
	if err := w.WriteData(o.Redirects); err != nil {
		return err
	}
	if err := w.WriteData(o.Echos); err != nil {
		return err
	}
	if err := w.WriteData(o.EchoReps); err != nil {
		return err
	}
	if err := w.WriteData(o.Timestamps); err != nil {
		return err
	}
	if err := w.WriteData(o.TimestampReps); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrMasks); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrMaskReps); err != nil {
		return err
	}
	return nil
}
func (o *MIBICMPStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Msgs); err != nil {
		return err
	}
	if err := w.ReadData(&o.Errors); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationUnreachs); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeExcds); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParameterProbs); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceQuenchs); err != nil {
		return err
	}
	if err := w.ReadData(&o.Redirects); err != nil {
		return err
	}
	if err := w.ReadData(&o.Echos); err != nil {
		return err
	}
	if err := w.ReadData(&o.EchoReps); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timestamps); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimestampReps); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrMasks); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrMaskReps); err != nil {
		return err
	}
	return nil
}

// MIBICMPInfo structure represents MIBICMPINFO RPC structure.
type MIBICMPInfo struct {
	ICMPInStats  *MIBICMPStats `idl:"name:icmpInStats" json:"icmp_in_stats"`
	ICMPOutStats *MIBICMPStats `idl:"name:icmpOutStats" json:"icmp_out_stats"`
}

func (o *MIBICMPInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBICMPInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ICMPInStats != nil {
		if err := o.ICMPInStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MIBICMPStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ICMPOutStats != nil {
		if err := o.ICMPOutStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MIBICMPStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBICMPInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ICMPInStats == nil {
		o.ICMPInStats = &MIBICMPStats{}
	}
	if err := o.ICMPInStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ICMPOutStats == nil {
		o.ICMPOutStats = &MIBICMPStats{}
	}
	if err := o.ICMPOutStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MIBICMP structure represents MIB_ICMP RPC structure.
type MIBICMP struct {
	Stats *MIBICMPInfo `idl:"name:stats" json:"stats"`
}

func (o *MIBICMP) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBICMP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Stats != nil {
		if err := o.Stats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MIBICMPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBICMP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Stats == nil {
		o.Stats = &MIBICMPInfo{}
	}
	if err := o.Stats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MIBInterfaceNumber structure represents MIB_IFNUMBER RPC structure.
type MIBInterfaceNumber struct {
	Value uint32 `idl:"name:dwValue" json:"value"`
}

func (o *MIBInterfaceNumber) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBInterfaceNumber) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Value); err != nil {
		return err
	}
	return nil
}
func (o *MIBInterfaceNumber) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Value); err != nil {
		return err
	}
	return nil
}

// MIBInterfaceRow structure represents MIB_IFROW RPC structure.
type MIBInterfaceRow struct {
	Name               []uint16 `idl:"name:wszName" json:"name"`
	Index              uint32   `idl:"name:dwIndex" json:"index"`
	Type               uint32   `idl:"name:dwType" json:"type"`
	MTU                uint32   `idl:"name:dwMtu" json:"mtu"`
	Speed              uint32   `idl:"name:dwSpeed" json:"speed"`
	PhysicalAddrLength uint32   `idl:"name:dwPhysAddrLen" json:"physical_addr_length"`
	PhysicalAddr       []byte   `idl:"name:bPhysAddr" json:"physical_addr"`
	AdminStatus        uint32   `idl:"name:dwAdminStatus" json:"admin_status"`
	OperatorStatus     uint32   `idl:"name:dwOperStatus" json:"operator_status"`
	LastChange         uint32   `idl:"name:dwLastChange" json:"last_change"`
	InOctets           uint32   `idl:"name:dwInOctets" json:"in_octets"`
	InUcastPackets     uint32   `idl:"name:dwInUcastPkts" json:"in_ucast_packets"`
	InNUcastPackets    uint32   `idl:"name:dwInNUcastPkts" json:"in_n_ucast_packets"`
	InDiscards         uint32   `idl:"name:dwInDiscards" json:"in_discards"`
	InErrors           uint32   `idl:"name:dwInErrors" json:"in_errors"`
	InUnknownProtos    uint32   `idl:"name:dwInUnknownProtos" json:"in_unknown_protos"`
	OutOctets          uint32   `idl:"name:dwOutOctets" json:"out_octets"`
	OutUcastPackets    uint32   `idl:"name:dwOutUcastPkts" json:"out_ucast_packets"`
	OutNUcastPackets   uint32   `idl:"name:dwOutNUcastPkts" json:"out_n_ucast_packets"`
	OutDiscards        uint32   `idl:"name:dwOutDiscards" json:"out_discards"`
	OutErrors          uint32   `idl:"name:dwOutErrors" json:"out_errors"`
	OutQLength         uint32   `idl:"name:dwOutQLen" json:"out_q_length"`
	DescriptionLength  uint32   `idl:"name:dwDescrLen" json:"description_length"`
	Description        []byte   `idl:"name:bDescr" json:"description"`
}

func (o *MIBInterfaceRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBInterfaceRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.MTU); err != nil {
		return err
	}
	if err := w.WriteData(o.Speed); err != nil {
		return err
	}
	if err := w.WriteData(o.PhysicalAddrLength); err != nil {
		return err
	}
	for i1 := range o.PhysicalAddr {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.PhysicalAddr[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PhysicalAddr); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AdminStatus); err != nil {
		return err
	}
	if err := w.WriteData(o.OperatorStatus); err != nil {
		return err
	}
	if err := w.WriteData(o.LastChange); err != nil {
		return err
	}
	if err := w.WriteData(o.InOctets); err != nil {
		return err
	}
	if err := w.WriteData(o.InUcastPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.InNUcastPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.InDiscards); err != nil {
		return err
	}
	if err := w.WriteData(o.InErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.InUnknownProtos); err != nil {
		return err
	}
	if err := w.WriteData(o.OutOctets); err != nil {
		return err
	}
	if err := w.WriteData(o.OutUcastPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.OutNUcastPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.OutDiscards); err != nil {
		return err
	}
	if err := w.WriteData(o.OutErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.OutQLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DescriptionLength); err != nil {
		return err
	}
	for i1 := range o.Description {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.Description[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Description); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MIBInterfaceRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.Name = make([]uint16, 256)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.MTU); err != nil {
		return err
	}
	if err := w.ReadData(&o.Speed); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhysicalAddrLength); err != nil {
		return err
	}
	o.PhysicalAddr = make([]byte, 8)
	for i1 := range o.PhysicalAddr {
		i1 := i1
		if err := w.ReadData(&o.PhysicalAddr[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AdminStatus); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperatorStatus); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastChange); err != nil {
		return err
	}
	if err := w.ReadData(&o.InOctets); err != nil {
		return err
	}
	if err := w.ReadData(&o.InUcastPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.InNUcastPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.InDiscards); err != nil {
		return err
	}
	if err := w.ReadData(&o.InErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.InUnknownProtos); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutOctets); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutUcastPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutNUcastPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutDiscards); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutQLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DescriptionLength); err != nil {
		return err
	}
	o.Description = make([]byte, 256)
	for i1 := range o.Description {
		i1 := i1
		if err := w.ReadData(&o.Description[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MIBInterfaceStatus structure represents MIB_IFSTATUS RPC structure.
type MIBInterfaceStatus struct {
	InterfaceIndex    uint32 `idl:"name:dwIfIndex" json:"interface_index"`
	AdminStatus       uint32 `idl:"name:dwAdminStatus" json:"admin_status"`
	OperationalStatus uint32 `idl:"name:dwOperationalStatus" json:"operational_status"`
	MHBeatActive      bool   `idl:"name:bMHbeatActive" json:"m_hbeat_active"`
	MHBeatAlive       bool   `idl:"name:bMHbeatAlive" json:"m_hbeat_alive"`
}

func (o *MIBInterfaceStatus) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBInterfaceStatus) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminStatus); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationalStatus); err != nil {
		return err
	}
	if !o.MHBeatActive {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.MHBeatAlive {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInterfaceStatus) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminStatus); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationalStatus); err != nil {
		return err
	}
	var _bMHBeatActive int32
	if err := w.ReadData(&_bMHBeatActive); err != nil {
		return err
	}
	o.MHBeatActive = _bMHBeatActive != 0
	var _bMHBeatAlive int32
	if err := w.ReadData(&_bMHBeatAlive); err != nil {
		return err
	}
	o.MHBeatAlive = _bMHBeatAlive != 0
	return nil
}

// MIBInterfaceTable structure represents MIB_IFTABLE RPC structure.
type MIBInterfaceTable struct {
	EntriesLength uint32             `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBInterfaceRow `idl:"name:table" json:"table"`
}

func (o *MIBInterfaceTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBInterfaceTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBInterfaceRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBInterfaceRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInterfaceTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBInterfaceRow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBInterfaceRow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIPAddrRow structure represents MIB_IPADDRROW RPC structure.
type MIBIPAddrRow struct {
	Addr      uint32 `idl:"name:dwAddr" json:"addr"`
	Index     uint32 `idl:"name:dwIndex" json:"index"`
	Mask      uint32 `idl:"name:dwMask" json:"mask"`
	BCastAddr uint32 `idl:"name:dwBCastAddr" json:"b_cast_addr"`
	ReasmSize uint32 `idl:"name:dwReasmSize" json:"reasm_size"`
	_         uint16 `idl:"name:unused1"`
	Type      uint16 `idl:"name:wType" json:"type"`
}

func (o *MIBIPAddrRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPAddrRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Addr); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.BCastAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.ReasmSize); err != nil {
		return err
	}
	// reserved unused1
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPAddrRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Addr); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.BCastAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReasmSize); err != nil {
		return err
	}
	// reserved unused1
	var _unused1 uint16
	if err := w.ReadData(&_unused1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MIBIPAddrTable structure represents MIB_IPADDRTABLE RPC structure.
type MIBIPAddrTable struct {
	EntriesLength uint32          `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPAddrRow `idl:"name:table" json:"table"`
}

func (o *MIBIPAddrTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPAddrTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPAddrRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPAddrRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIPAddrTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPAddrRow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPAddrRow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIPForwardNumber structure represents MIB_IPFORWARDNUMBER RPC structure.
type MIBIPForwardNumber struct {
	Value uint32 `idl:"name:dwValue" json:"value"`
}

func (o *MIBIPForwardNumber) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPForwardNumber) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Value); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPForwardNumber) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Value); err != nil {
		return err
	}
	return nil
}

// MIBIPForwardTable structure represents MIB_IPFORWARDTABLE RPC structure.
type MIBIPForwardTable struct {
	EntriesLength uint32             `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPForwardRow `idl:"name:table" json:"table"`
	_             []byte             `idl:"name:reserved"`
}

func (o *MIBIPForwardTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPForwardTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPForwardRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPForwardRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	for i1 := 0; uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPForwardTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPForwardRow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPForwardRow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	var _reserved []byte
	_reserved = make([]byte, 8)
	for i1 := range _reserved {
		i1 := i1
		if err := w.ReadData(&_reserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MIBIPMulticastBoundary structure represents MIB_IPMCAST_BOUNDARY RPC structure.
type MIBIPMulticastBoundary struct {
	InterfaceIndex uint32 `idl:"name:dwIfIndex" json:"interface_index"`
	GroupAddress   uint32 `idl:"name:dwGroupAddress" json:"group_address"`
	GroupMask      uint32 `idl:"name:dwGroupMask" json:"group_mask"`
	Status         uint32 `idl:"name:dwStatus" json:"status"`
}

func (o *MIBIPMulticastBoundary) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastBoundary) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupMask); err != nil {
		return err
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastBoundary) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	return nil
}

// MIBIPMulticastBoundaryTable structure represents MIB_IPMCAST_BOUNDARY_TABLE RPC structure.
type MIBIPMulticastBoundaryTable struct {
	EntriesLength uint32                    `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPMulticastBoundary `idl:"name:table" json:"table"`
}

func (o *MIBIPMulticastBoundaryTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastBoundaryTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPMulticastBoundary{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPMulticastBoundary{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIPMulticastBoundaryTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPMulticastBoundary, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPMulticastBoundary{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIPMulticastGlobal structure represents MIB_IPMCAST_GLOBAL RPC structure.
type MIBIPMulticastGlobal struct {
	Enable uint32 `idl:"name:dwEnable" json:"enable"`
}

func (o *MIBIPMulticastGlobal) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastGlobal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Enable); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastGlobal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	return nil
}

// MIBIPMulticastInterfaceEntry structure represents MIB_IPMCAST_IF_ENTRY RPC structure.
type MIBIPMulticastInterfaceEntry struct {
	InterfaceIndex     uint32 `idl:"name:dwIfIndex" json:"interface_index"`
	TTL                uint32 `idl:"name:dwTtl" json:"ttl"`
	Protocol           uint32 `idl:"name:dwProtocol" json:"protocol"`
	RateLimit          uint32 `idl:"name:dwRateLimit" json:"rate_limit"`
	InMulticastOctets  uint32 `idl:"name:ulInMcastOctets" json:"in_multicast_octets"`
	OutMulticastOctets uint32 `idl:"name:ulOutMcastOctets" json:"out_multicast_octets"`
}

func (o *MIBIPMulticastInterfaceEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastInterfaceEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.TTL); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RateLimit); err != nil {
		return err
	}
	if err := w.WriteData(o.InMulticastOctets); err != nil {
		return err
	}
	if err := w.WriteData(o.OutMulticastOctets); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastInterfaceEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.TTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RateLimit); err != nil {
		return err
	}
	if err := w.ReadData(&o.InMulticastOctets); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutMulticastOctets); err != nil {
		return err
	}
	return nil
}

// MIBIPMulticastInterfaceTable structure represents MIB_IPMCAST_IF_TABLE RPC structure.
type MIBIPMulticastInterfaceTable struct {
	EntriesLength uint32                          `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPMulticastInterfaceEntry `idl:"name:table" json:"table"`
}

func (o *MIBIPMulticastInterfaceTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastInterfaceTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPMulticastInterfaceEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPMulticastInterfaceEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIPMulticastInterfaceTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPMulticastInterfaceEntry, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPMulticastInterfaceEntry{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIPMulticastOutgoingInterface structure represents MIB_IPMCAST_OIF RPC structure.
type MIBIPMulticastOutgoingInterface struct {
	OutInterfaceIndex uint32 `idl:"name:dwOutIfIndex" json:"out_interface_index"`
	NextHopAddr       uint32 `idl:"name:dwNextHopAddr" json:"next_hop_addr"`
	_                 []byte `idl:"name:pvReserved"`
	_                 uint32 `idl:"name:dwReserved"`
}

func (o *MIBIPMulticastOutgoingInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastOutgoingInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OutInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.NextHopAddr); err != nil {
		return err
	}
	// reserved pvReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastOutgoingInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.NextHopAddr); err != nil {
		return err
	}
	// reserved pvReserved
	var _pvReserved []byte
	_ptr_pvReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		// FIXME: unknown type pvReserved
		return nil
	})
	_s_pvReserved := func(ptr interface{}) { _pvReserved = *ptr.(*[]byte) }
	if err := w.ReadPointer(&_pvReserved, _s_pvReserved, _ptr_pvReserved); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MIBIPMulticastMulticastForwardingEntry structure represents MIB_IPMCAST_MFE RPC structure.
type MIBIPMulticastMulticastForwardingEntry struct {
	Group               uint32                             `idl:"name:dwGroup" json:"group"`
	Source              uint32                             `idl:"name:dwSource" json:"source"`
	SourceMask          uint32                             `idl:"name:dwSrcMask" json:"source_mask"`
	UpstreamNeighbor    uint32                             `idl:"name:dwUpStrmNgbr" json:"upstream_neighbor"`
	InInterfaceIndex    uint32                             `idl:"name:dwInIfIndex" json:"in_interface_index"`
	InInterfaceProtocol uint32                             `idl:"name:dwInIfProtocol" json:"in_interface_protocol"`
	RouteProtocol       uint32                             `idl:"name:dwRouteProtocol" json:"route_protocol"`
	RouteNetwork        uint32                             `idl:"name:dwRouteNetwork" json:"route_network"`
	RouteMask           uint32                             `idl:"name:dwRouteMask" json:"route_mask"`
	UpTime              uint32                             `idl:"name:ulUpTime" json:"up_time"`
	ExpiryTime          uint32                             `idl:"name:ulExpiryTime" json:"expiry_time"`
	Timeout             uint32                             `idl:"name:ulTimeOut" json:"timeout"`
	NumOutInterface     uint32                             `idl:"name:ulNumOutIf" json:"num_out_interface"`
	Flags               uint32                             `idl:"name:fFlags" json:"flags"`
	_                   uint32                             `idl:"name:dwReserved"`
	RgmioOutInfo        []*MIBIPMulticastOutgoingInterface `idl:"name:rgmioOutInfo" json:"rgmio_out_info"`
}

func (o *MIBIPMulticastMulticastForwardingEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastMulticastForwardingEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Group); err != nil {
		return err
	}
	if err := w.WriteData(o.Source); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceMask); err != nil {
		return err
	}
	if err := w.WriteData(o.UpstreamNeighbor); err != nil {
		return err
	}
	if err := w.WriteData(o.InInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.InInterfaceProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteNetwork); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteMask); err != nil {
		return err
	}
	if err := w.WriteData(o.UpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.ExpiryTime); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.NumOutInterface); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	for i1 := range o.RgmioOutInfo {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.RgmioOutInfo[i1] != nil {
			if err := o.RgmioOutInfo[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPMulticastOutgoingInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.RgmioOutInfo); uint64(i1) < 1; i1++ {
		if err := (&MIBIPMulticastOutgoingInterface{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIPMulticastMulticastForwardingEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Group); err != nil {
		return err
	}
	if err := w.ReadData(&o.Source); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpstreamNeighbor); err != nil {
		return err
	}
	if err := w.ReadData(&o.InInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.InInterfaceProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteNetwork); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExpiryTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumOutInterface); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	o.RgmioOutInfo = make([]*MIBIPMulticastOutgoingInterface, 1)
	for i1 := range o.RgmioOutInfo {
		i1 := i1
		if o.RgmioOutInfo[i1] == nil {
			o.RgmioOutInfo[i1] = &MIBIPMulticastOutgoingInterface{}
		}
		if err := o.RgmioOutInfo[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIPMulticastOutgoingInterfaceStats structure represents MIB_IPMCAST_OIF_STATS RPC structure.
type MIBIPMulticastOutgoingInterfaceStats struct {
	OutInterfaceIndex uint32 `idl:"name:dwOutIfIndex" json:"out_interface_index"`
	NextHopAddr       uint32 `idl:"name:dwNextHopAddr" json:"next_hop_addr"`
	DialContext       []byte `idl:"name:pvDialContext" json:"dial_context"`
	TTLTooLow         uint32 `idl:"name:ulTtlTooLow" json:"ttl_too_low"`
	FragNeeded        uint32 `idl:"name:ulFragNeeded" json:"frag_needed"`
	OutPackets        uint32 `idl:"name:ulOutPackets" json:"out_packets"`
	OutDiscards       uint32 `idl:"name:ulOutDiscards" json:"out_discards"`
}

func (o *MIBIPMulticastOutgoingInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastOutgoingInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OutInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.NextHopAddr); err != nil {
		return err
	}
	if o.DialContext != nil {
		_ptr_pvDialContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			// FIXME unknown type pvDialContext
			return nil
		})
		if err := w.WritePointer(&o.DialContext, _ptr_pvDialContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TTLTooLow); err != nil {
		return err
	}
	if err := w.WriteData(o.FragNeeded); err != nil {
		return err
	}
	if err := w.WriteData(o.OutPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.OutDiscards); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastOutgoingInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.NextHopAddr); err != nil {
		return err
	}
	_ptr_pvDialContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		// FIXME: unknown type pvDialContext
		return nil
	})
	_s_pvDialContext := func(ptr interface{}) { o.DialContext = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.DialContext, _s_pvDialContext, _ptr_pvDialContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.TTLTooLow); err != nil {
		return err
	}
	if err := w.ReadData(&o.FragNeeded); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutDiscards); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MIBIPMulticastMulticastForwardingEntryStats structure represents MIB_IPMCAST_MFE_STATS RPC structure.
type MIBIPMulticastMulticastForwardingEntryStats struct {
	Group                     uint32                                  `idl:"name:dwGroup" json:"group"`
	Source                    uint32                                  `idl:"name:dwSource" json:"source"`
	SourceMask                uint32                                  `idl:"name:dwSrcMask" json:"source_mask"`
	UpstreamNeighbor          uint32                                  `idl:"name:dwUpStrmNgbr" json:"upstream_neighbor"`
	InInterfaceIndex          uint32                                  `idl:"name:dwInIfIndex" json:"in_interface_index"`
	InInterfaceProtocol       uint32                                  `idl:"name:dwInIfProtocol" json:"in_interface_protocol"`
	RouteProtocol             uint32                                  `idl:"name:dwRouteProtocol" json:"route_protocol"`
	RouteNetwork              uint32                                  `idl:"name:dwRouteNetwork" json:"route_network"`
	RouteMask                 uint32                                  `idl:"name:dwRouteMask" json:"route_mask"`
	UpTime                    uint32                                  `idl:"name:ulUpTime" json:"up_time"`
	ExpiryTime                uint32                                  `idl:"name:ulExpiryTime" json:"expiry_time"`
	NumOutInterface           uint32                                  `idl:"name:ulNumOutIf" json:"num_out_interface"`
	InPackets                 uint32                                  `idl:"name:ulInPkts" json:"in_packets"`
	InOctets                  uint32                                  `idl:"name:ulInOctets" json:"in_octets"`
	PacketsDifferentInterface uint32                                  `idl:"name:ulPktsDifferentIf" json:"packets_different_interface"`
	QueueOverflow             uint32                                  `idl:"name:ulQueueOverflow" json:"queue_overflow"`
	OutStats                  []*MIBIPMulticastOutgoingInterfaceStats `idl:"name:rgmiosOutStats" json:"out_stats"`
}

func (o *MIBIPMulticastMulticastForwardingEntryStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastMulticastForwardingEntryStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Group); err != nil {
		return err
	}
	if err := w.WriteData(o.Source); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceMask); err != nil {
		return err
	}
	if err := w.WriteData(o.UpstreamNeighbor); err != nil {
		return err
	}
	if err := w.WriteData(o.InInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.InInterfaceProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteNetwork); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteMask); err != nil {
		return err
	}
	if err := w.WriteData(o.UpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.ExpiryTime); err != nil {
		return err
	}
	if err := w.WriteData(o.NumOutInterface); err != nil {
		return err
	}
	if err := w.WriteData(o.InPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.InOctets); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketsDifferentInterface); err != nil {
		return err
	}
	if err := w.WriteData(o.QueueOverflow); err != nil {
		return err
	}
	for i1 := range o.OutStats {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.OutStats[i1] != nil {
			if err := o.OutStats[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPMulticastOutgoingInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.OutStats); uint64(i1) < 1; i1++ {
		if err := (&MIBIPMulticastOutgoingInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIPMulticastMulticastForwardingEntryStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Group); err != nil {
		return err
	}
	if err := w.ReadData(&o.Source); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpstreamNeighbor); err != nil {
		return err
	}
	if err := w.ReadData(&o.InInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.InInterfaceProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteNetwork); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExpiryTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumOutInterface); err != nil {
		return err
	}
	if err := w.ReadData(&o.InPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.InOctets); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketsDifferentInterface); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueueOverflow); err != nil {
		return err
	}
	o.OutStats = make([]*MIBIPMulticastOutgoingInterfaceStats, 1)
	for i1 := range o.OutStats {
		i1 := i1
		if o.OutStats[i1] == nil {
			o.OutStats[i1] = &MIBIPMulticastOutgoingInterfaceStats{}
		}
		if err := o.OutStats[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIPMulticastScope structure represents MIB_IPMCAST_SCOPE RPC structure.
type MIBIPMulticastScope struct {
	GroupAddress uint32   `idl:"name:dwGroupAddress" json:"group_address"`
	GroupMask    uint32   `idl:"name:dwGroupMask" json:"group_mask"`
	NameBuffer   []uint16 `idl:"name:snNameBuffer" json:"name_buffer"`
	Status       uint32   `idl:"name:dwStatus" json:"status"`
	_            []byte   `idl:"name:reserved"`
}

func (o *MIBIPMulticastScope) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastScope) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupMask); err != nil {
		return err
	}
	for i1 := range o.NameBuffer {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.NameBuffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NameBuffer); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	// reserved reserved
	for i1 := 0; uint64(i1) < 492; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPMulticastScope) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupMask); err != nil {
		return err
	}
	o.NameBuffer = make([]uint16, 256)
	for i1 := range o.NameBuffer {
		i1 := i1
		if err := w.ReadData(&o.NameBuffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	// reserved reserved
	var _reserved []byte
	_reserved = make([]byte, 492)
	for i1 := range _reserved {
		i1 := i1
		if err := w.ReadData(&_reserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MIBIPNetRow structure represents MIB_IPNETROW RPC structure.
type MIBIPNetRow struct {
	Index              uint32 `idl:"name:dwIndex" json:"index"`
	PhysicalAddrLength uint32 `idl:"name:dwPhysAddrLen" json:"physical_addr_length"`
	PhysicalAddr       []byte `idl:"name:bPhysAddr" json:"physical_addr"`
	Addr               uint32 `idl:"name:dwAddr" json:"addr"`
	Type               uint32 `idl:"name:dwType" json:"type"`
}

func (o *MIBIPNetRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPNetRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.PhysicalAddrLength); err != nil {
		return err
	}
	for i1 := range o.PhysicalAddr {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.PhysicalAddr[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PhysicalAddr); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Addr); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPNetRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhysicalAddrLength); err != nil {
		return err
	}
	o.PhysicalAddr = make([]byte, 8)
	for i1 := range o.PhysicalAddr {
		i1 := i1
		if err := w.ReadData(&o.PhysicalAddr[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Addr); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	return nil
}

// MIBIPNetTable structure represents MIB_IPNETTABLE RPC structure.
type MIBIPNetTable struct {
	EntriesLength uint32         `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPNetRow `idl:"name:table" json:"table"`
	_             []byte         `idl:"name:reserved"`
}

func (o *MIBIPNetTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPNetTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPNetRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPNetRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	for i1 := 0; uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPNetTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPNetRow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPNetRow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	var _reserved []byte
	_reserved = make([]byte, 8)
	for i1 := range _reserved {
		i1 := i1
		if err := w.ReadData(&_reserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MIBIPStats structure represents MIB_IPSTATS RPC structure.
type MIBIPStats struct {
	Forwarding      MIBIPStatsForwarding `idl:"name:Forwarding" json:"forwarding"`
	DefaultTTL      uint32               `idl:"name:dwDefaultTTL" json:"default_ttl"`
	InReceives      uint32               `idl:"name:dwInReceives" json:"in_receives"`
	InHdrErrors     uint32               `idl:"name:dwInHdrErrors" json:"in_hdr_errors"`
	InAddrErrors    uint32               `idl:"name:dwInAddrErrors" json:"in_addr_errors"`
	ForwDatagrams   uint32               `idl:"name:dwForwDatagrams" json:"forw_datagrams"`
	InUnknownProtos uint32               `idl:"name:dwInUnknownProtos" json:"in_unknown_protos"`
	InDiscards      uint32               `idl:"name:dwInDiscards" json:"in_discards"`
	InDelivers      uint32               `idl:"name:dwInDelivers" json:"in_delivers"`
	OutRequests     uint32               `idl:"name:dwOutRequests" json:"out_requests"`
	RoutingDiscards uint32               `idl:"name:dwRoutingDiscards" json:"routing_discards"`
	OutDiscards     uint32               `idl:"name:dwOutDiscards" json:"out_discards"`
	OutNoRoutes     uint32               `idl:"name:dwOutNoRoutes" json:"out_no_routes"`
	ReasmTimeout    uint32               `idl:"name:dwReasmTimeout" json:"reasm_timeout"`
	ReasmReqds      uint32               `idl:"name:dwReasmReqds" json:"reasm_reqds"`
	ReasmOKs        uint32               `idl:"name:dwReasmOks" json:"reasm_oks"`
	ReasmFails      uint32               `idl:"name:dwReasmFails" json:"reasm_fails"`
	FragOKs         uint32               `idl:"name:dwFragOks" json:"frag_oks"`
	FragFails       uint32               `idl:"name:dwFragFails" json:"frag_fails"`
	FragCreates     uint32               `idl:"name:dwFragCreates" json:"frag_creates"`
	InterfaceLength uint32               `idl:"name:dwNumIf" json:"interface_length"`
	AddrLength      uint32               `idl:"name:dwNumAddr" json:"addr_length"`
	RoutesLength    uint32               `idl:"name:dwNumRoutes" json:"routes_length"`
}

func (o *MIBIPStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Forwarding)); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultTTL); err != nil {
		return err
	}
	if err := w.WriteData(o.InReceives); err != nil {
		return err
	}
	if err := w.WriteData(o.InHdrErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.InAddrErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.ForwDatagrams); err != nil {
		return err
	}
	if err := w.WriteData(o.InUnknownProtos); err != nil {
		return err
	}
	if err := w.WriteData(o.InDiscards); err != nil {
		return err
	}
	if err := w.WriteData(o.InDelivers); err != nil {
		return err
	}
	if err := w.WriteData(o.OutRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.RoutingDiscards); err != nil {
		return err
	}
	if err := w.WriteData(o.OutDiscards); err != nil {
		return err
	}
	if err := w.WriteData(o.OutNoRoutes); err != nil {
		return err
	}
	if err := w.WriteData(o.ReasmTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.ReasmReqds); err != nil {
		return err
	}
	if err := w.WriteData(o.ReasmOKs); err != nil {
		return err
	}
	if err := w.WriteData(o.ReasmFails); err != nil {
		return err
	}
	if err := w.WriteData(o.FragOKs); err != nil {
		return err
	}
	if err := w.WriteData(o.FragFails); err != nil {
		return err
	}
	if err := w.WriteData(o.FragCreates); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrLength); err != nil {
		return err
	}
	if err := w.WriteData(o.RoutesLength); err != nil {
		return err
	}
	return nil
}
func (o *MIBIPStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Forwarding)); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultTTL); err != nil {
		return err
	}
	if err := w.ReadData(&o.InReceives); err != nil {
		return err
	}
	if err := w.ReadData(&o.InHdrErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.InAddrErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForwDatagrams); err != nil {
		return err
	}
	if err := w.ReadData(&o.InUnknownProtos); err != nil {
		return err
	}
	if err := w.ReadData(&o.InDiscards); err != nil {
		return err
	}
	if err := w.ReadData(&o.InDelivers); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoutingDiscards); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutDiscards); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutNoRoutes); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReasmTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReasmReqds); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReasmOKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReasmFails); err != nil {
		return err
	}
	if err := w.ReadData(&o.FragOKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.FragFails); err != nil {
		return err
	}
	if err := w.ReadData(&o.FragCreates); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoutesLength); err != nil {
		return err
	}
	return nil
}

// MIBMulticastForwardingEntryStatsTable structure represents MIB_MFE_STATS_TABLE RPC structure.
type MIBMulticastForwardingEntryStatsTable struct {
	EntriesLength uint32                                         `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPMulticastMulticastForwardingEntryStats `idl:"name:table" json:"table"`
}

func (o *MIBMulticastForwardingEntryStatsTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBMulticastForwardingEntryStatsTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPMulticastMulticastForwardingEntryStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPMulticastMulticastForwardingEntryStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBMulticastForwardingEntryStatsTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPMulticastMulticastForwardingEntryStats, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPMulticastMulticastForwardingEntryStats{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBMulticastForwardingEntryTable structure represents MIB_MFE_TABLE RPC structure.
type MIBMulticastForwardingEntryTable struct {
	EntriesLength uint32                                    `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIPMulticastMulticastForwardingEntry `idl:"name:table" json:"table"`
}

func (o *MIBMulticastForwardingEntryTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBMulticastForwardingEntryTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIPMulticastMulticastForwardingEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIPMulticastMulticastForwardingEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBMulticastForwardingEntryTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIPMulticastMulticastForwardingEntry, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIPMulticastMulticastForwardingEntry{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBOpaqueInfo structure represents MIB_OPAQUE_INFO RPC structure.
type MIBOpaqueInfo struct {
	ID    uint32 `idl:"name:dwId" json:"id"`
	Align uint64 `idl:"name:ullAlign" json:"align"`
}

func (o *MIBOpaqueInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBOpaqueInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.Align); err != nil {
		return err
	}
	return nil
}
func (o *MIBOpaqueInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Align); err != nil {
		return err
	}
	return nil
}

// MIBOpaqueQuery structure represents MIB_OPAQUE_QUERY RPC structure.
type MIBOpaqueQuery struct {
	VarID    uint32   `idl:"name:dwVarId" json:"var_id"`
	VarIndex []uint32 `idl:"name:rgdwVarIndex" json:"var_index"`
}

func (o *MIBOpaqueQuery) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBOpaqueQuery) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.VarID); err != nil {
		return err
	}
	for i1 := range o.VarIndex {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.VarIndex[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.VarIndex); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBOpaqueQuery) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.VarID); err != nil {
		return err
	}
	o.VarIndex = make([]uint32, 1)
	for i1 := range o.VarIndex {
		i1 := i1
		if err := w.ReadData(&o.VarIndex[i1]); err != nil {
			return err
		}
	}
	return nil
}

// MIBProxyArp structure represents MIB_PROXYARP RPC structure.
type MIBProxyArp struct {
	Address        uint32 `idl:"name:dwAddress" json:"address"`
	Mask           uint32 `idl:"name:dwMask" json:"mask"`
	InterfaceIndex uint32 `idl:"name:dwIfIndex" json:"interface_index"`
}

func (o *MIBProxyArp) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBProxyArp) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Address); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *MIBProxyArp) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Address); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}

// MIBTCPRow structure represents MIB_TCPROW RPC structure.
type MIBTCPRow struct {
	State      MIBTCPState `idl:"name:State" json:"state"`
	LocalAddr  uint32      `idl:"name:dwLocalAddr" json:"local_addr"`
	LocalPort  uint32      `idl:"name:dwLocalPort" json:"local_port"`
	RemoteAddr uint32      `idl:"name:dwRemoteAddr" json:"remote_addr"`
	RemotePort uint32      `idl:"name:dwRemotePort" json:"remote_port"`
}

func (o *MIBTCPRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBTCPRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.State)); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalPort); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.RemotePort); err != nil {
		return err
	}
	return nil
}
func (o *MIBTCPRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.State)); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemotePort); err != nil {
		return err
	}
	return nil
}

// MIBTCPStats structure represents MIB_TCPSTATS RPC structure.
type MIBTCPStats struct {
	RTOAlgorithm    TCPRTOAlgorithm `idl:"name:RtoAlgorithm" json:"rto_algorithm"`
	RTOMin          uint32          `idl:"name:dwRtoMin" json:"rto_min"`
	RTOMax          uint32          `idl:"name:dwRtoMax" json:"rto_max"`
	MaxConn         uint32          `idl:"name:dwMaxConn" json:"max_conn"`
	ActiveOpens     uint32          `idl:"name:dwActiveOpens" json:"active_opens"`
	PassiveOpens    uint32          `idl:"name:dwPassiveOpens" json:"passive_opens"`
	AttemptFails    uint32          `idl:"name:dwAttemptFails" json:"attempt_fails"`
	EstabResets     uint32          `idl:"name:dwEstabResets" json:"estab_resets"`
	CurrentEstab    uint32          `idl:"name:dwCurrEstab" json:"current_estab"`
	InSegs          uint32          `idl:"name:dwInSegs" json:"in_segs"`
	OutSegs         uint32          `idl:"name:dwOutSegs" json:"out_segs"`
	RetransimitSegs uint32          `idl:"name:dwRetransSegs" json:"retransimit_segs"`
	InErrs          uint32          `idl:"name:dwInErrs" json:"in_errs"`
	OutRsts         uint32          `idl:"name:dwOutRsts" json:"out_rsts"`
	ConnsLength     uint32          `idl:"name:dwNumConns" json:"conns_length"`
}

func (o *MIBTCPStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBTCPStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RTOAlgorithm)); err != nil {
		return err
	}
	if err := w.WriteData(o.RTOMin); err != nil {
		return err
	}
	if err := w.WriteData(o.RTOMax); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxConn); err != nil {
		return err
	}
	if err := w.WriteData(o.ActiveOpens); err != nil {
		return err
	}
	if err := w.WriteData(o.PassiveOpens); err != nil {
		return err
	}
	if err := w.WriteData(o.AttemptFails); err != nil {
		return err
	}
	if err := w.WriteData(o.EstabResets); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentEstab); err != nil {
		return err
	}
	if err := w.WriteData(o.InSegs); err != nil {
		return err
	}
	if err := w.WriteData(o.OutSegs); err != nil {
		return err
	}
	if err := w.WriteData(o.RetransimitSegs); err != nil {
		return err
	}
	if err := w.WriteData(o.InErrs); err != nil {
		return err
	}
	if err := w.WriteData(o.OutRsts); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnsLength); err != nil {
		return err
	}
	return nil
}
func (o *MIBTCPStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RTOAlgorithm)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTOMin); err != nil {
		return err
	}
	if err := w.ReadData(&o.RTOMax); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxConn); err != nil {
		return err
	}
	if err := w.ReadData(&o.ActiveOpens); err != nil {
		return err
	}
	if err := w.ReadData(&o.PassiveOpens); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttemptFails); err != nil {
		return err
	}
	if err := w.ReadData(&o.EstabResets); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentEstab); err != nil {
		return err
	}
	if err := w.ReadData(&o.InSegs); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutSegs); err != nil {
		return err
	}
	if err := w.ReadData(&o.RetransimitSegs); err != nil {
		return err
	}
	if err := w.ReadData(&o.InErrs); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutRsts); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnsLength); err != nil {
		return err
	}
	return nil
}

// MIBTCPTable structure represents MIB_TCPTABLE RPC structure.
type MIBTCPTable struct {
	EntriesLength uint32       `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBTCPRow `idl:"name:table" json:"table"`
	_             []byte       `idl:"name:reserved"`
}

func (o *MIBTCPTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBTCPTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBTCPRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBTCPRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	for i1 := 0; uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MIBTCPTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBTCPRow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBTCPRow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	var _reserved []byte
	_reserved = make([]byte, 8)
	for i1 := range _reserved {
		i1 := i1
		if err := w.ReadData(&_reserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MIBUDPRow structure represents MIB_UDPROW RPC structure.
type MIBUDPRow struct {
	LocalAddr uint32 `idl:"name:dwLocalAddr" json:"local_addr"`
	LocalPort uint32 `idl:"name:dwLocalPort" json:"local_port"`
}

func (o *MIBUDPRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBUDPRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalPort); err != nil {
		return err
	}
	return nil
}
func (o *MIBUDPRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalPort); err != nil {
		return err
	}
	return nil
}

// MIBUDPStats structure represents MIB_UDPSTATS RPC structure.
type MIBUDPStats struct {
	InDatagrams  uint32 `idl:"name:dwInDatagrams" json:"in_datagrams"`
	NoPorts      uint32 `idl:"name:dwNoPorts" json:"no_ports"`
	InErrors     uint32 `idl:"name:dwInErrors" json:"in_errors"`
	OutDatagrams uint32 `idl:"name:dwOutDatagrams" json:"out_datagrams"`
	AddrsLength  uint32 `idl:"name:dwNumAddrs" json:"addrs_length"`
}

func (o *MIBUDPStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBUDPStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InDatagrams); err != nil {
		return err
	}
	if err := w.WriteData(o.NoPorts); err != nil {
		return err
	}
	if err := w.WriteData(o.InErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.OutDatagrams); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrsLength); err != nil {
		return err
	}
	return nil
}
func (o *MIBUDPStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InDatagrams); err != nil {
		return err
	}
	if err := w.ReadData(&o.NoPorts); err != nil {
		return err
	}
	if err := w.ReadData(&o.InErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutDatagrams); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrsLength); err != nil {
		return err
	}
	return nil
}

// MIBUDPTable structure represents MIB_UDPTABLE RPC structure.
type MIBUDPTable struct {
	EntriesLength uint32       `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBUDPRow `idl:"name:table" json:"table"`
	_             []byte       `idl:"name:reserved"`
}

func (o *MIBUDPTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBUDPTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	for i1 := range o.Table {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Table[i1] != nil {
			if err := o.Table[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBUDPRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBUDPRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	for i1 := 0; uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *MIBUDPTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBUDPRow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBUDPRow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved reserved
	var _reserved []byte
	_reserved = make([]byte, 8)
	for i1 := range _reserved {
		i1 := i1
		if err := w.ReadData(&_reserved[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// Server0 structure represents MPR_SERVER_0 RPC structure.
type Server0 struct {
	LANOnlyMode bool   `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime      uint32 `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts  uint32 `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse  uint32 `idl:"name:dwPortsInUse" json:"ports_in_use"`
}

func (o *Server0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Server0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.LANOnlyMode {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalPorts); err != nil {
		return err
	}
	if err := w.WriteData(o.PortsInUse); err != nil {
		return err
	}
	return nil
}
func (o *Server0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bLANOnlyMode int32
	if err := w.ReadData(&_bLANOnlyMode); err != nil {
		return err
	}
	o.LANOnlyMode = _bLANOnlyMode != 0
	if err := w.ReadData(&o.UpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalPorts); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortsInUse); err != nil {
		return err
	}
	return nil
}

// Server1 structure represents MPR_SERVER_1 RPC structure.
type Server1 struct {
	PPTPPortsLength uint32 `idl:"name:dwNumPptpPorts" json:"pptp_ports_length"`
	PPTPPortFlags   uint32 `idl:"name:dwPptpPortFlags" json:"pptp_port_flags"`
	L2TPPortsLength uint32 `idl:"name:dwNumL2tpPorts" json:"l2tp_ports_length"`
	L2TPPortFlags   uint32 `idl:"name:dwL2tpPortFlags" json:"l2tp_port_flags"`
}

func (o *Server1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Server1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PPTPPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PPTPPortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.L2TPPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.L2TPPortFlags); err != nil {
		return err
	}
	return nil
}
func (o *Server1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2TPPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2TPPortFlags); err != nil {
		return err
	}
	return nil
}

// Server2 structure represents MPR_SERVER_2 RPC structure.
type Server2 struct {
	PPTPPortsLength uint32 `idl:"name:dwNumPptpPorts" json:"pptp_ports_length"`
	PPTPPortFlags   uint32 `idl:"name:dwPptpPortFlags" json:"pptp_port_flags"`
	L2TPPortsLength uint32 `idl:"name:dwNumL2tpPorts" json:"l2tp_ports_length"`
	L2TPPortFlags   uint32 `idl:"name:dwL2tpPortFlags" json:"l2tp_port_flags"`
	SSTPPortsLength uint32 `idl:"name:dwNumSstpPorts" json:"sstp_ports_length"`
	SSTPPortFlags   uint32 `idl:"name:dwSstpPortFlags" json:"sstp_port_flags"`
}

func (o *Server2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Server2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PPTPPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PPTPPortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.L2TPPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.L2TPPortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.SSTPPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.SSTPPortFlags); err != nil {
		return err
	}
	return nil
}
func (o *Server2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2TPPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2TPPortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.SSTPPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.SSTPPortFlags); err != nil {
		return err
	}
	return nil
}

// PPPNbfcpInfo structure represents PPP_NBFCP_INFO RPC structure.
type PPPNbfcpInfo struct {
	Error       uint32   `idl:"name:dwError" json:"error"`
	Workstation []uint16 `idl:"name:wszWksta" json:"workstation"`
}

func (o *PPPNbfcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPNbfcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	for i1 := range o.Workstation {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.Workstation[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Workstation); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *PPPNbfcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	o.Workstation = make([]uint16, 17)
	for i1 := range o.Workstation {
		i1 := i1
		if err := w.ReadData(&o.Workstation[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// PPPIPCPInfo structure represents PPP_IPCP_INFO RPC structure.
type PPPIPCPInfo struct {
	Error         uint32   `idl:"name:dwError" json:"error"`
	Address       []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
}

func (o *PPPIPCPInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPIPCPInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteAddress {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteAddress); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *PPPIPCPInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	o.Address = make([]uint16, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	o.RemoteAddress = make([]uint16, 16)
	for i1 := range o.RemoteAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// PPPIPCPInfo2 structure represents PPP_IPCP_INFO2 RPC structure.
type PPPIPCPInfo2 struct {
	Error         uint32   `idl:"name:dwError" json:"error"`
	Address       []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
	Options       uint32   `idl:"name:dwOptions" json:"options"`
	RemoteOptons  uint32   `idl:"name:dwRemoteOptons" json:"remote_optons"`
}

func (o *PPPIPCPInfo2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPIPCPInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteAddress {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteAddress); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteOptons); err != nil {
		return err
	}
	return nil
}
func (o *PPPIPCPInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	o.Address = make([]uint16, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	o.RemoteAddress = make([]uint16, 16)
	for i1 := range o.RemoteAddress {
		i1 := i1
		if err := w.ReadData(&o.RemoteAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteOptons); err != nil {
		return err
	}
	return nil
}

// PPPIpxcpInfo structure represents PPP_IPXCP_INFO RPC structure.
type PPPIpxcpInfo struct {
	Error   uint32   `idl:"name:dwError" json:"error"`
	Address []uint16 `idl:"name:wszAddress" json:"address"`
}

func (o *PPPIpxcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPIpxcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *PPPIpxcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	o.Address = make([]uint16, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// PPPIPv6CreatePartitionInfo structure represents PPP_IPV6_CP_INFO RPC structure.
type PPPIPv6CreatePartitionInfo struct {
	Version           uint32 `idl:"name:dwVersion" json:"version"`
	Size              uint32 `idl:"name:dwSize" json:"size"`
	Error             uint32 `idl:"name:dwError" json:"error"`
	InterfaceID       []byte `idl:"name:bInterfaceIdentifier" json:"interface_id"`
	RemoteInterfaceID []byte `idl:"name:bRemoteInterfaceIdentifier" json:"remote_interface_id"`
	Options           uint32 `idl:"name:dwOptions" json:"options"`
	RemoteOptions     uint32 `idl:"name:dwRemoteOptions" json:"remote_options"`
	Prefix            []byte `idl:"name:bPrefix" json:"prefix"`
	PrefixLength      uint32 `idl:"name:dwPrefixLength" json:"prefix_length"`
}

func (o *PPPIPv6CreatePartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPIPv6CreatePartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	for i1 := range o.InterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteInterfaceID); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteOptions); err != nil {
		return err
	}
	for i1 := range o.Prefix {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Prefix[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Prefix); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrefixLength); err != nil {
		return err
	}
	return nil
}
func (o *PPPIPv6CreatePartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	o.InterfaceID = make([]byte, 8)
	for i1 := range o.InterfaceID {
		i1 := i1
		if err := w.ReadData(&o.InterfaceID[i1]); err != nil {
			return err
		}
	}
	o.RemoteInterfaceID = make([]byte, 8)
	for i1 := range o.RemoteInterfaceID {
		i1 := i1
		if err := w.ReadData(&o.RemoteInterfaceID[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteOptions); err != nil {
		return err
	}
	o.Prefix = make([]byte, 8)
	for i1 := range o.Prefix {
		i1 := i1
		if err := w.ReadData(&o.Prefix[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PrefixLength); err != nil {
		return err
	}
	return nil
}

// PPPATCPInfo structure represents PPP_ATCP_INFO RPC structure.
type PPPATCPInfo struct {
	Error   uint32   `idl:"name:dwError" json:"error"`
	Address []uint16 `idl:"name:wszAddress" json:"address"`
}

func (o *PPPATCPInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPATCPInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *PPPATCPInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	o.Address = make([]uint16, 33)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// PPPCCPInfo structure represents PPP_CCP_INFO RPC structure.
type PPPCCPInfo struct {
	Error                      uint32 `idl:"name:dwError" json:"error"`
	CompressionAlgorithm       uint32 `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	Options                    uint32 `idl:"name:dwOptions" json:"options"`
	RemoteCompressionAlgorithm uint32 `idl:"name:dwRemoteCompressionAlgorithm" json:"remote_compression_algorithm"`
	RemoteOptions              uint32 `idl:"name:dwRemoteOptions" json:"remote_options"`
}

func (o *PPPCCPInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPCCPInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteOptions); err != nil {
		return err
	}
	return nil
}
func (o *PPPCCPInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteOptions); err != nil {
		return err
	}
	return nil
}

// PPPLCPInfo structure represents PPP_LCP_INFO RPC structure.
type PPPLCPInfo struct {
	Error                        uint32 `idl:"name:dwError" json:"error"`
	AuthenticationProtocol       uint32 `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	AuthenticationData           uint32 `idl:"name:dwAuthenticationData" json:"authentication_data"`
	RemoteAuthenticationProtocol uint32 `idl:"name:dwRemoteAuthenticationProtocol" json:"remote_authentication_protocol"`
	RemoteAuthenticationData     uint32 `idl:"name:dwRemoteAuthenticationData" json:"remote_authentication_data"`
	TerminateReason              uint32 `idl:"name:dwTerminateReason" json:"terminate_reason"`
	RemoteTerminateReason        uint32 `idl:"name:dwRemoteTerminateReason" json:"remote_terminate_reason"`
	Options                      uint32 `idl:"name:dwOptions" json:"options"`
	RemoteOptions                uint32 `idl:"name:dwRemoteOptions" json:"remote_options"`
	EAPTypeID                    uint32 `idl:"name:dwEapTypeId" json:"eap_type_id"`
	RemoteEAPTypeID              uint32 `idl:"name:dwRemoteEapTypeId" json:"remote_eap_type_id"`
}

func (o *PPPLCPInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPLCPInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationData); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAuthenticationProtocol); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAuthenticationData); err != nil {
		return err
	}
	if err := w.WriteData(o.TerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.EAPTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteEAPTypeID); err != nil {
		return err
	}
	return nil
}
func (o *PPPLCPInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationData); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAuthenticationProtocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAuthenticationData); err != nil {
		return err
	}
	if err := w.ReadData(&o.TerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.EAPTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteEAPTypeID); err != nil {
		return err
	}
	return nil
}

// PPPInfo structure represents PPP_INFO RPC structure.
type PPPInfo struct {
	NBF *PPPNbfcpInfo `idl:"name:nbf" json:"nbf"`
	IP  *PPPIPCPInfo  `idl:"name:ip" json:"ip"`
	IPX *PPPIpxcpInfo `idl:"name:ipx" json:"ipx"`
	AT  *PPPATCPInfo  `idl:"name:at" json:"at"`
}

func (o *PPPInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.NBF != nil {
		if err := o.NBF.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPNbfcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IP != nil {
		if err := o.IP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPIPCPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IPX != nil {
		if err := o.IPX.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPIpxcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AT != nil {
		if err := o.AT.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPATCPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PPPInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.NBF == nil {
		o.NBF = &PPPNbfcpInfo{}
	}
	if err := o.NBF.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IP == nil {
		o.IP = &PPPIPCPInfo{}
	}
	if err := o.IP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IPX == nil {
		o.IPX = &PPPIpxcpInfo{}
	}
	if err := o.IPX.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AT == nil {
		o.AT = &PPPATCPInfo{}
	}
	if err := o.AT.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PPPInfo2 structure represents PPP_INFO_2 RPC structure.
type PPPInfo2 struct {
	NBF *PPPNbfcpInfo `idl:"name:nbf" json:"nbf"`
	IP  *PPPIPCPInfo2 `idl:"name:ip" json:"ip"`
	IPX *PPPIpxcpInfo `idl:"name:ipx" json:"ipx"`
	AT  *PPPATCPInfo  `idl:"name:at" json:"at"`
	CCP *PPPCCPInfo   `idl:"name:ccp" json:"ccp"`
	LCP *PPPLCPInfo   `idl:"name:lcp" json:"lcp"`
}

func (o *PPPInfo2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.NBF != nil {
		if err := o.NBF.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPNbfcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IP != nil {
		if err := o.IP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPIPCPInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IPX != nil {
		if err := o.IPX.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPIpxcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AT != nil {
		if err := o.AT.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPATCPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CCP != nil {
		if err := o.CCP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPCCPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LCP != nil {
		if err := o.LCP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPLCPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PPPInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.NBF == nil {
		o.NBF = &PPPNbfcpInfo{}
	}
	if err := o.NBF.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IP == nil {
		o.IP = &PPPIPCPInfo2{}
	}
	if err := o.IP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IPX == nil {
		o.IPX = &PPPIpxcpInfo{}
	}
	if err := o.IPX.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AT == nil {
		o.AT = &PPPATCPInfo{}
	}
	if err := o.AT.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CCP == nil {
		o.CCP = &PPPCCPInfo{}
	}
	if err := o.CCP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LCP == nil {
		o.LCP = &PPPLCPInfo{}
	}
	if err := o.LCP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PPPInfo3 structure represents PPP_INFO_3 RPC structure.
type PPPInfo3 struct {
	NBF  *PPPNbfcpInfo               `idl:"name:nbf" json:"nbf"`
	IP   *PPPIPCPInfo2               `idl:"name:ip" json:"ip"`
	IPv6 *PPPIPv6CreatePartitionInfo `idl:"name:ipv6" json:"ipv6"`
	CCP  *PPPCCPInfo                 `idl:"name:ccp" json:"ccp"`
	LCP  *PPPLCPInfo                 `idl:"name:lcp" json:"lcp"`
}

func (o *PPPInfo3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PPPInfo3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.NBF != nil {
		if err := o.NBF.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPNbfcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IP != nil {
		if err := o.IP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPIPCPInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IPv6 != nil {
		if err := o.IPv6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPIPv6CreatePartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CCP != nil {
		if err := o.CCP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPCCPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LCP != nil {
		if err := o.LCP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPLCPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PPPInfo3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.NBF == nil {
		o.NBF = &PPPNbfcpInfo{}
	}
	if err := o.NBF.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IP == nil {
		o.IP = &PPPIPCPInfo2{}
	}
	if err := o.IP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IPv6 == nil {
		o.IPv6 = &PPPIPv6CreatePartitionInfo{}
	}
	if err := o.IPv6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CCP == nil {
		o.CCP = &PPPCCPInfo{}
	}
	if err := o.CCP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LCP == nil {
		o.LCP = &PPPLCPInfo{}
	}
	if err := o.LCP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASIPort0 structure represents RASI_PORT_0 RPC structure.
type RASIPort0 struct {
	Port               uint32           `idl:"name:dwPort" json:"port"`
	Connection         uint32           `idl:"name:dwConnection" json:"connection"`
	PortCondition      RASPortCondition `idl:"name:dwPortCondition" json:"port_condition"`
	TotalNumberOfCalls uint32           `idl:"name:dwTotalNumberOfCalls" json:"total_number_of_calls"`
	ConnectDuration    uint32           `idl:"name:dwConnectDuration" json:"connect_duration"`
	PortName           []uint16         `idl:"name:wszPortName" json:"port_name"`
	MediaName          []uint16         `idl:"name:wszMediaName" json:"media_name"`
	DeviceName         []uint16         `idl:"name:wszDeviceName" json:"device_name"`
	DeviceType         []uint16         `idl:"name:wszDeviceType" json:"device_type"`
}

func (o *RASIPort0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASIPort0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PortCondition)); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalNumberOfCalls); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectDuration); err != nil {
		return err
	}
	for i1 := range o.PortName {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.PortName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PortName); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.MediaName {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.MediaName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.MediaName); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *RASIPort0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PortCondition)); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalNumberOfCalls); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectDuration); err != nil {
		return err
	}
	o.PortName = make([]uint16, 17)
	for i1 := range o.PortName {
		i1 := i1
		if err := w.ReadData(&o.PortName[i1]); err != nil {
			return err
		}
	}
	o.MediaName = make([]uint16, 17)
	for i1 := range o.MediaName {
		i1 := i1
		if err := w.ReadData(&o.MediaName[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]uint16, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.DeviceType = make([]uint16, 17)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// RASIPort1 structure represents RASI_PORT_1 RPC structure.
type RASIPort1 struct {
	Port                 uint32               `idl:"name:dwPort" json:"port"`
	Connection           uint32               `idl:"name:dwConnection" json:"connection"`
	HardwareCondition    RASHardwareCondition `idl:"name:dwHardwareCondition" json:"hardware_condition"`
	LineSpeed            uint32               `idl:"name:dwLineSpeed" json:"line_speed"`
	BytesXmited          uint32               `idl:"name:dwBytesXmited" json:"bytes_xmited"`
	BytesRcved           uint32               `idl:"name:dwBytesRcved" json:"bytes_rcved"`
	FramesXmited         uint32               `idl:"name:dwFramesXmited" json:"frames_xmited"`
	FramesRcved          uint32               `idl:"name:dwFramesRcved" json:"frames_rcved"`
	CRCError             uint32               `idl:"name:dwCrcErr" json:"crc_error"`
	TimeoutError         uint32               `idl:"name:dwTimeoutErr" json:"timeout_error"`
	AlignmentError       uint32               `idl:"name:dwAlignmentErr" json:"alignment_error"`
	HardwareOverrunError uint32               `idl:"name:dwHardwareOverrunErr" json:"hardware_overrun_error"`
	FramingError         uint32               `idl:"name:dwFramingErr" json:"framing_error"`
	BufferOverrunError   uint32               `idl:"name:dwBufferOverrunErr" json:"buffer_overrun_error"`
	CompressionRatioIn   uint32               `idl:"name:dwCompressionRatioIn" json:"compression_ratio_in"`
	CompressionRatioOut  uint32               `idl:"name:dwCompressionRatioOut" json:"compression_ratio_out"`
}

func (o *RASIPort1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASIPort1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.HardwareCondition)); err != nil {
		return err
	}
	if err := w.WriteData(o.LineSpeed); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.CRCError); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutError); err != nil {
		return err
	}
	if err := w.WriteData(o.AlignmentError); err != nil {
		return err
	}
	if err := w.WriteData(o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.FramingError); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioOut); err != nil {
		return err
	}
	return nil
}
func (o *RASIPort1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.HardwareCondition)); err != nil {
		return err
	}
	if err := w.ReadData(&o.LineSpeed); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.CRCError); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutError); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlignmentError); err != nil {
		return err
	}
	if err := w.ReadData(&o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramingError); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioOut); err != nil {
		return err
	}
	return nil
}

// RASIConnection0 structure represents RASI_CONNECTION_0 RPC structure.
type RASIConnection0 struct {
	Connection      uint32              `idl:"name:dwConnection" json:"connection"`
	Interface       uint32              `idl:"name:dwInterface" json:"interface"`
	ConnectDuration uint32              `idl:"name:dwConnectDuration" json:"connect_duration"`
	InterfaceType   RouterInterfaceType `idl:"name:dwInterfaceType" json:"interface_type"`
	ConnectionFlags uint32              `idl:"name:dwConnectionFlags" json:"connection_flags"`
	InterfaceName   []uint16            `idl:"name:wszInterfaceName" json:"interface_name"`
	UserName        []uint16            `idl:"name:wszUserName" json:"user_name"`
	LogonDomain     []uint16            `idl:"name:wszLogonDomain" json:"logon_domain"`
	RemoteComputer  []uint16            `idl:"name:wszRemoteComputer" json:"remote_computer"`
}

func (o *RASIConnection0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASIConnection0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectDuration); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectionFlags); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.UserName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.UserName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UserName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.LogonDomain {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LogonDomain[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LogonDomain); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RemoteComputer {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.RemoteComputer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteComputer); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *RASIConnection0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectDuration); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectionFlags); err != nil {
		return err
	}
	o.InterfaceName = make([]uint16, 257)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	o.UserName = make([]uint16, 257)
	for i1 := range o.UserName {
		i1 := i1
		if err := w.ReadData(&o.UserName[i1]); err != nil {
			return err
		}
	}
	o.LogonDomain = make([]uint16, 16)
	for i1 := range o.LogonDomain {
		i1 := i1
		if err := w.ReadData(&o.LogonDomain[i1]); err != nil {
			return err
		}
	}
	o.RemoteComputer = make([]uint16, 17)
	for i1 := range o.RemoteComputer {
		i1 := i1
		if err := w.ReadData(&o.RemoteComputer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// RASIConnection1 structure represents RASI_CONNECTION_1 RPC structure.
type RASIConnection1 struct {
	Connection           uint32   `idl:"name:dwConnection" json:"connection"`
	Interface            uint32   `idl:"name:dwInterface" json:"interface"`
	PPPInfo              *PPPInfo `idl:"name:PppInfo" json:"ppp_info"`
	BytesXmited          uint32   `idl:"name:dwBytesXmited" json:"bytes_xmited"`
	BytesRcved           uint32   `idl:"name:dwBytesRcved" json:"bytes_rcved"`
	FramesXmited         uint32   `idl:"name:dwFramesXmited" json:"frames_xmited"`
	FramesRcved          uint32   `idl:"name:dwFramesRcved" json:"frames_rcved"`
	CRCError             uint32   `idl:"name:dwCrcErr" json:"crc_error"`
	TimeoutError         uint32   `idl:"name:dwTimeoutErr" json:"timeout_error"`
	AlignmentError       uint32   `idl:"name:dwAlignmentErr" json:"alignment_error"`
	HardwareOverrunError uint32   `idl:"name:dwHardwareOverrunErr" json:"hardware_overrun_error"`
	FramingError         uint32   `idl:"name:dwFramingErr" json:"framing_error"`
	BufferOverrunError   uint32   `idl:"name:dwBufferOverrunErr" json:"buffer_overrun_error"`
	CompressionRatioIn   uint32   `idl:"name:dwCompressionRatioIn" json:"compression_ratio_in"`
	CompressionRatioOut  uint32   `idl:"name:dwCompressionRatioOut" json:"compression_ratio_out"`
}

func (o *RASIConnection1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASIConnection1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if o.PPPInfo != nil {
		if err := o.PPPInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BytesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesXmited); err != nil {
		return err
	}
	if err := w.WriteData(o.FramesRcved); err != nil {
		return err
	}
	if err := w.WriteData(o.CRCError); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutError); err != nil {
		return err
	}
	if err := w.WriteData(o.AlignmentError); err != nil {
		return err
	}
	if err := w.WriteData(o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.FramingError); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionRatioOut); err != nil {
		return err
	}
	return nil
}
func (o *RASIConnection1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	if o.PPPInfo == nil {
		o.PPPInfo = &PPPInfo{}
	}
	if err := o.PPPInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesXmited); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramesRcved); err != nil {
		return err
	}
	if err := w.ReadData(&o.CRCError); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutError); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlignmentError); err != nil {
		return err
	}
	if err := w.ReadData(&o.HardwareOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.FramingError); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferOverrunError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioIn); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionRatioOut); err != nil {
		return err
	}
	return nil
}

// RASIConnection2 structure represents RASI_CONNECTION_2 RPC structure.
type RASIConnection2 struct {
	Connection    uint32              `idl:"name:dwConnection" json:"connection"`
	UserName      []uint16            `idl:"name:wszUserName" json:"user_name"`
	InterfaceType RouterInterfaceType `idl:"name:dwInterfaceType" json:"interface_type"`
	GUID          *dtyp.GUID          `idl:"name:guid" json:"guid"`
	PPPInfo2      *PPPInfo2           `idl:"name:PppInfo2" json:"ppp_info2"`
}

func (o *RASIConnection2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASIConnection2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	for i1 := range o.UserName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.UserName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UserName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PPPInfo2 != nil {
		if err := o.PPPInfo2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RASIConnection2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	o.UserName = make([]uint16, 257)
	for i1 := range o.UserName {
		i1 := i1
		if err := w.ReadData(&o.UserName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PPPInfo2 == nil {
		o.PPPInfo2 = &PPPInfo2{}
	}
	if err := o.PPPInfo2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASIConnection3 structure represents RASI_CONNECTION_3 RPC structure.
type RASIConnection3 struct {
	Version            uint32              `idl:"name:dwVersion" json:"version"`
	Size               uint32              `idl:"name:dwSize" json:"size"`
	Connection         uint32              `idl:"name:dwConnection" json:"connection"`
	UserName           []uint16            `idl:"name:wszUserName" json:"user_name"`
	InterfaceType      RouterInterfaceType `idl:"name:dwInterfaceType" json:"interface_type"`
	GUID               *dtyp.GUID          `idl:"name:guid" json:"guid"`
	PPPInfo3           *PPPInfo3           `idl:"name:PppInfo3" json:"ppp_info3"`
	RASQuarantineState RASQuarantineState  `idl:"name:rasQuarState" json:"ras_quarantine_state"`
	Timer              *dtyp.Filetime      `idl:"name:timer" json:"timer"`
}

func (o *RASIConnection3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASIConnection3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Connection); err != nil {
		return err
	}
	for i1 := range o.UserName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.UserName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UserName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PPPInfo3 != nil {
		if err := o.PPPInfo3.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PPPInfo3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.RASQuarantineState)); err != nil {
		return err
	}
	if o.Timer != nil {
		if err := o.Timer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RASIConnection3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	o.UserName = make([]uint16, 257)
	for i1 := range o.UserName {
		i1 := i1
		if err := w.ReadData(&o.UserName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PPPInfo3 == nil {
		o.PPPInfo3 = &PPPInfo3{}
	}
	if err := o.PPPInfo3.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RASQuarantineState)); err != nil {
		return err
	}
	if o.Timer == nil {
		o.Timer = &dtyp.Filetime{}
	}
	if err := o.Timer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Interface0 structure represents MPRI_INTERFACE_0 RPC structure.
type Interface0 struct {
	InterfaceName         []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface             uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled               bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType         RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState       RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError             uint32                `idl:"name:dwLastError" json:"last_error"`
}

func (o *Interface0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Interface0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ConnectionState)); err != nil {
		return err
	}
	if err := w.WriteData(o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.WriteData(o.LastError); err != nil {
		return err
	}
	return nil
}
func (o *Interface0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.InterfaceName = make([]uint16, 257)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ConnectionState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastError); err != nil {
		return err
	}
	return nil
}

// Interface1 structure represents MPRI_INTERFACE_1 RPC structure.
type Interface1 struct {
	InterfaceName           []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface               uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled                 bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType           RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState         RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons   uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError               uint32                `idl:"name:dwLastError" json:"last_error"`
	DialoutHoursRestriction string                `idl:"name:lpwsDialoutHoursRestriction" json:"dialout_hours_restriction"`
}

func (o *Interface1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Interface1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ConnectionState)); err != nil {
		return err
	}
	if err := w.WriteData(o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.WriteData(o.LastError); err != nil {
		return err
	}
	if o.DialoutHoursRestriction != "" {
		_ptr_lpwsDialoutHoursRestriction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.DialoutHoursRestriction); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DialoutHoursRestriction, _ptr_lpwsDialoutHoursRestriction); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Interface1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	o.InterfaceName = make([]uint16, 257)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ConnectionState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastError); err != nil {
		return err
	}
	_ptr_lpwsDialoutHoursRestriction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.DialoutHoursRestriction); err != nil {
			return err
		}
		return nil
	})
	_s_lpwsDialoutHoursRestriction := func(ptr interface{}) { o.DialoutHoursRestriction = *ptr.(*string) }
	if err := w.ReadPointer(&o.DialoutHoursRestriction, _s_lpwsDialoutHoursRestriction, _ptr_lpwsDialoutHoursRestriction); err != nil {
		return err
	}
	return nil
}

// Interface2 structure represents MPRI_INTERFACE_2 RPC structure.
type Interface2 struct {
	InterfaceName            []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface                uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled                  bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType            RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState          RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons    uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError                uint32                `idl:"name:dwLastError" json:"last_error"`
	Options                  uint32                `idl:"name:dwfOptions" json:"options"`
	LocalPhoneNumber         []uint16              `idl:"name:szLocalPhoneNumber" json:"local_phone_number"`
	Alternates               string                `idl:"name:szAlternates" json:"alternates"`
	IPAddr                   uint32                `idl:"name:ipaddr" json:"ip_addr"`
	IPAddrDNS                uint32                `idl:"name:ipaddrDns" json:"ip_addr_dns"`
	IPAddrDNSAlt             uint32                `idl:"name:ipaddrDnsAlt" json:"ip_addr_dns_alt"`
	IPAddrWINS               uint32                `idl:"name:ipaddrWins" json:"ip_addr_wins"`
	IPAddrWINSAlt            uint32                `idl:"name:ipaddrWinsAlt" json:"ip_addr_wins_alt"`
	NetProtocols             uint32                `idl:"name:dwfNetProtocols" json:"net_protocols"`
	DeviceType               []uint16              `idl:"name:szDeviceType" json:"device_type"`
	DeviceName               []uint16              `idl:"name:szDeviceName" json:"device_name"`
	X25PadType               []uint16              `idl:"name:szX25PadType" json:"x25_pad_type"`
	X25Address               []uint16              `idl:"name:szX25Address" json:"x25_address"`
	X25Facilities            []uint16              `idl:"name:szX25Facilities" json:"x25_facilities"`
	X25UserData              []uint16              `idl:"name:szX25UserData" json:"x25_user_data"`
	Channels                 uint32                `idl:"name:dwChannels" json:"channels"`
	SubEntries               uint32                `idl:"name:dwSubEntries" json:"sub_entries"`
	DialMode                 uint32                `idl:"name:dwDialMode" json:"dial_mode"`
	DialExtraPercent         uint32                `idl:"name:dwDialExtraPercent" json:"dial_extra_percent"`
	DialExtraSampleSeconds   uint32                `idl:"name:dwDialExtraSampleSeconds" json:"dial_extra_sample_seconds"`
	HangUpExtraPercent       uint32                `idl:"name:dwHangUpExtraPercent" json:"hang_up_extra_percent"`
	HangUpExtraSampleSeconds uint32                `idl:"name:dwHangUpExtraSampleSeconds" json:"hang_up_extra_sample_seconds"`
	IdleDisconnectSeconds    uint32                `idl:"name:dwIdleDisconnectSeconds" json:"idle_disconnect_seconds"`
	Type                     uint32                `idl:"name:dwType" json:"type"`
	EncryptionType           uint32                `idl:"name:dwEncryptionType" json:"encryption_type"`
	CustomAuthKey            uint32                `idl:"name:dwCustomAuthKey" json:"custom_auth_key"`
	CustomAuthDataSize       uint32                `idl:"name:dwCustomAuthDataSize" json:"custom_auth_data_size"`
	CustomAuthData           uint8                 `idl:"name:lpbCustomAuthData" json:"custom_auth_data"`
	ID                       *dtyp.GUID            `idl:"name:guidId" json:"id"`
	VPNStrategy              uint32                `idl:"name:dwVpnStrategy" json:"vpn_strategy"`
}

func (o *Interface2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Interface2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ConnectionState)); err != nil {
		return err
	}
	if err := w.WriteData(o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.WriteData(o.LastError); err != nil {
		return err
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	for i1 := range o.LocalPhoneNumber {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.LocalPhoneNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalPhoneNumber); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.Alternates != "" {
		_ptr_szAlternates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Alternates); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Alternates, _ptr_szAlternates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrDNS); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrDNSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrWINSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.NetProtocols); err != nil {
		return err
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25PadType {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.X25PadType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25PadType); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25Address {
		i1 := i1
		if uint64(i1) >= 201 {
			break
		}
		if err := w.WriteData(o.X25Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25Address); uint64(i1) < 201; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25Facilities {
		i1 := i1
		if uint64(i1) >= 201 {
			break
		}
		if err := w.WriteData(o.X25Facilities[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25Facilities); uint64(i1) < 201; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25UserData {
		i1 := i1
		if uint64(i1) >= 201 {
			break
		}
		if err := w.WriteData(o.X25UserData[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25UserData); uint64(i1) < 201; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Channels); err != nil {
		return err
	}
	if err := w.WriteData(o.SubEntries); err != nil {
		return err
	}
	if err := w.WriteData(o.DialMode); err != nil {
		return err
	}
	if err := w.WriteData(o.DialExtraPercent); err != nil {
		return err
	}
	if err := w.WriteData(o.DialExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.HangUpExtraPercent); err != nil {
		return err
	}
	if err := w.WriteData(o.HangUpExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleDisconnectSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionType); err != nil {
		return err
	}
	if err := w.WriteData(o.CustomAuthKey); err != nil {
		return err
	}
	if err := w.WriteData(o.CustomAuthDataSize); err != nil {
		return err
	}
	// XXX pointer to primitive type, default behavior is to write non-null pointer.
	// if this behavior is not desired, use goext_default_null([cond]) attribute.
	_ptr_lpbCustomAuthData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.CustomAuthData); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.CustomAuthData, _ptr_lpbCustomAuthData); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VPNStrategy); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *Interface2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	o.InterfaceName = make([]uint16, 257)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ConnectionState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastError); err != nil {
		return err
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	o.LocalPhoneNumber = make([]uint16, 129)
	for i1 := range o.LocalPhoneNumber {
		i1 := i1
		if err := w.ReadData(&o.LocalPhoneNumber[i1]); err != nil {
			return err
		}
	}
	_ptr_szAlternates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Alternates); err != nil {
			return err
		}
		return nil
	})
	_s_szAlternates := func(ptr interface{}) { o.Alternates = *ptr.(*string) }
	if err := w.ReadPointer(&o.Alternates, _s_szAlternates, _ptr_szAlternates); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrDNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrDNSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrWINSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetProtocols); err != nil {
		return err
	}
	o.DeviceType = make([]uint16, 17)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]uint16, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.X25PadType = make([]uint16, 33)
	for i1 := range o.X25PadType {
		i1 := i1
		if err := w.ReadData(&o.X25PadType[i1]); err != nil {
			return err
		}
	}
	o.X25Address = make([]uint16, 201)
	for i1 := range o.X25Address {
		i1 := i1
		if err := w.ReadData(&o.X25Address[i1]); err != nil {
			return err
		}
	}
	o.X25Facilities = make([]uint16, 201)
	for i1 := range o.X25Facilities {
		i1 := i1
		if err := w.ReadData(&o.X25Facilities[i1]); err != nil {
			return err
		}
	}
	o.X25UserData = make([]uint16, 201)
	for i1 := range o.X25UserData {
		i1 := i1
		if err := w.ReadData(&o.X25UserData[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Channels); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubEntries); err != nil {
		return err
	}
	if err := w.ReadData(&o.DialMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.DialExtraPercent); err != nil {
		return err
	}
	if err := w.ReadData(&o.DialExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.HangUpExtraPercent); err != nil {
		return err
	}
	if err := w.ReadData(&o.HangUpExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleDisconnectSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionType); err != nil {
		return err
	}
	if err := w.ReadData(&o.CustomAuthKey); err != nil {
		return err
	}
	if err := w.ReadData(&o.CustomAuthDataSize); err != nil {
		return err
	}
	_ptr_lpbCustomAuthData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.CustomAuthData); err != nil {
			return err
		}
		return nil
	})
	_s_lpbCustomAuthData := func(ptr interface{}) { o.CustomAuthData = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.CustomAuthData, _s_lpbCustomAuthData, _ptr_lpbCustomAuthData); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &dtyp.GUID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.VPNStrategy); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// Interface3 structure represents MPRI_INTERFACE_3 RPC structure.
type Interface3 struct {
	InterfaceName            []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface                uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled                  bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType            RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState          RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons    uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError                uint32                `idl:"name:dwLastError" json:"last_error"`
	Options                  uint32                `idl:"name:dwfOptions" json:"options"`
	LocalPhoneNumber         []uint16              `idl:"name:szLocalPhoneNumber" json:"local_phone_number"`
	Alternates               string                `idl:"name:szAlternates" json:"alternates"`
	IPAddr                   uint32                `idl:"name:ipaddr" json:"ip_addr"`
	IPAddrDNS                uint32                `idl:"name:ipaddrDns" json:"ip_addr_dns"`
	IPAddrDNSAlt             uint32                `idl:"name:ipaddrDnsAlt" json:"ip_addr_dns_alt"`
	IPAddrWINS               uint32                `idl:"name:ipaddrWins" json:"ip_addr_wins"`
	IPAddrWINSAlt            uint32                `idl:"name:ipaddrWinsAlt" json:"ip_addr_wins_alt"`
	NetProtocols             uint32                `idl:"name:dwfNetProtocols" json:"net_protocols"`
	DeviceType               []uint16              `idl:"name:szDeviceType" json:"device_type"`
	DeviceName               []uint16              `idl:"name:szDeviceName" json:"device_name"`
	X25PadType               []uint16              `idl:"name:szX25PadType" json:"x25_pad_type"`
	X25Address               []uint16              `idl:"name:szX25Address" json:"x25_address"`
	X25Facilities            []uint16              `idl:"name:szX25Facilities" json:"x25_facilities"`
	X25UserData              []uint16              `idl:"name:szX25UserData" json:"x25_user_data"`
	Channels                 uint32                `idl:"name:dwChannels" json:"channels"`
	SubEntries               uint32                `idl:"name:dwSubEntries" json:"sub_entries"`
	DialMode                 uint32                `idl:"name:dwDialMode" json:"dial_mode"`
	DialExtraPercent         uint32                `idl:"name:dwDialExtraPercent" json:"dial_extra_percent"`
	DialExtraSampleSeconds   uint32                `idl:"name:dwDialExtraSampleSeconds" json:"dial_extra_sample_seconds"`
	HangUpExtraPercent       uint32                `idl:"name:dwHangUpExtraPercent" json:"hang_up_extra_percent"`
	HangUpExtraSampleSeconds uint32                `idl:"name:dwHangUpExtraSampleSeconds" json:"hang_up_extra_sample_seconds"`
	IdleDisconnectSeconds    uint32                `idl:"name:dwIdleDisconnectSeconds" json:"idle_disconnect_seconds"`
	Type                     uint32                `idl:"name:dwType" json:"type"`
	EncryptionType           uint32                `idl:"name:dwEncryptionType" json:"encryption_type"`
	CustomAuthKey            uint32                `idl:"name:dwCustomAuthKey" json:"custom_auth_key"`
	CustomAuthDataSize       uint32                `idl:"name:dwCustomAuthDataSize" json:"custom_auth_data_size"`
	CustomAuthData           uint8                 `idl:"name:lpbCustomAuthData" json:"custom_auth_data"`
	ID                       *dtyp.GUID            `idl:"name:guidId" json:"id"`
	VPNStrategy              uint32                `idl:"name:dwVpnStrategy" json:"vpn_strategy"`
	AddressCount             uint32                `idl:"name:AddressCount" json:"address_count"`
	Ipv6addrDNS              *IN6Addr              `idl:"name:ipv6addrDns" json:"ipv6addr_dns"`
	Ipv6addrDNSAlt           *IN6Addr              `idl:"name:ipv6addrDnsAlt" json:"ipv6addr_dns_alt"`
	Ipv6addr                 *IN6Addr              `idl:"name:ipv6addr" json:"ipv6addr"`
}

func (o *Interface3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Interface3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Interface); err != nil {
		return err
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.InterfaceType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ConnectionState)); err != nil {
		return err
	}
	if err := w.WriteData(o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.WriteData(o.LastError); err != nil {
		return err
	}
	if err := w.WriteData(o.Options); err != nil {
		return err
	}
	for i1 := range o.LocalPhoneNumber {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.LocalPhoneNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalPhoneNumber); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.Alternates != "" {
		_ptr_szAlternates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Alternates); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Alternates, _ptr_szAlternates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrDNS); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrDNSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddrWINSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.NetProtocols); err != nil {
		return err
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25PadType {
		i1 := i1
		if uint64(i1) >= 33 {
			break
		}
		if err := w.WriteData(o.X25PadType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25PadType); uint64(i1) < 33; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25Address {
		i1 := i1
		if uint64(i1) >= 201 {
			break
		}
		if err := w.WriteData(o.X25Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25Address); uint64(i1) < 201; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25Facilities {
		i1 := i1
		if uint64(i1) >= 201 {
			break
		}
		if err := w.WriteData(o.X25Facilities[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25Facilities); uint64(i1) < 201; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.X25UserData {
		i1 := i1
		if uint64(i1) >= 201 {
			break
		}
		if err := w.WriteData(o.X25UserData[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.X25UserData); uint64(i1) < 201; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Channels); err != nil {
		return err
	}
	if err := w.WriteData(o.SubEntries); err != nil {
		return err
	}
	if err := w.WriteData(o.DialMode); err != nil {
		return err
	}
	if err := w.WriteData(o.DialExtraPercent); err != nil {
		return err
	}
	if err := w.WriteData(o.DialExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.HangUpExtraPercent); err != nil {
		return err
	}
	if err := w.WriteData(o.HangUpExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleDisconnectSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionType); err != nil {
		return err
	}
	if err := w.WriteData(o.CustomAuthKey); err != nil {
		return err
	}
	if err := w.WriteData(o.CustomAuthDataSize); err != nil {
		return err
	}
	// XXX pointer to primitive type, default behavior is to write non-null pointer.
	// if this behavior is not desired, use goext_default_null([cond]) attribute.
	_ptr_lpbCustomAuthData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.CustomAuthData); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.CustomAuthData, _ptr_lpbCustomAuthData); err != nil {
		return err
	}
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VPNStrategy); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressCount); err != nil {
		return err
	}
	if o.Ipv6addrDNS != nil {
		if err := o.Ipv6addrDNS.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IN6Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Ipv6addrDNSAlt != nil {
		if err := o.Ipv6addrDNSAlt.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IN6Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Ipv6addr != nil {
		_ptr_ipv6addr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Ipv6addr != nil {
				if err := o.Ipv6addr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IN6Addr{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Ipv6addr, _ptr_ipv6addr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Interface3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	o.InterfaceName = make([]uint16, 257)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadEnum((*uint16)(&o.InterfaceType)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ConnectionState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.UnReachabilityReasons); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastError); err != nil {
		return err
	}
	if err := w.ReadData(&o.Options); err != nil {
		return err
	}
	o.LocalPhoneNumber = make([]uint16, 129)
	for i1 := range o.LocalPhoneNumber {
		i1 := i1
		if err := w.ReadData(&o.LocalPhoneNumber[i1]); err != nil {
			return err
		}
	}
	_ptr_szAlternates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Alternates); err != nil {
			return err
		}
		return nil
	})
	_s_szAlternates := func(ptr interface{}) { o.Alternates = *ptr.(*string) }
	if err := w.ReadPointer(&o.Alternates, _s_szAlternates, _ptr_szAlternates); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrDNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrDNSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddrWINSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetProtocols); err != nil {
		return err
	}
	o.DeviceType = make([]uint16, 17)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]uint16, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.X25PadType = make([]uint16, 33)
	for i1 := range o.X25PadType {
		i1 := i1
		if err := w.ReadData(&o.X25PadType[i1]); err != nil {
			return err
		}
	}
	o.X25Address = make([]uint16, 201)
	for i1 := range o.X25Address {
		i1 := i1
		if err := w.ReadData(&o.X25Address[i1]); err != nil {
			return err
		}
	}
	o.X25Facilities = make([]uint16, 201)
	for i1 := range o.X25Facilities {
		i1 := i1
		if err := w.ReadData(&o.X25Facilities[i1]); err != nil {
			return err
		}
	}
	o.X25UserData = make([]uint16, 201)
	for i1 := range o.X25UserData {
		i1 := i1
		if err := w.ReadData(&o.X25UserData[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Channels); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubEntries); err != nil {
		return err
	}
	if err := w.ReadData(&o.DialMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.DialExtraPercent); err != nil {
		return err
	}
	if err := w.ReadData(&o.DialExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.HangUpExtraPercent); err != nil {
		return err
	}
	if err := w.ReadData(&o.HangUpExtraSampleSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleDisconnectSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionType); err != nil {
		return err
	}
	if err := w.ReadData(&o.CustomAuthKey); err != nil {
		return err
	}
	if err := w.ReadData(&o.CustomAuthDataSize); err != nil {
		return err
	}
	_ptr_lpbCustomAuthData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.CustomAuthData); err != nil {
			return err
		}
		return nil
	})
	_s_lpbCustomAuthData := func(ptr interface{}) { o.CustomAuthData = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.CustomAuthData, _s_lpbCustomAuthData, _ptr_lpbCustomAuthData); err != nil {
		return err
	}
	if o.ID == nil {
		o.ID = &dtyp.GUID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.VPNStrategy); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressCount); err != nil {
		return err
	}
	if o.Ipv6addrDNS == nil {
		o.Ipv6addrDNS = &IN6Addr{}
	}
	if err := o.Ipv6addrDNS.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Ipv6addrDNSAlt == nil {
		o.Ipv6addrDNSAlt = &IN6Addr{}
	}
	if err := o.Ipv6addrDNSAlt.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ipv6addr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Ipv6addr == nil {
			o.Ipv6addr = &IN6Addr{}
		}
		if err := o.Ipv6addr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ipv6addr := func(ptr interface{}) { o.Ipv6addr = *ptr.(**IN6Addr) }
	if err := w.ReadPointer(&o.Ipv6addr, _s_ipv6addr, _ptr_ipv6addr); err != nil {
		return err
	}
	return nil
}

// Device0 structure represents MPR_DEVICE_0 RPC structure.
type Device0 struct {
	DeviceType []uint16 `idl:"name:szDeviceType" json:"device_type"`
	DeviceName []uint16 `idl:"name:szDeviceName" json:"device_name"`
}

func (o *Device0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Device0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Device0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	o.DeviceType = make([]uint16, 17)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]uint16, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Device1 structure represents MPR_DEVICE_1 RPC structure.
type Device1 struct {
	DeviceType       []uint16 `idl:"name:szDeviceType" json:"device_type"`
	DeviceName       []uint16 `idl:"name:szDeviceName" json:"device_name"`
	LocalPhoneNumber []uint16 `idl:"name:szLocalPhoneNumber" json:"local_phone_number"`
	Alternates       string   `idl:"name:szAlternates" json:"alternates"`
}

func (o *Device1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Device1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.LocalPhoneNumber {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.LocalPhoneNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalPhoneNumber); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.Alternates != "" {
		_ptr_szAlternates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Alternates); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Alternates, _ptr_szAlternates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Device1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	o.DeviceType = make([]uint16, 17)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]uint16, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.LocalPhoneNumber = make([]uint16, 129)
	for i1 := range o.LocalPhoneNumber {
		i1 := i1
		if err := w.ReadData(&o.LocalPhoneNumber[i1]); err != nil {
			return err
		}
	}
	_ptr_szAlternates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Alternates); err != nil {
			return err
		}
		return nil
	})
	_s_szAlternates := func(ptr interface{}) { o.Alternates = *ptr.(*string) }
	if err := w.ReadPointer(&o.Alternates, _s_szAlternates, _ptr_szAlternates); err != nil {
		return err
	}
	return nil
}

// Credentialsex1 structure represents MPR_CREDENTIALSEX_1 RPC structure.
type Credentialsex1 struct {
	Size   uint32 `idl:"name:dwSize" json:"size"`
	Offset uint32 `idl:"name:dwOffset" json:"offset"`
	Data   []byte `idl:"name:bData" json:"data"`
}

func (o *Credentialsex1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Credentialsex1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *Credentialsex1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	o.Data = make([]byte, 1)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// InterfaceFilterInfo structure represents IFFILTER_INFO RPC structure.
type InterfaceFilterInfo struct {
	EnableFragChk bool `idl:"name:bEnableFragChk" json:"enable_frag_chk"`
}

func (o *InterfaceFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.EnableFragChk {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bEnableFragChk int32
	if err := w.ReadData(&_bEnableFragChk); err != nil {
		return err
	}
	o.EnableFragChk = _bEnableFragChk != 0
	return nil
}

// Filter0 structure represents MPR_FILTER_0 RPC structure.
type Filter0 struct {
	Enable bool `idl:"name:fEnable" json:"enable"`
}

func (o *Filter0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Filter0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.Enable {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Filter0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bEnable int32
	if err := w.ReadData(&_bEnable); err != nil {
		return err
	}
	o.Enable = _bEnable != 0
	return nil
}

// IPXGlobalInfo structure represents IPX_GLOBAL_INFO RPC structure.
type IPXGlobalInfo struct {
	RoutingTableHashSize uint32 `idl:"name:RoutingTableHashSize" json:"routing_table_hash_size"`
	EventLogMask         uint32 `idl:"name:EventLogMask" json:"event_log_mask"`
}

func (o *IPXGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RoutingTableHashSize); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogMask); err != nil {
		return err
	}
	return nil
}
func (o *IPXGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoutingTableHashSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogMask); err != nil {
		return err
	}
	return nil
}

// IPXInterfaceInfo structure represents IPX_IF_INFO RPC structure.
type IPXInterfaceInfo struct {
	AdministratorState uint32 `idl:"name:AdministratorState" json:"administrator_state"`
	NetBIOSAccept      uint32 `idl:"name:NetbiosAccept" json:"netbios_accept"`
	NetBIOSDeliver     uint32 `idl:"name:NetbiosDeliver" json:"netbios_deliver"`
}

func (o *IPXInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AdministratorState); err != nil {
		return err
	}
	if err := w.WriteData(o.NetBIOSAccept); err != nil {
		return err
	}
	if err := w.WriteData(o.NetBIOSDeliver); err != nil {
		return err
	}
	return nil
}
func (o *IPXInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdministratorState); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetBIOSAccept); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetBIOSDeliver); err != nil {
		return err
	}
	return nil
}

// IpxwanInterfaceInfo structure represents IPXWAN_IF_INFO RPC structure.
type IpxwanInterfaceInfo struct {
	Adminstate uint32 `idl:"name:Adminstate" json:"adminstate"`
}

func (o *IpxwanInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxwanInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Adminstate); err != nil {
		return err
	}
	return nil
}
func (o *IpxwanInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Adminstate); err != nil {
		return err
	}
	return nil
}

// IPXStaticRouteInfo structure represents IPX_STATIC_ROUTE_INFO RPC structure.
type IPXStaticRouteInfo struct {
	Field1            *IPXStaticRouteInfo_Field1 `idl:"name:" json:""`
	TickCount         uint16                     `idl:"name:TickCount" json:"tick_count"`
	HopCount          uint16                     `idl:"name:HopCount" json:"hop_count"`
	NextHopMACAddress []byte                     `idl:"name:NextHopMacAddress" json:"next_hop_mac_address"`
}

func (o *IPXStaticRouteInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXStaticRouteInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// FIXME unknown type
	if err := w.WriteData(o.TickCount); err != nil {
		return err
	}
	if err := w.WriteData(o.HopCount); err != nil {
		return err
	}
	for i1 := range o.NextHopMACAddress {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.NextHopMACAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NextHopMACAddress); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPXStaticRouteInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// FIXME: unknown type
	if err := w.ReadData(&o.TickCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.HopCount); err != nil {
		return err
	}
	o.NextHopMACAddress = make([]byte, 6)
	for i1 := range o.NextHopMACAddress {
		i1 := i1
		if err := w.ReadData(&o.NextHopMACAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

type IPXStaticRouteInfo_Field1 struct {
	DwordAlign uint32 `idl:"name:DwordAlign" json:"dword_align"`
	Network    []byte `idl:"name:Network" json:"network"`
}

// IPXStaticServiceInfo structure represents IPX_STATIC_SERVICE_INFO RPC structure.
type IPXStaticServiceInfo IPXServerEntry

func (o *IPXStaticServiceInfo) IPXServerEntry() *IPXServerEntry { return (*IPXServerEntry)(o) }

func (o *IPXStaticServiceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXStaticServiceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Network {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Network[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Network); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Node {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.Node[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Node); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Socket {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if err := w.WriteData(o.Socket[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Socket); uint64(i1) < 2; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HopCount); err != nil {
		return err
	}
	return nil
}
func (o *IPXStaticServiceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	o.Name = make([]byte, 48)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	o.Network = make([]byte, 4)
	for i1 := range o.Network {
		i1 := i1
		if err := w.ReadData(&o.Network[i1]); err != nil {
			return err
		}
	}
	o.Node = make([]byte, 6)
	for i1 := range o.Node {
		i1 := i1
		if err := w.ReadData(&o.Node[i1]); err != nil {
			return err
		}
	}
	o.Socket = make([]byte, 2)
	for i1 := range o.Socket {
		i1 := i1
		if err := w.ReadData(&o.Socket[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.HopCount); err != nil {
		return err
	}
	return nil
}

// IPXServerEntry structure represents IPX_SERVER_ENTRY RPC structure.
type IPXServerEntry struct {
	Type     uint16 `idl:"name:Type" json:"type"`
	Name     []byte `idl:"name:Name" json:"name"`
	Network  []byte `idl:"name:Network" json:"network"`
	Node     []byte `idl:"name:Node" json:"node"`
	Socket   []byte `idl:"name:Socket" json:"socket"`
	HopCount uint16 `idl:"name:HopCount" json:"hop_count"`
}

func (o *IPXServerEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXServerEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Network {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Network[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Network); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Node {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.Node[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Node); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Socket {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if err := w.WriteData(o.Socket[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Socket); uint64(i1) < 2; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HopCount); err != nil {
		return err
	}
	return nil
}
func (o *IPXServerEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	o.Name = make([]byte, 48)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	o.Network = make([]byte, 4)
	for i1 := range o.Network {
		i1 := i1
		if err := w.ReadData(&o.Network[i1]); err != nil {
			return err
		}
	}
	o.Node = make([]byte, 6)
	for i1 := range o.Node {
		i1 := i1
		if err := w.ReadData(&o.Node[i1]); err != nil {
			return err
		}
	}
	o.Socket = make([]byte, 2)
	for i1 := range o.Socket {
		i1 := i1
		if err := w.ReadData(&o.Socket[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.HopCount); err != nil {
		return err
	}
	return nil
}

// IPXStaticNetBIOSNameInfo structure represents IPX_STATIC_NETBIOS_NAME_INFO RPC structure.
type IPXStaticNetBIOSNameInfo struct {
	Field1 *IPXStaticNetBIOSNameInfo_Field1 `idl:"name:" json:""`
}

func (o *IPXStaticNetBIOSNameInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXStaticNetBIOSNameInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// FIXME unknown type
	return nil
}
func (o *IPXStaticNetBIOSNameInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// FIXME: unknown type
	return nil
}

type IPXStaticNetBIOSNameInfo_Field1 struct {
	DwordAlign uint32 `idl:"name:DwordAlign" json:"dword_align"`
	Name       []byte `idl:"name:Name" json:"name"`
}

// IPXAdapterInfo structure represents IPX_ADAPTER_INFO RPC structure.
type IPXAdapterInfo struct {
	PacketType  uint32   `idl:"name:PacketType" json:"packet_type"`
	AdapterName []uint16 `idl:"name:AdapterName" json:"adapter_name"`
}

func (o *IPXAdapterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXAdapterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketType); err != nil {
		return err
	}
	for i1 := range o.AdapterName {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.AdapterName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AdapterName); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPXAdapterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketType); err != nil {
		return err
	}
	o.AdapterName = make([]uint16, 48)
	for i1 := range o.AdapterName {
		i1 := i1
		if err := w.ReadData(&o.AdapterName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IPXTrafficFilterGlobalInfo structure represents IPX_TRAFFIC_FILTER_GLOBAL_INFO RPC structure.
type IPXTrafficFilterGlobalInfo struct {
	FilterAction uint32 `idl:"name:FilterAction" json:"filter_action"`
}

func (o *IPXTrafficFilterGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXTrafficFilterGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.FilterAction); err != nil {
		return err
	}
	return nil
}
func (o *IPXTrafficFilterGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.FilterAction); err != nil {
		return err
	}
	return nil
}

// IPXTrafficFilterInfo structure represents IPX_TRAFFIC_FILTER_INFO RPC structure.
type IPXTrafficFilterInfo struct {
	FilterDefinition       uint32 `idl:"name:FilterDefinition" json:"filter_definition"`
	DestinationNetwork     []byte `idl:"name:DestinationNetwork" json:"destination_network"`
	DestinationNetworkMask []byte `idl:"name:DestinationNetworkMask" json:"destination_network_mask"`
	DestinationNode        []byte `idl:"name:DestinationNode" json:"destination_node"`
	DestinationSocket      []byte `idl:"name:DestinationSocket" json:"destination_socket"`
	SourceNetwork          []byte `idl:"name:SourceNetwork" json:"source_network"`
	SourceNetworkMask      []byte `idl:"name:SourceNetworkMask" json:"source_network_mask"`
	SourceNode             []byte `idl:"name:SourceNode" json:"source_node"`
	SourceSocket           []byte `idl:"name:SourceSocket" json:"source_socket"`
	PacketType             uint8  `idl:"name:PacketType" json:"packet_type"`
}

func (o *IPXTrafficFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXTrafficFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.FilterDefinition); err != nil {
		return err
	}
	for i1 := range o.DestinationNetwork {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.DestinationNetwork[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DestinationNetwork); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DestinationNetworkMask {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.DestinationNetworkMask[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DestinationNetworkMask); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DestinationNode {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.DestinationNode[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DestinationNode); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DestinationSocket {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if err := w.WriteData(o.DestinationSocket[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DestinationSocket); uint64(i1) < 2; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.SourceNetwork {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.SourceNetwork[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SourceNetwork); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.SourceNetworkMask {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.SourceNetworkMask[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SourceNetworkMask); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.SourceNode {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.SourceNode[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SourceNode); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.SourceSocket {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if err := w.WriteData(o.SourceSocket[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SourceSocket); uint64(i1) < 2; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PacketType); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPXTrafficFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.FilterDefinition); err != nil {
		return err
	}
	o.DestinationNetwork = make([]byte, 4)
	for i1 := range o.DestinationNetwork {
		i1 := i1
		if err := w.ReadData(&o.DestinationNetwork[i1]); err != nil {
			return err
		}
	}
	o.DestinationNetworkMask = make([]byte, 4)
	for i1 := range o.DestinationNetworkMask {
		i1 := i1
		if err := w.ReadData(&o.DestinationNetworkMask[i1]); err != nil {
			return err
		}
	}
	o.DestinationNode = make([]byte, 6)
	for i1 := range o.DestinationNode {
		i1 := i1
		if err := w.ReadData(&o.DestinationNode[i1]); err != nil {
			return err
		}
	}
	o.DestinationSocket = make([]byte, 2)
	for i1 := range o.DestinationSocket {
		i1 := i1
		if err := w.ReadData(&o.DestinationSocket[i1]); err != nil {
			return err
		}
	}
	o.SourceNetwork = make([]byte, 4)
	for i1 := range o.SourceNetwork {
		i1 := i1
		if err := w.ReadData(&o.SourceNetwork[i1]); err != nil {
			return err
		}
	}
	o.SourceNetworkMask = make([]byte, 4)
	for i1 := range o.SourceNetworkMask {
		i1 := i1
		if err := w.ReadData(&o.SourceNetworkMask[i1]); err != nil {
			return err
		}
	}
	o.SourceNode = make([]byte, 6)
	for i1 := range o.SourceNode {
		i1 := i1
		if err := w.ReadData(&o.SourceNode[i1]); err != nil {
			return err
		}
	}
	o.SourceSocket = make([]byte, 2)
	for i1 := range o.SourceSocket {
		i1 := i1
		if err := w.ReadData(&o.SourceSocket[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PacketType); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// InterfaceTableIndex structure represents IF_TABLE_INDEX RPC structure.
type InterfaceTableIndex struct {
	InterfaceIndex uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
}

func (o *InterfaceTableIndex) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceTableIndex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceTableIndex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}

// RoutingTableIndex structure represents ROUTING_TABLE_INDEX RPC structure.
type RoutingTableIndex struct {
	Network []byte `idl:"name:Network" json:"network"`
}

func (o *RoutingTableIndex) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RoutingTableIndex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Network {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Network[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Network); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *RoutingTableIndex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Network = make([]byte, 4)
	for i1 := range o.Network {
		i1 := i1
		if err := w.ReadData(&o.Network[i1]); err != nil {
			return err
		}
	}
	return nil
}

// StaticRoutesTableIndex structure represents STATIC_ROUTES_TABLE_INDEX RPC structure.
type StaticRoutesTableIndex struct {
	InterfaceIndex uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
	Network        []byte `idl:"name:Network" json:"network"`
}

func (o *StaticRoutesTableIndex) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *StaticRoutesTableIndex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	for i1 := range o.Network {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Network[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Network); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *StaticRoutesTableIndex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	o.Network = make([]byte, 4)
	for i1 := range o.Network {
		i1 := i1
		if err := w.ReadData(&o.Network[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// ServicesTableIndex structure represents SERVICES_TABLE_INDEX RPC structure.
type ServicesTableIndex struct {
	ServiceType uint16 `idl:"name:ServiceType" json:"service_type"`
	ServiceName []byte `idl:"name:ServiceName" json:"service_name"`
}

func (o *ServicesTableIndex) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServicesTableIndex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.ServiceType); err != nil {
		return err
	}
	for i1 := range o.ServiceName {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.ServiceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ServiceName); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(2); err != nil {
		return err
	}
	return nil
}
func (o *ServicesTableIndex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServiceType); err != nil {
		return err
	}
	o.ServiceName = make([]byte, 48)
	for i1 := range o.ServiceName {
		i1 := i1
		if err := w.ReadData(&o.ServiceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(2); err != nil {
		return err
	}
	return nil
}

// StaticServicesTableIndex structure represents STATIC_SERVICES_TABLE_INDEX RPC structure.
type StaticServicesTableIndex struct {
	InterfaceIndex uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
	ServiceType    uint16 `idl:"name:ServiceType" json:"service_type"`
	ServiceName    []byte `idl:"name:ServiceName" json:"service_name"`
}

func (o *StaticServicesTableIndex) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *StaticServicesTableIndex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.ServiceType); err != nil {
		return err
	}
	for i1 := range o.ServiceName {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.ServiceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ServiceName); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *StaticServicesTableIndex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServiceType); err != nil {
		return err
	}
	o.ServiceName = make([]byte, 48)
	for i1 := range o.ServiceName {
		i1 := i1
		if err := w.ReadData(&o.ServiceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

type IPXMIBIndex struct {
	InterfaceTableIndex      *InterfaceTableIndex      `idl:"name:InterfaceTableIndex" json:"interface_table_index"`
	RoutingTableIndex        *RoutingTableIndex        `idl:"name:RoutingTableIndex" json:"routing_table_index"`
	StaticRoutesTableIndex   *StaticRoutesTableIndex   `idl:"name:StaticRoutesTableIndex" json:"static_routes_table_index"`
	ServicesTableIndex       *ServicesTableIndex       `idl:"name:ServicesTableIndex" json:"services_table_index"`
	StaticServicesTableIndex *StaticServicesTableIndex `idl:"name:StaticServicesTableIndex" json:"static_services_table_index"`
}

// IPXMIBGetInputData structure represents IPX_MIB_GET_INPUT_DATA RPC structure.
type IPXMIBGetInputData struct {
	TableID  uint32       `idl:"name:TableId" json:"table_id"`
	MIBIndex *IPXMIBIndex `idl:"name:MibIndex" json:"mib_index"`
}

func (o *IPXMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	// FIXME unknown type MibIndex
	return nil
}
func (o *IPXMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	// FIXME: unknown type MibIndex
	return nil
}

// IpxmibBase structure represents IPXMIB_BASE RPC structure.
type IpxmibBase struct {
	OperatorState    uint32 `idl:"name:OperState" json:"operator_state"`
	PrimaryNetNumber []byte `idl:"name:PrimaryNetNumber" json:"primary_net_number"`
	Node             []byte `idl:"name:Node" json:"node"`
	SystemName       []byte `idl:"name:SysName" json:"system_name"`
	MaxPathSplits    uint32 `idl:"name:MaxPathSplits" json:"max_path_splits"`
	InterfaceCount   uint32 `idl:"name:IfCount" json:"interface_count"`
	DestinationCount uint32 `idl:"name:DestCount" json:"destination_count"`
	ServCount        uint32 `idl:"name:ServCount" json:"serv_count"`
}

func (o *IpxmibBase) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxmibBase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.OperatorState); err != nil {
		return err
	}
	for i1 := range o.PrimaryNetNumber {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.PrimaryNetNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PrimaryNetNumber); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Node {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.Node[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Node); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.SystemName {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.SystemName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SystemName); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MaxPathSplits); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceCount); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ServCount); err != nil {
		return err
	}
	return nil
}
func (o *IpxmibBase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperatorState); err != nil {
		return err
	}
	o.PrimaryNetNumber = make([]byte, 4)
	for i1 := range o.PrimaryNetNumber {
		i1 := i1
		if err := w.ReadData(&o.PrimaryNetNumber[i1]); err != nil {
			return err
		}
	}
	o.Node = make([]byte, 6)
	for i1 := range o.Node {
		i1 := i1
		if err := w.ReadData(&o.Node[i1]); err != nil {
			return err
		}
	}
	o.SystemName = make([]byte, 48)
	for i1 := range o.SystemName {
		i1 := i1
		if err := w.ReadData(&o.SystemName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.MaxPathSplits); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServCount); err != nil {
		return err
	}
	return nil
}

// IPXInterfaceStats structure represents IPX_IF_STATS RPC structure.
type IPXInterfaceStats struct {
	InterfaceOperatorState uint32 `idl:"name:IfOperState" json:"interface_operator_state"`
	MaxPacketSize          uint32 `idl:"name:MaxPacketSize" json:"max_packet_size"`
	InHdrErrors            uint32 `idl:"name:InHdrErrors" json:"in_hdr_errors"`
	InFiltered             uint32 `idl:"name:InFiltered" json:"in_filtered"`
	InNoRoutes             uint32 `idl:"name:InNoRoutes" json:"in_no_routes"`
	InDiscards             uint32 `idl:"name:InDiscards" json:"in_discards"`
	InDelivers             uint32 `idl:"name:InDelivers" json:"in_delivers"`
	OutFiltered            uint32 `idl:"name:OutFiltered" json:"out_filtered"`
	OutDiscards            uint32 `idl:"name:OutDiscards" json:"out_discards"`
	OutDelivers            uint32 `idl:"name:OutDelivers" json:"out_delivers"`
	NetBIOSReceived        uint32 `idl:"name:NetbiosReceived" json:"netbios_received"`
	NetBIOSSent            uint32 `idl:"name:NetbiosSent" json:"netbios_sent"`
}

func (o *IPXInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceOperatorState); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxPacketSize); err != nil {
		return err
	}
	if err := w.WriteData(o.InHdrErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.InFiltered); err != nil {
		return err
	}
	if err := w.WriteData(o.InNoRoutes); err != nil {
		return err
	}
	if err := w.WriteData(o.InDiscards); err != nil {
		return err
	}
	if err := w.WriteData(o.InDelivers); err != nil {
		return err
	}
	if err := w.WriteData(o.OutFiltered); err != nil {
		return err
	}
	if err := w.WriteData(o.OutDiscards); err != nil {
		return err
	}
	if err := w.WriteData(o.OutDelivers); err != nil {
		return err
	}
	if err := w.WriteData(o.NetBIOSReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.NetBIOSSent); err != nil {
		return err
	}
	return nil
}
func (o *IPXInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceOperatorState); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxPacketSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.InHdrErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.InFiltered); err != nil {
		return err
	}
	if err := w.ReadData(&o.InNoRoutes); err != nil {
		return err
	}
	if err := w.ReadData(&o.InDiscards); err != nil {
		return err
	}
	if err := w.ReadData(&o.InDelivers); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutFiltered); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutDiscards); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutDelivers); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetBIOSReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetBIOSSent); err != nil {
		return err
	}
	return nil
}

// IPXInterface structure represents IPX_INTERFACE RPC structure.
type IPXInterface struct {
	InterfaceIndex          uint32             `idl:"name:InterfaceIndex" json:"interface_index"`
	AdministratorState      uint32             `idl:"name:AdministratorState" json:"administrator_state"`
	AdapterIndex            uint32             `idl:"name:AdapterIndex" json:"adapter_index"`
	InterfaceName           []byte             `idl:"name:InterfaceName" json:"interface_name"`
	InterfaceType           uint32             `idl:"name:InterfaceType" json:"interface_type"`
	MediaType               uint32             `idl:"name:MediaType" json:"media_type"`
	NetNumber               []byte             `idl:"name:NetNumber" json:"net_number"`
	MACAddress              []byte             `idl:"name:MacAddress" json:"mac_address"`
	Delay                   uint32             `idl:"name:Delay" json:"delay"`
	Throughput              uint32             `idl:"name:Throughput" json:"throughput"`
	NetBIOSAccept           uint32             `idl:"name:NetbiosAccept" json:"netbios_accept"`
	NetBIOSDeliver          uint32             `idl:"name:NetbiosDeliver" json:"netbios_deliver"`
	EnableIPXWANNegotiation uint32             `idl:"name:EnableIpxWanNegotiation" json:"enable_ipx_wan_negotiation"`
	InterfaceStats          *IPXInterfaceStats `idl:"name:IfStats" json:"interface_stats"`
}

func (o *IPXInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.AdministratorState); err != nil {
		return err
	}
	if err := w.WriteData(o.AdapterIndex); err != nil {
		return err
	}
	for i1 := range o.InterfaceName {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InterfaceName); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InterfaceType); err != nil {
		return err
	}
	if err := w.WriteData(o.MediaType); err != nil {
		return err
	}
	for i1 := range o.NetNumber {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.NetNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NetNumber); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.MACAddress {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.MACAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.MACAddress); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Delay); err != nil {
		return err
	}
	if err := w.WriteData(o.Throughput); err != nil {
		return err
	}
	if err := w.WriteData(o.NetBIOSAccept); err != nil {
		return err
	}
	if err := w.WriteData(o.NetBIOSDeliver); err != nil {
		return err
	}
	if err := w.WriteData(o.EnableIPXWANNegotiation); err != nil {
		return err
	}
	if o.InterfaceStats != nil {
		if err := o.InterfaceStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPXInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPXInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdministratorState); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdapterIndex); err != nil {
		return err
	}
	o.InterfaceName = make([]byte, 48)
	for i1 := range o.InterfaceName {
		i1 := i1
		if err := w.ReadData(&o.InterfaceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.InterfaceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.MediaType); err != nil {
		return err
	}
	o.NetNumber = make([]byte, 4)
	for i1 := range o.NetNumber {
		i1 := i1
		if err := w.ReadData(&o.NetNumber[i1]); err != nil {
			return err
		}
	}
	o.MACAddress = make([]byte, 6)
	for i1 := range o.MACAddress {
		i1 := i1
		if err := w.ReadData(&o.MACAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Delay); err != nil {
		return err
	}
	if err := w.ReadData(&o.Throughput); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetBIOSAccept); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetBIOSDeliver); err != nil {
		return err
	}
	if err := w.ReadData(&o.EnableIPXWANNegotiation); err != nil {
		return err
	}
	if o.InterfaceStats == nil {
		o.InterfaceStats = &IPXInterfaceStats{}
	}
	if err := o.InterfaceStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IPXRoute structure represents IPX_ROUTE RPC structure.
type IPXRoute struct {
	InterfaceIndex    uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
	Protocol          uint32 `idl:"name:Protocol" json:"protocol"`
	Network           []byte `idl:"name:Network" json:"network"`
	TickCount         uint16 `idl:"name:TickCount" json:"tick_count"`
	HopCount          uint16 `idl:"name:HopCount" json:"hop_count"`
	NextHopMACAddress []byte `idl:"name:NextHopMacAddress" json:"next_hop_mac_address"`
	Flags             uint32 `idl:"name:Flags" json:"flags"`
}

func (o *IPXRoute) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXRoute) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	for i1 := range o.Network {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Network[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Network); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TickCount); err != nil {
		return err
	}
	if err := w.WriteData(o.HopCount); err != nil {
		return err
	}
	for i1 := range o.NextHopMACAddress {
		i1 := i1
		if uint64(i1) >= 6 {
			break
		}
		if err := w.WriteData(o.NextHopMACAddress[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NextHopMACAddress); uint64(i1) < 6; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *IPXRoute) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	o.Network = make([]byte, 4)
	for i1 := range o.Network {
		i1 := i1
		if err := w.ReadData(&o.Network[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.TickCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.HopCount); err != nil {
		return err
	}
	o.NextHopMACAddress = make([]byte, 6)
	for i1 := range o.NextHopMACAddress {
		i1 := i1
		if err := w.ReadData(&o.NextHopMACAddress[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// IPXService structure represents IPX_SERVICE RPC structure.
type IPXService struct {
	InterfaceIndex uint32          `idl:"name:InterfaceIndex" json:"interface_index"`
	Protocol       uint32          `idl:"name:Protocol" json:"protocol"`
	Server         *IPXServerEntry `idl:"name:Server" json:"server"`
}

func (o *IPXService) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXService) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if o.Server != nil {
		if err := o.Server.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPXServerEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPXService) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if o.Server == nil {
		o.Server = &IPXServerEntry{}
	}
	if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

type IPXMIBRow struct {
	Interface *IPXInterface `idl:"name:Interface" json:"interface"`
	Route     *IPXRoute     `idl:"name:Route" json:"route"`
	Service   *IPXService   `idl:"name:Service" json:"service"`
}

// IPXMIBSetInputData structure represents IPX_MIB_SET_INPUT_DATA RPC structure.
type IPXMIBSetInputData struct {
	TableID uint32     `idl:"name:TableId" json:"table_id"`
	MIBRow  *IPXMIBRow `idl:"name:MibRow" json:"mib_row"`
}

func (o *IPXMIBSetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPXMIBSetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	// FIXME unknown type MibRow
	return nil
}
func (o *IPXMIBSetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	// FIXME: unknown type MibRow
	return nil
}

// SAPServiceFilterInfo structure represents SAP_SERVICE_FILTER_INFO RPC structure.
type SAPServiceFilterInfo struct {
	Field1      *SAPServiceFilterInfo_Field1 `idl:"name:" json:""`
	ServiceName []byte                       `idl:"name:ServiceName" json:"service_name"`
}

func (o *SAPServiceFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPServiceFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	// FIXME unknown type
	for i1 := range o.ServiceName {
		i1 := i1
		if uint64(i1) >= 48 {
			break
		}
		if err := w.WriteData(o.ServiceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ServiceName); uint64(i1) < 48; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(2); err != nil {
		return err
	}
	return nil
}
func (o *SAPServiceFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	// FIXME: unknown type
	o.ServiceName = make([]byte, 48)
	for i1 := range o.ServiceName {
		i1 := i1
		if err := w.ReadData(&o.ServiceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(2); err != nil {
		return err
	}
	return nil
}

type SAPServiceFilterInfo_Field1 struct {
	ServiceType      uint16 `idl:"name:ServiceType" json:"service_type"`
	ServiceTypeAlign uint32 `idl:"name:ServiceType_align" json:"service_type_align"`
}

// SAPInterfaceFilters structure represents SAP_IF_FILTERS RPC structure.
type SAPInterfaceFilters struct {
	SupplyFilterAction uint32                  `idl:"name:SupplyFilterAction" json:"supply_filter_action"`
	SupplyFilterCount  uint32                  `idl:"name:SupplyFilterCount" json:"supply_filter_count"`
	ListenFilterAction uint32                  `idl:"name:ListenFilterAction" json:"listen_filter_action"`
	ListenFilterCount  uint32                  `idl:"name:ListenFilterCount" json:"listen_filter_count"`
	ServiceFilter      []*SAPServiceFilterInfo `idl:"name:ServiceFilter" json:"service_filter"`
}

func (o *SAPInterfaceFilters) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterfaceFilters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SupplyFilterAction); err != nil {
		return err
	}
	if err := w.WriteData(o.SupplyFilterCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ListenFilterAction); err != nil {
		return err
	}
	if err := w.WriteData(o.ListenFilterCount); err != nil {
		return err
	}
	for i1 := range o.ServiceFilter {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.ServiceFilter[i1] != nil {
			if err := o.ServiceFilter[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SAPServiceFilterInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ServiceFilter); uint64(i1) < 1; i1++ {
		if err := (&SAPServiceFilterInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterfaceFilters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SupplyFilterAction); err != nil {
		return err
	}
	if err := w.ReadData(&o.SupplyFilterCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ListenFilterAction); err != nil {
		return err
	}
	if err := w.ReadData(&o.ListenFilterCount); err != nil {
		return err
	}
	o.ServiceFilter = make([]*SAPServiceFilterInfo, 1)
	for i1 := range o.ServiceFilter {
		i1 := i1
		if o.ServiceFilter[i1] == nil {
			o.ServiceFilter[i1] = &SAPServiceFilterInfo{}
		}
		if err := o.ServiceFilter[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// SAPInterfaceInfo structure represents SAP_IF_INFO RPC structure.
type SAPInterfaceInfo struct {
	AdminState             uint32 `idl:"name:AdminState" json:"admin_state"`
	UpdateMode             uint32 `idl:"name:UpdateMode" json:"update_mode"`
	PacketType             uint32 `idl:"name:PacketType" json:"packet_type"`
	Supply                 uint32 `idl:"name:Supply" json:"supply"`
	Listen                 uint32 `idl:"name:Listen" json:"listen"`
	GetNearestServerReply  uint32 `idl:"name:GetNearestServerReply" json:"get_nearest_server_reply"`
	PeriodicUpdateInterval uint32 `idl:"name:PeriodicUpdateInterval" json:"periodic_update_interval"`
	AgeIntervalMultiplier  uint32 `idl:"name:AgeIntervalMultiplier" json:"age_interval_multiplier"`
}

func (o *SAPInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminState); err != nil {
		return err
	}
	if err := w.WriteData(o.UpdateMode); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketType); err != nil {
		return err
	}
	if err := w.WriteData(o.Supply); err != nil {
		return err
	}
	if err := w.WriteData(o.Listen); err != nil {
		return err
	}
	if err := w.WriteData(o.GetNearestServerReply); err != nil {
		return err
	}
	if err := w.WriteData(o.PeriodicUpdateInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AgeIntervalMultiplier); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminState); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpdateMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Supply); err != nil {
		return err
	}
	if err := w.ReadData(&o.Listen); err != nil {
		return err
	}
	if err := w.ReadData(&o.GetNearestServerReply); err != nil {
		return err
	}
	if err := w.ReadData(&o.PeriodicUpdateInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AgeIntervalMultiplier); err != nil {
		return err
	}
	return nil
}

// SAPInterfaceConfig structure represents SAP_IF_CONFIG RPC structure.
type SAPInterfaceConfig struct {
	SAPInterfaceInfo    *SAPInterfaceInfo    `idl:"name:SapIfInfo" json:"sap_interface_info"`
	SAPInterfaceFilters *SAPInterfaceFilters `idl:"name:SapIfFilters" json:"sap_interface_filters"`
}

func (o *SAPInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.SAPInterfaceInfo != nil {
		if err := o.SAPInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAPInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SAPInterfaceFilters != nil {
		if err := o.SAPInterfaceFilters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAPInterfaceFilters{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAPInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.SAPInterfaceInfo == nil {
		o.SAPInterfaceInfo = &SAPInterfaceInfo{}
	}
	if err := o.SAPInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SAPInterfaceFilters == nil {
		o.SAPInterfaceFilters = &SAPInterfaceFilters{}
	}
	if err := o.SAPInterfaceFilters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAPMIBBase structure represents SAP_MIB_BASE RPC structure.
type SAPMIBBase struct {
	SAPOperatorState uint32 `idl:"name:SapOperState" json:"sap_operator_state"`
}

func (o *SAPMIBBase) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPMIBBase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SAPOperatorState); err != nil {
		return err
	}
	return nil
}
func (o *SAPMIBBase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SAPOperatorState); err != nil {
		return err
	}
	return nil
}

// SAPInterfaceStats structure represents SAP_IF_STATS RPC structure.
type SAPInterfaceStats struct {
	SAPInterfaceOperatorState uint32 `idl:"name:SapIfOperState" json:"sap_interface_operator_state"`
	SAPInterfaceInputPackets  uint32 `idl:"name:SapIfInputPackets" json:"sap_interface_input_packets"`
	SAPInterfaceOutputPackets uint32 `idl:"name:SapIfOutputPackets" json:"sap_interface_output_packets"`
}

func (o *SAPInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SAPInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.WriteData(o.SAPInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.SAPInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SAPInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.ReadData(&o.SAPInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.SAPInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}

// SAPInterface structure represents SAP_INTERFACE RPC structure.
type SAPInterface struct {
	InterfaceIndex    uint32             `idl:"name:InterfaceIndex" json:"interface_index"`
	SAPInterfaceInfo  *SAPInterfaceInfo  `idl:"name:SapIfInfo" json:"sap_interface_info"`
	SAPInterfaceStats *SAPInterfaceStats `idl:"name:SapIfStats" json:"sap_interface_stats"`
}

func (o *SAPInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if o.SAPInterfaceInfo != nil {
		if err := o.SAPInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAPInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SAPInterfaceStats != nil {
		if err := o.SAPInterfaceStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAPInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAPInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if o.SAPInterfaceInfo == nil {
		o.SAPInterfaceInfo = &SAPInterfaceInfo{}
	}
	if err := o.SAPInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SAPInterfaceStats == nil {
		o.SAPInterfaceStats = &SAPInterfaceStats{}
	}
	if err := o.SAPInterfaceStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAPMIBGetInputData structure represents SAP_MIB_GET_INPUT_DATA RPC structure.
type SAPMIBGetInputData struct {
	TableID        uint32 `idl:"name:TableId" json:"table_id"`
	InterfaceIndex uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
}

func (o *SAPMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *SAPMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}

// SAPMIBSetInputData structure represents SAP_MIB_SET_INPUT_DATA RPC structure.
type SAPMIBSetInputData struct {
	TableID      uint32        `idl:"name:TableId" json:"table_id"`
	SAPInterface *SAPInterface `idl:"name:SapInterface" json:"sap_interface"`
}

func (o *SAPMIBSetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPMIBSetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	if o.SAPInterface != nil {
		if err := o.SAPInterface.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAPInterface{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAPMIBSetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	if o.SAPInterface == nil {
		o.SAPInterface = &SAPInterface{}
	}
	if err := o.SAPInterface.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RIPMIBBase structure represents RIPMIB_BASE RPC structure.
type RIPMIBBase struct {
	RIPOperatorState uint32 `idl:"name:RIPOperState" json:"rip_operator_state"`
}

func (o *RIPMIBBase) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPMIBBase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RIPOperatorState); err != nil {
		return err
	}
	return nil
}
func (o *RIPMIBBase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RIPOperatorState); err != nil {
		return err
	}
	return nil
}

// RIPInterfaceStats structure represents RIP_IF_STATS RPC structure.
type RIPInterfaceStats struct {
	RIPInterfaceOperatorState uint32 `idl:"name:RipIfOperState" json:"rip_interface_operator_state"`
	RIPInterfaceInputPackets  uint32 `idl:"name:RipIfInputPackets" json:"rip_interface_input_packets"`
	RIPInterfaceOutputPackets uint32 `idl:"name:RipIfOutputPackets" json:"rip_interface_output_packets"`
}

func (o *RIPInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RIPInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.WriteData(o.RIPInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.RIPInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RIPInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.ReadData(&o.RIPInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.RIPInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}

// RIPInterfaceInfo structure represents RIP_IF_INFO RPC structure.
type RIPInterfaceInfo struct {
	AdminState             uint32 `idl:"name:AdminState" json:"admin_state"`
	UpdateMode             uint32 `idl:"name:UpdateMode" json:"update_mode"`
	PacketType             uint32 `idl:"name:PacketType" json:"packet_type"`
	Supply                 uint32 `idl:"name:Supply" json:"supply"`
	Listen                 uint32 `idl:"name:Listen" json:"listen"`
	PeriodicUpdateInterval uint32 `idl:"name:PeriodicUpdateInterval" json:"periodic_update_interval"`
	AgeIntervalMultiplier  uint32 `idl:"name:AgeIntervalMultiplier" json:"age_interval_multiplier"`
}

func (o *RIPInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminState); err != nil {
		return err
	}
	if err := w.WriteData(o.UpdateMode); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketType); err != nil {
		return err
	}
	if err := w.WriteData(o.Supply); err != nil {
		return err
	}
	if err := w.WriteData(o.Listen); err != nil {
		return err
	}
	if err := w.WriteData(o.PeriodicUpdateInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AgeIntervalMultiplier); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminState); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpdateMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Supply); err != nil {
		return err
	}
	if err := w.ReadData(&o.Listen); err != nil {
		return err
	}
	if err := w.ReadData(&o.PeriodicUpdateInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AgeIntervalMultiplier); err != nil {
		return err
	}
	return nil
}

// RIPInterface structure represents RIP_INTERFACE RPC structure.
type RIPInterface struct {
	InterfaceIndex    uint32             `idl:"name:InterfaceIndex" json:"interface_index"`
	RIPInterfaceInfo  *RIPInterfaceInfo  `idl:"name:RipIfInfo" json:"rip_interface_info"`
	RIPInterfaceStats *RIPInterfaceStats `idl:"name:RipIfStats" json:"rip_interface_stats"`
}

func (o *RIPInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if o.RIPInterfaceInfo != nil {
		if err := o.RIPInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RIPInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RIPInterfaceStats != nil {
		if err := o.RIPInterfaceStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RIPInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RIPInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if o.RIPInterfaceInfo == nil {
		o.RIPInterfaceInfo = &RIPInterfaceInfo{}
	}
	if err := o.RIPInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RIPInterfaceStats == nil {
		o.RIPInterfaceStats = &RIPInterfaceStats{}
	}
	if err := o.RIPInterfaceStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RIPMIBGetInputData structure represents RIP_MIB_GET_INPUT_DATA RPC structure.
type RIPMIBGetInputData struct {
	TableID        uint32 `idl:"name:TableId" json:"table_id"`
	InterfaceIndex uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
}

func (o *RIPMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *RIPMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}

// RIPMIBSetInputData structure represents RIP_MIB_SET_INPUT_DATA RPC structure.
type RIPMIBSetInputData struct {
	TableID      uint32        `idl:"name:TableId" json:"table_id"`
	RIPInterface *RIPInterface `idl:"name:RipInterface" json:"rip_interface"`
}

func (o *RIPMIBSetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPMIBSetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	if o.RIPInterface != nil {
		if err := o.RIPInterface.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RIPInterface{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RIPMIBSetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	if o.RIPInterface == nil {
		o.RIPInterface = &RIPInterface{}
	}
	if err := o.RIPInterface.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// EAPTLSHash structure represents EAPTLS_HASH RPC structure.
type EAPTLSHash struct {
	HashLength uint32 `idl:"name:cbHash" json:"hash_length"`
	Hash       []byte `idl:"name:pbHash" json:"hash"`
}

func (o *EAPTLSHash) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EAPTLSHash) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.HashLength); err != nil {
		return err
	}
	for i1 := range o.Hash {
		i1 := i1
		if uint64(i1) >= 20 {
			break
		}
		if err := w.WriteData(o.Hash[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Hash); uint64(i1) < 20; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *EAPTLSHash) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.HashLength); err != nil {
		return err
	}
	o.Hash = make([]byte, 20)
	for i1 := range o.Hash {
		i1 := i1
		if err := w.ReadData(&o.Hash[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// EAPTLSUserProperties structure represents EAPTLS_USER_PROPERTIES RPC structure.
type EAPTLSUserProperties struct {
	_             uint32      `idl:"name:reserved"`
	Version       uint32      `idl:"name:dwVersion" json:"version"`
	Size          uint32      `idl:"name:dwSize" json:"size"`
	Flags         uint32      `idl:"name:fFlags" json:"flags"`
	Hash          *EAPTLSHash `idl:"name:Hash" json:"hash"`
	DiffUser      string      `idl:"name:pwszDiffUser" json:"diff_user"`
	PINOffset     uint32      `idl:"name:dwPinOffset" json:"pin_offset"`
	PIN           string      `idl:"name:pwszPin" json:"pin"`
	Length        uint16      `idl:"name:usLength" json:"length"`
	MaximumLength uint16      `idl:"name:usMaximumLength" json:"maximum_length"`
	Seed          uint8       `idl:"name:ucSeed" json:"seed"`
	String        []uint16    `idl:"name:awszString" json:"string"`
}

func (o *EAPTLSUserProperties) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EAPTLSUserProperties) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	// reserved reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Hash != nil {
		if err := o.Hash.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EAPTLSHash{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DiffUser != "" {
		_ptr_pwszDiffUser := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.DiffUser); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DiffUser, _ptr_pwszDiffUser); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PINOffset); err != nil {
		return err
	}
	if o.PIN != "" {
		_ptr_pwszPin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.PIN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PIN, _ptr_pwszPin); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Seed); err != nil {
		return err
	}
	for i1 := range o.String {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.String[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.String); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *EAPTLSUserProperties) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	// reserved reserved
	var _reserved uint32
	if err := w.ReadData(&_reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.Hash == nil {
		o.Hash = &EAPTLSHash{}
	}
	if err := o.Hash.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pwszDiffUser := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.DiffUser); err != nil {
			return err
		}
		return nil
	})
	_s_pwszDiffUser := func(ptr interface{}) { o.DiffUser = *ptr.(*string) }
	if err := w.ReadPointer(&o.DiffUser, _s_pwszDiffUser, _ptr_pwszDiffUser); err != nil {
		return err
	}
	if err := w.ReadData(&o.PINOffset); err != nil {
		return err
	}
	_ptr_pwszPin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.PIN); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPin := func(ptr interface{}) { o.PIN = *ptr.(*string) }
	if err := w.ReadPointer(&o.PIN, _s_pwszPin, _ptr_pwszPin); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Seed); err != nil {
		return err
	}
	o.String = make([]uint16, 1)
	for i1 := range o.String {
		i1 := i1
		if err := w.ReadData(&o.String[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// IPBOOTPGlobalConfig structure represents IPBOOTP_GLOBAL_CONFIG RPC structure.
type IPBOOTPGlobalConfig struct {
	LoggingLevel     uint32 `idl:"name:GC_LoggingLevel" json:"logging_level"`
	MaxRecvQueueSize uint32 `idl:"name:GC_MaxRecvQueueSize" json:"max_recv_queue_size"`
	ServerCount      uint32 `idl:"name:GC_ServerCount" json:"server_count"`
}

func (o *IPBOOTPGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.ServerCount); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServerCount); err != nil {
		return err
	}
	return nil
}

// IPBOOTPInterfaceConfig structure represents IPBOOTP_IF_CONFIG RPC structure.
type IPBOOTPInterfaceConfig struct {
	State               uint32 `idl:"name:IC_State" json:"state"`
	RelayMode           uint32 `idl:"name:IC_RelayMode" json:"relay_mode"`
	MaxHopCount         uint32 `idl:"name:IC_MaxHopCount" json:"max_hop_count"`
	MinSecondsSinceBoot uint32 `idl:"name:IC_MinSecondsSinceBoot" json:"min_seconds_since_boot"`
}

func (o *IPBOOTPInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.RelayMode); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxHopCount); err != nil {
		return err
	}
	if err := w.WriteData(o.MinSecondsSinceBoot); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelayMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxHopCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinSecondsSinceBoot); err != nil {
		return err
	}
	return nil
}

// IPBOOTPMIBGetInputData structure represents IPBOOTP_MIB_GET_INPUT_DATA RPC structure.
type IPBOOTPMIBGetInputData struct {
	TypeID         uint32 `idl:"name:IMGID_TypeID" json:"type_id"`
	InterfaceIndex uint32 `idl:"name:IMGID_IfIndex" json:"interface_index"`
}

func (o *IPBOOTPMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}

// IPBOOTPMIBGetOutputData structure represents IPBOOTP_MIB_GET_OUTPUT_DATA RPC structure.
type IPBOOTPMIBGetOutputData struct {
	TypeID         uint32 `idl:"name:IMGOD_TypeID" json:"type_id"`
	InterfaceIndex uint32 `idl:"name:IMGOD_IfIndex" json:"interface_index"`
	Buffer         []byte `idl:"name:IMGOD_Buffer" json:"buffer"`
}

func (o *IPBOOTPMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IPBOOTPInterfaceStats structure represents IPBOOTP_IF_STATS RPC structure.
type IPBOOTPInterfaceStats struct {
	State             uint32 `idl:"name:IS_State" json:"state"`
	SendFailures      uint32 `idl:"name:IS_SendFailures" json:"send_failures"`
	ReceiveFailures   uint32 `idl:"name:IS_ReceiveFailures" json:"receive_failures"`
	ArpUpdateFailures uint32 `idl:"name:IS_ArpUpdateFailures" json:"arp_update_failures"`
	RequestsReceived  uint32 `idl:"name:IS_RequestsReceived" json:"requests_received"`
	RequestsDiscarded uint32 `idl:"name:IS_RequestsDiscarded" json:"requests_discarded"`
	RepliesReceived   uint32 `idl:"name:IS_RepliesReceived" json:"replies_received"`
	RepliesDiscarded  uint32 `idl:"name:IS_RepliesDiscarded" json:"replies_discarded"`
}

func (o *IPBOOTPInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.SendFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.ReceiveFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.ArpUpdateFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestsDiscarded); err != nil {
		return err
	}
	if err := w.WriteData(o.RepliesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.RepliesDiscarded); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.SendFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReceiveFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.ArpUpdateFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestsDiscarded); err != nil {
		return err
	}
	if err := w.ReadData(&o.RepliesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.RepliesDiscarded); err != nil {
		return err
	}
	return nil
}

// IPBOOTPInterfaceBinding structure represents IPBOOTP_IF_BINDING RPC structure.
type IPBOOTPInterfaceBinding struct {
	State     uint32 `idl:"name:IB_State" json:"state"`
	AddrCount uint32 `idl:"name:IB_AddrCount" json:"addr_count"`
}

func (o *IPBOOTPInterfaceBinding) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPInterfaceBinding) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrCount); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPInterfaceBinding) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrCount); err != nil {
		return err
	}
	return nil
}

// IPBOOTPIPAddress structure represents IPBOOTP_IP_ADDRESS RPC structure.
type IPBOOTPIPAddress struct {
	Address uint32 `idl:"name:IA_Address" json:"address"`
	Netmask uint32 `idl:"name:IA_Netmask" json:"netmask"`
}

func (o *IPBOOTPIPAddress) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPIPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Address); err != nil {
		return err
	}
	if err := w.WriteData(o.Netmask); err != nil {
		return err
	}
	return nil
}
func (o *IPBOOTPIPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Address); err != nil {
		return err
	}
	if err := w.ReadData(&o.Netmask); err != nil {
		return err
	}
	return nil
}

// DHCPV6RMIBGetOutputData structure represents DHCPV6R_MIB_GET_OUTPUT_DATA RPC structure.
type DHCPV6RMIBGetOutputData struct {
	TypeID         uint32 `idl:"name:IMGOD_TypeID" json:"type_id"`
	InterfaceIndex uint32 `idl:"name:IMGOD_IfIndex" json:"interface_index"`
	Buffer         []byte `idl:"name:IMGOD_Buffer" json:"buffer"`
}

func (o *DHCPV6RMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// DHCPV6RInterfaceStats structure represents DHCPV6R_IF_STATS RPC structure.
type DHCPV6RInterfaceStats struct {
	State             uint32 `idl:"name:IS_State" json:"state"`
	SendFailures      uint32 `idl:"name:IS_SendFailures" json:"send_failures"`
	ReceiveFailures   uint32 `idl:"name:IS_ReceiveFailures" json:"receive_failures"`
	RequestsReceived  uint32 `idl:"name:IS_RequestsReceived" json:"requests_received"`
	RequestsDiscarded uint32 `idl:"name:IS_RequestsDiscarded" json:"requests_discarded"`
	RepliesReceived   uint32 `idl:"name:IS_RepliesReceived" json:"replies_received"`
	RepliesDiscarded  uint32 `idl:"name:IS_RepliesDiscarded" json:"replies_discarded"`
}

func (o *DHCPV6RInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.SendFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.ReceiveFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestsDiscarded); err != nil {
		return err
	}
	if err := w.WriteData(o.RepliesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.RepliesDiscarded); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.SendFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReceiveFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestsDiscarded); err != nil {
		return err
	}
	if err := w.ReadData(&o.RepliesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.RepliesDiscarded); err != nil {
		return err
	}
	return nil
}

// DHCPV6RMIBGetInputData structure represents DHCPV6R_MIB_GET_INPUT_DATA RPC structure.
type DHCPV6RMIBGetInputData struct {
	TypeID         uint32 `idl:"name:IMGID_TypeID" json:"type_id"`
	InterfaceIndex uint32 `idl:"name:IMGID_IfIndex" json:"interface_index"`
}

func (o *DHCPV6RMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	return nil
}

// DHCPV6RGlobalConfig structure represents DHCPV6R_GLOBAL_CONFIG RPC structure.
type DHCPV6RGlobalConfig struct {
	LoggingLevel     uint32 `idl:"name:GC_LoggingLevel" json:"logging_level"`
	MaxRecvQueueSize uint32 `idl:"name:GC_MaxRecvQueueSize" json:"max_recv_queue_size"`
	ServerCount      uint32 `idl:"name:GC_ServerCount" json:"server_count"`
}

func (o *DHCPV6RGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.ServerCount); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServerCount); err != nil {
		return err
	}
	return nil
}

// DHCPV6RInterfaceConfig structure represents DHCPV6R_IF_CONFIG RPC structure.
type DHCPV6RInterfaceConfig struct {
	State          uint32 `idl:"name:IC_State" json:"state"`
	RelayMode      uint32 `idl:"name:IC_RelayMode" json:"relay_mode"`
	MaxHopCount    uint32 `idl:"name:IC_MaxHopCount" json:"max_hop_count"`
	MinElapsedTime uint32 `idl:"name:IC_MinElapsedTime" json:"min_elapsed_time"`
}

func (o *DHCPV6RInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.RelayMode); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxHopCount); err != nil {
		return err
	}
	if err := w.WriteData(o.MinElapsedTime); err != nil {
		return err
	}
	return nil
}
func (o *DHCPV6RInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelayMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxHopCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinElapsedTime); err != nil {
		return err
	}
	return nil
}

// IPRIPMIBGetInputData structure represents IPRIP_MIB_GET_INPUT_DATA RPC structure.
type IPRIPMIBGetInputData struct {
	TypeID uint32                       `idl:"name:IMGID_TypeID" json:"type_id"`
	Field2 *IPRIPMIBGetInputData_Field2 `idl:"name:" json:""`
}

func (o *IPRIPMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	// FIXME unknown type
	return nil
}
func (o *IPRIPMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	// FIXME: unknown type
	return nil
}

type IPRIPMIBGetInputData_Field2 struct {
	InterfaceIndex uint32 `idl:"name:IMGID_IfIndex" json:"interface_index"`
	PeerAddress    uint32 `idl:"name:IMGID_PeerAddress" json:"peer_address"`
}

// IPRIPMIBGetOutputData structure represents IPRIP_MIB_GET_OUTPUT_DATA RPC structure.
type IPRIPMIBGetOutputData struct {
	TypeID uint32                        `idl:"name:IMGOD_TypeID" json:"type_id"`
	Field2 *IPRIPMIBGetOutputData_Field2 `idl:"name:" json:""`
	Buffer []byte                        `idl:"name:IMGOD_Buffer" json:"buffer"`
}

func (o *IPRIPMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	// FIXME unknown type
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	// FIXME: unknown type
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

type IPRIPMIBGetOutputData_Field2 struct {
	InterfaceIndex uint32 `idl:"name:IMGOD_IfIndex" json:"interface_index"`
	PeerAddress    uint32 `idl:"name:IMGOD_PeerAddress" json:"peer_address"`
}

// IPRIPGlobalStats structure represents IPRIP_GLOBAL_STATS RPC structure.
type IPRIPGlobalStats struct {
	GsSystemRouteChanges uint32 `idl:"name:GS_SystemRouteChanges" json:"gs_system_route_changes"`
	GsTotalResponsesSent uint32 `idl:"name:GS_TotalResponsesSent" json:"gs_total_responses_sent"`
}

func (o *IPRIPGlobalStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPGlobalStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.GsSystemRouteChanges); err != nil {
		return err
	}
	if err := w.WriteData(o.GsTotalResponsesSent); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPGlobalStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GsSystemRouteChanges); err != nil {
		return err
	}
	if err := w.ReadData(&o.GsTotalResponsesSent); err != nil {
		return err
	}
	return nil
}

// IPRIPGlobalConfig structure represents IPRIP_GLOBAL_CONFIG RPC structure.
type IPRIPGlobalConfig struct {
	LoggingLevel               uint32 `idl:"name:GC_LoggingLevel" json:"logging_level"`
	MaxRecvQueueSize           uint32 `idl:"name:GC_MaxRecvQueueSize" json:"max_recv_queue_size"`
	MaxSendQueueSize           uint32 `idl:"name:GC_MaxSendQueueSize" json:"max_send_queue_size"`
	MinTriggeredUpdateInterval uint32 `idl:"name:GC_MinTriggeredUpdateInterval" json:"min_triggered_update_interval"`
	PeerFilterMode             uint32 `idl:"name:GC_PeerFilterMode" json:"peer_filter_mode"`
	PeerFilterCount            uint32 `idl:"name:GC_PeerFilterCount" json:"peer_filter_count"`
}

func (o *IPRIPGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxSendQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.MinTriggeredUpdateInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.PeerFilterMode); err != nil {
		return err
	}
	if err := w.WriteData(o.PeerFilterCount); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxSendQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinTriggeredUpdateInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.PeerFilterMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.PeerFilterCount); err != nil {
		return err
	}
	return nil
}

// IPRIPInterfaceStats structure represents IPRIP_IF_STATS RPC structure.
type IPRIPInterfaceStats struct {
	State                      uint32 `idl:"name:IS_State" json:"state"`
	SendFailures               uint32 `idl:"name:IS_SendFailures" json:"send_failures"`
	ReceiveFailures            uint32 `idl:"name:IS_ReceiveFailures" json:"receive_failures"`
	RequestsSent               uint32 `idl:"name:IS_RequestsSent" json:"requests_sent"`
	RequestsReceived           uint32 `idl:"name:IS_RequestsReceived" json:"requests_received"`
	ResponsesSent              uint32 `idl:"name:IS_ResponsesSent" json:"responses_sent"`
	ResponsesReceived          uint32 `idl:"name:IS_ResponsesReceived" json:"responses_received"`
	BadResponsePacketsReceived uint32 `idl:"name:IS_BadResponsePacketsReceived" json:"bad_response_packets_received"`
	BadResponseEntriesReceived uint32 `idl:"name:IS_BadResponseEntriesReceived" json:"bad_response_entries_received"`
	TriggeredUpdatesSent       uint32 `idl:"name:IS_TriggeredUpdatesSent" json:"triggered_updates_sent"`
}

func (o *IPRIPInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.SendFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.ReceiveFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestsSent); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponsesSent); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponsesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.BadResponsePacketsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.BadResponseEntriesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.TriggeredUpdatesSent); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.SendFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReceiveFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestsSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponsesSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponsesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadResponsePacketsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadResponseEntriesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.TriggeredUpdatesSent); err != nil {
		return err
	}
	return nil
}

// IPRIPInterfaceConfig structure represents IPRIP_IF_CONFIG RPC structure.
type IPRIPInterfaceConfig struct {
	State                   uint32 `idl:"name:IC_State" json:"state"`
	Metric                  uint32 `idl:"name:IC_Metric" json:"metric"`
	UpdateMode              uint32 `idl:"name:IC_UpdateMode" json:"update_mode"`
	AcceptMode              uint32 `idl:"name:IC_AcceptMode" json:"accept_mode"`
	AnnounceMode            uint32 `idl:"name:IC_AnnounceMode" json:"announce_mode"`
	ProtocolFlags           uint32 `idl:"name:IC_ProtocolFlags" json:"protocol_flags"`
	RouteExpirationInterval uint32 `idl:"name:IC_RouteExpirationInterval" json:"route_expiration_interval"`
	RouteRemovalInterval    uint32 `idl:"name:IC_RouteRemovalInterval" json:"route_removal_interval"`
	FullUpdateInterval      uint32 `idl:"name:IC_FullUpdateInterval" json:"full_update_interval"`
	AuthenticationType      uint32 `idl:"name:IC_AuthenticationType" json:"authentication_type"`
	AuthenticationKey       []byte `idl:"name:IC_AuthenticationKey" json:"authentication_key"`
	RouteTag                uint16 `idl:"name:IC_RouteTag" json:"route_tag"`
	UnicastPeerMode         uint32 `idl:"name:IC_UnicastPeerMode" json:"unicast_peer_mode"`
	AcceptFilterMode        uint32 `idl:"name:IC_AcceptFilterMode" json:"accept_filter_mode"`
	AnnounceFilterMode      uint32 `idl:"name:IC_AnnounceFilterMode" json:"announce_filter_mode"`
	UnicastPeerCount        uint32 `idl:"name:IC_UnicastPeerCount" json:"unicast_peer_count"`
	AcceptFilterCount       uint32 `idl:"name:IC_AcceptFilterCount" json:"accept_filter_count"`
	AnnounceFilterCount     uint32 `idl:"name:IC_AnnounceFilterCount" json:"announce_filter_count"`
}

func (o *IPRIPInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Metric); err != nil {
		return err
	}
	if err := w.WriteData(o.UpdateMode); err != nil {
		return err
	}
	if err := w.WriteData(o.AcceptMode); err != nil {
		return err
	}
	if err := w.WriteData(o.AnnounceMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ProtocolFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteExpirationInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.RouteRemovalInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.FullUpdateInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationType); err != nil {
		return err
	}
	for i1 := range o.AuthenticationKey {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.AuthenticationKey[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AuthenticationKey); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RouteTag); err != nil {
		return err
	}
	if err := w.WriteData(o.UnicastPeerMode); err != nil {
		return err
	}
	if err := w.WriteData(o.AcceptFilterMode); err != nil {
		return err
	}
	if err := w.WriteData(o.AnnounceFilterMode); err != nil {
		return err
	}
	if err := w.WriteData(o.UnicastPeerCount); err != nil {
		return err
	}
	if err := w.WriteData(o.AcceptFilterCount); err != nil {
		return err
	}
	if err := w.WriteData(o.AnnounceFilterCount); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Metric); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpdateMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.AcceptMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.AnnounceMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtocolFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteExpirationInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouteRemovalInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.FullUpdateInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationType); err != nil {
		return err
	}
	o.AuthenticationKey = make([]byte, 16)
	for i1 := range o.AuthenticationKey {
		i1 := i1
		if err := w.ReadData(&o.AuthenticationKey[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.RouteTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.UnicastPeerMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.AcceptFilterMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.AnnounceFilterMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.UnicastPeerCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.AcceptFilterCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.AnnounceFilterCount); err != nil {
		return err
	}
	return nil
}

// IPRIPRouteFilter structure represents IPRIP_ROUTE_FILTER RPC structure.
type IPRIPRouteFilter struct {
	RfLoAddress uint32 `idl:"name:RF_LoAddress" json:"rf_lo_address"`
	RfHiAddress uint32 `idl:"name:RF_HiAddress" json:"rf_hi_address"`
}

func (o *IPRIPRouteFilter) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPRouteFilter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RfLoAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.RfHiAddress); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPRouteFilter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RfLoAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.RfHiAddress); err != nil {
		return err
	}
	return nil
}

// IPRIPInterfaceBinding structure represents IPRIP_IF_BINDING RPC structure.
type IPRIPInterfaceBinding struct {
	State     uint32 `idl:"name:IB_State" json:"state"`
	AddrCount uint32 `idl:"name:IB_AddrCount" json:"addr_count"`
}

func (o *IPRIPInterfaceBinding) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPInterfaceBinding) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrCount); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPInterfaceBinding) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrCount); err != nil {
		return err
	}
	return nil
}

// IPRIPIPAddress structure represents IPRIP_IP_ADDRESS RPC structure.
type IPRIPIPAddress struct {
	Address uint32 `idl:"name:IA_Address" json:"address"`
	Netmask uint32 `idl:"name:IA_Netmask" json:"netmask"`
}

func (o *IPRIPIPAddress) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPIPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Address); err != nil {
		return err
	}
	if err := w.WriteData(o.Netmask); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPIPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Address); err != nil {
		return err
	}
	if err := w.ReadData(&o.Netmask); err != nil {
		return err
	}
	return nil
}

// IPRIPPeerStats structure represents IPRIP_PEER_STATS RPC structure.
type IPRIPPeerStats struct {
	PsLastPeerRouteTag           uint32 `idl:"name:PS_LastPeerRouteTag" json:"ps_last_peer_route_tag"`
	PsLastPeerUpdateTickCount    uint32 `idl:"name:PS_LastPeerUpdateTickCount" json:"ps_last_peer_update_tick_count"`
	PsLastPeerUpdateVersion      uint32 `idl:"name:PS_LastPeerUpdateVersion" json:"ps_last_peer_update_version"`
	PsBadResponsePacketsFromPeer uint32 `idl:"name:PS_BadResponsePacketsFromPeer" json:"ps_bad_response_packets_from_peer"`
	PsBadResponseEntriesFromPeer uint32 `idl:"name:PS_BadResponseEntriesFromPeer" json:"ps_bad_response_entries_from_peer"`
}

func (o *IPRIPPeerStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPPeerStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PsLastPeerRouteTag); err != nil {
		return err
	}
	if err := w.WriteData(o.PsLastPeerUpdateTickCount); err != nil {
		return err
	}
	if err := w.WriteData(o.PsLastPeerUpdateVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.PsBadResponsePacketsFromPeer); err != nil {
		return err
	}
	if err := w.WriteData(o.PsBadResponseEntriesFromPeer); err != nil {
		return err
	}
	return nil
}
func (o *IPRIPPeerStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PsLastPeerRouteTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.PsLastPeerUpdateTickCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.PsLastPeerUpdateVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.PsBadResponsePacketsFromPeer); err != nil {
		return err
	}
	if err := w.ReadData(&o.PsBadResponseEntriesFromPeer); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGetInputData structure represents IGMP_MIB_GET_INPUT_DATA RPC structure.
type IGMPMIBGetInputData struct {
	TypeID         uint32 `idl:"name:TypeId" json:"type_id"`
	Flags          uint16 `idl:"name:Flags" json:"flags"`
	Signature      uint16 `idl:"name:Signature" json:"signature"`
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	RASClientAddr  uint32 `idl:"name:RasClientAddr" json:"ras_client_addr"`
	GroupAddr      uint32 `idl:"name:GroupAddr" json:"group_addr"`
	Count          uint32 `idl:"name:Count" json:"count"`
}

func (o *IGMPMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Signature); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.RASClientAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Signature); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.RASClientAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGetOutputData structure represents IGMP_MIB_GET_OUTPUT_DATA RPC structure.
type IGMPMIBGetOutputData struct {
	TypeID uint32 `idl:"name:TypeId" json:"type_id"`
	Flags  uint32 `idl:"name:Flags" json:"flags"`
	Count  uint32 `idl:"name:Count" json:"count"`
	Buffer []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *IGMPMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGlobalConfig structure represents IGMP_MIB_GLOBAL_CONFIG RPC structure.
type IGMPMIBGlobalConfig struct {
	Version        uint32 `idl:"name:Version" json:"version"`
	LoggingLevel   uint32 `idl:"name:LoggingLevel" json:"logging_level"`
	RASClientStats uint32 `idl:"name:RasClientStats" json:"ras_client_stats"`
}

func (o *IGMPMIBGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.RASClientStats); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.RASClientStats); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGlobalStats structure represents IGMP_MIB_GLOBAL_STATS RPC structure.
type IGMPMIBGlobalStats struct {
	CurrentGroupMemberships uint32 `idl:"name:CurrentGroupMemberships" json:"current_group_memberships"`
	GroupMembershipsAdded   uint32 `idl:"name:GroupMembershipsAdded" json:"group_memberships_added"`
}

func (o *IGMPMIBGlobalStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGlobalStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentGroupMemberships); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupMembershipsAdded); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGlobalStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentGroupMemberships); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupMembershipsAdded); err != nil {
		return err
	}
	return nil
}

// IGMPMIBInterfaceBinding structure represents IGMP_MIB_IF_BINDING RPC structure.
type IGMPMIBInterfaceBinding struct {
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	InterfaceType  uint32 `idl:"name:IfType" json:"interface_type"`
	State          uint32 `idl:"name:State" json:"state"`
	AddrCount      uint32 `idl:"name:AddrCount" json:"addr_count"`
}

func (o *IGMPMIBInterfaceBinding) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceBinding) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceType); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrCount); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceBinding) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrCount); err != nil {
		return err
	}
	return nil
}

// IGMPMIBInterfaceConfig structure represents IGMP_MIB_IF_CONFIG RPC structure.
type IGMPMIBInterfaceConfig struct {
	Version                     uint32 `idl:"name:Version" json:"version"`
	InterfaceIndex              uint32 `idl:"name:IfIndex" json:"interface_index"`
	IPAddr                      uint32 `idl:"name:IpAddr" json:"ip_addr"`
	InterfaceType               uint32 `idl:"name:IfType" json:"interface_type"`
	Flags                       uint32 `idl:"name:Flags" json:"flags"`
	IGMPProtocolType            uint32 `idl:"name:IgmpProtocolType" json:"igmp_protocol_type"`
	RobustnessVariable          uint32 `idl:"name:RobustnessVariable" json:"robustness_variable"`
	StartupQueryInterval        uint32 `idl:"name:StartupQueryInterval" json:"startup_query_interval"`
	StartupQueryCount           uint32 `idl:"name:StartupQueryCount" json:"startup_query_count"`
	GenQueryInterval            uint32 `idl:"name:GenQueryInterval" json:"gen_query_interval"`
	GenQueryMaxResponseTime     uint32 `idl:"name:GenQueryMaxResponseTime" json:"gen_query_max_response_time"`
	LastMemQueryInterval        uint32 `idl:"name:LastMemQueryInterval" json:"last_mem_query_interval"`
	LastMemQueryCount           uint32 `idl:"name:LastMemQueryCount" json:"last_mem_query_count"`
	OtherQuerierPresentInterval uint32 `idl:"name:OtherQuerierPresentInterval" json:"other_querier_present_interval"`
	GroupMembershipTimeout      uint32 `idl:"name:GroupMembershipTimeout" json:"group_membership_timeout"`
	StaticGroupsLength          uint32 `idl:"name:NumStaticGroups" json:"static_groups_length"`
}

func (o *IGMPMIBInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceType); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.IGMPProtocolType); err != nil {
		return err
	}
	if err := w.WriteData(o.RobustnessVariable); err != nil {
		return err
	}
	if err := w.WriteData(o.StartupQueryInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.StartupQueryCount); err != nil {
		return err
	}
	if err := w.WriteData(o.GenQueryInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.GenQueryMaxResponseTime); err != nil {
		return err
	}
	if err := w.WriteData(o.LastMemQueryInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.LastMemQueryCount); err != nil {
		return err
	}
	if err := w.WriteData(o.OtherQuerierPresentInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupMembershipTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.StaticGroupsLength); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.IGMPProtocolType); err != nil {
		return err
	}
	if err := w.ReadData(&o.RobustnessVariable); err != nil {
		return err
	}
	if err := w.ReadData(&o.StartupQueryInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.StartupQueryCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.GenQueryInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.GenQueryMaxResponseTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastMemQueryInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastMemQueryCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.OtherQuerierPresentInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupMembershipTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.StaticGroupsLength); err != nil {
		return err
	}
	return nil
}

// IGMPMIBInterfaceGroupsList structure represents IGMP_MIB_IF_GROUPS_LIST RPC structure.
type IGMPMIBInterfaceGroupsList struct {
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	IPAddr         uint32 `idl:"name:IpAddr" json:"ip_addr"`
	InterfaceType  uint32 `idl:"name:IfType" json:"interface_type"`
	GroupsLength   uint32 `idl:"name:NumGroups" json:"groups_length"`
	Buffer         []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *IGMPMIBInterfaceGroupsList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceGroupsList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceType); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupsLength); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceGroupsList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupsLength); err != nil {
		return err
	}
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGroupInfo structure represents IGMP_MIB_GROUP_INFO RPC structure.
type IGMPMIBGroupInfo struct {
	Field1                *IGMPMIBGroupInfo_Field1 `idl:"name:" json:""`
	IPAddr                uint32                   `idl:"name:IpAddr" json:"ip_addr"`
	GroupUpTime           uint32                   `idl:"name:GroupUpTime" json:"group_up_time"`
	GroupExpiryTime       uint32                   `idl:"name:GroupExpiryTime" json:"group_expiry_time"`
	LastReporter          uint32                   `idl:"name:LastReporter" json:"last_reporter"`
	V1HostPresentTimeLeft uint32                   `idl:"name:V1HostPresentTimeLeft" json:"v1_host_present_time_left"`
	Flags                 uint32                   `idl:"name:Flags" json:"flags"`
}

func (o *IGMPMIBGroupInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// FIXME unknown type
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupUpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupExpiryTime); err != nil {
		return err
	}
	if err := w.WriteData(o.LastReporter); err != nil {
		return err
	}
	if err := w.WriteData(o.V1HostPresentTimeLeft); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// FIXME: unknown type
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupUpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupExpiryTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastReporter); err != nil {
		return err
	}
	if err := w.ReadData(&o.V1HostPresentTimeLeft); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

type IGMPMIBGroupInfo_Field1 struct {
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	GroupAddr      uint32 `idl:"name:GroupAddr" json:"group_addr"`
}

// IGMPMIBInterfaceStats structure represents IGMP_MIB_IF_STATS RPC structure.
type IGMPMIBInterfaceStats struct {
	InterfaceIndex            uint32 `idl:"name:IfIndex" json:"interface_index"`
	IPAddr                    uint32 `idl:"name:IpAddr" json:"ip_addr"`
	InterfaceType             uint32 `idl:"name:IfType" json:"interface_type"`
	State                     uint8  `idl:"name:State" json:"state"`
	QuerierState              uint8  `idl:"name:QuerierState" json:"querier_state"`
	IGMPProtocolType          uint32 `idl:"name:IgmpProtocolType" json:"igmp_protocol_type"`
	QuerierIPAddr             uint32 `idl:"name:QuerierIpAddr" json:"querier_ip_addr"`
	ProxyInterfaceIndex       uint32 `idl:"name:ProxyIfIndex" json:"proxy_interface_index"`
	QuerierPresentTimeLeft    uint32 `idl:"name:QuerierPresentTimeLeft" json:"querier_present_time_left"`
	LastQuerierChangeTime     uint32 `idl:"name:LastQuerierChangeTime" json:"last_querier_change_time"`
	V1QuerierPresentTimeLeft  uint32 `idl:"name:V1QuerierPresentTimeLeft" json:"v1_querier_present_time_left"`
	UpTime                    uint32 `idl:"name:Uptime" json:"up_time"`
	TotalIGMPPacketsReceived  uint32 `idl:"name:TotalIgmpPacketsReceived" json:"total_igmp_packets_received"`
	TotalIGMPPacketsForRouter uint32 `idl:"name:TotalIgmpPacketsForRouter" json:"total_igmp_packets_for_router"`
	GeneralQueriesReceived    uint32 `idl:"name:GeneralQueriesReceived" json:"general_queries_received"`
	WrongVersionQueries       uint32 `idl:"name:WrongVersionQueries" json:"wrong_version_queries"`
	JoinsReceived             uint32 `idl:"name:JoinsReceived" json:"joins_received"`
	LeavesReceived            uint32 `idl:"name:LeavesReceived" json:"leaves_received"`
	CurrentGroupMemberships   uint32 `idl:"name:CurrentGroupMemberships" json:"current_group_memberships"`
	GroupMembershipsAdded     uint32 `idl:"name:GroupMembershipsAdded" json:"group_memberships_added"`
	WrongChecksumPackets      uint32 `idl:"name:WrongChecksumPackets" json:"wrong_checksum_packets"`
	ShortPacketsReceived      uint32 `idl:"name:ShortPacketsReceived" json:"short_packets_received"`
	LongPacketsReceived       uint32 `idl:"name:LongPacketsReceived" json:"long_packets_received"`
	PacketsWithoutRouterAlert uint32 `idl:"name:PacketsWithoutRtrAlert" json:"packets_without_router_alert"`
}

func (o *IGMPMIBInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceType); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.QuerierState); err != nil {
		return err
	}
	if err := w.WriteData(o.IGMPProtocolType); err != nil {
		return err
	}
	if err := w.WriteData(o.QuerierIPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.ProxyInterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.QuerierPresentTimeLeft); err != nil {
		return err
	}
	if err := w.WriteData(o.LastQuerierChangeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.V1QuerierPresentTimeLeft); err != nil {
		return err
	}
	if err := w.WriteData(o.UpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalIGMPPacketsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalIGMPPacketsForRouter); err != nil {
		return err
	}
	if err := w.WriteData(o.GeneralQueriesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.WrongVersionQueries); err != nil {
		return err
	}
	if err := w.WriteData(o.JoinsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.LeavesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentGroupMemberships); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupMembershipsAdded); err != nil {
		return err
	}
	if err := w.WriteData(o.WrongChecksumPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.ShortPacketsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.LongPacketsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketsWithoutRouterAlert); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuerierState); err != nil {
		return err
	}
	if err := w.ReadData(&o.IGMPProtocolType); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuerierIPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProxyInterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuerierPresentTimeLeft); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastQuerierChangeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.V1QuerierPresentTimeLeft); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalIGMPPacketsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalIGMPPacketsForRouter); err != nil {
		return err
	}
	if err := w.ReadData(&o.GeneralQueriesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.WrongVersionQueries); err != nil {
		return err
	}
	if err := w.ReadData(&o.JoinsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.LeavesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentGroupMemberships); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupMembershipsAdded); err != nil {
		return err
	}
	if err := w.ReadData(&o.WrongChecksumPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.ShortPacketsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.LongPacketsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketsWithoutRouterAlert); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGroupInterfacesList structure represents IGMP_MIB_GROUP_IFS_LIST RPC structure.
type IGMPMIBGroupInterfacesList struct {
	GroupAddr        uint32 `idl:"name:GroupAddr" json:"group_addr"`
	InterfacesLength uint32 `idl:"name:NumInterfaces" json:"interfaces_length"`
	Buffer           []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *IGMPMIBGroupInterfacesList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupInterfacesList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfacesLength); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupInterfacesList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfacesLength); err != nil {
		return err
	}
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGroupSourceInfoV3 structure represents IGMP_MIB_GROUP_SOURCE_INFO_V3 RPC structure.
type IGMPMIBGroupSourceInfoV3 struct {
	Source           uint32 `idl:"name:Source" json:"source"`
	SourceExpiryTime uint32 `idl:"name:SourceExpiryTime" json:"source_expiry_time"`
	SourceUpTime     uint32 `idl:"name:SourceUpTime" json:"source_up_time"`
	Flags            uint32 `idl:"name:Flags" json:"flags"`
}

func (o *IGMPMIBGroupSourceInfoV3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupSourceInfoV3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Source); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceExpiryTime); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceUpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupSourceInfoV3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Source); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceExpiryTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceUpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// IGMPMIBGroupInfoV3 structure represents IGMP_MIB_GROUP_INFO_V3 RPC structure.
type IGMPMIBGroupInfoV3 struct {
	Field1                *IGMPMIBGroupInfoV3_Field1 `idl:"name:" json:""`
	IPAddr                uint32                     `idl:"name:IpAddr" json:"ip_addr"`
	GroupUpTime           uint32                     `idl:"name:GroupUpTime" json:"group_up_time"`
	GroupExpiryTime       uint32                     `idl:"name:GroupExpiryTime" json:"group_expiry_time"`
	LastReporter          uint32                     `idl:"name:LastReporter" json:"last_reporter"`
	V1HostPresentTimeLeft uint32                     `idl:"name:V1HostPresentTimeLeft" json:"v1_host_present_time_left"`
	Flags                 uint32                     `idl:"name:Flags" json:"flags"`
	Version               uint32                     `idl:"name:Version" json:"version"`
	Size                  uint32                     `idl:"name:Size" json:"size"`
	FilterType            uint32                     `idl:"name:FilterType" json:"filter_type"`
	V2HostPresentTimeLeft uint32                     `idl:"name:V2HostPresentTimeLeft" json:"v2_host_present_time_left"`
	SourcesLength         uint32                     `idl:"name:NumSources" json:"sources_length"`
}

func (o *IGMPMIBGroupInfoV3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupInfoV3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// FIXME unknown type
	if err := w.WriteData(o.IPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupUpTime); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupExpiryTime); err != nil {
		return err
	}
	if err := w.WriteData(o.LastReporter); err != nil {
		return err
	}
	if err := w.WriteData(o.V1HostPresentTimeLeft); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.FilterType); err != nil {
		return err
	}
	if err := w.WriteData(o.V2HostPresentTimeLeft); err != nil {
		return err
	}
	if err := w.WriteData(o.SourcesLength); err != nil {
		return err
	}
	return nil
}
func (o *IGMPMIBGroupInfoV3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// FIXME: unknown type
	if err := w.ReadData(&o.IPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupUpTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupExpiryTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastReporter); err != nil {
		return err
	}
	if err := w.ReadData(&o.V1HostPresentTimeLeft); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.FilterType); err != nil {
		return err
	}
	if err := w.ReadData(&o.V2HostPresentTimeLeft); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourcesLength); err != nil {
		return err
	}
	return nil
}

type IGMPMIBGroupInfoV3_Field1 struct {
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	GroupAddr      uint32 `idl:"name:GroupAddr" json:"group_addr"`
}

// InterfaceRouteEntry structure represents INTERFACE_ROUTE_ENTRY RPC structure.
type InterfaceRouteEntry struct {
	Index     uint32              `idl:"name:dwIndex" json:"index"`
	RouteInfo *InterfaceRouteInfo `idl:"name:routeInfo" json:"route_info"`
}

func (o *InterfaceRouteEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceRouteEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if o.RouteInfo != nil {
		if err := o.RouteInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceRouteInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceRouteEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if o.RouteInfo == nil {
		o.RouteInfo = &InterfaceRouteInfo{}
	}
	if err := o.RouteInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IPNATMIBQuery structure represents IP_NAT_MIB_QUERY RPC structure.
type IPNATMIBQuery struct {
	OID    uint32                `idl:"name:Oid" json:"oid"`
	Field2 *IPNATMIBQuery_Field2 `idl:"name:" json:""`
}

func (o *IPNATMIBQuery) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATMIBQuery) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.OID); err != nil {
		return err
	}
	// FIXME unknown type
	return nil
}
func (o *IPNATMIBQuery) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.OID); err != nil {
		return err
	}
	// FIXME: unknown type
	return nil
}

type IPNATMIBQuery_Field2 struct {
	Index uint32 `idl:"name:Index" json:"index"`
	Data  uint8  `idl:"name:Data" json:"data"`
}

// IPNATDirection type represents IP_NAT_DIRECTION RPC enumeration.
type IPNATDirection uint16

var (
	IpnatDirectionNATInboundDirection  IPNATDirection = 0
	IpnatDirectionNATOutboundDirection IPNATDirection = 1
)

func (o IPNATDirection) String() string {
	switch o {
	case IpnatDirectionNATInboundDirection:
		return "IpnatDirectionNATInboundDirection"
	case IpnatDirectionNATOutboundDirection:
		return "IpnatDirectionNATOutboundDirection"
	}
	return "Invalid"
}

// IPNATSessionMapping structure represents IP_NAT_SESSION_MAPPING RPC structure.
type IPNATSessionMapping struct {
	Protocol       uint8          `idl:"name:Protocol" json:"protocol"`
	PrivateAddress uint32         `idl:"name:PrivateAddress" json:"private_address"`
	PrivatePort    uint16         `idl:"name:PrivatePort" json:"private_port"`
	PublicAddress  uint32         `idl:"name:PublicAddress" json:"public_address"`
	PublicPort     uint16         `idl:"name:PublicPort" json:"public_port"`
	RemoteAddress  uint32         `idl:"name:RemoteAddress" json:"remote_address"`
	RemotePort     uint16         `idl:"name:RemotePort" json:"remote_port"`
	Direction      IPNATDirection `idl:"name:Direction" json:"direction"`
	IdleTime       uint32         `idl:"name:IdleTime" json:"idle_time"`
}

func (o *IPNATSessionMapping) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATSessionMapping) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivateAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivatePort); err != nil {
		return err
	}
	if err := w.WriteData(o.PublicAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.PublicPort); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.RemotePort); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleTime); err != nil {
		return err
	}
	return nil
}
func (o *IPNATSessionMapping) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivateAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivatePort); err != nil {
		return err
	}
	if err := w.ReadData(&o.PublicAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.PublicPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemotePort); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleTime); err != nil {
		return err
	}
	return nil
}

// IPNATEnumerateSessionMappings structure represents IP_NAT_ENUMERATE_SESSION_MAPPINGS RPC structure.
type IPNATEnumerateSessionMappings struct {
	Index              uint32                 `idl:"name:Index" json:"index"`
	EnumerateContext   []uint32               `idl:"name:EnumerateContext" json:"enumerate_context"`
	EnumerateCount     uint32                 `idl:"name:EnumerateCount" json:"enumerate_count"`
	EnumerateTotalHint uint32                 `idl:"name:EnumerateTotalHint" json:"enumerate_total_hint"`
	EnumerateTable     []*IPNATSessionMapping `idl:"name:EnumerateTable" json:"enumerate_table"`
}

func (o *IPNATEnumerateSessionMappings) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATEnumerateSessionMappings) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	for i1 := range o.EnumerateContext {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.EnumerateContext[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.EnumerateContext); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EnumerateCount); err != nil {
		return err
	}
	if err := w.WriteData(o.EnumerateTotalHint); err != nil {
		return err
	}
	for i1 := range o.EnumerateTable {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.EnumerateTable[i1] != nil {
			if err := o.EnumerateTable[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&IPNATSessionMapping{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.EnumerateTable); uint64(i1) < 1; i1++ {
		if err := (&IPNATSessionMapping{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPNATEnumerateSessionMappings) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	o.EnumerateContext = make([]uint32, 4)
	for i1 := range o.EnumerateContext {
		i1 := i1
		if err := w.ReadData(&o.EnumerateContext[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.EnumerateCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.EnumerateTotalHint); err != nil {
		return err
	}
	o.EnumerateTable = make([]*IPNATSessionMapping, 1)
	for i1 := range o.EnumerateTable {
		i1 := i1
		if o.EnumerateTable[i1] == nil {
			o.EnumerateTable[i1] = &IPNATSessionMapping{}
		}
		if err := o.EnumerateTable[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// IPNATInterfaceStatistics structure represents IP_NAT_INTERFACE_STATISTICS RPC structure.
type IPNATInterfaceStatistics struct {
	TotalMappings   uint32 `idl:"name:TotalMappings" json:"total_mappings"`
	InboundMappings uint32 `idl:"name:InboundMappings" json:"inbound_mappings"`
	BytesForward    uint64 `idl:"name:BytesForward" json:"bytes_forward"`
	BytesReverse    uint64 `idl:"name:BytesReverse" json:"bytes_reverse"`
	PacketsForward  uint64 `idl:"name:PacketsForward" json:"packets_forward"`
	PacketsReverse  uint64 `idl:"name:PacketsReverse" json:"packets_reverse"`
	RejectsForward  uint64 `idl:"name:RejectsForward" json:"rejects_forward"`
	RejectsReverse  uint64 `idl:"name:RejectsReverse" json:"rejects_reverse"`
}

func (o *IPNATInterfaceStatistics) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATInterfaceStatistics) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalMappings); err != nil {
		return err
	}
	if err := w.WriteData(o.InboundMappings); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesForward); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesReverse); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketsForward); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketsReverse); err != nil {
		return err
	}
	if err := w.WriteData(o.RejectsForward); err != nil {
		return err
	}
	if err := w.WriteData(o.RejectsReverse); err != nil {
		return err
	}
	return nil
}
func (o *IPNATInterfaceStatistics) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalMappings); err != nil {
		return err
	}
	if err := w.ReadData(&o.InboundMappings); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesForward); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesReverse); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketsForward); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketsReverse); err != nil {
		return err
	}
	if err := w.ReadData(&o.RejectsForward); err != nil {
		return err
	}
	if err := w.ReadData(&o.RejectsReverse); err != nil {
		return err
	}
	return nil
}

// IPDNSProxyMIBQuery structure represents IP_DNS_PROXY_MIB_QUERY RPC structure.
type IPDNSProxyMIBQuery struct {
	OID    uint32                     `idl:"name:Oid" json:"oid"`
	Field2 *IPDNSProxyMIBQuery_Field2 `idl:"name:" json:""`
}

func (o *IPDNSProxyMIBQuery) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPDNSProxyMIBQuery) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.OID); err != nil {
		return err
	}
	// FIXME unknown type
	return nil
}
func (o *IPDNSProxyMIBQuery) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.OID); err != nil {
		return err
	}
	// FIXME: unknown type
	return nil
}

type IPDNSProxyMIBQuery_Field2 struct {
	Index uint32 `idl:"name:Index" json:"index"`
	Data  uint8  `idl:"name:Data" json:"data"`
}

// IPDNSProxyStatistics structure represents IP_DNS_PROXY_STATISTICS RPC structure.
type IPDNSProxyStatistics struct {
	MessagesIgnored   uint32 `idl:"name:MessagesIgnored" json:"messages_ignored"`
	QueriesReceived   uint32 `idl:"name:QueriesReceived" json:"queries_received"`
	ResponsesReceived uint32 `idl:"name:ResponsesReceived" json:"responses_received"`
	QueriesSent       uint32 `idl:"name:QueriesSent" json:"queries_sent"`
	ResponsesSent     uint32 `idl:"name:ResponsesSent" json:"responses_sent"`
}

func (o *IPDNSProxyStatistics) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPDNSProxyStatistics) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MessagesIgnored); err != nil {
		return err
	}
	if err := w.WriteData(o.QueriesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponsesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.QueriesSent); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponsesSent); err != nil {
		return err
	}
	return nil
}
func (o *IPDNSProxyStatistics) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessagesIgnored); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueriesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponsesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueriesSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponsesSent); err != nil {
		return err
	}
	return nil
}

// IPAutoDHCPMIBQuery structure represents IP_AUTO_DHCP_MIB_QUERY RPC structure.
type IPAutoDHCPMIBQuery struct {
	OID    uint32                     `idl:"name:Oid" json:"oid"`
	Field2 *IPAutoDHCPMIBQuery_Field2 `idl:"name:" json:""`
	_      uint32                     `idl:"name:Reserved"`
}

func (o *IPAutoDHCPMIBQuery) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPMIBQuery) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.OID); err != nil {
		return err
	}
	// FIXME unknown type
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPMIBQuery) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.OID); err != nil {
		return err
	}
	// FIXME: unknown type
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	return nil
}

type IPAutoDHCPMIBQuery_Field2 struct {
	Index uint32 `idl:"name:Index" json:"index"`
	Data  uint8  `idl:"name:Data" json:"data"`
}

// IPAutoDHCPStatistics structure represents IP_AUTO_DHCP_STATISTICS RPC structure.
type IPAutoDHCPStatistics struct {
	MessagesIgnored   uint32 `idl:"name:MessagesIgnored" json:"messages_ignored"`
	BOOTPOffersSent   uint32 `idl:"name:BootpOffersSent" json:"bootp_offers_sent"`
	DiscoversReceived uint32 `idl:"name:DiscoversReceived" json:"discovers_received"`
	InformsReceived   uint32 `idl:"name:InformsReceived" json:"informs_received"`
	OffersSent        uint32 `idl:"name:OffersSent" json:"offers_sent"`
	RequestsReceived  uint32 `idl:"name:RequestsReceived" json:"requests_received"`
	ACKsSent          uint32 `idl:"name:AcksSent" json:"acks_sent"`
	NAKsSent          uint32 `idl:"name:NaksSent" json:"naks_sent"`
	DeclinesReceived  uint32 `idl:"name:DeclinesReceived" json:"declines_received"`
	ReleasesReceived  uint32 `idl:"name:ReleasesReceived" json:"releases_received"`
}

func (o *IPAutoDHCPStatistics) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPStatistics) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MessagesIgnored); err != nil {
		return err
	}
	if err := w.WriteData(o.BOOTPOffersSent); err != nil {
		return err
	}
	if err := w.WriteData(o.DiscoversReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.InformsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.OffersSent); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.ACKsSent); err != nil {
		return err
	}
	if err := w.WriteData(o.NAKsSent); err != nil {
		return err
	}
	if err := w.WriteData(o.DeclinesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.ReleasesReceived); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPStatistics) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessagesIgnored); err != nil {
		return err
	}
	if err := w.ReadData(&o.BOOTPOffersSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.DiscoversReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.InformsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.OffersSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACKsSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.NAKsSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeclinesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReleasesReceived); err != nil {
		return err
	}
	return nil
}

// MIBDaMessage structure represents MIB_DA_MSG RPC structure.
type MIBDaMessage struct {
	OperationCode uint32   `idl:"name:op_code" json:"operation_code"`
	ReturnCode    uint32   `idl:"name:ret_code" json:"return_code"`
	InSNMPID      []uint32 `idl:"name:in_snmp_id" json:"in_snmp_id"`
	ObjectID      []uint32 `idl:"name:obj_id" json:"object_id"`
	AttributeID   uint32   `idl:"name:attr_id" json:"attribute_id"`
	InstanceID    []uint32 `idl:"name:inst_id" json:"instance_id"`
	NextSNMPID    []uint32 `idl:"name:next_snmp_id" json:"next_snmp_id"`
	Creator       uint32   `idl:"name:creator" json:"creator"`
	AttributeType uint32   `idl:"name:attr_type" json:"attribute_type"`
	InstanceCnt   uint32   `idl:"name:inst_cnt" json:"instance_cnt"`
	MapFlag       uint32   `idl:"name:map_flag" json:"map_flag"`
	Data          []uint64 `idl:"name:data" json:"data"`
}

func (o *MIBDaMessage) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBDaMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationCode); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	for i1 := range o.InSNMPID {
		i1 := i1
		if uint64(i1) >= 44 {
			break
		}
		if err := w.WriteData(o.InSNMPID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InSNMPID); uint64(i1) < 44; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	for i1 := range o.ObjectID {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.ObjectID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ObjectID); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AttributeID); err != nil {
		return err
	}
	for i1 := range o.InstanceID {
		i1 := i1
		if uint64(i1) >= 23 {
			break
		}
		if err := w.WriteData(o.InstanceID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InstanceID); uint64(i1) < 23; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	for i1 := range o.NextSNMPID {
		i1 := i1
		if uint64(i1) >= 44 {
			break
		}
		if err := w.WriteData(o.NextSNMPID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NextSNMPID); uint64(i1) < 44; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Creator); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeType); err != nil {
		return err
	}
	if err := w.WriteData(o.InstanceCnt); err != nil {
		return err
	}
	if err := w.WriteData(o.MapFlag); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(ndr.Uint3264(o.Data[i1])); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint64(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBDaMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	o.InSNMPID = make([]uint32, 44)
	for i1 := range o.InSNMPID {
		i1 := i1
		if err := w.ReadData(&o.InSNMPID[i1]); err != nil {
			return err
		}
	}
	o.ObjectID = make([]uint32, 17)
	for i1 := range o.ObjectID {
		i1 := i1
		if err := w.ReadData(&o.ObjectID[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.AttributeID); err != nil {
		return err
	}
	o.InstanceID = make([]uint32, 23)
	for i1 := range o.InstanceID {
		i1 := i1
		if err := w.ReadData(&o.InstanceID[i1]); err != nil {
			return err
		}
	}
	o.NextSNMPID = make([]uint32, 44)
	for i1 := range o.NextSNMPID {
		i1 := i1
		if err := w.ReadData(&o.NextSNMPID[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Creator); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeType); err != nil {
		return err
	}
	if err := w.ReadData(&o.InstanceCnt); err != nil {
		return err
	}
	if err := w.ReadData(&o.MapFlag); err != nil {
		return err
	}
	o.Data = make([]uint64, 32)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData((*ndr.Uint3264)(&o.Data[i1])); err != nil {
			return err
		}
	}
	return nil
}

// IPAutoDHCPGlobalInfo structure represents IP_AUTO_DHCP_GLOBAL_INFO RPC structure.
type IPAutoDHCPGlobalInfo struct {
	LoggingLevel   uint32 `idl:"name:LoggingLevel" json:"logging_level"`
	Flags          uint32 `idl:"name:Flags" json:"flags"`
	LeaseTime      uint32 `idl:"name:LeaseTime" json:"lease_time"`
	ScopeNetwork   uint32 `idl:"name:ScopeNetwork" json:"scope_network"`
	ScopeMask      uint32 `idl:"name:ScopeMask" json:"scope_mask"`
	ExclusionCount uint32 `idl:"name:ExclusionCount" json:"exclusion_count"`
	ExclusionArray uint32 `idl:"name:ExclusionArray" json:"exclusion_array"`
}

func (o *IPAutoDHCPGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.LeaseTime); err != nil {
		return err
	}
	if err := w.WriteData(o.ScopeNetwork); err != nil {
		return err
	}
	if err := w.WriteData(o.ScopeMask); err != nil {
		return err
	}
	if err := w.WriteData(o.ExclusionCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ExclusionArray); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.LeaseTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScopeNetwork); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScopeMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExclusionCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExclusionArray); err != nil {
		return err
	}
	return nil
}

// IPAutoDHCPInterfaceInfo structure represents IP_AUTO_DHCP_INTERFACE_INFO RPC structure.
type IPAutoDHCPInterfaceInfo struct {
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *IPAutoDHCPInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *IPAutoDHCPInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// IPDNSProxyGlobalInfo structure represents IP_DNS_PROXY_GLOBAL_INFO RPC structure.
type IPDNSProxyGlobalInfo struct {
	LoggingLevel   uint32 `idl:"name:LoggingLevel" json:"logging_level"`
	Flags          uint32 `idl:"name:Flags" json:"flags"`
	TimeoutSeconds uint32 `idl:"name:TimeoutSeconds" json:"timeout_seconds"`
}

func (o *IPDNSProxyGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPDNSProxyGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutSeconds); err != nil {
		return err
	}
	return nil
}
func (o *IPDNSProxyGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutSeconds); err != nil {
		return err
	}
	return nil
}

// IPDNSProxyInterfaceInfo structure represents IP_DNS_PROXY_INTERFACE_INFO RPC structure.
type IPDNSProxyInterfaceInfo struct {
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *IPDNSProxyInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPDNSProxyInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *IPDNSProxyInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// IPNATGlobalInfo structure represents IP_NAT_GLOBAL_INFO RPC structure.
type IPNATGlobalInfo struct {
	LoggingLevel uint32                 `idl:"name:LoggingLevel" json:"logging_level"`
	Flags        uint32                 `idl:"name:Flags" json:"flags"`
	Header       *RouterInfoBlockHeader `idl:"name:Header" json:"header"`
}

func (o *IPNATGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RouterInfoBlockHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPNATGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &RouterInfoBlockHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IPNATTimeout structure represents IP_NAT_TIMEOUT RPC structure.
type IPNATTimeout struct {
	TCPTimeoutSeconds uint32 `idl:"name:TCPTimeoutSeconds" json:"tcp_timeout_seconds"`
	UDPTimeoutSeconds uint32 `idl:"name:UDPTimeoutSeconds" json:"udp_timeout_seconds"`
}

func (o *IPNATTimeout) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATTimeout) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TCPTimeoutSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.UDPTimeoutSeconds); err != nil {
		return err
	}
	return nil
}
func (o *IPNATTimeout) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TCPTimeoutSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.UDPTimeoutSeconds); err != nil {
		return err
	}
	return nil
}

// IPNATInterfaceInfo structure represents IP_NAT_INTERFACE_INFO RPC structure.
type IPNATInterfaceInfo struct {
	Index  uint32                 `idl:"name:Index" json:"index"`
	Flags  uint32                 `idl:"name:Flags" json:"flags"`
	Header *RouterInfoBlockHeader `idl:"name:Header" json:"header"`
}

func (o *IPNATInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RouterInfoBlockHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPNATInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &RouterInfoBlockHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IPNATAddressRange structure represents IP_NAT_ADDRESS_RANGE RPC structure.
type IPNATAddressRange struct {
	StartAddress uint32 `idl:"name:StartAddress" json:"start_address"`
	EndAddress   uint32 `idl:"name:EndAddress" json:"end_address"`
	SubnetMask   uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
}

func (o *IPNATAddressRange) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATAddressRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StartAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.EndAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	return nil
}
func (o *IPNATAddressRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StartAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.EndAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	return nil
}

// IPNATPortMapping structure represents IP_NAT_PORT_MAPPING RPC structure.
type IPNATPortMapping struct {
	Protocol       uint8  `idl:"name:Protocol" json:"protocol"`
	PublicPort     uint16 `idl:"name:PublicPort" json:"public_port"`
	PublicAddress  uint32 `idl:"name:PublicAddress" json:"public_address"`
	PrivatePort    uint16 `idl:"name:PrivatePort" json:"private_port"`
	PrivateAddress uint32 `idl:"name:PrivateAddress" json:"private_address"`
}

func (o *IPNATPortMapping) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATPortMapping) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Protocol); err != nil {
		return err
	}
	if err := w.WriteData(o.PublicPort); err != nil {
		return err
	}
	if err := w.WriteData(o.PublicAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivatePort); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivateAddress); err != nil {
		return err
	}
	return nil
}
func (o *IPNATPortMapping) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Protocol); err != nil {
		return err
	}
	if err := w.ReadData(&o.PublicPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.PublicAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivatePort); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivateAddress); err != nil {
		return err
	}
	return nil
}

// IPNATAddressMapping structure represents IP_NAT_ADDRESS_MAPPING RPC structure.
type IPNATAddressMapping struct {
	PrivateAddress       uint32 `idl:"name:PrivateAddress" json:"private_address"`
	PublicAddress        uint32 `idl:"name:PublicAddress" json:"public_address"`
	AllowInboundSessions bool   `idl:"name:AllowInboundSessions" json:"allow_inbound_sessions"`
}

func (o *IPNATAddressMapping) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPNATAddressMapping) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivateAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.PublicAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowInboundSessions); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IPNATAddressMapping) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivateAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.PublicAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowInboundSessions); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IPAlgorithmGlobalInfo structure represents IP_ALG_GLOBAL_INFO RPC structure.
type IPAlgorithmGlobalInfo struct {
	LoggingLevel uint32 `idl:"name:LoggingLevel" json:"logging_level"`
	Flags        uint32 `idl:"name:Flags" json:"flags"`
}

func (o *IPAlgorithmGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPAlgorithmGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *IPAlgorithmGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// RIPGlobalInfo structure represents RIP_GLOBAL_INFO RPC structure.
type RIPGlobalInfo struct {
	EventLogMask uint32 `idl:"name:EventLogMask" json:"event_log_mask"`
}

func (o *RIPGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogMask); err != nil {
		return err
	}
	return nil
}
func (o *RIPGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogMask); err != nil {
		return err
	}
	return nil
}

// RIPRouteFilterInfo structure represents RIP_ROUTE_FILTER_INFO RPC structure.
type RIPRouteFilterInfo struct {
	Network []byte `idl:"name:Network" json:"network"`
	Mask    []byte `idl:"name:Mask" json:"mask"`
}

func (o *RIPRouteFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPRouteFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Network {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Network[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Network); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Mask {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Mask[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Mask); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *RIPRouteFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Network = make([]byte, 4)
	for i1 := range o.Network {
		i1 := i1
		if err := w.ReadData(&o.Network[i1]); err != nil {
			return err
		}
	}
	o.Mask = make([]byte, 4)
	for i1 := range o.Mask {
		i1 := i1
		if err := w.ReadData(&o.Mask[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RIPInterfaceFilters structure represents RIP_IF_FILTERS RPC structure.
type RIPInterfaceFilters struct {
	SupplyFilterAction uint32                `idl:"name:SupplyFilterAction" json:"supply_filter_action"`
	SupplyFilterCount  uint32                `idl:"name:SupplyFilterCount" json:"supply_filter_count"`
	ListenFilterAction uint32                `idl:"name:ListenFilterAction" json:"listen_filter_action"`
	ListenFilterCount  uint32                `idl:"name:ListenFilterCount" json:"listen_filter_count"`
	RouteFilter        []*RIPRouteFilterInfo `idl:"name:RouteFilter" json:"route_filter"`
}

func (o *RIPInterfaceFilters) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterfaceFilters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SupplyFilterAction); err != nil {
		return err
	}
	if err := w.WriteData(o.SupplyFilterCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ListenFilterAction); err != nil {
		return err
	}
	if err := w.WriteData(o.ListenFilterCount); err != nil {
		return err
	}
	for i1 := range o.RouteFilter {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.RouteFilter[i1] != nil {
			if err := o.RouteFilter[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RIPRouteFilterInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.RouteFilter); uint64(i1) < 1; i1++ {
		if err := (&RIPRouteFilterInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterfaceFilters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SupplyFilterAction); err != nil {
		return err
	}
	if err := w.ReadData(&o.SupplyFilterCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ListenFilterAction); err != nil {
		return err
	}
	if err := w.ReadData(&o.ListenFilterCount); err != nil {
		return err
	}
	o.RouteFilter = make([]*RIPRouteFilterInfo, 1)
	for i1 := range o.RouteFilter {
		i1 := i1
		if o.RouteFilter[i1] == nil {
			o.RouteFilter[i1] = &RIPRouteFilterInfo{}
		}
		if err := o.RouteFilter[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// RIPInterfaceConfig structure represents RIP_IF_CONFIG RPC structure.
type RIPInterfaceConfig struct {
	RIPInterfaceInfo    *RIPInterfaceInfo    `idl:"name:RipIfInfo" json:"rip_interface_info"`
	RIPInterfaceFilters *RIPInterfaceFilters `idl:"name:RipIfFilters" json:"rip_interface_filters"`
}

func (o *RIPInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RIPInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.RIPInterfaceInfo != nil {
		if err := o.RIPInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RIPInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RIPInterfaceFilters != nil {
		if err := o.RIPInterfaceFilters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RIPInterfaceFilters{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RIPInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.RIPInterfaceInfo == nil {
		o.RIPInterfaceInfo = &RIPInterfaceInfo{}
	}
	if err := o.RIPInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RIPInterfaceFilters == nil {
		o.RIPInterfaceFilters = &RIPInterfaceFilters{}
	}
	if err := o.RIPInterfaceFilters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAPGlobalInfo structure represents SAP_GLOBAL_INFO RPC structure.
type SAPGlobalInfo struct {
	EventLogMask uint32 `idl:"name:EventLogMask" json:"event_log_mask"`
}

func (o *SAPGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SAPGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogMask); err != nil {
		return err
	}
	return nil
}
func (o *SAPGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogMask); err != nil {
		return err
	}
	return nil
}

// OSPFRouteFilter structure represents OSPF_ROUTE_FILTER RPC structure.
type OSPFRouteFilter struct {
	Address uint32 `idl:"name:dwAddress" json:"address"`
	Mask    uint32 `idl:"name:dwMask" json:"mask"`
}

func (o *OSPFRouteFilter) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFRouteFilter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Address); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	return nil
}
func (o *OSPFRouteFilter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Address); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	return nil
}

// OSPFFilterAction type represents OSPF_FILTER_ACTION RPC enumeration.
type OSPFFilterAction uint16

var (
	OSPFFilterActionDrop   OSPFFilterAction = 0
	OSPFFilterActionAccept OSPFFilterAction = 1
)

func (o OSPFFilterAction) String() string {
	switch o {
	case OSPFFilterActionDrop:
		return "OSPFFilterActionDrop"
	case OSPFFilterActionAccept:
		return "OSPFFilterActionAccept"
	}
	return "Invalid"
}

// OSPFRouteFilterInfo structure represents OSPF_ROUTE_FILTER_INFO RPC structure.
type OSPFRouteFilterInfo struct {
	Type          uint32             `idl:"name:type" json:"type"`
	ActionOnMatch OSPFFilterAction   `idl:"name:ofaActionOnMatch" json:"action_on_match"`
	FiltersLength uint32             `idl:"name:dwNumFilters" json:"filters_length"`
	Filters       []*OSPFRouteFilter `idl:"name:pFilters" json:"filters"`
}

func (o *OSPFRouteFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFRouteFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ActionOnMatch)); err != nil {
		return err
	}
	if err := w.WriteData(o.FiltersLength); err != nil {
		return err
	}
	for i1 := range o.Filters {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Filters[i1] != nil {
			if err := o.Filters[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OSPFRouteFilter{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Filters); uint64(i1) < 1; i1++ {
		if err := (&OSPFRouteFilter{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OSPFRouteFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ActionOnMatch)); err != nil {
		return err
	}
	if err := w.ReadData(&o.FiltersLength); err != nil {
		return err
	}
	o.Filters = make([]*OSPFRouteFilter, 1)
	for i1 := range o.Filters {
		i1 := i1
		if o.Filters[i1] == nil {
			o.Filters[i1] = &OSPFRouteFilter{}
		}
		if err := o.Filters[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OSPFProtoFilterInfo structure represents OSPF_PROTO_FILTER_INFO RPC structure.
type OSPFProtoFilterInfo struct {
	Type           uint32           `idl:"name:type" json:"type"`
	ActionOnMatch  OSPFFilterAction `idl:"name:ofaActionOnMatch" json:"action_on_match"`
	ProtoIDsLength uint32           `idl:"name:dwNumProtoIds" json:"proto_ids_length"`
	ProtoID        []uint32         `idl:"name:pdwProtoId" json:"proto_id"`
}

func (o *OSPFProtoFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFProtoFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ActionOnMatch)); err != nil {
		return err
	}
	if err := w.WriteData(o.ProtoIDsLength); err != nil {
		return err
	}
	for i1 := range o.ProtoID {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.ProtoID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ProtoID); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *OSPFProtoFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ActionOnMatch)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProtoIDsLength); err != nil {
		return err
	}
	o.ProtoID = make([]uint32, 1)
	for i1 := range o.ProtoID {
		i1 := i1
		if err := w.ReadData(&o.ProtoID[i1]); err != nil {
			return err
		}
	}
	return nil
}

// OSPFGlobalParam structure represents OSPF_GLOBAL_PARAM RPC structure.
type OSPFGlobalParam struct {
	Type           uint32 `idl:"name:type" json:"type"`
	Create         uint32 `idl:"name:create" json:"create"`
	Enable         uint32 `idl:"name:enable" json:"enable"`
	RouterID       uint32 `idl:"name:routerId" json:"router_id"`
	ASBorderRouter uint32 `idl:"name:ASBrdrRtr" json:"as_border_router"`
	LogLevel       uint32 `idl:"name:logLevel" json:"log_level"`
}

func (o *OSPFGlobalParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFGlobalParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Create); err != nil {
		return err
	}
	if err := w.WriteData(o.Enable); err != nil {
		return err
	}
	if err := w.WriteData(o.RouterID); err != nil {
		return err
	}
	if err := w.WriteData(o.ASBorderRouter); err != nil {
		return err
	}
	if err := w.WriteData(o.LogLevel); err != nil {
		return err
	}
	return nil
}
func (o *OSPFGlobalParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Create); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouterID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ASBorderRouter); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogLevel); err != nil {
		return err
	}
	return nil
}

// OSPFAreaParam structure represents OSPF_AREA_PARAM RPC structure.
type OSPFAreaParam struct {
	Type           uint32 `idl:"name:type" json:"type"`
	Create         uint32 `idl:"name:create" json:"create"`
	Enable         uint32 `idl:"name:enable" json:"enable"`
	AreaID         uint32 `idl:"name:areaId" json:"area_id"`
	AuthType       uint32 `idl:"name:authType" json:"auth_type"`
	ImportAsExtern uint32 `idl:"name:importASExtern" json:"import_as_extern"`
	StubMetric     uint32 `idl:"name:stubMetric" json:"stub_metric"`
	ImportSumAdv   uint32 `idl:"name:importSumAdv" json:"import_sum_adv"`
}

func (o *OSPFAreaParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFAreaParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Create); err != nil {
		return err
	}
	if err := w.WriteData(o.Enable); err != nil {
		return err
	}
	if err := w.WriteData(o.AreaID); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthType); err != nil {
		return err
	}
	if err := w.WriteData(o.ImportAsExtern); err != nil {
		return err
	}
	if err := w.WriteData(o.StubMetric); err != nil {
		return err
	}
	if err := w.WriteData(o.ImportSumAdv); err != nil {
		return err
	}
	return nil
}
func (o *OSPFAreaParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Create); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	if err := w.ReadData(&o.AreaID); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthType); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImportAsExtern); err != nil {
		return err
	}
	if err := w.ReadData(&o.StubMetric); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImportSumAdv); err != nil {
		return err
	}
	return nil
}

// OSPFAreaRangeParam structure represents OSPF_AREA_RANGE_PARAM RPC structure.
type OSPFAreaRangeParam struct {
	Type      uint32 `idl:"name:type" json:"type"`
	Create    uint32 `idl:"name:create" json:"create"`
	Enable    uint32 `idl:"name:enable" json:"enable"`
	AreaID    uint32 `idl:"name:areaId" json:"area_id"`
	RangeNet  uint32 `idl:"name:rangeNet" json:"range_net"`
	RangeMask uint32 `idl:"name:rangeMask" json:"range_mask"`
}

func (o *OSPFAreaRangeParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFAreaRangeParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Create); err != nil {
		return err
	}
	if err := w.WriteData(o.Enable); err != nil {
		return err
	}
	if err := w.WriteData(o.AreaID); err != nil {
		return err
	}
	if err := w.WriteData(o.RangeNet); err != nil {
		return err
	}
	if err := w.WriteData(o.RangeMask); err != nil {
		return err
	}
	return nil
}
func (o *OSPFAreaRangeParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Create); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	if err := w.ReadData(&o.AreaID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RangeNet); err != nil {
		return err
	}
	if err := w.ReadData(&o.RangeMask); err != nil {
		return err
	}
	return nil
}

// OSPFVirtualInterfaceParam structure represents OSPF_VIRT_INTERFACE_PARAM RPC structure.
type OSPFVirtualInterfaceParam struct {
	Type                    uint32 `idl:"name:type" json:"type"`
	Create                  uint32 `idl:"name:create" json:"create"`
	Enable                  uint32 `idl:"name:enable" json:"enable"`
	TransitAreaID           uint32 `idl:"name:transitAreaId" json:"transit_area_id"`
	VirtualNeighborRouterID uint32 `idl:"name:virtNeighborRouterId" json:"virtual_neighbor_router_id"`
	TransitDelay            uint32 `idl:"name:transitDelay" json:"transit_delay"`
	RetransimitInterval     uint32 `idl:"name:retransInterval" json:"retransimit_interval"`
	HelloInterval           uint32 `idl:"name:helloInterval" json:"hello_interval"`
	DeadInterval            uint32 `idl:"name:deadInterval" json:"dead_interval"`
	Password                []byte `idl:"name:password" json:"password"`
}

func (o *OSPFVirtualInterfaceParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFVirtualInterfaceParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Create); err != nil {
		return err
	}
	if err := w.WriteData(o.Enable); err != nil {
		return err
	}
	if err := w.WriteData(o.TransitAreaID); err != nil {
		return err
	}
	if err := w.WriteData(o.VirtualNeighborRouterID); err != nil {
		return err
	}
	if err := w.WriteData(o.TransitDelay); err != nil {
		return err
	}
	if err := w.WriteData(o.RetransimitInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.HelloInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DeadInterval); err != nil {
		return err
	}
	for i1 := range o.Password {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Password[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Password); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *OSPFVirtualInterfaceParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Create); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransitAreaID); err != nil {
		return err
	}
	if err := w.ReadData(&o.VirtualNeighborRouterID); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransitDelay); err != nil {
		return err
	}
	if err := w.ReadData(&o.RetransimitInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.HelloInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeadInterval); err != nil {
		return err
	}
	o.Password = make([]byte, 8)
	for i1 := range o.Password {
		i1 := i1
		if err := w.ReadData(&o.Password[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// OSPFInterfaceParam structure represents OSPF_INTERFACE_PARAM RPC structure.
type OSPFInterfaceParam struct {
	Type                uint32 `idl:"name:type" json:"type"`
	Create              uint32 `idl:"name:create" json:"create"`
	Enable              uint32 `idl:"name:enable" json:"enable"`
	InterfaceIPAddr     uint32 `idl:"name:intfIpAddr" json:"interface_ip_addr"`
	InterfaceSubnetMask uint32 `idl:"name:intfSubnetMask" json:"interface_subnet_mask"`
	AreaID              uint32 `idl:"name:areaId" json:"area_id"`
	InterfaceType       uint32 `idl:"name:intfType" json:"interface_type"`
	RouterPriority      uint32 `idl:"name:routerPriority" json:"router_priority"`
	TransitDelay        uint32 `idl:"name:transitDelay" json:"transit_delay"`
	RetransimitInterval uint32 `idl:"name:retransInterval" json:"retransimit_interval"`
	HelloInterval       uint32 `idl:"name:helloInterval" json:"hello_interval"`
	DeadInterval        uint32 `idl:"name:deadInterval" json:"dead_interval"`
	PollInterval        uint32 `idl:"name:pollInterval" json:"poll_interval"`
	MetricCost          uint32 `idl:"name:metricCost" json:"metric_cost"`
	Password            []byte `idl:"name:password" json:"password"`
	MTUSize             uint32 `idl:"name:mtuSize" json:"mtu_size"`
}

func (o *OSPFInterfaceParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFInterfaceParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Create); err != nil {
		return err
	}
	if err := w.WriteData(o.Enable); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceSubnetMask); err != nil {
		return err
	}
	if err := w.WriteData(o.AreaID); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceType); err != nil {
		return err
	}
	if err := w.WriteData(o.RouterPriority); err != nil {
		return err
	}
	if err := w.WriteData(o.TransitDelay); err != nil {
		return err
	}
	if err := w.WriteData(o.RetransimitInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.HelloInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DeadInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.PollInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.MetricCost); err != nil {
		return err
	}
	for i1 := range o.Password {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Password[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Password); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MTUSize); err != nil {
		return err
	}
	return nil
}
func (o *OSPFInterfaceParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Create); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceSubnetMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.AreaID); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceType); err != nil {
		return err
	}
	if err := w.ReadData(&o.RouterPriority); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransitDelay); err != nil {
		return err
	}
	if err := w.ReadData(&o.RetransimitInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.HelloInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeadInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.PollInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetricCost); err != nil {
		return err
	}
	o.Password = make([]byte, 8)
	for i1 := range o.Password {
		i1 := i1
		if err := w.ReadData(&o.Password[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.MTUSize); err != nil {
		return err
	}
	return nil
}

// OSPFNBMANeighborParam structure represents OSPF_NBMA_NEIGHBOR_PARAM RPC structure.
type OSPFNBMANeighborParam struct {
	Type             uint32 `idl:"name:type" json:"type"`
	Create           uint32 `idl:"name:create" json:"create"`
	Enable           uint32 `idl:"name:enable" json:"enable"`
	NeighborIPAddr   uint32 `idl:"name:neighborIpAddr" json:"neighbor_ip_addr"`
	InterfaceIPAddr  uint32 `idl:"name:intfIpAddr" json:"interface_ip_addr"`
	NeighborPriority uint32 `idl:"name:neighborPriority" json:"neighbor_priority"`
}

func (o *OSPFNBMANeighborParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OSPFNBMANeighborParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Create); err != nil {
		return err
	}
	if err := w.WriteData(o.Enable); err != nil {
		return err
	}
	if err := w.WriteData(o.NeighborIPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIPAddr); err != nil {
		return err
	}
	if err := w.WriteData(o.NeighborPriority); err != nil {
		return err
	}
	return nil
}
func (o *OSPFNBMANeighborParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Create); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	if err := w.ReadData(&o.NeighborIPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIPAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.NeighborPriority); err != nil {
		return err
	}
	return nil
}

// RASDeviceType type represents RASDEVICETYPE RPC enumeration.
type RASDeviceType uint32

var (
	RASDeviceTypeModem       RASDeviceType = 0
	RASDeviceTypeX25         RASDeviceType = 1
	RASDeviceTypeISDN        RASDeviceType = 2
	RASDeviceTypeSerial      RASDeviceType = 3
	RASDeviceTypeFrameRelay  RASDeviceType = 4
	RASDeviceTypeATM         RASDeviceType = 5
	RASDeviceTypeSONET       RASDeviceType = 6
	RASDeviceTypeSW56        RASDeviceType = 7
	RASDeviceTypeTunnelPPTP  RASDeviceType = 8
	RASDeviceTypeTunnelL2TP  RASDeviceType = 9
	RASDeviceTypeIRDA        RASDeviceType = 10
	RASDeviceTypeParallel    RASDeviceType = 11
	RASDeviceTypeOther       RASDeviceType = 12
	RASDeviceTypePPPoE       RASDeviceType = 13
	RASDeviceTypeTunnelSSTP  RASDeviceType = 14
	RASDeviceTypeTunnelIKEv2 RASDeviceType = 15
	RASDeviceTypeTunnel      RASDeviceType = 65536
	RASDeviceTypeDirect      RASDeviceType = 131072
	RASDeviceTypeNullModem   RASDeviceType = 262144
	RASDeviceTypeBroadband   RASDeviceType = 524288
)

func (o RASDeviceType) String() string {
	switch o {
	case RASDeviceTypeModem:
		return "RASDeviceTypeModem"
	case RASDeviceTypeX25:
		return "RASDeviceTypeX25"
	case RASDeviceTypeISDN:
		return "RASDeviceTypeISDN"
	case RASDeviceTypeSerial:
		return "RASDeviceTypeSerial"
	case RASDeviceTypeFrameRelay:
		return "RASDeviceTypeFrameRelay"
	case RASDeviceTypeATM:
		return "RASDeviceTypeATM"
	case RASDeviceTypeSONET:
		return "RASDeviceTypeSONET"
	case RASDeviceTypeSW56:
		return "RASDeviceTypeSW56"
	case RASDeviceTypeTunnelPPTP:
		return "RASDeviceTypeTunnelPPTP"
	case RASDeviceTypeTunnelL2TP:
		return "RASDeviceTypeTunnelL2TP"
	case RASDeviceTypeIRDA:
		return "RASDeviceTypeIRDA"
	case RASDeviceTypeParallel:
		return "RASDeviceTypeParallel"
	case RASDeviceTypeOther:
		return "RASDeviceTypeOther"
	case RASDeviceTypePPPoE:
		return "RASDeviceTypePPPoE"
	case RASDeviceTypeTunnelSSTP:
		return "RASDeviceTypeTunnelSSTP"
	case RASDeviceTypeTunnelIKEv2:
		return "RASDeviceTypeTunnelIKEv2"
	case RASDeviceTypeTunnel:
		return "RASDeviceTypeTunnel"
	case RASDeviceTypeDirect:
		return "RASDeviceTypeDirect"
	case RASDeviceTypeNullModem:
		return "RASDeviceTypeNullModem"
	case RASDeviceTypeBroadband:
		return "RASDeviceTypeBroadband"
	}
	return "Invalid"
}

// RASMANStatus type represents RASMAN_STATUS RPC enumeration.
type RASMANStatus uint16

var (
	RASMANStatusOpen        RASMANStatus = 0
	RASMANStatusClosed      RASMANStatus = 1
	RASMANStatusUnavailable RASMANStatus = 2
	RASMANStatusRemoved     RASMANStatus = 3
)

func (o RASMANStatus) String() string {
	switch o {
	case RASMANStatusOpen:
		return "RASMANStatusOpen"
	case RASMANStatusClosed:
		return "RASMANStatusClosed"
	case RASMANStatusUnavailable:
		return "RASMANStatusUnavailable"
	case RASMANStatusRemoved:
		return "RASMANStatusRemoved"
	}
	return "Invalid"
}

// RequestTypes type represents ReqTypes RPC enumeration.
type RequestTypes uint16

var (
	RequestTypesPortEnum             RequestTypes = 21
	RequestTypesGetInfo              RequestTypes = 22
	RequestTypesGetDeviceConfig      RequestTypes = 73
	RequestTypesSetDeviceConfigInfo  RequestTypes = 94
	RequestTypesGetDeviceConfigInfo  RequestTypes = 95
	RequestTypesGetCalledID          RequestTypes = 105
	RequestTypesSetCalledID          RequestTypes = 106
	RequestTypesGetNDISWANDriverCaps RequestTypes = 111
)

func (o RequestTypes) String() string {
	switch o {
	case RequestTypesPortEnum:
		return "RequestTypesPortEnum"
	case RequestTypesGetInfo:
		return "RequestTypesGetInfo"
	case RequestTypesGetDeviceConfig:
		return "RequestTypesGetDeviceConfig"
	case RequestTypesSetDeviceConfigInfo:
		return "RequestTypesSetDeviceConfigInfo"
	case RequestTypesGetDeviceConfigInfo:
		return "RequestTypesGetDeviceConfigInfo"
	case RequestTypesGetCalledID:
		return "RequestTypesGetCalledID"
	case RequestTypesSetCalledID:
		return "RequestTypesSetCalledID"
	case RequestTypesGetNDISWANDriverCaps:
		return "RequestTypesGetNDISWANDriverCaps"
	}
	return "Invalid"
}

// RASMANState type represents RASMAN_STATE RPC enumeration.
type RASMANState uint16

var (
	RASMANStateConnecting      RASMANState = 0
	RASMANStateListening       RASMANState = 1
	RASMANStateConnected       RASMANState = 2
	RASMANStateDisconnecting   RASMANState = 3
	RASMANStateDisconnected    RASMANState = 4
	RASMANStateListenCompleted RASMANState = 5
)

func (o RASMANState) String() string {
	switch o {
	case RASMANStateConnecting:
		return "RASMANStateConnecting"
	case RASMANStateListening:
		return "RASMANStateListening"
	case RASMANStateConnected:
		return "RASMANStateConnected"
	case RASMANStateDisconnecting:
		return "RASMANStateDisconnecting"
	case RASMANStateDisconnected:
		return "RASMANStateDisconnected"
	case RASMANStateListenCompleted:
		return "RASMANStateListenCompleted"
	}
	return "Invalid"
}

// RASMANDisconnectType type represents RASMAN_DISCONNECT_TYPE RPC enumeration.
type RASMANDisconnectType uint16

var (
	RASMANDisconnectTypeUserRequested       RASMANDisconnectType = 0
	RASMANDisconnectTypeRemoteDisconnection RASMANDisconnectType = 1
	RASMANDisconnectTypeHardwareFailure     RASMANDisconnectType = 2
	RASMANDisconnectTypeNotDisconnected     RASMANDisconnectType = 3
)

func (o RASMANDisconnectType) String() string {
	switch o {
	case RASMANDisconnectTypeUserRequested:
		return "RASMANDisconnectTypeUserRequested"
	case RASMANDisconnectTypeRemoteDisconnection:
		return "RASMANDisconnectTypeRemoteDisconnection"
	case RASMANDisconnectTypeHardwareFailure:
		return "RASMANDisconnectTypeHardwareFailure"
	case RASMANDisconnectTypeNotDisconnected:
		return "RASMANDisconnectTypeNotDisconnected"
	}
	return "Invalid"
}

// RASMANUsage type represents RASMAN_USAGE RPC enumeration.
type RASMANUsage uint16

var (
	RASMANUsageCallNone           RASMANUsage = 0
	RASMANUsageCallIn             RASMANUsage = 1
	RASMANUsageCallOut            RASMANUsage = 2
	RASMANUsageCallRouter         RASMANUsage = 4
	RASMANUsageCallLogon          RASMANUsage = 8
	RASMANUsageCallOutOnly        RASMANUsage = 16
	RASMANUsageCallInOnly         RASMANUsage = 32
	RASMANUsageCallOutboundRouter RASMANUsage = 64
)

func (o RASMANUsage) String() string {
	switch o {
	case RASMANUsageCallNone:
		return "RASMANUsageCallNone"
	case RASMANUsageCallIn:
		return "RASMANUsageCallIn"
	case RASMANUsageCallOut:
		return "RASMANUsageCallOut"
	case RASMANUsageCallRouter:
		return "RASMANUsageCallRouter"
	case RASMANUsageCallLogon:
		return "RASMANUsageCallLogon"
	case RASMANUsageCallOutOnly:
		return "RASMANUsageCallOutOnly"
	case RASMANUsageCallInOnly:
		return "RASMANUsageCallInOnly"
	case RASMANUsageCallOutboundRouter:
		return "RASMANUsageCallOutboundRouter"
	}
	return "Invalid"
}

// RequestBuffer structure represents RequestBuffer RPC structure.
type RequestBuffer struct {
	Index       uint32       `idl:"name:RB_PCBIndex" json:"index"`
	RequestType RequestTypes `idl:"name:RB_Reqtype" json:"request_type"`
	_           uint32       `idl:"name:RB_Dummy"`
	Done        uint32       `idl:"name:RB_Done" json:"done"`
	Alignment   int64        `idl:"name:Alignment" json:"alignment"`
	Buffer      []byte       `idl:"name:RB_Buffer" json:"buffer"`
}

func (o *RequestBuffer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RequestBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RequestType)); err != nil {
		return err
	}
	// reserved RB_Dummy
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Done); err != nil {
		return err
	}
	if err := w.WriteData(o.Alignment); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *RequestBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RequestType)); err != nil {
		return err
	}
	// reserved RB_Dummy
	var _RB_Dummy uint32
	if err := w.ReadData(&_RB_Dummy); err != nil {
		return err
	}
	if err := w.ReadData(&o.Done); err != nil {
		return err
	}
	if err := w.ReadData(&o.Alignment); err != nil {
		return err
	}
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// DeviceConfigInfo structure represents DeviceConfigInfo RPC structure.
type DeviceConfigInfo struct {
	ReturnCode   uint32 `idl:"name:retcode" json:"return_code"`
	Version      uint32 `idl:"name:dwVersion" json:"version"`
	BufferLength uint32 `idl:"name:cbBuffer" json:"buffer_length"`
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	Data         []byte `idl:"name:abdata" json:"data"`
}

func (o *DeviceConfigInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DeviceConfigInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferLength); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *DeviceConfigInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	o.Data = make([]byte, 1)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// RASDeviceInfo structure represents RAS_DEVICE_INFO RPC structure.
type RASDeviceInfo struct {
	Version               uint32        `idl:"name:dwVersion" json:"version"`
	Write                 bool          `idl:"name:fWrite" json:"write"`
	RASEnabled            bool          `idl:"name:fRasEnabled" json:"ras_enabled"`
	RouterEnabled         bool          `idl:"name:fRouterEnabled" json:"router_enabled"`
	RouterOutboundEnabled bool          `idl:"name:fRouterOutboundEnabled" json:"router_outbound_enabled"`
	TAPILineID            uint32        `idl:"name:dwTapiLineId" json:"tapi_line_id"`
	Error                 uint32        `idl:"name:dwError" json:"error"`
	EndpointsLength       uint32        `idl:"name:dwNumEndPoints" json:"endpoints_length"`
	MaxOutCalls           uint32        `idl:"name:dwMaxOutCalls" json:"max_out_calls"`
	MaxInCalls            uint32        `idl:"name:dwMaxInCalls" json:"max_in_calls"`
	MinWANEndpoints       uint32        `idl:"name:dwMinWanEndPoints" json:"min_wan_endpoints"`
	MaxWANEndpoints       uint32        `idl:"name:dwMaxWanEndPoints" json:"max_wan_endpoints"`
	DeviceType            RASDeviceType `idl:"name:eDeviceType" json:"device_type"`
	Device                *dtyp.GUID    `idl:"name:guidDevice" json:"device"`
	PortName              []byte        `idl:"name:szPortName" json:"port_name"`
	DeviceName            []byte        `idl:"name:szDeviceName" json:"device_name"`
	UnicodeDeviceName     []uint16      `idl:"name:wszUnicodeDeviceName" json:"unicode_device_name"`
}

func (o *RASDeviceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASDeviceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if !o.Write {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.RASEnabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.RouterEnabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.RouterOutboundEnabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TAPILineID); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	if err := w.WriteData(o.EndpointsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxOutCalls); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxInCalls); err != nil {
		return err
	}
	if err := w.WriteData(o.MinWANEndpoints); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxWANEndpoints); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.DeviceType)); err != nil {
		return err
	}
	if o.Device != nil {
		if err := o.Device.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.PortName {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.PortName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PortName); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.UnicodeDeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.UnicodeDeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UnicodeDeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *RASDeviceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	var _bWrite int32
	if err := w.ReadData(&_bWrite); err != nil {
		return err
	}
	o.Write = _bWrite != 0
	var _bRASEnabled int32
	if err := w.ReadData(&_bRASEnabled); err != nil {
		return err
	}
	o.RASEnabled = _bRASEnabled != 0
	var _bRouterEnabled int32
	if err := w.ReadData(&_bRouterEnabled); err != nil {
		return err
	}
	o.RouterEnabled = _bRouterEnabled != 0
	var _bRouterOutboundEnabled int32
	if err := w.ReadData(&_bRouterOutboundEnabled); err != nil {
		return err
	}
	o.RouterOutboundEnabled = _bRouterOutboundEnabled != 0
	if err := w.ReadData(&o.TAPILineID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	if err := w.ReadData(&o.EndpointsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxOutCalls); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxInCalls); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinWANEndpoints); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxWANEndpoints); err != nil {
		return err
	}
	_eDeviceType := uint16(o.DeviceType)
	if err := w.ReadEnum(&_eDeviceType); err != nil {
		return err
	}
	o.DeviceType = RASDeviceType(_eDeviceType)
	if o.Device == nil {
		o.Device = &dtyp.GUID{}
	}
	if err := o.Device.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.PortName = make([]byte, 17)
	for i1 := range o.PortName {
		i1 := i1
		if err := w.ReadData(&o.PortName[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]byte, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.UnicodeDeviceName = make([]uint16, 129)
	for i1 := range o.UnicodeDeviceName {
		i1 := i1
		if err := w.ReadData(&o.UnicodeDeviceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// RASCalledidInfo structure represents RAS_CALLEDID_INFO RPC structure.
type RASCalledidInfo struct {
	Size     uint32 `idl:"name:dwSize" json:"size"`
	CalledID []byte `idl:"name:bCalledId" json:"called_id"`
}

func (o *RASCalledidInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASCalledidInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	for i1 := range o.CalledID {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.CalledID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.CalledID); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *RASCalledidInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	o.CalledID = make([]byte, 1)
	for i1 := range o.CalledID {
		i1 := i1
		if err := w.ReadData(&o.CalledID[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// GetSetCalledID structure represents GetSetCalledId RPC structure.
type GetSetCalledID struct {
	ReturnCode    uint32           `idl:"name:retcode" json:"return_code"`
	Write         bool             `idl:"name:fWrite" json:"write"`
	Size          uint32           `idl:"name:dwSize" json:"size"`
	Device        *dtyp.GUID       `idl:"name:guidDevice" json:"device"`
	RASDeviceInfo *RASDeviceInfo   `idl:"name:rdi" json:"ras_device_info"`
	Info          *RASCalledidInfo `idl:"name:rciInfo" json:"info"`
}

func (o *GetSetCalledID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *GetSetCalledID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if !o.Write {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Device != nil {
		if err := o.Device.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RASDeviceInfo != nil {
		if err := o.RASDeviceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASDeviceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Info != nil {
		if err := o.Info.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASCalledidInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetSetCalledID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	var _bWrite int32
	if err := w.ReadData(&_bWrite); err != nil {
		return err
	}
	o.Write = _bWrite != 0
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if o.Device == nil {
		o.Device = &dtyp.GUID{}
	}
	if err := o.Device.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RASDeviceInfo == nil {
		o.RASDeviceInfo = &RASDeviceInfo{}
	}
	if err := o.RASDeviceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Info == nil {
		o.Info = &RASCalledidInfo{}
	}
	if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASNDISWANDriverInfo structure represents RAS_NDISWAN_DRIVER_INFO RPC structure.
type RASNDISWANDriverInfo struct {
	DriverCaps uint32 `idl:"name:DriverCaps" json:"driver_caps"`
	_          uint32 `idl:"name:Reserved"`
}

func (o *RASNDISWANDriverInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASNDISWANDriverInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DriverCaps); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	return nil
}
func (o *RASNDISWANDriverInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DriverCaps); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	return nil
}

// GetNDISWANDriverCaps structure represents GetNdiswanDriverCapsStruct RPC structure.
type GetNDISWANDriverCaps struct {
	ReturnCode        uint32                `idl:"name:retcode" json:"return_code"`
	NDISWANDriverInfo *RASNDISWANDriverInfo `idl:"name:NdiswanDriverInfo" json:"ndiswan_driver_info"`
}

func (o *GetNDISWANDriverCaps) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *GetNDISWANDriverCaps) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if o.NDISWANDriverInfo != nil {
		if err := o.NDISWANDriverInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASNDISWANDriverInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetNDISWANDriverCaps) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if o.NDISWANDriverInfo == nil {
		o.NDISWANDriverInfo = &RASNDISWANDriverInfo{}
	}
	if err := o.NDISWANDriverInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GetDevConfig structure represents GetDevConfigStruct RPC structure.
type GetDevConfig struct {
	ReturnCode uint32 `idl:"name:retcode" json:"return_code"`
	DeviceType []byte `idl:"name:devicetype" json:"device_type"`
	Size       uint32 `idl:"name:size" json:"size"`
	Config     []byte `idl:"name:config" json:"config"`
}

func (o *GetDevConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *GetDevConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	for i1 := range o.Config {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Config[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Config); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *GetDevConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	o.DeviceType = make([]byte, 17)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	o.Config = make([]byte, 1)
	for i1 := range o.Config {
		i1 := i1
		if err := w.ReadData(&o.Config[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// Enum structure represents Enum RPC structure.
type Enum struct {
	ReturnCode uint32 `idl:"name:retcode" json:"return_code"`
	Size       uint32 `idl:"name:size" json:"size"`
	Entries    uint32 `idl:"name:entries" json:"entries"`
	Buffer     []byte `idl:"name:buffer" json:"buffer"`
}

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
	if err := w.WriteData(o.ReturnCode); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *Enum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	o.Buffer = make([]byte, 1)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// RASMANPort32 structure represents RASMAN_PORT_32 RPC structure.
type RASMANPort32 struct {
	Port            uint32        `idl:"name:P_Port" json:"port"`
	PortName        []byte        `idl:"name:P_PortName" json:"port_name"`
	Status          RASMANStatus  `idl:"name:P_Status" json:"status"`
	RASDeviceType   RASDeviceType `idl:"name:P_rdtDeviceType" json:"ras_device_type"`
	ConfiguredUsage RASMANUsage   `idl:"name:P_ConfiguredUsage" json:"configured_usage"`
	CurrentUsage    RASMANUsage   `idl:"name:P_CurrentUsage" json:"current_usage"`
	MediaName       []byte        `idl:"name:P_MediaName" json:"media_name"`
	DeviceType      []byte        `idl:"name:P_DeviceType" json:"device_type"`
	DeviceName      []byte        `idl:"name:P_DeviceName" json:"device_name"`
	LineDeviceID    uint32        `idl:"name:P_LineDeviceId" json:"line_device_id"`
	AddressID       uint32        `idl:"name:P_AddressId" json:"address_id"`
}

func (o *RASMANPort32) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASMANPort32) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	for i1 := range o.PortName {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.PortName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PortName); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RASDeviceType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ConfiguredUsage)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.CurrentUsage)); err != nil {
		return err
	}
	for i1 := range o.MediaName {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.MediaName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.MediaName); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LineDeviceID); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressID); err != nil {
		return err
	}
	return nil
}
func (o *RASMANPort32) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	o.PortName = make([]byte, 16)
	for i1 := range o.PortName {
		i1 := i1
		if err := w.ReadData(&o.PortName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	_eRASDeviceType := uint16(o.RASDeviceType)
	if err := w.ReadEnum(&_eRASDeviceType); err != nil {
		return err
	}
	o.RASDeviceType = RASDeviceType(_eRASDeviceType)
	if err := w.ReadEnum((*uint16)(&o.ConfiguredUsage)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.CurrentUsage)); err != nil {
		return err
	}
	o.MediaName = make([]byte, 16)
	for i1 := range o.MediaName {
		i1 := i1
		if err := w.ReadData(&o.MediaName[i1]); err != nil {
			return err
		}
	}
	o.DeviceType = make([]byte, 16)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]byte, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.LineDeviceID); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressID); err != nil {
		return err
	}
	return nil
}

// RASMANInfo structure represents RASMAN_INFO RPC structure.
type RASMANInfo struct {
	PortStatus           RASMANStatus         `idl:"name:RI_PortStatus" json:"port_status"`
	ConnState            RASMANState          `idl:"name:RI_ConnState" json:"conn_state"`
	LinkSpeed            uint32               `idl:"name:RI_LinkSpeed" json:"link_speed"`
	LastError            uint32               `idl:"name:RI_LastError" json:"last_error"`
	CurrentUsage         RASMANUsage          `idl:"name:RI_CurrentUsage" json:"current_usage"`
	DeviceTypeConnecting []byte               `idl:"name:RI_DeviceTypeConnecting" json:"device_type_connecting"`
	DeviceConnecting     []byte               `idl:"name:RI_DeviceConnecting" json:"device_connecting"`
	DeviceType           []byte               `idl:"name:RI_szDeviceType" json:"device_type"`
	DeviceName           []byte               `idl:"name:RI_szDeviceName" json:"device_name"`
	PortName             []byte               `idl:"name:RI_szPortName" json:"port_name"`
	DisconnectType       RASMANDisconnectType `idl:"name:RI_DisconnectType" json:"disconnect_type"`
	OwnershipFlag        uint32               `idl:"name:RI_OwnershipFlag" json:"ownership_flag"`
	ConnectDuration      uint32               `idl:"name:RI_ConnectDuration" json:"connect_duration"`
	BytesReceived        uint32               `idl:"name:RI_BytesReceived" json:"bytes_received"`
	Phonebook            []byte               `idl:"name:RI_Phonebook" json:"phonebook"`
	PhoneEntry           []byte               `idl:"name:RI_PhoneEntry" json:"phone_entry"`
	ConnectionHandle     []byte               `idl:"name:RI_ConnectionHandle" json:"connection_handle"`
	SubEntry             uint32               `idl:"name:RI_SubEntry" json:"sub_entry"`
	RASDeviceType        RASDeviceType        `idl:"name:RI_rdtDeviceType" json:"ras_device_type"`
	GUIDEntry            *dtyp.GUID           `idl:"name:RI_GuidEntry" json:"guid_entry"`
	SessionID            uint32               `idl:"name:RI_dwSessionId" json:"session_id"`
	Flags                uint32               `idl:"name:RI_dwFlags" json:"flags"`
	CorrelationGUID      *dtyp.GUID           `idl:"name:RI_CorrelationGuid" json:"correlation_guid"`
}

func (o *RASMANInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASMANInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PortStatus)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ConnState)); err != nil {
		return err
	}
	if err := w.WriteData(o.LinkSpeed); err != nil {
		return err
	}
	if err := w.WriteData(o.LastError); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.CurrentUsage)); err != nil {
		return err
	}
	for i1 := range o.DeviceTypeConnecting {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.DeviceTypeConnecting[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceTypeConnecting); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceConnecting {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceConnecting[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceConnecting); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceType {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.DeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceType); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.PortName {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.PortName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PortName); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.DisconnectType)); err != nil {
		return err
	}
	if err := w.WriteData(o.OwnershipFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectDuration); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesReceived); err != nil {
		return err
	}
	for i1 := range o.Phonebook {
		i1 := i1
		if uint64(i1) >= 261 {
			break
		}
		if err := w.WriteData(o.Phonebook[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Phonebook); uint64(i1) < 261; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.PhoneEntry {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.PhoneEntry[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PhoneEntry); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.ConnectionHandle != nil {
		_ptr_RI_ConnectionHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			// FIXME unknown type RI_ConnectionHandle
			return nil
		})
		if err := w.WritePointer(&o.ConnectionHandle, _ptr_RI_ConnectionHandle); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SubEntry); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RASDeviceType)); err != nil {
		return err
	}
	if o.GUIDEntry != nil {
		if err := o.GUIDEntry.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.CorrelationGUID != nil {
		if err := o.CorrelationGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *RASMANInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PortStatus)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ConnState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.LinkSpeed); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastError); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.CurrentUsage)); err != nil {
		return err
	}
	o.DeviceTypeConnecting = make([]byte, 16)
	for i1 := range o.DeviceTypeConnecting {
		i1 := i1
		if err := w.ReadData(&o.DeviceTypeConnecting[i1]); err != nil {
			return err
		}
	}
	o.DeviceConnecting = make([]byte, 129)
	for i1 := range o.DeviceConnecting {
		i1 := i1
		if err := w.ReadData(&o.DeviceConnecting[i1]); err != nil {
			return err
		}
	}
	o.DeviceType = make([]byte, 16)
	for i1 := range o.DeviceType {
		i1 := i1
		if err := w.ReadData(&o.DeviceType[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]byte, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.PortName = make([]byte, 17)
	for i1 := range o.PortName {
		i1 := i1
		if err := w.ReadData(&o.PortName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadEnum((*uint16)(&o.DisconnectType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.OwnershipFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectDuration); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesReceived); err != nil {
		return err
	}
	o.Phonebook = make([]byte, 261)
	for i1 := range o.Phonebook {
		i1 := i1
		if err := w.ReadData(&o.Phonebook[i1]); err != nil {
			return err
		}
	}
	o.PhoneEntry = make([]byte, 257)
	for i1 := range o.PhoneEntry {
		i1 := i1
		if err := w.ReadData(&o.PhoneEntry[i1]); err != nil {
			return err
		}
	}
	_ptr_RI_ConnectionHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		// FIXME: unknown type RI_ConnectionHandle
		return nil
	})
	_s_RI_ConnectionHandle := func(ptr interface{}) { o.ConnectionHandle = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ConnectionHandle, _s_RI_ConnectionHandle, _ptr_RI_ConnectionHandle); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubEntry); err != nil {
		return err
	}
	_eRASDeviceType := uint16(o.RASDeviceType)
	if err := w.ReadEnum(&_eRASDeviceType); err != nil {
		return err
	}
	o.RASDeviceType = RASDeviceType(_eRASDeviceType)
	if o.GUIDEntry == nil {
		o.GUIDEntry = &dtyp.GUID{}
	}
	if err := o.GUIDEntry.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.CorrelationGUID == nil {
		o.CorrelationGUID = &dtyp.GUID{}
	}
	if err := o.CorrelationGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// Info structure represents Info RPC structure.
type Info struct {
	Field1 *Info_Field1 `idl:"name:" json:""`
	Info   *RASMANInfo  `idl:"name:info" json:"info"`
}

func (o *Info) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Info) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	// FIXME unknown type
	if o.Info != nil {
		if err := o.Info.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASMANInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	// FIXME: unknown type
	if o.Info == nil {
		o.Info = &RASMANInfo{}
	}
	if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type Info_Field1 struct {
	ReturnCode   uint32 `idl:"name:retcode" json:"return_code"`
	PaddingField []byte `idl:"name:paddingField" json:"padding_field"`
}

// CallbackList structure represents RASRPC_CALLBACKLIST RPC structure.
type CallbackList struct {
	PortName   []uint16      `idl:"name:pszPortName" json:"port_name"`
	DeviceName []uint16      `idl:"name:pszDeviceName" json:"device_name"`
	Number     []uint16      `idl:"name:pszNumber" json:"number"`
	DeviceType uint32        `idl:"name:dwDeviceType" json:"device_type"`
	Next       *CallbackList `idl:"name:pNext;pointer:unique" json:"next"`
}

func (o *CallbackList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CallbackList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	for i1 := range o.PortName {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.PortName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PortName); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Number {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.Number[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Number); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CallbackList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CallbackList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	o.PortName = make([]uint16, 17)
	for i1 := range o.PortName {
		i1 := i1
		if err := w.ReadData(&o.PortName[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]uint16, 129)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.Number = make([]uint16, 129)
	for i1 := range o.Number {
		i1 := i1
		if err := w.ReadData(&o.Number[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &CallbackList{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**CallbackList) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	return nil
}

// StringList structure represents RASRPC_STRINGLIST RPC structure.
type StringList struct {
	String []uint16    `idl:"name:psz" json:"string"`
	Next   *StringList `idl:"name:pNext;pointer:unique" json:"next"`
}

func (o *StringList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *StringList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	for i1 := range o.String {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.String[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.String); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&StringList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StringList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	o.String = make([]uint16, 256)
	for i1 := range o.String {
		i1 := i1
		if err := w.ReadData(&o.String[i1]); err != nil {
			return err
		}
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &StringList{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**StringList) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	return nil
}

// Locationlist structure represents RASRPC_LOCATIONLIST RPC structure.
type Locationlist struct {
	LocationID uint32        `idl:"name:dwLocationId" json:"location_id"`
	Prefix     uint32        `idl:"name:iPrefix" json:"prefix"`
	Suffix     uint32        `idl:"name:iSuffix" json:"suffix"`
	Next       *Locationlist `idl:"name:pNext;pointer:unique" json:"next"`
}

func (o *Locationlist) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Locationlist) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.LocationID); err != nil {
		return err
	}
	if err := w.WriteData(o.Prefix); err != nil {
		return err
	}
	if err := w.WriteData(o.Suffix); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Locationlist{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Locationlist) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocationID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Prefix); err != nil {
		return err
	}
	if err := w.ReadData(&o.Suffix); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Locationlist{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Locationlist) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	return nil
}

// User structure represents RASRPC_PBUSER RPC structure.
type User struct {
	OperatorDial             bool          `idl:"name:fOperatorDial" json:"operator_dial"`
	PreviewPhoneNumber       bool          `idl:"name:fPreviewPhoneNumber" json:"preview_phone_number"`
	UseLocation              bool          `idl:"name:fUseLocation" json:"use_location"`
	ShowLights               bool          `idl:"name:fShowLights" json:"show_lights"`
	ShowConnectStatus        bool          `idl:"name:fShowConnectStatus" json:"show_connect_status"`
	CloseOnDial              bool          `idl:"name:fCloseOnDial" json:"close_on_dial"`
	AllowLogonPhonebookEdits bool          `idl:"name:fAllowLogonPhonebookEdits" json:"allow_logon_phonebook_edits"`
	AllowLogonLocationEdits  bool          `idl:"name:fAllowLogonLocationEdits" json:"allow_logon_location_edits"`
	SkipConnectComplete      bool          `idl:"name:fSkipConnectComplete" json:"skip_connect_complete"`
	NewEntryWizard           bool          `idl:"name:fNewEntryWizard" json:"new_entry_wizard"`
	RedialAttempts           uint32        `idl:"name:dwRedialAttempts" json:"redial_attempts"`
	RedialSeconds            uint32        `idl:"name:dwRedialSeconds" json:"redial_seconds"`
	IdleDisconnectSeconds    uint32        `idl:"name:dwIdleDisconnectSeconds" json:"idle_disconnect_seconds"`
	RedialOnLinkFailure      bool          `idl:"name:fRedialOnLinkFailure" json:"redial_on_link_failure"`
	PopupOnTopWhenRedialing  bool          `idl:"name:fPopupOnTopWhenRedialing" json:"popup_on_top_when_redialing"`
	ExpandAutoDialQuery      bool          `idl:"name:fExpandAutoDialQuery" json:"expand_auto_dial_query"`
	CallbackMode             uint32        `idl:"name:dwCallbackMode" json:"callback_mode"`
	Callbacks                *CallbackList `idl:"name:pCallbacks;pointer:unique" json:"callbacks"`
	LastCallbackByCaller     []uint16      `idl:"name:pszLastCallbackByCaller" json:"last_callback_by_caller"`
	PhonebookMode            uint32        `idl:"name:dwPhonebookMode" json:"phonebook_mode"`
	PersonalFile             []uint16      `idl:"name:pszPersonalFile" json:"personal_file"`
	AlternatePath            []uint16      `idl:"name:pszAlternatePath" json:"alternate_path"`
	Phonebooks               *StringList   `idl:"name:pPhonebooks;pointer:unique" json:"phonebooks"`
	AreaCodes                *StringList   `idl:"name:pAreaCodes;pointer:unique" json:"area_codes"`
	UseAreaAndCountry        bool          `idl:"name:fUseAreaAndCountry" json:"use_area_and_country"`
	Prefixes                 *StringList   `idl:"name:pPrefixes;pointer:unique" json:"prefixes"`
	Suffixes                 *StringList   `idl:"name:pSuffixes;pointer:unique" json:"suffixes"`
	Locations                *Locationlist `idl:"name:pLocations;pointer:unique" json:"locations"`
	XPhonebook               uint32        `idl:"name:dwXPhonebook" json:"x_phonebook"`
	YPhonebook               uint32        `idl:"name:dwYPhonebook" json:"y_phonebook"`
	DefaultEntry             []uint16      `idl:"name:pszDefaultEntry" json:"default_entry"`
	Initialized              bool          `idl:"name:fInitialized" json:"initialized"`
	Dirty                    bool          `idl:"name:fDirty" json:"dirty"`
}

func (o *User) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *User) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if !o.OperatorDial {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.PreviewPhoneNumber {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.UseLocation {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.ShowLights {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.ShowConnectStatus {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.CloseOnDial {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.AllowLogonPhonebookEdits {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.AllowLogonLocationEdits {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.SkipConnectComplete {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.NewEntryWizard {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RedialAttempts); err != nil {
		return err
	}
	if err := w.WriteData(o.RedialSeconds); err != nil {
		return err
	}
	if err := w.WriteData(o.IdleDisconnectSeconds); err != nil {
		return err
	}
	if !o.RedialOnLinkFailure {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.PopupOnTopWhenRedialing {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.ExpandAutoDialQuery {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CallbackMode); err != nil {
		return err
	}
	if o.Callbacks != nil {
		_ptr_pCallbacks := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Callbacks != nil {
				if err := o.Callbacks.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CallbackList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Callbacks, _ptr_pCallbacks); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	for i1 := range o.LastCallbackByCaller {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.LastCallbackByCaller[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LastCallbackByCaller); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PhonebookMode); err != nil {
		return err
	}
	for i1 := range o.PersonalFile {
		i1 := i1
		if uint64(i1) >= 260 {
			break
		}
		if err := w.WriteData(o.PersonalFile[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PersonalFile); uint64(i1) < 260; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.AlternatePath {
		i1 := i1
		if uint64(i1) >= 260 {
			break
		}
		if err := w.WriteData(o.AlternatePath[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AlternatePath); uint64(i1) < 260; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if o.Phonebooks != nil {
		_ptr_pPhonebooks := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Phonebooks != nil {
				if err := o.Phonebooks.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&StringList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Phonebooks, _ptr_pPhonebooks); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.AreaCodes != nil {
		_ptr_pAreaCodes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AreaCodes != nil {
				if err := o.AreaCodes.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&StringList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AreaCodes, _ptr_pAreaCodes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.UseAreaAndCountry {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.Prefixes != nil {
		_ptr_pPrefixes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Prefixes != nil {
				if err := o.Prefixes.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&StringList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Prefixes, _ptr_pPrefixes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Suffixes != nil {
		_ptr_pSuffixes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Suffixes != nil {
				if err := o.Suffixes.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&StringList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Suffixes, _ptr_pSuffixes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Locations != nil {
		_ptr_pLocations := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Locations != nil {
				if err := o.Locations.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Locationlist{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Locations, _ptr_pLocations); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.XPhonebook); err != nil {
		return err
	}
	if err := w.WriteData(o.YPhonebook); err != nil {
		return err
	}
	for i1 := range o.DefaultEntry {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.DefaultEntry[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DefaultEntry); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if !o.Initialized {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.Dirty {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *User) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	var _bOperatorDial int32
	if err := w.ReadData(&_bOperatorDial); err != nil {
		return err
	}
	o.OperatorDial = _bOperatorDial != 0
	var _bPreviewPhoneNumber int32
	if err := w.ReadData(&_bPreviewPhoneNumber); err != nil {
		return err
	}
	o.PreviewPhoneNumber = _bPreviewPhoneNumber != 0
	var _bUseLocation int32
	if err := w.ReadData(&_bUseLocation); err != nil {
		return err
	}
	o.UseLocation = _bUseLocation != 0
	var _bShowLights int32
	if err := w.ReadData(&_bShowLights); err != nil {
		return err
	}
	o.ShowLights = _bShowLights != 0
	var _bShowConnectStatus int32
	if err := w.ReadData(&_bShowConnectStatus); err != nil {
		return err
	}
	o.ShowConnectStatus = _bShowConnectStatus != 0
	var _bCloseOnDial int32
	if err := w.ReadData(&_bCloseOnDial); err != nil {
		return err
	}
	o.CloseOnDial = _bCloseOnDial != 0
	var _bAllowLogonPhonebookEdits int32
	if err := w.ReadData(&_bAllowLogonPhonebookEdits); err != nil {
		return err
	}
	o.AllowLogonPhonebookEdits = _bAllowLogonPhonebookEdits != 0
	var _bAllowLogonLocationEdits int32
	if err := w.ReadData(&_bAllowLogonLocationEdits); err != nil {
		return err
	}
	o.AllowLogonLocationEdits = _bAllowLogonLocationEdits != 0
	var _bSkipConnectComplete int32
	if err := w.ReadData(&_bSkipConnectComplete); err != nil {
		return err
	}
	o.SkipConnectComplete = _bSkipConnectComplete != 0
	var _bNewEntryWizard int32
	if err := w.ReadData(&_bNewEntryWizard); err != nil {
		return err
	}
	o.NewEntryWizard = _bNewEntryWizard != 0
	if err := w.ReadData(&o.RedialAttempts); err != nil {
		return err
	}
	if err := w.ReadData(&o.RedialSeconds); err != nil {
		return err
	}
	if err := w.ReadData(&o.IdleDisconnectSeconds); err != nil {
		return err
	}
	var _bRedialOnLinkFailure int32
	if err := w.ReadData(&_bRedialOnLinkFailure); err != nil {
		return err
	}
	o.RedialOnLinkFailure = _bRedialOnLinkFailure != 0
	var _bPopupOnTopWhenRedialing int32
	if err := w.ReadData(&_bPopupOnTopWhenRedialing); err != nil {
		return err
	}
	o.PopupOnTopWhenRedialing = _bPopupOnTopWhenRedialing != 0
	var _bExpandAutoDialQuery int32
	if err := w.ReadData(&_bExpandAutoDialQuery); err != nil {
		return err
	}
	o.ExpandAutoDialQuery = _bExpandAutoDialQuery != 0
	if err := w.ReadData(&o.CallbackMode); err != nil {
		return err
	}
	_ptr_pCallbacks := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Callbacks == nil {
			o.Callbacks = &CallbackList{}
		}
		if err := o.Callbacks.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pCallbacks := func(ptr interface{}) { o.Callbacks = *ptr.(**CallbackList) }
	if err := w.ReadPointer(&o.Callbacks, _s_pCallbacks, _ptr_pCallbacks); err != nil {
		return err
	}
	o.LastCallbackByCaller = make([]uint16, 129)
	for i1 := range o.LastCallbackByCaller {
		i1 := i1
		if err := w.ReadData(&o.LastCallbackByCaller[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PhonebookMode); err != nil {
		return err
	}
	o.PersonalFile = make([]uint16, 260)
	for i1 := range o.PersonalFile {
		i1 := i1
		if err := w.ReadData(&o.PersonalFile[i1]); err != nil {
			return err
		}
	}
	o.AlternatePath = make([]uint16, 260)
	for i1 := range o.AlternatePath {
		i1 := i1
		if err := w.ReadData(&o.AlternatePath[i1]); err != nil {
			return err
		}
	}
	_ptr_pPhonebooks := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Phonebooks == nil {
			o.Phonebooks = &StringList{}
		}
		if err := o.Phonebooks.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pPhonebooks := func(ptr interface{}) { o.Phonebooks = *ptr.(**StringList) }
	if err := w.ReadPointer(&o.Phonebooks, _s_pPhonebooks, _ptr_pPhonebooks); err != nil {
		return err
	}
	_ptr_pAreaCodes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AreaCodes == nil {
			o.AreaCodes = &StringList{}
		}
		if err := o.AreaCodes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pAreaCodes := func(ptr interface{}) { o.AreaCodes = *ptr.(**StringList) }
	if err := w.ReadPointer(&o.AreaCodes, _s_pAreaCodes, _ptr_pAreaCodes); err != nil {
		return err
	}
	var _bUseAreaAndCountry int32
	if err := w.ReadData(&_bUseAreaAndCountry); err != nil {
		return err
	}
	o.UseAreaAndCountry = _bUseAreaAndCountry != 0
	_ptr_pPrefixes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Prefixes == nil {
			o.Prefixes = &StringList{}
		}
		if err := o.Prefixes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pPrefixes := func(ptr interface{}) { o.Prefixes = *ptr.(**StringList) }
	if err := w.ReadPointer(&o.Prefixes, _s_pPrefixes, _ptr_pPrefixes); err != nil {
		return err
	}
	_ptr_pSuffixes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Suffixes == nil {
			o.Suffixes = &StringList{}
		}
		if err := o.Suffixes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pSuffixes := func(ptr interface{}) { o.Suffixes = *ptr.(**StringList) }
	if err := w.ReadPointer(&o.Suffixes, _s_pSuffixes, _ptr_pSuffixes); err != nil {
		return err
	}
	_ptr_pLocations := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Locations == nil {
			o.Locations = &Locationlist{}
		}
		if err := o.Locations.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pLocations := func(ptr interface{}) { o.Locations = *ptr.(**Locationlist) }
	if err := w.ReadPointer(&o.Locations, _s_pLocations, _ptr_pLocations); err != nil {
		return err
	}
	if err := w.ReadData(&o.XPhonebook); err != nil {
		return err
	}
	if err := w.ReadData(&o.YPhonebook); err != nil {
		return err
	}
	o.DefaultEntry = make([]uint16, 257)
	for i1 := range o.DefaultEntry {
		i1 := i1
		if err := w.ReadData(&o.DefaultEntry[i1]); err != nil {
			return err
		}
	}
	var _bInitialized int32
	if err := w.ReadData(&_bInitialized); err != nil {
		return err
	}
	o.Initialized = _bInitialized != 0
	var _bDirty int32
	if err := w.ReadData(&_bDirty); err != nil {
		return err
	}
	o.Dirty = _bDirty != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// RemoteSetDNSConfig structure represents IRemoteSetDnsConfig RPC structure.
type RemoteSetDNSConfig dcom.InterfacePointer

func (o *RemoteSetDNSConfig) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteSetDNSConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteSetDNSConfig) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteSetDNSConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteSetDNSConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteIPv6Config structure represents IRemoteIPV6Config RPC structure.
type RemoteIPv6Config dcom.InterfacePointer

func (o *RemoteIPv6Config) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteIPv6Config) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteIPv6Config) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteIPv6Config) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteIPv6Config) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteICFICSConfig structure represents IRemoteICFICSConfig RPC structure.
type RemoteICFICSConfig dcom.InterfacePointer

func (o *RemoteICFICSConfig) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteICFICSConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteICFICSConfig) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteICFICSConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteICFICSConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteSSTPCertCheck structure represents IRemoteSstpCertCheck RPC structure.
type RemoteSSTPCertCheck dcom.InterfacePointer

func (o *RemoteSSTPCertCheck) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteSSTPCertCheck) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteSSTPCertCheck) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteSSTPCertCheck) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteSSTPCertCheck) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteNetworkConfig structure represents IRemoteNetworkConfig RPC structure.
type RemoteNetworkConfig dcom.InterfacePointer

func (o *RemoteNetworkConfig) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteNetworkConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteNetworkConfig) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteNetworkConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteNetworkConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteStringIDConfig structure represents IRemoteStringIdConfig RPC structure.
type RemoteStringIDConfig dcom.InterfacePointer

func (o *RemoteStringIDConfig) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteStringIDConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteStringIDConfig) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteStringIDConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteStringIDConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteRouterRestart structure represents IRemoteRouterRestart RPC structure.
type RemoteRouterRestart dcom.InterfacePointer

func (o *RemoteRouterRestart) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteRouterRestart) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteRouterRestart) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteRouterRestart) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteRouterRestart) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

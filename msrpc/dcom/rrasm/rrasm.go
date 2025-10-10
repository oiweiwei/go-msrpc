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

// MaxSstpHashSize represents the MAX_SSTP_HASH_SIZE RPC constant
var MaxSstpHashSize = 32

// MaxGroupLength represents the MAX_GROUP_LEN RPC constant
var MaxGroupLength = 64

// MaxAdaptorNameLength represents the MAX_ADAPTOR_NAME_LEN RPC constant
var MaxAdaptorNameLength = 48

// RasrpcMaxEntryName represents the RASRPC_MaxEntryName RPC constant
var RasrpcMaxEntryName = 256

// RasrpcMaxPortName represents the RASRPC_MaxPortName RPC constant
var RasrpcMaxPortName = 16

// RasrpcMaxDeviceName represents the RASRPC_MaxDeviceName RPC constant
var RasrpcMaxDeviceName = 128

// RasrpcMaxPhoneNumber represents the RASRPC_MaxPhoneNumber RPC constant
var RasrpcMaxPhoneNumber = 128

// RasrpcMaxPath represents the RASRPC_MAX_PATH RPC constant
var RasrpcMaxPath = 260

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

// MIBIpforwardType type represents MIB_IPFORWARD_TYPE RPC enumeration.
type MIBIpforwardType uint16

var (
	MIBIpforwardTypeIprouteTypeOther    MIBIpforwardType = 1
	MIBIpforwardTypeIprouteTypeInvalid  MIBIpforwardType = 2
	MIBIpforwardTypeIprouteTypeDirect   MIBIpforwardType = 3
	MIBIpforwardTypeIprouteTypeIndirect MIBIpforwardType = 4
)

func (o MIBIpforwardType) String() string {
	switch o {
	case MIBIpforwardTypeIprouteTypeOther:
		return "MIBIpforwardTypeIprouteTypeOther"
	case MIBIpforwardTypeIprouteTypeInvalid:
		return "MIBIpforwardTypeIprouteTypeInvalid"
	case MIBIpforwardTypeIprouteTypeDirect:
		return "MIBIpforwardTypeIprouteTypeDirect"
	case MIBIpforwardTypeIprouteTypeIndirect:
		return "MIBIpforwardTypeIprouteTypeIndirect"
	}
	return "Invalid"
}

// MIBIpforwardProto type represents MIB_IPFORWARD_PROTO RPC enumeration.
type MIBIpforwardProto uint16

var (
	MIBIpforwardProtoIPProtoOther          MIBIpforwardProto = 1
	MIBIpforwardProtoIPProtoLocal          MIBIpforwardProto = 2
	MIBIpforwardProtoIPProtoNetmgmt        MIBIpforwardProto = 3
	MIBIpforwardProtoIPProtoICMP           MIBIpforwardProto = 4
	MIBIpforwardProtoIPProtoEgp            MIBIpforwardProto = 5
	MIBIpforwardProtoIPProtoGgp            MIBIpforwardProto = 6
	MIBIpforwardProtoIPProtoHello          MIBIpforwardProto = 7
	MIBIpforwardProtoIPProtoRip            MIBIpforwardProto = 8
	MIBIpforwardProtoIPProtoIsIs           MIBIpforwardProto = 9
	MIBIpforwardProtoIPProtoEsIs           MIBIpforwardProto = 10
	MIBIpforwardProtoIPProtoCisco          MIBIpforwardProto = 11
	MIBIpforwardProtoIPProtoBbn            MIBIpforwardProto = 12
	MIBIpforwardProtoIPProtoOspf           MIBIpforwardProto = 13
	MIBIpforwardProtoIPProtoBgp            MIBIpforwardProto = 14
	MIBIpforwardProtoIPProtoNTAutostatic   MIBIpforwardProto = 10002
	MIBIpforwardProtoIPProtoNTStatic       MIBIpforwardProto = 10006
	MIBIpforwardProtoIPProtoNTStaticNonDod MIBIpforwardProto = 10007
)

func (o MIBIpforwardProto) String() string {
	switch o {
	case MIBIpforwardProtoIPProtoOther:
		return "MIBIpforwardProtoIPProtoOther"
	case MIBIpforwardProtoIPProtoLocal:
		return "MIBIpforwardProtoIPProtoLocal"
	case MIBIpforwardProtoIPProtoNetmgmt:
		return "MIBIpforwardProtoIPProtoNetmgmt"
	case MIBIpforwardProtoIPProtoICMP:
		return "MIBIpforwardProtoIPProtoICMP"
	case MIBIpforwardProtoIPProtoEgp:
		return "MIBIpforwardProtoIPProtoEgp"
	case MIBIpforwardProtoIPProtoGgp:
		return "MIBIpforwardProtoIPProtoGgp"
	case MIBIpforwardProtoIPProtoHello:
		return "MIBIpforwardProtoIPProtoHello"
	case MIBIpforwardProtoIPProtoRip:
		return "MIBIpforwardProtoIPProtoRip"
	case MIBIpforwardProtoIPProtoIsIs:
		return "MIBIpforwardProtoIPProtoIsIs"
	case MIBIpforwardProtoIPProtoEsIs:
		return "MIBIpforwardProtoIPProtoEsIs"
	case MIBIpforwardProtoIPProtoCisco:
		return "MIBIpforwardProtoIPProtoCisco"
	case MIBIpforwardProtoIPProtoBbn:
		return "MIBIpforwardProtoIPProtoBbn"
	case MIBIpforwardProtoIPProtoOspf:
		return "MIBIpforwardProtoIPProtoOspf"
	case MIBIpforwardProtoIPProtoBgp:
		return "MIBIpforwardProtoIPProtoBgp"
	case MIBIpforwardProtoIPProtoNTAutostatic:
		return "MIBIpforwardProtoIPProtoNTAutostatic"
	case MIBIpforwardProtoIPProtoNTStatic:
		return "MIBIpforwardProtoIPProtoNTStatic"
	case MIBIpforwardProtoIPProtoNTStaticNonDod:
		return "MIBIpforwardProtoIPProtoNTStaticNonDod"
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
	MibtcpStateClosed    MIBTCPState = 1
	MibtcpStateListen    MIBTCPState = 2
	MibtcpStateSynSent   MIBTCPState = 3
	MibtcpStateSynRcvd   MIBTCPState = 4
	MibtcpStateEstab     MIBTCPState = 5
	MibtcpStateInWait1   MIBTCPState = 6
	MibtcpStateInWait2   MIBTCPState = 7
	MibtcpStateCloseWait MIBTCPState = 8
	MibtcpStateClosing   MIBTCPState = 9
	MibtcpStateLastAck   MIBTCPState = 10
	MibtcpStateTimeWait  MIBTCPState = 11
	MibtcpStateDeleteTcb MIBTCPState = 12
)

func (o MIBTCPState) String() string {
	switch o {
	case MibtcpStateClosed:
		return "MibtcpStateClosed"
	case MibtcpStateListen:
		return "MibtcpStateListen"
	case MibtcpStateSynSent:
		return "MibtcpStateSynSent"
	case MibtcpStateSynRcvd:
		return "MibtcpStateSynRcvd"
	case MibtcpStateEstab:
		return "MibtcpStateEstab"
	case MibtcpStateInWait1:
		return "MibtcpStateInWait1"
	case MibtcpStateInWait2:
		return "MibtcpStateInWait2"
	case MibtcpStateCloseWait:
		return "MibtcpStateCloseWait"
	case MibtcpStateClosing:
		return "MibtcpStateClosing"
	case MibtcpStateLastAck:
		return "MibtcpStateLastAck"
	case MibtcpStateTimeWait:
		return "MibtcpStateTimeWait"
	case MibtcpStateDeleteTcb:
		return "MibtcpStateDeleteTcb"
	}
	return "Invalid"
}

// TCPRtoAlgorithm type represents TCP_RTO_ALGORITHM RPC enumeration.
type TCPRtoAlgorithm uint16

var (
	TCPRtoAlgorithmMibtcpRtoOther    TCPRtoAlgorithm = 1
	TCPRtoAlgorithmMibtcpRtoConstant TCPRtoAlgorithm = 2
	TCPRtoAlgorithmMibtcpRtoRsre     TCPRtoAlgorithm = 3
	TCPRtoAlgorithmMibtcpRtoVanj     TCPRtoAlgorithm = 4
)

func (o TCPRtoAlgorithm) String() string {
	switch o {
	case TCPRtoAlgorithmMibtcpRtoOther:
		return "TCPRtoAlgorithmMibtcpRtoOther"
	case TCPRtoAlgorithmMibtcpRtoConstant:
		return "TCPRtoAlgorithmMibtcpRtoConstant"
	case TCPRtoAlgorithmMibtcpRtoRsre:
		return "TCPRtoAlgorithmMibtcpRtoRsre"
	case TCPRtoAlgorithmMibtcpRtoVanj:
		return "TCPRtoAlgorithmMibtcpRtoVanj"
	}
	return "Invalid"
}

// In6Addr structure represents IN6_ADDR RPC structure.
type In6Addr struct {
	Union *In6Addr_Union `idl:"name:u" json:"union"`
}

func (o *In6Addr) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *In6Addr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	// FIXME unknown type u
	return nil
}
func (o *In6Addr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	// FIXME: unknown type u
	return nil
}

type In6Addr_Union struct {
	Byte []byte   `idl:"name:Byte" json:"byte"`
	Word []uint16 `idl:"name:Word" json:"word"`
}

// DimInformationContainer structure represents DIM_INFORMATION_CONTAINER RPC structure.
type DimInformationContainer struct {
	BufferSize uint32 `idl:"name:dwBufferSize" json:"buffer_size"`
	Buffer     []byte `idl:"name:pBuffer;size_is:(dwBufferSize)" json:"buffer"`
}

func (o *DimInformationContainer) xxx_PreparePayload(ctx context.Context) error {
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
func (o *DimInformationContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DimInformationContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MprapiObjectHeaderIDL structure represents MPRAPI_OBJECT_HEADER_IDL RPC structure.
type MprapiObjectHeaderIDL struct {
	Revision uint8  `idl:"name:revision" json:"revision"`
	Type     uint8  `idl:"name:type" json:"type"`
	Size     uint16 `idl:"name:size" json:"size"`
}

func (o *MprapiObjectHeaderIDL) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprapiObjectHeaderIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MprapiObjectHeaderIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppProjectionInfo1 structure represents PPP_PROJECTION_INFO_1 RPC structure.
type PppProjectionInfo1 struct {
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
	LcpError                     uint32   `idl:"name:dwLcpError" json:"lcp_error"`
	AuthenticationProtocol       uint32   `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	AuthenticationData           uint32   `idl:"name:dwAuthenticationData" json:"authentication_data"`
	RemoteAuthenticationProtocol uint32   `idl:"name:dwRemoteAuthenticationProtocol" json:"remote_authentication_protocol"`
	RemoteAuthenticationData     uint32   `idl:"name:dwRemoteAuthenticationData" json:"remote_authentication_data"`
	LcpTerminateReason           uint32   `idl:"name:dwLcpTerminateReason" json:"lcp_terminate_reason"`
	LcpRemoteTerminateReason     uint32   `idl:"name:dwLcpRemoteTerminateReason" json:"lcp_remote_terminate_reason"`
	LcpOptions                   uint32   `idl:"name:dwLcpOptions" json:"lcp_options"`
	LcpRemoteOptions             uint32   `idl:"name:dwLcpRemoteOptions" json:"lcp_remote_options"`
	EapTypeID                    uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
	RemoteEapTypeID              uint32   `idl:"name:dwRemoteEapTypeId" json:"remote_eap_type_id"`
	CcpError                     uint32   `idl:"name:dwCcpError" json:"ccp_error"`
	CompressionAlgorithm         uint32   `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	CcpOptions                   uint32   `idl:"name:dwCcpOptions" json:"ccp_options"`
	RemoteCompressionAlgorithm   uint32   `idl:"name:dwRemoteCompressionAlgorithm" json:"remote_compression_algorithm"`
	CcpRemoteOptions             uint32   `idl:"name:dwCcpRemoteOptions" json:"ccp_remote_options"`
}

func (o *PppProjectionInfo1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppProjectionInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.LcpError); err != nil {
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
	if err := w.WriteData(o.LcpTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LcpRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LcpOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.LcpRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.EapTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteEapTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.CcpError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CcpOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CcpRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *PppProjectionInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.LcpError); err != nil {
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
	if err := w.ReadData(&o.LcpTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LcpRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LcpOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.LcpRemoteOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.EapTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteEapTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CcpError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CcpOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CcpRemoteOptions); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// PppProjectionInfo2 structure represents PPP_PROJECTION_INFO_2 RPC structure.
type PppProjectionInfo2 struct {
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
	LcpError                     uint32   `idl:"name:dwLcpError" json:"lcp_error"`
	AuthenticationProtocol       uint32   `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	AuthenticationData           uint32   `idl:"name:dwAuthenticationData" json:"authentication_data"`
	RemoteAuthenticationProtocol uint32   `idl:"name:dwRemoteAuthenticationProtocol" json:"remote_authentication_protocol"`
	RemoteAuthenticationData     uint32   `idl:"name:dwRemoteAuthenticationData" json:"remote_authentication_data"`
	LcpTerminateReason           uint32   `idl:"name:dwLcpTerminateReason" json:"lcp_terminate_reason"`
	LcpRemoteTerminateReason     uint32   `idl:"name:dwLcpRemoteTerminateReason" json:"lcp_remote_terminate_reason"`
	LcpOptions                   uint32   `idl:"name:dwLcpOptions" json:"lcp_options"`
	LcpRemoteOptions             uint32   `idl:"name:dwLcpRemoteOptions" json:"lcp_remote_options"`
	EapTypeID                    uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
	EmbeddedEapTypeID            uint32   `idl:"name:dwEmbeddedEAPTypeId" json:"embedded_eap_type_id"`
	RemoteEapTypeID              uint32   `idl:"name:dwRemoteEapTypeId" json:"remote_eap_type_id"`
	CcpError                     uint32   `idl:"name:dwCcpError" json:"ccp_error"`
	CompressionAlgorithm         uint32   `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	CcpOptions                   uint32   `idl:"name:dwCcpOptions" json:"ccp_options"`
	RemoteCompressionAlgorithm   uint32   `idl:"name:dwRemoteCompressionAlgorithm" json:"remote_compression_algorithm"`
	CcpRemoteOptions             uint32   `idl:"name:dwCcpRemoteOptions" json:"ccp_remote_options"`
}

func (o *PppProjectionInfo2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppProjectionInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.LcpError); err != nil {
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
	if err := w.WriteData(o.LcpTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LcpRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.WriteData(o.LcpOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.LcpRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.EapTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.EmbeddedEapTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteEapTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.CcpError); err != nil {
		return err
	}
	if err := w.WriteData(o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CcpOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.WriteData(o.CcpRemoteOptions); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *PppProjectionInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.LcpError); err != nil {
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
	if err := w.ReadData(&o.LcpTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LcpRemoteTerminateReason); err != nil {
		return err
	}
	if err := w.ReadData(&o.LcpOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.LcpRemoteOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.EapTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.EmbeddedEapTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteEapTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CcpError); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CcpOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteCompressionAlgorithm); err != nil {
		return err
	}
	if err := w.ReadData(&o.CcpRemoteOptions); err != nil {
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
	EapTypeID              uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
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
	if err := w.WriteData(o.EapTypeID); err != nil {
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
	if err := w.ReadData(&o.EapTypeID); err != nil {
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
	EapTypeID              uint32   `idl:"name:dwEapTypeId" json:"eap_type_id"`
	EmbeddedEapTypeID      uint32   `idl:"name:dwEmbeddedEAPTypeId" json:"embedded_eap_type_id"`
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
	if err := w.WriteData(o.EapTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.EmbeddedEapTypeID); err != nil {
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
	if err := w.ReadData(&o.EapTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.EmbeddedEapTypeID); err != nil {
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
	// *ProjectionInfoIDL1_PppProjectionInfo
	// *ProjectionInfoIDL1_IKEv2ProjectionInfo
	Value is_ProjectionInfoIDL1 `json:"value"`
}

func (o *ProjectionInfoIDL1) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ProjectionInfoIDL1_PppProjectionInfo:
		if value != nil {
			return value.PppProjectionInfo
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
		_o, _ := o.Value.(*ProjectionInfoIDL1_PppProjectionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProjectionInfoIDL1_PppProjectionInfo{}).MarshalNDR(ctx, w); err != nil {
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
		o.Value = &ProjectionInfoIDL1_PppProjectionInfo{}
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

// ProjectionInfoIDL1_PppProjectionInfo structure represents PROJECTION_INFO_IDL_1 RPC union arm.
//
// It has following labels: 1
type ProjectionInfoIDL1_PppProjectionInfo struct {
	PppProjectionInfo *PppProjectionInfo1 `idl:"name:PppProjectionInfo" json:"ppp_projection_info"`
}

func (*ProjectionInfoIDL1_PppProjectionInfo) is_ProjectionInfoIDL1() {}

func (o *ProjectionInfoIDL1_PppProjectionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PppProjectionInfo != nil {
		if err := o.PppProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppProjectionInfo1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProjectionInfoIDL1_PppProjectionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PppProjectionInfo == nil {
		o.PppProjectionInfo = &PppProjectionInfo1{}
	}
	if err := o.PppProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
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
	// *ProjectionInfoIDL2_PppProjectionInfo
	// *ProjectionInfoIDL2_IKEv2ProjectionInfo
	Value is_ProjectionInfoIDL2 `json:"value"`
}

func (o *ProjectionInfoIDL2) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ProjectionInfoIDL2_PppProjectionInfo:
		if value != nil {
			return value.PppProjectionInfo
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
		_o, _ := o.Value.(*ProjectionInfoIDL2_PppProjectionInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProjectionInfoIDL2_PppProjectionInfo{}).MarshalNDR(ctx, w); err != nil {
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
		o.Value = &ProjectionInfoIDL2_PppProjectionInfo{}
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

// ProjectionInfoIDL2_PppProjectionInfo structure represents PROJECTION_INFO_IDL_2 RPC union arm.
//
// It has following labels: 1
type ProjectionInfoIDL2_PppProjectionInfo struct {
	PppProjectionInfo *PppProjectionInfo2 `idl:"name:PppProjectionInfo" json:"ppp_projection_info"`
}

func (*ProjectionInfoIDL2_PppProjectionInfo) is_ProjectionInfoIDL2() {}

func (o *ProjectionInfoIDL2_PppProjectionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PppProjectionInfo != nil {
		if err := o.PppProjectionInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppProjectionInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProjectionInfoIDL2_PppProjectionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PppProjectionInfo == nil {
		o.PppProjectionInfo = &PppProjectionInfo2{}
	}
	if err := o.PppProjectionInfo.UnmarshalNDR(ctx, w); err != nil {
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
	Header                *MprapiObjectHeaderIDL `idl:"name:Header" json:"header"`
	ConnectDuration       uint32                 `idl:"name:dwConnectDuration" json:"connect_duration"`
	InterfaceType         RouterInterfaceType    `idl:"name:dwInterfaceType" json:"interface_type"`
	ConnectionFlags       uint32                 `idl:"name:dwConnectionFlags" json:"connection_flags"`
	InterfaceName         []uint16               `idl:"name:wszInterfaceName" json:"interface_name"`
	UserName              []uint16               `idl:"name:wszUserName" json:"user_name"`
	LogonDomain           []uint16               `idl:"name:wszLogonDomain" json:"logon_domain"`
	RemoteComputer        []uint16               `idl:"name:wszRemoteComputer" json:"remote_computer"`
	GUID                  *dtyp.GUID             `idl:"name:guid" json:"guid"`
	RASQuarantineState    RASQuarantineState     `idl:"name:rasQuarState" json:"ras_quarantine_state"`
	ProbationTime         *dtyp.Filetime         `idl:"name:probationTime" json:"probation_time"`
	BytesXmited           uint32                 `idl:"name:dwBytesXmited" json:"bytes_xmited"`
	BytesRcved            uint32                 `idl:"name:dwBytesRcved" json:"bytes_rcved"`
	FramesXmited          uint32                 `idl:"name:dwFramesXmited" json:"frames_xmited"`
	FramesRcved           uint32                 `idl:"name:dwFramesRcved" json:"frames_rcved"`
	CRCError              uint32                 `idl:"name:dwCrcErr" json:"crc_error"`
	TimeoutError          uint32                 `idl:"name:dwTimeoutErr" json:"timeout_error"`
	AlignmentError        uint32                 `idl:"name:dwAlignmentErr" json:"alignment_error"`
	HardwareOverrunError  uint32                 `idl:"name:dwHardwareOverrunErr" json:"hardware_overrun_error"`
	FramingError          uint32                 `idl:"name:dwFramingErr" json:"framing_error"`
	BufferOverrunError    uint32                 `idl:"name:dwBufferOverrunErr" json:"buffer_overrun_error"`
	CompressionRatioIn    uint32                 `idl:"name:dwCompressionRatioIn" json:"compression_ratio_in"`
	CompressionRatioOut   uint32                 `idl:"name:dwCompressionRatioOut" json:"compression_ratio_out"`
	SwitchOversLength     uint32                 `idl:"name:dwNumSwitchOvers" json:"switch_overs_length"`
	RemoteEndpointAddress []uint16               `idl:"name:wszRemoteEndpointAddress" json:"remote_endpoint_address"`
	LocalEndpointAddress  []uint16               `idl:"name:wszLocalEndpointAddress" json:"local_endpoint_address"`
	ProjectionInfo        *ProjectionInfoIDL1    `idl:"name:ProjectionInfo" json:"projection_info"`
	HConnection           uint32                 `idl:"name:hConnection" json:"h_connection"`
	HInterface            uint32                 `idl:"name:hInterface" json:"h_interface"`
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
	if err := w.WriteData(o.HConnection); err != nil {
		return err
	}
	if err := w.WriteData(o.HInterface); err != nil {
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
		o.Header = &MprapiObjectHeaderIDL{}
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
	if err := w.ReadData(&o.HConnection); err != nil {
		return err
	}
	if err := w.ReadData(&o.HInterface); err != nil {
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
	HConnection           uint32              `idl:"name:hConnection" json:"h_connection"`
	HInterface            uint32              `idl:"name:hInterface" json:"h_interface"`
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
	if err := w.WriteData(o.HConnection); err != nil {
		return err
	}
	if err := w.WriteData(o.HInterface); err != nil {
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
	if err := w.ReadData(&o.HConnection); err != nil {
		return err
	}
	if err := w.ReadData(&o.HInterface); err != nil {
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
	IsEkuoid bool   `idl:"name:IsEKUOID" json:"is_ekuoid"`
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
	if !o.IsEkuoid {
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
	var _bIsEkuoid int32
	if err := w.ReadData(&_bIsEkuoid); err != nil {
		return err
	}
	o.IsEkuoid = _bIsEkuoid != 0
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

// RouterCustomL2tpPolicy0 structure represents ROUTER_CUSTOM_L2TP_POLICY_0 RPC structure.
type RouterCustomL2tpPolicy0 struct {
	IntegrityMethod         uint32 `idl:"name:dwIntegrityMethod" json:"integrity_method"`
	EncryptionMethod        uint32 `idl:"name:dwEncryptionMethod" json:"encryption_method"`
	CipherTransformConstant uint32 `idl:"name:dwCipherTransformConstant" json:"cipher_transform_constant"`
	AuthTransformConstant   uint32 `idl:"name:dwAuthTransformConstant" json:"auth_transform_constant"`
	PFSGroup                uint32 `idl:"name:dwPfsGroup" json:"pfs_group"`
	DHGroup                 uint32 `idl:"name:dwDhGroup" json:"dh_group"`
}

func (o *RouterCustomL2tpPolicy0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RouterCustomL2tpPolicy0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RouterCustomL2tpPolicy0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MprInterfaceCustominfoex0 structure represents MPR_IF_CUSTOMINFOEX_0 RPC structure.
type MprInterfaceCustominfoex0 struct {
	Header            *MprapiObjectHeaderIDL             `idl:"name:Header" json:"header"`
	Flags             uint32                             `idl:"name:dwFlags" json:"flags"`
	CustomIKEv2Config *RouterIKEv2InterfaceCustomConfig0 `idl:"name:customIkev2Config" json:"custom_ikev2_config"`
}

func (o *MprInterfaceCustominfoex0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprInterfaceCustominfoex0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MprInterfaceCustominfoex0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
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

// MprInterfaceCustominfoex1 structure represents MPR_IF_CUSTOMINFOEX_1 RPC structure.
type MprInterfaceCustominfoex1 struct {
	Header            *MprapiObjectHeaderIDL             `idl:"name:Header" json:"header"`
	Flags             uint32                             `idl:"name:dwFlags" json:"flags"`
	CustomIKEv2Config *RouterIKEv2InterfaceCustomConfig1 `idl:"name:customIkev2Config" json:"custom_ikev2_config"`
}

func (o *MprInterfaceCustominfoex1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprInterfaceCustominfoex1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MprInterfaceCustominfoex1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
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

// MprInterfaceCustominfoexIDL structure represents MPR_IF_CUSTOMINFOEX_IDL RPC union.
type MprInterfaceCustominfoexIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *MprInterfaceCustominfoexIDL_InterfaceConfigObj1
	// *MprInterfaceCustominfoexIDL_InterfaceConfigObj2
	Value is_MprInterfaceCustominfoexIDL `json:"value"`
}

func (o *MprInterfaceCustominfoexIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MprInterfaceCustominfoexIDL_InterfaceConfigObj1:
		if value != nil {
			return value.InterfaceConfigObj1
		}
	case *MprInterfaceCustominfoexIDL_InterfaceConfigObj2:
		if value != nil {
			return value.InterfaceConfigObj2
		}
	}
	return nil
}

type is_MprInterfaceCustominfoexIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MprInterfaceCustominfoexIDL()
}

func (o *MprInterfaceCustominfoexIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		_o, _ := o.Value.(*MprInterfaceCustominfoexIDL_InterfaceConfigObj1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprInterfaceCustominfoexIDL_InterfaceConfigObj1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*MprInterfaceCustominfoexIDL_InterfaceConfigObj2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprInterfaceCustominfoexIDL_InterfaceConfigObj2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *MprInterfaceCustominfoexIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Value = &MprInterfaceCustominfoexIDL_InterfaceConfigObj1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &MprInterfaceCustominfoexIDL_InterfaceConfigObj2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// MprInterfaceCustominfoexIDL_InterfaceConfigObj1 structure represents MPR_IF_CUSTOMINFOEX_IDL RPC union arm.
//
// It has following labels: 1
type MprInterfaceCustominfoexIDL_InterfaceConfigObj1 struct {
	InterfaceConfigObj1 *MprInterfaceCustominfoex0 `idl:"name:IfConfigObj1" json:"interface_config_obj1"`
}

func (*MprInterfaceCustominfoexIDL_InterfaceConfigObj1) is_MprInterfaceCustominfoexIDL() {}

func (o *MprInterfaceCustominfoexIDL_InterfaceConfigObj1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.InterfaceConfigObj1 != nil {
		if err := o.InterfaceConfigObj1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprInterfaceCustominfoex0{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprInterfaceCustominfoexIDL_InterfaceConfigObj1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.InterfaceConfigObj1 == nil {
		o.InterfaceConfigObj1 = &MprInterfaceCustominfoex0{}
	}
	if err := o.InterfaceConfigObj1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprInterfaceCustominfoexIDL_InterfaceConfigObj2 structure represents MPR_IF_CUSTOMINFOEX_IDL RPC union arm.
//
// It has following labels: 2
type MprInterfaceCustominfoexIDL_InterfaceConfigObj2 struct {
	InterfaceConfigObj2 *MprInterfaceCustominfoex1 `idl:"name:IfConfigObj2" json:"interface_config_obj2"`
}

func (*MprInterfaceCustominfoexIDL_InterfaceConfigObj2) is_MprInterfaceCustominfoexIDL() {}

func (o *MprInterfaceCustominfoexIDL_InterfaceConfigObj2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.InterfaceConfigObj2 != nil {
		if err := o.InterfaceConfigObj2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprInterfaceCustominfoex1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprInterfaceCustominfoexIDL_InterfaceConfigObj2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.InterfaceConfigObj2 == nil {
		o.InterfaceConfigObj2 = &MprInterfaceCustominfoex1{}
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

// L2tpTunnelConfigParams1 structure represents L2TP_TUNNEL_CONFIG_PARAMS_1 RPC structure.
type L2tpTunnelConfigParams1 struct {
	IdleTimeout                uint32                   `idl:"name:dwIdleTimeout" json:"idle_timeout"`
	EncryptionType             uint32                   `idl:"name:dwEncryptionType" json:"encryption_type"`
	SALifeTime                 uint32                   `idl:"name:dwSaLifeTime" json:"sa_life_time"`
	SADataSizeForRenegotiation uint32                   `idl:"name:dwSaDataSizeForRenegotiation" json:"sa_data_size_for_renegotiation"`
	CustomPolicy               *RouterCustomL2tpPolicy0 `idl:"name:customPolicy" json:"custom_policy"`
}

func (o *L2tpTunnelConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *L2tpTunnelConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
				if err := (&RouterCustomL2tpPolicy0{}).MarshalNDR(ctx, w); err != nil {
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
func (o *L2tpTunnelConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
			o.CustomPolicy = &RouterCustomL2tpPolicy0{}
		}
		if err := o.CustomPolicy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_customPolicy := func(ptr interface{}) { o.CustomPolicy = *ptr.(**RouterCustomL2tpPolicy0) }
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

// L2tpConfigParams1 structure represents L2TP_CONFIG_PARAMS_1 RPC structure.
type L2tpConfigParams1 struct {
	PortsLength uint32 `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags   uint32 `idl:"name:dwPortFlags" json:"port_flags"`
}

func (o *L2tpConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *L2tpConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *L2tpConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// L2tpConfigParams2 structure represents L2TP_CONFIG_PARAMS_2 RPC structure.
type L2tpConfigParams2 struct {
	PortsLength            uint32                   `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags              uint32                   `idl:"name:dwPortFlags" json:"port_flags"`
	TunnelConfigParamFlags uint32                   `idl:"name:dwTunnelConfigParamFlags" json:"tunnel_config_param_flags"`
	TunnelConfigParams     *L2tpTunnelConfigParams1 `idl:"name:TunnelConfigParams" json:"tunnel_config_params"`
}

func (o *L2tpConfigParams2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *L2tpConfigParams2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&L2tpTunnelConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *L2tpConfigParams2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.TunnelConfigParams = &L2tpTunnelConfigParams1{}
	}
	if err := o.TunnelConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SstpCertInfo1 structure represents SSTP_CERT_INFO_1 RPC structure.
type SstpCertInfo1 struct {
	IsDefault bool       `idl:"name:isDefault" json:"is_default"`
	CertBlob  *CertBlob1 `idl:"name:certBlob" json:"cert_blob"`
}

func (o *SstpCertInfo1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SstpCertInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SstpCertInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// SstpConfigParams1 structure represents SSTP_CONFIG_PARAMS_1 RPC structure.
type SstpConfigParams1 struct {
	PortsLength     uint32         `idl:"name:dwNumPorts" json:"ports_length"`
	PortFlags       uint32         `idl:"name:dwPortFlags" json:"port_flags"`
	IsUseHTTPS      bool           `idl:"name:isUseHttps" json:"is_use_https"`
	CertAlgorithm   uint32         `idl:"name:certAlgorithm" json:"cert_algorithm"`
	SstpCertDetails *SstpCertInfo1 `idl:"name:sstpCertDetails" json:"sstp_cert_details"`
}

func (o *SstpConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SstpConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.SstpCertDetails != nil {
		if err := o.SstpCertDetails.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SstpCertInfo1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SstpConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.SstpCertDetails == nil {
		o.SstpCertDetails = &SstpCertInfo1{}
	}
	if err := o.SstpCertDetails.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprapiTunnelConfigParams1 structure represents MPRAPI_TUNNEL_CONFIG_PARAMS_1 RPC structure.
type MprapiTunnelConfigParams1 struct {
	IKEConfigParams  *IKEv2ConfigParams1 `idl:"name:IkeConfigParams" json:"ike_config_params"`
	PPTPConfigParams *PPTPConfigParams1  `idl:"name:PptpConfigParams" json:"pptp_config_params"`
	L2tpConfigParams *L2tpConfigParams1  `idl:"name:L2tpConfigParams" json:"l2tp_config_params"`
	SstpConfigParams *SstpConfigParams1  `idl:"name:SstpConfigParams" json:"sstp_config_params"`
}

func (o *MprapiTunnelConfigParams1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprapiTunnelConfigParams1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.L2tpConfigParams != nil {
		if err := o.L2tpConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&L2tpConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SstpConfigParams != nil {
		if err := o.SstpConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SstpConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprapiTunnelConfigParams1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.L2tpConfigParams == nil {
		o.L2tpConfigParams = &L2tpConfigParams1{}
	}
	if err := o.L2tpConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SstpConfigParams == nil {
		o.SstpConfigParams = &SstpConfigParams1{}
	}
	if err := o.SstpConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprapiTunnelConfigParams2 structure represents MPRAPI_TUNNEL_CONFIG_PARAMS_2 RPC structure.
type MprapiTunnelConfigParams2 struct {
	IKEConfigParams  *IKEv2ConfigParams2 `idl:"name:IkeConfigParams" json:"ike_config_params"`
	PPTPConfigParams *PPTPConfigParams1  `idl:"name:PptpConfigParams" json:"pptp_config_params"`
	L2tpConfigParams *L2tpConfigParams1  `idl:"name:L2tpConfigParams" json:"l2tp_config_params"`
	SstpConfigParams *SstpConfigParams1  `idl:"name:SstpConfigParams" json:"sstp_config_params"`
}

func (o *MprapiTunnelConfigParams2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprapiTunnelConfigParams2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.L2tpConfigParams != nil {
		if err := o.L2tpConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&L2tpConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SstpConfigParams != nil {
		if err := o.SstpConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SstpConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprapiTunnelConfigParams2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.L2tpConfigParams == nil {
		o.L2tpConfigParams = &L2tpConfigParams1{}
	}
	if err := o.L2tpConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SstpConfigParams == nil {
		o.SstpConfigParams = &SstpConfigParams1{}
	}
	if err := o.SstpConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprapiTunnelConfigParams3 structure represents MPRAPI_TUNNEL_CONFIG_PARAMS_3 RPC structure.
type MprapiTunnelConfigParams3 struct {
	IKEConfigParams  *IKEv2ConfigParams3 `idl:"name:IkeConfigParams" json:"ike_config_params"`
	PPTPConfigParams *PPTPConfigParams1  `idl:"name:PptpConfigParams" json:"pptp_config_params"`
	L2tpConfigParams *L2tpConfigParams2  `idl:"name:L2tpConfigParams" json:"l2tp_config_params"`
	SstpConfigParams *SstpConfigParams1  `idl:"name:SstpConfigParams" json:"sstp_config_params"`
}

func (o *MprapiTunnelConfigParams3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprapiTunnelConfigParams3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.L2tpConfigParams != nil {
		if err := o.L2tpConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&L2tpConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SstpConfigParams != nil {
		if err := o.SstpConfigParams.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SstpConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprapiTunnelConfigParams3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.L2tpConfigParams == nil {
		o.L2tpConfigParams = &L2tpConfigParams2{}
	}
	if err := o.L2tpConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SstpConfigParams == nil {
		o.SstpConfigParams = &SstpConfigParams1{}
	}
	if err := o.SstpConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerEx1 structure represents MPR_SERVER_EX_1 RPC structure.
type MprServerEx1 struct {
	Header       *MprapiObjectHeaderIDL     `idl:"name:Header" json:"header"`
	LANOnlyMode  bool                       `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime       uint32                     `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts   uint32                     `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse   uint32                     `idl:"name:dwPortsInUse" json:"ports_in_use"`
	_            uint32                     `idl:"name:Reserved"`
	ConfigParams *MprapiTunnelConfigParams1 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *MprServerEx1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServerEx1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&MprapiTunnelConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerEx1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
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
		o.ConfigParams = &MprapiTunnelConfigParams1{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerEx2 structure represents MPR_SERVER_EX_2 RPC structure.
type MprServerEx2 struct {
	Header       *MprapiObjectHeaderIDL     `idl:"name:Header" json:"header"`
	LANOnlyMode  bool                       `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime       uint32                     `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts   uint32                     `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse   uint32                     `idl:"name:dwPortsInUse" json:"ports_in_use"`
	_            uint32                     `idl:"name:Reserved"`
	ConfigParams *MprapiTunnelConfigParams2 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *MprServerEx2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServerEx2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&MprapiTunnelConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerEx2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
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
		o.ConfigParams = &MprapiTunnelConfigParams2{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerEx3 structure represents MPR_SERVER_EX_3 RPC structure.
type MprServerEx3 struct {
	Header       *MprapiObjectHeaderIDL     `idl:"name:Header" json:"header"`
	LANOnlyMode  bool                       `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime       uint32                     `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts   uint32                     `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse   uint32                     `idl:"name:dwPortsInUse" json:"ports_in_use"`
	_            uint32                     `idl:"name:Reserved"`
	ConfigParams *MprapiTunnelConfigParams3 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *MprServerEx3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServerEx3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&MprapiTunnelConfigParams3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerEx3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
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
		o.ConfigParams = &MprapiTunnelConfigParams3{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerExIDL structure represents MPR_SERVER_EX_IDL RPC union.
type MprServerExIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *MprServerExIDL_ServerConfig1
	// *MprServerExIDL_ServerConfig2
	// *MprServerExIDL_ServerConfig3
	Value is_MprServerExIDL `json:"value"`
}

func (o *MprServerExIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MprServerExIDL_ServerConfig1:
		if value != nil {
			return value.ServerConfig1
		}
	case *MprServerExIDL_ServerConfig2:
		if value != nil {
			return value.ServerConfig2
		}
	case *MprServerExIDL_ServerConfig3:
		if value != nil {
			return value.ServerConfig3
		}
	}
	return nil
}

type is_MprServerExIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MprServerExIDL()
}

func (o *MprServerExIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		_o, _ := o.Value.(*MprServerExIDL_ServerConfig1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprServerExIDL_ServerConfig1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*MprServerExIDL_ServerConfig2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprServerExIDL_ServerConfig2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(3):
		_o, _ := o.Value.(*MprServerExIDL_ServerConfig3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprServerExIDL_ServerConfig3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *MprServerExIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Value = &MprServerExIDL_ServerConfig1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &MprServerExIDL_ServerConfig2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(3):
		o.Value = &MprServerExIDL_ServerConfig3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// MprServerExIDL_ServerConfig1 structure represents MPR_SERVER_EX_IDL RPC union arm.
//
// It has following labels: 1
type MprServerExIDL_ServerConfig1 struct {
	ServerConfig1 *MprServerEx1 `idl:"name:ServerConfig1" json:"server_config1"`
}

func (*MprServerExIDL_ServerConfig1) is_MprServerExIDL() {}

func (o *MprServerExIDL_ServerConfig1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerConfig1 != nil {
		if err := o.ServerConfig1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprServerEx1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerExIDL_ServerConfig1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerConfig1 == nil {
		o.ServerConfig1 = &MprServerEx1{}
	}
	if err := o.ServerConfig1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerExIDL_ServerConfig2 structure represents MPR_SERVER_EX_IDL RPC union arm.
//
// It has following labels: 2
type MprServerExIDL_ServerConfig2 struct {
	ServerConfig2 *MprServerEx2 `idl:"name:ServerConfig2" json:"server_config2"`
}

func (*MprServerExIDL_ServerConfig2) is_MprServerExIDL() {}

func (o *MprServerExIDL_ServerConfig2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerConfig2 != nil {
		if err := o.ServerConfig2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprServerEx2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerExIDL_ServerConfig2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerConfig2 == nil {
		o.ServerConfig2 = &MprServerEx2{}
	}
	if err := o.ServerConfig2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerExIDL_ServerConfig3 structure represents MPR_SERVER_EX_IDL RPC union arm.
//
// It has following labels: 3
type MprServerExIDL_ServerConfig3 struct {
	ServerConfig3 *MprServerEx3 `idl:"name:ServerConfig3" json:"server_config3"`
}

func (*MprServerExIDL_ServerConfig3) is_MprServerExIDL() {}

func (o *MprServerExIDL_ServerConfig3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerConfig3 != nil {
		if err := o.ServerConfig3.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprServerEx3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerExIDL_ServerConfig3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerConfig3 == nil {
		o.ServerConfig3 = &MprServerEx3{}
	}
	if err := o.ServerConfig3.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerSetConfigEx1 structure represents MPR_SERVER_SET_CONFIG_EX_1 RPC structure.
type MprServerSetConfigEx1 struct {
	Header                *MprapiObjectHeaderIDL     `idl:"name:Header" json:"header"`
	SetConfigForProtocols uint32                     `idl:"name:setConfigForProtocols" json:"set_config_for_protocols"`
	ConfigParams          *MprapiTunnelConfigParams1 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *MprServerSetConfigEx1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServerSetConfigEx1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&MprapiTunnelConfigParams1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerSetConfigEx1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &MprapiTunnelConfigParams1{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerSetConfigEx2 structure represents MPR_SERVER_SET_CONFIG_EX_2 RPC structure.
type MprServerSetConfigEx2 struct {
	Header                *MprapiObjectHeaderIDL     `idl:"name:Header" json:"header"`
	SetConfigForProtocols uint32                     `idl:"name:setConfigForProtocols" json:"set_config_for_protocols"`
	ConfigParams          *MprapiTunnelConfigParams2 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *MprServerSetConfigEx2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServerSetConfigEx2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&MprapiTunnelConfigParams2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerSetConfigEx2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &MprapiTunnelConfigParams2{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerSetConfigEx3 structure represents MPR_SERVER_SET_CONFIG_EX_3 RPC structure.
type MprServerSetConfigEx3 struct {
	Header                *MprapiObjectHeaderIDL     `idl:"name:Header" json:"header"`
	SetConfigForProtocols uint32                     `idl:"name:setConfigForProtocols" json:"set_config_for_protocols"`
	ConfigParams          *MprapiTunnelConfigParams3 `idl:"name:ConfigParams" json:"config_params"`
}

func (o *MprServerSetConfigEx3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServerSetConfigEx3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&MprapiTunnelConfigParams3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerSetConfigEx3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MprapiObjectHeaderIDL{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SetConfigForProtocols); err != nil {
		return err
	}
	if o.ConfigParams == nil {
		o.ConfigParams = &MprapiTunnelConfigParams3{}
	}
	if err := o.ConfigParams.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerSetConfigExIDL structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union.
type MprServerSetConfigExIDL struct {
	Revision uint8
	// Types that are assignable to Value
	//
	// *MprServerSetConfigExIDL_ServerSetConfig1
	// *MprServerSetConfigExIDL_ServerSetConfig2
	// *MprServerSetConfigExIDL_ServerSetConfig3
	Value is_MprServerSetConfigExIDL `json:"value"`
}

func (o *MprServerSetConfigExIDL) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MprServerSetConfigExIDL_ServerSetConfig1:
		if value != nil {
			return value.ServerSetConfig1
		}
	case *MprServerSetConfigExIDL_ServerSetConfig2:
		if value != nil {
			return value.ServerSetConfig2
		}
	case *MprServerSetConfigExIDL_ServerSetConfig3:
		if value != nil {
			return value.ServerSetConfig3
		}
	}
	return nil
}

type is_MprServerSetConfigExIDL interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MprServerSetConfigExIDL()
}

func (o *MprServerSetConfigExIDL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		_o, _ := o.Value.(*MprServerSetConfigExIDL_ServerSetConfig1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprServerSetConfigExIDL_ServerSetConfig1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*MprServerSetConfigExIDL_ServerSetConfig2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprServerSetConfigExIDL_ServerSetConfig2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(3):
		_o, _ := o.Value.(*MprServerSetConfigExIDL_ServerSetConfig3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MprServerSetConfigExIDL_ServerSetConfig3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

func (o *MprServerSetConfigExIDL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Value = &MprServerSetConfigExIDL_ServerSetConfig1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &MprServerSetConfigExIDL_ServerSetConfig2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(3):
		o.Value = &MprServerSetConfigExIDL_ServerSetConfig3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.Revision)
	}
	return nil
}

// MprServerSetConfigExIDL_ServerSetConfig1 structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union arm.
//
// It has following labels: 1
type MprServerSetConfigExIDL_ServerSetConfig1 struct {
	ServerSetConfig1 *MprServerSetConfigEx1 `idl:"name:ServerSetConfig1" json:"server_set_config1"`
}

func (*MprServerSetConfigExIDL_ServerSetConfig1) is_MprServerSetConfigExIDL() {}

func (o *MprServerSetConfigExIDL_ServerSetConfig1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerSetConfig1 != nil {
		if err := o.ServerSetConfig1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprServerSetConfigEx1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerSetConfigExIDL_ServerSetConfig1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerSetConfig1 == nil {
		o.ServerSetConfig1 = &MprServerSetConfigEx1{}
	}
	if err := o.ServerSetConfig1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerSetConfigExIDL_ServerSetConfig2 structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union arm.
//
// It has following labels: 2
type MprServerSetConfigExIDL_ServerSetConfig2 struct {
	ServerSetConfig2 *MprServerSetConfigEx2 `idl:"name:ServerSetConfig2" json:"server_set_config2"`
}

func (*MprServerSetConfigExIDL_ServerSetConfig2) is_MprServerSetConfigExIDL() {}

func (o *MprServerSetConfigExIDL_ServerSetConfig2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerSetConfig2 != nil {
		if err := o.ServerSetConfig2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprServerSetConfigEx2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerSetConfigExIDL_ServerSetConfig2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerSetConfig2 == nil {
		o.ServerSetConfig2 = &MprServerSetConfigEx2{}
	}
	if err := o.ServerSetConfig2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MprServerSetConfigExIDL_ServerSetConfig3 structure represents MPR_SERVER_SET_CONFIG_EX_IDL RPC union arm.
//
// It has following labels: 3
type MprServerSetConfigExIDL_ServerSetConfig3 struct {
	ServerSetConfig3 *MprServerSetConfigEx3 `idl:"name:ServerSetConfig3" json:"server_set_config3"`
}

func (*MprServerSetConfigExIDL_ServerSetConfig3) is_MprServerSetConfigExIDL() {}

func (o *MprServerSetConfigExIDL_ServerSetConfig3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServerSetConfig3 != nil {
		if err := o.ServerSetConfig3.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MprServerSetConfigEx3{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MprServerSetConfigExIDL_ServerSetConfig3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ServerSetConfig3 == nil {
		o.ServerSetConfig3 = &MprServerSetConfigEx3{}
	}
	if err := o.ServerSetConfig3.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASUpdateConnection1IDL structure represents RAS_UPDATE_CONNECTION_1_IDL RPC structure.
type RASUpdateConnection1IDL struct {
	Header                *MprapiObjectHeaderIDL `idl:"name:Header" json:"header"`
	InterfaceIndex        uint32                 `idl:"name:dwIfIndex" json:"interface_index"`
	RemoteEndpointAddress []uint16               `idl:"name:wszRemoteEndpointAddress" json:"remote_endpoint_address"`
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
		if err := (&MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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
		o.Header = &MprapiObjectHeaderIDL{}
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

// DimInterfaceContainer structure represents DIM_INTERFACE_CONTAINER RPC structure.
type DimInterfaceContainer struct {
	GetInterfaceInfo  uint32 `idl:"name:fGetInterfaceInfo" json:"get_interface_info"`
	InterfaceInfoSize uint32 `idl:"name:dwInterfaceInfoSize" json:"interface_info_size"`
	InterfaceInfo     []byte `idl:"name:pInterfaceInfo;size_is:(dwInterfaceInfoSize)" json:"interface_info"`
	GetGlobalInfo     uint32 `idl:"name:fGetGlobalInfo" json:"get_global_info"`
	GlobalInfoSize    uint32 `idl:"name:dwGlobalInfoSize" json:"global_info_size"`
	GlobalInfo        []byte `idl:"name:pGlobalInfo;size_is:(dwGlobalInfoSize)" json:"global_info"`
}

func (o *DimInterfaceContainer) xxx_PreparePayload(ctx context.Context) error {
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
func (o *DimInterfaceContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DimInterfaceContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RtrTocEntry structure represents RTR_TOC_ENTRY RPC structure.
type RtrTocEntry struct {
	InfoType uint32 `idl:"name:InfoType" json:"info_type"`
	InfoSize uint32 `idl:"name:InfoSize" json:"info_size"`
	Count    uint32 `idl:"name:Count" json:"count"`
	Offset   uint32 `idl:"name:Offset" json:"offset"`
}

func (o *RtrTocEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RtrTocEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RtrTocEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RtrInfoBlockHeader structure represents RTR_INFO_BLOCK_HEADER RPC structure.
type RtrInfoBlockHeader struct {
	Version         uint32         `idl:"name:Version" json:"version"`
	Size            uint32         `idl:"name:Size" json:"size"`
	TocEntriesCount uint32         `idl:"name:TocEntriesCount" json:"toc_entries_count"`
	TocEntry        []*RtrTocEntry `idl:"name:TocEntry" json:"toc_entry"`
}

func (o *RtrInfoBlockHeader) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RtrInfoBlockHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.TocEntriesCount); err != nil {
		return err
	}
	for i1 := range o.TocEntry {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.TocEntry[i1] != nil {
			if err := o.TocEntry[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RtrTocEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.TocEntry); uint64(i1) < 1; i1++ {
		if err := (&RtrTocEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RtrInfoBlockHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.TocEntriesCount); err != nil {
		return err
	}
	o.TocEntry = make([]*RtrTocEntry, 1)
	for i1 := range o.TocEntry {
		i1 := i1
		if o.TocEntry[i1] == nil {
			o.TocEntry[i1] = &RtrTocEntry{}
		}
		if err := o.TocEntry[i1].UnmarshalNDR(ctx, w); err != nil {
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
	Version         uint32        `idl:"name:dwVersion" json:"version"`
	FiltersLength   uint32        `idl:"name:dwNumFilters" json:"filters_length"`
	FaDefaultAction ForwardAction `idl:"name:faDefaultAction" json:"fa_default_action"`
	FiFilter        []*FilterInfo `idl:"name:fiFilter" json:"fi_filter"`
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
	if err := w.WriteEnum(uint16(o.FaDefaultAction)); err != nil {
		return err
	}
	for i1 := range o.FiFilter {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.FiFilter[i1] != nil {
			if err := o.FiFilter[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FilterInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.FiFilter); uint64(i1) < 1; i1++ {
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
	if err := w.ReadEnum((*uint16)(&o.FaDefaultAction)); err != nil {
		return err
	}
	o.FiFilter = make([]*FilterInfo, 1)
	for i1 := range o.FiFilter {
		i1 := i1
		if o.FiFilter[i1] == nil {
			o.FiFilter[i1] = &FilterInfo{}
		}
		if err := o.FiFilter[i1].UnmarshalNDR(ctx, w); err != nil {
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
	Version         uint32          `idl:"name:dwVersion" json:"version"`
	FiltersLength   uint32          `idl:"name:dwNumFilters" json:"filters_length"`
	FaDefaultAction ForwardAction   `idl:"name:faDefaultAction" json:"fa_default_action"`
	FiFilter        []*FilterInfoV6 `idl:"name:fiFilter" json:"fi_filter"`
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
	if err := w.WriteEnum(uint16(o.FaDefaultAction)); err != nil {
		return err
	}
	for i1 := range o.FiFilter {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.FiFilter[i1] != nil {
			if err := o.FiFilter[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FilterInfoV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.FiFilter); uint64(i1) < 1; i1++ {
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
	if err := w.ReadEnum((*uint16)(&o.FaDefaultAction)); err != nil {
		return err
	}
	o.FiFilter = make([]*FilterInfoV6, 1)
	for i1 := range o.FiFilter {
		i1 := i1
		if o.FiFilter[i1] == nil {
			o.FiFilter[i1] = &FilterInfoV6{}
		}
		if err := o.FiFilter[i1].UnmarshalNDR(ctx, w); err != nil {
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
	DestinationPrefix       *In6Addr `idl:"name:DestinationPrefix" json:"destination_prefix"`
	DestinationPrefixLength uint32   `idl:"name:DestPrefixLength" json:"destination_prefix_length"`
	NextHopAddress          *In6Addr `idl:"name:NextHopAddress" json:"next_hop_address"`
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
		if err := (&In6Addr{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&In6Addr{}).MarshalNDR(ctx, w); err != nil {
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
		o.DestinationPrefix = &In6Addr{}
	}
	if err := o.DestinationPrefix.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationPrefixLength); err != nil {
		return err
	}
	if o.NextHopAddress == nil {
		o.NextHopAddress = &In6Addr{}
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
	ProtocolsLength   uint32            `idl:"name:dwNumProtocols" json:"protocols_length"`
	PpmProtocolMetric []*ProtocolMetric `idl:"name:ppmProtocolMetric" json:"ppm_protocol_metric"`
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
	for i1 := range o.PpmProtocolMetric {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.PpmProtocolMetric[i1] != nil {
			if err := o.PpmProtocolMetric[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProtocolMetric{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.PpmProtocolMetric); uint64(i1) < 1; i1++ {
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
	o.PpmProtocolMetric = make([]*ProtocolMetric, 1)
	for i1 := range o.PpmProtocolMetric {
		i1 := i1
		if o.PpmProtocolMetric[i1] == nil {
			o.PpmProtocolMetric[i1] = &ProtocolMetric{}
		}
		if err := o.PpmProtocolMetric[i1].UnmarshalNDR(ctx, w); err != nil {
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
	ProtocolsLength   uint32              `idl:"name:dwNumProtocols" json:"protocols_length"`
	PpmProtocolMetric []*ProtocolMetricEx `idl:"name:ppmProtocolMetric" json:"ppm_protocol_metric"`
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
	for i1 := range o.PpmProtocolMetric {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.PpmProtocolMetric[i1] != nil {
			if err := o.PpmProtocolMetric[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProtocolMetricEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.PpmProtocolMetric); uint64(i1) < 1; i1++ {
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
	o.PpmProtocolMetric = make([]*ProtocolMetricEx, 1)
	for i1 := range o.PpmProtocolMetric {
		i1 := i1
		if o.PpmProtocolMetric[i1] == nil {
			o.PpmProtocolMetric[i1] = &ProtocolMetricEx{}
		}
		if err := o.PpmProtocolMetric[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// RtrDiscInfo structure represents RTR_DISC_INFO RPC structure.
type RtrDiscInfo struct {
	MaxAdvtInterval uint16 `idl:"name:wMaxAdvtInterval" json:"max_advt_interval"`
	MinAdvtInterval uint16 `idl:"name:wMinAdvtInterval" json:"min_advt_interval"`
	AdvtLifetime    uint16 `idl:"name:wAdvtLifetime" json:"advt_lifetime"`
	Advertise       bool   `idl:"name:bAdvertise" json:"advertise"`
	PrefLevel       int32  `idl:"name:lPrefLevel" json:"pref_level"`
}

func (o *RtrDiscInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RtrDiscInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxAdvtInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.MinAdvtInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.AdvtLifetime); err != nil {
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
func (o *RtrDiscInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxAdvtInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinAdvtInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdvtLifetime); err != nil {
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

// MCastHbeatInfo structure represents MCAST_HBEAT_INFO RPC structure.
type MCastHbeatInfo struct {
	Group        []uint16 `idl:"name:pwszGroup" json:"group"`
	Active       bool     `idl:"name:bActive" json:"active"`
	DeadInterval uint32   `idl:"name:ulDeadInterval" json:"dead_interval"`
	ByProtocol   uint8    `idl:"name:byProtocol" json:"by_protocol"`
	Port         uint16   `idl:"name:wPort" json:"port"`
}

func (o *MCastHbeatInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MCastHbeatInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MCastHbeatInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBMCastLimitRow structure represents MIB_MCAST_LIMIT_ROW RPC structure.
type MIBMCastLimitRow struct {
	TTL       uint32 `idl:"name:dwTtl" json:"ttl"`
	RateLimit uint32 `idl:"name:dwRateLimit" json:"rate_limit"`
}

func (o *MIBMCastLimitRow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBMCastLimitRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBMCastLimitRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpinipConfigInfo structure represents IPINIP_CONFIG_INFO RPC structure.
type IpinipConfigInfo struct {
	RemoteAddress uint32 `idl:"name:dwRemoteAddress" json:"remote_address"`
	LocalAddress  uint32 `idl:"name:dwLocalAddress" json:"local_address"`
	ByTTL         uint8  `idl:"name:byTtl" json:"by_ttl"`
}

func (o *IpinipConfigInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpinipConfigInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpinipConfigInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DimMIBEntryContainer structure represents DIM_MIB_ENTRY_CONTAINER RPC structure.
type DimMIBEntryContainer struct {
	MIBInEntrySize  uint32 `idl:"name:dwMibInEntrySize" json:"mib_in_entry_size"`
	MIBInEntry      []byte `idl:"name:pMibInEntry;size_is:(dwMibInEntrySize)" json:"mib_in_entry"`
	MIBOutEntrySize uint32 `idl:"name:dwMibOutEntrySize" json:"mib_out_entry_size"`
	MIBOutEntry     []byte `idl:"name:pMibOutEntry;size_is:(dwMibOutEntrySize)" json:"mib_out_entry"`
}

func (o *DimMIBEntryContainer) xxx_PreparePayload(ctx context.Context) error {
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
func (o *DimMIBEntryContainer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DimMIBEntryContainer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBIpforwardrow structure represents MIB_IPFORWARDROW RPC structure.
type MIBIpforwardrow struct {
	ForwardDestination    uint32            `idl:"name:dwForwardDest" json:"forward_destination"`
	ForwardMask           uint32            `idl:"name:dwForwardMask" json:"forward_mask"`
	ForwardPolicy         uint32            `idl:"name:dwForwardPolicy" json:"forward_policy"`
	ForwardNextHop        uint32            `idl:"name:dwForwardNextHop" json:"forward_next_hop"`
	ForwardInterfaceIndex uint32            `idl:"name:dwForwardIfIndex" json:"forward_interface_index"`
	ForwardType           MIBIpforwardType  `idl:"name:ForwardType" json:"forward_type"`
	ForwardProto          MIBIpforwardProto `idl:"name:ForwardProto" json:"forward_proto"`
	ForwardAge            uint32            `idl:"name:dwForwardAge" json:"forward_age"`
	ForwardNextHopAs      uint32            `idl:"name:dwForwardNextHopAS" json:"forward_next_hop_as"`
	ForwardMetric1        uint32            `idl:"name:dwForwardMetric1" json:"forward_metric1"`
	ForwardMetric2        uint32            `idl:"name:dwForwardMetric2" json:"forward_metric2"`
	ForwardMetric3        uint32            `idl:"name:dwForwardMetric3" json:"forward_metric3"`
	ForwardMetric4        uint32            `idl:"name:dwForwardMetric4" json:"forward_metric4"`
	ForwardMetric5        uint32            `idl:"name:dwForwardMetric5" json:"forward_metric5"`
}

func (o *MIBIpforwardrow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpforwardrow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIpforwardrow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBIpdestrow structure represents MIB_IPDESTROW RPC structure.
type MIBIpdestrow struct {
	ForwardRow        *MIBIpforwardrow `idl:"name:ForwardRow" json:"forward_row"`
	ForwardPreference uint32           `idl:"name:dwForwardPreference" json:"forward_preference"`
	ForwardViewSet    uint32           `idl:"name:dwForwardViewSet" json:"forward_view_set"`
}

func (o *MIBIpdestrow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpdestrow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&MIBIpforwardrow{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MIBIpdestrow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ForwardRow == nil {
		o.ForwardRow = &MIBIpforwardrow{}
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

// MIBIpdesttable structure represents MIB_IPDESTTABLE RPC structure.
type MIBIpdesttable struct {
	EntriesLength uint32          `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpdestrow `idl:"name:table" json:"table"`
}

func (o *MIBIpdesttable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpdesttable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpdestrow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpdestrow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIpdesttable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpdestrow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpdestrow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBRoutestate structure represents MIB_ROUTESTATE RPC structure.
type MIBRoutestate struct {
	RoutesSetToStack bool `idl:"name:bRoutesSetToStack" json:"routes_set_to_stack"`
}

func (o *MIBRoutestate) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBRoutestate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBRoutestate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Mibicmpstats structure represents MIBICMPSTATS RPC structure.
type Mibicmpstats struct {
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

func (o *Mibicmpstats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Mibicmpstats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Mibicmpstats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Mibicmpinfo structure represents MIBICMPINFO RPC structure.
type Mibicmpinfo struct {
	ICMPInStats  *Mibicmpstats `idl:"name:icmpInStats" json:"icmp_in_stats"`
	ICMPOutStats *Mibicmpstats `idl:"name:icmpOutStats" json:"icmp_out_stats"`
}

func (o *Mibicmpinfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Mibicmpinfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&Mibicmpstats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ICMPOutStats != nil {
		if err := o.ICMPOutStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Mibicmpstats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Mibicmpinfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ICMPInStats == nil {
		o.ICMPInStats = &Mibicmpstats{}
	}
	if err := o.ICMPInStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ICMPOutStats == nil {
		o.ICMPOutStats = &Mibicmpstats{}
	}
	if err := o.ICMPOutStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MIBICMP structure represents MIB_ICMP RPC structure.
type MIBICMP struct {
	Stats *Mibicmpinfo `idl:"name:stats" json:"stats"`
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
		if err := (&Mibicmpinfo{}).MarshalNDR(ctx, w); err != nil {
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
		o.Stats = &Mibicmpinfo{}
	}
	if err := o.Stats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MIBIfnumber structure represents MIB_IFNUMBER RPC structure.
type MIBIfnumber struct {
	Value uint32 `idl:"name:dwValue" json:"value"`
}

func (o *MIBIfnumber) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIfnumber) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIfnumber) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Value); err != nil {
		return err
	}
	return nil
}

// MIBIfrow structure represents MIB_IFROW RPC structure.
type MIBIfrow struct {
	Name              []uint16 `idl:"name:wszName" json:"name"`
	Index             uint32   `idl:"name:dwIndex" json:"index"`
	Type              uint32   `idl:"name:dwType" json:"type"`
	Mtu               uint32   `idl:"name:dwMtu" json:"mtu"`
	Speed             uint32   `idl:"name:dwSpeed" json:"speed"`
	PhysAddrLength    uint32   `idl:"name:dwPhysAddrLen" json:"phys_addr_length"`
	PhysAddr          []byte   `idl:"name:bPhysAddr" json:"phys_addr"`
	AdminStatus       uint32   `idl:"name:dwAdminStatus" json:"admin_status"`
	OperatorStatus    uint32   `idl:"name:dwOperStatus" json:"operator_status"`
	LastChange        uint32   `idl:"name:dwLastChange" json:"last_change"`
	InOctets          uint32   `idl:"name:dwInOctets" json:"in_octets"`
	InUcastPkts       uint32   `idl:"name:dwInUcastPkts" json:"in_ucast_pkts"`
	InNUcastPkts      uint32   `idl:"name:dwInNUcastPkts" json:"in_n_ucast_pkts"`
	InDiscards        uint32   `idl:"name:dwInDiscards" json:"in_discards"`
	InErrors          uint32   `idl:"name:dwInErrors" json:"in_errors"`
	InUnknownProtos   uint32   `idl:"name:dwInUnknownProtos" json:"in_unknown_protos"`
	OutOctets         uint32   `idl:"name:dwOutOctets" json:"out_octets"`
	OutUcastPkts      uint32   `idl:"name:dwOutUcastPkts" json:"out_ucast_pkts"`
	OutNUcastPkts     uint32   `idl:"name:dwOutNUcastPkts" json:"out_n_ucast_pkts"`
	OutDiscards       uint32   `idl:"name:dwOutDiscards" json:"out_discards"`
	OutErrors         uint32   `idl:"name:dwOutErrors" json:"out_errors"`
	OutQLength        uint32   `idl:"name:dwOutQLen" json:"out_q_length"`
	DescriptionLength uint32   `idl:"name:dwDescrLen" json:"description_length"`
	Description       []byte   `idl:"name:bDescr" json:"description"`
}

func (o *MIBIfrow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIfrow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.Mtu); err != nil {
		return err
	}
	if err := w.WriteData(o.Speed); err != nil {
		return err
	}
	if err := w.WriteData(o.PhysAddrLength); err != nil {
		return err
	}
	for i1 := range o.PhysAddr {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.PhysAddr[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PhysAddr); uint64(i1) < 8; i1++ {
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
	if err := w.WriteData(o.InUcastPkts); err != nil {
		return err
	}
	if err := w.WriteData(o.InNUcastPkts); err != nil {
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
	if err := w.WriteData(o.OutUcastPkts); err != nil {
		return err
	}
	if err := w.WriteData(o.OutNUcastPkts); err != nil {
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
func (o *MIBIfrow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.Mtu); err != nil {
		return err
	}
	if err := w.ReadData(&o.Speed); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhysAddrLength); err != nil {
		return err
	}
	o.PhysAddr = make([]byte, 8)
	for i1 := range o.PhysAddr {
		i1 := i1
		if err := w.ReadData(&o.PhysAddr[i1]); err != nil {
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
	if err := w.ReadData(&o.InUcastPkts); err != nil {
		return err
	}
	if err := w.ReadData(&o.InNUcastPkts); err != nil {
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
	if err := w.ReadData(&o.OutUcastPkts); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutNUcastPkts); err != nil {
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

// MIBIfstatus structure represents MIB_IFSTATUS RPC structure.
type MIBIfstatus struct {
	InterfaceIndex    uint32 `idl:"name:dwIfIndex" json:"interface_index"`
	AdminStatus       uint32 `idl:"name:dwAdminStatus" json:"admin_status"`
	OperationalStatus uint32 `idl:"name:dwOperationalStatus" json:"operational_status"`
	MHbeatActive      bool   `idl:"name:bMHbeatActive" json:"m_hbeat_active"`
	MHbeatAlive       bool   `idl:"name:bMHbeatAlive" json:"m_hbeat_alive"`
}

func (o *MIBIfstatus) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIfstatus) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if !o.MHbeatActive {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.MHbeatAlive {
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
func (o *MIBIfstatus) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	var _bMHbeatActive int32
	if err := w.ReadData(&_bMHbeatActive); err != nil {
		return err
	}
	o.MHbeatActive = _bMHbeatActive != 0
	var _bMHbeatAlive int32
	if err := w.ReadData(&_bMHbeatAlive); err != nil {
		return err
	}
	o.MHbeatAlive = _bMHbeatAlive != 0
	return nil
}

// MIBIftable structure represents MIB_IFTABLE RPC structure.
type MIBIftable struct {
	EntriesLength uint32      `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIfrow `idl:"name:table" json:"table"`
}

func (o *MIBIftable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIftable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIfrow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIfrow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIftable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIfrow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIfrow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIpaddrrow structure represents MIB_IPADDRROW RPC structure.
type MIBIpaddrrow struct {
	Addr      uint32 `idl:"name:dwAddr" json:"addr"`
	Index     uint32 `idl:"name:dwIndex" json:"index"`
	Mask      uint32 `idl:"name:dwMask" json:"mask"`
	BCastAddr uint32 `idl:"name:dwBCastAddr" json:"b_cast_addr"`
	ReasmSize uint32 `idl:"name:dwReasmSize" json:"reasm_size"`
	_         uint16 `idl:"name:unused1"`
	Type      uint16 `idl:"name:wType" json:"type"`
}

func (o *MIBIpaddrrow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpaddrrow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIpaddrrow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBIpaddrtable structure represents MIB_IPADDRTABLE RPC structure.
type MIBIpaddrtable struct {
	EntriesLength uint32          `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpaddrrow `idl:"name:table" json:"table"`
}

func (o *MIBIpaddrtable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpaddrtable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpaddrrow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpaddrrow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIpaddrtable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpaddrrow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpaddrrow{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIpforwardnumber structure represents MIB_IPFORWARDNUMBER RPC structure.
type MIBIpforwardnumber struct {
	Value uint32 `idl:"name:dwValue" json:"value"`
}

func (o *MIBIpforwardnumber) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpforwardnumber) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIpforwardnumber) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Value); err != nil {
		return err
	}
	return nil
}

// MIBIpforwardtable structure represents MIB_IPFORWARDTABLE RPC structure.
type MIBIpforwardtable struct {
	EntriesLength uint32             `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpforwardrow `idl:"name:table" json:"table"`
	_             []byte             `idl:"name:reserved"`
}

func (o *MIBIpforwardtable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpforwardtable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpforwardrow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpforwardrow{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MIBIpforwardtable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpforwardrow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpforwardrow{}
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

// MIBIpmcastBoundary structure represents MIB_IPMCAST_BOUNDARY RPC structure.
type MIBIpmcastBoundary struct {
	InterfaceIndex uint32 `idl:"name:dwIfIndex" json:"interface_index"`
	GroupAddress   uint32 `idl:"name:dwGroupAddress" json:"group_address"`
	GroupMask      uint32 `idl:"name:dwGroupMask" json:"group_mask"`
	Status         uint32 `idl:"name:dwStatus" json:"status"`
}

func (o *MIBIpmcastBoundary) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastBoundary) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIpmcastBoundary) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBIpmcastBoundaryTable structure represents MIB_IPMCAST_BOUNDARY_TABLE RPC structure.
type MIBIpmcastBoundaryTable struct {
	EntriesLength uint32                `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpmcastBoundary `idl:"name:table" json:"table"`
}

func (o *MIBIpmcastBoundaryTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastBoundaryTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpmcastBoundary{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpmcastBoundary{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIpmcastBoundaryTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpmcastBoundary, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpmcastBoundary{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIpmcastGlobal structure represents MIB_IPMCAST_GLOBAL RPC structure.
type MIBIpmcastGlobal struct {
	Enable uint32 `idl:"name:dwEnable" json:"enable"`
}

func (o *MIBIpmcastGlobal) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastGlobal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIpmcastGlobal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enable); err != nil {
		return err
	}
	return nil
}

// MIBIpmcastInterfaceEntry structure represents MIB_IPMCAST_IF_ENTRY RPC structure.
type MIBIpmcastInterfaceEntry struct {
	InterfaceIndex uint32 `idl:"name:dwIfIndex" json:"interface_index"`
	TTL            uint32 `idl:"name:dwTtl" json:"ttl"`
	Protocol       uint32 `idl:"name:dwProtocol" json:"protocol"`
	RateLimit      uint32 `idl:"name:dwRateLimit" json:"rate_limit"`
	InMCastOctets  uint32 `idl:"name:ulInMcastOctets" json:"in_mcast_octets"`
	OutMCastOctets uint32 `idl:"name:ulOutMcastOctets" json:"out_mcast_octets"`
}

func (o *MIBIpmcastInterfaceEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastInterfaceEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.InMCastOctets); err != nil {
		return err
	}
	if err := w.WriteData(o.OutMCastOctets); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastInterfaceEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.InMCastOctets); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutMCastOctets); err != nil {
		return err
	}
	return nil
}

// MIBIpmcastInterfaceTable structure represents MIB_IPMCAST_IF_TABLE RPC structure.
type MIBIpmcastInterfaceTable struct {
	EntriesLength uint32                      `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpmcastInterfaceEntry `idl:"name:table" json:"table"`
}

func (o *MIBIpmcastInterfaceTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastInterfaceTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpmcastInterfaceEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpmcastInterfaceEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIpmcastInterfaceTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpmcastInterfaceEntry, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpmcastInterfaceEntry{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIpmcastOif structure represents MIB_IPMCAST_OIF RPC structure.
type MIBIpmcastOif struct {
	OutInterfaceIndex uint32 `idl:"name:dwOutIfIndex" json:"out_interface_index"`
	NextHopAddr       uint32 `idl:"name:dwNextHopAddr" json:"next_hop_addr"`
	_                 []byte `idl:"name:pvReserved"`
	_                 uint32 `idl:"name:dwReserved"`
}

func (o *MIBIpmcastOif) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastOif) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIpmcastOif) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBIpmcastMfe structure represents MIB_IPMCAST_MFE RPC structure.
type MIBIpmcastMfe struct {
	Group               uint32           `idl:"name:dwGroup" json:"group"`
	Source              uint32           `idl:"name:dwSource" json:"source"`
	SourceMask          uint32           `idl:"name:dwSrcMask" json:"source_mask"`
	UpStrmNgbr          uint32           `idl:"name:dwUpStrmNgbr" json:"up_strm_ngbr"`
	InInterfaceIndex    uint32           `idl:"name:dwInIfIndex" json:"in_interface_index"`
	InInterfaceProtocol uint32           `idl:"name:dwInIfProtocol" json:"in_interface_protocol"`
	RouteProtocol       uint32           `idl:"name:dwRouteProtocol" json:"route_protocol"`
	RouteNetwork        uint32           `idl:"name:dwRouteNetwork" json:"route_network"`
	RouteMask           uint32           `idl:"name:dwRouteMask" json:"route_mask"`
	UpTime              uint32           `idl:"name:ulUpTime" json:"up_time"`
	ExpiryTime          uint32           `idl:"name:ulExpiryTime" json:"expiry_time"`
	Timeout             uint32           `idl:"name:ulTimeOut" json:"timeout"`
	NumOutInterface     uint32           `idl:"name:ulNumOutIf" json:"num_out_interface"`
	Flags               uint32           `idl:"name:fFlags" json:"flags"`
	_                   uint32           `idl:"name:dwReserved"`
	RgmioOutInfo        []*MIBIpmcastOif `idl:"name:rgmioOutInfo" json:"rgmio_out_info"`
}

func (o *MIBIpmcastMfe) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastMfe) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.UpStrmNgbr); err != nil {
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
			if err := (&MIBIpmcastOif{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.RgmioOutInfo); uint64(i1) < 1; i1++ {
		if err := (&MIBIpmcastOif{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIpmcastMfe) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.UpStrmNgbr); err != nil {
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
	o.RgmioOutInfo = make([]*MIBIpmcastOif, 1)
	for i1 := range o.RgmioOutInfo {
		i1 := i1
		if o.RgmioOutInfo[i1] == nil {
			o.RgmioOutInfo[i1] = &MIBIpmcastOif{}
		}
		if err := o.RgmioOutInfo[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIpmcastOifStats structure represents MIB_IPMCAST_OIF_STATS RPC structure.
type MIBIpmcastOifStats struct {
	OutInterfaceIndex uint32 `idl:"name:dwOutIfIndex" json:"out_interface_index"`
	NextHopAddr       uint32 `idl:"name:dwNextHopAddr" json:"next_hop_addr"`
	DialContext       []byte `idl:"name:pvDialContext" json:"dial_context"`
	TTLTooLow         uint32 `idl:"name:ulTtlTooLow" json:"ttl_too_low"`
	FragNeeded        uint32 `idl:"name:ulFragNeeded" json:"frag_needed"`
	OutPackets        uint32 `idl:"name:ulOutPackets" json:"out_packets"`
	OutDiscards       uint32 `idl:"name:ulOutDiscards" json:"out_discards"`
}

func (o *MIBIpmcastOifStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastOifStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBIpmcastOifStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBIpmcastMfeStats structure represents MIB_IPMCAST_MFE_STATS RPC structure.
type MIBIpmcastMfeStats struct {
	Group                  uint32                `idl:"name:dwGroup" json:"group"`
	Source                 uint32                `idl:"name:dwSource" json:"source"`
	SourceMask             uint32                `idl:"name:dwSrcMask" json:"source_mask"`
	UpStrmNgbr             uint32                `idl:"name:dwUpStrmNgbr" json:"up_strm_ngbr"`
	InInterfaceIndex       uint32                `idl:"name:dwInIfIndex" json:"in_interface_index"`
	InInterfaceProtocol    uint32                `idl:"name:dwInIfProtocol" json:"in_interface_protocol"`
	RouteProtocol          uint32                `idl:"name:dwRouteProtocol" json:"route_protocol"`
	RouteNetwork           uint32                `idl:"name:dwRouteNetwork" json:"route_network"`
	RouteMask              uint32                `idl:"name:dwRouteMask" json:"route_mask"`
	UpTime                 uint32                `idl:"name:ulUpTime" json:"up_time"`
	ExpiryTime             uint32                `idl:"name:ulExpiryTime" json:"expiry_time"`
	NumOutInterface        uint32                `idl:"name:ulNumOutIf" json:"num_out_interface"`
	InPkts                 uint32                `idl:"name:ulInPkts" json:"in_pkts"`
	InOctets               uint32                `idl:"name:ulInOctets" json:"in_octets"`
	PktsDifferentInterface uint32                `idl:"name:ulPktsDifferentIf" json:"pkts_different_interface"`
	QueueOverflow          uint32                `idl:"name:ulQueueOverflow" json:"queue_overflow"`
	RgmiosOutStats         []*MIBIpmcastOifStats `idl:"name:rgmiosOutStats" json:"rgmios_out_stats"`
}

func (o *MIBIpmcastMfeStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastMfeStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.UpStrmNgbr); err != nil {
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
	if err := w.WriteData(o.InPkts); err != nil {
		return err
	}
	if err := w.WriteData(o.InOctets); err != nil {
		return err
	}
	if err := w.WriteData(o.PktsDifferentInterface); err != nil {
		return err
	}
	if err := w.WriteData(o.QueueOverflow); err != nil {
		return err
	}
	for i1 := range o.RgmiosOutStats {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.RgmiosOutStats[i1] != nil {
			if err := o.RgmiosOutStats[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MIBIpmcastOifStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.RgmiosOutStats); uint64(i1) < 1; i1++ {
		if err := (&MIBIpmcastOifStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBIpmcastMfeStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.UpStrmNgbr); err != nil {
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
	if err := w.ReadData(&o.InPkts); err != nil {
		return err
	}
	if err := w.ReadData(&o.InOctets); err != nil {
		return err
	}
	if err := w.ReadData(&o.PktsDifferentInterface); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueueOverflow); err != nil {
		return err
	}
	o.RgmiosOutStats = make([]*MIBIpmcastOifStats, 1)
	for i1 := range o.RgmiosOutStats {
		i1 := i1
		if o.RgmiosOutStats[i1] == nil {
			o.RgmiosOutStats[i1] = &MIBIpmcastOifStats{}
		}
		if err := o.RgmiosOutStats[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBIpmcastScope structure represents MIB_IPMCAST_SCOPE RPC structure.
type MIBIpmcastScope struct {
	GroupAddress uint32   `idl:"name:dwGroupAddress" json:"group_address"`
	GroupMask    uint32   `idl:"name:dwGroupMask" json:"group_mask"`
	SnNameBuffer []uint16 `idl:"name:snNameBuffer" json:"sn_name_buffer"`
	Status       uint32   `idl:"name:dwStatus" json:"status"`
	_            []byte   `idl:"name:reserved"`
}

func (o *MIBIpmcastScope) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpmcastScope) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	for i1 := range o.SnNameBuffer {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.SnNameBuffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SnNameBuffer); uint64(i1) < 256; i1++ {
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
func (o *MIBIpmcastScope) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupMask); err != nil {
		return err
	}
	o.SnNameBuffer = make([]uint16, 256)
	for i1 := range o.SnNameBuffer {
		i1 := i1
		if err := w.ReadData(&o.SnNameBuffer[i1]); err != nil {
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

// MIBIpnetrow structure represents MIB_IPNETROW RPC structure.
type MIBIpnetrow struct {
	Index          uint32 `idl:"name:dwIndex" json:"index"`
	PhysAddrLength uint32 `idl:"name:dwPhysAddrLen" json:"phys_addr_length"`
	PhysAddr       []byte `idl:"name:bPhysAddr" json:"phys_addr"`
	Addr           uint32 `idl:"name:dwAddr" json:"addr"`
	Type           uint32 `idl:"name:dwType" json:"type"`
}

func (o *MIBIpnetrow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpnetrow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.PhysAddrLength); err != nil {
		return err
	}
	for i1 := range o.PhysAddr {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.PhysAddr[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PhysAddr); uint64(i1) < 8; i1++ {
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
func (o *MIBIpnetrow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhysAddrLength); err != nil {
		return err
	}
	o.PhysAddr = make([]byte, 8)
	for i1 := range o.PhysAddr {
		i1 := i1
		if err := w.ReadData(&o.PhysAddr[i1]); err != nil {
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

// MIBIpnettable structure represents MIB_IPNETTABLE RPC structure.
type MIBIpnettable struct {
	EntriesLength uint32         `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpnetrow `idl:"name:table" json:"table"`
	_             []byte         `idl:"name:reserved"`
}

func (o *MIBIpnettable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBIpnettable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpnetrow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpnetrow{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MIBIpnettable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpnetrow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpnetrow{}
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
	ReasmOks        uint32               `idl:"name:dwReasmOks" json:"reasm_oks"`
	ReasmFails      uint32               `idl:"name:dwReasmFails" json:"reasm_fails"`
	FragOks         uint32               `idl:"name:dwFragOks" json:"frag_oks"`
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
	if err := w.WriteData(o.ReasmOks); err != nil {
		return err
	}
	if err := w.WriteData(o.ReasmFails); err != nil {
		return err
	}
	if err := w.WriteData(o.FragOks); err != nil {
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
	if err := w.ReadData(&o.ReasmOks); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReasmFails); err != nil {
		return err
	}
	if err := w.ReadData(&o.FragOks); err != nil {
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

// MIBMfeStatsTable structure represents MIB_MFE_STATS_TABLE RPC structure.
type MIBMfeStatsTable struct {
	EntriesLength uint32                `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpmcastMfeStats `idl:"name:table" json:"table"`
}

func (o *MIBMfeStatsTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBMfeStatsTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpmcastMfeStats{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpmcastMfeStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBMfeStatsTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpmcastMfeStats, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpmcastMfeStats{}
		}
		if err := o.Table[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// MIBMfeTable structure represents MIB_MFE_TABLE RPC structure.
type MIBMfeTable struct {
	EntriesLength uint32           `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBIpmcastMfe `idl:"name:table" json:"table"`
}

func (o *MIBMfeTable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBMfeTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBIpmcastMfe{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBIpmcastMfe{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBMfeTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBIpmcastMfe, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBIpmcastMfe{}
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
	VarID        uint32   `idl:"name:dwVarId" json:"var_id"`
	RgdwVarIndex []uint32 `idl:"name:rgdwVarIndex" json:"rgdw_var_index"`
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
	for i1 := range o.RgdwVarIndex {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.RgdwVarIndex[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RgdwVarIndex); uint64(i1) < 1; i1++ {
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
	o.RgdwVarIndex = make([]uint32, 1)
	for i1 := range o.RgdwVarIndex {
		i1 := i1
		if err := w.ReadData(&o.RgdwVarIndex[i1]); err != nil {
			return err
		}
	}
	return nil
}

// MIBProxyarp structure represents MIB_PROXYARP RPC structure.
type MIBProxyarp struct {
	Address        uint32 `idl:"name:dwAddress" json:"address"`
	Mask           uint32 `idl:"name:dwMask" json:"mask"`
	InterfaceIndex uint32 `idl:"name:dwIfIndex" json:"interface_index"`
}

func (o *MIBProxyarp) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBProxyarp) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBProxyarp) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBTcprow structure represents MIB_TCPROW RPC structure.
type MIBTcprow struct {
	State      MIBTCPState `idl:"name:State" json:"state"`
	LocalAddr  uint32      `idl:"name:dwLocalAddr" json:"local_addr"`
	LocalPort  uint32      `idl:"name:dwLocalPort" json:"local_port"`
	RemoteAddr uint32      `idl:"name:dwRemoteAddr" json:"remote_addr"`
	RemotePort uint32      `idl:"name:dwRemotePort" json:"remote_port"`
}

func (o *MIBTcprow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBTcprow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBTcprow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBTcpstats structure represents MIB_TCPSTATS RPC structure.
type MIBTcpstats struct {
	RtoAlgorithm TCPRtoAlgorithm `idl:"name:RtoAlgorithm" json:"rto_algorithm"`
	RtoMin       uint32          `idl:"name:dwRtoMin" json:"rto_min"`
	RtoMax       uint32          `idl:"name:dwRtoMax" json:"rto_max"`
	MaxConn      uint32          `idl:"name:dwMaxConn" json:"max_conn"`
	ActiveOpens  uint32          `idl:"name:dwActiveOpens" json:"active_opens"`
	PassiveOpens uint32          `idl:"name:dwPassiveOpens" json:"passive_opens"`
	AttemptFails uint32          `idl:"name:dwAttemptFails" json:"attempt_fails"`
	EstabResets  uint32          `idl:"name:dwEstabResets" json:"estab_resets"`
	CurrEstab    uint32          `idl:"name:dwCurrEstab" json:"curr_estab"`
	InSegs       uint32          `idl:"name:dwInSegs" json:"in_segs"`
	OutSegs      uint32          `idl:"name:dwOutSegs" json:"out_segs"`
	RetransSegs  uint32          `idl:"name:dwRetransSegs" json:"retrans_segs"`
	InErrs       uint32          `idl:"name:dwInErrs" json:"in_errs"`
	OutRsts      uint32          `idl:"name:dwOutRsts" json:"out_rsts"`
	ConnsLength  uint32          `idl:"name:dwNumConns" json:"conns_length"`
}

func (o *MIBTcpstats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBTcpstats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RtoAlgorithm)); err != nil {
		return err
	}
	if err := w.WriteData(o.RtoMin); err != nil {
		return err
	}
	if err := w.WriteData(o.RtoMax); err != nil {
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
	if err := w.WriteData(o.CurrEstab); err != nil {
		return err
	}
	if err := w.WriteData(o.InSegs); err != nil {
		return err
	}
	if err := w.WriteData(o.OutSegs); err != nil {
		return err
	}
	if err := w.WriteData(o.RetransSegs); err != nil {
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
func (o *MIBTcpstats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RtoAlgorithm)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RtoMin); err != nil {
		return err
	}
	if err := w.ReadData(&o.RtoMax); err != nil {
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
	if err := w.ReadData(&o.CurrEstab); err != nil {
		return err
	}
	if err := w.ReadData(&o.InSegs); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutSegs); err != nil {
		return err
	}
	if err := w.ReadData(&o.RetransSegs); err != nil {
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

// MIBTcptable structure represents MIB_TCPTABLE RPC structure.
type MIBTcptable struct {
	EntriesLength uint32       `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBTcprow `idl:"name:table" json:"table"`
	_             []byte       `idl:"name:reserved"`
}

func (o *MIBTcptable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBTcptable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBTcprow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBTcprow{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MIBTcptable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBTcprow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBTcprow{}
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

// MIBUdprow structure represents MIB_UDPROW RPC structure.
type MIBUdprow struct {
	LocalAddr uint32 `idl:"name:dwLocalAddr" json:"local_addr"`
	LocalPort uint32 `idl:"name:dwLocalPort" json:"local_port"`
}

func (o *MIBUdprow) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBUdprow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBUdprow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBUdpstats structure represents MIB_UDPSTATS RPC structure.
type MIBUdpstats struct {
	InDatagrams  uint32 `idl:"name:dwInDatagrams" json:"in_datagrams"`
	NoPorts      uint32 `idl:"name:dwNoPorts" json:"no_ports"`
	InErrors     uint32 `idl:"name:dwInErrors" json:"in_errors"`
	OutDatagrams uint32 `idl:"name:dwOutDatagrams" json:"out_datagrams"`
	AddrsLength  uint32 `idl:"name:dwNumAddrs" json:"addrs_length"`
}

func (o *MIBUdpstats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBUdpstats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MIBUdpstats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MIBUdptable structure represents MIB_UDPTABLE RPC structure.
type MIBUdptable struct {
	EntriesLength uint32       `idl:"name:dwNumEntries" json:"entries_length"`
	Table         []*MIBUdprow `idl:"name:table" json:"table"`
	_             []byte       `idl:"name:reserved"`
}

func (o *MIBUdptable) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MIBUdptable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&MIBUdprow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Table); uint64(i1) < 1; i1++ {
		if err := (&MIBUdprow{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MIBUdptable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	o.Table = make([]*MIBUdprow, 1)
	for i1 := range o.Table {
		i1 := i1
		if o.Table[i1] == nil {
			o.Table[i1] = &MIBUdprow{}
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

// MprServer0 structure represents MPR_SERVER_0 RPC structure.
type MprServer0 struct {
	LANOnlyMode bool   `idl:"name:fLanOnlyMode" json:"lan_only_mode"`
	UpTime      uint32 `idl:"name:dwUpTime" json:"up_time"`
	TotalPorts  uint32 `idl:"name:dwTotalPorts" json:"total_ports"`
	PortsInUse  uint32 `idl:"name:dwPortsInUse" json:"ports_in_use"`
}

func (o *MprServer0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServer0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MprServer0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MprServer1 structure represents MPR_SERVER_1 RPC structure.
type MprServer1 struct {
	PPTPPortsLength uint32 `idl:"name:dwNumPptpPorts" json:"pptp_ports_length"`
	PPTPPortFlags   uint32 `idl:"name:dwPptpPortFlags" json:"pptp_port_flags"`
	L2tpPortsLength uint32 `idl:"name:dwNumL2tpPorts" json:"l2tp_ports_length"`
	L2tpPortFlags   uint32 `idl:"name:dwL2tpPortFlags" json:"l2tp_port_flags"`
}

func (o *MprServer1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServer1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.L2tpPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.L2tpPortFlags); err != nil {
		return err
	}
	return nil
}
func (o *MprServer1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2tpPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2tpPortFlags); err != nil {
		return err
	}
	return nil
}

// MprServer2 structure represents MPR_SERVER_2 RPC structure.
type MprServer2 struct {
	PPTPPortsLength uint32 `idl:"name:dwNumPptpPorts" json:"pptp_ports_length"`
	PPTPPortFlags   uint32 `idl:"name:dwPptpPortFlags" json:"pptp_port_flags"`
	L2tpPortsLength uint32 `idl:"name:dwNumL2tpPorts" json:"l2tp_ports_length"`
	L2tpPortFlags   uint32 `idl:"name:dwL2tpPortFlags" json:"l2tp_port_flags"`
	SstpPortsLength uint32 `idl:"name:dwNumSstpPorts" json:"sstp_ports_length"`
	SstpPortFlags   uint32 `idl:"name:dwSstpPortFlags" json:"sstp_port_flags"`
}

func (o *MprServer2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprServer2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.L2tpPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.L2tpPortFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.SstpPortsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.SstpPortFlags); err != nil {
		return err
	}
	return nil
}
func (o *MprServer2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPTPPortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2tpPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.L2tpPortFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.SstpPortsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.SstpPortFlags); err != nil {
		return err
	}
	return nil
}

// PppNbfcpInfo structure represents PPP_NBFCP_INFO RPC structure.
type PppNbfcpInfo struct {
	Error       uint32   `idl:"name:dwError" json:"error"`
	Workstation []uint16 `idl:"name:wszWksta" json:"workstation"`
}

func (o *PppNbfcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppNbfcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PppNbfcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppIpcpInfo structure represents PPP_IPCP_INFO RPC structure.
type PppIpcpInfo struct {
	Error         uint32   `idl:"name:dwError" json:"error"`
	Address       []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
}

func (o *PppIpcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppIpcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PppIpcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppIpcpInfo2 structure represents PPP_IPCP_INFO2 RPC structure.
type PppIpcpInfo2 struct {
	Error         uint32   `idl:"name:dwError" json:"error"`
	Address       []uint16 `idl:"name:wszAddress" json:"address"`
	RemoteAddress []uint16 `idl:"name:wszRemoteAddress" json:"remote_address"`
	Options       uint32   `idl:"name:dwOptions" json:"options"`
	RemoteOptons  uint32   `idl:"name:dwRemoteOptons" json:"remote_optons"`
}

func (o *PppIpcpInfo2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppIpcpInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PppIpcpInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppIpxcpInfo structure represents PPP_IPXCP_INFO RPC structure.
type PppIpxcpInfo struct {
	Error   uint32   `idl:"name:dwError" json:"error"`
	Address []uint16 `idl:"name:wszAddress" json:"address"`
}

func (o *PppIpxcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppIpxcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PppIpxcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppIPv6CreatePartitionInfo structure represents PPP_IPV6_CP_INFO RPC structure.
type PppIPv6CreatePartitionInfo struct {
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

func (o *PppIPv6CreatePartitionInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppIPv6CreatePartitionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PppIPv6CreatePartitionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppAtcpInfo structure represents PPP_ATCP_INFO RPC structure.
type PppAtcpInfo struct {
	Error   uint32   `idl:"name:dwError" json:"error"`
	Address []uint16 `idl:"name:wszAddress" json:"address"`
}

func (o *PppAtcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppAtcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PppAtcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppCcpInfo structure represents PPP_CCP_INFO RPC structure.
type PppCcpInfo struct {
	Error                      uint32 `idl:"name:dwError" json:"error"`
	CompressionAlgorithm       uint32 `idl:"name:dwCompressionAlgorithm" json:"compression_algorithm"`
	Options                    uint32 `idl:"name:dwOptions" json:"options"`
	RemoteCompressionAlgorithm uint32 `idl:"name:dwRemoteCompressionAlgorithm" json:"remote_compression_algorithm"`
	RemoteOptions              uint32 `idl:"name:dwRemoteOptions" json:"remote_options"`
}

func (o *PppCcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppCcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PppCcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PppLcpInfo structure represents PPP_LCP_INFO RPC structure.
type PppLcpInfo struct {
	Error                        uint32 `idl:"name:dwError" json:"error"`
	AuthenticationProtocol       uint32 `idl:"name:dwAuthenticationProtocol" json:"authentication_protocol"`
	AuthenticationData           uint32 `idl:"name:dwAuthenticationData" json:"authentication_data"`
	RemoteAuthenticationProtocol uint32 `idl:"name:dwRemoteAuthenticationProtocol" json:"remote_authentication_protocol"`
	RemoteAuthenticationData     uint32 `idl:"name:dwRemoteAuthenticationData" json:"remote_authentication_data"`
	TerminateReason              uint32 `idl:"name:dwTerminateReason" json:"terminate_reason"`
	RemoteTerminateReason        uint32 `idl:"name:dwRemoteTerminateReason" json:"remote_terminate_reason"`
	Options                      uint32 `idl:"name:dwOptions" json:"options"`
	RemoteOptions                uint32 `idl:"name:dwRemoteOptions" json:"remote_options"`
	EapTypeID                    uint32 `idl:"name:dwEapTypeId" json:"eap_type_id"`
	RemoteEapTypeID              uint32 `idl:"name:dwRemoteEapTypeId" json:"remote_eap_type_id"`
}

func (o *PppLcpInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppLcpInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.EapTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteEapTypeID); err != nil {
		return err
	}
	return nil
}
func (o *PppLcpInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.EapTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteEapTypeID); err != nil {
		return err
	}
	return nil
}

// PppInfo structure represents PPP_INFO RPC structure.
type PppInfo struct {
	Nbf *PppNbfcpInfo `idl:"name:nbf" json:"nbf"`
	IP  *PppIpcpInfo  `idl:"name:ip" json:"ip"`
	Ipx *PppIpxcpInfo `idl:"name:ipx" json:"ipx"`
	AT  *PppAtcpInfo  `idl:"name:at" json:"at"`
}

func (o *PppInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Nbf != nil {
		if err := o.Nbf.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppNbfcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IP != nil {
		if err := o.IP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppIpcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Ipx != nil {
		if err := o.Ipx.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppIpxcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AT != nil {
		if err := o.AT.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppAtcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PppInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Nbf == nil {
		o.Nbf = &PppNbfcpInfo{}
	}
	if err := o.Nbf.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IP == nil {
		o.IP = &PppIpcpInfo{}
	}
	if err := o.IP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Ipx == nil {
		o.Ipx = &PppIpxcpInfo{}
	}
	if err := o.Ipx.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AT == nil {
		o.AT = &PppAtcpInfo{}
	}
	if err := o.AT.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PppInfo2 structure represents PPP_INFO_2 RPC structure.
type PppInfo2 struct {
	Nbf *PppNbfcpInfo `idl:"name:nbf" json:"nbf"`
	IP  *PppIpcpInfo2 `idl:"name:ip" json:"ip"`
	Ipx *PppIpxcpInfo `idl:"name:ipx" json:"ipx"`
	AT  *PppAtcpInfo  `idl:"name:at" json:"at"`
	Ccp *PppCcpInfo   `idl:"name:ccp" json:"ccp"`
	Lcp *PppLcpInfo   `idl:"name:lcp" json:"lcp"`
}

func (o *PppInfo2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Nbf != nil {
		if err := o.Nbf.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppNbfcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IP != nil {
		if err := o.IP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppIpcpInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Ipx != nil {
		if err := o.Ipx.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppIpxcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AT != nil {
		if err := o.AT.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppAtcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Ccp != nil {
		if err := o.Ccp.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppCcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Lcp != nil {
		if err := o.Lcp.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppLcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PppInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Nbf == nil {
		o.Nbf = &PppNbfcpInfo{}
	}
	if err := o.Nbf.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IP == nil {
		o.IP = &PppIpcpInfo2{}
	}
	if err := o.IP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Ipx == nil {
		o.Ipx = &PppIpxcpInfo{}
	}
	if err := o.Ipx.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AT == nil {
		o.AT = &PppAtcpInfo{}
	}
	if err := o.AT.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Ccp == nil {
		o.Ccp = &PppCcpInfo{}
	}
	if err := o.Ccp.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Lcp == nil {
		o.Lcp = &PppLcpInfo{}
	}
	if err := o.Lcp.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PppInfo3 structure represents PPP_INFO_3 RPC structure.
type PppInfo3 struct {
	Nbf  *PppNbfcpInfo               `idl:"name:nbf" json:"nbf"`
	IP   *PppIpcpInfo2               `idl:"name:ip" json:"ip"`
	IPv6 *PppIPv6CreatePartitionInfo `idl:"name:ipv6" json:"ipv6"`
	Ccp  *PppCcpInfo                 `idl:"name:ccp" json:"ccp"`
	Lcp  *PppLcpInfo                 `idl:"name:lcp" json:"lcp"`
}

func (o *PppInfo3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PppInfo3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Nbf != nil {
		if err := o.Nbf.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppNbfcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IP != nil {
		if err := o.IP.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppIpcpInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.IPv6 != nil {
		if err := o.IPv6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppIPv6CreatePartitionInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Ccp != nil {
		if err := o.Ccp.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppCcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Lcp != nil {
		if err := o.Lcp.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppLcpInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PppInfo3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Nbf == nil {
		o.Nbf = &PppNbfcpInfo{}
	}
	if err := o.Nbf.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IP == nil {
		o.IP = &PppIpcpInfo2{}
	}
	if err := o.IP.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.IPv6 == nil {
		o.IPv6 = &PppIPv6CreatePartitionInfo{}
	}
	if err := o.IPv6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Ccp == nil {
		o.Ccp = &PppCcpInfo{}
	}
	if err := o.Ccp.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Lcp == nil {
		o.Lcp = &PppLcpInfo{}
	}
	if err := o.Lcp.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RasiPort0 structure represents RASI_PORT_0 RPC structure.
type RasiPort0 struct {
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

func (o *RasiPort0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasiPort0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RasiPort0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RasiPort1 structure represents RASI_PORT_1 RPC structure.
type RasiPort1 struct {
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

func (o *RasiPort1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasiPort1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RasiPort1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RasiConnection0 structure represents RASI_CONNECTION_0 RPC structure.
type RasiConnection0 struct {
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

func (o *RasiConnection0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasiConnection0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RasiConnection0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RasiConnection1 structure represents RASI_CONNECTION_1 RPC structure.
type RasiConnection1 struct {
	Connection           uint32   `idl:"name:dwConnection" json:"connection"`
	Interface            uint32   `idl:"name:dwInterface" json:"interface"`
	PppInfo              *PppInfo `idl:"name:PppInfo" json:"ppp_info"`
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

func (o *RasiConnection1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasiConnection1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.PppInfo != nil {
		if err := o.PppInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppInfo{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RasiConnection1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connection); err != nil {
		return err
	}
	if err := w.ReadData(&o.Interface); err != nil {
		return err
	}
	if o.PppInfo == nil {
		o.PppInfo = &PppInfo{}
	}
	if err := o.PppInfo.UnmarshalNDR(ctx, w); err != nil {
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

// RasiConnection2 structure represents RASI_CONNECTION_2 RPC structure.
type RasiConnection2 struct {
	Connection    uint32              `idl:"name:dwConnection" json:"connection"`
	UserName      []uint16            `idl:"name:wszUserName" json:"user_name"`
	InterfaceType RouterInterfaceType `idl:"name:dwInterfaceType" json:"interface_type"`
	GUID          *dtyp.GUID          `idl:"name:guid" json:"guid"`
	PppInfo2      *PppInfo2           `idl:"name:PppInfo2" json:"ppp_info2"`
}

func (o *RasiConnection2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasiConnection2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.PppInfo2 != nil {
		if err := o.PppInfo2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RasiConnection2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.PppInfo2 == nil {
		o.PppInfo2 = &PppInfo2{}
	}
	if err := o.PppInfo2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RasiConnection3 structure represents RASI_CONNECTION_3 RPC structure.
type RasiConnection3 struct {
	Version            uint32              `idl:"name:dwVersion" json:"version"`
	Size               uint32              `idl:"name:dwSize" json:"size"`
	Connection         uint32              `idl:"name:dwConnection" json:"connection"`
	UserName           []uint16            `idl:"name:wszUserName" json:"user_name"`
	InterfaceType      RouterInterfaceType `idl:"name:dwInterfaceType" json:"interface_type"`
	GUID               *dtyp.GUID          `idl:"name:guid" json:"guid"`
	PppInfo3           *PppInfo3           `idl:"name:PppInfo3" json:"ppp_info3"`
	RASQuarantineState RASQuarantineState  `idl:"name:rasQuarState" json:"ras_quarantine_state"`
	Timer              *dtyp.Filetime      `idl:"name:timer" json:"timer"`
}

func (o *RasiConnection3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasiConnection3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.PppInfo3 != nil {
		if err := o.PppInfo3.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PppInfo3{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RasiConnection3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.PppInfo3 == nil {
		o.PppInfo3 = &PppInfo3{}
	}
	if err := o.PppInfo3.UnmarshalNDR(ctx, w); err != nil {
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

// MpriInterface0 structure represents MPRI_INTERFACE_0 RPC structure.
type MpriInterface0 struct {
	InterfaceName         []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface             uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled               bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType         RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState       RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError             uint32                `idl:"name:dwLastError" json:"last_error"`
}

func (o *MpriInterface0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MpriInterface0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MpriInterface0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MpriInterface1 structure represents MPRI_INTERFACE_1 RPC structure.
type MpriInterface1 struct {
	InterfaceName               []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface                   uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled                     bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType               RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState             RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons       uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError                   uint32                `idl:"name:dwLastError" json:"last_error"`
	LpwsDialoutHoursRestriction string                `idl:"name:lpwsDialoutHoursRestriction" json:"lpws_dialout_hours_restriction"`
}

func (o *MpriInterface1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MpriInterface1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.LpwsDialoutHoursRestriction != "" {
		_ptr_lpwsDialoutHoursRestriction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.LpwsDialoutHoursRestriction); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LpwsDialoutHoursRestriction, _ptr_lpwsDialoutHoursRestriction); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MpriInterface1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		if err := ndr.ReadUTF16String(ctx, w, &o.LpwsDialoutHoursRestriction); err != nil {
			return err
		}
		return nil
	})
	_s_lpwsDialoutHoursRestriction := func(ptr interface{}) { o.LpwsDialoutHoursRestriction = *ptr.(*string) }
	if err := w.ReadPointer(&o.LpwsDialoutHoursRestriction, _s_lpwsDialoutHoursRestriction, _ptr_lpwsDialoutHoursRestriction); err != nil {
		return err
	}
	return nil
}

// MpriInterface2 structure represents MPRI_INTERFACE_2 RPC structure.
type MpriInterface2 struct {
	InterfaceName            []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface                uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled                  bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType            RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState          RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons    uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError                uint32                `idl:"name:dwLastError" json:"last_error"`
	DwfOptions               uint32                `idl:"name:dwfOptions" json:"dwf_options"`
	LocalPhoneNumber         []uint16              `idl:"name:szLocalPhoneNumber" json:"local_phone_number"`
	Alternates               string                `idl:"name:szAlternates" json:"alternates"`
	Ipaddr                   uint32                `idl:"name:ipaddr" json:"ipaddr"`
	IpaddrDNS                uint32                `idl:"name:ipaddrDns" json:"ipaddr_dns"`
	IpaddrDNSAlt             uint32                `idl:"name:ipaddrDnsAlt" json:"ipaddr_dns_alt"`
	IpaddrWINS               uint32                `idl:"name:ipaddrWins" json:"ipaddr_wins"`
	IpaddrWINSAlt            uint32                `idl:"name:ipaddrWinsAlt" json:"ipaddr_wins_alt"`
	DwfNetProtocols          uint32                `idl:"name:dwfNetProtocols" json:"dwf_net_protocols"`
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
	VpnStrategy              uint32                `idl:"name:dwVpnStrategy" json:"vpn_strategy"`
}

func (o *MpriInterface2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MpriInterface2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DwfOptions); err != nil {
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
	if err := w.WriteData(o.Ipaddr); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrDNS); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrDNSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrWINSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.DwfNetProtocols); err != nil {
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
	if err := w.WriteData(o.VpnStrategy); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *MpriInterface2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DwfOptions); err != nil {
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
	if err := w.ReadData(&o.Ipaddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrDNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrDNSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrWINSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.DwfNetProtocols); err != nil {
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
	if err := w.ReadData(&o.VpnStrategy); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MpriInterface3 structure represents MPRI_INTERFACE_3 RPC structure.
type MpriInterface3 struct {
	InterfaceName            []uint16              `idl:"name:wszInterfaceName" json:"interface_name"`
	Interface                uint32                `idl:"name:dwInterface" json:"interface"`
	Enabled                  bool                  `idl:"name:fEnabled" json:"enabled"`
	InterfaceType            RouterInterfaceType   `idl:"name:dwIfType" json:"interface_type"`
	ConnectionState          RouterConnectionState `idl:"name:dwConnectionState" json:"connection_state"`
	UnReachabilityReasons    uint32                `idl:"name:fUnReachabilityReasons" json:"un_reachability_reasons"`
	LastError                uint32                `idl:"name:dwLastError" json:"last_error"`
	DwfOptions               uint32                `idl:"name:dwfOptions" json:"dwf_options"`
	LocalPhoneNumber         []uint16              `idl:"name:szLocalPhoneNumber" json:"local_phone_number"`
	Alternates               string                `idl:"name:szAlternates" json:"alternates"`
	Ipaddr                   uint32                `idl:"name:ipaddr" json:"ipaddr"`
	IpaddrDNS                uint32                `idl:"name:ipaddrDns" json:"ipaddr_dns"`
	IpaddrDNSAlt             uint32                `idl:"name:ipaddrDnsAlt" json:"ipaddr_dns_alt"`
	IpaddrWINS               uint32                `idl:"name:ipaddrWins" json:"ipaddr_wins"`
	IpaddrWINSAlt            uint32                `idl:"name:ipaddrWinsAlt" json:"ipaddr_wins_alt"`
	DwfNetProtocols          uint32                `idl:"name:dwfNetProtocols" json:"dwf_net_protocols"`
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
	VpnStrategy              uint32                `idl:"name:dwVpnStrategy" json:"vpn_strategy"`
	AddressCount             uint32                `idl:"name:AddressCount" json:"address_count"`
	Ipv6addrDNS              *In6Addr              `idl:"name:ipv6addrDns" json:"ipv6addr_dns"`
	Ipv6addrDNSAlt           *In6Addr              `idl:"name:ipv6addrDnsAlt" json:"ipv6addr_dns_alt"`
	Ipv6addr                 *In6Addr              `idl:"name:ipv6addr" json:"ipv6addr"`
}

func (o *MpriInterface3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MpriInterface3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.DwfOptions); err != nil {
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
	if err := w.WriteData(o.Ipaddr); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrDNS); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrDNSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrWINS); err != nil {
		return err
	}
	if err := w.WriteData(o.IpaddrWINSAlt); err != nil {
		return err
	}
	if err := w.WriteData(o.DwfNetProtocols); err != nil {
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
	if err := w.WriteData(o.VpnStrategy); err != nil {
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
		if err := (&In6Addr{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Ipv6addrDNSAlt != nil {
		if err := o.Ipv6addrDNSAlt.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&In6Addr{}).MarshalNDR(ctx, w); err != nil {
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
				if err := (&In6Addr{}).MarshalNDR(ctx, w); err != nil {
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
func (o *MpriInterface3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.DwfOptions); err != nil {
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
	if err := w.ReadData(&o.Ipaddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrDNS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrDNSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrWINS); err != nil {
		return err
	}
	if err := w.ReadData(&o.IpaddrWINSAlt); err != nil {
		return err
	}
	if err := w.ReadData(&o.DwfNetProtocols); err != nil {
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
	if err := w.ReadData(&o.VpnStrategy); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressCount); err != nil {
		return err
	}
	if o.Ipv6addrDNS == nil {
		o.Ipv6addrDNS = &In6Addr{}
	}
	if err := o.Ipv6addrDNS.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Ipv6addrDNSAlt == nil {
		o.Ipv6addrDNSAlt = &In6Addr{}
	}
	if err := o.Ipv6addrDNSAlt.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ipv6addr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Ipv6addr == nil {
			o.Ipv6addr = &In6Addr{}
		}
		if err := o.Ipv6addr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ipv6addr := func(ptr interface{}) { o.Ipv6addr = *ptr.(**In6Addr) }
	if err := w.ReadPointer(&o.Ipv6addr, _s_ipv6addr, _ptr_ipv6addr); err != nil {
		return err
	}
	return nil
}

// MprDevice0 structure represents MPR_DEVICE_0 RPC structure.
type MprDevice0 struct {
	DeviceType []uint16 `idl:"name:szDeviceType" json:"device_type"`
	DeviceName []uint16 `idl:"name:szDeviceName" json:"device_name"`
}

func (o *MprDevice0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprDevice0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MprDevice0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MprDevice1 structure represents MPR_DEVICE_1 RPC structure.
type MprDevice1 struct {
	DeviceType       []uint16 `idl:"name:szDeviceType" json:"device_type"`
	DeviceName       []uint16 `idl:"name:szDeviceName" json:"device_name"`
	LocalPhoneNumber []uint16 `idl:"name:szLocalPhoneNumber" json:"local_phone_number"`
	Alternates       string   `idl:"name:szAlternates" json:"alternates"`
}

func (o *MprDevice1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprDevice1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MprDevice1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MprCredentialsex1 structure represents MPR_CREDENTIALSEX_1 RPC structure.
type MprCredentialsex1 struct {
	Size   uint32 `idl:"name:dwSize" json:"size"`
	Offset uint32 `idl:"name:dwOffset" json:"offset"`
	Data   []byte `idl:"name:bData" json:"data"`
}

func (o *MprCredentialsex1) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprCredentialsex1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MprCredentialsex1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IffilterInfo structure represents IFFILTER_INFO RPC structure.
type IffilterInfo struct {
	EnableFragChk bool `idl:"name:bEnableFragChk" json:"enable_frag_chk"`
}

func (o *IffilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IffilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IffilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MprFilter0 structure represents MPR_FILTER_0 RPC structure.
type MprFilter0 struct {
	Enable bool `idl:"name:fEnable" json:"enable"`
}

func (o *MprFilter0) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MprFilter0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MprFilter0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxGlobalInfo structure represents IPX_GLOBAL_INFO RPC structure.
type IpxGlobalInfo struct {
	RoutingTableHashSize uint32 `idl:"name:RoutingTableHashSize" json:"routing_table_hash_size"`
	EventLogMask         uint32 `idl:"name:EventLogMask" json:"event_log_mask"`
}

func (o *IpxGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxInterfaceInfo structure represents IPX_IF_INFO RPC structure.
type IpxInterfaceInfo struct {
	AdministratorState uint32 `idl:"name:AdministratorState" json:"administrator_state"`
	NetBIOSAccept      uint32 `idl:"name:NetbiosAccept" json:"netbios_accept"`
	NetBIOSDeliver     uint32 `idl:"name:NetbiosDeliver" json:"netbios_deliver"`
}

func (o *IpxInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxStaticRouteInfo structure represents IPX_STATIC_ROUTE_INFO RPC structure.
type IpxStaticRouteInfo struct {
	Field1            *IpxStaticRouteInfo_Field1 `idl:"name:" json:""`
	TickCount         uint16                     `idl:"name:TickCount" json:"tick_count"`
	HopCount          uint16                     `idl:"name:HopCount" json:"hop_count"`
	NextHopMACAddress []byte                     `idl:"name:NextHopMacAddress" json:"next_hop_mac_address"`
}

func (o *IpxStaticRouteInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxStaticRouteInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxStaticRouteInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type IpxStaticRouteInfo_Field1 struct {
	DwordAlign uint32 `idl:"name:DwordAlign" json:"dword_align"`
	Network    []byte `idl:"name:Network" json:"network"`
}

// IpxStaticServiceInfo structure represents IPX_STATIC_SERVICE_INFO RPC structure.
type IpxStaticServiceInfo IpxServerEntry

func (o *IpxStaticServiceInfo) IpxServerEntry() *IpxServerEntry { return (*IpxServerEntry)(o) }

func (o *IpxStaticServiceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxStaticServiceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxStaticServiceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxServerEntry structure represents IPX_SERVER_ENTRY RPC structure.
type IpxServerEntry struct {
	Type     uint16 `idl:"name:Type" json:"type"`
	Name     []byte `idl:"name:Name" json:"name"`
	Network  []byte `idl:"name:Network" json:"network"`
	Node     []byte `idl:"name:Node" json:"node"`
	Socket   []byte `idl:"name:Socket" json:"socket"`
	HopCount uint16 `idl:"name:HopCount" json:"hop_count"`
}

func (o *IpxServerEntry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxServerEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxServerEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxStaticNetBIOSNameInfo structure represents IPX_STATIC_NETBIOS_NAME_INFO RPC structure.
type IpxStaticNetBIOSNameInfo struct {
	Field1 *IpxStaticNetBIOSNameInfo_Field1 `idl:"name:" json:""`
}

func (o *IpxStaticNetBIOSNameInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxStaticNetBIOSNameInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// FIXME unknown type
	return nil
}
func (o *IpxStaticNetBIOSNameInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// FIXME: unknown type
	return nil
}

type IpxStaticNetBIOSNameInfo_Field1 struct {
	DwordAlign uint32 `idl:"name:DwordAlign" json:"dword_align"`
	Name       []byte `idl:"name:Name" json:"name"`
}

// IpxAdapterInfo structure represents IPX_ADAPTER_INFO RPC structure.
type IpxAdapterInfo struct {
	PacketType  uint32   `idl:"name:PacketType" json:"packet_type"`
	AdapterName []uint16 `idl:"name:AdapterName" json:"adapter_name"`
}

func (o *IpxAdapterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxAdapterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxAdapterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxTrafficFilterGlobalInfo structure represents IPX_TRAFFIC_FILTER_GLOBAL_INFO RPC structure.
type IpxTrafficFilterGlobalInfo struct {
	FilterAction uint32 `idl:"name:FilterAction" json:"filter_action"`
}

func (o *IpxTrafficFilterGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxTrafficFilterGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxTrafficFilterGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.FilterAction); err != nil {
		return err
	}
	return nil
}

// IpxTrafficFilterInfo structure represents IPX_TRAFFIC_FILTER_INFO RPC structure.
type IpxTrafficFilterInfo struct {
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

func (o *IpxTrafficFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxTrafficFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxTrafficFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type IpxMIBIndex struct {
	InterfaceTableIndex      *InterfaceTableIndex      `idl:"name:InterfaceTableIndex" json:"interface_table_index"`
	RoutingTableIndex        *RoutingTableIndex        `idl:"name:RoutingTableIndex" json:"routing_table_index"`
	StaticRoutesTableIndex   *StaticRoutesTableIndex   `idl:"name:StaticRoutesTableIndex" json:"static_routes_table_index"`
	ServicesTableIndex       *ServicesTableIndex       `idl:"name:ServicesTableIndex" json:"services_table_index"`
	StaticServicesTableIndex *StaticServicesTableIndex `idl:"name:StaticServicesTableIndex" json:"static_services_table_index"`
}

// IpxMIBGetInputData structure represents IPX_MIB_GET_INPUT_DATA RPC structure.
type IpxMIBGetInputData struct {
	TableID  uint32       `idl:"name:TableId" json:"table_id"`
	MIBIndex *IpxMIBIndex `idl:"name:MibIndex" json:"mib_index"`
}

func (o *IpxMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxInterfaceStats structure represents IPX_IF_STATS RPC structure.
type IpxInterfaceStats struct {
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

func (o *IpxInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxInterface structure represents IPX_INTERFACE RPC structure.
type IpxInterface struct {
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
	EnableIpxWANNegotiation uint32             `idl:"name:EnableIpxWanNegotiation" json:"enable_ipx_wan_negotiation"`
	InterfaceStats          *IpxInterfaceStats `idl:"name:IfStats" json:"interface_stats"`
}

func (o *IpxInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.EnableIpxWANNegotiation); err != nil {
		return err
	}
	if o.InterfaceStats != nil {
		if err := o.InterfaceStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IpxInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IpxInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.EnableIpxWANNegotiation); err != nil {
		return err
	}
	if o.InterfaceStats == nil {
		o.InterfaceStats = &IpxInterfaceStats{}
	}
	if err := o.InterfaceStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IpxRoute structure represents IPX_ROUTE RPC structure.
type IpxRoute struct {
	InterfaceIndex    uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
	Protocol          uint32 `idl:"name:Protocol" json:"protocol"`
	Network           []byte `idl:"name:Network" json:"network"`
	TickCount         uint16 `idl:"name:TickCount" json:"tick_count"`
	HopCount          uint16 `idl:"name:HopCount" json:"hop_count"`
	NextHopMACAddress []byte `idl:"name:NextHopMacAddress" json:"next_hop_mac_address"`
	Flags             uint32 `idl:"name:Flags" json:"flags"`
}

func (o *IpxRoute) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxRoute) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxRoute) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpxService structure represents IPX_SERVICE RPC structure.
type IpxService struct {
	InterfaceIndex uint32          `idl:"name:InterfaceIndex" json:"interface_index"`
	Protocol       uint32          `idl:"name:Protocol" json:"protocol"`
	Server         *IpxServerEntry `idl:"name:Server" json:"server"`
}

func (o *IpxService) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxService) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&IpxServerEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IpxService) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Server = &IpxServerEntry{}
	}
	if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

type IpxMIBRow struct {
	Interface *IpxInterface `idl:"name:Interface" json:"interface"`
	Route     *IpxRoute     `idl:"name:Route" json:"route"`
	Service   *IpxService   `idl:"name:Service" json:"service"`
}

// IpxMIBSetInputData structure represents IPX_MIB_SET_INPUT_DATA RPC structure.
type IpxMIBSetInputData struct {
	TableID uint32     `idl:"name:TableId" json:"table_id"`
	MIBRow  *IpxMIBRow `idl:"name:MibRow" json:"mib_row"`
}

func (o *IpxMIBSetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpxMIBSetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpxMIBSetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	// FIXME: unknown type MibRow
	return nil
}

// SapServiceFilterInfo structure represents SAP_SERVICE_FILTER_INFO RPC structure.
type SapServiceFilterInfo struct {
	Field1      *SapServiceFilterInfo_Field1 `idl:"name:" json:""`
	ServiceName []byte                       `idl:"name:ServiceName" json:"service_name"`
}

func (o *SapServiceFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapServiceFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SapServiceFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type SapServiceFilterInfo_Field1 struct {
	ServiceType      uint16 `idl:"name:ServiceType" json:"service_type"`
	ServiceTypeAlign uint32 `idl:"name:ServiceType_align" json:"service_type_align"`
}

// SapInterfaceFilters structure represents SAP_IF_FILTERS RPC structure.
type SapInterfaceFilters struct {
	SupplyFilterAction uint32                  `idl:"name:SupplyFilterAction" json:"supply_filter_action"`
	SupplyFilterCount  uint32                  `idl:"name:SupplyFilterCount" json:"supply_filter_count"`
	ListenFilterAction uint32                  `idl:"name:ListenFilterAction" json:"listen_filter_action"`
	ListenFilterCount  uint32                  `idl:"name:ListenFilterCount" json:"listen_filter_count"`
	ServiceFilter      []*SapServiceFilterInfo `idl:"name:ServiceFilter" json:"service_filter"`
}

func (o *SapInterfaceFilters) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapInterfaceFilters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&SapServiceFilterInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.ServiceFilter); uint64(i1) < 1; i1++ {
		if err := (&SapServiceFilterInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *SapInterfaceFilters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	o.ServiceFilter = make([]*SapServiceFilterInfo, 1)
	for i1 := range o.ServiceFilter {
		i1 := i1
		if o.ServiceFilter[i1] == nil {
			o.ServiceFilter[i1] = &SapServiceFilterInfo{}
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

// SapInterfaceInfo structure represents SAP_IF_INFO RPC structure.
type SapInterfaceInfo struct {
	AdminState             uint32 `idl:"name:AdminState" json:"admin_state"`
	UpdateMode             uint32 `idl:"name:UpdateMode" json:"update_mode"`
	PacketType             uint32 `idl:"name:PacketType" json:"packet_type"`
	Supply                 uint32 `idl:"name:Supply" json:"supply"`
	Listen                 uint32 `idl:"name:Listen" json:"listen"`
	GetNearestServerReply  uint32 `idl:"name:GetNearestServerReply" json:"get_nearest_server_reply"`
	PeriodicUpdateInterval uint32 `idl:"name:PeriodicUpdateInterval" json:"periodic_update_interval"`
	AgeIntervalMultiplier  uint32 `idl:"name:AgeIntervalMultiplier" json:"age_interval_multiplier"`
}

func (o *SapInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SapInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// SapInterfaceConfig structure represents SAP_IF_CONFIG RPC structure.
type SapInterfaceConfig struct {
	SapInterfaceInfo    *SapInterfaceInfo    `idl:"name:SapIfInfo" json:"sap_interface_info"`
	SapInterfaceFilters *SapInterfaceFilters `idl:"name:SapIfFilters" json:"sap_interface_filters"`
}

func (o *SapInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.SapInterfaceInfo != nil {
		if err := o.SapInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SapInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SapInterfaceFilters != nil {
		if err := o.SapInterfaceFilters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SapInterfaceFilters{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SapInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.SapInterfaceInfo == nil {
		o.SapInterfaceInfo = &SapInterfaceInfo{}
	}
	if err := o.SapInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SapInterfaceFilters == nil {
		o.SapInterfaceFilters = &SapInterfaceFilters{}
	}
	if err := o.SapInterfaceFilters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SapMIBBase structure represents SAP_MIB_BASE RPC structure.
type SapMIBBase struct {
	SapOperatorState uint32 `idl:"name:SapOperState" json:"sap_operator_state"`
}

func (o *SapMIBBase) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapMIBBase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SapOperatorState); err != nil {
		return err
	}
	return nil
}
func (o *SapMIBBase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SapOperatorState); err != nil {
		return err
	}
	return nil
}

// SapInterfaceStats structure represents SAP_IF_STATS RPC structure.
type SapInterfaceStats struct {
	SapInterfaceOperatorState uint32 `idl:"name:SapIfOperState" json:"sap_interface_operator_state"`
	SapInterfaceInputPackets  uint32 `idl:"name:SapIfInputPackets" json:"sap_interface_input_packets"`
	SapInterfaceOutputPackets uint32 `idl:"name:SapIfOutputPackets" json:"sap_interface_output_packets"`
}

func (o *SapInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SapInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.WriteData(o.SapInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.SapInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}
func (o *SapInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SapInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.ReadData(&o.SapInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.SapInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}

// SapInterface structure represents SAP_INTERFACE RPC structure.
type SapInterface struct {
	InterfaceIndex    uint32             `idl:"name:InterfaceIndex" json:"interface_index"`
	SapInterfaceInfo  *SapInterfaceInfo  `idl:"name:SapIfInfo" json:"sap_interface_info"`
	SapInterfaceStats *SapInterfaceStats `idl:"name:SapIfStats" json:"sap_interface_stats"`
}

func (o *SapInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if o.SapInterfaceInfo != nil {
		if err := o.SapInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SapInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SapInterfaceStats != nil {
		if err := o.SapInterfaceStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SapInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SapInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if o.SapInterfaceInfo == nil {
		o.SapInterfaceInfo = &SapInterfaceInfo{}
	}
	if err := o.SapInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SapInterfaceStats == nil {
		o.SapInterfaceStats = &SapInterfaceStats{}
	}
	if err := o.SapInterfaceStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SapMIBGetInputData structure represents SAP_MIB_GET_INPUT_DATA RPC structure.
type SapMIBGetInputData struct {
	TableID        uint32 `idl:"name:TableId" json:"table_id"`
	InterfaceIndex uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
}

func (o *SapMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SapMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// SapMIBSetInputData structure represents SAP_MIB_SET_INPUT_DATA RPC structure.
type SapMIBSetInputData struct {
	TableID      uint32        `idl:"name:TableId" json:"table_id"`
	SapInterface *SapInterface `idl:"name:SapInterface" json:"sap_interface"`
}

func (o *SapMIBSetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapMIBSetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	if o.SapInterface != nil {
		if err := o.SapInterface.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SapInterface{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SapMIBSetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	if o.SapInterface == nil {
		o.SapInterface = &SapInterface{}
	}
	if err := o.SapInterface.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RipmibBase structure represents RIPMIB_BASE RPC structure.
type RipmibBase struct {
	RipOperatorState uint32 `idl:"name:RIPOperState" json:"rip_operator_state"`
}

func (o *RipmibBase) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipmibBase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RipOperatorState); err != nil {
		return err
	}
	return nil
}
func (o *RipmibBase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RipOperatorState); err != nil {
		return err
	}
	return nil
}

// RipInterfaceStats structure represents RIP_IF_STATS RPC structure.
type RipInterfaceStats struct {
	RipInterfaceOperatorState uint32 `idl:"name:RipIfOperState" json:"rip_interface_operator_state"`
	RipInterfaceInputPackets  uint32 `idl:"name:RipIfInputPackets" json:"rip_interface_input_packets"`
	RipInterfaceOutputPackets uint32 `idl:"name:RipIfOutputPackets" json:"rip_interface_output_packets"`
}

func (o *RipInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RipInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.WriteData(o.RipInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.RipInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}
func (o *RipInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RipInterfaceOperatorState); err != nil {
		return err
	}
	if err := w.ReadData(&o.RipInterfaceInputPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.RipInterfaceOutputPackets); err != nil {
		return err
	}
	return nil
}

// RipInterfaceInfo structure represents RIP_IF_INFO RPC structure.
type RipInterfaceInfo struct {
	AdminState             uint32 `idl:"name:AdminState" json:"admin_state"`
	UpdateMode             uint32 `idl:"name:UpdateMode" json:"update_mode"`
	PacketType             uint32 `idl:"name:PacketType" json:"packet_type"`
	Supply                 uint32 `idl:"name:Supply" json:"supply"`
	Listen                 uint32 `idl:"name:Listen" json:"listen"`
	PeriodicUpdateInterval uint32 `idl:"name:PeriodicUpdateInterval" json:"periodic_update_interval"`
	AgeIntervalMultiplier  uint32 `idl:"name:AgeIntervalMultiplier" json:"age_interval_multiplier"`
}

func (o *RipInterfaceInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipInterfaceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RipInterfaceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RipInterface structure represents RIP_INTERFACE RPC structure.
type RipInterface struct {
	InterfaceIndex    uint32             `idl:"name:InterfaceIndex" json:"interface_index"`
	RipInterfaceInfo  *RipInterfaceInfo  `idl:"name:RipIfInfo" json:"rip_interface_info"`
	RipInterfaceStats *RipInterfaceStats `idl:"name:RipIfStats" json:"rip_interface_stats"`
}

func (o *RipInterface) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIndex); err != nil {
		return err
	}
	if o.RipInterfaceInfo != nil {
		if err := o.RipInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RipInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RipInterfaceStats != nil {
		if err := o.RipInterfaceStats.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RipInterfaceStats{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RipInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIndex); err != nil {
		return err
	}
	if o.RipInterfaceInfo == nil {
		o.RipInterfaceInfo = &RipInterfaceInfo{}
	}
	if err := o.RipInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RipInterfaceStats == nil {
		o.RipInterfaceStats = &RipInterfaceStats{}
	}
	if err := o.RipInterfaceStats.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RipMIBGetInputData structure represents RIP_MIB_GET_INPUT_DATA RPC structure.
type RipMIBGetInputData struct {
	TableID        uint32 `idl:"name:TableId" json:"table_id"`
	InterfaceIndex uint32 `idl:"name:InterfaceIndex" json:"interface_index"`
}

func (o *RipMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RipMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RipMIBSetInputData structure represents RIP_MIB_SET_INPUT_DATA RPC structure.
type RipMIBSetInputData struct {
	TableID      uint32        `idl:"name:TableId" json:"table_id"`
	RipInterface *RipInterface `idl:"name:RipInterface" json:"rip_interface"`
}

func (o *RipMIBSetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipMIBSetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TableID); err != nil {
		return err
	}
	if o.RipInterface != nil {
		if err := o.RipInterface.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RipInterface{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RipMIBSetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TableID); err != nil {
		return err
	}
	if o.RipInterface == nil {
		o.RipInterface = &RipInterface{}
	}
	if err := o.RipInterface.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// EaptlsHash structure represents EAPTLS_HASH RPC structure.
type EaptlsHash struct {
	HashLength uint32 `idl:"name:cbHash" json:"hash_length"`
	Hash       []byte `idl:"name:pbHash" json:"hash"`
}

func (o *EaptlsHash) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EaptlsHash) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EaptlsHash) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EaptlsUserProperties structure represents EAPTLS_USER_PROPERTIES RPC structure.
type EaptlsUserProperties struct {
	_             uint32      `idl:"name:reserved"`
	Version       uint32      `idl:"name:dwVersion" json:"version"`
	Size          uint32      `idl:"name:dwSize" json:"size"`
	Flags         uint32      `idl:"name:fFlags" json:"flags"`
	Hash          *EaptlsHash `idl:"name:Hash" json:"hash"`
	DiffUser      string      `idl:"name:pwszDiffUser" json:"diff_user"`
	PinOffset     uint32      `idl:"name:dwPinOffset" json:"pin_offset"`
	Pin           string      `idl:"name:pwszPin" json:"pin"`
	Length        uint16      `idl:"name:usLength" json:"length"`
	MaximumLength uint16      `idl:"name:usMaximumLength" json:"maximum_length"`
	UcSeed        uint8       `idl:"name:ucSeed" json:"uc_seed"`
	AwszString    []uint16    `idl:"name:awszString" json:"awsz_string"`
}

func (o *EaptlsUserProperties) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EaptlsUserProperties) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&EaptlsHash{}).MarshalNDR(ctx, w); err != nil {
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
	if err := w.WriteData(o.PinOffset); err != nil {
		return err
	}
	if o.Pin != "" {
		_ptr_pwszPin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Pin); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Pin, _ptr_pwszPin); err != nil {
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
	if err := w.WriteData(o.UcSeed); err != nil {
		return err
	}
	for i1 := range o.AwszString {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.AwszString[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AwszString); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *EaptlsUserProperties) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Hash = &EaptlsHash{}
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
	if err := w.ReadData(&o.PinOffset); err != nil {
		return err
	}
	_ptr_pwszPin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Pin); err != nil {
			return err
		}
		return nil
	})
	_s_pwszPin := func(ptr interface{}) { o.Pin = *ptr.(*string) }
	if err := w.ReadPointer(&o.Pin, _s_pwszPin, _ptr_pwszPin); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.UcSeed); err != nil {
		return err
	}
	o.AwszString = make([]uint16, 1)
	for i1 := range o.AwszString {
		i1 := i1
		if err := w.ReadData(&o.AwszString[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// IpbootpGlobalConfig structure represents IPBOOTP_GLOBAL_CONFIG RPC structure.
type IpbootpGlobalConfig struct {
	GCLoggingLevel     uint32 `idl:"name:GC_LoggingLevel" json:"gc_logging_level"`
	GCMaxRecvQueueSize uint32 `idl:"name:GC_MaxRecvQueueSize" json:"gc_max_recv_queue_size"`
	GCServerCount      uint32 `idl:"name:GC_ServerCount" json:"gc_server_count"`
}

func (o *IpbootpGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.GCLoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.GCMaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.GCServerCount); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCLoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCMaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCServerCount); err != nil {
		return err
	}
	return nil
}

// IpbootpInterfaceConfig structure represents IPBOOTP_IF_CONFIG RPC structure.
type IpbootpInterfaceConfig struct {
	ICState               uint32 `idl:"name:IC_State" json:"ic_state"`
	ICRelayMode           uint32 `idl:"name:IC_RelayMode" json:"ic_relay_mode"`
	ICMaxHopCount         uint32 `idl:"name:IC_MaxHopCount" json:"ic_max_hop_count"`
	ICMinSecondsSinceBoot uint32 `idl:"name:IC_MinSecondsSinceBoot" json:"ic_min_seconds_since_boot"`
}

func (o *IpbootpInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ICState); err != nil {
		return err
	}
	if err := w.WriteData(o.ICRelayMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICMaxHopCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ICMinSecondsSinceBoot); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICState); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICRelayMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICMaxHopCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICMinSecondsSinceBoot); err != nil {
		return err
	}
	return nil
}

// IpbootpMIBGetInputData structure represents IPBOOTP_MIB_GET_INPUT_DATA RPC structure.
type IpbootpMIBGetInputData struct {
	ImgidTypeID         uint32 `idl:"name:IMGID_TypeID" json:"imgid_type_id"`
	ImgidInterfaceIndex uint32 `idl:"name:IMGID_IfIndex" json:"imgid_interface_index"`
}

func (o *IpbootpMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgidTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgidInterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgidTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgidInterfaceIndex); err != nil {
		return err
	}
	return nil
}

// IpbootpMIBGetOutputData structure represents IPBOOTP_MIB_GET_OUTPUT_DATA RPC structure.
type IpbootpMIBGetOutputData struct {
	ImgodTypeID         uint32 `idl:"name:IMGOD_TypeID" json:"imgod_type_id"`
	ImgodInterfaceIndex uint32 `idl:"name:IMGOD_IfIndex" json:"imgod_interface_index"`
	ImgodBuffer         []byte `idl:"name:IMGOD_Buffer" json:"imgod_buffer"`
}

func (o *IpbootpMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgodTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgodInterfaceIndex); err != nil {
		return err
	}
	for i1 := range o.ImgodBuffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.ImgodBuffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ImgodBuffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgodTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgodInterfaceIndex); err != nil {
		return err
	}
	o.ImgodBuffer = make([]byte, 1)
	for i1 := range o.ImgodBuffer {
		i1 := i1
		if err := w.ReadData(&o.ImgodBuffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IpbootpInterfaceStats structure represents IPBOOTP_IF_STATS RPC structure.
type IpbootpInterfaceStats struct {
	IsState             uint32 `idl:"name:IS_State" json:"is_state"`
	IsSendFailures      uint32 `idl:"name:IS_SendFailures" json:"is_send_failures"`
	IsReceiveFailures   uint32 `idl:"name:IS_ReceiveFailures" json:"is_receive_failures"`
	IsArpUpdateFailures uint32 `idl:"name:IS_ArpUpdateFailures" json:"is_arp_update_failures"`
	IsRequestsReceived  uint32 `idl:"name:IS_RequestsReceived" json:"is_requests_received"`
	IsRequestsDiscarded uint32 `idl:"name:IS_RequestsDiscarded" json:"is_requests_discarded"`
	IsRepliesReceived   uint32 `idl:"name:IS_RepliesReceived" json:"is_replies_received"`
	IsRepliesDiscarded  uint32 `idl:"name:IS_RepliesDiscarded" json:"is_replies_discarded"`
}

func (o *IpbootpInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IsState); err != nil {
		return err
	}
	if err := w.WriteData(o.IsSendFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.IsReceiveFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.IsArpUpdateFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRequestsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRequestsDiscarded); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRepliesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRepliesDiscarded); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsState); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsSendFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsReceiveFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsArpUpdateFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRequestsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRequestsDiscarded); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRepliesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRepliesDiscarded); err != nil {
		return err
	}
	return nil
}

// IpbootpInterfaceBinding structure represents IPBOOTP_IF_BINDING RPC structure.
type IpbootpInterfaceBinding struct {
	IbState     uint32 `idl:"name:IB_State" json:"ib_state"`
	IbAddrCount uint32 `idl:"name:IB_AddrCount" json:"ib_addr_count"`
}

func (o *IpbootpInterfaceBinding) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpInterfaceBinding) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IbState); err != nil {
		return err
	}
	if err := w.WriteData(o.IbAddrCount); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpInterfaceBinding) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IbState); err != nil {
		return err
	}
	if err := w.ReadData(&o.IbAddrCount); err != nil {
		return err
	}
	return nil
}

// IpbootpIPAddress structure represents IPBOOTP_IP_ADDRESS RPC structure.
type IpbootpIPAddress struct {
	IaAddress uint32 `idl:"name:IA_Address" json:"ia_address"`
	IaNetmask uint32 `idl:"name:IA_Netmask" json:"ia_netmask"`
}

func (o *IpbootpIPAddress) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpIPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IaAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.IaNetmask); err != nil {
		return err
	}
	return nil
}
func (o *IpbootpIPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IaAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.IaNetmask); err != nil {
		return err
	}
	return nil
}

// V6rMIBGetOutputData structure represents DHCPV6R_MIB_GET_OUTPUT_DATA RPC structure.
type V6rMIBGetOutputData struct {
	ImgodTypeID         uint32 `idl:"name:IMGOD_TypeID" json:"imgod_type_id"`
	ImgodInterfaceIndex uint32 `idl:"name:IMGOD_IfIndex" json:"imgod_interface_index"`
	ImgodBuffer         []byte `idl:"name:IMGOD_Buffer" json:"imgod_buffer"`
}

func (o *V6rMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *V6rMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgodTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgodInterfaceIndex); err != nil {
		return err
	}
	for i1 := range o.ImgodBuffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.ImgodBuffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ImgodBuffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *V6rMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgodTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgodInterfaceIndex); err != nil {
		return err
	}
	o.ImgodBuffer = make([]byte, 1)
	for i1 := range o.ImgodBuffer {
		i1 := i1
		if err := w.ReadData(&o.ImgodBuffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// V6rInterfaceStats structure represents DHCPV6R_IF_STATS RPC structure.
type V6rInterfaceStats struct {
	IsState             uint32 `idl:"name:IS_State" json:"is_state"`
	IsSendFailures      uint32 `idl:"name:IS_SendFailures" json:"is_send_failures"`
	IsReceiveFailures   uint32 `idl:"name:IS_ReceiveFailures" json:"is_receive_failures"`
	IsRequestsReceived  uint32 `idl:"name:IS_RequestsReceived" json:"is_requests_received"`
	IsRequestsDiscarded uint32 `idl:"name:IS_RequestsDiscarded" json:"is_requests_discarded"`
	IsRepliesReceived   uint32 `idl:"name:IS_RepliesReceived" json:"is_replies_received"`
	IsRepliesDiscarded  uint32 `idl:"name:IS_RepliesDiscarded" json:"is_replies_discarded"`
}

func (o *V6rInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *V6rInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IsState); err != nil {
		return err
	}
	if err := w.WriteData(o.IsSendFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.IsReceiveFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRequestsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRequestsDiscarded); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRepliesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRepliesDiscarded); err != nil {
		return err
	}
	return nil
}
func (o *V6rInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsState); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsSendFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsReceiveFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRequestsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRequestsDiscarded); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRepliesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRepliesDiscarded); err != nil {
		return err
	}
	return nil
}

// V6rMIBGetInputData structure represents DHCPV6R_MIB_GET_INPUT_DATA RPC structure.
type V6rMIBGetInputData struct {
	ImgidTypeID         uint32 `idl:"name:IMGID_TypeID" json:"imgid_type_id"`
	ImgidInterfaceIndex uint32 `idl:"name:IMGID_IfIndex" json:"imgid_interface_index"`
}

func (o *V6rMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *V6rMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgidTypeID); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgidInterfaceIndex); err != nil {
		return err
	}
	return nil
}
func (o *V6rMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgidTypeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgidInterfaceIndex); err != nil {
		return err
	}
	return nil
}

// V6rGlobalConfig structure represents DHCPV6R_GLOBAL_CONFIG RPC structure.
type V6rGlobalConfig struct {
	GCLoggingLevel     uint32 `idl:"name:GC_LoggingLevel" json:"gc_logging_level"`
	GCMaxRecvQueueSize uint32 `idl:"name:GC_MaxRecvQueueSize" json:"gc_max_recv_queue_size"`
	GCServerCount      uint32 `idl:"name:GC_ServerCount" json:"gc_server_count"`
}

func (o *V6rGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *V6rGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.GCLoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.GCMaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.GCServerCount); err != nil {
		return err
	}
	return nil
}
func (o *V6rGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCLoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCMaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCServerCount); err != nil {
		return err
	}
	return nil
}

// V6rInterfaceConfig structure represents DHCPV6R_IF_CONFIG RPC structure.
type V6rInterfaceConfig struct {
	ICState          uint32 `idl:"name:IC_State" json:"ic_state"`
	ICRelayMode      uint32 `idl:"name:IC_RelayMode" json:"ic_relay_mode"`
	ICMaxHopCount    uint32 `idl:"name:IC_MaxHopCount" json:"ic_max_hop_count"`
	ICMinElapsedTime uint32 `idl:"name:IC_MinElapsedTime" json:"ic_min_elapsed_time"`
}

func (o *V6rInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *V6rInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ICState); err != nil {
		return err
	}
	if err := w.WriteData(o.ICRelayMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICMaxHopCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ICMinElapsedTime); err != nil {
		return err
	}
	return nil
}
func (o *V6rInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICState); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICRelayMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICMaxHopCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICMinElapsedTime); err != nil {
		return err
	}
	return nil
}

// IpripMIBGetInputData structure represents IPRIP_MIB_GET_INPUT_DATA RPC structure.
type IpripMIBGetInputData struct {
	ImgidTypeID uint32                       `idl:"name:IMGID_TypeID" json:"imgid_type_id"`
	Field2      *IpripMIBGetInputData_Field2 `idl:"name:" json:""`
}

func (o *IpripMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgidTypeID); err != nil {
		return err
	}
	// FIXME unknown type
	return nil
}
func (o *IpripMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgidTypeID); err != nil {
		return err
	}
	// FIXME: unknown type
	return nil
}

type IpripMIBGetInputData_Field2 struct {
	ImgidInterfaceIndex uint32 `idl:"name:IMGID_IfIndex" json:"imgid_interface_index"`
	ImgidPeerAddress    uint32 `idl:"name:IMGID_PeerAddress" json:"imgid_peer_address"`
}

// IpripMIBGetOutputData structure represents IPRIP_MIB_GET_OUTPUT_DATA RPC structure.
type IpripMIBGetOutputData struct {
	ImgodTypeID uint32                        `idl:"name:IMGOD_TypeID" json:"imgod_type_id"`
	Field2      *IpripMIBGetOutputData_Field2 `idl:"name:" json:""`
	ImgodBuffer []byte                        `idl:"name:IMGOD_Buffer" json:"imgod_buffer"`
}

func (o *IpripMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ImgodTypeID); err != nil {
		return err
	}
	// FIXME unknown type
	for i1 := range o.ImgodBuffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.ImgodBuffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ImgodBuffer); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *IpripMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImgodTypeID); err != nil {
		return err
	}
	// FIXME: unknown type
	o.ImgodBuffer = make([]byte, 1)
	for i1 := range o.ImgodBuffer {
		i1 := i1
		if err := w.ReadData(&o.ImgodBuffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

type IpripMIBGetOutputData_Field2 struct {
	ImgodInterfaceIndex uint32 `idl:"name:IMGOD_IfIndex" json:"imgod_interface_index"`
	ImgodPeerAddress    uint32 `idl:"name:IMGOD_PeerAddress" json:"imgod_peer_address"`
}

// IpripGlobalStats structure represents IPRIP_GLOBAL_STATS RPC structure.
type IpripGlobalStats struct {
	GsSystemRouteChanges uint32 `idl:"name:GS_SystemRouteChanges" json:"gs_system_route_changes"`
	GsTotalResponsesSent uint32 `idl:"name:GS_TotalResponsesSent" json:"gs_total_responses_sent"`
}

func (o *IpripGlobalStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripGlobalStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpripGlobalStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpripGlobalConfig structure represents IPRIP_GLOBAL_CONFIG RPC structure.
type IpripGlobalConfig struct {
	GCLoggingLevel               uint32 `idl:"name:GC_LoggingLevel" json:"gc_logging_level"`
	GCMaxRecvQueueSize           uint32 `idl:"name:GC_MaxRecvQueueSize" json:"gc_max_recv_queue_size"`
	GCMaxSendQueueSize           uint32 `idl:"name:GC_MaxSendQueueSize" json:"gc_max_send_queue_size"`
	GCMinTriggeredUpdateInterval uint32 `idl:"name:GC_MinTriggeredUpdateInterval" json:"gc_min_triggered_update_interval"`
	GCPeerFilterMode             uint32 `idl:"name:GC_PeerFilterMode" json:"gc_peer_filter_mode"`
	GCPeerFilterCount            uint32 `idl:"name:GC_PeerFilterCount" json:"gc_peer_filter_count"`
}

func (o *IpripGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.GCLoggingLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.GCMaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.GCMaxSendQueueSize); err != nil {
		return err
	}
	if err := w.WriteData(o.GCMinTriggeredUpdateInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.GCPeerFilterMode); err != nil {
		return err
	}
	if err := w.WriteData(o.GCPeerFilterCount); err != nil {
		return err
	}
	return nil
}
func (o *IpripGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCLoggingLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCMaxRecvQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCMaxSendQueueSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCMinTriggeredUpdateInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCPeerFilterMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.GCPeerFilterCount); err != nil {
		return err
	}
	return nil
}

// IpripInterfaceStats structure represents IPRIP_IF_STATS RPC structure.
type IpripInterfaceStats struct {
	IsState                      uint32 `idl:"name:IS_State" json:"is_state"`
	IsSendFailures               uint32 `idl:"name:IS_SendFailures" json:"is_send_failures"`
	IsReceiveFailures            uint32 `idl:"name:IS_ReceiveFailures" json:"is_receive_failures"`
	IsRequestsSent               uint32 `idl:"name:IS_RequestsSent" json:"is_requests_sent"`
	IsRequestsReceived           uint32 `idl:"name:IS_RequestsReceived" json:"is_requests_received"`
	IsResponsesSent              uint32 `idl:"name:IS_ResponsesSent" json:"is_responses_sent"`
	IsResponsesReceived          uint32 `idl:"name:IS_ResponsesReceived" json:"is_responses_received"`
	IsBadResponsePacketsReceived uint32 `idl:"name:IS_BadResponsePacketsReceived" json:"is_bad_response_packets_received"`
	IsBadResponseEntriesReceived uint32 `idl:"name:IS_BadResponseEntriesReceived" json:"is_bad_response_entries_received"`
	IsTriggeredUpdatesSent       uint32 `idl:"name:IS_TriggeredUpdatesSent" json:"is_triggered_updates_sent"`
}

func (o *IpripInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IsState); err != nil {
		return err
	}
	if err := w.WriteData(o.IsSendFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.IsReceiveFailures); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRequestsSent); err != nil {
		return err
	}
	if err := w.WriteData(o.IsRequestsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsResponsesSent); err != nil {
		return err
	}
	if err := w.WriteData(o.IsResponsesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsBadResponsePacketsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsBadResponseEntriesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.IsTriggeredUpdatesSent); err != nil {
		return err
	}
	return nil
}
func (o *IpripInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsState); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsSendFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsReceiveFailures); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRequestsSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsRequestsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsResponsesSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsResponsesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsBadResponsePacketsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsBadResponseEntriesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsTriggeredUpdatesSent); err != nil {
		return err
	}
	return nil
}

// IpripInterfaceConfig structure represents IPRIP_IF_CONFIG RPC structure.
type IpripInterfaceConfig struct {
	ICState                   uint32 `idl:"name:IC_State" json:"ic_state"`
	ICMetric                  uint32 `idl:"name:IC_Metric" json:"ic_metric"`
	ICUpdateMode              uint32 `idl:"name:IC_UpdateMode" json:"ic_update_mode"`
	ICAcceptMode              uint32 `idl:"name:IC_AcceptMode" json:"ic_accept_mode"`
	ICAnnounceMode            uint32 `idl:"name:IC_AnnounceMode" json:"ic_announce_mode"`
	ICProtocolFlags           uint32 `idl:"name:IC_ProtocolFlags" json:"ic_protocol_flags"`
	ICRouteExpirationInterval uint32 `idl:"name:IC_RouteExpirationInterval" json:"ic_route_expiration_interval"`
	ICRouteRemovalInterval    uint32 `idl:"name:IC_RouteRemovalInterval" json:"ic_route_removal_interval"`
	ICFullUpdateInterval      uint32 `idl:"name:IC_FullUpdateInterval" json:"ic_full_update_interval"`
	ICAuthenticationType      uint32 `idl:"name:IC_AuthenticationType" json:"ic_authentication_type"`
	ICAuthenticationKey       []byte `idl:"name:IC_AuthenticationKey" json:"ic_authentication_key"`
	ICRouteTag                uint16 `idl:"name:IC_RouteTag" json:"ic_route_tag"`
	ICUnicastPeerMode         uint32 `idl:"name:IC_UnicastPeerMode" json:"ic_unicast_peer_mode"`
	ICAcceptFilterMode        uint32 `idl:"name:IC_AcceptFilterMode" json:"ic_accept_filter_mode"`
	ICAnnounceFilterMode      uint32 `idl:"name:IC_AnnounceFilterMode" json:"ic_announce_filter_mode"`
	ICUnicastPeerCount        uint32 `idl:"name:IC_UnicastPeerCount" json:"ic_unicast_peer_count"`
	ICAcceptFilterCount       uint32 `idl:"name:IC_AcceptFilterCount" json:"ic_accept_filter_count"`
	ICAnnounceFilterCount     uint32 `idl:"name:IC_AnnounceFilterCount" json:"ic_announce_filter_count"`
}

func (o *IpripInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ICState); err != nil {
		return err
	}
	if err := w.WriteData(o.ICMetric); err != nil {
		return err
	}
	if err := w.WriteData(o.ICUpdateMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICAcceptMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICAnnounceMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICProtocolFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.ICRouteExpirationInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.ICRouteRemovalInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.ICFullUpdateInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.ICAuthenticationType); err != nil {
		return err
	}
	for i1 := range o.ICAuthenticationKey {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.ICAuthenticationKey[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ICAuthenticationKey); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ICRouteTag); err != nil {
		return err
	}
	if err := w.WriteData(o.ICUnicastPeerMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICAcceptFilterMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICAnnounceFilterMode); err != nil {
		return err
	}
	if err := w.WriteData(o.ICUnicastPeerCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ICAcceptFilterCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ICAnnounceFilterCount); err != nil {
		return err
	}
	return nil
}
func (o *IpripInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICState); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICMetric); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICUpdateMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICAcceptMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICAnnounceMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICProtocolFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICRouteExpirationInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICRouteRemovalInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICFullUpdateInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICAuthenticationType); err != nil {
		return err
	}
	o.ICAuthenticationKey = make([]byte, 16)
	for i1 := range o.ICAuthenticationKey {
		i1 := i1
		if err := w.ReadData(&o.ICAuthenticationKey[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.ICRouteTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICUnicastPeerMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICAcceptFilterMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICAnnounceFilterMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICUnicastPeerCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICAcceptFilterCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ICAnnounceFilterCount); err != nil {
		return err
	}
	return nil
}

// IpripRouteFilter structure represents IPRIP_ROUTE_FILTER RPC structure.
type IpripRouteFilter struct {
	RfLoAddress uint32 `idl:"name:RF_LoAddress" json:"rf_lo_address"`
	RfHiAddress uint32 `idl:"name:RF_HiAddress" json:"rf_hi_address"`
}

func (o *IpripRouteFilter) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripRouteFilter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpripRouteFilter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IpripInterfaceBinding structure represents IPRIP_IF_BINDING RPC structure.
type IpripInterfaceBinding struct {
	IbState     uint32 `idl:"name:IB_State" json:"ib_state"`
	IbAddrCount uint32 `idl:"name:IB_AddrCount" json:"ib_addr_count"`
}

func (o *IpripInterfaceBinding) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripInterfaceBinding) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IbState); err != nil {
		return err
	}
	if err := w.WriteData(o.IbAddrCount); err != nil {
		return err
	}
	return nil
}
func (o *IpripInterfaceBinding) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IbState); err != nil {
		return err
	}
	if err := w.ReadData(&o.IbAddrCount); err != nil {
		return err
	}
	return nil
}

// IpripIPAddress structure represents IPRIP_IP_ADDRESS RPC structure.
type IpripIPAddress struct {
	IaAddress uint32 `idl:"name:IA_Address" json:"ia_address"`
	IaNetmask uint32 `idl:"name:IA_Netmask" json:"ia_netmask"`
}

func (o *IpripIPAddress) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripIPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IaAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.IaNetmask); err != nil {
		return err
	}
	return nil
}
func (o *IpripIPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IaAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.IaNetmask); err != nil {
		return err
	}
	return nil
}

// IpripPeerStats structure represents IPRIP_PEER_STATS RPC structure.
type IpripPeerStats struct {
	PsLastPeerRouteTag           uint32 `idl:"name:PS_LastPeerRouteTag" json:"ps_last_peer_route_tag"`
	PsLastPeerUpdateTickCount    uint32 `idl:"name:PS_LastPeerUpdateTickCount" json:"ps_last_peer_update_tick_count"`
	PsLastPeerUpdateVersion      uint32 `idl:"name:PS_LastPeerUpdateVersion" json:"ps_last_peer_update_version"`
	PsBadResponsePacketsFromPeer uint32 `idl:"name:PS_BadResponsePacketsFromPeer" json:"ps_bad_response_packets_from_peer"`
	PsBadResponseEntriesFromPeer uint32 `idl:"name:PS_BadResponseEntriesFromPeer" json:"ps_bad_response_entries_from_peer"`
}

func (o *IpripPeerStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IpripPeerStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IpripPeerStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBGetInputData structure represents IGMP_MIB_GET_INPUT_DATA RPC structure.
type IgmpMIBGetInputData struct {
	TypeID         uint32 `idl:"name:TypeId" json:"type_id"`
	Flags          uint16 `idl:"name:Flags" json:"flags"`
	Signature      uint16 `idl:"name:Signature" json:"signature"`
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	RASClientAddr  uint32 `idl:"name:RasClientAddr" json:"ras_client_addr"`
	GroupAddr      uint32 `idl:"name:GroupAddr" json:"group_addr"`
	Count          uint32 `idl:"name:Count" json:"count"`
}

func (o *IgmpMIBGetInputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGetInputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGetInputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBGetOutputData structure represents IGMP_MIB_GET_OUTPUT_DATA RPC structure.
type IgmpMIBGetOutputData struct {
	TypeID uint32 `idl:"name:TypeId" json:"type_id"`
	Flags  uint32 `idl:"name:Flags" json:"flags"`
	Count  uint32 `idl:"name:Count" json:"count"`
	Buffer []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *IgmpMIBGetOutputData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGetOutputData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGetOutputData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBGlobalConfig structure represents IGMP_MIB_GLOBAL_CONFIG RPC structure.
type IgmpMIBGlobalConfig struct {
	Version        uint32 `idl:"name:Version" json:"version"`
	LoggingLevel   uint32 `idl:"name:LoggingLevel" json:"logging_level"`
	RASClientStats uint32 `idl:"name:RasClientStats" json:"ras_client_stats"`
}

func (o *IgmpMIBGlobalConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGlobalConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGlobalConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBGlobalStats structure represents IGMP_MIB_GLOBAL_STATS RPC structure.
type IgmpMIBGlobalStats struct {
	CurrentGroupMemberships uint32 `idl:"name:CurrentGroupMemberships" json:"current_group_memberships"`
	GroupMembershipsAdded   uint32 `idl:"name:GroupMembershipsAdded" json:"group_memberships_added"`
}

func (o *IgmpMIBGlobalStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGlobalStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGlobalStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBInterfaceBinding structure represents IGMP_MIB_IF_BINDING RPC structure.
type IgmpMIBInterfaceBinding struct {
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	InterfaceType  uint32 `idl:"name:IfType" json:"interface_type"`
	State          uint32 `idl:"name:State" json:"state"`
	AddrCount      uint32 `idl:"name:AddrCount" json:"addr_count"`
}

func (o *IgmpMIBInterfaceBinding) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBInterfaceBinding) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBInterfaceBinding) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBInterfaceConfig structure represents IGMP_MIB_IF_CONFIG RPC structure.
type IgmpMIBInterfaceConfig struct {
	Version                     uint32 `idl:"name:Version" json:"version"`
	InterfaceIndex              uint32 `idl:"name:IfIndex" json:"interface_index"`
	IPAddr                      uint32 `idl:"name:IpAddr" json:"ip_addr"`
	InterfaceType               uint32 `idl:"name:IfType" json:"interface_type"`
	Flags                       uint32 `idl:"name:Flags" json:"flags"`
	IgmpProtocolType            uint32 `idl:"name:IgmpProtocolType" json:"igmp_protocol_type"`
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

func (o *IgmpMIBInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.IgmpProtocolType); err != nil {
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
func (o *IgmpMIBInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.IgmpProtocolType); err != nil {
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

// IgmpMIBInterfaceGroupsList structure represents IGMP_MIB_IF_GROUPS_LIST RPC structure.
type IgmpMIBInterfaceGroupsList struct {
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	IPAddr         uint32 `idl:"name:IpAddr" json:"ip_addr"`
	InterfaceType  uint32 `idl:"name:IfType" json:"interface_type"`
	GroupsLength   uint32 `idl:"name:NumGroups" json:"groups_length"`
	Buffer         []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *IgmpMIBInterfaceGroupsList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBInterfaceGroupsList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBInterfaceGroupsList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBGroupInfo structure represents IGMP_MIB_GROUP_INFO RPC structure.
type IgmpMIBGroupInfo struct {
	Field1                *IgmpMIBGroupInfo_Field1 `idl:"name:" json:""`
	IPAddr                uint32                   `idl:"name:IpAddr" json:"ip_addr"`
	GroupUpTime           uint32                   `idl:"name:GroupUpTime" json:"group_up_time"`
	GroupExpiryTime       uint32                   `idl:"name:GroupExpiryTime" json:"group_expiry_time"`
	LastReporter          uint32                   `idl:"name:LastReporter" json:"last_reporter"`
	V1HostPresentTimeLeft uint32                   `idl:"name:V1HostPresentTimeLeft" json:"v1_host_present_time_left"`
	Flags                 uint32                   `idl:"name:Flags" json:"flags"`
}

func (o *IgmpMIBGroupInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGroupInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGroupInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type IgmpMIBGroupInfo_Field1 struct {
	InterfaceIndex uint32 `idl:"name:IfIndex" json:"interface_index"`
	GroupAddr      uint32 `idl:"name:GroupAddr" json:"group_addr"`
}

// IgmpMIBInterfaceStats structure represents IGMP_MIB_IF_STATS RPC structure.
type IgmpMIBInterfaceStats struct {
	InterfaceIndex            uint32 `idl:"name:IfIndex" json:"interface_index"`
	IPAddr                    uint32 `idl:"name:IpAddr" json:"ip_addr"`
	InterfaceType             uint32 `idl:"name:IfType" json:"interface_type"`
	State                     uint8  `idl:"name:State" json:"state"`
	QuerierState              uint8  `idl:"name:QuerierState" json:"querier_state"`
	IgmpProtocolType          uint32 `idl:"name:IgmpProtocolType" json:"igmp_protocol_type"`
	QuerierIPAddr             uint32 `idl:"name:QuerierIpAddr" json:"querier_ip_addr"`
	ProxyInterfaceIndex       uint32 `idl:"name:ProxyIfIndex" json:"proxy_interface_index"`
	QuerierPresentTimeLeft    uint32 `idl:"name:QuerierPresentTimeLeft" json:"querier_present_time_left"`
	LastQuerierChangeTime     uint32 `idl:"name:LastQuerierChangeTime" json:"last_querier_change_time"`
	V1QuerierPresentTimeLeft  uint32 `idl:"name:V1QuerierPresentTimeLeft" json:"v1_querier_present_time_left"`
	Uptime                    uint32 `idl:"name:Uptime" json:"uptime"`
	TotalIgmpPacketsReceived  uint32 `idl:"name:TotalIgmpPacketsReceived" json:"total_igmp_packets_received"`
	TotalIgmpPacketsForRouter uint32 `idl:"name:TotalIgmpPacketsForRouter" json:"total_igmp_packets_for_router"`
	GeneralQueriesReceived    uint32 `idl:"name:GeneralQueriesReceived" json:"general_queries_received"`
	WrongVersionQueries       uint32 `idl:"name:WrongVersionQueries" json:"wrong_version_queries"`
	JoinsReceived             uint32 `idl:"name:JoinsReceived" json:"joins_received"`
	LeavesReceived            uint32 `idl:"name:LeavesReceived" json:"leaves_received"`
	CurrentGroupMemberships   uint32 `idl:"name:CurrentGroupMemberships" json:"current_group_memberships"`
	GroupMembershipsAdded     uint32 `idl:"name:GroupMembershipsAdded" json:"group_memberships_added"`
	WrongChecksumPackets      uint32 `idl:"name:WrongChecksumPackets" json:"wrong_checksum_packets"`
	ShortPacketsReceived      uint32 `idl:"name:ShortPacketsReceived" json:"short_packets_received"`
	LongPacketsReceived       uint32 `idl:"name:LongPacketsReceived" json:"long_packets_received"`
	PacketsWithoutRtrAlert    uint32 `idl:"name:PacketsWithoutRtrAlert" json:"packets_without_rtr_alert"`
}

func (o *IgmpMIBInterfaceStats) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBInterfaceStats) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.IgmpProtocolType); err != nil {
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
	if err := w.WriteData(o.Uptime); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalIgmpPacketsReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalIgmpPacketsForRouter); err != nil {
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
	if err := w.WriteData(o.PacketsWithoutRtrAlert); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBInterfaceStats) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.IgmpProtocolType); err != nil {
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
	if err := w.ReadData(&o.Uptime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalIgmpPacketsReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalIgmpPacketsForRouter); err != nil {
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
	if err := w.ReadData(&o.PacketsWithoutRtrAlert); err != nil {
		return err
	}
	return nil
}

// IgmpMIBGroupInterfacesList structure represents IGMP_MIB_GROUP_IFS_LIST RPC structure.
type IgmpMIBGroupInterfacesList struct {
	GroupAddr        uint32 `idl:"name:GroupAddr" json:"group_addr"`
	InterfacesLength uint32 `idl:"name:NumInterfaces" json:"interfaces_length"`
	Buffer           []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *IgmpMIBGroupInterfacesList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGroupInterfacesList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGroupInterfacesList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBGroupSourceInfoV3 structure represents IGMP_MIB_GROUP_SOURCE_INFO_V3 RPC structure.
type IgmpMIBGroupSourceInfoV3 struct {
	Source           uint32 `idl:"name:Source" json:"source"`
	SourceExpiryTime uint32 `idl:"name:SourceExpiryTime" json:"source_expiry_time"`
	SourceUpTime     uint32 `idl:"name:SourceUpTime" json:"source_up_time"`
	Flags            uint32 `idl:"name:Flags" json:"flags"`
}

func (o *IgmpMIBGroupSourceInfoV3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGroupSourceInfoV3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGroupSourceInfoV3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IgmpMIBGroupInfoV3 structure represents IGMP_MIB_GROUP_INFO_V3 RPC structure.
type IgmpMIBGroupInfoV3 struct {
	Field1                *IgmpMIBGroupInfoV3_Field1 `idl:"name:" json:""`
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

func (o *IgmpMIBGroupInfoV3) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IgmpMIBGroupInfoV3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IgmpMIBGroupInfoV3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type IgmpMIBGroupInfoV3_Field1 struct {
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
	InSnmpID      []uint32 `idl:"name:in_snmp_id" json:"in_snmp_id"`
	ObjectID      []uint32 `idl:"name:obj_id" json:"object_id"`
	AttributeID   uint32   `idl:"name:attr_id" json:"attribute_id"`
	InstanceID    []uint32 `idl:"name:inst_id" json:"instance_id"`
	NextSnmpID    []uint32 `idl:"name:next_snmp_id" json:"next_snmp_id"`
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
	for i1 := range o.InSnmpID {
		i1 := i1
		if uint64(i1) >= 44 {
			break
		}
		if err := w.WriteData(o.InSnmpID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.InSnmpID); uint64(i1) < 44; i1++ {
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
	for i1 := range o.NextSnmpID {
		i1 := i1
		if uint64(i1) >= 44 {
			break
		}
		if err := w.WriteData(o.NextSnmpID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NextSnmpID); uint64(i1) < 44; i1++ {
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
	o.InSnmpID = make([]uint32, 44)
	for i1 := range o.InSnmpID {
		i1 := i1
		if err := w.ReadData(&o.InSnmpID[i1]); err != nil {
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
	o.NextSnmpID = make([]uint32, 44)
	for i1 := range o.NextSnmpID {
		i1 := i1
		if err := w.ReadData(&o.NextSnmpID[i1]); err != nil {
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
	LoggingLevel uint32              `idl:"name:LoggingLevel" json:"logging_level"`
	Flags        uint32              `idl:"name:Flags" json:"flags"`
	Header       *RtrInfoBlockHeader `idl:"name:Header" json:"header"`
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
		if err := (&RtrInfoBlockHeader{}).MarshalNDR(ctx, w); err != nil {
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
		o.Header = &RtrInfoBlockHeader{}
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
	Index  uint32              `idl:"name:Index" json:"index"`
	Flags  uint32              `idl:"name:Flags" json:"flags"`
	Header *RtrInfoBlockHeader `idl:"name:Header" json:"header"`
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
		if err := (&RtrInfoBlockHeader{}).MarshalNDR(ctx, w); err != nil {
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
		o.Header = &RtrInfoBlockHeader{}
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

// RipGlobalInfo structure represents RIP_GLOBAL_INFO RPC structure.
type RipGlobalInfo struct {
	EventLogMask uint32 `idl:"name:EventLogMask" json:"event_log_mask"`
}

func (o *RipGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RipGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogMask); err != nil {
		return err
	}
	return nil
}

// RipRouteFilterInfo structure represents RIP_ROUTE_FILTER_INFO RPC structure.
type RipRouteFilterInfo struct {
	Network []byte `idl:"name:Network" json:"network"`
	Mask    []byte `idl:"name:Mask" json:"mask"`
}

func (o *RipRouteFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipRouteFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RipRouteFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RipInterfaceFilters structure represents RIP_IF_FILTERS RPC structure.
type RipInterfaceFilters struct {
	SupplyFilterAction uint32                `idl:"name:SupplyFilterAction" json:"supply_filter_action"`
	SupplyFilterCount  uint32                `idl:"name:SupplyFilterCount" json:"supply_filter_count"`
	ListenFilterAction uint32                `idl:"name:ListenFilterAction" json:"listen_filter_action"`
	ListenFilterCount  uint32                `idl:"name:ListenFilterCount" json:"listen_filter_count"`
	RouteFilter        []*RipRouteFilterInfo `idl:"name:RouteFilter" json:"route_filter"`
}

func (o *RipInterfaceFilters) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipInterfaceFilters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
			if err := (&RipRouteFilterInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.RouteFilter); uint64(i1) < 1; i1++ {
		if err := (&RipRouteFilterInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *RipInterfaceFilters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	o.RouteFilter = make([]*RipRouteFilterInfo, 1)
	for i1 := range o.RouteFilter {
		i1 := i1
		if o.RouteFilter[i1] == nil {
			o.RouteFilter[i1] = &RipRouteFilterInfo{}
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

// RipInterfaceConfig structure represents RIP_IF_CONFIG RPC structure.
type RipInterfaceConfig struct {
	RipInterfaceInfo    *RipInterfaceInfo    `idl:"name:RipIfInfo" json:"rip_interface_info"`
	RipInterfaceFilters *RipInterfaceFilters `idl:"name:RipIfFilters" json:"rip_interface_filters"`
}

func (o *RipInterfaceConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RipInterfaceConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.RipInterfaceInfo != nil {
		if err := o.RipInterfaceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RipInterfaceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RipInterfaceFilters != nil {
		if err := o.RipInterfaceFilters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RipInterfaceFilters{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RipInterfaceConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.RipInterfaceInfo == nil {
		o.RipInterfaceInfo = &RipInterfaceInfo{}
	}
	if err := o.RipInterfaceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RipInterfaceFilters == nil {
		o.RipInterfaceFilters = &RipInterfaceFilters{}
	}
	if err := o.RipInterfaceFilters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SapGlobalInfo structure represents SAP_GLOBAL_INFO RPC structure.
type SapGlobalInfo struct {
	EventLogMask uint32 `idl:"name:EventLogMask" json:"event_log_mask"`
}

func (o *SapGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SapGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SapGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogMask); err != nil {
		return err
	}
	return nil
}

// OspfRouteFilter structure represents OSPF_ROUTE_FILTER RPC structure.
type OspfRouteFilter struct {
	Address uint32 `idl:"name:dwAddress" json:"address"`
	Mask    uint32 `idl:"name:dwMask" json:"mask"`
}

func (o *OspfRouteFilter) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfRouteFilter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *OspfRouteFilter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// OspfFilterAction type represents OSPF_FILTER_ACTION RPC enumeration.
type OspfFilterAction uint16

var (
	OspfFilterActionDrop   OspfFilterAction = 0
	OspfFilterActionAccept OspfFilterAction = 1
)

func (o OspfFilterAction) String() string {
	switch o {
	case OspfFilterActionDrop:
		return "OspfFilterActionDrop"
	case OspfFilterActionAccept:
		return "OspfFilterActionAccept"
	}
	return "Invalid"
}

// OspfRouteFilterInfo structure represents OSPF_ROUTE_FILTER_INFO RPC structure.
type OspfRouteFilterInfo struct {
	Type             uint32             `idl:"name:type" json:"type"`
	OfaActionOnMatch OspfFilterAction   `idl:"name:ofaActionOnMatch" json:"ofa_action_on_match"`
	FiltersLength    uint32             `idl:"name:dwNumFilters" json:"filters_length"`
	Filters          []*OspfRouteFilter `idl:"name:pFilters" json:"filters"`
}

func (o *OspfRouteFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfRouteFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.OfaActionOnMatch)); err != nil {
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
			if err := (&OspfRouteFilter{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Filters); uint64(i1) < 1; i1++ {
		if err := (&OspfRouteFilter{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OspfRouteFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.OfaActionOnMatch)); err != nil {
		return err
	}
	if err := w.ReadData(&o.FiltersLength); err != nil {
		return err
	}
	o.Filters = make([]*OspfRouteFilter, 1)
	for i1 := range o.Filters {
		i1 := i1
		if o.Filters[i1] == nil {
			o.Filters[i1] = &OspfRouteFilter{}
		}
		if err := o.Filters[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OspfProtoFilterInfo structure represents OSPF_PROTO_FILTER_INFO RPC structure.
type OspfProtoFilterInfo struct {
	Type             uint32           `idl:"name:type" json:"type"`
	OfaActionOnMatch OspfFilterAction `idl:"name:ofaActionOnMatch" json:"ofa_action_on_match"`
	ProtoIDsLength   uint32           `idl:"name:dwNumProtoIds" json:"proto_ids_length"`
	ProtoID          []uint32         `idl:"name:pdwProtoId" json:"proto_id"`
}

func (o *OspfProtoFilterInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfProtoFilterInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.OfaActionOnMatch)); err != nil {
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
func (o *OspfProtoFilterInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.OfaActionOnMatch)); err != nil {
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

// OspfGlobalParam structure represents OSPF_GLOBAL_PARAM RPC structure.
type OspfGlobalParam struct {
	Type      uint32 `idl:"name:type" json:"type"`
	Create    uint32 `idl:"name:create" json:"create"`
	Enable    uint32 `idl:"name:enable" json:"enable"`
	RouterID  uint32 `idl:"name:routerId" json:"router_id"`
	AsBrdrRtr uint32 `idl:"name:ASBrdrRtr" json:"as_brdr_rtr"`
	LogLevel  uint32 `idl:"name:logLevel" json:"log_level"`
}

func (o *OspfGlobalParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfGlobalParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.AsBrdrRtr); err != nil {
		return err
	}
	if err := w.WriteData(o.LogLevel); err != nil {
		return err
	}
	return nil
}
func (o *OspfGlobalParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.AsBrdrRtr); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogLevel); err != nil {
		return err
	}
	return nil
}

// OspfAreaParam structure represents OSPF_AREA_PARAM RPC structure.
type OspfAreaParam struct {
	Type           uint32 `idl:"name:type" json:"type"`
	Create         uint32 `idl:"name:create" json:"create"`
	Enable         uint32 `idl:"name:enable" json:"enable"`
	AreaID         uint32 `idl:"name:areaId" json:"area_id"`
	AuthType       uint32 `idl:"name:authType" json:"auth_type"`
	ImportAsExtern uint32 `idl:"name:importASExtern" json:"import_as_extern"`
	StubMetric     uint32 `idl:"name:stubMetric" json:"stub_metric"`
	ImportSumAdv   uint32 `idl:"name:importSumAdv" json:"import_sum_adv"`
}

func (o *OspfAreaParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfAreaParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *OspfAreaParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// OspfAreaRangeParam structure represents OSPF_AREA_RANGE_PARAM RPC structure.
type OspfAreaRangeParam struct {
	Type      uint32 `idl:"name:type" json:"type"`
	Create    uint32 `idl:"name:create" json:"create"`
	Enable    uint32 `idl:"name:enable" json:"enable"`
	AreaID    uint32 `idl:"name:areaId" json:"area_id"`
	RangeNet  uint32 `idl:"name:rangeNet" json:"range_net"`
	RangeMask uint32 `idl:"name:rangeMask" json:"range_mask"`
}

func (o *OspfAreaRangeParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfAreaRangeParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *OspfAreaRangeParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// OspfVirtInterfaceParam structure represents OSPF_VIRT_INTERFACE_PARAM RPC structure.
type OspfVirtInterfaceParam struct {
	Type                 uint32 `idl:"name:type" json:"type"`
	Create               uint32 `idl:"name:create" json:"create"`
	Enable               uint32 `idl:"name:enable" json:"enable"`
	TransitAreaID        uint32 `idl:"name:transitAreaId" json:"transit_area_id"`
	VirtNeighborRouterID uint32 `idl:"name:virtNeighborRouterId" json:"virt_neighbor_router_id"`
	TransitDelay         uint32 `idl:"name:transitDelay" json:"transit_delay"`
	RetransInterval      uint32 `idl:"name:retransInterval" json:"retrans_interval"`
	HelloInterval        uint32 `idl:"name:helloInterval" json:"hello_interval"`
	DeadInterval         uint32 `idl:"name:deadInterval" json:"dead_interval"`
	Password             []byte `idl:"name:password" json:"password"`
}

func (o *OspfVirtInterfaceParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfVirtInterfaceParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.VirtNeighborRouterID); err != nil {
		return err
	}
	if err := w.WriteData(o.TransitDelay); err != nil {
		return err
	}
	if err := w.WriteData(o.RetransInterval); err != nil {
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
func (o *OspfVirtInterfaceParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.VirtNeighborRouterID); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransitDelay); err != nil {
		return err
	}
	if err := w.ReadData(&o.RetransInterval); err != nil {
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

// OspfInterfaceParam structure represents OSPF_INTERFACE_PARAM RPC structure.
type OspfInterfaceParam struct {
	Type                uint32 `idl:"name:type" json:"type"`
	Create              uint32 `idl:"name:create" json:"create"`
	Enable              uint32 `idl:"name:enable" json:"enable"`
	InterfaceIPAddr     uint32 `idl:"name:intfIpAddr" json:"interface_ip_addr"`
	InterfaceSubnetMask uint32 `idl:"name:intfSubnetMask" json:"interface_subnet_mask"`
	AreaID              uint32 `idl:"name:areaId" json:"area_id"`
	InterfaceType       uint32 `idl:"name:intfType" json:"interface_type"`
	RouterPriority      uint32 `idl:"name:routerPriority" json:"router_priority"`
	TransitDelay        uint32 `idl:"name:transitDelay" json:"transit_delay"`
	RetransInterval     uint32 `idl:"name:retransInterval" json:"retrans_interval"`
	HelloInterval       uint32 `idl:"name:helloInterval" json:"hello_interval"`
	DeadInterval        uint32 `idl:"name:deadInterval" json:"dead_interval"`
	PollInterval        uint32 `idl:"name:pollInterval" json:"poll_interval"`
	MetricCost          uint32 `idl:"name:metricCost" json:"metric_cost"`
	Password            []byte `idl:"name:password" json:"password"`
	MtuSize             uint32 `idl:"name:mtuSize" json:"mtu_size"`
}

func (o *OspfInterfaceParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfInterfaceParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RetransInterval); err != nil {
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
	if err := w.WriteData(o.MtuSize); err != nil {
		return err
	}
	return nil
}
func (o *OspfInterfaceParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RetransInterval); err != nil {
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
	if err := w.ReadData(&o.MtuSize); err != nil {
		return err
	}
	return nil
}

// OspfNbmaNeighborParam structure represents OSPF_NBMA_NEIGHBOR_PARAM RPC structure.
type OspfNbmaNeighborParam struct {
	Type             uint32 `idl:"name:type" json:"type"`
	Create           uint32 `idl:"name:create" json:"create"`
	Enable           uint32 `idl:"name:enable" json:"enable"`
	NeighborIPAddr   uint32 `idl:"name:neighborIpAddr" json:"neighbor_ip_addr"`
	InterfaceIPAddr  uint32 `idl:"name:intfIpAddr" json:"interface_ip_addr"`
	NeighborPriority uint32 `idl:"name:neighborPriority" json:"neighbor_priority"`
}

func (o *OspfNbmaNeighborParam) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OspfNbmaNeighborParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *OspfNbmaNeighborParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Rasdevicetype type represents RASDEVICETYPE RPC enumeration.
type Rasdevicetype uint32

var (
	RasdevicetypeRdtModem       Rasdevicetype = 0
	RasdevicetypeRdtX25         Rasdevicetype = 1
	RasdevicetypeRdtISDN        Rasdevicetype = 2
	RasdevicetypeRdtSerial      Rasdevicetype = 3
	RasdevicetypeRdtFrameRelay  Rasdevicetype = 4
	RasdevicetypeRdtAtm         Rasdevicetype = 5
	RasdevicetypeRdtSonet       Rasdevicetype = 6
	RasdevicetypeRdtSw56        Rasdevicetype = 7
	RasdevicetypeRdtTunnelPPTP  Rasdevicetype = 8
	RasdevicetypeRdtTunnelL2tp  Rasdevicetype = 9
	RasdevicetypeRdtIrda        Rasdevicetype = 10
	RasdevicetypeRdtParallel    Rasdevicetype = 11
	RasdevicetypeRdtOther       Rasdevicetype = 12
	RasdevicetypeRdtPpPoE       Rasdevicetype = 13
	RasdevicetypeRdtTunnelSstp  Rasdevicetype = 14
	RasdevicetypeRdtTunnelIKEv2 Rasdevicetype = 15
	RasdevicetypeRdtTunnel      Rasdevicetype = 65536
	RasdevicetypeRdtDirect      Rasdevicetype = 131072
	RasdevicetypeRdtNullModem   Rasdevicetype = 262144
	RasdevicetypeRdtBroadband   Rasdevicetype = 524288
)

func (o Rasdevicetype) String() string {
	switch o {
	case RasdevicetypeRdtModem:
		return "RasdevicetypeRdtModem"
	case RasdevicetypeRdtX25:
		return "RasdevicetypeRdtX25"
	case RasdevicetypeRdtISDN:
		return "RasdevicetypeRdtISDN"
	case RasdevicetypeRdtSerial:
		return "RasdevicetypeRdtSerial"
	case RasdevicetypeRdtFrameRelay:
		return "RasdevicetypeRdtFrameRelay"
	case RasdevicetypeRdtAtm:
		return "RasdevicetypeRdtAtm"
	case RasdevicetypeRdtSonet:
		return "RasdevicetypeRdtSonet"
	case RasdevicetypeRdtSw56:
		return "RasdevicetypeRdtSw56"
	case RasdevicetypeRdtTunnelPPTP:
		return "RasdevicetypeRdtTunnelPPTP"
	case RasdevicetypeRdtTunnelL2tp:
		return "RasdevicetypeRdtTunnelL2tp"
	case RasdevicetypeRdtIrda:
		return "RasdevicetypeRdtIrda"
	case RasdevicetypeRdtParallel:
		return "RasdevicetypeRdtParallel"
	case RasdevicetypeRdtOther:
		return "RasdevicetypeRdtOther"
	case RasdevicetypeRdtPpPoE:
		return "RasdevicetypeRdtPpPoE"
	case RasdevicetypeRdtTunnelSstp:
		return "RasdevicetypeRdtTunnelSstp"
	case RasdevicetypeRdtTunnelIKEv2:
		return "RasdevicetypeRdtTunnelIKEv2"
	case RasdevicetypeRdtTunnel:
		return "RasdevicetypeRdtTunnel"
	case RasdevicetypeRdtDirect:
		return "RasdevicetypeRdtDirect"
	case RasdevicetypeRdtNullModem:
		return "RasdevicetypeRdtNullModem"
	case RasdevicetypeRdtBroadband:
		return "RasdevicetypeRdtBroadband"
	}
	return "Invalid"
}

// RasmanStatus type represents RASMAN_STATUS RPC enumeration.
type RasmanStatus uint16

var (
	RasmanStatusOpen        RasmanStatus = 0
	RasmanStatusClosed      RasmanStatus = 1
	RasmanStatusUnavailable RasmanStatus = 2
	RasmanStatusRemoved     RasmanStatus = 3
)

func (o RasmanStatus) String() string {
	switch o {
	case RasmanStatusOpen:
		return "RasmanStatusOpen"
	case RasmanStatusClosed:
		return "RasmanStatusClosed"
	case RasmanStatusUnavailable:
		return "RasmanStatusUnavailable"
	case RasmanStatusRemoved:
		return "RasmanStatusRemoved"
	}
	return "Invalid"
}

// RequestTypes type represents ReqTypes RPC enumeration.
type RequestTypes uint16

var (
	RequestTypesReqtypePortenum             RequestTypes = 21
	RequestTypesReqtypeGetinfo              RequestTypes = 22
	RequestTypesReqtypeGetdevconfig         RequestTypes = 73
	RequestTypesReqtypeSetdeviceconfiginfo  RequestTypes = 94
	RequestTypesReqtypeGetdeviceconfiginfo  RequestTypes = 95
	RequestTypesReqtypeGetcalledid          RequestTypes = 105
	RequestTypesReqtypeSetcalledid          RequestTypes = 106
	RequestTypesReqtypeGetndiswandrivercaps RequestTypes = 111
)

func (o RequestTypes) String() string {
	switch o {
	case RequestTypesReqtypePortenum:
		return "RequestTypesReqtypePortenum"
	case RequestTypesReqtypeGetinfo:
		return "RequestTypesReqtypeGetinfo"
	case RequestTypesReqtypeGetdevconfig:
		return "RequestTypesReqtypeGetdevconfig"
	case RequestTypesReqtypeSetdeviceconfiginfo:
		return "RequestTypesReqtypeSetdeviceconfiginfo"
	case RequestTypesReqtypeGetdeviceconfiginfo:
		return "RequestTypesReqtypeGetdeviceconfiginfo"
	case RequestTypesReqtypeGetcalledid:
		return "RequestTypesReqtypeGetcalledid"
	case RequestTypesReqtypeSetcalledid:
		return "RequestTypesReqtypeSetcalledid"
	case RequestTypesReqtypeGetndiswandrivercaps:
		return "RequestTypesReqtypeGetndiswandrivercaps"
	}
	return "Invalid"
}

// RasmanState type represents RASMAN_STATE RPC enumeration.
type RasmanState uint16

var (
	RasmanStateConnecting      RasmanState = 0
	RasmanStateListening       RasmanState = 1
	RasmanStateConnected       RasmanState = 2
	RasmanStateDisconnecting   RasmanState = 3
	RasmanStateDisconnected    RasmanState = 4
	RasmanStateListencompleted RasmanState = 5
)

func (o RasmanState) String() string {
	switch o {
	case RasmanStateConnecting:
		return "RasmanStateConnecting"
	case RasmanStateListening:
		return "RasmanStateListening"
	case RasmanStateConnected:
		return "RasmanStateConnected"
	case RasmanStateDisconnecting:
		return "RasmanStateDisconnecting"
	case RasmanStateDisconnected:
		return "RasmanStateDisconnected"
	case RasmanStateListencompleted:
		return "RasmanStateListencompleted"
	}
	return "Invalid"
}

// RasmanDisconnectType type represents RASMAN_DISCONNECT_TYPE RPC enumeration.
type RasmanDisconnectType uint16

var (
	RasmanDisconnectTypeUserRequested       RasmanDisconnectType = 0
	RasmanDisconnectTypeRemoteDisconnection RasmanDisconnectType = 1
	RasmanDisconnectTypeHardwareFailure     RasmanDisconnectType = 2
	RasmanDisconnectTypeNotDisconnected     RasmanDisconnectType = 3
)

func (o RasmanDisconnectType) String() string {
	switch o {
	case RasmanDisconnectTypeUserRequested:
		return "RasmanDisconnectTypeUserRequested"
	case RasmanDisconnectTypeRemoteDisconnection:
		return "RasmanDisconnectTypeRemoteDisconnection"
	case RasmanDisconnectTypeHardwareFailure:
		return "RasmanDisconnectTypeHardwareFailure"
	case RasmanDisconnectTypeNotDisconnected:
		return "RasmanDisconnectTypeNotDisconnected"
	}
	return "Invalid"
}

// RasmanUsage type represents RASMAN_USAGE RPC enumeration.
type RasmanUsage uint16

var (
	RasmanUsageCallNone           RasmanUsage = 0
	RasmanUsageCallIn             RasmanUsage = 1
	RasmanUsageCallOut            RasmanUsage = 2
	RasmanUsageCallRouter         RasmanUsage = 4
	RasmanUsageCallLogon          RasmanUsage = 8
	RasmanUsageCallOutOnly        RasmanUsage = 16
	RasmanUsageCallInOnly         RasmanUsage = 32
	RasmanUsageCallOutboundRouter RasmanUsage = 64
)

func (o RasmanUsage) String() string {
	switch o {
	case RasmanUsageCallNone:
		return "RasmanUsageCallNone"
	case RasmanUsageCallIn:
		return "RasmanUsageCallIn"
	case RasmanUsageCallOut:
		return "RasmanUsageCallOut"
	case RasmanUsageCallRouter:
		return "RasmanUsageCallRouter"
	case RasmanUsageCallLogon:
		return "RasmanUsageCallLogon"
	case RasmanUsageCallOutOnly:
		return "RasmanUsageCallOutOnly"
	case RasmanUsageCallInOnly:
		return "RasmanUsageCallInOnly"
	case RasmanUsageCallOutboundRouter:
		return "RasmanUsageCallOutboundRouter"
	}
	return "Invalid"
}

// RequestBuffer structure represents RequestBuffer RPC structure.
type RequestBuffer struct {
	RbPcbIndex uint32       `idl:"name:RB_PCBIndex" json:"rb_pcb_index"`
	RbReqtype  RequestTypes `idl:"name:RB_Reqtype" json:"rb_reqtype"`
	RbDummy    uint32       `idl:"name:RB_Dummy" json:"rb_dummy"`
	RbDone     uint32       `idl:"name:RB_Done" json:"rb_done"`
	Alignment  int64        `idl:"name:Alignment" json:"alignment"`
	RbBuffer   []byte       `idl:"name:RB_Buffer" json:"rb_buffer"`
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
	if err := w.WriteData(o.RbPcbIndex); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RbReqtype)); err != nil {
		return err
	}
	if err := w.WriteData(o.RbDummy); err != nil {
		return err
	}
	if err := w.WriteData(o.RbDone); err != nil {
		return err
	}
	if err := w.WriteData(o.Alignment); err != nil {
		return err
	}
	for i1 := range o.RbBuffer {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.RbBuffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RbBuffer); uint64(i1) < 1; i1++ {
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
	if err := w.ReadData(&o.RbPcbIndex); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RbReqtype)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RbDummy); err != nil {
		return err
	}
	if err := w.ReadData(&o.RbDone); err != nil {
		return err
	}
	if err := w.ReadData(&o.Alignment); err != nil {
		return err
	}
	o.RbBuffer = make([]byte, 1)
	for i1 := range o.RbBuffer {
		i1 := i1
		if err := w.ReadData(&o.RbBuffer[i1]); err != nil {
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
	Retcode      uint32 `idl:"name:retcode" json:"retcode"`
	Version      uint32 `idl:"name:dwVersion" json:"version"`
	BufferLength uint32 `idl:"name:cbBuffer" json:"buffer_length"`
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	Abdata       []byte `idl:"name:abdata" json:"abdata"`
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
	if err := w.WriteData(o.Retcode); err != nil {
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
	for i1 := range o.Abdata {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Abdata[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Abdata); uint64(i1) < 1; i1++ {
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
	if err := w.ReadData(&o.Retcode); err != nil {
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
	o.Abdata = make([]byte, 1)
	for i1 := range o.Abdata {
		i1 := i1
		if err := w.ReadData(&o.Abdata[i1]); err != nil {
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
	EndPointsLength       uint32        `idl:"name:dwNumEndPoints" json:"end_points_length"`
	MaxOutCalls           uint32        `idl:"name:dwMaxOutCalls" json:"max_out_calls"`
	MaxInCalls            uint32        `idl:"name:dwMaxInCalls" json:"max_in_calls"`
	MinWANEndPoints       uint32        `idl:"name:dwMinWanEndPoints" json:"min_wan_end_points"`
	MaxWANEndPoints       uint32        `idl:"name:dwMaxWanEndPoints" json:"max_wan_end_points"`
	DeviceType            Rasdevicetype `idl:"name:eDeviceType" json:"device_type"`
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
	if err := w.WriteData(o.EndPointsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxOutCalls); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxInCalls); err != nil {
		return err
	}
	if err := w.WriteData(o.MinWANEndPoints); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxWANEndPoints); err != nil {
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
	if err := w.ReadData(&o.EndPointsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxOutCalls); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxInCalls); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinWANEndPoints); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxWANEndPoints); err != nil {
		return err
	}
	_eDeviceType := uint16(o.DeviceType)
	if err := w.ReadEnum(&_eDeviceType); err != nil {
		return err
	}
	o.DeviceType = Rasdevicetype(_eDeviceType)
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
	Retcode uint32           `idl:"name:retcode" json:"retcode"`
	Write   bool             `idl:"name:fWrite" json:"write"`
	Size    uint32           `idl:"name:dwSize" json:"size"`
	Device  *dtyp.GUID       `idl:"name:guidDevice" json:"device"`
	Rdi     *RASDeviceInfo   `idl:"name:rdi" json:"rdi"`
	RciInfo *RASCalledidInfo `idl:"name:rciInfo" json:"rci_info"`
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
	if err := w.WriteData(o.Retcode); err != nil {
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
	if o.Rdi != nil {
		if err := o.Rdi.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASDeviceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RciInfo != nil {
		if err := o.RciInfo.MarshalNDR(ctx, w); err != nil {
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
	if err := w.ReadData(&o.Retcode); err != nil {
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
	if o.Rdi == nil {
		o.Rdi = &RASDeviceInfo{}
	}
	if err := o.Rdi.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RciInfo == nil {
		o.RciInfo = &RASCalledidInfo{}
	}
	if err := o.RciInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RASNdiswanDriverInfo structure represents RAS_NDISWAN_DRIVER_INFO RPC structure.
type RASNdiswanDriverInfo struct {
	DriverCaps uint32 `idl:"name:DriverCaps" json:"driver_caps"`
	_          uint32 `idl:"name:Reserved"`
}

func (o *RASNdiswanDriverInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RASNdiswanDriverInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RASNdiswanDriverInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// GetNdiswanDriverCaps structure represents GetNdiswanDriverCapsStruct RPC structure.
type GetNdiswanDriverCaps struct {
	Retcode           uint32                `idl:"name:retcode" json:"retcode"`
	NdiswanDriverInfo *RASNdiswanDriverInfo `idl:"name:NdiswanDriverInfo" json:"ndiswan_driver_info"`
}

func (o *GetNdiswanDriverCaps) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *GetNdiswanDriverCaps) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Retcode); err != nil {
		return err
	}
	if o.NdiswanDriverInfo != nil {
		if err := o.NdiswanDriverInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RASNdiswanDriverInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetNdiswanDriverCaps) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Retcode); err != nil {
		return err
	}
	if o.NdiswanDriverInfo == nil {
		o.NdiswanDriverInfo = &RASNdiswanDriverInfo{}
	}
	if err := o.NdiswanDriverInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GetDevConfig structure represents GetDevConfigStruct RPC structure.
type GetDevConfig struct {
	Retcode    uint32 `idl:"name:retcode" json:"retcode"`
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
	if err := w.WriteData(o.Retcode); err != nil {
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
	if err := w.ReadData(&o.Retcode); err != nil {
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
	Retcode uint32 `idl:"name:retcode" json:"retcode"`
	Size    uint32 `idl:"name:size" json:"size"`
	Entries uint32 `idl:"name:entries" json:"entries"`
	Buffer  []byte `idl:"name:buffer" json:"buffer"`
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
	if err := w.WriteData(o.Retcode); err != nil {
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
	if err := w.ReadData(&o.Retcode); err != nil {
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

// RasmanPort32 structure represents RASMAN_PORT_32 RPC structure.
type RasmanPort32 struct {
	PPort            uint32        `idl:"name:P_Port" json:"p_port"`
	PPortName        []byte        `idl:"name:P_PortName" json:"p_port_name"`
	PStatus          RasmanStatus  `idl:"name:P_Status" json:"p_status"`
	PRdtDeviceType   Rasdevicetype `idl:"name:P_rdtDeviceType" json:"p_rdt_device_type"`
	PConfiguredUsage RasmanUsage   `idl:"name:P_ConfiguredUsage" json:"p_configured_usage"`
	PCurrentUsage    RasmanUsage   `idl:"name:P_CurrentUsage" json:"p_current_usage"`
	PMediaName       []byte        `idl:"name:P_MediaName" json:"p_media_name"`
	PDeviceType      []byte        `idl:"name:P_DeviceType" json:"p_device_type"`
	PDeviceName      []byte        `idl:"name:P_DeviceName" json:"p_device_name"`
	PLineDeviceID    uint32        `idl:"name:P_LineDeviceId" json:"p_line_device_id"`
	PAddressID       uint32        `idl:"name:P_AddressId" json:"p_address_id"`
}

func (o *RasmanPort32) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasmanPort32) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PPort); err != nil {
		return err
	}
	for i1 := range o.PPortName {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.PPortName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PPortName); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.PStatus)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PRdtDeviceType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PConfiguredUsage)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PCurrentUsage)); err != nil {
		return err
	}
	for i1 := range o.PMediaName {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.PMediaName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PMediaName); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.PDeviceType {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.PDeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PDeviceType); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.PDeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.PDeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PDeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PLineDeviceID); err != nil {
		return err
	}
	if err := w.WriteData(o.PAddressID); err != nil {
		return err
	}
	return nil
}
func (o *RasmanPort32) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PPort); err != nil {
		return err
	}
	o.PPortName = make([]byte, 16)
	for i1 := range o.PPortName {
		i1 := i1
		if err := w.ReadData(&o.PPortName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadEnum((*uint16)(&o.PStatus)); err != nil {
		return err
	}
	_ePRdtDeviceType := uint16(o.PRdtDeviceType)
	if err := w.ReadEnum(&_ePRdtDeviceType); err != nil {
		return err
	}
	o.PRdtDeviceType = Rasdevicetype(_ePRdtDeviceType)
	if err := w.ReadEnum((*uint16)(&o.PConfiguredUsage)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PCurrentUsage)); err != nil {
		return err
	}
	o.PMediaName = make([]byte, 16)
	for i1 := range o.PMediaName {
		i1 := i1
		if err := w.ReadData(&o.PMediaName[i1]); err != nil {
			return err
		}
	}
	o.PDeviceType = make([]byte, 16)
	for i1 := range o.PDeviceType {
		i1 := i1
		if err := w.ReadData(&o.PDeviceType[i1]); err != nil {
			return err
		}
	}
	o.PDeviceName = make([]byte, 129)
	for i1 := range o.PDeviceName {
		i1 := i1
		if err := w.ReadData(&o.PDeviceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PLineDeviceID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PAddressID); err != nil {
		return err
	}
	return nil
}

// RasmanInfo structure represents RASMAN_INFO RPC structure.
type RasmanInfo struct {
	RiPortStatus           RasmanStatus         `idl:"name:RI_PortStatus" json:"ri_port_status"`
	RiConnState            RasmanState          `idl:"name:RI_ConnState" json:"ri_conn_state"`
	RiLinkSpeed            uint32               `idl:"name:RI_LinkSpeed" json:"ri_link_speed"`
	RiLastError            uint32               `idl:"name:RI_LastError" json:"ri_last_error"`
	RiCurrentUsage         RasmanUsage          `idl:"name:RI_CurrentUsage" json:"ri_current_usage"`
	RiDeviceTypeConnecting []byte               `idl:"name:RI_DeviceTypeConnecting" json:"ri_device_type_connecting"`
	RiDeviceConnecting     []byte               `idl:"name:RI_DeviceConnecting" json:"ri_device_connecting"`
	RiStringDeviceType     []byte               `idl:"name:RI_szDeviceType" json:"ri_string_device_type"`
	RiStringDeviceName     []byte               `idl:"name:RI_szDeviceName" json:"ri_string_device_name"`
	RiStringPortName       []byte               `idl:"name:RI_szPortName" json:"ri_string_port_name"`
	RiDisconnectType       RasmanDisconnectType `idl:"name:RI_DisconnectType" json:"ri_disconnect_type"`
	RiOwnershipFlag        uint32               `idl:"name:RI_OwnershipFlag" json:"ri_ownership_flag"`
	RiConnectDuration      uint32               `idl:"name:RI_ConnectDuration" json:"ri_connect_duration"`
	RiBytesReceived        uint32               `idl:"name:RI_BytesReceived" json:"ri_bytes_received"`
	RiPhonebook            []byte               `idl:"name:RI_Phonebook" json:"ri_phonebook"`
	RiPhoneEntry           []byte               `idl:"name:RI_PhoneEntry" json:"ri_phone_entry"`
	RiConnectionHandle     []byte               `idl:"name:RI_ConnectionHandle" json:"ri_connection_handle"`
	RiSubEntry             uint32               `idl:"name:RI_SubEntry" json:"ri_sub_entry"`
	RiRdtDeviceType        Rasdevicetype        `idl:"name:RI_rdtDeviceType" json:"ri_rdt_device_type"`
	RiGUIDEntry            *dtyp.GUID           `idl:"name:RI_GuidEntry" json:"ri_guid_entry"`
	RiDwSessionID          uint32               `idl:"name:RI_dwSessionId" json:"ri_dw_session_id"`
	RiDwFlags              uint32               `idl:"name:RI_dwFlags" json:"ri_dw_flags"`
	RiCorrelationGUID      *dtyp.GUID           `idl:"name:RI_CorrelationGuid" json:"ri_correlation_guid"`
}

func (o *RasmanInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasmanInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RiPortStatus)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RiConnState)); err != nil {
		return err
	}
	if err := w.WriteData(o.RiLinkSpeed); err != nil {
		return err
	}
	if err := w.WriteData(o.RiLastError); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RiCurrentUsage)); err != nil {
		return err
	}
	for i1 := range o.RiDeviceTypeConnecting {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RiDeviceTypeConnecting[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RiDeviceTypeConnecting); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RiDeviceConnecting {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.RiDeviceConnecting[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RiDeviceConnecting); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RiStringDeviceType {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RiStringDeviceType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RiStringDeviceType); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RiStringDeviceName {
		i1 := i1
		if uint64(i1) >= 129 {
			break
		}
		if err := w.WriteData(o.RiStringDeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RiStringDeviceName); uint64(i1) < 129; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RiStringPortName {
		i1 := i1
		if uint64(i1) >= 17 {
			break
		}
		if err := w.WriteData(o.RiStringPortName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RiStringPortName); uint64(i1) < 17; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.RiDisconnectType)); err != nil {
		return err
	}
	if err := w.WriteData(o.RiOwnershipFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.RiConnectDuration); err != nil {
		return err
	}
	if err := w.WriteData(o.RiBytesReceived); err != nil {
		return err
	}
	for i1 := range o.RiPhonebook {
		i1 := i1
		if uint64(i1) >= 261 {
			break
		}
		if err := w.WriteData(o.RiPhonebook[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RiPhonebook); uint64(i1) < 261; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.RiPhoneEntry {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.RiPhoneEntry[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RiPhoneEntry); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.RiConnectionHandle != nil {
		_ptr_RI_ConnectionHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			// FIXME unknown type RI_ConnectionHandle
			return nil
		})
		if err := w.WritePointer(&o.RiConnectionHandle, _ptr_RI_ConnectionHandle); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RiSubEntry); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.RiRdtDeviceType)); err != nil {
		return err
	}
	if o.RiGUIDEntry != nil {
		if err := o.RiGUIDEntry.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RiDwSessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.RiDwFlags); err != nil {
		return err
	}
	if o.RiCorrelationGUID != nil {
		if err := o.RiCorrelationGUID.MarshalNDR(ctx, w); err != nil {
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
func (o *RasmanInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RiPortStatus)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RiConnState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiLinkSpeed); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiLastError); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.RiCurrentUsage)); err != nil {
		return err
	}
	o.RiDeviceTypeConnecting = make([]byte, 16)
	for i1 := range o.RiDeviceTypeConnecting {
		i1 := i1
		if err := w.ReadData(&o.RiDeviceTypeConnecting[i1]); err != nil {
			return err
		}
	}
	o.RiDeviceConnecting = make([]byte, 129)
	for i1 := range o.RiDeviceConnecting {
		i1 := i1
		if err := w.ReadData(&o.RiDeviceConnecting[i1]); err != nil {
			return err
		}
	}
	o.RiStringDeviceType = make([]byte, 16)
	for i1 := range o.RiStringDeviceType {
		i1 := i1
		if err := w.ReadData(&o.RiStringDeviceType[i1]); err != nil {
			return err
		}
	}
	o.RiStringDeviceName = make([]byte, 129)
	for i1 := range o.RiStringDeviceName {
		i1 := i1
		if err := w.ReadData(&o.RiStringDeviceName[i1]); err != nil {
			return err
		}
	}
	o.RiStringPortName = make([]byte, 17)
	for i1 := range o.RiStringPortName {
		i1 := i1
		if err := w.ReadData(&o.RiStringPortName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadEnum((*uint16)(&o.RiDisconnectType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiOwnershipFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiConnectDuration); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiBytesReceived); err != nil {
		return err
	}
	o.RiPhonebook = make([]byte, 261)
	for i1 := range o.RiPhonebook {
		i1 := i1
		if err := w.ReadData(&o.RiPhonebook[i1]); err != nil {
			return err
		}
	}
	o.RiPhoneEntry = make([]byte, 257)
	for i1 := range o.RiPhoneEntry {
		i1 := i1
		if err := w.ReadData(&o.RiPhoneEntry[i1]); err != nil {
			return err
		}
	}
	_ptr_RI_ConnectionHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		// FIXME: unknown type RI_ConnectionHandle
		return nil
	})
	_s_RI_ConnectionHandle := func(ptr interface{}) { o.RiConnectionHandle = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.RiConnectionHandle, _s_RI_ConnectionHandle, _ptr_RI_ConnectionHandle); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiSubEntry); err != nil {
		return err
	}
	_eRiRdtDeviceType := uint16(o.RiRdtDeviceType)
	if err := w.ReadEnum(&_eRiRdtDeviceType); err != nil {
		return err
	}
	o.RiRdtDeviceType = Rasdevicetype(_eRiRdtDeviceType)
	if o.RiGUIDEntry == nil {
		o.RiGUIDEntry = &dtyp.GUID{}
	}
	if err := o.RiGUIDEntry.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiDwSessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RiDwFlags); err != nil {
		return err
	}
	if o.RiCorrelationGUID == nil {
		o.RiCorrelationGUID = &dtyp.GUID{}
	}
	if err := o.RiCorrelationGUID.UnmarshalNDR(ctx, w); err != nil {
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
	Info   *RasmanInfo  `idl:"name:info" json:"info"`
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
		if err := (&RasmanInfo{}).MarshalNDR(ctx, w); err != nil {
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
		o.Info = &RasmanInfo{}
	}
	if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type Info_Field1 struct {
	Retcode      uint32 `idl:"name:retcode" json:"retcode"`
	PaddingField []byte `idl:"name:paddingField" json:"padding_field"`
}

// RasrpcCallbacklist structure represents RASRPC_CALLBACKLIST RPC structure.
type RasrpcCallbacklist struct {
	PortName   []uint16            `idl:"name:pszPortName" json:"port_name"`
	DeviceName []uint16            `idl:"name:pszDeviceName" json:"device_name"`
	Number     []uint16            `idl:"name:pszNumber" json:"number"`
	DeviceType uint32              `idl:"name:dwDeviceType" json:"device_type"`
	Next       *RasrpcCallbacklist `idl:"name:pNext;pointer:unique" json:"next"`
}

func (o *RasrpcCallbacklist) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasrpcCallbacklist) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
				if err := (&RasrpcCallbacklist{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RasrpcCallbacklist) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
			o.Next = &RasrpcCallbacklist{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**RasrpcCallbacklist) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	return nil
}

// RasrpcStringlist structure represents RASRPC_STRINGLIST RPC structure.
type RasrpcStringlist struct {
	Psz  []uint16          `idl:"name:psz" json:"psz"`
	Next *RasrpcStringlist `idl:"name:pNext;pointer:unique" json:"next"`
}

func (o *RasrpcStringlist) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasrpcStringlist) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	for i1 := range o.Psz {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.Psz[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Psz); uint64(i1) < 256; i1++ {
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
				if err := (&RasrpcStringlist{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RasrpcStringlist) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	o.Psz = make([]uint16, 256)
	for i1 := range o.Psz {
		i1 := i1
		if err := w.ReadData(&o.Psz[i1]); err != nil {
			return err
		}
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &RasrpcStringlist{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**RasrpcStringlist) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	return nil
}

// RasrpcLocationlist structure represents RASRPC_LOCATIONLIST RPC structure.
type RasrpcLocationlist struct {
	LocationID uint32              `idl:"name:dwLocationId" json:"location_id"`
	Prefix     uint32              `idl:"name:iPrefix" json:"prefix"`
	Suffix     uint32              `idl:"name:iSuffix" json:"suffix"`
	Next       *RasrpcLocationlist `idl:"name:pNext;pointer:unique" json:"next"`
}

func (o *RasrpcLocationlist) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasrpcLocationlist) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
				if err := (&RasrpcLocationlist{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RasrpcLocationlist) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
			o.Next = &RasrpcLocationlist{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**RasrpcLocationlist) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	return nil
}

// RasrpcPbuser structure represents RASRPC_PBUSER RPC structure.
type RasrpcPbuser struct {
	OperatorDial             bool                `idl:"name:fOperatorDial" json:"operator_dial"`
	PreviewPhoneNumber       bool                `idl:"name:fPreviewPhoneNumber" json:"preview_phone_number"`
	UseLocation              bool                `idl:"name:fUseLocation" json:"use_location"`
	ShowLights               bool                `idl:"name:fShowLights" json:"show_lights"`
	ShowConnectStatus        bool                `idl:"name:fShowConnectStatus" json:"show_connect_status"`
	CloseOnDial              bool                `idl:"name:fCloseOnDial" json:"close_on_dial"`
	AllowLogonPhonebookEdits bool                `idl:"name:fAllowLogonPhonebookEdits" json:"allow_logon_phonebook_edits"`
	AllowLogonLocationEdits  bool                `idl:"name:fAllowLogonLocationEdits" json:"allow_logon_location_edits"`
	SkipConnectComplete      bool                `idl:"name:fSkipConnectComplete" json:"skip_connect_complete"`
	NewEntryWizard           bool                `idl:"name:fNewEntryWizard" json:"new_entry_wizard"`
	RedialAttempts           uint32              `idl:"name:dwRedialAttempts" json:"redial_attempts"`
	RedialSeconds            uint32              `idl:"name:dwRedialSeconds" json:"redial_seconds"`
	IdleDisconnectSeconds    uint32              `idl:"name:dwIdleDisconnectSeconds" json:"idle_disconnect_seconds"`
	RedialOnLinkFailure      bool                `idl:"name:fRedialOnLinkFailure" json:"redial_on_link_failure"`
	PopupOnTopWhenRedialing  bool                `idl:"name:fPopupOnTopWhenRedialing" json:"popup_on_top_when_redialing"`
	ExpandAutoDialQuery      bool                `idl:"name:fExpandAutoDialQuery" json:"expand_auto_dial_query"`
	CallbackMode             uint32              `idl:"name:dwCallbackMode" json:"callback_mode"`
	Callbacks                *RasrpcCallbacklist `idl:"name:pCallbacks;pointer:unique" json:"callbacks"`
	LastCallbackByCaller     []uint16            `idl:"name:pszLastCallbackByCaller" json:"last_callback_by_caller"`
	PhonebookMode            uint32              `idl:"name:dwPhonebookMode" json:"phonebook_mode"`
	PersonalFile             []uint16            `idl:"name:pszPersonalFile" json:"personal_file"`
	AlternatePath            []uint16            `idl:"name:pszAlternatePath" json:"alternate_path"`
	Phonebooks               *RasrpcStringlist   `idl:"name:pPhonebooks;pointer:unique" json:"phonebooks"`
	AreaCodes                *RasrpcStringlist   `idl:"name:pAreaCodes;pointer:unique" json:"area_codes"`
	UseAreaAndCountry        bool                `idl:"name:fUseAreaAndCountry" json:"use_area_and_country"`
	Prefixes                 *RasrpcStringlist   `idl:"name:pPrefixes;pointer:unique" json:"prefixes"`
	Suffixes                 *RasrpcStringlist   `idl:"name:pSuffixes;pointer:unique" json:"suffixes"`
	Locations                *RasrpcLocationlist `idl:"name:pLocations;pointer:unique" json:"locations"`
	XPhonebook               uint32              `idl:"name:dwXPhonebook" json:"x_phonebook"`
	YPhonebook               uint32              `idl:"name:dwYPhonebook" json:"y_phonebook"`
	DefaultEntry             []uint16            `idl:"name:pszDefaultEntry" json:"default_entry"`
	Initialized              bool                `idl:"name:fInitialized" json:"initialized"`
	Dirty                    bool                `idl:"name:fDirty" json:"dirty"`
}

func (o *RasrpcPbuser) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RasrpcPbuser) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
				if err := (&RasrpcCallbacklist{}).MarshalNDR(ctx, w); err != nil {
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
				if err := (&RasrpcStringlist{}).MarshalNDR(ctx, w); err != nil {
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
				if err := (&RasrpcStringlist{}).MarshalNDR(ctx, w); err != nil {
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
				if err := (&RasrpcStringlist{}).MarshalNDR(ctx, w); err != nil {
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
				if err := (&RasrpcStringlist{}).MarshalNDR(ctx, w); err != nil {
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
				if err := (&RasrpcLocationlist{}).MarshalNDR(ctx, w); err != nil {
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
func (o *RasrpcPbuser) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
			o.Callbacks = &RasrpcCallbacklist{}
		}
		if err := o.Callbacks.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pCallbacks := func(ptr interface{}) { o.Callbacks = *ptr.(**RasrpcCallbacklist) }
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
			o.Phonebooks = &RasrpcStringlist{}
		}
		if err := o.Phonebooks.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pPhonebooks := func(ptr interface{}) { o.Phonebooks = *ptr.(**RasrpcStringlist) }
	if err := w.ReadPointer(&o.Phonebooks, _s_pPhonebooks, _ptr_pPhonebooks); err != nil {
		return err
	}
	_ptr_pAreaCodes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AreaCodes == nil {
			o.AreaCodes = &RasrpcStringlist{}
		}
		if err := o.AreaCodes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pAreaCodes := func(ptr interface{}) { o.AreaCodes = *ptr.(**RasrpcStringlist) }
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
			o.Prefixes = &RasrpcStringlist{}
		}
		if err := o.Prefixes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pPrefixes := func(ptr interface{}) { o.Prefixes = *ptr.(**RasrpcStringlist) }
	if err := w.ReadPointer(&o.Prefixes, _s_pPrefixes, _ptr_pPrefixes); err != nil {
		return err
	}
	_ptr_pSuffixes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Suffixes == nil {
			o.Suffixes = &RasrpcStringlist{}
		}
		if err := o.Suffixes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pSuffixes := func(ptr interface{}) { o.Suffixes = *ptr.(**RasrpcStringlist) }
	if err := w.ReadPointer(&o.Suffixes, _s_pSuffixes, _ptr_pSuffixes); err != nil {
		return err
	}
	_ptr_pLocations := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Locations == nil {
			o.Locations = &RasrpcLocationlist{}
		}
		if err := o.Locations.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pLocations := func(ptr interface{}) { o.Locations = *ptr.(**RasrpcLocationlist) }
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

// RemoteIcficsConfig structure represents IRemoteICFICSConfig RPC structure.
type RemoteIcficsConfig dcom.InterfacePointer

func (o *RemoteIcficsConfig) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteIcficsConfig) xxx_PreparePayload(ctx context.Context) error {
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

func (o *RemoteIcficsConfig) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteIcficsConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteIcficsConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteSstpCertCheck structure represents IRemoteSstpCertCheck RPC structure.
type RemoteSstpCertCheck dcom.InterfacePointer

func (o *RemoteSstpCertCheck) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteSstpCertCheck) xxx_PreparePayload(ctx context.Context) error {
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

func (o *RemoteSstpCertCheck) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteSstpCertCheck) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteSstpCertCheck) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

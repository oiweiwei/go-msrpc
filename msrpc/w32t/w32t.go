// The w32t package implements the W32T client protocol.
//
// # Introduction
//
// The W32Time Remote Protocol is a remote procedure call (RPC) interface for controlling
// and monitoring a time service that implements the Network Time Protocol (NTP) Authentication
// Extensions [MS-SNTP].
//
// # Overview
//
// The W32Time Remote Protocol is an RPC-based protocol used for controlling and monitoring
// a time service that implements the Network Time Protocol (NTP) Authentication Extensions
// specified in [MS-SNTP].
//
// The client side of the W32Time Remote Protocol is an application that issues method
// calls on the RPC interface.
//
// The server side of the W32Time Remote Protocol provides methods for controlling and
// monitoring the client and server instances of the locally hosted NTP Authentication
// Extensions [MS-SNTP] implementation.<1>
package w32t

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
	GoPackage = "w32t"
)

// NTPPeerInfo structure represents W32TIME_NTP_PEER_INFO RPC structure.
//
// The W32TIME_NTP_PEER_INFO structure defines the current state of a time peer for
// an NTP time provider.
type NTPPeerInfo struct {
	// ulSize:  The size, in bytes, of this structure.<8>
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulResolveAttempts:  The ResolvedAttempts element value (see section 3.2.1.3).
	ResolveAttempts uint32 `idl:"name:ulResolveAttempts" json:"resolve_attempts"`
	// u64TimeRemaining:  The TimeRemaining element value (see section 3.2.1.3).
	TimeRemaining uint64 `idl:"name:u64TimeRemaining" json:"time_remaining"`
	// u64LastSuccessfulSync:  The LastSuccessfulSync element value (see section 3.2.1.3).
	LastSuccessfulSync uint64 `idl:"name:u64LastSuccessfulSync" json:"last_successful_sync"`
	// ulLastSyncError:  The LastSyncError element value (see section 3.2.1.3).
	LastSyncError uint32 `idl:"name:ulLastSyncError" json:"last_sync_error"`
	// ulLastSyncErrorMsgId:  The LastSyncErrorMessageId element value (see section 3.2.1.3).
	LastSyncErrorMessageID uint32 `idl:"name:ulLastSyncErrorMsgId" json:"last_sync_error_message_id"`
	// ulValidDataCounter:  The ValidDataCounter element value (see section 3.2.1.3).
	ValidDataCounter uint32 `idl:"name:ulValidDataCounter" json:"valid_data_counter"`
	// ulAuthTypeMsgId:  The AuthenticationTypeMessageId element value (see section 3.2.1.3).
	AuthTypeMessageID uint32 `idl:"name:ulAuthTypeMsgId" json:"auth_type_message_id"`
	// wszUniqueName:  The PeerName element value (see section 3.2.1.3).
	UniqueName string `idl:"name:wszUniqueName;string;pointer:unique" json:"unique_name"`
	// ulMode:  This time peer's current NTP association mode, as specified in [RFC1305]
	// section 3.2.1, "Common Variables".
	Mode uint8 `idl:"name:ulMode" json:"mode"`
	// ulStratum:  This time peer's stratum level, which indicates the distance between
	// this time peer and a reference source. This value is compared with other peers' stratum
	// levels to ensure that a machine closer to a reference source is not synchronized
	// to a machine that is farther away, as specified in [RFC1305] section 2.2, "Network
	// Configurations".
	Stratum uint8 `idl:"name:ulStratum" json:"stratum"`
	// ulReachability:  An 8-bit shift register that contains this time peer's reachability,
	// as specified in [RFC1305] section 3.2.3, "Peer Variables".
	Reachability uint8 `idl:"name:ulReachability" json:"reachability"`
	// ulPeerPollInterval:  This time peer's poll interval, expressed as specified in [RFC1305],
	// using units of seconds given as exponents to a power of two. For example, a value
	// of six indicates a minimum interval of 64 seconds.
	PeerPollInterval uint8 `idl:"name:ulPeerPollInterval" json:"peer_poll_interval"`
	// ulHostPollInterval:  The interval at which the NTP service provider is polling this
	// time peer, expressed as specified in [RFC1305], using units of seconds given as exponents
	// to a power of two. For example, a value of six indicates a minimum interval of 64
	// seconds.
	HostPollInterval uint8 `idl:"name:ulHostPollInterval" json:"host_poll_interval"`
}

func (o *NTPPeerInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTPPeerInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.ResolveAttempts); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeRemaining); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSuccessfulSync); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSyncError); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSyncErrorMessageID); err != nil {
		return err
	}
	if err := w.WriteData(o.ValidDataCounter); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthTypeMessageID); err != nil {
		return err
	}
	if o.UniqueName != "" {
		_ptr_wszUniqueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.UniqueName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UniqueName, _ptr_wszUniqueName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Mode); err != nil {
		return err
	}
	if err := w.WriteData(o.Stratum); err != nil {
		return err
	}
	if err := w.WriteData(o.Reachability); err != nil {
		return err
	}
	if err := w.WriteData(o.PeerPollInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.HostPollInterval); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *NTPPeerInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResolveAttempts); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeRemaining); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSuccessfulSync); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSyncError); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSyncErrorMessageID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValidDataCounter); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthTypeMessageID); err != nil {
		return err
	}
	_ptr_wszUniqueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.UniqueName); err != nil {
			return err
		}
		return nil
	})
	_s_wszUniqueName := func(ptr interface{}) { o.UniqueName = *ptr.(*string) }
	if err := w.ReadPointer(&o.UniqueName, _s_wszUniqueName, _ptr_wszUniqueName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mode); err != nil {
		return err
	}
	if err := w.ReadData(&o.Stratum); err != nil {
		return err
	}
	if err := w.ReadData(&o.Reachability); err != nil {
		return err
	}
	if err := w.ReadData(&o.PeerPollInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.HostPollInterval); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// NTPProviderData structure represents W32TIME_NTP_PROVIDER_DATA RPC structure.
//
// The W32TIME_NTP_PROVIDER_DATA structure defines the state of an NTP time provider.
type NTPProviderData struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulError:  The LastError element value (see section 3.2.1.2). Because the values transmitted
	// in this field are implementation-specific, all nonzero values MUST be treated as
	// equivalent for protocol purposes.<6>
	Error uint32 `idl:"name:ulError" json:"error"`
	// ulErrorMsgId:  The LastErrorMessageId element value (see section 3.2.1.2). The values
	// in this field are implementation-specific. If an implementation receives a value
	// it does not understand, the implementation MUST ignore the value.<7>
	ErrorMessageID uint32 `idl:"name:ulErrorMsgId" json:"error_message_id"`
	// cPeerInfo:  The number of active time peers that synchronize with this NTP time provider.
	// This value also indicates the number of structures in pPeerInfo.
	PeerInfoCount uint32 `idl:"name:cPeerInfo" json:"peer_info_count"`
	// pPeerInfo:  The PeerList element value (see section 3.2.1.2). A pointer to W32TIME_NTP_PEER_INFO
	// structures representing the time peers with which this time provider is currently
	// synchronizing.
	PeerInfo []*NTPPeerInfo `idl:"name:pPeerInfo;size_is:(cPeerInfo)" json:"peer_info"`
}

func (o *NTPProviderData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.PeerInfo != nil && o.PeerInfoCount == 0 {
		o.PeerInfoCount = uint32(len(o.PeerInfo))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTPProviderData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	if err := w.WriteData(o.ErrorMessageID); err != nil {
		return err
	}
	if err := w.WriteData(o.PeerInfoCount); err != nil {
		return err
	}
	if o.PeerInfo != nil || o.PeerInfoCount > 0 {
		_ptr_pPeerInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.PeerInfoCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.PeerInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.PeerInfo[i1] != nil {
					if err := o.PeerInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NTPPeerInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.PeerInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&NTPPeerInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PeerInfo, _ptr_pPeerInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTPProviderData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	if err := w.ReadData(&o.ErrorMessageID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PeerInfoCount); err != nil {
		return err
	}
	_ptr_pPeerInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.PeerInfoCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.PeerInfoCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.PeerInfo", sizeInfo[0])
		}
		o.PeerInfo = make([]*NTPPeerInfo, sizeInfo[0])
		for i1 := range o.PeerInfo {
			i1 := i1
			if o.PeerInfo[i1] == nil {
				o.PeerInfo[i1] = &NTPPeerInfo{}
			}
			if err := o.PeerInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pPeerInfo := func(ptr interface{}) { o.PeerInfo = *ptr.(*[]*NTPPeerInfo) }
	if err := w.ReadPointer(&o.PeerInfo, _s_pPeerInfo, _ptr_pPeerInfo); err != nil {
		return err
	}
	return nil
}

// HardwareProviderData structure represents W32TIME_HARDWARE_PROVIDER_DATA RPC structure.
//
// The W32TIME_HARDWARE_PROVIDER_DATA structure contains operational information about
// a hardware time provider, such as a cesium or atomic clock.
type HardwareProviderData struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulError:  The LastError element value (see section 3.2.1.2). Because the values transmitted
	// in this field are implementation specific, all nonzero values MUST be treated as
	// equivalent for the purposes of this protocol.<4>
	Error uint32 `idl:"name:ulError" json:"error"`
	// ulErrorMsgId:  The LastErrorMessageId element value (see section 3.2.1.2). The values
	// in this field are implementation-specific. If an implementation receives a value
	// it does not understand, the implementation MUST ignore the value.<5>
	ErrorMessageID uint32 `idl:"name:ulErrorMsgId" json:"error_message_id"`
	// wszReferenceIdentifier:  The Reference Clock Identifier that identifies the time
	// source for this time service, as specified in [RFC1305] Appendix A, "NTP Data Format".
	ReferenceID string `idl:"name:wszReferenceIdentifier;string;pointer:unique" json:"reference_id"`
}

func (o *HardwareProviderData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *HardwareProviderData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	if err := w.WriteData(o.ErrorMessageID); err != nil {
		return err
	}
	if o.ReferenceID != "" {
		_ptr_wszReferenceIdentifier := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ReferenceID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ReferenceID, _ptr_wszReferenceIdentifier); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *HardwareProviderData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	if err := w.ReadData(&o.ErrorMessageID); err != nil {
		return err
	}
	_ptr_wszReferenceIdentifier := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ReferenceID); err != nil {
			return err
		}
		return nil
	})
	_s_wszReferenceIdentifier := func(ptr interface{}) { o.ReferenceID = *ptr.(*string) }
	if err := w.ReadPointer(&o.ReferenceID, _s_wszReferenceIdentifier, _ptr_wszReferenceIdentifier); err != nil {
		return err
	}
	return nil
}

// ProviderData structure represents W32TIME_PROVIDER_DATA RPC union.
type ProviderData struct {
	// Types that are assignable to Value
	//
	// *ProviderData_NTP
	// *ProviderData_Hardware
	Value is_ProviderData `json:"value"`
}

func (o *ProviderData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ProviderData_NTP:
		if value != nil {
			return value.NTPProviderData
		}
	case *ProviderData_Hardware:
		if value != nil {
			return value.HardwareProviderData
		}
	}
	return nil
}

type is_ProviderData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ProviderData()
}

func (o *ProviderData) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ProviderData_NTP:
		return uint32(0)
	case *ProviderData_Hardware:
		return uint32(1)
	}
	return uint32(0)
}

func (o *ProviderData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
	case uint32(0):
		_o, _ := o.Value.(*ProviderData_NTP)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProviderData_NTP{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*ProviderData_Hardware)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProviderData_Hardware{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ProviderData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
	case uint32(0):
		o.Value = &ProviderData_NTP{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &ProviderData_Hardware{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ProviderData_NTP structure represents W32TIME_PROVIDER_DATA RPC union arm.
//
// It has following labels: 0
type ProviderData_NTP struct {
	NTPProviderData *NTPProviderData `idl:"name:pNtpProviderData" json:"ntp_provider_data"`
}

func (*ProviderData_NTP) is_ProviderData() {}

func (o *ProviderData_NTP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.NTPProviderData != nil {
		_ptr_pNtpProviderData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NTPProviderData != nil {
				if err := o.NTPProviderData.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NTPProviderData{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NTPProviderData, _ptr_pNtpProviderData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProviderData_NTP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pNtpProviderData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NTPProviderData == nil {
			o.NTPProviderData = &NTPProviderData{}
		}
		if err := o.NTPProviderData.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNtpProviderData := func(ptr interface{}) { o.NTPProviderData = *ptr.(**NTPProviderData) }
	if err := w.ReadPointer(&o.NTPProviderData, _s_pNtpProviderData, _ptr_pNtpProviderData); err != nil {
		return err
	}
	return nil
}

// ProviderData_Hardware structure represents W32TIME_PROVIDER_DATA RPC union arm.
//
// It has following labels: 1
type ProviderData_Hardware struct {
	HardwareProviderData *HardwareProviderData `idl:"name:pHardwareProviderData" json:"hardware_provider_data"`
}

func (*ProviderData_Hardware) is_ProviderData() {}

func (o *ProviderData_Hardware) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.HardwareProviderData != nil {
		_ptr_pHardwareProviderData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.HardwareProviderData != nil {
				if err := o.HardwareProviderData.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&HardwareProviderData{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.HardwareProviderData, _ptr_pHardwareProviderData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProviderData_Hardware) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pHardwareProviderData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.HardwareProviderData == nil {
			o.HardwareProviderData = &HardwareProviderData{}
		}
		if err := o.HardwareProviderData.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pHardwareProviderData := func(ptr interface{}) { o.HardwareProviderData = *ptr.(**HardwareProviderData) }
	if err := w.ReadPointer(&o.HardwareProviderData, _s_pHardwareProviderData, _ptr_pHardwareProviderData); err != nil {
		return err
	}
	return nil
}

// ProviderInfo structure represents W32TIME_PROVIDER_INFO RPC structure.
//
// The W32TIME_PROVIDER_INFO structure defines information about a selected time provider
// (either an NTP time provider or a hardware time provider).
type ProviderInfo struct {
	// ulProviderType:  The type of time provider, which MUST be one of the following values.
	//
	// All other values are reserved for future use and servers SHOULD NOT send them.
	//
	//	+-------+------------------------+
	//	|       |                        |
	//	| VALUE |        MEANING         |
	//	|       |                        |
	//	+-------+------------------------+
	//	+-------+------------------------+
	//	|     0 | NTP time provider      |
	//	+-------+------------------------+
	//	|     1 | Hardware time provider |
	//	+-------+------------------------+
	ProviderType uint32 `idl:"name:ulProviderType" json:"provider_type"`
	// ProviderData:  A W32TIME_PROVIDER_DATA union that contains information about the
	// time provider.
	ProviderData *ProviderData `idl:"name:ProviderData;switch_is:ulProviderType" json:"provider_data"`
}

func (o *ProviderInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ProviderInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ProviderType); err != nil {
		return err
	}
	_swProviderData := uint32(o.ProviderType)
	if o.ProviderData != nil {
		if err := o.ProviderData.MarshalUnionNDR(ctx, w, _swProviderData); err != nil {
			return err
		}
	} else {
		if err := (&ProviderData{}).MarshalUnionNDR(ctx, w, _swProviderData); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProviderInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProviderType); err != nil {
		return err
	}
	if o.ProviderData == nil {
		o.ProviderData = &ProviderData{}
	}
	_swProviderData := uint32(o.ProviderType)
	if err := o.ProviderData.UnmarshalUnionNDR(ctx, w, _swProviderData); err != nil {
		return err
	}
	return nil
}

// Entry structure represents W32TIME_ENTRY RPC structure.
//
// The W32TIME_ENTRY structure defines the general entry as a possible extension to
// other time service data structures. This structure has no current use.
type Entry struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// wszName:  A null-terminated string that indicates the name of the entry.
	Name string `idl:"name:wszName;string;pointer:unique" json:"name"`
	// wszValue:  A null-terminated string that indicates the value of the entry.
	Value string `idl:"name:wszValue;string;pointer:unique" json:"value"`
	// wszHelp:  A null-terminated string that indicates the display text of the entry.
	Help string `idl:"name:wszHelp;string;pointer:unique" json:"help"`
}

func (o *Entry) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Entry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Value != "" {
		_ptr_wszValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Value); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_wszValue); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Help != "" {
		_ptr_wszHelp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Help); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Help, _ptr_wszHelp); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Entry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Value); err != nil {
			return err
		}
		return nil
	})
	_s_wszValue := func(ptr interface{}) { o.Value = *ptr.(*string) }
	if err := w.ReadPointer(&o.Value, _s_wszValue, _ptr_wszValue); err != nil {
		return err
	}
	_ptr_wszHelp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Help); err != nil {
			return err
		}
		return nil
	})
	_s_wszHelp := func(ptr interface{}) { o.Help = *ptr.(*string) }
	if err := w.ReadPointer(&o.Help, _s_wszHelp, _ptr_wszHelp); err != nil {
		return err
	}
	return nil
}

// NTPClientProviderConfigData structure represents W32TIME_NTPCLIENT_PROVIDER_CONFIG_DATA RPC structure.
//
// The W32TIME_NTPCLIENT_PROVIDER_CONFIG_DATA structure contains configuration data
// about an NtpClient time provider.
//
// The structure is defined to match the NtpClient time provider's configuration of
// the W32Time implementation. Fields in the structure that do not apply to other implementations
// SHOULD<10> have their corresponding configuration-setting type fields set to W32TIME_CONFIGURATION_SETTING_UNDEFINED.
type NTPClientProviderConfigData struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulAllowNonstandardModeCombinations:  An integer that indicates whether mode combinations
	// that would result in an error action as defined in [RFC1305] Table 5 (Modes and Actions)
	// are allowed.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | Indicates that mode combinations that would result in an error action are not    |
	//	|       | allowed.                                                                         |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | Indicates that mode combinations that would result in an error action are        |
	//	|       | allowed.                                                                         |
	//	+-------+----------------------------------------------------------------------------------+
	AllowNonstandardModeCombinations uint32 `idl:"name:ulAllowNonstandardModeCombinations" json:"allow_nonstandard_mode_combinations"`
	// ulCrossSiteSyncFlags:  The CrossSiteSyncFlags element value (see section 3.2.1.2.1).
	CrossSiteSyncFlags uint32 `idl:"name:ulCrossSiteSyncFlags" json:"cross_site_sync_flags"`
	// ulResolvePeerBackoffMinutes:  The ResolvePeerBackoffMinutes element value (see [MS-SNTP]
	// section 3.1.1).
	ResolvePeerBackoffMinutes uint32 `idl:"name:ulResolvePeerBackoffMinutes" json:"resolve_peer_backoff_minutes"`
	// ulResolvePeerBackoffMaxTimes:  The ResolvePeerBackoffMaxTimes element value (see
	// [MS-SNTP] section 3.1.1).
	ResolvePeerBackoffMaxTimes uint32 `idl:"name:ulResolvePeerBackoffMaxTimes" json:"resolve_peer_backoff_max_times"`
	// ulCompatibilityFlags:  The CompatibilityFlags element value (see section 3.2.1.2.1).
	CompatibilityFlags uint32 `idl:"name:ulCompatibilityFlags" json:"compatibility_flags"`
	// ulEventLogFlags:  The NTPEventLogFlags element value (see section 3.2.1.2.1).
	EventLogFlags uint32 `idl:"name:ulEventLogFlags" json:"event_log_flags"`
	// ulLargeSampleSkew:  The LargeSampleSkew element value (see section 3.2.1.2.1).
	LargeSampleSkew uint32 `idl:"name:ulLargeSampleSkew" json:"large_sample_skew"`
	// ulSpecialPollInterval:  An integer that indicates a special poll interval, in seconds,
	// for manual time synchronization.
	SpecialPollInterval uint32 `idl:"name:ulSpecialPollInterval" json:"special_poll_interval"`
	// wszType:  A case-insensitive, null-terminated string that indicates the time synchronization
	// behavior of the time service. The string MUST have one of the allowable values listed
	// for the TimeSourceType Abstract Data Model variable described in [MS-SNTP] section
	// 3.1.1.
	Type string `idl:"name:wszType;string;pointer:unique" json:"type"`
	// wszNtpServer:  A case-insensitive, null-terminated string that indicates a space-delimited
	// list of time sources that the time service can synchronize with. Each time source
	// MUST be in the following form.
	//
	// <Time Source>[,<Bitwise Flag>]
	//
	// The "Time Source" MUST be in the form of a fully qualified domain name (FQDN) or
	// an IP address.
	//
	// The "Bitwise Flag", if included, MUST be a bitwise OR of zero or more of the following
	// flags.
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|                        |                                                                                  |
	//	|         VALUE          |                                     MEANING                                      |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| SpecialInterval 0x01   | The time service uses the polling interval for this time source, as defined by   |
	//	|                        | the value of the ulSpecialPollInterval member.                                   |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| UseAsFallbackOnly 0x02 | The time service uses this time source only when all other time sources have     |
	//	|                        | failed.                                                                          |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| SymmetricActive 0x04   | The time service uses the symmetric active mode when communicating with this     |
	//	|                        | time source.                                                                     |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| Client 0x08            | The time service uses the client mode when communicating with this time source.  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//
	// Multiple time sources are delineated by a space. For two time sources, the following
	// form would be used.
	//
	// <Time Source #1>,<Bitwise Flags #1> <Time Source #2>,<Bitwise Flags #2>
	NTPServer string `idl:"name:wszNtpServer;string;pointer:unique" json:"ntp_server"`
	// ulAllowNonstandardModeCombinationsFlag:  An integer that indicates the source of
	// the configuration setting for ulAllowNonstandardModeCombinations, as specified in
	// section 2.2.6.
	AllowNonstandardModeCombinationsFlag uint32 `idl:"name:ulAllowNonstandardModeCombinationsFlag" json:"allow_nonstandard_mode_combinations_flag"`
	// ulCrossSiteSyncFlagsFlag:  An integer that indicates the source of the configuration
	// setting for ulCrossSiteSyncFlags, as specified in section 2.2.6.
	CrossSiteSyncFlagsFlag uint32 `idl:"name:ulCrossSiteSyncFlagsFlag" json:"cross_site_sync_flags_flag"`
	// ulResolvePeerBackoffMinutesFlag:  An integer that indicates the source of the configuration
	// setting for ulResolvePeerBackoffMinutes, as specified in section 2.2.6.
	ResolvePeerBackoffMinutesFlag uint32 `idl:"name:ulResolvePeerBackoffMinutesFlag" json:"resolve_peer_backoff_minutes_flag"`
	// ulResolvePeerBackoffMaxTimesFlag:  An integer that indicates the source of the configuration
	// setting for ulResolvePeerBackoffMaxTimes, as specified in section 2.2.6.
	ResolvePeerBackoffMaxTimesFlag uint32 `idl:"name:ulResolvePeerBackoffMaxTimesFlag" json:"resolve_peer_backoff_max_times_flag"`
	// ulCompatibilityFlagsFlag:  An integer that indicates the source of the configuration
	// setting for ulCompatibilityFlags, as specified in section 2.2.6.
	CompatibilityFlagsFlag uint32 `idl:"name:ulCompatibilityFlagsFlag" json:"compatibility_flags_flag"`
	// ulEventLogFlagsFlag:  An integer that indicates the source of the configuration setting
	// for ulEventLogFlags, as specified in section 2.2.6.
	EventLogFlagsFlag uint32 `idl:"name:ulEventLogFlagsFlag" json:"event_log_flags_flag"`
	// ulLargeSampleSkewFlag:  An integer that indicates the source of the configuration
	// setting for ulLargeSampleSkew, as specified in section 2.2.6.
	LargeSampleSkewFlag uint32 `idl:"name:ulLargeSampleSkewFlag" json:"large_sample_skew_flag"`
	// ulSpecialPollIntervalFlag:  An integer that indicates the source of the configuration
	// setting for ulSpecialPollInterval, as specified in section 2.2.6.
	SpecialPollIntervalFlag uint32 `idl:"name:ulSpecialPollIntervalFlag" json:"special_poll_interval_flag"`
	// ulTypeFlag:  An integer that indicates the source of the configuration setting for
	// wszType, as specified in section 2.2.6.
	TypeFlag uint32 `idl:"name:ulTypeFlag" json:"type_flag"`
	// ulNtpServerFlag:  An integer that indicates the source of the configuration setting
	// for wszNtpServer, as specified in section 2.2.6.
	NTPServerFlag uint32 `idl:"name:ulNtpServerFlag" json:"ntp_server_flag"`
	// cEntries:  An integer that indicates the number of additional configuration entries
	// in pEntries.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// pEntries:  A pointer to W32TIME_ENTRY structures that represent additional configuration
	// entries.
	Entries []*Entry `idl:"name:pEntries;size_is:(cEntries)" json:"entries"`
}

func (o *NTPClientProviderConfigData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Entries != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.Entries))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTPClientProviderConfigData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowNonstandardModeCombinations); err != nil {
		return err
	}
	if err := w.WriteData(o.CrossSiteSyncFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.ResolvePeerBackoffMinutes); err != nil {
		return err
	}
	if err := w.WriteData(o.ResolvePeerBackoffMaxTimes); err != nil {
		return err
	}
	if err := w.WriteData(o.CompatibilityFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.LargeSampleSkew); err != nil {
		return err
	}
	if err := w.WriteData(o.SpecialPollInterval); err != nil {
		return err
	}
	if o.Type != "" {
		_ptr_wszType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Type); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Type, _ptr_wszType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.NTPServer != "" {
		_ptr_wszNtpServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.NTPServer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.NTPServer, _ptr_wszNtpServer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AllowNonstandardModeCombinationsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.CrossSiteSyncFlagsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.ResolvePeerBackoffMinutesFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.ResolvePeerBackoffMaxTimesFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.CompatibilityFlagsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogFlagsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.LargeSampleSkewFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.SpecialPollIntervalFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.NTPServerFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	if o.Entries != nil || o.EntriesCount > 0 {
		_ptr_pEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Entries {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Entries[i1] != nil {
					if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Entries, _ptr_pEntries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTPClientProviderConfigData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowNonstandardModeCombinations); err != nil {
		return err
	}
	if err := w.ReadData(&o.CrossSiteSyncFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResolvePeerBackoffMinutes); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResolvePeerBackoffMaxTimes); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompatibilityFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.LargeSampleSkew); err != nil {
		return err
	}
	if err := w.ReadData(&o.SpecialPollInterval); err != nil {
		return err
	}
	_ptr_wszType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Type); err != nil {
			return err
		}
		return nil
	})
	_s_wszType := func(ptr interface{}) { o.Type = *ptr.(*string) }
	if err := w.ReadPointer(&o.Type, _s_wszType, _ptr_wszType); err != nil {
		return err
	}
	_ptr_wszNtpServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.NTPServer); err != nil {
			return err
		}
		return nil
	})
	_s_wszNtpServer := func(ptr interface{}) { o.NTPServer = *ptr.(*string) }
	if err := w.ReadPointer(&o.NTPServer, _s_wszNtpServer, _ptr_wszNtpServer); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowNonstandardModeCombinationsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.CrossSiteSyncFlagsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResolvePeerBackoffMinutesFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResolvePeerBackoffMaxTimesFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompatibilityFlagsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogFlagsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.LargeSampleSkewFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.SpecialPollIntervalFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.NTPServerFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	_ptr_pEntries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*Entry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &Entry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEntries := func(ptr interface{}) { o.Entries = *ptr.(*[]*Entry) }
	if err := w.ReadPointer(&o.Entries, _s_pEntries, _ptr_pEntries); err != nil {
		return err
	}
	return nil
}

// NTPServerProviderConfigData structure represents W32TIME_NTPSERVER_PROVIDER_CONFIG_DATA RPC structure.
//
// The W32TIME_NTPSERVER_PROVIDER_CONFIG_DATA structure contains configuration data
// about an NtpServer time provider.
//
// The structure is defined to match the NtpServer time provider's configuration of
// the W32Time implementation as described in [WTSREF]. Fields in the structure that
// do not apply to other implementations SHOULD<11> have their corresponding configuration-setting
// type fields set to W32TIME_CONFIGURATION_SETTING_UNDEFINED.
type NTPServerProviderConfigData struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulAllowNonstandardModeCombinations:  An integer that indicates whether or not nonstandard
	// mode combinations are allowed.
	//
	//	+-------+---------------------------------------------------------------+
	//	|       |                                                               |
	//	| VALUE |                            MEANING                            |
	//	|       |                                                               |
	//	+-------+---------------------------------------------------------------+
	//	+-------+---------------------------------------------------------------+
	//	|     0 | Indicates that nonstandard mode combinations are not allowed. |
	//	+-------+---------------------------------------------------------------+
	//	|     1 | Indicates that nonstandard mode combinations are allowed.     |
	//	+-------+---------------------------------------------------------------+
	AllowNonstandardModeCombinations uint32 `idl:"name:ulAllowNonstandardModeCombinations" json:"allow_nonstandard_mode_combinations"`
	// ulAllowNonstandardModeCombinationsFlag:  An integer that indicates the source of
	// the configuration setting for ulAllowNonstandardModeCombinations, as specified in
	// section 2.2.6.
	AllowNonstandardModeCombinationsFlag uint32 `idl:"name:ulAllowNonstandardModeCombinationsFlag" json:"allow_nonstandard_mode_combinations_flag"`
	// ulEventLogFlags:  An integer that indicates the combination of flags that determines
	// how the time provider logs events into an event log. This MUST be the following value.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| NCELF_LogServerResponseError 0x00000008 | Log an event when the time provider fails to validate a request for              |
	//	|                                         | authenticated time synchronization.                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	EventLogFlags uint32 `idl:"name:ulEventLogFlags" json:"event_log_flags"`
	// ulEventLogFlagsFlag:  An integer that indicates the source of the configuration setting
	// for ulEventLogFlags, as specified in section 2.2.6.
	EventLogFlagsFlag uint32 `idl:"name:ulEventLogFlagsFlag" json:"event_log_flags_flag"`
	// cEntries:  An integer that indicates the number of additional configuration entries
	// in pEntries.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// pEntries:  A pointer to W32TIME_ENTRY structures that represent additional configuration
	// entries.
	Entries []*Entry `idl:"name:pEntries;size_is:(cEntries)" json:"entries"`
}

func (o *NTPServerProviderConfigData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Entries != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.Entries))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTPServerProviderConfigData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowNonstandardModeCombinations); err != nil {
		return err
	}
	if err := w.WriteData(o.AllowNonstandardModeCombinationsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogFlagsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	if o.Entries != nil || o.EntriesCount > 0 {
		_ptr_pEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Entries {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Entries[i1] != nil {
					if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Entries, _ptr_pEntries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTPServerProviderConfigData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowNonstandardModeCombinations); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowNonstandardModeCombinationsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogFlagsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	_ptr_pEntries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*Entry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &Entry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEntries := func(ptr interface{}) { o.Entries = *ptr.(*[]*Entry) }
	if err := w.ReadPointer(&o.Entries, _s_pEntries, _ptr_pEntries); err != nil {
		return err
	}
	return nil
}

// ProviderConfigData structure represents W32TIME_PROVIDER_CONFIG_DATA RPC union.
type ProviderConfigData struct {
	// Types that are assignable to Value
	//
	// *ProviderConfigData_NTPClient
	// *ProviderConfigData_NTPServer
	Value is_ProviderConfigData `json:"value"`
}

func (o *ProviderConfigData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ProviderConfigData_NTPClient:
		if value != nil {
			return value.NTPClientProviderConfigData
		}
	case *ProviderConfigData_NTPServer:
		if value != nil {
			return value.NTPServerProviderConfigData
		}
	}
	return nil
}

type is_ProviderConfigData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ProviderConfigData()
}

func (o *ProviderConfigData) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ProviderConfigData_NTPClient:
		return uint32(0)
	case *ProviderConfigData_NTPServer:
		return uint32(1)
	}
	return uint32(0)
}

func (o *ProviderConfigData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
	case uint32(0):
		_o, _ := o.Value.(*ProviderConfigData_NTPClient)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProviderConfigData_NTPClient{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*ProviderConfigData_NTPServer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProviderConfigData_NTPServer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ProviderConfigData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
	case uint32(0):
		o.Value = &ProviderConfigData_NTPClient{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &ProviderConfigData_NTPServer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ProviderConfigData_NTPClient structure represents W32TIME_PROVIDER_CONFIG_DATA RPC union arm.
//
// It has following labels: 0
type ProviderConfigData_NTPClient struct {
	NTPClientProviderConfigData *NTPClientProviderConfigData `idl:"name:pNtpClientProviderConfigData" json:"ntp_client_provider_config_data"`
}

func (*ProviderConfigData_NTPClient) is_ProviderConfigData() {}

func (o *ProviderConfigData_NTPClient) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.NTPClientProviderConfigData != nil {
		_ptr_pNtpClientProviderConfigData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NTPClientProviderConfigData != nil {
				if err := o.NTPClientProviderConfigData.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NTPClientProviderConfigData{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NTPClientProviderConfigData, _ptr_pNtpClientProviderConfigData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProviderConfigData_NTPClient) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pNtpClientProviderConfigData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NTPClientProviderConfigData == nil {
			o.NTPClientProviderConfigData = &NTPClientProviderConfigData{}
		}
		if err := o.NTPClientProviderConfigData.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNtpClientProviderConfigData := func(ptr interface{}) { o.NTPClientProviderConfigData = *ptr.(**NTPClientProviderConfigData) }
	if err := w.ReadPointer(&o.NTPClientProviderConfigData, _s_pNtpClientProviderConfigData, _ptr_pNtpClientProviderConfigData); err != nil {
		return err
	}
	return nil
}

// ProviderConfigData_NTPServer structure represents W32TIME_PROVIDER_CONFIG_DATA RPC union arm.
//
// It has following labels: 1
type ProviderConfigData_NTPServer struct {
	NTPServerProviderConfigData *NTPServerProviderConfigData `idl:"name:pNtpServerProviderConfigData" json:"ntp_server_provider_config_data"`
}

func (*ProviderConfigData_NTPServer) is_ProviderConfigData() {}

func (o *ProviderConfigData_NTPServer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.NTPServerProviderConfigData != nil {
		_ptr_pNtpServerProviderConfigData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NTPServerProviderConfigData != nil {
				if err := o.NTPServerProviderConfigData.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NTPServerProviderConfigData{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NTPServerProviderConfigData, _ptr_pNtpServerProviderConfigData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProviderConfigData_NTPServer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pNtpServerProviderConfigData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NTPServerProviderConfigData == nil {
			o.NTPServerProviderConfigData = &NTPServerProviderConfigData{}
		}
		if err := o.NTPServerProviderConfigData.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNtpServerProviderConfigData := func(ptr interface{}) { o.NTPServerProviderConfigData = *ptr.(**NTPServerProviderConfigData) }
	if err := w.ReadPointer(&o.NTPServerProviderConfigData, _s_pNtpServerProviderConfigData, _ptr_pNtpServerProviderConfigData); err != nil {
		return err
	}
	return nil
}

// ProviderConfig structure represents W32TIME_PROVIDER_CONFIG RPC structure.
//
// The W32TIME_PROVIDER_CONFIG structure defines configuration data for a selected time
// provider.
type ProviderConfig struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulProviderType:  The type of time provider, which MUST be one of the following values.
	//
	//	+---------------------------------------------------+-----------------------------+
	//	|                                                   |                             |
	//	|                       VALUE                       |           MEANING           |
	//	|                                                   |                             |
	//	+---------------------------------------------------+-----------------------------+
	//	+---------------------------------------------------+-----------------------------+
	//	| W32TIME_NTPCLIENT_PROVIDER_CONFIG_DATA 0x00000000 | NtpClient NTP time provider |
	//	+---------------------------------------------------+-----------------------------+
	//	| W32TIME_NTPSERVER_PROVIDER_CONFIG_DATA 0x00000001 | NtpServer NTP time provider |
	//	+---------------------------------------------------+-----------------------------+
	ProviderType uint32 `idl:"name:ulProviderType" json:"provider_type"`
	// pProviderConfigData:  A W32TIME_PROVIDER_CONFIG_DATA union that contains configuration
	// data about the time provider.
	ProviderConfigData *ProviderConfigData `idl:"name:pProviderConfigData;switch_is:ulProviderType" json:"provider_config_data"`
}

func (o *ProviderConfig) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ProviderConfig) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.ProviderType); err != nil {
		return err
	}
	if o.ProviderConfigData != nil {
		_ptr_pProviderConfigData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swProviderConfigData := uint32(o.ProviderType)
			if o.ProviderConfigData != nil {
				if err := o.ProviderConfigData.MarshalUnionNDR(ctx, w, _swProviderConfigData); err != nil {
					return err
				}
			} else {
				if err := (&ProviderConfigData{}).MarshalUnionNDR(ctx, w, _swProviderConfigData); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderConfigData, _ptr_pProviderConfigData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProviderConfig) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProviderType); err != nil {
		return err
	}
	_ptr_pProviderConfigData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ProviderConfigData == nil {
			o.ProviderConfigData = &ProviderConfigData{}
		}
		_swProviderConfigData := uint32(o.ProviderType)
		if err := o.ProviderConfigData.UnmarshalUnionNDR(ctx, w, _swProviderConfigData); err != nil {
			return err
		}
		return nil
	})
	_s_pProviderConfigData := func(ptr interface{}) { o.ProviderConfigData = *ptr.(**ProviderConfigData) }
	if err := w.ReadPointer(&o.ProviderConfigData, _s_pProviderConfigData, _ptr_pProviderConfigData); err != nil {
		return err
	}
	return nil
}

// ConfigurationBasic structure represents W32TIME_CONFIGURATION_BASIC RPC structure.
//
// The W32TIME_CONFIGURATION_BASIC structure defines the basic configuration data of
// the time service.
//
// The structure is defined to match the basic configuration of the W32Time implementation,
// as described in [WTSREF]. Fields in the structure that are not valid in other implementations
// SHOULD have their corresponding configuration-setting type fields set to W32TIME_CONFIGURATION_SETTING_UNDEFINED.<12>
type ConfigurationBasic struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulEventLogFlags:  The EventLogFlags element value (see section 3.2.1.1).
	EventLogFlags uint32 `idl:"name:ulEventLogFlags" json:"event_log_flags"`
	// ulAnnounceFlags:  An integer that indicates the combination of flags that determines
	// how the time service advertises itself as a time server. The value MUST be a bitwise
	// OR of zero or more of the following flags.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| Timeserv_Announce_No 0x00000000            | Not a time server.                                                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| Timeserv_Announce_Yes 0x00000001           | Always advertised as a time server.                                              |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| Timeserv_Announce_Auto 0x00000002          | Advertising as a time server is decided automatically: only when the server is   |
	//	|                                            | synchronized.                                                                    |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| Reliable_Timeserv_Announce_Yes 0x00000004  | Always advertised as a reliable time server.                                     |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| Reliable_Timeserv_Announce_Auto 0x00000008 | Advertising as a time server is decided automatically: only when the server is   |
	//	|                                            | synchronized and is a reliable time server.                                      |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// All other values are reserved for future use.
	AnnounceFlags uint32 `idl:"name:ulAnnounceFlags" json:"announce_flags"`
	// ulTimeJumpAuditOffset:  The TimeJumpAuditOffset element value (see section 3.2.1.1).
	TimeJumpAuditOffset uint32 `idl:"name:ulTimeJumpAuditOffset" json:"time_jump_audit_offset"`
	// ulMinPollInterval:  An integer that indicates the minimum poll interval of domain
	// time synchronization, expressed as specified in [RFC1305] section 3.2.7, using units
	// of seconds given as exponents to a power of two.
	MinPollInterval uint32 `idl:"name:ulMinPollInterval" json:"min_poll_interval"`
	// ulMaxPollInterval:  An integer that indicates the maximum poll interval of domain
	// time synchronization, expressed as specified in [RFC1305] section 3.2.7, using units
	// of seconds given as exponents to a power of two.
	MaxPollInterval uint32 `idl:"name:ulMaxPollInterval" json:"max_poll_interval"`
	// ulMaxNegPhaseCorrection:  The MaxNegPhaseCorrection element value (see section 3.2.1.1).
	MaxNegPhaseCorrection uint32 `idl:"name:ulMaxNegPhaseCorrection" json:"max_neg_phase_correction"`
	// ulMaxPosPhaseCorrection:  The MaxPosPhaseCorrection element value (see section 3.2.1.1).
	MaxPosPhaseCorrection uint32 `idl:"name:ulMaxPosPhaseCorrection" json:"max_pos_phase_correction"`
	// ulMaxAllowedPhaseOffset:  The MaxAllowedPhaseOffset element value (see section 3.2.1.1).
	MaxAllowedPhaseOffset uint32 `idl:"name:ulMaxAllowedPhaseOffset" json:"max_allowed_phase_offset"`
	// ulEventLogFlagsFlag:  An integer that indicates the source of the configuration setting
	// for ulEventLogFlags, as specified in section 2.2.6.
	EventLogFlagsFlag uint32 `idl:"name:ulEventLogFlagsFlag" json:"event_log_flags_flag"`
	// ulAnnounceFlagsFlag:  An integer that indicates the source of the configuration setting
	// for ulAnnounceFlags, as specified in section 2.2.6.
	AnnounceFlagsFlag uint32 `idl:"name:ulAnnounceFlagsFlag" json:"announce_flags_flag"`
	// ulTimeJumpAuditOffsetFlag:  An integer that indicates the source of the configuration
	// setting for ulTimeJumpAuditOffset, as specified in section 2.2.6.
	TimeJumpAuditOffsetFlag uint32 `idl:"name:ulTimeJumpAuditOffsetFlag" json:"time_jump_audit_offset_flag"`
	// ulMinPollIntervalFlag:  An integer that indicates the source of the configuration
	// setting for ulMinPollInterval, as specified in section 2.2.6.
	MinPollIntervalFlag uint32 `idl:"name:ulMinPollIntervalFlag" json:"min_poll_interval_flag"`
	// ulMaxPollIntervalFlag:  An integer that indicates the source of the configuration
	// setting for ulMaxPollInterval, as specified in section 2.2.6.
	MaxPollIntervalFlag uint32 `idl:"name:ulMaxPollIntervalFlag" json:"max_poll_interval_flag"`
	// ulMaxNegPhaseCorrectionFlag:  An integer that indicates the source of the configuration
	// setting for ulMaxNegPhaseCorrection, as specified in section 2.2.6.
	MaxNegPhaseCorrectionFlag uint32 `idl:"name:ulMaxNegPhaseCorrectionFlag" json:"max_neg_phase_correction_flag"`
	// ulMaxPosPhaseCorrectionFlag:  An integer that indicates the source of the configuration
	// setting for ulMaxPosPhaseCorrection, as specified in section 2.2.6.
	MaxPosPhaseCorrectionFlag uint32 `idl:"name:ulMaxPosPhaseCorrectionFlag" json:"max_pos_phase_correction_flag"`
	// ulMaxAllowedPhaseOffsetFlag:  An integer that indicates the source of the configuration
	// setting for ulMaxAllowedPhaseOffset, as specified in section 2.2.6.
	MaxAllowedPhaseOffsetFlag uint32 `idl:"name:ulMaxAllowedPhaseOffsetFlag" json:"max_allowed_phase_offset_flag"`
}

func (o *ConfigurationBasic) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationBasic) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.AnnounceFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeJumpAuditOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.MinPollInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxPollInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxNegPhaseCorrection); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxPosPhaseCorrection); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxAllowedPhaseOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.EventLogFlagsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.AnnounceFlagsFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeJumpAuditOffsetFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.MinPollIntervalFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxPollIntervalFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxNegPhaseCorrectionFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxPosPhaseCorrectionFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxAllowedPhaseOffsetFlag); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationBasic) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.AnnounceFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeJumpAuditOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinPollInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxPollInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxNegPhaseCorrection); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxPosPhaseCorrection); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxAllowedPhaseOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.EventLogFlagsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.AnnounceFlagsFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeJumpAuditOffsetFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinPollIntervalFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxPollIntervalFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxNegPhaseCorrectionFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxPosPhaseCorrectionFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxAllowedPhaseOffsetFlag); err != nil {
		return err
	}
	return nil
}

// ConfigurationAdvanced structure represents W32TIME_CONFIGURATION_ADVANCED RPC structure.
//
// The W32TIME_CONFIGURATION_ADVANCED structure defines the advanced configuration data
// of the time service.<13>
//
// The structure is defined to match the advanced configuration of the W32Time implementation.
// Fields in the structure that are not valid in other implementations SHOULD have their
// corresponding configuration-setting type fields set to W32TIME_CONFIGURATION_SETTING_UNDEFINED.
// For more information on W32Time, see [WTSREF].
type ConfigurationAdvanced struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulFrequencyCorrectRate:  The FrequencyCorrectRate element value (see section 3.2.1.1).
	FrequencyCorrectRate uint32 `idl:"name:ulFrequencyCorrectRate" json:"frequency_correct_rate"`
	// ulPollAdjustFactor:  The PollAdjustFactor element value (see section 3.2.1.1).
	PollAdjustFactor uint32 `idl:"name:ulPollAdjustFactor" json:"poll_adjust_factor"`
	// ulLargePhaseOffset:  An integer that indicates the threshold that determines whether
	// or not a time sample indicates a spike, in 100-nanosecond units. If the time difference
	// of the time sample is more than the value, the sample indicates a possible spike,
	// in which case the time service changes its state, as specified in section 2.2.7.
	// For more information on spike detection, see [NTP-TR9733i] and [NTP-TR9733] section
	// 4.
	LargePhaseOffset uint32 `idl:"name:ulLargePhaseOffset" json:"large_phase_offset"`
	// ulSpikeWatchPeriod:  An integer that indicates the time interval, in seconds, that
	// determines how long the time service watches a spike condition. If time samples constantly
	// indicate spikes in this time interval, the time service becomes unsynchronized, in
	// which case the time service MUST change its state, as described in section 2.2.7.
	// For more information on spike detection, see [NTP-TR9733i] and [NTP-TR9733] section
	// 4.3.
	SpikeWatchPeriod uint32 `idl:"name:ulSpikeWatchPeriod" json:"spike_watch_period"`
	// ulLocalClockDispersion:  An integer that indicates the local clock dispersion, in
	// seconds. The root dispersion is set to this value if the time service runs as a primary
	// server, or if the root dispersion is invalid in the received response. For details
	// on dispersion and root dispersion, see [RFC1305] section 3.2.
	LocalClockDispersion uint32 `idl:"name:ulLocalClockDispersion" json:"local_clock_dispersion"`
	// ulHoldPeriod:  An integer that indicates the number of time samples during which
	// the spike detection is disabled when the time service is in the HOLD state, as specified
	// in section 2.2.7. For more information on the HOLD state, see [NTP-TR9733i] and [NTP-TR9733]
	// section 4.3.
	HoldPeriod uint32 `idl:"name:ulHoldPeriod" json:"hold_period"`
	// ulPhaseCorrectRate:  The PhaseCorrectRate element value (see section 3.2.1.1).
	PhaseCorrectRate uint32 `idl:"name:ulPhaseCorrectRate" json:"phase_correct_rate"`
	// ulUpdateInterval:  The UpdateInterval element value (see section 3.2.1.1).
	UpdateInterval uint32 `idl:"name:ulUpdateInterval" json:"update_interval"`
	// ulFrequencyCorrectRateFlag:  An integer that indicates the source of the configuration
	// setting for ulFrequencyCorrectRate, as specified in section 2.2.6.
	FrequencyCorrectRateFlag uint32 `idl:"name:ulFrequencyCorrectRateFlag" json:"frequency_correct_rate_flag"`
	// ulPollAdjustFactorFlag:  An integer that indicates the source of the configuration
	// setting for ulPollAdjustFactor, as specified in section 2.2.6.
	PollAdjustFactorFlag uint32 `idl:"name:ulPollAdjustFactorFlag" json:"poll_adjust_factor_flag"`
	// ulLargePhaseOffsetFlag:  An integer that indicates the source of the configuration
	// setting for ulLargePhaseOffset, as specified in section 2.2.6.
	LargePhaseOffsetFlag uint32 `idl:"name:ulLargePhaseOffsetFlag" json:"large_phase_offset_flag"`
	// ulSpikeWatchPeriodFlag:  An integer that indicates the source of the configuration
	// setting for ulSpikeWatchPeriod, as specified in section 2.2.6.
	SpikeWatchPeriodFlag uint32 `idl:"name:ulSpikeWatchPeriodFlag" json:"spike_watch_period_flag"`
	// ulLocalClockDispersionFlag:  An integer that indicates the source of the configuration
	// setting for ulLocalClockDispersion, as specified in section 2.2.6.
	LocalClockDispersionFlag uint32 `idl:"name:ulLocalClockDispersionFlag" json:"local_clock_dispersion_flag"`
	// ulHoldPeriodFlag:  An integer that indicates the source of the configuration setting
	// for ulHoldPeriod, as specified in section 2.2.6.
	HoldPeriodFlag uint32 `idl:"name:ulHoldPeriodFlag" json:"hold_period_flag"`
	// ulPhaseCorrectRateFlag:  An integer that indicates the source of the configuration
	// setting for ulPhaseCorrectRate, as specified in section 2.2.6.
	PhaseCorrectRateFlag uint32 `idl:"name:ulPhaseCorrectRateFlag" json:"phase_correct_rate_flag"`
	// ulUpdateIntervalFlag:  An integer that indicates the source of the configuration
	// setting for ulUpdateInterval, as specified in section 2.2.6.
	UpdateIntervalFlag uint32 `idl:"name:ulUpdateIntervalFlag" json:"update_interval_flag"`
}

func (o *ConfigurationAdvanced) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationAdvanced) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.FrequencyCorrectRate); err != nil {
		return err
	}
	if err := w.WriteData(o.PollAdjustFactor); err != nil {
		return err
	}
	if err := w.WriteData(o.LargePhaseOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.SpikeWatchPeriod); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalClockDispersion); err != nil {
		return err
	}
	if err := w.WriteData(o.HoldPeriod); err != nil {
		return err
	}
	if err := w.WriteData(o.PhaseCorrectRate); err != nil {
		return err
	}
	if err := w.WriteData(o.UpdateInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.FrequencyCorrectRateFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.PollAdjustFactorFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.LargePhaseOffsetFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.SpikeWatchPeriodFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalClockDispersionFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.HoldPeriodFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.PhaseCorrectRateFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.UpdateIntervalFlag); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationAdvanced) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.FrequencyCorrectRate); err != nil {
		return err
	}
	if err := w.ReadData(&o.PollAdjustFactor); err != nil {
		return err
	}
	if err := w.ReadData(&o.LargePhaseOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.SpikeWatchPeriod); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalClockDispersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.HoldPeriod); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhaseCorrectRate); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpdateInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.FrequencyCorrectRateFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.PollAdjustFactorFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.LargePhaseOffsetFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.SpikeWatchPeriodFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalClockDispersionFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.HoldPeriodFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.PhaseCorrectRateFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpdateIntervalFlag); err != nil {
		return err
	}
	return nil
}

// ConfigurationDefault structure represents W32TIME_CONFIGURATION_DEFAULT RPC structure.
//
// The W32TIME_CONFIGURATION_DEFAULT structure defines the default configuration data
// of the time service as described in [WTSREF].
//
// The structure is defined to match the default configuration of the W32Time implementation.
// Fields in the structure that are not valid in other implementations SHOULD have their
// corresponding configuration-setting type fields set to W32TIME_CONFIGURATION_SETTING_UNDEFINED.<14>
type ConfigurationDefault struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// wszFileLogName:  The FileLogName element value (see section 3.2.1.1).
	FileLogName string `idl:"name:wszFileLogName;string;pointer:unique" json:"file_log_name"`
	// wszFileLogEntries:  The FileLogEntries element value (see section 3.2.1.1).
	FileLogEntries string `idl:"name:wszFileLogEntries;string;pointer:unique" json:"file_log_entries"`
	// ulFileLogSize:  The FileLogSize element value (see section 3.2.1.1).
	FileLogSize uint32 `idl:"name:ulFileLogSize" json:"file_log_size"`
	// ulFileLogFlags:  The FileLogFlags element value (see section 3.2.1.1).
	FileLogFlags uint32 `idl:"name:ulFileLogFlags" json:"file_log_flags"`
	// ulFileLogNameFlag:  An integer that indicates the source of the configuration setting
	// for wszFileLogName, as specified in section 2.2.6.
	FileLogNameFlag uint32 `idl:"name:ulFileLogNameFlag" json:"file_log_name_flag"`
	// ulFileLogEntriesFlag:  An integer that indicates the source of the configuration
	// setting for wszFileLogEntries, as specified in section 2.2.6.
	FileLogEntriesFlag uint32 `idl:"name:ulFileLogEntriesFlag" json:"file_log_entries_flag"`
	// ulFileLogSizeFlag:  An integer that indicates the source of the configuration setting
	// for ulFileLogSize, as specified in section 2.2.6.
	FileLogSizeFlag uint32 `idl:"name:ulFileLogSizeFlag" json:"file_log_size_flag"`
	// ulFileLogFlagsFlag:  An integer that indicates the source of the configuration setting
	// for ulFileLogFlags, as specified in section 2.2.6.
	FileLogFlagsFlag uint32 `idl:"name:ulFileLogFlagsFlag" json:"file_log_flags_flag"`
}

func (o *ConfigurationDefault) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationDefault) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.FileLogName != "" {
		_ptr_wszFileLogName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FileLogName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FileLogName, _ptr_wszFileLogName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.FileLogEntries != "" {
		_ptr_wszFileLogEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FileLogEntries); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FileLogEntries, _ptr_wszFileLogEntries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FileLogSize); err != nil {
		return err
	}
	if err := w.WriteData(o.FileLogFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.FileLogNameFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.FileLogEntriesFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.FileLogSizeFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.FileLogFlagsFlag); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationDefault) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_wszFileLogName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileLogName); err != nil {
			return err
		}
		return nil
	})
	_s_wszFileLogName := func(ptr interface{}) { o.FileLogName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FileLogName, _s_wszFileLogName, _ptr_wszFileLogName); err != nil {
		return err
	}
	_ptr_wszFileLogEntries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileLogEntries); err != nil {
			return err
		}
		return nil
	})
	_s_wszFileLogEntries := func(ptr interface{}) { o.FileLogEntries = *ptr.(*string) }
	if err := w.ReadPointer(&o.FileLogEntries, _s_wszFileLogEntries, _ptr_wszFileLogEntries); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileLogSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileLogFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileLogNameFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileLogEntriesFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileLogSizeFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileLogFlagsFlag); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ConfigurationProvider structure represents W32TIME_CONFIGURATION_PROVIDER RPC structure.
//
// The W32TIME_CONFIGURATION_PROVIDER structure defines the configuration data of an
// NTP time provider.
type ConfigurationProvider struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// ulInputProvider:  An integer that indicates whether the provider is capable of retrieving
	// time samples.
	//
	//	+-------+------------------------------------------------------------------------+
	//	|       |                                                                        |
	//	| VALUE |                                MEANING                                 |
	//	|       |                                                                        |
	//	+-------+------------------------------------------------------------------------+
	//	+-------+------------------------------------------------------------------------+
	//	|     0 | Indicates that the provider is not capable of retrieving time samples. |
	//	+-------+------------------------------------------------------------------------+
	//	|     1 | Indicates that the provider is capable of retrieving time samples.     |
	//	+-------+------------------------------------------------------------------------+
	InputProvider uint32 `idl:"name:ulInputProvider" json:"input_provider"`
	// ulEnabled:  An integer that indicates whether or not the provider is enabled.
	//
	//	+-------+------------------------------------------+
	//	|       |                                          |
	//	| VALUE |                 MEANING                  |
	//	|       |                                          |
	//	+-------+------------------------------------------+
	//	+-------+------------------------------------------+
	//	|     0 | Indicates that the provider is disabled. |
	//	+-------+------------------------------------------+
	//	|     1 | Indicates that the provider is enabled.  |
	//	+-------+------------------------------------------+
	Enabled uint32 `idl:"name:ulEnabled" json:"enabled"`
	// wszDllName:  The ProviderDllName element value (see section 3.2.1.2).
	DLLName string `idl:"name:wszDllName;string;pointer:unique" json:"dll_name"`
	// wszProviderName:  The ProviderName element value (see section 3.2.1.2)
	ProviderName string `idl:"name:wszProviderName;string;pointer:unique" json:"provider_name"`
	// ulDllNameFlag:  An integer indicating the source of the configuration setting for
	// wszDllName, as specified in section 2.2.6.
	DLLNameFlag uint32 `idl:"name:ulDllNameFlag" json:"dll_name_flag"`
	// ulProviderNameFlag:  An integer indicating the source of the configuration setting
	// for wszProviderName, as specified in section 2.2.6.
	ProviderNameFlag uint32 `idl:"name:ulProviderNameFlag" json:"provider_name_flag"`
	// ulInputProviderFlag:  An integer indicating the source of the configuration setting
	// for ulInputProvider, as specified in section 2.2.6.
	InputProviderFlag uint32 `idl:"name:ulInputProviderFlag" json:"input_provider_flag"`
	// ulEnabledFlag:  An integer indicating the source of the configuration setting for
	// ulEnabled, as specified in section 2.2.6.
	EnabledFlag uint32 `idl:"name:ulEnabledFlag" json:"enabled_flag"`
	// pProviderConfig:  A pointer to the W32TIME_PROVIDER_CONFIG structure.
	ProviderConfig *ProviderConfig `idl:"name:pProviderConfig" json:"provider_config"`
}

func (o *ConfigurationProvider) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationProvider) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.InputProvider); err != nil {
		return err
	}
	if err := w.WriteData(o.Enabled); err != nil {
		return err
	}
	if o.DLLName != "" {
		_ptr_wszDllName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DLLName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DLLName, _ptr_wszDllName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProviderName != "" {
		_ptr_wszProviderName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProviderName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderName, _ptr_wszProviderName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DLLNameFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.ProviderNameFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.InputProviderFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.EnabledFlag); err != nil {
		return err
	}
	if o.ProviderConfig != nil {
		_ptr_pProviderConfig := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ProviderConfig != nil {
				if err := o.ProviderConfig.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ProviderConfig{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderConfig, _ptr_pProviderConfig); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConfigurationProvider) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.InputProvider); err != nil {
		return err
	}
	if err := w.ReadData(&o.Enabled); err != nil {
		return err
	}
	_ptr_wszDllName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DLLName); err != nil {
			return err
		}
		return nil
	})
	_s_wszDllName := func(ptr interface{}) { o.DLLName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DLLName, _s_wszDllName, _ptr_wszDllName); err != nil {
		return err
	}
	_ptr_wszProviderName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProviderName); err != nil {
			return err
		}
		return nil
	})
	_s_wszProviderName := func(ptr interface{}) { o.ProviderName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProviderName, _s_wszProviderName, _ptr_wszProviderName); err != nil {
		return err
	}
	if err := w.ReadData(&o.DLLNameFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProviderNameFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.InputProviderFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.EnabledFlag); err != nil {
		return err
	}
	_ptr_pProviderConfig := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ProviderConfig == nil {
			o.ProviderConfig = &ProviderConfig{}
		}
		if err := o.ProviderConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pProviderConfig := func(ptr interface{}) { o.ProviderConfig = *ptr.(**ProviderConfig) }
	if err := w.ReadPointer(&o.ProviderConfig, _s_pProviderConfig, _ptr_pProviderConfig); err != nil {
		return err
	}
	return nil
}

// ConfigurationInfo structure represents W32TIME_CONFIGURATION_INFO RPC structure.
//
// The W32TIME_CONFIGURATION_INFO structure defines the configuration data of the time
// service.
type ConfigurationInfo struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// basicConfig:  The W32TIME_CONFIGURATION_BASIC structure that represents the basic
	// time service configuration data.
	BasicConfig *ConfigurationBasic `idl:"name:basicConfig" json:"basic_config"`
	// advancedConfig:  The W32TIME_CONFIGURATION_ADVANCED structure that represents the
	// advanced time service configuration data.
	AdvancedConfig *ConfigurationAdvanced `idl:"name:advancedConfig" json:"advanced_config"`
	// defaultConfig:  The W32TIME_CONFIGURATION_DEFAULT structure that represents the default
	// time service configuration data.
	DefaultConfig *ConfigurationDefault `idl:"name:defaultConfig" json:"default_config"`
	// cProviderConfig:  The number of time providers that are configured in the time service.
	// This value also indicates the number of structures in pProviderConfig.
	ProviderConfigCount uint32 `idl:"name:cProviderConfig" json:"provider_config_count"`
	// pProviderConfig:  An array of W32TIME_CONFIGURATION_PROVIDER structures that represent
	// the configuration data of time providers that are configured in the time service.
	ProviderConfig []*ConfigurationProvider `idl:"name:pProviderConfig;size_is:(cProviderConfig)" json:"provider_config"`
	// cEntries:  An integer that indicates the number of additional configuration entries
	// in pEntries.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// pEntries:  A pointer to W32TIME_ENTRY structures that represent additional configuration
	// entries.
	Entries []*Entry `idl:"name:pEntries;size_is:(cEntries)" json:"entries"`
}

func (o *ConfigurationInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ProviderConfig != nil && o.ProviderConfigCount == 0 {
		o.ProviderConfigCount = uint32(len(o.ProviderConfig))
	}
	if o.Entries != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.Entries))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ConfigurationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.BasicConfig != nil {
		if err := o.BasicConfig.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ConfigurationBasic{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AdvancedConfig != nil {
		if err := o.AdvancedConfig.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ConfigurationAdvanced{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DefaultConfig != nil {
		if err := o.DefaultConfig.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ConfigurationDefault{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProviderConfigCount); err != nil {
		return err
	}
	if o.ProviderConfig != nil || o.ProviderConfigCount > 0 {
		_ptr_pProviderConfig := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ProviderConfigCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ProviderConfig {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ProviderConfig[i1] != nil {
					_ptr_pProviderConfig := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.ProviderConfig[i1] != nil {
							if err := o.ProviderConfig[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ConfigurationProvider{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.ProviderConfig[i1], _ptr_pProviderConfig); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ProviderConfig); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderConfig, _ptr_pProviderConfig); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	if o.Entries != nil || o.EntriesCount > 0 {
		_ptr_pEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Entries {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Entries[i1] != nil {
					if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Entries, _ptr_pEntries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConfigurationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if o.BasicConfig == nil {
		o.BasicConfig = &ConfigurationBasic{}
	}
	if err := o.BasicConfig.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AdvancedConfig == nil {
		o.AdvancedConfig = &ConfigurationAdvanced{}
	}
	if err := o.AdvancedConfig.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DefaultConfig == nil {
		o.DefaultConfig = &ConfigurationDefault{}
	}
	if err := o.DefaultConfig.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProviderConfigCount); err != nil {
		return err
	}
	_ptr_pProviderConfig := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ProviderConfigCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ProviderConfigCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ProviderConfig", sizeInfo[0])
		}
		o.ProviderConfig = make([]*ConfigurationProvider, sizeInfo[0])
		for i1 := range o.ProviderConfig {
			i1 := i1
			_ptr_pProviderConfig := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ProviderConfig[i1] == nil {
					o.ProviderConfig[i1] = &ConfigurationProvider{}
				}
				if err := o.ProviderConfig[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_pProviderConfig := func(ptr interface{}) { o.ProviderConfig[i1] = *ptr.(**ConfigurationProvider) }
			if err := w.ReadPointer(&o.ProviderConfig[i1], _s_pProviderConfig, _ptr_pProviderConfig); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pProviderConfig := func(ptr interface{}) { o.ProviderConfig = *ptr.(*[]*ConfigurationProvider) }
	if err := w.ReadPointer(&o.ProviderConfig, _s_pProviderConfig, _ptr_pProviderConfig); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	_ptr_pEntries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*Entry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &Entry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEntries := func(ptr interface{}) { o.Entries = *ptr.(*[]*Entry) }
	if err := w.ReadPointer(&o.Entries, _s_pEntries, _ptr_pEntries); err != nil {
		return err
	}
	return nil
}

// StatusInfo structure represents W32TIME_STATUS_INFO RPC structure.
//
// The W32TIME_STATUS_INFO structure defines the current status data of the time service.
type StatusInfo struct {
	// ulSize:  The size, in bytes, of this structure.
	Size uint32 `idl:"name:ulSize" json:"size"`
	// eLeapIndicator:  An integer that warns of an impending leap second in the last minute
	// of the current day, as specified in [RFC1305] section 3.2.
	LeapIndicator uint32 `idl:"name:eLeapIndicator" json:"leap_indicator"`
	// nStratum:  An integer that indicates the stratum level of the local clock in the
	// time service, as specified in [RFC1305] section 3.2.
	Stratum uint32 `idl:"name:nStratum" json:"stratum"`
	// nPollInterval:  An integer that indicates the poll interval of the time service,
	// expressed as specified in [RFC1305] section 3.2, using units of seconds given as
	// exponents to a power of two.
	PollInterval int32 `idl:"name:nPollInterval" json:"poll_interval"`
	// refidSource:  A 32-bit code that identifies the particular reference clock of the
	// time source that the time service is synchronizing with, as specified in [RFC1305]
	// section 3.2.
	SourceID uint32 `idl:"name:refidSource" json:"source_id"`
	// qwLastSyncTicks:  The LastSyncTicks element value (see section 3.2.1.1).
	LastSyncTicks uint64 `idl:"name:qwLastSyncTicks" json:"last_sync_ticks"`
	// toRootDelay:  A 64-bit signed integer that indicates the total round-trip delay to
	// the primary time source, as specified in [RFC1305] section 3.2, in 100-nanosecond
	// units.
	ToRootDelay int64 `idl:"name:toRootDelay" json:"to_root_delay"`
	// tpRootDispersion:  A 64-bit unsigned integer that indicates the root dispersion,
	// as specified in [RFC1305] section 3.2, in 100-nanosecond units.
	RootDispersion uint64 `idl:"name:tpRootDispersion" json:"root_dispersion"`
	// nClockPrecision:  An integer that indicates the time resolution of the local system
	// clock, expressed in the same format as poll intervals that are specified in [RFC1305]
	// section 3.2, using units of seconds given as exponents to a power of two.
	ClockPrecision int32 `idl:"name:nClockPrecision" json:"clock_precision"`
	// wszSource:  The TimeSourceIPAddress element value (see section 3.2.1.1).
	Source string `idl:"name:wszSource;string;pointer:unique" json:"source"`
	// toSysPhaseOffset:  The SysPhaseOffset element value (see section 3.2.1.1).
	ToSystemPhaseOffset int64 `idl:"name:toSysPhaseOffset" json:"to_system_phase_offset"`
	// ulLcState:  The CurrentState element value (see section 3.2.1.1).
	LCState uint32 `idl:"name:ulLcState" json:"lc_state"`
	// ulTSFlags:  The TimeSourceFlags element value (see section 3.2.1.1).
	TSFlags uint32 `idl:"name:ulTSFlags" json:"ts_flags"`
	// ulClockRate:  The ClockRate element value (see section 3.2.1.1).
	ClockRate uint32 `idl:"name:ulClockRate" json:"clock_rate"`
	// ulNetlogonServiceBits:  An unsigned 32-bit integer that contains information about
	// the functionality that the time service provides, as specified in section 3.2.5.2.
	NetlogonServiceBits uint32 `idl:"name:ulNetlogonServiceBits" json:"netlogon_service_bits"`
	// eLastSyncResult:  An integer that indicates the TimeSync_ReturnResult code, as specified
	// in section 3.2.5.1.
	LastSyncResult uint32 `idl:"name:eLastSyncResult" json:"last_sync_result"`
	// tpTimeLastGoodSync:  The TimeLastGoodSync element value (see section 3.2.1.1).
	TimeLastGoodSync uint64 `idl:"name:tpTimeLastGoodSync" json:"time_last_good_sync"`
	// cEntries:  The number of additional configuration entries in pEntries.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// pEntries:  A pointer to W32TIME_ENTRY structures that represent additional configuration
	// entries.
	Entries []*Entry `idl:"name:pEntries;size_is:(cEntries)" json:"entries"`
}

func (o *StatusInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Entries != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.Entries))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *StatusInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.LeapIndicator); err != nil {
		return err
	}
	if err := w.WriteData(o.Stratum); err != nil {
		return err
	}
	if err := w.WriteData(o.PollInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceID); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSyncTicks); err != nil {
		return err
	}
	if err := w.WriteData(o.ToRootDelay); err != nil {
		return err
	}
	if err := w.WriteData(o.RootDispersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ClockPrecision); err != nil {
		return err
	}
	if o.Source != "" {
		_ptr_wszSource := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Source); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Source, _ptr_wszSource); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ToSystemPhaseOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.LCState); err != nil {
		return err
	}
	if err := w.WriteData(o.TSFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.ClockRate); err != nil {
		return err
	}
	if err := w.WriteData(o.NetlogonServiceBits); err != nil {
		return err
	}
	if err := w.WriteData(o.LastSyncResult); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeLastGoodSync); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	if o.Entries != nil || o.EntriesCount > 0 {
		_ptr_pEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Entries {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Entries[i1] != nil {
					if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Entry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Entries, _ptr_pEntries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *StatusInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.LeapIndicator); err != nil {
		return err
	}
	if err := w.ReadData(&o.Stratum); err != nil {
		return err
	}
	if err := w.ReadData(&o.PollInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceID); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSyncTicks); err != nil {
		return err
	}
	if err := w.ReadData(&o.ToRootDelay); err != nil {
		return err
	}
	if err := w.ReadData(&o.RootDispersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClockPrecision); err != nil {
		return err
	}
	_ptr_wszSource := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Source); err != nil {
			return err
		}
		return nil
	})
	_s_wszSource := func(ptr interface{}) { o.Source = *ptr.(*string) }
	if err := w.ReadPointer(&o.Source, _s_wszSource, _ptr_wszSource); err != nil {
		return err
	}
	if err := w.ReadData(&o.ToSystemPhaseOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.LCState); err != nil {
		return err
	}
	if err := w.ReadData(&o.TSFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClockRate); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetlogonServiceBits); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastSyncResult); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeLastGoodSync); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	_ptr_pEntries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*Entry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &Entry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEntries := func(ptr interface{}) { o.Entries = *ptr.(*[]*Entry) }
	if err := w.ReadPointer(&o.Entries, _s_pEntries, _ptr_pEntries); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

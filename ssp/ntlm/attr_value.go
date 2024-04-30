package ntlm

import (
	"context"
	"encoding/binary"

	"github.com/oiweiwei/go-msrpc/ssp/ntlm/internal"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

// AttrValues represents the map of attribute-value pairs.
type AttrValues map[AttrValueID]*Value

func (avl AttrValues) ToValue() *Value {

	val := Value{}

	for attrID, attrVal := range avl {
		switch attrID {
		case AttrNetBIOSComputerName:
			val.NetBIOSComputerName = attrVal.NetBIOSComputerName
		case AttrNetBIOSDomainName:
			val.NetBIOSDomainName = attrVal.NetBIOSDomainName
		case AttrDNSComputerName:
			val.DNSComputerName = attrVal.DNSComputerName
		case AttrDNSDomainName:
			val.DNSDomainName = attrVal.DNSDomainName
		case AttrDNSTreeName:
			val.DNSTreeName = attrVal.DNSTreeName
		case AttrFlags:
			val.Flag = attrVal.Flag
		case AttrTimestamp:
			val.Timestamp = attrVal.Timestamp
		case AttrSingleHost:
			val.SingleHostData = attrVal.SingleHostData
		case AttrTargetName:
			val.TargetName = attrVal.TargetName
		case AttrChannelBindings:
			val.ChannelBindings = attrVal.ChannelBindings
		}
	}
	return &val
}

// HasAttr function returns 'true' if the attribute is present.
func (avl AttrValues) HasAttr(attr AttrValueID) bool {
	_, ok := avl[attr]
	return ok
}

// Flag function returns the flag attribute.
func (avl AttrValues) Flag() AttrFlag {
	value := avl[AttrFlags]
	if value != nil {
		return value.Flag
	}
	return AttrFlag(0)
}

// Timestamp function returns the timestamp attribute.
func (avl AttrValues) Timestamp() Filetime {
	value := avl[AttrTimestamp]
	if value != nil {
		return value.Timestamp
	}
	return Filetime(0)
}

// Marshal function marshals the attribute-value pairs.
func (avl *AttrValues) Marshal(ctx context.Context) ([]byte, error) {

	e := internal.NewCodec(ctx, nil)

	if *avl != nil {
		for _, attr := range AttrIDList[1:] {

			value, ok := (*avl)[attr]
			if !ok {
				continue
			}

			raw, err := value.MarshalAttr(ctx, attr)
			if err != nil {
				return nil, err
			}

			e.WriteData(ctx, (uint16)(attr))
			e.WriteData(ctx, (uint16)(len(raw)))
			e.WriteBytes(ctx, raw, len(raw))
		}
	}

	// write EOL.
	e.WriteData(ctx, (uint16)(AttrEOL))
	e.WriteData(ctx, (uint16)(0))

	return e.Bytes(ctx)
}

// Unmarshal function unmarshals the attribute value pairs.
func (avl *AttrValues) Unmarshal(ctx context.Context, b []byte) error {

	e := internal.NewCodec(ctx, b)

	*avl = make(map[AttrValueID]*Value)

	for {

		var (
			attr   AttrValueID
			length uint16
			raw    []byte
		)

		e.ReadData(ctx, (*uint16)(&attr))
		e.ReadData(ctx, (*uint16)(&length))

		if attr == AttrEOL {
			break
		}

		(*avl)[attr] = new(Value)

		e.ReadBytes(ctx, &raw, int(length))

		if err := (*avl)[attr].UnmarshalAttr(ctx, attr, raw); err != nil {
			return err
		}
	}

	return e.ReadAll(ctx)
}

// The attribute value.
type Value struct {
	// The server's NetBIOS computer name. The name MUST be in Unicode,
	// and is not null-terminated.
	NetBIOSComputerName string
	// The server's NetBIOS domain name. The name MUST be in Unicode,
	// and is not null-terminated.
	NetBIOSDomainName string
	// The fully qualified domain name (FQDN) of the computer.
	// The name MUST be in Unicode, and is not null-terminated.
	DNSComputerName string
	// The fully qualified domain name (FQDN) of the computer.
	DNSDomainName string
	// The FQDN of the forest. The name MUST be in Unicode,
	// and is not null-terminated.
	DNSTreeName string
	// A 32-bit value indicating server or client configuration.
	Flag AttrFlag
	// A FILETIME structure in little-endian byte order that
	// contains the server local time. This structure is
	// always sent in the ChallengeMessage.
	Timestamp Filetime
	// A SingleHostData structure. The Value field contains a
	// platform-specific blob, as well as a MachineID created at
	// computer startup to identify the calling machine.
	SingleHostData *SingleHostData
	// The SPN of the target server. The name MUST be in Unicode and is
	// not null-terminated.
	TargetName string
	// A channel bindings hash. The Value field contains an MD5 hash of
	// a gss_channel_bindings_struct. An all-zero value of the hash is
	// used to indicate absence of channel bindings.
	ChannelBindings []byte
}

// MarshalAttr function marshals the given attribute.
func (v *Value) MarshalAttr(ctx context.Context, attrID AttrValueID) ([]byte, error) {

	if v == nil {
		// marshal in any case.
		v = &Value{}
	}

	var (
		raw []byte
		err error
	)

	switch attrID {
	case AttrNetBIOSComputerName:
		raw, err = utf16le.Encode(v.NetBIOSComputerName)
	case AttrNetBIOSDomainName:
		raw, err = utf16le.Encode(v.NetBIOSDomainName)
	case AttrDNSComputerName:
		raw, err = utf16le.Encode(v.DNSComputerName)
	case AttrDNSDomainName:
		raw, err = utf16le.Encode(v.DNSDomainName)
	case AttrDNSTreeName:
		raw, err = utf16le.Encode(v.DNSTreeName)
	case AttrFlags:
		raw = make([]byte, 4)
		binary.LittleEndian.PutUint32(raw, (uint32)(v.Flag))
	case AttrTimestamp:
		raw = make([]byte, 8)
		binary.LittleEndian.PutUint64(raw, (uint64)(v.Timestamp))
	case AttrSingleHost:
		raw, err = v.SingleHostData.Marshal(ctx)
	case AttrTargetName:
		raw, err = utf16le.Encode(v.TargetName)
	case AttrChannelBindings:
		raw = make([]byte, 16)
		copy(raw, v.ChannelBindings)
	}

	if err != nil {
		return nil, err
	}

	return raw, nil
}

// UnmarshalAttr function unmarshals the given attribute.
func (v *Value) UnmarshalAttr(ctx context.Context, attrID AttrValueID, b []byte) error {

	var err error

	if len(b) == 0 {
		return nil
	}

	switch attrID {
	case AttrNetBIOSComputerName:
		v.NetBIOSComputerName, err = utf16le.Decode(b)
	case AttrNetBIOSDomainName:
		v.NetBIOSDomainName, err = utf16le.Decode(b)
	case AttrDNSComputerName:
		v.DNSComputerName, err = utf16le.Decode(b)
	case AttrDNSDomainName:
		v.DNSDomainName, err = utf16le.Decode(b)
	case AttrDNSTreeName:
		v.DNSTreeName, err = utf16le.Decode(b)
	case AttrFlags:
		v.Flag = (AttrFlag)(binary.LittleEndian.Uint32(b))
	case AttrTimestamp:
		v.Timestamp = Filetime(binary.LittleEndian.Uint64(b))
	case AttrSingleHost:
		v.SingleHostData = &SingleHostData{}
		err = v.SingleHostData.Unmarshal(ctx, b)
	case AttrTargetName:
		v.TargetName, err = utf16le.Decode(b)
	case AttrChannelBindings:
		v.ChannelBindings = make([]byte, len(b))
		copy(v.ChannelBindings, b)
	}

	return err
}

// The attribute-value pair identifier.
type AttrValueID uint16

const (
	// Indicates that this is the last AV_PAIR in the list.
	// This type of information MUST be present in the AV pair list.
	AttrEOL AttrValueID = 0x0000
	// The server's NetBIOS computer name. The name MUST be in Unicode,
	// and is not null-terminated.
	// This type of information MUST be present in the AV pair list.
	AttrNetBIOSComputerName AttrValueID = 0x0001
	// The server's NetBIOS domain name. The name MUST be in Unicode,
	// and is not null-terminated.
	// This type of information MUST be present in the AV_pair list.
	AttrNetBIOSDomainName AttrValueID = 0x0002
	// The fully qualified domain name (FQDN) of the computer.
	// The name MUST be in Unicode, and is not null-terminated.
	AttrDNSComputerName AttrValueID = 0x0003
	// The FQDN of the domain. The name MUST be in Unicode,
	// and is not null-terminated.
	AttrDNSDomainName AttrValueID = 0x0004
	// The FQDN of the forest. The name MUST be in Unicode,
	// and is not null-terminated.
	AttrDNSTreeName AttrValueID = 0x0005
	// A 32-bit value indicating server or client configuration.
	AttrFlags AttrValueID = 0x0006
	// A FILETIME structure in little-endian byte order that
	// contains the server local time. This structure is
	// always sent in the ChallengeMessage.
	AttrTimestamp AttrValueID = 0x0007
	// A SingleHostData structure. The Value field contains a
	// platform-specific blob, as well as a MachineID created at
	// computer startup to identify the calling machine.
	AttrSingleHost AttrValueID = 0x0008
	// The SPN of the target server. The name MUST be in Unicode and is
	// not null-terminated.
	AttrTargetName AttrValueID = 0x0009
	// A channel bindings hash. The Value field contains an MD5 hash of
	// a gss_channel_bindings_struct. An all-zero value of the hash is
	// used to indicate absence of channel bindings.
	AttrChannelBindings AttrValueID = 0x000a
)

var (
	// The list of allowed attributes.
	AttrIDList = [...]AttrValueID{
		AttrEOL,
		AttrNetBIOSComputerName,
		AttrNetBIOSDomainName,
		AttrDNSComputerName,
		AttrDNSDomainName,
		AttrDNSTreeName,
		AttrFlags,
		AttrTimestamp,
		AttrSingleHost,
		AttrTargetName,
		AttrChannelBindings,
	}
)

// A 32-bit value indicating server or client configuration.
type AttrFlag uint32

// IsSet function returns `true` if flag is set.
func (f AttrFlag) IsSet(ff AttrFlag) bool {
	return f&ff != 0
}

const (
	// Indicates to the client that the account authentication is
	// constrained.
	AccountAuthenticationConstrained AttrFlag = 1 << 0
	// Indicates that the client is providing message integrity in
	// the MIC field in the AuthenticateMessage.
	MICProvided = 1 << 1
	// Indicates that the client is providing a target SPN generated
	// from an untrusted source.
	SPNFromUntrustedSource = 1 << 2
)

// A 256-bit random number created at computer startup to identify
// the calling machine.
type MachineID []byte

// The Single_Host_Data structure allows a client to send machine-specific
// information within an authentication exchange to services on the same
// machine. The client can produce additional information to be processed
// in an implementation-specific way when the client and server are on the
// same host. If the server and client platforms are different or if they
// are on different hosts, then the information MUST be ignored. Any fields
// after the MachineID field MUST be ignored on receipt.
type SingleHostData struct {
	// A 32-bit unsigned integer that defines the length, in bytes, of
	// the Value field in the AV_PAIR structure.
	Size uint32
	// A 32-bit integer value containing 0x00000000.
	_ uint32
	//  An 8-byte platform-specific blob containing info only relevant when
	// the client and the server are on the same host.
	CustomData []byte
	// A 256-bit random number created at computer startup to identify
	// the calling machine.
	MachineID MachineID
}

// Marshal function marshals the SingleHostData structure.
func (v *SingleHostData) Marshal(ctx context.Context) ([]byte, error) {

	if v == nil {
		// marshal in any case.
		v = &SingleHostData{}
	}

	e := internal.NewCodec(ctx, nil)

	// size.
	e.WriteData(ctx, uint32(0))
	// pad.
	e.WriteData(ctx, uint32(0))
	// custom_data.
	e.WriteBytes(ctx, v.CustomData, 8)
	// machine_id.
	e.WriteBytes(ctx, ([]byte)(v.MachineID), 32)

	b, err := e.Bytes(ctx)
	if err != nil {
		return nil, err
	}

	// re-encode size.
	v.Size = uint32(len(b))
	binary.LittleEndian.PutUint32(b, v.Size)

	return b, nil
}

// Unmarshal function unmarshals the SingleHostData structure.
func (v *SingleHostData) Unmarshal(ctx context.Context, b []byte) error {

	e := internal.NewCodec(ctx, b)

	// size.
	e.ReadData(ctx, &v.Size)
	// pad.
	_pad32 := uint32(0)
	e.ReadData(ctx, &_pad32)
	// custom_data.
	e.ReadBytes(ctx, &v.CustomData, 8)
	// machine_id.
	e.ReadBytes(ctx, (*[]byte)(&v.MachineID), 32)

	return e.ReadAll(ctx)
}

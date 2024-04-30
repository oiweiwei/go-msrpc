package dcetypes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
)

var (
	// https://github.com/boundary/wireshark/blob/master/epan/dissectors/packet-dcerpc-epm.c#L250
	ProtocolOSI_OID           = 0x00
	ProtocolDNA_SESSCTL       = 0x02
	ProtocolDNA_SESSCTL_V3    = 0x03
	ProtocolDNA_NSP           = 0x04
	ProtocolOSI_TP4           = 0x05
	ProtocolOSI_CLNS          = 0x06
	ProtocolTCP               = 0x07
	ProtocolUDP               = 0x08
	ProtocolIP                = 0x09
	ProtocolRPC_CL            = 0x0a
	ProtocolRPC_CO            = 0x0b
	ProtocolSPX               = 0x0c /* from DCOM spec (is this correct?) */
	ProtocolUUID              = 0x0d
	ProtocolIPX               = 0x0e /* from DCOM spec (is this correct?) */
	ProtocolNamedPipe         = 0x0f
	ProtocolLRPC              = 0x10
	ProtocolNetBIOS           = 0x11
	ProtocolNetBEUI           = 0x12
	ProtocolNetwareSPX        = 0x13
	ProtocolNetwareIPX        = 0x14
	ProtocolAppleTalkStream   = 0x16
	ProtocolAppleTalkDatagram = 0x17
	ProtocolAppleTalk         = 0x18
	ProtocolNetBIOS2          = 0x19
	ProtocolVinesSPP          = 0x1a
	ProtocolVinesIPC          = 0x1b
	ProtocolStreetTalk        = 0x1c
	ProtocolHTTP              = 0x1f
	ProtocolUnixDomain        = 0x20
	ProtocolNULL              = 0x21
	ProtocolNetBIOS3          = 0x22
)

type Floor struct {
	Protocol     uint8      `json:"protocol"`
	UUID         *uuid.UUID `json:"uuid,omitempty"`
	VersionMajor uint16     `json:"version_major"`
	Data         []byte     `json:"data"`
}

func (f *Floor) VersionMinor() uint16 {
	return binary.LittleEndian.Uint16(f.Data)
}

func (f *Floor) Port() uint16 {
	return binary.BigEndian.Uint16(f.Data)
}

func (f *Floor) Str() string {
	return string(f.Data[:len(f.Data)-1])
}

func (f *Floor) IP() net.IP {
	return net.IP(f.Data)
}

func (f *Floor) Hex() string {
	return hex.EncodeToString(f.Data)
}

func (f *Floor) String() string {
	switch int(f.Protocol) {
	case ProtocolUUID:
		return fmt.Sprintf("0x%02x: %s v%d.%d", f.Protocol, f.UUID, f.VersionMajor, f.VersionMinor())
	case ProtocolNamedPipe, ProtocolLRPC:
		return fmt.Sprintf("0x%02x: pipe:%s", f.Protocol, f.Str())
	case ProtocolNetBIOS, ProtocolNetBIOS2, ProtocolNetBIOS3:
		return fmt.Sprintf("0x%02x: nb:%s", f.Protocol, f.Str())
	case ProtocolTCP:
		return fmt.Sprintf("0x%02x: tcp:%d", f.Protocol, f.Port())
	case ProtocolUDP:
		return fmt.Sprintf("0x%02x: udp:%d", f.Protocol, f.Port())
	case ProtocolHTTP:
		return fmt.Sprintf("0x%02x: http:%d", f.Protocol, f.Port())
	case ProtocolIP:
		return fmt.Sprintf("0x%02x: ip:%s", f.Protocol, f.IP())
	case ProtocolRPC_CO:
		return fmt.Sprintf("0x%02x: rpc v5.%d", f.Protocol, f.VersionMinor())
	case ProtocolRPC_CL:
		return fmt.Sprintf("0x%02x: rpc v4.%d", f.Protocol, f.VersionMinor())
	default:
		return fmt.Sprintf("0x%02x: data:%s", f.Protocol, f.Hex())
	}
}

func FloorsToTower(floors []*Floor) *Tower {

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, uint16(len(floors)))

	for _, f := range floors {
		if f.Protocol == uint8(ProtocolUUID) {
			binary.Write(buf, binary.LittleEndian, uint16(19))
			binary.Write(buf, binary.LittleEndian, f.Protocol)
			if f.UUID == nil {
				f.UUID = &uuid.UUID{}
			}
			binary.Write(buf, binary.LittleEndian, f.UUID)
			binary.Write(buf, binary.LittleEndian, f.VersionMajor)
		} else {
			binary.Write(buf, binary.LittleEndian, uint16(1))
			binary.Write(buf, binary.LittleEndian, f.Protocol)
		}
		binary.Write(buf, binary.LittleEndian, uint16(len(f.Data)))
		buf.Write(f.Data)
	}

	return &Tower{
		TowerOctetString: buf.Bytes(),
	}
}

func (o *Tower) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Floors())
}

// Binding function returns the binding representation of the tower.
func (o *Tower) Binding() *dcerpc.Binding {

	var b dcerpc.Binding

	for i, floor := range o.Floors() {
		switch int(floor.Protocol) {
		case ProtocolUUID:
			if i != 0 {
				// set the transfer syntax identifier.
				b.TransferSyntaxID.IfUUID = floor.UUID
				b.TransferSyntaxID.IfVersionMajor = floor.VersionMajor
				b.TransferSyntaxID.IfVersionMinor = floor.VersionMinor()
				break
			}
			// set syntax identifier.
			b.SyntaxID.IfUUID = floor.UUID
			b.SyntaxID.IfVersionMajor = floor.VersionMajor
			b.SyntaxID.IfVersionMinor = floor.VersionMinor()
		case ProtocolTCP:
			// the tcp port.
			b.StringBinding.Endpoint = strconv.Itoa(int(floor.Port()))
			b.StringBinding.ProtocolSequence = dcerpc.ProtocolSequenceIPTCP
		case ProtocolUDP:
			// the udp port.
			b.StringBinding.Endpoint = strconv.Itoa(int(floor.Port()))
			b.StringBinding.ProtocolSequence = dcerpc.ProtocolSequenceIPUDP
		case ProtocolHTTP:
			// the http port.
			b.StringBinding.Endpoint = strconv.Itoa(int(floor.Port()))
			b.StringBinding.ProtocolSequence = dcerpc.ProtocolSequenceHTTP
		case ProtocolNamedPipe:
			// the named pipe name.
			b.StringBinding.Endpoint = floor.Str()
			b.StringBinding.ProtocolSequence = dcerpc.ProtocolSequenceNamedPipe
		case ProtocolLRPC:
			// the local interprocess communication endpoint.
			b.StringBinding.Endpoint = floor.Str()
			b.StringBinding.ProtocolSequence = dcerpc.ProtocolSequenceLRPC
		case ProtocolIP:
			// the ip address.
			b.StringBinding.NetworkAddress = floor.IP().String()
		case ProtocolNetBIOS:
			// the netbios computer name.
			b.StringBinding.ComputerName = floor.Str()
		}
	}

	return &b
}

// Floors function returns the decoded tower floors.
func (o *Tower) Floors() []*Floor {

	buf := bytes.NewBuffer(o.TowerOctetString)

	var l uint16
	binary.Read(buf, binary.LittleEndian, &l)

	floors := make([]*Floor, l)

	for i := range floors {

		floors[i] = &Floor{}

		binary.Read(buf, binary.LittleEndian, &l)

		if l >= 1 {
			binary.Read(buf, binary.LittleEndian, &floors[i].Protocol)
			l -= 1
		}
		if l >= 16 {
			floors[i].UUID = new(uuid.UUID)
			binary.Read(buf, binary.LittleEndian, floors[i].UUID)
			l -= 16
		}
		if l >= 2 {
			binary.Read(buf, binary.LittleEndian, &floors[i].VersionMajor)
			l -= 2
		}

		binary.Read(buf, binary.LittleEndian, &l)
		floors[i].Data = make([]byte, l)
		buf.Read(floors[i].Data)
	}

	return floors
}

// package netlogon implements the Netlogon secure channel client security
// service client as described in
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrpc/fb50db72-7f71-478d-a180-12eb0ca3b36b.
//
// This package also contains client-side GSSAPI bindings (InitSecurityContext, Wrap, Unwrap and so on).
package netlogon

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/oiweiwei/go-msrpc/text/encoding/oem"
)

type NegFlag uint32

func (c NegFlag) IsSet(cc NegFlag) bool {
	return c&cc != 0
}

const (
	// Buffer contains a NetBIOS domain name as an OEM_STRING.
	NegFlagNetBIOSDomainName NegFlag = 1 << (31 - 31)
	// Buffer contains a NetBIOS computer name as an OEM_STRING.
	NegFlagNetBIOSComputerName NegFlag = 1 << (31 - 30)
	// Buffer contains a DNS domain name as a compressed UTF-8 string,
	// as specified in [RFC1035].
	NegFlagDNSDomainName NegFlag = 1 << (31 - 29)
	// Buffer contains a DNS host name as a compressed UTF-8 string.
	NegFlagDNSHostName NegFlag = 1 << (31 - 28)
	// Buffer contains a NetBIOS computer name as a compressed UTF-8
	// string.
	NegFlagNetBIOSComputerNameUTF8 NegFlag = 1 << (31 - 27)
)

type Cap uint32

const (
	CapR0                     = 1 << (31 - 0)
	CapSecureRPC              = 1 << (31 - 1)
	CapR2                     = 1 << (31 - 2)
	CapR3                     = 1 << (31 - 3)
	CapR4                     = 1 << (31 - 4)
	CapR5                     = 1 << (31 - 5)
	CapR6                     = 1 << (31 - 6)
	CapAES_SHA2               = 1 << (31 - 7)
	CapR8                     = 1 << (31 - 8)
	CapR9                     = 1 << (31 - 9)
	CapRODCPassThrough        = 1 << (31 - 10)
	CapIgnoreNT4EmuADM        = 1 << (31 - 11)
	CapCrossForestTrust       = 1 << (31 - 12)
	CapNetrLogonGetDomainInfo = 1 << (31 - 13)
	CapNetrServerPasswordSet2 = 1 << (31 - 14)
	CapR15                    = 1 << (31 - 15)
	CapTransitiveTrust        = 1 << (31 - 16)
	CapStrongKey              = 1 << (31 - 17)
	CapAvoidSAReplication     = 1 << (31 - 18)
	CapAvoidUAReplication     = 1 << (31 - 19)
	CapConcurrentRPC          = 1 << (31 - 20)
	CapGenericPassthrough     = 1 << (31 - 21)
	CapNetrLogonSendToSam     = 1 << (31 - 22)
	CapRefusePasswdChange     = 1 << (31 - 23)
	CapNetrDatabaseRedo       = 1 << (31 - 24)
	CapNotReqValidationLevel2 = 1 << (31 - 25)
	CapRestartFullSync        = 1 << (31 - 26)
	CapBDCHandlingChangelog   = 1 << (31 - 27)
	CapR28                    = 1 << (31 - 28)
	CapRC4                    = 1 << (31 - 29)
	CapPersistentDBUpdate     = 1 << (31 - 30)
	CapR31                    = 1 << (31 - 31)
)

func (c Cap) IsSet(cc Cap) bool {
	return c&cc != 0
}

type MessageType uint32

const (
	// This is a negotiate request message.
	MessageTypeRequest MessageType = 0x00000000
	// This is a negotiate response message.
	MessageTypeResponse MessageType = 0x00000001
)

// The NL_AUTH_MESSAGE structure is a token containing information that
// is part of the first message in establishing a security context between
// a client and a server. It is used for establishing the secure session
// when Netlogon functions as a security support provider (SSP).
type AuthMessage struct {
	// A 32-bit unsigned integer. This value is used to indicate whether
	// the message is a negotiate request message sent from a client to
	// a server, or a negotiate response message sent from the server to
	// the client.
	MessageType MessageType
	// A set of bit flags indicating the principal names carried in the
	// request. A flag is TRUE (or set) if its value is equal to 1
	Flags NegFlag
	// NetBIOS domain name as an OEM_STRING.
	NetBIOSDomainName string
	// NetBIOS computer name as an OEM_STRING.
	NetBIOSComputerName string
	// DNS domain name as a compressed UTF-8 string, as specified in [RFC1035].
	DNSDomainName string
	// DNS host name as a compressed UTF-8 string.
	DNSHostName string
}

func (m *AuthMessage) Marshal() ([]byte, error) {

	w := bytes.NewBuffer(nil)

	if m.NetBIOSDomainName != "" {
		m.Flags = m.Flags | NegFlagNetBIOSDomainName
	}

	if m.NetBIOSComputerName != "" {
		m.Flags = m.Flags | NegFlagNetBIOSComputerName | NegFlagNetBIOSComputerNameUTF8
	}

	if m.DNSDomainName != "" {
		m.Flags = m.Flags | NegFlagDNSDomainName
	}

	if m.DNSHostName != "" {
		m.Flags = m.Flags | NegFlagDNSHostName
	}

	binary.Write(w, binary.LittleEndian, (uint32)(m.MessageType))
	binary.Write(w, binary.LittleEndian, (uint32)(m.Flags))

	if m.Flags.IsSet(NegFlagNetBIOSDomainName) {
		b, err := oem.Encode(m.NetBIOSDomainName)
		if err != nil {
			return nil, err
		}
		w.Write(b)
		w.WriteByte(0)
	}

	if m.Flags.IsSet(NegFlagNetBIOSComputerName) {
		b, err := oem.Encode(m.NetBIOSComputerName)
		if err != nil {
			return nil, err
		}
		w.Write(b)
		w.WriteByte(0)
	}

	if m.Flags.IsSet(NegFlagDNSDomainName) {
		for _, label := range strings.Split(strings.TrimRight(m.DNSDomainName, "."), ".") {
			w.WriteByte(uint8(len(label)))
			w.Write([]byte(label))
		}
		w.WriteByte(0)
	}

	if m.Flags.IsSet(NegFlagDNSHostName) {
		for _, label := range strings.Split(strings.TrimRight(m.DNSHostName, "."), ".") {
			w.WriteByte(uint8(len(label)))
			w.Write([]byte(label))
		}
		w.WriteByte(0)
	}

	if m.Flags.IsSet(NegFlagNetBIOSComputerNameUTF8) {
		w.WriteByte(uint8(len(m.NetBIOSComputerName)))
		w.Write([]byte(m.NetBIOSComputerName))
		w.WriteByte(0)
	}

	return w.Bytes(), nil
}

func (m *AuthMessage) Unmarshal(b []byte) error {

	var err error

	r := bytes.NewBuffer(b)

	binary.Read(r, binary.LittleEndian, (*uint32)(&m.MessageType))
	binary.Read(r, binary.LittleEndian, (*uint32)(&m.Flags))

	if m.Flags.IsSet(NegFlagNetBIOSDomainName) {
		if m.NetBIOSDomainName, err = readNBName(r, true); err != nil {
			return err
		}
	}

	if m.Flags.IsSet(NegFlagNetBIOSComputerName) {
		if m.NetBIOSComputerName, err = readNBName(r, true); err != nil {
			return err
		}
	}

	if m.Flags.IsSet(NegFlagDNSDomainName) {
		if m.DNSDomainName, err = readDNSName(r); err != nil {
			return err
		}
	}

	if m.Flags.IsSet(NegFlagDNSHostName) {
		if m.DNSHostName, err = readDNSName(r); err != nil {
			return err
		}
	}

	if m.Flags.IsSet(NegFlagNetBIOSComputerNameUTF8) {
		if m.NetBIOSComputerName, err = readNBName(r, false); err != nil {
			return err
		}
	}

	return nil
}

func readDNSName(r *bytes.Buffer) (string, error) {

	l, err := r.ReadBytes(0)
	if err != nil {
		return "", err
	}

	var labels []string
	for len(l) > 0 {
		sz := int(l[0])
		if sz+1 > len(l) {
			return "", fmt.Errorf("invalid dns label length")
		}
		labels, l = append(labels, string(l[1:sz+1])), l[sz+1:]
	}

	return strings.Join(labels, "."), nil
}

func readNBName(r *bytes.Buffer, isOEM bool) (string, error) {
	l, err := r.ReadBytes(0)
	if err != nil {
		return "", err
	}

	if !isOEM {
		return string(l[:len(l)-1]), nil
	}

	s, err := oem.Decode(l[:len(l)-1])
	if err != nil {
		return "", err
	}

	return s, nil
}

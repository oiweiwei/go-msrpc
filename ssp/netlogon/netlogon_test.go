package netlogon

import "testing"

func TestAuthMessageRoundTrip(t *testing.T) {
	in := &AuthMessage{
		MessageType:         MessageTypeRequest,
		NetBIOSDomainName:   "SF",
		NetBIOSComputerName: "LABV2-SF",
		DNSDomainName:       "sf.example.com",
		DNSHostName:         "labv2-sf.sf.example.com",
	}
	b, err := in.Marshal()
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var out AuthMessage
	if err := out.Unmarshal(b); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if out.NetBIOSComputerName != in.NetBIOSComputerName {
		t.Errorf("NetBIOSComputerName = %q, want %q", out.NetBIOSComputerName, in.NetBIOSComputerName)
	}
	if out.NetBIOSDomainName != in.NetBIOSDomainName {
		t.Errorf("NetBIOSDomainName = %q, want %q", out.NetBIOSDomainName, in.NetBIOSDomainName)
	}
}

func TestAuthMessageUnmarshalUTF8ComputerNameOnly(t *testing.T) {
	// The UTF-8 computer name is a compressed (RFC 1035 label) string: length, bytes, terminating zero.
	b := []byte{
		0, 0, 0, 0, // NL_AUTH_MESSAGE request
		0x10, 0, 0, 0, // NETBIOS_COMPUTER_NAME_UTF8 only
		8, 'L', 'A', 'B', 'V', '2', '-', 'S', 'F', 0,
	}
	var m AuthMessage
	if err := m.Unmarshal(b); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if m.NetBIOSComputerName != "LABV2-SF" {
		t.Errorf("NetBIOSComputerName = %q, want %q", m.NetBIOSComputerName, "LABV2-SF")
	}
}

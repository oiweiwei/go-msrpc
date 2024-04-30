package ntlm

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/oiweiwei/go-msrpc/ssp/credential"
)

func TestNTOWF(t *testing.T) {

	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/7795bd0e-fd5e-43ec-bd9c-994704d8ee26

	v2 := &V2{Config: &Config{}}

	ret, err := v2.NTOWF(context.Background(), credential.NewFromPassword("Domain\\User", "Password"))
	if err != nil {
		t.Errorf("test_ntowf: unexpected error: %v", err)
	}

	exp := "0c868a403bfd7a93a3001ef22ef02e3f"

	if val := hex.EncodeToString(ret); val != exp {
		t.Errorf("test_ntowf: mismatched value %s != %s", val, exp)
	}

}

func TestSessionBaseKey(t *testing.T) {

	v2 := &V2{Config: &Config{}}

	cm := &ChallengeMessage{
		Negotiate: NegotiateKeyExchange |
			Negotiate56 |
			Negotiate128 |
			NegotiateVersion |
			NegotiateTargetInfo |
			NegotiateExtendedSessionSecurity |
			TargetTypeServer |
			NegotiateAlwaysSign |
			NegotiateNTLM |
			NegotiateSeal |
			NegotiateSign |
			NegotiateOEM |
			NegotiateUnicode,
		TargetInfo: AttrValues{
			AttrNetBIOSComputerName: &Value{NetBIOSComputerName: "Server"},
			AttrNetBIOSDomainName:   &Value{NetBIOSDomainName: "Domain"},
		},
		ServerChallenge: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
	}

	nonce := []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	resp, err := v2.ChallengeResponse(context.Background(), credential.NewFromPassword("Domain\\User", "Password"),
		cm, nonce)
	if err != nil {
		t.Errorf("test_session_key: unexpected error: %v", err)
	}

	fmt.Println(hex.EncodeToString(resp.SessionBaseKey))
	fmt.Println(hex.Dump(resp.Tmp))
	fmt.Println(hex.Dump(resp.LM))
	fmt.Println(hex.Dump(resp.NT))

}

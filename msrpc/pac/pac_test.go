package pac

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jcmturner/gokrb5/v8/test/testdata"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

func TestPAC(t *testing.T) {

	b, err := hex.DecodeString(testdata.MarshaledPAC_AD_WIN2K_PAC)

	if err != nil {
		t.Fatal(err)
	}

	var pac PACType

	if err := ndr.Unmarshal(b, &pac, ndr.Opaque); err != nil {
		t.Fatal(err)
	}

	j := func(v any) string {
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			panic(err)
		}
		return string(b)
	}

	fmt.Println(j(pac))

	data := make(map[string]any)

	for _, e := range pac.Buffers {

		bb := b[e.Offset : e.Offset+uint64(e.BufferLength)]

		switch e.Type {
		case 1:

			var logon KerberosValidationInfo
			if err := ndr.UnmarshalWithTypeSerializationV1(bb, ndr.UnmarshalerPointer(&logon)); err != nil {
				t.Fatal(err)
			}

			data["KerberosValidationInfo"] = logon

		case 10:
			var client PACClientInfo
			if err := ndr.Unmarshal(bb, &client, ndr.Opaque); err != nil {
				t.Fatal(err)
			}

			data["PACClientInfo"] = client

		case 12:
			var upn UPNDNSInfo
			if err := ndr.Unmarshal(bb, &upn, ndr.Opaque); err != nil {
				t.Fatal(err)
			}

			data["UPNDNSInfo"] = upn
		case 6:
			var chk PACSignatureData
			if err := ndr.Unmarshal(bb, &chk, ndr.Opaque); err != nil {
				t.Fatal(err)
			}
			data["PACSignatureData"] = chk
		case 7:
			var chk PACSignatureData
			if err := ndr.Unmarshal(bb, &chk, ndr.Opaque); err != nil {
				t.Fatal(err)
			}
			data["KDC_PACSignatureData"] = chk
		}

	}

	fmt.Println(j(data))
}

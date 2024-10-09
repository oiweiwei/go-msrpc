package pac

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jcmturner/gokrb5/v8/test/testdata"

	pac_testdata "github.com/oiweiwei/go-msrpc/msrpc/pac/testdata"

	"github.com/oiweiwei/go-msrpc/msrpc/adts/claims/claims/v1"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var j = func(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestPAC(t *testing.T) {

	for _, vec := range []string{
		testdata.MarshaledPAC_AD_WIN2K_PAC,
		pac_testdata.MarshaledPAC_AD_WIN2K_PAC,
	} {

		b, err := hex.DecodeString(vec)

		if err != nil {
			t.Fatal(err)
		}

		var pac PAC

		if err := pac.Unmarshal(b); err != nil {
			t.Fatal(err)
		}

		b, err = pac.Marshal()
		if err != nil {
			t.Fatal(err)
		}

		pac = PAC{}

		if err = pac.Unmarshal(b); err != nil {
			t.Fatal(err)
		}

		fmt.Println(j(pac))
	}
}

func TestClaimsSet(t *testing.T) {

	var meta claims.ClaimsSetMetadata

	for _, vec := range []string{
		testdata.MarshaledPAC_ClientClaimsInfoStr,
		testdata.MarshaledPAC_ClientClaimsInfoMultiStr,
		testdata.MarshaledPAC_ClientClaimsInfoInt,
		testdata.MarshaledPAC_ClientClaimsInfoMulti,
		testdata.MarshaledPAC_ClientClaimsInfoMultiUint,
	} {

		b, err := hex.DecodeString(vec)
		if err != nil {
			t.Fatal(err)
		}

		if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&meta)); err != nil {
			t.Fatal(err)
		}

		j := func(v any) string {
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				panic(err)
			}
			return string(b)
		}

		cls, err := meta.Claims()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(j(cls.ClaimsMaps()))
	}

}

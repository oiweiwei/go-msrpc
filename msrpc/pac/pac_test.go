package pac

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jcmturner/gokrb5/v8/test/testdata"
	"github.com/jcmturner/gokrb5/v8/types"

	ndr "github.com/oiweiwei/go-msrpc/ndr"

	"github.com/oiweiwei/go-msrpc/msrpc/adts/claims/claims/v1"

	pac_testdata "github.com/oiweiwei/go-msrpc/msrpc/pac/testdata"
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

func TestCredentialsInfo(t *testing.T) {

	for _, vec := range [][2]string{
		{pac_testdata.MarshaledPAC_CredentialsInfo_Key, pac_testdata.MarshaledPAC_CredentialsInfo},
	} {

		b, err := hex.DecodeString(vec[1])
		if err != nil {
			t.Fatal(err)
		}

		var ci PACCredentialInfo

		if err := ndr.Unmarshal(b, &ci, ndr.Opaque); err != nil {
			t.Fatal(err)
		}

		bkey, err := hex.DecodeString(vec[0])
		if err != nil {
			t.Fatal(err)
		}

		key := types.EncryptionKey{
			KeyType:  int32(ci.EncryptionType),
			KeyValue: bkey,
		}

		pkg, err := ci.DecryptCredentialData(key)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(j(pkg))
	}

}

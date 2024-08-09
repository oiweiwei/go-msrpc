package pac

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jcmturner/gokrb5/v8/test/testdata"
)

func TestPAC(t *testing.T) {

	b, err := hex.DecodeString(testdata.MarshaledPAC_AD_WIN2K_PAC)

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

	j := func(v any) string {
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			panic(err)
		}
		return string(b)
	}

	fmt.Println(j(pac))
}

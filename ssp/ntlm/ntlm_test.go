package ntlm

import (
	"context"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestNegotiateMessage(t *testing.T) {

	m := &NegotiateMessage{
		Negotiate:   NegotiateVersion | NegotiateOEM | NegotiateUnicode | NegotiateOEMDomainSupplied | NegotiateOEMWorkstationSupplied,
		DomainName:  "MSAD.COM",
		Workstation: "PC-12345",
	}

	b, err := m.Marshal(context.Background())
	fmt.Println(err)
	fmt.Println(hex.Dump(b))

	dm := &NegotiateMessage{}

	fmt.Println(dm.Unmarshal(context.Background(), b))

	fmt.Println(dm)

	fmt.Println("deep_equal", reflect.DeepEqual(dm, m))
}

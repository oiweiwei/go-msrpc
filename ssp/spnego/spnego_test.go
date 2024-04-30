package spnego

import (
	"context"
	"encoding/asn1"
	"fmt"
	"reflect"
	"testing"
)

func TestNegTokenInit2(t *testing.T) {

	negInit2 := &NegTokenInit{
		MechTypes:    []asn1.ObjectIdentifier{MechanismTypeSPNEGO},
		HintName:     HintName,
		HintAddress:  []byte{1, 2, 3},
		MechToken:    []byte{1, 2, 3, 4},
		MechTokenMIC: []byte{5, 6, 7, 8},
	}

	b, err := negInit2.Marshal(context.Background())
	if err != nil {
		t.Errorf("init2: marshal: %v", err)
	}

	negInit2Exp := &NegTokenInit{}

	if err := negInit2Exp.Unmarshal(context.Background(), b); err != nil {
		t.Errorf("init2: unmarshal: %v", err)
	}

	if !reflect.DeepEqual(negInit2, negInit2Exp) {
		t.Errorf("neg init 2 marshal/unmarshal does not match")
	}

	fmt.Println(negInit2)
	fmt.Println(negInit2Exp)
}

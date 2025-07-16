package drsuapi

import (
	"encoding/asn1"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestData = []asn1.ObjectIdentifier{
	asn1.ObjectIdentifier{1, 2, 840, 113556, 1, 4, 94},
	asn1.ObjectIdentifier{2, 5, 6, 2},
	asn1.ObjectIdentifier{1, 2, 840, 113556, 1, 2, 1},
	asn1.ObjectIdentifier{1, 2, 840, 113556, 1, 3, 223},
	asn1.ObjectIdentifier{1, 2, 840, 113556, 1, 5, 7000, 53},
}

func TestPrefixTable(t *testing.T) {

	prefixes := make(PrefixTable, 0)

	attrs := make([]uint32, len(TestData))

	for i, oid := range TestData {
		attr, err := prefixes.OIDToAttribute(oid)
		t.Logf("Converting OID %v to attribute: %d", oid, attr)
		assert.NoError(t, err, "Failed to convert OID to attribute")
		attrs[i] = attr
	}

	for i := range TestData {
		oid, err := prefixes.AttributeToOID(attrs[i])
		t.Logf("Converting attribute %d to OID: %v", attrs[i], oid)
		assert.NoError(t, err, "Failed to convert attribute to OID")
		assert.Equal(t, oid, TestData[i], "Converted OID does not match original OID")
	}
}

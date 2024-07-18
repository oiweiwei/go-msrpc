package drsuapi

import (
	"encoding/asn1"
	"fmt"
)

type PrefixTable map[uint32][]byte

func (o *SchemaPrefixTable) Build() PrefixTable {
	table := PrefixTable{}
	if o == nil {
		return table
	}
	for i1 := range o.PrefixEntry {
		table[o.PrefixEntry[i1].Index] = o.PrefixEntry[i1].Prefix.Elements
	}
	return table
}

func (o PrefixTable) AttributeToOID(attrType uint32) (asn1.ObjectIdentifier, error) {

	if o == nil {
		return nil, fmt.Errorf("prefix table is empty")
	}

	upper, lower := attrType/65536, attrType%65536

	prefix, ok := o[upper]
	if !ok {
		return asn1.ObjectIdentifier{}, fmt.Errorf("prefix for id %d was not found", upper)
	}

	raw := make([]byte, len(prefix))
	copy(raw, prefix)

	if lower < 128 {
		raw = append(raw, uint8(lower))
	} else {
		if lower >= 32768 {
			lower -= 32768
		}
		raw = append(raw, uint8((lower/128)%128+128), uint8(lower%128))
	}

	return RawPrefixToOID(raw)
}

func RawPrefixToOID(b []byte) (asn1.ObjectIdentifier, error) {

	raw := asn1.RawValue{
		Tag:   asn1.TagOID,
		Bytes: b,
	}

	b, err := asn1.Marshal(raw)
	if err != nil {
		return nil, err
	}

	var oid asn1.ObjectIdentifier

	if _, err := asn1.Unmarshal(b, &oid); err != nil {
		return nil, err
	}

	return oid, nil
}

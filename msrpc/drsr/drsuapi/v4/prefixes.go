package drsuapi

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"sort"
)

type PrefixTable map[uint32][]byte

func (o PrefixTable) SchemaPrefixTable() *SchemaPrefixTable {

	ret := SchemaPrefixTable{
		PrefixEntry: make([]*PrefixTableEntry, 0, len(o)),
	}

	for i, prefix := range o {
		entry := &PrefixTableEntry{
			Index: uint32(i),
			Prefix: &OID{
				Elements: prefix,
			},
		}
		ret.PrefixEntry = append(ret.PrefixEntry, entry)
	}

	sort.Slice(ret.PrefixEntry, func(i, j int) bool {
		return ret.PrefixEntry[i].Index < ret.PrefixEntry[j].Index
	})

	return &ret
}

func (o PrefixTable) OIDToAttribute(oid asn1.ObjectIdentifier) (uint32, error) {

	if o == nil {
		return 0, fmt.Errorf("prefix table is nil")
	}

	if len(oid) < 1 {
		return 0, fmt.Errorf("oid is too short: %d elements", len(oid))
	}

	last := oid[len(oid)-1]

	raw, err := OIDToRawPrefix(oid)
	if err != nil {
		return 0, fmt.Errorf("failed to convert oid to raw prefix: %w", err)
	}

	if last < 128 {
		if len(raw) < 1 {
			return 0, fmt.Errorf("raw prefix is too short: %d bytes", len(raw))
		}
		raw = raw[:len(raw)-1]
	} else {
		if len(raw) < 2 {
			return 0, fmt.Errorf("raw prefix is too short: %d bytes", len(raw))
		}
		raw = raw[:len(raw)-2]
	}

	lower := last % 16384
	if last > 16384 {
		lower += 32768
	}

	upper, found := uint32(0), false

	for i, prefix := range o {
		if bytes.Equal(prefix, raw) {
			upper, found = i, true
			break
		}
		if i >= upper {
			// use the next available index. (in case if table does not contain monotonically increasing indices)
			upper = i + 1
		}
	}

	if !found {
		o[upper] = raw
	}

	return upper<<16 | uint32(lower), nil
}

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

func OIDToRawPrefix(oid asn1.ObjectIdentifier) ([]byte, error) {

	b, err := asn1.Marshal(oid)
	if err != nil {
		return nil, err
	}

	raw := asn1.RawValue{}
	if _, err := asn1.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if raw.Tag != asn1.TagOID {
		return nil, fmt.Errorf("object identifier does not have OID tag")
	}

	return raw.Bytes, nil
}

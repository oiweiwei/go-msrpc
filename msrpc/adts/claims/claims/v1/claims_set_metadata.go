package claims

import (
	"fmt"

	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

// Claims function returns the parsed claims set from the ClaimsSetMetadata.
func (o *ClaimsSetMetadata) Claims(opts ...any) (*ClaimsSet, error) {

	var (
		cls ClaimsSet
		raw []byte
		err error
	)

	if o.CompressionFormat != ClaimsCompressionFormatNone {
		for _, opt := range opts {
			if opt, ok := opt.(Compression); ok && opt.Format() == o.CompressionFormat {
				if raw, err = opt.Decompress(o.ClaimsSet); err != nil {
					return nil, fmt.Errorf("claims set decompression failed: %w", err)
				}
			}
		}
	} else {
		raw = o.ClaimsSet
	}

	if raw == nil && o.ClaimsSet != nil && o.CompressionFormat != ClaimsCompressionFormatNone {
		return nil, fmt.Errorf("claims set is compressed but no decompression method provided")
	}

	if err := ndr.UnmarshalWithTypeSerializationV1(raw, ndr.UnmarshalerPointer(&cls)); err != nil {
		return nil, err
	}

	return &cls, nil
}

// Compression is an interface for compressing and decompressing byte slices.
type Compression interface {
	// Compress compresses a byte slice.
	Compress([]byte) ([]byte, error)
	// Decompress decompresses a byte slice.
	Decompress([]byte) ([]byte, error)
	// Format returns the compression format.
	Format() ClaimsCompressionFormat
}

// BuildClaimsSetMetadata builds a ClaimsSetMetadata from a ClaimsSet.
// The ClaimsSet is serialized and optionally compressed based on the provided options.
func BuildClaimsSetMetadata(claims *ClaimsSet, opts ...any) (*ClaimsSetMetadata, error) {

	var (
		cls ClaimsSetMetadata
		raw []byte
		err error
	)

	if raw, err = ndr.MarshalWithTypeSerializationV1(claims); err != nil {
		return nil, err
	}

	cls.UncompressedClaimsSetSize = uint32(len(raw))

	for _, opt := range opts {
		if opt, ok := opt.(Compression); ok {
			if raw, err = opt.Compress(raw); err != nil {
				return nil, fmt.Errorf("claims set compression failed: %w", err)
			}
			cls.CompressionFormat = opt.Format()
			break
		}
	}

	cls.ClaimsSet = raw

	return &cls, nil
}

type ClaimsMap struct {
	SourceType ClaimsSourceType `json:"source_type"`
	Claims     map[string]any   `json:"claims"`
}

// Map returns a map of claims.
func (o *ClaimsSet) ClaimsMaps() []*ClaimsMap {

	ret := make([]*ClaimsMap, 0, len(o.ClaimsArrays))

	for i1 := range o.ClaimsArrays {
		m := make(map[string]any)
		for i2 := range o.ClaimsArrays[i1].ClaimEntries {

			entry := o.ClaimsArrays[i1].ClaimEntries[i2]

			value := entry.Values.GetValue()
			if value == nil {
				continue
			}

			switch value := entry.Values.GetValue().(type) {
			case *ClaimEntry_Values_ClaimEntryString:
				if len(value.StringValues) == 1 {
					m[entry.ID] = value.StringValues[0]
				} else {
					m[entry.ID] = value.StringValues
				}
			case *ClaimEntry_Values_ClaimEntryUint64:
				if len(value.Uint64Values) == 1 {
					m[entry.ID] = value.Uint64Values[0]
				} else {
					m[entry.ID] = value.Uint64Values
				}
			case *ClaimEntry_Values_ClaimEntryInt64:
				if len(value.Int64Values) == 1 {
					m[entry.ID] = value.Int64Values[0]
				} else {
					m[entry.ID] = value.Int64Values
				}
			case *ClaimEntry_Values_ClaimEntryBoolean:
				vv := make([]bool, len(value.BooleanValues))
				for i := range value.BooleanValues {
					vv[i] = value.BooleanValues[i] != 0
				}
				if len(vv) == 1 {
					m[entry.ID] = vv[0]
				} else {
					m[entry.ID] = vv
				}
			}
		}

		if len(m) > 0 {
			ret = append(ret, &ClaimsMap{Claims: m, SourceType: o.ClaimsArrays[i1].ClaimsSourceType})
		}
	}

	return ret
}

// ClaimsSetFromClaimsMaps function returns a ClaimsSet from a slice of ClaimsMap.
func ClaimsSetFromClaimsMaps(ms ...*ClaimsMap) *ClaimsSet {

	claimsSet := &ClaimsSet{}

	for _, m := range ms {

		array := &ClaimsArray{ClaimsSourceType: m.SourceType}

		for k, v := range m.Claims {

			var val is_ClaimEntry_Values

			switch v := v.(type) {
			case string:
				val = &ClaimEntry_Values_ClaimEntryString{StringValues: []string{v}}
			case []string:
				val = &ClaimEntry_Values_ClaimEntryString{StringValues: v}
			case uint64:
				val = &ClaimEntry_Values_ClaimEntryUint64{Uint64Values: []uint64{v}}
			case []uint64:
				val = &ClaimEntry_Values_ClaimEntryUint64{Uint64Values: v}
			case int64:
				val = &ClaimEntry_Values_ClaimEntryInt64{Int64Values: []int64{v}}
			case []int64:
				val = &ClaimEntry_Values_ClaimEntryInt64{Int64Values: v}
			case bool:
				var vv uint64
				if v {
					vv = 1
				}
				val = &ClaimEntry_Values_ClaimEntryBoolean{BooleanValues: []uint64{vv}}
			case []bool:
				vv := make([]uint64, len(v))
				for i := range v {
					if v[i] {
						vv[i] = 1
					}
				}
				val = &ClaimEntry_Values_ClaimEntryBoolean{BooleanValues: vv}
			case int:
				val = &ClaimEntry_Values_ClaimEntryInt64{Int64Values: []int64{int64(v)}}
			case []int:
				vv := make([]int64, len(v))
				for i := range v {
					vv[i] = int64(v[i])
				}
				val = &ClaimEntry_Values_ClaimEntryInt64{Int64Values: vv}
			}

			array.ClaimEntries = append(array.ClaimEntries, &ClaimEntry{ID: k, Values: &ClaimEntry_Values{Value: val}})
		}

		claimsSet.ClaimsArrays = append(claimsSet.ClaimsArrays, array)
	}

	return claimsSet
}

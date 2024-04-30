package claims

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
)

var (
	// import guard
	GoPackage = "adts/claims"
)

var (
	// Syntax UUID
	ClaimsSyntaxUUID = &uuid.UUID{TimeLow: 0xbba9cb76, TimeMid: 0xeb0c, TimeHiAndVersion: 0x462c, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x1b, Node: [6]uint8{0x5d, 0x8c, 0x34, 0x41, 0x57, 0x1}}
	// Syntax ID
	ClaimsSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ClaimsSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// Claims interface.
type ClaimsClient interface {

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

// ClaimType type represents CLAIM_TYPE RPC enumeration.
//
// The CLAIM_TYPE enumeration enumerates various value types of a claim.
type ClaimType uint16

var (
	// CLAIM_TYPE_INT64:  The value type of the claim is LONG64.
	ClaimTypeInt64 ClaimType = 1
	// CLAIM_TYPE_UINT64:  The value type of the claim is ULONG64.
	ClaimTypeUint64 ClaimType = 2
	// CLAIM_TYPE_STRING:  The value type of the claim is a null-terminated string of Unicode
	// characters.
	ClaimTypeString ClaimType = 3
	// CLAIM_TYPE_BOOLEAN:  The value type of the claim is ULONG64; a value is set to 1
	// to specify TRUE, or 0 to specify FALSE.
	ClaimTypeBoolean ClaimType = 6
)

func (o ClaimType) String() string {
	switch o {
	case ClaimTypeInt64:
		return "ClaimTypeInt64"
	case ClaimTypeUint64:
		return "ClaimTypeUint64"
	case ClaimTypeString:
		return "ClaimTypeString"
	case ClaimTypeBoolean:
		return "ClaimTypeBoolean"
	}
	return "Invalid"
}

// ClaimsSourceType type represents CLAIMS_SOURCE_TYPE RPC enumeration.
//
// The CLAIMS_SOURCE_TYPE enumeration specifies the source of the claims.
//
// Note  No semantics are to be attached to these values other than those specified
// in section 3.
type ClaimsSourceType uint16

var (
	ClaimsSourceTypeAD          ClaimsSourceType = 1
	ClaimsSourceTypeCertificate ClaimsSourceType = 2
)

func (o ClaimsSourceType) String() string {
	switch o {
	case ClaimsSourceTypeAD:
		return "ClaimsSourceTypeAD"
	case ClaimsSourceTypeCertificate:
		return "ClaimsSourceTypeCertificate"
	}
	return "Invalid"
}

// ClaimsCompressionFormat type represents CLAIMS_COMPRESSION_FORMAT RPC enumeration.
//
// The CLAIMS_COMPRESSION_FORMAT enumeration specifies the source of the compression
// algorithm that is used for encoding claims in a CLAIMS_SET_METADATA structure.
type ClaimsCompressionFormat uint16

var (
	// COMPRESSION_FORMAT_NONE:  No compression.
	ClaimsCompressionFormatNone ClaimsCompressionFormat = 0
	// COMPRESSION_FORMAT_LZNT1:  The LZNT1 compression algorithm is used. For more information,
	// see [MS-XCA] section 2.5.
	ClaimsCompressionFormatLZNT1 ClaimsCompressionFormat = 2
	// COMPRESSION_FORMAT_XPRESS:  The Xpress LZ77 compression algorithm is used. For more
	// information, see [MS-XCA] sections 2.3 and 2.4.
	ClaimsCompressionFormatXPress ClaimsCompressionFormat = 3
	// COMPRESSION_FORMAT_XPRESS_HUFF:  The Xpress LZ77+Huffman compression algorithm is
	// used. For more details, see [MS-XCA] sections 2.1 and 2.2.
	ClaimsCompressionFormatXPressHuff ClaimsCompressionFormat = 4
)

func (o ClaimsCompressionFormat) String() string {
	switch o {
	case ClaimsCompressionFormatNone:
		return "ClaimsCompressionFormatNone"
	case ClaimsCompressionFormatLZNT1:
		return "ClaimsCompressionFormatLZNT1"
	case ClaimsCompressionFormatXPress:
		return "ClaimsCompressionFormatXPress"
	case ClaimsCompressionFormatXPressHuff:
		return "ClaimsCompressionFormatXPressHuff"
	}
	return "Invalid"
}

// ClaimEntry structure represents CLAIM_ENTRY RPC structure.
//
// The CLAIM_ENTRY structure specifies a single claim.
type ClaimEntry struct {
	// Id:  Specifies the claim identifier.
	ID string `idl:"name:Id" json:"id"`
	// Type:  Specifies the type of the data in the Values union. Refer to section 2.2.18.2
	// for allowed values and their interpretation.
	Type ClaimType `idl:"name:Type" json:"type"`
	// Values:  A union of arrays of the various types of claim values that a CLAIM_ENTRY
	// can contain. The actual type of the elements is specified by the Type member.
	//
	// ValueCount:  Specifies the number of array elements in the Int64Values member.
	//
	// Int64Values:  An array of LONG64 values of the claim. The array has ValueCount elements.
	//
	// ValueCount:  Specifies the number of array elements in the Uint64Values member.
	//
	// Uint64Values:  An array of ULONG64 values of the claim. The array has ValueCount
	// elements.
	//
	// ValueCount:  Specifies the number of array elements in the StringValues member.
	//
	// StringValues:  An array of null-terminated, Unicode string values of the claim. The
	// array has ValueCount elements.
	//
	// ValueCount:  Specifies the number of array elements in the BooleanValues member.
	Values *ClaimEntry_Values `idl:"name:Values;switch_is:Type" json:"values"`
}

func (o *ClaimEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ID != "" {
		_ptr_Id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ID, _ptr_Id); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	_swValues := uint16(o.Type)
	if o.Values != nil {
		if err := o.Values.MarshalUnionNDR(ctx, w, _swValues); err != nil {
			return err
		}
	} else {
		if err := (&ClaimEntry_Values{}).MarshalUnionNDR(ctx, w, _swValues); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_Id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ID); err != nil {
			return err
		}
		return nil
	})
	_s_Id := func(ptr interface{}) { o.ID = *ptr.(*string) }
	if err := w.ReadPointer(&o.ID, _s_Id, _ptr_Id); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.Values == nil {
		o.Values = &ClaimEntry_Values{}
	}
	_swValues := uint16(o.Type)
	if err := o.Values.UnmarshalUnionNDR(ctx, w, _swValues); err != nil {
		return err
	}
	return nil
}

// ClaimEntry_Values structure represents CLAIM_ENTRY union anonymous member.
//
// The CLAIM_ENTRY structure specifies a single claim.
type ClaimEntry_Values struct {
	// Types that are assignable to Value
	//
	// *ClaimEntry_Values_ClaimEntryInt64
	// *ClaimEntry_Values_ClaimEntryUint64
	// *ClaimEntry_Values_ClaimEntryString
	// *ClaimEntry_Values_ClaimEntryBoolean
	// *ClaimEntry_Values_DefaultClaimEntry
	Value is_ClaimEntry_Values `json:"value"`
}

func (o *ClaimEntry_Values) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ClaimEntry_Values_ClaimEntryInt64:
		if value != nil {
			return value.ClaimEntryInt64
		}
	case *ClaimEntry_Values_ClaimEntryUint64:
		if value != nil {
			return value.ClaimEntryUint64
		}
	case *ClaimEntry_Values_ClaimEntryString:
		if value != nil {
			return value.ClaimEntryString
		}
	case *ClaimEntry_Values_ClaimEntryBoolean:
		if value != nil {
			return value.ClaimEntryBoolean
		}
	}
	return nil
}

type is_ClaimEntry_Values interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ClaimEntry_Values()
}

func (o *ClaimEntry_Values) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ClaimEntry_Values_ClaimEntryInt64:
		return uint16(1)
	case *ClaimEntry_Values_ClaimEntryUint64:
		return uint16(2)
	case *ClaimEntry_Values_ClaimEntryString:
		return uint16(3)
	case *ClaimEntry_Values_ClaimEntryBoolean:
		return uint16(6)
	}
	return uint16(0)
}

func (o *ClaimEntry_Values) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*ClaimEntry_Values_ClaimEntryInt64)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClaimEntry_Values_ClaimEntryInt64{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*ClaimEntry_Values_ClaimEntryUint64)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClaimEntry_Values_ClaimEntryUint64{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*ClaimEntry_Values_ClaimEntryString)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClaimEntry_Values_ClaimEntryString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(6):
		_o, _ := o.Value.(*ClaimEntry_Values_ClaimEntryBoolean)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ClaimEntry_Values_ClaimEntryBoolean{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		// no-op
	}
	return nil
}

func (o *ClaimEntry_Values) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &ClaimEntry_Values_ClaimEntryInt64{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &ClaimEntry_Values_ClaimEntryUint64{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &ClaimEntry_Values_ClaimEntryString{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(6):
		o.Value = &ClaimEntry_Values_ClaimEntryBoolean{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &ClaimEntry_Values_DefaultClaimEntry{}
	}
	return nil
}

// ClaimEntry_Values_ClaimEntryInt64 structure represents ClaimEntry_Values RPC union arm.
//
// It has following labels: 1
type ClaimEntry_Values_ClaimEntryInt64 struct {
	ClaimEntryInt64 *ClaimEntry_Values_ClaimEntryInt64 `idl:"name:CLAIM_ENTRY_INT64" json:"claim_entry_int64"`
}

func (*ClaimEntry_Values_ClaimEntryInt64) is_ClaimEntry_Values() {}

func (o *ClaimEntry_Values_ClaimEntryInt64) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClaimEntryInt64 != nil {
		if err := o.ClaimEntryInt64.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClaimEntry_Values_ClaimEntryInt64{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimEntry_Values_ClaimEntryInt64) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ClaimEntryInt64 == nil {
		o.ClaimEntryInt64 = &ClaimEntry_Values_ClaimEntryInt64{}
	}
	if err := o.ClaimEntryInt64.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClaimEntry_Values_ClaimEntryUint64 structure represents ClaimEntry_Values RPC union arm.
//
// It has following labels: 2
type ClaimEntry_Values_ClaimEntryUint64 struct {
	ClaimEntryUint64 *ClaimEntry_Values_ClaimEntryUint64 `idl:"name:CLAIM_ENTRY_UINT64" json:"claim_entry_uint64"`
}

func (*ClaimEntry_Values_ClaimEntryUint64) is_ClaimEntry_Values() {}

func (o *ClaimEntry_Values_ClaimEntryUint64) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClaimEntryUint64 != nil {
		if err := o.ClaimEntryUint64.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClaimEntry_Values_ClaimEntryUint64{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimEntry_Values_ClaimEntryUint64) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ClaimEntryUint64 == nil {
		o.ClaimEntryUint64 = &ClaimEntry_Values_ClaimEntryUint64{}
	}
	if err := o.ClaimEntryUint64.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClaimEntry_Values_ClaimEntryString structure represents ClaimEntry_Values RPC union arm.
//
// It has following labels: 3
type ClaimEntry_Values_ClaimEntryString struct {
	ClaimEntryString *ClaimEntry_Values_ClaimEntryString `idl:"name:CLAIM_ENTRY_STRING" json:"claim_entry_string"`
}

func (*ClaimEntry_Values_ClaimEntryString) is_ClaimEntry_Values() {}

func (o *ClaimEntry_Values_ClaimEntryString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClaimEntryString != nil {
		if err := o.ClaimEntryString.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClaimEntry_Values_ClaimEntryString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimEntry_Values_ClaimEntryString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ClaimEntryString == nil {
		o.ClaimEntryString = &ClaimEntry_Values_ClaimEntryString{}
	}
	if err := o.ClaimEntryString.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClaimEntry_Values_ClaimEntryBoolean structure represents ClaimEntry_Values RPC union arm.
//
// It has following labels: 6
type ClaimEntry_Values_ClaimEntryBoolean struct {
	ClaimEntryBoolean *ClaimEntry_Values_ClaimEntryBoolean `idl:"name:CLAIM_ENTRY_BOOLEAN" json:"claim_entry_boolean"`
}

func (*ClaimEntry_Values_ClaimEntryBoolean) is_ClaimEntry_Values() {}

func (o *ClaimEntry_Values_ClaimEntryBoolean) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClaimEntryBoolean != nil {
		if err := o.ClaimEntryBoolean.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClaimEntry_Values_ClaimEntryBoolean{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimEntry_Values_ClaimEntryBoolean) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ClaimEntryBoolean == nil {
		o.ClaimEntryBoolean = &ClaimEntry_Values_ClaimEntryBoolean{}
	}
	if err := o.ClaimEntryBoolean.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClaimEntry_Values_DefaultClaimEntry structure represents ClaimEntry_Values RPC default union arm.
type ClaimEntry_Values_DefaultClaimEntry struct {
}

func (*ClaimEntry_Values_DefaultClaimEntry) is_ClaimEntry_Values() {}

func (o *ClaimEntry_Values_DefaultClaimEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *ClaimEntry_Values_DefaultClaimEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// ClaimsArray structure represents CLAIMS_ARRAY RPC structure.
//
// The CLAIMS_ARRAY structure specifies an array of CLAIM_ENTRY structures and the associated
// claims source type.
type ClaimsArray struct {
	// usClaimsSourceType:  Specifies the source of the claims.
	ClaimsSourceType ClaimsSourceType `idl:"name:usClaimsSourceType" json:"claims_source_type"`
	// ulClaimsCount:  Specifies the number of CLAIM_ENTRY elements in the ClaimEntries
	// member of this structure.
	ClaimsCount uint32 `idl:"name:ulClaimsCount" json:"claims_count"`
	// ClaimEntries:  An array that contains ulClaimsCount number of CLAIM_ENTRY elements.
	ClaimEntries []*ClaimEntry `idl:"name:ClaimEntries;size_is:(ulClaimsCount)" json:"claim_entries"`
}

func (o *ClaimsArray) xxx_PreparePayload(ctx context.Context) error {
	if o.ClaimEntries != nil && o.ClaimsCount == 0 {
		o.ClaimsCount = uint32(len(o.ClaimEntries))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimsArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.ClaimsSourceType)); err != nil {
		return err
	}
	if err := w.WriteData(o.ClaimsCount); err != nil {
		return err
	}
	if o.ClaimEntries != nil || o.ClaimsCount > 0 {
		_ptr_ClaimEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ClaimsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ClaimEntries {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ClaimEntries[i1] != nil {
					if err := o.ClaimEntries[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ClaimEntry{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ClaimEntries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ClaimEntry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ClaimEntries, _ptr_ClaimEntries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimsArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.ClaimsSourceType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClaimsCount); err != nil {
		return err
	}
	_ptr_ClaimEntries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ClaimsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ClaimsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ClaimEntries", sizeInfo[0])
		}
		o.ClaimEntries = make([]*ClaimEntry, sizeInfo[0])
		for i1 := range o.ClaimEntries {
			i1 := i1
			if o.ClaimEntries[i1] == nil {
				o.ClaimEntries[i1] = &ClaimEntry{}
			}
			if err := o.ClaimEntries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ClaimEntries := func(ptr interface{}) { o.ClaimEntries = *ptr.(*[]*ClaimEntry) }
	if err := w.ReadPointer(&o.ClaimEntries, _s_ClaimEntries, _ptr_ClaimEntries); err != nil {
		return err
	}
	return nil
}

// ClaimsSet structure represents CLAIMS_SET RPC structure.
//
// The CLAIMS_SET structure specifies CLAIMS_ARRAY structures, each from a different
// claims source.
type ClaimsSet struct {
	// ulClaimsArrayCount:  Specifies the number of CLAIMS_ARRAY elements that are in the
	// ClaimsArrays member. This field MUST be greater than or equal to 1.
	ClaimsArrayCount uint32 `idl:"name:ulClaimsArrayCount" json:"claims_array_count"`
	// ClaimsArrays:  An array containing ulClaimsArrayCount number of CLAIMS_ARRAY structures.
	ClaimsArrays []*ClaimsArray `idl:"name:ClaimsArrays;size_is:(ulClaimsArrayCount)" json:"claims_arrays"`
	// usReservedType:  This field is not used.
	_ uint16 `idl:"name:usReservedType"`
	// ulReservedFieldSize:  Specifies the length, in bytes, of the ReservedField member.
	_ uint32 `idl:"name:ulReservedFieldSize"`
	// ReservedField:  A byte array containing ulReservedFieldSize bytes.
	_ []byte `idl:"name:ReservedField;size_is:(ulReservedFieldSize)"`
}

func (o *ClaimsSet) xxx_PreparePayload(ctx context.Context) error {
	if o.ClaimsArrays != nil && o.ClaimsArrayCount == 0 {
		o.ClaimsArrayCount = uint32(len(o.ClaimsArrays))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimsSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClaimsArrayCount); err != nil {
		return err
	}
	if o.ClaimsArrays != nil || o.ClaimsArrayCount > 0 {
		_ptr_ClaimsArrays := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ClaimsArrayCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ClaimsArrays {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ClaimsArrays[i1] != nil {
					if err := o.ClaimsArrays[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ClaimsArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ClaimsArrays); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ClaimsArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ClaimsArrays, _ptr_ClaimsArrays); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved usReservedType
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved ulReservedFieldSize
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved ReservedField
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	return nil
}
func (o *ClaimsSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClaimsArrayCount); err != nil {
		return err
	}
	_ptr_ClaimsArrays := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ClaimsArrayCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ClaimsArrayCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ClaimsArrays", sizeInfo[0])
		}
		o.ClaimsArrays = make([]*ClaimsArray, sizeInfo[0])
		for i1 := range o.ClaimsArrays {
			i1 := i1
			if o.ClaimsArrays[i1] == nil {
				o.ClaimsArrays[i1] = &ClaimsArray{}
			}
			if err := o.ClaimsArrays[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ClaimsArrays := func(ptr interface{}) { o.ClaimsArrays = *ptr.(*[]*ClaimsArray) }
	if err := w.ReadPointer(&o.ClaimsArrays, _s_ClaimsArrays, _ptr_ClaimsArrays); err != nil {
		return err
	}
	// reserved usReservedType
	var _usReservedType uint16
	if err := w.ReadData(&_usReservedType); err != nil {
		return err
	}
	// reserved ulReservedFieldSize
	var _ulReservedFieldSize uint32
	if err := w.ReadData(&_ulReservedFieldSize); err != nil {
		return err
	}
	// reserved ReservedField
	var _ReservedField []byte
	_ptr_ReservedField := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _ReservedField", sizeInfo[0])
		}
		_ReservedField = make([]byte, sizeInfo[0])
		for i1 := range _ReservedField {
			i1 := i1
			if err := w.ReadData(&_ReservedField[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ReservedField := func(ptr interface{}) { _ReservedField = *ptr.(*[]byte) }
	if err := w.ReadPointer(&_ReservedField, _s_ReservedField, _ptr_ReservedField); err != nil {
		return err
	}
	return nil
}

// ClaimsSetMetadata structure represents CLAIMS_SET_METADATA RPC structure.
//
// The CLAIMS_SET_METADATA structure specifies an encoded CLAIMS_SET structure with
// information about the encoding.
type ClaimsSetMetadata struct {
	// ulClaimsSetSize:  Contains the size, in bytes, of the ClaimsSet member.
	ClaimsSetSize uint32 `idl:"name:ulClaimsSetSize" json:"claims_set_size"`
	// ClaimsSet:  A byte array of length ulClaimsSetSize bytes. This field contains a CLAIMS_SET
	// structure that is encoded as described in section 3.1.1.11.2.5.
	ClaimsSet []byte `idl:"name:ClaimsSet;size_is:(ulClaimsSetSize)" json:"claims_set"`
	// usCompressionFormat:  Specifies the compression algorithm used for encoding a CLAIMS_SET
	// structure, as specified in section 3.1.1.11.2.5.
	CompressionFormat ClaimsCompressionFormat `idl:"name:usCompressionFormat" json:"compression_format"`
	// ulUncompressedClaimsSetSize:  Specifies the size of the partially encoded CLAIMS_SET
	// structure before compression, the fully encoded version of which is stored in the
	// ClaimsSet member.
	UncompressedClaimsSetSize uint32 `idl:"name:ulUncompressedClaimsSetSize" json:"uncompressed_claims_set_size"`
	// usReservedType:  The server MUST set this member to 0. The client MUST ignore this
	// member.
	_ uint16 `idl:"name:usReservedType"`
	// ulReservedFieldSize:  Specifies the size, in bytes, of the ReservedField member.
	_ uint32 `idl:"name:ulReservedFieldSize"`
	// ReservedField:  A byte array containing ulReservedFieldSize elements.
	_ []byte `idl:"name:ReservedField;size_is:(ulReservedFieldSize)"`
}

func (o *ClaimsSetMetadata) xxx_PreparePayload(ctx context.Context) error {
	if o.ClaimsSet != nil && o.ClaimsSetSize == 0 {
		o.ClaimsSetSize = uint32(len(o.ClaimsSet))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClaimsSetMetadata) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClaimsSetSize); err != nil {
		return err
	}
	if o.ClaimsSet != nil || o.ClaimsSetSize > 0 {
		_ptr_ClaimsSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ClaimsSetSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ClaimsSet {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ClaimsSet[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ClaimsSet); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ClaimsSet, _ptr_ClaimsSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.CompressionFormat)); err != nil {
		return err
	}
	if err := w.WriteData(o.UncompressedClaimsSetSize); err != nil {
		return err
	}
	// reserved usReservedType
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved ulReservedFieldSize
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved ReservedField
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	return nil
}
func (o *ClaimsSetMetadata) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClaimsSetSize); err != nil {
		return err
	}
	_ptr_ClaimsSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ClaimsSetSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ClaimsSetSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ClaimsSet", sizeInfo[0])
		}
		o.ClaimsSet = make([]byte, sizeInfo[0])
		for i1 := range o.ClaimsSet {
			i1 := i1
			if err := w.ReadData(&o.ClaimsSet[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ClaimsSet := func(ptr interface{}) { o.ClaimsSet = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ClaimsSet, _s_ClaimsSet, _ptr_ClaimsSet); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.CompressionFormat)); err != nil {
		return err
	}
	if err := w.ReadData(&o.UncompressedClaimsSetSize); err != nil {
		return err
	}
	// reserved usReservedType
	var _usReservedType uint16
	if err := w.ReadData(&_usReservedType); err != nil {
		return err
	}
	// reserved ulReservedFieldSize
	var _ulReservedFieldSize uint32
	if err := w.ReadData(&_ulReservedFieldSize); err != nil {
		return err
	}
	// reserved ReservedField
	var _ReservedField []byte
	_ptr_ReservedField := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _ReservedField", sizeInfo[0])
		}
		_ReservedField = make([]byte, sizeInfo[0])
		for i1 := range _ReservedField {
			i1 := i1
			if err := w.ReadData(&_ReservedField[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ReservedField := func(ptr interface{}) { _ReservedField = *ptr.(*[]byte) }
	if err := w.ReadPointer(&_ReservedField, _s_ReservedField, _ptr_ReservedField); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultClaimsClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultClaimsClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewClaimsClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClaimsClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClaimsSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultClaimsClient{cc: cc}, nil
}

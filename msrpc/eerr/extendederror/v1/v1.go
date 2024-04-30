package extendederror

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
	GoPackage = "eerr"
)

var (
	// Syntax UUID
	ExtendedErrorSyntaxUUID = &uuid.UUID{TimeLow: 0x14a8831c, TimeMid: 0xbc82, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x64, Node: [6]uint8{0x0, 0x8, 0xc7, 0x45, 0x7e, 0x5d}}
	// Syntax ID
	ExtendedErrorSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ExtendedErrorSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// ExtendedError interface.
type ExtendedErrorClient interface {

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

// ANSIString structure represents EEAString RPC structure.
//
// The EEAString structure encodes strings of ANSI characters, as specified in [ISO/IEC-8859-1],
// that contain troubleshooting information.
type ANSIString struct {
	// nLength:  This field MUST contain the size of pString in bytes.
	Length int16 `idl:"name:nLength" json:"length"`
	// pString:  A NULL-terminated ANSI string that contains troubleshooting information.
	String []byte `idl:"name:pString;size_is:(nLength)" json:"string"`
}

func (o *ANSIString) xxx_PreparePayload(ctx context.Context) error {
	if o.String != nil && o.Length == 0 {
		o.Length = int16(len(o.String))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ANSIString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.String != nil || o.Length > 0 {
		_ptr_pString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.String {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.String[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.String); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_pString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ANSIString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_pString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.String", sizeInfo[0])
		}
		o.String = make([]byte, sizeInfo[0])
		for i1 := range o.String {
			i1 := i1
			if err := w.ReadData(&o.String[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pString := func(ptr interface{}) { o.String = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.String, _s_pString, _ptr_pString); err != nil {
		return err
	}
	return nil
}

// UnicodeString structure represents EEUString RPC structure.
//
// The EEUString structure encodes Unicode strings that contain troubleshooting information.
// The EEComputerName structure uses this type.
type UnicodeString struct {
	// nLength:  This field MUST contain the length of pString in characters.
	Length int16 `idl:"name:nLength" json:"length"`
	// pString:  A NULL-terminated Unicode string that contains troubleshooting information.
	String []uint16 `idl:"name:pString;size_is:(nLength)" json:"string"`
}

func (o *UnicodeString) xxx_PreparePayload(ctx context.Context) error {
	if o.String != nil && o.Length == 0 {
		o.Length = int16(len(o.String))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.String != nil || o.Length > 0 {
		_ptr_pString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.String {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.String[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.String); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_pString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_pString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.String", sizeInfo[0])
		}
		o.String = make([]uint16, sizeInfo[0])
		for i1 := range o.String {
			i1 := i1
			if err := w.ReadData(&o.String[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pString := func(ptr interface{}) { o.String = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.String, _s_pString, _ptr_pString); err != nil {
		return err
	}
	return nil
}

// BinaryInfo structure represents BinaryEEInfo RPC structure.
//
// The BinaryEEInfo structure encodes binary data that contains troubleshooting information.
type BinaryInfo struct {
	// nSize:  This field MUST contain the size of pBlob in bytes.
	Size int16 `idl:"name:nSize" json:"size"`
	// pBlob:  Binary data that contains troubleshooting information.
	Blob []byte `idl:"name:pBlob;size_is:(nSize)" json:"blob"`
}

func (o *BinaryInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.Blob != nil && o.Size == 0 {
		o.Size = int16(len(o.Blob))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BinaryInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Blob != nil || o.Size > 0 {
		_ptr_pBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Blob {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Blob[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Blob, _ptr_pBlob); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BinaryInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_pBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
		}
		o.Blob = make([]byte, sizeInfo[0])
		for i1 := range o.Blob {
			i1 := i1
			if err := w.ReadData(&o.Blob[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pBlob := func(ptr interface{}) { o.Blob = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Blob, _s_pBlob, _ptr_pBlob); err != nil {
		return err
	}
	return nil
}

// ExtendedErrorParamTypesInternal type represents ExtendedErrorParamTypesInternal RPC enumeration.
//
// The ExtendedErrorParamTypesInternal enumeration defines the values that are valid
// for the Type field in the ExtendedErrorParam structure.
type ExtendedErrorParamTypesInternal uint16

var (
	// eeptiAnsiString:  The ANSIString member of the union is valid.
	ExtendedErrorParamTypesInternalANSIString ExtendedErrorParamTypesInternal = 1
	// eeptiUnicodeString:  The UnicodeString member of the union is valid.
	ExtendedErrorParamTypesInternalUnicodeString ExtendedErrorParamTypesInternal = 2
	// eeptiLongVal:  The LVal member of the union is valid. LVal is used to encode a long.
	ExtendedErrorParamTypesInternalLongValue ExtendedErrorParamTypesInternal = 3
	// eeptiShortValue:  The IVal member of the union is valid. IVal is used to encode
	// a short.
	ExtendedErrorParamTypesInternalShortValue ExtendedErrorParamTypesInternal = 4
	// eeptiPointerValue:  The PVal member of the union is valid. PVal is used to encode
	// an __int64.
	ExtendedErrorParamTypesInternalPointerValue ExtendedErrorParamTypesInternal = 5
	// eeptiNone:  No additional details are present in this parameter.
	ExtendedErrorParamTypesInternalNone ExtendedErrorParamTypesInternal = 6
	// eeptiBinary:  The Blob member of the union is valid.
	ExtendedErrorParamTypesInternalBinary ExtendedErrorParamTypesInternal = 7
)

func (o ExtendedErrorParamTypesInternal) String() string {
	switch o {
	case ExtendedErrorParamTypesInternalANSIString:
		return "ExtendedErrorParamTypesInternalANSIString"
	case ExtendedErrorParamTypesInternalUnicodeString:
		return "ExtendedErrorParamTypesInternalUnicodeString"
	case ExtendedErrorParamTypesInternalLongValue:
		return "ExtendedErrorParamTypesInternalLongValue"
	case ExtendedErrorParamTypesInternalShortValue:
		return "ExtendedErrorParamTypesInternalShortValue"
	case ExtendedErrorParamTypesInternalPointerValue:
		return "ExtendedErrorParamTypesInternalPointerValue"
	case ExtendedErrorParamTypesInternalNone:
		return "ExtendedErrorParamTypesInternalNone"
	case ExtendedErrorParamTypesInternalBinary:
		return "ExtendedErrorParamTypesInternalBinary"
	}
	return "Invalid"
}

// ExtendedErrorParam structure represents ExtendedErrorParam RPC structure.
//
// The ExtendedErrorParam structure contains a parameter, as described in section 1.3.1,
// that provides additional details about the error record.
type ExtendedErrorParam struct {
	// Type:  Indicates which member of the union is valid. ExtendedErrorParamTypesInternal
	// lists all of the possible values.
	Type               ExtendedErrorParamTypesInternal        `idl:"name:Type" json:"type"`
	ExtendedErrorParam *ExtendedErrorParam_ExtendedErrorParam `idl:"name:ExtendedErrorParam;switch_is:Type" json:"extended_error_param"`
}

func (o *ExtendedErrorParam) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtendedErrorParam) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	_swExtendedErrorParam := int16(o.Type)
	if o.ExtendedErrorParam != nil {
		if err := o.ExtendedErrorParam.MarshalUnionNDR(ctx, w, _swExtendedErrorParam); err != nil {
			return err
		}
	} else {
		if err := (&ExtendedErrorParam_ExtendedErrorParam{}).MarshalUnionNDR(ctx, w, _swExtendedErrorParam); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtendedErrorParam) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.ExtendedErrorParam == nil {
		o.ExtendedErrorParam = &ExtendedErrorParam_ExtendedErrorParam{}
	}
	_swExtendedErrorParam := int16(o.Type)
	if err := o.ExtendedErrorParam.UnmarshalUnionNDR(ctx, w, _swExtendedErrorParam); err != nil {
		return err
	}
	return nil
}

// ExtendedErrorParam_ExtendedErrorParam structure represents ExtendedErrorParam union anonymous member.
//
// The ExtendedErrorParam structure contains a parameter, as described in section 1.3.1,
// that provides additional details about the error record.
type ExtendedErrorParam_ExtendedErrorParam struct {
	// Types that are assignable to Value
	//
	// *ExtendedErrorParam_ANSIString
	// *ExtendedErrorParam_UnicodeString
	// *ExtendedErrorParam_LValue
	// *ExtendedErrorParam_Value
	// *ExtendedErrorParam_PValue
	// *ExtendedErrorParam_6
	// *ExtendedErrorParam_Blob
	Value is_ExtendedErrorParam_ExtendedErrorParam `json:"value"`
}

func (o *ExtendedErrorParam_ExtendedErrorParam) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ExtendedErrorParam_ANSIString:
		if value != nil {
			return value.ANSIString
		}
	case *ExtendedErrorParam_UnicodeString:
		if value != nil {
			return value.UnicodeString
		}
	case *ExtendedErrorParam_LValue:
		if value != nil {
			return value.LValue
		}
	case *ExtendedErrorParam_Value:
		if value != nil {
			return value.Value
		}
	case *ExtendedErrorParam_PValue:
		if value != nil {
			return value.PValue
		}
	case *ExtendedErrorParam_Blob:
		if value != nil {
			return value.Blob
		}
	}
	return nil
}

type is_ExtendedErrorParam_ExtendedErrorParam interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ExtendedErrorParam_ExtendedErrorParam()
}

func (o *ExtendedErrorParam_ExtendedErrorParam) NDRSwitchValue(sw int16) int16 {
	if o == nil {
		return int16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ExtendedErrorParam_ANSIString:
		return int16(1)
	case *ExtendedErrorParam_UnicodeString:
		return int16(2)
	case *ExtendedErrorParam_LValue:
		return int16(3)
	case *ExtendedErrorParam_Value:
		return int16(4)
	case *ExtendedErrorParam_PValue:
		return int16(5)
	case *ExtendedErrorParam_6:
		return int16(6)
	case *ExtendedErrorParam_Blob:
		return int16(7)
	}
	return int16(0)
}

func (o *ExtendedErrorParam_ExtendedErrorParam) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw int16) error {
	if err := w.WriteSwitch(int16(sw)); err != nil {
		return err
	}
	switch sw {
	case int16(1):
		_o, _ := o.Value.(*ExtendedErrorParam_ANSIString)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExtendedErrorParam_ANSIString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int16(2):
		_o, _ := o.Value.(*ExtendedErrorParam_UnicodeString)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExtendedErrorParam_UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int16(3):
		_o, _ := o.Value.(*ExtendedErrorParam_LValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExtendedErrorParam_LValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int16(4):
		_o, _ := o.Value.(*ExtendedErrorParam_Value)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExtendedErrorParam_Value{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int16(5):
		_o, _ := o.Value.(*ExtendedErrorParam_PValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExtendedErrorParam_PValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int16(6):
	case int16(7):
		_o, _ := o.Value.(*ExtendedErrorParam_Blob)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExtendedErrorParam_Blob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ExtendedErrorParam_ExtendedErrorParam) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw int16) error {
	if err := w.ReadSwitch((*int16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case int16(1):
		o.Value = &ExtendedErrorParam_ANSIString{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int16(2):
		o.Value = &ExtendedErrorParam_UnicodeString{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int16(3):
		o.Value = &ExtendedErrorParam_LValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int16(4):
		o.Value = &ExtendedErrorParam_Value{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int16(5):
		o.Value = &ExtendedErrorParam_PValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int16(6):
		o.Value = &ExtendedErrorParam_6{}
	case int16(7):
		o.Value = &ExtendedErrorParam_Blob{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ExtendedErrorParam_ANSIString structure represents ExtendedErrorParam_ExtendedErrorParam RPC union arm.
//
// It has following labels: 1
type ExtendedErrorParam_ANSIString struct {
	// AnsiString:  A parameter of type EEAString.
	ANSIString *ANSIString `idl:"name:AnsiString" json:"ansi_string"`
}

func (*ExtendedErrorParam_ANSIString) is_ExtendedErrorParam_ExtendedErrorParam() {}

func (o *ExtendedErrorParam_ANSIString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ANSIString != nil {
		if err := o.ANSIString.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ANSIString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtendedErrorParam_ANSIString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ANSIString == nil {
		o.ANSIString = &ANSIString{}
	}
	if err := o.ANSIString.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ExtendedErrorParam_UnicodeString structure represents ExtendedErrorParam_ExtendedErrorParam RPC union arm.
//
// It has following labels: 2
type ExtendedErrorParam_UnicodeString struct {
	// UnicodeString:  A parameter of type EEUString.
	UnicodeString *UnicodeString `idl:"name:UnicodeString" json:"unicode_string"`
}

func (*ExtendedErrorParam_UnicodeString) is_ExtendedErrorParam_ExtendedErrorParam() {}

func (o *ExtendedErrorParam_UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UnicodeString != nil {
		if err := o.UnicodeString.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtendedErrorParam_UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UnicodeString == nil {
		o.UnicodeString = &UnicodeString{}
	}
	if err := o.UnicodeString.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ExtendedErrorParam_LValue structure represents ExtendedErrorParam_ExtendedErrorParam RPC union arm.
//
// It has following labels: 3
type ExtendedErrorParam_LValue struct {
	// LVal:  This parameter MUST be used to encode long values that contain troubleshooting
	// information.
	LValue int32 `idl:"name:LVal" json:"l_value"`
}

func (*ExtendedErrorParam_LValue) is_ExtendedErrorParam_ExtendedErrorParam() {}

func (o *ExtendedErrorParam_LValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.LValue); err != nil {
		return err
	}
	return nil
}
func (o *ExtendedErrorParam_LValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.LValue); err != nil {
		return err
	}
	return nil
}

// ExtendedErrorParam_Value structure represents ExtendedErrorParam_ExtendedErrorParam RPC union arm.
//
// It has following labels: 4
type ExtendedErrorParam_Value struct {
	// IVal:  This parameter MUST be used to encode integer values that contain troubleshooting
	// information.
	Value int16 `idl:"name:IVal" json:"value"`
}

func (*ExtendedErrorParam_Value) is_ExtendedErrorParam_ExtendedErrorParam() {}

func (o *ExtendedErrorParam_Value) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Value); err != nil {
		return err
	}
	return nil
}
func (o *ExtendedErrorParam_Value) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Value); err != nil {
		return err
	}
	return nil
}

// ExtendedErrorParam_PValue structure represents ExtendedErrorParam_ExtendedErrorParam RPC union arm.
//
// It has following labels: 5
type ExtendedErrorParam_PValue struct {
	// PVal:  This parameter MUST be used to encode 64-bit integer values that contain troubleshooting
	// information.
	PValue int64 `idl:"name:PVal" json:"p_value"`
}

func (*ExtendedErrorParam_PValue) is_ExtendedErrorParam_ExtendedErrorParam() {}

func (o *ExtendedErrorParam_PValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.PValue); err != nil {
		return err
	}
	return nil
}
func (o *ExtendedErrorParam_PValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.PValue); err != nil {
		return err
	}
	return nil
}

// ExtendedErrorParam_6 structure represents ExtendedErrorParam_ExtendedErrorParam RPC union arm.
//
// It has following labels: 6
type ExtendedErrorParam_6 struct {
}

func (*ExtendedErrorParam_6) is_ExtendedErrorParam_ExtendedErrorParam() {}

func (o *ExtendedErrorParam_6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *ExtendedErrorParam_6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// ExtendedErrorParam_Blob structure represents ExtendedErrorParam_ExtendedErrorParam RPC union arm.
//
// It has following labels: 7
type ExtendedErrorParam_Blob struct {
	// Blob:  A parameter of type BinaryEEInfo.
	Blob *BinaryInfo `idl:"name:Blob" json:"blob"`
}

func (*ExtendedErrorParam_Blob) is_ExtendedErrorParam_ExtendedErrorParam() {}

func (o *ExtendedErrorParam_Blob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Blob != nil {
		if err := o.Blob.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&BinaryInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtendedErrorParam_Blob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Blob == nil {
		o.Blob = &BinaryInfo{}
	}
	if err := o.Blob.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ComputerNamePresentType type represents EEComputerNamePresentType RPC enumeration.
type ComputerNamePresentType uint16

var (
	ComputerNamePresentTypePresent    ComputerNamePresentType = 1
	ComputerNamePresentTypeNotPresent ComputerNamePresentType = 2
)

func (o ComputerNamePresentType) String() string {
	switch o {
	case ComputerNamePresentTypePresent:
		return "ComputerNamePresentTypePresent"
	case ComputerNamePresentTypeNotPresent:
		return "ComputerNamePresentTypeNotPresent"
	}
	return "Invalid"
}

// ComputerName structure represents EEComputerName RPC structure.
//
// The EEComputerName structure identifies the network node on which the error record
// was generated.
type ComputerName struct {
	// Type:  Indicates the contents of a union.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| eecnpPresent 1   | Network Node Identifier Name member of the union is valid and contains a network |
	//	|                  | node identifier.                                                                 |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| eecnNotPresent 2 | No Network Node Identifier This structure does not contain a network node        |
	//	|                  | identifier.                                                                      |
	//	+------------------+----------------------------------------------------------------------------------+
	Type         ComputerNamePresentType    `idl:"name:Type" json:"type"`
	ComputerName *ComputerName_ComputerName `idl:"name:ComputerName;switch_is:Type" json:"computer_name"`
}

func (o *ComputerName) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ComputerName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	_swComputerName := int16(o.Type)
	if o.ComputerName != nil {
		if err := o.ComputerName.MarshalUnionNDR(ctx, w, _swComputerName); err != nil {
			return err
		}
	} else {
		if err := (&ComputerName_ComputerName{}).MarshalUnionNDR(ctx, w, _swComputerName); err != nil {
			return err
		}
	}
	return nil
}
func (o *ComputerName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.ComputerName == nil {
		o.ComputerName = &ComputerName_ComputerName{}
	}
	_swComputerName := int16(o.Type)
	if err := o.ComputerName.UnmarshalUnionNDR(ctx, w, _swComputerName); err != nil {
		return err
	}
	return nil
}

// ComputerName_ComputerName structure represents EEComputerName union anonymous member.
//
// The EEComputerName structure identifies the network node on which the error record
// was generated.
type ComputerName_ComputerName struct {
	// Types that are assignable to Value
	//
	// *ComputerName_Name
	// *ComputerName_2
	Value is_ComputerName_ComputerName `json:"value"`
}

func (o *ComputerName_ComputerName) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ComputerName_Name:
		if value != nil {
			return value.Name
		}
	}
	return nil
}

type is_ComputerName_ComputerName interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ComputerName_ComputerName()
}

func (o *ComputerName_ComputerName) NDRSwitchValue(sw int16) int16 {
	if o == nil {
		return int16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ComputerName_Name:
		return int16(1)
	case *ComputerName_2:
		return int16(2)
	}
	return int16(0)
}

func (o *ComputerName_ComputerName) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw int16) error {
	if err := w.WriteSwitch(int16(sw)); err != nil {
		return err
	}
	switch sw {
	case int16(1):
		_o, _ := o.Value.(*ComputerName_Name)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ComputerName_Name{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int16(2):
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ComputerName_ComputerName) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw int16) error {
	if err := w.ReadSwitch((*int16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case int16(1):
		o.Value = &ComputerName_Name{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int16(2):
		o.Value = &ComputerName_2{}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ComputerName_Name structure represents ComputerName_ComputerName RPC union arm.
//
// It has following labels: 1
type ComputerName_Name struct {
	// Name:  Unicode string that identifies the network node on which the error record
	// was generated. The format in which the network node is identified is implementation-specific,
	// and this information MUST be used for display purposes only. This specification does
	// not define what the format is. Software agents who use this structure SHOULD use
	// a network node identifier that is unique within a specific topology and is descriptive
	// to a human reader. If Type is equal to eecnpNotPresent, the error record MUST be
	// interpreted as generated on the local network node.
	Name *UnicodeString `idl:"name:Name" json:"name"`
}

func (*ComputerName_Name) is_ComputerName_ComputerName() {}

func (o *ComputerName_Name) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ComputerName_Name) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Name == nil {
		o.Name = &UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ComputerName_2 structure represents ComputerName_ComputerName RPC union arm.
//
// It has following labels: 2
type ComputerName_2 struct {
}

func (*ComputerName_2) is_ComputerName_ComputerName() {}

func (o *ComputerName_2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *ComputerName_2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// ExtendedErrorInfo structure represents ExtendedErrorInfo RPC structure.
//
// The ExtendedErrorInfo structure represents an error record.
type ExtendedErrorInfo struct {
	// Next:  An error record for the immediate error cause for this error record. For the
	// root error, it MUST be set to NULL.
	Next *ExtendedErrorInfo `idl:"name:Next" json:"next"`
	// ComputerName:  Network node identifier as specified in section 2.2.1.7.
	ComputerName *ComputerName `idl:"name:ComputerName" json:"computer_name"`
	// ProcessID:  The ID of the process in which the error occurred.
	ProcessID uint32 `idl:"name:ProcessID" json:"process_id"`
	// TimeStamp:  Time at which the error record was generated, which is expressed as the
	// number of 100-nanosecond intervals since January 1, 1601. It MUST be interpreted
	// as Coordinated Universal Time (UTC).
	Timestamp int64 `idl:"name:TimeStamp" json:"timestamp"`
	// GeneratingComponent:  Component or protocol layer identifier where the error occurred
	// as described in section 1.3.1.
	GeneratingComponent uint32 `idl:"name:GeneratingComponent" json:"generating_component"`
	// Status:  Error code as described in section 1.3.1.
	Status uint32 `idl:"name:Status" json:"status"`
	// DetectionLocation:  Location where the error occurred as described in section 1.3.1.
	DetectionLocation uint16 `idl:"name:DetectionLocation" json:"detection_location"`
	// Flags:  One or more flags that specify the presence or absence of other error records
	// in the error sequence.
	//
	//	+--------+----------------------------------------------------------------------------------+
	//	|        |                                                                                  |
	//	| VALUE  |                                     MEANING                                      |
	//	|        |                                                                                  |
	//	+--------+----------------------------------------------------------------------------------+
	//	+--------+----------------------------------------------------------------------------------+
	//	| 0x0000 | All of the error records from the error sequence are present in the encoding.    |
	//	+--------+----------------------------------------------------------------------------------+
	//	| 0x0001 | One or more error records from the error sequence before the current record are  |
	//	|        | not present in the encoding.                                                     |
	//	+--------+----------------------------------------------------------------------------------+
	//	| 0x0002 | One or more error records from the error sequence after the current record are   |
	//	|        | not present in the encoding.                                                     |
	//	+--------+----------------------------------------------------------------------------------+
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// nLen:  Number of elements in the Params array. MUST be less than or equal to 4.
	Length int16 `idl:"name:nLen" json:"length"`
	// Params:  Array of error parameters as described in the data model in section 1.3.1.
	Params []*ExtendedErrorParam `idl:"name:Params;size_is:(nLen)" json:"params"`
}

func (o *ExtendedErrorInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.Params != nil && o.Length == 0 {
		o.Length = int16(len(o.Params))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtendedErrorInfo) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Length)
	return []uint64{
		dimSize1,
	}
}
func (o *ExtendedErrorInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_Next := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ExtendedErrorInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_Next); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ComputerName != nil {
		if err := o.ComputerName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ComputerName{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProcessID); err != nil {
		return err
	}
	if err := w.WriteData(o.Timestamp); err != nil {
		return err
	}
	if err := w.WriteData(o.GeneratingComponent); err != nil {
		return err
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	if err := w.WriteData(o.DetectionLocation); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	for i1 := range o.Params {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Params[i1] != nil {
			if err := o.Params[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ExtendedErrorParam{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Params); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&ExtendedErrorParam{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtendedErrorInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	_ptr_Next := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &ExtendedErrorInfo{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Next := func(ptr interface{}) { o.Next = *ptr.(**ExtendedErrorInfo) }
	if err := w.ReadPointer(&o.Next, _s_Next, _ptr_Next); err != nil {
		return err
	}
	if o.ComputerName == nil {
		o.ComputerName = &ComputerName{}
	}
	if err := o.ComputerName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProcessID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timestamp); err != nil {
		return err
	}
	if err := w.ReadData(&o.GeneratingComponent); err != nil {
		return err
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	if err := w.ReadData(&o.DetectionLocation); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Length > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Length)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Params", sizeInfo[0])
	}
	o.Params = make([]*ExtendedErrorParam, sizeInfo[0])
	for i1 := range o.Params {
		i1 := i1
		if o.Params[i1] == nil {
			o.Params[i1] = &ExtendedErrorParam{}
		}
		if err := o.Params[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

type xxx_DefaultExtendedErrorClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultExtendedErrorClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewExtendedErrorClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ExtendedErrorClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ExtendedErrorSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultExtendedErrorClient{cc: cc}, nil
}

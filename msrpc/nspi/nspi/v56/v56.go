package nspi

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
	_ = dcetypes.GoPackage
)

var (
	// import guard
	GoPackage = "nspi"
)

var (
	// Syntax UUID
	NspiSyntaxUUID = &uuid.UUID{TimeLow: 0xf5cc5a18, TimeMid: 0x4264, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x59, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2f, 0x84, 0x26}}
	// Syntax ID
	NspiSyntaxV56_0 = &dcerpc.SyntaxID{IfUUID: NspiSyntaxUUID, IfVersionMajor: 56, IfVersionMinor: 0}
)

// nspi interface.
type NspiClient interface {

	// The NspiBind method initiates a session between a client and the NSPI server.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	Bind(context.Context, *BindRequest, ...dcerpc.CallOption) (*BindResponse, error)

	// The NspiUnbind method destroys the context handle. No other action is taken.
	//
	// Return Values: The server returns a DWORD value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	Unbind(context.Context, *UnbindRequest, ...dcerpc.CallOption) (*UnbindResponse, error)

	// The NspiUpdateStat method updates the STAT block representing position in a table
	// to reflect positioning changes requested by the client.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	UpdateStat(context.Context, *UpdateStatRequest, ...dcerpc.CallOption) (*UpdateStatResponse, error)

	// The NspiQueryRows method returns to the client a number of rows from a specified
	// table. The server MUST return no more rows than the number specified in the input
	// parameter Count. Although the protocol places no further boundary or requirements
	// on the minimum number of rows the server returns, implementations SHOULD return as
	// many rows as possible subject to this maximum limit to improve usability of the NSPI
	// server for clients.
	//
	// Return Values:  The server returns a long value specifying the return status of
	// the method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	QueryRows(context.Context, *QueryRowsRequest, ...dcerpc.CallOption) (*QueryRowsResponse, error)

	// The NspiSeekEntries method searches for and sets the logical position in a specific
	// table to the first entry greater than or equal to a specified value. Optionally,
	// it might also return information about rows in the table.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	SeekEntries(context.Context, *SeekEntriesRequest, ...dcerpc.CallOption) (*SeekEntriesResponse, error)

	// The NspiGetMatches method returns an Explicit Table. The rows in the table are chosen
	// based on a two possible criteria: a restriction applied to an address book container
	// or the values of a property on a single object that hold references to other objects.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetMatches(context.Context, *GetMatchesRequest, ...dcerpc.CallOption) (*GetMatchesResponse, error)

	// The NspiResortRestriction method applies a sort order to the objects in a restricted
	// address book container.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResortRestriction(context.Context, *ResortRestrictionRequest, ...dcerpc.CallOption) (*ResortRestrictionResponse, error)

	// The NspiDNToMId method maps a set of DN to a set of MId.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	DNToMID(context.Context, *DNToMIDRequest, ...dcerpc.CallOption) (*DNToMIDResponse, error)

	// The NspiGetPropList method returns a list of all the properties that have values
	// on a specified object.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetPropertyList(context.Context, *GetPropertyListRequest, ...dcerpc.CallOption) (*GetPropertyListResponse, error)

	// The NspiGetProps method returns an address book row containing a set of the properties
	// and values that exist on an object.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The NspiCompareMIds method compares the position in an address book container of
	// two objects identified by MId and returns the value of the comparison.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	CompareMIDs(context.Context, *CompareMIDsRequest, ...dcerpc.CallOption) (*CompareMIDsResponse, error)

	// The NspiModProps method is used to modify the properties of an object in the address
	// book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ModifyProperties(context.Context, *ModifyPropertiesRequest, ...dcerpc.CallOption) (*ModifyPropertiesResponse, error)

	// The NspiGetSpecialTable method returns the rows of a special table to the client.
	// The special table can be an Address Book Hierarchy Table or an Address Creation Table.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetSpecialTable(context.Context, *GetSpecialTableRequest, ...dcerpc.CallOption) (*GetSpecialTableResponse, error)

	// The NspiGetTemplateInfo method returns information about template objects in the
	// address book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetTemplateInfo(context.Context, *GetTemplateInfoRequest, ...dcerpc.CallOption) (*GetTemplateInfoResponse, error)

	// The NspiModLinkAtt method modifies the values of a specific property of a specific
	// row in the address book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ModifyLinkAttribute(context.Context, *ModifyLinkAttributeRequest, ...dcerpc.CallOption) (*ModifyLinkAttributeResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// The NspiQueryColumns method returns a list of all the properties the NSPI server
	// is aware of. It returns this list as an array of proptags.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	QueryColumns(context.Context, *QueryColumnsRequest, ...dcerpc.CallOption) (*QueryColumnsResponse, error)

	// The NspiGetNamesFromIDs method returns a list of property names for a set of proptags.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetNamesFromIDs(context.Context, *GetNamesFromIDsRequest, ...dcerpc.CallOption) (*GetNamesFromIDsResponse, error)

	// The NspiGetIDsFromNames method returns a list of proptags for a set of property names.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetIDsFromNames(context.Context, *GetIDsFromNamesRequest, ...dcerpc.CallOption) (*GetIDsFromNamesResponse, error)

	// The NspiResolveNames method takes a set of string values in an 8-bit character set
	// and performs ANR (as specified in 3.1.1.6) on those strings. The server reports the
	// MId that are the result of the ANR process. Certain property values are returned
	// for any valid MIds identified by the ANR process.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResolveNames(context.Context, *ResolveNamesRequest, ...dcerpc.CallOption) (*ResolveNamesResponse, error)

	// The NspiResolveNamesW method takes a set of string values in the Unicode character
	// set and performs ANR (as specified in 3.1.1.6) on those strings. The server reports
	// the MId that are the result of the ANR process. Certain property values are returned
	// for any valid MIds identified by the ANR process.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResolveNamesW(context.Context, *ResolveNamesWRequest, ...dcerpc.CallOption) (*ResolveNamesWResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// FlatUID structure represents FlatUID_r RPC structure.
//
// The FlatUID_r structure is an encoding of the FlatUID data structure defined in [MS-OXCDATA].
// The semantic meaning is unchanged from the FlatUID data structure.
type FlatUID struct {
	// ab:  Encodes the ordered bytes of the FlatUID data structure.
	Value []byte `idl:"name:ab" json:"value"`
}

func (o *FlatUID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FlatUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Value {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Value[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Value); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *FlatUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Value = make([]byte, 16)
	for i1 := range o.Value {
		i1 := i1
		if err := w.ReadData(&o.Value[i1]); err != nil {
			return err
		}
	}
	return nil
}

// PropertyTagArray structure represents PropertyTagArray_r RPC structure.
//
// The PropertyTagArray_r structure is an encoding of the PropTagArray data structure
// defined in [MS-OXCDATA]. The permissible number of proptag values in the PropertyTagArray_r
// structure exceeds that of the PropTagArray data structure. The semantic meaning is
// otherwise unchanged from the PropTagArray data structure.
type PropertyTagArray struct {
	// cValues:  Encodes the Count field of PropTagArray. This field MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// aulPropTag:  Encodes the PropertyTags field of PropTagArray.
	PropertyTag []uint32 `idl:"name:aulPropTag;size_is:((cValues+1));length_is:(cValues)" json:"property_tag"`
}

func (o *PropertyTagArray) xxx_PreparePayload(ctx context.Context) error {
	if o.PropertyTag != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32((len(o.PropertyTag) - 1))
		if len(o.PropertyTag) < 1 {
			o.ValuesCount = uint32(0)
		}
	}
	if o.PropertyTag != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.PropertyTag))
	}
	if len(o.PropertyTag) > int(100001) {
		return fmt.Errorf("PropertyTag is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *PropertyTagArray) NDRSizeInfo() []uint64 {
	dimSize1 := uint64((o.ValuesCount + 1))
	return []uint64{
		dimSize1,
	}
}
func (o *PropertyTagArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	dimLength1 := uint64(o.ValuesCount)
	if dimLength1 > sizeInfo[0] {
		dimLength1 = sizeInfo[0]
	} else {
		sizeInfo[0] = dimLength1
	}
	if err := w.WriteSize(0); err != nil {
		return err
	}
	if err := w.WriteSize(dimLength1); err != nil {
		return err
	}
	for i1 := range o.PropertyTag {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.PropertyTag[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PropertyTag); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyTagArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	for sz1 := range sizeInfo {
		if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
			return err
		}
		if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
			return err
		}
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.PropertyTag", sizeInfo[0])
	}
	o.PropertyTag = make([]uint32, sizeInfo[0])
	for i1 := range o.PropertyTag {
		i1 := i1
		if err := w.ReadData(&o.PropertyTag[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Binary structure represents Binary_r RPC structure.
//
// The Binary_r structure encodes an array of uninterpreted bytes.
type Binary struct {
	// cb:  The number of uninterpreted bytes represented in this structure. This value
	// MUST NOT exceed 2,097,152.
	Length uint32 `idl:"name:cb" json:"length"`
	// lpb:  The uninterpreted bytes.
	Value []byte `idl:"name:lpb;size_is:(cb)" json:"value"`
}

func (o *Binary) xxx_PreparePayload(ctx context.Context) error {
	if o.Value != nil && o.Length == 0 {
		o.Length = uint32(len(o.Value))
	}
	if o.Length > uint32(2097152) {
		return fmt.Errorf("Length is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Binary) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.Value != nil || o.Length > 0 {
		_ptr_lpb := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Value {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Value[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Value); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_lpb); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Binary) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_lpb := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return fmt.Errorf("buffer overflow for size %d of array o.Value", sizeInfo[0])
		}
		o.Value = make([]byte, sizeInfo[0])
		for i1 := range o.Value {
			i1 := i1
			if err := w.ReadData(&o.Value[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpb := func(ptr interface{}) { o.Value = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Value, _s_lpb, _ptr_lpb); err != nil {
		return err
	}
	return nil
}

// ShortArray structure represents ShortArray_r RPC structure.
//
// The ShortArray_r structure encodes an array of 16-bit integers.
type ShortArray struct {
	// cValues:  The number of 16-bit integer values represented in the ShortArray_r structure.
	// This value MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lpi:  The 16-bit integer values.
	Values []int16 `idl:"name:lpi;size_is:(cValues)" json:"values"`
}

func (o *ShortArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Values != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.Values))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ShortArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.Values != nil || o.ValuesCount > 0 {
		_ptr_lpi := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Values[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(int16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_lpi); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ShortArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lpi := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]int16, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			if err := w.ReadData(&o.Values[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpi := func(ptr interface{}) { o.Values = *ptr.(*[]int16) }
	if err := w.ReadPointer(&o.Values, _s_lpi, _ptr_lpi); err != nil {
		return err
	}
	return nil
}

// LongArray structure represents LongArray_r RPC structure.
//
// The LongArray_r structure encodes an array of 32-bit integers.
type LongArray struct {
	// cValues:  The number of 32-bit integers represented in this structure. This value
	// MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lpl:  The 32-bit integer values.
	Values []int32 `idl:"name:lpl;size_is:(cValues)" json:"values"`
}

func (o *LongArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Values != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.Values))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LongArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.Values != nil || o.ValuesCount > 0 {
		_ptr_lpl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Values[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(int32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_lpl); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LongArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lpl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]int32, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			if err := w.ReadData(&o.Values[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpl := func(ptr interface{}) { o.Values = *ptr.(*[]int32) }
	if err := w.ReadPointer(&o.Values, _s_lpl, _ptr_lpl); err != nil {
		return err
	}
	return nil
}

// CharStringArray structure represents StringArray_r RPC structure.
//
// The StringArray_r structure encodes an array of references to 8-bit character strings.
type CharStringArray struct {
	// cValues:  The number of 8-bit character strings references represented in the StringArray_r
	// structure. This value MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lppszA:  The 8-bit character string references. The strings referred to are NULL-terminated.
	Values []string `idl:"name:lppszA;size_is:(cValues);string" json:"values"`
}

func (o *CharStringArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Values != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.Values))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CharStringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.Values != nil || o.ValuesCount > 0 {
		_ptr_lppszA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Values[i1] != "" {
					_ptr_lppszA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteCharNString(ctx, w, o.Values[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.Values[i1], _ptr_lppszA); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_lppszA); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CharStringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lppszA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]string, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			_ptr_lppszA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadCharNString(ctx, w, &o.Values[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_lppszA := func(ptr interface{}) { o.Values[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Values[i1], _s_lppszA, _ptr_lppszA); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lppszA := func(ptr interface{}) { o.Values = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.Values, _s_lppszA, _ptr_lppszA); err != nil {
		return err
	}
	return nil
}

// BinaryArray structure represents BinaryArray_r RPC structure.
//
// The BinaryArray_r structure is an array of Binary_r data structures.
type BinaryArray struct {
	// cValues:  The number of Binary_r data structures represented in the BinaryArray_r
	// structure. This value MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lpbin:  The Binary_r data structures.
	Values []*Binary `idl:"name:lpbin;size_is:(cValues)" json:"values"`
}

func (o *BinaryArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Values != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.Values))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BinaryArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.Values != nil || o.ValuesCount > 0 {
		_ptr_lpbin := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Values[i1] != nil {
					if err := o.Values[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Binary{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Binary{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_lpbin); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BinaryArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lpbin := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]*Binary, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			if o.Values[i1] == nil {
				o.Values[i1] = &Binary{}
			}
			if err := o.Values[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpbin := func(ptr interface{}) { o.Values = *ptr.(*[]*Binary) }
	if err := w.ReadPointer(&o.Values, _s_lpbin, _ptr_lpbin); err != nil {
		return err
	}
	return nil
}

// FlatUIDArray structure represents FlatUIDArray_r RPC structure.
//
// The FlatUIDArray_r structure encodes an array of FlatUID_r data structures.
type FlatUIDArray struct {
	// cValues:  The number of FlatUID_r structures represented in the FlatUIDArray_r structure.
	// This value MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lpguid:   The FlatUID_r data structures.
	GUID []*FlatUID `idl:"name:lpguid;size_is:(cValues)" json:"guid"`
}

func (o *FlatUIDArray) xxx_PreparePayload(ctx context.Context) error {
	if o.GUID != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.GUID))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FlatUIDArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.GUID != nil || o.ValuesCount > 0 {
		_ptr_lpguid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.GUID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.GUID[i1] != nil {
					_ptr_lpguid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.GUID[i1] != nil {
							if err := o.GUID[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&FlatUID{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.GUID[i1], _ptr_lpguid); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.GUID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GUID, _ptr_lpguid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FlatUIDArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lpguid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.GUID", sizeInfo[0])
		}
		o.GUID = make([]*FlatUID, sizeInfo[0])
		for i1 := range o.GUID {
			i1 := i1
			_ptr_lpguid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.GUID[i1] == nil {
					o.GUID[i1] = &FlatUID{}
				}
				if err := o.GUID[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_lpguid := func(ptr interface{}) { o.GUID[i1] = *ptr.(**FlatUID) }
			if err := w.ReadPointer(&o.GUID[i1], _s_lpguid, _ptr_lpguid); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpguid := func(ptr interface{}) { o.GUID = *ptr.(*[]*FlatUID) }
	if err := w.ReadPointer(&o.GUID, _s_lpguid, _ptr_lpguid); err != nil {
		return err
	}
	return nil
}

// StringArray structure represents WStringArray_r RPC structure.
//
// The WStringArray_r structure encodes an array of references to Unicode strings.
type StringArray struct {
	// cValues:  The number of Unicode character strings references represented in the WStringArray_r
	// structure. This value MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lppszW:  The Unicode character string references. The strings referred to are NULL-terminated.
	Values []string `idl:"name:lppszW;size_is:(cValues);string" json:"values"`
}

func (o *StringArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Values != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.Values))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.Values != nil || o.ValuesCount > 0 {
		_ptr_lppszW := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Values[i1] != "" {
					_ptr_lppszW := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.Values[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.Values[i1], _ptr_lppszW); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_lppszW); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lppszW := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]string, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			_ptr_lppszW := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.Values[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_lppszW := func(ptr interface{}) { o.Values[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Values[i1], _s_lppszW, _ptr_lppszW); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lppszW := func(ptr interface{}) { o.Values = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.Values, _s_lppszW, _ptr_lppszW); err != nil {
		return err
	}
	return nil
}

// DateTimeArray structure represents DateTimeArray_r RPC structure.
//
// The DateTimeArray_r structure encodes an array of FILETIME structures.
type DateTimeArray struct {
	// cValues:  The number of FILETIME data structures represented in the DateTimeArray_r
	// structure. This value MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lpft:  The FILETIME data structures.
	Values []*dtyp.Filetime `idl:"name:lpft;size_is:(cValues)" json:"values"`
}

func (o *DateTimeArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Values != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.Values))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DateTimeArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.Values != nil || o.ValuesCount > 0 {
		_ptr_lpft := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Values[i1] != nil {
					if err := o.Values[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_lpft); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DateTimeArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lpft := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]*dtyp.Filetime, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			if o.Values[i1] == nil {
				o.Values[i1] = &dtyp.Filetime{}
			}
			if err := o.Values[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpft := func(ptr interface{}) { o.Values = *ptr.(*[]*dtyp.Filetime) }
	if err := w.ReadPointer(&o.Values, _s_lpft, _ptr_lpft); err != nil {
		return err
	}
	return nil
}

// PropertyValue structure represents PropertyValue_r RPC structure.
//
// The PropertyValue_r structure is an encoding of the PropertyValue data structure
// defined in [MS-OXCDATA].
//
// For property values with uninterpreted byte values, the permissible number of bytes
// in the PropertyValue_r structure exceeds that of the PropertyValue data structure.
// For property values with multiple values, the permissible number of values in the
// PropertyValue_r structure exceeds that of the PropertyValue data structure. The semantic
// meaning is otherwise unchanged from the PropertyValue data structure.
type PropertyValue struct {
	// ulPropTag:  Encodes the proptag of the property whose value is represented by the
	// PropertyValue_r data structure.
	PropertyTag uint32 `idl:"name:ulPropTag" json:"property_tag"`
	// ulReserved:  Reserved. All clients and servers MUST set this value to the constant
	// 0x00000000.
	_ uint32 `idl:"name:ulReserved"`
	// Value:  Encodes the actual value of the property represented by the PropertyValue_r
	// data structure. The type value held is specified by the Property Type of the proptag
	// in the field ulPropTag.
	Value *PropertyValueUnion `idl:"name:Value;switch_is:(ulPropTag 65535 &)" json:"value"`
}

func (o *PropertyValue) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag); err != nil {
		return err
	}
	// reserved ulReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	_swValue := int32((o.PropertyTag & 65535))
	if o.Value != nil {
		if err := o.Value.MarshalUnionNDR(ctx, w, _swValue); err != nil {
			return err
		}
	} else {
		if err := (&PropertyValueUnion{}).MarshalUnionNDR(ctx, w, _swValue); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag); err != nil {
		return err
	}
	// reserved ulReserved
	var _ulReserved uint32
	if err := w.ReadData(&_ulReserved); err != nil {
		return err
	}
	if o.Value == nil {
		o.Value = &PropertyValueUnion{}
	}
	_swValue := int32((o.PropertyTag & 65535))
	if err := o.Value.UnmarshalUnionNDR(ctx, w, _swValue); err != nil {
		return err
	}
	return nil
}

// PropertyRow structure represents PropertyRow_r RPC structure.
//
// The PropertyRow_r structure is an encoding of the StandardPropertyRow data structure
// defined in [MS-OXCDATA]. The semantic meaning is unchanged from the StandardPropertyRow
// data structure.
type PropertyRow struct {
	// Reserved:  Reserved. All clients and servers MUST set this value to the constant
	// 0x00000000.
	_ uint32 `idl:"name:Reserved"`
	// cValues:  The number of PropertyValue_r structures represented in the PropertyRow_r
	// structure. This value MUST NOT exceed 100,000.
	ValuesCount uint32 `idl:"name:cValues" json:"values_count"`
	// lpProps:  Encodes the ValueArray field of the StandardPropertyRow data structure.
	Properties []*PropertyValue `idl:"name:lpProps;size_is:(cValues)" json:"properties"`
}

func (o *PropertyRow) xxx_PreparePayload(ctx context.Context) error {
	if o.Properties != nil && o.ValuesCount == 0 {
		o.ValuesCount = uint32(len(o.Properties))
	}
	if o.ValuesCount > uint32(100000) {
		return fmt.Errorf("ValuesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.ValuesCount); err != nil {
		return err
	}
	if o.Properties != nil || o.ValuesCount > 0 {
		_ptr_lpProps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValuesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Properties {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Properties[i1] != nil {
					if err := o.Properties[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Properties); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PropertyValue{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Properties, _ptr_lpProps); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValuesCount); err != nil {
		return err
	}
	_ptr_lpProps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValuesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValuesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Properties", sizeInfo[0])
		}
		o.Properties = make([]*PropertyValue, sizeInfo[0])
		for i1 := range o.Properties {
			i1 := i1
			if o.Properties[i1] == nil {
				o.Properties[i1] = &PropertyValue{}
			}
			if err := o.Properties[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpProps := func(ptr interface{}) { o.Properties = *ptr.(*[]*PropertyValue) }
	if err := w.ReadPointer(&o.Properties, _s_lpProps, _ptr_lpProps); err != nil {
		return err
	}
	return nil
}

// PropertyRowSet structure represents PropertyRowSet_r RPC structure.
//
// The PropertyRowSet_r structure is an encoding of the PropertyRowSet data structure
// defined in [MS-OXCDATA] section 2.19.2, PropertyRowSet.
//
// The permissible number of PropertyRows in the PropertyRowSet_r data structure exceeds
// that of the PropertyRowSet data structure. The semantic meaning is otherwise unchanged
// from the PropertyRowSet data structure.
type PropertyRowSet struct {
	// cRows:   Encodes the RowCount field of the PropertyRowSet data structures. This value
	// MUST NOT exceed 100,000.
	RowsCount uint32 `idl:"name:cRows" json:"rows_count"`
	// aRow:  Encodes the Rows field of the PropertyRowSet data structure.
	Row []*PropertyRow `idl:"name:aRow;size_is:(cRows)" json:"row"`
}

func (o *PropertyRowSet) xxx_PreparePayload(ctx context.Context) error {
	if o.Row != nil && o.RowsCount == 0 {
		o.RowsCount = uint32(len(o.Row))
	}
	if o.RowsCount > uint32(100000) {
		return fmt.Errorf("RowsCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *PropertyRowSet) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.RowsCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PropertyRowSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RowsCount); err != nil {
		return err
	}
	for i1 := range o.Row {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Row[i1] != nil {
			if err := o.Row[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Row); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&PropertyRow{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyRowSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RowsCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.RowsCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.RowsCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Row", sizeInfo[0])
	}
	o.Row = make([]*PropertyRow, sizeInfo[0])
	for i1 := range o.Row {
		i1 := i1
		if o.Row[i1] == nil {
			o.Row[i1] = &PropertyRow{}
		}
		if err := o.Row[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// Restriction structure represents Restriction_r RPC structure.
//
// The Restriction_r structure is an encoding of the Restriction filters defined in
// [MS-OXCDATA].
//
// The permissible number of Restriction structures encoded in AndRestriction_r and
// OrRestriction_r data structures recursively included in the Restriction_r data type
// exceeds that of the AndRestriction_r and OrRestriction_r data structures recursively
// included in the Restriction filters. The semantic meaning is otherwise unchanged
// from the Restriction filters.
type Restriction struct {
	// rt:  Encodes the RestrictType field common to all restriction structures.
	RestrictionType uint32 `idl:"name:rt" json:"restriction_type"`
	// res:  Encodes the actual restriction specified by the type in the rt field.
	Restriction *RestrictionUnion `idl:"name:res;switch_is:rt" json:"restriction"`
}

func (o *Restriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Restriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RestrictionType); err != nil {
		return err
	}
	_swRestriction := int32(o.RestrictionType)
	if o.Restriction != nil {
		if err := o.Restriction.MarshalUnionNDR(ctx, w, _swRestriction); err != nil {
			return err
		}
	} else {
		if err := (&RestrictionUnion{}).MarshalUnionNDR(ctx, w, _swRestriction); err != nil {
			return err
		}
	}
	return nil
}
func (o *Restriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RestrictionType); err != nil {
		return err
	}
	if o.Restriction == nil {
		o.Restriction = &RestrictionUnion{}
	}
	_swRestriction := int32(o.RestrictionType)
	if err := o.Restriction.UnmarshalUnionNDR(ctx, w, _swRestriction); err != nil {
		return err
	}
	return nil
}

// AndRestriction structure represents AndRestriction_r RPC structure.
//
// The AndRestriction_r, OrRestriction_r restriction types share a single RPC encoding.
// The AndOrRestriction_r structure is an encoding of the both the AndRestriction data
// structure and the OrRestriction data structure defined in [MS-OXCDATA]. These two
// data structures share the same data layout, so a single encoding is included in the
// NSPI Protocol. The sense of the data structure's use is derived from the context
// of its inclusion in the RestrictionUnion_r data structure defined in this specification.
//
// The permissible number of Restriction structures in the AndRestriction_r and OrRestriction_r
// data structures exceeds that of the AndRestriction and OrRestriction structures.
// The semantic meaning is otherwise unchanged from the AndRestriction and OrRestriction
// data structures, as context dictates.
type AndRestriction struct {
	// cRes:  Encodes the RestrictCount field of the AndRestriction and OrRestriction data
	// structures. This value MUST NOT exceed 100,000.
	RestrictionCount uint32 `idl:"name:cRes" json:"restriction_count"`
	// lpRes:  Encodes the Restricts field of the AndRestriction and OrRestriction data
	// structures.
	Restriction []*Restriction `idl:"name:lpRes;size_is:(cRes)" json:"restriction"`
}

func (o *AndRestriction) xxx_PreparePayload(ctx context.Context) error {
	if o.Restriction != nil && o.RestrictionCount == 0 {
		o.RestrictionCount = uint32(len(o.Restriction))
	}
	if o.RestrictionCount > uint32(100000) {
		return fmt.Errorf("RestrictionCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AndRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RestrictionCount); err != nil {
		return err
	}
	if o.Restriction != nil || o.RestrictionCount > 0 {
		_ptr_lpRes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RestrictionCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Restriction {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Restriction[i1] != nil {
					if err := o.Restriction[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Restriction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Restriction); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Restriction{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Restriction, _ptr_lpRes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AndRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RestrictionCount); err != nil {
		return err
	}
	_ptr_lpRes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RestrictionCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RestrictionCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Restriction", sizeInfo[0])
		}
		o.Restriction = make([]*Restriction, sizeInfo[0])
		for i1 := range o.Restriction {
			i1 := i1
			if o.Restriction[i1] == nil {
				o.Restriction[i1] = &Restriction{}
			}
			if err := o.Restriction[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpRes := func(ptr interface{}) { o.Restriction = *ptr.(*[]*Restriction) }
	if err := w.ReadPointer(&o.Restriction, _s_lpRes, _ptr_lpRes); err != nil {
		return err
	}
	return nil
}

// OrRestriction structure represents OrRestriction_r RPC structure.
type OrRestriction struct {
	RestrictionCount uint32         `idl:"name:cRes" json:"restriction_count"`
	Restriction      []*Restriction `idl:"name:lpRes;size_is:(cRes)" json:"restriction"`
}

func (o *OrRestriction) xxx_PreparePayload(ctx context.Context) error {
	if o.Restriction != nil && o.RestrictionCount == 0 {
		o.RestrictionCount = uint32(len(o.Restriction))
	}
	if o.RestrictionCount > uint32(100000) {
		return fmt.Errorf("RestrictionCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OrRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RestrictionCount); err != nil {
		return err
	}
	if o.Restriction != nil || o.RestrictionCount > 0 {
		_ptr_lpRes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RestrictionCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Restriction {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Restriction[i1] != nil {
					if err := o.Restriction[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Restriction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Restriction); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Restriction{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Restriction, _ptr_lpRes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OrRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RestrictionCount); err != nil {
		return err
	}
	_ptr_lpRes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RestrictionCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RestrictionCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Restriction", sizeInfo[0])
		}
		o.Restriction = make([]*Restriction, sizeInfo[0])
		for i1 := range o.Restriction {
			i1 := i1
			if o.Restriction[i1] == nil {
				o.Restriction[i1] = &Restriction{}
			}
			if err := o.Restriction[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpRes := func(ptr interface{}) { o.Restriction = *ptr.(*[]*Restriction) }
	if err := w.ReadPointer(&o.Restriction, _s_lpRes, _ptr_lpRes); err != nil {
		return err
	}
	return nil
}

// NotRestriction structure represents NotRestriction_r RPC structure.
//
// The NotRestriction_r structure is an encoding of the NotRestriction data structure
// defined in [MS-OXCDATA]. The semantic meaning is unchanged from the NotRestriction
// data structure.
type NotRestriction struct {
	// lpRes:  Encodes the Restriction field of the NotRestriction data structure.
	Restriction *Restriction `idl:"name:lpRes" json:"restriction"`
}

func (o *NotRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NotRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Restriction != nil {
		_ptr_lpRes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Restriction != nil {
				if err := o.Restriction.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Restriction{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Restriction, _ptr_lpRes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NotRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_lpRes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Restriction == nil {
			o.Restriction = &Restriction{}
		}
		if err := o.Restriction.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpRes := func(ptr interface{}) { o.Restriction = *ptr.(**Restriction) }
	if err := w.ReadPointer(&o.Restriction, _s_lpRes, _ptr_lpRes); err != nil {
		return err
	}
	return nil
}

// ContentRestriction structure represents ContentRestriction_r RPC structure.
type ContentRestriction struct {
	FuzzyLevel  uint32         `idl:"name:ulFuzzyLevel" json:"fuzzy_level"`
	PropertyTag uint32         `idl:"name:ulPropTag" json:"property_tag"`
	Property    *PropertyValue `idl:"name:lpProp" json:"property"`
}

func (o *ContentRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContentRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.FuzzyLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag); err != nil {
		return err
	}
	if o.Property != nil {
		_ptr_lpProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Property != nil {
				if err := o.Property.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PropertyValue{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Property, _ptr_lpProp); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContentRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.FuzzyLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag); err != nil {
		return err
	}
	_ptr_lpProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Property == nil {
			o.Property = &PropertyValue{}
		}
		if err := o.Property.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpProp := func(ptr interface{}) { o.Property = *ptr.(**PropertyValue) }
	if err := w.ReadPointer(&o.Property, _s_lpProp, _ptr_lpProp); err != nil {
		return err
	}
	return nil
}

// BitMaskRestriction structure represents BitMaskRestriction_r RPC structure.
//
// The BitMaskRestriction_r structure is an encoding of the BitMaskRestriction data
// structure defined in [MS-OXCDATA]. The semantic meaning is unchanged from the BitMaskRestriction
// data structure.
type BitMaskRestriction struct {
	// relBMR:  Encodes the BitmapRelOp field of the BitMaskRestriction data structure.
	BitMaskRestriction uint32 `idl:"name:relBMR" json:"bit_mask_restriction"`
	// ulPropTag:  Encodes the PropTag field of the BitMaskRestriction data structure.
	PropertyTag uint32 `idl:"name:ulPropTag" json:"property_tag"`
	// ulMask:  Encodes the Mask field of the BitMaskRestriction data structure.
	Mask uint32 `idl:"name:ulMask" json:"mask"`
}

func (o *BitMaskRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BitMaskRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.BitMaskRestriction); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	return nil
}
func (o *BitMaskRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.BitMaskRestriction); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	return nil
}

// PropertyRestriction structure represents PropertyRestriction_r RPC structure.
//
// The PropertyRestriction_r structure is an encoding of the PropertyRestriction data
// structure defined in [MS-OXCDATA]. The semantic meaning is unchanged from the PropertyRestriction
// data structure.
type PropertyRestriction struct {
	// relop:   Encodes the RelOp field of the PropertyRestriction data structure.
	Op uint32 `idl:"name:relop" json:"op"`
	// ulPropTag:  Encodes the PropTag field of the PropertyRestriction data structure.
	PropertyTag uint32 `idl:"name:ulPropTag" json:"property_tag"`
	// lpProp:  Encodes the TaggedValue field of the PropertyRestriction data structure.
	Property *PropertyValue `idl:"name:lpProp" json:"property"`
}

func (o *PropertyRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Op); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag); err != nil {
		return err
	}
	if o.Property != nil {
		_ptr_lpProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Property != nil {
				if err := o.Property.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PropertyValue{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Property, _ptr_lpProp); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Op); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag); err != nil {
		return err
	}
	_ptr_lpProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Property == nil {
			o.Property = &PropertyValue{}
		}
		if err := o.Property.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpProp := func(ptr interface{}) { o.Property = *ptr.(**PropertyValue) }
	if err := w.ReadPointer(&o.Property, _s_lpProp, _ptr_lpProp); err != nil {
		return err
	}
	return nil
}

// ComparePropertiesRestriction structure represents ComparePropsRestriction_r RPC structure.
//
// The ComparePropsRestriction_r structure is an encoding of the ComparePropsRestriction
// data structure defined in [MS-OXCDATA]. The semantic meaning is unchanged from the
// ComparePropsRestriction data structure.
type ComparePropertiesRestriction struct {
	// relop:  Encodes the RelOp field of the ComparePropsRestriction data structure.
	Op uint32 `idl:"name:relop" json:"op"`
	// ulPropTag1:  Encodes the PropTag1 field of the ComparePropsRestriction data structure.
	PropertyTag1 uint32 `idl:"name:ulPropTag1" json:"property_tag1"`
	// ulPropTag2:  Encodes the PropTag2 field of the ComparePropsRestriction data structure.
	PropertyTag2 uint32 `idl:"name:ulPropTag2" json:"property_tag2"`
}

func (o *ComparePropertiesRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ComparePropertiesRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Op); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag1); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag2); err != nil {
		return err
	}
	return nil
}
func (o *ComparePropertiesRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Op); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag1); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag2); err != nil {
		return err
	}
	return nil
}

// SubRestriction structure represents SubRestriction_r RPC structure.
//
// The SubRestriction_r structure is an encoding of the SubObjectRestriction data structure
// defined in [MS-OXCDATA]. The semantic meaning is unchanged from the SubObjectRestriction
// data structure.
type SubRestriction struct {
	// ulSubObject:  Encodes the SubObject field of the SubObjectRestriction data structure.
	SubObject uint32 `idl:"name:ulSubObject" json:"sub_object"`
	// lpRes:  Encodes the Restriction field of the SubObjectRestriction data structure.
	Restriction *Restriction `idl:"name:lpRes" json:"restriction"`
}

func (o *SubRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SubObject); err != nil {
		return err
	}
	if o.Restriction != nil {
		_ptr_lpRes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Restriction != nil {
				if err := o.Restriction.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Restriction{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Restriction, _ptr_lpRes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubObject); err != nil {
		return err
	}
	_ptr_lpRes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Restriction == nil {
			o.Restriction = &Restriction{}
		}
		if err := o.Restriction.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpRes := func(ptr interface{}) { o.Restriction = *ptr.(**Restriction) }
	if err := w.ReadPointer(&o.Restriction, _s_lpRes, _ptr_lpRes); err != nil {
		return err
	}
	return nil
}

// SizeRestriction structure represents SizeRestriction_r RPC structure.
//
// The SizeRestriction_r structure is an encoding of the SizeRestriction data structure
// defined in [MS-OXCDATA]. The semantic meaning is unchanged from the SizeRestriction
// data structure.
type SizeRestriction struct {
	// relop:  Encodes the RelOp field of the SizeRestriction data structure.
	Op uint32 `idl:"name:relop" json:"op"`
	// ulPropTag:  Encodes the PropTag field of the SizeRestriction data structure.
	PropertyTag uint32 `idl:"name:ulPropTag" json:"property_tag"`
	// cb:  Encodes the Size field of the SizeRestriction data structure.
	Length uint32 `idl:"name:cb" json:"length"`
}

func (o *SizeRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SizeRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Op); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	return nil
}
func (o *SizeRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Op); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	return nil
}

// ExistRestriction structure represents ExistRestriction_r RPC structure.
//
// The ExistRestriction_r structure is an encoding of the ExistRestriction data structure
// defined in [MS-OXCDATA]. The semantic meaning is unchanged from the ExistRestriction
// data structure.
type ExistRestriction struct {
	// ulReserved1:  Reserved. All clients and servers MUST set this value to the constant
	// 0x00000000.
	_ uint32 `idl:"name:ulReserved1"`
	// ulPropTag:  Encodes the PropTag field of the ExistRestriction data structure.
	PropertyTag uint32 `idl:"name:ulPropTag" json:"property_tag"`
	// ulReserved2:  Reserved. All clients and servers MUST set this value to the constant
	// 0x00000000.
	_ uint32 `idl:"name:ulReserved2"`
}

func (o *ExistRestriction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExistRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// reserved ulReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyTag); err != nil {
		return err
	}
	// reserved ulReserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	return nil
}
func (o *ExistRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// reserved ulReserved1
	var _ulReserved1 uint32
	if err := w.ReadData(&_ulReserved1); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyTag); err != nil {
		return err
	}
	// reserved ulReserved2
	var _ulReserved2 uint32
	if err := w.ReadData(&_ulReserved2); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion structure represents RestrictionUnion_r RPC union.
//
// The RestrictionUnion_r structure encodes a single instance of any type of restriction.
// It is an aggregation data structure, allowing a single parameter to a Name Service
// Provider Interface (NSPI) method to contain any type of restriction data structure.
type RestrictionUnion struct {
	// Types that are assignable to Value
	//
	// *RestrictionUnion_And
	// *RestrictionUnion_Or
	// *RestrictionUnion_Not
	// *RestrictionUnion_Content
	// *RestrictionUnion_Property
	// *RestrictionUnion_CompareProperties
	// *RestrictionUnion_BitMask
	// *RestrictionUnion_Size
	// *RestrictionUnion_Exist
	// *RestrictionUnion_SubRestriction
	Value is_RestrictionUnion `json:"value"`
}

func (o *RestrictionUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RestrictionUnion_And:
		if value != nil {
			return value.And
		}
	case *RestrictionUnion_Or:
		if value != nil {
			return value.Or
		}
	case *RestrictionUnion_Not:
		if value != nil {
			return value.Not
		}
	case *RestrictionUnion_Content:
		if value != nil {
			return value.Content
		}
	case *RestrictionUnion_Property:
		if value != nil {
			return value.Property
		}
	case *RestrictionUnion_CompareProperties:
		if value != nil {
			return value.CompareProperties
		}
	case *RestrictionUnion_BitMask:
		if value != nil {
			return value.BitMask
		}
	case *RestrictionUnion_Size:
		if value != nil {
			return value.Size
		}
	case *RestrictionUnion_Exist:
		if value != nil {
			return value.Exist
		}
	case *RestrictionUnion_SubRestriction:
		if value != nil {
			return value.SubRestriction
		}
	}
	return nil
}

type is_RestrictionUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RestrictionUnion()
}

func (o *RestrictionUnion) NDRSwitchValue(sw int32) int32 {
	if o == nil {
		return int32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *RestrictionUnion_And:
		return int32(0)
	case *RestrictionUnion_Or:
		return int32(1)
	case *RestrictionUnion_Not:
		return int32(2)
	case *RestrictionUnion_Content:
		return int32(3)
	case *RestrictionUnion_Property:
		return int32(4)
	case *RestrictionUnion_CompareProperties:
		return int32(5)
	case *RestrictionUnion_BitMask:
		return int32(6)
	case *RestrictionUnion_Size:
		return int32(7)
	case *RestrictionUnion_Exist:
		return int32(8)
	case *RestrictionUnion_SubRestriction:
		return int32(9)
	}
	return int32(0)
}

func (o *RestrictionUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw int32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(int32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(0):
		_o, _ := o.Value.(*RestrictionUnion_And)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_And{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(1):
		_o, _ := o.Value.(*RestrictionUnion_Or)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_Or{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(2):
		_o, _ := o.Value.(*RestrictionUnion_Not)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_Not{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(3):
		_o, _ := o.Value.(*RestrictionUnion_Content)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_Content{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4):
		_o, _ := o.Value.(*RestrictionUnion_Property)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_Property{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(5):
		_o, _ := o.Value.(*RestrictionUnion_CompareProperties)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_CompareProperties{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(6):
		_o, _ := o.Value.(*RestrictionUnion_BitMask)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_BitMask{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(7):
		_o, _ := o.Value.(*RestrictionUnion_Size)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_Size{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(8):
		_o, _ := o.Value.(*RestrictionUnion_Exist)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_Exist{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(9):
		_o, _ := o.Value.(*RestrictionUnion_SubRestriction)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RestrictionUnion_SubRestriction{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *RestrictionUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw int32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*int32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(0):
		o.Value = &RestrictionUnion_And{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(1):
		o.Value = &RestrictionUnion_Or{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(2):
		o.Value = &RestrictionUnion_Not{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(3):
		o.Value = &RestrictionUnion_Content{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4):
		o.Value = &RestrictionUnion_Property{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(5):
		o.Value = &RestrictionUnion_CompareProperties{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(6):
		o.Value = &RestrictionUnion_BitMask{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(7):
		o.Value = &RestrictionUnion_Size{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(8):
		o.Value = &RestrictionUnion_Exist{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(9):
		o.Value = &RestrictionUnion_SubRestriction{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// RestrictionUnion_And structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 0
type RestrictionUnion_And struct {
	// resAnd:  RestrictionUnion_r contains an encoding of an AndRestriction.
	And *AndRestriction `idl:"name:resAnd" json:"and"`
}

func (*RestrictionUnion_And) is_RestrictionUnion() {}

func (o *RestrictionUnion_And) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.And != nil {
		if err := o.And.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AndRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_And) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.And == nil {
		o.And = &AndRestriction{}
	}
	if err := o.And.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_Or structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 1
type RestrictionUnion_Or struct {
	// resOr:  RestrictionUnion_r contains an encoding of an OrRestriction.
	Or *OrRestriction `idl:"name:resOr" json:"or"`
}

func (*RestrictionUnion_Or) is_RestrictionUnion() {}

func (o *RestrictionUnion_Or) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Or != nil {
		if err := o.Or.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OrRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_Or) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Or == nil {
		o.Or = &OrRestriction{}
	}
	if err := o.Or.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_Not structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 2
type RestrictionUnion_Not struct {
	// resNot:  RestrictionUnion_r contains an encoding of a NotRestriction.
	Not *NotRestriction `idl:"name:resNot" json:"not"`
}

func (*RestrictionUnion_Not) is_RestrictionUnion() {}

func (o *RestrictionUnion_Not) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Not != nil {
		if err := o.Not.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NotRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_Not) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Not == nil {
		o.Not = &NotRestriction{}
	}
	if err := o.Not.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_Content structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 3
type RestrictionUnion_Content struct {
	// resContent:  RestrictionUnion_r contains an encoding of a ContentRestriction.
	Content *ContentRestriction `idl:"name:resContent" json:"content"`
}

func (*RestrictionUnion_Content) is_RestrictionUnion() {}

func (o *RestrictionUnion_Content) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Content != nil {
		if err := o.Content.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ContentRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_Content) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Content == nil {
		o.Content = &ContentRestriction{}
	}
	if err := o.Content.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_Property structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 4
type RestrictionUnion_Property struct {
	// resProperty:  RestrictionUnion_r contains an encoding of a PropertyRestriction.
	Property *PropertyRestriction `idl:"name:resProperty" json:"property"`
}

func (*RestrictionUnion_Property) is_RestrictionUnion() {}

func (o *RestrictionUnion_Property) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Property != nil {
		if err := o.Property.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PropertyRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_Property) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Property == nil {
		o.Property = &PropertyRestriction{}
	}
	if err := o.Property.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_CompareProperties structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 5
type RestrictionUnion_CompareProperties struct {
	// resCompareProps:  RestrictionUnion_r contains an encoding of a CompareRestriction.
	CompareProperties *ComparePropertiesRestriction `idl:"name:resCompareProps" json:"compare_properties"`
}

func (*RestrictionUnion_CompareProperties) is_RestrictionUnion() {}

func (o *RestrictionUnion_CompareProperties) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CompareProperties != nil {
		if err := o.CompareProperties.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ComparePropertiesRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_CompareProperties) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.CompareProperties == nil {
		o.CompareProperties = &ComparePropertiesRestriction{}
	}
	if err := o.CompareProperties.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_BitMask structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 6
type RestrictionUnion_BitMask struct {
	// resBitMask:  RestrictionUnion_r contains an encoding of a BitMaskRestriction.
	BitMask *BitMaskRestriction `idl:"name:resBitMask" json:"bit_mask"`
}

func (*RestrictionUnion_BitMask) is_RestrictionUnion() {}

func (o *RestrictionUnion_BitMask) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BitMask != nil {
		if err := o.BitMask.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&BitMaskRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_BitMask) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.BitMask == nil {
		o.BitMask = &BitMaskRestriction{}
	}
	if err := o.BitMask.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_Size structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 7
type RestrictionUnion_Size struct {
	// resSize:  RestrictionUnion_r contains an encoding of a SizeRestriction.
	Size *SizeRestriction `idl:"name:resSize" json:"size"`
}

func (*RestrictionUnion_Size) is_RestrictionUnion() {}

func (o *RestrictionUnion_Size) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Size != nil {
		if err := o.Size.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SizeRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_Size) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Size == nil {
		o.Size = &SizeRestriction{}
	}
	if err := o.Size.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_Exist structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 8
type RestrictionUnion_Exist struct {
	// resExist:  RestrictionUnion_r contains an encoding of an ExistRestriction.
	Exist *ExistRestriction `idl:"name:resExist" json:"exist"`
}

func (*RestrictionUnion_Exist) is_RestrictionUnion() {}

func (o *RestrictionUnion_Exist) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Exist != nil {
		if err := o.Exist.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ExistRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_Exist) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Exist == nil {
		o.Exist = &ExistRestriction{}
	}
	if err := o.Exist.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RestrictionUnion_SubRestriction structure represents RestrictionUnion_r RPC union arm.
//
// It has following labels: 9
type RestrictionUnion_SubRestriction struct {
	// resSubRestriction:  RestrictionUnion_r contains an encoding of a SubRestriction.
	SubRestriction *SubRestriction `idl:"name:resSubRestriction" json:"sub_restriction"`
}

func (*RestrictionUnion_SubRestriction) is_RestrictionUnion() {}

func (o *RestrictionUnion_SubRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SubRestriction != nil {
		if err := o.SubRestriction.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SubRestriction{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RestrictionUnion_SubRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SubRestriction == nil {
		o.SubRestriction = &SubRestriction{}
	}
	if err := o.SubRestriction.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyName structure represents PropertyName_r RPC structure.
//
// The PropertyName_r structure is an encoding of the PropertyName data structure defined
// in [MS-OXCDATA]. The semantic meaning is unchanged from the PropertyName data structure.
type PropertyName struct {
	// lpguid:  Encodes the GUID field of the PropertyName data structure. This field is
	// encoded as a FlatUID_r data structure.
	GUID *FlatUID `idl:"name:lpguid" json:"guid"`
	// ulReserved:  Reserved. All clients and servers MUST set this value to the constant
	// 0x00000000.
	_ uint32 `idl:"name:ulReserved"`
	// lID:  Encodes the lID field of the PropertyName data structure. In addition to the
	// definition of the LID field, this field is always present in the PropertyName_r data
	// structure; it is not optional.
	ID int32 `idl:"name:lID" json:"id"`
}

func (o *PropertyName) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.GUID != nil {
		_ptr_lpguid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.GUID != nil {
				if err := o.GUID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&FlatUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GUID, _ptr_lpguid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved ulReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	return nil
}
func (o *PropertyName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_lpguid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.GUID == nil {
			o.GUID = &FlatUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpguid := func(ptr interface{}) { o.GUID = *ptr.(**FlatUID) }
	if err := w.ReadPointer(&o.GUID, _s_lpguid, _ptr_lpguid); err != nil {
		return err
	}
	// reserved ulReserved
	var _ulReserved uint32
	if err := w.ReadData(&_ulReserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	return nil
}

// PropertyNameSet structure represents PropertyNameSet_r RPC structure.
//
// The PropertyNameSet_r structure is used to aggregate a number of PropertyName_r structures
// into a single data structure.
type PropertyNameSet struct {
	// cNames:  The number of PropertyName_r structures in this aggregation. The value MUST
	// NOT exceed 100,000.
	NamesCount uint32 `idl:"name:cNames" json:"names_count"`
	// aNames:  The list of PropertyName_r structures in this aggregation.
	Names []*PropertyName `idl:"name:aNames;size_is:(cNames)" json:"names"`
}

func (o *PropertyNameSet) xxx_PreparePayload(ctx context.Context) error {
	if o.Names != nil && o.NamesCount == 0 {
		o.NamesCount = uint32(len(o.Names))
	}
	if o.NamesCount > uint32(100000) {
		return fmt.Errorf("NamesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *PropertyNameSet) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.NamesCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PropertyNameSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.NamesCount); err != nil {
		return err
	}
	for i1 := range o.Names {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Names[i1] != nil {
			if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&PropertyName{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyNameSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.NamesCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.NamesCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.NamesCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
	}
	o.Names = make([]*PropertyName, sizeInfo[0])
	for i1 := range o.Names {
		i1 := i1
		if o.Names[i1] == nil {
			o.Names[i1] = &PropertyName{}
		}
		if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CharStringsArray structure represents StringsArray_r RPC structure.
//
// The StringsArray_r structure is used to aggregate a number of character type strings
// into a single data structure.
type CharStringsArray struct {
	// Count:  The number of character string structures in this aggregation. The value
	// MUST NOT exceed 100,000.
	Count uint32 `idl:"name:Count" json:"count"`
	// Strings:  The list of character type strings in this aggregation. The strings in
	// this list are NULL-terminated.
	Strings string `idl:"name:Strings;size_is:(Count);string" json:"strings"`
}

func (o *CharStringsArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Strings != "" && o.Count == 0 {
		o.Count = uint32(len(o.Strings))
	}
	if o.Count > uint32(100000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CharStringsArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Strings != "" || o.Count > 0 {
		_ptr_Strings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.CharNLen(o.Strings)
			if dimLength1 > sizeInfo[0] {
				dimLength1 = sizeInfo[0]
			} else {
				sizeInfo[0] = dimLength1
			}
			if err := w.WriteSize(0); err != nil {
				return err
			}
			if err := w.WriteSize(dimLength1); err != nil {
				return err
			}
			_Strings_buf := []byte(o.Strings)
			if uint64(len(_Strings_buf)) > sizeInfo[0]-1 {
				_Strings_buf = _Strings_buf[:sizeInfo[0]-1]
			}
			if o.Strings != ndr.ZeroString {
				_Strings_buf = append(_Strings_buf, byte(0))
			}
			for i1 := range _Strings_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Strings_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Strings_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Strings, _ptr_Strings); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CharStringsArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_Strings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Strings_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Strings_buf", sizeInfo[0])
		}
		_Strings_buf = make([]byte, sizeInfo[0])
		for i1 := range _Strings_buf {
			i1 := i1
			if err := w.ReadData(&_Strings_buf[i1]); err != nil {
				return err
			}
		}
		o.Strings = strings.TrimRight(string(_Strings_buf), ndr.ZeroString)
		return nil
	})
	_s_Strings := func(ptr interface{}) { o.Strings = *ptr.(*string) }
	if err := w.ReadPointer(&o.Strings, _s_Strings, _ptr_Strings); err != nil {
		return err
	}
	return nil
}

// StringsArray structure represents WStringsArray_r RPC structure.
//
// The WStringsArray_r structure is used to aggregate a number of wchar_t type strings
// into a single data structure.
type StringsArray struct {
	// Count:  The number of character strings structures in this aggregation. The value
	// MUST NOT exceed 100,000.
	Count uint32 `idl:"name:Count" json:"count"`
	// Strings:  The list of wchar_t type strings in this aggregation. The strings in this
	// list are NULL-terminated.
	Strings string `idl:"name:Strings;size_is:(Count);string" json:"strings"`
}

func (o *StringsArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Strings != "" && o.Count == 0 {
		o.Count = uint32(len(o.Strings))
	}
	if o.Count > uint32(100000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StringsArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Strings != "" || o.Count > 0 {
		_ptr_Strings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.CharNLen(o.Strings)
			if dimLength1 > sizeInfo[0] {
				dimLength1 = sizeInfo[0]
			} else {
				sizeInfo[0] = dimLength1
			}
			if err := w.WriteSize(0); err != nil {
				return err
			}
			if err := w.WriteSize(dimLength1); err != nil {
				return err
			}
			_Strings_buf := []byte(o.Strings)
			if uint64(len(_Strings_buf)) > sizeInfo[0]-1 {
				_Strings_buf = _Strings_buf[:sizeInfo[0]-1]
			}
			if o.Strings != ndr.ZeroString {
				_Strings_buf = append(_Strings_buf, byte(0))
			}
			for i1 := range _Strings_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Strings_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Strings_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Strings, _ptr_Strings); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StringsArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_Strings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Strings_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Strings_buf", sizeInfo[0])
		}
		_Strings_buf = make([]byte, sizeInfo[0])
		for i1 := range _Strings_buf {
			i1 := i1
			if err := w.ReadData(&_Strings_buf[i1]); err != nil {
				return err
			}
		}
		o.Strings = strings.TrimRight(string(_Strings_buf), ndr.ZeroString)
		return nil
	})
	_s_Strings := func(ptr interface{}) { o.Strings = *ptr.(*string) }
	if err := w.ReadPointer(&o.Strings, _s_Strings, _ptr_Strings); err != nil {
		return err
	}
	return nil
}

// Stat structure represents STAT RPC structure.
//
// The STAT structure is used to specify the state of a table and location information
// that applies to that table. It appears as both an input parameter and an output parameter
// in many Name Service Provider Interface (NSPI) methods.
type Stat struct {
	// SortType:  This field contains a DWORD representing a sort order. The client sets
	// this field to specify the sort type of this table. Possible values are specified
	// in Table Sort Orders (section 2.2.10). The server MUST NOT modify this field.
	SortType uint32 `idl:"name:SortType" json:"sort_type"`
	// ContainerID:  This field contains a MId. The client sets this field to specify the
	// MId of the address book container that this STAT represents. The client obtains these
	// MIds from the server's address book hierarchy table. The server MUST NOT modify this
	// field in any NSPI method except NspiGetMatches.
	ContainerID uint32 `idl:"name:ContainerID" json:"container_id"`
	// CurrentRec:  This field contains a MId. The client sets this field to specify a beginning
	// position in the table for the start of an NSPI method. The server sets this field
	// to report the end position in the table after processing an NSPI method.
	CurrentRecord uint32 `idl:"name:CurrentRec" json:"current_record"`
	// Delta:  This field contains a long value. The client sets this field to specify an
	// offset from the beginning position in the table for the start of an NSPI method.
	// If the NSPI method returns a success value, the server MUST set this field to 0.
	Delta int32 `idl:"name:Delta" json:"delta"`
	// NumPos:  This field contains a DWORD value specifying a position in the table. The
	// client sets this field to specify a fractional position for the beginning position
	// in the table for the start of an NSPI method (section 3.1.1.4.2). The server sets
	// this field to specify the approximate fractional position at the end of an NSPI method.
	// This value is a zero index; the first element in a table has the numeric position
	// 0. Although the protocol places no boundary or requirements on the accuracy of the
	// approximation the server reports, it is recommended that implementations maximize
	// the accuracy of the approximation to improve usability of the NSPI server for clients.
	PosLength uint32 `idl:"name:NumPos" json:"pos_length"`
	// TotalRecs:  This field contains a DWORD specifying the number of rows in the table.
	// The client sets this field to specify a fractional position for the beginning position
	// in the table for the start of an NSPI method (section 3.1.1.4.2). The server sets
	// this field to specify the total number of rows in the table. Unlike the NumPos field,
	// the server MUST report this number accurately; an approximation is insufficient.
	TotalRecords uint32 `idl:"name:TotalRecs" json:"total_records"`
	// CodePage:  This field contains a DWORD value representing a codepage. The client
	// sets this field to specify the codepage the client uses for non-Unicode strings.
	// The server MUST use this value during string handling (section 3.1.1.2). The server
	// MUST NOT modify this field.
	CodePage uint32 `idl:"name:CodePage" json:"code_page"`
	// TemplateLocale:  This field contains a DWORD value representing a language code identifier
	// (LCID). The client sets this field to specify the LCID associated with the template
	// the client wishes the server to return. The server MUST NOT modify this field.
	TemplateLocale uint32 `idl:"name:TemplateLocale" json:"template_locale"`
	// SortLocale:  This field contains a DWORD value representing an LCID. The client sets
	// this field to specify the LCID that it wishes the server to use when sorting any
	// strings. The server MUST use this value during sorting (section 3.1.1.2). The server
	// MUST NOT modify this field.
	SortLocale uint32 `idl:"name:SortLocale" json:"sort_locale"`
}

func (o *Stat) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Stat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SortType); err != nil {
		return err
	}
	if err := w.WriteData(o.ContainerID); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentRecord); err != nil {
		return err
	}
	if err := w.WriteData(o.Delta); err != nil {
		return err
	}
	if err := w.WriteData(o.PosLength); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalRecords); err != nil {
		return err
	}
	if err := w.WriteData(o.CodePage); err != nil {
		return err
	}
	if err := w.WriteData(o.TemplateLocale); err != nil {
		return err
	}
	if err := w.WriteData(o.SortLocale); err != nil {
		return err
	}
	return nil
}
func (o *Stat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SortType); err != nil {
		return err
	}
	if err := w.ReadData(&o.ContainerID); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentRecord); err != nil {
		return err
	}
	if err := w.ReadData(&o.Delta); err != nil {
		return err
	}
	if err := w.ReadData(&o.PosLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalRecords); err != nil {
		return err
	}
	if err := w.ReadData(&o.CodePage); err != nil {
		return err
	}
	if err := w.ReadData(&o.TemplateLocale); err != nil {
		return err
	}
	if err := w.ReadData(&o.SortLocale); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion structure represents PROP_VAL_UNION RPC union.
//
// The PROP_VAL_UNION structure encodes a single instance of any type of property value.
// It is an aggregation data structure, allowing a single parameter to an NSPI method
// to contain any type of property value.
type PropertyValueUnion struct {
	// Types that are assignable to Value
	//
	// *PropertyValueUnion_Int16
	// *PropertyValueUnion_Int32
	// *PropertyValueUnion_Bool
	// *PropertyValueUnion_CharString
	// *PropertyValueUnion_Binary
	// *PropertyValueUnion_String
	// *PropertyValueUnion_GUID
	// *PropertyValueUnion_DateTime
	// *PropertyValueUnion_Error
	// *PropertyValueUnion_Int16Array
	// *PropertyValueUnion_Int32Array
	// *PropertyValueUnion_CharStringArray
	// *PropertyValueUnion_BinaryArray
	// *PropertyValueUnion_GUIDArray
	// *PropertyValueUnion_StringArray
	// *PropertyValueUnion_DateTimeArray
	// *PropertyValueUnion_Reserved
	Value is_PropertyValueUnion `json:"value"`
}

func (o *PropertyValueUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PropertyValueUnion_Int16:
		if value != nil {
			return value.Int16
		}
	case *PropertyValueUnion_Int32:
		if value != nil {
			return value.Int32
		}
	case *PropertyValueUnion_Bool:
		if value != nil {
			return value.Bool
		}
	case *PropertyValueUnion_CharString:
		if value != nil {
			return value.CharString
		}
	case *PropertyValueUnion_Binary:
		if value != nil {
			return value.Binary
		}
	case *PropertyValueUnion_String:
		if value != nil {
			return value.String
		}
	case *PropertyValueUnion_GUID:
		if value != nil {
			return value.GUID
		}
	case *PropertyValueUnion_DateTime:
		if value != nil {
			return value.DateTime
		}
	case *PropertyValueUnion_Error:
		if value != nil {
			return value.Error
		}
	case *PropertyValueUnion_Int16Array:
		if value != nil {
			return value.Int16Array
		}
	case *PropertyValueUnion_Int32Array:
		if value != nil {
			return value.Int32Array
		}
	case *PropertyValueUnion_CharStringArray:
		if value != nil {
			return value.CharStringArray
		}
	case *PropertyValueUnion_BinaryArray:
		if value != nil {
			return value.BinaryArray
		}
	case *PropertyValueUnion_GUIDArray:
		if value != nil {
			return value.GUIDArray
		}
	case *PropertyValueUnion_StringArray:
		if value != nil {
			return value.StringArray
		}
	case *PropertyValueUnion_DateTimeArray:
		if value != nil {
			return value.DateTimeArray
		}
	}
	return nil
}

type is_PropertyValueUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PropertyValueUnion()
}

func (o *PropertyValueUnion) NDRSwitchValue(sw int32) int32 {
	if o == nil {
		return int32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PropertyValueUnion_Int16:
		return int32(2)
	case *PropertyValueUnion_Int32:
		return int32(3)
	case *PropertyValueUnion_Bool:
		return int32(11)
	case *PropertyValueUnion_CharString:
		return int32(30)
	case *PropertyValueUnion_Binary:
		return int32(258)
	case *PropertyValueUnion_String:
		return int32(31)
	case *PropertyValueUnion_GUID:
		return int32(72)
	case *PropertyValueUnion_DateTime:
		return int32(64)
	case *PropertyValueUnion_Error:
		return int32(10)
	case *PropertyValueUnion_Int16Array:
		return int32(4098)
	case *PropertyValueUnion_Int32Array:
		return int32(4099)
	case *PropertyValueUnion_CharStringArray:
		return int32(4126)
	case *PropertyValueUnion_BinaryArray:
		return int32(4354)
	case *PropertyValueUnion_GUIDArray:
		return int32(4168)
	case *PropertyValueUnion_StringArray:
		return int32(4127)
	case *PropertyValueUnion_DateTimeArray:
		return int32(4160)
	case *PropertyValueUnion_Reserved:
		switch sw {
		case int32(1),
			int32(13):
			return sw
		}
		return int32(1)
	}
	return int32(0)
}

func (o *PropertyValueUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw int32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(int32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(2):
		_o, _ := o.Value.(*PropertyValueUnion_Int16)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Int16{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(3):
		_o, _ := o.Value.(*PropertyValueUnion_Int32)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Int32{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(11):
		_o, _ := o.Value.(*PropertyValueUnion_Bool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Bool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(30):
		_o, _ := o.Value.(*PropertyValueUnion_CharString)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_CharString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(258):
		_o, _ := o.Value.(*PropertyValueUnion_Binary)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Binary{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(31):
		_o, _ := o.Value.(*PropertyValueUnion_String)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(72):
		_o, _ := o.Value.(*PropertyValueUnion_GUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(64):
		_o, _ := o.Value.(*PropertyValueUnion_DateTime)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_DateTime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(10):
		_o, _ := o.Value.(*PropertyValueUnion_Error)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Error{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4098):
		_o, _ := o.Value.(*PropertyValueUnion_Int16Array)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Int16Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4099):
		_o, _ := o.Value.(*PropertyValueUnion_Int32Array)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Int32Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4126):
		_o, _ := o.Value.(*PropertyValueUnion_CharStringArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_CharStringArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4354):
		_o, _ := o.Value.(*PropertyValueUnion_BinaryArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_BinaryArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4168):
		_o, _ := o.Value.(*PropertyValueUnion_GUIDArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_GUIDArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4127):
		_o, _ := o.Value.(*PropertyValueUnion_StringArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_StringArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(4160):
		_o, _ := o.Value.(*PropertyValueUnion_DateTimeArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_DateTimeArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(1),
		int32(13):
		_o, _ := o.Value.(*PropertyValueUnion_Reserved)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValueUnion_Reserved{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *PropertyValueUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw int32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*int32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(2):
		o.Value = &PropertyValueUnion_Int16{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(3):
		o.Value = &PropertyValueUnion_Int32{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(11):
		o.Value = &PropertyValueUnion_Bool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(30):
		o.Value = &PropertyValueUnion_CharString{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(258):
		o.Value = &PropertyValueUnion_Binary{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(31):
		o.Value = &PropertyValueUnion_String{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(72):
		o.Value = &PropertyValueUnion_GUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(64):
		o.Value = &PropertyValueUnion_DateTime{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(10):
		o.Value = &PropertyValueUnion_Error{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4098):
		o.Value = &PropertyValueUnion_Int16Array{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4099):
		o.Value = &PropertyValueUnion_Int32Array{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4126):
		o.Value = &PropertyValueUnion_CharStringArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4354):
		o.Value = &PropertyValueUnion_BinaryArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4168):
		o.Value = &PropertyValueUnion_GUIDArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4127):
		o.Value = &PropertyValueUnion_StringArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(4160):
		o.Value = &PropertyValueUnion_DateTimeArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(1),
		int32(13):
		o.Value = &PropertyValueUnion_Reserved{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// PropertyValueUnion_Int16 structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 2
type PropertyValueUnion_Int16 struct {
	// i:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single 16-bit integer value.
	Int16 int16 `idl:"name:i" json:"int16"`
}

func (*PropertyValueUnion_Int16) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Int16) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Int16); err != nil {
		return err
	}
	return nil
}
func (o *PropertyValueUnion_Int16) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Int16); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_Int32 structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 3
type PropertyValueUnion_Int32 struct {
	// l:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single 32-bit integer value.
	Int32 int32 `idl:"name:l" json:"int32"`
}

func (*PropertyValueUnion_Int32) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Int32) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Int32); err != nil {
		return err
	}
	return nil
}
func (o *PropertyValueUnion_Int32) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Int32); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_Bool structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 11
type PropertyValueUnion_Bool struct {
	// b:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single Boolean value. The client and server MUST NOT set this to values other than
	// 1 or 0.
	Bool uint16 `idl:"name:b" json:"bool"`
}

func (*PropertyValueUnion_Bool) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Bool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Bool); err != nil {
		return err
	}
	return nil
}
func (o *PropertyValueUnion_Bool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Bool); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_CharString structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 30
type PropertyValueUnion_CharString struct {
	// lpszA:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single 8-bit character string value. This value is NULL-terminated.
	CharString string `idl:"name:lpszA;string" json:"char_string"`
}

func (*PropertyValueUnion_CharString) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_CharString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CharString != "" {
		_ptr_lpszA := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.CharString); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CharString, _ptr_lpszA); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_CharString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_lpszA := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.CharString); err != nil {
			return err
		}
		return nil
	})
	_s_lpszA := func(ptr interface{}) { o.CharString = *ptr.(*string) }
	if err := w.ReadPointer(&o.CharString, _s_lpszA, _ptr_lpszA); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_Binary structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 258
type PropertyValueUnion_Binary struct {
	// bin:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single binary data value. The number of bytes that can be encoded in this structure
	// is 2,097,152.
	Binary *Binary `idl:"name:bin" json:"binary"`
}

func (*PropertyValueUnion_Binary) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Binary) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Binary != nil {
		if err := o.Binary.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Binary{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_Binary) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Binary == nil {
		o.Binary = &Binary{}
	}
	if err := o.Binary.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_String structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 31
type PropertyValueUnion_String struct {
	// lpszW:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single Unicode string value. This value is NULL-terminated.
	String string `idl:"name:lpszW;string" json:"string"`
}

func (*PropertyValueUnion_String) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.String != "" {
		_ptr_lpszW := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.String); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_lpszW); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_lpszW := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.String); err != nil {
			return err
		}
		return nil
	})
	_s_lpszW := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_lpszW, _ptr_lpszW); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_GUID structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 72
type PropertyValueUnion_GUID struct {
	// lpguid:  PROP_VAL_UNION contains an encoding of the value of a property that can
	// contain a single GUID value. The value is encoded as a FlatUID_r data structure.
	GUID *FlatUID `idl:"name:lpguid" json:"guid"`
}

func (*PropertyValueUnion_GUID) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_GUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GUID != nil {
		_ptr_lpguid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.GUID != nil {
				if err := o.GUID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&FlatUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GUID, _ptr_lpguid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_GUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_lpguid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.GUID == nil {
			o.GUID = &FlatUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpguid := func(ptr interface{}) { o.GUID = *ptr.(**FlatUID) }
	if err := w.ReadPointer(&o.GUID, _s_lpguid, _ptr_lpguid); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_DateTime structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 64
type PropertyValueUnion_DateTime struct {
	// ft:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single 64-bit integer value. The value is encoded as a FILETIME structure.
	DateTime *dtyp.Filetime `idl:"name:ft" json:"date_time"`
}

func (*PropertyValueUnion_DateTime) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_DateTime) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DateTime != nil {
		if err := o.DateTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_DateTime) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DateTime == nil {
		o.DateTime = &dtyp.Filetime{}
	}
	if err := o.DateTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_Error structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 10
type PropertyValueUnion_Error struct {
	// err:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// a single PtypErrorCode value.
	Error int32 `idl:"name:err" json:"error"`
}

func (*PropertyValueUnion_Error) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Error) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	return nil
}
func (o *PropertyValueUnion_Error) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_Int16Array structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 4098
type PropertyValueUnion_Int16Array struct {
	// MVi:  PROP_VAL_UNION contains an encoding of the values of a property that can contain
	// multiple 16-bit integer values. The number of values that can be encoded in this
	// structure is 100,000.
	Int16Array *ShortArray `idl:"name:MVi" json:"int16_array"`
}

func (*PropertyValueUnion_Int16Array) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Int16Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Int16Array != nil {
		if err := o.Int16Array.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ShortArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_Int16Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Int16Array == nil {
		o.Int16Array = &ShortArray{}
	}
	if err := o.Int16Array.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_Int32Array structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 4099
type PropertyValueUnion_Int32Array struct {
	// MVl:  PROP_VAL_UNION contains an encoding of the values of a property that can contain
	// multiple 32-bit integer values. The number of values that can be encoded in this
	// structure is 100,000.
	Int32Array *LongArray `idl:"name:MVl" json:"int32_array"`
}

func (*PropertyValueUnion_Int32Array) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Int32Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Int32Array != nil {
		if err := o.Int32Array.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LongArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_Int32Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Int32Array == nil {
		o.Int32Array = &LongArray{}
	}
	if err := o.Int32Array.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_CharStringArray structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 4126
type PropertyValueUnion_CharStringArray struct {
	// MVszA:  PROP_VAL_UNION contains an encoding of the values of a property that can
	// contain multiple 8-bit character string values. These string values are NULL-terminated.
	// The number of values that can be encoded in this structure is 100,000.
	CharStringArray *CharStringArray `idl:"name:MVszA" json:"char_string_array"`
}

func (*PropertyValueUnion_CharStringArray) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_CharStringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CharStringArray != nil {
		if err := o.CharStringArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CharStringArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_CharStringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.CharStringArray == nil {
		o.CharStringArray = &CharStringArray{}
	}
	if err := o.CharStringArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_BinaryArray structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 4354
type PropertyValueUnion_BinaryArray struct {
	// MVbin:  PROP_VAL_UNION contains an encoding of the values of a property that can
	// contain multiple binary data values. The number of bytes that can be encoded in each
	// value of this structure is 2,097,152. The number of values that can be encoded in
	// this structure is 100,000.
	BinaryArray *BinaryArray `idl:"name:MVbin" json:"binary_array"`
}

func (*PropertyValueUnion_BinaryArray) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_BinaryArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BinaryArray != nil {
		if err := o.BinaryArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&BinaryArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_BinaryArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.BinaryArray == nil {
		o.BinaryArray = &BinaryArray{}
	}
	if err := o.BinaryArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_GUIDArray structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 4168
type PropertyValueUnion_GUIDArray struct {
	// MVguid:  PROP_VAL_UNION contains an encoding of the values of a property that can
	// contain multiple GUID values. The values are encoded as FlatUID_r data structures.
	// The number of values that can be encoded in this structure is 100,000.
	GUIDArray *FlatUIDArray `idl:"name:MVguid" json:"guid_array"`
}

func (*PropertyValueUnion_GUIDArray) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_GUIDArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GUIDArray != nil {
		if err := o.GUIDArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FlatUIDArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_GUIDArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GUIDArray == nil {
		o.GUIDArray = &FlatUIDArray{}
	}
	if err := o.GUIDArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_StringArray structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 4127
type PropertyValueUnion_StringArray struct {
	// MVszW:  PROP_VAL_UNION contains an encoding of the values of a property that can
	// contain multiple Unicode string values. These string values are NULL-terminated.
	// The number of values that can be encoded in this structure is 100,000.
	StringArray *StringArray `idl:"name:MVszW" json:"string_array"`
}

func (*PropertyValueUnion_StringArray) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_StringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StringArray != nil {
		if err := o.StringArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StringArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_StringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.StringArray == nil {
		o.StringArray = &StringArray{}
	}
	if err := o.StringArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_DateTimeArray structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 4160
type PropertyValueUnion_DateTimeArray struct {
	// MVft:  PROP_VAL_UNION contains an encoding of the value of a property that can contain
	// multiple 64-bit integer values. The values are encoded as FILETIME structures. The
	// number of values that can be encoded in this structure is 100,000.
	DateTimeArray *DateTimeArray `idl:"name:MVft" json:"date_time_array"`
}

func (*PropertyValueUnion_DateTimeArray) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_DateTimeArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DateTimeArray != nil {
		if err := o.DateTimeArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTimeArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyValueUnion_DateTimeArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DateTimeArray == nil {
		o.DateTimeArray = &DateTimeArray{}
	}
	if err := o.DateTimeArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyValueUnion_Reserved structure represents PROP_VAL_UNION RPC union arm.
//
// It has following labels: 1, 13
type PropertyValueUnion_Reserved struct {
	// lReserved:  Reserved. All clients and servers MUST set this value to the constant
	// 0x00000000.
	_ int32 `idl:"name:lReserved"`
}

func (*PropertyValueUnion_Reserved) is_PropertyValueUnion() {}

func (o *PropertyValueUnion_Reserved) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	// reserved lReserved
	if err := w.WriteData(int32(0)); err != nil {
		return err
	}
	return nil
}
func (o *PropertyValueUnion_Reserved) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	// reserved lReserved
	var _lReserved int32
	if err := w.ReadData(&_lReserved); err != nil {
		return err
	}
	return nil
}

// Handle structure represents NSPI_HANDLE RPC structure.
type Handle dcetypes.ContextHandle

func (o *Handle) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Handle) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Handle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Handle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultNspiClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultNspiClient) Bind(ctx context.Context, in *BindRequest, opts ...dcerpc.CallOption) (*BindResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BindResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) Unbind(ctx context.Context, in *UnbindRequest, opts ...dcerpc.CallOption) (*UnbindResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnbindResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) UpdateStat(ctx context.Context, in *UpdateStatRequest, opts ...dcerpc.CallOption) (*UpdateStatResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateStatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) QueryRows(ctx context.Context, in *QueryRowsRequest, opts ...dcerpc.CallOption) (*QueryRowsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryRowsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) SeekEntries(ctx context.Context, in *SeekEntriesRequest, opts ...dcerpc.CallOption) (*SeekEntriesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SeekEntriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) GetMatches(ctx context.Context, in *GetMatchesRequest, opts ...dcerpc.CallOption) (*GetMatchesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMatchesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) ResortRestriction(ctx context.Context, in *ResortRestrictionRequest, opts ...dcerpc.CallOption) (*ResortRestrictionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResortRestrictionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) DNToMID(ctx context.Context, in *DNToMIDRequest, opts ...dcerpc.CallOption) (*DNToMIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DNToMIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) GetPropertyList(ctx context.Context, in *GetPropertyListRequest, opts ...dcerpc.CallOption) (*GetPropertyListResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertyListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) CompareMIDs(ctx context.Context, in *CompareMIDsRequest, opts ...dcerpc.CallOption) (*CompareMIDsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CompareMIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) ModifyProperties(ctx context.Context, in *ModifyPropertiesRequest, opts ...dcerpc.CallOption) (*ModifyPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ModifyPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) GetSpecialTable(ctx context.Context, in *GetSpecialTableRequest, opts ...dcerpc.CallOption) (*GetSpecialTableResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSpecialTableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) GetTemplateInfo(ctx context.Context, in *GetTemplateInfoRequest, opts ...dcerpc.CallOption) (*GetTemplateInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTemplateInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) ModifyLinkAttribute(ctx context.Context, in *ModifyLinkAttributeRequest, opts ...dcerpc.CallOption) (*ModifyLinkAttributeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ModifyLinkAttributeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) QueryColumns(ctx context.Context, in *QueryColumnsRequest, opts ...dcerpc.CallOption) (*QueryColumnsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryColumnsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) GetNamesFromIDs(ctx context.Context, in *GetNamesFromIDsRequest, opts ...dcerpc.CallOption) (*GetNamesFromIDsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNamesFromIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) GetIDsFromNames(ctx context.Context, in *GetIDsFromNamesRequest, opts ...dcerpc.CallOption) (*GetIDsFromNamesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetIDsFromNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) ResolveNames(ctx context.Context, in *ResolveNamesRequest, opts ...dcerpc.CallOption) (*ResolveNamesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResolveNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) ResolveNamesW(ctx context.Context, in *ResolveNamesWRequest, opts ...dcerpc.CallOption) (*ResolveNamesWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResolveNamesWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNspiClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultNspiClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewNspiClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (NspiClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(NspiSyntaxV56_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultNspiClient{cc: cc}, nil
}

// xxx_BindOperation structure represents the NspiBind operation
type xxx_BindOperation struct {
	Flags         uint32   `idl:"name:dwFlags" json:"flags"`
	Stat          *Stat    `idl:"name:pStat" json:"stat"`
	ServerGUID    *FlatUID `idl:"name:pServerGuid;pointer:unique" json:"server_guid"`
	ContextHandle *Handle  `idl:"name:contextHandle;pointer:ref" json:"context_handle"`
	Return        int32    `idl:"name:Return" json:"return"`
}

func (o *xxx_BindOperation) OpNum() int { return 0 }

func (o *xxx_BindOperation) OpName() string { return "/nspi/v56/NspiBind" }

func (o *xxx_BindOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pServerGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=FlatUID_r}(struct))
	{
		if o.ServerGUID != nil {
			_ptr_pServerGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerGUID != nil {
					if err := o.ServerGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&FlatUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerGUID, _ptr_pServerGuid); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pServerGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=FlatUID_r}(struct))
	{
		_ptr_pServerGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerGUID == nil {
				o.ServerGUID = &FlatUID{}
			}
			if err := o.ServerGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pServerGuid := func(ptr interface{}) { o.ServerGUID = *ptr.(**FlatUID) }
		if err := w.ReadPointer(&o.ServerGUID, _s_pServerGuid, _ptr_pServerGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pServerGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=FlatUID_r}(struct))
	{
		if o.ServerGUID != nil {
			_ptr_pServerGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerGUID != nil {
					if err := o.ServerGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&FlatUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerGUID, _ptr_pServerGuid); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// contextHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pServerGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=FlatUID_r}(struct))
	{
		_ptr_pServerGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerGUID == nil {
				o.ServerGUID = &FlatUID{}
			}
			if err := o.ServerGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pServerGuid := func(ptr interface{}) { o.ServerGUID = *ptr.(**FlatUID) }
		if err := w.ReadPointer(&o.ServerGUID, _s_pServerGuid, _ptr_pServerGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// contextHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BindRequest structure represents the NspiBind operation request
type BindRequest struct {
	// dwFlags:  A DWORD value containing a set of bit flags. The server MUST ignore values
	// other than the bit flag fAnonymousLogin.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to specify both input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// pServerGuid: The value NULL or a pointer to a GUID value associated with the specific
	// NSPI server.
	ServerGUID *FlatUID `idl:"name:pServerGuid;pointer:unique" json:"server_guid"`
}

func (o *BindRequest) xxx_ToOp(ctx context.Context, op *xxx_BindOperation) *xxx_BindOperation {
	if op == nil {
		op = &xxx_BindOperation{}
	}
	if o == nil {
		return op
	}
	op.Flags = o.Flags
	op.Stat = o.Stat
	op.ServerGUID = o.ServerGUID
	return op
}

func (o *BindRequest) xxx_FromOp(ctx context.Context, op *xxx_BindOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.Stat = op.Stat
	o.ServerGUID = op.ServerGUID
}
func (o *BindRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BindRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BindResponse structure represents the NspiBind operation response
type BindResponse struct {
	// pServerGuid: The value NULL or a pointer to a GUID value associated with the specific
	// NSPI server.
	ServerGUID *FlatUID `idl:"name:pServerGuid;pointer:unique" json:"server_guid"`
	// contextHandle: An RPC context handle as specified in section 2.3.9.
	ContextHandle *Handle `idl:"name:contextHandle;pointer:ref" json:"context_handle"`
	// Return: The NspiBind return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BindResponse) xxx_ToOp(ctx context.Context, op *xxx_BindOperation) *xxx_BindOperation {
	if op == nil {
		op = &xxx_BindOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerGUID = o.ServerGUID
	op.ContextHandle = o.ContextHandle
	op.Return = o.Return
	return op
}

func (o *BindResponse) xxx_FromOp(ctx context.Context, op *xxx_BindOperation) {
	if o == nil {
		return
	}
	o.ServerGUID = op.ServerGUID
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
}
func (o *BindResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BindResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnbindOperation structure represents the NspiUnbind operation
type xxx_UnbindOperation struct {
	ContextHandle *Handle `idl:"name:contextHandle" json:"context_handle"`
	_             uint32  `idl:"name:Reserved"`
	Return        uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_UnbindOperation) OpNum() int { return 1 }

func (o *xxx_UnbindOperation) OpName() string { return "/nspi/v56/NspiUnbind" }

func (o *xxx_UnbindOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnbindOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// contextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnbindOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// contextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnbindOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnbindOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// contextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnbindOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// contextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UnbindRequest structure represents the NspiUnbind operation request
type UnbindRequest struct {
	// contextHandle: An RPC context handle as specified in section 2.3.9.
	ContextHandle *Handle `idl:"name:contextHandle" json:"context_handle"`
}

func (o *UnbindRequest) xxx_ToOp(ctx context.Context, op *xxx_UnbindOperation) *xxx_UnbindOperation {
	if op == nil {
		op = &xxx_UnbindOperation{}
	}
	if o == nil {
		return op
	}
	op.ContextHandle = o.ContextHandle
	return op
}

func (o *UnbindRequest) xxx_FromOp(ctx context.Context, op *xxx_UnbindOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
}
func (o *UnbindRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UnbindRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnbindOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnbindResponse structure represents the NspiUnbind operation response
type UnbindResponse struct {
	// contextHandle: An RPC context handle as specified in section 2.3.9.
	ContextHandle *Handle `idl:"name:contextHandle" json:"context_handle"`
	// Return: The NspiUnbind return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UnbindResponse) xxx_ToOp(ctx context.Context, op *xxx_UnbindOperation) *xxx_UnbindOperation {
	if op == nil {
		op = &xxx_UnbindOperation{}
	}
	if o == nil {
		return op
	}
	op.ContextHandle = o.ContextHandle
	op.Return = o.Return
	return op
}

func (o *UnbindResponse) xxx_FromOp(ctx context.Context, op *xxx_UnbindOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
}
func (o *UnbindResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UnbindResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnbindOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateStatOperation structure represents the NspiUpdateStat operation
type xxx_UpdateStatOperation struct {
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	_      uint32  `idl:"name:Reserved"`
	Stat   *Stat   `idl:"name:pStat" json:"stat"`
	Delta  int32   `idl:"name:plDelta;pointer:unique" json:"delta"`
	Return int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateStatOperation) OpNum() int { return 2 }

func (o *xxx_UpdateStatOperation) OpName() string { return "/nspi/v56/NspiUpdateStat" }

func (o *xxx_UpdateStatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateStatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// plDelta {in, out} (1:{pointer=unique}*(1)(int32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_plDelta := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Delta); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Delta, _ptr_plDelta); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateStatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// plDelta {in, out} (1:{pointer=unique}*(1)(int32))
	{
		_ptr_plDelta := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Delta); err != nil {
				return err
			}
			return nil
		})
		_s_plDelta := func(ptr interface{}) { o.Delta = *ptr.(*int32) }
		if err := w.ReadPointer(&o.Delta, _s_plDelta, _ptr_plDelta); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateStatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateStatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// plDelta {in, out} (1:{pointer=unique}*(1)(int32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_plDelta := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Delta); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Delta, _ptr_plDelta); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateStatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// plDelta {in, out} (1:{pointer=unique}*(1)(int32))
	{
		_ptr_plDelta := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Delta); err != nil {
				return err
			}
			return nil
		})
		_s_plDelta := func(ptr interface{}) { o.Delta = *ptr.(*int32) }
		if err := w.ReadPointer(&o.Delta, _s_plDelta, _ptr_plDelta); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateStatRequest structure represents the NspiUpdateStat operation request
type UpdateStatRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to specify both input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// plDelta: The value NULL or a pointer to a long value indicating movement within the
	// address book container specified by the input parameter pStat.
	Delta int32 `idl:"name:plDelta;pointer:unique" json:"delta"`
}

func (o *UpdateStatRequest) xxx_ToOp(ctx context.Context, op *xxx_UpdateStatOperation) *xxx_UpdateStatOperation {
	if op == nil {
		op = &xxx_UpdateStatOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.Delta = o.Delta
	return op
}

func (o *UpdateStatRequest) xxx_FromOp(ctx context.Context, op *xxx_UpdateStatOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.Delta = op.Delta
}
func (o *UpdateStatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UpdateStatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateStatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateStatResponse structure represents the NspiUpdateStat operation response
type UpdateStatResponse struct {
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to specify both input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// plDelta: The value NULL or a pointer to a long value indicating movement within the
	// address book container specified by the input parameter pStat.
	Delta int32 `idl:"name:plDelta;pointer:unique" json:"delta"`
	// Return: The NspiUpdateStat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UpdateStatResponse) xxx_ToOp(ctx context.Context, op *xxx_UpdateStatOperation) *xxx_UpdateStatOperation {
	if op == nil {
		op = &xxx_UpdateStatOperation{}
	}
	if o == nil {
		return op
	}
	op.Stat = o.Stat
	op.Delta = o.Delta
	op.Return = o.Return
	return op
}

func (o *UpdateStatResponse) xxx_FromOp(ctx context.Context, op *xxx_UpdateStatOperation) {
	if o == nil {
		return
	}
	o.Stat = op.Stat
	o.Delta = op.Delta
	o.Return = op.Return
}
func (o *UpdateStatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UpdateStatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateStatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryRowsOperation structure represents the NspiQueryRows operation
type xxx_QueryRowsOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	Flags        uint32            `idl:"name:dwFlags" json:"flags"`
	Stat         *Stat             `idl:"name:pStat" json:"stat"`
	ETableCount  uint32            `idl:"name:dwETableCount" json:"e_table_count"`
	ETable       []uint32          `idl:"name:lpETable;size_is:(dwETableCount);pointer:unique" json:"e_table"`
	Count        uint32            `idl:"name:Count" json:"count"`
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	Rows         *PropertyRowSet   `idl:"name:ppRows" json:"rows"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryRowsOperation) OpNum() int { return 3 }

func (o *xxx_QueryRowsOperation) OpName() string { return "/nspi/v56/NspiQueryRows" }

func (o *xxx_QueryRowsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ETable != nil && o.ETableCount == 0 {
		o.ETableCount = uint32(len(o.ETable))
	}
	if o.ETableCount > uint32(100000) {
		return fmt.Errorf("ETableCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRowsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwETableCount {in} (1:{range=(0,100000), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ETableCount); err != nil {
			return err
		}
	}
	// lpETable {in} (1:{pointer=unique}*(1))(2:{alias=DWORD}[dim:0,size_is=dwETableCount](uint32))
	{
		if o.ETable != nil || o.ETableCount > 0 {
			_ptr_lpETable := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ETableCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ETable {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ETable[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ETable); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ETable, _ptr_lpETable); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Count {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRowsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwETableCount {in} (1:{range=(0,100000), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ETableCount); err != nil {
			return err
		}
	}
	// lpETable {in} (1:{pointer=unique}*(1))(2:{alias=DWORD}[dim:0,size_is=dwETableCount](uint32))
	{
		_ptr_lpETable := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ETable", sizeInfo[0])
			}
			o.ETable = make([]uint32, sizeInfo[0])
			for i1 := range o.ETable {
				i1 := i1
				if err := w.ReadData(&o.ETable[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpETable := func(ptr interface{}) { o.ETable = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.ETable, _s_lpETable, _ptr_lpETable); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Count {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRowsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRowsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		if o.Rows != nil {
			_ptr_ppRows := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rows != nil {
					if err := o.Rows.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRowSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rows, _ptr_ppRows); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryRowsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		_ptr_ppRows := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rows == nil {
				o.Rows = &PropertyRowSet{}
			}
			if err := o.Rows.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRows := func(ptr interface{}) { o.Rows = *ptr.(**PropertyRowSet) }
		if err := w.ReadPointer(&o.Rows, _s_ppRows, _ptr_ppRows); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryRowsRequest structure represents the NspiQueryRows operation request
type QueryRowsRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags:  A DWORD value, containing a set of bit flags. The server MUST ignore values
	// other than the bit flags fEphID and fSkipObjects.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to specify both input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// dwETableCount: A DWORD value containing the number values in the input parameter
	// lpETable. This value is limited to 100,000.
	ETableCount uint32 `idl:"name:dwETableCount" json:"e_table_count"`
	// lpETable:  An array of DWORD values, representing an Explicit Table (see Explicit
	// Tables (section 3.1.1.3.2)).
	ETable []uint32 `idl:"name:lpETable;size_is:(dwETableCount);pointer:unique" json:"e_table"`
	// Count: A DWORD value containing the number of rows the client is requesting.
	Count uint32 `idl:"name:Count" json:"count"`
	// pPropTags: The value NULL or a reference to a PropertyTagArray_r value, containing
	// a list of the proptags of the properties that client requires to be returned for
	// each row returned.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
}

func (o *QueryRowsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryRowsOperation) *xxx_QueryRowsOperation {
	if op == nil {
		op = &xxx_QueryRowsOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	op.Stat = o.Stat
	op.ETableCount = o.ETableCount
	op.ETable = o.ETable
	op.Count = o.Count
	op.PropertyTags = o.PropertyTags
	return op
}

func (o *QueryRowsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryRowsOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
	o.Stat = op.Stat
	o.ETableCount = op.ETableCount
	o.ETable = op.ETable
	o.Count = op.Count
	o.PropertyTags = op.PropertyTags
}
func (o *QueryRowsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryRowsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryRowsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryRowsResponse structure represents the NspiQueryRows operation response
type QueryRowsResponse struct {
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to specify both input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// ppRows: A reference to a PropertyRowSet_r value. Contains the address book container
	// rows the server returns in response to the request.
	Rows *PropertyRowSet `idl:"name:ppRows" json:"rows"`
	// Return: The NspiQueryRows return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryRowsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryRowsOperation) *xxx_QueryRowsOperation {
	if op == nil {
		op = &xxx_QueryRowsOperation{}
	}
	if o == nil {
		return op
	}
	op.Stat = o.Stat
	op.Rows = o.Rows
	op.Return = o.Return
	return op
}

func (o *QueryRowsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryRowsOperation) {
	if o == nil {
		return
	}
	o.Stat = op.Stat
	o.Rows = op.Rows
	o.Return = op.Return
}
func (o *QueryRowsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryRowsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryRowsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SeekEntriesOperation structure represents the NspiSeekEntries operation
type xxx_SeekEntriesOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	_            uint32            `idl:"name:Reserved"`
	Stat         *Stat             `idl:"name:pStat" json:"stat"`
	Target       *PropertyValue    `idl:"name:pTarget" json:"target"`
	ETable       *PropertyTagArray `idl:"name:lpETable;pointer:unique" json:"e_table"`
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	Rows         *PropertyRowSet   `idl:"name:ppRows" json:"rows"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_SeekEntriesOperation) OpNum() int { return 4 }

func (o *xxx_SeekEntriesOperation) OpName() string { return "/nspi/v56/NspiSeekEntries" }

func (o *xxx_SeekEntriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekEntriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pTarget {in} (1:{pointer=ref}*(1))(2:{alias=PropertyValue_r}(struct))
	{
		if o.Target != nil {
			if err := o.Target.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpETable {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.ETable != nil {
			_ptr_lpETable := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ETable != nil {
					if err := o.ETable.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ETable, _ptr_lpETable); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekEntriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pTarget {in} (1:{pointer=ref}*(1))(2:{alias=PropertyValue_r}(struct))
	{
		if o.Target == nil {
			o.Target = &PropertyValue{}
		}
		if err := o.Target.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpETable {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_lpETable := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ETable == nil {
				o.ETable = &PropertyTagArray{}
			}
			if err := o.ETable.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpETable := func(ptr interface{}) { o.ETable = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.ETable, _s_lpETable, _ptr_lpETable); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekEntriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekEntriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		if o.Rows != nil {
			_ptr_ppRows := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rows != nil {
					if err := o.Rows.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRowSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rows, _ptr_ppRows); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekEntriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		_ptr_ppRows := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rows == nil {
				o.Rows = &PropertyRowSet{}
			}
			if err := o.Rows.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRows := func(ptr interface{}) { o.Rows = *ptr.(**PropertyRowSet) }
		if err := w.ReadPointer(&o.Rows, _s_ppRows, _ptr_ppRows); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SeekEntriesRequest structure represents the NspiSeekEntries operation request
type SeekEntriesRequest struct {
	// hRpc:  An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to both specify input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// pTarget: A PropertyValue_r value holding the value that is being sought.
	Target *PropertyValue `idl:"name:pTarget" json:"target"`
	// lpETable: The value NULL or a PropertyTagArray_r value. It holds a list of Mids that
	// comprises a restricted address book container.
	ETable *PropertyTagArray `idl:"name:lpETable;pointer:unique" json:"e_table"`
	// pPropTags: The value NULL or a reference to a PropertyTagArray_r value. Contains
	// list of the proptags of the columns that client wants to be returned for each row
	// returned.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
}

func (o *SeekEntriesRequest) xxx_ToOp(ctx context.Context, op *xxx_SeekEntriesOperation) *xxx_SeekEntriesOperation {
	if op == nil {
		op = &xxx_SeekEntriesOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.Target = o.Target
	op.ETable = o.ETable
	op.PropertyTags = o.PropertyTags
	return op
}

func (o *SeekEntriesRequest) xxx_FromOp(ctx context.Context, op *xxx_SeekEntriesOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.Target = op.Target
	o.ETable = op.ETable
	o.PropertyTags = op.PropertyTags
}
func (o *SeekEntriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SeekEntriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SeekEntriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SeekEntriesResponse structure represents the NspiSeekEntries operation response
type SeekEntriesResponse struct {
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to both specify input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// ppRows: A reference to a PropertyRowSet_r value. Contains the address book container
	// rows the server returns in response to the request.
	Rows *PropertyRowSet `idl:"name:ppRows" json:"rows"`
	// Return: The NspiSeekEntries return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SeekEntriesResponse) xxx_ToOp(ctx context.Context, op *xxx_SeekEntriesOperation) *xxx_SeekEntriesOperation {
	if op == nil {
		op = &xxx_SeekEntriesOperation{}
	}
	if o == nil {
		return op
	}
	op.Stat = o.Stat
	op.Rows = o.Rows
	op.Return = o.Return
	return op
}

func (o *SeekEntriesResponse) xxx_FromOp(ctx context.Context, op *xxx_SeekEntriesOperation) {
	if o == nil {
		return
	}
	o.Stat = op.Stat
	o.Rows = op.Rows
	o.Return = op.Return
}
func (o *SeekEntriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SeekEntriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SeekEntriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMatchesOperation structure represents the NspiGetMatches operation
type xxx_GetMatchesOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	_            uint32            `idl:"name:Reserved1"`
	Stat         *Stat             `idl:"name:pStat" json:"stat"`
	_            *PropertyTagArray `idl:"name:pReserved;pointer:unique"`
	_            uint32            `idl:"name:Reserved2"`
	Filter       *Restriction      `idl:"name:Filter;pointer:unique" json:"filter"`
	PropertyName *PropertyName     `idl:"name:lpPropName;pointer:unique" json:"property_name"`
	Requested    uint32            `idl:"name:ulRequested" json:"requested"`
	OutMIDs      *PropertyTagArray `idl:"name:ppOutMIds" json:"out_m_ids"`
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	Rows         *PropertyRowSet   `idl:"name:ppRows" json:"rows"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMatchesOperation) OpNum() int { return 5 }

func (o *xxx_GetMatchesOperation) OpName() string { return "/nspi/v56/NspiGetMatches" }

func (o *xxx_GetMatchesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMatchesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved1 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved1
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pReserved {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		// reserved pReserved
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Reserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved2
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// Filter {in} (1:{pointer=unique}*(1))(2:{alias=Restriction_r}(struct))
	{
		if o.Filter != nil {
			_ptr_Filter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Filter != nil {
					if err := o.Filter.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Restriction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Filter, _ptr_Filter); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpPropName {in} (1:{pointer=unique}*(1))(2:{alias=PropertyName_r}(struct))
	{
		if o.PropertyName != nil {
			_ptr_lpPropName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyName{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_lpPropName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ulRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Requested); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMatchesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved1 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved1
		var _Reserved1 uint32
		if err := w.ReadData(&_Reserved1); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pReserved {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		// reserved pReserved
		var _pReserved *PropertyTagArray
		_ptr_pReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _pReserved == nil {
				_pReserved = &PropertyTagArray{}
			}
			if err := _pReserved.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pReserved := func(ptr interface{}) { _pReserved = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&_pReserved, _s_pReserved, _ptr_pReserved); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Reserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved2
		var _Reserved2 uint32
		if err := w.ReadData(&_Reserved2); err != nil {
			return err
		}
	}
	// Filter {in} (1:{pointer=unique}*(1))(2:{alias=Restriction_r}(struct))
	{
		_ptr_Filter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Filter == nil {
				o.Filter = &Restriction{}
			}
			if err := o.Filter.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Filter := func(ptr interface{}) { o.Filter = *ptr.(**Restriction) }
		if err := w.ReadPointer(&o.Filter, _s_Filter, _ptr_Filter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpPropName {in} (1:{pointer=unique}*(1))(2:{alias=PropertyName_r}(struct))
	{
		_ptr_lpPropName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &PropertyName{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpPropName := func(ptr interface{}) { o.PropertyName = *ptr.(**PropertyName) }
		if err := w.ReadPointer(&o.PropertyName, _s_lpPropName, _ptr_lpPropName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ulRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Requested); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMatchesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMatchesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppOutMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.OutMIDs != nil {
			_ptr_ppOutMIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutMIDs != nil {
					if err := o.OutMIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutMIDs, _ptr_ppOutMIds); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		if o.Rows != nil {
			_ptr_ppRows := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rows != nil {
					if err := o.Rows.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRowSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rows, _ptr_ppRows); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMatchesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppOutMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppOutMIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OutMIDs == nil {
				o.OutMIDs = &PropertyTagArray{}
			}
			if err := o.OutMIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutMIds := func(ptr interface{}) { o.OutMIDs = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.OutMIDs, _s_ppOutMIds, _ptr_ppOutMIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		_ptr_ppRows := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rows == nil {
				o.Rows = &PropertyRowSet{}
			}
			if err := o.Rows.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRows := func(ptr interface{}) { o.Rows = *ptr.(**PropertyRowSet) }
		if err := w.ReadPointer(&o.Rows, _s_ppRows, _ptr_ppRows); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetMatchesRequest structure represents the NspiGetMatches operation request
type GetMatchesRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat: A reference to a STAT block describing a logical position in a specific address
	// book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// Filter: The value NULL or a Restriction_r value. Holds a logical restriction to apply
	// to the rows in the address book container specified in the pStat parameter.
	Filter *Restriction `idl:"name:Filter;pointer:unique" json:"filter"`
	// lpPropName: The value NULL or a PropertyName_r value. Holds the property to be opened
	// as a restricted address book container.
	PropertyName *PropertyName `idl:"name:lpPropName;pointer:unique" json:"property_name"`
	// ulRequested: A DWORD value. Contains the maximum number of rows to return in a restricted
	// address book container.
	Requested uint32 `idl:"name:ulRequested" json:"requested"`
	// pPropTags:  The value NULL or a reference to a PropertyTagArray_r value. Contains
	// list of the proptags of the columns that client wants to be returned for each row
	// returned.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
}

func (o *GetMatchesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMatchesOperation) *xxx_GetMatchesOperation {
	if op == nil {
		op = &xxx_GetMatchesOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.Filter = o.Filter
	op.PropertyName = o.PropertyName
	op.Requested = o.Requested
	op.PropertyTags = o.PropertyTags
	return op
}

func (o *GetMatchesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMatchesOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.Filter = op.Filter
	o.PropertyName = op.PropertyName
	o.Requested = op.Requested
	o.PropertyTags = op.PropertyTags
}
func (o *GetMatchesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMatchesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMatchesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMatchesResponse structure represents the NspiGetMatches operation response
type GetMatchesResponse struct {
	// pStat: A reference to a STAT block describing a logical position in a specific address
	// book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// ppOutMIds:  A PropertyTagArray_r value. On return, it holds a list of MId that comprise
	// a restricted address book container.
	OutMIDs *PropertyTagArray `idl:"name:ppOutMIds" json:"out_m_ids"`
	// ppRows:  A reference to a PropertyRowSet_r value. Contains the address book container
	// rows the server returns in response to the request.
	Rows *PropertyRowSet `idl:"name:ppRows" json:"rows"`
	// Return: The NspiGetMatches return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMatchesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMatchesOperation) *xxx_GetMatchesOperation {
	if op == nil {
		op = &xxx_GetMatchesOperation{}
	}
	if o == nil {
		return op
	}
	op.Stat = o.Stat
	op.OutMIDs = o.OutMIDs
	op.Rows = o.Rows
	op.Return = o.Return
	return op
}

func (o *GetMatchesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMatchesOperation) {
	if o == nil {
		return
	}
	o.Stat = op.Stat
	o.OutMIDs = op.OutMIDs
	o.Rows = op.Rows
	o.Return = op.Return
}
func (o *GetMatchesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMatchesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMatchesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResortRestrictionOperation structure represents the NspiResortRestriction operation
type xxx_ResortRestrictionOperation struct {
	Handle  *Handle           `idl:"name:hRpc" json:"handle"`
	_       uint32            `idl:"name:Reserved"`
	Stat    *Stat             `idl:"name:pStat" json:"stat"`
	InMIDs  *PropertyTagArray `idl:"name:pInMIds" json:"in_m_ids"`
	OutMIDs *PropertyTagArray `idl:"name:ppOutMIds" json:"out_m_ids"`
	Return  int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ResortRestrictionOperation) OpNum() int { return 6 }

func (o *xxx_ResortRestrictionOperation) OpName() string { return "/nspi/v56/NspiResortRestriction" }

func (o *xxx_ResortRestrictionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResortRestrictionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pInMIds {in} (1:{pointer=ref}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.InMIDs != nil {
			if err := o.InMIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppOutMIds {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.OutMIDs != nil {
			_ptr_ppOutMIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutMIDs != nil {
					if err := o.OutMIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutMIDs, _ptr_ppOutMIds); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResortRestrictionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pInMIds {in} (1:{pointer=ref}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.InMIDs == nil {
			o.InMIDs = &PropertyTagArray{}
		}
		if err := o.InMIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppOutMIds {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppOutMIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OutMIDs == nil {
				o.OutMIDs = &PropertyTagArray{}
			}
			if err := o.OutMIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutMIds := func(ptr interface{}) { o.OutMIDs = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.OutMIDs, _s_ppOutMIds, _ptr_ppOutMIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResortRestrictionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResortRestrictionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppOutMIds {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.OutMIDs != nil {
			_ptr_ppOutMIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutMIDs != nil {
					if err := o.OutMIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutMIDs, _ptr_ppOutMIds); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResortRestrictionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pStat {in, out} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppOutMIds {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppOutMIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OutMIDs == nil {
				o.OutMIDs = &PropertyTagArray{}
			}
			if err := o.OutMIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutMIds := func(ptr interface{}) { o.OutMIDs = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.OutMIDs, _s_ppOutMIds, _ptr_ppOutMIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResortRestrictionRequest structure represents the NspiResortRestriction operation request
type ResortRestrictionRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat: A reference to a STAT block describing a logical position in a specific address
	// book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// pInMIds: A PropertyTagArray_r value. It holds a list of MIds that comprise a restricted
	// address book container.
	InMIDs *PropertyTagArray `idl:"name:pInMIds" json:"in_m_ids"`
	// ppOutMIds: A PropertyTagArray_r value. On return, it holds a list of MIds that comprise
	// a restricted address book container.
	OutMIDs *PropertyTagArray `idl:"name:ppOutMIds" json:"out_m_ids"`
}

func (o *ResortRestrictionRequest) xxx_ToOp(ctx context.Context, op *xxx_ResortRestrictionOperation) *xxx_ResortRestrictionOperation {
	if op == nil {
		op = &xxx_ResortRestrictionOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.InMIDs = o.InMIDs
	op.OutMIDs = o.OutMIDs
	return op
}

func (o *ResortRestrictionRequest) xxx_FromOp(ctx context.Context, op *xxx_ResortRestrictionOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.InMIDs = op.InMIDs
	o.OutMIDs = op.OutMIDs
}
func (o *ResortRestrictionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResortRestrictionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResortRestrictionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResortRestrictionResponse structure represents the NspiResortRestriction operation response
type ResortRestrictionResponse struct {
	// pStat: A reference to a STAT block describing a logical position in a specific address
	// book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// ppOutMIds: A PropertyTagArray_r value. On return, it holds a list of MIds that comprise
	// a restricted address book container.
	OutMIDs *PropertyTagArray `idl:"name:ppOutMIds" json:"out_m_ids"`
	// Return: The NspiResortRestriction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResortRestrictionResponse) xxx_ToOp(ctx context.Context, op *xxx_ResortRestrictionOperation) *xxx_ResortRestrictionOperation {
	if op == nil {
		op = &xxx_ResortRestrictionOperation{}
	}
	if o == nil {
		return op
	}
	op.Stat = o.Stat
	op.OutMIDs = o.OutMIDs
	op.Return = o.Return
	return op
}

func (o *ResortRestrictionResponse) xxx_FromOp(ctx context.Context, op *xxx_ResortRestrictionOperation) {
	if o == nil {
		return
	}
	o.Stat = op.Stat
	o.OutMIDs = op.OutMIDs
	o.Return = op.Return
}
func (o *ResortRestrictionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResortRestrictionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResortRestrictionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DNToMIDOperation structure represents the NspiDNToMId operation
type xxx_DNToMIDOperation struct {
	Handle  *Handle           `idl:"name:hRpc" json:"handle"`
	_       uint32            `idl:"name:Reserved"`
	Names   *CharStringsArray `idl:"name:pNames" json:"names"`
	OutMIDs *PropertyTagArray `idl:"name:ppOutMIds" json:"out_m_ids"`
	Return  int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_DNToMIDOperation) OpNum() int { return 7 }

func (o *xxx_DNToMIDOperation) OpName() string { return "/nspi/v56/NspiDNToMId" }

func (o *xxx_DNToMIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DNToMIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pNames {in} (1:{pointer=ref}*(1))(2:{alias=StringsArray_r}(struct))
	{
		if o.Names != nil {
			if err := o.Names.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CharStringsArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DNToMIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pNames {in} (1:{pointer=ref}*(1))(2:{alias=StringsArray_r}(struct))
	{
		if o.Names == nil {
			o.Names = &CharStringsArray{}
		}
		if err := o.Names.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DNToMIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DNToMIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppOutMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.OutMIDs != nil {
			_ptr_ppOutMIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutMIDs != nil {
					if err := o.OutMIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutMIDs, _ptr_ppOutMIds); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DNToMIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppOutMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppOutMIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OutMIDs == nil {
				o.OutMIDs = &PropertyTagArray{}
			}
			if err := o.OutMIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutMIds := func(ptr interface{}) { o.OutMIDs = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.OutMIDs, _s_ppOutMIds, _ptr_ppOutMIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DNToMIDRequest structure represents the NspiDNToMId operation request
type DNToMIDRequest struct {
	// hRpc:  An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pNames: A StringsArray_r value. It holds a list of strings containing DNs, according
	// to [MS-OXOABK].
	Names *CharStringsArray `idl:"name:pNames" json:"names"`
}

func (o *DNToMIDRequest) xxx_ToOp(ctx context.Context, op *xxx_DNToMIDOperation) *xxx_DNToMIDOperation {
	if op == nil {
		op = &xxx_DNToMIDOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Names = o.Names
	return op
}

func (o *DNToMIDRequest) xxx_FromOp(ctx context.Context, op *xxx_DNToMIDOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Names = op.Names
}
func (o *DNToMIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DNToMIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DNToMIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DNToMIDResponse structure represents the NspiDNToMId operation response
type DNToMIDResponse struct {
	// ppOutMIds: A PropertyTagArray_r value. On return, it holds a list of MIds.
	OutMIDs *PropertyTagArray `idl:"name:ppOutMIds" json:"out_m_ids"`
	// Return: The NspiDNToMId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DNToMIDResponse) xxx_ToOp(ctx context.Context, op *xxx_DNToMIDOperation) *xxx_DNToMIDOperation {
	if op == nil {
		op = &xxx_DNToMIDOperation{}
	}
	if o == nil {
		return op
	}
	op.OutMIDs = o.OutMIDs
	op.Return = o.Return
	return op
}

func (o *DNToMIDResponse) xxx_FromOp(ctx context.Context, op *xxx_DNToMIDOperation) {
	if o == nil {
		return
	}
	o.OutMIDs = op.OutMIDs
	o.Return = op.Return
}
func (o *DNToMIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DNToMIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DNToMIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertyListOperation structure represents the NspiGetPropList operation
type xxx_GetPropertyListOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	Flags        uint32            `idl:"name:dwFlags" json:"flags"`
	MID          uint32            `idl:"name:dwMId" json:"m_id"`
	CodePage     uint32            `idl:"name:CodePage" json:"code_page"`
	PropertyTags *PropertyTagArray `idl:"name:ppPropTags" json:"property_tags"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertyListOperation) OpNum() int { return 8 }

func (o *xxx_GetPropertyListOperation) OpName() string { return "/nspi/v56/NspiGetPropList" }

func (o *xxx_GetPropertyListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// dwMId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MID); err != nil {
			return err
		}
	}
	// CodePage {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CodePage); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// dwMId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MID); err != nil {
			return err
		}
	}
	// CodePage {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CodePage); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppPropTags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_ppPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_ppPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppPropTags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_ppPropTags, _ptr_ppPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPropertyListRequest structure represents the NspiGetPropList operation request
type GetPropertyListRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags: A DWORD value, containing a set of bit flags. The server MUST ignore values
	// other than the bit flag fSkipObjects.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwMId: A DWORD value, containing a MId.
	MID uint32 `idl:"name:dwMId" json:"m_id"`
	// CodePage: The codepage in which the client wishes the server to express string values
	// properties.
	CodePage uint32 `idl:"name:CodePage" json:"code_page"`
}

func (o *GetPropertyListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertyListOperation) *xxx_GetPropertyListOperation {
	if op == nil {
		op = &xxx_GetPropertyListOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	op.MID = o.MID
	op.CodePage = o.CodePage
	return op
}

func (o *GetPropertyListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyListOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
	o.MID = op.MID
	o.CodePage = op.CodePage
}
func (o *GetPropertyListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertyListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertyListResponse structure represents the NspiGetPropList operation response
type GetPropertyListResponse struct {
	// ppPropTags: A PropertyTagArray_r value. On return, it holds a list of properties.
	PropertyTags *PropertyTagArray `idl:"name:ppPropTags" json:"property_tags"`
	// Return: The NspiGetPropList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertyListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertyListOperation) *xxx_GetPropertyListOperation {
	if op == nil {
		op = &xxx_GetPropertyListOperation{}
	}
	if o == nil {
		return op
	}
	op.PropertyTags = o.PropertyTags
	op.Return = o.Return
	return op
}

func (o *GetPropertyListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyListOperation) {
	if o == nil {
		return
	}
	o.PropertyTags = op.PropertyTags
	o.Return = op.Return
}
func (o *GetPropertyListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertyListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesOperation structure represents the NspiGetProps operation
type xxx_GetPropertiesOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	Flags        uint32            `idl:"name:dwFlags" json:"flags"`
	Stat         *Stat             `idl:"name:pStat" json:"stat"`
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	Rows         *PropertyRow      `idl:"name:ppRows" json:"rows"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 9 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/nspi/v56/NspiGetProps" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRow_r}(struct))
	{
		if o.Rows != nil {
			_ptr_ppRows := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rows != nil {
					if err := o.Rows.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRow{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rows, _ptr_ppRows); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRow_r}(struct))
	{
		_ptr_ppRows := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rows == nil {
				o.Rows = &PropertyRow{}
			}
			if err := o.Rows.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRows := func(ptr interface{}) { o.Rows = *ptr.(**PropertyRow) }
		if err := w.ReadPointer(&o.Rows, _s_ppRows, _ptr_ppRows); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPropertiesRequest structure represents the NspiGetProps operation request
type GetPropertiesRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags: A DWORD value, containing a set of bit flags. The server MUST ignore values
	// other than the bit flags fEphID and fSkipObjects.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to both specify input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// pPropTags: The value NULL or a reference to a PropertyTagArray_r value. Contains
	// list of the proptags of the properties that the client wants to be returned.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	op.Stat = o.Stat
	op.PropertyTags = o.PropertyTags
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
	o.Stat = op.Stat
	o.PropertyTags = op.PropertyTags
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the NspiGetProps operation response
type GetPropertiesResponse struct {
	// ppRows: A reference to a PropertyRow_r value. Contains the address book container
	// row the server returns in response to the request.
	Rows *PropertyRow `idl:"name:ppRows" json:"rows"`
	// Return: The NspiGetProps return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Rows = o.Rows
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.Rows = op.Rows
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CompareMIDsOperation structure represents the NspiCompareMIds operation
type xxx_CompareMIDsOperation struct {
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	_      uint32  `idl:"name:Reserved"`
	Stat   *Stat   `idl:"name:pStat" json:"stat"`
	MID1   uint32  `idl:"name:MId1" json:"m_id1"`
	MID2   uint32  `idl:"name:MId2" json:"m_id2"`
	Result int32   `idl:"name:plResult" json:"result"`
	Return int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CompareMIDsOperation) OpNum() int { return 10 }

func (o *xxx_CompareMIDsOperation) OpName() string { return "/nspi/v56/NspiCompareMIds" }

func (o *xxx_CompareMIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CompareMIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MId1 {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MID1); err != nil {
			return err
		}
	}
	// MId2 {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MID2); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CompareMIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MId1 {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MID1); err != nil {
			return err
		}
	}
	// MId2 {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MID2); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CompareMIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CompareMIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// plResult {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CompareMIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// plResult {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CompareMIDsRequest structure represents the NspiCompareMIds operation request
type CompareMIDsRequest struct {
	// hRpc:  An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat: pStat: A reference to a STAT block describing a logical position in a specific
	// address book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// MId1: A DWORD value containing a MId.
	MID1 uint32 `idl:"name:MId1" json:"m_id1"`
	// MId2: A DWORD value containing a MId.
	MID2 uint32 `idl:"name:MId2" json:"m_id2"`
}

func (o *CompareMIDsRequest) xxx_ToOp(ctx context.Context, op *xxx_CompareMIDsOperation) *xxx_CompareMIDsOperation {
	if op == nil {
		op = &xxx_CompareMIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.MID1 = o.MID1
	op.MID2 = o.MID2
	return op
}

func (o *CompareMIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_CompareMIDsOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.MID1 = op.MID1
	o.MID2 = op.MID2
}
func (o *CompareMIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CompareMIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CompareMIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CompareMIDsResponse structure represents the NspiCompareMIds operation response
type CompareMIDsResponse struct {
	// plResult: A DWORD value. On return, it contains the result of the comparison.
	Result int32 `idl:"name:plResult" json:"result"`
	// Return: The NspiCompareMIds return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CompareMIDsResponse) xxx_ToOp(ctx context.Context, op *xxx_CompareMIDsOperation) *xxx_CompareMIDsOperation {
	if op == nil {
		op = &xxx_CompareMIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *CompareMIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_CompareMIDsOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *CompareMIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CompareMIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CompareMIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyPropertiesOperation structure represents the NspiModProps operation
type xxx_ModifyPropertiesOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	_            uint32            `idl:"name:Reserved"`
	Stat         *Stat             `idl:"name:pStat" json:"stat"`
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	Row          *PropertyRow      `idl:"name:pRow" json:"row"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyPropertiesOperation) OpNum() int { return 11 }

func (o *xxx_ModifyPropertiesOperation) OpName() string { return "/nspi/v56/NspiModProps" }

func (o *xxx_ModifyPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pRow {in} (1:{pointer=ref}*(1))(2:{alias=PropertyRow_r}(struct))
	{
		if o.Row != nil {
			if err := o.Row.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyRow{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pRow {in} (1:{pointer=ref}*(1))(2:{alias=PropertyRow_r}(struct))
	{
		if o.Row == nil {
			o.Row = &PropertyRow{}
		}
		if err := o.Row.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ModifyPropertiesRequest structure represents the NspiModProps operation request
type ModifyPropertiesRequest struct {
	// hRpc:  An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat:  A reference to a STAT block describing a logical position in a specific
	// address book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// pPropTags: The value NULL or a reference to a PropertyTagArray_r. Contains list of
	// the proptags of the columns that client requests all values to be removed from.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	// pRow: A PropertyRow_r value. Contains an address book row.
	Row *PropertyRow `idl:"name:pRow" json:"row"`
}

func (o *ModifyPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyPropertiesOperation) *xxx_ModifyPropertiesOperation {
	if op == nil {
		op = &xxx_ModifyPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.PropertyTags = o.PropertyTags
	op.Row = o.Row
	return op
}

func (o *ModifyPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyPropertiesOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.PropertyTags = op.PropertyTags
	o.Row = op.Row
}
func (o *ModifyPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyPropertiesResponse structure represents the NspiModProps operation response
type ModifyPropertiesResponse struct {
	// Return: The NspiModProps return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyPropertiesOperation) *xxx_ModifyPropertiesOperation {
	if op == nil {
		op = &xxx_ModifyPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ModifyPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyPropertiesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ModifyPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSpecialTableOperation structure represents the NspiGetSpecialTable operation
type xxx_GetSpecialTableOperation struct {
	Handle  *Handle         `idl:"name:hRpc" json:"handle"`
	Flags   uint32          `idl:"name:dwFlags" json:"flags"`
	Stat    *Stat           `idl:"name:pStat" json:"stat"`
	Version uint32          `idl:"name:lpVersion" json:"version"`
	Rows    *PropertyRowSet `idl:"name:ppRows" json:"rows"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSpecialTableOperation) OpNum() int { return 12 }

func (o *xxx_GetSpecialTableOperation) OpName() string { return "/nspi/v56/NspiGetSpecialTable" }

func (o *xxx_GetSpecialTableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSpecialTableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpVersion {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSpecialTableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpVersion {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSpecialTableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSpecialTableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpVersion {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		if o.Rows != nil {
			_ptr_ppRows := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rows != nil {
					if err := o.Rows.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRowSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rows, _ptr_ppRows); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSpecialTableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpVersion {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		_ptr_ppRows := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rows == nil {
				o.Rows = &PropertyRowSet{}
			}
			if err := o.Rows.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRows := func(ptr interface{}) { o.Rows = *ptr.(**PropertyRowSet) }
		if err := w.ReadPointer(&o.Rows, _s_ppRows, _ptr_ppRows); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSpecialTableRequest structure represents the NspiGetSpecialTable operation request
type GetSpecialTableRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags:  A DWORD value containing a set of bit flags. The server MUST ignore values
	// other than the bit flags NspiAddressCreationTemplates and NspiUnicodeStrings.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pStat: A pointer to a STAT block describing a logical position in a specific address
	// book container. This parameter is used to both specify input parameters from the
	// client and return values from the NSPI server.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// lpVersion: A reference to a DWORD. On input, holds the value of the version number
	// of the address book hierarchy table that the client has.
	Version uint32 `idl:"name:lpVersion" json:"version"`
}

func (o *GetSpecialTableRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSpecialTableOperation) *xxx_GetSpecialTableOperation {
	if op == nil {
		op = &xxx_GetSpecialTableOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	op.Stat = o.Stat
	op.Version = o.Version
	return op
}

func (o *GetSpecialTableRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSpecialTableOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
	o.Stat = op.Stat
	o.Version = op.Version
}
func (o *GetSpecialTableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSpecialTableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSpecialTableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSpecialTableResponse structure represents the NspiGetSpecialTable operation response
type GetSpecialTableResponse struct {
	// lpVersion: A reference to a DWORD. On input, holds the value of the version number
	// of the address book hierarchy table that the client has.
	Version uint32 `idl:"name:lpVersion" json:"version"`
	// ppRows:  An PropertyRowSet_r. On return, holds the rows for the table that the client
	// is requesting.
	Rows *PropertyRowSet `idl:"name:ppRows" json:"rows"`
	// Return: The NspiGetSpecialTable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSpecialTableResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSpecialTableOperation) *xxx_GetSpecialTableOperation {
	if op == nil {
		op = &xxx_GetSpecialTableOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	op.Rows = o.Rows
	op.Return = o.Return
	return op
}

func (o *GetSpecialTableResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSpecialTableOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.Rows = op.Rows
	o.Return = op.Return
}
func (o *GetSpecialTableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSpecialTableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSpecialTableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTemplateInfoOperation structure represents the NspiGetTemplateInfo operation
type xxx_GetTemplateInfoOperation struct {
	Handle   *Handle      `idl:"name:hRpc" json:"handle"`
	Flags    uint32       `idl:"name:dwFlags" json:"flags"`
	Type     uint32       `idl:"name:ulType" json:"type"`
	DN       string       `idl:"name:pDN;string;pointer:unique" json:"dn"`
	CodePage uint32       `idl:"name:dwCodePage" json:"code_page"`
	LocaleID uint32       `idl:"name:dwLocaleID" json:"locale_id"`
	Data     *PropertyRow `idl:"name:ppData" json:"data"`
	Return   int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTemplateInfoOperation) OpNum() int { return 13 }

func (o *xxx_GetTemplateInfoOperation) OpName() string { return "/nspi/v56/NspiGetTemplateInfo" }

func (o *xxx_GetTemplateInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTemplateInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// ulType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// pDN {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](char))
	{
		if o.DN != "" {
			_ptr_pDN := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.DN); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DN, _ptr_pDN); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwCodePage {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CodePage); err != nil {
			return err
		}
	}
	// dwLocaleID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTemplateInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// ulType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// pDN {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](char))
	{
		_ptr_pDN := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.DN); err != nil {
				return err
			}
			return nil
		})
		_s_pDN := func(ptr interface{}) { o.DN = *ptr.(*string) }
		if err := w.ReadPointer(&o.DN, _s_pDN, _ptr_pDN); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwCodePage {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CodePage); err != nil {
			return err
		}
	}
	// dwLocaleID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTemplateInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTemplateInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRow_r}(struct))
	{
		if o.Data != nil {
			_ptr_ppData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Data != nil {
					if err := o.Data.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRow{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_ppData); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTemplateInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRow_r}(struct))
	{
		_ptr_ppData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Data == nil {
				o.Data = &PropertyRow{}
			}
			if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppData := func(ptr interface{}) { o.Data = *ptr.(**PropertyRow) }
		if err := w.ReadPointer(&o.Data, _s_ppData, _ptr_ppData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTemplateInfoRequest structure represents the NspiGetTemplateInfo operation request
type GetTemplateInfoRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags: A DWORD value containing a set of bit flags. The server MUST ignore values
	// other than the bit flags TI_HELPFILE_NAME, TI_HELPFILE_CONTENTS, TI_SCRIPT, TI_TEMPLATE,
	// and TI_EMT.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// ulType: A DWORD value. Specifies the display type of the template for which information
	// is requested.
	Type uint32 `idl:"name:ulType" json:"type"`
	// pDN: The value NULL or the DN of the template requested. The value is NULL-terminated.
	DN string `idl:"name:pDN;string;pointer:unique" json:"dn"`
	// dwCodePage: A DWORD value. Specifies the codepage of the template for which information
	// is requested.
	CodePage uint32 `idl:"name:dwCodePage" json:"code_page"`
	// dwLocaleID: A DWORD value. Specifies the LCID of the template for which information
	// is requested.
	LocaleID uint32 `idl:"name:dwLocaleID" json:"locale_id"`
}

func (o *GetTemplateInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTemplateInfoOperation) *xxx_GetTemplateInfoOperation {
	if op == nil {
		op = &xxx_GetTemplateInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	op.Type = o.Type
	op.DN = o.DN
	op.CodePage = o.CodePage
	op.LocaleID = o.LocaleID
	return op
}

func (o *GetTemplateInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTemplateInfoOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
	o.Type = op.Type
	o.DN = op.DN
	o.CodePage = op.CodePage
	o.LocaleID = op.LocaleID
}
func (o *GetTemplateInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTemplateInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTemplateInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTemplateInfoResponse structure represents the NspiGetTemplateInfo operation response
type GetTemplateInfoResponse struct {
	// ppData: A reference to a PropertyRow_r value. On return, it contains the information
	// requested.
	Data *PropertyRow `idl:"name:ppData" json:"data"`
	// Return: The NspiGetTemplateInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTemplateInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTemplateInfoOperation) *xxx_GetTemplateInfoOperation {
	if op == nil {
		op = &xxx_GetTemplateInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Data = o.Data
	op.Return = o.Return
	return op
}

func (o *GetTemplateInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTemplateInfoOperation) {
	if o == nil {
		return
	}
	o.Data = op.Data
	o.Return = op.Return
}
func (o *GetTemplateInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTemplateInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTemplateInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyLinkAttributeOperation structure represents the NspiModLinkAtt operation
type xxx_ModifyLinkAttributeOperation struct {
	Handle      *Handle      `idl:"name:hRpc" json:"handle"`
	Flags       uint32       `idl:"name:dwFlags" json:"flags"`
	PropertyTag uint32       `idl:"name:ulPropTag" json:"property_tag"`
	MID         uint32       `idl:"name:dwMId" json:"m_id"`
	EntryIDs    *BinaryArray `idl:"name:lpEntryIds" json:"entry_ids"`
	Return      int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyLinkAttributeOperation) OpNum() int { return 14 }

func (o *xxx_ModifyLinkAttributeOperation) OpName() string { return "/nspi/v56/NspiModLinkAtt" }

func (o *xxx_ModifyLinkAttributeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyLinkAttributeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// ulPropTag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PropertyTag); err != nil {
			return err
		}
	}
	// dwMId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MID); err != nil {
			return err
		}
	}
	// lpEntryIds {in} (1:{pointer=ref}*(1))(2:{alias=BinaryArray_r}(struct))
	{
		if o.EntryIDs != nil {
			if err := o.EntryIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&BinaryArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyLinkAttributeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// ulPropTag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PropertyTag); err != nil {
			return err
		}
	}
	// dwMId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MID); err != nil {
			return err
		}
	}
	// lpEntryIds {in} (1:{pointer=ref}*(1))(2:{alias=BinaryArray_r}(struct))
	{
		if o.EntryIDs == nil {
			o.EntryIDs = &BinaryArray{}
		}
		if err := o.EntryIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyLinkAttributeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyLinkAttributeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyLinkAttributeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ModifyLinkAttributeRequest structure represents the NspiModLinkAtt operation request
type ModifyLinkAttributeRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags: A DWORD value containing a set of bit flags. The server MUST ignore values
	// other than the bit flag fDelete.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// ulPropTag: A DWORD value. Contains the proptag of the property that the client wishes
	// to modify.
	PropertyTag uint32 `idl:"name:ulPropTag" json:"property_tag"`
	// dwMId:  A DWORD value containing the MId of the address book row that the client
	// wishes to modify.
	MID uint32 `idl:"name:dwMId" json:"m_id"`
	// lpEntryIds: A BinaryArray value. Contains a list of Entry IDs to be used to modify
	// the requested property on the requested address book row. These Entry IDs can be
	// either Ephemeral Entry IDs or Permanent Entry IDs or both.
	EntryIDs *BinaryArray `idl:"name:lpEntryIds" json:"entry_ids"`
}

func (o *ModifyLinkAttributeRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyLinkAttributeOperation) *xxx_ModifyLinkAttributeOperation {
	if op == nil {
		op = &xxx_ModifyLinkAttributeOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	op.PropertyTag = o.PropertyTag
	op.MID = o.MID
	op.EntryIDs = o.EntryIDs
	return op
}

func (o *ModifyLinkAttributeRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyLinkAttributeOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
	o.PropertyTag = op.PropertyTag
	o.MID = op.MID
	o.EntryIDs = op.EntryIDs
}
func (o *ModifyLinkAttributeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyLinkAttributeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyLinkAttributeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyLinkAttributeResponse structure represents the NspiModLinkAtt operation response
type ModifyLinkAttributeResponse struct {
	// Return: The NspiModLinkAtt return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyLinkAttributeResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyLinkAttributeOperation) *xxx_ModifyLinkAttributeOperation {
	if op == nil {
		op = &xxx_ModifyLinkAttributeOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ModifyLinkAttributeResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyLinkAttributeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ModifyLinkAttributeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyLinkAttributeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyLinkAttributeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryColumnsOperation structure represents the NspiQueryColumns operation
type xxx_QueryColumnsOperation struct {
	Handle  *Handle           `idl:"name:hRpc" json:"handle"`
	_       uint32            `idl:"name:Reserved"`
	Flags   uint32            `idl:"name:dwFlags" json:"flags"`
	Columns *PropertyTagArray `idl:"name:ppColumns" json:"columns"`
	Return  int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryColumnsOperation) OpNum() int { return 16 }

func (o *xxx_QueryColumnsOperation) OpName() string { return "/nspi/v56/NspiQueryColumns" }

func (o *xxx_QueryColumnsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryColumnsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryColumnsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryColumnsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryColumnsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppColumns {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.Columns != nil {
			_ptr_ppColumns := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Columns != nil {
					if err := o.Columns.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Columns, _ptr_ppColumns); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryColumnsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppColumns {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppColumns := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Columns == nil {
				o.Columns = &PropertyTagArray{}
			}
			if err := o.Columns.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppColumns := func(ptr interface{}) { o.Columns = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.Columns, _s_ppColumns, _ptr_ppColumns); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryColumnsRequest structure represents the NspiQueryColumns operation request
type QueryColumnsRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags: A DWORD value containing a set of bit flags. The server MUST ignore values
	// other than the bit flag NspiUnicodeProptypes.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
}

func (o *QueryColumnsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryColumnsOperation) *xxx_QueryColumnsOperation {
	if op == nil {
		op = &xxx_QueryColumnsOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	return op
}

func (o *QueryColumnsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryColumnsOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
}
func (o *QueryColumnsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryColumnsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryColumnsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryColumnsResponse structure represents the NspiQueryColumns operation response
type QueryColumnsResponse struct {
	// ppColumns: A reference to a PropertyTagArray_r structure. On return, contains a list
	// of proptags.
	Columns *PropertyTagArray `idl:"name:ppColumns" json:"columns"`
	// Return: The NspiQueryColumns return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryColumnsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryColumnsOperation) *xxx_QueryColumnsOperation {
	if op == nil {
		op = &xxx_QueryColumnsOperation{}
	}
	if o == nil {
		return op
	}
	op.Columns = o.Columns
	op.Return = o.Return
	return op
}

func (o *QueryColumnsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryColumnsOperation) {
	if o == nil {
		return
	}
	o.Columns = op.Columns
	o.Return = op.Return
}
func (o *QueryColumnsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryColumnsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryColumnsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNamesFromIDsOperation structure represents the NspiGetNamesFromIDs operation
type xxx_GetNamesFromIDsOperation struct {
	Handle               *Handle           `idl:"name:hRpc" json:"handle"`
	_                    uint32            `idl:"name:Reserved"`
	GUID                 *FlatUID          `idl:"name:lpguid;pointer:unique" json:"guid"`
	PropertyTags         *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	ReturnedPropertyTags *PropertyTagArray `idl:"name:ppReturnedPropTags" json:"returned_property_tags"`
	Names                *PropertyNameSet  `idl:"name:ppNames" json:"names"`
	Return               int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNamesFromIDsOperation) OpNum() int { return 17 }

func (o *xxx_GetNamesFromIDsOperation) OpName() string { return "/nspi/v56/NspiGetNamesFromIDs" }

func (o *xxx_GetNamesFromIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesFromIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// lpguid {in} (1:{pointer=unique}*(1))(2:{alias=FlatUID_r}(struct))
	{
		if o.GUID != nil {
			_ptr_lpguid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.GUID != nil {
					if err := o.GUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&FlatUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_lpguid); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesFromIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// lpguid {in} (1:{pointer=unique}*(1))(2:{alias=FlatUID_r}(struct))
	{
		_ptr_lpguid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.GUID == nil {
				o.GUID = &FlatUID{}
			}
			if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpguid := func(ptr interface{}) { o.GUID = *ptr.(**FlatUID) }
		if err := w.ReadPointer(&o.GUID, _s_lpguid, _ptr_lpguid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesFromIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesFromIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppReturnedPropTags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.ReturnedPropertyTags != nil {
			_ptr_ppReturnedPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnedPropertyTags != nil {
					if err := o.ReturnedPropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnedPropertyTags, _ptr_ppReturnedPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppNames {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyNameSet_r}(struct))
	{
		if o.Names != nil {
			_ptr_ppNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Names != nil {
					if err := o.Names.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyNameSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Names, _ptr_ppNames); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesFromIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppReturnedPropTags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppReturnedPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnedPropertyTags == nil {
				o.ReturnedPropertyTags = &PropertyTagArray{}
			}
			if err := o.ReturnedPropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppReturnedPropTags := func(ptr interface{}) { o.ReturnedPropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.ReturnedPropertyTags, _s_ppReturnedPropTags, _ptr_ppReturnedPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppNames {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyNameSet_r}(struct))
	{
		_ptr_ppNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Names == nil {
				o.Names = &PropertyNameSet{}
			}
			if err := o.Names.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppNames := func(ptr interface{}) { o.Names = *ptr.(**PropertyNameSet) }
		if err := w.ReadPointer(&o.Names, _s_ppNames, _ptr_ppNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNamesFromIDsRequest structure represents the NspiGetNamesFromIDs operation request
type GetNamesFromIDsRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// lpguid: The value NULL or a FlatUID_r value. Specifies the property set about which
	// the client is requesting information.
	GUID *FlatUID `idl:"name:lpguid;pointer:unique" json:"guid"`
	// pPropTags: The value NULL or a PropertyTagArray_r value. Specifies the specific Property
	// IDs about which the client is requesting information.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
}

func (o *GetNamesFromIDsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNamesFromIDsOperation) *xxx_GetNamesFromIDsOperation {
	if op == nil {
		op = &xxx_GetNamesFromIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.GUID = o.GUID
	op.PropertyTags = o.PropertyTags
	return op
}

func (o *GetNamesFromIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNamesFromIDsOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.GUID = op.GUID
	o.PropertyTags = op.PropertyTags
}
func (o *GetNamesFromIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNamesFromIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamesFromIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNamesFromIDsResponse structure represents the NspiGetNamesFromIDs operation response
type GetNamesFromIDsResponse struct {
	// ppReturnedPropTags: Contains an SPropTagArray. On return, it contains a list of all
	// the proptags in the property set specified in the input parameter lpguid. If lpguid
	// is NULL, this value MUST be NULL.
	ReturnedPropertyTags *PropertyTagArray `idl:"name:ppReturnedPropTags" json:"returned_property_tags"`
	// ppNames:  A PropertyNameSet_r value. On return, it contains a list of property names
	// satisfying the client's request.
	Names *PropertyNameSet `idl:"name:ppNames" json:"names"`
	// Return: The NspiGetNamesFromIDs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNamesFromIDsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNamesFromIDsOperation) *xxx_GetNamesFromIDsOperation {
	if op == nil {
		op = &xxx_GetNamesFromIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.ReturnedPropertyTags = o.ReturnedPropertyTags
	op.Names = o.Names
	op.Return = o.Return
	return op
}

func (o *GetNamesFromIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNamesFromIDsOperation) {
	if o == nil {
		return
	}
	o.ReturnedPropertyTags = op.ReturnedPropertyTags
	o.Names = op.Names
	o.Return = op.Return
}
func (o *GetNamesFromIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNamesFromIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamesFromIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIDsFromNamesOperation structure represents the NspiGetIDsFromNames operation
type xxx_GetIDsFromNamesOperation struct {
	Handle             *Handle           `idl:"name:hRpc" json:"handle"`
	_                  uint32            `idl:"name:Reserved"`
	Flags              uint32            `idl:"name:dwFlags" json:"flags"`
	PropertyNamesCount uint32            `idl:"name:cPropNames" json:"property_names_count"`
	Names              []*PropertyName   `idl:"name:pNames;size_is:(cPropNames)" json:"names"`
	PropertyTags       *PropertyTagArray `idl:"name:ppPropTags" json:"property_tags"`
	Return             int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIDsFromNamesOperation) OpNum() int { return 18 }

func (o *xxx_GetIDsFromNamesOperation) OpName() string { return "/nspi/v56/NspiGetIDsFromNames" }

func (o *xxx_GetIDsFromNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Names != nil && o.PropertyNamesCount == 0 {
		o.PropertyNamesCount = uint32(len(o.Names))
	}
	if o.PropertyNamesCount > uint32(100000) {
		return fmt.Errorf("PropertyNamesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsFromNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// cPropNames {in} (1:{range=(0,100000), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PropertyNamesCount); err != nil {
			return err
		}
	}
	// pNames {in} (1:{pointer=ref}*(1)[dim:0,size_is=cPropNames]*(1))(2:{alias=PropertyName_r}(struct))
	{
		dimSize1 := uint64(o.PropertyNamesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Names {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Names[i1] != nil {
				_ptr_pNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Names[i1] != nil {
						if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&PropertyName{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Names[i1], _ptr_pNames); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsFromNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// cPropNames {in} (1:{range=(0,100000), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PropertyNamesCount); err != nil {
			return err
		}
	}
	// pNames {in} (1:{pointer=ref}*(1)[dim:0,size_is=cPropNames]*(1))(2:{alias=PropertyName_r}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*PropertyName, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			_ptr_pNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Names[i1] == nil {
					o.Names[i1] = &PropertyName{}
				}
				if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_pNames := func(ptr interface{}) { o.Names[i1] = *ptr.(**PropertyName) }
			if err := w.ReadPointer(&o.Names[i1], _s_pNames, _ptr_pNames); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsFromNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsFromNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppPropTags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_ppPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_ppPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsFromNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppPropTags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_ppPropTags, _ptr_ppPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetIDsFromNamesRequest structure represents the NspiGetIDsFromNames operation request
type GetIDsFromNamesRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// dwFlags: A DWORD value. All clients MUST set this value to either 0 or the flag NspiVerifyNames.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// cPropNames: A DWORD value containing the number property names supplied by the client.
	// The value is limited to 100,000.
	PropertyNamesCount uint32 `idl:"name:cPropNames" json:"property_names_count"`
	// pNames: A reference to a PropertyName_r value. Contains a list of property names
	// supplied by the client.
	Names []*PropertyName `idl:"name:pNames;size_is:(cPropNames)" json:"names"`
}

func (o *GetIDsFromNamesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIDsFromNamesOperation) *xxx_GetIDsFromNamesOperation {
	if op == nil {
		op = &xxx_GetIDsFromNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Flags = o.Flags
	op.PropertyNamesCount = o.PropertyNamesCount
	op.Names = o.Names
	return op
}

func (o *GetIDsFromNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIDsFromNamesOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Flags = op.Flags
	o.PropertyNamesCount = op.PropertyNamesCount
	o.Names = op.Names
}
func (o *GetIDsFromNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIDsFromNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIDsFromNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIDsFromNamesResponse structure represents the NspiGetIDsFromNames operation response
type GetIDsFromNamesResponse struct {
	// ppPropTags: A reference to a PropertyTagArray_r value. On successful return to the
	// client, contains a list of proptags.
	PropertyTags *PropertyTagArray `idl:"name:ppPropTags" json:"property_tags"`
	// Return: The NspiGetIDsFromNames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIDsFromNamesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIDsFromNamesOperation) *xxx_GetIDsFromNamesOperation {
	if op == nil {
		op = &xxx_GetIDsFromNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.PropertyTags = o.PropertyTags
	op.Return = o.Return
	return op
}

func (o *GetIDsFromNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIDsFromNamesOperation) {
	if o == nil {
		return
	}
	o.PropertyTags = op.PropertyTags
	o.Return = op.Return
}
func (o *GetIDsFromNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIDsFromNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIDsFromNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResolveNamesOperation structure represents the NspiResolveNames operation
type xxx_ResolveNamesOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	_            uint32            `idl:"name:Reserved"`
	Stat         *Stat             `idl:"name:pStat" json:"stat"`
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	Strings      *CharStringsArray `idl:"name:paStr" json:"strings"`
	MIDs         *PropertyTagArray `idl:"name:ppMIds" json:"m_ids"`
	Rows         *PropertyRowSet   `idl:"name:ppRows" json:"rows"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ResolveNamesOperation) OpNum() int { return 19 }

func (o *xxx_ResolveNamesOperation) OpName() string { return "/nspi/v56/NspiResolveNames" }

func (o *xxx_ResolveNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// paStr {in} (1:{pointer=ref}*(1))(2:{alias=StringsArray_r}(struct))
	{
		if o.Strings != nil {
			if err := o.Strings.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CharStringsArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// paStr {in} (1:{pointer=ref}*(1))(2:{alias=StringsArray_r}(struct))
	{
		if o.Strings == nil {
			o.Strings = &CharStringsArray{}
		}
		if err := o.Strings.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.MIDs != nil {
			_ptr_ppMIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MIDs != nil {
					if err := o.MIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MIDs, _ptr_ppMIds); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		if o.Rows != nil {
			_ptr_ppRows := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rows != nil {
					if err := o.Rows.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRowSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rows, _ptr_ppRows); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppMIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MIDs == nil {
				o.MIDs = &PropertyTagArray{}
			}
			if err := o.MIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppMIds := func(ptr interface{}) { o.MIDs = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.MIDs, _s_ppMIds, _ptr_ppMIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		_ptr_ppRows := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rows == nil {
				o.Rows = &PropertyRowSet{}
			}
			if err := o.Rows.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRows := func(ptr interface{}) { o.Rows = *ptr.(**PropertyRowSet) }
		if err := w.ReadPointer(&o.Rows, _s_ppRows, _ptr_ppRows); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResolveNamesRequest structure represents the NspiResolveNames operation request
type ResolveNamesRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat: A reference to a STAT block describing a logical position in a specific address
	// book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// pPropTags: The value NULL or a reference to a PropertyTagArray_r value containing
	// a list of the proptags of the columns that the client requests to be returned for
	// each row returned.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	// paStr: A StringsArray_r value. Specifies the values the client is requesting the
	// server to do ANR on. The server MUST apply any necessary character set conversion
	// as specified in String Handling (section 3.1.1.2).
	Strings *CharStringsArray `idl:"name:paStr" json:"strings"`
}

func (o *ResolveNamesRequest) xxx_ToOp(ctx context.Context, op *xxx_ResolveNamesOperation) *xxx_ResolveNamesOperation {
	if op == nil {
		op = &xxx_ResolveNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.PropertyTags = o.PropertyTags
	op.Strings = o.Strings
	return op
}

func (o *ResolveNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_ResolveNamesOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.PropertyTags = op.PropertyTags
	o.Strings = op.Strings
}
func (o *ResolveNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResolveNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResolveNamesResponse structure represents the NspiResolveNames operation response
type ResolveNamesResponse struct {
	// ppMIds: A PropertyTagArray_r value. On return, contains a list of MIds matching the
	// array of strings, as specified in the input parameter paStr.
	MIDs *PropertyTagArray `idl:"name:ppMIds" json:"m_ids"`
	// ppRows: A reference to a PropertyRowSet_r value. Contains the address book container
	// rows that the server returns in response to the request.
	Rows *PropertyRowSet `idl:"name:ppRows" json:"rows"`
	// Return: The NspiResolveNames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResolveNamesResponse) xxx_ToOp(ctx context.Context, op *xxx_ResolveNamesOperation) *xxx_ResolveNamesOperation {
	if op == nil {
		op = &xxx_ResolveNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.MIDs = o.MIDs
	op.Rows = o.Rows
	op.Return = o.Return
	return op
}

func (o *ResolveNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_ResolveNamesOperation) {
	if o == nil {
		return
	}
	o.MIDs = op.MIDs
	o.Rows = op.Rows
	o.Return = op.Return
}
func (o *ResolveNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResolveNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResolveNamesWOperation structure represents the NspiResolveNamesW operation
type xxx_ResolveNamesWOperation struct {
	Handle       *Handle           `idl:"name:hRpc" json:"handle"`
	_            uint32            `idl:"name:Reserved"`
	Stat         *Stat             `idl:"name:pStat" json:"stat"`
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	Strings      *StringsArray     `idl:"name:paWStr" json:"strings"`
	MIDs         *PropertyTagArray `idl:"name:ppMIds" json:"m_ids"`
	Rows         *PropertyRowSet   `idl:"name:ppRows" json:"rows"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ResolveNamesWOperation) OpNum() int { return 20 }

func (o *xxx_ResolveNamesWOperation) OpName() string { return "/nspi/v56/NspiResolveNamesW" }

func (o *xxx_ResolveNamesWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat != nil {
			if err := o.Stat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Stat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.PropertyTags != nil {
			_ptr_pPropTags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyTags != nil {
					if err := o.PropertyTags.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyTags, _ptr_pPropTags); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// paWStr {in} (1:{pointer=ref}*(1))(2:{alias=WStringsArray_r}(struct))
	{
		if o.Strings != nil {
			if err := o.Strings.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&StringsArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRpc {in} (1:{context_handle, alias=NSPI_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Handle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pStat {in} (1:{pointer=ref}*(1))(2:{alias=STAT}(struct))
	{
		if o.Stat == nil {
			o.Stat = &Stat{}
		}
		if err := o.Stat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pPropTags {in} (1:{pointer=unique}*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_pPropTags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyTags == nil {
				o.PropertyTags = &PropertyTagArray{}
			}
			if err := o.PropertyTags.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pPropTags := func(ptr interface{}) { o.PropertyTags = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.PropertyTags, _s_pPropTags, _ptr_pPropTags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// paWStr {in} (1:{pointer=ref}*(1))(2:{alias=WStringsArray_r}(struct))
	{
		if o.Strings == nil {
			o.Strings = &StringsArray{}
		}
		if err := o.Strings.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		if o.MIDs != nil {
			_ptr_ppMIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MIDs != nil {
					if err := o.MIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyTagArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MIDs, _ptr_ppMIds); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		if o.Rows != nil {
			_ptr_ppRows := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rows != nil {
					if err := o.Rows.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyRowSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rows, _ptr_ppRows); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveNamesWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppMIds {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyTagArray_r}(struct))
	{
		_ptr_ppMIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MIDs == nil {
				o.MIDs = &PropertyTagArray{}
			}
			if err := o.MIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppMIds := func(ptr interface{}) { o.MIDs = *ptr.(**PropertyTagArray) }
		if err := w.ReadPointer(&o.MIDs, _s_ppMIds, _ptr_ppMIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppRows {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PropertyRowSet_r}(struct))
	{
		_ptr_ppRows := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rows == nil {
				o.Rows = &PropertyRowSet{}
			}
			if err := o.Rows.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRows := func(ptr interface{}) { o.Rows = *ptr.(**PropertyRowSet) }
		if err := w.ReadPointer(&o.Rows, _s_ppRows, _ptr_ppRows); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResolveNamesWRequest structure represents the NspiResolveNamesW operation request
type ResolveNamesWRequest struct {
	// hRpc: An RPC context handle as specified in section 2.3.9.
	Handle *Handle `idl:"name:hRpc" json:"handle"`
	// pStat: A reference to a STAT block describing a logical position in a specific address
	// book container.
	Stat *Stat `idl:"name:pStat" json:"stat"`
	// pPropTags: The value NULL or a reference to a PropertyTagArray_r containing a list
	// of the proptags of the columns that the client requests to be returned for each row
	// returned.
	PropertyTags *PropertyTagArray `idl:"name:pPropTags;pointer:unique" json:"property_tags"`
	// paWStr: A WStringsArray_r value. Specifies the values on which the client is requesting
	// that the server perform ANR. The server MUST apply any necessary character set conversion
	// as specified in String Handling (section 3.1.1.2).
	Strings *StringsArray `idl:"name:paWStr" json:"strings"`
}

func (o *ResolveNamesWRequest) xxx_ToOp(ctx context.Context, op *xxx_ResolveNamesWOperation) *xxx_ResolveNamesWOperation {
	if op == nil {
		op = &xxx_ResolveNamesWOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Stat = o.Stat
	op.PropertyTags = o.PropertyTags
	op.Strings = o.Strings
	return op
}

func (o *ResolveNamesWRequest) xxx_FromOp(ctx context.Context, op *xxx_ResolveNamesWOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Stat = op.Stat
	o.PropertyTags = op.PropertyTags
	o.Strings = op.Strings
}
func (o *ResolveNamesWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResolveNamesWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveNamesWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResolveNamesWResponse structure represents the NspiResolveNamesW operation response
type ResolveNamesWResponse struct {
	// ppMIds: A PropertyTagArray_r value. On return, contains a list of MIds matching the
	// array of strings, as specified in the input parameter paWStr
	MIDs *PropertyTagArray `idl:"name:ppMIds" json:"m_ids"`
	// ppRows: A reference to a PropertyRowSet_r. Contains the address book container rows
	// that the server returns in response to the request.
	Rows *PropertyRowSet `idl:"name:ppRows" json:"rows"`
	// Return: The NspiResolveNamesW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResolveNamesWResponse) xxx_ToOp(ctx context.Context, op *xxx_ResolveNamesWOperation) *xxx_ResolveNamesWOperation {
	if op == nil {
		op = &xxx_ResolveNamesWOperation{}
	}
	if o == nil {
		return op
	}
	op.MIDs = o.MIDs
	op.Rows = o.Rows
	op.Return = o.Return
	return op
}

func (o *ResolveNamesWResponse) xxx_FromOp(ctx context.Context, op *xxx_ResolveNamesWOperation) {
	if o == nil {
		return
	}
	o.MIDs = op.MIDs
	o.Rows = op.Rows
	o.Return = op.Return
}
func (o *ResolveNamesWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResolveNamesWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveNamesWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

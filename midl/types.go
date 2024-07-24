package midl

import (
	"fmt"
	"sort"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

// TypeStore interface represents the type storage to
// resolve kind of TypeRef.
type TypeStore interface {
	// LookupType function must resolve the type reference
	// and return type or 'false'
	LookupType(n string) (*Type, bool)
}

// Interface structure represents the interface definition.
type Interface struct {
	// The interface name.
	Name string `json:"name,omitempty"`
	// The interface base.
	BaseName string `json:"base_name,omitempty"`
	// The interface base.
	Base *Interface `json:"-"`
	// The interface attributes.
	Attrs *InterfaceAttr `json:"attrs,omitempty"`
	// Body is an interface body (including export symbols and function calls).
	Body InterfaceBody `json:"body,omitempty"`
}

func (iff *Interface) IsObject() bool {
	return iff != nil && (iff.BaseName != "" || iff.Attrs.Object)
}

// Exports function ...
func (iff *Interface) Exports() []*Export {
	return iff.Body.Exports()
}

// InterfaceBody structure represents the interface body definitions.
type InterfaceBody struct {
	// Imports is a list of file paths for the imported symbols.
	Imports []string `json:"imports"`
	// Export is a map for the exported symbols.
	Export map[string]*Export `json:"exports"`
	// Operations is a list of the interface operations.
	Operations []*Operation `json:"operations"`
}

// Exports function ...
func (b InterfaceBody) Exports() []*Export {

	exportSyms := make([]*Export, 0, len(b.Export))
	for _, sym := range b.Export {
		exportSyms = append(exportSyms, sym)
	}

	sort.Slice(exportSyms, func(i, j int) bool {
		return exportSyms[i].Position < exportSyms[j].Position
	})

	return exportSyms
}

// Export structure represents the exported symbol.
type Export struct {
	// The positional number.
	Position int
	// The export symbol name.
	Name string
	// The constant value, or nil.
	Const *Const `json:"const,omitempty"`
	// The type value, or nil.
	Type *Type `json:"type,omitempty"`
	// The list of aliases of the tagged declarator.
	Aliases []string `json:"aliases,omitempty"`
}

// Operation structure represents the RPC call.
type Operation struct {
	// The operation name.
	Name string `json:"name"`
	// The operation number.
	OpNum int `json:"op_num"`
	// The operation return type.
	Type *Type `json:"type"`
	// The operation parameter list.
	Params []*Param `json:"params"`
	// The operation attributes.
	Attrs *OperationAttr `json:"attrs,omitempty"`
}

func (o *Operation) GetName() string {
	if o.Attrs.PropGet {
		return "get_" + o.Name
	}
	if o.Attrs.PropPut {
		return "put_" + o.Name
	}
	if o.Attrs.PropPutRef {
		return "putref_" + o.Name
	}
	return o.Name
}

// Const structure represents the constant declaration.
// This declaration is obtained either via `const` expression
// or via #define statement.
type Const struct {
	// The constant name.
	Name string `json:"name"`
	// The constant type.
	Type Kind `json:"kind"`
	// The constant value.
	Value Expr `json:"value"`
}

// Param structure represents the operation/function parameter
// declaration.
type Param struct {
	// The parameter name.
	Name string `json:"name,omitempty"`
	// The parameter type.
	Type *Type `json:"type,omitempty"`
	// The parameter attributes.
	Attrs *ParamAttr `json:"attrs,omitempty"`
}

func (p *Param) IsHandle() bool {
	if p.Type.Is(TypeAttribute) {
		return p.Type.Elem.Is(TypeHandle)
	}
	return p.Type.Is(TypeHandle)
}

// Array structure represents the array type.
type Array struct {
	// The array bounds.
	Bound ArrayBound
}

// IsFixed function returns true if array is of fixed size.
func (a Array) IsFixed() bool {
	return a.Bound.Lower >= 0 && a.Bound.Upper >= 0
}

// Size function returns the fixed size array.
func (a Array) Size() int64 {
	return a.Bound.Upper - a.Bound.Lower + 1
}

// ArrayBound structure represents the array bounds
// declaration.
type ArrayBound struct {
	// The upper and lower bounds for the array declaration.
	Upper, Lower int64
}

// Func structure represents the function type.
type Func struct {
	// The function parameters list.
	Params []*Param
}

// Field structure represents the union/structure field
// declaration.
type Field struct {
	// The field position.
	Position int `json:"-"`
	// The field name.
	Name string `json:"name,omitempty"`
	// The field attributes.
	Attrs *FieldAttr `json:"attrs,omitempty"`
	// The field type.
	Type *Type `json:"type,omitempty"`
	// Default Value.
	DefaultValue Expr `json:"default_value,omitempty"`
}

func (f *Field) IsString() bool {
	return f.Attrs.Usage.IsString || (f.Type.Attrs != nil && f.Type.Attrs.Usage.IsString)
}

func (f *Field) IsHandle() bool {
	if f.Type.Is(TypeAttribute) {
		return f.Type.Elem.Is(TypeHandle)
	}
	return f.Type.Is(TypeHandle)
}

// Enum structure represents the enumeration type.
type Enum struct {
	// Requires 32-bit uint type.
	Is32 bool
	// The enumeration elements.
	Elems []*Element `json:"elements"`
}

// Element structure represents the enumeration element.
type Element struct {
	// The enumeration element label.
	Value string `json:"value"`
	// The enumeration element identifier.
	ID int `json:"id"`
}

// Union structure represents the union type.
type Union struct {
	// The list of union switch/case statements.
	Body   []*UnionCase `json:"body,omitempty"`
	Switch *UnionSwitch `json:"switch,omitempty"`
}

func (u *Union) IsEncapsulated() bool {
	return u.Switch != nil
}

// UnionCase structure is a union case statement, if Label is nil,
// this is default case.
type UnionCase struct {
	Position  int           `json:"position,omitempty"`
	Labels    []interface{} `json:"labels,omitempty"`
	Arms      []*Field      `json:"arms,omitempty"`
	IsDefault bool          `json:"is_default,omitempty"`
}

// UnionSwitch structure represents the union switch
// (switch_is declaration).
type UnionSwitch struct {
	Name string `json:"name"`
	Type *Type  `json:"type,omitempty"`
}

// Struct structure represents the structure type.
type Struct struct {
	// The structure fields.
	Fields []*Field `json:"fields,omitempty"`
}

// LastField function
func (s *Struct) LastField() *Field {
	return s.Fields[len(s.Fields)-1]
}

// Type structure represents the type declaration.
type Type struct {
	Kind      Kind       `json:"kind,omitempty"`
	Charset   Charset    `json:"charset,omitempty"`
	Name      string     `json:"name,omitempty"`
	Tag       string     `json:"tag,omitempty"`
	Func      *Func      `json:"func,omitempty"`
	Array     *Array     `json:"array,omitempty"`
	Enum      *Enum      `json:"enum,omitempty"`
	Struct    *Struct    `json:"struct,omitempty"`
	Union     *Union     `json:"union,omitempty"`
	Interface *Interface `json:"-"`
	Attrs     *TypeAttr  `json:"attrs,omitempty"`
	Alias     string     `json:"alias,omitempty"`
	Elem      *Type      `json:"elem,omitempty"`
}

func (t *Type) flat() []*Type {

	var list []*Type

	for t := t; t != nil; t = t.Elem {
		pt := *t
		pt.Elem = nil
		list = append(list, &pt)
	}

	return list
}

func (t *Type) Base() *Type {
	for t.Elem != nil {
		t = t.Elem
	}
	return t
}

func (t *Type) Append(tt *Type) *Type {
	if t == nil {
		return tt
	}
	if base := t.Base(); !base.Is(TypePointer) && !base.Is(TypeArray) && !base.Is(TypeFunc) && !base.Is(TypeAttribute) {
		panic(fmt.Sprintf("invalid type %s for Type.Append operation", base.Kind))
	} else {
		base.Elem = tt
	}
	return t
}

func (t *Type) AppendAfterPointer(tt *Type) *Type {

	top := t

	for t.Elem != nil && t.Elem.Elem != nil && t.Elem.Elem.Elem != nil {
		t = t.Elem
	}
	t.Elem = tt
	return top
}

// TypeName function returns the name of the type.
func (t *Type) TypeName() string {
	if t.Name == "" {
		if !t.Is(TypeTag) {
			return "UNKNOWN"
		}
		return TagName(t.Kind, t.Tag)
	}
	return t.Name
}

// TagName function transforms the kind and tag to a tagged name.
func TagName(kind Kind, tag string) string {
	n := ""
	switch kind {
	case TypeEnum:
		n = "_enum_"
	case TypeStruct:
		n = "_struct_"
	case TypeUnion:
		n = "_union_"
	case TypeCUnion:
		n = "_c_union_"
	}
	return n + tag
}

// Is function returns true if type is of a given kind. Note,
// that synthetic type `TypeTag` is used to determine if type
// can be a tagged type.
func (t *Type) Is(kind Kind) bool {
	if t == nil {
		return false
	}
	if kind == TypeTag {
		return t.Kind == TypeStruct || t.Kind == TypeEnum || t.Kind == TypeUnion || t.Kind == TypeCUnion
	}
	return t.Kind == kind
}

func (t *Type) IsPrimitiveType() bool {

	switch t.Kind {
	case TypeChar,
		TypeUChar,
		TypeWChar,
		TypeBoolean,
		TypeInt8,
		TypeUint8,
		TypeInt16,
		TypeUint16,
		TypeInt32,
		TypeInt32_64,
		TypeUint32_64,
		TypeUint32,
		TypeError,
		TypeInt64,
		TypeUint64,
		TypeFloat32,
		TypeFloat64:
		// TypeHandle,
		// TypeDispInterface:
		return true
	}

	return false
}

// InterfaceAttr ...
type InterfaceAttr struct {
	UUID           *uuid.UUID
	Version        *Version
	Endpoints      []string
	Exceptions     []string
	Local          bool
	PointerDefault PointerType
	MSUnion        bool
	Object         bool
	HelpString     string
	Dual           bool
	Hidden         bool
	Nonextensible  bool
	ODL            bool
	OLEAutomation  bool
}

// TypeAttr ...
type TypeAttr struct {
	TransmitAs              *Type
	Handle                  bool
	SwitchType              *Type
	Usage                   Usage
	Format                  Format
	Pointer                 PointerType
	V1Enum                  bool
	Range                   *Range
	DisableConsistencyCheck bool
	WireMarshal             string
	Public                  bool
	Alias                   string
	// names, pointers to keep track of lost data merged.
	Names    []string
	Pointers []PointerType
	Parent   string
	Pad      uint64
	IsLayout bool
}

func (t *TypeAttr) EnumType() Kind {
	if t.V1Enum {
		return TypeUint32
	}
	return TypeUint16
}

func (t *TypeAttr) withAlias(a string) *TypeAttr {
	t.Alias = a
	return t
}

func (t *TypeAttr) clone() *TypeAttr {
	return &TypeAttr{
		t.TransmitAs,
		t.Handle,
		t.SwitchType,
		t.Usage,
		t.Format,
		t.Pointer,
		t.V1Enum,
		t.Range,
		t.DisableConsistencyCheck,
		t.WireMarshal,
		t.Public,
		t.Alias,
		append([]string{}, t.Names...),
		append([]PointerType{}, t.Pointers...),
		t.Parent,
		t.Pad,
		t.IsLayout,
	}
}

func (t *TypeAttr) Merge(tt *TypeAttr) *TypeAttr {

	if t = t.clone(); tt == nil {
		return t
	}

	if tt.TransmitAs != nil {
		t.TransmitAs = tt.TransmitAs
	}
	if tt.Handle {
		t.Handle = tt.Handle
	}
	if tt.SwitchType != nil {
		t.SwitchType = tt.SwitchType
	}
	if tt.Usage.IsString {
		t.Usage.IsString = tt.Usage.IsString
	}
	if tt.Usage.ContextHandle {
		t.Usage.ContextHandle = tt.Usage.ContextHandle
	}
	if tt.Format.NullTerminated {
		t.Format.NullTerminated = tt.Format.NullTerminated
	}
	if tt.Format.UTF8 {
		t.Format.UTF8 = tt.Format.UTF8
	}
	if tt.Format.Rune {
		t.Format.Rune = tt.Format.Rune
	}
	if tt.Format.Hex {
		t.Format.Hex = tt.Format.Hex
	}
	if tt.Format.MultiSize {
		t.Format.MultiSize = tt.Format.MultiSize
	}
	if t.Pointer == PointerTypeNone || (tt.Pointer != PointerTypeNone && tt.Pointer != PointerTypeRefWeak) {
		// keep track of older pointer information.
		if t.Pointer != PointerTypeNone {
			t.Pointers = append(t.Pointers, t.Pointer)
		}
		if tt.Pointer == PointerTypeRefWeak {
			t.Pointer = PointerTypeRef
		} else {
			t.Pointer = tt.Pointer
		}
	}
	if tt.V1Enum {
		t.V1Enum = tt.V1Enum
	}
	if tt.Range != nil {
		t.Range = tt.Range
	}
	if tt.DisableConsistencyCheck {
		t.DisableConsistencyCheck = tt.DisableConsistencyCheck
	}
	if tt.WireMarshal != "" {
		t.WireMarshal = tt.WireMarshal
	}
	if tt.Public {
		t.Public = tt.Public
	}
	if tt.Alias != "" {
		// keep track of alias information.
		if t.Alias != "" && t.Alias != tt.Alias {
			t.Names = append(t.Names, t.Alias)
		}
		t.Alias = tt.Alias
	}

	if tt.Parent != "" {
		t.Parent = tt.Parent
	}
	if tt.Pad != 0 {
		t.Pad = tt.Pad
	}
	if tt.IsLayout {
		t.IsLayout = tt.IsLayout
	}

	return t
}

// OperationAttr ...
type OperationAttr struct {
	Idempotent       bool
	Broadcast        bool
	Maybe            bool
	ReflectDeletions bool
	Usage            Usage
	Format           Format
	Pointer          PointerType
	Callback         bool
	PropGet          bool
	PropPut          bool
	PropPutRef       bool
	ID               Expr
	Restricted       bool
	CallAs           string
}

// ParamAttr ...
type ParamAttr struct {
	*FieldAttr
	Direction               Direction
	DisableConsistencyCheck bool
	IIDIs                   Expr
	Retval                  bool
	DefaultValue            Expr
	Optional                bool
	Annotation              string
}

type SizeAttr struct {
	MinIs  []Expr
	MaxIs  []Expr
	SizeIs []Expr
}

// FieldAttr ...
type FieldAttr struct {
	FirstIs     []Expr
	LastIs      []Expr
	LengthIs    []Expr
	MinIs       []Expr
	MaxIs       []Expr
	SizeIs      []Expr
	Usage       Usage
	Format      Format
	SwitchIs    Expr
	Ignore      bool
	Pointer     PointerType
	Range       *Range
	SwitchType  *Type
	Safearray   *Type
	Layout      []*Field
	NoSizeLimit bool
	IsLayout    bool
}

func (f *FieldAttr) SizeAttr() *SizeAttr {
	if f.SizeIs == nil && f.MaxIs == nil && f.MinIs == nil {
		return nil
	}
	return &SizeAttr{f.MinIs, f.MaxIs, f.SizeIs}
}

func (f *FieldAttr) Dim() []Dim {

	max := max(f.NoSizeLimit, f.SizeIs, f.MaxIs, f.MinIs, f.FirstIs, f.LastIs, f.LengthIs)

	dims := make([]Dim, max)

	for i := 0; i < max; i++ {
		if len(f.SizeIs) > i {
			dims[i].SizeIs = f.SizeIs[i]
		}
		if len(f.MaxIs) > i {
			dims[i].MaxIs = f.MaxIs[i]
		}
		if len(f.MinIs) > i {
			dims[i].MinIs = f.MinIs[i]
		}
		if len(f.LengthIs) > i {
			dims[i].LengthIs = f.LengthIs[i]
		}
		if len(f.FirstIs) > i {
			dims[i].FirstIs = f.FirstIs[i]
		}
		if len(f.LastIs) > i {
			dims[i].LastIs = f.LastIs[i]
		}
		if f.NoSizeLimit {
			dims[i].NoSizeLimit = f.NoSizeLimit
		}
	}
	return dims
}

func max(noSizeLimit bool, exprs ...[]Expr) int {
	var max int
	if noSizeLimit {
		max++
	}
	for _, expr := range exprs {
		if len(expr) > max {
			max = len(expr)
		}
	}
	return max
}

func (f *FieldAttr) TypeAttr() *TypeAttr {
	if f != nil {
		return &TypeAttr{
			SwitchType: f.SwitchType,
			Range:      f.Range,
			Pointer:    f.Pointer,
			Usage:      f.Usage,
			Format:     f.Format,
			IsLayout:   f.IsLayout,
		}
	}
	return &TypeAttr{}
}

// Usage structure represents the field, parameter usage attribute.
type Usage struct {
	ContextHandle bool
	IsString      bool
}

type Format struct {
	NullTerminated bool
	UTF8           bool
	MultiSize      bool
	Rune           bool
	Hex            bool
}

type Direction struct {
	// Determines the operation/function parameter direction.
	In, Out bool
}

// Version structure represents the interface version structure.
type Version struct {
	Minor, Major uint16
}

func (v *Version) String() string {
	if v == nil {
		return "v0"
	}

	if v.Minor == 0 {
		return fmt.Sprintf("v%d", v.Major)
	}

	return fmt.Sprintf("v%d.%d", v.Major, v.Minor)
}

// Range structure repsents the range attribute.
type Range struct {
	Min, Max int64
}

type PointerType int

const (
	PointerTypeNone PointerType = iota
	PointerTypeRef
	PointerTypePtr
	PointerTypeUnique
	PointerTypeRefWeak
)

// Kind is the type selector.
type Kind int

const (
	TypeInvalid Kind = iota
	TypeInt8
	TypeUint8
	TypeInt16
	TypeUint16
	TypeInt32
	TypeInt64
	TypeUint32
	TypeUint64
	TypeInt32_64  // 32 for NDR, 64 for NDR64 representation.
	TypeUint32_64 // 32 for NDR, 64 for NDR64 representation.
	TypeFloat32
	TypeFloat64
	TypeCharset
	TypeUChar
	TypeChar
	TypeWChar
	TypeEnum
	TypeString
	TypeStruct
	TypeUnion
	TypeCUnion
	TypeFunc
	TypeArray
	TypePointer
	TypePipe
	TypeBoolean
	TypeVoid // constant declaration only.
	TypeHandle
	TypeError
	TypeInterface
	TypeDispInterface
	TypeRef       // type reference, resolved via lookup.
	TypeAttribute //
	TypeTag
)

func (k Kind) Signed() bool {
	return k == TypeInt8 || k == TypeInt16 || k == TypeInt32 || k == TypeInt64 || k == TypeInt32_64 ||
		k == TypeFloat32 || k == TypeFloat64 || k == TypeChar || k == TypeWChar
}

// Charset is a predefined character set, as per DCE/RPC [C706].
type Charset int

const (
	CharsetNone Charset = iota
	CharsetISO_Latin_1
	CharsetISO_Multilingual
	CharsetISO_UCS
)

func IsIntegerType(kind Kind) bool {
	switch kind {
	case TypeInt8, TypeUint8, TypeInt16, TypeUint16, TypeInt32, TypeUint32, TypeInt64, TypeUint64:
		return true
	}
	return false
}

func PrimitiveTypeSize(kind Kind) int {

	var ret int

	switch kind {
	case TypeBoolean, TypeChar, TypeUChar, TypeInt8, TypeUint8:
		ret = 1
	case TypeUint16, TypeInt16, TypeWChar:
		ret = 2
	case TypeFloat32, TypeUint32, TypeInt32, TypeError:
		ret = 4
	case TypeFloat64, TypeUint64, TypeInt64:
		ret = 8
	case TypeInt32_64, TypeUint32_64:
		return 5
	}

	return ret
}

// TypeSize ...
func TypeSize(s TypeStore, t *Type) int {

	var ret int

	switch t.Kind {
	case TypeBoolean, TypeChar, TypeUChar, TypeInt8, TypeUint8:
		ret = 1
	case TypeUint16, TypeInt16, TypeWChar:
		ret = 2
	case TypeFloat32, TypeUint32, TypeInt32:
		ret = 4
	case TypeFloat64, TypeUint64, TypeInt64:
		ret = 8
	case TypeEnum:
		if t.Attrs.V1Enum {
			ret = 4
		} else {
			ret = 2
		}
	case TypeStruct:
		for _, f := range t.Struct.Fields {
			if sz := TypeSize(s, f.Type); sz == 0 {
				return 0
			} else {
				ret += sz
			}
		}
	case TypeRef:
		tt, ok := s.LookupType(t.Name)
		if !ok {
			return 0
		}
		ret = TypeSize(s, tt)
	case TypePointer:
		ret = 4
	case TypeArray:
		if t.Array.IsFixed() {
			ret = int(t.Array.Bound.Upper) * TypeSize(s, t.Elem)
		}
	case TypeAttribute:
		return TypeSize(s, t.Elem)
	default:
		ret = 0
	}

	return ret
}

package midl

// parser_type.go contains definitions for parser helper types.

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

type pTagID struct {
	Tag string
	ID  Expr
}

type pTagIDs []pTagID

type pDeclarator struct {
	Name string
	Type *Type
}

type pDeclarators []*pDeclarator

type pTypedef struct {
	Type        *Type
	Attrs       *TypeAttr
	Declarators pDeclarators
}

func SetSwitchType(fields []*Field) {
	for _, f := range fields {
		if f.Attrs.SwitchIs.Empty() {
			continue
		}
		if f.Attrs.SwitchType != nil {
			continue
		}

		if f.Type.Attrs != nil && f.Type.Attrs.SwitchType != nil {
			f.Attrs.SwitchType = f.Type.Attrs.SwitchType
			continue
		}

		for _, ff := range fields {
			if f.Attrs.SwitchIs.Ident().String() == ff.Name {
				f.Attrs.SwitchType = ff.Type
			}
		}
	}
}

// SetTagName function tries to set the tag name for the type.
func (p *pTypedef) SetTagName() bool {

	if !p.Type.Is(TypeTag) {
		p.Type.Tag = ""
		return false
	}

	if p.Type.Tag != "" {
		return true
	}

	for _, decl := range p.Declarators {
		if decl.Type == nil {
			p.Type.Tag = decl.Name
			break
		}
	}

	if p.Type.Tag == "" {
		if len(p.Declarators) > 0 {
			p.Type.Tag = p.Declarators[0].Name
		} else {
			p.Type.Tag = MD5(fmt.Sprintf("%+v", p))
		}
	}

	return true
}

func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (p *pTypedef) pToExport_() []*Export {

	ret := []*Export{}

	// try assign attribute alias.
	for _, decl := range p.Declarators {
		// handle void * context_handle
		if p.Attrs.Usage.ContextHandle && decl.Type.Is(TypePointer) {
			decl.Type = decl.Type.Elem
		}
	}

	for _, decl := range p.Declarators {
		if decl.Type != nil {
			continue
		}
		p.Attrs.Alias = decl.Name
		break
	}

	if p.SetTagName() {
		// if type is taggable, create exported
		// symbol with name of the tag.
		ret = append(ret, &Export{Name: p.Type.TypeName(), Type: p.Type})
		for _, decl := range p.Declarators {
			if decl.Type == nil {
				// add aliases to the exported symbol for lookup.
				ret[0].Aliases = append(ret[0].Aliases, decl.Name)
			}
		}
	}

	// the current alias for line of:
	// typedef X *PY, Y, Z, *PZ, such that PY will have an
	// alias of Y and PZ of Z.
	alias := p.Attrs.Alias

	for _, decl := range p.Declarators {
		if decl.Type == nil {
			// this alias will be the main one for the
			// next pointer declarators.
			alias = decl.Name
		} else if alias != "" {
			// add alias right after the type declarator.
			decl.Type = decl.Type.Append(&Type{
				Kind:  TypeAttribute,
				Attrs: p.Attrs.clone().withAlias(alias),
			})
		}

		// type is an attribute.
		decl.Type = &Type{
			Kind:  TypeAttribute,
			Attrs: p.Attrs.clone().withAlias(decl.Name),
			Elem:  decl.Type,
		}

		ret = append(ret, &Export{
			Name: decl.Name,
			Type: decl.Type.Append(p.Type),
		})
	}

	return ret
}

type pAttrType struct {
	Type int
	Attr pAttr
}

type pAttr struct {
	FirstIs                 []Expr
	LastIs                  []Expr
	LengthIs                []Expr
	MinIs                   []Expr
	MaxIs                   []Expr
	SizeIs                  []Expr
	Pointer                 PointerType
	Usage                   Usage
	Format                  Format
	SwitchIs                Expr
	Ignore                  bool
	Direction               Direction
	TransmitAs              *Type
	SwitchType              *Type
	Handle                  bool
	Idempotent              bool
	Broadcast               bool
	Maybe                   bool
	ReflectDeletions        bool
	UUID                    *uuid.UUID
	Version                 *Version
	Endpoints               []string
	Exceptions              []string
	Local                   bool
	PointerDefault          PointerType
	V1Enum                  bool
	MSUnion                 bool
	Range                   *Range
	DisableConsistencyCheck bool
	Object                  bool
	Callback                bool
	IIDIs                   Expr
	Retval                  bool
	HelpString              string
	Dual                    bool
	PropGet                 bool
	PropPut                 bool
	PropPutRef              bool
	ID                      Expr
	Hidden                  bool
	Nonextensible           bool
	Restricted              bool
	DefaultValue            Expr
	ODL                     bool
	OLEAutomation           bool
	Default                 bool
	Source                  bool
	Optional                bool
	AppObject               bool
	Annotation              string
	CallAs                  string
	WireMarshal             string
	Public                  bool
	Safearray               *Type
	Alias                   string
	Pad                     uint64
	Layout                  []*Field
	NoSizeLimit             bool
	IsLayout                bool
	NullIf                  Expr
}

func (a pAttr) Set(at pAttrType) pAttr {
	switch at.Type {
	case FIRST_IS:
		a.FirstIs = at.Attr.FirstIs
	case LAST_IS:
		a.LastIs = at.Attr.LastIs
	case LENGTH_IS:
		a.LengthIs = at.Attr.LengthIs
	case MIN_IS:
		a.MinIs = at.Attr.MinIs
	case MAX_IS:
		a.MaxIs = at.Attr.MaxIs
	case SIZE_IS:
		a.SizeIs = at.Attr.SizeIs
	case POINTER_REF:
		a.Pointer = PointerTypeRef
	case POINTER_PTR:
		a.Pointer = PointerTypePtr
	case POINTER_UNIQUE:
		a.Pointer = PointerTypeUnique
	case USAGE_STRING:
		a.Usage.IsString = true
	case USAGE_CONTEXT_HANDLE:
		a.Usage.ContextHandle = true
	case FORMAT_NULL_TERMINATED:
		a.Format.NullTerminated = true
	case FORMAT_UTF8:
		a.Format.UTF8 = true
	case FORMAT_MULTI_SIZE:
		a.Format.MultiSize = true
	case FORMAT_RUNE:
		a.Format.Rune = true
	case FORMAT_HEX:
		a.Format.Hex = true
	case SWITCH_IS:
		a.SwitchIs = at.Attr.SwitchIs
	case IGNORE:
		a.Ignore = true
	case IN:
		a.Direction.In = true
	case OUT:
		a.Direction.Out = true
	case TRANSMIT_AS:
		a.TransmitAs = at.Attr.TransmitAs
	case SWITCH_TYPE:
		a.SwitchType = at.Attr.SwitchType
	case HANDLE:
		a.Handle = true
	case IDEMPOTENT:
		a.Idempotent = true
	case BROADCAST:
		a.Broadcast = true
	case MAYBE:
		a.Maybe = true
	case REFLECT_DELETIONS:
		a.ReflectDeletions = true
	case UUID:
		a.UUID = at.Attr.UUID
	case VERSION:
		a.Version = at.Attr.Version
	case ENDPOINT:
		a.Endpoints = at.Attr.Endpoints
	case EXCEPTIONS:
		a.Exceptions = at.Attr.Exceptions
	case LOCAL:
		a.Local = true
	case POINTER_DEFAULT:
		a.PointerDefault = at.Attr.PointerDefault
	case V1_ENUM:
		a.V1Enum = true
	case MS_UNION:
		a.MSUnion = true
	case RANGE:
		a.Range = at.Attr.Range
	case DISABLE_CONSISTENCY_CHECK:
		a.DisableConsistencyCheck = true
	case OBJECT:
		a.Object = true
	case CALLBACK:
		a.Callback = true
	case RETVAL:
		a.Retval = true
	case IID_IS:
		a.IIDIs = at.Attr.IIDIs
	case HELP_STRING:
		a.HelpString = at.Attr.HelpString
	case DUAL:
		a.Dual = true
	case PROPGET:
		a.PropGet = true
	case PROPPUT:
		a.PropPut = true
	case PROPPUTREF:
		a.PropPutRef = true
	case ID:
		a.ID = at.Attr.ID
	case HIDDEN:
		a.Hidden = true
	case NONEXTENSIBLE:
		a.Nonextensible = true
	case RESTRICTED:
		a.Restricted = true
	case DEFAULT_VALUE:
		a.DefaultValue = at.Attr.DefaultValue
	case ODL:
		a.ODL = true
	case OLEAUTOMATION:
		a.OLEAutomation = true
	case OPTIONAL:
		a.Optional = true
	case APPOBJECT:
		a.AppObject = true
	case ANNOTATION:
		a.Annotation = at.Attr.Annotation
	case CALL_AS:
		a.CallAs = at.Attr.CallAs
	case WIRE_MARSHAL:
		a.WireMarshal = at.Attr.WireMarshal
	case PUBLIC:
		a.Public = true
	case SAFEARRAY:
		a.Safearray = at.Attr.Safearray
	case PAD:
		a.Pad = at.Attr.Pad
	case GOEXT_LAYOUT:
		a.Layout = at.Attr.Layout
	case GOEXT_NO_SIZE_LIMIT:
		a.NoSizeLimit = at.Attr.NoSizeLimit
	case GOEXT_NULL_IF:
		a.NullIf = at.Attr.NullIf
	default:
		panic(fmt.Sprintf("unknown attribute: %d", at.Type))
	}

	return a
}

func (a pAttr) MapAttr() map[int]interface{} {

	var ret = make(map[int]interface{})

	if a.FirstIs != nil {
		ret[FIRST_IS] = a.FirstIs
	}
	if a.LastIs != nil {
		ret[LAST_IS] = a.LastIs
	}
	if a.LengthIs != nil {
		ret[LENGTH_IS] = a.LengthIs
	}
	if a.MinIs != nil {
		ret[MIN_IS] = a.MinIs
	}
	if a.MaxIs != nil {
		ret[MAX_IS] = a.MaxIs
	}
	if a.SizeIs != nil {
		ret[SIZE_IS] = a.SizeIs
	}
	if a.Pointer != PointerTypeNone {
		ret[POINTER] = a.Pointer
	}
	if a.Usage.IsString {
		ret[USAGE_STRING] = a.Usage.IsString
	}
	if a.Usage.ContextHandle {
		ret[USAGE_CONTEXT_HANDLE] = a.Usage.ContextHandle
	}
	if a.Format.NullTerminated {
		ret[FORMAT_NULL_TERMINATED] = a.Format.NullTerminated
	}
	if a.Format.UTF8 {
		ret[FORMAT_UTF8] = a.Format.UTF8
	}
	if a.Format.MultiSize {
		ret[FORMAT_MULTI_SIZE] = a.Format.MultiSize
	}
	if a.Format.Rune {
		ret[FORMAT_RUNE] = a.Format.Rune
	}
	if a.Format.Hex {
		ret[FORMAT_HEX] = a.Format.Hex
	}
	if !a.SwitchIs.Empty() {
		ret[SWITCH_IS] = a.SwitchIs
	}
	if a.Ignore {
		ret[IGNORE] = a.Ignore
	}
	if a.Direction.In {
		ret[IN] = a.Direction.In
	}
	if a.Direction.Out {
		ret[OUT] = a.Direction.Out
	}
	if a.TransmitAs != nil {
		ret[TRANSMIT_AS] = a.TransmitAs
	}
	if a.SwitchType != nil {
		ret[SWITCH_TYPE] = a.SwitchType
	}
	if a.Handle {
		ret[HANDLE] = a.Handle
	}
	if a.Idempotent {
		ret[IDEMPOTENT] = a.Idempotent
	}
	if a.Broadcast {
		ret[BROADCAST] = a.Broadcast
	}
	if a.Maybe {
		ret[MAYBE] = a.Maybe
	}
	if a.ReflectDeletions {
		ret[REFLECT_DELETIONS] = a.ReflectDeletions
	}
	if a.UUID != nil {
		ret[UUID] = a.UUID
	}
	if a.Version != nil {
		ret[VERSION] = a.Version
	}
	if a.Endpoints != nil {
		ret[ENDPOINT] = a.Endpoints
	}
	if a.Exceptions != nil {
		ret[EXCEPTIONS] = a.Exceptions
	}
	if a.Local {
		ret[LOCAL] = a.Local
	}
	if a.PointerDefault != PointerTypeNone {
		ret[POINTER_DEFAULT] = a.PointerDefault
	}
	if a.V1Enum {
		ret[V1_ENUM] = a.V1Enum
	}
	if a.MSUnion {
		ret[MS_UNION] = a.MSUnion
	}
	if a.Range != nil {
		ret[RANGE] = a.Range
	}
	if a.DisableConsistencyCheck {
		ret[DISABLE_CONSISTENCY_CHECK] = a.DisableConsistencyCheck
	}
	if a.Object {
		ret[OBJECT] = a.Object
	}
	if a.Callback {
		ret[CALLBACK] = a.Callback
	}
	if !a.IIDIs.Empty() {
		ret[IID_IS] = a.IIDIs
	}
	if a.Retval {
		ret[RETVAL] = a.Retval
	}
	if a.HelpString != "" {
		ret[HELP_STRING] = a.HelpString
	}
	if a.Dual {
		ret[DUAL] = a.Dual
	}
	if a.PropGet {
		ret[PROPGET] = a.PropGet
	}
	if a.PropPut {
		ret[PROPPUT] = a.PropPut
	}
	if a.PropPutRef {
		ret[PROPPUTREF] = a.PropPutRef
	}
	if !a.ID.Empty() {
		ret[ID] = a.ID
	}
	if a.Hidden {
		ret[HIDDEN] = a.Hidden
	}
	if a.Nonextensible {
		ret[NONEXTENSIBLE] = a.Nonextensible
	}
	if a.Restricted {
		ret[RESTRICTED] = a.Restricted
	}
	if !a.DefaultValue.Empty() {
		ret[DEFAULT_VALUE] = a.DefaultValue
	}
	if a.ODL {
		ret[ODL] = a.ODL
	}
	if a.OLEAutomation {
		ret[OLEAUTOMATION] = a.OLEAutomation
	}
	if a.Default {
		ret[DEFAULT] = a.Default
	}
	if a.Source {
		ret[SOURCE] = a.Source
	}
	if a.Optional {
		ret[OPTIONAL] = a.Optional
	}
	if a.AppObject {
		ret[APPOBJECT] = a.AppObject
	}
	if a.CallAs != "" {
		ret[CALL_AS] = a.CallAs
	}
	if a.Annotation != "" {
		ret[ANNOTATION] = a.Annotation
	}
	if a.WireMarshal != "" {
		ret[WIRE_MARSHAL] = a.WireMarshal
	}
	if a.Public {
		ret[PUBLIC] = a.Public
	}
	if a.Safearray != nil {
		ret[SAFEARRAY] = a.Safearray
	}
	if a.Pad != 0 {
		ret[PAD] = a.Pad
	}
	if a.Layout != nil {
		ret[GOEXT_LAYOUT] = a.Layout
	}
	if a.NoSizeLimit {
		ret[GOEXT_NO_SIZE_LIMIT] = a.NoSizeLimit
	}

	if !a.NullIf.Empty() {
		ret[GOEXT_NULL_IF] = a.NullIf
	}

	return ret
}

func (a pAttr) Merge(ma pAttr) pAttr {

	if ma.FirstIs != nil {
		a.FirstIs = ma.FirstIs
	}
	if ma.LastIs != nil {
		a.LastIs = ma.LastIs
	}
	if ma.LengthIs != nil {
		a.LengthIs = ma.LengthIs
	}
	if ma.MinIs != nil {
		a.MinIs = ma.MinIs
	}
	if ma.MaxIs != nil {
		a.MaxIs = ma.MaxIs
	}
	if ma.SizeIs != nil {
		a.SizeIs = ma.SizeIs
	}
	if ma.Pointer != PointerTypeNone {
		a.Pointer = ma.Pointer
	}
	if ma.Usage.IsString {
		a.Usage.IsString = ma.Usage.IsString
	}
	if ma.Usage.ContextHandle {
		a.Usage.ContextHandle = ma.Usage.ContextHandle
	}
	if ma.Format.NullTerminated {
		a.Format.NullTerminated = ma.Format.NullTerminated
	}
	if ma.Format.MultiSize {
		a.Format.MultiSize = ma.Format.MultiSize
	}
	if ma.Format.UTF8 {
		a.Format.UTF8 = ma.Format.UTF8
	}
	if ma.Format.Rune {
		a.Format.Rune = ma.Format.Rune
	}
	if ma.Format.Hex {
		a.Format.Hex = ma.Format.Hex
	}
	if !ma.SwitchIs.Empty() {
		a.SwitchIs = ma.SwitchIs
	}
	if ma.Ignore {
		a.Ignore = ma.Ignore
	}
	if ma.Direction.In {
		a.Direction.In = ma.Direction.In
	}
	if ma.Direction.Out {
		a.Direction.Out = ma.Direction.Out
	}
	if ma.TransmitAs != nil {
		a.TransmitAs = ma.TransmitAs
	}
	if ma.SwitchType != nil {
		a.SwitchType = ma.SwitchType
	}
	if ma.Handle {
		a.Handle = ma.Handle
	}
	if ma.Idempotent {
		a.Idempotent = ma.Idempotent
	}
	if ma.Broadcast {
		a.Broadcast = ma.Broadcast
	}
	if ma.Maybe {
		a.Maybe = ma.Maybe
	}
	if ma.ReflectDeletions {
		a.ReflectDeletions = ma.ReflectDeletions
	}
	if ma.UUID != nil {
		a.UUID = ma.UUID
	}
	if ma.Version != nil {
		a.Version = ma.Version
	}
	if ma.Endpoints != nil {
		a.Endpoints = ma.Endpoints
	}
	if ma.Exceptions != nil {
		a.Exceptions = ma.Exceptions
	}
	if ma.Local {
		a.Local = ma.Local
	}
	if ma.PointerDefault != PointerTypeNone {
		a.PointerDefault = ma.PointerDefault
	}
	if ma.V1Enum {
		a.V1Enum = ma.V1Enum
	}
	if ma.MSUnion {
		a.MSUnion = ma.MSUnion
	}
	if ma.Range != nil {
		a.Range = ma.Range
	}
	if ma.DisableConsistencyCheck {
		a.DisableConsistencyCheck = ma.DisableConsistencyCheck
	}
	if ma.Object {
		a.Object = ma.Object
	}
	if ma.Callback {
		a.Callback = ma.Callback
	}
	if !ma.IIDIs.Empty() {
		a.IIDIs = ma.IIDIs
	}
	if ma.Retval {
		a.Retval = ma.Retval
	}
	if ma.HelpString != "" {
		a.HelpString = ma.HelpString
	}
	if ma.Dual {
		a.Dual = ma.Dual
	}
	if ma.PropGet {
		a.PropGet = ma.PropGet
	}
	if ma.PropPut {
		a.PropPut = ma.PropPut
	}
	if ma.PropPutRef {
		a.PropPutRef = ma.PropPutRef
	}
	if !ma.ID.Empty() {
		a.ID = ma.ID
	}
	if ma.Hidden {
		a.Hidden = ma.Hidden
	}
	if ma.Nonextensible {
		a.Nonextensible = ma.Nonextensible
	}
	if ma.Restricted {
		a.Restricted = ma.Restricted
	}
	if !ma.DefaultValue.Empty() {
		a.DefaultValue = ma.DefaultValue
	}
	if ma.ODL {
		a.ODL = ma.ODL
	}
	if ma.OLEAutomation {
		a.OLEAutomation = ma.OLEAutomation
	}
	if ma.Default {
		a.Default = ma.Default
	}
	if ma.Source {
		a.Source = ma.Source
	}
	if ma.Optional {
		a.Optional = ma.Optional
	}
	if ma.AppObject {
		a.AppObject = ma.AppObject
	}
	if ma.CallAs != "" {
		a.CallAs = ma.CallAs
	}
	if ma.Annotation != "" {
		a.Annotation = ma.Annotation
	}
	if ma.WireMarshal != "" {
		a.WireMarshal = ma.WireMarshal
	}
	if ma.Public {
		a.Public = ma.Public
	}
	if ma.Safearray != nil {
		a.Safearray = ma.Safearray
	}
	if ma.Pad != 0 {
		a.Pad = ma.Pad
	}
	if ma.Layout != nil {
		a.Layout = ma.Layout
	}
	if ma.NoSizeLimit {
		a.NoSizeLimit = ma.NoSizeLimit
	}
	if !ma.NullIf.Empty() {
		a.NullIf = ma.NullIf
	}

	return a
}

// Param function returns parameter attributes.
func (a pAttr) Param() *ParamAttr {
	return &ParamAttr{
		a.Field(),
		a.Direction,
		a.DisableConsistencyCheck,
		a.IIDIs,
		a.Retval,
		a.DefaultValue,
		a.Optional,
		a.Annotation,
	}
}

// Operation function returns operation attributes.
func (a pAttr) Operation() *OperationAttr {
	return &OperationAttr{
		a.Idempotent,
		a.Broadcast,
		a.Maybe,
		a.ReflectDeletions,
		a.Usage,
		a.Format,
		a.Pointer,
		a.Callback,
		a.PropGet,
		a.PropPut,
		a.PropPutRef,
		a.ID,
		a.Restricted,
		a.CallAs,
	}
}

func (a pAttr) ComClass() *ComClassAttr {
	return &ComClassAttr{
		a.Interface(),
		a.AppObject,
	}
}

func (a pAttr) DispatchInterface() *DispatchInterfaceAttr {
	return &DispatchInterfaceAttr{
		a.Interface(),
	}
}

func (a pAttr) Library() *LibraryAttr {
	return &LibraryAttr{
		a.Interface(),
	}
}

func (a pAttr) ComInterface() *ComInterfaceAttr {
	return &ComInterfaceAttr{
		a.Default,
		a.Source,
	}
}

var pAllowedFieldAttr = []int{
	FIRST_IS,
	LAST_IS,
	LENGTH_IS,
	MIN_IS,
	MAX_IS,
	SIZE_IS,
	USAGE_STRING,
	USAGE_CONTEXT_HANDLE,
	FORMAT,
	FORMAT_NULL_TERMINATED,
	FORMAT_UTF8,
	FORMAT_RUNE,
	FORMAT_HEX,
	FORMAT_MULTI_SIZE,
	SWITCH_IS,
	IGNORE,
	RANGE,
	POINTER,
	SWITCH_TYPE,
	GOEXT_NULL_IF,
}

var pAllowedParamAttr = []int{
	FIRST_IS,
	LAST_IS,
	LENGTH_IS,
	MIN_IS,
	MAX_IS,
	SIZE_IS,
	USAGE_STRING,
	USAGE_CONTEXT_HANDLE,
	FORMAT,
	FORMAT_NULL_TERMINATED,
	FORMAT_UTF8,
	FORMAT_RUNE,
	FORMAT_HEX,
	FORMAT_MULTI_SIZE,
	SWITCH_IS,
	IGNORE,
	RANGE,
	POINTER,
	SWITCH_TYPE,
	IN,
	OUT,
	DISABLE_CONSISTENCY_CHECK,
	IID_IS,
	RETVAL,
	OPTIONAL,
	GOEXT_NULL_IF,
}

var pAllowedTypeAttr = []int{
	TRANSMIT_AS,
	HANDLE,
	SWITCH_TYPE,
	USAGE_STRING,
	USAGE_CONTEXT_HANDLE,
	FORMAT,
	FORMAT_NULL_TERMINATED,
	FORMAT_UTF8,
	FORMAT_RUNE,
	FORMAT_HEX,
	FORMAT_MULTI_SIZE,
	POINTER,
	V1_ENUM,
	RANGE,
	DISABLE_CONSISTENCY_CHECK,
	WIRE_MARSHAL,
}

var pAllowedInterfaceAttr = []int{
	UUID,
	VERSION,
	ENDPOINT,
	EXCEPTIONS,
	LOCAL,
	POINTER_DEFAULT,
	MS_UNION,
	OBJECT,
	HELP_STRING,
	DUAL,
	HIDDEN,
	NONEXTENSIBLE,
	ODL,
	OLEAUTOMATION,
}

// Field function returns field attributes.
func (a pAttr) Field() *FieldAttr {
	return &FieldAttr{
		a.FirstIs,
		a.LastIs,
		a.LengthIs,
		a.MinIs,
		a.MaxIs,
		a.SizeIs,
		a.Usage,
		a.Format,
		a.SwitchIs,
		a.Ignore,
		a.Pointer,
		a.Range,
		a.SwitchType,
		a.Safearray,
		a.Layout,
		a.NoSizeLimit,
		a.IsLayout,
		a.NullIf,
	}
}

// Interface function returns interface attributes.
func (a pAttr) Interface() *InterfaceAttr {
	return &InterfaceAttr{
		a.UUID,
		a.Version,
		a.Endpoints,
		a.Exceptions,
		a.Local,
		a.PointerDefault,
		a.MSUnion,
		a.Object,
		a.HelpString,
		a.Dual,
		a.Hidden,
		a.Nonextensible,
		a.ODL,
		a.OLEAutomation,
	}
}

// Type function returns type attributes.
func (a pAttr) Type() *TypeAttr {
	return &TypeAttr{
		a.TransmitAs,
		a.Handle,
		a.SwitchType,
		a.Usage,
		a.Format,
		a.Pointer,
		a.V1Enum,
		a.Range,
		a.DisableConsistencyCheck,
		a.WireMarshal,
		a.Public,
		a.Alias,
		nil,
		nil,
		"",
		a.Pad,
		a.IsLayout,
	}
}

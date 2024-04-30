package midl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"runtime"
	"strings"
)

var json1 = struct {
	Marshal func(v interface{}) ([]byte, error)
}{
	Marshal: func(v interface{}) ([]byte, error) {
		var (
			b = bytes.NewBuffer(nil)
			e = json.NewEncoder(b)
		)

		e.SetIndent("", "  ")
		e.SetEscapeHTML(false)

		if err := e.Encode(v); err != nil {
			return nil, err
		}

		return b.Bytes(), nil
	},
}

func (a *Array) MarshalJSON() ([]byte, error) {

	var upper, lower string

	if a.IsFixed() {
		return json1.Marshal(fmt.Sprintf("[%d]", a.Size()))
	}

	if a.Bound.Lower == -1 {
		lower = "*"
	} else {
		lower = fmt.Sprintf("%d", a.Bound.Lower)
	}

	if a.Bound.Upper == 0 {
		upper = "*"
	} else {
		upper = fmt.Sprintf("%d", a.Bound.Upper)
	}

	if upper == lower && upper == "*" {
		return json.Marshal(fmt.Sprintf("[*]"))
	}

	return json1.Marshal(fmt.Sprintf("[%s:%s]", lower, upper))
}

func (u Usage) String() string {

	var ret []string

	if u.ContextHandle {
		ret = append(ret, "context_handle")
	}

	if u.IsString {
		ret = append(ret, "string")
	}

	return strings.Join(ret, ", ")
}

func (f Format) String() string {

	var ret []string

	if f.NullTerminated {
		ret = append(ret, "null")
	}

	if f.UTF8 {
		ret = append(ret, "utf8")
	}

	if f.Rune {
		ret = append(ret, "rune")
	}

	if f.MultiSize {
		ret = append(ret, "multi_size")
	}

	return strings.Join(ret, ", ")
}

func (a *ComInterfaceAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *ComInterfaceAttr) String() string {

	var ret []string

	if a.Source {
		ret = append(ret, "source")
	}

	if a.Default {
		ret = append(ret, "default")
	}

	return strings.Join(ret, ", ")
}

func (a *LibraryAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *LibraryAttr) String() string {

	var ret []string

	if a.InterfaceAttr != nil {
		ret = append(ret, a.InterfaceAttr.String())
	}

	return strings.Join(ret, ", ")
}

func (a *DispatchInterfaceAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *DispatchInterfaceAttr) String() string {

	var ret []string

	if a.InterfaceAttr != nil {
		ret = append(ret, a.InterfaceAttr.String())
	}

	return strings.Join(ret, ", ")
}

func (a *ComClassAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *ComClassAttr) String() string {

	var ret []string

	if a.InterfaceAttr != nil {
		ret = append(ret, a.InterfaceAttr.String())
	}

	if a.AppObject {
		ret = append(ret, "appobject")
	}

	return strings.Join(ret, ", ")
}

func (a *InterfaceAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *InterfaceAttr) String() string {

	var ret []string

	if a.UUID != nil {
		ret = append(ret, fmt.Sprintf("uuid=%s", a.UUID))
	}

	if a.Version != nil {
		ret = append(ret, fmt.Sprintf("version=%d.%d", a.Version.Major, a.Version.Minor))
	}

	if len(a.Endpoints) > 0 {
		ret = append(ret, fmt.Sprintf("endpoints=[%v]", a.Endpoints))
	}

	if len(a.Exceptions) > 0 {
		ret = append(ret, fmt.Sprintf("exceptions=[%v]", a.Exceptions))
	}

	if a.Local {
		ret = append(ret, "local")
	}

	if a.PointerDefault != 0 {
		ret = append(ret, fmt.Sprintf("pointer_default=%s", a.PointerDefault))
	}

	if a.MSUnion {
		ret = append(ret, "ms_union")
	}

	if a.Object {
		ret = append(ret, "object")
	}

	if a.HelpString != "" {
		ret = append(ret, fmt.Sprintf("{%s}", a.HelpString))
	}

	if a.Dual {
		ret = append(ret, "dual")
	}

	if a.Hidden {
		ret = append(ret, "hidden")
	}

	if a.Nonextensible {
		ret = append(ret, "nonextensible")
	}

	if a.ODL {
		ret = append(ret, "odl")
	}

	if a.OLEAutomation {
		ret = append(ret, "ole_automation")
	}

	return strings.Join(ret, ", ")
}

func (a *Type) MarshalJSON() ([]byte, error) {

	type mType Type

	var list []*mType

	if runtime.Callers(0, make([]uintptr, 100)) == 100 {
		return []byte("null"), nil
	}

	for _, t := range a.flat() {
		list = append(list, (*mType)(t))
	}

	return json1.Marshal(list)
}

func (a *TypeAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *TypeAttr) String() string {

	var ret []string

	if a.TransmitAs != nil {
		ret = append(ret, fmt.Sprintf("transmit_as=%s", a.TransmitAs.Kind))
	}

	if a.Handle {
		ret = append(ret, "handle")
	}

	if a.SwitchType != nil {
		for _, scope := range a.SwitchType.Scopes() {
			typName := scope.Types[0].Name
			if typName == "" {
				typName = scope.Types[0].Kind.String()
			}

			ret = append(ret, fmt.Sprintf("switch_type={%s}(%s)", scope.Attr.String(), typName))
			break
		}
	}

	if usage := a.Usage.String(); usage != "" {
		ret = append(ret, usage)
	}

	if format := a.Format.String(); format != "" {
		ret = append(ret, format)
	}

	if a.Pointer != PointerTypeNone && a.Pointer != PointerTypeRefWeak {
		ret = append(ret, fmt.Sprintf("pointer=%s", a.Pointer))
	}

	if a.V1Enum {
		ret = append(ret, "v1_enum")
	}

	if a.Range != nil {
		ret = append(ret, fmt.Sprintf("range=(%d,%d)", a.Range.Min, a.Range.Max))
	}

	if a.DisableConsistencyCheck {
		ret = append(ret, "disable_consistency_check")
	}

	if a.WireMarshal != "" {
		ret = append(ret, fmt.Sprintf("wire_marshal=%s", a.WireMarshal))
	}
	if a.Public {
		ret = append(ret, "public")
	}
	if a.Alias != "" {
		ret = append(ret, fmt.Sprintf("alias=%s", a.Alias))
	}

	if len(a.Names) > 0 {
		ret = append(ret, fmt.Sprintf("names=%s", strings.Join(a.Names, ";")))
	}

	if len(a.Pointers) > 0 {
		ret = append(ret, fmt.Sprintf("pointers=%s", a.Pointers))
	}

	return strings.Join(ret, ", ")
}

func (a *FieldAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *FieldAttr) String() string {

	var ret []string

	for _, v := range []struct {
		name  string
		value []Expr
	}{
		{"first_is", a.FirstIs},
		{"last_is", a.LastIs},
		{"length_is", a.LengthIs},
		{"min_is", a.MinIs},
		{"max_is", a.MaxIs},
		{"size_is", a.SizeIs},
	} {
		if len(v.value) > 0 {
			var expr []string
			for i, v := range v.value {
				expr = append(expr, fmt.Sprintf("%d:{%s}", i, v))
			}
			ret = append(ret, fmt.Sprintf("%s=(%s)", v.name, strings.Join(expr, ",")))
		}
	}

	if usage := a.Usage.String(); usage != "" {
		ret = append(ret, usage)
	}

	if format := a.Format.String(); format != "" {
		ret = append(ret, format)
	}

	if !a.SwitchIs.Empty() {
		ret = append(ret, fmt.Sprintf("switch_is={%s}", a.SwitchIs))
	}

	if a.Ignore {
		ret = append(ret, "ignore")
	}

	if a.Pointer != PointerTypeNone && a.Pointer != PointerTypeRefWeak {
		ret = append(ret, fmt.Sprintf("pointer=%s", a.Pointer))
	}

	if a.Range != nil {
		ret = append(ret, fmt.Sprintf("range=(%d,%d)", a.Range.Min, a.Range.Max))
	}

	if a.SwitchType != nil {
		for _, scope := range a.SwitchType.Scopes() {
			ret = append(ret, fmt.Sprintf("switch_type={%s}", scope.Attr.String()))
		}
	}

	if a.Safearray != nil {
		if a.Safearray.Kind == TypeRef {
			ret = append(ret, fmt.Sprintf("safearray=%s", a.Safearray.Name))
		} else {
			ret = append(ret, fmt.Sprintf("safearray=%s", a.Safearray.Kind))
		}
	}

	return strings.Join(ret, ", ")
}

func (a *OperationAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (u *UnionCase) String() string {

	ret := []string{}

	for _, label := range u.Labels {
		ret = append(ret, fmt.Sprintf("%v", label))
	}

	return strings.Join(ret, ", ")
}

func (a *OperationAttr) String() string {

	var ret []string

	for _, v := range []struct {
		value bool
		name  string
	}{
		{a.Idempotent, "idempotent"},
		{a.Broadcast, "broadcast"},
		{a.Maybe, "maybe"},
		{a.ReflectDeletions, "reflect_deletions"},
		{a.Usage.String() != "", a.Usage.String()},
		{a.Format.String() != "", a.Format.String()},
		{a.Pointer != PointerTypeNone && a.Pointer != PointerTypeRefWeak, fmt.Sprintf("pointer=%s", a.Pointer)},
		{a.Callback, "callback"},
		{a.PropGet, "propget"},
		{a.PropPut, "propput"},
		{a.PropPutRef, "propputref"},
		{!a.ID.Empty(), fmt.Sprintf("id={%s}", a.ID)},
		{a.Restricted, "restricted"},
		{a.CallAs != "", fmt.Sprintf("call_as=%s", a.CallAs)},
	} {
		if v.value {
			ret = append(ret, v.name)
		}
	}

	return strings.Join(ret, ", ")
}

func (a *ParamAttr) MarshalJSON() ([]byte, error) {
	return json1.Marshal(a.String())
}

func (a *ParamAttr) String() string {

	var ret []string

	if direction := a.Direction.String(); direction != "" {
		ret = append(ret, direction)
	}

	if a.DisableConsistencyCheck {
		ret = append(ret, "disable_consistency_check")
	}

	if !a.IIDIs.Empty() {
		ret = append(ret, fmt.Sprintf("iid_is={%s}", a.IIDIs))
	}

	if a.Retval {
		ret = append(ret, "retval")
	}

	if !a.DefaultValue.Empty() {
		ret = append(ret, fmt.Sprintf("default_value={%s}", a.DefaultValue))
	}

	if a.FieldAttr != nil {
		if field := a.FieldAttr.String(); field != "" {
			ret = append(ret, field)
		}
	}

	if a.Optional {
		ret = append(ret, "optional")
	}

	if a.Annotation != "" {
		ret = append(ret, fmt.Sprintf("{%s}", a.Annotation))
	}

	return strings.Join(ret, ", ")
}

func (d Direction) String() string {

	var ret []string

	if d.In {
		ret = append(ret, "in")
	}

	if d.Out {
		ret = append(ret, "out")
	}

	return strings.Join(ret, ", ")
}

func (p PointerType) String() string {

	var ret string

	switch p {
	case PointerTypePtr:
		ret = "ptr"
	case PointerTypeRef:
		ret = "ref"
	case PointerTypeUnique:
		ret = "unique"
	case PointerTypeRefWeak:
	default:
		ret = "invalid"
	}

	return ret
}

func (t Kind) MarshalJSON() ([]byte, error) {
	return json1.Marshal(t.String())
}

func (t Kind) String() (ret string) {
	switch t {
	case TypeFloat32:
		ret = "float32"
	case TypeFloat64:
		ret = "float64"
	case TypeInt8:
		ret = "int8"
	case TypeInt16:
		ret = "int16"
	case TypeInt32:
		ret = "int32"
	case TypeInt32_64:
		ret = "int32_64"
	case TypeUint32_64:
		ret = "uint32_64"
	case TypeInt64:
		ret = "int64"
	case TypeUint8:
		ret = "uint8"
	case TypeUint16:
		ret = "uint16"
	case TypeUint32:
		ret = "uint32"
	case TypeUint64:
		ret = "uint64"
	case TypeBoolean:
		ret = "bool"
	case TypeVoid:
		ret = "void"
	case TypeHandle:
		ret = "handle"
	case TypeError:
		ret = "error_status_t"
	case TypeCharset:
		ret = "charset"
	case TypeChar:
		ret = "char"
	case TypeUChar:
		ret = "uchar"
	case TypeWChar:
		ret = "wchar"
	case TypeString:
		ret = "string"
	case TypeStruct:
		ret = "struct"
	case TypeUnion:
		ret = "union"
	case TypeCUnion:
		ret = "c_union"
	case TypeFunc:
		ret = "function"
	case TypeArray:
		ret = "array"
	case TypePointer:
		ret = "pointer"
	case TypeEnum:
		ret = "enum"
	case TypePipe:
		ret = "pipe"
	case TypeInterface:
		ret = "interface"
	case TypeDispInterface:
		ret = "dispinterface"
	case TypeRef:
		ret = "ref"
	case TypeAttribute:
		ret = "attribute"
	default:
		ret = "invalid"
	}

	return ret
}

func (c Charset) String() (ret string) {
	switch c {
	case CharsetISO_UCS:
		ret = "ISO_UCS"
	case CharsetISO_Multilingual:
		ret = "ISO_MULTILINGUAL"
	case CharsetISO_Latin_1:
		ret = "ISO_LATIN_1"
	default:
		ret = "NONE"
	}
	return ret
}

func (e Expr) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e Expr) String() string {

	if e.Empty() {
		return ""
	}

	if !e.reqParam {
		switch e.Value.(type) {
		case string:
			return fmt.Sprintf("\"%v\"", e.Value)
		default:
			return fmt.Sprintf("%v", e.Value)
		}
	}

	return e.Expr.String()
}

func (e Expr) Expression(namer ...func(string) string) string {

	if e.Empty() {
		return ""
	}

	if !e.reqParam {
		switch e.Value.(type) {
		case string:
			return fmt.Sprintf("\"%v\"", e.Value)
		default:
			return fmt.Sprintf("%v", e.Value)
		}
	}

	return e.Expr.Expression(namer...)
}

var opToString = map[int]string{
	LOGICAL_AND: "&&",
	LOGICAL_OR:  "||",
	GE:          ">=",
	GT:          ">",
	LE:          "<=",
	LT:          "<",
	EQ:          "=",
	NE:          "!=",
	OR:          "|",
	XOR:         "^",
	AND:         "&",
	LSH:         "<<",
	RSH:         ">>",
	UPLUS:       "+",
	UMINUS:      "-",
	UNEG:        "~",
	UNOT:        "!",
	UMUL:        "*",
	'+':         "+",
	'-':         "-",
	'*':         "*",
	'/':         "/",
}

var noop = func(s string) string { return s }

func (t *ExprTree) Expression(namer ...func(string) string) string {

	nr := noop
	if len(namer) > 0 {
		nr = namer[0]
	}

	var ret string

	if t == nil {
		return ""
	}

	switch t.Op {
	case TERNARY:
		if t.Cond.Op == UMUL || t.Cond.Op == IDENT {
			ret = fmt.Sprintf(t.Lval.Expression(nr))
		} else {
			ret = fmt.Sprintf("(%s?%s:%s)", t.Cond.Expression(nr), t.Lval.Expression(nr), t.Rval.Expression(nr))
		}
	case LOGICAL_AND, LOGICAL_OR, LE, LT, GE, GT, EQ, NE, OR, XOR, AND, LSH, RSH:
		ret = fmt.Sprintf("(%s%s%s)", t.Lval.Expression(nr), opToString[t.Op], t.Rval.Expression(nr))
	case '+', '-', '*', '/', '%':
		ret = fmt.Sprintf("(%s%s%s)", t.Lval.Expression(nr), string(rune(t.Op)), t.Rval.Expression(nr))
	case UNEG:
		ret = fmt.Sprintf("(%suint32(%s))", "^", t.Lval.Expression(nr))
	case UPLUS, UMINUS, UNOT:
		ret = fmt.Sprintf("(%s%s)", opToString[t.Op], t.Lval.Expression(nr))
	case UMUL:
		ret = fmt.Sprintf("%s", t.Lval.Expression(nr))
	case IDENT:
		ret = nr(fmt.Sprintf("%v", t.Value))
	default:
		ret = fmt.Sprintf("%v", t.Value)
	}

	return ret
}

func (t *ExprTree) String() string {

	var ret string

	if t == nil {
		return ""
	}

	switch t.Op {
	case TERNARY:
		ret = fmt.Sprintf("(%s %s %s ?:)", t.Cond, t.Lval, t.Rval)
	case LOGICAL_AND, LOGICAL_OR, LE, LT, GE, GT, EQ, NE, OR, XOR, AND, LSH, RSH:
		ret = fmt.Sprintf("(%s %s %s)", t.Lval, t.Rval, opToString[t.Op])
	case '+', '-', '*', '/', '%':
		ret = fmt.Sprintf("(%s %s %s)", t.Lval, t.Rval, opToString[t.Op])
	case UPLUS, UNEG, UMINUS, UNOT:
		ret = fmt.Sprintf("(%s %s)", t.Lval, opToString[t.Op])
	case UMUL:
		ret = fmt.Sprintf("*%s", t.Lval)
	default:
		ret = fmt.Sprintf("%v", t.Value)
	}

	return ret
}

func (d Dim) String() string {
	ret := []string{fmt.Sprintf("dim:%d", d.Dimension)}

	for _, size := range []struct {
		Name  string
		Value Expr
	}{
		{"size_is", d.SizeIs},
		{"max_is", d.MaxIs},
		{"min_is", d.MinIs},
		{"length_is", d.LengthIs},
		{"first_is", d.FirstIs},
		{"last_is", d.LastIs},
	} {
		if !size.Value.Empty() {
			ret = append(ret, fmt.Sprintf("%s=%s", size.Name, size.Value.Expression()))
		}
	}
	if d.IsString {
		ret = append(ret, "string")
	}
	if d.IsNullTerminated {
		ret = append(ret, "null")
	}
	if d.IsMultiSize {
		ret = append(ret, "multisz")
	}
	if d.IsUTF8 {
		ret = append(ret, "utf8")
	}

	return strings.Join(ret, ",")
}

func InterfaceToExport(iff *Interface) *Type {

	return &Type{
		Kind:  TypeAttribute,
		Attrs: &TypeAttr{Alias: iff.Name},
		Elem: &Type{
			Kind:      TypeInterface,
			Interface: iff,
		},
	}
}

func Hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & ^uint32(0b11<<30) /* unset sign and higher order */)
}

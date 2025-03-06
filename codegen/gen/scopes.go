package gen

import (
	"strings"

	"github.com/oiweiwei/go-msrpc/midl"
)

// Scopes structure is a scope list iterator. Within single scope
// pointer default attributes are applied.
type Scopes struct {
	scopePos, typePos int
	scopes            []*midl.Scope
	prev              *midl.Type
	alias             string
}

// clone function ...
func (i *Scopes) clone() *Scopes {
	return &Scopes{
		scopePos: i.scopePos,
		typePos:  i.typePos,
		scopes:   i.scopes,
		prev:     i.prev,
		alias:    i.alias,
	}
}

func (i *Scopes) IsV1Enum() bool {
	for _, scope := range i.scopes {
		if scope.Attr.V1Enum {
			return true
		}
	}
	return false
}

func (i *Scopes) EnumType() midl.Kind {
	for _, scope := range i.scopes {
		if scope.Attr.V1Enum {
			return scope.Attr.EnumType()
		}
	}
	return i.scopes[0].Attr.EnumType()
}

// NewScope function ...
func NewScopes(scopes []*midl.Scope) *Scopes {
	i := &Scopes{scopes: scopes, typePos: -1}
	return i.Next()
}

func (i *Scopes) Base() *Scopes {

	i = i.clone()

	for i != nil && i.Type().Elem != nil {
		i = i.Next()
	}

	return i
}

// Next function advances the scope to the next type.
func (i *Scopes) Next() *Scopes {

	if i == nil || len(i.scopes) == 0 {
		return nil
	}

	i = i.clone()

	if i.typePos >= 0 {
		i.prev = i.Type()
	}

	if i.typePos++; i.typePos >= len(i.scopes[i.scopePos].Types) {
		i.typePos = 0
		for {
			if i.scopePos++; i.scopePos >= len(i.scopes) {
				return nil
			}
			if len(i.scopes[i.scopePos].Types) != 0 {
				break
			}
		}
	}
	return i
}

// Type function returns the current type.
func (i *Scopes) Type() *midl.Type {
	return i.scopes[i.scopePos].Types[i.typePos].Type
}

func (i *Scopes) Union() *midl.Union {
	return i.Type().Union
}

func (i *Scopes) Struct() *midl.Struct {
	return i.Type().Struct
}

func (i *Scopes) Array() *midl.Array {
	return i.Type().Array
}

func (i *Scopes) Enum() *midl.Enum {
	return i.Type().Enum
}

// ScopedType function return type with scope attributes.
func (i *Scopes) ScopedType() *midl.Type {
	typ := *i.Type()
	typ.Attrs = (&midl.TypeAttr{}).Merge(typ.Attrs).Merge(i.Scope())
	return &typ
}

// Is function ...
func (i *Scopes) Is(kind midl.Kind) bool {
	if i != nil {
		return i.Type().Is(kind)
	}
	return false
}

func (i *Scopes) Kind() midl.Kind {
	return i.Type().Kind
}

// Scope function returns the attributes for the
// current scoped set of types.
func (i *Scopes) Scope() *midl.TypeAttr {
	return i.scopes[i.scopePos].Attr
}

func (i *Scopes) WithAlias(a string) *Scopes {
	i = i.clone()
	i.alias = a
	return i
}

func (i *Scopes) Alias() string {
	if i.alias != "" {
		return i.alias
	}
	return i.scopes[i.scopePos].Attr.Alias
}

// IsBool function returns `true` if type can be rendered
// as a boolean.
func (i *Scopes) IsBool() bool {
	return i.Type().IsPrimitiveType() && strings.HasPrefix(strings.ToUpper(i.Alias()), "BOOL")
}

func (i *Scopes) IsNotFixedArray() bool {
	return i != nil && i.Array() != nil && (!i.Array().IsFixed() || i.Dim().IsString)
}

// Dim function returns the dimenisonal information for
// current scoped type.
func (i *Scopes) Dim() midl.Dim {
	return i.scopes[i.scopePos].Types[i.typePos].Dim
}

// PointerCount function returns the number of pointers following
// the current pointer (+ current pointer itself).
func (i *Scopes) PointerCount() int {
	return i.scopes[i.scopePos].Types[i.typePos].Pointer
}

// IsTopLevelPointer function returns `true` if current scope points
// to a first pointer in a possible chain of pointers.
func (i *Scopes) IsTopLevelPointer() bool {
	for pos := i.typePos - 1; pos >= 0; pos-- {
		if i.scopes[i.scopePos].Types[pos].Is(midl.TypePointer) {
			return false
		}
	}
	return true
}

// IsConformant function returns `true` if array is conformant. That is,
// array is conformant if it's not fixed (and length needed to be determined
// in a runtime).
func (i *Scopes) IsConformant() bool {

	for scopes := i; scopes != nil; scopes = scopes.Next() {
		if scopes.Is(midl.TypeArray) {
			if !scopes.Type().Array.IsFixed() && !scopes.Dim().NoSizeLimit {
				return true
			}
		} else if scopes.Is(midl.TypeStruct) {
			return NewScopes(scopes.StructLastField().Scopes()).IsConformant()
		} else {
			break
		}
	}
	return false
}

// IsVarying function checks if the given array is varying. Varying array is an array
// that has length_is, first_is, last_is, and for some reason it has a string property.
// (or it is a pointer to wchar_t, char, which is a non-null terminated string.)
func (i *Scopes) IsVarying() bool {
	for scopes := i; scopes != nil; scopes = scopes.Next() {
		if scopes.Is(midl.TypeArray) {
			if dim := scopes.Dim(); (dim.IsString && dim.IsNullTerminated) || !dim.LengthIs.Empty() || !dim.FirstIs.Empty() || !dim.LastIs.Empty() {
				return true
			}
		} else {
			break
		}
	}
	return false
}

// IsTopLevelArray function returns true if given array should have a conformance
// data written. It can be top-level if the data type is a pointer, or a fixed array.
// The top level array may contain varying information which needs to be written.
func (i *Scopes) IsTopLevelArray() bool {
	return i.prev != nil && (i.prev.Is(midl.TypePointer) || i.prev.Is(midl.TypeArray) && i.prev.Array.IsFixed())
}

func (i *Scopes) Alignment() int {
	a := i.alignment(false)
	if a != 5 {
		return a
	}

	return a + i.alignment(true)
}

func (i *Scopes) alignment(opaque bool) int {

	switch t := i.ScopedType(); {
	case t.Is(midl.TypeEnum):
		if t.Attrs.V1Enum {
			return 4
		}
		return 2
	case t.IsPrimitiveType():
		return midl.PrimitiveTypeSize(t.Kind)
	case t.Is(midl.TypePointer):
		if opaque {
			return 1
		}
		return 5
	case t.Is(midl.TypeArray):
		a := i.Next().alignment(opaque)
		va := 5
		if opaque {
			va = 1
		}
		if i.IsVarying() {
			if a < va {
				a = va
			}
		}
		return a
	case t.Is(midl.TypeStruct):
		a := 0
		for _, f := range t.Struct.Fields {
			if f.Attrs.Ignore {
				continue
			}
			if fa := NewScopes(f.Scopes()).alignment(opaque); fa > a {
				a = fa
			}
		}
		return a
	case t.Is(midl.TypeUnion) || t.Is(midl.TypeCUnion):
		a := 0
		if i.Scope().SwitchType != nil {
			a = NewScopes(i.Scope().SwitchType.Scopes()).alignment(opaque)
		}
		for _, c := range t.Union.Body {
			for _, f := range c.Arms {
				if cfa := NewScopes(f.Scopes()).alignment(opaque); cfa > a {
					a = cfa
				}
			}
			if t.Is(midl.TypeCUnion) {
				break
			}
		}
		return a
	}
	return 0
}

func (i *Scopes) GoName() string {
	return GoName(i.Alias())
}

func (i *Scopes) StructLastField() *midl.Field {
	if !i.Is(midl.TypeStruct) {
		return nil
	}
	return i.Struct().LastField()
}

func (i *Scopes) IsDeferrable() bool {
	switch {
	case i.Is(midl.TypePointer):
		return true
	case i.Is(midl.TypeArray):
		return i.Next().IsDeferrable()
	case i.Is(midl.TypeStruct):
		for _, field := range i.Struct().Fields {
			if field.Attrs.Ignore {
				continue
			}
			if NewScopes(field.Scopes()).IsDeferrable() {
				return true
			}
		}
	case i.Is(midl.TypeUnion) || i.Is(midl.TypeCUnion):
		for _, cs := range i.Union().Body {
			for _, arm := range cs.Arms {
				if NewScopes(arm.Scopes()).IsDeferrable() {
					return true
				}
			}
		}
	}
	return false
}

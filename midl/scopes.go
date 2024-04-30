package midl

import "math/big"

type Scope struct {
	Attr  *TypeAttr
	Types []*ScopedType
}

type ScopedType struct {
	*Type   `json:"type"`
	Dim     Dim `json:"dim"`
	Pointer int `json:"pointer"`
}

type Dim struct {
	MinIs            Expr
	MaxIs            Expr
	SizeIs           Expr
	FirstIs          Expr
	LastIs           Expr
	LengthIs         Expr
	IsString         bool
	IsNullTerminated bool
	IsMultiSize      bool
	IsUTF8           bool
	Dimension        int
	NoSizeLimit      bool
}

type Size struct {
	MinIs  Expr
	MaxIs  Expr
	SizeIs Expr
}

func (d Size) Empty() bool {
	return d.MinIs.Empty() && d.MaxIs.Empty() && d.SizeIs.Empty()
}

func (d Size) Is() (expr Expr) {
	if expr = d.MaxIs; !expr.Empty() {
		if !d.MinIs.Empty() {
			expr, _ = expr.Sub(d.MinIs)
		}
		expr, _ = expr.Add(NewValue(big.NewInt(1)))
		return expr
	}
	return d.SizeIs
}

func (d Dim) Size() Size {
	return Size{d.MinIs, d.MaxIs, d.SizeIs}
}

func (d Dim) Empty() bool {
	return d.MinIs.Empty() && d.MaxIs.Empty() && d.SizeIs.Empty() &&
		d.FirstIs.Empty() && d.LastIs.Empty() && d.LengthIs.Empty()
}

func (f *Field) Scopes() []*Scope {

	var ok bool

	// first determine if there are any dimensional attributes.
	dims, scopes := f.Attrs.Dim(), f.Type.Scopes()
	if len(scopes) == 0 {
		return scopes
	}

	// count the number of pointers that should point to the
	// conformant array.
	ptrCount := len(dims)
	for i, pointer := len(scopes)-1, 1; i >= 0; i-- {
		for j := len(scopes[i].Types) - 1; j >= 0; j-- {
			if scopes[i].Types[j].Is(TypePointer) {
				scopes[i].Types[j].Pointer = pointer
				pointer++
			} else {
				pointer = 1
			}
			if array := scopes[i].Types[j]; array.Kind == TypeArray && !array.Array.IsFixed() {
				ptrCount--
			}
		}
	}

	// overwrite the attributes.
	// this will override the pointer at the reference site.
	scopes[0].Attr = scopes[0].Attr.Merge(f.Attrs.TypeAttr())

	dimCount := len(dims)

	// assign sizes to proper dimensions.
	// if needed add arrays.
dim_loop:
	for i := 0; i < len(scopes) && len(dims) > 0; i++ {

		for j := 0; j < len(scopes[i].Types); j++ {

			switch typ := scopes[i].Types[j]; {
			case typ.Is(TypeArray):

				if scopes[i].Types[j].Array.IsFixed() {
					continue
				}
				scopes[i].Types[j].Dim = dims[len(dims)-dimCount]

			case typ.Is(TypePointer):

				if ptrCount <= 0 {
					// ignore this pointer since there is an array
					// to attach the sizing information.
					continue
				}
				ptrCount--

				if dims[len(dims)-dimCount].Empty() {
					// skip over the empty dimension. break the switch statement.
					break
				}

				// find next type in this or next scope.
				if i, j, ok = nextType(scopes, i, j); !ok {
					break dim_loop
				}

				// https://learn.microsoft.com/en-us/windows/win32/Midl/size-is
				// insert the absent array right after the pointer.

				extendScopes(scopes, i, j)[i].Types[j] = &ScopedType{
					Type: &Type{
						Kind:  TypeArray,
						Array: &Array{Bound: ArrayBound{Lower: -1}},
						Elem:  scopes[i].Types[j+1].Type,
					},
					Dim: dims[len(dims)-dimCount],
				}

				j++

			default:
				continue
			}

			if dimCount--; dimCount <= 0 {
				break dim_loop
			}
		}
	}

	isString := f.IsString()
	isNullTerminated := isString
	isChar := false

	// go in reverse direction and determine where to set the string.
string_loop:
	for i := len(scopes) - 1; i >= 0; i-- {
		for j := len(scopes[i].Types) - 1; j >= 0; j-- {
			switch typ := scopes[i].Types[j]; {
			case typ.Is(TypeChar) || typ.Is(TypeWChar):
				// keep track of isChar, to avoid rendering pointer to char as a uint8.
				isString, isChar = true, typ.Is(TypeChar)
			case typ.Is(TypePointer):

				if !f.IsString() && (scopes[i].Attr.Pointer == PointerTypeRef || scopes[i].Attr.Pointer == PointerTypeRefWeak) {
					break string_loop
				}

				if !isString {
					break string_loop
				}

				// insert string pointing to the next element.
				if i, j, ok = nextType(scopes, i, j); !ok {
					break string_loop
				}

				extendScopes(scopes, i, j)[i].Types[j] = &ScopedType{
					Type: &Type{
						Kind:  TypeArray,
						Array: &Array{Bound: ArrayBound{Lower: -1}},
						Elem:  scopes[i].Types[j+1].Type,
					},
					Dim: Dim{IsString: true, IsNullTerminated: isNullTerminated},
				}

				break string_loop

			case typ.Is(TypeArray):

				if !isString {
					break string_loop
				}
				// break loop if we determined the non-null terminated string of
				// char (which is byte buffer, or fixed array).
				if !f.IsString() && (isChar || typ.Array.IsFixed()) {
					break string_loop
				}

				// set is_string
				scopes[i].Types[j].Dim.IsString = true
				scopes[i].Types[j].Dim.IsNullTerminated = isNullTerminated
				break string_loop
			}
		}
	}

	// put proper dimension number.
	dimension := 0
	for _, scope := range scopes {
		for _, typ := range scope.Types {
			if typ.Kind == TypeArray {
				typ.Dim.Dimension = dimension
				dimension++
			} else {
				dimension = 0
			}
		}
	}

	// put proper pointer dereference number.
	// (go in reverse order).
	// so that pointer(**ptr) -> 2, pointer(*ptr) -> 1
	for i, pointer := len(scopes)-1, 1; i >= 0; i-- {
		for j := len(scopes[i].Types) - 1; j >= 0; j-- {
			if scopes[i].Types[j].Is(TypePointer) {
				scopes[i].Types[j].Pointer = pointer
				pointer++
			} else {
				pointer = 1
			}
		}
	}

	return scopes
}

func nextType(scopes []*Scope, i, j int) (int, int, bool) {
	j++
	for i < len(scopes) && j >= len(scopes[i].Types) {
		i++
		j = 0
	}
	return i, j, i < len(scopes)
}

func extendScopes(scopes []*Scope, i, j int) []*Scope {
	scopes[i].Types = append(scopes[i].Types[:j+1], scopes[i].Types[j:]...)
	return scopes

}

// Scopes function returns the list types grouped by scope.
// Within a single scope pointer defaults are applied.
func (t *Type) Scopes() []*Scope {

	scopes := []*Scope{}

	var curScope *Scope

	for t := t; t != nil; t = t.Elem {
		if t.Is(TypeAttribute) {
			attr := t.Attrs.clone()
			for t = t.Elem; t.Is(TypeAttribute); t = t.Elem {
				attr = t.Attrs.Merge(attr)
			}
			if curScope != nil {
				scopes = append(scopes, curScope)
			}
			curScope = &Scope{Attr: attr}
		}

		if curScope == nil {
			curScope = &Scope{Attr: &TypeAttr{}}
		}

		curScope.Types = append(curScope.Types, &ScopedType{Type: t})
	}

	if curScope != nil {
		scopes = append(scopes, curScope)
	}

	return scopes
}

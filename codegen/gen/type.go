package gen

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/oiweiwei/go-msrpc/codegen/doc"
	"github.com/oiweiwei/go-msrpc/midl"
)

type TypeGenerator struct {
	*Generator
	*Scopes
	GoTypeName string
	Doc        *doc.TypeDoc
	OrigType   *midl.Type
}

func (p *Generator) NewTypeGenerator(ctx context.Context, typ *midl.Type) *TypeGenerator {
	scopes := NewScopes(typ.Scopes())
	docs, ok := p.Doc.Type(scopes.Alias())
	if !ok {
		docs, _ = p.Doc.Type(scopes.Scope().Parent)
	}
	return &TypeGenerator{
		Generator:  p,
		Scopes:     scopes,
		GoTypeName: scopes.GoName(),
		Doc:        docs,
		OrigType:   typ,
	}
}

func (p *TypeGenerator) Skip() bool {
	return p.Scopes == nil || p.Scope().Alias == ""
}

func (p *TypeGenerator) GenEnum(ctx context.Context) {

	p.P("//", p.GoTypeName, "type", "represents", RPCName(p.Alias()), "RPC", "enumeration.")
	if p.Doc != nil {
		p.P("//")
		p.GenComment(ctx, p.Doc.Doc)
	}

	if p.Enum().Is32 {
		// accomodate all possible values.
		p.P("type", p.GoTypeName, "uint32")
	} else {
		p.P("type", p.GoTypeName, p.EnumType())
	}
	p.P()
	p.P("var", "(")
	for _, enum := range p.Enum().Elems {
		if doc, ok := p.Doc.Field(enum.Value); ok {
			p.GenComment(ctx, doc.Doc)
		}
		p.P(GoMergeNames(p.GoTypeName, GoName(enum.Value)), p.GoTypeName, "=", strconv.Itoa(enum.ID))
	}
	p.P(")")
	p.P()
	p.Block("func", "(o "+p.GoTypeName+")", "String()", "string", func() {
		p.Block("switch o", func() {
			for _, enum := range p.Enum().Elems {
				n := GoMergeNames(p.GoTypeName, GoName(enum.Value))
				p.P("case", n, ":")
				p.P("return", p.Q(n))
			}
		})
		p.P("return", p.Q("Invalid"))
	})
}

func (p *TypeGenerator) GenCUnion(ctx context.Context) {

	p.Structure(p.GoTypeName, func() {
		for _, cases := range p.Union().Body {
			for _, field := range cases.Arms {
				p.GenStructField(ctx, field)
			}
		}
	})

	for _, cases := range p.Union().Body {
		p.GenSubTypes(ctx, cases.Arms)
	}
}

func (p *TypeGenerator) GenStructField(ctx context.Context, field *midl.Field) {
	if field.IsHandle() {
		return
	}
	if doc, ok := p.Doc.Field(field.Name); ok {
		p.GenComment(ctx, doc.Doc)
	}
	notes := ""
	if !field.IsString() && p.GoFieldTypeName(ctx, p.Scope(), field) == "string" {
		// notes = `// field is not declared as a string` + p.GenStructFieldComment(ctx, field)
	}
	p.P(p.GoFieldName(field), p.GoFieldTypeName(ctx, p.Scope(), field), p.GenTag(ctx, field), notes)
	for _, field := range field.Attrs.Layout {
		p.GenStructField(ctx, field)
	}
}

func (p *TypeGenerator) GenStructFieldComment(ctx context.Context, field *midl.Field, defaultPtr ...string) string {
	ret := ""
	for i, scope := range field.Scopes() {
		ret += fmt.Sprintf("(%d:", i+1)
		attr := scope.Attr.String()
		if len(defaultPtr) > 0 && len(scope.Types) > 0 && scope.Types[0].Is(midl.TypePointer) && scope.Attr.Pointer == midl.PointerTypeNone {
			if attr != "" {
				attr += ","
			}
			attr += "pointer=" + defaultPtr[0]
		}

		if attr != "" {
			ret += "{" + attr + "}"
		}
		for _, typ := range scope.Types {
			switch {
			case typ.Is(midl.TypePointer):
				ret += "*"
				ret += fmt.Sprintf("(%d)", typ.Pointer)
			case typ.Is(midl.TypeArray):
				ret += "["
				if typ.Array.IsFixed() {
					ret += fmt.Sprintf("%d", typ.Array.Size())
				} else {
					ret += typ.Dim.String()
				}
				ret += "]"
			default:
				ret += "(" + typ.Kind.String() + ")"
			}
		}
		ret += ")"
	}

	return ret
}

func (p *TypeGenerator) GenStructPreparePayloadAfterHook(ctx context.Context) {
	p.If("hook, ok := (interface{})(o).(interface { AfterPreparePayload(context.Context) error }); ok", func() {
		p.CheckErr(p.B("hook.AfterPreparePayload", "ctx"))
	})
}

func (p *TypeGenerator) GenStructPreparePayloadFields(ctx context.Context, fields []*midl.Field) {

	ranged := make([]*midl.Field, 0)
	defaulted := make([]*midl.Field, 0)

	for _, field := range fields {

		defaulted = append(defaulted, field)

		skip := p.GoFieldName(field) == "_"

		if field.Attrs.Range != nil && !skip {
			ranged = append(ranged, field)
		}

		if field.Attrs.SizeAttr() == nil && field.Attrs.LengthIs == nil {
			continue
		}

		scopes := NewScopes(field.Scopes())
		for ; scopes != nil; scopes = scopes.Next() {
			if scopes.Is(midl.TypeArray) {
				break
			}
		}

		if scopes := scopes; field.Attrs.SizeAttr() != nil {
			for ; scopes != nil; scopes = scopes.Next() {
				if sz := scopes.Dim().Size(); !sz.Empty() {
					if !skip {
						p.GenStructPreparePayloadSizeExpr(ctx, field, scopes, sz.Is())
					}
					// check minimum value borders.
					if to, val, ok := p.GetMinVarValue(ctx, sz.Is()); ok {
						p.If(to.Expression(p.EO), "<", val, func() {
							p.P(to.Expression(p.EO), "=", val)
						})
					}
					break
				}
			}
		}

		if scopes := scopes; field.Attrs.LengthIs != nil {
			for ; scopes != nil; scopes = scopes.Next() {
				if sz := scopes.Dim().LengthIs; !sz.Empty() {
					if !skip {
						p.GenStructPreparePayloadSizeExpr(ctx, field, scopes, sz)
					}
					if to, val, ok := p.GetMinVarValue(ctx, sz); ok {
						p.If(to.Expression(p.EO), "<", val, func() {
							p.P(to.Expression(p.EO), "=", val)
						})
					}
					break
				}
			}
		}
	}

ranged_loop:
	for _, field := range ranged {

		rng := field.Attrs.Range
		fL := p.O(p.GoFieldName(field))
		fT := p.GoFieldTypeName(ctx, p.Scope(), field)

		// process range statement for array.
		for scopes := NewScopes(field.Scopes()); scopes != nil; scopes = scopes.Next() {
			if scopes.Is(midl.TypeArray) {
				p.If(p.Len(fL), ">", p.B("int", rng.Max), func() {
					p.P(`return fmt.Errorf("` + p.GoFieldName(field) + ` is out of range")`)
				})
				continue ranged_loop
			}
		}

		// process range statement for primitive type.
		if field.Type.Base().Kind.Signed() || rng.Min != 0 {
			p.If(fL, "<", p.B(fT, rng.Min), "||", fL, ">", p.B(fT, rng.Max), func() {
				p.P(`return fmt.Errorf("` + p.GoFieldName(field) + ` is out of range")`)
			})
		} else if rng.Min == 0 {
			p.If(fL, ">", p.B(fT, rng.Max), func() {
				p.P(`return fmt.Errorf("` + p.GoFieldName(field) + ` is out of range")`)
			})
		}
	}

	for _, field := range defaulted {
		if !field.DefaultValue.IsZero() {
			p.If(p.EO(field.Name), "==", p.GoTypeZeroValue(ctx, p.Scope(), field, NewScopes(field.Scopes())), func() {
				expr := field.DefaultValue
				p.P(p.EO(field.Name), "=", p.B(
					p.GoFieldTypeName(ctx, p.Scope(), field),
					p.GenExpr(ctx, expr, p.LookupExprField(ctx, expr), ""),
				))
			})
		}
	}
}

func (p *TypeGenerator) LookupExprField(ctx context.Context, expr midl.Expr) *midl.Field {

	if p.Struct() == nil {
		return nil
	}

	i, ok := expr.ResolveTo(midl.Expr{})
	if !ok {
		return nil
	}

	for _, field := range p.Struct().Fields {
		if i.Value.(string) == field.Name {
			return field
		}
	}

	return nil
}

func (p *TypeGenerator) GenExpr(ctx context.Context, expr midl.Expr, field *midl.Field, typ string, n ...string) string {

	isBool := false
	if field != nil {
		isBool = NewScopes(field.Scopes()).IsBool()
	}

	if field == nil || (!isBool && (expr.Expr.Op != midl.TERNARY || (expr.Expr.Cond.Op == midl.UMUL || expr.Expr.Cond.Op == midl.IDENT))) {

		out := expr.Expression(p.EO)
		if out == "o._" {
			return "0"
		}
		return out
	}

	varName := "_expr" + field.Name + strings.Join(n, "")

	if typ == "" {
		typ = p.GoFieldTypeName(ctx, p.Scope(), field)
	}

	p.P(varName, ":=", typ+"(0)")

	var cond, lval, rval *midl.ExprTree
	if isBool {
		cond = &midl.ExprTree{Op: midl.IDENT, Value: field.Name}
		lval = &midl.ExprTree{Value: big.NewInt(1)}
		rval = &midl.ExprTree{Value: big.NewInt(0)}
	} else {
		cond, lval, rval = expr.Expr.Cond, expr.Expr.Lval, expr.Expr.Rval
	}

	chk := ""
	if !isBool && !midl.IsBoolOp(expr.Expr.Cond.Op) {
		chk = "!= 0"
	}

	p.If(cond.Expression(p.EO), chk, func() {
		p.P(varName, "=", p.B(typ, lval.Expression(p.EO)))
	}, p.Else(func() {
		p.P(varName, "=", p.B(typ, rval.Expression(p.EO)))
	}))

	return varName
}

func (p *TypeGenerator) GenStructPreparePayload(ctx context.Context) {

	p.P("")
	p.P("func", "(o *"+p.GoTypeName+")", p.XXX()+"PreparePayload(ctx context.Context)", "error", "{")

	p.GenStructPreparePayloadFields(ctx, p.Struct().Fields)
	p.GenStructPreparePayloadAfterHook(ctx)

	p.P("return", "nil")
	p.P("}")
}

func (p *TypeGenerator) GenStructPreparePayloadSizeExpr(ctx context.Context, field *midl.Field, scopes *Scopes, sz midl.Expr) {

	n := p.GoFieldName(field)

	lenIdent := midl.NewIdent(p.DataLen(ctx, field, NewScopes(field.Scopes()), p.O(n), ""))

	expr, ok := sz.Resolve(lenIdent)
	if !ok {
		if expr, ok = sz.ResolveTo(lenIdent); !ok {
			p.P("//", "cannot evaluate expression", sz.String())
			return
		}
	}

	target := fmt.Sprintf("%s", expr.Value)
	for _, szField := range p.Struct().Fields {
		if szField.Name != target {
			continue
		}
		szN := GoName(szField.Name)
		szT := p.GoFieldTypeName(ctx, p.Scope(), szField)
		// ... if o.Data != nil && o.DataLen == 0 ...
		p.If(p.O(n), "!=", p.GoTypeZeroValue(ctx, p.Scope(), field, scopes), "&&", p.O(szN), "==", "0", func() {
			// TODO: handle complex expressions.
			p.P(p.O(szN), "=", p.B(szT, expr.Expression()))
			// check variable borders.
			if to, val, ok := p.GetMinVarValue(ctx, expr); ok {
				p.If(to.Expression(), "<", val, func() {
					p.P(p.O(szN), "=", p.B(szT, "0"))
				})
			}
		})
	}
}

func (p *TypeGenerator) GenSubTypes(ctx context.Context, fields []*midl.Field) {

	for _, field := range fields {
		p.GenSubType(ctx, field)
		for _, field := range field.Attrs.Layout {
			p.GenSubType(ctx, field)
		}
	}
}

func (p *TypeGenerator) GenSubType(ctx context.Context, field *midl.Field) {

	switch t := field.Type.Base(); {
	case t.Is(midl.TypeStruct) || t.Is(midl.TypeUnion) || t.Is(midl.TypeCUnion):

		if LookupName(ctx, NewScopes(field.Scopes()).Base()) != "" {
			return
		}

		n := p.EmbeddedTypeName(ctx, p.Scope(), field)

		if strings.Contains(n, p.XXX()) {
			panic("invalid type " + n)
		}

		parent := p.Scope().Parent
		if parent == "" {
			parent = p.Alias()
		}

		p.GenType(ctx, &midl.Type{
			Kind:  midl.TypeAttribute,
			Attrs: field.Attrs.TypeAttr().Merge(&midl.TypeAttr{Alias: Escape(n), Parent: parent}),
			Elem:  t,
		})
	}
}

func (p *TypeGenerator) GenInterface(ctx context.Context) {
	tt := *midl.LookupType("MInterfacePointer")
	tt.Attrs = tt.Attrs.Merge(p.OrigType.Attrs)
	p.NewTypeGenerator(ctx, &tt).GenStruct(ctx)
}

func (p *TypeGenerator) GenStruct(ctx context.Context) {

	if p.Scope().Parent != "" {
		p.P("//", p.GoTypeName, "structure", "represents", RPCName(p.Scope().Parent), "structure", "anonymous", "member.")
	} else {
		p.P("//", p.GoTypeName, "structure", "represents", RPCName(p.Alias()), "RPC", "structure.")
	}
	if p.Doc != nil {
		p.P("//")
		p.GenComment(ctx, p.Doc.Doc)
	}

	if n := p.Scope().Names; len(n) > 0 {
		names := p.GoScopeTypeName(ctx, p.Scope(), nil, p.Scopes.WithAlias(n[0]), true)
		p.P("type", p.GoTypeName, names)
		p.P()
		p.P("func", "(o *"+p.GoTypeName+")", GoName(n[0]), "()", "*"+names, "{", "return", p.B("(*"+names+")", "o"), "}")
	} else {
		p.Structure(p.GoTypeName, func() {
			for _, field := range p.Struct().Fields {
				p.GenStructField(ctx, field)
			}
		})
	}

	p.GenStructPreparePayload(ctx)
	p.GenStructNDRSizeInfo(ctx)
	p.GenLayout(ctx)
	p.GenStructMarshalNDR(ctx)
	p.GenStructUnmarshalNDR(ctx)
	p.GenSubTypes(ctx, p.Struct().Fields)

}

func (p *TypeGenerator) GenLayout(ctx context.Context) {
	if p.Scope().IsLayout {
		p.P("func", "(o *"+p.GoTypeName+")", "NDRLayout()", "{}")
	}
}

func (p *TypeGenerator) GenStructNDRSizeInfo(ctx context.Context) {

	if !p.IsConformant() {
		return
	}

	p.P()
	p.P("func", "(o *"+p.GoTypeName+")", "NDRSizeInfo() []uint64 {")

	last := p.Struct().LastField()
	if scopes := NewScopes(last.Scopes()); !scopes.Is(midl.TypeArray) {
		p.P("return", p.O(p.GoFieldName(last))+".NDRSizeInfo()")
	} else {

		var dims []string

		for i := 0; scopes.IsNotFixedArray(); scopes, i = scopes.Next(), i+1 {
			dims = append(dims, p.GenArraySizeVar(ctx, last, i, scopes, "Size", scopes.Dim().Size().Is()))
		}
		p.P("return []uint64{")
		for _, dim := range dims {
			p.P(dim, ",")
		}
		p.P("}")
	}

	p.P("}")
}

func (p *TypeGenerator) GenStructMarshalNDRSizePreamble(ctx context.Context, field *midl.Field) {
	if !p.IsConformant() {
		return
	}
	p.P("sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)")
	p.If("!ok", func() {
		p.P("sizeInfo", "=", p.O("NDRSizeInfo()"))
		p.Range("sz1", "sizeInfo", func() {
			p.GenWriteSize(ctx, p.IVar("sizeInfo", "sz1"))
		})
		p.P("ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)")
	})
}

func (p *TypeGenerator) GenStructMarshalNDR(ctx context.Context) {

	trailingField := len(p.Struct().Fields) - 1
	if p.IsConformant() || p.IsVarying() {
		trailingField--
	}

	p.Block("func", "(o *"+p.GoTypeName+")", "MarshalNDR(ctx context.Context, w ndr.Writer)", "error", func() {
		// prepare payload.
		p.CheckErr("o." + p.XXX() + "PreparePayload(ctx)")
		// write size information.
		p.GenStructMarshalNDRSizePreamble(ctx, p.StructLastField())
		// align the structure.
		p.GenDoAlignmentMarshalNDR(ctx, p.Alignment())
		// marshal fields.
		for i, field := range p.Struct().Fields {
			// marshal field.
			p.GenFieldMarshalNDR(ctx, field, NewScopes(field.Scopes()))
			// add trailing gap.
			if i == trailingField {
				if f := NewScopes(field.Scopes()); f.alignment(false) != p.alignment(false) {
					p.GenDoTrailingGapMarshalNDR(ctx, p.Alignment())
				}
			}
		}
		if p.Scope().Pad != 0 {
			p.P("//", "pad", p.Scope().Pad)
			p.GenDoAlignmentMarshalNDR(ctx, int(p.Scope().Pad))
		}
		p.P("return nil")
	})

}

func (p *TypeGenerator) GenFieldMarshalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {

	if scopes == nil || field.Attrs.Ignore {
		// end.
		return
	}

	// reserved field.
	if p.GoFieldName(field) == "_" {
		p.P("//", "reserved", field.Name)
		p.GenZeroFieldMarshalNDR(WithVarName(ctx, ZeroLen), field, scopes, index...)
		return
	}

	// get variable name.
	name := p.VarName(ctx, index...)
	if name == "" {
		name = p.O(p.IVar(p.GoFieldName(field), index...))
	}

	switch {

	case scopes.Type().Is(midl.TypeHandle):

		p.P("//", "skip", "[handle]", field.Name)

	case scopes.Type().IsPrimitiveType():

		// marshal primitive type.

		if scopes.IsBool() && midl.PrimitiveTypeSize(scopes.Kind()) > 1 {
			typeName := GoPrimitiveTypeName(scopes.Kind())
			p.If("!", name, func() {
				p.CheckErr(p.B("w.WriteData", p.B(typeName, "0")))
			}, p.Else(func() {
				p.CheckErr(p.B("w.WriteData", p.B(typeName, "1")))
			}))
		} else if scopes.Is(midl.TypeInt32_64) {
			p.CheckErr(p.B("w.WriteData", p.B("ndr.Int3264", name)))
		} else if scopes.Is(midl.TypeUint32_64) {
			p.CheckErr(p.B("w.WriteData", p.B("ndr.Uint3264", name)))
		} else if field.Attrs.Format.Rune {
			p.CheckErr(p.B("w.WriteData", p.B("uint16", name)))
		} else {
			p.CheckErr(p.B("w.WriteData", name))
		}

	case scopes.Is(midl.TypeEnum):

		// marshal enum.

		p.CheckErr(p.B("w.WriteEnum", p.B(scopes.EnumType(), name)))

	case scopes.Is(midl.TypeUnion):

		// marshal union.

		switchVar := "_sw" + p.ToVar(name)
		if !scopes.Union().IsEncapsulated() {
			p.GenerateSwitchIs(ctx, field, scopes, switchVar)
		}

		p.If(name, "!=", "nil", func() {
			if !scopes.Union().IsEncapsulated() {
				p.CheckErr(p.B(name+".MarshalUnionNDR", "ctx", "w", switchVar))
			} else {
				p.CheckErr(p.B(name+".MarshalNDR", "ctx", "w"))
			}
		}, p.Else(func() {
			p.GenZeroUnionFieldMarshalNDR(ctx, field, scopes, index...)
		}))

	case scopes.Is(midl.TypeStruct), scopes.Is(midl.TypeInterface):

		// marshal structure.

		p.If(name, "!=", "nil", func() {
			p.CheckErr(p.B(name+".MarshalNDR", "ctx", "w"))
		}, p.Else(func() {
			p.GenZeroStructFieldMarshalNDR(ctx, field, scopes, index...)
		}))

	case scopes.Is(midl.TypePointer):

		// marshal pointer.

		sizeChk := ""
		if next := scopes.Next(); next.Is(midl.TypeArray) && !next.Dim().Size().Empty() {
			// immitate memory allocation for conformance array.
			sizeChk = fmt.Sprintf("|| %s > 0", p.GenExpr(ctx, next.Dim().Size().Is(), p.LookupExprField(ctx, next.Dim().Size().Is()), ""))
		}

		if next := scopes.Next(); next.Type().IsPrimitiveType() {
			if field.Attrs.DefaultNull != nil {
				nullChk := ""
				if len(field.Attrs.DefaultNull) > 0 {
					for _, expr := range field.Attrs.DefaultNull {
						chk := p.GenExpr(ctx, expr, p.LookupExprField(ctx, expr), "")
						if nullChk == "" {
							nullChk = fmt.Sprintf("!%s", chk)
						} else {
							nullChk = fmt.Sprintf("%s || !%s", nullChk, chk)
						}
					}
				} else {
					nullChk = fmt.Sprintf("%s != %v", name, p.GoTypeZeroValue(ctx, p.Scope(), field, scopes))
				}

				p.If(nullChk, func() {
					fN := "_ptr_" + field.Name
					p.P(fN, ":=", "ndr.MarshalNDRFunc", "(", "func(ctx context.Context, w ndr.Writer) error {")
					p.GenFieldMarshalNDR(ctx, field, scopes.Next(), index...)
					p.P("return nil")
					p.P("})")
					p.CheckErr(p.B("w.WritePointer", p.Amp(name), fN))
				}, p.Else(func() {
					p.GenZeroPointerFieldMarshalNDR(ctx, field, scopes, index...)
				}))
			} else {
				p.P("//", "XXX", "pointer to primitive type, default behavior is to write non-null pointer.")
				p.P("//", "if this behavior is not desired, use goext_default_null([cond]) attribute.")
				fN := "_ptr_" + field.Name
				p.P(fN, ":=", "ndr.MarshalNDRFunc", "(", "func(ctx context.Context, w ndr.Writer) error {")
				p.GenFieldMarshalNDR(ctx, field, scopes.Next(), index...)
				p.P("return nil")
				p.P("})")
				p.CheckErr(p.B("w.WritePointer", p.Amp(name), fN))
			}
		} else {
			p.If(name, "!=", p.GoTypeZeroValue(ctx, p.Scope(), field, scopes), sizeChk, func() {
				fN := "_ptr_" + field.Name
				p.P(fN, ":=", "ndr.MarshalNDRFunc", "(", "func(ctx context.Context, w ndr.Writer) error {")
				p.GenFieldMarshalNDR(ctx, field, scopes.Next(), index...)
				p.P("return nil")
				p.P("})")
				p.CheckErr(p.B("w.WritePointer", p.Amp(name), fN))
			}, p.Else(func() {
				p.GenZeroPointerFieldMarshalNDR(ctx, field, scopes, index...)
			}))
		}

	case scopes.Is(midl.TypeArray):

		// marshal array.

		if !p.GenMarshalSizeInfo(ctx, field, scopes, index...) {
			break
		}

		idx := p.Var("i", len(index)+1)

		if scopes.Dim().IsString {
			name = p.GenWriteStringBuffer(ctx, scopes, field, index...)
			ctx, index = WithVarName(ctx, name), nil
		}

		if scopes.Dim().NoSizeLimit {
			p.Range(idx, name, func() {
				p.P(idx, ":=", idx)
				p.GenFieldMarshalNDR(ctx, field, scopes.Next(), append(index, idx)...)
			})
			break
		}

		p.Range(idx, name, func() {
			// iterate over the array.
			p.P(idx, ":=", idx)
			if scopes.Array().IsFixed() && !scopes.Dim().IsString {
				// break if index exceeds the fixed-array limit.
				p.If(p.B("uint64", idx), ">=", scopes.Array().Size(), func() {
					p.P("break")
				})
			} else {
				// break if index exceeds the conformance/varyingness limit.
				p.If(p.B("uint64", idx), ">=", p.IVar("sizeInfo", scopes.Dim().Dimension), func() {
					p.P("break")
				})
			}
			p.GenFieldMarshalNDR(ctx, field, scopes.Next(), append(index, idx)...)
		})

		p.GenZeroFieldMarshalNDR(ctx, field, scopes, index...)

	default:
		p.P("//", "FIXME", "unknown type", field.Name)
	}
}

func (p *TypeGenerator) GenMarshalSizeInfo(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) bool {

	// get variable name.
	name := p.VarName(ctx, index...)
	if name == "" {
		name = p.O(p.IVar(p.GoFieldName(field), index...))
	}

	isConformant, isVarying := scopes.IsConformant(), scopes.IsVarying()

	if field.Position == 0 || scopes.IsTopLevelArray() {

		// for non-top level array size information is handled via NDRSizeInfo.

		if dim := scopes.Dim(); dim.IsString && !dim.NoSizeLimit && dim.Size().Empty() && dim.LengthIs.Empty() {

			// top-level string.

			if name == ZeroLen {
				name = `""`
			}

			switch scopes.Next().Kind() {
			case midl.TypeWChar, midl.TypeUint16, midl.TypeInt16:
				if dim.IsNullTerminated {
					p.CheckErr(p.B("ndr.WriteUTF16NString", "ctx", "w", name))
				} else {
					p.CheckErr(p.B("ndr.WriteUTF16String", "ctx", "w", name))
				}
			case midl.TypeUChar, midl.TypeChar, midl.TypeUint8:
				if dim.IsNullTerminated {
					p.CheckErr(p.B("ndr.WriteCharNString", "ctx", "w", name))
				} else {
					p.CheckErr(p.B("ndr.WriteCharString", "ctx", "w", name))
				}
			}

			return false

			// else, string length should be retrieved from external parameters
			// size is, length is, and so on.
		}

		if isConformant {

			dims := []string{}
			for scopes, i := scopes, len(index); scopes.IsNotFixedArray(); scopes, i = scopes.Next(), i+1 {
				dims = append(dims, p.GenArraySizeVar(ctx, field, i, scopes, "Size", scopes.Dim().Size().Is()))
				// write max_count.
				p.GenWriteSize(ctx, dims[len(dims)-1])
			}

			p.Block("sizeInfo := []uint64", func() {
				for _, dim := range dims {
					p.P(dim, ",")
				}
			})
		}
	}

	if isVarying {

		if !isConformant {
			p.Block("sizeInfo := []uint64", func() {
				for scopes, i := scopes, len(index); scopes.IsNotFixedArray(); scopes, i = scopes.Next(), i+1 {
					p.P(0, ",")
				}
			})
		}

		for scopes, i := scopes, len(index); scopes.IsNotFixedArray(); scopes, i = scopes.Next(), i+1 {
			dim := p.GenArraySizeVar(ctx, field, i, scopes, "Length", scopes.Dim().LengthIs)
			if !isConformant {
				p.P(p.IVar("sizeInfo", scopes.Dim().Dimension), "=", dim)
			} else {
				p.If(dim, ">", p.IVar("sizeInfo", scopes.Dim().Dimension), func() {
					p.P(dim, "=", p.IVar("sizeInfo", scopes.Dim().Dimension))
				}, p.Else(func() {
					// reset the size information back for padding and etc.
					p.P(p.IVar("sizeInfo", scopes.Dim().Dimension), "=", dim)
				}))
			}
			// write offset.
			p.GenWriteSize(ctx, 0)
			// write actual_count.
			p.GenWriteSize(ctx, dim)
		}
	}

	return true
}

func (p *TypeGenerator) GetMinVarValue(ctx context.Context, sizeExpr midl.Expr) (midl.Expr, int64, bool) {
	expr, ok := sizeExpr.Eval(midl.NewArgs(big.NewInt(0)))
	if !ok {
		return midl.Expr{}, 0, false
	}
	val, ok := expr.Int64()
	if !ok || val >= 0 {
		return midl.Expr{}, 0, false
	}
	to, ok := sizeExpr.ResolveTo(midl.Expr{})
	if !ok {
		return midl.Expr{}, 0, false
	}

	return to, -val, true
}

func (p *TypeGenerator) GenArraySizeVar(ctx context.Context, field *midl.Field, i int, scopes *Scopes, name string, sizeExpr midl.Expr) string {

	if !scopes.IsNotFixedArray() {
		return ""
	}

	dimVar := p.Var("dim"+name, i+1)

	if !sizeExpr.Empty() {
		p.P(dimVar, ":=", p.B("uint64", p.GenExpr(ctx, sizeExpr, p.LookupExprField(ctx, sizeExpr), "uint64")))
		// check minimum value borders.
		if to, val, ok := p.GetMinVarValue(ctx, sizeExpr); ok {
			p.If(p.GenExpr(ctx, to, p.LookupExprField(ctx, to), "") /*to.Expression(p.EO) */, "<", val, func() {
				p.P(dimVar, "=", "uint64(0)")
			})
		}
		return dimVar
	}

	if i == 0 {
		// short-path.
		p.P(dimVar, ":=", p.DataLen(ctx, field, scopes, p.O(p.GoFieldName(field))))
		return dimVar
	}

	// calculate size on the fly.
	p.P(dimVar, ":=", "uint64(0)")
	p.GenRange(ctx, p.O(p.GoFieldName(field)), 0, i, func(ctx context.Context, n string) {
		lenVar := p.Var("len" + Title(dimVar))
		p.If(lenVar, ":=", p.DataLen(ctx, field, scopes, n), ";", lenVar, ">", dimVar, func() {
			p.P(dimVar, "=", lenVar)
		})
	})

	return dimVar
}

// GenWriteStringBuffer function ...
func (p *TypeGenerator) GenWriteStringBuffer(ctx context.Context, scopes *Scopes, field *midl.Field, index ...interface{}) string {

	name := p.O(p.IVar(p.GoFieldName(field), index...))

	// XXX: either use ZeroString as a pointer discriminator.
	// p.P(name, "=", p.B("strings.TrimPrefix", name, "ndr.ZeroString"))

	varName := p.BufVar(name)

	if field.Attrs.Format.MultiSize {
		p.P("var", varName, "[]uint16")
		p.Range("i1", name, func() {
			p.P(varName, "=", p.B("append", varName, p.B("utf16.Encode", p.B("[]rune", p.Var(name, "[i1]")))+"..."))
			p.P(varName, "=", p.B("append", varName, p.B("uint16", "0")))
		})
		p.P(varName, "=", p.B("append", varName, p.B("uint16", "0")))
		return varName
	}

	switch scopes := scopes.Next(); scopes.Kind() {
	case midl.TypeWChar, midl.TypeUint16, midl.TypeInt16:
		p.P(varName, ":=", p.B("utf16.Encode", p.B("[]rune", name)))
	case midl.TypeChar, midl.TypeUChar, midl.TypeInt8, midl.TypeUint8:
		p.P(varName, ":=", p.B("[]byte", name))
	default:
		p.P(`return errors.New("cannot handle ` + scopes.Kind().String() + ` to string type conversion")`)
	}

	if scopes.Dim().NoSizeLimit {
		// skip for unlimited size types.
		return varName
	}

	if scopes.Dim().IsNullTerminated {

		// declared with [string] attribute. size written should be one less
		// to include zero.

		if scopes.Array().IsFixed() {
			size := scopes.Array().Size()
			p.If(p.B("uint64", p.Len(varName)), ">", size, "-", "1", func() {
				p.P(varName, "=", varName+fmt.Sprintf("[:%d-1]", size))
			})
		} else {
			sizeInfo := p.IVar("sizeInfo", scopes.Dim().Dimension)
			p.If(p.B("uint64", p.Len(varName)), ">", sizeInfo, "-", "1", func() {
				p.P(varName, "=", varName+p.IVar("", ":"+sizeInfo+"-1"))
			})
		}

		zeroValue := p.B(GoPrimitiveTypeName(scopes.Next().Kind()), "0")

		// append zeroValue only if string is not zero.
		p.If(name, "!=", "ndr.ZeroString", func() {
			p.P(varName, "=", p.B("append", varName, zeroValue))
		})

	} else {

		if scopes.Array().IsFixed() {
			size := scopes.Array().Size()
			p.If(p.B("uint64", p.Len(varName)), ">", size, func() {
				p.P(varName, "=", varName+fmt.Sprintf("[:%d]", size))
			})
		} else {
			sizeInfo := p.IVar("sizeInfo", scopes.Dim().Dimension)
			p.If(p.B("uint64", p.Len(varName)), ">", sizeInfo, func() {
				p.P(varName, "=", varName+p.IVar("", ":"+sizeInfo))
			})
		}
	}

	return varName
}

func (p *TypeGenerator) GoFieldName(field *midl.Field) string {

	if field.Name == "" {
		if field.Type.Base().Is(midl.TypeUnion) {
			field.Name = p.GoTypeName
		}
	}

	return GoFieldName(field)

}

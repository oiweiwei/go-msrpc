package gen

import (
	"context"

	"github.com/oiweiwei/go-msrpc/midl"
)

func (p *TypeGenerator) GenZeroUnionFieldMarshalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {

	name := p.VarName(ctx)
	if name == "" {
		name = p.O(p.IVar(p.GoFieldName(field), index...))
	}

	typeName := p.GoScopeTypeName(ctx, p.Scope(), field, scopes, true)

	if scopes.Union().IsEncapsulated() {
		p.CheckErr(p.B("("+p.Amp(typeName)+"{}).MarshalNDR", "ctx", "w"))
	} else {
		switchVar := "_sw" + p.ToVar(name)
		p.CheckErr(p.B("("+p.Amp(typeName)+"{}).MarshalUnionNDR", "ctx", "w", switchVar))
	}
}

func (p *TypeGenerator) GenZeroStructFieldMarshalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {

	typeName := p.GoScopeTypeName(ctx, p.Scope(), field, scopes, true)

	p.CheckErr(p.B("("+p.Amp(typeName)+"{}).MarshalNDR", "ctx", "w"))
}

func (p *TypeGenerator) GenZeroPointerFieldMarshalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {
	p.CheckErr(p.B("w.WritePointer", "nil"))
}

func (p *TypeGenerator) GenZeroEnumFieldMarshalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {

	p.CheckErr(p.B("w.WriteData", p.B(scopes.Scope().EnumType(), "0")))
}

func (p *TypeGenerator) GenZeroArrayFieldMarhalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {

	name := p.VarName(ctx)
	if name == "" {
		name = p.O(p.IVar(p.GoFieldName(field), index...))
	}

	ctx = WithVarName(ctx, ZeroLen)

	idx := p.Var("i", len(index)+1)

	if scopes.Array().IsFixed() && (!scopes.Dim().IsString || p.IsNotLastFixedArray(ctx, scopes, field)) {
		p.Block("for", idx, ":=", p.Len(name), ";", p.B("uint64", idx), "<", scopes.Array().Size(), ";", idx, "++", func() {
			p.GenZeroFieldMarshalNDR(ctx, field, scopes.Next(), append(index, idx)...)
		})
	} else if scopes.IsConformant() {
		p.Block("for", idx, ":=", p.Len(name), ";", p.B("uint64", idx), "<", p.IVar("sizeInfo", scopes.Dim().Dimension), ";", idx, "++", func() {
			p.GenZeroFieldMarshalNDR(ctx, field, scopes.Next(), append(index, idx)...)
		})
	}
}

func (p *TypeGenerator) GenZeroFieldMarshalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {

	if scopes == nil {
		return
	}

	switch {

	case scopes.Type().IsPrimitiveType():
		p.CheckErr(p.B("w.WriteData", p.GoTypeZeroValue(ctx, p.Scope(), field, scopes)))
	case scopes.Is(midl.TypeEnum):
		p.GenZeroEnumFieldMarshalNDR(ctx, field, scopes, index...)
	case scopes.Is(midl.TypeUnion):
		p.GenZeroUnionFieldMarshalNDR(ctx, field, scopes, index...)
	case scopes.Is(midl.TypeStruct), scopes.Is(midl.TypeInterface):
		p.GenZeroStructFieldMarshalNDR(ctx, field, scopes, index...)
	case scopes.Is(midl.TypePointer):
		p.GenZeroPointerFieldMarshalNDR(ctx, field, scopes, index...)
	case scopes.Is(midl.TypeArray):
		if p.GoFieldName(field) == "_" {
			// reserved field size information.
			if !p.GenMarshalSizeInfo(ctx, field, scopes, index...) {
				break
			}
		}
		p.GenZeroArrayFieldMarhalNDR(ctx, field, scopes, index...)
	}
}

// GoTypeZeroValue function returns the zero value for the given type. (nil, "", intX(0), uintX(0), false).
func (p *Generator) GoTypeZeroValue(ctx context.Context, attr *midl.TypeAttr, field *midl.Field, scopes *Scopes) string {

	for scopes != nil {
		switch scopes.Kind() {
		case midl.TypeArray, midl.TypeUnion, midl.TypeCUnion, midl.TypeStruct, midl.TypeInterface, midl.TypeVoid:
			// catch constructed types.
			if scopes.Dim().IsString && !field.Attrs.Format.MultiSize {
				return `""`
			}
			return "nil"
		case midl.TypePointer:
			scopes = scopes.Next()
			continue
		case midl.TypePipe:
			return "nil"
		}

		break
	}

	// primitive type.

	n := p.GoScopeTypeName(ctx, attr, field, scopes)
	if n == "bool" {
		return "false"
	}

	return n + "(0)"
}

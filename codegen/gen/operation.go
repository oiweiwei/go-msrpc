package gen

import (
	"context"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl"
)

type ParamGenerator struct {
	*Generator
	*TypeGenerator
}

func (p *Generator) NewParamGenerator(ctx context.Context, typ *midl.Type) *ParamGenerator {
	return &ParamGenerator{
		Generator:     p,
		TypeGenerator: p.NewTypeGenerator(ctx, typ),
	}
}

var (
	AnyParam = 0
	InParam  = 1
	OutParam = 2
)

func ParamName(dir int) string {
	switch dir {
	case InParam:
		return "Request"
	case OutParam:
		return "Response"
	}

	return ""
}

func (p *Generator) GenOperation(ctx context.Context, op *midl.Operation) {
	if p.IsUnusedOp(op.Name) {
		return
	}
	ctx = WithOp(ctx)
	p.GenOperationStruct(ctx, op, AnyParam)
	p.GenOperationMarshalNDR(ctx, op, InParam)
	p.GenOperationUnmarshalNDR(ctx, op, InParam)
	p.GenOperationMarshalNDR(ctx, op, OutParam)
	p.GenOperationUnmarshalNDR(ctx, op, OutParam)
	p.GenOperationStruct(ctx, op, InParam)
	p.GenOperationStruct(ctx, op, OutParam)
}

func (p *Generator) GenOperationMarshalNDR(ctx context.Context, op *midl.Operation, dir int) {

	pN := ParamName(dir)
	fN := "MarshalNDR" + pN
	opN := p.OpName(ctx, op, AnyParam)

	strct := &midl.Type{
		Kind: midl.TypeAttribute,
		Attrs: &midl.TypeAttr{
			Alias: Escape(p.OpName(ctx, op, dir)),
		},
		Elem: &midl.Type{
			Kind: midl.TypeStruct, Struct: &midl.Struct{}},
	}

	for _, param := range p.OperationParams(ctx, op) {
		if !p.IsDir(ctx, param.Attrs.Direction, dir) || param.IsHandle() {
			continue
		}
		strct.Elem.Struct.Fields = append(strct.Elem.Struct.Fields, &midl.Field{
			Name:  param.Name,
			Type:  param.Type,
			Attrs: param.Attrs.FieldAttr,
		})
	}

	pp := p.NewTypeGenerator(ctx, strct)

	p.P()
	p.Block("func", "(o *"+opN+")", p.XXX()+"Prepare"+pN+"Payload", "(ctx context.Context)", "error", func() {
		pp.GenStructPreparePayloadFields(ctx, strct.Elem.Struct.Fields)
		p.PrepareOperationPayloadHook(ctx, dir)
		p.P("return nil")
	})

	p.P()
	p.Block("func", "(o *"+opN+")", fN, "(ctx context.Context, w ndr.Writer)", "error", func() {

		p.CheckErr("o." + p.XXX() + "Prepare" + pN + "Payload(ctx)")

		for _, param := range p.OperationParams(ctx, op) {

			if !p.IsDir(ctx, param.Attrs.Direction, dir) || param.IsHandle() {
				continue
			}

			field := &midl.Field{
				Name:  param.Name,
				Type:  param.Type,
				Attrs: param.Attrs.FieldAttr,
			}

			scopes := NewScopes(field.Scopes())
			if scopes.Is(midl.TypePointer) &&
				(scopes.Scope().Pointer == midl.PointerTypeNone || scopes.Scope().Pointer == midl.PointerTypeRef || scopes.Scope().Pointer == midl.PointerTypeRefWeak) &&
				// pointers to interfaces are always unique pointers.
				!scopes.Next().Is(midl.TypeInterface) {
				// skip ref pointer.
				scopes = scopes.Next()
			}

			p.P("//", field.Name, pp.GenParamComment(ctx, param), pp.GenStructFieldComment(ctx, field))
			p.Block("", func() {
				pp.GenFieldMarshalNDR(ctx, field, scopes)
				if scopes.IsDeferrable() {
					p.CheckErr(p.B("w.WriteDeferred"))
				}
			})

		}

		p.P("return nil")
	})

}

func (p *Generator) PrepareOperationPayloadHook(ctx context.Context, dir int) {
	pN := ParamName(dir)
	p.If("hook, ok := (interface{})(o).(interface { AfterPrepare"+pN+"Payload(context.Context) error }); ok", func() {
		p.CheckErr(p.B("hook.AfterPrepare"+pN+"Payload", "ctx"))
	})
}

func (p *Generator) GenOperationUnmarshalNDR(ctx context.Context, op *midl.Operation, dir int) {

	fn := "UnmarshalNDR"
	if dir == InParam {
		fn += "Request"
	} else if dir == OutParam {
		fn += "Response"
	}

	p.P()
	p.Block("func", "(o *"+p.OpName(ctx, op, AnyParam)+")", fn, "(ctx context.Context, w ndr.Reader)", "error", func() {

		strct := &midl.Type{
			Kind: midl.TypeAttribute,
			Attrs: &midl.TypeAttr{
				Alias: Escape(p.OpName(ctx, op, dir)),
			},
			Elem: &midl.Type{
				Kind: midl.TypeStruct, Struct: &midl.Struct{}},
		}

		pp := p.NewTypeGenerator(ctx, strct)

		for _, param := range p.OperationParams(ctx, op) {

			if !p.IsDir(ctx, param.Attrs.Direction, dir) || param.IsHandle() {
				continue
			}

			field := &midl.Field{
				Name:  param.Name,
				Type:  param.Type,
				Attrs: param.Attrs.FieldAttr,
			}

			scopes := NewScopes(field.Scopes())
			if scopes.Is(midl.TypePointer) && (scopes.Scope().Pointer == midl.PointerTypeNone || scopes.Scope().Pointer == midl.PointerTypeRef || scopes.Scope().Pointer == midl.PointerTypeRefWeak) &&
				// pointers to interfaces are always unique pointers.
				!scopes.Next().Is(midl.TypeInterface) {
				scopes = scopes.Next()
			}

			p.P("//", field.Name, pp.GenParamComment(ctx, param), pp.GenStructFieldComment(ctx, field, "ref"))
			p.P("{")
			pp.GenFieldUnmarshalNDR(ctx, field, scopes)
			if scopes.IsDeferrable() {
				p.CheckErr(p.B("w.ReadDeferred"))
			}
			p.P("}")

		}

		p.P("return nil")
	})

}

func (p *Generator) GenParamComment(ctx context.Context, param *midl.Param) string {
	if param.Attrs == nil {
		return "{}"
	}
	attr := *param.Attrs
	attr.FieldAttr = nil
	return "{" + (&attr).String() + "}"
}

func (p *Generator) MethodName(ctx context.Context, op *midl.Operation) string {
	n := GoName(op.Name)
	if strings.HasPrefix(op.Name, "_") {
		n = "_" + n
	}
	if op.Attrs.PropGet {
		n = "Get" + n
	}
	if op.Attrs.PropPut {
		n = "Set" + n
	}
	if op.Attrs.PropPutRef {
		n = "SetByRef" + n
	}
	return n
}

func (p *Generator) OpName(ctx context.Context, op *midl.Operation, dir int) string {
	n := p.MethodName(ctx, op)
	if dir == InParam {
		n += "Request"
		return n
	} else if dir == OutParam {
		n += "Response"
		return n
	}
	return p.XXX() + n + "Operation"
}

func (p *Generator) IsDir(ctx context.Context, pdir midl.Direction, dir int) bool {
	return dir == AnyParam || (dir == InParam && pdir.In) || (dir == OutParam && pdir.Out)
}

func (p *Generator) ReturnValue() string {
	return "Return"
}

func (p *Generator) OperationReturnValue(ctx context.Context, op *midl.Operation) *midl.Param {
	return &midl.Param{
		Name:  p.ReturnValue(),
		Type:  op.Type,
		Attrs: &midl.ParamAttr{FieldAttr: &midl.FieldAttr{}, Direction: midl.Direction{Out: true}},
	}
}

func (p *Generator) OperationParams(ctx context.Context, op *midl.Operation) []*midl.Param {

	params := make([]*midl.Param, len(op.Params))
	copy(params, op.Params)

	if !op.Type.Is(midl.TypeVoid) {
		params = append(params, p.OperationReturnValue(ctx, op))
	}

	if !Interface(ctx).IsObject() {
		return params
	}

	return append([]*midl.Param{
		{
			Name:  "This",
			Type:  midl.LookupType("ORPCTHIS"),
			Attrs: &midl.ParamAttr{FieldAttr: &midl.FieldAttr{}, Direction: midl.Direction{In: true}},
		},
		{
			Name:  "That",
			Type:  midl.LookupType("ORPCTHAT"),
			Attrs: &midl.ParamAttr{FieldAttr: &midl.FieldAttr{}, Direction: midl.Direction{Out: true}},
		},
	}, params...)
}

func (p *Generator) GetExprIdents(ctx context.Context, expr ...midl.Expr) []string {
	ret := make([]string, 0, len(expr))
	for _, expr := range expr {
		val, ok := expr.Ident().Value.(string)
		if ok {
			ret = append(ret, val)
		}
	}
	return ret
}

func (p *Generator) GetOutputParamImplicitDependents(ctx context.Context, op *midl.Operation, dir int) []*midl.Param {

	if dir != OutParam {
		return nil
	}

	deps, seen := []*midl.Param{}, make(map[string]bool)
	params := make(map[string]*midl.Param)

	for _, param := range p.OperationParams(ctx, op) {
		params[param.Name] = param
	}

	for _, param := range p.OperationParams(ctx, op) {
		if !param.Attrs.Direction.Out {
			continue
		}
		for _, expr := range [][]midl.Expr{
			param.Attrs.SizeIs,
			param.Attrs.LengthIs,
			param.Attrs.FirstIs,
			param.Attrs.LastIs,
			param.Attrs.MaxIs,
			param.Attrs.MinIs,
			[]midl.Expr{param.Attrs.SwitchIs},
		} {
			for _, ident := range p.GetExprIdents(ctx, expr...) {
				if p, ok := params[ident]; ok && p.Attrs.Direction.In && !p.Attrs.Direction.Out && !seen[p.Name] {
					deps, seen[p.Name] = append(deps, p), true
				}
			}
		}
	}

	return deps
}

func (p *Generator) GenOperationStruct(ctx context.Context, op *midl.Operation, dir int) {

	p.P()
	p.P("//", p.OpName(ctx, op, dir), "structure", "represents", "the", RPCName(op.Name),
		"operation", strings.ToLower(ParamName(dir)))

	doc, _ := p.Doc.Type(op.Name)

	// generate go structure for the in/out/any parameters.
	p.Structure(p.OpName(ctx, op, dir), func() {
		for _, param := range p.GetOutputParamImplicitDependents(ctx, op, dir) {
			p.P("//", "XXX:", param.Name, "is an implicit input depedency for output parameters")
			p.NewParamGenerator(ctx, param.Type).GenStructField(ctx, &midl.Field{
				Name:  param.Name,
				Attrs: param.Attrs.FieldAttr,
				Type:  param.Type,
			})
		}

		for _, param := range p.OperationParams(ctx, op) {
			if !p.IsDir(ctx, param.Attrs.Direction, dir) {
				continue
			}
			if param.IsHandle() {
				// skip handle.
				continue
			}
			if dir != AnyParam && GoName(param.Name) == "_" {
				// skip reserved.
				continue
			}
			if dir != AnyParam {
				// gen docstring.
				if doc, ok := doc.ObjectField(param.Name); ok {
					p.GenComment(ctx, doc.Doc)
				} else if param.Name == p.ReturnValue() {
					p.P("//", p.ReturnValue()+":", "The", op.Name, "return value.")
				}
			}
			// generate structure field.
			p.NewParamGenerator(ctx, param.Type).GenStructField(ctx, &midl.Field{
				Name:  param.Name,
				Attrs: param.Attrs.FieldAttr,
				Type:  param.Type,
			})
		}
	})

	if dir == AnyParam {
		p.GenOpNum(ctx, op, dir)
		p.GenOpName(ctx, op, dir)
		return
	}

	p.GenOperationToOp(ctx, op, dir)
	p.GenOperationFromOp(ctx, op, dir)

	dirN := ParamName(dir)

	p.Block("func", "(o *"+p.OpName(ctx, op, dir)+")", "MarshalNDR", "(ctx context.Context, w ndr.Writer)", "error", func() {
		p.P("return", "o."+p.XXX()+"ToOp(ctx, nil).MarshalNDR"+dirN+"(ctx, w)")
	})

	p.Block("func", "(o *"+p.OpName(ctx, op, dir)+")", "UnmarshalNDR", "(ctx context.Context, r ndr.Reader)", "error", func() {
		p.P("_o", ":=", p.Amp(p.OpName(ctx, op, AnyParam))+"{}")
		p.CheckErr("_o.UnmarshalNDR" + dirN + "(ctx, r)")
		p.P("o." + p.XXX() + "FromOp(ctx, _o)")
		p.P("return", "nil")
	})

}

func (p *Generator) GenOpNum(ctx context.Context, op *midl.Operation, dir int) {
	p.P()
	p.P("func", "(o *"+p.OpName(ctx, op, dir)+")", "OpNum()", "int", "{", "return", op.OpNum, "}")
}

func (p *Generator) GenOpName(ctx context.Context, op *midl.Operation, dir int) {
	p.P()
	p.P("func", "(o *"+p.OpName(ctx, op, dir)+")", "OpName()", "string", "{",
		"return", p.Q("/"+Interface(ctx).Name+"/"+Interface(ctx).Attrs.Version.String()+"/"+op.Name),
		"}")
}

func (p *Generator) GenOperationToOp(ctx context.Context, op *midl.Operation, dir int) {

	p.P()
	p.P("func", "(o *"+p.OpName(ctx, op, dir)+")", p.XXX()+"ToOp(ctx context.Context, op *"+p.OpName(ctx, op, AnyParam), ")", "*"+p.OpName(ctx, op, AnyParam), "{")
	p.If("op == nil", func() {
		p.P("op", "=", "&"+p.OpName(ctx, op, AnyParam)+"{}")
	})
	p.If("o == nil", func() {
		p.P("return", "op")
	})

	if implicit := p.GetOutputParamImplicitDependents(ctx, op, dir); len(implicit) > 0 {
		p.P("//", "XXX:", "implicit input dependencies for output parameters")
		for _, param := range implicit {
			n := GoName(param.Name)
			if n == "_" {
				continue
			}
			f := &midl.Field{
				Name:  param.Name,
				Type:  param.Type,
				Attrs: param.Attrs.FieldAttr,
			}

			// if op.(n) is zero value, you can set it.
			p.If("op."+n, "==", p.GoTypeZeroValue(ctx, f.Type.Attrs, f, NewScopes(f.Scopes())), func() {
				p.P("op."+n, "=", p.O(n))
			})
		}
		p.P()
	}

	for _, param := range p.OperationParams(ctx, op) {
		if !p.IsDir(ctx, param.Attrs.Direction, dir) || param.IsHandle() {
			continue
		}
		n := GoName(param.Name)
		if n == "_" {
			continue
		}
		p.P("op."+n, "=", p.O(n))

	}
	p.P("return", "op")
	p.P("}")
}

func (p *Generator) GenOperationFromOp(ctx context.Context, op *midl.Operation, dir int) {
	p.P()
	p.P("func", "(o *"+p.OpName(ctx, op, dir)+")", p.XXX()+"FromOp(ctx context.Context, op *"+p.OpName(ctx, op, AnyParam), ")", "{")
	p.If("o == nil", func() {
		p.P("return")
	})
	if implicit := p.GetOutputParamImplicitDependents(ctx, op, dir); len(implicit) > 0 {
		p.P("//", "XXX:", "implicit input dependencies for output parameters")
		for _, param := range implicit {
			n := GoName(param.Name)
			if n == "_" {
				continue
			}
			p.P(p.O(n), "=", "op."+n)
		}
		p.P()
	}

	for _, param := range p.OperationParams(ctx, op) {
		if !p.IsDir(ctx, param.Attrs.Direction, dir) || param.IsHandle() {
			continue
		}
		n := GoName(param.Name)
		if n == "_" {
			continue
		}
		p.P(p.O(n), "=", "op."+n)

	}
	p.P("}")
}

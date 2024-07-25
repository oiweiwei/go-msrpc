package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl"
)

func (p *TypeGenerator) GenUnion(ctx context.Context) {

	sw := p.Union().Switch

	if p.Scope().Parent != "" {
		p.P("//", p.GoTypeName, "structure", "represents", RPCName(p.Scope().Parent), "union", "anonymous", "member.")
	} else {
		p.P("//", p.GoTypeName, "structure", "represents", RPCName(p.Alias()), "RPC", "union.")
	}
	if p.Doc != nil {
		p.P("//")
		p.GenComment(ctx, p.Doc.Doc)
	}

	p.Structure(p.GoTypeName, func() {
		if sw != nil {
			p.P(GoName(sw.Name), p.SwitchType(ctx, p.Scopes))
		}
		p.CB(func() {
			p.P("Types that are assignable to Value")
			p.P()
			for _, cases := range p.Union().Body {
				p.P("*" + p.UnionArmName(ctx, cases))
			}
		})
		p.P("Value", "is_"+p.GoTypeName, "`json:\"value\"`")
	})

	p.GenerateUnionGetValue(ctx)

	p.P()
	p.P("type", "is_"+p.GoTypeName, "interface", "{")
	p.P("ndr.Marshaler")
	p.P("ndr.Unmarshaler")
	p.P("is_"+p.GoTypeName, "()")
	p.P("}")

	if !p.Union().IsEncapsulated() {
		p.GenerateNDRSwitchValue(ctx)
	}
	p.GenLayout(ctx)
	p.GenUnionMarshalNDR(ctx)
	p.GenUnionUnmarshalNDR(ctx)

	for _, cases := range p.Union().Body {
		// generate interface definition.
		p.GenUnionArm(ctx, cases)
	}
}

func (p *TypeGenerator) GenerateUnionGetValue(ctx context.Context) {

	p.P()

	p.Block("func", "(o *"+p.GoTypeName+")", "GetValue()", "any", func() {
		p.If("o == nil", func() {
			p.P("return", "nil")
		})

		p.Block("switch", "value", ":=", "(interface{})(o.Value).(type)", func() {
			for _, cases := range p.Union().Body {
				if len(cases.Arms) == 0 || p.GoFieldName(cases.Arms[0]) == "_" {
					continue
				}

				// determine for-sure.
				p.P("case", "*"+p.UnionArmName(ctx, cases), ":")
				p.If("value", "!=", "nil", func() {
					p.P("return", "value."+p.GoFieldName(cases.Arms[0]))
				})
			}
		})

		p.P("return", "nil")
	})

}

func (p *TypeGenerator) SwitchTypeZeroValue(ctx context.Context, switchType string) any {
	switch switchType {
	case "string":
		return `""`
	}
	return 0
}

func (p *TypeGenerator) GenerateNDRSwitchValue(ctx context.Context) {

	switchType := p.SwitchType(ctx, p.Scopes)

	p.P()
	p.Block("func", "(o *"+p.GoTypeName+")", "NDRSwitchValue(sw "+switchType+")", switchType, func() {
		p.If("o == nil", func() {
			p.P("return", p.B(switchType, p.SwitchTypeZeroValue(ctx, switchType)))
		})
		p.Block("switch", "(interface{})(o.Value).(type)", func() {
			for _, cases := range p.Union().Body {
				if len(cases.Labels) == 0 {
					continue
				}
				// determine for-sure.
				p.P("case", "*"+p.UnionArmName(ctx, cases), ":")
				if len(cases.Labels) > 1 {
					p.Block("switch", "sw", func() {
						var labels []string
						for _, label := range cases.Labels {
							labels = append(labels, p.B(switchType, label))
						}
						p.P("case", strings.Join(labels, ",\n"), ":")
						p.P("return", "sw")
					})
				}
				p.P("return", p.B(switchType, cases.Labels[0]))
			}
		})
		p.P("return", p.B(switchType, p.SwitchTypeZeroValue(ctx, switchType)))
	})
}

func (p *TypeGenerator) GenerateSwitchIs(ctx context.Context, field *midl.Field, scopes *Scopes, switchVar string) {
	switchIs, switchType := field.Attrs.SwitchIs, p.SwitchType(ctx, scopes)
	p.P(switchVar, ":=", p.B(switchType, p.GenExpr(ctx, switchIs, p.LookupExprField(ctx, switchIs), switchType)))
}

func (p *TypeGenerator) SwitchType(ctx context.Context, scopes *Scopes) string {

	var switchType string

	if p.Union() != nil && p.Union().IsEncapsulated() {
		switchType = p.Generator.GoFieldTypeName(ctx, p.Scope(), &midl.Field{
			Attrs: &midl.FieldAttr{},
			Type:  p.Union().Switch.Type,
		})
	} else {
		if scope := NewScopes(scopes.Scope().SwitchType.Scopes()); scope != nil && scope.Is(midl.TypeEnum) {
			switchType = scope.EnumType().String()
		} else {
			switchType = p.Generator.GoFieldTypeName(ctx, p.Scope(), &midl.Field{
				Attrs: &midl.FieldAttr{},
				Type:  scopes.Scope().SwitchType,
			})
		}
	}

	return switchType
}

func (p *TypeGenerator) GenUnionMarshalNDR(ctx context.Context) {

	p.P()

	switchType := p.SwitchType(ctx, p.Scopes)

	sw := p.Union().Switch

	if sw != nil {
		p.P("func", "(o *"+p.GoTypeName+")", "MarshalNDR(ctx context.Context, w ndr.Writer)", "error", "{")
	} else {
		p.P("func", "(o *"+p.GoTypeName+")", "MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw "+switchType+")", "error", "{")
	}

	// XXX: perform more strict type checking.
	// p.P("_switch, ok := sw.(" + switchType + ")")
	// p.If("!ok", func() {
	//	p.P("return", `fmt.Errorf("incompatible switch type %%T, expected `+switchType+`", sw)`)
	// })

	swVar := "sw"
	if sw != nil {
		swVar = p.O(GoName(sw.Name))
	}

	if sw != nil {
		p.CheckErr(p.B("w.WriteData", p.B(switchType, swVar)))
	} else {
		p.CheckErr(p.B("w.WriteSwitch", p.B(switchType, swVar)))
	}

	if MSUnion(ctx) && sw == nil {
		p.P("//", "ms_union")
		p.GenDoAlignmentMarshalNDR(ctx, p.Alignment())
	}

	p.Block("switch", swVar, func() {
		var defaultCase *midl.UnionCase
		for _, cases := range p.Union().Body {
			if cases.IsDefault {
				defaultCase = cases
				continue
			}
			labels := []string{}
			for _, label := range cases.Labels {
				labels = append(labels, p.B(switchType, label))
			}
			p.P("case", strings.Join(labels, ",\n"), ":")
			if len(cases.Arms) == 0 {
				continue
			}

			armName := p.UnionArmName(ctx, cases)

			if p.IsArmPointerToPrimitiveType(ctx, cases) {
				// special case, take care of the nil value for the union arm containing
				// primitive type.
				p.GenUnionArmPrimitiveTypeMarshalNDR(ctx, cases, armName)
			} else if len(cases.Arms) > 0 {
				p.P("_o, _ := o.Value.(*" + armName + ")")
				p.If("_o != nil", func() {
					p.CheckErr(p.B("_o.MarshalNDR", "ctx", "w"))
				}, p.Else(func() {
					p.CheckErr(p.B("("+p.Amp(armName+"{}")+")"+".MarshalNDR", "ctx", "w"))
				}))
			} else {
				p.P("//", "no-op")
			}
		}

		p.P("default", ":")
		if defaultCase == nil {
			p.P("return", `fmt.Errorf("unsupported switch case value %v", `+swVar+`)`)
		} else if len(defaultCase.Arms) > 0 {
			armName := p.UnionArmName(ctx, defaultCase)
			if p.IsArmPointerToPrimitiveType(ctx, defaultCase) {
				// special case, take care of the nil value for the union arm containing
				// primitive type.
				p.GenUnionArmPrimitiveTypeMarshalNDR(ctx, defaultCase, armName)
			} else if len(defaultCase.Arms) > 0 {
				p.P("_o, _ := o.Value.(*" + armName + ")")
				p.If("_o != nil", func() {
					p.CheckErr(p.B("_o.MarshalNDR", "ctx", "w"))
				}, p.Else(func() {
					if p.IsArmPointerToPrimitiveType(ctx, defaultCase) {
						p.GenZeroPointerFieldMarshalNDR(ctx, nil, nil)
					} else {
						p.CheckErr(p.B("("+p.Amp(armName+"{}")+")"+".MarshalNDR", "ctx", "w"))
					}
				}))
			} else {
				p.P("//", "no-op")
			}
		}
	})
	p.P("return nil")
	p.P("}")

}

func (p *TypeGenerator) GenUnionArmPrimitiveTypeMarshalNDR(ctx context.Context, cases *midl.UnionCase, armName string) {
	p.P("_o, _ := o.Value.(*" + armName + ")")
	p.If("_o", "!=", "nil", func() {
		p.P("_ptr_o", ":=", "ndr.MarshalNDRFunc", "(", "func(ctx context.Context, w ndr.Writer) error {")
		p.CheckErr(p.B("_o.MarshalNDR", "ctx", "w"))
		p.P("return nil")
		p.P("})")
		p.CheckErr(p.B("w.WritePointer", p.Amp("_o"), "_ptr_o"))
	}, p.Else(func() {
		p.GenZeroPointerFieldMarshalNDR(ctx, nil, nil)
	}))
}

func (p *TypeGenerator) IsArmPointerToPrimitiveType(ctx context.Context, cases *midl.UnionCase) bool {

	if len(cases.Arms) != 1 {
		return false
	}

	scopes := NewScopes(cases.Arms[0].Scopes())
	return scopes.Is(midl.TypePointer) && scopes.Next().Type().IsPrimitiveType()
}

func (p *TypeGenerator) UnionArmName(ctx context.Context, cases *midl.UnionCase) string {

	caseName := p.GoTypeName

	if parts := strings.Split(caseName, "_"); len(parts) == 2 && strings.Contains(parts[0], parts[1]) {
		caseName = parts[0]
	} else if len(parts) == 2 && strings.Contains(parts[1], strings.ReplaceAll(parts[0], "Enum", "")) {
		caseName = parts[1]
	} else if len(parts) == 2 && strings.Contains(parts[1], strings.ReplaceAll(parts[0], "Record", "")) {
		caseName = parts[1]
	} else if len(parts) == 2 && parts[1] == "Value" {
		caseName = parts[0]
	}

	if len(cases.Arms) > 0 {
		// use arm name as a union arm type name.
		n := GoName(cases.Arms[0].Name)

		if strings.Contains(n, caseName) || strings.Contains(caseName, n) {
			n = strings.ReplaceAll(n, caseName, "")
		}

		if n == "_" {
			n = GoNameNoReserved(cases.Arms[0].Name)
		}

		if n != "" {
			return caseName + "_" + n
		}
	}

	if cases.IsDefault {
		// there must be at most one default case for a union.
		return caseName + "_Default" + strings.Split(p.GoTypeName, "_")[0]
	}

	// try find suitable name by enumeration label.
	if switchType := p.Scope().SwitchType; switchType != nil {
		if switchType = switchType.Base(); switchType.Is(midl.TypeEnum) {
			for _, e := range switchType.Enum.Elems {
				for _, label := range cases.Labels {
					val, ok := label.(midl.Expr).Int64()
					if !ok {
						continue
					}
					if e.ID == int(val) {
						return caseName + "_" + GoName(e.Value)
					}
				}
			}
		}
	}

	if len(cases.Labels) > 0 {
		return caseName + "_" + fmt.Sprintf("%v", cases.Labels[0])
	}

	// A union arm that consists solely of a terminating semicolon
	// is legal and specifies a null arm. (no name, no labels).
	return caseName + "_" + fmt.Sprintf("%d", cases.Position)
}

func (p *TypeGenerator) GenUnionArmInterface(ctx context.Context, cases *midl.UnionCase) {
	p.P()
	p.P("func", "(*"+p.UnionArmName(ctx, cases)+")", "is_"+p.GoTypeName, "()", "{}")
}

func (p *TypeGenerator) GenUnionArm(ctx context.Context, cases *midl.UnionCase) {

	if cases.IsDefault && len(cases.Arms) == 0 {
		return
	}

	caseName := p.UnionArmName(ctx, cases)

	p.P()
	if cases.IsDefault {
		p.P("//", caseName, "structure", "represents", RPCName(p.Alias()), "RPC", "default", "union arm.")
	} else {
		p.P("//", caseName, "structure", "represents", RPCName(p.Alias()), "RPC", "union arm.")
		p.P("//")
		p.P("//", "It has following labels:", cases.String())
	}

	p.Structure(caseName, func() {
		for _, field := range p.ArmFields(ctx, cases) {
			p.GenStructField(ctx, field)
		}
	})

	p.GenUnionArmInterface(ctx, cases)
	p.GenUnionArmMarshalNDR(ctx, cases, caseName)
	p.GenUnionArmUnmarshalNDR(ctx, cases, caseName)

	// skip embedded struct, because it will be generated as an arm.
	if !p.IsEmbeddedArmStruct(ctx, cases) {
		p.GenSubTypes(ctx, cases.Arms)
	}
}

func (p *TypeGenerator) ArmFields(ctx context.Context, cases *midl.UnionCase) []*midl.Field {

	// if p.IsEmbeddedArmStruct(ctx, cases) && cases.Arms[0].Type.Struct != nil {
	//	return cases.Arms[0].Type.Struct.Fields
	// }

	return cases.Arms
}

func (p *TypeGenerator) IsEmbeddedArmStruct(ctx context.Context, cases *midl.UnionCase) bool {
	return len(cases.Arms) > 0 && (cases.Arms[0].Type.Is(midl.TypeStruct) || cases.Arms[0].Type.Is(midl.TypeUnion)) && cases.Arms[0].Type.Name == "" && cases.Arms[0].Type.Tag == ""
}

func (p *TypeGenerator) GenUnionArmMarshalNDR(ctx context.Context, arm *midl.UnionCase, caseName string) {

	p.P()
	p.P("func", "(o *"+caseName+")", "MarshalNDR(ctx context.Context, w ndr.Writer)", "error", "{")

	for _, field := range p.ArmFields(ctx, arm) {
		scopes := NewScopes(field.Scopes())
		if p.IsArmPointerToPrimitiveType(ctx, arm) {
			scopes = scopes.Next()
		}
		p.GenFieldMarshalNDR(ctx, field, scopes)
	}

	p.P("return nil")
	p.P("}")
}

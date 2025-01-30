package gen

import (
	"context"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl"
)

func (p *TypeGenerator) GenUnionUnmarshalNDR(ctx context.Context) {

	p.P()

	switchType := p.SwitchType(ctx, p.Scopes)

	sw := p.Union().Switch

	if sw != nil {
		p.P("func", "(o *"+p.GoTypeName+")", "UnmarshalNDR(ctx context.Context, w ndr.Reader)", "error", "{")
	} else {
		p.P("func", "(o *"+p.GoTypeName+")", "UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw "+switchType+")", "error", "{")
	}

	swVar := "sw"
	if sw != nil {
		swVar = p.O(GoName(sw.Name))
	}

	p.GenDoUnionAlignemntUnmarshalNDR(ctx, p.Alignment())

	if sw != nil {
		if p.IsEnumSwitch(ctx, p.Scopes) {
			p.CheckErr(p.B("w.ReadEnum", p.B(p.BPtr(switchType), p.Amp(swVar))))
		} else {
			p.CheckErr(p.B("w.ReadData", p.B(p.BPtr(switchType), p.Amp(swVar))))
		}
	} else {
		if p.IsEnumSwitch(ctx, p.Scopes) {
			p.CheckErr(p.B("w.ReadSwitch", p.B("ndr.Enum", p.B(p.BPtr(switchType), p.Amp(swVar)))))
		} else {
			p.CheckErr(p.B("w.ReadSwitch", p.B(p.BPtr(switchType), p.Amp(swVar))))
		}
	}

	if MSUnion(ctx) && sw == nil {
		p.P("//", "ms_union")
		p.GenDoAlignmentUnmarshalNDR(ctx, p.Alignment())
	} else {
		p.GenDoUnionAlignemntUnmarshalNDR(ctx, p.Alignment())
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
			armName := p.UnionArmName(ctx, cases)
			if p.IsArmPointerToPrimitiveType(ctx, cases) {
				p.GenUnionArmPrimitiveTypeUnmarshalNDR(ctx, cases, armName)
			} else {
				p.P("o.Value", "=", p.Amp(armName)+"{}")
				if len(cases.Arms) > 0 {
					p.CheckErr(p.B("o.Value.UnmarshalNDR", "ctx", "w"))
				}
			}
		}

		p.P("default", ":")
		if defaultCase == nil {
			p.P("return", `fmt.Errorf("unsupported switch case value %v", `+swVar+`)`)
		} else if len(defaultCase.Arms) > 0 {
			armName := p.UnionArmName(ctx, defaultCase)
			if p.IsArmPointerToPrimitiveType(ctx, defaultCase) {
				p.GenUnionArmPrimitiveTypeUnmarshalNDR(ctx, defaultCase, armName)
			} else {
				p.P("o.Value", "=", p.Amp(armName)+"{}")
				if len(defaultCase.Arms) > 0 {
					p.CheckErr(p.B("o.Value.UnmarshalNDR", "ctx", "w"))
				}
			}
		}
	})
	p.P("return", "nil")
	p.P("}")

}

func (p *TypeGenerator) GenUnionArmPrimitiveTypeUnmarshalNDR(ctx context.Context, cases *midl.UnionCase, armName string) {
	// render deferred unmarshalling function.
	p.P("_ptr_o", ":=", "ndr.UnmarshalNDRFunc", "(", "func(ctx context.Context, w ndr.Reader) error {")
	p.P("o.Value", "=", p.Amp(armName)+"{}")
	p.CheckErr(p.B("o.Value.UnmarshalNDR", "ctx", "w"))
	p.P("return nil")
	p.P("})")
	p.P("_s_o", ":=", "func(ptr interface{}) {", "o.Value", "=", "*ptr.(**", armName, ")", "}")
	p.CheckErr(p.B("w.ReadPointer", p.Amp("o.Value"), "_s_o", "_ptr_o"))
}

func (p *TypeGenerator) GenUnionArmUnmarshalNDR(ctx context.Context, arm *midl.UnionCase, caseName string) {

	p.P("func", "(o *"+caseName+")", "UnmarshalNDR(ctx context.Context, w ndr.Reader)", "error", "{")

	if p.IsEmbeddedArmStruct(ctx, arm) {
		p.GenDoAlignmentUnmarshalNDR(ctx, NewScopes(arm.Arms[0].Scopes()).Alignment())
	}

	for _, field := range p.ArmFields(ctx, arm) {
		scopes := NewScopes(field.Scopes())
		if p.IsArmPointerToPrimitiveType(ctx, arm) {
			scopes = scopes.Next()
		}
		p.GenFieldUnmarshalNDR(ctx, field, scopes)
	}

	p.P("return nil")
	p.P("}")
}

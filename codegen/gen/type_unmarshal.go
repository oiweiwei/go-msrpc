package gen

import (
	"context"

	"github.com/oiweiwei/go-msrpc/midl"
)

func (p *TypeGenerator) GenStructUnmarshalNDR(ctx context.Context) {

	trailingField := len(p.Struct().Fields) - 1
	if p.IsConformant() || p.IsVarying() {
		trailingField--
	}

	p.Block("func", "(o *"+p.GoTypeName+")", "UnmarshalNDR(ctx context.Context, w ndr.Reader)", "error", func() {
		// size checks.
		p.GenStructUnmarshalNDRSizeCheck(ctx)
		// write size information.
		p.GenStructUnmarshalNDRSizePreamble(ctx, p.StructLastField())
		// align the structure.
		p.GenDoAlignmentUnmarshalNDR(ctx, p.Alignment())
		// unmarshal fields.
		for i, field := range p.Struct().Fields {
			p.GenFieldUnmarshalNDR(ctx, field, NewScopes(field.Scopes()))
			if i == trailingField {
				if f := NewScopes(field.Scopes()); f.alignment(false) != p.alignment(false) {
					p.GenDoTrailingGapUnmarshalNDR(ctx, p.Alignment())
				}
			}
		}
		if p.Scope().Pad != 0 {
			p.P("//", "pad", p.Scope().Pad)
			p.GenDoAlignmentUnmarshalNDR(ctx, int(p.Scope().Pad))
		}
		p.P("return nil")
	})

}

func (p *TypeGenerator) GenStructUnmarshalNDRSizeCheck(ctx context.Context) {
	if rng := p.Scope().Range; rng != nil {
		if rng.Min == rng.Max && rng.Min > 0 {
			p.If("w.Len()", "!=", rng.Min, "/* size-is check */", func() {
				p.P("return", "nil")
			})
		} else {
			if rng.Min > 0 {
				p.If("w.Len()", "<", rng.Min, "/* min-is check */", func() {
					p.P("return", "nil")
				})
			}
			if rng.Max > 0 {
				p.If("w.Len()", ">", rng.Max, "/* max-is check */", func() {
					p.P("return", "nil")
				})
			}
		}
	}
}

func (p *TypeGenerator) GenStructUnmarshalNDRSizePreamble(ctx context.Context, field *midl.Field) {
	if !p.IsConformant() {
		return
	}
	p.P("sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)")
	p.If("!ok", func() {
		p.P("sizeInfo", "=", p.O("NDRSizeInfo()"))
		p.Range("i1", "sizeInfo", func() {
			p.GenReadSize(ctx, "sizeInfo[i1]")
		})
		p.P("ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)")
	})
}

func (p *Generator) GenReadSize(ctx context.Context, sz interface{}) {
	p.CheckErr(p.B("w.ReadSize", p.Amp(sz)))
}

func (p *TypeGenerator) GenFieldUnmarshalNDR(ctx context.Context, field *midl.Field, scopes *Scopes, index ...interface{}) {

	if scopes == nil || field.Attrs.Ignore {
		// end.
		return
	}

	var name string

	// reserved field.
	if p.GoFieldName(field) == "_" && !IsReserved(ctx) {
		ctx = WithReserved(ctx)
		typeName := p.GoScopeTypeName(ctx, p.Scope(), field, scopes)
		p.P("//", "reserved", field.Name)
		name = "_" + field.Name
		p.P("var", name, typeName)
		ctx, index = WithVarName(ctx, name), nil
	}

	// get variable name.
	if name = p.VarName(ctx, index...); name == "" {
		name = p.O(p.IVar(p.GoFieldName(field), index...))
	}

	switch {

	case scopes.Type().Is(midl.TypeHandle):

		p.P("//", "skip", "[handle]", field.Name)

	case scopes.Type().IsPrimitiveType():

		// marshal primitive type.

		if scopes.IsBool() && midl.PrimitiveTypeSize(scopes.Kind()) > 1 {

			typeName := GoPrimitiveTypeName(scopes.Kind())

			varName := "_b" + p.ToVar(name)
			p.P("var", varName, typeName)

			p.CheckErr(p.B("w.ReadData", p.Amp(varName)))

			p.P(name, "=", varName, "!=", "0")

		} else if scopes.Is(midl.TypeInt32_64) {
			p.CheckErr(p.B("w.ReadData", p.B(p.BPtr("ndr.Int3264"), p.Amp(name))))
		} else if scopes.Is(midl.TypeUint32_64) {
			p.CheckErr(p.B("w.ReadData", p.B(p.BPtr("ndr.Uint3264"), p.Amp(name))))
		} else if field.Attrs.Format.Rune {
			varName := "_" + p.ToVar(name)
			p.P("var", varName, "uint16")
			p.CheckErr(p.B("w.ReadData", p.Amp(varName)))
			p.P(name, "=", p.B("rune", varName))
		} else {
			p.CheckErr(p.B("w.ReadData", p.Amp(name)))
		}

	case scopes.Is(midl.TypeEnum):

		// marshal enum.
		if !scopes.IsV1Enum() && scopes.Enum().Is32 {
			varName := "_e" + p.ToVar(name)
			p.P(varName, ":=", p.B(scopes.EnumType(), name))
			p.CheckErr(p.B("w.ReadEnum", p.Amp(varName)))
			p.P(name, "=", p.B(p.GoScopeTypeName(ctx, p.Scope(), field, scopes), varName))
		} else {
			p.CheckErr(p.B("w.ReadEnum", p.B(p.BPtr(scopes.EnumType()), p.Amp(name))))
		}

	case scopes.Is(midl.TypeUnion):

		// marshal union.

		typeName := p.GoScopeTypeName(ctx, p.Scope(), field, scopes, true)

		p.If(name, "==", "nil", func() {
			p.P(name, "=", p.Amp(typeName), "{}")
		})

		switchVar := "_sw" + p.ToVar(name)
		if !scopes.Union().IsEncapsulated() {
			p.GenerateSwitchIs(ctx, field, scopes, switchVar)
		}

		if !scopes.Union().IsEncapsulated() {
			p.CheckErr(p.B(name+".UnmarshalUnionNDR", "ctx", "w", switchVar))
		} else {
			p.CheckErr(p.B(name+".UnmarshalNDR", "ctx", "w"))
		}

	case scopes.Is(midl.TypeStruct), scopes.Is(midl.TypeInterface):

		// marshal structure.

		typeName := p.GoScopeTypeName(ctx, p.Scope(), field, scopes, true)

		p.If(name, "==", "nil", func() {
			p.P(name, "=", p.Amp(typeName), "{}")
		})

		p.CheckErr(p.B(name+".UnmarshalNDR", "ctx", "w"))

	case scopes.Is(midl.TypePointer):

		fN := "_ptr_" + field.Name

		// render deferred unmarshalling function.
		p.P(fN, ":=", "ndr.UnmarshalNDRFunc", "(", "func(ctx context.Context, w ndr.Reader) error {")
		p.GenFieldUnmarshalNDR(ctx, field, scopes.Next(), index...)
		p.P("return nil")
		p.P("})")

		// render setter function.
		sN := "_s" + "_" + field.Name
		decl := p.GoScopeTypeName(ctx, p.Scope(), field, scopes)
		p.P(sN, ":=", "func(ptr interface{}) {", name, "=", "*ptr.(*", decl, ")", "}")

		// read the pointer.
		p.CheckErr(p.B("w.ReadPointer", p.Amp(name), sN, fN))

	case scopes.Is(midl.TypeArray):

		// marshal array.

		isConformant, isVarying := scopes.IsConformant(), scopes.IsVarying()

		dim := 0
		for scopes := scopes; scopes.IsNotFixedArray(); scopes = scopes.Next() {
			dim++
		}

		if field.Position == 0 || scopes.IsTopLevelArray() {

			if dim := scopes.Dim(); dim.IsString && !dim.NoSizeLimit && dim.Size().Empty() && dim.LengthIs.Empty() {

				// unmarshal string using std-library.

				switch scopes.Next().Kind() {
				case midl.TypeWChar, midl.TypeUint16, midl.TypeInt16:
					if dim.IsNullTerminated {
						p.CheckErr(p.B("ndr.ReadUTF16NString", "ctx", "w", p.Amp(name)))
					} else {
						p.CheckErr(p.B("ndr.ReadUTF16String", "ctx", "w", p.Amp(name)))
					}
				case midl.TypeUChar, midl.TypeChar, midl.TypeUint8, midl.TypeInt8:
					if dim.IsNullTerminated {
						p.CheckErr(p.B("ndr.ReadCharNString", "ctx", "w", p.Amp(name)))
					} else {
						p.CheckErr(p.B("ndr.ReadCharString", "ctx", "w", p.Amp(name)))
					}
				}
				break
			}

			if isConformant {
				// initialize sizeInfo.
				p.Block("sizeInfo := []uint64", func() {
					for i := 0; i < dim; i++ {
						p.P("0", ",")
					}
				})

				p.Range("sz1", "sizeInfo", func() {
					// read max count.
					p.GenReadSize(ctx, p.IVar("sizeInfo", "sz1"))
				})
			}
		}

		// XXX: for non-wire marshaling try to set size-info.
		if isConformant && !isVarying && !IsOp(ctx) && !IsReserved(ctx) {
			if sz := scopes.Dim().Size(); !sz.Empty() {
				// check minimum value borders.
				szVar := p.IVar("sizeInfo", scopes.Dim().Dimension)
				p.P("//", "XXX: for opaque unmarshaling")
				if to, val, ok := p.GetMinVarValue(ctx, sz.Is()); ok {
					exprTo := p.GenExpr(ctx, to, field, "uint64", "_to")
					p.If(exprTo /*to.Expression(p.EO) */, ">", val, "&&", szVar, "==", "0", func() {
						exprSz := p.GenExpr(ctx, sz.Is(), field, "", "_sz")
						p.P(szVar, "=", p.B("uint64", exprSz /* sz.Expression(p.EO)*/))
					})
				} else {
					exprSz := p.GenExpr(ctx, sz.Is(), field, "uint64")
					p.If(exprSz /* sz.Expression(p.EO) */, ">", "0", "&&", szVar, "==", "0", func() {
						p.P(szVar, "=", p.B("uint64", exprSz /* sz.Expression(p.EO)*/))
					})
				}
			}
		}

		if isVarying && !p.IsNotLastFixedArray(ctx, scopes, field) {

			if !isConformant {
				// initialize sizeInfo.
				p.Block("sizeInfo := []uint64", func() {
					for i := 0; i < dim; i++ {
						p.P("0", ",")
					}
				})
			}

			p.Range("sz1", "sizeInfo", func() {
				// read offset.
				p.GenReadSize(ctx, p.IVar("sizeInfo", "sz1"))
				// read actual count.
				p.GenReadSize(ctx, p.IVar("sizeInfo", "sz1"))
			})
		}

		idx := p.Var("i", len(index)+1)

		origName, decl := name, p.GoScopeTypeName(ctx, p.Scope(), field, scopes)

		if scopes.Dim().IsString {

			name = p.BufVar(name)
			ctx, index = WithVarName(ctx, name), nil

			// replace decl with buffer type name.
			switch scopes := scopes.Next(); scopes.Kind() {
			case midl.TypeWChar, midl.TypeUint16, midl.TypeInt16:
				decl = "[]uint16"
			case midl.TypeChar, midl.TypeUChar, midl.TypeInt8, midl.TypeUint8:
				decl = "[]byte"
			}

			p.P("var", name, decl)
		}

		if scopes.Dim().NoSizeLimit {
			p.Block("for", idx+" := 0;", "w.Len() > 0;", idx+"++", func() {
				p.P(idx, ":=", idx)
				p.P(name, "=", p.B("append", name, p.GoTypeZeroValue(ctx, p.Scope(), field, scopes.Next())))
				p.GenFieldUnmarshalNDR(ctx, field, scopes.Next(), append(index, idx)...)
			})

			if scopes.Dim().IsString {
				if field.Attrs.Format.MultiSize {
					tmpName := "_tmp" + name
					p.P(tmpName, ":=", p.B("string", p.B("utf16.Decode", name)))
					p.If(tmpName, "=", p.B("strings.TrimRight", tmpName, "ndr.ZeroString"), ";", tmpName, "!=", `""`, func() {
						p.P(origName, "=", p.B("strings.Split", tmpName, "ndr.ZeroString"))
					})
					break
				}

				// convert utf16/byte array back to string.
				switch scopes := scopes.Next(); scopes.Kind() {
				case midl.TypeWChar, midl.TypeUint16, midl.TypeInt16:
					p.P(origName, "=", p.B("strings.TrimRight", p.B("string", p.B("utf16.Decode", name)), "ndr.ZeroString"))
				case midl.TypeChar, midl.TypeUChar, midl.TypeInt8, midl.TypeUint8:
					p.P(origName, "=", p.B("strings.TrimRight", p.B("string", name), "ndr.ZeroString"))
				}
			}
			break
		}

		if scopes.Array().IsFixed() && (!scopes.Dim().IsString || p.IsNotLastFixedArray(ctx, scopes, field)) {
			p.P(name, "=", p.B("make", decl, scopes.Array().Size()))
		} else {
			szVar := p.IVar("sizeInfo", scopes.Dim().Dimension)
			p.If(szVar, ">", p.B("uint64", "w.Len()"), "/* sanity-check */", func() {
				p.P("return", `fmt.Errorf("buffer overflow for size %d of array `+name+`", `, szVar, `)`)
			})
			p.P(name, "=", "make(", decl, ",", szVar, ")")
		}

		p.Range(idx, name, func() {
			// iterate over the array.
			p.P(idx, ":=", idx)
			p.GenFieldUnmarshalNDR(ctx, field, scopes.Next(), append(index, idx)...)
		})

		if scopes.Dim().IsString {

			if field.Attrs.Format.MultiSize {
				tmpName := "_tmp" + name
				p.P(tmpName, ":=", p.B("string", p.B("utf16.Decode", name)))
				p.If(tmpName, "=", p.B("strings.TrimRight", tmpName, "ndr.ZeroString"), ";", tmpName, "!=", `""`, func() {
					p.P(origName, "=", p.B("strings.Split", tmpName, "ndr.ZeroString"))
				})
				break
			}

			// convert utf16/byte array back to string.
			switch scopes := scopes.Next(); scopes.Kind() {
			case midl.TypeWChar, midl.TypeUint16, midl.TypeInt16:
				p.P(origName, "=", p.B("strings.TrimRight", p.B("string", p.B("utf16.Decode", name)), "ndr.ZeroString"))
			case midl.TypeChar, midl.TypeUChar, midl.TypeInt8, midl.TypeUint8:
				p.P(origName, "=", p.B("strings.TrimRight", p.B("string", name), "ndr.ZeroString"))
			}
		}

	case scopes.Is(midl.TypePipe):

		{
			p.If(name, "==", "nil", func() {
				p.P(name, "=", p.B(p.GoScopeTypeNameWithN(ctx, p.Scope(), field, scopes, true, "New")))
			})

			field := p.GenPipeField(ctx, scopes.Type())
			scopes := NewScopes(field.Scopes())

			p.Range(func() {
				p.P("var", "_chunk", p.GoFieldTypeName(ctx, NewScopes(field.Scopes()).Scope(), field))
				p.GenFieldUnmarshalNDR(WithVarName(ctx, "_chunk"), field, scopes)
				p.If(p.Len("_chunk"), "==", "0", "/* end */", func() {
					p.P("break")
				})

				p.P(name + ".Append(_chunk)")
			})
		}

	default:
		p.P("//", "FIXME:", "unknown type", field.Name)
	}

	if scopes.Is(midl.TypeArray) {

		if field.Attrs.Format.Hex {
			varName := "_hex_" + field.Name
			p.P(varName, ",", "err", ":=", p.B(p.AddImport(HexImport)+"."+"DecodeString", p.B("string", name)))
			p.If("err", "!=", "nil", func() {
				p.P("return", "err")
			})
			p.P(name, "=", varName)
		}

		if len(field.Attrs.Layout) > 0 {
			for _, field := range field.Attrs.Layout {
				fN := "_layout_" + field.Name
				p.P(fN, ":=", "ndr.UnmarshalNDRFunc", "(", "func(ctx context.Context, w ndr.Reader) error {")
				p.GenFieldUnmarshalNDR(ctx, field, NewScopes(field.Scopes()), index...)
				p.P("return nil")
				p.P("})")
				p.If(p.Len(name), ">", "0", func() {
					p.CheckErr(p.B(p.B("w.WithBytes", name)+".Unmarshal", "ctx", fN))
				})
			}
		}
	}

	if IsReserved(ctx) && (scopes.IsBool() && midl.PrimitiveTypeSize(scopes.Kind()) > 1) {
		// XXX: guard.
		p.P("_", "=", name)
	}

}

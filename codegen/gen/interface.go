package gen

import (
	"context"
	"fmt"
	"regexp"

	"github.com/oiweiwei/go-msrpc/midl"
)

var (
	notUsedOpRe = regexp.MustCompile("(NotUsedOnWire)|(NotImplemented)|(LSA_DP)|(LSA_TM)|(Unused)|(Opnum\\d+Reserved)|(Reserved_Opnum\\d+)|(Reserved\\d+)|(Opnum\\d+NotUsedByProtocol)")
)

func (p *Generator) IsUnusedOp(n string) bool {
	return notUsedOpRe.MatchString(n)
}

func (p *Generator) GenServerHandle(ctx context.Context, iff *midl.Interface) {

	n := GoName(iff.Name) + "Server"

	p.P()
	p.Block("func", p.B("Register"+n, "conn dcerpc.Conn", "o "+n, "opts ...dcerpc.Option"), func() {
		p.P(p.B("conn.RegisterServer", p.B("New"+n+"Handle", "o"),
			p.B("append", "opts", p.B("dcerpc.WithAbstractSyntax", p.SyntaxName(ctx, iff)))+"..."))

	})

	p.P()
	p.Block("func", p.B("New"+n+"Handle", "o "+n), "dcerpc.ServerHandle", func() {
		p.Block("return", p.B("func", "ctx context.Context", "opNum int", "r ndr.Reader"), p.B("", "dcerpc.Operation", "error"), func() {
			p.P("return", p.B(n+"Handle", "ctx", "o", "opNum", "r"))
		})
	})

	p.P()
	p.Block("func", p.B(n+"Handle", "ctx context.Context", "o "+n, "opNum int", "r ndr.Reader"), p.B("", "dcerpc.Operation", "error"), func() {

		first := 0
		if len(iff.Body.Operations) > 0 {
			first = iff.Body.Operations[0].OpNum
		}

		if first > 0 {
			p.If("opNum", "<", first, func() {
				p.P("//", iff.BaseName, "base method.")
				p.P("return", p.B(p.GoInterfaceTypeName(ctx, iff.Base, GoName(iff.BaseName)+"ServerHandle"), "ctx", "o", "opNum", "r"))
			})
		}

		p.Block("switch opNum", func() {
			for _, op := range iff.Body.Operations {
				p.P("case", op.OpNum, ":", "//", op.Name)
				if p.IsUnusedOp(op.Name) {
					p.P("//", op.Name)
					p.P("return", "nil, nil")
					continue
				}
				p.P("in", ":=", p.Amp(p.OpName(ctx, op, InParam))+"{}")
				p.If("err := in.UnmarshalNDR(ctx, r);", "err != nil", func() {
					p.P("return", "nil, err")
				})
				p.P("resp, err", ":=", p.B("o."+p.MethodName(ctx, op), "ctx", "in"))
				p.P("return", "resp."+p.XXX()+"ToOp(ctx)", ", err")
			}
		})

		p.P("return", "nil, nil")
	})
}

func (p *Generator) GenServerInterface(ctx context.Context, iff *midl.Interface) {

	n := GoName(iff.Name) + "Server"
	baseN := GoName(iff.BaseName) + "Server"

	p.P()
	p.P("//", iff.Name, "server interface.")
	p.Block("type", n, "interface", func() {
		if iff.IsObject() && iff.Base != nil {
			p.P()
			p.P("//", iff.BaseName, "base class.")
			p.P(p.GoInterfaceTypeName(ctx, iff.Base, baseN))
		}
		for _, op := range iff.Body.Operations {
			// gen operation doc.
			doc, ok := p.Doc.Type(op.Name)
			if ok {
				p.P()
				p.GenComment(ctx, doc.Doc)
			} else {
				p.P()
				p.P("//", op.Name, "operation.")
			}
			for _, field := range []string{"Return Values", "Exceptions Thrown"} {
				if doc, ok := doc.Field(field); ok {
					p.P("//")
					p.GenComment(ctx, doc.Doc)
				}
			}
			if p.IsUnusedOp(op.Name) {
				p.P("//", p.MethodName(ctx, op))
				continue
			}

			p.P(p.B(p.MethodName(ctx, op),
				"context.Context",
				"*"+p.OpName(ctx, op, InParam),
			),
				p.B("", "*"+p.OpName(ctx, op, OutParam), "error"))
		}
	})
}

func (p *Generator) GenClient(ctx context.Context, iff *midl.Interface) {

	n := GoName(iff.Name) + "Client"
	baseN := GoName(iff.BaseName) + "Client"

	dn := p.XXX() + "Default" + n

	p.P()
	p.Structure(dn, func() {
		if iff.IsObject() && iff.Base != nil {
			p.P(p.GoInterfaceTypeName(ctx, iff.Base, baseN))
		}
		p.P("cc", "dcerpc.Conn")
		if iff.IsObject() {
			p.P("ipid", "*dcom.IPID")
		}
	})

	if iff.IsObject() && iff.Base != nil {
		p.Block("func", "(o *"+dn+")", p.B(GoName(iff.BaseName)), p.GoInterfaceTypeName(ctx, iff.Base, baseN), func() {
			p.P("return", p.O(GoName(iff.BaseName)+"Client"))
		})
	}

	for _, op := range iff.Body.Operations {

		if p.IsUnusedOp(op.Name) {
			continue
		}

		p.P()
		p.P("func", "(o *"+dn+")",
			p.B(p.MethodName(ctx, op),
				"ctx context.Context", "in *"+p.OpName(ctx, op, InParam), "opts ...dcerpc.CallOption",
			),
			p.B("",
				"*"+p.OpName(ctx, op, OutParam), "error",
			), "{")
		p.P("op", ":=", p.B("in."+p.XXX()+"ToOp", "ctx"))
		if iff.IsObject() {
			p.If("_, ok := dcom.HasIPID(opts); !ok", func() {
				p.If("o.ipid != nil", func() {
					p.P("opts = append(opts, dcom.WithIPID(o.ipid))")
				}, p.Else(func() {
					p.P("return", "nil, fmt.Errorf(\"%s: ipid is missing\", op.OpName())")
				}))
			})
		}
		p.If("err", ":=", p.B("o.cc.Invoke", "ctx", "op", "opts..."), ";", "err != nil", func() {
			p.P("return", "nil, err")
		})
		p.P("out", ":=", p.Amp(p.OpName(ctx, op, OutParam)+"{}"))
		p.P(p.B("out."+p.XXX()+"FromOp", "ctx", "op"))

		if !op.Type.Is(midl.TypeVoid) {

			param := p.OperationReturnValue(ctx, op)

			field := &midl.Field{
				Name:  param.Name,
				Type:  param.Type,
				Attrs: param.Attrs.FieldAttr,
			}

			p.If("op."+GoName(field.Name), "!=", p.GoTypeZeroValue(ctx, nil, field, NewScopes(field.Scopes())), func() {
				p.P("return", "out", ",", p.B("fmt.Errorf", `"%s: %w"`, "op.OpName()", p.B("errors.New", "ctx", "op."+GoName(field.Name))))
			})
		}

		p.P("return", "out, nil")
		p.P("}")
	}

	p.P()
	p.Block("func", "(o *"+dn+")", "AlterContext(ctx context.Context, opts ...dcerpc.Option)", "error", func() {
		p.P("return o.cc.AlterContext(ctx, opts...)")
	})

	if iff.IsObject() {
		p.P()
		p.Block("func", "(o *"+dn+")", "IPID(ctx context.Context, ipid *dcom.IPID)", n, func() {
			p.If("ipid == nil", func() {
				p.P("ipid = &dcom.IPID{}")
			})
			p.Block("return", "&", dn, func() {
				if iff.Base != nil {
					p.P(GoName(iff.BaseName)+"Client", ":", p.O(GoName(iff.BaseName)+"Client"+"."+p.B("IPID", "ctx", "ipid")), ",")
				}
				p.P("cc", ":", "o.cc,")
				p.P("ipid", ":", "ipid,")
			})
		})
	}

	p.Block("func", p.B("New"+n, "ctx context.Context", "cc dcerpc.Conn", "opts ...dcerpc.Option"), p.B("", n, "error"), func() {

		if iff.IsObject() {

			p.P("var", "err", "error")

			p.If("!dcom.IsSuperclass(opts)", func() {
				p.P("cc, err", "=", p.B("cc.Bind", "ctx", p.B("append", "opts", p.B("dcerpc.WithAbstractSyntax", p.SyntaxName(ctx, iff)))+"..."), ";")
				p.If("err != nil", func() {
					p.P("return", "nil, err")
				})
			})

			if iff.Base != nil {
				p.P("base, err", ":=", p.B(p.GoInterfaceTypeName(ctx, iff.Base, "New"+baseN),
					"ctx", "cc", p.B("append", "opts", p.B("dcom.Superclass", "cc"))+"..."))
				p.If("err != nil", func() {
					p.P("return", "nil, err")
				})
			}

			if iff.Base != nil {
				p.P("ipid, ok := dcom.HasIPID(opts)")
				p.If("ok", func() {
					p.P("base", "=", p.B("base.IPID", "ctx", "ipid"))
				})
			} else {
				p.P("ipid, _ := dcom.HasIPID(opts)")
			}

			p.Block("return", "&", dn, func() {
				if iff.Base != nil {
					p.P(GoName(iff.BaseName)+"Client", ":", "base", ",")
				}
				p.P("cc", ":", "cc,")
				p.P("ipid", ":", "ipid,")
			}, ",", "nil")

		} else {
			p.P("cc, err", ":=", p.B("cc.Bind", "ctx", p.B("append", "opts", p.B("dcerpc.WithAbstractSyntax", p.SyntaxName(ctx, iff)))+"..."), ";")
			p.If("err != nil", func() {
				p.P("return", "nil, err")
			})
			p.P("return", p.Amp(dn+"{cc: cc}"), ",", "nil")
		}
	})
}

func (p *Generator) GenClientInterface(ctx context.Context, iff *midl.Interface) {
	n := GoName(iff.Name) + "Client"
	baseN := GoName(iff.BaseName) + "Client"
	p.P()
	p.P("//", iff.Name, "interface.")
	p.Block("type", n, "interface", func() {
		if iff.IsObject() && iff.Base != nil {
			// p.P()
			// p.P("//", iff.BaseName, "base class.")
			// p.P(p.GoInterfaceTypeName(ctx, iff.Base, baseN))
			p.P()
			p.P("//", iff.BaseName, "retrieval method.")
			p.P(p.B(GoName(iff.BaseName)), p.GoInterfaceTypeName(ctx, iff.Base, baseN))
		}
		for _, op := range iff.Body.Operations {
			// gen operation doc.
			doc, ok := p.Doc.Type(op.Name)
			if ok {
				p.P()
				p.GenComment(ctx, doc.Doc)
			} else {
				p.P()
				p.P("//", op.Name, "operation.")
			}
			for _, field := range []string{"Return Values", "Exceptions Thrown"} {
				if doc, ok := doc.Field(field); ok {
					p.P("//")
					p.GenComment(ctx, doc.Doc)
				}
			}
			if p.IsUnusedOp(op.Name) {
				p.P("//", p.MethodName(ctx, op))
				continue
			}

			p.P(p.B(p.MethodName(ctx, op),
				"context.Context",
				"*"+p.OpName(ctx, op, InParam),
				"...dcerpc.CallOption",
			),
				p.B("", "*"+p.OpName(ctx, op, OutParam), "error"))
		}

		p.P()
		p.P("//", "AlterContext alters the client context.")
		p.P(p.B("AlterContext", "context.Context", "...dcerpc.Option"), "error")

		if iff.IsObject() {
			p.P()
			p.P("//", "IPID sets the object interface identifier.")
			p.P(p.B("IPID", "context.Context", "*dcom.IPID"), p.B("", n))
		}
	})
}

func (p *Generator) SyntaxName(ctx context.Context, iff *midl.Interface) string {
	n := GoName(iff.Name) + "Syntax"
	ver := iff.Attrs.Version
	if ver == nil {
		ver = &midl.Version{}
	}
	return fmt.Sprintf("%sV%d_%d", n, ver.Major, ver.Minor)
}

func (p *Generator) SyntaxNameWithImport(ctx context.Context, iff *midl.Interface) string {
	n := p.GoInterfaceTypeName(ctx, iff, "") + "Syntax"
	ver := iff.Attrs.Version
	if ver == nil {
		ver = &midl.Version{}
	}
	return fmt.Sprintf("%sV%d_%d", n, ver.Major, ver.Minor)
}

func (p *Generator) GenInterfaceID(ctx context.Context, iff *midl.Interface) {

	if iff.Attrs.UUID == nil {
		return
	}

	n := GoName(iff.Name) + "Syntax"

	p.P()
	p.P("var", "(")
	if iff.IsObject() {
		p.GoPackageName(ctx, NewScopes(midl.LookupType("IID").Scopes()))
		p.P("//", iff.Name, "interface", "identifier", iff.Attrs.UUID)
		p.P(GoName(iff.Name)+"IID", "=", "&dcom.IID", UUIDToGUID(iff.Attrs.UUID))
	}

	p.P("//", "Syntax UUID")
	p.P(n+"UUID", "=", fmt.Sprintf("%#v", iff.Attrs.UUID))
	ver := iff.Attrs.Version
	if ver == nil {
		ver = &midl.Version{Major: 0, Minor: 0}
	}
	p.P("//", "Syntax ID")
	p.P(n+fmt.Sprintf("V%d_%d", ver.Major, ver.Minor), "=", "&dcerpc.SyntaxID{",
		"IfUUID:", n+"UUID", ",", "IfVersionMajor:", ver.Major, ",", "IfVersionMinor:", ver.Minor, "}")
	p.P(")")
}

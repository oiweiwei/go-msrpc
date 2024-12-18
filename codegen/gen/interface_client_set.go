package gen

import (
	"context"
	"path/filepath"

	"github.com/oiweiwei/go-msrpc/midl"
)

func (p *Generator) GenClientSet(ctx context.Context, f *midl.File) {

	n := "Client"

	if !f.PkgIs("dcom") {
		p.AddImport(Import{
			Name:  "dcom_client",
			Path:  filepath.Join(p.ImportsPath, "dcom", "client"),
			Guard: "dcom_client.GoPackage",
		})
	}
	p.AddImport(Import{
		Name:  "dcom",
		Path:  filepath.Join(p.ImportsPath, "dcom"),
		Guard: "dcom.GoPackage",
	})

	var dcom *midl.File

	for _, f := range midl.Files() {
		if f.PkgIs("dcom") {
			dcom = f
			break
		}
	}

	p.P("//", f.GoPkg, "client set.")
	p.Block("type", n, "interface", func() {
		if !f.PkgIs("dcom") {
			p.P()
			p.P("//", "DCOM common interfaces")
			for _, iff := range dcom.Interfaces {
				if iff.Body.Operations != nil && iff.Base != nil {
					p.P(GoName(iff.Name), "()", p.GoInterfaceTypeName(ctx, iff, "")+"Client")
				}
			}
		}
		p.P()
		p.P("//", "Package specific interfaces")
		for _, iff := range f.Interfaces {
			if iff.Body.Operations != nil && iff.Base != nil {
				p.P(GoName(iff.Name), "()", p.GoInterfaceTypeName(ctx, iff, "")+"Client")
			}
		}
		p.P("//", "AlterContext alters the client context.")
		p.P(p.B("AlterContext", "context.Context", "...dcerpc.Option"), "error")
		p.P()
		p.P("//", "Conn returns the client connection (unsafe)")
		p.P(p.B("Conn"), "dcerpc.Conn")
		p.P()
		p.P("//", "IPID sets the object interface identifier.")
		p.P(p.B("IPID", "context.Context", "*dcom.IPID"), p.B("", n))
	})

	dn := p.XXX() + "Default" + n
	p.Block("type", dn, "struct", func() {
		p.P("cc", "dcerpc.Conn")
		p.P()
		if !f.PkgIs("dcom") {
			p.P("dcomClient", "dcom_client.Client")
			p.P()
		}
		for _, iff := range f.Interfaces {
			if iff.Body.Operations != nil && iff.Base != nil {
				p.P(GoNamePrivate(iff.Name), p.GoInterfaceTypeName(ctx, iff, "")+"Client")
			}
		}
	})

	if !f.PkgIs("dcom") {
		for _, iff := range dcom.Interfaces {
			if iff.Body.Operations != nil && iff.Base != nil {
				p.P()
				p.P("func", "(o *"+dn+")", GoName(iff.Name), "()", p.GoInterfaceTypeName(ctx, iff, "")+"Client", "{")
				p.P("return", p.O("dcomClient"+"."+GoName(iff.Name)), "()")
				p.P("}")
			}
		}
	}

	for _, iff := range f.Interfaces {
		if iff.Body.Operations != nil && iff.Base != nil {
			p.P()
			p.P("func", "(o *"+dn+")", GoName(iff.Name), "()", p.GoInterfaceTypeName(ctx, iff, "")+"Client", "{")
			p.P("return", p.O(GoNamePrivate(iff.Name)))
			p.P("}")
		}
	}

	p.P()

	p.Block("func", "NewClient", "(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option)", p.B("", n, "error"), func() {

		p.P()
		p.P("opts", "=", "append", "(", "opts", ",")
		for _, iff := range f.Interfaces {
			if iff.Body.Operations != nil && iff.Base != nil {
				p.P(p.B("dcerpc.WithAbstractSyntax", p.SyntaxNameWithImport(ctx, iff)), ",")
			}
		}
		if !f.PkgIs("dcom") {
			for _, iff := range dcom.Interfaces {
				if iff.Body.Operations != nil && iff.Base != nil {
					p.P(p.B("dcerpc.WithAbstractSyntax", p.SyntaxNameWithImport(ctx, iff)), ",")
				}
			}
		}
		p.P(")")
		p.P()
		p.P("cc, err", ":=", p.B("cc.Bind", "ctx", "opts..."), ";")
		p.If("err != nil", func() {
			p.P("return", "nil, err")
		})
		p.P()
		p.P("o", ":=", p.Amp(dn+"{cc: cc}"))
		p.P()
		if !f.PkgIs("dcom") {
			p.P("dcomClient", ",", "err", ":=", p.B("dcom_client.NewClient", "ctx", "cc", p.B("append", "opts", p.B("dcerpc.WithNoBind", "cc"))+"..."))
			p.If("err != nil", func() {
				p.P("return", "nil, err")
			})
			p.P("o.dcomClient", "=", "dcomClient")
		}
		p.P()
		p.P("sub, ok := cc.(dcerpc.SubConn)")
		p.If("!ok", func() {
			p.P("return", "nil, fmt.Errorf(\"sub-conn is not supported\")")
		})
		p.P()
		for _, iff := range f.Interfaces {
			if iff.Body.Operations != nil && iff.Base != nil {
				p.P()
				iffName := GoNamePrivate(iff.Name)
				p.P(iffName+"SubConn, err", ":=", p.B("sub.SubConn", "ctx", p.SyntaxNameWithImport(ctx, iff)))
				p.If("err != nil", func() {
					p.P("//", "XXX: use main subconnection as a last resort")
					p.P("//", "it was noticed that we can reuse the main connection for dcom interfaces")
					p.P(iffName+"SubConn", "=", "sub")
				})
				p.P()
				p.P("o."+GoNamePrivate(iff.Name), ",", "err", "=", p.B(p.GoInterfaceTypeName(ctx, iff, "New"+GoName(iff.Name))+"Client", "ctx", iffName+"SubConn", p.B("append", "opts", p.B("dcerpc.WithNoBind", iffName+"SubConn"))+"..."))
			}
		}

		p.P("return", "o", ",", "nil")
	})

	p.P()
	p.Block("func", "(o *"+dn+")", "AlterContext(ctx context.Context, opts ...dcerpc.Option)", "error", func() {
		p.P("return o.cc.AlterContext(ctx, opts...)")
	})
	p.P()
	p.Block("func", "(o *"+dn+")", "Conn()", "dcerpc.Conn", func() {
		p.P("return o.cc")
	})
	p.P()
	p.Block("func", "(o *"+dn+")", "IPID(ctx context.Context, ipid *dcom.IPID)", n, func() {
		p.If("ipid == nil", func() {
			p.P("ipid = &dcom.IPID{}")
		})
		p.Block("return", "&", dn, func() {
			if !f.PkgIs("dcom") {
				p.P("dcomClient", ":", p.O("dcomClient."+p.B("IPID", "ctx", "ipid")), ",")
			}
			for _, iff := range f.Interfaces {
				if iff.Body.Operations != nil && iff.Base != nil {
					p.P(GoNamePrivate(iff.Name), ":", p.O(GoNamePrivate(iff.Name)+"."+p.B("IPID", "ctx", "ipid")), ",")
				}
			}
			p.P("cc", ":", "o.cc,")
		})
	})

}

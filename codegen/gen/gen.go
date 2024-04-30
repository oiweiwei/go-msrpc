package gen

import (
	"context"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/oiweiwei/go-msrpc/codegen/doc"
	"github.com/oiweiwei/go-msrpc/midl"
	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

func FieldName(i int, n string) string {
	if n == "" {
		return "Field" + strconv.Itoa(i)
	}
	return n
}

type GeneratorOptions struct {
	UnmarshalRequest  bool
	UnmarshalResponse bool
}

type Generator struct {
	out         *FileBuffer
	inCB        bool // in comment block.
	Trace       bool
	MultiFile   bool
	Format      bool
	ImportsPath string
	Dir         string
	Doc         *doc.Doc
	Cache       string
	Opts        GeneratorOptions
	CheckErr    func(...interface{})
	Files       []*FileBuffer
}

func trimExt(p string) string {
	return strings.TrimSuffix(p, filepath.Ext(p))
}

func (p *Generator) Reset(ctx context.Context, source interface{}, n ...string) {

	if p.out != nil {
		p.Files = append(p.Files, p.out)
	}

	if source == nil {
		return
	}

	switch source := source.(type) {
	case *midl.File:
		path := trimExt(source.Path)
		p.out = NewFileBuffer(path, source.GoPkg)
		if len(n) == 0 {
			p.out.IsRoot = true
		}
	case *midl.Interface:
		path := filepath.Join(trimExt(File(ctx).Path), strings.ToLower(source.Name), source.Attrs.Version.String())
		p.out = NewFileBuffer(path, strings.ToLower(source.Name), n...)
	}
}

func (p *Generator) P(args ...interface{}) {
	if p.inCB {
		fmt.Fprintln(p.out.Out, append([]interface{}{"//"}, args...)...)
	} else {
		fmt.Fprintln(p.out.Out, args...)
	}
}

func (p *Generator) T(args ...interface{}) {
	if p.Trace {
		fmt.Fprintln(p.out.Out, args...)
	}
}

func MSUnion(ctx context.Context) bool {
	iff := Interface(ctx)
	if iff != nil {
		return iff.Attrs.MSUnion
	}
	return false
}

func PointerDefault(ctx context.Context) midl.PointerType {
	iff := Interface(ctx)
	if iff != nil && iff.Attrs.PointerDefault != midl.PointerTypeNone {
		return iff.Attrs.PointerDefault
	}
	return midl.PointerTypeUnique
}

func (p *Generator) Gen(ctx context.Context, fn string) error {

	p.CheckErr = CheckErr1(p)

	f, err := midl.NewFile(fn, "").Load()
	if err != nil {
		return err
	}

	p.Doc, err = doc.BuildDoc(fn, doc.WithCache(p.Cache))
	if err != nil {
		return err
	}

	ctx = WithFile(ctx, f)

	p.Reset(ctx, f)

	for _, sym := range f.Exports() {
		if sym.Const != nil {
			p.GenConst(ctx, sym.Const)
		}
	}

	for _, sym := range f.Exports() {
		if sym.Type != nil {
			p.GenType(ctx, sym.Type)
		}
	}

	if len(f.ComClasses) > 0 {
		_ = p.GoPackageName(ctx, NewScopes(midl.LookupType("CLSID").Scopes()))
	}

	for _, sym := range f.ComClasses {
		if sym.Attrs.UUID != nil {
			p.P("//", sym.Name, "class", "identifier", sym.Attrs.UUID)
			p.P("var", GoName(sym.Name)+"ClassID", "=", "&dcom.ClassID", UUIDToGUID(sym.Attrs.UUID))
		}
	}

	for _, iff := range f.Interfaces {

		ctx := WithInterface(ctx, iff)

		p.Reset(ctx, iff)

		p.GenInterfaceID(ctx, iff)
		p.GenClientInterface(ctx, iff)

		for _, sym := range iff.Exports() {
			if sym.Const != nil {
				p.GenConst(ctx, sym.Const)
			}
		}

		for _, sym := range iff.Exports() {
			if sym.Type != nil {
				p.GenType(ctx, sym.Type)
			}
		}

		if iff.Attrs.UUID == nil {
			p.P("//", "XXX: no declaration")
			continue
		}

		p.GenClient(ctx, iff)

		for _, op := range iff.Body.Operations {
			p.GenOperation(ctx, op)
		}

		p.Reset(ctx, iff, "server")

		p.GenServerInterface(ctx, iff)
		p.GenServerHandle(ctx, iff)
	}

	p.Reset(ctx, nil)

	for _, file := range p.Files {

		b := file.Reset()
		p.out = file

		if p.out.IsRoot {
			p.P("//", "The", p.out.PackageName, "package", "implements", "the", strings.ToUpper(p.out.PackageName), "client", "protocol.")
			if intro, ok := p.Doc.Type("Introduction"); ok {
				p.P("//")
				p.P("//", "# Introduction")
				p.P("//")
				p.GenComment(ctx, intro.Doc)
			}
			if overview, ok := p.Doc.Type("Overview (synopsis)"); ok {
				p.P("//")
				p.P("//", "# Overview")
				p.P("//")
				p.GenComment(ctx, overview.Doc)
			}
		}

		p.GenPackage(ctx, f)
		p.GenImports(ctx)
		p.GenImportGuards(ctx)

		if file.FileName == "" {
			p.GenPackageImportGuard(ctx)
		}

		b = append(file.Reset(), b...)
		if p.Format {
			if b, err = format.Source(b); err != nil {
				return err
			}
		}

		p := filepath.Join(p.Dir, strings.ReplaceAll(file.Path, "-", "_"))
		fmt.Fprintln(os.Stderr, "create dir", p)

		if err := os.MkdirAll(p, 0777); err != nil {
			return err
		}

		base := file.FileName
		if base == "" {
			base = filepath.Base(p)
		}

		fileName := filepath.Join(p, base) + ".go"

		fmt.Fprintln(os.Stderr, "create file", fileName)

		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		fmt.Fprintf(f, "%s", string(b))

		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (p *Generator) XXX() string {
	return "xxx_"
}

func (p *Generator) GenConst(ctx context.Context, c *midl.Const) {
	p.P()
	p.P("//", GoName(c.Name), "represents the", c.Name, "RPC constant")
	p.P("var", GoName(c.Name), "=", c.Value.Expression(GoName))
}

func (p *Generator) GenType(ctx context.Context, t *midl.Type) {

	pp := p.NewTypeGenerator(ctx, t)

	if pp.Skip() {
		return
	}

	switch {
	case pp.Is(midl.TypeStruct):
		pp.GenStruct(ctx)
	case pp.Is(midl.TypeCUnion):
		pp.GenCUnion(ctx)
	case pp.Is(midl.TypeUnion):
		pp.GenUnion(ctx)
	case pp.Is(midl.TypeEnum):
		pp.GenEnum(ctx)
	case pp.Is(midl.TypeInterface):
		pp.GenInterface(ctx)
	}
}

func (p *Generator) GenDoAlignmentMarshalNDR(ctx context.Context, alignment int) {
	if alignment > 1 {
		p.CheckErr(p.B("w.WriteAlign", alignment))
	}
}

func (p *Generator) GenDoAlignmentUnmarshalNDR(ctx context.Context, alignment int) {
	if alignment > 1 {
		p.CheckErr(p.B("w.ReadAlign", alignment))
	}
}

/*
When no pointer attribute is associated with a pointer that is a member
of a structure or union, the MIDL compiler assigns pointer attributes
using the following priority rules (1 is highest):

1. Attributes explicitly applied to the pointer type
2. Attributes explicitly applied to the pointer parameter or member
3. (?) The pointer_default attribute in the IDL file that defines the type
4. The pointer_default attribute in the IDL file that imports the type
5. Unique (Microsoft RPC default mode)
*/
func (p *Generator) GenPointerType(ctx context.Context, scopes *Scopes) midl.PointerType {

	var ptr midl.PointerType

	if scopes.IsTopLevelPointer() {
		if ptr = scopes.Scope().Pointer; ptr != midl.PointerTypeNone {
			return ptr
		}
	}

	if iff := Interface(ctx); iff != nil {
		if ptr = iff.Attrs.PointerDefault; ptr != midl.PointerTypeNone {
			return ptr
		}
	}

	return midl.PointerTypeUnique
}

func (p *Generator) GenWriteSize(ctx context.Context, sz interface{}) {
	p.CheckErr(p.B("w.WriteSize", sz))
}

func (p *Generator) CB(f func()) {
	p.inCB = true
	f()
	p.inCB = false
}

func (p *Generator) GenComment(ctx context.Context, docs []string) {
	p.CB(func() {
		for _, doc := range docs {
			for _, ddoc := range strings.Split(doc, "\n") {
				ddoc = strings.ReplaceAll(ddoc, "… ", "... ")
				if strings.Contains(ddoc, "|") || strings.Contains(ddoc, "+---") {
					p.P("\t", ddoc)
					continue
				}
				line := ""
				for _, word := range strings.Split(ddoc, " ") {
					line += word + " "
					if len(line) > 80 {
						p.P(line)
						line = ""
					}
				}
				if len(line) > 0 {
					p.P(line)
				}
			}
			p.P()
		}
	})
}

func (p *Generator) GenTag(ctx context.Context, field *midl.Field) string {
	tag := []string{`name:` + field.Name}

	for _, size := range []struct {
		E []midl.Expr
		N string
	}{
		{field.Attrs.SizeIs, "size_is"},
		{field.Attrs.MaxIs, "max_is"},
		{field.Attrs.MinIs, "min_is"},
		{field.Attrs.LengthIs, "length_is"},
		{field.Attrs.FirstIs, "first_is"},
		{field.Attrs.LastIs, "last_is"},
	} {
		if size.E == nil {
			continue
		}

		szs := []string{}
		for _, sz := range size.E {
			szs = append(szs, sz.Expression())
		}
		tag = append(tag, size.N+":("+strings.Join(szs, ", ")+")")
	}

	if field.Attrs.Usage.IsString {
		tag = append(tag, "string")
	}

	if !field.Attrs.SwitchIs.Empty() {
		tag = append(tag, "switch_is:"+field.Attrs.SwitchIs.String())
	}

	if field.Attrs.Pointer != midl.PointerTypeNone && field.Attrs.Pointer != midl.PointerTypeRefWeak {
		tag = append(tag, "pointer:"+field.Attrs.Pointer.String())
	}

	n := GoSnakeCase(field.Name)
	if n == "_" {
		return "`idl:\"" + strings.Join(tag, ";") + "\"`"
	}

	return "`idl:\"" + strings.Join(tag, ";") + "\" json:\"" + n + "\"`"
}

func UUIDToGUID(u *uuid.UUID) string {

	return fmt.Sprintf("{Data1: 0x%08x, Data2: 0x%04x, Data3: 0x%04x, Data4: "+
		"[]byte{0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x}}",
		u.TimeLow, u.TimeMid, u.TimeHiAndVersion,
		u.ClockSeqHiAndReserved, u.ClockSeqLow, u.Node[0], u.Node[1], u.Node[2], u.Node[3], u.Node[4], u.Node[5])

}

package gen

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/oiweiwei/go-msrpc/codegen/go_names"
	"github.com/oiweiwei/go-msrpc/midl"
)

var (
	GoName           = go_names.GoName
	GoNamePrivate    = go_names.GoNamePrivate
	GoNameNoReserved = go_names.GoNameNoReserved
	GoMergeNames     = go_names.GoMergeNames
	GoSnakeCase      = go_names.GoSnakeCase
	LexName          = go_names.GoLexName
	Escape           = go_names.Escape
	RPCName          = go_names.Unescape
	Title            = go_names.Title
)

func GoHex(n int) func(uint64) string {
	return func(v uint64) string {
		return fmt.Sprintf("0x%0*X", n*2, v)
	}
}

var defaultVersion = &midl.Version{}

func (p *Generator) GoInterfaceTypeName(ctx context.Context, iff *midl.Interface, n string) string {

	if n == "" {
		n = GoName(iff.Name)
	}

	for _, f := range midl.Files() {

		if _, ok := f.LookupType(iff.Name); !ok {
			continue
		}

		pkg, ver := filepath.Join(f.GoPkg, strings.ToLower(iff.Name)), iff.Attrs.Version
		if ver == nil {
			ver = defaultVersion
		}

		pkgName := lastPart(pkg)

		p.AddImport(Import{
			Name:  pkgName,
			Path:  filepath.Join(p.ImportsPath, pkg, ver.String()),
			Guard: pkgName + "." + "GoPackage",
		})

		return pkgName + "." + n
	}

	return n
}

func (p *Generator) GoScopeTypeName(ctx context.Context, attr *midl.TypeAttr, field *midl.Field, scopes *Scopes, trim ...bool) string {
	n := p.GoTypeName(ctx, attr, field, scopes)
	if len(trim) > 0 {
		return strings.TrimLeft(n, "*")
	}
	return n
}

func (p *Generator) GoFieldTypeName(ctx context.Context, attr *midl.TypeAttr, field *midl.Field) string {
	return p.GoTypeName(ctx, attr, field, NewScopes(field.Scopes()))
}

func (p *Generator) EmbeddedTypeName(ctx context.Context, attrs *midl.TypeAttr, field *midl.Field) string {
	return GoName(attrs.Alias) + "_" + GoFieldName(field)
}

func (p *Generator) GoTypeName(ctx context.Context, attrs *midl.TypeAttr, field *midl.Field, scopes *Scopes) string {

	var ret string

go_type_name_loop:
	for scopes := scopes; scopes != nil; scopes = scopes.Next() {

		switch {
		case scopes.Is(midl.TypeArray):

			if scopes.Dim().IsString {
				// string is a terminal type.
				ret += "string"

				if field.Attrs.Format.MultiSize {
					ret = "[]" + ret
				}

				break go_type_name_loop
			}

			// array declaration. add [] if it is not a string.
			ret = "[]" + ret
			continue

		case scopes.Is(midl.TypePointer):
			// pointer declaration. noop.
			continue

		case scopes.Is(midl.TypeVoid):
			ret = "[]byte"
			break go_type_name_loop
		}

		// primitive type.
		if scopes.Type().IsPrimitiveType() {
			if scopes.IsBool() {
				ret += "bool"
				break go_type_name_loop
			}

			if field.Attrs.Format.Rune {
				ret += "rune"
				break go_type_name_loop
			}

			ret += GoPrimitiveTypeName(scopes.Kind())
			// terminal type.
			break go_type_name_loop
		}

		// constructed (exported type) type.

		if !scopes.Is(midl.TypeEnum) {
			// always with pointer.
			ret += "*"
		}

		// lookup type name.
		name := LookupName(ctx, scopes)
		if name == "" {
			ret += p.EmbeddedTypeName(ctx, attrs, field)
		} else {
			ret += p.GoPackageName(ctx, scopes) + name
		}

		break go_type_name_loop
	}

	if ret == "byte" {
		ret = "uint8"
	}

	if strings.Contains(ret, "[]") {
		ret = strings.ReplaceAll(ret, "uint8", "byte")
	}

	return ret
}

func (p *Generator) GoPackageName(ctx context.Context, scopes *Scopes) string {

	n := scopes.Alias()
	if n == "" {
		n = scopes.Type().TypeName()
	}

	if n == "" {
		panic("empty name")
	}

	pkg, ver, ok := GoPackage(ctx, n)
	if !ok {
		panic("unknown type " + n)
	}

	if pkg == "" {
		return ""
	}

	if strings.Contains(pkg, "guiddef") {
		pkg = "dtyp"
	}

	pkgName := lastPart(pkg)

	if ver != nil {
		p.AddImport(Import{
			Name:  pkgName,
			Path:  filepath.Join(p.ImportsPath, pkg, ver.String()),
			Guard: pkgName + "." + "GoPackage",
		})
	} else {
		p.AddImport(Import{
			Name:  pkgName,
			Path:  filepath.Join(p.ImportsPath, pkg),
			Guard: pkgName + "." + "GoPackage",
		})
	}

	return pkgName + "."
}

func lastPart(n string) string {
	parts := strings.Split(n, "/")
	return parts[len(parts)-1]
}

func GoPackage(ctx context.Context, n string) (string, *midl.Version, bool) {

	var defaultVersion = &midl.Version{}

	if iff := Interface(ctx); iff != nil {

		// we are rendering interface.

		if _, ok := iff.Body.Export[n]; ok {
			// type is local to interface.
			return "", nil, true
		}

		for _, iff := range File(ctx).Interfaces {
			// type is in some other interface.
			// (in the same file).
			if _, ok := iff.Body.Export[n]; ok {

				if iff.Attrs.Version == nil {
					return filepath.Join(File(ctx).GoPkg, strings.ToLower(iff.Name)), defaultVersion, true
				}

				return filepath.Join(File(ctx).GoPkg, strings.ToLower(iff.Name)), iff.Attrs.Version, true
			}
		}

		if File(ctx).IsLocal(n) {
			// type is file-level definition.
			return File(ctx).GoPkg, nil, true
		}
	} else {
		if _, ok := File(ctx).Export[n]; ok {
			return "", nil, true
		}
	}

	for _, f := range midl.Files() {
		if _, ok := f.Export[n]; ok {
			return f.GoPkg, nil, true
		}
		for _, iff := range f.Interfaces {
			if _, ok := iff.Body.Export[n]; ok {
				// type is local to interface.
				if iff.Attrs.Version == nil {
					return filepath.Join(f.GoPkg, strings.ToLower(iff.Name)), defaultVersion, true
				}
				return filepath.Join(f.GoPkg, strings.ToLower(iff.Name)), iff.Attrs.Version, true
			}
		}
	}

	return "", nil, false
}

func LookupName(ctx context.Context, scopes *Scopes) string {

	if alias := scopes.Alias(); alias != "" {
		if strings.HasPrefix(alias, "PDNS") {
			// FIXME: dnsp.idl
			alias = scopes.Scope().Names[0]
		}
		return GoName(alias)
	}

	// find alias based on tagged type.
	for _, alias := range File(ctx).LookupAlias(scopes.Type().TypeName()) {
		if alias != "" {
			return GoName(alias)
		}
	}

	return ""
}

func GoPrimitiveTypeName(t midl.Kind) string {

	switch t {
	case midl.TypeWChar:
		return "uint16"
	case midl.TypeUChar, midl.TypeChar:
		return "byte"
	case midl.TypeError:
		return "uint32"
	case midl.TypeInt32_64:
		return "int64"
	case midl.TypeUint32_64:
		return "uint64"
	}

	return t.String()
}

func GoFieldName(field *midl.Field) string {
	return GoName(FieldName(field.Position, field.Name))
}

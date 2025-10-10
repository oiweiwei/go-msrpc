package gen

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl"
)

func (p *Generator) GenPackageImportGuard(ctx context.Context) {
	p.P("")
	p.P("var (")
	p.P("//", "import guard")
	p.P("GoPackage", "=", p.Q(File(ctx).GoPkg))
	p.P(")")
}

func (p *Generator) GetImports() []Import {

	std, local := []Import{}, []Import{}

	for _, imp := range p.out.Imports {
		if imp.Path == "" {
			std = append(std, imp)
		} else {
			local = append(local, imp)
		}
	}

	sort.Slice(std, func(i, j int) bool {
		return std[i].Name < std[j].Name
	})

	sort.Slice(local, func(i, j int) bool {
		return local[i].Name < local[j].Name
	})

	if len(local) == 0 {
		return std
	}

	return append(append(std, Import{}), local...)
}

type Import struct {
	Name, Path, Guard string
}

var (
	HexImport = Import{
		Name:  "encoding/hex",
		Guard: "hex.DecodeString",
	}

	SyncImport = Import{
		Name:  "sync",
		Guard: "&sync.Mutex{}",
	}

	DefaultImports = []Import{
		{
			Name:  "context",
			Guard: "context.Background",
		},
		{
			Name:  "fmt",
			Guard: "fmt.Errorf",
		},
		{
			Name:  "unicode/utf16",
			Guard: "utf16.Encode",
		},
		{
			Name:  "strings",
			Guard: "strings.TrimPrefix",
		},
		{
			Name:  "ndr",
			Guard: "ndr.ZeroString",
			Path:  "github.com/oiweiwei/go-msrpc/ndr",
		},
		{
			Name:  "uuid",
			Guard: "(*uuid.UUID)(nil)",
			Path:  "github.com/oiweiwei/go-msrpc/midl/uuid",
		},
		{
			Name:  "dcerpc",
			Guard: "(*dcerpc.SyntaxID)(nil)",
			Path:  "github.com/oiweiwei/go-msrpc/dcerpc",
		},
		{
			Name:  "errors",
			Guard: "(*errors.Error)(nil)",
			Path:  "github.com/oiweiwei/go-msrpc/dcerpc/errors",
		},
	}
)

func (p *Generator) GenPackage(ctx context.Context, f *midl.File) {
	p.P("package", p.out.PackageName)
}

func (p *Generator) GenImports(ctx context.Context) {
	p.P("")
	p.P("import", "(")
	for _, imp := range p.GetImports() {
		if imp.Name == "" {
			p.P("")
			continue
		}
		if imp.Path != "" {
			p.P(imp.Name, p.Q(imp.Path))
		} else {
			p.P(p.Q(imp.Name))
		}
	}
	p.P(")")
}

func (p *Generator) GenImportGuards(ctx context.Context) {
	p.P("")
	p.P("var", "(")
	for _, imp := range p.out.Imports {
		p.P("_", "=", imp.Guard)
	}
	p.P(")")
}

func (p *Generator) AddImport(imp Import) string {

	for _, imps := range p.out.Imports {
		if imps.Name == imp.Name {
			return imps.Name[strings.LastIndex(imps.Name, "/")+1:]
		}
	}
	fmt.Println("[ADD IMPORT]", imp)
	p.out.Imports = append(p.out.Imports, imp)
	return imp.Name[strings.LastIndex(imp.Name, "/")+1:]
}

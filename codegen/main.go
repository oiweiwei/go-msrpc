package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"context"

	"github.com/oiweiwei/go-msrpc/midl"

	"github.com/oiweiwei/go-msrpc/codegen/gen"
)

var (
	fn    string
	I     string
	j     bool
	dir   string
	frmt  bool
	trace bool
	cache string
)

func init() {
	flag.BoolVar(&frmt, "format", true, "use format")
	flag.BoolVar(&trace, "trace", false, "use format")
	flag.StringVar(&fn, "f", "", "filename")
	flag.StringVar(&I, "I", "github.com/oiweiwei/go-msrpc/example", "import path")
	flag.StringVar(&dir, "dir", "msrpc/", "the generation dir")
	flag.StringVar(&cache, "doc-cache", ".cache/doc/", "the cache directory for doc")
	flag.BoolVar(&j, "j", false, "json output")
	flag.Parse()
}

func main() {

	if j {
		file, err := midl.NewFile(fn, "").Load()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		b, _ := json.MarshalIndent(file, "", "  ")
		fmt.Fprintln(os.Stdout, string(b))
		os.Exit(0)
	}

	midl.RPCErrorVerbose = true

	p := &gen.Generator{
		ImportsPath: I,
		Format:      frmt,
		Trace:       trace,
		Dir:         dir,
		Cache:       cache,
	}
	if err := p.Gen(context.Background(), fn); err != nil {
		fmt.Fprintf(os.Stderr, "gen: %v\n", err)
		os.Exit(1)
	}
}

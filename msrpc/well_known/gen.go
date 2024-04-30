//go:build exclude

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"strings"

	"encoding/csv"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

var (
	file string
	pkg  string
	out  string
)

func init() {
	flag.StringVar(&file, "f", "./data/well_known_endpoints.tsv", "source file")
	flag.StringVar(&pkg, "pkg", "well_known", "go package name")
	flag.StringVar(&out, "o", "./well_known_endpoints.go", "go output file")
	flag.Parse()
}

type Descriptor struct {
	UUID     uuid.UUID
	Domain   string
	Details  string
	Name     string
	Version  dcerpc.Version
	Endpoint string
}

func (d Descriptor) UUIDName() string {

	n := ""
	if d.Domain != "" {
		for _, part := range strings.Split(d.Domain, "-") {
			n += part
		}
	}

	n += strings.Title(d.Name)
	return n
}

func main() {

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	descs := []*Descriptor{}

	r := csv.NewReader(f)

	r.Comma = '\t'

	for {
		line, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		desc := &Descriptor{
			Domain:   line[1],
			Details:  line[2],
			Name:     line[3],
			Endpoint: line[5],
		}

		u, err := uuid.Unmarshal(line[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "uuid", err)
			os.Exit(1)
		}

		desc.UUID = *u

		if _, err = fmt.Sscanf(line[4], "v%d.%d", &desc.Version.Major, &desc.Version.Minor); err != nil {
			fmt.Fprintln(os.Stderr, "scanf", err)
			os.Exit(1)
		}

		descs = append(descs, desc)
	}

	b := bytes.NewBuffer(nil)
	p := func(args ...interface{}) { fmt.Fprintln(b, args...) }

	p("package", pkg)
	p()
	p("import", "(")
	p(fmt.Sprintf("%q", "github.com/oiweiwei/go-msrpc/midl/uuid"))
	p(")")
	p()
	p("var", "(")
	for _, desc := range descs {
		p(desc.UUIDName(), "=", fmt.Sprintf("%#v", desc.UUID))
	}
	p(")")

	p()
	p("type", "UUID", "uuid.UUID")
	p()
	p("func", "(u UUID)", "Describe", "()", "string", "{")
	p("switch", "(uuid.UUID)(u)", "{")
	for _, desc := range descs {
		p("case", desc.UUIDName(), ":")
		p("return", fmt.Sprintf(`"%s: %s: %s"`, desc.Domain, desc.Details, desc.Name))
	}
	p("}")
	p("return", `""`)

	p("}")
	p()
	p("func", "(u UUID)", "WellKnownEndpoint", "()", "[]string", "{")
	p("switch", "(uuid.UUID)(u)", "{")
	for _, desc := range descs {
		if desc.Endpoint != "" {
			p("case", desc.UUIDName(), ":")
			p("return", fmt.Sprintf("%#+v", strings.Split(strings.ReplaceAll(desc.Endpoint, "named_pipe:", ""), " ")))
		}
	}
	p("}")
	p("return", "nil")
	p("}")

	p("func", "(u UUID)", "Name", "()", "string", "{")
	p("switch", "(uuid.UUID)(u)", "{")
	for _, desc := range descs {
		if desc.Name != "" {
			p("case", desc.UUIDName(), ":")
			p("return", fmt.Sprintf("%q", desc.Name))
		}
	}
	p("}")
	p("return", `""`)
	p("}")

	source, err := format.Source(b.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, "format", err)
		os.Exit(1)
	}

	if f, err = os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		fmt.Fprintln(os.Stderr, "format", err)
		os.Exit(1)
	}

	fmt.Fprintf(f, "%s", string(source))

}

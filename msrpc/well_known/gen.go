//go:build exclude

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"gopkg.in/yaml.v3"
)

var (
	file string
	pkg  string
	out  string
)

func init() {
	flag.StringVar(&file, "f", "./data/well_known_endpoints.yaml", "source file")
	flag.StringVar(&pkg, "pkg", "well_known", "go package name")
	flag.StringVar(&out, "o", "./well_known_endpoints.go", "go output file")
	flag.Parse()
}

type YAMLEntry struct {
	UUID        string   `yaml:"uuid"`
	Domain      string   `yaml:"domain"`
	Annotation  string   `yaml:"annotation"`
	IfName      string   `yaml:"if_name"`
	IfVersions  []string `yaml:"if_versions"`
	Endpoints   []string `yaml:"endpoints"`
	Keywords    []string `yaml:"keywords"`
}

type Descriptor struct {
	UUID      uuid.UUID
	Domain    string
	Annotation string
	Name      string
	Endpoints []string
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

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var entries []YAMLEntry
	if err = yaml.Unmarshal(data, &entries); err != nil {
		fmt.Fprintln(os.Stderr, "yaml:", err)
		os.Exit(1)
	}

	descs := make([]*Descriptor, 0, len(entries))
	for _, e := range entries {
		u, err := uuid.Unmarshal(e.UUID)
		if err != nil {
			fmt.Fprintln(os.Stderr, "uuid:", err)
			os.Exit(1)
		}
		descs = append(descs, &Descriptor{
			UUID:       *u,
			Domain:     e.Domain,
			Annotation: e.Annotation,
			Name:       e.IfName,
			Endpoints:  e.Endpoints,
		})
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
		p("return", fmt.Sprintf(`"%s: %s: %s"`, desc.Domain, desc.Annotation, desc.Name))
	}
	p("}")
	p("return", `""`)
	p("}")

	p()
	p("func", "(u UUID)", "WellKnownEndpoint", "()", "[]string", "{")
	p("switch", "(uuid.UUID)(u)", "{")
	for _, desc := range descs {
		if len(desc.Endpoints) > 0 {
			p("case", desc.UUIDName(), ":")
			p("return", fmt.Sprintf("%#v", desc.Endpoints))
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
		fmt.Fprintln(os.Stderr, "format:", err)
		os.Exit(1)
	}

	out_f, err := os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "open:", err)
		os.Exit(1)
	}

	fmt.Fprintf(out_f, "%s", string(source))
}

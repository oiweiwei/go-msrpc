//go:build exclude

// This example generates the MOF file from the wmio binary data.
package main

/*
	for f in data/*.dump; do go run wmio.go -f "$f" > "${f/dump/mof}" && clang-format -i "${f/dump/mof}"; done
*/

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmio"
)

var (
	f string
	j bool
	v bool
)

func init() {
	flag.StringVar(&f, "f", "data/wmio_win32_process.dump", "cim object hexstream")
	flag.BoolVar(&j, "j", false, "json-format")
	flag.BoolVar(&v, "v", false, "validate encoding")
	flag.Parse()
}

func main() {

	_, p, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintln(os.Stderr, "cannot determine current dir")
		return
	}

	hexdata, err := os.ReadFile(filepath.Join(filepath.Dir(p), f))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	b, err := hex.DecodeString(strings.TrimSpace(string(hexdata)))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	obj, err := wmio.Unmarshal(b)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if v {
		b, err := wmio.Marshal(obj)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		obj, err = wmio.Unmarshal(b)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	jj := func(d any) string { b, _ := json.MarshalIndent(d, "", "  "); return string(b) }

	if j {
		fmt.Fprintln(os.Stdout, jj(obj))
	} else {
		if obj.Decoration != nil {
			fmt.Println("//", "@namespace", `"`+obj.Decoration.Namespace+`"`)
			fmt.Println("//", "@servername", `"`+obj.Decoration.ServerName+`"`)
		}
		if obj.Instance != nil {
			ClsToPrint(obj.Instance.CurrentClass, wmio.Methods{}, wmio.Class{})
			InsToPrint(*obj.Instance)
		}
		if obj.Class != nil {
			ClsToPrint(obj.Class.ParentClass, obj.Class.ParentClassMethods, wmio.Class{})
			ClsToPrint(obj.Class.CurrentClass, obj.Class.CurrentClassMethods, obj.Class.ParentClass)
		}
	}

}

func InsToPrint(ins wmio.Instance) {

	P := fmt.Println

	qs, _ := QualToString(ins.Qualifiers)
	P(qs)

	sort.Slice(ins.Properties, func(i, j int) bool {
		return ins.Properties[i].Order < ins.Properties[j].Order
	})

	P("instance", "of", ins.ClassName, "{")
	for _, p := range ins.Properties {
		if p.Value.Type.IsArray() {
			P(PropToString(p), "=", "{", p.Value, "}", ";")
		} else {
			P(PropToString(p), "=", p.Value, ";")
		}
	}
	P("};")

}

func ClsToPrint(cls wmio.Class, methods wmio.Methods, parent wmio.Class) {

	P := fmt.Println

	qs, _ := QualToString(cls.Qualifiers)

	P(qs)

	sort.Slice(cls.Properties, func(i, j int) bool {
		return cls.Properties[i].Order < cls.Properties[j].Order
	})

	if len(cls.Derivation) > 0 {
		P("class", cls.Name, ":", strings.Join(cls.Derivation, " "), "{")
	} else {
		P("class", cls.Name, "{")
	}
	for _, p := range cls.Properties {
		P()
		if int(p.ClassOfOrigin) < len(cls.Derivation) {
			P("//", "@inherited", cls.Derivation[p.ClassOfOrigin])
		}
		if !p.Nullable && !p.InheritDefault {
			if p.Value.Type.IsArray() {
				P(PropToString(p), "=", "{", p.Value, "}", ";")
			} else {
				P(PropToString(p), "=", p.Value, ";")
			}
		} else {
			P(PropToString(p), ";")
		}
	}

	for _, m := range methods.Methods {
		qs, _ := QualToString(m.Qualifiers)
		P()
		P(qs)

		var out []*wmio.Property

		if m.OutputSignature.Class != nil {
			out = m.OutputSignature.Class.CurrentClass.Properties
		}

		sort.Slice(out, func(i, j int) bool {
			return out[i].Order < out[j].Order
		})

		ret := "void"
		for _, p := range out {
			if p.Name == "ReturnValue" {
				_, ret = QualToString(p.Qualifiers)
			}
		}

		var in []*wmio.Property

		if m.InputSignature.Class != nil {
			in = m.InputSignature.Class.CurrentClass.Properties
		}

		sort.Slice(in, func(i, j int) bool {
			return in[i].Order < in[j].Order
		})

		P(ret, m.Name, "(")

		sep := ""
		for _, p := range in {
			if p.InheritDefault {
				P("//", "@default", p.Value)
			}
			P(sep)
			sep = ","
			P(PropToString(p))
		}

		for _, p := range out {
			if p.Name == "ReturnValue" {
				continue
			}
			if p.InheritDefault {
				P("//", "@default", p.Value)
			}
			P(sep)
			sep = ","
			P(PropToString(p))
		}

		P(");")
	}
	P("};")
}

func QualToString(qs []*wmio.Qualifier) (string, string) {

	var tq string

	qq := []*wmio.Qualifier{}
	for _, q := range qs {
		if q.Name == "CIMTYPE" {
			tq = q.Value.Value.(string)
			continue
		}
		qq = append(qq, q)
	}

	sb := strings.Builder{}

	sep := ""
	if len(qq) > 1 {
		sb.WriteString("[")
		for _, q := range qq {
			if q.Name != "CIMTYPE" {
				sb.WriteString(sep)
				sep = ", "
				sb.WriteString(q.String())
			}
		}
		sb.WriteString("] ")
	}

	return sb.String(), tq[strings.Index(tq, ":")+1:]

}

func PropToString(p *wmio.Property) string {

	qs, tq := QualToString(p.Qualifiers)

	sb := strings.Builder{}

	sb.WriteString(qs)
	sb.WriteString(" ")
	sb.WriteString(tq)
	sb.WriteString(" ")
	sb.WriteString(p.Name)

	return sb.String()
}

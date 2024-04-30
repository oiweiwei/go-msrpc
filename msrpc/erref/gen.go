//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	out string
	pkg string
	url string
)

func init() {
	flag.StringVar(&out, "o", "-", "output file")
	flag.StringVar(&pkg, "pkg", "", "output file")
	flag.StringVar(&url, "url", "", "target url")
	flag.Parse()
}

type Err struct {
	Code    string
	Name    string
	Details string
}

func Text(p *goquery.Selection) string {

	s := strings.Split(p.Text(), "\n")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}

	return strings.TrimSpace(strings.Join(s, " "))
}

func main() {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	first := true

	ret := []*Err{}
	e := &Err{}

	doc.Find("table tbody tr").Each(func(_ int, tr *goquery.Selection) {
		tr.Find("td > p").Each(func(i int, p *goquery.Selection) {
			switch txt := Text(p); i {
			case 0:
				if !first {
					ret, e = append(ret, e), &Err{}
				}
				first = false
				e.Code = txt
				parts := strings.Split(txt, " ")
				if len(parts) == 1 {
					break
				}
				e.Code, txt = parts[0], parts[1]
				fallthrough
			case 1:
				if e.Name == "" {
					e.Name = txt
					break
				}
				fallthrough
			case 2:
				if e.Details == "" {
					e.Details = txt
					break
				}
				e.Details += "; " + Text(p)
			}
		})
	})

	ret = append(ret, e)

	pp := &P{}
	p := pp.P

	p("package", pkg)
	p()

	p("var (")
	for _, e := range ret {
		p(GoName(e.Name), "=", "&Error", "{", e.Code, ",", fmt.Sprintf("%q", e.Name), ",", fmt.Sprintf("%q", e.Details), "}")
	}
	p(")")

	p()

	seen := map[string]struct{}{}

	p("func", "FromCode(code uint32)", "error", "{")
	p()
	p("if", "code", "==", "0", "{")
	p("return", "nil")
	p("}")
	p()
	p("switch", "code", "{")
	for _, e := range ret {

		if _, ok := seen[e.Code]; ok {
			continue
		}

		p("case", e.Code, ":")
		p("return", GoName(e.Name))

		seen[e.Code] = struct{}{}
	}
	p("}")
	p("return", "nil")
	p("}")

	b, err := format.Source(pp.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	f := os.Stdout

	switch out {
	case "-", "stdout":
	default:
		f, err = os.OpenFile(out, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	fmt.Fprintln(f, string(b))
}

type P struct {
	bytes.Buffer
}

func (p *P) P(args ...interface{}) {
	fmt.Fprintln(&p.Buffer, args...)
}

func GoName(n string) (ret string) {
	for _, part := range strings.Split(n, "_") {
		ret += strings.Title(strings.ToLower(part))
	}
	return ret
}

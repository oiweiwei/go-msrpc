//go:build ignore

package main

import (
	"bytes"
	"context"
	"encoding/asn1"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

var (
	u       string
	cache   string
	dump    bool
	backoff time.Duration
	mask    = make(map[string]struct{})
)

func init() {
	flag.StringVar(&u, "url", "https://learn.microsoft.com/en-us/windows/win32/adschema/", "the base url")
	flag.StringVar(&cache, "cache", ".cache/adschema.json", "the cache path")
	flag.BoolVar(&dump, "dump", false, "dump the attributes")
	flag.DurationVar(&backoff, "backoff", 500*time.Millisecond, "the backoff factor")
	var attrs string
	flag.StringVar(&attrs, "attrs", "", "the attributes to dump")

	flag.Parse()

	if attrs != "" {
		for _, attr := range strings.Split(attrs, ",") {
			mask[attr] = struct{}{}
		}
	}
}

type Attrs map[string]map[string]string

func (a Attrs) Keys() []string {
	keys := make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (a Attrs) Get(n string) (Attr, bool) {

	var err error

	mattr, ok := a[n]
	if !ok {
		return Attr{}, false
	}

	attr := Attr{
		CN:          mattr["CN"],
		LDAPName:    mattr["Ldap-Display-Name"],
		Description: mattr["Description"],
		Syntax:      mattr["Syntax"],
	}

	for _, v := range strings.Split(mattr["Attribute-Id"], ".") {
		iv, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return Attr{}, false
		}
		attr.AttributeID = append(attr.AttributeID, int(iv))
	}

	if attr.SystemID, err = uuid.Parse(mattr["System-Id-Guid"]); err != nil {
		return Attr{}, false
	}

	return attr, true
}

type Attr struct {
	CN          string                `json:"cn"`
	LDAPName    string                `json:"ldap_name"`
	AttributeID asn1.ObjectIdentifier `json:"attribute_id"`
	SystemID    *uuid.UUID            `json:"system_id"`
	Syntax      string                `json:"syntax"`
	Description string                `json:"description"`
}

func isSelectedAttr(n string) bool {
	if len(mask) == 0 {
		return true
	}
	_, ok := mask[n]
	return ok
}

type P struct {
	bytes.Buffer
}

func (p *P) P(args ...interface{}) {
	fmt.Fprintln(&p.Buffer, args...)
}

func main() {

	ctx := context.Background()

	var (
		attrs Attrs
		err   error
	)

	if dump {
		if attrs, err = queryAttrs(ctx); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if cache != "" {
			if err = dumpAttrs(ctx, cache, attrs); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}

	if attrs == nil {
		if attrs, err = loadAttrs(ctx, cache); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	p := &P{}

	p.P("package", "ad")
	p.P()

	p.P("import", "(")
	p.P(`"encoding/asn1"`)
	p.P()
	p.P(`"github.com/oiweiwei/go-msrpc/midl/uuid"`)
	p.P(")")
	p.P()

	p.P("var attributes = map[string]Attr{")
	for _, a := range attrs.Keys() {
		attr, ok := attrs.Get(a)
		if !ok {
			fmt.Fprintln(os.Stderr, "failed to get", a)
			os.Exit(1)
		}

		p.P(fmt.Sprintf("%q", a), ":", "{")
		p.P("CN:", fmt.Sprintf("%q", attr.CN), ",")
		p.P("LDAPName:", fmt.Sprintf("%q", attr.LDAPName), ",")
		p.P("AttributeID:", fmt.Sprintf("%#v", attr.AttributeID), ",")
		p.P("SystemID:", fmt.Sprintf("%#v", attr.SystemID), ",")
		p.P("Syntax:", fmt.Sprintf("Syntax{Name: %q, Converter: %s}", attr.Syntax, GetConverterSyntax(attr.Syntax)), ",")
		p.P("Description:", fmt.Sprintf("%q", attr.Description), ",")
		p.P("},")
	}
	p.P("}")

	b, err := format.Source(p.Buffer.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, string(b))

}

func GetConverterSyntax(s string) string {

	switch s {
	case "String(Unicode)":
		return "StringUnicode{}"
	case "String(NT-Sec-Desc)":
		return "SecurityDescriptor{}"
	case "Enumeration":
		return "Enumeration{}"
	case "String(Generalized-Time)":
		return "GeneralizedTime{}"
	case "Interval":
		return "Interval{}"
	case "Boolean":
		return "Boolean{}"
	case "Object(DS-DN)":
		return "DSDN{}"
	case "String(Object-Identifier)":
		return "ObjectIdentifier{}"
	case "String(Sid)":
		return "SID{}"
	default:
		return "Raw{}"
	}
}

func loadAttrs(ctx context.Context, from string) (Attrs, error) {

	f, err := os.Open(from)
	if err != nil {
		return nil, err
	}

	var attrs Attrs

	if err := json.NewDecoder(f).Decode(&attrs); err != nil {
		return nil, err
	}

	if len(attrs) == 0 {
		return nil, fmt.Errorf("empty cache")
	}

	return attrs, nil
}

func dumpAttrs(ctx context.Context, to string, attrs Attrs) error {

	if err := os.MkdirAll(filepath.Dir(to), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(attrs)

}

func queryAttrs(ctx context.Context) (Attrs, error) {

	doc, err := goquery.NewDocument(u + "attributes-all")
	if err != nil {
		return nil, err
	}

	attr := Attrs{}
	attrMu := new(sync.Mutex)

	wg := new(sync.WaitGroup)
	in := make(chan string, 8)

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go worker(ctx, wg, in, func(k string, value map[string]string) {
			attrMu.Lock()
			defer attrMu.Unlock()
			attr[k] = value
		})
	}

	doc.Find("div.content ul li a").Each(func(i int, a *goquery.Selection) {
		href, ok := a.Attr("href")
		if !ok || !isSelectedAttr(href) {
			return
		}
		in <- href
	})

	close(in)
	wg.Wait()

	return attr, nil
}

func worker(ctx context.Context, wg *sync.WaitGroup, in <-chan string, setAttr func(string, map[string]string)) {

	defer wg.Done()

	for href := range in {
		attrs := map[string]string{}
		for attempt := 1; ; attempt++ {

			doc, err := goquery.NewDocument(u + href)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			doc.Find("div.content > p").First().Each(func(_ int, p *goquery.Selection) {
				attrs["Description"] = p.Text()
			})

			doc.Find("div table").Each(func(i int, table *goquery.Selection) {
				table.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {
					var key, value string
					tr.Find("td").Each(func(i int, td *goquery.Selection) {
						switch i {
						case 0:
							switch td.Text() {
							case "CN", "Ldap-Display-Name", "Attribute-Id", "System-Id-Guid", "Syntax":
								key = td.Text()
							}
						case 1:
							value = td.Text()
						}
					})
					if key != "" && value != "" {
						attrs[key] = value
					}
				})
			})
			if len(attrs) > 0 {
				fmt.Fprintln(os.Stderr, "done", u+href)
				break
			}
			// backoff.
			time.Sleep(time.Duration(attempt) * backoff)
		}

		setAttr(href, attrs)
	}
}

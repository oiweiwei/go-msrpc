package doc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
	"jaytaylor.com/html2text"

	"github.com/PuerkitoBio/goquery"
)

var (
	LogOutput = os.Stderr
)

const (
	openSpecsURL = "https://learn.microsoft.com/en-us/openspecs/windows_protocols"
)

type Doc struct {
	index *index
	Types []*TypeDoc `json:"types"`
	cache string     `json:"-"`
}

func MergeType(t1, t2 *TypeDoc) *TypeDoc {
	if t1 == nil {
		return t2
	}

	t := &TypeDoc{Name: t1.Name}
	t.Doc = append(t.Doc, t1.Doc...)
	if len(t1.Doc) == 0 || len(t2.Doc) == 0 || t1.Doc[0] != t2.Doc[0] {
		t.Doc = append(t.Doc, t2.Doc...)
	}

	fields := map[string]*FieldDoc{}

	for _, field := range append(append(([]*FieldDoc)(nil), t1.Fields...), t2.Fields...) {

		f, ok := fields[field.Name]
		if !ok {
			f = &FieldDoc{Name: field.Name}
			fields[field.Name] = f
			t.Fields = append(t.Fields, f)
		}

		if len(f.Doc) == 0 || len(field.Doc) == 0 || f.Doc[0] != field.Doc[0] {
			f.Doc = append(f.Doc, field.Doc...)
		}
	}

	return t
}

func (d *Doc) Type(n ...string) (*TypeDoc, bool) {
	if d == nil {
		return nil, false
	}

	var out *TypeDoc

	for _, typ := range d.Types {
		if typ != nil {
			for _, nn := range n {
				if nn == typ.Name {
					out = MergeType(out, typ)
				}
			}
		}
	}

	return out, out != nil
}

func BuildDoc(n string, opts ...func(*Doc)) (*Doc, error) {

	n = strings.TrimSuffix(filepath.Base(n), filepath.Ext(n))
	fmt.Fprintln(LogOutput, "[DEBUG]", "reading", n)

	doc := &Doc{}

	for _, o := range opts {
		o(doc)
	}

	if err := doc.buildDoc(n); err != nil {
		return nil, err
	}

	return doc, nil
}

func WithCache(c string) func(*Doc) {
	return func(d *Doc) {
		d.cache = c
	}
}

func (d *Doc) buildDoc(n string) error {

	var err error

	if d.index, err = buildIndex(n, d.cache); err != nil {
		fmt.Fprintln(LogOutput, "[ERROR]", n, err)
		return nil
	}

	if d.Types, err = d.load(); err != nil {
		return err
	}

	fmt.Fprintln(LogOutput, "[DEBUG]", "found", len(d.Types), "types")

	return nil
}

type task struct {
	name, ref string
}

func (d *Doc) worker(taskChan <-chan *task, outChan chan<- *TypeDoc, done func()) {

	defer done()

	for task := range taskChan {
		td, _ := d.handle(task.name, task.ref)
		outChan <- td
	}
}

func NewDocument(url string) (*goquery.Document, error) {

	openSpecsURL := openSpecsURL

	for i := 0; i < strings.Count(url, "/")-1; i++ {
		openSpecsURL = openSpecsURL[:strings.LastIndex(openSpecsURL, "/")]
	}

	url = MustURLJoinPath(openSpecsURL, url)

	fmt.Fprintln(LogOutput, "[DEBUG]", url)

	for {
		cli := &http.Client{Timeout: 1 * time.Second}
		resp, err := cli.Get(url)
		if err != nil {
			fmt.Fprintln(LogOutput, "[ERROR]", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(LogOutput, "[ERROR]", err)
			continue
		}

		r := bytes.NewReader(bytes.ReplaceAll(body, []byte(`</b>:`), []byte(`:</b>`)))

		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			fmt.Fprintln(LogOutput, "[ERROR]", err)
			continue
		}
		if doc == nil {
			continue
		}

		return doc, nil
	}
}

func (d *Doc) NewDocumentFromCache(url string) (*goquery.Document, error) {

	f, err := os.Open(filepath.Join(d.cache, url))
	if err != nil {
		return nil, err
	}

	defer f.Close()

	node, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromNode(node), nil
}

func (d *Doc) DocumentToCache(url string, nodes []*html.Node) error {

	f, err := os.OpenFile(filepath.Join(d.cache, url), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	w := (io.Writer)(f)

	if _, err := w.Write([]byte(`<div class="content">`)); err != nil {
		return err
	}

	for _, node := range nodes {
		if err := html.Render(w, node); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(`</div>`)); err != nil {
		return err
	}

	return nil
}

func (d *Doc) newDocument(url string) (*goquery.Document, error) {

	if d.cache != "" {
		if doc, err := d.NewDocumentFromCache(url); err != nil {
			fmt.Fprintln(LogOutput, "[INFO]", "cache", err)
		} else {
			return doc, nil
		}
	}

	doc, err := NewDocument(url)
	if err != nil {
		fmt.Fprintln(LogOutput, "[ERROR]", url)
		return nil, err
	}

	if d.cache != "" {
		if err := d.DocumentToCache(url, doc.Find("main > div.content > :not(h1, div, nav)").Nodes); err != nil {
			fmt.Fprintln(LogOutput, "[ERROR]", url)
		}
	}

	nodes := doc.Find("main > div.content").Nodes
	if len(nodes) < 1 {
		return nil, fmt.Errorf("cannot find div.content element")
	}

	return goquery.NewDocumentFromNode(nodes[0]), nil
}

var (
	nameRe = regexp.MustCompile(`([A-Za-z0-9_ \xa0]+)(?:<\d+>)?(?:\((?:(?:variable)|(?:\d+ (?:bytes?|words?)))\))?\s*:`)
)

func parseName(n string) (string, bool) {
	m := nameRe.FindStringSubmatch(n)
	if len(m) > 1 {
		return strings.TrimSpace(m[1]), true
	}
	return n, false
}

func (d *Doc) handle(name, ref string) (*TypeDoc, error) {

	fmt.Fprintln(LogOutput, "[DEBUG]", "fetching", name, ref)

	td := &TypeDoc{Name: name, Ref: ref}

	url := filepath.Join(d.index.n, ref)
	if strings.Contains(ref, "/") {
		url = ref
	}

	doc, err := d.newDocument(url)
	if err != nil {
		fmt.Fprintln(LogOutput, "[ERROR]", name, ref, err)
		return td, nil
	}

	fmt.Fprintln(LogOutput, "[DEBUG]", name, "done")

	doc.Find("div.content > :not(h1, div, nav)").Each(func(_ int, child *goquery.Selection) {
		switch goquery.NodeName(child) {
		case "p":
			// field description or struct description.
			found := false
			child.Find("b").Each(func(_ int, b *goquery.Selection) {
				if name, ok := parseName(String(b.Text())); ok {
					found = true
					td.AddField(strings.TrimSuffix(name, ":"), "")
				}
			})

			if !found {
				child.Find("i").Each(func(_ int, i *goquery.Selection) {
					if name, ok := parseName(String(i.Text())); ok {
						found = true
						td.AddField(strings.TrimSuffix(name, ":"), "")
					}
				})
			}

			txt := String(child.Text())
			if !found {
				switch {
				case strings.HasPrefix(txt, "When processing"),
					strings.HasPrefix(txt, "The following statements define the sequence"),
					strings.HasPrefix(txt, "Message processing for"),
					strings.HasPrefix(txt, "The processing for this method"),
					strings.HasPrefix(txt, "The server MUST"),
					strings.HasPrefix(txt, "The RPC server MUST"),
					strings.HasPrefix(txt, "The behavior required when receiving"),
					strings.HasPrefix(txt, "The following are semantic checks"),
					strings.HasPrefix(txt, "Processing rules"),
					strings.HasPrefix(txt, "The processing rules"),
					strings.HasPrefix(txt, "Processing:"),
					strings.HasPrefix(txt, "Processing instructions:"),
					strings.HasPrefix(txt, "While processing this"),
					strings.HasPrefix(txt, "Sequential Processing Rules:"),
					strings.HasPrefix(txt, "When W32Time"),
					strings.HasPrefix(txt, "The CA server MUST"),
					strings.HasPrefix(txt, "The following processing rules apply"),
					strings.HasPrefix(txt, "In order to perform"),
					strings.HasPrefix(txt, "The following is an overview"):
					td.AddField("Call Processing", "")
				case strings.HasPrefix(txt, "In response"),
					strings.HasPrefix(txt, "The response of the server"),
					strings.HasPrefix(txt, "When a RAZA server receives this message"),
					strings.HasPrefix(txt, "On receipt of this message"),
					strings.HasPrefix(txt, "When this method is"):
					td.AddField("Call Response", "")
				case strings.HasPrefix(txt, "Upon receiving"),
					strings.HasPrefix(txt, "On receiving"),
					strings.HasPrefix(txt, "When the server receives"),
					strings.HasPrefix(txt, "After receiving this message"),
					strings.HasPrefix(txt, "Upon receipt"):
					td.AddField("Call Received", "")
				case strings.Contains(txt, "Exceptions Thrown"):
					td.AddField("Exceptions Thrown", "")
				case strings.Contains(txt, "Server Operations"):
					td.AddField("Server Operations", "")
				case strings.HasPrefix(txt, "The following definitions"),
					strings.HasPrefix(txt, "These definitions"):
					td.AddField("Call Definitions", "")
				case strings.HasPrefix(txt, "The following validation"):
					td.AddField("Call Validations", "")
				case strings.HasPrefix(txt, "The security principal"),
					strings.HasPrefix(txt, "This method obtains the identity"):
					td.AddField("Security", "")
				case strings.HasPrefix(txt, "Sections"):
					td.AddField("Normative", "")
				}
			}
			td.AddDoc(txt)

		case "dl":

			l := ""

			var f func(*html.Node)
			f = func(n *html.Node) {
				if n.Type == html.TextNode {
					l += n.Data
				} else if n.Type == html.ElementNode {
					switch n.Data {
					case "dl", "dt", "dd", "ul", "ol":
						td.AddDoc(String(l))
						l = ""
					case "table":
						td.AddDoc(String(l))
						l = ""
						td.AddDoc(renderTable(&goquery.Selection{
							Nodes: []*html.Node{n},
						}))
						return
					case "img":
						td.AddDoc(String(l))
						l = ""
						td.AddDoc(renderImage(&goquery.Selection{
							Nodes: []*html.Node{n},
						}))
						return
					case "li":
						strL := String(l)
						if strL != "" {
							td.AddDoc("  *  " + String(l))
						}

						l = ""
					}
				}

				if n.FirstChild != nil {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						f(c)
					}
				}
			}

			for _, n := range child.Nodes {
				f(n)
			}

		case "ul":
			txt := renderHTML(child)
			add := false
			for _, line := range strings.Split(txt, "\n") {
				if line = strings.TrimSpace(line); line == "" {
					continue
				}
				if strings.Count(line, "*") == 1 {
					add = true
					continue
				}

				if add {
					add, line = false, "*"+" "+line
					td.AddDoc(line)
				} else {
					td.AddDoc(line)
				}
			}

		case "ol":
			td.AddDoc(renderHTML(child))
		case "table":
			td.AddDoc(renderTable(child))
		}
	})

	return td, nil
}

var (
	workerNum = 16
)

func (d *Doc) load() ([]*TypeDoc, error) {

	ret := []*TypeDoc{}

	if d.cache != "" {
		if err := os.MkdirAll(filepath.Join(d.cache, d.index.n), 0755); err != nil {
			fmt.Fprintln(LogOutput, "[ERROR]", err)
		}
	}

	taskChan, outChan := make(chan *task, workerNum), make(chan *TypeDoc, workerNum)
	wg := new(sync.WaitGroup)

	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go d.worker(taskChan, outChan, wg.Done)
	}

	enum := d.index.enumerate()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(enum); i++ {
			if td := <-outChan; td != nil {
				ret = append(ret, td)
			}
		}
	}()

	for _, nameRef := range enum {
		taskChan <- &task{nameRef[0], nameRef[1]}
	}

	close(taskChan)

	wg.Wait()

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Ref < ret[j].Ref
	})

	return ret, nil
}

func renderImage(img *goquery.Selection) string {
	alt, ok := img.Attr("alt")
	if !ok {
		return ""
	}
	src, ok := img.Attr("src")
	if !ok {
		return ""
	}
	return "[" + alt + "](" + src + ")"
}

func renderHTML(elem *goquery.Selection) string {
	ret, err := html2text.FromHTMLNode(elem.Nodes[0])
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}
	return ret
}

func renderTable(table *goquery.Selection) string {

	table.Find("td").Each(func(_ int, tx *goquery.Selection) {
		// normalize text.
		tx.SetText(String(tx.Text()))
	})

	buf := bytes.NewBuffer(nil)

	html, _ := goquery.OuterHtml(table)

	opts := html2text.Options{
		PrettyTables:        true,
		OmitLinks:           true,
		PrettyTablesOptions: html2text.NewPrettyTablesOptions(),
	}

	opts.PrettyTablesOptions.RowLine = true
	opts.PrettyTablesOptions.ColWidth = 80

	if strings.Contains(html, "colspan=") {

		// render table header separately.
		html, _ = goquery.OuterHtml(table.Find("th").First().Parent())
		html = "<table>" + html + "</table>"

		header, _ := html2text.FromString(html, opts)
		last := ""
		for _, line := range strings.Split(header, "\n") {
			last = line
			fmt.Fprintln(buf, line)
		}

		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if i == 0 {
				return
			}
			tr.Find("td").Each(func(i int, td *goquery.Selection) {
				attr, ok := td.Attr("colspan")
				if !ok {
					return
				}
				colspan, err := strconv.Atoi(attr)
				if err != nil {
					fmt.Fprintf(LogOutput, "colonspan: %v", err)
				}
				txt := String(td.Text())
				repeat := (colspan-1)*4 + 1 - len(txt)
				if repeat <= 0 {
					repeat = 10
				}
				fmt.Fprintf(buf, "| %s %s", txt, strings.Repeat(" ", repeat))
			})
			fmt.Fprintf(buf, "|\n")
			fmt.Fprintln(buf, last)

		})
	} else {
		html, _ = html2text.FromString(html, opts)
		fmt.Fprintln(buf, html)
	}

	return buf.String()
}

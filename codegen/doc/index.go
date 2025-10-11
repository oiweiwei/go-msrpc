package doc

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func init() {

	_, p, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot determine current dir")
	}

	var err error

	if indexes, err = readIndexURLs(filepath.Join(filepath.Dir(p), "data", "index.tsv")); err != nil {
		panic(err)
	}

	if types, err = readUnionURLs(filepath.Join(filepath.Dir(p), "data", "types.tsv")); err != nil {
		panic(err)
	}
}

var (
	indexes [][]string
	types   map[string]map[string][]string
)

func readIndexURLs(p string) ([][]string, error) {

	f, err := os.Open(filepath.Join(filepath.Dir(p), "index.tsv"))
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'

	return r.ReadAll()
}

func readUnionURLs(p string) (map[string]map[string][]string, error) {

	f, err := os.Open(filepath.Join(filepath.Dir(p), "types.tsv"))
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	ret := make(map[string]map[string][]string)

	for _, r := range records {

		if len(r) != 3 {
			continue
		}

		if ret[r[0]] == nil {
			ret[r[0]] = make(map[string][]string)
		}

		r[1] += " " + "structure"

		ret[r[0]][r[1]] = append(ret[r[0]][r[1]], r[2])
	}

	return ret, nil
}

func indexURL(n string) (string, string) {
	n = strings.TrimPrefix(n, "ms-")
	for _, idx := range indexes {
		if idxn := idx[0][strings.LastIndex(idx[0], "/")+1:]; idxn == n || idxn == "mc-"+n {
			if strings.Contains(idx[1], "/") {
				ridx := strings.Split(idx[1], "/")
				return ridx[0], ridx[1]
			}
			if strings.HasPrefix(idxn, "mc-") {
				return "mc-" + n, idx[1]
			}
			if strings.Count(idx[0], "/") > 0 {
				return strings.ReplaceAll(idx[0], n, "ms-"+n), idx[1]
			}
			return "ms-" + n, idx[1]
		}
	}
	return "", ""
}

func buildIndex(n string, cache string) (*index, error) {
	idx := &index{cache: cache}
	_, err := idx.load(n)
	if err != nil {
		return nil, err
	}
	return idx, nil
}

type index struct {
	n     string
	cache string
	idx   map[string]map[string]string
}

func (i *index) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.idx)
}

func (i *index) UnmarshalJSON(b []byte) error {

	var idx map[string]map[string]string

	if err := json.Unmarshal(b, &idx); err != nil {
		return err
	}

	i.idx = idx
	return nil
}

var (
	re = regexp.MustCompile(`^(?:[A-Za-z_0-9]+::)?_?([A-Za-z_0-9]+)(?: [mM]ethod)?(?: \(get\))?(?: \([Oo]pnum:? ?\d+\))? ?(?:[mM]ethod|[sS]tructure|Structurestructure|[eE]numeration|[pP]acket)$`)
)

func (i *index) enumerate() [][2]string {

	ret := make([][2]string, 0, len(i.idx))

	for en, e := range i.idx {
		if m := re.FindStringSubmatch(strings.TrimSpace(strings.ReplaceAll(en, "\n", ""))); len(m) > 1 {
			ret = append(ret, [2]string{m[1], e[en]})
		} else if strings.Contains(en, "Introduction") || strings.Contains(en, "Overview") {
			ret = append(ret, [2]string{en, e[en]})
		}
	}

	return ret
}

func (i *index) toCache() error {

	f, err := os.OpenFile(filepath.Join(i.cache, i.n, "index.json"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(i)
}

func (i *index) fromCache() error {

	f, err := os.Open(filepath.Join(i.cache, i.n, "index.json"))
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewDecoder(f).Decode(i)
}

func (i *index) parse(u string) error {

	if i.cache != "" {
		if err := i.fromCache(); err != nil {
			fmt.Fprintln(LogOutput, "[INFO]", "index cache", err)
		} else {
			return nil
		}
	}

	if i.idx == nil {
		i.idx = make(map[string]map[string]string)
	}

	fmt.Fprintln(LogOutput, "[DEBUG]", "fetching", "index", u)

	doc, err := NewDocument(u)
	if err != nil {
		return fmt.Errorf("parse index: %v", err)
	}

	doc.Find("div p").Each(func(_ int, s *goquery.Selection) {

		txt := String(s.Text())

		if i.idx[txt] == nil {
			i.idx[txt] = make(map[string]string)
		}

		s.Find("a").Each(func(_ int, s *goquery.Selection) {
			if href, _ := s.Attr("href"); href != "" {
				i.idx[txt][String(s.Text())] = href
			}
		})

		if len(i.idx[txt]) == 0 {
			delete(i.idx, txt)
		}
	})

	delete(i.idx, "")

	if i.cache != "" {
		if err := i.toCache(); err != nil {
			fmt.Fprintln(LogOutput, "[INFO]", "index cache", err)
		}
	}

	return nil
}

func (i *index) load(n string) (string, error) {

	var iref string

	if i.n, iref = indexURL(n); n == "" || iref == "" {
		return n, fmt.Errorf("index not found")
	}

	fmt.Fprintln(LogOutput, "[DEBUG]", i.n, iref)

	if err := i.parse(filepath.Join(i.n, iref)); err != nil {
		return n, err
	}

	if t, ok := types[strings.TrimPrefix(i.n, "ms-")]; ok {
		for n, refs := range t {
			i.idx[n] = map[string]string{n: refs[0]}
		}
	}

	return n, nil
}

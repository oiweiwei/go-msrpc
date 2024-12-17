package go_names

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	Title = cases.Title(language.English).String
)

func Escape(n string) string {
	return "%" + n
}

func Unescape(n string) string {
	return strings.TrimPrefix(n, "%")
}

type Namer struct {
	*Config
}

func GoName(n string) string {
	return strings.Join((&Namer{defaultConfig}).GoName(n), "")
}

func GoNamePrivate(n string) string {
	nn := (&Namer{defaultConfig}).GoName(n)
	if len(nn) > 0 {
		nn[0] = strings.ToLower(nn[0])
	}
	ret := strings.Join(nn, "")
	switch ret {
	case "import":
		ret += "1"
	}
	return ret

}

func GoNameNoReserved(n string) string {
	return strings.Join((&Namer{defaultConfig}).GoNameNoReserved(n), "")
}

func GoSnakeCase(n string) string {
	out := (&Namer{defaultConfig}).GoName(n)
	for i := range out {
		out[i] = strings.ToLower(out[i])
	}
	return strings.Join(out, "_")
}

func GoLexName(n string) []string {
	return (&Namer{defaultConfig}).GoName(n)
}

func GoMergeNames(n1, n2 string) string {
	return (&Namer{defaultConfig}).MergeNames(n1, n2)
}

func (nr *Namer) MergeNames(n1, n2 string) string {

	l1, l2 := nr.GoName(n1), nr.GoName(n2)

	i, j := 0, 0
	for i, j = 0, 0; j < len(l1) && i < len(l2); i, j = i+1, j+1 {

		p1, p2 := strings.ToLower(l1[j]), strings.ToLower(l2[i])
		if p1 == p2 {
			continue
		}

		if len(p1) > 2 && len(p2) > 2 && (strings.HasPrefix(p1, p2) || strings.HasPrefix(p2, p1)) {
			continue
		}

		break
	}

	if len(l2[i:]) == 0 {
		i--
	}

	for len(l1) > 0 && i < len(l2) && l1[len(l1)-1] == l2[i] {
		i++
	}

	acronym := ""
	for _, lex := range l1 {
		acronym += strings.ToLower(lex[:1])
	}

	if i < len(l2) && strings.ToLower(l2[i]) == acronym {
		i++
	}

	// typeflagtype
	if len(l1) > 1 && i+1 < len(l2) && strings.ToLower(l2[i]) == strings.ToLower(l1[len(l1)-2]) {
		i++
		for i < len(l2) && l2[i] == l1[len(l1)-1] {
			i++
		}
	}

	return strings.Join(nr.GoName(strings.Join(append(l1, l2[i:]...), "_")), "")
}

func (nr *Namer) GoName(n string) []string {
	return nr.goName(n, true)
}

func (nr *Namer) GoNameNoReserved(n string) []string {
	return nr.goName(n, false)
}

func (nr *Namer) goName(n string, reserved bool) []string {

	if n == "" {
		return nil
	}

	if ret := Unescape(n); ret != n {
		return []string{ret}
	}

	if n, ok := nr.Rename[n]; ok {
		nn := make([]string, len(n))
		copy(nn, n)
		return nn
	}

	for _, prefix := range nr.Trim.PrefixAll {
		if n = strings.TrimPrefix(n, prefix); n == "" {
			n = prefix
		}
	}

	for _, suffix := range nr.Trim.Suffix {
		n = strings.TrimSuffix(n, suffix)
	}

	ret := []string{}

	toks := LexName(n)

	end := ""

	for i, tok := range toks {

		if tok == "" {
			continue
		}

		ltok := strings.ToLower(tok)

		if i == 0 {
			if tok, ok := nr.Rotate[tok]; ok {
				end = tok
				continue
			}
			if _, ok := nr.Trim.Prefix[tok]; ok {
				continue
			}
		}

		if _, ok := nr.Trim.Word[ltok]; ok {
			continue
		}

		if tok, ok := nr.Abbr[ltok]; ok {
			ret = append(ret, tok)
			continue
		}

		ret = append(ret, Title(ltok))
	}

	if reserved {
		if reservedRe.MatchString(strings.Join(ret, "")) {
			ret = []string{"_"}
			return ret
		}
	}

	if end != "" {
	endloop:
		for j := len(ret) - 1; j >= 0; j-- {
			if sizeLength(ret[j]) && sizeLength(end) {
				ret[j], end = end, ""
				break
			}
			for i := len(end); i >= 3 && len(ret) > 0; i-- {
				if strings.HasSuffix(ret[j], end[:i]) {
					ret[j], end = end, ""
					break endloop
				}
			}
		}
		if end != "" {
			ret = append(ret, end)
		}
	}
	if len(ret) == 0 {
		if len(toks) > 0 {
			tok := toks[len(toks)-1]
			if abbr, ok := nr.Abbr[strings.ToLower(tok)]; ok {
				ret = append(ret, abbr)
			} else {
				ret = append(ret, Title(strings.ToLower(tok)))
			}
		}
	}

	if len(ret) > 2 {
		for i := 0; i < len(ret)-1; i++ {
			if ret[i] == "Context" && ret[i+1] == "Handle" {
				ret = append(ret[:i], ret[i+2:]...)
			}
		}
	}

	return ret
}

func sizeLength(s string) bool {
	s = strings.ToLower(s)
	return s == "size" || s == "length"
}

var (
	reservedRe = regexp.MustCompile(`^(Word)?(Reserved?|Unused?|Dummy)(String|Long|Array|MustBeZero|Field|Pad)?(\d+)?(Type|Size)?$`)
)

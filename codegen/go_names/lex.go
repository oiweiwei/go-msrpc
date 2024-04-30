package go_names

import (
	"strings"
	"unicode"
)

func LexName(n string) []string {
	return (Lexer{defaultConfig}).lexName(n)
}

type Lexer struct {
	*Config
}

func (l Lexer) lookBehind(out []string, cur []rune) ([]string, []rune) {

	if len(out) < 1 {
		return out, cur
	}

	lb, ok := l.LookBehind[strings.ToLower(out[len(out)-1])]
	if !ok {
		return out, cur
	}

	ret, ok := lb[strings.ToLower(string(cur))]
	if !ok {
		return out, cur
	}

	if len(ret) == 0 {
		return out[:len(out)-1], nil
	}

	if len(ret) > 1 {
		return append(out[:len(out)-1], ret[:len(ret)-1]...), []rune(ret[len(ret)-1])
	}

	return out[:len(out)-1], []rune(ret[0])
}

func (l *Lexer) lexBetween(cur []rune) []string {
	lb, ok := l.Split[strings.ToLower(string(cur))]
	if !ok {
		return []string{string(cur)}
	}
	return lb
}

func (l Lexer) lexName(n string) []string {

	nr := []rune(n)
	cur := []rune{}
	out := []string{}

	for i := 0; i < len(nr); i++ {

		switch {

		case nr[i] == '_':

			// '_' indicates the gap between lexems.

			if len(cur) > 0 {
				// determine if previous symbol needs to be modified.
				out, cur = l.lookBehind(out, cur)
				// append to the outgoing sequence.
				out, cur = append(out, l.lexBetween(cur)...), nil
			}

			continue

		case unicode.IsUpper(nr[i]):

			// upper-case indicates the gap between lexems.
			// several upper-case letters are treated as a single lexem.
			// upper-case letter followed by lowercase indicates the
			// new lexem.

			if (i+1 < len(nr) && unicode.IsLower(nr[i+1])) || (i-1 >= 0 && unicode.IsLower(nr[i-1])) {
				if len(cur) > 0 {
					// determine if previous symbol needs to be modified.
					out, cur = l.lookBehind(out, cur)
					// append to the outgoing sequence.
					out, cur = append(out, l.lexBetween(cur)...), nil
				}
			}
		}

		cur = append(cur, nr[i])
	}

	if len(cur) > 0 {
		// determine if previous symbol needs to be modified.
		out, cur = l.lookBehind(out, cur)
		// append to the outgoing sequence.
		out = append(out, l.lexBetween(cur)...)
	}

	for i := range out {
		out[i] = out[i] // strings.ToLower(out[i])
	}

	return out
}

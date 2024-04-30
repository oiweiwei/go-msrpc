package midl

// lexer.go implements the goyacc lexer interface.

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

var (
	// ErrEOF ...
	ErrEOF = errors.New("unexpected EOF")
)

// Lexer ...
type Lexer struct {
	errLex    error
	pos, line int
	sqb       int
	rb        int
	cb        int
	s         string
	uuidNext  bool
}

// NewLexer ...
func NewLexer(s string) *Lexer {
	return &Lexer{s: s}
}

const (
	errStr = "syntax error: %s: line: %d: near '%s'"
)

// Error ...
func (l *Lexer) Error(s string) {
	l.errLex = fmt.Errorf(errStr, s, l.line+1, strings.Split(l.s, "\n")[l.line])
}

// Errorf ...
func (l *Lexer) Errorf(frmt string, args ...interface{}) {
	l.errLex = fmt.Errorf(frmt, args...)
}

// Push function pushes the string on current position to re-enter the line.
func (l *Lexer) Push(s string) {
	l.s = l.s[:l.pos] + s + "\n" + l.s[l.pos:]
}

func (l *Lexer) skipSpace() (c rune) {

	// skip spaces.
	for c = l.next(); unicode.IsSpace(c); c = l.next() {
		if c == '\n' {
			l.line++
		}
	}

	return c
}

// Lex ...
func (l *Lexer) Lex(lval *RPCSymType) int {

	if l.errLex != nil {
		return 0
	}

	for {

		var (
			c, c1 rune
			ret   int
		)

		switch c = l.skipSpace(); c {
		case '[':
			l.sqb++
			return int(c)
		case ']':
			l.sqb--
			return int(c)
		case '(':
			l.rb++
			return int(c)
		case ')':
			l.rb--
			return int(c)
		case '{':
			l.cb++
			return int(c)
		case '}':
			l.cb--
			return int(c)
		case '*', '%', '+', '-', ':', ';':
			return int(c)
		case 0:
			return 0
		}

		// next next character. ignore eof here.
		switch c1 = l.next(); c {
		case '/':
			switch c1 {
			case '/':
				// read comment line.
				if !l.skipLine() {
					return 0
				}
				// redo lexer.
				continue
			case '*':
				if !l.skipCommentBlock() {
					return 0
				}
				// redo lexer.
				continue
			}
		case '&':
			switch c1 {
			case '&':
				ret = LOGICAL_AND
			}
		case '|':
			switch c1 {
			case '|':
				ret = LOGICAL_OR
			}
		case '<':
			switch c1 {
			case '<':
				ret = LSH
			case '=':
				ret = LE
			}
		case '>':
			switch c1 {
			case '>':
				ret = RSH
			case '=':
				ret = GE
			}
		case '=':
			switch c1 {
			case '=':
				ret = EQ
			}
		case '!':
			switch c1 {
			case '=':
				ret = NE
			}
		case '.':
			switch c1 {
			case '.':
				ret = RNG
			}
		case 'L':
			switch c1 {
			case '"':
				// shift 'L' sym. and one more to go the
				// '"' position.
				c, c1 = c1, l.next()
			}
		case '?':
			switch c1 {
			case '?':
				switch c1 = l.next(); c1 {
				case '<':
					// ??<.
					ret = int('{')
				case '>':
					// ??>.
					ret = int('}')
				default:
					// shift 3-rd symbol back.
					l.backup()
					c1 = '?'
				}
			}
		}

		// we have a 2-sym character.
		if ret != 0 {
			return ret
		}

		// we do not use the next character, backup it back
		// for further consideration.
		if ret == 0 && c1 != 0 {
			l.backup()
		}

		// check one digit symbol.
		switch c {
		case '&':
			return AND
		case '|':
			return OR
		case '^':
			return XOR
		case '>':
			return GT
		case '<':
			return LT
		case '=', '.', '!', '?':
			return int(c)
		}

		// we do not use the first character, backup it back
		// so that other types will read it again.
		l.backup()

		switch {
		case c == '"':
			return l.lexString(lval)
		case c == '\'':
			return l.lexChar(lval)
		case unicode.IsDigit(c):
			return l.lexInt(lval)
		case c == '#':
			if c = l.next(); c == 0 {
				l.Errorf("lex: lexPragma: %v", ErrEOF)
				return 0
			}
			if c = rune(l.lexIdent(lval)); c == 0 {
				l.Errorf("lex: lexPragma: %v", ErrEOF)
				return 0
			}
			switch lval.Ident {
			case "define":
				return PRAGMA_DEFINE
			default:
				// skip pragma.
				if !l.skipLine() {
					return 0
				}
				continue
			}

		case isIdent(c):
			return l.lexIdent(lval)
		}

		// read character back and return as an integer.
		return int(l.next())
	}
}

// next function advances position on one character and
// returns the character.
func (l *Lexer) next() (c rune) {
	if l.pos == len(l.s) {
		return 0
	}

	c, l.pos = rune(l.s[l.pos]), l.pos+1
	return c
}

// backup function steps for one position back.
func (l *Lexer) backup() {
	l.pos--
}

// TokName function maps the token to the token name.
func (l *Lexer) TokName(c int) (string, bool) {

	if l.sqb > 0 && l.rb == 0 {
		if _, ok := SQBReservedTok[c]; ok {
			return "", false
		}
	}
	if _, ok := ReservedTok[c]; ok {
		return "", false
	}
	ret, ok := TokName[c]
	return ret, ok
}

// lexInt function parses integer literal (eg 123, 0xfffe).
func (l *Lexer) lexInt(lval *RPCSymType) int {

	var (
		ret  *big.Int = big.NewInt(0)
		c    rune
		text []rune
	)

	for c = l.next(); isAlpha(c); c = l.next() {
		text = append(text, c)
	}

	if !isAlpha(c) && c != 0 {
		l.backup()
	}

	if err := ret.UnmarshalText([]byte(string(text))); err != nil {
		l.Errorf("lex: lexInt: %v", err)
		return 0
	}

	lval.Int = ret
	return INT_LITERAL
}

// lexString function parses the string literal (eg "some string").
func (l *Lexer) lexString(lval *RPCSymType) int {

	var (
		ret string
		c   rune
	)

	if c = l.next(); c == 0 {
		l.Errorf("lex: lexString: %v", ErrEOF)
		return 0
	} else if c != '"' {
		l.Errorf("lex: lexString: invalid string literal: missing '\"' at the beginning of the sequence")
		return 0
	}

	for c = l.next(); c != '"'; c = l.next() {

		switch c {
		case '\\':
			// process escape character.
			if c = l.next(); c == 0 {
				l.Errorf("lex: lexString: %v", ErrEOF)
				return 0
			} else if c == '\'' {
				l.Errorf("lex: lexString: invalid string escape \\%c", c)
				return 0
			}
			var ok bool
			if c, ok = escape(c); !ok {
				l.Errorf("lex: lexString:  invalid string escape \\%c", c)
				return 0
			}
		case 0:
			// EOF.
			l.Errorf("lex: lexString: unexpected EOF")
			return 0
		}

		ret += string(c)
	}

	lval.String = ret
	return STRING_LITERAL
}

// lexChar function parses the character literal (eg 'a', '\n').
func (l *Lexer) lexChar(lval *RPCSymType) int {

	var (
		ret rune
		c   rune
	)

	if c = l.next(); c == 0 {
		l.Errorf("lex: lexChar: %v", ErrEOF)
		return 0
	} else if c != '\'' {
		l.Errorf("lex: lexChar: invalid char literal: missing ' at the beginning of the sequence")
		return 0
	}

	switch c = l.next(); c {
	case '\'':
		l.Errorf("lex: lexChar: empty char literal")
	case '\\':
		// process escape character.
		if c = l.next(); c == 0 {
			l.Errorf("lex: lexChar: %v", ErrEOF)
			return 0
		} else if c == '"' {
			l.Errorf("lex: lexChar: invalid char escape \\%c", c)
			return 0
		}
		var ok bool
		if c, ok = escape(c); !ok {
			l.Errorf("lex: lexChar: invalid char escape \\%c", c)
			return 0
		}
	case 0:
		l.Errorf("lex: lexChar: %v", ErrEOF)
		return 0
	}

	ret = c

	if c = l.next(); c == 0 {
		l.Errorf("lex: lexChar: %v", ErrEOF)
		return 0
	} else if c != '\'' {
		l.Errorf("lex: lexChar: invalid char literal: missing ' at the end of the sequence")
		return 0
	}

	lval.Char = ret
	return CHARACTER_LITERAL
}

// lexIdent function parses the identifier and returns either an
// IDENT or reserved token.
func (l *Lexer) lexIdent(lval *RPCSymType) int {

	var (
		c    rune
		text []rune
	)

	for c = l.next(); isAlpha(c); c = l.next() {
		text = append(text, c)
	}

	if !isAlpha(c) && c != 0 {
		l.backup()
	}

	lval.Ident = string(text)

	if ret, ok := NameTok[lval.Ident]; ok {

		// outside of square brackets or inside and within round brackets
		// this can be an ident. (but only within curly brackets definition).
		if l.cb > 0 {
			if l.sqb == 0 {
				if _, ok := ReservedTok[ret]; !ok {
					return IDENT
				}
			} else if l.rb > 0 {
				if _, ok := SQBReservedTok[ret]; !ok {
					return IDENT
				}
			}
		}

		if ret == UUID {
			return l.lexUUID(lval)
		}

		return ret
	}

	return IDENT
}

// lexUUID function parses the UUID string representation.
func (l *Lexer) lexUUID(lval *RPCSymType) int {

	var (
		text []rune
		c    rune
		err  error
	)

	if tok := l.Lex(lval); tok != '(' {
		l.Errorf("unexpected symbol %v for uuid attribute", tok)
		return 0
	}

	if c = l.skipSpace(); c == 0 {
		return 0
	}

	expQ := c == '"'
	if !expQ {
		l.backup()
	}

	for i := 0; i < 36; i++ {
		switch c = l.next(); c {
		case 0:
			l.Errorf("lex: lexUUID: %v", ErrEOF)
			return 0
		default:
			text = append(text, c)
		}
	}

	var u *uuid.UUID

	if u, err = uuid.Unmarshal(string(text)); err != nil {
		l.Errorf("lex: lexUUID: %v", err)
		return 0
	}

	if expQ {
		if c := l.next(); c != '"' {
			l.Errorf("unexpected symbol %v, expected quote", c)
			return 0
		}
	}

	if tok := l.Lex(lval); tok != ')' {
		l.Errorf("unexpected symbol %v for uuid attribute", tok)
		return 0
	}

	lval.UUID = u

	return UUID
}

// skipLine function skips the line till the \n character and
// returns 'true' if EOF hasn't been reached.
func (l *Lexer) skipLine() bool {
	for c := l.next(); c != '\n'; c = l.next() {
		if c == 0 {
			return false
		}
	}
	l.line++
	return true
}

// skipCommentBlock function skips over comment block (/* */) and
// returns 'true' if EOF hasn't been reached.
func (l *Lexer) skipCommentBlock() bool {
	for c := l.next(); ; c = l.next() {
		switch c {
		case '*':
			if c1 := l.next(); c1 == '/' {
				// comment end, goto start.
				return true
			} else if c1 == 0 {
				l.Errorf("lex: %v", ErrEOF)
				return false
			} else {
				l.backup()
			}
		case '\n':
			l.line++
		case 0:
			l.Errorf("lex: %v", ErrEOF)
			return false
		}
	}
}

// esc is an escape character mapping.
var esc = map[rune]rune{
	'a':  '\a',
	'b':  '\b',
	'f':  '\f',
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
	'v':  '\v',
	'\\': '\\',
	'\'': '\'',
	'"':  '"',
}

// escape function escapes the rune and returns 'true' if
// escape character exists.
func escape(c rune) (rune, bool) {
	ret, ok := esc[c]
	return ret, ok
}

// isIdent function returns 'true' when character is valid
// identifier start character.
func isIdent(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// isAlpha function returns 'true' when character is valid
// alphanumeric character.
func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		(c == '_')
}

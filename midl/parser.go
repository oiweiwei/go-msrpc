package midl

// parser.go contains the IDL parser implementation.

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {

	/*

		_, p, _, ok := runtime.Caller(0)
		if !ok {
			panic("cannot determine current dir")
		}

	*/

	// load global default.
	_, err := NewFile("dcetypes.idl", "").Load()
	if err != nil {
		panic("unable to load dcetypes.idl: " + err.Error())
	}
}

// Parse ...
func Parse(s string) (*File, error) {

	var (
		p = NewParser(s)
	)

	if RPCParse(p); p.errLex != nil {
		return nil, p.errLex
	}

	// save exportSyms to the file for further reference.
	p.result.exportSyms = p.exportSyms

	// resolve defered type rederences.
	for _, ref := range p.refs {
		typ, ok := p.LookupType(ref.Name)
		if !ok {
			name := strings.Replace(ref.Name, "_struct_", "_union_", 1)
			if typ, ok = p.LookupType(name); !ok {
				return nil, fmt.Errorf("unable to resolve type %s", ref.Name)
			}
		}

		*ref = *typ
	}

	return p.result, nil
}

// setResult ...
func setResult(p interface{}, f *File) {
	p.(*Parser).result = f
}

// exportSyms ...
func exportSyms(p interface{}, n string, e *Export) {
	pr := p.(*Parser)
	pr.curPos++
	e.Position = pr.curPos
	pr.exportSyms[n] = e
}

// tokName ...
func tokName(p interface{}, c int) (string, bool) {
	ret, ok := p.(*Parser).TokName(c)
	if ok {
		p.(*Parser).errLex = nil
	}
	return ret, ok
}

// pushLex ...
func pushLex(p interface{}, s string) {
	p.(*Parser).Push(s)
}

// pushRef ...
func pushRef(p interface{}, typ *Type) *Type {
	p.(*Parser).Defer(typ)
	return typ
}

// lookupType ...
func lookupType(p interface{}, n string) (*Type, bool) {
	typ, ok := p.(*Parser).LookupType(n)
	return typ, ok
}

func resolvePointers(p interface{}, pointerDefault PointerType) {
}

// lookupConst ...
func lookupConst(p interface{}, n string) (*Const, bool) {
	c, ok := p.(*Parser).LookupConst(n)
	return c, ok
}

// storeConst ...
func storeConst(p interface{}, n string, val Expr, typ ...Kind) {
	p.(*Parser).Store(n, val, typ...)
}

type Parser struct {
	// The lexer entity.
	*Lexer
	// The result.
	result *File
	// The export symbols for the resulting file.
	exportSyms map[string]*Export
	// The list of delayed references.
	refs []*Type
	// The current exported symbol position.
	curPos int
}

// NewParser function returns the parser entity.
func NewParser(s string) *Parser {
	return &Parser{
		Lexer:      NewLexer(s),
		exportSyms: make(map[string]*Export),
	}
}

// Defer function defers the type resolution until the parse
// end.
func (p *Parser) Defer(ref *Type) {
	p.refs = append(p.refs, ref)
}

// LookupType ...
func (p *Parser) LookupType(n string) (*Type, bool) {

	var (
		e  *Export
		ok bool
	)

	if e, ok = p.exportSyms[n]; !ok {
		for _, f := range Files() {
			if e, ok = f.exportSyms[n]; !ok {
				continue
			}
			break
		}
	}

	if !ok {
		return nil, false
	} else if e.Type == nil {
		log.Warnf("parse: %s: not a type reference", n)
		return nil, false
	}

	return e.Type, true
}

// LookupConst ...
func (p *Parser) LookupConst(n string) (*Const, bool) {

	var (
		e  *Export
		f  *File
		ok bool
	)

	if e, ok = p.exportSyms[n]; !ok {
		for _, f = range Files() {
			if e, ok = f.exportSyms[n]; !ok {
				continue
			}
			break
		}
	}
	if !ok {
		return nil, false
	} else if e.Const == nil {
		log.Warnf("not a const reference %s", n)
		return nil, false
	}

	return e.Const, true
}

// LookupExpr ...
func (p *Parser) LookupExpr(n string) (Expr, bool) {

	c, ok := p.LookupConst(n)
	if !ok {
		return Expr{}, false
	}

	return c.Value, true
}

// Store ...
func (p *Parser) Store(n string, val Expr, typ ...Kind) {
	e := &Export{
		Name: n,
		Const: &Const{
			Name:  n,
			Value: val,
		},
	}
	if len(typ) > 0 {
		e.Const.Type = typ[0]
	}
	p.exportSyms[n] = e
}

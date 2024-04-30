package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl"
)

type fileCtx struct{}

func WithFile(ctx context.Context, f *midl.File) context.Context {
	return context.WithValue(ctx, fileCtx{}, f)
}

func File(ctx context.Context) *midl.File {
	f, _ := ctx.Value(fileCtx{}).(*midl.File)
	return f
}

type reservedCtx struct{}

func WithReserved(ctx context.Context) context.Context {
	return context.WithValue(ctx, reservedCtx{}, true)
}

func IsReserved(ctx context.Context) bool {
	ok, _ := ctx.Value(reservedCtx{}).(bool)
	return ok
}

type interfaceCtx struct{}

func WithInterface(ctx context.Context, iff *midl.Interface) context.Context {
	return context.WithValue(ctx, interfaceCtx{}, iff)
}

func Interface(ctx context.Context) *midl.Interface {
	iff, _ := ctx.Value(interfaceCtx{}).(*midl.Interface)
	return iff
}

type fieldCtx struct{}

func WithField(ctx context.Context, field *midl.Field) context.Context {
	return context.WithValue(ctx, fieldCtx{}, field)
}

func Field(ctx context.Context) *midl.Field {
	field, _ := ctx.Value(fieldCtx{}).(*midl.Field)
	return field
}

var (
	Obj = "o"
)

func (p *Generator) GenRange(ctx context.Context, n string, i, upto int, fn func(context.Context, string)) {
	if i >= upto {
		fn(ctx, n)
		return
	}
	idx := fmt.Sprintf("i%d", i+1)
	p.Range(idx, n, func() {
		p.GenRange(ctx, p.IVar(n, idx), i+1, upto, fn)
	})
}

func (p *Generator) If(args ...interface{}) {
	p.Block(append([]interface{}{"if"}, args...)...)
}

var (
	CheckErr1 = func(p *Generator) func(...interface{}) {
		return func(args ...interface{}) {
			p.Block(append(append([]interface{}{"if", "err", ":="}, args...), ";", "err != nil", func() {
				p.P("return", "err")
			}))
		}
	}

	CheckErr2 = func(p *Generator) func(...interface{}) {
		return func(args ...interface{}) {
			p.P(args...)
		}
	}
)

func (p *Generator) Structure(args ...interface{}) {
	p.Block(append([]interface{}{"type", args[0], "struct"}, args[1:]...)...)
}

func (p *Generator) Range(args ...interface{}) {

	switch len(args) {
	case 0, 1:
		panic("unknown range")
	case 2:
		p.Block(append([]interface{}{"for", "range"}, args...)...)
	case 3:
		p.Block(append([]interface{}{"for", args[0], ":=", "range"}, args[1:]...)...)
	default:
		p.Block(append([]interface{}{"for", args[0], ",", args[1], ":=", "range"}, args[2:]...)...)
	}
}

func (p *Generator) ElseIf(args ...interface{}) []interface{} {
	return append([]interface{}{"else", "if"}, args...)
}

func (p *Generator) Else(args ...interface{}) []interface{} {
	return append([]interface{}{"else"}, args...)
}

func (p *Generator) Block(args ...interface{}) {

	var (
		printable []interface{}
	)

	for _, arg := range args {
		switch arg := arg.(type) {
		case func():
			if len(printable) > 0 {
				p.P(append(printable, "{")...)
			}
			printable = []interface{}{"}"}
			arg()
		case []interface{}:
			p.Block(append(printable, arg...)...)
			printable = nil
		default:
			printable = append(printable, arg)
		}
	}

	if len(printable) > 0 {
		p.P(printable...)
	}

}

// Var function generates the variable name.
func (p *Generator) Var(n string, idx ...interface{}) string {
	if len(idx) > 0 {
		return fmt.Sprintf("%s%v", n, idx[0])
	}
	return n
}

func (p *Generator) Q(n string) string {
	return `"` + n + `"`
}

func (p *Generator) B(pre interface{}, s ...interface{}) string {

	ss := make([]string, len(s))
	for i := range s {
		ss[i] = fmt.Sprintf("%v", s[i])
	}

	return fmt.Sprintf("%v(%s)", pre, strings.Join(ss, ","))
}

func (p *Generator) Amp(v interface{}) string {
	return "&" + fmt.Sprintf("%v", v)
}

func (p *Generator) BPtr(ptr interface{}) string {
	return fmt.Sprintf("(*%v)", ptr)
}

var (
	varReplacer = strings.NewReplacer("[", "_", "]", "", Obj+".", "")
)

func (p *Generator) ToVar(s string) string {
	return varReplacer.Replace(s)
}

func (p *Generator) IVar(n string, idxs ...interface{}) string {
	for _, idx := range idxs {
		n += fmt.Sprintf("[%v]", idx)
	}
	return n
}

// EO function generates the O with namer.
func (p *Generator) EO(n string) string {
	return p.O(GoName(n))
}

// O function generates the "o."
func (p *Generator) O(n string) string {
	return Obj + "." + n
}

var (
	ZeroLen = "%ZERO%"
)

func (p *Generator) Len(n string) string {
	if strings.HasPrefix(n, ZeroLen) {
		return "0"
	}
	return "len(" + n + ")"
}

// R function generates the repeat with start end sequence.
func (p *Generator) R(start, s, end string, count int) string {
	return strings.Repeat(start, count) + s + strings.Repeat(end, count)
}

func (p *Generator) BufVar(name string) string {
	return "_" + p.ToVar(name) + "_buf"
}

type varNameCtx struct{}

func WithVarName(ctx context.Context, n string) context.Context {
	return context.WithValue(ctx, varNameCtx{}, n)
}

func (p *Generator) VarName(ctx context.Context, index ...interface{}) string {
	n, _ := ctx.Value(varNameCtx{}).(string)
	if n != "" {
		return p.IVar(n, index...)
	}
	return ""
}

type isOp struct{}

func WithOp(ctx context.Context) context.Context {
	return context.WithValue(ctx, isOp{}, true)
}

func IsOp(ctx context.Context) bool {
	ok, _ := ctx.Value(isOp{}).(bool)
	return ok
}

package midl

import (
	"fmt"
	"math/big"
)

type ExprStore interface {
	LookupExpr(string) (Expr, bool)
}

type Args struct {
	args []Expr
}

func NewArgs(args ...interface{}) ExprStore {
	exprStore := &Args{}
	for _, arg := range args {
		exprStore.args = append(exprStore.args, NewValue(arg))
	}
	return exprStore
}

func (a *Args) LookupExpr(string) (Expr, bool) {
	var arg Expr
	if len(a.args) > 0 {
		arg, a.args = a.args[0], a.args[1:]
		return arg, true
	}
	return arg, false
}

func IsBoolOp(op int) bool {
	switch op {
	case LOGICAL_OR, LOGICAL_AND, GE, GT, LE, LT, EQ, NE:
		return true
	}
	return false
}

var binaryToFunc = map[int]func(Expr, Expr) (Expr, bool){
	'+':         Expr.Add,
	'-':         Expr.Sub,
	'*':         Expr.Mul,
	'/':         Expr.Div,
	'%':         Expr.Rem,
	AND:         Expr.And,
	OR:          Expr.Or,
	XOR:         Expr.Xor,
	LSH:         Expr.Lsh,
	RSH:         Expr.Rsh,
	GE:          Expr.Ge,
	GT:          Expr.Gt,
	LE:          Expr.Le,
	LT:          Expr.Lt,
	EQ:          Expr.Eq,
	NE:          Expr.Ne,
	LOGICAL_OR:  Expr.LogicalOr,
	LOGICAL_AND: Expr.LogicalAnd,
}

var inverse = map[int]int{
	UPLUS:  UMINUS,
	UMINUS: UPLUS,
	'+':    '-',
	'-':    '+',
	'*':    '/',
	'/':    '*',
}

var unaryToFunc = map[int]func(Expr) (Expr, bool){
	UPLUS:  Expr.Positive,
	UMINUS: Expr.Negative,
	UNEG:   Expr.Neg,
	UNOT:   Expr.Not,
	UMUL:   Expr.Ptr,
}

var ternaryToFunc = map[int]func(Expr, Expr, Expr) (Expr, bool){
	TERNARY: Expr.Ter,
}

// Exprs is a list of expressions.
type Exprs []Expr

// CanEval function returns true if expression can be evaluated
// immediately.
func CanEval(es ...Expr) bool {
	for _, e := range es {
		if !e.CanEval() {
			return false
		}
	}
	return true
}

// NewExpr function builds a new expression with expression tree.
func NewExpr(op int, valuer func() (interface{}, bool), es ...Expr) (Expr, bool) {

	var (
		ret Expr
		ok  bool
	)

	switch ret.reqParam = !CanEval(es...); len(es) {
	case 1:
		ret.Expr = &ExprTree{Op: op, Lval: es[0].Expr}
	case 2:
		ret.Expr = &ExprTree{Op: op, Lval: es[0].Expr, Rval: es[1].Expr}
	case 3:
		ret.Expr = &ExprTree{Op: op, Cond: es[0].Expr, Lval: es[1].Expr, Rval: es[2].Expr}
	default:
		panic("unknown expression")
	}

	// we cannot evaluate the value at this point, so return.
	if ret.reqParam {
		return ret, true
	}

	// extract value.
	if ret.Value, ok = valuer(); !ok {
		return ret, false
	}

	return ret, true
}

// bigInt2 is a helper method to extract two *big.Int values.
func bigInt2(lval, rval Expr) (*big.Int, *big.Int, bool) {
	lv, ok := lval.BigInt()
	if !ok {
		return nil, nil, false
	}
	rv, ok := rval.BigInt()
	if !ok {
		return nil, nil, false
	}
	return lv, rv, true
}

// ExprTree is an expression tree build when expression cannot
// be evaluated immediately.
type ExprTree struct {
	// Op is an expression tree operation.
	Op int
	// Value is a terminal value of the expression tree.
	Value interface{}
	// Cond, Lval, Rval are the expression tree children.
	// (Cond used for ternary only, Lval is a value for unary operation).
	Cond, Lval, Rval *ExprTree
}

func (e *ExprTree) ResolveTo(to Expr) (Expr, bool) {

	if e == nil {
		return Expr{}, false
	}

	if e.Op == IDENT {
		if to.IsZero() {
			return Expr{Value: e.Value, reqParam: true, Expr: e}, true
		}
		return Expr{Value: e.Value, reqParam: true, Expr: &ExprTree{Op: IDENT, Value: to.Value}}, true
	}

	if expr, ok := e.Cond.ResolveTo(to); ok {
		return expr, true
	}

	if expr, ok := e.Lval.ResolveTo(to); ok {
		return expr, true
	}

	if expr, ok := e.Rval.ResolveTo(to); ok {
		return expr, true
	}

	return Expr{}, false
}

func (e *ExprTree) Resolve(expr *ExprTree) (Expr, *ExprTree, bool) {

	if e == nil || (e.Op != IDENT && inverse[e.Op] == 0) {
		return Expr{}, nil, false
	}

	if e.Op == UMINUS {
		e = &ExprTree{
			Op: '/',
			Lval: &ExprTree{
				Value: big.NewInt(-1),
			},
			Rval: e.Lval,
		}
	} else if e.Op == UPLUS {
		e = &ExprTree{
			Lval: e.Lval,
		}
	}

	if e.Op == IDENT {
		return Expr{Value: e.Value}, expr, true
	}

	if ident, expr, ok := e.Lval.Resolve(&ExprTree{
		Rval: e.Rval,
		Lval: expr,
		Op:   inverse[e.Op],
	}); ok {
		return ident, expr, true
	}

	if ident, expr, ok := e.Rval.Resolve(&ExprTree{
		Lval: expr,
		Rval: e.Lval,
		Op:   inverse[e.Op],
	}); ok {
		return ident, expr, true
	}

	return Expr{}, nil, false
}

// Eval function evaluates the expression tree with Store for lookup.
func (e *ExprTree) Eval(l ExprStore) (Expr, bool) {

	switch e.Op {
	case 0:
		return Expr{Value: e.Value}, true
	case IDENT:
		if l == nil {
			return Expr{}, false
		}
		// lookup value.
		exp, ok := l.LookupExpr(e.Value.(string))
		if !ok {
			return Expr{}, false
		}
		return exp, true
	case '+', '-', '*', '/', '%', AND, OR, XOR, LSH, RSH, GE, GT, LT, LE, EQ, NE, LOGICAL_AND, LOGICAL_OR:
		lv, ok := e.Lval.Eval(l)
		if !ok {
			return Expr{}, false
		}
		rv, ok := e.Rval.Eval(l)
		if !ok {
			return Expr{}, false
		}
		return binaryToFunc[e.Op](lv, rv)
	case UPLUS, UNEG, UMINUS, UNOT:
		v, ok := e.Lval.Eval(l)
		if !ok {
			return Expr{}, false
		}
		return unaryToFunc[e.Op](v)
	case TERNARY:
		cond, ok := e.Cond.Eval(l)
		if !ok {
			return Expr{}, false
		}
		lv, ok := e.Lval.Eval(l)
		if !ok {
			return Expr{}, false
		}
		rv, ok := e.Rval.Eval(l)
		if !ok {
			return Expr{}, false
		}
		return ternaryToFunc[TERNARY](cond, lv, rv)
	}

	return Expr{}, false
}

// Expr is an expression.
type Expr struct {
	// Value is an expression value.
	Value interface{}
	// CanEval flag indicates whether the expression can be evluated.
	reqParam bool
	// Expr is a parsed expression tree.
	Expr *ExprTree
}

func (e Expr) Resolve(in Expr) (Expr, bool) {

	ident, expr, ok := e.Expr.Resolve(in.Expr)
	if !ok {
		return Expr{}, false
	}

	_, reqParam := e.Eval(nil)

	ret := Expr{Value: ident.Value, reqParam: !reqParam, Expr: expr}

	return ret, true
}

func (e Expr) IsZero() bool {
	return e.Value == nil && e.Expr == nil
}

func (e Expr) IsIdent() bool {
	return !e.IsZero() && (e.Expr.Op == IDENT || e.Expr.Op == UMUL && e.Expr.Lval.Op == IDENT)
}

// Eval function evaluates the expression and returns evaluated expression.
func (e Expr) Eval(l ExprStore) (Expr, bool) {
	if !e.reqParam {
		return e, true
	}
	exp, ok := e.Expr.Eval(l)
	if !ok {
		return Expr{}, false
	}
	return exp, true
}

func (e Expr) CanEval() bool {
	return !e.reqParam
}

// BigInt function returns the *big.Int representation of the value.
func (e Expr) BigInt() (*big.Int, bool) {
	val, _ := e.Value.(*big.Int)
	return val, val != nil
}

// Bool function returns boolean representation of the expression.
func (e Expr) Bool() bool {
	var ret bool
	switch v := e.Value.(type) {
	case bool:
		ret = v
	case string:
		ret = v != ""
	case rune:
		ret = v != 0
	case *big.Int:
		ret = v.Uint64() != 0
	default:
		ret = e.Value != nil
	}
	return ret
}

// String function returns string representation of the expression.
func (e Expr) Str() (string, bool) {
	ret, ok := e.Value.(string)
	return ret, ok
}

// Char function returns the char representation of the expression.
func (e Expr) Char() (rune, bool) {
	ret, ok := e.Value.(rune)
	return ret, ok
}

// Null function returns the null representation of the expression.
func (e Expr) Null() bool {
	return e.Value == nil
}

// Int function returns int64 representation of the expression.
func (e Expr) Int64() (int64, bool) {
	val, ok := e.BigInt()
	if !ok || !val.IsInt64() {
		return 0, false
	}
	return val.Int64(), true
}

// Uint function returns uint64 representation of the expression.
func (e Expr) Uint64() (uint64, bool) {
	val, ok := e.BigInt()
	if !ok || !val.IsUint64() {
		return 0, false
	}
	return val.Uint64(), true
}

// LogicalOr function performs logical-or (||) operation on expression.
func (lval Expr) LogicalOr(rval Expr) (Expr, bool) {
	return NewExpr(LOGICAL_OR, func() (interface{}, bool) {
		return lval.Bool() || rval.Bool(), true
	}, lval, rval)
}

// LogicalAnd function performs logical-and (&&) operation on expression.
func (lval Expr) LogicalAnd(rval Expr) (Expr, bool) {
	return NewExpr(LOGICAL_AND, func() (interface{}, bool) {
		return lval.Bool() && rval.Bool(), true
	}, lval, rval)
}

// Ternary function performs ternary (?:) operation on expression.
func (cond Expr) Ter(lval, rval Expr) (Expr, bool) {
	return NewExpr(TERNARY, func() (interface{}, bool) {
		if cond.Bool() {
			return lval.Value, true
		}
		return rval.Value, true
	}, cond, lval, rval)
}

// Or function performs the bitwise or operation on expression.
func (lval Expr) Or(rval Expr) (Expr, bool) {
	return NewExpr(OR, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return big.NewInt(0).Or(lv, rv), true
	}, lval, rval)
}

// Or function performs the bitwise and operation on expression.
func (lval Expr) And(rval Expr) (Expr, bool) {
	return NewExpr(AND, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return big.NewInt(0).And(lv, rv), true
	}, lval, rval)
}

// Or function performs the bitwise exclusive-or (xor) operation on expression.
func (lval Expr) Xor(rval Expr) (Expr, bool) {
	return NewExpr(XOR, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return big.NewInt(0).Xor(lv, rv), true
	}, lval, rval)
}

// Le function compares two integers and returns `true` if lval <= rval.
func (lval Expr) Le(rval Expr) (Expr, bool) {
	return NewExpr(LE, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return lv.Cmp(rv) <= 0, true
	}, lval, rval)
}

// Ge function compares two integers and returns `true` if lval >= rval.
func (lval Expr) Ge(rval Expr) (Expr, bool) {
	return NewExpr(GE, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return lv.Cmp(rv) >= 0, true
	}, lval, rval)
}

// Lt function compares two integers and returns `true` if lval < rval.
func (lval Expr) Lt(rval Expr) (Expr, bool) {
	return NewExpr(LT, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return lv.Cmp(rv) < 0, true
	}, lval, rval)
}

// Gt function compares two integers and returns `true` if lval > rval.
func (lval Expr) Gt(rval Expr) (Expr, bool) {
	return NewExpr(GT, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return lv.Cmp(rv) > 0, true
	}, lval, rval)
}

// Eq function compares two values and returns `true` if lval is equal rval.
func (lval Expr) Eq(rval Expr) (Expr, bool) {
	return NewExpr(EQ, func() (interface{}, bool) {
		switch lval.Value.(type) {
		case *big.Int:
			lv, rv, ok := bigInt2(lval, rval)
			if !ok {
				return nil, false
			}
			return lv.Cmp(rv) == 0, true
		case bool:
			return lval.Bool() == rval.Bool(), true
		case string:
			lv, ok := lval.Str()
			if !ok {
				return nil, false
			}
			rv, ok := rval.Str()
			if !ok {
				return nil, false
			}
			return lv == rv, true
		case rune:
			lv, ok := lval.Char()
			if !ok {
				return nil, false
			}
			rv, ok := rval.Char()
			if !ok {
				return nil, false
			}
			return lv == rv, true
		default:
			if ok := lval.Null() && rval.Null(); !ok {
				return nil, false
			}
			return true, true
		}
	}, lval, rval)
}

// Ne function compares two values and returns `true` if `lval` is not equal `rval`.
func (lval Expr) Ne(rval Expr) (Expr, bool) {
	ret, ok := lval.Eq(rval)
	if !ok {
		return ret, false
	}
	if ret.Expr.Op = NE; ret.CanEval() {
		ret.Value = !ret.Bool()
	}
	return ret, true
}

// Lsh function performs the left-shift of the `lval` on `rval` amount.
// `rval` must be a valid unsigned integer.
func (lval Expr) Lsh(rval Expr) (Expr, bool) {
	return NewExpr(LSH, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok || !rv.IsUint64() {
			return nil, false
		}
		return big.NewInt(0).Lsh(lv, uint(rv.Uint64())), true
	}, lval, rval)
}

// Rsh function performs the right-shift of the `lval` on `rval` amount.
// `rval` must be a valid unsigned integer.
func (lval Expr) Rsh(rval Expr) (Expr, bool) {
	return NewExpr(RSH, func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok || !rv.IsUint64() {
			return nil, false
		}
		return big.NewInt(0).Rsh(lv, uint(rv.Uint64())), true
	}, lval, rval)
}

// Add function performs addition of two integers.
func (lval Expr) Add(rval Expr) (Expr, bool) {
	return NewExpr('+', func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return big.NewInt(0).Add(lv, rv), true
	}, lval, rval)
}

// Sub function performs integer substraction (`lval` - `rval`).
func (lval Expr) Sub(rval Expr) (Expr, bool) {
	return NewExpr('-', func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return big.NewInt(0).Sub(lv, rv), true
	}, lval, rval)
}

// Mul function performs integer multiplication (`lval` * `rval`).
func (lval Expr) Mul(rval Expr) (Expr, bool) {
	return NewExpr('*', func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok {
			return nil, false
		}
		return big.NewInt(0).Mul(lv, rv), true
	}, lval, rval)
}

// Div function performs integer division (`lval` / `rval`).
// Note that `rval` must evaluate to the non-zero value.
func (lval Expr) Div(rval Expr) (Expr, bool) {
	return NewExpr('/', func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok || rv.Uint64() == 0 {
			return nil, false
		}
		return big.NewInt(0).Div(lv, rv), true
	}, lval, rval)
}

// Rem function performs C-like integer reminder operation (`lval` % `rval`).
// Note that `rval` must evaluate to the non-zero value.
func (lval Expr) Rem(rval Expr) (Expr, bool) {
	return NewExpr('%', func() (interface{}, bool) {
		lv, rv, ok := bigInt2(lval, rval)
		if !ok || rv.Uint64() == 0 {
			return nil, false
		}
		return big.NewInt(0).Rem(lv, rv), true
	}, lval, rval)
}

// Positive function returns integer.
func (val Expr) Positive() (Expr, bool) {
	return NewExpr(UPLUS, func() (interface{}, bool) {
		val, ok := val.BigInt()
		if !ok {
			return nil, false
		}
		return val, true
	}, val)
}

// Negative function returns -integer.
func (val Expr) Negative() (Expr, bool) {
	return NewExpr(UPLUS, func() (interface{}, bool) {
		val, ok := val.BigInt()
		if !ok {
			return nil, false
		}
		return big.NewInt(0).Neg(val), true
	}, val)
}

// Ptr function doesn't do anything special, needed to
// resemble correct tree with pointer value.
func (val Expr) Ptr() (Expr, bool) {
	return NewExpr(UMUL, func() (interface{}, bool) {
		return val.Value, true
	}, val)
}

// Neg function performs integer bitwise negation.
func (val Expr) Neg() (Expr, bool) {
	return NewExpr(UNEG, func() (interface{}, bool) {
		val, ok := val.BigInt()
		if !ok {
			return nil, false
		}
		return big.NewInt(0).Not(val), true
	}, val)
}

// Not function performs boolean negation.
func (val Expr) Not() (Expr, bool) {
	return NewExpr(UNEG, func() (interface{}, bool) {
		return !val.Bool(), true
	}, val)
}

// NewValue function creates a new expression node value.
func NewValue(v interface{}) Expr {
	return Expr{Value: v, Expr: &ExprTree{Value: v}}
}

// NewIdent function creates a new expression variable.
func NewIdent(v string) Expr {
	return Expr{reqParam: true, Value: v, Expr: &ExprTree{Op: IDENT, Value: v}}
}

// Coerce function tries to cast the given Expr to the kind
// provided and returns the same expression.
func (e Expr) Coerce(kind Kind) (Expr, error) {

	var (
		ret Expr
		val interface{}
	)

	if !e.CanEval() {
		return ret, fmt.Errorf("cannot evaluate a value")
	}

	switch kind {
	case TypeInt8, TypeInt16, TypeInt32, TypeInt64:
		v, ok := e.BigInt()
		if !ok {
			return ret, fmt.Errorf("cannot coerce the expression to the integer type")
		}

		switch kind {
		case TypeInt8:
			if v.BitLen() > 8 {
				return ret, fmt.Errorf("8-bit integer overflow")
			}
		case TypeInt16:
			if v.BitLen() > 16 {
				return ret, fmt.Errorf("16-bit integer overflow")
			}
		case TypeInt32, TypeFloat32:
			if v.BitLen() > 32 {
				return ret, fmt.Errorf("32-bit integer overflow")
			}
		case TypeInt64, TypeFloat64:
			if v.BitLen() > 64 {
				return ret, fmt.Errorf("64-bit integer overflow")
			}
		default:
			return ret, fmt.Errorf("unknown size integer type")
		}
		val = big.NewInt(0).SetInt64(v.Int64())
	case TypeUint8, TypeUint16, TypeUint32, TypeUint64:
		v, ok := e.BigInt()
		if !ok {
			return ret, fmt.Errorf("cannot coerce the expression to the integer type")
		}
		switch kind {
		case TypeUint8:
			if v.BitLen() > 8 {
				return ret, fmt.Errorf("8-bit integer overflow")
			}
		case TypeUint16:
			if v.BitLen() > 16 {
				return ret, fmt.Errorf("16-bit integer overflow")
			}
		case TypeUint32:
			if v.BitLen() > 32 {
				return ret, fmt.Errorf("32-bit integer overflow")
			}
		case TypeUint64:
			if v.BitLen() > 64 {
				return ret, fmt.Errorf("64-bit integer overflow")
			}
		default:
			return ret, fmt.Errorf("unknown size integer type")
		}
		val = big.NewInt(0).SetUint64(v.Uint64())
	case TypeFloat32, TypeFloat64:
		v, ok := e.BigInt()
		if !ok {
			return ret, fmt.Errorf("cannot coerce the expression to the floating point type")
		}
		val = big.NewFloat(0).SetInt(v)
	case TypeString:
		v, ok := e.Str()
		if !ok {
			return ret, fmt.Errorf("cannot coerce the expression to the string type")
		}
		val = v
	case TypeBoolean:
		val = e.Bool()
	case TypeChar:
		v, ok := e.Char()
		if !ok {
			return ret, fmt.Errorf("cannot coerce the expression to the char type")
		}
		val = v
	case TypeVoid:
		if ok := e.Null(); !ok {
			return ret, fmt.Errorf("cannot coerce the expression to the void type")
		}
		val = nil
	default:
		return ret, fmt.Errorf("unsupported type %s", kind)
	}

	return Expr{Value: val, Expr: e.Expr}, nil
}

// Empty function returns 'true' when expression is empty.
func (val Expr) Empty() bool {
	return val.Expr == nil
}

func (val Expr) Ident() Expr {
	expr, _ := val.ResolveTo(Expr{})
	return expr
}

func (val Expr) ResolveTo(to Expr) (Expr, bool) {

	if val.CanEval() {
		return Expr{}, false
	}

	return val.Expr.ResolveTo(to)
}

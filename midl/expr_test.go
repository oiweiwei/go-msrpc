package midl

import (
	"fmt"
	"math/big"
	"testing"
)

type tExprStore map[string]Expr

func (s tExprStore) LookupExpr(v string) (Expr, bool) {
	e, ok := s[v]
	return e, ok
}

func TestExprTreeResolve(t *testing.T) {

	e := Expr{
		reqParam: true,
		Expr: &ExprTree{
			Op: '*',
			Rval: &ExprTree{
				Value: big.NewInt(4),
			},
			Lval: &ExprTree{
				Op: '+',
				Lval: &ExprTree{
					Op: UMINUS,
					Lval: &ExprTree{
						Op:    IDENT,
						Value: "x",
					},
				},
				Rval: &ExprTree{
					Value: big.NewInt(2),
				},
			},
		}}

	fmt.Println(e.String())

	eval, ok := e.Eval(tExprStore{"x": Expr{Value: big.NewInt(5)}})
	fmt.Println(eval.String(), ok)

	fmt.Println("y", "=", e.Expression())

	// r := &ExprTree{Value: big.NewInt(28)}
	rx, ok := e.Resolve(NewIdent("y"))
	fmt.Println(rx.Value, "=", rx.Expression(), ok)

}

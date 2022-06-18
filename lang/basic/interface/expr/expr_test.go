package expr

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr     Expr
		env      Env
		expected string
	}{
		// sqrt(A/pi)
		{
			call{
				fn: "sqrt",
				args: []Expr{
					binary{
						op: '/',
						x:  Var("A"),
						y:  Var("pi"),
					},
				},
			},
			Env{"A": 87616, "pi": math.Pi}, "167"},
	}

	for _, test := range tests {
		// 仅在表达式变更时才输出
		got := fmt.Sprintf("%.6g", test.expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.expected {
			t.Errorf("%s.Eval() in %v = %q, expected %q\n",
				test.expr, test.env, got, test.expected)
		}
	}
}

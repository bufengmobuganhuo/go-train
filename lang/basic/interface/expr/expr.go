package expr

import (
	"fmt"
	"math"
	"strings"
)

// 表达式
type Expr interface {
	// 返回表达式在Env下的计算结果
	Eval(env Env) float64
	// 检查表达式是否符合要求
	Check(vars map[Var]bool) error
}

// Var表示一个变量，比如x
type Var string

// 表示数字常量
type literal float64

// 表示一元操作符表达式，比如-x
type unary struct {
	// +或者-
	op rune
	// 表达式
	x Expr
}

// 二元操作符表达式，比如x+y
type binary struct {
	// 操作符，加减乘除
	op   rune
	x, y Expr
}

// 表示函数调用表达式，比如sin(x)
type call struct {
	// 要调用的函数，有sin, sqrt, pow
	fn string
	// 函数调用的参数
	args []Expr
}

// 一个上下文，用于将变量映射到对应的数值
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator:%q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator:%q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

// 函数名和需要的参数个数
var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func main() {
}

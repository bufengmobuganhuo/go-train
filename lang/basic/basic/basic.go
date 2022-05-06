package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 不可使用 ":="在函数外定义变量，包内可见的变量
var (
	aa = 3
	ss = "kkk"
)

// 变量的零值
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 自动推断类型
func variableTypeDeduction() {
	var a, b, s, c = 3, 4, "abc", true
	fmt.Println(a, b, s, c)
}

// 省略关键词var， ":=" 可以声明并赋值，后续使用时使用"="赋值，而不能再使用":="，否则就是重复定义
// 推荐方法
func variableShorter() {
	a, b, c, s := 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func euler() {
	// 欧拉公式
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
	// e的指数的特殊表示
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	// 类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 定义常量
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 定义枚举
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
	)
	fmt.Println(cpp, java, python)
	// 从0开始，后面递增
	const (
		ruby = iota
		c
		golang
	)
	fmt.Println(ruby, c, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss)

	euler()
	triangle()
	consts()
	enums()
}

package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 如果有错误，返回一个error
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// 只接收一个结果
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	// 会自动返回上述两个变量，如果返回的个数较多，则不建议
	return
}

// 传入一个入参(int, int)返回为int的函数
func apply(op func(int, int) int, a, b int) int {
	// 获取函数的指针
	pointer := reflect.ValueOf(op).Pointer()
	// 获取函数名
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling func %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// 因为值传递的原因，这样的交换不起作用
func swap(a, b int) {
	b, a = a, b
}

// 传入指针（变量的地址）是可以的
func swapWithPointer(aPointer, bPointer *int) {
	// *aPointer：表示指针aPointer指向的内容
	*aPointer, *bPointer = *bPointer, *aPointer
}

func main() {
	fmt.Println(eval(1, 3, "_"))
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 2, 3))

	// 直接传入一个函数
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 2))

	fmt.Println(sum(2, 3, 4))

	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)
	// &a：取变量a的指针
	swapWithPointer(&a, &b)
	fmt.Println(a, b)
}

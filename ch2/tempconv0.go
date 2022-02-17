package main

import "fmt"

// 声明摄氏温度类型，底层是float64
type Celsius float64

// 华氏温度
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	c := AbsoluteZeroC
	fmt.Printf("%v\n", c)
}

// 将摄氏温度转化为华氏温度
func CToF(c Celsius) Fahrenheit {
	// 如果二者具有相同的底层类型，则可以互相转换
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 类似Java的toString()方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g 摄氏度", c)
}

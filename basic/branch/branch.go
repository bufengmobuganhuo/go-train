package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	// if语句中可以有多个赋值语句
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(grade(0))
	fmt.Println(grade(79))
	fmt.Println(grade(94))
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
		// Panic中断程序执行，报错
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

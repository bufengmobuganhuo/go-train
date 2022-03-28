package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		// 获取到异常
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("this is an error"))
}

func main() {
	tryRecover()
}

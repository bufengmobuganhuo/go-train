package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		// 通道关闭后，这个也会执行完成
		for natural := range naturals {
			squares <- (natural * natural)
		}
		close(squares)
	}()

	for square := range squares {
		fmt.Println(square)
	}
}

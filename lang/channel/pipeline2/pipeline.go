package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for natural := range in {
		out <- natural * natural
	}
	close(out)
}

func printer(in <-chan int) {
	for square := range in {
		fmt.Println(square)
	}
}

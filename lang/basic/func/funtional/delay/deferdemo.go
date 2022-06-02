package main

import (
	"log"
	"time"
)

func main() {
	//bigSlowOperation()
	log.Println("result ", double(2))
}

func double(x int) (result int) {
	defer func() {
		result += x
	}()
	return x + x
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", start)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

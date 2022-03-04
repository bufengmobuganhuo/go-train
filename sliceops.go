package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

// slice的各种操作
func main() {
	fmt.Println("creating slice")
	// 定义一个slice，默认是nil
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	// 定义slice，并且赋值
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// 定义长度16的slice
	s2 := make([]int, 16)

	// 定义len=10, cap=32的slice
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copying slice")
	// 把s1拷贝到s2
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("deleting elements from slice")
	// 删除s2[3], 第二个为可变参数
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front=", front)

	fmt.Println("poping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail=", tail)
	printSlice(s2)
}

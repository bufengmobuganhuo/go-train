package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 左闭右开区间
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s := arr[2:]
	fmt.Println("arr[2:]=", s)
	fmt.Println("arr[:]=", arr[:])

	fmt.Println("after updateSlice(s)")
	// 切片是原数组的一个视图，修改slice会影响原数组，
	// 修改的是切片的第0个
	updateSlice(s)
	// 切片的第0个，对应的是原数组的第2个
	fmt.Println(s)
	fmt.Println(arr)

	fmt.Println("extending slice")
	arr[2] = 2
	s1 := arr[2:6]
	// 会扩展s1，获取到底层数组的元素
	s2 := s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 此时会替换arr最后一位，后面再append只会生成新的slice，而不会再修改数组
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	// s4, s5不再是对arr的一个view，而是一个新数组的view（类似List，旧数组如果没有使用会被回收掉）
	fmt.Println("s3, s4, s5=", s3, s4, s5)
	fmt.Println("arr=", arr)

}

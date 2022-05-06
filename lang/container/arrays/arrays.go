package main

import "fmt"

func changeArr(arr [3]int) {
	arr[0] = 100
}

func changeArrByPointer(arr *[3]int) {
	arr[0] = 100
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	// 不指明数组长度
	arr3 := [...]int{1, 3, 5}

	fmt.Println(arr1, arr2, arr3)

	// 遍历数组的办法
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for i, _ := range arr3 {
		fmt.Println(arr3[i])
	}

	// 这里会报错，因为长度不同，被认为是不同的type
	changeArr(arr2)
	// 数组是值传递（拷贝一份数组到函数中）
	fmt.Println(arr2[0])
	// 传递指针就可以修改原数组
	changeArrByPointer(&arr2)
	fmt.Println(arr2[0])
}

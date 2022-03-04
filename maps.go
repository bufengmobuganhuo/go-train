package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	// m2 == empty map
	m2 := make(map[string]int)

	// m3 == empty map
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("traversing map")
	// map的遍历，遍历是没有顺序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	// 不存在时，得到类型的零值
	causeName, ok := m["cause"]
	fmt.Println(causeName, ok)

	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println("before deleting:", name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("after deleting:", name, ok)

}

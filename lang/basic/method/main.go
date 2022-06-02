package main

import (
	"fmt"

	"mengyu.com/gotrain/lang/basic/method/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Remove(144)
	fmt.Println(x.String())

	x.Clear()

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
}

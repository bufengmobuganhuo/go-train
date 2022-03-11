package main

import (
	"fmt"

	"mengyu.com/gotrain/tree"
)

func main() {
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	// 自动传入一个指针
	root.Right.Left.SetValueWithPointer(4)
	root.Right.Left.Print()
	fmt.Println("\ntraverse")
	root.Traverse()
	fmt.Println()

	pRoot := &root
	pRoot.SetValueWithPointer(4)
	// 虽然pRoot是一个地址，而print的入参是一个值传递的变量，这里会把对应的变量找到，传入print
	pRoot.Print()
}

package main

import (
	"fmt"

	"mengyu.com/gotrain/tree"
)

// 内嵌的方式扩展已有struct，和java的继承作用一样，但不是继承
// 是一种语法糖
type EmbededNode struct {
	*tree.Node
}

// 对已有类型进行扩展
func (myNode *EmbededNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := EmbededNode{myNode.Left}
	right := EmbededNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Node.Print()
}

// 功能和重载一样，但不是重载
func (EmbededNode *EmbededNode) Traverse() {
	fmt.Println("This method is shadowed")
}

func main() {
	root := EmbededNode{&tree.Node{Value: 3}}
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
	// 类似执行父类的方法
	root.Node.Traverse()
	fmt.Println()

	fmt.Println("post order")
	root.postOrder()

	pRoot := &root
	pRoot.SetValueWithPointer(4)
	// 虽然pRoot是一个地址，而print的入参是一个值传递的变量，这里会把对应的变量找到，传入print
	pRoot.Print()
}

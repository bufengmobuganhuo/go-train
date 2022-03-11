package tree

import "fmt"

// 首字母大写，对其他包可见
type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// 为结构定义的方法
// 定义一个专门为打印Node的方法，类似toString()方法
// 也是传值
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 是一个值传递，这里的修改不会影响外面的值
func (node Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) SetValueWithPointer(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}

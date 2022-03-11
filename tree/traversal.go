package tree

// 只要是相同的包，就可以使用tree.*的形式获取到，无论在不在一个文件中
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

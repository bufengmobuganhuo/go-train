package queue

// 表示任何类型
type Queue []interface{}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 这里只允许返回int类型
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

package stack

import (
	"alg/list"
)

type Stack struct {
	list.SingleLinkList
}

// 入栈
func (st *Stack) Push(data interface{}) error {
	err := st.Insert(0, data)
	return err
}

// 出栈
func (st *Stack) Pop() (*list.SingleNode, error) {
	return st.Delete(st.Size - 1)
}

package list

import "errors"

type Node struct {
	Last *Node
	Next *Node
	Data interface{}
}

type List struct {
	Head *Node
	Last *Node
	size int
}

// 获取节点
func (linkList *List) Get(index int) (*Node, error) {
	err := linkList.checkIndex(index)
	if err != nil {
		return nil, err
	}
	temp := linkList.Head
	i := 0
	for true {
		if i == index || temp.Next == nil {
			return temp, err
		}
		temp = temp.Next
		i++
	}
	return nil, err
}

// 插入数据
func (linkList *List) Insert(index int, data interface{}) error {
	err := linkList.checkIndex(index)
	if err != nil {
		return err
	}

	insertNode := new(Node)
	insertNode.Data = data

	if linkList.size == 0 {
		// 首次插入
		linkList.Last = insertNode
		linkList.Head = insertNode
	} else if index == linkList.size {
		lastNode, err := linkList.Get(index - 1)
		if err != nil {
			return err
		}
		lastNode.Next = insertNode
		insertNode.Last = lastNode
		linkList.Last = insertNode
	} else {
		prevNode, err := linkList.Get(index)
		if err != nil {
			return err
		}
		if index == 0 {
			insertNode.Next = prevNode
			prevNode.Last = insertNode
			linkList.Head = insertNode
		}
		insertNode.Last = prevNode.Last
		insertNode.Next = prevNode
		prevNode.Last.Next = insertNode
		prevNode.Last = insertNode
	}
	linkList.size++
	return nil
}

// 校验列表范围
func (linkList *List) checkIndex(index int) error {
	if index < 0 || index > linkList.size {
		return errors.New("元素超出列表范围")
	}
	return nil
}

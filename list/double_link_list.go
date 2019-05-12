package list

import (
	"errors"
)

type Node struct {
	Last *Node
	Next *Node
	Data interface{}
}

type List struct {
	Head *Node
	End  *Node
	Size int
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

	if linkList.Size == 0 {
		// 首次插入
		linkList.End = insertNode
		linkList.Head = insertNode
	} else if index == linkList.Size {
		EndNode, _ := linkList.Get(index - 1)
		EndNode.Next = insertNode
		insertNode.Last = EndNode
		linkList.End = insertNode
	} else {
		prevNode, _ := linkList.Get(index)
		if index == 0 {
			insertNode.Next = prevNode
			prevNode.Last = insertNode
			linkList.Head = insertNode
		} else {
			insertNode.Last = prevNode.Last
			insertNode.Next = prevNode
			prevNode.Last.Next = insertNode
			prevNode.Last = insertNode
		}
	}
	linkList.Size++
	return nil
}

// 删除
func (linkList *List) Delete(index int) (*Node, error) {
	err := linkList.checkIndex(index)
	if err != nil {
		return nil, err
	}
	delNode := new(Node)
	if index == 0 {
		delNode = linkList.Head
		linkList.Head = linkList.Head.Next
	} else if index == linkList.Size-1 {
		delNode = linkList.End
		prevNode, _ := linkList.Get(index - 1)
		prevNode.Next = nil
		linkList.End = prevNode
	} else {
		prevNode, _ := linkList.Get(index - 1)
		delNode = prevNode.Next
		delNode.Next.Last = prevNode
		prevNode.Next = prevNode.Next.Next
	}
	linkList.Size--
	return delNode, err
}

// 更新
func (linkList *List) Update(index int, data interface{}) error {
	err := linkList.checkIndex(index)
	if err != nil {
		return err
	}
	prevNode, _ := linkList.Get(index)
	prevNode.Data = data
	return nil
}

// 遍历输出
func (linkList *List) Output() {
	temp := linkList.Head
	for true {
		if temp != nil {
			temp = temp.Next
		} else {
			break
		}
	}
}

// 按照顺序遍历索引
func (linkList *List) Foreach() []*Node {
	var ret []*Node
	temp := linkList.Head
	for true {
		if temp != nil {
			ret = append(ret, temp)
			temp = temp.Next
		} else {
			break
		}
	}
	return ret
}

// 校验列表范围
func (linkList *List) checkIndex(index int) error {
	if index < 0 || index > linkList.Size {
		return errors.New("元素超出列表范围")
	}
	return nil
}

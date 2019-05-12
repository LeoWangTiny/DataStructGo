package list

import (
	"errors"
	"fmt"
)

// 单向链表
type SingleLinkList struct {
	Head *SingleNode // 单向链表的头节点指针
	End  *SingleNode // 单向链表的尾节点指针
	Size int         // 单向链表的节点个数
}

// 单向链表节点
type SingleNode struct {
	Next *SingleNode // 下一个节点的指针
	Data interface{} // 存放的数据
}

// 插入
func (sl *SingleLinkList) Insert(index int, data interface{}) error {
	err := sl.checkIndex(index)
	if err != nil {
		return err
	}
	insertNode := new(SingleNode)
	insertNode.Data = data
	if sl.Size == 0 {
		// 头部插入
		sl.Head = insertNode
		sl.End = insertNode
	} else if sl.Size == index {
		// 尾部插入
		sl.End.Next = insertNode
		sl.End = insertNode
	} else {
		// 中间插入
		if index == 0 {
			insertNode.Next = sl.Head
			sl.Head = insertNode
		} else {
			prevNode, _ := sl.Get(index - 1)
			insertNode.Next = prevNode.Next
			prevNode.Next = insertNode
		}
	}
	sl.Size++
	return nil
}

// 更新
func (sl *SingleLinkList) Update(index int, data interface{}) error {
	err := sl.checkIndex(index)
	if err != nil {
		return err
	}
	prevNode, _ := sl.Get(index)
	prevNode.Data = data
	return nil
}

// 删除
func (sl *SingleLinkList) Delete(index int) (*SingleNode, error) {
	err := sl.checkIndex(index)
	if err != nil {
		return nil, err
	}
	delNode := new(SingleNode)
	if index == 0 {
		delNode = sl.Head
		sl.Head = sl.Head.Next
	} else if index == sl.Size-1 {
		delNode = sl.End
		prevNode, _ := sl.Get(index - 1)
		prevNode.Next = nil
		sl.End = prevNode
	} else {
		prevNode, _ := sl.Get(index - 1)
		delNode = prevNode.Next
		prevNode.Next = prevNode.Next.Next
	}
	sl.Size--
	return delNode, err
}

// 查询
func (sl *SingleLinkList) Get(index int) (*SingleNode, error) {
	err := sl.checkIndex(index)
	if err != nil {
		return nil, err
	}
	temp := sl.Head
	i := 0
	if index == 0 {
		return sl.Head, nil
	} else if index == sl.Size-1 {
		return sl.End, nil
	}
	for true {
		if i == index {
			return temp, nil
		}
		temp = temp.Next
		i++
	}
	return nil, nil
}

// 校验索引
func (sl *SingleLinkList) checkIndex(index int) error {
	if index < 0 || index > sl.Size {
		return errors.New("索引超出链表范围！")
	}
	return nil
}

// 按照顺序遍历索引
func (sl *SingleLinkList) Foreach() []*SingleNode {
	var ret []*SingleNode
	temp := sl.Head
	for i := 0; i < sl.Size; i++ {
		ret = append(ret, temp)
		temp = temp.Next
	}
	return ret
}

// 遍历输出
func (sl *SingleLinkList) Output() {
	i := 0
	temp := sl.Head
	for true {
		if i == sl.Size {
			break
		}
		fmt.Println(temp)
		temp = temp.Next
		i++
	}
}

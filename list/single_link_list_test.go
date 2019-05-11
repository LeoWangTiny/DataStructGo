package list

import (
	"fmt"
	"testing"
)

var sl = new(SingleLinkList)

// 测试添加
func TestSingleLinkList_Insert(t *testing.T) {
	// 头部插入
	err1 := sl.Insert(0, 1)
	err2 := sl.Insert(0, 2)
	err3 := sl.Insert(0, 3)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error(err1)
		t.Error(err2)
		t.Error(err3)
	}
	fmt.Println(sl.Head, sl.Head.Next, sl.Head.Next.Next)
	// 中间插入
	err4 := sl.Insert(1, 300)
	// 尾部插入
	err5 := sl.Insert(4, 400)
	if err4 != nil || err5 != nil {
		t.Error(err4)
		t.Error(err5)
	}
	fmt.Println(sl.Head, sl.Head.Next, sl.Head.Next.Next, sl.Head.Next.Next.Next, sl.Head.Next.Next.Next.Next)
}

// 测试删除
func TestSingleLinkList_Delete(t *testing.T) {
	_ = sl.Insert(0, 1)
	_ = sl.Insert(0, 2)
	_ = sl.Insert(0, 3)
	_ = sl.Insert(0, 4)
	_ = sl.Insert(0, 5)
	// 头部删除
	fmt.Println(sl.Head, sl.Head.Next, sl.Head.Next.Next, sl.Head.Next.Next.Next, sl.Head.Next.Next.Next.Next)
	delNode, _ := sl.Delete(0)
	fmt.Println(delNode)
	fmt.Println(sl.Head, sl.Head.Next, sl.Head.Next.Next, sl.Head.Next.Next.Next)
	// 尾部删除
	delNode, _ = sl.Delete(3)
	fmt.Println(delNode)
	fmt.Println(sl.Head, sl.Head.Next, sl.Head.Next.Next)
	// 中间删除
	delNode, _ = sl.Delete(1)
	fmt.Println(delNode)
	fmt.Println(sl.Head, sl.Head.Next)
}

// 获取
func TestSingleLinkList_Get(t *testing.T) {
	_ = sl.Insert(0, 1)
	_ = sl.Insert(0, 2)
	_ = sl.Insert(0, 3)
	_ = sl.Insert(0, 4)
	fmt.Println(sl.Get(0))
	fmt.Println(sl.Get(1))
	fmt.Println(sl.Get(2))
	fmt.Println(sl.Get(3))
}

// 更新
func TestSingleLinkList_Update(t *testing.T) {
	_ = sl.Insert(0, 1)
	_ = sl.Insert(0, 2)
	_ = sl.Insert(0, 3)
	fmt.Println(sl.Get(0))
	fmt.Println(sl.Get(1))
	fmt.Println(sl.Get(2))
	_ = sl.Insert(0, 100)
	_ = sl.Insert(1, 200)
	_ = sl.Insert(2, 300)
	fmt.Println(sl.Get(0))
	fmt.Println(sl.Get(1))
	fmt.Println(sl.Get(2))
}

// 遍历
func TestSingleLinkList_Foreach(t *testing.T) {
	_ = sl.Insert(0, 1)
	_ = sl.Insert(0, 2)
	_ = sl.Insert(0, 3)
	fmt.Println(sl.Foreach())
}

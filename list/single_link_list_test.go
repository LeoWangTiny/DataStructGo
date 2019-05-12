package list

import (
	"fmt"
	"testing"
)

var sl = new(SingleLinkList)

// 测试添加
func TestSingleLinkList_Insert(t *testing.T) {
	err0 := sl.Insert(-1, 1)
	fmt.Println(err0)
	// 头部插入
	err1 := sl.Insert(0, 1)
	err2 := sl.Insert(0, 2)
	err3 := sl.Insert(0, 3)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error(err1)
		t.Error(err2)
		t.Error(err3)
	}
	sl.Output()
	// 中间插入
	err4 := sl.Insert(1, 300)
	// 尾部插入
	err5 := sl.Insert(4, 400)
	if err4 != nil || err5 != nil {
		t.Error(err4)
		t.Error(err5)
	}
	sl.Output()
}

// 测试删除
func TestSingleLinkList_Delete(t *testing.T) {
	_ = sl.Insert(0, 1)
	_ = sl.Insert(0, 2)
	_ = sl.Insert(0, 3)
	_ = sl.Insert(0, 4)
	_ = sl.Insert(0, 5)
	sl.Output()
	_, err0 := sl.Delete(-1)
	fmt.Println(err0)
	// 尾部删除
	delNodeEnd, _ := sl.Delete(4)
	fmt.Println(delNodeEnd)
	sl.Output()
	// 头部删除
	delNode, _ := sl.Delete(0)
	fmt.Println(delNode)
	sl.Output()
	// 中间删除
	delNode, _ = sl.Delete(1)
	fmt.Println(delNode)
	sl.Output()
}

// 获取
func TestSingleLinkList_Get(t *testing.T) {
	_ = sl.Insert(0, 1)
	_ = sl.Insert(0, 2)
	_ = sl.Insert(0, 3)
	_ = sl.Insert(0, 4)
	err0 := sl.Insert(-1, 1)
	fmt.Println(err0)
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
	_ = sl.Update(0, 100)
	_ = sl.Update(1, 200)
	_ = sl.Update(2, 300)
	err := sl.Update(-1, 300)
	fmt.Println(sl.Get(0))
	fmt.Println(sl.Get(1))
	fmt.Println(sl.Get(2))
	fmt.Println(err)
}

// 遍历
func TestSingleLinkList_Foreach(t *testing.T) {
	_ = sl.Insert(0, 1)
	_ = sl.Insert(0, 2)
	_ = sl.Insert(0, 3)
	fmt.Println(sl.Foreach())
}

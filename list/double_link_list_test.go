package list

import (
	"fmt"
	"testing"
)

var dsl = new(List)

// 测试添加
func TestDoubleLinkList_Insert(t *testing.T) {
	err0 := dsl.Insert(-1, 1)
	fmt.Println(err0)
	// 头部插入
	err1 := dsl.Insert(0, 1)
	err2 := dsl.Insert(0, 2)
	err3 := dsl.Insert(0, 3)
	dsl.Output()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error(err1)
		t.Error(err2)
		t.Error(err3)
	}
	// 中间插入
	err4 := dsl.Insert(1, 300)
	// 尾部插入
	err5 := dsl.Insert(4, 400)
	if err4 != nil || err5 != nil {
		t.Error(err4)
		t.Error(err5)
	}
	dsl.Output()
}

// 测试删除
func TestDoubleLinkList_Delete(t *testing.T) {
	_ = dsl.Insert(0, 1)
	_ = dsl.Insert(0, 2)
	_ = dsl.Insert(0, 3)
	_ = dsl.Insert(0, 4)
	_ = dsl.Insert(0, 5)
	dsl.Output()
	_, err0 := dsl.Delete(-1)
	fmt.Println(err0)
	// 尾部删除
	delNodeEnd, _ := dsl.Delete(4)
	fmt.Println(delNodeEnd)
	dsl.Output()
	// 头部删除
	delNode, _ := dsl.Delete(0)
	fmt.Println(delNode)
	dsl.Output()
	// 中间删除
	delNode, _ = dsl.Delete(1)
	fmt.Println(delNode)
	sl.Output()
}

// 获取
func TestDoubleLinkList_Get(t *testing.T) {
	_ = dsl.Insert(0, 1)
	_ = dsl.Insert(0, 2)
	_ = dsl.Insert(0, 3)
	_ = dsl.Insert(0, 4)
	_, err0 := dsl.Get(-1)
	fmt.Println(err0)
	fmt.Println(dsl.Get(0))
	fmt.Println(dsl.Get(1))
	fmt.Println(dsl.Get(2))
	fmt.Println(dsl.Get(3))
}

// 更新
func TestDoubleLinkList_Update(t *testing.T) {
	_ = dsl.Insert(0, 1)
	_ = dsl.Insert(0, 2)
	_ = dsl.Insert(0, 3)
	fmt.Println(dsl.Get(0))
	fmt.Println(dsl.Get(1))
	fmt.Println(dsl.Get(2))
	_ = dsl.Update(0, 100)
	_ = dsl.Update(1, 200)
	_ = dsl.Update(2, 300)
	err := dsl.Update(-1, 300)
	fmt.Println(dsl.Get(0))
	fmt.Println(dsl.Get(1))
	fmt.Println(dsl.Get(2))
	fmt.Println(err)
}

// 遍历
func TestDoubleLinkList_Foreach(t *testing.T) {
	_ = dsl.Insert(0, 1)
	_ = dsl.Insert(0, 2)
	_ = dsl.Insert(0, 3)
	fmt.Println(dsl.Foreach())
}

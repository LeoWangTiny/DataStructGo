package stack

import (
	"fmt"
	"testing"
)

var st = new(Stack)

func TestStack_Push(t *testing.T) {
	err := st.Push(1)
	if err != nil {
		t.Error(err)
	}
	err = st.Push(2)
	if err != nil {
		t.Error(err)
	}
	err = st.Push(3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(st)
}

func TestStack_Pop(t *testing.T) {
	err := st.Push(1)
	if err != nil {
		t.Error(err)
	}
	err = st.Push(2)
	if err != nil {
		t.Error(err)
	}
	err = st.Push(3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(st)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
}

package main

import (
	"data_struct/queue"
	"fmt"
)

func main() {
	mt := queue.MockShed{}
	mt.Init([]int{1,2,3,4,5})
	fmt.Println(mt)
	for {
		ele, err := mt.Pop()
		if err != nil {
			break
		} else {
			fmt.Println(ele)
		}
	}
	fmt.Println(mt)
}

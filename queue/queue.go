package queue

import (
	"errors"
)

type Queue struct {
	data   []int
	Length int
}

func (q *Queue) Init(data []int) {
	q.data = data
	q.Length = len(q.data)
}

func (q *Queue) Push(element int) {
	q.data = append(q.data, element)
	q.Length++
}

func (q *Queue) Pop() (int, error) {
	if len(q.data) == 0 {
		return 0, errors.New("队列已经空了")
	}
	ret := q.data[0]
	q.data = q.data[1:]
	q.Length--
	return ret, nil
}

package shed

import (
	"errors"
)

type Shed struct {
	data   []int
	Length int
}

func (s *Shed) Init(data []int) {
	s.data, s.Length = data, len(data)
}

func (s *Shed) Push(element int) {
	s.data = append(s.data, element)
	s.Length++
}

func (s *Shed) Pop() (int, error) {
	if s.Length == 0 {
		return 0, errors.New("栈已经空了")
	}
	ret := s.data[s.Length-1]
	s.data = s.data[:s.Length-1]
	s.Length--
	return ret, nil
}

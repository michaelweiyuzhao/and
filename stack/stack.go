package stack

import (
	"errors"
)

type Stack struct {
	numElements uint
	val []int
}

func (s *Stack) Empty() bool {
	return s.numElements == 0
}

func (s *Stack) Push(v int) error {
	s.numElements++
	s.val = append(s.val, v)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("Empty stack")
	}
	s.numElements--
	return s.val[s.numElements], nil
}

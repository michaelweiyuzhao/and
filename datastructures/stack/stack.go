package stack

import (
	"errors"
)

var ErrorStackEmpty error = errors.New("Stack is empty")

type Stack struct {
	numElements uint
	val []int
}

func (s *Stack) Size() uint {
	return s.numElements
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
		return 0, ErrorStackEmpty
	}
	s.numElements--
	return s.val[s.numElements], nil
}

func (s *Stack) Peek() (int, error) {
	if s.Empty() {
		return 0, ErrorStackEmpty
	}
	return s.val[s.numElements - 1], nil
}

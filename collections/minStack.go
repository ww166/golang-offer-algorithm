package collections

import "fmt"

type MinStack struct {
	data    Stack
	minData Stack
}

type ()

func (s *MinStack) Push(value interface{}) {
	s.data.Push(value)

	fmt.Println(fmt.Sprintf("value: %v", value))
	fmt.Println(fmt.Sprintf("minData.Peek: %v", s.minData.Peek()))

	if fmt.Sprintf("%v", value) <= fmt.Sprintf("%v", s.minData.Peek()) {
		s.minData.Push(value)
	}
}

/**
	Pop the value from the top of stack
**/
func (s *MinStack) Pop() interface{} {
	l := s.data.Length()
	if l > 0 {
		value := s.data.Pop()
		if value == s.minData.Peek() {
			s.minData.Pop()
		}
		return value
	}

	return nil
}

/**
	Fetch the value from the top of stack
**/
func (s *MinStack) Peek() interface{} {
	return s.data.Peek()
}

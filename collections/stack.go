package collections

import "reflect"

type Stack struct {
	data []interface{}
}

/**
	Push value to stack
**/
func (s *Stack) Push(value interface{}) {
	if len(s.data) > 0 {
		typ := reflect.TypeOf(value)
		typ2 := reflect.TypeOf(s.data[0])
		if typ.PkgPath()+"#"+typ.Name() != typ2.PkgPath()+"#"+typ2.Name() {
			panic("pushing different type onto stack")
		}
	}
	s.data = append(s.data, value)
}

/**
	Pop the value from the top of stack
**/
func (s *Stack) Pop() interface{} {
	l := len(s.data)
	if l > 0 {
		value := s.data[l-1]
		s.data[l-1] = nil
		s.data = s.data[:l-1]

		return value
	}

	return nil

}

/**
	Fetch the value from the top of stack
**/
func (s *Stack) Peek() interface{} {
	l := len(s.data)
	if l > 0 {
		value := s.data[l-1]
		return value
	}

	return nil
}

func (s *Stack) Length() int {
	return len(s.data)
}

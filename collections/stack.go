package collections

import "reflect"

type Stack struct {
	Data []interface{}
}

/**
	Push value to stack
**/
func (s *Stack) Push(value interface{}) {
	if len(s.Data) > 0 {
		typ := reflect.TypeOf(value)
		typ2 := reflect.TypeOf(s.Data[0])
		if typ.PkgPath()+"#"+typ.Name() != typ2.PkgPath()+"#"+typ2.Name() {
			panic("pushing different type onto stack")
		}
	}
	s.Data = append(s.Data, value)
}

/**
	Pop the value from the top of stack
**/
func (s *Stack) Pop() interface{} {
	l := len(s.Data)
	if l > 0 {
		value := s.Data[l-1]
		s.Data[l-1] = nil
		s.Data = s.Data[:l-1]

		return value
	}

	return nil

}

/**
	Fetch the value from the top of stack
**/
func (s *Stack) Peek() interface{} {
	l := len(s.Data)
	if l > 0 {
		value := s.Data[l-1]
		return value
	}

	return nil
}

func (s *Stack) Length() int {
	return len(s.Data)
}

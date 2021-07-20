package collections

import "reflect"

type DataInterface interface {
}

type Data struct {
	data []DataInterface
}

type Interface interface {
	Init(*Data) *Interface
}

type Stack struct {
	Self *Data
}

/**
Initialize data
*/
func (s *Stack) Init(other *Data) *Data {
	s.Self = other
	return s.Self
}

/**
	Push value to stack
**/
func (s *Stack) Push(value DataInterface) {
	if len(s.Self.data) > 0 {
		typ := reflect.TypeOf(value)
		typ2 := reflect.TypeOf(s.Self.data[0])
		if typ.PkgPath()+"#"+typ.Name() != typ2.PkgPath()+"#"+typ2.Name() {
			panic("pushing different type onto stack")
		}
	}
	s.Self.data = append(s.Self.data, value)
}

/**
	Pop the value from the top of stack
**/
func (s *Stack) Pop() DataInterface {
	l := len(s.Self.data)
	if l > 0 {
		value := s.Self.data[l-1]
		s.Self.data[l-1] = nil
		s.Self.data = s.Self.data[:l-1]

		return value
	}

	return nil

}

/**
	Fetch the value from the top of stack
**/
func (s *Stack) Peek() DataInterface {
	l := len(s.Self.data)
	if l > 0 {
		value := s.Self.data[l-1]
		return value
	}

	return nil
}

func (s *Stack) Length() int {
	return len(s.Self.data)
}

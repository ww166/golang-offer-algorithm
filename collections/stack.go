package collections

import (
	"errors"
	"reflect"
)

type Stack []any

/**
	Push value to stack
**/
func (s *Stack) Push(value any) {
	if len(*s) > 0 {
		typ := reflect.TypeOf(value)
		typ2 := reflect.TypeOf((*s)[0])
		if typ.PkgPath()+"#"+typ.Name() != typ2.PkgPath()+"#"+typ2.Name() {
			panic("pushing different type onto stack")
		}
	}
	*s = append(*s, value)
}

/**
	Pop the value from the top of stack
**/
func (s *Stack) Pop() (any, error) {
	l := len(*s)
	if l > 0 {
		value := (*s)[l-1]
		(*s)[l-1] = nil
		*s = (*s)[:l-1]

		return value, nil
	}

	return nil, errors.New("empty stack")

}

/**
	Fetch the value from the top of stack
**/
func (s *Stack) Peek() (any, error) {
	l := len(*s)
	if l > 0 {
		value := (*s)[l-1]
		return value, nil
	}

	return nil, errors.New("empty stack")
}

func (s *Stack) Length() int {
	return len(*s)
}

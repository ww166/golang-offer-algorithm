package examples

import (
	"x.localhost/collections"
)

/*
	双栈排序
*/

type _int int
type _float64 float64
type _string string

type LessInterface interface {
	collections.DataInterface
	less(a interface{}) bool
}

func (s _int) less(a interface{}) bool {
	return s < a.(_int)
}

func (s _float64) less(a interface{}) bool {
	return s < a.(_float64)
}

func (s _string) less(a interface{}) bool {
	return s < a.(_string)
}

type LessableStack struct {
	collections.Stack
}

func (s *LessableStack) Init(other *collections.StackData) *collections.StackData {
	s.Self = other
	return s.Self
}

func DoubleStackSort(source LessableStack) {
	help := LessableStack{}
	helpData := &collections.StackData{}
	help.Init(helpData)

	for {
		v := source.Pop()
		if nil != v {
			peek := help.Peek()

			if peek != nil && v.(LessInterface).less(peek) {
				popValue := help.Pop()
				for {
					if popValue == nil {
						help.Push(v)
						break
					}

					lessThanPop := v.(LessInterface).less(popValue)

					if lessThanPop {
						source.Push(popValue)
					} else {
						help.Push(v)
						break
					}
					popValue = help.Pop()
				}

			} else {
				help.Push(v)
			}

		} else {
			break
		}
	}

	for {
		v := help.Pop()
		if nil != v {
			source.Push(v)
		} else {
			break
		}
	}
}

func IntDoubleStackSort(a ...int) LessableStack {
	source_stack := LessableStack{}
	stackData := &collections.StackData{}
	source_stack.Init(stackData)

	for _, v := range a {
		_v := _int(v)
		source_stack.Push(_v)
	}

	DoubleStackSort(source_stack)

	return source_stack
}

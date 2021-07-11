package examples

import (
	"fmt"

	"x.localhost/collections"
)

/*
	双栈排序
*/

type _int int
type _float64 float64
type _string string

type Interface interface {
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

	data []Interface
}

func (s *LessableStack) init() {
	s.Stack.Data = []
}

func DoubleStackSort(source LessableStack) {
	help := LessableStack{}
	help.init()

	for {
		v := source.Pop()
		if nil != v {
			peek := help.Peek()

			if peek != nil && v.(Interface).less(peek) {
				for popValue := help.Pop(); popValue != nil; {
					if v.(Interface).less(popValue) {
						source.Push(popValue)
					} else {
						help.Push(v)
						break
					}

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
	help_stack := LessableStack{}

	source_stack.init()
	help_stack.init()

	for _, v := range a {
		_v := _int(v)
		source_stack.Push(_v)
	}

	DoubleStackSort(source_stack)

	for v := help_stack.Pop(); v != nil; {
		fmt.Println(v)
	}

	return help_stack
}

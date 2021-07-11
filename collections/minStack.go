package collections

import (
	"fmt"
)

type _int int
type _float64 float64
type _string string

type LessOrEqual interface {
	lte(a, b interface{}, defaultValue bool) bool
}

func (s _int) lte(a, b interface{}, defaultValue bool) bool {
	aValue, aOk := a.(int)
	bValue, bOk := b.(int)

	if aOk && bOk {
		return aValue <= bValue
	}

	return defaultValue
}

func (s _float64) lte(a, b interface{}, defaultValue bool) bool {
	aValue, aOk := a.(float64)
	bValue, bOk := b.(float64)

	if aOk && bOk {
		return aValue <= bValue
	}

	return defaultValue
}

func (s _string) lte(a, b interface{}, defaultValue bool) bool {
	aValue, aOk := a.(string)
	bValue, bOk := b.(string)

	if aOk && bOk {
		return aValue <= bValue
	}

	return defaultValue

}

type MinStack struct {
	data    Stack
	minData Stack

	lessOrEqual LessOrEqual
}

// Create MinStack with lte function

func (s *MinStack) New(l LessOrEqual) {
	s.lessOrEqual = l
}

func (s *MinStack) Push(value interface{}) {
	peekValue := s.minData.Peek()
	s.data.Push(value)

	fmt.Println(fmt.Sprintf("value: %v", value))
	fmt.Println(fmt.Sprintf("minData.Peek: %v", peekValue))

	if s.lessOrEqual.lte(value, peekValue, true) {
		fmt.Println(fmt.Sprintf("value: %v is lte %v", value, peekValue))
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
func (s *MinStack) Peek() LessOrEqual {
	return s.data.Peek().(LessOrEqual)
}

func (s *MinStack) Print() {
	println("++++++++++++++++++++++++++++++++++++")
	if nil != s.data.Data {
		for _, v := range s.data.Data {
			print("%v", v)
			print(", ")
		}
	}

	println("----------------------")

	if nil != s.minData.Data {
		for _, v := range s.minData.Data {
			print("%v", v)
			print(", ")
		}
	}

	println("++++++++++++++++++++++++++++++++++++")
}

/*
	整形 最小值堆栈
*/
type IntMinStack struct {
	MinStack

	lessOrEqual _int
}

/*
	初始化 整形 最小值堆栈
	主要目的： 设置 整形比较接口函数 lte
*/
func (s *IntMinStack) Init() {
	s.MinStack.lessOrEqual = s.lessOrEqual
}

/*
	64位 浮点型 最小值堆栈
*/
type Float64MinStack struct {
	MinStack

	lessOrEqual _float64
}

/*
	初始化 64位浮点型 最小值堆栈
	主要目的： 设置 64位浮点型 比较接口函数 lte
*/
func (s *Float64MinStack) Init() {
	s.MinStack.lessOrEqual = s.lessOrEqual
}

/*
	字符串类型 最小值堆栈
*/
type StringMinStack struct {
	MinStack

	lessOrEqual _string
}

/*
	初始化 字符串类型 最小值堆栈
	主要目的： 设置 字符串类型 比较接口函数 lte
*/
func (s *StringMinStack) Init() {
	s.MinStack.lessOrEqual = s.lessOrEqual
}

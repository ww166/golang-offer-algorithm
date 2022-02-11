package collections

import (
	"constraints"
	"errors"
	"fmt"
	"reflect"
)

type _int int
type _float64 float64
type _string string

//type LTE func[V constraints.Ordered](a, b V) bool

func LTE[V constraints.Ordered](a, b V) bool {
	return a <= b
}

type MinStack[V constraints.Ordered] struct {
	data    Stack
	minData Stack
}

func (s *MinStack[V]) New() {

}

func (s *MinStack[V]) Push(value V) {
	peekValue, err := s.minData.Peek()
	s.data.Push(value)

	fmt.Println(fmt.Sprintf("value: %v", value))
	fmt.Println(fmt.Sprintf("minData.Peek: %v", peekValue))

	if err == nil {
		fmt.Println(fmt.Sprintf("value: %v is lte %v", value, peekValue))
		s.minData.Push(value)
	} else {
		//convertedPeekValue := reflect.ValueOf(value).Convert(reflect.TypeOf(peekValue))
		//convertedValue := reflect.ValueOf(value).Convert(reflect.TypeOf(value))
		convertedPeekValue := reflect.ValueOf(peekValue).Interface().(V)
		convertedValue := reflect.ValueOf(value).Interface().(V)

		if convertedValue <= convertedPeekValue {
			fmt.Println(fmt.Sprintf("value: %v is lte %v", value, peekValue))
			s.minData.Push(value)
		}
	}
}

/**
	Pop the value from the top of stack
**/
func (s *MinStack[V]) Pop() (V, error) {
	l := s.data.Length()
	if l > 0 {
		value, _ := s.data.Pop()
		p, _ := s.minData.Peek()
		if value == p {
			p, _ := s.minData.Pop()
			print(p)
		}
		return reflect.ValueOf(value).Interface().(V), nil
	}

	return *new(V), errors.New("empty stack")
}

/**
	Fetch the value from the top of stack
**/
func (s *MinStack[V]) Peek() (V, error) {
	v, err := s.data.Peek()
	return reflect.ValueOf(v).Interface().(V), err
}

func (s *MinStack[V]) Print() {
	println("++++++++++++++++++++++++++++++++++++")
	if nil != s.data {
		for _, v := range s.data {
			print("%v", v)
			print(", ")
		}
	}

	println("----------------------")

	if nil != s.minData {
		for _, v := range s.minData {
			print("%v", v)
			print(", ")
		}
	}

	println("++++++++++++++++++++++++++++++++++++")
}

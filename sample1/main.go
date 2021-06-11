package main

import (
	"fmt"

	"x.localhost/collections"
)

func main() {
	arr := []int{4, 3, 1, 4, 6, 2, 1, 3, 2, 5, 1}

	myStack := &collections.Stack{}

	for _, v := range arr {
		myStack.Push(v)
	}

	fmt.Println(myStack)

	for range arr {
		myStack.Pop()
		fmt.Println(myStack)
	}

	myMinStack := &collections.MinStack{}

	for _, v := range arr {
		myMinStack.Push(v)
	}

	fmt.Println(myMinStack)

	for range arr {
		myMinStack.Pop()
		fmt.Println(myMinStack)
	}
}

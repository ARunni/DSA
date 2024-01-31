package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}
func (s *Stack) pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) size() int {
	return len(s.items)
}

func deleteMiddle(stack *Stack) {
	if stack.size() == 0 {
		fmt.Println("stack is empty")
		return
	}

	middle := stack.size() / 2
	tempStack := &Stack{}

	for i := 0; i < middle; i++ {
		tempStack.push(stack.pop())
	}
	// Skipping middle element
	stack.pop()

	for tempStack.size() > 0 {
		stack.push(tempStack.pop())
	}
}

func printStack(stack *Stack) {
	for i := len(stack.items) - 1; i >= 0; i-- {
		fmt.Printf("%v", stack.items[i])
	}
	fmt.Println()
}

func main()  {
	stack := &Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	fmt.Println("Original Stack")
	printStack(stack)

	fmt.Println("stack after deleting middle element")
	printStack(stack)
	
}
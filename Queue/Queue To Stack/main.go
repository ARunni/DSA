package main

import "fmt"

func QueueToStack(queue []int) map[int]int {
	stack := make(map[int]int)
	for i, v := range queue {
		stack[i] = v
	}
	return stack
}

func main() {
	queue := []int{7, 0, 3, 5, 5, 2, 1}
	fmt.Println("the queue ", queue)
	stack := QueueToStack(queue)
	fmt.Println("the stack ", stack)
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	current := head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println("Original List:")
	printList(head)

	reversed := reverseList(head)
	fmt.Println("Reversed List:")
	printList(reversed)
}

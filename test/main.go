package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func insert(rn *TreeNode, value int) *TreeNode {
	if rn == nil {
		return &TreeNode{Value: value}
	}
	if value <= rn.Value {
		rn.Left = insert(rn.Left, value)
	} else {
		rn.Right = insert(rn.Right, value)
	}
	return rn
}

func PostOrderDisplay(rn *TreeNode) {
	if rn != nil {
		PostOrderDisplay(rn.Left)
		PostOrderDisplay(rn.Right)
		fmt.Println(rn.Value)

	}
}
func InOrderDisplay(rn *TreeNode) {
	if rn != nil {
		InOrderDisplay(rn.Left)
		fmt.Println(rn.Value)
		InOrderDisplay(rn.Right)

	}
}
func PreOrderDisplay(rn *TreeNode) {
	if rn != nil {
		fmt.Println(rn.Value)
		InOrderDisplay(rn.Left)
		InOrderDisplay(rn.Right)

	}
}

func Max(rn *TreeNode) *TreeNode {
	if rn.Right != nil {
		rn = Max(rn.Right)
	}
	return rn
}
func Min(rn *TreeNode) *TreeNode {
	if rn.Left != nil {
		rn = Min(rn.Left)
	}
	return rn
}

func Delete(rn *TreeNode, value int) *TreeNode {
	if rn != nil {
		if value < rn.Value {
			rn.Left = Delete(rn.Left, value)
		} else if value > rn.Value {
			rn.Right = Delete(rn.Right, value)
		} else {
			if rn.Left == nil {
				return rn.Right
			} else if rn.Right == nil {
				return rn.Left
			}
			rn.Value = Min(rn.Right).Value
			rn.Right = Delete(rn.Right, rn.Value)
		}
	}
	return rn
}

func Search(rn *TreeNode, value int) *TreeNode {
	if rn == nil || rn.Value == value {
		return rn

	}
	if value < rn.Value {
		return Search(rn.Left, value)
	}
	return Search(rn.Right, value)

}

func main() {
	var root *TreeNode
	root = insert(root, 6)
	root = insert(root, 9)
	root = insert(root, 25)
	root = insert(root, 7)
	PostOrderDisplay(root)
	fmt.Println("")
	InOrderDisplay(root)
	fmt.Println("")
	fmt.Println(Max(root).Value)
	fmt.Println("min")
	fmt.Println(Min(root).Value)
	fmt.Println("del")
	Delete(root, 7)
	InOrderDisplay(root)
	fmt.Println("search")

	fmt.Println("")
	c := Search(root, 5)
	if c != nil {
		fmt.Println("found")
	} else {

		fmt.Println("not found")
	}
}

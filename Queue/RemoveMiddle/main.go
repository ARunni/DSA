package main

import "fmt"

type queue struct {
	arr  []int
	size int
}

func (q *queue) Enqueue(data int) {
	q.arr = append(q.arr, data)
	q.size++
}
func (q *queue) Dequueue() int {
	firstelement := q.arr[0]
	q.arr = q.arr[1:]
	return firstelement
}
func (q *queue) removeMiddle(midIndex, currentIndex int) {
	if currentIndex == midIndex {
		mid := q.Dequueue()
		fmt.Println("mid element is ", mid)
		return
	}
	temp := q.Dequueue()
	q.removeMiddle(midIndex, currentIndex+1)
	q.Enqueue(temp)
}

func main() {
	q := queue{}
	q.Enqueue(2)
	q.Enqueue(6)
	q.Enqueue(3)
	q.Enqueue(1)
	q.Enqueue(4)
	q.Enqueue(8)
	q.Enqueue(5)
	q.Enqueue(9)
	currentIndex := 0
	midIndex := q.size / 2
	fmt.Println(q.arr)
	q.removeMiddle(midIndex, currentIndex)
	fmt.Println(q.arr)
}

package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enquueue(item int) {
	q.data = append(q.data, item)
}
func (q *Queue) Dequueue() int {
	if len(q.data) == 0 {
		return 0
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value
}
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
func (q *Queue) Display() {
	fmt.Println("Queue elements :", q.data)
}

func main() {
	queue := Queue{}
	queue.Enquueue(30)
	queue.Enquueue(36)
	queue.Enquueue(27)
	queue.Enquueue(90)
	queue.Enquueue(58)
	queue.Display()
	queue.Dequueue()
	queue.Display()
}

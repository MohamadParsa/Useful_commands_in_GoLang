package main

import "fmt"

type stack []int

func (s stack) isEmpty() bool     { return len(s) == 0 }
func (s *stack) push(newItem int) { *s = append(*s, newItem) }
func (s *stack) pop() int {
	if (*s).isEmpty() {
		return 0
	}
	head := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return head
}

type queue []int

func (q queue) isEmpty() bool        { return len(q) == 0 }
func (q *queue) enqueue(newItem int) { *q = append(*q, newItem) }
func (q *queue) dequeue() int {
	if (*q).isEmpty() {
		return 0
	}
	first := (*q)[0]
	*q = (*q)[1:len(*q)]
	return first
}
func main() {
	fmt.Println("stack")
	s := stack{}
	fmt.Println(s.pop())
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())

	fmt.Println("queue")
	q := queue{}
	fmt.Println(q.dequeue())
	q.enqueue(1)
	q.enqueue(2)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())

}

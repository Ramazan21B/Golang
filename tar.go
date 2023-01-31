package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}
type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Top() int {
	if s.size == 0 {
		return 0
	}
	return s.top.value
}
func (s *Stack) Push(value int) {
	node := &Node{value: value, next: s.top}
	s.top = node
	s.size++
}
func (s *Stack) Pop() int {
	if s.size == 0 {
		return 0
	}
	value := s.top.value
	s.top = s.top.next
	return value
}
func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Clear() {
	for i := 0; i < s.Size(); i++ {
		s.Pop()
	}
}
func (s *Stack) Increment() {
	cur := s.top
	for cur != nil {
		cur.value++
		cur = cur.next
	}
}
func (s *Stack) Contains() bool {
	if s.Size() == 0 {
		return false
	} else {
		return true
	}

}
func main() {
	st := &Stack{}
	st.Push(12)
	st.Increment()
	fmt.Println(st.Top())
}

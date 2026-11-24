package main

import "fmt"

// Stack is a fixed-capacity LIFO stack over an array, the classic
// interview formulation where overflow and underflow are explicit.
type Stack struct {
	data []int
	top  int // index of the next free slot
}

// NewStack returns a stack that holds at most capacity values.
func NewStack(capacity int) *Stack {
	return &Stack{data: make([]int, capacity)}
}

// Push adds v on top; false means overflow (stack full).
func (s *Stack) Push(v int) bool {
	if s.top == len(s.data) {
		return false
	}
	s.data[s.top] = v
	s.top++
	return true
}

// Pop removes and returns the top value; false means underflow (empty).
func (s *Stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	s.top--
	return s.data[s.top], true
}

// Top returns the top value without removing it.
func (s *Stack) Top() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	return s.data[s.top-1], true
}

// Size reports how many values are on the stack.
func (s *Stack) Size() int { return s.top }

func main() {
	s := NewStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Push(4)) // overflow
	for s.Size() > 0 {
		v, _ := s.Pop()
		fmt.Print(v, " ")
	}
	fmt.Println()
}

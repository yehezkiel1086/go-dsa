package main

import "fmt"

/**
Push: Adds a new element on the stack.
Pop: Removes and returns the top element from the stack.
Peek: Returns the top (last) element on the stack.
isEmpty: Checks if the stack is empty.
Size: Finds the number of elements in the stack.
*/

type t_stack interface {
	push(int)
	pop() int
	peek() int
	checkEmpty() bool
	size() int
	String() string
}

type Stack struct {
	Top int
	List []int // list of positive integers (-1 means empty)
}

func (s *Stack) push(v int) {
	s.List = append(s.List, v)
	s.Top += 1
}

func (s *Stack) pop() int {
	if len(s.List) != 0 {
		last := s.List[s.size() - 1] // get last element
		s.List = s.List[:s.size() - 1] // remove last element
		s.Top -= 1
		return last
	}
	return -1
}

func (s *Stack) peek() int {
	return s.List[s.size() - 1]
}

func (s *Stack) checkEmpty() bool {
	return s.size() == 0
}

func (s *Stack) size() int {
	return len(s.List)
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.List)
}

func main() {
	lst := []int{1, 2, 3, 4, 5}

	var s t_stack = &Stack{
		Top: 0,
		List: lst,
	}

	// add values
	s.push(10)
	s.push(9)
	s.push(8)
	s.push(7)
	s.push(6)
	fmt.Println(s)

	// pop values
	popped := s.pop()
	fmt.Println(popped, s)

	// peek value
	v := s.peek()
	fmt.Println(v, s)

	// check empty
	fmt.Println(s.checkEmpty())

	// get size
	fmt.Println(s.size())
}

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack[T any] struct {
	zeroVal T
	s       []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		s: []T{},
	}
}

func (s *Stack[T]) Len() int {
	return len(s.s)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Push(n T) {
	s.s = append(s.s, n)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return s.zeroVal, errors.New("Stack Underflow")
	}

	n := s.s[s.Len()-1]
	s.s = s.s[:s.Len()-1]

	return n, nil
}

func printRPN(tokens []string) string {
	// Constraints dictates that the tokens are always valid
	stack := NewStack[string]()

	for _, t := range tokens {
		switch t {
		case "+":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(fmt.Sprintf("(%s + %s)", a, b))
		case "-":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(fmt.Sprintf("(%s - %s)", a, b))
		case "*":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(fmt.Sprintf("(%s * %s)", a, b))
		case "/":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(fmt.Sprintf("(%s / %s)", a, b))
		default:
			stack.Push(t)
		}
	}

	solution, _ := stack.Pop()
	return solution
}

func evalRPN(tokens []string) int {
	// Constraints dictates that the tokens are always valid
	stack := NewStack[int]()

	for _, t := range tokens {
		switch t {
		case "+":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a + b)
		case "-":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a - b)
		case "*":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a * b)
		case "/":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a / b)
		default:
			tNum, _ := strconv.Atoi(t)
			stack.Push(tNum)
		}
	}

	solution, _ := stack.Pop()
	return solution
}

func main() {
	var arr []string

	arr = []string{"2", "1", "+", "3", "*"}
	fmt.Printf("%s = %d\n", printRPN(arr), evalRPN(arr)) // Output: 9

	arr = []string{"4", "13", "5", "/", "+"}
	fmt.Printf("%s = %d\n", printRPN(arr), evalRPN(arr)) // Output: 6

	arr = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Printf("%s = %d\n", printRPN(arr), evalRPN(arr)) // Output: 22
}

package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	slice []int
}

func (s *stack) isEmpty() bool {
	return len((*s).slice) == 0
}

func (s *stack) push(v int) {
	(*s).slice = append((*s).slice, v)
}
func (s *stack) pop() int {
	if (*s).isEmpty() {
		return 0
	}
	lastIndex := len((*s).slice) - 1
	element := (*s).slice[lastIndex]    //getting last value
	(*s).slice = (*s).slice[:lastIndex] //assigning new slice without last element
	return element
}

func main() {
	fmt.Println("res:", evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println("res:", evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println("res:", evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))

}
func evalRPN(tokens []string) int {
	var s stack
	for _, v := range tokens {
		switch v {
		case "+":
			s.push(s.pop() + s.pop())
		case "-":
			first := s.pop()
			sec := s.pop()
			s.push(sec - first)
		case "*":
			s.push(s.pop() * s.pop())
		case "/":
			first := s.pop()
			sec := s.pop()
			s.push(sec / first)
		default:
			i, _ := strconv.Atoi(v)
			s.push(i)
		}
	}
	return s.pop()
}

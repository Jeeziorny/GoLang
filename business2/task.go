package main

import "fmt"

type Task interface {
	solve()
	getOperator() string
	toString() string
}

type simpleTask struct {
	operator string
	arg1 float64
	arg2 float64
	result float64
}

func (s *simpleTask) getOperator() string {
	return s.operator
}

func (s *simpleTask) solve() {
	switch s.operator {
	case "+": s.result = s.arg1 + s.arg2
	case "*": s.result = s.arg1 * s.arg2
	default : fmt.Printf("Incorrect operator in solve")
	}
}

func (s *simpleTask) toString() string {
	return "Task:  " +
		fmt.Sprintf("%.2f", s.arg1) +
		" " +s.operator+" "+fmt.Sprintf("%.2f", s.arg2)+
		" = "+fmt.Sprintf("%.2f", s.result)
}
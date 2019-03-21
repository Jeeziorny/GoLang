package main

import "fmt"

type Task interface {
	solve() float64
	toString() string
}

type simpleTask struct {
	operator string
	arg1 float64
	arg2 float64
}

func (s simpleTask) solve() float64 {
	switch s.operator {
	case "+": return s.arg1 + s.arg2
	case "-": return s.arg1 - s.arg2
	case "/": return s.arg1 / s.arg2
	case "*": return s.arg1 * s.arg2
	default : return 0
	}
}

func (s simpleTask) toString() string {
	return "Task: " + " " + fmt.Sprintf("%.2f", s.arg1) +" " +s.operator+" "+fmt.Sprintf("%.2f", s.arg2)
}

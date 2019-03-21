package main

import (
	"fmt"
	"sync"
)

type taskList interface {
	showTaskList()
}

type simpleTaskList struct {
	tasklist []Task
	mux sync.Mutex
}

func (s simpleTaskList) showTaskList() {
	s.mux.Lock()
	for _, tempTask := range s.tasklist {
		fmt.Println(tempTask.toString())
	}
	s.mux.Unlock()
}

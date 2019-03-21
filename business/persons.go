package main

import (
	"busconst"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func boss(list *simpleTaskList) {
	operators := [4]string{"+", "-", "/", "*"}
	var curOperator = rand.Intn(4)
	for {
		tempTask := simpleTask{operators[curOperator],
			math.Floor(rand.Float64()*100/10),
			math.Floor(rand.Float64()*100/10)}
		list.mux.Lock()
		list.tasklist = append(list.tasklist, tempTask)
		list.mux.Unlock()
		if loud {
			fmt.Println("Boss has appended task: "+tempTask.toString())
		}
		curOperator = rand.Intn(4)
		time.Sleep(time.Duration(busconst.BossPerformance)*time.Millisecond)
	}
}

func worker(tasks *simpleTaskList, magazine *simpleMagazine, id int) {
	for {
		tasks.mux.Lock()
		if len(tasks.tasklist) > 0 {
			tempTask := tasks.tasklist[0]
			tasks.tasklist = tasks.tasklist[1:]
			tasks.mux.Unlock()
			magazine.mux.Lock()
			magazine.products = append(magazine.products,
				fmt.Sprintf(tempTask.toString() + "= %.2f", tempTask.solve()))
			magazine.mux.Unlock()
			if loud {
				fmt.Println("Worker no. ", id, " solved "+tempTask.toString())
			}
		} else {
			tasks.mux.Unlock()
		}
		time.Sleep(time.Duration(busconst.WorkerPerformance) * time.Millisecond)
	}
}

func customer(magazine *simpleMagazine, id int) {
	for {
		magazine.mux.Lock()
		if len(magazine.products) > 0 {
			tempProduct := magazine.products[0]
			magazine.products = magazine.products[1:]
			magazine.mux.Unlock()
			if loud {
				fmt.Println("Customer no. ", id, " bought: "+tempProduct)
			}
		} else {
			magazine.mux.Unlock()
		}
		time.Sleep(time.Duration(busconst.WorkerPerformance)*time.Millisecond)
	}
}


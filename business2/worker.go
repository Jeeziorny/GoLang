package main

import (
	"fmt"
	"math/rand"
	"time"
)

func workerFactory() func(id int,
						  listOutRequest chan <- bool,
						  listOutStream <- chan Task,
						  warehouse chan <- Task,
						  addingMachines []chan Task,
						  multMachines []chan Task) {
	if decide() {
		fmt.Println("Nervous worker: start")
		return nervousWorker
	} else {
		fmt.Println("Calm worker: start")
		return calmWorker
	}
}

func nervousWorker(id int,
				   listOutRequest chan <- bool,
				   listOutStream <- chan Task,
				   warehouse chan <- Task,
				   addingMachines []chan Task,
				   multMachines []chan Task) {
	var t Task
	done := 0
	for {
		listOutRequest <- true
		select {
		case t = <- listOutStream:
			if loud {
				fmt.Println("nW", id, " take ", t.toString())
			}
			if t.getOperator() == "+" {
				nervousSolutionProcess(addingMachines[:], &t)
				warehouse <- t
			} else {
				nervousSolutionProcess(multMachines[:], &t)
				warehouse <- t
			}
		}
		done++
		nChan <- id
		if loud {
			fmt.Println("nW", id, " leave ", t.toString())
		}
		time.Sleep(time.Duration(1000.0 - 1000.0*WorkPerformance+ float64(rand.Intn(100)))*time.Millisecond)
	}
}

func nervousSolutionProcess(machines []chan Task, t *Task) {
	var done = false
	for !done {
		for _, ch := range machines {
			select {
			case ch <- *t:
				select {
				case *t = <- ch:
					done = true
				}
			}
			if done { break }
		}
	}
}

func calmWorker(id int,
				listOutRequest chan <- bool,
				listOutStream <- chan Task,
				warehouse chan <- Task,
				addingMachines []chan Task,
				multMachines []chan Task) {
	var t Task
	done := 0
	for {
		listOutRequest <- true
		select {
		case t = <- listOutStream:
			if loud {
				fmt.Println("cW", id, " take ", t.toString())
			}
			if t.getOperator() == "+" {
				calmSolutionProcess(addingMachines[:], &t)
				warehouse <- t
			} else {
				calmSolutionProcess(multMachines[:], &t)
				warehouse <- t
			}
		}
		done++
		cChan <- id
		if loud {
			fmt.Println("cW", id, " leave ", t.toString())
		}
		time.Sleep(time.Duration(1000.0 - 1000.0*WorkPerformance+ float64(rand.Intn(100)))*time.Millisecond)
	}
}

func calmSolutionProcess(machines []chan Task, t *Task) {
	var done = false
	n := len(machines)
	n = rand.Intn(n)
	for !done {
		select {
		case machines[n] <- *t:
			select {
			case *t = <- machines[n]:
				done = true
			}
		}
	}
}

func decide() bool {
	return rand.Intn(2) == 0
}

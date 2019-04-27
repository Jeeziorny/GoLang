package main

import (
	"fmt"
	"math/rand"
	"time"
)

func machine(feeder chan Task, id int) {
	var task Task
	for {
		select {
		case task = <- feeder:
			if loud {
				fmt.Println("M ", id, "get: ", task.toString())
			}
			task.solve()
			time.Sleep(time.Duration(1000.0 - 1000.0*MachPerformance+ float64(rand.Intn(100)))*time.Millisecond)
			if loud {
				fmt.Println("M ", id, "set: ", task.toString())
			}
			feeder <- task
		}
	}
}
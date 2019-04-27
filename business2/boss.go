package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func boss(taskList chan <- Task) {
	for {
		var t Task = &simpleTask{
			getOperator(),
			math.Floor(rand.Float64() * 100 / 10),
			math.Floor(rand.Float64() * 100 / 10),
			-1}
		if loud {
			fmt.Println("Boss add", t.toString())
		}
		taskList <- t
		time.Sleep(time.Duration(1000.0 - 1000.0*BossPerformance+ float64(rand.Intn(100)))*time.Millisecond)
	}
}

func getOperator() string {
	if rand.Intn(2) == 0 {
		return "+"
	} else {
		return "*"
	}
}
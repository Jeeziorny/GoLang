package main

import (
	"fmt"
	"math/rand"
	"time"
)

var loud = false
var MACHINES_NUMBER = 2
var nChan = make(chan int, 100)
var cChan = make(chan int, 100)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	setMode()
	srun()
}

func srun() {
	listInputStream := make(chan Task, 100)
	listOutputStream:= make(chan Task, 100)
	warehouseInputStream := make(chan Task, 100)
	warehouseOutputStream:= make(chan Task, 100)

	listOutRequest := make(chan bool, 100)
	warehouseOutRequest := make(chan bool, 100)

	listPrintRequest := make(chan bool, 100)
	warehousePrintRequest := make(chan bool, 100)

	workersStatRequest := make(chan bool, 5)

	go statManager(workersStatRequest)

	go TList(listInputStream,
			listOutputStream,
			listOutRequest,
			listPrintRequest)
	go TList(warehouseInputStream,
			warehouseOutputStream,
			warehouseOutRequest,
			warehousePrintRequest)

	go boss(listInputStream)

	machineId := 0
	var addMachines [amNum]chan Task
	for i := range addMachines {
		addMachines[i] = make(chan Task)
		go machine(addMachines[i], machineId)
		machineId++
	}

	var mulMachines [mmNum]chan Task
	for i := range mulMachines {
		mulMachines[i] = make(chan Task)
		go machine(mulMachines[i], machineId)
		machineId++
	}

	for i := 0; i < WNum; i++ {
		worker := workerFactory()
		go worker(i,
				  listOutRequest,
				  listOutputStream,
				  warehouseInputStream,
				  addMachines[:],
				  mulMachines[:])
	}

	for i := 0; i < CNum; i++ {
		go customer(warehouseOutRequest, warehouseOutputStream, i)
	}

	if !loud {
		for {
			fmt.Println("Press: ")
			fmt.Println("1. for browse the magazine")
			fmt.Println("2. for browse the taskList")
			fmt.Println("3. for check stats")
			var choose int
			fmt.Scan(&choose)
			switch choose {
			case 1:
				warehousePrintRequest <- true
			case 2:
				listPrintRequest <- true
			case 3:
				workersStatRequest <- true
			default:

			}
		}
	}

	for {
		time.Sleep(time.Second * 100)
	}
}

func setMode() {
	fmt.Println("Press: ")
	fmt.Println("1. for talkative mode")
	fmt.Println("2. for quiet mode")
	var mode int
	fmt.Scanf("%d", &mode)
	switch mode {
	case 1 :
		loud = true
	case 2 :
		loud = false
	default:
		fmt.Printf("Bye")
	}
}
package main

import (
	"busconst"
	"fmt"
	"math/rand"
	"time"
)

var loud bool = false

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//setMode()
	run()
}

func run() {
	list := simpleTaskList{tasklist: make([]Task, 0)}
	magazine := simpleMagazine{products: make([]string, 0)}
	//boss
	go boss(&list)
	//workers
	for i := 0 ; i < busconst.NumOfWorkers ; i++ {
		go worker(&list, &magazine, i)
		time.Sleep(time.Duration(rand.Intn(5))*time.Millisecond)
	}
	//customers
	for i := 0 ; i < busconst.NumOfCustomers ; i++ {
		go customer(&magazine, i)
	}
	setMode()
	if !loud {
		go interaction(&list, &magazine)
	}
	time.Sleep(60 * time.Second)
}

func interaction(list *simpleTaskList, magazine *simpleMagazine) {
	for {
		fmt.Println("Press: ")
		fmt.Println("1. for browse the magazine")
		fmt.Println("2. for browse the taskList")
		var choose int
		fmt.Scan(&choose)
		switch choose {
		case 1:
			magazine.showMagazine()
		case 2:
			list.showTaskList()
		default:
			fmt.Println("Incorrect input")
		}
	}
}

func setMode() {
	fmt.Println("Press: ")
	fmt.Println("1. for talkative mode")
	fmt.Println("2. for quiet mode")
	var mode int
	fmt.Scanf("%d", &mode)
	switch(mode) {
	case 1 :
		loud = true
	case 2 :
		loud = false
	default:
		fmt.Printf("Bye")
	}
}
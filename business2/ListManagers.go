package main

import "fmt"

func TList(inStream <- chan Task,
	       outStream chan <- Task,
	       outRequest <- chan bool,
	       printRequest chan bool) {
	var myTasks []Task
	var tempTask Task
	pendingRequest := 0
	for {
		for pendingRequest > 0 && len(myTasks) > 0 {
			outStream <- myTasks[0]
			myTasks = myTasks[1:]
			pendingRequest--
		}
		select {
		case tempTask = <- inStream:
			myTasks = append(myTasks, tempTask)
		case _ = <- outRequest:
			pendingRequest++
			if len(myTasks) > 0 {
				outStream <- myTasks[0]
				myTasks = myTasks[1:]
			}
		case _ = <- printRequest:
			for _, t := range myTasks {
				fmt.Println(t.toString())
			}
		}
	}
}

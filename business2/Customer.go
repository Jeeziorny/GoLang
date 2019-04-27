package main

import "fmt"

func customer(warehouseOutRequest chan <- bool,
			  warehouseOutStream <- chan Task,
			  id int) {
	var t Task
	for {
		warehouseOutRequest <- true
		select {
		case t = <- warehouseOutStream:
			if loud {
				fmt.Println("C", id, " ", t.toString())
			}
		}
	}
}

package main

import "fmt"

func statManager(request <- chan bool) {
	nStats := make(map[int]int)
	cStats := make(map[int]int)
	var id int
	for {
		select {
		case _ = <- request:
			fmt.Println("Nervous workers:")
			for key, value := range nStats {
				fmt.Println("id = ", key, " val: ", value)
			}
			fmt.Println("Calm workers:")
			for key, value := range cStats {
				fmt.Println("id = ", key, " val: ", value)
			}
		case id = <- cChan:
			_, exists := cStats[id]
			if exists {
				cStats[id]++
			} else {
				cStats[id] = 1
			}
		case id = <- nChan:
			_, exists := nStats[id]
			if exists {
				nStats[id]++
			} else {
				nStats[id] = 1
			}
		}
	}
}

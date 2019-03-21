package main

import (
	"fmt"
	"sync"
)

type Magazine interface {
	showMagazine()
}

type simpleMagazine struct {
	products []string
	mux sync.Mutex
}

func (s simpleMagazine) showMagazine() {
	s.mux.Lock()
	for _, s := range s.products {
		fmt.Println(s)
	}
	s.mux.Unlock()
}

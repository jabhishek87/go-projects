package main

import (
	"fmt"
)

func subChan() {
	villians := []string{"Amrish", "Pran", "Prem", "Rocky", "Billy"}

	// Un buffered channel
	// finishedBeating := make(chan)

	// buffered channel
	finishedBeating := make(chan bool)

	for _, villian := range villians {
		go attackCH(villian, finishedBeating)
		fmt.Println(<-finishedBeating)
	}
}

func attackCH(vName string, finish chan bool) {
	fmt.Println("Hero Punishing to ", vName)
	finish <- true
}

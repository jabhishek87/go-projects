package main

import (
	"fmt"
	"sync"
)

func main6() {
	var beeper sync.WaitGroup
	evilNinjas := []string{"Tommy", "Johnny", "Bobby"}
	beeper.Add(len(evilNinjas))
	for _, ninja := range evilNinjas {
		go attack6(ninja, &beeper)
	}

	beeper.Wait()
	fmt.Println("Mission Complete")
	// time.Sleep(2 * time.Second)
}

func attack6(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", evilNinja)
	beeper.Done()
}

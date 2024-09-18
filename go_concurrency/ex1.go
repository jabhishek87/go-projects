package main

import (
	"fmt"
	"time"
)

func subMod() {
	villians := []string{"Amrish", "Pran", "Prem", "Rocky", "Billy"}

	for _, villian := range villians {
		go attack(villian)
	}
}

func attack(vName string) {
	fmt.Println("Hero Punishing to ", vName)
	time.Sleep(time.Second * 1)
}

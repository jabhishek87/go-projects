package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bufChan2() {
	channel := make(chan string)
	go throwingNinjaStar2(channel)
	for msg := range channel {
		fmt.Println(msg)
	}
}

func throwingNinjaStar2(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 10
	for i := range numRounds {
		score := rand.Intn(10)
		channel <- fmt.Sprint(i, " scored: ", score)
	}
	close(channel)
}

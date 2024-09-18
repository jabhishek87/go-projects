package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bufChan() {
	channel := make(chan string)
	numRounds := 10
	go throwingNinjaStar(channel, numRounds)
	for range numRounds {
		fmt.Println(<-channel)
	}
}

func throwingNinjaStar(channel chan string, numRounds int) {
	rand.Seed(time.Now().UnixNano())
	for i := range numRounds {
		score := rand.Intn(10)
		channel <- fmt.Sprint(i, " scored: ", score)
	}

}

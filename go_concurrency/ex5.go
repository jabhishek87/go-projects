package main

import "fmt"

func main5() {
	ninja1, ninja2 := make(chan string), make(chan string)
	go captailElect(ninja1, "ninja1")
	go captailElect(ninja2, "ninja2")

	select {
	case msg := <-ninja1:
		fmt.Println(msg)
	case msg := <-ninja2:
		fmt.Println(msg)
	default:
		fmt.Println("Neither")
	}
	voteSimulation()
}

func captailElect(ninja chan string, message string) {
	ninja <- message
}

func voteSimulation() {
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)

	var ninja1Count, ninja2Count int

	for range 1000 {
		select {
		case <-ninja1:
			ninja1Count++

		case <-ninja2:
			ninja2Count++
		}
	}

	fmt.Printf("ninja1Count: %d, ninja2Count: %d\n", ninja1Count, ninja2Count)

}

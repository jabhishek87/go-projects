package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	lock  sync.Mutex
	i     int
)

func main7() {
	iterations := 10000
	for range iterations {
		// fmt.Println(num)
		go goInc()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Resulted count is :", count)
}

func goInc() {
	lock.Lock()
	count++
	lock.Unlock()

}

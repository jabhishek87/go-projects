package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func mul(num int) int {
	time.Sleep(time.Second)
	return num * 2
}

func processData(wg *sync.WaitGroup, result *[]int, i int) {
	defer wg.Done()
	res := mul(i)
	*result = append(*result, res)
}

func processDataCon(wg *sync.WaitGroup, result *int, i int) {
	defer wg.Done()
	res := mul(i)
	*result = res
}

func processDataMutexQuick(wg *sync.WaitGroup, result *[]int, i int) {

	defer wg.Done()
	res := mul(i)
	lock.Lock()
	*result = append(*result, res)
	lock.Unlock()
}

func processDataMutex(wg *sync.WaitGroup, result *[]int, i int) {
	lock.Lock()
	defer wg.Done()
	res := mul(i)
	*result = append(*result, res)
	lock.Unlock()
}

func concurR(num int) {
	var wg sync.WaitGroup

	startTime := time.Now()
	result := make([]int, num)

	for i := range num {
		wg.Add(1)
		go processDataCon(&wg, &result[i], i)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(startTime))
}

func concurMutex(num int) {
	var wg sync.WaitGroup

	startTime := time.Now()
	result := []int{}

	for i := range num {
		wg.Add(1)
		go processDataMutexQuick(&wg, &result, i)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(startTime))
}

func paraMutex(num int) {
	var wg sync.WaitGroup

	startTime := time.Now()
	result := []int{}

	for i := range num {
		wg.Add(1)
		go processDataMutex(&wg, &result, i)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(startTime))
}

func inconsistentFunc(num int) {
	var wg sync.WaitGroup

	startTime := time.Now()
	result := []int{}

	for i := range num {
		wg.Add(1)
		go processData(&wg, &result, i)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(startTime))
}

func parallelLoop(num int) {
	startTime := time.Now()
	result := []int{}

	for i := range num {
		result = append(result, mul(i))
	}
	fmt.Println(result)
	fmt.Println(time.Since(startTime))
}

func main() {

	// https://www.youtube.com/watch?v=Bk1c30avsuU

	n := 10
	// it wil take n Seconds
	// parallelLoop(n)

	// it will run in few seconds but give incosistent result
	//inconsistentFunc(n)

	// run same as parallel
	// paraMutex(n)

	// run cocurrently increase the num and see no difference
	concurMutex(n)

	// confinement no need og lock and un lock
	concurR(n)
}

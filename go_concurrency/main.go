package main

import (
	"fmt"
	"time"
)

// https://www.youtube.com/watch?v=2HsMsbMDwsg&list=PLve39GJ2D71wSwRQLp_h8B60pKgS85StC&index=1
func main() {
	startTime := time.Now()
	defer func() {
		fmt.Println("Total time taken " + time.Since(startTime).String())
	}()
	// go subMod()
	// go subChan()
	// bufChan()
	// main5()
	main7()
	// time.Sleep(time.Second * 2)
}

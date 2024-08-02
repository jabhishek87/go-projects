package main

func main() {

	tri := triangle{height: 3, base: 3}
	squ := square{sideLength: 4}

	printArea(tri)
	printArea(squ)

	// // int2
	// f, err := os.Open(os.Args[1])

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	// io.Copy(os.Stdout, f)

}

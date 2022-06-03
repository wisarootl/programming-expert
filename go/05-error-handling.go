package main

import "fmt"

func main() {
	//  === Error Handling ===
	// most of error in Go is in complie-time
	// panic is run-time error
	fmt.Println("\n=== Error Handling ===")
	fmt.Println("hello")

	// defer the execution to the end of function
	// defer will be also executed before panic
	defer fmt.Println("defer")
	fmt.Println("goodbye")
	panic("I failed")
	fmt.Println("sth after panic")
}
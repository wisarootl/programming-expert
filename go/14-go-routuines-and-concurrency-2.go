// GoRoutine is similar to thread in python. all concurrent execution of multiple pieces code at the same time
// GoRoutine do not have limitation of "Global Interpreter Rock" like Python.
// GoRoutines is literally run in parallel
package main

import (
	"fmt"
	"time"
)

// channel : return for GoRoutine
// rv = channel
func add(a int, b int, rv chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println(a + b)
	rv <- a + b
}

func main() {
	//  === Go Routines and Concurrency (cont.) ===
	fmt.Println("\n=== Go Routines and Concurrency ===")
	returnChan := make(chan int)
	go add(1, 2, returnChan)
	rv := <- returnChan
	go add(3, 4, returnChan)
	rv = <- returnChan
	go add(6, 8, returnChan)
	rv = <- returnChan
	go add(9, 8, returnChan)
	rv = <- returnChan

	fmt.Println("done")
	fmt.Println(rv)
	fmt.Println(returnChan)
}
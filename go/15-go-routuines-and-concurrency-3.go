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
// directed channel
// rv <- chan : receive-only channel (<-rv)
// rv chan<- : send-only channel (rv<-a)
func add(a int, b int, delay int, rv chan int) {
	time.Sleep(time.Duration(delay) * time.Second)
	rv <- a + b
}

func main() {
	//  === Go Routines and Concurrency (cont.) ===
	fmt.Println("\n=== Go Routines and Concurrency ===")
	chan1 := make(chan int)
	chan2 := make(chan int)

	go add(5, 5, 5, chan1)
	go add(10, 10, 2, chan2)

	
	fmt.Println("\n= wait 5 sec get 10 and 20")
	rv1 := <-chan1
	rv2 := <-chan2
	fmt.Println(rv1)
	fmt.Println(rv2)

	fmt.Println("\n= wait 2 sec get 20 and wait 3 sec get 10")
	go add(5, 5, 5, chan1)
	go add(10, 10, 2, chan2)
	rv2 = <-chan2
	fmt.Println(rv2)
	rv1 = <-chan1
	fmt.Println(rv1)
	
	fmt.Println("\n= get faster one first (10)")
	go add(5, 5, 1, chan1)
	go add(10, 10, 2, chan2)
	select {
	case rv1 = <-chan1:
		fmt.Println(rv1)
	case rv2 = <-chan2:
		fmt.Println(rv2)
	}
	// kill channel chan2
	rv2 = <-chan2

	fmt.Println("\n= get faster one first (10) and then get slower one (20)")
	go add(5, 5, 1, chan1)
	go add(10, 10, 2, chan2)

	for i := 0; i < 2; i++ {
		select {
		case rv1 = <-chan1:
			fmt.Println(rv1)
		case rv2 = <-chan2:
			fmt.Println(rv2)
		}
	}

}
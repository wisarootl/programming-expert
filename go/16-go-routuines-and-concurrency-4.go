package main

import "fmt"

func main() {
	//  === Go Routines and Concurrency (cont.) ===
	fmt.Println("\n=== Go Routines and Concurrency ===")
	// buffer channel
	ch := make(chan int, 2) // buffer = 2
	// ch := make(chan int) // if not specify buffer like this, we will get deadlock
	ch <- 1
	fmt.Println(<-ch)
}
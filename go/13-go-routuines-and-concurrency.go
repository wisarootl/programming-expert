// GoRoutine is similar to thread in python. all concurrent execution of multiple pieces code at the same time
// GoRoutine do not have limitation of "Global Interpreter Rock" like Python.
// GoRoutines is literally run in parallel
package main

import (
	"fmt"
	"time"
)

func run1() {
	time.Sleep(3 * time.Second)
	fmt.Println("run1")
}

func run2() {
	time.Sleep(2 * time.Second)
	fmt.Println("run2")
}

func run3() {
	time.Sleep(1 * time.Second)
	fmt.Println("run3")
}

func main() {
	//  === Go Routines and Concurrency ===
	fmt.Println("\n=== Go Routines and Concurrency ===")
	// run1()
	// run2()
	// run3()
	go run1()
	go run2()
	go run3()
	// time.Sleep(4 * time.Second)
	fmt.Println("done")
	time.Sleep(4 * time.Second)
}
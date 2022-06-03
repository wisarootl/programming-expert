package main

import "fmt"

func doPanic() {
	panic("fail")
}

func handlePanic() {
	r := recover()
	if r != nil {
		fmt.Println(r)
		fmt.Println("r")
	}
}

func main() {
	//  === Error Handling (cont.) ===
	fmt.Println("\n=== Error Handling ===")
	defer handlePanic()
	fmt.Println("start")
	doPanic()
	fmt.Println("end")
}
package main

import "fmt"

func SumPartialRange(start uint, end uint, result chan<- uint) {
	sum := uint(0)
	for i := start; i <= end; i++ {
		sum += i
	}
	result <- sum
}

func SumToN(n uint) uint {
	chan1 := make(chan uint)
	chan2 := make(chan uint)

	go SumPartialRange(0, n/2, chan1)
	go SumPartialRange(n/2+1, n, chan2)

	sum := <-chan1 + <-chan2
	return sum
}

func main() {
	fmt.Println(SumToN(5))
}
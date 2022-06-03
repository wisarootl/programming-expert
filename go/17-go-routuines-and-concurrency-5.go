package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func increment(counter *Counter, ch chan<- bool) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.value = counter.value + 1
	fmt.Println(counter.value)
	ch <- true
}

func main() {
	//  === Go Routines and Concurrency (cont.) ===
	fmt.Println("\n=== Go Routines and Concurrency ===")
	counter := Counter{value: 0, mutex: sync.Mutex{}}
	ch := make(chan bool)

	for i := 0; i < 100; i++ {
		go increment(&counter, ch)
	}

	results := 0
	for results < 100 {
		<-ch
		results++
	}
	time.Sleep(2 * time.Second)
}
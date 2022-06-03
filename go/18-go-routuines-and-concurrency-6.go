package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func increment(counter *Counter, wg *sync.WaitGroup) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.value = counter.value + 1
	fmt.Println(counter.value)
	wg.Done()
}

func main() {
	//  === Go Routines and Concurrency (cont.) ===
	fmt.Println("\n=== Go Routines and Concurrency ===")
	// wait group
	wg := sync.WaitGroup{}
	counter := Counter{value: 0, mutex: sync.Mutex{}}
	
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go increment(&counter, &wg)
	}

	wg.Wait()
}
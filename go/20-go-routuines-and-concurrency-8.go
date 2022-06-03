package main

import (
	"fmt"
	"sync"
)

type MappedSums struct {
	mutex sync.Mutex
	sums  map[string]uint
}

func (m *MappedSums) Increment(key string, value uint) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.sums[key]
	if !ok {
		m.sums[key] = 0
	}
	m.sums[key] += value
}

func CreateMappedSums(keys []string, sumTo []uint) map[string]uint {
	mappedSums := MappedSums{sums: make(map[string]uint)}
	var wg sync.WaitGroup

	doIncrement := func(key string, n uint) {
		for i := uint(1); i <= n; i++ {
			mappedSums.Increment(key, i)
		}
		wg.Done()
	}

	wg.Add(len(keys))

	for i, key := range keys {
		go doIncrement(key, sumTo[i])
	}

	wg.Wait()

	return mappedSums.sums
}

func main() {
	keys := []string{"a", "b", "c"}
	sumTo := []uint{10, 10, 5}
	fmt.Println(CreateMappedSums(keys, sumTo))
}
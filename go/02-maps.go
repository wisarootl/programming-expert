package main

import "fmt"

func main() {
	// === Maps ===
	fmt.Println("\n=== Maps ===")
	// map is like a dictionary in python
	var mp map[string]int
	fmt.Println(mp)

	mp2 := map[string]int{"h": 1, "a": 3}
	fmt.Println(mp2)

	mp3 := map[string][]int{"h": {1}, "a": {3}}
	fmt.Println(mp3)

	mp4 := make(map[int]rune)
	mp4[5] = 2
	fmt.Println(mp4)
	value, is_key_existing := mp4[5]
	fmt.Println(value, is_key_existing)
	delete(mp4, 5)
	fmt.Println(mp4)

	mp5 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "z": 5, "f": 6}
	for key, value := range mp5 {
		fmt.Println(key,value)
	}

	n := 100
	divisibleMap := make(map[int]uint)

	for num := 1; num <= n; num++ {
		for d := 1; d <= 5; d++ {
			if num % d == 0 {
				divisibleMap[d]++
			}
		}
	}

	fmt.Println(divisibleMap)

}
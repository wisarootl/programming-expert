package main

import "fmt"

type MultiplyStringResult struct {
	str string
	idx int
}

func multiplyString(str string, factor uint, idx int, result chan<- MultiplyStringResult) {
	newString := ""

	for i := uint(0); i < factor; i++ {
		newString = newString + str
	}

	resultStruct := MultiplyStringResult{newString, idx}
	result <- resultStruct
}

func MultiplyStringsConcurrently(strings []string, factor uint) []string {
	multipliedStrings := make([]string, len(strings))
	resultsChannel := make(chan MultiplyStringResult)

	for idx, str := range strings {
		go multiplyString(str, factor, idx, resultsChannel)
	}

	resultsGathered := 0
	for resultsGathered < len(strings) {
		multipliedString := <-resultsChannel
		resultsGathered++
		multipliedStrings[multipliedString.idx] = multipliedString.str
	}

	return multipliedStrings
}

func main() {
	strings := []string{"bird", "plane", "superman"}
	factor := uint(3)
	fmt.Println(MultiplyStringsConcurrently(strings, factor))
}
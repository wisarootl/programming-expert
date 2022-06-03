package main

import (
	"errors"
	"fmt"
	"strconv"
)

func SumNumericStrings(strings []string) (int, error) {
	sum := 0
	for _, str := range strings {
		if len(str) >= 20 {
			return 0, errors.New("At least one of the strings in the input array contains 20 characters or more.")
		}

		value, err := strconv.Atoi(str)
		
		if err != nil {
			return 0, errors.New("At least one of the strings in the input array does not represent an integer.")
		}

		sum += value
	}
	return sum, nil
}


func main() {
	//  === Error Handling (cont.) ===
	fmt.Println("\n=== Error Handling ===")
	strings := []string{"1234", "234", "1", "7", "0"}
	fmt.Println(SumNumericStrings(strings))
}

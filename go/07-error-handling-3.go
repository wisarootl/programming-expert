package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	//  === Error Handling (cont.) ===
	fmt.Println("\n=== Error Handling ===")
	result, err := divide(1, 0)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}
}
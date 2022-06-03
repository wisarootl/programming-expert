package main

import (
	"fmt"
)

func test(a int, b int) (x int, y int) {
	x = a
	y = b
	return
}

// ...int is similar to *args in python
func sumInts(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func returnFunc(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func callFunc(callable func(string) string, arg string) {
	result := callable(arg)
	fmt.Println(result)
}

func main() {
	//  === Functions ===
	fmt.Println("\n=== Functions ===")
	// Named return values
	a, b := test(1, 2)
	fmt.Println(a, b)


	// Variadic function (accept ... (equivalent to *args in python))
	fmt.Println("\n= Variadic function")
	nums := []int{1, 2, 3, 4, 5, 6}
	// sum := sumInts(1, 2, 3, 4, 5, 6)
	sum := sumInts(nums...)
	fmt.Println(sum)

	// functions as type
	fmt.Println("\n= functions as type")
	fn := returnFunc(100)
	value := fn(50)
	fmt.Println(value)

	// function closure
	fmt.Println("\n= function closure")
	pos := adder()
	neg := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
	
	// function as parameters
	fmt.Println("\n= function as parameters")
	myFunc := func(str string) string {
		return str + "world!"
	}
	callFunc(myFunc, "hello ")
}

// statically typed (Go) language : requires the programmer to specify the type of each variable,
// either explicitly or implicitly (using something like the := operator in Go).
// These types are then checked by the compiler before the code is executed. This typically
// results in less run-time errors as many typing errors are caught at compile time.
// Go is an example of a statically typed language.

// dynamically typed (Python) language : does not require the programmer to specify the types of their variables.
// This is because most type checking is done at run-time, as opposed to compile time.
// Dynamically typed languages are usually more flexible and allow for variables to change
// the types they store throughout the program. However, dynamically typed languages are prone
// to more run-time errors and can often be difficult to debug.
// Python is an example of a dynamically typed language.

// Strongly typed language (Python, Go) vs. Weakly typed language

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	
	// === Introduction ===
	fmt.Println("=== Introduction ===")
	// fmt.Println("Hello World")
	/*
		hello
		world
	*/
	fmt.Println("Hello World")

	// === Data Types ===
	fmt.Println("\n=== Data Types ===")
	// uint8 : use 8 bit to represent unsigned int with out sign (range = 0 to 255)
	// int8 : use 8 bit to represent int with sign (range = -128 to 127)
	// int int8 int16 int32 int64
	// rune = int32
	// byte = int8 (can store single char)

	// float32
	// float64

	// bool (true, false)

	// string "hello" : string have to be in double quotation mark

	// nil : none type

	// complex128, complex64

	// var x int : we can declare without value. Go will assign default value
	var x int = 2 // declare but not use will get error
	fmt.Println(x)

	const x2 float32 = 9.01 // declare as const, the value cannot
	fmt.Println(x2)

	// === Implicit Assignment ===
	fmt.Println("\n=== Implicit Assignment ===")
	// y := 1 // equivalent to var y int = 1
	y := uint(1)
	// : implicitly guess type of y by the assigned value
	// fmt.Println(y)
	fmt.Printf("%T\n", y)

	// === console output ===
	fmt.Println("\n=== Console Output ===")
	fmt.Println("hello", x2, 2.45) // print line
	fmt.Println("hello", 2, 2.45)

	fmt.Printf("value : %v\ntype : %T\nbinary representation : %b\n", x, x, x) // print format
	fmt.Printf("exponent : %e\n", 100.001) // exponent
	fmt.Printf("float : %f\n", 1.001)
	fmt.Printf("float : %.2f\n", 1.001) // round
	fmt.Printf("string : %s\n", "hi there")
	fmt.Printf("digit : %d\n", 48) // digit or int
	str := fmt.Sprintf("%T", 1) // same as Printf. It just return string
	fmt.Println(str)
	
	// === Arithmetic Operators ===
	fmt.Println("\n=== Arithmetic Operators ===")
	a := 5
	b := int8(2)
	c := float64(a) / float64(b)
	c++
	c--
	fmt.Println(c)
	a2 := math.Max(5, 10)
	// math.Pow(10, 2) = 10^2
	fmt.Println(a2)
	str2 := "1234"
	a, err := strconv.Atoi(str2)
	fmt.Println(a, err)
	
	str3 := "1234"
	a3, err3 := strconv.ParseInt(str3, 10, 64) // (string, base, bit size)
	fmt.Println(a3, err3)

	// === Conditions And Conditionals ===
	fmt.Println("\n=== Conditions And Conditionals ===")
	a4 := 1.2
	b4 := 1
	c4 := a4 < float64(b4) && b4 < 1.0
	// &&, ||, !false, 
	fmt.Println(c4)

	a5 := 1

	if a5 < 2 {
		fmt.Println("a is less than 2")
	} else if a5 == 3 {
		fmt.Println("a is equal to 2")
	} else {
		fmt.Println("a is greater than 2")
	}

	// === Switch Statement ===
	fmt.Println("\n=== Switch Statement ===")
	a6 := 3
	switch a6 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}

	// === Looping (For & While) ===
	fmt.Println("\n=== Looping (For & While) ===")
	// for loop
	fmt.Println("= for loop")
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	// while loop
	fmt.Println("\n= while loop")
	x6 := 0
	for x6 < 3 {
		x6++
		fmt.Println(x6) 
	}

	// string
	fmt.Println("\n= string")
	str6 := "hello" // string in Go == slice (list) of bytes (bytes store ASCII code)
	fmt.Println(str6[0]) // str6[0] return ASCII code of "h"
	fmt.Println(string(str6[0]))
	fmt.Printf("%T\n", str6[0])
	// ASCII use 1 byte (256 char)
	// UTF8 use 4 bytes
	for i := 0; i < len(str6); i++ {
		fmt.Printf("%c ",str6[i])
	}
	fmt.Println("")
	for i, char := range str6 {
		fmt.Println(i, char, string(char))
	}
	fmt.Println("")
	for _, char := range str6 {
		fmt.Printf("%T\n", char)
		break
	}

	// === Arrays and Slices ===
	fmt.Println("\n=== Arrays and Slices ===")
	fmt.Println("= array")
	var arr [2]int
	arr[0] = 100
	fmt.Println(arr)
	fmt.Println(len(arr))

	arr2 := [2]int{4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{4, 5, 7, 7, 6, 6, 2}
	fmt.Println(arr3)

	arr4 := [2][2]string{{"hello", "world"}, {"code", "go"}}
	fmt.Println(arr4)

	for _, arr := range arr4 {
		for j, str := range arr {
			fmt.Println(j, str)
		}
	}

	fmt.Println("\n= slice")
	// slice is more flexible than array
		arr5 := [5]int{1, 2, 3, 4, 5}
		sliceArr := arr5[:4]
		fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
		//  slice
		//     	pointer -> a[0]
		//  		length -> 4
		//			capacity -> 5
		sliceArr = sliceArr[1:]
		fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
		sliceArr[0] = 100
		arr5[3] = 200
		fmt.Println(arr5, sliceArr)

		arr6 := []string{"hello", "world"}
		arr6 = arr6[:1]
		fmt.Println(arr6, len(arr6), cap(arr6))
		arr6 = arr6[:2]
		fmt.Println(arr6, len(arr6), cap(arr6))

		arr7 := []string{}

		for i := 0; i < 10; i++ {
			arr7 = append(arr7, "hello")
			fmt.Println(arr7, len(arr7), cap(arr7))
		}

		arr8 := make([][]int, 5, 10) // make specify size and capacity of slice
		fmt.Println(arr8, len(arr8), cap(arr8))

		arr9 := []bool{}
		arr10 := []bool{true, false, false}
		arr9 = append(arr9, arr10...) // ... = * in python (breakdown element)
		fmt.Println(arr9)

		for i, val := range arr9 {
			fmt.Println(i, val)
		}

		arr11 := []bool{true, false, false}
		test(arr11)
		fmt.Println(arr11)

		arr12 := [3]bool{true, false, false}
		test2(arr12)
		fmt.Println(arr12)

}

// pass slice (more flexible than array) (it is similar to list in Python)
func test(x []bool) {
	x[0] = false
}

// pass array (fix size) and it will pass a copy of array. thus, change inside func will not affect original array
func test2(x [3]bool) {
	x[0] = false
}
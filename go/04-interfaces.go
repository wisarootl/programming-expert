package main

import "fmt"

type Name interface {
	GetName() string
	// LastName(x string) string
}

type Person struct {
	firstName string
	lastName string
}

func (p Person) GetName() string {
	return p.firstName + " " +  p.lastName
}

type Employee struct {
	name string
}

func (e Employee) GetName() string {
	return e.name
}

func main() {
	//  === Interfaces- ===
	fmt.Println("\n=== Interfaces ===") 
	var name Name = Person{"Time", "Ruscica"} 
	// name = variable name
	// Name = interface name

	fmt.Println(name)
	fmt.Println(name.GetName())

	names := []Name{Employee{"Tim"}, Person{"John", "Doe"}}
	for _, name := range names {
		// all class in Name interface will have GetName() method
		fmt.Println(name.GetName())
	}
	fmt.Println(names)
}

// array of class of the same interface
type Addable interface {
	getNumericValue() float64
}

func AddAll(items []Addable) float64 {
	sum := float64(0)

	for _, item := range items {
		sum += item.getNumericValue()
	}

	return sum
}
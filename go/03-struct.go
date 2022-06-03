package main

import "fmt"

// name of function / class
// start name with capital = export (can be used out size the main file or public)
// start name w/o capital = not export (can not be used outside)
type Person struct {
	name string
	age uint
	clothing Clothing // inheritance
}

type Clothing struct {
	showSize uint
	shirtColor string
}

func (c Clothing) GetShoeSize() uint {
	return c.showSize
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) SetName(newName string) {
	p.name = newName
	fmt.Println(p)
}

func (p Person) GetShoeSize() uint {
	return p.clothing.GetShoeSize()
}

type Person2 struct {
	name string
	age uint
}

func (p Person2) Equal(p2 Person2) bool {
	return p.age == p2.age && p.name == p2.name
}

func main() {
	// === Struct ===
	fmt.Println("\n=== Struct ===") // similar to class in python
	p1 := Person{"Tim", 21, Clothing{12, "Blue"}} // equivalent to {age:21, name:"Tim"}
	fmt.Println(p1)

	var p2 Person
	p2.name = "Tim"
	fmt.Println(p2)

	fmt.Println(p1.GetName())
	p1.SetName("Joey")
	fmt.Println(p1.GetName()) /// not modified outside method itself

	p3 := Person{"Tim", 21, Clothing{12, "Blue"}}
	fmt.Println(p3.clothing)

	fmt.Println("\n= inheritance")
	p4 := Person{"Tim", 21, Clothing{12, "Blue"}}
	fmt.Println(p4.clothing.shirtColor)
	shoeSize := p1.clothing.GetShoeSize()
	fmt.Println(shoeSize)
	shoeSize = p1.GetShoeSize()
	fmt.Println(shoeSize)

	slice := []Person2{{"Joe", 10}, {"Joe", 10}}
	fmt.Println(slice[0].Equal(slice[1]))

}
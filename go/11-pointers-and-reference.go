package main

import "fmt"

func change(x *int) {
	*x = 2
}

type Book struct {
	title	string
	id		int
}

func changeTitle(book Book, title string) {
	book.title = title
}

func changeTitlePointer(book *Book, title string) {
	// when use title with struct. Go. will automatically dereference.
	book.title = title
	// or we can explicitly dereference like below syntax
	// (*book).title = title
}

func (b *Book) changeTitleMethod(title string) {
	b.title = title
}

func main() {
	//  === Pointers and Reference ===
	fmt.Println("\n=== Pointers and Reference ===")
	x := 0
	y := &x // y = address of x
	// pointer = reference in go
	fmt.Println(y, x)
	fmt.Printf("%T %T\n", y, x)
	// type of y = * type of x (asterisk type of x)
	*y = 10 // dereference y
	fmt.Println(y, x)

	// pointer and reference are the same in Go
	fmt.Println("\n= function with pointer and dereference")
	y2 := 0
	change(&y2)
	fmt.Println(y2)

	// nest pointer
	fmt.Println("\n= nested pointer")
	x3 := []string{"hello", "world"}
	y3 := &x3
	w3 := &y3
	fmt.Println(x3, y3, w3)
	fmt.Println(x3, *y3, *w3)
	fmt.Println(x3, *y3, **w3)
	fmt.Printf("%T\n", w3)

	// struct with pointer and reference
	fmt.Println("\n= struct with pointer and reference")
	book := Book{"Tae is great!", 1}
	fmt.Println("init book :", book)
	changeTitle(book, "new title")
	fmt.Println("change title w/o pointer : ", book)
	changeTitlePointer(&book, "new title")
	fmt.Println("change title with pointer : ", book)
	book.changeTitleMethod("new title changed by method")
	fmt.Println(book)

	// slice of pointers
	fmt.Println("\n= slice of pointers")
	// books := []*Book{&Book{"a", 1}}
	// equivalent to
	books := []*Book{{"a", 1}}
	fmt.Println(books, books[0], *books[0])

	// int example
	x4 := 0
	y4 := 1
	s4 := []*int{&x4, &y4}

	fmt.Println(s4, *s4[0])

}
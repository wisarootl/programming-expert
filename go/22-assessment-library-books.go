package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	copies int
}
type Library struct {
	books []*Book
}

func (l *Library) findBook(id int) *Book {
	for _, book := range l.books {
		if book.id == id {
			return book
		}
	}
	return nil
}

func (l *Library) CheckoutBook(id int) (*Book, bool) {
	book := l.findBook(id)

	if book == nil || book.copies < 1 {
		return nil, false
	}

	book.copies -= 1
	return book, true
}

func (l *Library) ReturnBook(id int) bool {
	book := l.findBook(id)

	if book != nil {
		book.copies += 1
		return true
	}

	return false
}

func main() {
	books := []*Book{&Book{1, "Programming For Dummies", "Jason Bourne", 2}, &Book{2, "Routines Explained", "Billy Cyrus", 1}, &Book{3, "Anger Management", "Micheal Scott", 0}}
	library := Library{books}
	book, ok := library.CheckoutBook(1)
	fmt.Println(books[0], book)
	fmt.Println(true, ok)

	var nilBook *Book = nil
	book2, ok2 := library.CheckoutBook(4)
	fmt.Println(nilBook, book2)
	fmt.Println(false, ok2)

	ok3 := library.ReturnBook(2)
	fmt.Println(true, ok3)
}
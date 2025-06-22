package main

import "fmt"

type Book struct {
	Book_id   int
	Book_Name string
}

type Member struct {
	Member_id      int
	Borrowed_Books int
}

type LMS struct {
	Books   map[int]Book
	Members map[int]Member
}

func (lms *LMS) AddBook(book Book) {
	lms.Books[book.Book_id] = book
}

func main() {

	lms := LMS{
		Books:   make(map[int]Book),
		Members: make(map[int]Member),
	}
	book := Book{Book_id: 1, Book_Name: "Introduction to Go Programming"}

	lms.AddBook(book)
	fmt.Printf("%+v\n", lms.Books[1])
}

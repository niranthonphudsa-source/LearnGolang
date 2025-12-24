package niran

import "fmt"

// "fmt"

type Books struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Books

func BooksData() []Books {
	books = append(books, Books{ID: 1, Title: "2003", Author: "Niran"})
	fmt.Println(books)
	return books
}

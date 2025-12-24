package niran

// "fmt"

type Books struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Books

func CreateBook(b Books) []Books {
	books = append(books, Books{ID: b.ID, Title: b.Title, Author: b.Author})
	return books
}

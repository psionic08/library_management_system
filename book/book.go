package book

import "fmt"

type Book struct {
	Shelf    int32
	Title    string
	Author   string
	Id       string
	RenterId string
}

var books []Book

func AddBook(shelf int32, title string, author string, id string, renterId string) error {
	existingBook := getBookById(books, id)

	if existingBook != nil {
		return fmt.Errorf("book with id: %s already exists", id)
	}

	newBook := &Book{
		Shelf:    shelf,
		Title:    title,
		Author:   author,
		Id:       id,
		RenterId: renterId,
	}

	books = append(books, *newBook)
	return nil
}

func GetBooks() []Book {
	output := make([]Book, len(books))
	copy(output, books)
	return output
}

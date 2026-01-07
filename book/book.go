package book

import (
	"fmt"
	"library_management_system/user"
)

type Book struct {
	Shelf    int32
	Title    string
	Author   string
	Id       string
	RenterId string
}

var books []Book

func AddBook(shelf int32, title string, author string, id string) error {
	existingBook := getBookById(books, id)

	if existingBook != nil {
		return fmt.Errorf("book with id: %s already exists", id)
	}

	newBook := &Book{
		Shelf:    shelf,
		Title:    title,
		Author:   author,
		Id:       id,
		RenterId: "",
	}

	books = append(books, *newBook)
	return nil
}

func GetBooks() []Book {
	output := make([]Book, len(books))
	copy(output, books)
	return output
}

func LocateBook(id string) (*Book, error) {
	findBook := getBookById(books, id)
	if findBook == nil {
		return nil, fmt.Errorf("there is no book with id: %s", id)
	}
	clone := *findBook
	return &clone, nil
}

func RemoveBook(id string) error {
	for i, b := range books {
		if b.Id == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book with id: %s does not exist", id)
}

func RentBook(bookId string, userId string) error {
	existingUser, err := user.GetUserById(userId)
	if err != nil {
		return err
	}
	existingBook := getBookById(books, bookId)
	if existingBook == nil {
		return fmt.Errorf("book with id: %s does not exist", bookId)
	}
	if existingBook.RenterId != "" {
		return fmt.Errorf("the book has already been rented by User: %s", existingBook.RenterId)
	}
	existingBook.RenterId = existingUser.Id
	return nil
}

func ReturnBook(bookId string, userId string) error {
	existingUser, err := user.GetUserById(userId)
	if err != nil {
		return err
	}
	existingBook := getBookById(books, bookId)
	if existingBook == nil {
		return fmt.Errorf("book with id: %s does not exist", bookId)
	}
	if existingBook.RenterId == "" {
		return fmt.Errorf("the book was never rented out")
	}
	if existingBook.RenterId != userId {
		return fmt.Errorf("book: %s was not rented out to user: %s", bookId, existingUser.Id)
	}
	existingBook.RenterId = ""
	return nil
}

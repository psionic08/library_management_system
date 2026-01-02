package book

type Book struct {
	Shelf    int32
	Title    string
	Author   string
	Id       string
	RenterId string
}

var books []Book

func AddBook(shelf int32, title string, author string, id string, renterId string) {
	newBook := &Book{
		Shelf:    shelf,
		Title:    title,
		Author:   author,
		Id:       id,
		RenterId: renterId,
	}

	books = append(books, *newBook)
}

func GetBooks() []Book {
	return books
}

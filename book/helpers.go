package book

func getBookById(books []Book, id string) *Book {
	for i := range books {
		if books[i].Id == id {
			return &books[i]
		}
	}
	return nil
}

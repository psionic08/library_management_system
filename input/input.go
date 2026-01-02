package input

import (
	"bufio"
	"fmt"
	"library_management_system/book"
	"os"
)

func addBook(scanner *bufio.Scanner) {
	shelf := readInt32(scanner)
	title := readLine(scanner)
	author := readLine(scanner)
	id := readLine(scanner)
	renterId := readLine(scanner)

	book.AddBook(shelf, title, author, id, renterId)
}

func listBooks(scanner *bufio.Scanner) {
	books := book.GetBooks()

	fmt.Println("INDEX SHELF TITLE AUTHOR ID RENTERID")
	for i, bookItem := range books {
		fmt.Println(i+1, bookItem.Shelf, bookItem.Title, bookItem.Author, bookItem.Id, bookItem.RenterId)
	}
}

func ProgramInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Exit")
		fmt.Println("2. Add book")
		fmt.Println("3. List books")
		fmt.Print("Please enter your input: ")

		var input = readInt32(scanner)

		switch input {
		case 1:
			return
		case 2:
			addBook(scanner)
		case 3:
			listBooks(scanner)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

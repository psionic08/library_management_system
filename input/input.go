package input

import (
	"bufio"
	"fmt"
	"library_management_system/book"
	"library_management_system/user"
	"os"
)

func addBook(scanner *bufio.Scanner) {
	fmt.Print("Please enter the shelf number: ")
	shelf := readInt32(scanner)
	fmt.Print("Please enter the book title: ")
	title := readLine(scanner)
	fmt.Print("Please enter the author name: ")
	author := readLine(scanner)
	fmt.Print("Please enter the book id: ")
	id := readLine(scanner)
	fmt.Print("Please enter the id of renter: ")
	renterId := readLine(scanner)

	err := book.AddBook(shelf, title, author, id, renterId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Your book has been added successfully")
}

func listBooks() {
	books := book.GetBooks()

	fmt.Println("INDEX SHELF TITLE AUTHOR ID RENTERID")
	for i, bookItem := range books {
		fmt.Println(i+1, bookItem.Shelf, bookItem.Title, bookItem.Author, bookItem.Id, bookItem.RenterId)
	}
}

func addUser(scanner *bufio.Scanner) {
	fmt.Print("Please enter the user id: ")
	id := readLine(scanner)
	fmt.Print("Please enter the user name: ")
	name := readLine(scanner)

	err := user.AddUser(id, name)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User has been added successfully")
}

func listUsers() {
	users := user.GetUsers()

	fmt.Println("INDEX ID NAME")
	for i, currentUser := range users {
		fmt.Println(i+1, currentUser.Id, currentUser.Name)
	}
}

func ProgramInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Exit")
		fmt.Println("2. Add book")
		fmt.Println("3. List books")
		fmt.Println("4. Add user")
		fmt.Println("5. List users")
		fmt.Print("Please enter your input: ")

		var input = readInt32(scanner)

		switch input {
		case 1:
			return
		case 2:
			addBook(scanner)
		case 3:
			listBooks()
		case 4:
			addUser(scanner)
		case 5:
			listUsers()
		default:
			fmt.Println("Invalid choice")
		}
	}
}

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

	err := book.AddBook(shelf, title, author, id)
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

func locateBook(scanner *bufio.Scanner) {
	fmt.Printf("Please enter the book id to find: ")
	id := readLine(scanner)

	bookPointer, err := book.LocateBook(id)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Book with id: %s is located in shelf: %d \n", bookPointer.Id, bookPointer.Shelf)
}

func rentBook(scanner *bufio.Scanner) {
	fmt.Printf("Please enter the book id to be rented: ")
	bookId := readLine(scanner)
	fmt.Printf("Please enter the user id to rent the book to: ")
	userId := readLine(scanner)

	err := book.RentBook(bookId, userId)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Book: %s has been rented to user: %s \n", bookId, userId)
}

func returnBook(scanner *bufio.Scanner) {
	fmt.Printf("Please enter the book id to be renturned: ")
	bookId := readLine(scanner)
	fmt.Printf("Please enter the user id who is returning the book: ")
	userId := readLine(scanner)

	err := book.ReturnBook(bookId, userId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Book: %s has been returned by user: %s \n", bookId, userId)
}

func removeBook(scanner *bufio.Scanner) {
	fmt.Printf("Please enter the book id to be removed: ")
	id := readLine(scanner)

	err := book.RemoveBook(id)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The book has been removed")
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

func removeUser(scanner *bufio.Scanner) {
	fmt.Printf("Please enter the book id to be removed:")
	id := readLine(scanner)

	err := user.RemoveUser(id)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The user has been removed")
}

func ProgramInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Exit")
		fmt.Println("2. Add book")
		fmt.Println("3. Remove book")
		fmt.Println("4. List books")
		fmt.Println("5. Locate Book")
		fmt.Println("6, Rent book")
		fmt.Println("7. Return book")
		fmt.Println("8. Add user")
		fmt.Println("9. Remove user")
		fmt.Println("10. List users")
		fmt.Print("Please enter your input: ")

		var input = readInt32(scanner)

		switch input {
		case 1:
			return
		case 2:
			addBook(scanner)
		case 3:
			removeBook(scanner)
		case 4:
			listBooks()
		case 5:
			locateBook(scanner)
		case 6:
			rentBook(scanner)
		case 7:
			returnBook(scanner)
		case 8:
			addUser(scanner)
		case 9:
			removeUser(scanner)
		case 10:
			listUsers()
		default:
			fmt.Println("Invalid choice")
		}
	}
}

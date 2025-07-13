package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/task3/library_management/controllers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		option := strings.TrimSpace(input)

		switch option {
		case "1":
			controllers.AddBook(reader)
		case "2":
			controllers.RemoveBook(reader)
		case "3":
			controllers.BorrowBook(reader)
		case "4":
			controllers.ReturnBook(reader)
		case "5":
			controllers.ListAvailableBooks()
		case "6":
			controllers.ListBorrowedBooks(reader)
		case "0":
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
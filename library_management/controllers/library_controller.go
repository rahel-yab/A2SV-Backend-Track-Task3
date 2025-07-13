package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"example.com/task3/library_management/models"
	"example.com/task3/library_management/services"
)

func AddBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	id, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Enter Book Title: ")
	title := readLine(reader)
	fmt.Print("Enter Book Author: ")
	author := readLine(reader)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	services.AddBook(book)
	fmt.Println("Book added successfully.")
}

func RemoveBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to remove: ")
	id, _ := strconv.Atoi(readLine(reader))
	services.RemoveBook(id)
	fmt.Println("Book removed if it existed.")
}

func BorrowBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to borrow: ")
	bookID, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Enter Member ID: ")
	memberID, _ := strconv.Atoi(readLine(reader))
	err := services.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func ReturnBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to return: ")
	bookID, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Enter Member ID: ")
	memberID, _ := strconv.Atoi(readLine(reader))
	err := services.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func ListAvailableBooks() {
	books := services.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}
	fmt.Println("Available Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func ListBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	memberID, _ := strconv.Atoi(readLine(reader))
	books := services.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books for this member.")
		return
	}
	fmt.Println("Borrowed Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
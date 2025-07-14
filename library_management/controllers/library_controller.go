package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"example.com/task3/models"
	"example.com/task3/services"
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
	services.LibraryInstance.AddBook(book)
	fmt.Println("Book added successfully.")
}

func RemoveBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to remove: ")
	id, _ := strconv.Atoi(readLine(reader))
	err := services.LibraryInstance.RemoveBook(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book removed successfully.")
	}
}

func BorrowBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to borrow: ")
	bookID, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Enter Member ID: ")
	memberID, _ := strconv.Atoi(readLine(reader))
	err := services.LibraryInstance.BorrowBook(bookID, memberID)
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
	err := services.LibraryInstance.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func ListAvailableBooks() {
	books := services.LibraryInstance.ListAvailableBooks()
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
	books := services.LibraryInstance.ListBorrowedBooks(memberID)
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
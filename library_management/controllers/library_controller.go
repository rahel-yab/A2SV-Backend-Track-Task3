package main


type LibraryManager interface {
	AddBook(book models.book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.book
	ListBorrowedBooks(memberID int) []models.book
}
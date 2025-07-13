package models

type member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}
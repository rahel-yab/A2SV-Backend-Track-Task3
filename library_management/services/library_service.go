package services

import (
	"errors"

	"example.com/task3/library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

var books []models.Book
var members []models.Member

func AddBook(book models.Book) {
	book.Status = "Available"
	books = append(books, book)
}

func RemoveBook(bookID int) {
	for i, b := range books {
		if b.ID == bookID {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
}

func BorrowBook(bookID int, memberID int) error {
	for i, b := range books {
		if b.ID == bookID {
			if b.Status == "Borrowed" {
				return errors.New("book is already borrowed")
			}
			books[i].Status = "Borrowed"
			for j, m := range members {
				if m.ID == memberID {
					members[j].BorrowedBooks = append(members[j].BorrowedBooks, b)
					return nil
				}
			}
			newMember := models.Member{
				ID:           memberID,
				Name:         "Member",
				BorrowedBooks: []models.Book{b},
			}
			members = append(members, newMember)
			return nil
		}
	}
	return errors.New("book not found")
}

func ReturnBook(bookID int, memberID int) error {
	for i, b := range books {
		if b.ID == bookID {
			if b.Status != "Borrowed" {
				return errors.New("book is not currently borrowed")
			}
			books[i].Status = "Available"
			for j, m := range members {
				if m.ID == memberID {
					for k, bk := range m.BorrowedBooks {
						if bk.ID == bookID {
							members[j].BorrowedBooks = append(m.BorrowedBooks[:k], m.BorrowedBooks[k+1:]...)
							return nil
						}
					}
				}
			}
		}
	}
	return errors.New("book not found or member mismatch")
}

func ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, b := range books {
		if b.Status == "Available" {
			available = append(available, b)
		}
	}
	return available
}

func ListBorrowedBooks(memberID int) []models.Book {
	for _, m := range members {
		if m.ID == memberID {
			return m.BorrowedBooks
		}
	}
	return nil
}
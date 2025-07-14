package services

import (
	"errors"

	"example.com/task3/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

var library = Library{
	books:   make(map[int]models.Book),
	members: make(map[int]models.Member),
}

var LibraryInstance = &library

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}
	book.Status = "Borrowed"
	l.books[bookID] = book

	member, ok := l.members[memberID]
	if !ok {
		member = models.Member{
			ID:           memberID,
			Name:         "Member",
			BorrowedBooks: []models.Book{},
		}
	}
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status != "Borrowed" {
		return errors.New("book is not currently borrowed")
	}
	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	found := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return errors.New("book not borrowed by this member")
	}
	book.Status = "Available"
	l.books[bookID] = book
	l.members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, b := range l.books {
		if b.Status == "Available" {
			available = append(available, b)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.members[memberID]
	if !ok {
		return nil
	}
	return member.BorrowedBooks
}
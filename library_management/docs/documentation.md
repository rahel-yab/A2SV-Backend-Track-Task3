# Console-Based Library Management System

A simple, interactive command-line application for managing a library's books and members. This system allows users to add, remove, borrow, and return books, as well as view available and borrowed books by member.

## Features

- **Add Books:** Register new books with unique IDs, titles, and authors.
- **Remove Books:** Delete books from the library by their ID.
- **Borrow Books:** Members can borrow available books. The system tracks which member has borrowed which book.
- **Return Books:** Members can return borrowed books, making them available for others.
- **List Available Books:** View all books currently available for borrowing.
- **List Borrowed Books:** View all books borrowed by a specific member.

## Data Models

### Book

- `ID` (int): Unique identifier for the book.
- `Title` (string): Title of the book.
- `Author` (string): Author of the book.
- `Status` (string): Current status ("Available" or "Borrowed").

### Member

- `ID` (int): Unique identifier for the member.
- `Name` (string): Name of the member.
- `BorrowedBooks` ([]Book): List of books currently borrowed by the member.

## Folder Structure

```
library_management/
├── main.go                    # Entry point of the application.
├── controllers/
│   └── library_controller.go  # Handles console input and invokes the appropriate service methods.
├── models/
│   └── book.go                # Defines the Book struct.
│   └── member.go              # Defines the Member struct.
├── services/
│   └── library_service.go     # Contains business logic and data manipulation functions.
├── docs/
│   └── documentation.md       # Contains system documentation
└── go.mod                     # Defines the module and its dependencies.
```

## How to Run

1. Open a terminal and navigate to the `library_management` directory.
2. Run the application:
   ```sh
   go run main.go
   ```
3. Follow the on-screen prompts to interact with the system. Enter the number corresponding to the desired action and provide any requested information (e.g., book ID, member ID).

## Example Usage

- **Add a Book:**
  - Enter the book's ID, title, and author when prompted.
- **Borrow a Book:**
  - Enter the book ID and your member ID. If the member does not exist, a new member will be created.
- **Return a Book:**
  - Enter the book ID and your member ID to return a borrowed book.
- **List Available Books:**
  - View all books that are not currently borrowed.
- **List Borrowed Books:**
  - Enter your member ID to see all books you have borrowed.

## Technical Overview

- The application is organized using a simple MVC-like structure:
  - **Controllers:** Handle user input and output.
  - **Services:** Contain business logic and manage data.
  - **Models:** Define the data structures for books and members.

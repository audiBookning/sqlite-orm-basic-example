// Package books provides functions for managing books in a SQLite database.
package books

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID           int
	BookName     string
	CleanedTitle string
	URL          string
}

// NewDB initializes a new DB connection and creates the books table if it doesn't exist.
func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	sqlStmt := `
	create table if not exists books (id integer primary key, bookName text, cleanedTitle text, url text UNIQUE, added datetime default current_timestamp, edited datetime);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	return db, nil
}

// InsertBook inserts a new book into the database.
func InsertBook(db *sql.DB, book Book) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	stmt, err := tx.Prepare("insert into books(bookName, cleanedTitle, url) values(?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.BookName, book.CleanedTitle, book.URL)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

// UpdateBook updates an existing book in the database.
func UpdateBook(db *sql.DB, id int, book Book) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	stmt, err := tx.Prepare("update books set bookName = ?, cleanedTitle = ?, url = ?, edited = current_timestamp where id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.BookName, book.CleanedTitle, book.URL, id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

// DeleteBook deletes a book from the database.
func DeleteBook(db *sql.DB, id int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	stmt, err := tx.Prepare("delete from books where id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

// GetAllBooks retrieves all books from the database.
func GetAllBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query("select id, bookName, cleanedTitle, url from books")
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %v", err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.BookName, &book.CleanedTitle, &book.URL)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during iteration: %v", err)
	}
	return books, nil
}

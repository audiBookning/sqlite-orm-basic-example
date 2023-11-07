package books

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID           int
	BookName     string
	CleanedTitle string
	URL          string
}

func NewDB(dataSourceName string) (*sql.DB, error) {
	dir := filepath.Dir(dataSourceName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create directory: %v", err)
		}
	}

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

func GetAllBooks(db *sql.DB, page int, pageSize int) ([]Book, error) {
	offset := (page - 1) * pageSize
	rows, err := db.Query("SELECT id, bookName, cleanedTitle, url FROM books LIMIT ? OFFSET ?", pageSize, offset)
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

func GetFilteredTitles(db *sql.DB, page int, pageSize int, filter string) ([]Book, error) {
	offset := (page - 1) * pageSize
	filter = strings.ToLower(filter)
	rows, err := db.Query("SELECT id, bookName, cleanedTitle, url FROM books WHERE LOWER(cleanedTitle) LIKE ? LIMIT ? OFFSET ?", "%"+filter+"%", pageSize, offset)
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

func GetBooksByBookName(db *sql.DB, page int, pageSize int, filter string) ([]Book, error) {
	offset := (page - 1) * pageSize
	filter = strings.ToLower(filter)
	rows, err := db.Query("SELECT id, bookName, cleanedTitle, url FROM books WHERE LOWER(bookName) LIKE ? LIMIT ? OFFSET ?", "%"+filter+"%", pageSize, offset)
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

func GetUniqueBookNames(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT DISTINCT bookName FROM books")
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %v", err)
	}
	defer rows.Close()

	var bookNames []string
	for rows.Next() {
		var bookName string
		err := rows.Scan(&bookName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		bookNames = append(bookNames, bookName)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during iteration: %v", err)
	}
	return bookNames, nil
}

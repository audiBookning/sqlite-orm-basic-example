package main

import (
	"fmt"
	"log"
	"strings"

	"sqlite-basic-example/books"
)

func main() {
	log.Println("Opening database and creating table...")
	db, err := books.NewDB("./db/books.db")

	if err != nil {
		log.Fatalf("Failed to open database or create table: %v", err)
	}
	defer db.Close()
	log.Println("Database opened and table created successfully.")

	log.Println("Inserting data...")
	book := books.Book{
		BookName:     "bookName1",
		CleanedTitle: "cleanedTitle1",
		URL:          "url1",
	}
	err = books.InsertBook(db, book)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println("A book with this URL already exists.")
		} else {
			log.Fatalf("Failed to insert data: %v", err)
		}
	}
	log.Println("Data inserted successfully.")

	log.Println("Printing data...")
	allBooks, err := books.GetAllBooks(db)
	if err != nil {
		log.Fatalf("Failed to get all books: %v", err)
	}
	fmt.Println(allBooks)

	log.Println("Updating data...")
	book.BookName = "updatedBookName1"
	book.CleanedTitle = "updatedCleanedTitle1"
	book.URL = "updatedUrl1"
	err = books.UpdateBook(db, 1, book)
	if err != nil {
		log.Fatalf("Failed to update book: %v", err)
	}
	log.Println("Data updated successfully.")

	log.Println("Printing data...")
	allBooks, err = books.GetAllBooks(db)
	if err != nil {
		log.Fatalf("Failed to get all books: %v", err)
	}
	fmt.Println(allBooks)

	log.Println("Deleting data...")
	err = books.DeleteBook(db, 1)
	if err != nil {
		log.Fatalf("Failed to delete book: %v", err)
	}
	log.Println("Data deleted successfully.")

	log.Println("Printing data...")
	allBooks, err = books.GetAllBooks(db)
	if err != nil {
		log.Fatalf("Failed to get all books: %v", err)
	}
	fmt.Println(allBooks)
}

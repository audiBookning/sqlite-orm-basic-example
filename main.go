package main

import (
	"fmt"
	"log"
	"strings"

	"sqlite-basic-example/books"
)

func main() {
	// **************
	log.Println("Opening database and creating table...")
	db, err := books.NewDB("./db/books.db")

	if err != nil {
		log.Fatalf("Failed to open database or create table: %v", err)
	}
	defer db.Close()
	log.Println("Database opened and table created successfully.")

	// **************
	log.Println("InsertBook - Inserting data...")
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

	// **************
	log.Println("Printing data...")
	allBooks, err := books.GetAllBooks(db, 1, 50)
	if err != nil {
		log.Fatalf("Failed to get all books: %v", err)
	}
	fmt.Println(allBooks)

	log.Println("UpdateBook - Updating data...")
	book.BookName = "updatedBookName1"
	book.CleanedTitle = "updatedCleanedTitle1"
	book.URL = "updatedUrl1"
	err = books.UpdateBook(db, 1, book)
	if err != nil {
		log.Fatalf("Failed to update book: %v", err)
	}
	log.Println("Data updated successfully.")

	// **************
	log.Println("Printing data...")
	allBooks, err = books.GetAllBooks(db, 1, 50)
	if err != nil {
		log.Fatalf("Failed to get all books: %v", err)
	}
	fmt.Println(allBooks)

	// **************
	log.Println("Getting unique book names...")
	uniqueBookNames, err := books.GetUniqueBookNames(db)
	if err != nil {
		log.Fatalf("Failed to get unique book names: %v", err)
	}
	fmt.Println(uniqueBookNames)

	log.Println("Filtering and printing data by book name...")
	filteredBooksByBookName, err := books.GetBooksByBookName(db, 1, 50, "bookName1")
	if err != nil {
		log.Fatalf("Failed to get books by book name: %v", err)
	}
	fmt.Println(filteredBooksByBookName)

	// **************

	log.Println("GetFilteredTitles - Inserting more data...")
	book2 := books.Book{
		BookName:     "bookName2",
		CleanedTitle: "cleanedTitle2",
		URL:          "url2",
	}
	err = books.InsertBook(db, book2)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println("A book with this URL already exists.")
		} else {
			log.Fatalf("Failed to insert data: %v", err)
		}
	}
	log.Println("More data inserted successfully.")

	log.Println("Filtering and printing data...")
	filteredBooks, err := books.GetFilteredTitles(db, 1, 50, "cleanedTitle1")
	if err != nil {
		log.Fatalf("Failed to get filtered books: %v", err)
	}
	fmt.Println(filteredBooks)
	// **************
	log.Println("Deleting data...")
	err = books.DeleteBook(db, 1)
	if err != nil {
		log.Fatalf("Failed to delete book: %v", err)
	}
	log.Println("Data deleted successfully.")

	log.Println("Printing data...")
	allBooks, err = books.GetAllBooks(db, 1, 50)
	if err != nil {
		log.Fatalf("Failed to get all books: %v", err)
	}
	fmt.Println(allBooks)
}

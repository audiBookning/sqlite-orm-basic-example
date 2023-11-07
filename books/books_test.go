package books

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB("./testdb/TestNewDB.db")
	if err != nil {
		t.Errorf("Failed to open database or create table: %v", err)
	}
	defer db.Close()
}

func TestInsertBook(t *testing.T) {
	db, _ := NewDB("./testdb/TestInsertBook.db")
	defer db.Close()

	book := Book{
		BookName:     "Test Book",
		CleanedTitle: "Test Book",
		URL:          "http://test.com",
	}

	err := InsertBook(db, book)
	if err != nil {
		t.Errorf("Failed to insert book: %v", err)
	}
}

func TestUpdateBook(t *testing.T) {
	db, _ := NewDB("./testdb/TestUpdateBook.db")
	defer db.Close()

	book := Book{
		BookName:     "Updated Test Book",
		CleanedTitle: "Updated Test Book",
		URL:          "http://test2.com",
	}

	err := UpdateBook(db, 1, book)
	if err != nil {
		t.Errorf("Failed to update book: %v", err)
	}
}

func TestDeleteBook(t *testing.T) {
	db, _ := NewDB("./testdb/TestDeleteBook.db")
	defer db.Close()

	err := DeleteBook(db, 1)
	if err != nil {
		t.Errorf("Failed to delete book: %v", err)
	}
}

func TestGetAllBooks(t *testing.T) {
	db, _ := NewDB("./testdb/TestGetAllBooks.db")
	defer db.Close()
	book := Book{
		BookName:     "Test Book",
		CleanedTitle: "Test Book",
		URL:          "http://test.com",
	}
	_ = InsertBook(db, book)

	books, err := GetAllBooks(db, 1, 50)
	if err != nil {
		t.Errorf("Failed to get all books: %v", err)
	}

	if len(books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(books))
	}
}

func TestGetFilteredTitles(t *testing.T) {
	db, _ := NewDB("./testdb/TestGetFilteredTitles.db")
	defer db.Close()
	book := Book{
		BookName:     "Test Book",
		CleanedTitle: "Test Book",
		URL:          "http://test.com",
	}
	_ = InsertBook(db, book)

	filteredBooks, err := GetFilteredTitles(db, 1, 50, "Test Book")
	if err != nil {
		t.Errorf("Failed to get filtered books: %v", err)
	}

	if len(filteredBooks) != 1 {
		t.Errorf("Expected 1 book, got %d", len(filteredBooks))
	}
}

func TestGetBooksByBookName(t *testing.T) {
	db, _ := NewDB("./testdb/TestGetBooksByBookName.db")
	defer db.Close()

	book := Book{
		BookName:     "Test Book",
		CleanedTitle: "Test Book",
		URL:          "http://test.com",
	}
	_ = InsertBook(db, book)

	filteredBooksByBookName, err := GetBooksByBookName(db, 1, 50, "Test Book")
	if err != nil {
		t.Errorf("Failed to get books by book name: %v", err)
	}

	if len(filteredBooksByBookName) != 1 {
		t.Errorf("Expected 1 book, got %d", len(filteredBooksByBookName))
	}
}

func TestGetUniqueBookNames(t *testing.T) {
	db, _ := NewDB("./testdb/TestGetUniqueBookNames.db")
	defer db.Close()

	book := Book{
		BookName:     "Test Book",
		CleanedTitle: "Test Book",
		URL:          "http://test.com",
	}
	_ = InsertBook(db, book)

	uniqueBookNames, err := GetUniqueBookNames(db)
	if err != nil {
		t.Errorf("Failed to get unique book names: %v", err)
	}

	if len(uniqueBookNames) != 1 {
		t.Errorf("Expected 1 unique book name, got %d", len(uniqueBookNames))
	}

	expectedBookName := "Test Book"
	if !reflect.DeepEqual(uniqueBookNames[0], expectedBookName) {
		t.Errorf("Expected %s, got %s", expectedBookName, uniqueBookNames[0])
	}
}

func TestMain(m *testing.M) {
	// Setup

	// Call the test functions
	code := m.Run()

	// Cleanup after tests
	os.RemoveAll("./testdb")

	// Exit
	os.Exit(code)
}

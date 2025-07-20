package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"book_api/models"

	"github.com/gorilla/mux"
)

var books []models.Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var updatedBook models.Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = params["id"]
			books = append(books, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	// Sample data
	books = append(books, models.Book{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan", ISBN: "9780134190440"})
	books = append(books, models.Book{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", ISBN: "9780132350884"})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Book API. Visit /api/books"))
	}).Methods("GET")
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

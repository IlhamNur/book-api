package handler

import (
	"book-api/model"
	"book-api/store"
	"book-api/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// GetAllBooks returns all Books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	store := store.GetStore()
	store.Mu.RLock()
	defer store.Mu.RUnlock()

	Books := make([]model.Book, 0, len(store.Books))
	for _, book := range store.Books {
		Books = append(Books, book)
	}

	utils.JSON(w, http.StatusOK, Books)
}

// GetBookByID returns a book by its ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	store := store.GetStore()

	store.Mu.RLock()
	book, ok := store.Books[id]
	store.Mu.RUnlock()

	if !ok {
		utils.Error(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.JSON(w, http.StatusOK, book)
}

// CreateBook creates a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	book.ID = uuid.NewString()

	store := store.GetStore()
	store.Mu.Lock()
	store.Books[book.ID] = book
	store.Mu.Unlock()

	utils.JSON(w, http.StatusCreated, book)
}

// UpdateBook updates a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var update model.Book
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	store := store.GetStore()
	store.Mu.Lock()
	book, ok := store.Books[id]
	if !ok {
		store.Mu.Unlock()
		utils.Error(w, http.StatusNotFound, "Book not found")
		return
	}

	book.Title = update.Title
	book.Author = update.Author
	book.PublishedYear = update.PublishedYear
	store.Books[id] = book
	store.Mu.Unlock()

	utils.JSON(w, http.StatusOK, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	store := store.GetStore()

	store.Mu.Lock()
	_, ok := store.Books[id]
	if !ok {
		store.Mu.Unlock()
		utils.Error(w, http.StatusNotFound, "Book not found")
		return
	}
	delete(store.Books, id)
	store.Mu.Unlock()

	utils.JSON(w, http.StatusOK, map[string]string{"message": "Book deleted"})
}

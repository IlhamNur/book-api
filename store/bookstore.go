package store

import (
	"book-api/model"
	"sync"
)

// BookStore is a singleton storage for books in memory
type BookStore struct {
	Books map[string]model.Book
	Mu    sync.RWMutex
}

var instance *BookStore
var once sync.Once

// GetStore returns the singleton instance of BookStore
func GetStore() *BookStore {
	once.Do(func() {
		instance = &BookStore{
			Books: make(map[string]model.Book),
		}
	})
	return instance
}

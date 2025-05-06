package router

import (
	"book-api/handler"
	"book-api/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.LoggerMiddleware)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetAllBooks)
		r.Post("/", handler.CreateBook)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetBookByID)
			r.Put("/", handler.UpdateBook)
			r.Delete("/", handler.DeleteBook)
		})
	})

	return r
}

package app

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lulzshadowwalker/pararum/handler"
)

func initRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello, client!\n")
	})

	router.Route("/notes", generateNotesRoutes)
	return router
}

func generateNotesRoutes(r chi.Router) {
	n := handler.Note{}

	r.Get("/", n.Index)
	r.Post("/", n.Create)
	r.Get("/{id}", n.Show)
	r.Put("/{id}", n.Update)
	r.Delete("/{id}", n.Delete)
}

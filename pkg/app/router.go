package app

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lulzshadowwalker/pararum/pkg/handler"
	"github.com/lulzshadowwalker/pararum/pkg/repo"
)

func (a *App) initRouter() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "heyo client\n")
	})

	router.Route("/notes", a.generateNotesRoutes)
	a.router = router
}

func (a *App) generateNotesRoutes(r chi.Router) {
	n := handler.Note{
		Repo: repo.NoteRepo{
			Db: a.db,
		},
	}

	r.Get("/", n.Index)
	r.Post("/", n.Create)
	r.Get("/{id}", n.Show)
	r.Put("/{id}", n.Update)
	r.Delete("/{id}", n.Delete)
}

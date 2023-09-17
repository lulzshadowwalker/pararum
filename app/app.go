package app

import (
	"net/http"
)

type Application struct {
	router http.Handler
}

func New() *Application {
	return &Application{
		router: initRouter(),
	}
}

func (a *Application) Start() error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	return err
}

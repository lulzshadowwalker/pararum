package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{
		router: initRouter(),
	}
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	ch := make(chan error)
	go func() {
		defer close(ch)
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("server shutdown")
			} else {
				ch <- err
			}
		}
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		return server.Shutdown(timeout)
	}
}

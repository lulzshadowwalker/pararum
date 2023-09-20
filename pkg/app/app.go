package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	// "log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	router http.Handler
	db     *sql.DB
}

func New() *App {
	a := &App{
		db: registerDb(),
	}

	a.initRouter()

	return a
}

func registerDb() *sql.DB {
	uname, pwd, dbName := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")

	conStr := fmt.Sprintf("%s:%s@/%s", uname, pwd, dbName)
	_ = conStr
	db, err := sql.Open("mysql", "root:vZnRX$#aBC279@tcp(127.0.0.1:3306)/go_notes")
	if err != nil {
		log.Printf("failed to initialize application %q", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error connecting to database %q", err)
	}

	return db
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to datbase %q", err)
	}
	defer a.db.Close()
	log.Println("database has been initialized")

	ch := make(chan error, 1)
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

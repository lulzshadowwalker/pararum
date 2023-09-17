package main

import (
	"errors"
	"log"
	"net/http"

	application "github.com/lulzshadowwalker/pararum/app"
)

func main() {
	/// - [X] server setup
	/// - [ ] graceful termination
	/// - [ ]  CRUD ops
	///		* request validation
	/// 	* proper responses
	/// - [ ]  auth middleware

	app := application.New()
	if err := app.Start(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server shutdown")
		} else {
			log.Printf("error running the application %q\n", err)
		}
	}
}

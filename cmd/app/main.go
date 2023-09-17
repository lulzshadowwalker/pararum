package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/lulzshadowwalker/pararum/app"
)

func main() {
	/// - [X] server setup
	/// - [X] graceful termination
	/// - [ ]  CRUD ops
	///		* request validation
	/// 	* proper responses
	/// - [ ]  auth middleware

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := app.New()
	if err := app.Start(ctx); err != nil {
		log.Printf("error running the application %q\n", err)
	}
}

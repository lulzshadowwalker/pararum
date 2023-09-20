package main

import (
	"context"
	"os"
	"os/signal"

	_ "github.com/lulzshadowwalker/pararum/internal/config"
	"github.com/lulzshadowwalker/pararum/pkg/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := app.New()
	app.Start(ctx)
}

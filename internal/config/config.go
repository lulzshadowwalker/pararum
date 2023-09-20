package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("error loading application config %q\n", err)
	}
}

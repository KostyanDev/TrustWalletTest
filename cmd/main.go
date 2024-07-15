package main

import (
	"log"
	"os"

	"1/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

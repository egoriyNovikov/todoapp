package main

import (
	"log"

	"github.com/egoriynovikov/todoapp/internal/core/app"
)

func main() {
	application := app.NewApp()

	log.Println("Server starting on :8080")
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}

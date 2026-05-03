package main

import (
	"log"

	core_app "github.com/egoriynovikov/todoapp/internal/core"
)

func main() {
	application := core_app.NewApp()

	log.Println("Server starting on :8080")
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}

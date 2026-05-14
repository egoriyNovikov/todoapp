package main

import (
	"log"

	core_app "github.com/egoriynovikov/todoapp/internal/core"
	core_logger "github.com/egoriynovikov/todoapp/internal/core/logger"
)

func main() {
	logger, err := core_logger.New(false)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	application := core_app.NewApp()
	defer application.Close()

	log.Println("Server starting on :8080")
	if err := application.Run(logger); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/Met4tron/books-go/config"
	"github.com/Met4tron/books-go/pkg/logger"
)

func main() {
	logger.InitLogger()
	_, err := config.LoadConfig()

	if err != nil {
		logger.Error("Error to get env vars", err)
	}
}

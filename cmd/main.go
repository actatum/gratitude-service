package main

import (
	"fmt"
	"os"

	// Imported for swagger documentation
	_ "github.com/actatum/gratitude-board-service/docs"

	"github.com/actatum/gratitude-board-service/pkg/transport"
)

// @title Gratitude API
// @version 1.0
// @description API to handle CRUD operations for user/public messages.

// @host localhost:8080
// @BasePath /transport/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	if err := transport.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/actatum/gratitude-board-service/pkg/api"
)

func main() {
	if err := api.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

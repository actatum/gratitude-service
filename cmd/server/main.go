package main

import (
	"fmt"
	"github.com/actatum/gratitude-board-service/pkg/api/http"
	"os"
)

func main() {
	if err := http.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/LazarenkoA/go-arch-lint/internal/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	return app.Execute()
}

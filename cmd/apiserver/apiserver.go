package main

import (
	"os"

	"github.com/usmhk/goserver/internal/apiserver"
)

func main() {
	command := apiserver.NewAPIServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

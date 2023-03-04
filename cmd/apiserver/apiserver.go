package main

import (
	"os"

	"github.com/usmhk/goserver/internal/apiserver"
)

/**
 *
 * golang Software engineering
 *
 */
func main() {
	command := apiserver.NewAPIServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

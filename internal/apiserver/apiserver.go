package apiserver

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	appName = "APIServer"
)

func NewAPIServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          appName,
		Short:        `A good go practise demo.`,
		Long:         ``,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	return cmd
}

func run() error {
	//
	fmt.Println("hello")
	return nil
}

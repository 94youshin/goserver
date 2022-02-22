package apiserver

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/usmhk/goserver/pkg/version/verflag"
)

const (
	appName           = "APIServer"
	defaultConfigName = ""
)

var cfgFile string

func NewAPIServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   appName,
		Short: `A good go practise demo.`,
		Long: `A goog Go practical project, used to create user with basic information.

Find more apiserver information at:
	https://github.com/usmhk/goserver/blob/main/README.md`,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath, args)
				}
			}
			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test.yaml)")

	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {

	// create the gin engine
	g := gin.New()

	loadRouter(g)

	insecureServer := http.Server{
		Addr:    "0.0.0.0",
		Handler: g,
	}

	go func() {
		fmt.Printf("Start to listening the incoming requests on http address: %s\n", "0.0.0.0:8080")
		if err := insecureServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Join("", ""))
		viper.SetConfigName(defaultConfigName)
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

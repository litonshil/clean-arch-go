package cmd

import (
	"clean-arch/config"
	"clean-arch/infra/conn"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: "clean-arch",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	// load application configuration
	if err := config.Load(); err != nil {
		//log.Error(err)
		os.Exit(1)
	}

	conn.ConnectDB()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "sqs",
	Short: "Sqs is really nice project",
}

var (
	nonInteractiveMode bool
)

func init () {
	rootCmd.PersistentFlags().BoolVarP(&nonInteractiveMode, "non-interactive", "n", false, "Run in non interactive mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

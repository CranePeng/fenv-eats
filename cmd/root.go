package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "Eagle",
	Short:   "Eagle",
	Long:    "Eagle:Elastic Automatic Task System",
	Version: "0.1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "clivd",
		Short: "CLI test by vaclav-dvorak",
		Long:  "best CLI tool as hell by vaclav-dvorak",
	}
)

// Execute function execute cobra command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd())
}

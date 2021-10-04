package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version, commit, date = "dev", "dev", "now"
)

func versionCmd() *cobra.Command {
	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(cmd *cobra.Command, args []string) {
			const fmat = "%-20s %s\n"
			fmt.Printf(fmat, "Version", version)
			fmt.Printf(fmat, "Commit", commit)
			fmt.Printf(fmat, "Date", date)
		},
	}
	return &command
}

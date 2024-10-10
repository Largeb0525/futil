package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version info",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("futil version: %s\n", version)
	},
}

func initVersionCmd() {
	rootCmd.AddCommand(versionCmd)
}

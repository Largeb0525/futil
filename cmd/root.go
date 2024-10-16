package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "futil",
	Short:                 "File Utility",
	Long:                  "futil is a command-line tool for file operations, offering functionality for line counting and file checksum calculation.",
	DisableAutoGenTag:     true,
	DisableFlagsInUseLine: true,
}

func Execute() {
	initVersionCmd()
	initChecksumCmd()
	initLineCountCmd()
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

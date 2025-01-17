package cmd

import (
	"fmt"

	"github.com/Largeb0525/futil/internal"

	"github.com/spf13/cobra"
)

var file string

var lineCountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Long:  `Calculate the number of lines in a file`,
	Run: func(cmd *cobra.Command, args []string) {
		if file == "-" {
			fmt.Println(internal.CountLinesFromStdin())
		} else {
			err := internal.CheckFileValid(file)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				count, err := internal.CountLinesInFile(file)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println(count)
				}
			}
		}
	},
}

func initLineCountCmd() {
	lineCountCmd.Flags().StringVarP(&file, "file", "f", "", "the input file")
	rootCmd.AddCommand(lineCountCmd)
}

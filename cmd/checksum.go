package cmd

import (
	"fmt"
	"os"

	"github.com/Largeb0525/futil/internal"

	"github.com/spf13/cobra"
)

var algorithm string

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Print checksum of file",
	Long:  `Calculate the MD5, SHA1 or SHA256 checksum of a file`,
	Run: func(cmd *cobra.Command, args []string) {
		if file == "-" {
			fmt.Println(internal.ChecksumFromStdin(algorithm))
		} else {
			_, err := os.Stat(file)
			if os.IsNotExist(err) {
				fmt.Printf("error: No such file '%s'\n", file)
				return
			}
			hash, err := internal.CalculateChecksum(file, algorithm)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(hash)
			}
		}
	},
}

func initChecksumCmd() {
	checksumCmd.Flags().StringVarP(&file, "file", "f", "", "the input file")
	checksumCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "md5", "checksum algorithm (md5, sha1, sha256)")
	rootCmd.AddCommand(checksumCmd)
}

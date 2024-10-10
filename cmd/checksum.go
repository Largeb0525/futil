package cmd

import (
	"fmt"

	"github.com/Largeb0525/futil/internal"
	"github.com/spf13/cobra"
)

var useMD5, useSHA1, useSHA256 bool

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Print checksum of file",
	Long:  `Calculate the MD5, SHA1 or SHA256 checksum of a file`,
	Run: func(cmd *cobra.Command, args []string) {
		var algorithm string

		if useMD5 {
			algorithm = "md5"
		} else if useSHA1 {
			algorithm = "sha1"
		} else if useSHA256 {
			algorithm = "sha256"
		} else {
			fmt.Println("Please specify a checksum algorithm: --md5, --sha1, or --sha256")
			return
		}

		if file == "-" {
			fmt.Println(internal.ChecksumFromStdin(algorithm))
		} else {
			err := internal.CheckFileValid(file)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				hash, err := internal.CalculateChecksum(file, algorithm)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println(hash)
				}
			}
		}
	},
}

func initChecksumCmd() {
	checksumCmd.Flags().StringVarP(&file, "file", "f", "", "the input file (use '-' for stdin)")
	checksumCmd.Flags().BoolVar(&useMD5, "md5", false, "use MD5 algorithm")
	checksumCmd.Flags().BoolVar(&useSHA1, "sha1", false, "use SHA1 algorithm")
	checksumCmd.Flags().BoolVar(&useSHA256, "sha256", false, "use SHA256 algorithm")
	rootCmd.AddCommand(checksumCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// readingTimeCmd represents the readingTime command
var readingTimeCmd = &cobra.Command{
	Use:   "reading-time",
	Short: "Assign reading time tags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readingTime called")
	},
}

func init() {
	rootCmd.AddCommand(readingTimeCmd)
}

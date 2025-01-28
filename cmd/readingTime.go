package cmd

import (
	"github.com/kahnwong/wallabag-tagger/core"

	"github.com/spf13/cobra"
)

var readingTimeCmd = &cobra.Command{
	Use:   "reading-time",
	Short: "Assign reading time tags",
	Run: func(cmd *cobra.Command, args []string) {
		core.ReadingTime()
	},
}

func init() {
	core.WallabagInit()
	rootCmd.AddCommand(readingTimeCmd)
}

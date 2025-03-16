package cmd

import (
	"github.com/kahnwong/wallabag-tagger/core"

	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Apply tags via LLM",
	Run: func(cmd *cobra.Command, args []string) {
		core.LLMTags()
	},
}

func init() {
	core.WallabagInit()
	rootCmd.AddCommand(tagsCmd)
}

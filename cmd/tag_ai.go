package cmd

import (
	"github.com/kahnwong/wallabag-tagger/core"

	"github.com/spf13/cobra"
)

var tagAiCmd = &cobra.Command{
	Use:   "tag-ai",
	Short: "Tag articles for AI usage in the writing process",
	Run: func(cmd *cobra.Command, args []string) {
		core.AiTag()
	},
}

func init() {
	rootCmd.AddCommand(tagAiCmd)
}

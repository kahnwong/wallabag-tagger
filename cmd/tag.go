package cmd

import (
	"github.com/kahnwong/wallabag-tagger/core"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Apply tags via LLM",
	Run: func(cmd *cobra.Command, args []string) {
		core.LlmTag()
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}

package cmd

import (
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/Strubbl/wallabago/v9"
	"github.com/kahnwong/wallabag-tagger/core"

	"github.com/spf13/cobra"
)

type Tags struct {
	Tag []string `json:"tags"`
}

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Apply tags via LLM",
	Run: func(cmd *cobra.Command, args []string) {
		entries := core.WallabagGetEntries()
		for _, entry := range entries.Embedded.Items {
			if len(entry.Tags) > 1 || len(entry.Tags) == 0 {
				log.Info().Msgf("Skipping article: %s", entry.Title)
				continue
			}
			log.Info().Msgf("Processing article: %s", entry.Title)

			// get tags from llm
			tagsStr := core.GeminiGetTags(entry.Content)

			// convert json-string to Tags struct
			var tags Tags
			err := json.Unmarshal([]byte(tagsStr), &tags)
			if err != nil {
				log.Error().Msgf("Cannot unmarshal tags: %s", tagsStr)
			}

			// add tags prefix so it doesn't conflict with manually-assigned tags
			var tagsWithPrefix []string
			for _, tag := range tags.Tag {
				tagsWithPrefix = append(tagsWithPrefix, "llm-"+tag)
			}

			// update entry tags
			err = wallabago.AddEntryTags(entry.ID, tagsWithPrefix...)
			if err != nil {
				log.Error().Msgf("Cannot assign tags to article: %s", entry.Title)
			}
		}
	},
}

func init() {
	core.WallabagInit()
	rootCmd.AddCommand(tagsCmd)
}

package core

import (
	"encoding/json"
	"strings"

	"github.com/Strubbl/wallabago/v9"
	"github.com/rs/zerolog/log"
)

type Tags struct {
	Tag []string `json:"tags"`
}

func LLMTags() {
	entries := WallabagGetEntries()
	for _, entry := range entries.Embedded.Items {
		// skip if already tagged via LLM
		isSkip := false
		if len(entry.Tags) >= 1 { // if already has tags
			for _, tag := range entry.Tags {
				if strings.HasPrefix(tag.Label, "llm") {
					isSkip = true
					continue
				}
			}
			if isSkip {
				log.Info().Msgf("Skipping article: %s", entry.Title)
				continue
			}
		}
		log.Info().Msgf("Processing article: %s", entry.Title)

		// get tags from llm
		tagsStr, err := GeminiGetTags(entry.Content)

		if err == nil { // if successfully generated tags
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
	}
}

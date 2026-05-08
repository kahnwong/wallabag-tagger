package core

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/Strubbl/wallabago/v9"
)

type Tags struct {
	Tag []string `json:"tags"`
}

func isSkipEntry(entry wallabago.Item) bool {
	isSkip := false
	if len(entry.Tags) >= 1 { // if already has tags
		for _, tag := range entry.Tags {
			if strings.HasPrefix(tag.Label, "llm") {
				isSkip = true
				continue
			}
		}
	}

	return isSkip
}

func LLMTags() {
	entries := WallabagGetEntries()
	for _, entry := range entries.Embedded.Items {
		// skip if already tagged via LLM
		isSkip := isSkipEntry(entry)
		if isSkip {
			slog.Info("Skipping article", "title", entry.Title)
			continue
		}
		slog.Info("Processing article", "title", entry.Title)

		// get tags from llm
		tagsStr, err := GeminiGetTags(entry.Content)

		if err == nil { // if successfully generated tags
			// convert json-string to Tags struct
			var tags Tags
			err := json.Unmarshal([]byte(tagsStr), &tags)
			if err != nil {
				slog.Error("Cannot unmarshal tags", "tags", tagsStr)
			}

			// add tags prefix so it doesn't conflict with manually-assigned tags
			var tagsWithPrefix []string
			for _, tag := range tags.Tag {
				tagsWithPrefix = append(tagsWithPrefix, "llm-"+tag)
			}

			// update entry tags
			WallabagWriteTags(entry, tagsWithPrefix)
		}
	}
}

package core

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/microcosm-cc/bluemonday"
)

var topicTagPrefix = "llmtag-"

type Tags struct {
	Tags []string `json:"tags"`
}

func LlmTag() {
	entries := wallabag.GetEntries()
	for _, entry := range entries.Embedded.Items {
		// skip if already tagged via LLM
		isSkip := isSkipTagging(topicTagPrefix, entry)
		if isSkip {
			slog.Info("Skipping article", "title", entry.Title)
			continue
		}
		slog.Info("Processing article", "title", entry.Title)

		// prep prompt
		p := bluemonday.StripTagsPolicy()
		contentSanitized := p.Sanitize(
			entry.Content,
		)
		prompt := renderPrompt("resources/topic.txt", map[string]interface{}{
			"Content": contentSanitized,
		})

		// get tags from llm
		tagsStr, err := FetchLlmResponse(prompt)

		if err == nil { // if successfully generated tags
			// convert json-string to Tags struct
			var llmTags Tags
			err = json.Unmarshal([]byte(tagsStr), &llmTags)
			if err != nil {
				slog.Error("Cannot unmarshal tags", "tags", tagsStr)
			}

			// filter junk tags
			tags := filterTags(llmTags.Tags, []string{"data", "devops", "tools", "security", "software engineering", "web development", "leadership", "misc"})

			// add tag prefix so it doesn't conflict with manually-assigned tags
			var tagsWithPrefix []string
			for _, tag := range tags {
				tagsWithPrefix = append(tagsWithPrefix, fmt.Sprintf("%s%s", topicTagPrefix, tag))
			}

			// update entry tags
			wallabag.WriteTags(entry, tagsWithPrefix)
		}
	}
}

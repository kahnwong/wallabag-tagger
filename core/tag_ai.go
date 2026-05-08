package core

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

var aiTagPrefix = "llmai-"

func AiTag() {
	entries := wallabag.GetEntries()
	for _, entry := range entries.Embedded.Items {
		// skip if already tagged
		isSkip := isSkipTagging(aiTagPrefix, entry)
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
		prompt := renderPrompt("resources/ai.txt", map[string]interface{}{
			"Content": contentSanitized,
		})

		// get tag from llm
		isAiRaw, err := FetchLlmResponse(prompt)
		if err == nil { // if successfully obtained response
			// cleanup
			var isAi string
			isAiRaw = strings.ToLower(isAiRaw)
			if strings.HasPrefix(isAiRaw, "yes") {
				isAi = "yes"
			} else if strings.HasPrefix(isAiRaw, "no") {
				isAi = "no"
			}

			// write tag
			if isAi != "" {
				wallabag.WriteTags(entry, []string{fmt.Sprintf("%s%s", aiTagPrefix, isAi)})
			}
		}
	}
}

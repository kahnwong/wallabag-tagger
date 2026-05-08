package core

import (
	"strings"

	"github.com/Strubbl/wallabago/v9"
)

func isSkipTagging(tagPrefix string, entry wallabago.Item) bool {
	isSkip := false
	if len(entry.Tags) >= 1 { // if already has tags
		for _, tag := range entry.Tags {
			if strings.HasPrefix(tag.Label, tagPrefix) {
				isSkip = true
				continue
			}
		}
	}

	return isSkip
}

func filterTags(tags, possibleTags []string) []string {
	isPossible := make(map[string]bool)
	for _, v := range possibleTags {
		isPossible[v] = true
	}

	var filtered []string
	for _, r := range tags {
		if isPossible[r] {
			filtered = append(filtered, r)
		}
	}

	return filtered
}

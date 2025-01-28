package cmd

import (
	"fmt"
	"log"

	"github.com/Strubbl/wallabago/v9"
	"github.com/kahnwong/wallabag-tagger/core"

	"github.com/spf13/cobra"
)

func timeBinning(readingTime int) string {
	readingTimeTag := "readingTime-"
	switch {
	case readingTime <= 3:
		readingTimeTag += "3"
	case readingTime <= 5:
		readingTimeTag += "5"
	case readingTime <= 10:
		readingTimeTag += "10"
	case readingTime <= 15:
		readingTimeTag += "15"
	case readingTime <= 20:
		readingTimeTag += "20"
	case readingTime <= 25:
		readingTimeTag += "25"
	case readingTime <= 30:
		readingTimeTag += "30"
	default:
		readingTimeTag = ""
	}

	return readingTimeTag
}

var readingTimeCmd = &cobra.Command{
	Use:   "reading-time",
	Short: "Assign reading time tags",
	Run: func(cmd *cobra.Command, args []string) {
		// get entries
		entries := core.WallabagGetEntries(200)
		for _, entry := range entries.Embedded.Items {
			fmt.Printf("Processing article: %s\n", entry.Title)

			// assign reading time tags
			readingTimeTag := timeBinning(entry.ReadingTime)
			err := wallabago.AddEntryTags(entry.ID, readingTimeTag)
			if err != nil {
				log.Printf("Cannot assign tags to article: %s", entry.Title)
			}
		}
	},
}

func init() {
	core.WallabagInit()
	rootCmd.AddCommand(readingTimeCmd)
}

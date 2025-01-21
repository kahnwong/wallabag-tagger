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
	if readingTime <= 3 {
		readingTimeTag += "3"
	} else if readingTime <= 5 {
		readingTimeTag += "5"
	} else if readingTime <= 10 {
		readingTimeTag += "10"
	} else {
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

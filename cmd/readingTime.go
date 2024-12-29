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
	if readingTime <= 5 {
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
		entries, err := wallabago.GetEntries(
			wallabago.APICall,
			0, 0, "", "", 1, 200, "", 0, -1, "", "")
		if err != nil {
			log.Println("Cannot obtain articles from Wallabag")
		}

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

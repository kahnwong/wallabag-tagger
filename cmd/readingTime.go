package cmd

import (
	"fmt"
	"log"
	"sync"

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
		entries := core.WallabagGetEntries(400)

		// goroutines
		var wg sync.WaitGroup
		wg.Add(len(entries.Embedded.Items))
		for _, entry := range entries.Embedded.Items {
			go func() {
				fmt.Printf("Processing article: %s\n", entry.Title)

				// assign reading time tags
				readingTimeTag := timeBinning(entry.ReadingTime)
				err := wallabago.AddEntryTags(entry.ID, readingTimeTag)
				if err != nil {
					log.Printf("Cannot assign tags to article: %s", entry.Title)
				}

				wg.Done()
			}()
		}
		wg.Wait()
	},
}

func init() {
	core.WallabagInit()
	rootCmd.AddCommand(readingTimeCmd)
}

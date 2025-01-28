package core

import (
	"sync"

	"github.com/Strubbl/wallabago/v9"
	"github.com/rs/zerolog/log"
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

func ReadingTime() {
	// get entries
	entries := WallabagGetEntries(5)

	// goroutines
	var wg sync.WaitGroup
	wg.Add(len(entries.Embedded.Items))
	for _, entry := range entries.Embedded.Items {
		go func() {
			log.Info().Msgf("Processing article: %s", entry.Title)

			// assign reading time tag
			readingTimeTag := timeBinning(entry.ReadingTime)
			err := wallabago.AddEntryTags(entry.ID, readingTimeTag)
			if err != nil {
				log.Err(err).Msgf("Cannot assign tags to article: %s", entry.Title)
			}

			wg.Done()
		}()
	}
	wg.Wait()

}

package core

import (
	"log/slog"
	"os"

	"github.com/Strubbl/wallabago/v9"
	_ "github.com/joho/godotenv/autoload"
)

func WallabagInit() {
	wallabagConfig := wallabago.WallabagConfig{
		WallabagURL:  config.WallabagUrl,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		UserName:     config.Username,
		UserPassword: config.Password,
	}
	wallabago.SetConfig(wallabagConfig)
}

var wallabagFetchLimit = 300

func WallabagGetEntries() wallabago.Entries {
	entries, err := wallabago.GetEntries(
		wallabago.APICall,
		0, 0, "", "", 1, wallabagFetchLimit, "", 0, -1, "", "")
	if err != nil {
		slog.Error("Cannot obtain articles from Wallabag")
		os.Exit(1)
	}

	return entries
}

func WallabagWriteTags(entry wallabago.Item, tags []string) {
	err := wallabago.AddEntryTags(entry.ID, tags...)
	if err != nil {
		slog.Error("Cannot assign tags to article", "title", entry.Title, "error", err)
	}
}

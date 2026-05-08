package core

import (
	"log/slog"
	"os"

	"github.com/Strubbl/wallabago/v9"
	_ "github.com/joho/godotenv/autoload"
)

var wallabagFetchLimit = 300
var wallabag = NewWallabagClient()

type WallabagClient interface {
	GetEntries() wallabago.Entries
	WriteTags(entry wallabago.Item, tags []string)
}

type wallabagClient struct{}

func NewWallabagClient() WallabagClient {
	wallabagConfig := wallabago.WallabagConfig{
		WallabagURL:  config.WallabagUrl,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		UserName:     config.Username,
		UserPassword: config.Password,
	}
	wallabago.SetConfig(wallabagConfig)
	return &wallabagClient{}
}

func (w *wallabagClient) GetEntries() wallabago.Entries {
	entries, err := wallabago.GetEntries(
		wallabago.APICall,
		0, 0, "", "", 1, wallabagFetchLimit, "", 0, -1, "", "")
	if err != nil {
		slog.Error("Cannot obtain articles from Wallabag")
		os.Exit(1)
	}

	return entries
}

func (w *wallabagClient) WriteTags(entry wallabago.Item, tags []string) {
	err := wallabago.AddEntryTags(entry.ID, tags...)
	if err != nil {
		slog.Error("Cannot assign tags to article", "title", entry.Title, "error", err)
	}
}

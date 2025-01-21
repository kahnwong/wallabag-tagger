package core

import (
	"log"
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

func WallabagGetEntries(perPage int) wallabago.Entries {
	entries, err := wallabago.GetEntries(
		wallabago.APICall,
		0, 0, "", "", 1, perPage, "", 0, -1, "", "")
	if err != nil {
		log.Println("Cannot obtain articles from Wallabag")
		os.Exit(1)
	}

	return entries
}

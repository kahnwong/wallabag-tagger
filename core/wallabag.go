package core

import (
	"os"

	"github.com/Strubbl/wallabago/v9"
	_ "github.com/joho/godotenv/autoload"
)

func WallabagInit() {
	wallabagConfig := wallabago.WallabagConfig{
		WallabagURL:  os.Getenv("WALLABAG_URL"),
		ClientID:     os.Getenv("WALLABAG_CLIENT_ID"),
		ClientSecret: os.Getenv("WALLABAG_CLIENT_SECRET"),
		UserName:     os.Getenv("WALLABAG_USERNAME"),
		UserPassword: os.Getenv("WALLABAG_PASSWORD"),
	}
	wallabago.SetConfig(wallabagConfig)
}

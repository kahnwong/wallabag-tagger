package core

import (
	cliBase "github.com/kahnwong/cli-base"
)

var config = cliBase.ReadYamlSops[Config]("~/.config/wallabag-tagger/config.sops.yaml")

type Config struct {
	WallabagUrl    string `yaml:"WALLABAG_URL"`
	ClientID       string `yaml:"CLIENT_ID"`
	ClientSecret   string `yaml:"CLIENT_SECRET"`
	Username       string `yaml:"USERNAME"`
	Password       string `yaml:"PASSWORD"`
	GoogleAIApiKey string `yaml:"GOOGLE_AI_API_KEY"`
}

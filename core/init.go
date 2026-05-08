package core

import (
	"log/slog"
	"os"

	cliBase "github.com/kahnwong/cli-base-sops"
)

var config = func() *Config {
	cfg, err := cliBase.ReadYamlSops[Config]("~/.config/wallabag-tagger/config.sops.yaml")
	if err != nil {
		slog.Error("Failed to read config", "error", err)
		os.Exit(1)
	}
	return cfg
}()

type Config struct {
	WallabagUrl   string `yaml:"WALLABAG_URL"`
	ClientID      string `yaml:"CLIENT_ID"`
	ClientSecret  string `yaml:"CLIENT_SECRET"`
	Username      string `yaml:"USERNAME"`
	Password      string `yaml:"PASSWORD"`
	OpenAiBaseUrl string `yaml:"OPENAI_BASE_URL"`
	OpenaiApiKey  string `yaml:"OPENAI_API_KEY"`
	ModelName     string `yaml:"MODEL_NAME"`
}

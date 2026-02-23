package core

import (
	"bytes"
	"context"
	"embed"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/zerolog/log"
	"google.golang.org/genai"
)

//go:embed resources/*
var templatesFS embed.FS

func renderPrompt(templatePath string, data any) string {
	// init template
	tmpl, err := template.ParseFS(templatesFS, templatePath)
	if err != nil {
		log.Fatal().Msgf("Error parsing template: %s", templatePath)
	}

	// render template
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		log.Fatal().Msgf("Error rendering now playing: %s", templatePath)
	}

	return tpl.String()
}

func GeminiGetTags(content string) (string, error) {
	var err error

	// init client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.GoogleAIApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal().Msg("Failed to create GOOGLE AI client")
	}

	// submit
	p := bluemonday.StripTagsPolicy()
	contentSanitized := p.Sanitize(
		content,
	)
	prompt := renderPrompt("resources/prompt.txt", map[string]interface{}{
		"Content": contentSanitized,
	})

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
	}

	iter := client.Models.GenerateContentStream(ctx, "gemini-2.5-flash",
		[]*genai.Content{{Parts: []*genai.Part{{Text: prompt}}}},
		config)

	var output string
	for resp, err := range iter {
		if err != nil {
			log.Warn().Msg("Failed to generate text")
			continue
		}
		output += resp.Text()
	}

	return output, err
}

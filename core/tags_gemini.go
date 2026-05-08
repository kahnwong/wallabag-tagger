package core

import (
	"bytes"
	"context"
	"embed"
	"log/slog"
	"os"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/genai"
)

//go:embed resources/*
var templatesFS embed.FS

func renderPrompt(templatePath string, data any) string {
	// init template
	tmpl, err := template.ParseFS(templatesFS, templatePath)
	if err != nil {
		slog.Error("Error parsing template", "path", templatePath)
		os.Exit(1)
	}

	// render template
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		slog.Error("Error rendering template", "path", templatePath)
		os.Exit(1)
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
		slog.Error("Failed to create GOOGLE AI client", "error", err)
		os.Exit(1)
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
			slog.Warn("Failed to generate text", "error", err)
			continue
		}
		output += resp.Text()
	}

	return output, err
}

package core

import (
	"bytes"
	"context"
	"embed"
	"log/slog"
	"os"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
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

func FetchLlmResponse(content string) (string, error) {
	// init client
	ctx := context.Background()
	client := openai.NewClient(
		option.WithBaseURL(config.OpenAiBaseUrl),
		option.WithAPIKey(config.OpenaiApiKey),
	)

	// submit
	p := bluemonday.StripTagsPolicy()
	contentSanitized := p.Sanitize(
		content,
	)
	prompt := renderPrompt("resources/prompt.txt", map[string]interface{}{
		"Content": contentSanitized,
	})

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
		Model: config.ModelName,
	})

	if err != nil {
		return "", err
	}

	return resp.OutputText(), nil
}

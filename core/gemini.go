package core

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GeminiGetTags(content string) string {
	// init client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GoogleAIApiKey))
	if err != nil {
		fmt.Println("Failed to create GOOGLE AI client")
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "application/json"

	// submit
	p := bluemonday.StripTagsPolicy()
	contentSanitized := p.Sanitize(
		content,
	)
	prompt := fmt.Sprintf("You are a bot in a read-it-later app and your responsibility is to help with automatic tagging.\nPlease analyze the text between the sentences \"CONTENT START HERE\" and \"CONTENT END HERE\" and suggest relevant tags that describe its key themes, topics, and main ideas. The rules are:\n- Aim for a variety of tags, including broad categories, specific keywords, and potential sub-genres.\n- The tags language must be in english.\n- If it's a famous website you may also include a tag for the website. If the tag is not generic enough, don't include it.\n- The content can include text for cookie consent and privacy policy, ignore those while tagging.\n- Aim for 3-5 tags.\n- If there are no good tags, leave the array empty.\n- You must respond in JSON using this JSON schema:      Tags = {\"tags\": list[str]}      Return: Tags\nCONTENT START HERE\n\n<%s>\n\nCONTENT END HERE\nYou must respond in JSON with the key \"tags\" and the value is an array of string tags.", contentSanitized)
	iter := model.GenerateContentStream(ctx, genai.Text(prompt))

	var output string
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Failed to generate text")
		}

		if resp.Candidates != nil {
			for _, v := range resp.Candidates {
				for _, k := range v.Content.Parts {
					output += fmt.Sprint(k.(genai.Text))
				}
			}
		}
	}

	return output
}

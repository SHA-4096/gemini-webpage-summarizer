package websum

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var ctx context.Context
var client *genai.Client
var cs *genai.ChatSession

func InitGemini(apiKey string) {
	ctx = context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	var err error
	client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}
	//defer client.Close()
	model := client.GenerativeModel("gemini-pro")
	cs = model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Now you are an assitant to analyze contents from pages and giving summary of the webpage"),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("OK, I will try to understand the content of the page and give you a summary."),
			},
			Role: "model",
		},
	}
}

func ImportExistingClientAndChatSession(importedClient *genai.Client, importedCs *genai.ChatSession, importedCtx context.Context) {
	client = importedClient
	cs = importedCs
	ctx = importedCtx
}

func ChatWithGemini(query string) (string, error) {
	fmt.Println(query)
	resp, err := cs.SendMessage(ctx, genai.Text(query))
	if err != nil {
		return "something went wrong", err
	}
	if len(resp.Candidates) == 0 {
		return "something went wrong", nil
	}
	if len(resp.Candidates[0].Content.Parts) == 0 {
		return "something went wrong", nil
	}
	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}

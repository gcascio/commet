package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type OllamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaChatRequest struct {
	Model    string          `json:"model"`
	Stream   bool            `json:"stream"`
	Messages []OllamaMessage `json:"messages"`
}

type OllamaChatResponse struct {
	Message OllamaMessage `json:"message"`
	Error   string        `json:"error"`
}

const systemPrompt = `
Your mission is to create clean and comprehensive commit messages as per the conventional commit convention and explain WHAT were the changes and mainly WHY the changes were done.
I'll send you an output of 'git diff --staged' command, and you are to convert it into a commit message.
Do not preface the commit with anything. Conventional commit keywords: fix, feat, build, chore, ci, docs, style, refactor, perf, test.
Example:

fix: catch error when user is not found

The application was crashing when a user was not found in the DB.
This commit fixes the issue by adding a check for the user's existence before proceeding with the operation.
`

func GenerateCommitMessage(diff string) string {
	llmUrl := viper.GetString("llm")
	model := viper.GetString("model")

	request := OllamaChatRequest{
		Model:  model,
		Stream: false,
		Messages: []OllamaMessage{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: diff,
			},
		},
	}

	postBody, _ := json.Marshal(request)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(
		llmUrl,
		"application/json",
		responseBody,
	)

	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var data OllamaChatResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
		return "Error occurred creating the commit message"
	}

	if data.Error != "" {
		return "Error occurred: " + data.Error
	}

	return strings.TrimSpace(data.Message.Content)
}

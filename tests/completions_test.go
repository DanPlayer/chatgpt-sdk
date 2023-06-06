package tests

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/DanPlayer/chatgpt-sdk/v1"
)

func TestCompletions(t *testing.T) {
	ctx := context.Background()
	completions, err := ChatGpt.Completions(ctx, v1.CompletionsRequest{
		Model:       v1.GPT3TextDavinci003,
		Prompt:      "你好",
		Suffix:      "",
		MaxTokens:   7,
		Temperature: 0,
	})
	if err != nil {
		fmt.Printf("completions error: %s", err.Error())
		return
	}
	fmt.Println(completions)
}

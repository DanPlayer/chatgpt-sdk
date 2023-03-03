package tests

import (
	"fmt"
	v1 "github.com/DanPlayer/chatgpt-sdk/v1"
	"testing"
)

func TestCompletions(t *testing.T) {
	completions, err := ChatGpt.Completions(v1.CompletionsRequest{
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

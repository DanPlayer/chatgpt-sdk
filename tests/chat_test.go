package tests

import (
	"fmt"
	v1 "github.com/DanPlayer/chatgpt-sdk/v1"
	"testing"
)

func TestChat(t *testing.T) {
	response, err := ChatGpt.CreateChatCompletion(v1.CreateChatCompletionRequest{
		Model: v1.GPT3Dot5Turbo,
		Messages: []v1.ChatMessage{
			{
				Role:    "user",
				Content: "你好",
			},
		},
	})
	if err != nil {
		fmt.Printf("completions error: %s", err.Error())
		return
	}
	fmt.Println(response)
}

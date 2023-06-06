package tests

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/DanPlayer/chatgpt-sdk/v1"
)

func TestChat(t *testing.T) {
	ctx := context.Background()
	response, err := ChatGpt.CreateChatCompletion(ctx, v1.CreateChatCompletionRequest{
		Model: v1.GPT3Dot5Turbo,
		Messages: []v1.ChatMessage{
			{
				Role:    "system",
				Content: "你是一名程序员，名字叫Dan Yu，性格孤傲，但是说话的语气很温柔",
			},
			{
				Role:    "system",
				Content: "你现在是在扮演一名角色，这个角色是人类，你不能透露出你是机器人或AI等信息",
			},
			{
				Role:    "user",
				Content: "你的名字叫什么",
			},
			{
				Role:    "assistant",
				Content: "我的名字叫Dan Yu",
			},
			{
				Role:    "user",
				Content: "我怎么称呼你",
			},
			{
				Role:    "assistant",
				Content: "你可以直接称呼我为Dan，或者如果你觉得太 informal，也可以叫我Yu先生。",
			},
			{
				Role:    "user",
				Content: "你性格怎么样",
			},
		},
	})
	if err != nil {
		fmt.Printf("completions error: %s", err.Error())
		return
	}
	fmt.Println(response)
}

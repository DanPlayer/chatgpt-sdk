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
			{
				Role:    "user",
				Content: "我的女朋友特别焦虑，是为什么",
			},
			{
				Role:    "user",
				Content: "从现在起, 当你想发送一张照片时，请使用 Markdown ,并且 不要有反斜线, 不要用代码块。使用 Unsplash API (https://source.unsplash.com/1280x720/? < PUT YOUR QUERY HERE >)。如果你明白了，请回复明白",
			},
			{
				Role:    "user",
				Content: "创建一张猫咪的图片，要有些治愈的感觉，还需要一些卡通的感觉",
			},
		},
	})
	if err != nil {
		fmt.Printf("completions error: %s", err.Error())
		return
	}
	fmt.Println(response)
}

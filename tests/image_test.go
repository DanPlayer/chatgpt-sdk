package tests

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/DanPlayer/chatgpt-sdk/v1"
)

func TestCreateImage(t *testing.T) {
	ctx := context.Background()
	response, err := ChatGpt.CreateImage(ctx, v1.CreateImageRequest{
		Prompt: "A stop-motion animation still of a cute robot standing in the forest",
	})
	if err != nil {
		fmt.Printf("completions error: %s", err.Error())
		return
	}
	fmt.Println(response)
}

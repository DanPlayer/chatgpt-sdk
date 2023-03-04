package tests

import (
	"fmt"
	v1 "github.com/DanPlayer/chatgpt-sdk/v1"
	"testing"
)

func TestCreateImage(t *testing.T) {
	response, err := ChatGpt.CreateImage(v1.CreateImageRequest{
		Prompt: "A stop-motion animation still of a cute robot standing in the forest",
	})
	if err != nil {
		fmt.Printf("completions error: %s", err.Error())
		return
	}
	fmt.Println(response)
}

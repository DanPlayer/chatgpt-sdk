package v1

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	CreateEdits = "https://api.openai.com/v1/edits"
)

type CreateEditsRequest struct {
	Model       string  `json:"model"`
	Input       string  `json:"input"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	TopP        float32 `json:"top_p,omitempty"`
}

type CreateEditsResponse struct {
	Object  string              `json:"object"`
	Created int                 `json:"created"`
	Choices []CreateEditsChoice `json:"choices"`
	Usage   Usage               `json:"usage"`
}

type CreateEditsChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

// CreateEdits Creates a new edit for the provided input, instruction, and parameters.
func (chat *ChatGpt) CreateEdits(ctx context.Context, req CreateEditsRequest) (response CreateEditsResponse, err error) {
	resp, err := chat.Post(ctx, CreateEdits, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

package v1

import (
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
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Text  string `json:"text"`
		Index int    `json:"index"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// CreateEdits Creates a new edit for the provided input, instruction, and parameters.
func (chat *ChatGpt) CreateEdits(req CreateEditsRequest) (response CreateEditsResponse, err error) {
	resp, err := chat.Post(CreateEdits, req)
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

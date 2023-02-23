package v1

import (
	"encoding/json"
	"fmt"
)

const (
	Models = "https://api.openai.com/v1/models"
	Model  = "https://api.openai.com/v1/models/%s"
)

type ModelsResponse struct {
	Data []struct {
		Id         string        `json:"id"`
		Object     string        `json:"object"`
		OwnedBy    string        `json:"owned_by"`
		Permission []interface{} `json:"permission"`
	} `json:"data"`
	Object string `json:"object"`
}

type ModelResponse struct {
	Id         string        `json:"id"`
	Object     string        `json:"object"`
	OwnedBy    string        `json:"owned_by"`
	Permission []interface{} `json:"permission"`
}

// Models Lists the currently available models, and provides basic information about each one such as the owner and availability.
func (chat *ChatGpt) Models() (response ModelsResponse, err error) {
	resp, err := chat.Get(Models, nil)
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

// Model Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
// param model The ID of the model to use for this request
func (chat *ChatGpt) Model(model string) (response ModelResponse, err error) {
	resp, err := chat.Get(fmt.Sprintf(Model, model), nil)
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

package v1

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	Moderations = "https://api.openai.com/v1/moderations"
)

type CreateModerationsRequest struct {
	Input string `json:"input"`
	Model string `json:"model,omitempty"`
}

type CreateModerationsResponse struct {
	Id      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Categories struct {
			Hate            bool `json:"hate"`
			HateThreatening bool `json:"hate/threatening"`
			SelfHarm        bool `json:"self-harm"`
			Sexual          bool `json:"sexual"`
			SexualMinors    bool `json:"sexual/minors"`
			Violence        bool `json:"violence"`
			ViolenceGraphic bool `json:"violence/graphic"`
		} `json:"categories"`
		CategoryScores struct {
			Hate            float64 `json:"hate"`
			HateThreatening float64 `json:"hate/threatening"`
			SelfHarm        float64 `json:"self-harm"`
			Sexual          float64 `json:"sexual"`
			SexualMinors    float64 `json:"sexual/minors"`
			Violence        float64 `json:"violence"`
			ViolenceGraphic float64 `json:"violence/graphic"`
		} `json:"category_scores"`
		Flagged bool `json:"flagged"`
	} `json:"results"`
}

// CreateModerations Classifies if text violates OpenAI's Content Policy
func (chat *ChatGpt) CreateModerations(ctx context.Context, req CreateModerationsRequest) (response CreateModerationsResponse, err error) {
	resp, err := chat.Post(ctx, Moderations, req)
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

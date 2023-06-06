package v1

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	FineTunes      = "https://api.openai.com/v1/fine-tunes"
	FineTune       = "https://api.openai.com/v1/fine-tunes/%s"
	CancelFineTune = "https://api.openai.com/v1/fine-tunes/%s/cancel"
	FineTuneEvents = "https://api.openai.com/v1/fine-tunes/%s/events"
)

type CreateFineTunesRequest struct {
	TrainingFile                 string        `json:"training_file"`
	ValidationFile               string        `json:"validation_file,omitempty"`
	Model                        string        `json:"model,omitempty"`
	NEpochs                      int           `json:"n_epochs,omitempty"`
	BatchSize                    int           `json:"batch_size,omitempty"`
	LearningRateMultiplier       float32       `json:"learning_rate_multiplier,omitempty"`
	PromptLossWeight             float32       `json:"prompt_loss_weight,omitempty"`
	ComputeClassificationMetrics bool          `json:"compute_classification_metrics,omitempty"`
	ClassificationNClasses       int           `json:"classification_n_classes,omitempty"`
	ClassificationPositiveClass  string        `json:"classification_positive_class,omitempty"`
	ClassificationBetas          []interface{} `json:"classification_betas,omitempty"`
	Suffix                       string        `json:"suffix,omitempty"`
}

type CreateFineTunesResponse struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Model     string `json:"model"`
	CreatedAt int    `json:"created_at"`
	Events    []struct {
		Object    string `json:"object"`
		CreatedAt int    `json:"created_at"`
		Level     string `json:"level"`
		Message   string `json:"message"`
	} `json:"events"`
	FineTunedModel interface{} `json:"fine_tuned_model"`
	Hyperparams    struct {
		BatchSize              int     `json:"batch_size"`
		LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
		NEpochs                int     `json:"n_epochs"`
		PromptLossWeight       float64 `json:"prompt_loss_weight"`
	} `json:"hyperparams"`
	OrganizationId  string        `json:"organization_id"`
	ResultFiles     []interface{} `json:"result_files"`
	Status          string        `json:"status"`
	ValidationFiles []interface{} `json:"validation_files"`
	TrainingFiles   []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"training_files"`
	UpdatedAt int `json:"updated_at"`
}

// CreateFineTunes Creates a job that fine-tunes a specified model from a given dataset.
// Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.
func (chat *ChatGpt) CreateFineTunes(ctx context.Context, req CreateFineTunesRequest) (response CreateFineTunesResponse, err error) {
	resp, err := chat.Post(ctx, FineTunes, req)
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

type FineTunesResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Id             string      `json:"id"`
		Object         string      `json:"object"`
		Model          string      `json:"model"`
		CreatedAt      int         `json:"created_at"`
		FineTunedModel interface{} `json:"fine_tuned_model"`
		Hyperparams    struct {
		} `json:"hyperparams"`
		OrganizationId  string        `json:"organization_id"`
		ResultFiles     []interface{} `json:"result_files"`
		Status          string        `json:"status"`
		ValidationFiles []interface{} `json:"validation_files"`
		TrainingFiles   []interface{} `json:"training_files"`
		UpdatedAt       int           `json:"updated_at"`
	} `json:"data"`
}

// FineTunes List your organization's fine-tuning jobs
func (chat *ChatGpt) FineTunes(ctx context.Context) (response FineTunesResponse, err error) {
	resp, err := chat.Get(ctx, FineTunes, nil)
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

type FineTuneResponse struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Model     string `json:"model"`
	CreatedAt int    `json:"created_at"`
	Events    []struct {
		Object    string `json:"object"`
		CreatedAt int    `json:"created_at"`
		Level     string `json:"level"`
		Message   string `json:"message"`
	} `json:"events"`
	FineTunedModel string `json:"fine_tuned_model"`
	Hyperparams    struct {
		BatchSize              int     `json:"batch_size"`
		LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
		NEpochs                int     `json:"n_epochs"`
		PromptLossWeight       float64 `json:"prompt_loss_weight"`
	} `json:"hyperparams"`
	OrganizationId string `json:"organization_id"`
	ResultFiles    []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"result_files"`
	Status          string        `json:"status"`
	ValidationFiles []interface{} `json:"validation_files"`
	TrainingFiles   []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"training_files"`
	UpdatedAt int `json:"updated_at"`
}

// FineTune Gets info about the fine-tune job.
func (chat *ChatGpt) FineTune(ctx context.Context, id string) (response FineTuneResponse, err error) {
	resp, err := chat.Get(ctx, fmt.Sprintf(FineTune, id), nil)
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

type CancelFineTuneResponse struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Model     string `json:"model"`
	CreatedAt int    `json:"created_at"`
	Events    []struct {
	} `json:"events"`
	FineTunedModel interface{} `json:"fine_tuned_model"`
	Hyperparams    struct {
	} `json:"hyperparams"`
	OrganizationId  string        `json:"organization_id"`
	ResultFiles     []interface{} `json:"result_files"`
	Status          string        `json:"status"`
	ValidationFiles []interface{} `json:"validation_files"`
	TrainingFiles   []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"training_files"`
	UpdatedAt int `json:"updated_at"`
}

// CancelFineTune Immediately cancel a fine-tune job.
func (chat *ChatGpt) CancelFineTune(ctx context.Context, id string) (response CancelFineTuneResponse, err error) {
	resp, err := chat.Post(ctx, fmt.Sprintf(CancelFineTune, id), nil)
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

type FineTuneEventsResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string `json:"object"`
		CreatedAt int    `json:"created_at"`
		Level     string `json:"level"`
		Message   string `json:"message"`
	} `json:"data"`
}

func (chat *ChatGpt) FineTuneEvents(ctx context.Context, id string) (response FineTuneEventsResponse, err error) {
	resp, err := chat.Get(ctx, fmt.Sprintf(FineTuneEvents, id), nil)
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

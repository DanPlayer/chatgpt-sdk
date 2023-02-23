package v1

import (
	"encoding/json"
	"fmt"
)

const (
	Completions = "https://api.openai.com/v1/completions"
)

type CompletionsRequest struct {
	Model            string                 `json:"model"`                       // ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	Prompt           string                 `json:"prompt"`                      // The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays. Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.
	Suffix           string                 `json:"suffix,omitempty"`            // The suffix that comes after a completion of inserted text.
	MaxTokens        int                    `json:"max_tokens"`                  // The maximum number of tokens to generate in the completion. The token count of your prompt plus max_tokens cannot exceed the model's context length. Most models have a context length of 2048 tokens (except for the newest models, which support 4096).
	Temperature      float32                `json:"temperature"`                 // What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. We generally recommend altering this or top_p but not both.
	TopP             int                    `json:"top_p,omitempty"`             // An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered. We generally recommend altering this or temperature but not both.
	N                int                    `json:"n,omitempty"`                 // How many completions to generate for each prompt. Note: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.
	Stream           bool                   `json:"stream,omitempty"`            // Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a data: [DONE] message.
	Logprobs         interface{}            `json:"logprobs,omitempty"`          // logprobs integer Optional Defaults to null Include the log probabilities on the logprobs most likely tokens, as well the chosen tokens. For example, if logprobs is 5, the API will return a list of the 5 most likely tokens. The API will always return the logprob of the sampled token, so there may be up to logprobs+1 elements in the response. The maximum value for logprobs is 5. If you need more than this, please contact us through our Help center and describe your use case.
	Echo             bool                   `json:"echo,omitempty"`              // Echo back the prompt in addition to the completion
	Stop             string                 `json:"stop,omitempty"`              // Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.
	PresencePenalty  float32                `json:"presence_penalty,omitempty"`  // presence_penalty number Optional Defaults to 0 Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	FrequencyPenalty float32                `json:"frequency_penalty,omitempty"` // Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	BestOf           int                    `json:"best_of,omitempty"`           // best_of integer Optional Defaults to 1 Generates best_of completions server-side and returns the "best" (the one with the highest log probability per token). Results cannot be streamed. When used with n, best_of controls the number of candidate completions and n specifies how many to return – best_of must be greater than n. Note: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.
	LogitBias        map[string]interface{} `json:"logit_bias,omitempty"`        // Modify the likelihood of specified tokens appearing in the completion. Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this tokenizer tool (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token. As an example, you can pass {"50256": -100} to prevent the <|endoftext|> token from being generated.
	User             string                 `json:"user,omitempty"`              // A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
}

type CompletionsResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// Completions 用于生成文本。
// 您可以提供一个起始文本（prompt），然后ChatGPT将自动根据这个文本生成接下来的文本。
// 您可以控制生成的文本长度、温度（即创作性）、频率和置信度等参数。
func (chat *ChatGpt) Completions(req CompletionsRequest) (response CompletionsResponse, err error) {
	resp, err := chat.Post(Completions, req)
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

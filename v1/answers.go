package v1

import (
	"encoding/json"
	"fmt"
)

const (
	Answers = "https://api.openai.com/v1/answers"
)

type AnswersRequest struct {
	Model     string `json:"model"`  // 用于生成答案的 GPT 模型的 ID，例如 "davinci" 或 "curie"
	Prompt    string `json:"prompt"` //  (必需) 要回答的问题。
	Documents []struct {
		Text     string `json:"text"` // 用于生成答案的文档文本
		Metadata struct {
			Name   string `json:"name"`
			Source string `json:"source"`
		} `json:"metadata"`
	} `json:"documents"` // 一个列表，其中包含用于生成答案的文档对象
}

type AnswersResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Model   string `json:"model"`
	Created int    `json:"created"`
	Answers []struct {
		Document struct {
			Text     string `json:"text"`
			Metadata struct {
				Name   string `json:"name"`
				Source string `json:"source"`
			} `json:"metadata"`
		} `json:"document"`
		Text       string  `json:"text"`
		Confidence float64 `json:"confidence"`
	} `json:"answers"`
}

// Answers 用于自动回答问题。
// 您可以提供一个问题和一些上下文信息，然后ChatGPT将自动回答该问题。
// 这个API非常适合于构建智能客服、问答系统和搜索引擎等应用程序。
func (chat *ChatGpt) Answers(req AnswersRequest) (response AnswersResponse, err error) {
	resp, err := chat.Post(Answers, req)
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

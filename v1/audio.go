package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	Transcriptions = "https://api.openai.com/v1/audio/transcriptions"
	Translations   = "https://api.openai.com/v1/audio/translations"
)

type CreateTranscriptionsRequest struct {
	File           *os.File `json:"file"`                      // The audio file to transcribe, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	Model          string   `json:"model"`                     // ID of the model to use. Only whisper-1 is currently available.
	Prompt         string   `json:"prompt,omitempty"`          // An optional text to guide the model's style or continue a previous audio segment. The prompt should match the audio language.
	ResponseFormat string   `json:"response_format,omitempty"` // The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
	Temperature    float32  `json:"temperature,omitempty"`     // The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.
	Language       string   `json:"language,omitempty"`        // The language of the input audio. Supplying the input language in ISO-639-1 format will improve accuracy and latency.
}

type CreateTranscriptionsResponse struct {
	Text string `json:"text"`
}

func (chat *ChatGpt) CreateTranscriptions(req CreateTranscriptionsRequest) (response CreateTranscriptionsResponse, err error) {
	// 创建multipart/form-data格式的body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 设置file参数
	filePart, err := writer.CreateFormFile("file", req.File.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(filePart, req.File)

	// 设置其他参数
	if req.Model != "" {
		_ = writer.WriteField("model", req.Model)
	}
	if req.Prompt != "" {
		_ = writer.WriteField("prompt", req.Prompt)
	}
	if req.ResponseFormat != "" {
		_ = writer.WriteField("response_format", req.ResponseFormat)
	}
	if req.Temperature != 0 {
		_ = writer.WriteField("temperature", fmt.Sprintf("%f", req.Temperature))
	}
	if req.Language != "" {
		_ = writer.WriteField("language", req.Language)
	}

	// 结束body的编写
	_ = writer.Close()

	// 创建http请求
	request, err := http.NewRequest("POST", Transcriptions, &requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置header中的Authorization字段
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送http请求
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

type CreateTranslationsRequest struct {
	File           *os.File `json:"file"`                      // The audio file to translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	Model          string   `json:"model"`                     // ID of the model to use. Only whisper-1 is currently available.
	Prompt         string   `json:"prompt,omitempty"`          // An optional text to guide the model's style or continue a previous audio segment. The prompt should be in English.
	ResponseFormat string   `json:"response_format,omitempty"` // The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
	Temperature    float32  `json:"temperature,omitempty"`     // The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.
}

type CreateTranslationsResponse struct {
	Text string `json:"text"`
}

func (chat *ChatGpt) CreateTranslations(req CreateTranslationsRequest) (response CreateTranslationsResponse, err error) {
	// 创建multipart/form-data格式的body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 设置file参数
	filePart, err := writer.CreateFormFile("file", req.File.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(filePart, req.File)

	// 设置其他参数
	if req.Model != "" {
		_ = writer.WriteField("model", req.Model)
	}
	if req.Prompt != "" {
		_ = writer.WriteField("prompt", req.Prompt)
	}
	if req.ResponseFormat != "" {
		_ = writer.WriteField("response_format", req.ResponseFormat)
	}
	if req.Temperature != 0 {
		_ = writer.WriteField("temperature", fmt.Sprintf("%f", req.Temperature))
	}
	// 结束body的编写
	_ = writer.Close()

	// 创建http请求
	request, err := http.NewRequest("POST", Translations, &requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置header中的Authorization字段
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送http请求
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

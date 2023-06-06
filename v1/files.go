package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	Files       = "https://api.openai.com/v1/files"
	UploadFile  = "https://api.openai.com/v1/files"
	File        = "https://api.openai.com/v1/files/%s"
	FileContent = "https://api.openai.com/v1/files/%s/content"
)

type FilesResponse struct {
	Data []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"data"`
	Object string `json:"object"`
}

// Files are used to upload documents that can be used with features like Fine-tuning.
func (chat *ChatGpt) Files() (ctx context.Context, response FilesResponse, err error) {
	resp, err := chat.Get(ctx, Files, nil)
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

type UploadFileRequest struct {
	File    *os.File `json:"file"`    // If the purpose is set to "fine-tune", each line is a JSON record with "prompt" and "completion" fields representing your training examples.
	Purpose string   `json:"purpose"` // The intended purpose of the uploaded documents.  Use "fine-tune" for Fine-tuning. This allows us to validate the format of the uploaded file.
}

type UploadFileResponse struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

// UploadFile Upload a file that contains document(s) to be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you need to increase the storage limit.
func (chat *ChatGpt) UploadFile(ctx context.Context, req UploadFileRequest) (response UploadFileResponse, err error) {
	// 创建multipart/form-data格式的body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 设置file参数
	filePart, err := writer.CreateFormFile("image", req.File.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(filePart, req.File)

	// 设置其他参数
	if req.Purpose != "" {
		_ = writer.WriteField("purpose", req.Purpose)
	}

	// 结束body的编写
	_ = writer.Close()

	// 创建http请求
	request, err := http.NewRequestWithContext(ctx, "POST", UploadFile, &requestBody)
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

type DeleteFileResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

func (chat *ChatGpt) DeleteFile(ctx context.Context, id string) (response DeleteFileResponse, err error) {
	resp, err := chat.Delete(ctx, fmt.Sprintf(File, id), nil)
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

type RetrieveFileResponse struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

// RetrieveFile Returns information about a specific file.
func (chat *ChatGpt) RetrieveFile(ctx context.Context, id string) (response FilesResponse, err error) {
	resp, err := chat.Get(ctx, fmt.Sprintf(File, id), nil)
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

// RetrieveFileContent Returns the contents of the specified file
func (chat *ChatGpt) RetrieveFileContent(ctx context.Context, id string) ([]byte, error) {
	return chat.Get(ctx, fmt.Sprintf(FileContent, id), nil)
}

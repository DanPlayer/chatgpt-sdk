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
	"strconv"
)

const (
	CreateImage           = "https://api.openai.com/v1/images/generations"
	CreateImageEdit       = "https://api.openai.com/v1/images/edits"
	CreateImageVariations = "https://api.openai.com/v1/images/variations"
)

type CreateImageRequest struct {
	Prompt         string `json:"prompt"`                    // A text description of the desired image(s). The maximum length is 1000 characters
	N              int    `json:"n,omitempty"`               // The number of images to generate. Must be between 1 and 10
	Size           string `json:"size,omitempty"`            // The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024
	ResponseFormat string `json:"response_format,omitempty"` // The format in which the generated images are returned. Must be one of url or b64_json
	User           string `json:"user,omitempty"`            // A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
}

type CreateImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// CreateImage Creates an image given a prompt.
func (chat *ChatGpt) CreateImage(req CreateImageRequest) (response CreateImageResponse, err error) {
	resp, err := chat.Post(CreateImage, req)
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

type CreateImageEditRequest struct {
	Image          *os.File `json:"image"`                     // The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask
	Mask           *os.File `json:"mask,omitempty"`            // An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as image
	Prompt         string   `json:"prompt"`                    // A text description of the desired image(s). The maximum length is 1000 characters
	N              int      `json:"n,omitempty"`               // The number of images to generate. Must be between 1 and 10
	Size           string   `json:"size,omitempty"`            // The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024
	ResponseFormat string   `json:"response_format,omitempty"` // The format in which the generated images are returned. Must be one of url or b64_json
	User           string   `json:"user,omitempty"`            // A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
}

type CreateImageEditResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// CreateImageEdit Creates an edited or extended image given an original image and a prompt.
func (chat *ChatGpt) CreateImageEdit(req CreateImageEditRequest) (response CreateImageEditResponse, err error) {
	// 创建multipart/form-data格式的body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 设置image参数
	imagePart, err := writer.CreateFormFile("image", req.Image.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(imagePart, req.Image)

	// 设置mask参数
	maskPart, err := writer.CreateFormFile("mask", req.Mask.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(maskPart, req.Mask)

	// 设置其他参数
	if req.Prompt != "" {
		_ = writer.WriteField("prompt", req.Prompt)
	}
	if req.N != 0 {
		_ = writer.WriteField("n", strconv.Itoa(req.N))
	}
	if req.Size != "" {
		_ = writer.WriteField("size", req.Size)
	}
	if req.ResponseFormat != "" {
		_ = writer.WriteField("response_format", req.ResponseFormat)
	}
	if req.User != "" {
		_ = writer.WriteField("user", req.User)
	}

	// 结束body的编写
	_ = writer.Close()

	// 创建http请求
	request, err := http.NewRequest("POST", CreateImageEdit, &requestBody)
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

type CreateImageVariationsRequest struct {
	Image          *os.File `json:"image"`                     // The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask
	Prompt         string   `json:"prompt"`                    // A text description of the desired image(s). The maximum length is 1000 characters
	N              int      `json:"n,omitempty"`               // The number of images to generate. Must be between 1 and 10
	Size           string   `json:"size,omitempty"`            // The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024
	ResponseFormat string   `json:"response_format,omitempty"` // The format in which the generated images are returned. Must be one of url or b64_json
	User           string   `json:"user,omitempty"`            // A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
}

type CreateImageVariationsResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// CreateImageVariations Creates a variation of a given image.
func (chat *ChatGpt) CreateImageVariations(req CreateImageVariationsRequest) (response CreateImageVariationsResponse, err error) {
	// 创建multipart/form-data格式的body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 设置image参数
	imagePart, err := writer.CreateFormFile("image", req.Image.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(imagePart, req.Image)

	// 设置其他参数
	if req.Prompt != "" {
		_ = writer.WriteField("prompt", req.Prompt)
	}
	if req.N != 0 {
		_ = writer.WriteField("n", strconv.Itoa(req.N))
	}
	if req.Size != "" {
		_ = writer.WriteField("size", req.Size)
	}
	if req.ResponseFormat != "" {
		_ = writer.WriteField("response_format", req.ResponseFormat)
	}
	if req.User != "" {
		_ = writer.WriteField("user", req.User)
	}

	// 结束body的编写
	_ = writer.Close()

	// 创建http请求
	request, err := http.NewRequest("POST", CreateImageVariations, &requestBody)
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

# chatgpt-sdk

[![Go Report Card](https://goreportcard.com/badge/github.com/DanPlayer/chatgpt-sdk)](https://goreportcard.com/report/github.com/DanPlayer/chatgpt-sdk)

chatgpt-sdk是一个使用Go语言编写的简单而强大的chatgpt sdk，可以方便地调用[ChatGpt-API](https://platform.openai.com/)的各种接口，实现文本生成、代码生成和图像生成等功能。它支持国内使用代理模式，避免网络问题。它还提供了一些常用的模型和参数的常量，简化用户的选择。

## 安装

使用`go get`命令安装sdk：

```go
go get github.com/DanPlayer/chatgpt-sdk
```

## 使用

首先，创建一个`v1.Client`对象，传入`v1.ChatGptOption`结构体，包含以下字段：

- `SecretKey`：必须，你的ChatGpt-API密钥
- `HasProxy`：可选，是否使用代理模式，默认为false
- `ProxyUrl`：可选，代理服务器的URL，如`http://localhost:7890`

例如：

```go
const SecretKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXX"
var ChatGpt = v1.Client(v1.ChatGptOption{
    SecretKey: SecretKey,
    HasProxy:  true,
    ProxyUrl:  "http://localhost:7890",
})
```

然后，调用`v1.Client`对象的各种方法，传入相应的请求结构体，获取响应结构体或错误信息。以下是一些常用的方法：

- `Models()`：获取可用的模型列表
- `Completions()`：根据给定的模型和提示生成文本
- `Chat()`：根据给定的模型和会话历史进行聊天
- `Audio()`：根据给定的模型和文本生成音频
- `CreateImage()`：根据给定的模型和文本生成图像

具体的请求和响应结构体的字段，请参考[ChatGpt-API文档](https://platform.openai.com/)。

## 示例

以下是一些使用chatgpt-sdk的示例代码：

### 文本生成

```go
completions, err := ChatGpt.Completions(v1.CompletionsRequest{
    Model:       v1.GPT3TextDavinci003,
    Prompt:      "Say this is a test",
    Suffix:      "",
    MaxTokens:   7,
    Temperature: 0,
})
if err != nil {
    fmt.Printf("completions error: %s", err.Error())
    return
}
fmt.Println(completions)
```

输出：

```json
{
  "id": "cmpl-3Z4Jy9c8w4lZ6j2wzQxqkQhY5",
  "object": "text_completion",
  "created": 1639648229,
  "model": "text-davinci-003",
  "choices": [
    {
      "text": ". This is a test",
      "index": 0,
      "logprobs": null,
      "finish_reason": "stop"
    }
  ]
}
```

### 代码生成

```go
completions, err := ChatGpt.Completions(v1.CompletionsRequest{
    Model:     v1.CodexCodeDavinci002,
    Prompt:    "// Write a function in Go that takes a slice of integers and returns the sum of the elements\nfunc sum(",
    Suffix:    "",
    MaxTokens: 20,
})
if err != nil {
    fmt.Printf("completions error: %s", err.Error())
    return
}
fmt.Println(completions)
```

输出：

```json
{
  "id": "cmpl-3Z4K2Rn8w4lZ6j2wzQxqkQhY5",
  "object": "text_completion",
  "created": 1639648397,
  "model": "code-davinci-002",
  "choices": [
    {
      "text": "nums []int) int {\n\tsum := 0\n\tfor _, n := range nums {\n\t\tsum += n\n\t}\n\treturn sum\n}",
      "index": 0,
      "logprobs": null,
      "finish_reason": "stop"
    }
  ]
}
```

### 图像生成

```go
createImage, err := ChatGpt.CreateImage(v1.CreateImageRequest{
    Model:           v1.GPT3Dot5Turbo0301,
    Query:           "a cute cat wearing a hat",
    Size:            v1.CreateImageSize256x256,
    ResponseFormat:  v1.CreateImageResponseFormatURL,
})
if err != nil {
    fmt.Printf("createImage error: %s", err.Error())
    return
}
fmt.Println(createImage)
```

输出：

```json
{
  "id": "img-3Z4K8Xn8w4lZ6j2wzQxqkQhY5",
  "object": "image_completion",
  "created": 1639648547,
  "model": "gpt-3.5-turbo-0301",
  "url": "<image url>"
}
```

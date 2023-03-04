# chatgpt-sdk
simple chatgpt sdk，已经支持国内使用代理模式，调用接口

### 使用
```
go get github.com/DanPlayer/chatgpt-sdk
```

### ChatGpt-API具体文档
[文档地址](https://platform.openai.com/docs/api-reference/introduction)

```
// 一些可用模型的常量，更新的话可以从文档里查询，以原文档为准
const (
	GPT3Dot5Turbo0301       = "gpt-3.5-turbo-0301"
	GPT3Dot5Turbo           = "gpt-3.5-turbo"
	GPT3TextDavinci003      = "text-davinci-003"
	GPT3TextDavinci002      = "text-davinci-002"
	GPT3TextCurie001        = "text-curie-001"
	GPT3TextBabbage001      = "text-babbage-001"
	GPT3TextAda001          = "text-ada-001"
	GPT3TextDavinci001      = "text-davinci-001"
	GPT3DavinciInstructBeta = "davinci-instruct-beta"
	GPT3Davinci             = "davinci"
	GPT3CurieInstructBeta   = "curie-instruct-beta"
	GPT3Curie               = "curie"
	GPT3Ada                 = "ada"
	GPT3Babbage             = "babbage"
)

const (
	CodexCodeDavinci002 = "code-davinci-002"
	CodexCodeCushman001 = "code-cushman-001"
	CodexCodeDavinci001 = "code-davinci-001"
)

const (
	CreateImageSize256x256   = "256x256"
	CreateImageSize512x512   = "512x512"
	CreateImageSize1024x1024 = "1024x1024"
)

const (
	CreateImageResponseFormatURL     = "url"
	CreateImageResponseFormatB64JSON = "b64_json"
)
```

### 示例
```
const SecretKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXX"

var ChatGpt = v1.Client(v1.ChatGptOption{SecretKey: SecretKey, HasProxy: true, ProxyUrl: "http://localhost:7890"})

func Models() {
	models, err := ChatGpt.Models()
	if err != nil {
		fmt.Printf("models error: %s", err.Error())
		return
	}
	fmt.Println(models)
}

func TestCompletions() {
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
}
```
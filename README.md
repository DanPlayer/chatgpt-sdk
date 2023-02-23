# chatgpt-sdk
simple chatgpt sdk

### 使用
```
go get github.com/DanPlayer/chatgpt-sdk
```

### ChatGpt-API具体文档
[文档地址](https://platform.openai.com/docs/api-reference/introduction)

### 示例
```
const SecretKey = "XXXXXXXXXXXXXXXXXXXXX"

var ChatGpt = v1.Client(v1.ChatGptOption{SecretKey: SecretKey})

func Models() {
	models, err := ChatGpt.Models()
	if err != nil {
		fmt.Printf("models error: %s", err.Error())
		return
	}
	fmt.Println(models)
}
```
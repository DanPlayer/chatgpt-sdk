# chatgpt-sdk
simple chatgpt sdk

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
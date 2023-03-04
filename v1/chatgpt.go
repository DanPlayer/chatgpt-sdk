package v1

type ChatGpt struct {
	Authorization string
	HasProxy      bool
	ProxyUrl      string
}

type ChatGptOption struct {
	SecretKey string // OpenAI AuthKey
	HasProxy  bool   // proxy http
	ProxyUrl  string
}

func Client(option ChatGptOption) *ChatGpt {
	return &ChatGpt{
		Authorization: option.SecretKey,
		HasProxy:      option.HasProxy,
		ProxyUrl:      option.ProxyUrl,
	}
}

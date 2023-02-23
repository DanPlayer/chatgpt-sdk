package v1

type ChatGpt struct {
	Authorization string
}

type ChatGptOption struct {
	SecretKey string
}

func Client(option ChatGptOption) *ChatGpt {
	return &ChatGpt{
		Authorization: option.SecretKey,
	}
}

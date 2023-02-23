package tests

import v1 "github.com/DanPlayer/chatgpt-sdk/v1"

const SecretKey = "XXXXXXXXXXXXXXXXXXXXXXXXX"

var ChatGpt = v1.Client(v1.ChatGptOption{SecretKey: SecretKey})

package tests

import v1 "github.com/DanPlayer/chatgpt-sdk/v1"

const SecretKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXX"

var ChatGpt = v1.Client(v1.ChatGptOption{SecretKey: SecretKey, HasProxy: true, ProxyUrl: "http://localhost:7890"})

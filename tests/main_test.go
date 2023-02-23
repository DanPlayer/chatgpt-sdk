package tests

import v1 "github.com/DanPlayer/chatgpt-sdk/v1"

const SecretKey = "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"

var ChatGpt = v1.Client(v1.ChatGptOption{SecretKey: SecretKey})

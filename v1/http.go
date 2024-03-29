package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
)

func (chat *ChatGpt) Proxy() *http.Transport {
	// 设置代理地址
	proxyURL, err := neturl.Parse(chat.ProxyUrl)
	if err != nil {
		panic(err)
	}

	// 创建一个Transport
	return &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
}

func (chat *ChatGpt) Post(ctx context.Context, url string, data interface{}) ([]byte, error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))

	var client *http.Client
	if chat.HasProxy {
		client = &http.Client{
			Transport: chat.Proxy(),
		}
	} else {
		client = &http.Client{}
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (chat *ChatGpt) Get(ctx context.Context, url string, data interface{}) ([]byte, error) {
	var (
		requestBody []byte
		err         error
	)
	if data == nil {
		requestBody = nil
	} else {
		requestBody, err = json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return []byte{}, nil
		}
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))

	var client *http.Client
	if chat.HasProxy {
		client = &http.Client{
			Transport: chat.Proxy(),
		}
	} else {
		client = &http.Client{}
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (chat *ChatGpt) Delete(ctx context.Context, url string, data interface{}) ([]byte, error) {
	var (
		requestBody []byte
		err         error
	)
	if data == nil {
		requestBody = nil
	} else {
		requestBody, err = json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return []byte{}, nil
		}
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))

	var client *http.Client
	if chat.HasProxy {
		client = &http.Client{
			Transport: chat.Proxy(),
		}
	} else {
		client = &http.Client{}
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

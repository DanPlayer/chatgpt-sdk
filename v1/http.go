package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (chat *ChatGpt) Post(url string, data interface{}) ([]byte, error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (chat *ChatGpt) Get(url string, data interface{}) ([]byte, error) {
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

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (chat *ChatGpt) Delete(url string, data interface{}) ([]byte, error) {
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

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", chat.Authorization))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

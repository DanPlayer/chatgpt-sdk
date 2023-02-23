package tests

import (
	"fmt"
	"testing"
)

func TestModels(t *testing.T) {
	models, err := ChatGpt.Models()
	if err != nil {
		fmt.Printf("models error: %s", err.Error())
		return
	}
	fmt.Println(models)
}

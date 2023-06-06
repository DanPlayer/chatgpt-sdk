package tests

import (
	"context"
	"fmt"
	"testing"
)

func TestModels(t *testing.T) {
	ctx := context.Background()
	models, err := ChatGpt.Models(ctx)
	if err != nil {
		fmt.Printf("models error: %s", err.Error())
		return
	}
	fmt.Println(models)
}

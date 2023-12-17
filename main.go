package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context) (*string, error) {
	message := fmt.Sprintf("Hello %s!", "Vitor")
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}

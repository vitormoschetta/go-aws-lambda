package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event any) (any, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	response, ok := event.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received event of type %T, expected map[string]interface{}", event)
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return string(responseJSON), nil
}

func main() {
	lambda.Start(HandleRequest)
}

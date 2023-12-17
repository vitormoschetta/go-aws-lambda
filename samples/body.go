package samples

import (
	"context"
	"fmt"
)

func SampleRequestWithBody(ctx context.Context, event any) (any, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	eventMap, ok := event.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received event of type %T, expected map[string]interface{}", event)
	}

	body, ok := eventMap["body"].(string) // body is a json string
	if !ok {
		return nil, fmt.Errorf("received event with body of type %T, expected string", eventMap["body"])
	}

	return body, nil
}

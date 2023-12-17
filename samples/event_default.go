package samples

import (
	"context"
	"encoding/json"
	"fmt"
)

// Neste exemplo temos acesso ao evento completo, porem precisamos fazer o parse manual para cada nivel de profundidade do evento
func HandlerRequestDefault(ctx context.Context, event any) (any, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	eventMap, ok := event.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received event of type %T, expected map[string]interface{}", event)
	}

	headers, ok := eventMap["headers"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received headers of type %T, expected map[string]interface{}", eventMap["headers"])
	}

	body, ok := eventMap["body"].(string)
	if !ok {
		return nil, fmt.Errorf("received body of type %T, expected string", eventMap["body"])
	}

	query, ok := eventMap["rawQueryString"].(string)
	if !ok {
		return nil, fmt.Errorf("received rawQueryString of type %T, expected string", eventMap["rawQueryString"])
	}

	reqContext, ok := eventMap["requestContext"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received requestContext of type %T, expected map[string]interface{}", eventMap["requestContext"])
	}

	httpContext, ok := reqContext["http"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received http of type %T, expected map[string]interface{}", reqContext["http"])
	}

	path, ok := httpContext["path"].(string)
	if !ok {
		return nil, fmt.Errorf("received path of type %T, expected string", httpContext["path"])
	}

	httpMethod, ok := httpContext["method"].(string)
	if !ok {
		return nil, fmt.Errorf("received method of type %T, expected string", httpContext["method"])
	}

	resMap := map[string]interface{}{
		"headers": headers,
		"body":    body,
		"query":   query,
		"path":    path,
		"method":  httpMethod,
		"event":   event,
	}

	responseJSON, err := json.Marshal(resMap)
	if err != nil {
		return nil, err
	}

	return string(responseJSON), nil
}

package samples

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func SampleRequestWithRouter(ctx context.Context, event any) (any, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	eventMap, ok := event.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received event of type %T, expected map[string]interface{}", event)
	}

	reqContext, ok := eventMap["requestContext"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received requestContext of type %T, expected map[string]interface{}", eventMap["requestContext"])
	}

	httpContext, ok := reqContext["http"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received http of type %T, expected map[string]interface{}", reqContext["http"])
	}

	httpMethod, ok := httpContext["method"].(string)
	if !ok {
		return nil, fmt.Errorf("received method of type %T, expected string", httpContext["method"])
	}

	path, ok := httpContext["path"].(string)
	if !ok {
		return nil, fmt.Errorf("received path of type %T, expected string", httpContext["path"])
	}

	if httpMethod == "GET" && path == "/healthcheck" {
		return healthcheck()
	}

	if httpMethod == "GET" && strings.Contains(path, "/products") {
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 3 {
			return nil, fmt.Errorf("received path %s, expected /products/<id>", path)
		}

		id, err := strconv.Atoi(pathParts[2])
		if err != nil {
			return nil, fmt.Errorf("received path %s, expected /products/<id>", path)
		}

		return GetProduct(ctx, id)
	}

	if httpMethod == "POST" && strings.Contains(path, "/products") {
		body, ok := eventMap["body"].(string)
		if !ok {
			return nil, fmt.Errorf("received body of type %T, expected string", eventMap["body"])
		}

		var product Product
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			return nil, err
		}

		return PostProduct(ctx, product)
	}

	return GetRequestContext(ctx, event)
}

func healthcheck() (*string, error) {
	res := map[string]interface{}{
		"status": "ok",
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	resString := string(resJSON)
	return &resString, nil
}

func PostProduct(ctx context.Context, product Product) (*string, error) {
	res := map[string]interface{}{
		"data": product,
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	resString := string(resJSON)
	return &resString, nil
}

func GetProduct(ctx context.Context, id int) (*string, error) {
	res := map[string]interface{}{
		"data": Product{
			ID:    id,
			Name:  "Product " + strconv.Itoa(id),
			Price: 100.0,
		},
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	resString := string(resJSON)
	return &resString, nil
}

func GetRequestContext(ctx context.Context, event any) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	eventMap, ok := event.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("received event of type %T, expected map[string]interface{}", event)
	}

	responseJSON, err := json.Marshal(eventMap)
	if err != nil {
		return nil, err
	}

	responseString := string(responseJSON)
	return &responseString, nil
}

package samples

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

func SampleRequestWithRouter2(ctx context.Context, event any) (any, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	jsonParsed, err := gabs.ParseJSON(eventJSON)
	if err != nil {
		return nil, err
	}

	path, ok := jsonParsed.Path("requestContext.http.path").Data().(string)
	if !ok {
		return nil, fmt.Errorf("received path of type %T, expected string", jsonParsed.Path("requestContext.http.path").Data())
	}

	httpMethod, ok := jsonParsed.Path("requestContext.http.method").Data().(string)
	if !ok {
		return nil, fmt.Errorf("received httpMethod of type %T, expected string", jsonParsed.Path("requestContext.http.method").Data())
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
		body, ok := jsonParsed.Path("body").Data().(string)
		if !ok {
			return nil, fmt.Errorf("received body of type %T, expected string", jsonParsed.Path("body").Data())
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

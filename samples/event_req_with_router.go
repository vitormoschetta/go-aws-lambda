package samples

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func HandlerRequestUrlWithRouter(ctx context.Context, req events.LambdaFunctionURLRequest) (any, error) {
	var res Response
	var err error
	method := req.RequestContext.HTTP.Method
	path := req.RequestContext.HTTP.Path

	if method == "GET" && path == "/healthcheck" || path == "/" || path == "" {
		res, err = healthcheck()
	}

	if method == "GET" && strings.Contains(path, "/products") {
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 3 {
			return nil, fmt.Errorf("received path %s, expected /products/<id>", path)
		}

		id, err := strconv.Atoi(pathParts[2])
		if err != nil {
			return nil, fmt.Errorf("received path %s, expected /products/<id>", path)
		}

		res, err = GetProduct(ctx, id)
	}

	if method == "POST" && strings.Contains(path, "/products") {
		var product Product
		err := json.Unmarshal([]byte(req.Body), &product)
		if err != nil {
			return nil, err
		}

		res, err = PostProduct(ctx, product)
	}

	if err != nil {
		res = Response{
			Message: err.Error(),
		}
	}

	return res, nil
}

func healthcheck() (Response, error) {
	res := Response{
		Message: "ok",
	}

	return res, nil
}

func PostProduct(ctx context.Context, product Product) (Response, error) {
	res := Response{
		Message: "ok",
		Data:    product,
	}

	return res, nil
}

func GetProduct(ctx context.Context, id int) (Response, error) {
	res := Response{
		Message: "ok",
		Data: Product{
			ID:    id,
			Name:  "Product " + strconv.Itoa(id),
			Price: 100.0,
		},
	}

	return res, nil
}

package samples

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// Neste exemplo usamos o evento LambdaFunctionURLRequest, que é um tipo de evento já definido pelo SDK do AWS Lambda para Go
// Ele tem a mesma estrutura do evento default, porem já vem parseado
func HandlerRequestUrl(ctx context.Context, req events.LambdaFunctionURLRequest) (any, error) {
	mapRes := map[string]interface{}{
		"headers": req.Headers,
		"body":    req.Body,
		"query":   req.QueryStringParameters,
		"path":    req.RequestContext.HTTP.Path,
		"method":  req.RequestContext.HTTP.Method,
		"id":      req.RequestContext.RequestID,
		"event":   req,
	}

	return mapRes, nil
}

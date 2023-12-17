package samples

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// Neste exemplo usamos o evento LambdaFunctionURLRequestContext, que é um tipo de evento já definido pelo SDK do AWS Lambda para Go
// Ele tem acesso apenas ao contexto da requisição, porem já vem parseado
func HandlerRequestContext(ctx context.Context, reqCtx events.LambdaFunctionURLRequestContext) (any, error) {
	mapRes := map[string]interface{}{
		"path":   reqCtx.HTTP.Path,
		"method": reqCtx.HTTP.Method,
		"id":     reqCtx.RequestID,
	}

	return mapRes, nil
}

package main

import (
	"go-aws-lambda/samples"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// lambda.Start(samples.HandlerRequestDefault)
	// lambda.Start(samples.HandlerRequestUrl)
	lambda.Start(samples.HandlerRequestUrlWithRouter)
	// lambda.Start(samples.HandlerRequestContext)
}

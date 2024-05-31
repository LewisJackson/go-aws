package main

import (
	"go-aws/handler"
)

func main() {
	functionName := "http-function-url-tutorial"
	payload := []byte(`{}`)

	handler.HandleLambdaInvocation(functionName, payload)
}

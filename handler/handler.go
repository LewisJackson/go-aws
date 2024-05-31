package handler

import (
	"fmt"
	"go-aws/aws"
	"log"
)

// HandleLambdaInvocation handles the invocation of a Lambda function
func HandleLambdaInvocation(functionName string, payload []byte) {
	invoker, err := aws.NewLambdaInvoker()
	if err != nil {
		log.Fatalf("failed to create Lambda invoker: %v", err)
	}

	response, err := invoker.InvokeLambda(functionName, payload)
	if err != nil {
		log.Fatalf("failed to invoke Lambda function: %v", err)
	}

	fmt.Printf("Response: %s\n", response)
}

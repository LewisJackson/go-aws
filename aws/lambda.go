package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// LambdaInvoker is a struct that holds the Lambda client
type LambdaInvoker struct {
	Client *lambda.Client
}

// NewLambdaInvoker initializes a new LambdaInvoker
func NewLambdaInvoker() (*LambdaInvoker, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		return nil, err
	}

	return &LambdaInvoker{
		Client: lambda.NewFromConfig(cfg),
	}, nil
}

// InvokeLambda invokes a Lambda function with the given name and payload
func (li *LambdaInvoker) InvokeLambda(functionName string, payload []byte) ([]byte, error) {
	resp, err := li.Client.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
	})
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

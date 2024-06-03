package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDBClient is a struct that holds the DynamoDB client
type DynamoDBClient struct {
	Client *dynamodb.Client
}

// NewDynamoDBClient initializes a new DynamoDB client
func NewDynamoDBClient() (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		return nil, err
	}

	return &DynamoDBClient{
		Client: dynamodb.NewFromConfig(cfg),
	}, nil
}

// QueryTable performs a query operation on the DynamoDB table
func (ddb *DynamoDBClient) QueryTable(tableName string, keyConditionExpression string, expressionAttributeValues map[string]types.AttributeValue) ([]map[string]types.AttributeValue, error) {
	params := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    aws.String(keyConditionExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	}

	resp, err := ddb.Client.Query(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}

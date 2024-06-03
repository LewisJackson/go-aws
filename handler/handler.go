package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"go-aws/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

// ReadDataFromDynamoDB reads data from DynamoDB table
func ReadDataFromDynamoDB(tableName string, primaryKey string) {
	ddbClient, err := aws.NewDynamoDBClient()
	if err != nil {
		log.Fatalf("failed to create DynamoDB client: %v", err)
	}

	keyConditionExpression := "PK = :pk"
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{Value: primaryKey},
	}

	items, err := ddbClient.QueryTable(tableName, keyConditionExpression, expressionAttributeValues)
	if err != nil {
		log.Fatalf("failed to query DynamoDB table: %v", err)
	}

	// Unmarshal the items
	var unmarshalledItems []map[string]interface{}
	for _, item := range items {
		var unmarshalledItem map[string]interface{}
		err := attributevalue.UnmarshalMap(item, &unmarshalledItem)
		if err != nil {
			log.Fatalf("failed to unmarshal DynamoDB item: %v", err)
		}
		unmarshalledItems = append(unmarshalledItems, unmarshalledItem)
	}

	// Marshal the items to JSON
	jsonData, err := json.Marshal(unmarshalledItems)
	if err != nil {
		log.Fatalf("failed to marshal items to JSON: %v", err)
	}

	// Print JSON data
	fmt.Println("Query results (JSON):", string(jsonData))
}

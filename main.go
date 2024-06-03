package main

import (
	"fmt"
	"go-aws/config"
	"go-aws/handler"
)

func main() {
	cfg, err := config.LoadConfig("./configs/config.toml")
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		return
	}

	functionName := cfg.AWS.TableName
	payload := []byte(`{}`)
	handler.HandleLambdaInvocation(functionName, payload)

	tableName := cfg.AWS.TableName
	primaryKey := "form"
	handler.ReadDataFromDynamoDB(tableName, primaryKey)
}

package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type handler struct {
	DynamoDBTableName string
	DynamoDBTableHash string
	DynamoDBClient    dynamodbiface.DynamoDBAPI
}

func main() {
	sess := session.Must(session.NewSession())

	h := &handler{
		DynamoDBTableName: os.Getenv("TableName"),
		DynamoDBTableHash: os.Getenv("TableHash"),
		DynamoDBClient:    dynamodb.New(sess),
	}

	lambda.Start(h.run)
}

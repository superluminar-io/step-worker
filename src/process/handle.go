package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type configuration struct {
	BatchSize int64 `json:"BatchSize"`
}

type iterator struct {
	Index  int64  `json:"Index"`
	Cursor string `json:"Cursor"`
}

type payload struct {
	Comment       string        `json:"Comment"`
	Iterator      iterator      `json:"Iterator"`
	Configuration configuration `json:"Configuration"`
}

func (h *handler) run(ctx context.Context, e payload) (iterator, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(h.DynamoDBTableName),
		Limit:     aws.Int64(e.Configuration.BatchSize),
	}

	if e.Iterator.Cursor != "" {
		input.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			h.DynamoDBTableHash: {S: aws.String(e.Iterator.Cursor)},
		}
	}

	result, err := h.DynamoDBClient.Scan(input)

	if err != nil {
		return iterator{}, err
	}

	res := iterator{e.Iterator.Index + e.Configuration.BatchSize, ""}
	if result.LastEvaluatedKey != nil {
		res.Cursor = *result.LastEvaluatedKey[h.DynamoDBTableHash].S
	}

	return res, nil
}

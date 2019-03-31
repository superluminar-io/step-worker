package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type handler struct {
	DynamoDBClient dynamodbiface.DynamoDBAPI
}

type configuration struct {
	TableName     string `json:"TableName"`
	TableHash     string `json:"TableHash"`
	BatchSize     int64  `json:"BatchSize"`
	Segment       int64  `json:"Segment"`
	TotalSegments int64  `json:"TotalSegments"`
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
	// Configure DynamoDB ScanInput
	input := &dynamodb.ScanInput{
		TableName: aws.String(e.Configuration.TableName),
		Limit:     aws.Int64(e.Configuration.BatchSize),
	}

	// Configure DynamoDB Segment for parallel Scan operations
	if e.Configuration.TotalSegments > 1 {
		input.Segment = aws.Int64(e.Configuration.Segment)
		input.TotalSegments = aws.Int64(e.Configuration.TotalSegments)
	}

	// Add cursor to continue the Scan where the previous one stopped
	if e.Iterator.Cursor != "" {
		input.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			e.Configuration.TableHash: {S: aws.String(e.Iterator.Cursor)},
		}
	}

	result, err := h.DynamoDBClient.Scan(input)

	if err != nil {
		return iterator{}, err
	}

	res := iterator{
		Index: e.Iterator.Index + e.Configuration.BatchSize,
	}

	if result.LastEvaluatedKey != nil {
		res.Cursor = *result.LastEvaluatedKey[e.Configuration.TableHash].S
	}

	return res, nil
}

func main() {
	sess := session.Must(session.NewSession())

	h := &handler{
		DynamoDBClient: dynamodb.New(sess),
	}

	lambda.Start(h.run)
}

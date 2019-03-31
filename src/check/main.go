package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

type handler struct {
	StepFunctionClient *sfn.SFN
}

type execution struct {
	Arn    string `json:"Arn"`
	Status string `json:"Status"`
}

type executionList struct {
	Status string      `json:"Status"`
	List   []execution `json:"List"`
}

type payload struct {
	Executions executionList `json:"ExecutionList"`
}

func (h *handler) run(ctx context.Context, e payload) (executionList, error) {
	countFinishedItems := 0

	for i, v := range e.Executions.List {
		// Retrieve status for StateMachine Execution
		status, err := h.StepFunctionClient.DescribeExecution(&sfn.DescribeExecutionInput{
			ExecutionArn: aws.String(v.Arn),
		})

		if err == nil {
			// Increase counter for succeeded Executions
			if "SUCCEEDED" == *status.Status {
				countFinishedItems = countFinishedItems + 1
			}

			e.Executions.List[i].Status = *status.Status
		} else {
			e.Executions.List[i].Status = "FAILURE: " + err.Error()
		}
	}

	// If all Executions succeeded, the Puppeteer succeeded
	if countFinishedItems == len(e.Executions.List) {
		e.Executions.Status = "SUCCEEDED"
	}

	return e.Executions, nil
}

func main() {
	sess := session.Must(session.NewSession())

	handler := handler{
		StepFunctionClient: sfn.New(sess),
	}

	lambda.Start(handler.run)
}

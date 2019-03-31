package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

type handler struct {
	StepFunctionClient *sfn.SFN
}

type configuration struct {
	TableName string `json:"TableName"`
	TableHash string `json:"TableHash"`
	BatchSize int    `json:"BatchSize"`
	Workers   int    `json:"Workers"`
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
	Configuration configuration `json:"Configuration"`
}

func (h *handler) run(ctx context.Context, e payload) (executionList, error) {
	executions := executionList{
		Status: "RUNNING",
		List:   []execution{},
	}

	for i := 0; i < e.Configuration.Workers; i++ {
		// State StateMachine Executions for Segment
		exec, err := h.StepFunctionClient.StartExecution(
			&sfn.StartExecutionInput{
				Input: aws.String(input(inputData{
					TableName:     e.Configuration.TableName,
					TableHash:     e.Configuration.TableHash,
					BatchSize:     e.Configuration.BatchSize,
					Segment:       i,
					TotalSegments: e.Configuration.Workers,
				})),
				StateMachineArn: aws.String(os.Getenv("MachineWorker")),
			},
		)

		if err != nil {
			return executionList{Status: "FAILED"}, err
		}

		// Assume initial RUNNING state for StateMachine Execution if no error occured
		executions.List = append(
			executions.List,
			execution{
				Arn:    *exec.ExecutionArn,
				Status: "RUNNING",
			},
		)
	}

	return executions, nil
}

func main() {
	sess := session.Must(session.NewSession())

	handler := handler{
		StepFunctionClient: sfn.New(sess),
	}

	lambda.Start(handler.run)
}

package main

import "fmt"

type inputData struct {
	TableName     string
	TableHash     string
	BatchSize     int
	Segment       int
	TotalSegments int
}

func input(conf inputData) string {
	return fmt.Sprintf(
		`{
			"Comment": "Run State Machine to process DynamoDB",
			"Configuration": {
				"TableName": "%s",
				"TableHash": "%s",
				"BatchSize": %d,
				"Segment": %d,
				"TotalSegments": %d
			}
		}`,
		conf.TableName,
		conf.TableHash,
		conf.BatchSize,
		conf.Segment,
		conf.TotalSegments,
	)
}

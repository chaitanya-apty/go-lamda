package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

//HandleRequest - GO
func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var bodyData MyEvent
	err := json.Unmarshal([]byte(req.Body), &bodyData)
	if err != nil {
		errorResponse := events.APIGatewayProxyResponse{
			Body:       "Invalid Request",
			StatusCode: 200,
		}
		return errorResponse, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       "Welcome to λ' Mr." + bodyData.Name,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

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

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, name MyEvent) (events.APIGatewayProxyResponse, error) {
	res, _ := json.Marshal(Response{
		Message: "Hi Lambda",
	})
	return events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Body:            string(res),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

// functions/hello.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is your Netlify Function handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// You can access request details like headers or query parameters
	// For this simple example, we'll just return a greeting.

	responseMessage := fmt.Sprintf("Hello, World from Go! Your path was: %s", request.Path)

	body, err := json.Marshal(map[string]string{
		"message": responseMessage,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, fmt.Errorf("failed to marshal response: %w", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

func main() {
	// This is the entry point for your Netlify Function.
	lambda.Start(Handler)
}

package lambda

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
}

func SendError(statusCode int, errorMessage string) (events.APIGatewayProxyResponse, error) {

	responseBody := Response{
		Status:  statusCode,
		Message: errorMessage,
		Success: false,
	}
	body, _ := json.Marshal(responseBody)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}

func SendValidationError(statusCode int, validationMessage string) (events.APIGatewayProxyResponse, error) {
	responseBody := Response{
		Status:  statusCode,
		Message: validationMessage,
		Success: false,
	}
	body, _ := json.Marshal(responseBody)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}

func Send(statusCode int, data any) (events.APIGatewayProxyResponse, error) {

	responseBody := Response{
		Status:  statusCode,
		Data:    data,
		Success: true,
	}

	body, _ := json.Marshal(responseBody)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}

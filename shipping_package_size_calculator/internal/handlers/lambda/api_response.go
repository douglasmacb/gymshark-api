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

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func apiResponse(status int, body any) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return resp, nil
}

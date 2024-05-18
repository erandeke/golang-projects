package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(statusCode int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = statusCode

	stringBody, err := json.Marshal(body) // this should give me the json body lets test

	if err != nil {
		return nil, err
	}

	resp.Body = string(stringBody) // not sure why do we need this
	return &resp, nil

}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	log.Printf("Request: %+v", req)
	log.Printf("Request context: %+v", ctx)

	reqClaims := req.RequestContext.Authorizer["claims"]
	log.Printf("Request claims: %+v", reqClaims)

	email := reqClaims.(map[string]interface{})["email"].(string)
	log.Printf("Request email: %+v", email)

	//reqBody := req.Body
	reqBody := email
	log.Printf("Request body: %s", reqBody)

	reqMethod := req.HTTPMethod
	log.Printf("Request method: %s", reqMethod)

	body, err := json.Marshal(map[string]interface{}{
		"message": reqBody,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "main-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

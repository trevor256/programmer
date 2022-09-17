package main

import (
    "log"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func main() {
    lambda.Start(handler)
}

func handler (request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    log.Println("hello")

    response := events.APIGatewayProxyResponse{
        statusCode: 200,
    }

    return response, nil
}


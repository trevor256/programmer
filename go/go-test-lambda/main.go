package main

import (
    "log"
    "github.com/aws/aws-lambda-go/lambda"

)

func main() {
    lambda.Start(handler)
}

func handler (request events.APIGatewayProxyRequest) () {


}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Person struct {
	Name string `json:"name"`
}

var people = []Person{{"Jamila"}, {"Alex"}}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	m, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayV2HTTPResponse {
			StatusCode: 500,
			Body:       "oops something went wrong",
		}, nil
	}
	return events.APIGatewayV2HTTPResponse {
		StatusCode: 200,
		Body:       string(m),
	}, nil
}
package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type myrequest struct {
	Name string "json:name"
	Age  int    `json:age`
}

type myresponse struct {
	Message string "json:message"
}

func hello(event myrequest) (myresponse, error) {
	
	return myresponse{ "Message" : fmt.Sprintf("%s %s", event.Name, event.Age}
}

func main() {
	lambda.Start(hello)
}

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//LiveEvent Struct
type LiveEvent struct {
	Titel string`json:"title"`
	Presentor string`json:"presentor"`
	Description string`json:"description"`
    	DateBegin string`json:"datebegin"`
	DateEnd string`json:"dateend"`
	Live bool`json:"live"`
	Featured bool`json:"featured"`
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Received body: ", request.Body)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	liveEvent := LiveEvent{
		Titel: "The Big New Movie",
		Presentor: "Jan Michalowsky",
		Description: "Test Event",
		DateBegin: "2018-30-04 12:00",
		DateEnd: "2018-30-04 12:30",
		Live: true,
		Featured: true,
	}
	
	av, err := dynamodbattribute.MarshalMap(liveEvent)

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("BrandoTable"),
	}
	
	_, err = svc.PutItem(input)
	
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 200}, nil
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return events.APIGatewayProxyResponse{Body: "Success", StatusCode: 200}, nil

}

func main() {
	lambda.Start(Handler)
}

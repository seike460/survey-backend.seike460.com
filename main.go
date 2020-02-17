package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/caarlos0/env"
	"github.com/guregu/dynamo"
	"github.com/rs/xid"
	"github.com/seike460/survey-backend.seike460.com/models"
)

type config struct {
	Env string `env:"stage" envDefault:"dev"`
}

type surveys struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(
	ctx context.Context,
	request events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse,
	error) {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return returnErrorResponse(err)
	}

	surveys := surveys{}
	//  requestJsonをMapに変換
	err := json.Unmarshal([]byte(request.Body), &surveys)
	if err != nil {
		return returnErrorResponse(err)
	}

	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table(cfg.Env + "-surveys")
	uuid := xid.New()
	u := models.Survey{UUID: uuid, Msg: surveys.Message, Time: time.Now().UTC()}

	// put item
	err = table.Put(u).Run()
	if err != nil {
		return returnErrorResponse(err)
	}

	return events.APIGatewayProxyResponse{
			Body:       "{\"result\" : \"True\" ,  \"errorMsg\" : \"\"}",
			StatusCode: 200,
		},
		nil
}

func returnErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{
			Body:       "{\"result\" : \"False\" ,  \"errorMsg\" : \"" + err.Error() + "\"}",
			StatusCode: 400,
		},
		err
}

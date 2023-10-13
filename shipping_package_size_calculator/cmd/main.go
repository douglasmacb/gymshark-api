package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/config"
	handler "github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/handlers/lambda"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/repositories"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/services"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	logger, err := logging.New()
	if err != nil {
		return fmt.Errorf("error loading log config: %s", err)
	}

	dynamoDbTableName, err := config.DynamoDbTableNameFromEnv()
	if err != nil {
		return err
	}

	dynamodbClient := repositories.NewDynamoDbClient()
	repository := repositories.New(logger, dynamodbClient, dynamoDbTableName)

	srv := services.New(logger, repository)

	lambda.Start(handler.New(logger, srv).Handler)

	return nil
}

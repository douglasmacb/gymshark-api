package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	handler "github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/handlers/lambda"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
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

	srv := services.New(logger)
	lambda.Start(handler.New(logger, srv))

	return nil
}

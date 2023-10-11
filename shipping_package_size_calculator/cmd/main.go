package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	handler "github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/handlers/lambda"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	lambda.Start(handler.New())

	return nil
}

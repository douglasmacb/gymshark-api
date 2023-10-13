package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

type DynamoDbQuery interface {
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

func NewDynamoDbClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return dynamodb.NewFromConfig(cfg, func(options *dynamodb.Options) {
		options.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		options.RetryMode = aws.RetryModeStandard
		options.RetryMaxAttempts = 3
	})
}

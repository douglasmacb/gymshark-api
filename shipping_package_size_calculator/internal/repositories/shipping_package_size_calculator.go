package repositories

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
)

const (
	shippingPackageSizePk = "SHIPPING_PACKAGE_SIZE"
)

type ShippingPackageSizeCalculator struct {
	dynamoDbApi DynamoDbQuery
	logger      logging.Logger
	tableName   string
}

func New(log logging.Logger, dynamoDbApi DynamoDbQuery, tableName string) ShippingPackageSizeCalculator {
	return ShippingPackageSizeCalculator{
		logger:      log,
		dynamoDbApi: dynamoDbApi,
		tableName:   tableName,
	}
}

func (s ShippingPackageSizeCalculator) ShippingPackagesSizes(ctx context.Context) ([]int, error) {

	keyExpr := expression.Key(tablePartitionKeyName).Equal(expression.Value(shippingPackageSizePk))
	expr, err := expression.NewBuilder().WithKeyCondition(keyExpr).Build()

	qInput := &dynamodb.QueryInput{
		TableName:                 aws.String(s.tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	response, err := s.dynamoDbApi.Query(ctx, qInput)

	if err != nil {
		return nil, fmt.Errorf("error querying table in DynamoDB, %s", err)
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("no packages sizes found from DynamoDB")
	}

	var shippingPackagesSizes []int
	err = attributevalue.UnmarshalListOfMaps(response.Items, &shippingPackagesSizes)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling DynamoDB result to shipping packages sizes, %s", err)
	}

	return shippingPackagesSizes, nil
}

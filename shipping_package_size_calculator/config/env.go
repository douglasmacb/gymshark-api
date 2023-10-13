package config

import (
	"fmt"
	"os"
)

const packageSizesEnvPropertyName = "PACKAGE_SIZES"
const dynamoDbTableNameEnvPropertyName = "DYNAMODB_TABLE_NAME"

func PackageSizesFromEnv() (string, error) {
	var env = os.Getenv(packageSizesEnvPropertyName)
	if env == "" {
		return "", fmt.Errorf("missing mandatory environment variable %s", packageSizesEnvPropertyName)
	}

	return env, nil
}

func DynamoDbTableNameFromEnv() (string, error) {
	var env = os.Getenv(dynamoDbTableNameEnvPropertyName)
	if env == "" {
		return "", fmt.Errorf("missing mandatory environment variable %s", dynamoDbTableNameEnvPropertyName)
	}

	return env, nil
}

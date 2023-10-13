package config

import (
	"os"
	"testing"
)

func TestPackageSizesFromEnv(t *testing.T) {

	tests := []struct {
		name       string
		want       string
		wantErr    bool
		beforeTest func() error
	}{
		{
			name:    "no package sizes environment variable defined, return error",
			wantErr: true,
			beforeTest: func() error {
				return os.Unsetenv(packageSizesEnvPropertyName)
			},
		},
		{
			name:    "package sizes environment variable is empty, return error",
			wantErr: true,
			beforeTest: func() error {
				return os.Unsetenv(packageSizesEnvPropertyName)
			},
		},
		{
			name:    "package sizes environment variable is defined, ok",
			want:    "[250, 500, 1000, 2000]",
			wantErr: false,
			beforeTest: func() error {
				return os.Setenv(packageSizesEnvPropertyName, "[250, 500, 1000, 2000]")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				err := tt.beforeTest()
				if err != nil {
					t.Errorf("PackageSizesFromEnv() beforeTest = %v", err)
				}
			}
			got, err := PackageSizesFromEnv()
			if (err != nil) != tt.wantErr {
				t.Errorf("PackageSizesFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PackageSizesFromEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDbTableNameFromEnv(t *testing.T) {

	tests := []struct {
		name       string
		want       string
		wantErr    bool
		beforeTest func() error
	}{
		{
			name:    "no DynamoDB table name environment variable defined, return error",
			wantErr: true,
			beforeTest: func() error {
				return os.Unsetenv(dynamoDbTableNameEnvPropertyName)
			},
		},
		{
			name:    "DynamoDB table name environment variable is empty, return error",
			wantErr: true,
			beforeTest: func() error {
				return os.Unsetenv(dynamoDbTableNameEnvPropertyName)
			},
		},
		{
			name:    "DynamoDB table name environment variable is defined, ok",
			want:    "shipping",
			wantErr: false,
			beforeTest: func() error {
				return os.Setenv(dynamoDbTableNameEnvPropertyName, "shipping")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				err := tt.beforeTest()
				if err != nil {
					t.Errorf("DynamoDbTableNameFromEnv() beforeTest = %v", err)
				}
			}
			got, err := DynamoDbTableNameFromEnv()
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDbTableNameFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DynamoDbTableNameFromEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

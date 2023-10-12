package lambda

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"reflect"
	"testing"
)

func TestShippingPackageSizeCalculator_Handler(t *testing.T) {
	type fields struct {
		logger  logging.Logger
		service Service
	}
	type args struct {
		e events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShippingPackageSizeCalculator{
				logger:  tt.fields.logger,
				service: tt.fields.service,
			}
			got, err := s.Handler(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
		})
	}
}

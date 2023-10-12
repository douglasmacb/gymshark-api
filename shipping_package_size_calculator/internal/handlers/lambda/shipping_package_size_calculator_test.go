package lambda

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
	transport "github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/transport/lambda"
	"net/http"
	"reflect"
	"testing"
)

type mockService struct {
	Service
	shippingPackageSizeCalculatorResp  []models.ShippingPackage
	ShippingPackageSizeCalculatorError error
}

var shippingPackages = []models.ShippingPackage{
	{
		NumberOfItems: 1,
		Size:          250,
		IsFull:        false,
	},
}

func (m *mockService) ShippingPackageSizeCalculator(_ models.ShippingPackageSizeCalculator) ([]models.ShippingPackage, error) {
	if m.ShippingPackageSizeCalculatorError != nil {
		return nil, m.ShippingPackageSizeCalculatorError
	}
	return m.shippingPackageSizeCalculatorResp, nil
}

func mockErrorResponse(status int, message string) []byte {
	response := transport.Response{
		Success: false,
		Status:  status,
		Message: message,
	}
	responseBytes, _ := json.Marshal(response)

	return responseBytes
}

func mockSuccessResponse(data any) []byte {
	response := transport.Response{
		Success: true,
		Status:  http.StatusOK,
		Data:    data,
	}
	responseBytes, _ := json.Marshal(response)
	return responseBytes
}

func TestShippingPackageSizeCalculator_Handler(t *testing.T) {
	log, _ := logging.New()

	unmarshallResponseError := mockErrorResponse(http.StatusInternalServerError, ErrorFailedToUnmarshalRequestBody)
	negativeNumberOfItemsResponseError := mockErrorResponse(http.StatusBadRequest, ErrorInvalidNumberOfItemsOrdered)
	internalServerResponseError := mockErrorResponse(http.StatusInternalServerError, "error")
	noCompletePacksResponseError := mockErrorResponse(http.StatusNotFound, ErrorNoCompletePackagesFound)
	successResponse := mockSuccessResponse(shippingPackages)
	requestBody, _ := json.Marshal(models.ShippingPackageSizeCalculator{NumberOfItemsOrdered: 250})
	noCompletePacksRequestBody, _ := json.Marshal(models.ShippingPackageSizeCalculator{NumberOfItemsOrdered: 1})

	type fields struct {
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
		{
			name: "shipping package size calculator handler, should return an error if unmarshall throws an error",
			fields: fields{
				service: &mockService{},
			},
			args: args{
				e: events.APIGatewayProxyRequest{
					Body: "",
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       string(unmarshallResponseError),
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			wantErr: false,
		},
		{
			name: "shipping package size calculator handler, should return an error if numberOfItemsOrdered is negative",
			fields: fields{
				service: &mockService{},
			},
			args: args{
				e: events.APIGatewayProxyRequest{
					Body: string(negativeNumberOfItemsResponseError),
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       string(negativeNumberOfItemsResponseError),
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			wantErr: false,
		},
		{
			name: "shipping package size calculator handler, should return an error if service throws an error",
			fields: fields{
				service: &mockService{
					ShippingPackageSizeCalculatorError: errors.New("error"),
				},
			},
			args: args{
				e: events.APIGatewayProxyRequest{
					Body: string(requestBody),
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       string(internalServerResponseError),
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			wantErr: false,
		},
		{
			name: "shipping package size calculator handler, should return an error if no complete packages were found",
			fields: fields{
				service: &mockService{
					shippingPackageSizeCalculatorResp: []models.ShippingPackage{},
				},
			},
			args: args{
				e: events.APIGatewayProxyRequest{
					Body: string(noCompletePacksRequestBody),
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       string(noCompletePacksResponseError),
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			wantErr: false,
		},
		{
			name: "shipping package size calculator handler, should return packages with success",
			fields: fields{
				service: &mockService{
					shippingPackageSizeCalculatorResp: []models.ShippingPackage{
						{
							NumberOfItems: 1,
							Size:          250,
							IsFull:        false,
						},
					},
				},
			},
			args: args{
				e: events.APIGatewayProxyRequest{
					Body: string(requestBody),
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       string(successResponse),
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShippingPackageSizeCalculator{
				logger:  log,
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

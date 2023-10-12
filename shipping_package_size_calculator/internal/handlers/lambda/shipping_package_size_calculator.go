package lambda

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
	"net/http"
)

const (
	InvalidFieldNumberOfItemsOrdered = "Invalid or missing value for the 'numberOfItemsOrdered' field"
)

type Service interface {
	ShippingPackageSizeCalculator(e models.ShippingPackageSizeCalculator) ([]string, error)
}

type ShippingPackageSizeCalculator struct {
	logger  logging.Logger
	service Service
}

func New(log logging.Logger, srv Service) ShippingPackageSizeCalculator {
	return ShippingPackageSizeCalculator{
		logger:  log,
		service: srv,
	}
}

func (s ShippingPackageSizeCalculator) Handler(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var request models.ShippingPackageSizeCalculator

	if err := json.Unmarshal([]byte(e.Body), &request); err != nil {
		return apiResponse(http.StatusBadRequest,
			Response{
				Data: ErrorBody{
					ErrorMsg: aws.String(InvalidFieldNumberOfItemsOrdered),
				},
				Success: false,
				Status:  http.StatusBadRequest,
			})

	}

	s.logger.Info("Handling ShippingPackageSizeCalculator event", logging.Int("numberOfItemsOrdered", request.NumberOfItemsOrdered))

	packages, err := s.service.ShippingPackageSizeCalculator(request)
	if err != nil {
		return apiResponse(http.StatusBadRequest,
			Response{
				Message: err.Error(),
				Success: false,
				Status:  http.StatusBadRequest,
			})
	}

	return apiResponse(http.StatusOK,
		Response{
			Data:    packages,
			Success: true,
			Status:  http.StatusOK,
		})
}

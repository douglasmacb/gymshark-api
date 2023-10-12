package lambda

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
	transport "github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/transport/lambda"
	"net/http"
)

const (
	ErrorInvalidNumberOfItemsOrdered  = "Invalid value for numberOfItemsOrdered: it should be a positive integer"
	ErrorFailedToUnmarshalRequestBody = "Failed to unmarshall request body"
	ErrorNoCompletePackagesFound      = "No complete packages found for the given number of items"
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
		return transport.SendError(http.StatusInternalServerError, ErrorFailedToUnmarshalRequestBody)
	}

	s.logger.Info("Handling ShippingPackageSizeCalculator event", logging.Int("numberOfItemsOrdered", request.NumberOfItemsOrdered))

	if request.NumberOfItemsOrdered <= 0 {
		return transport.SendValidationError(http.StatusBadRequest, ErrorInvalidNumberOfItemsOrdered)
	}

	packages, err := s.service.ShippingPackageSizeCalculator(request)
	if err != nil {
		return transport.SendError(http.StatusInternalServerError, err.Error())
	}

	if len(packages) == 0 {
		return transport.SendError(http.StatusNotFound, ErrorNoCompletePackagesFound)
	}

	return transport.Send(http.StatusOK, packages)
}

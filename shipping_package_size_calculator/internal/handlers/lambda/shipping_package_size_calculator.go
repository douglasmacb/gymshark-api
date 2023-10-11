package lambda

import (
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
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

func (s ShippingPackageSizeCalculator) Handler(e models.ShippingPackageSizeCalculator) ([]string, error) {

	s.logger.Info("handling ShippingPackageSizeCalculator event", logging.Int("numberOfItemsOrdered", e.NumberOfItemsOrdered))

	return []string{""}, nil
}

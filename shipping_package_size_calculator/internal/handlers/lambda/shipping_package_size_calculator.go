package lambda

import (
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
)

type ShippingPackageSizeCalculator struct {
	logger logging.Logger
}

func New(log logging.Logger) ShippingPackageSizeCalculator {
	return ShippingPackageSizeCalculator{
		logger: log,
	}
}

func (s ShippingPackageSizeCalculator) Handler(e models.ShippingPackageSizeCalculator) ([]string, error) {

	s.logger.Info("handling ShippingPackageSizeCalculator event", logging.Int("numberOfItemsOrdered", e.NumberOfItemsOrdered))

	return []string{""}, nil
}

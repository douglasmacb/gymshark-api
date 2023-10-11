package services

import (
	"errors"

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

func (s ShippingPackageSizeCalculator) ShippingPackageSizeCalculator(e models.ShippingPackageSizeCalculator) ([]string, error) {
	s.logger.Info("serving ShippingPackageSizeCalculator event", logging.Int("numberOfItemsOrdered", e.NumberOfItemsOrdered))

	if e.NumberOfItemsOrdered <= 0 {
		s.logger.Error("numberOfItemsOrdered field is not provided or it is invalid")
		return nil, errors.New("numberOfItemsOrdered field should be provided and should be higher than zero")
	}

	return []string{}, nil
}

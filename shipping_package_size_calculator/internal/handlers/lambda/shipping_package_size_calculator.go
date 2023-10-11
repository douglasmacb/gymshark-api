package lambda

import "github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"

type ShippingPackageSizeCalculator struct {
}

func New() ShippingPackageSizeCalculator {
	return ShippingPackageSizeCalculator{}
}

func (s ShippingPackageSizeCalculator) Handler(e models.ShippingPackageSizeCalculator) ([]string, error) {
	return []string{""}, nil
}

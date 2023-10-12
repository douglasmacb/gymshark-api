package models

type ShippingPackageSizeCalculator struct {
	NumberOfItemsOrdered int `json:"numberOfItemsOrdered"`
}

type ShippingPackageSizeCalculatorResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ShippingPackage struct {
	NumberOfItems int
	IsFull        bool
}

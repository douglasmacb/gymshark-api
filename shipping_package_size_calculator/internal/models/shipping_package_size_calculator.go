package models

type ShippingPackageSizeCalculator struct {
	NumberOfItemsOrdered int `json:"numberOfItemsOrdered"`
}

type ShippingPackage struct {
	NumberOfItems int
	IsFull        bool
}

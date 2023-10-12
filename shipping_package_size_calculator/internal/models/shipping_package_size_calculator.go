package models

type ShippingPackageSizeCalculator struct {
	NumberOfItemsOrdered int `json:"numberOfItemsOrdered"`
}

type ShippingPackage struct {
	Quantity int  `json:"quantity"`
	Size     int  `json:"size"`
	IsFull   bool `json:"isFull"`
}

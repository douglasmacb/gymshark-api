package models

type ShippingPackageSizeCalculator struct {
	NumberOfItemsOrdered int `json:"numberOfItemsOrdered"`
}

type ShippingPackage struct {
	NumberOfItems int  `json:"numberOfItems"`
	Size          int  `json:"size"`
	IsFull        bool `json:"isFull"`
}

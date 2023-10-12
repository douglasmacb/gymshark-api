package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/config"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
	"math"
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
	s.logger.Info("Serving ShippingPackageSizeCalculator event", logging.Int("numberOfItemsOrdered", e.NumberOfItemsOrdered))

	if e.NumberOfItemsOrdered <= 0 {
		return nil, errors.New("invalid value for numberOfItemsOrdered: it should be a positive integer")
	}

	packages, err := calculate(e)
	if err != nil {
		return nil, err
	}

	if len(packages) == 0 {
		return nil, errors.New("no complete packages found for the given number of items ordered")
	}

	return packages, nil
}

// Load package sizes from environment variables or database (e.g., DynamoDB).
func loadPackageSizes() ([]int, error) {
	packageSizesFromEnv, err := config.PackageSizesFromEnv()
	if err != nil {
		return nil, err
	}

	// TODO: In the future, consider fetching package sizes from a database.
	// For simplicity, we're using environment variables for now.
	var packageSizes []int
	if err := json.Unmarshal([]byte(packageSizesFromEnv), &packageSizes); err != nil {
		return nil, fmt.Errorf("error unmarshaling: %s", err)
	}

	return packageSizes, nil
}

func calculate(e models.ShippingPackageSizeCalculator) ([]string, error) {
	// Load package sizes from environment variables or database (e.g., DynamoDB).
	packageSizes, err := loadPackageSizes()
	if err != nil {
		return nil, err
	}

	shippingPackages := calculateShippingPackages(e.NumberOfItemsOrdered, packageSizes)
	calculatedPackagesReadyForShipping := make([]string, 0, len(shippingPackages))

	for packageSize, shippingPackage := range shippingPackages {
		// Only whole packs can be sent.
		if shippingPackage.IsFull {
			itemDescription := fmt.Sprintf("%d x %d", shippingPackage.NumberOfItems, packageSize)
			calculatedPackagesReadyForShipping = append(calculatedPackagesReadyForShipping, itemDescription)
		}
	}

	return calculatedPackagesReadyForShipping, nil
}

// calculateShippingPackages calculates the number of shipping packages needed for a given number of items and package sizes.
func calculateShippingPackages(numberOfItemsOrdered int, packageSizes []int) map[int]models.ShippingPackage {
	shippingPackages := make(map[int]models.ShippingPackage)

	remainingItems := numberOfItemsOrdered

	for remainingItems > 0 {
		// Find the nearest package size for the number of items remaining
		nearestPackageSize, nearestPackageSizeIndex := findNearestWithIndex(packageSizes, remainingItems)

		// Retrieve the shipping package for the nearest package size.
		shippingPackage := shippingPackages[nearestPackageSize]

		// Set package full if there is no space left.
		if remainingItems >= nearestPackageSize {
			shippingPackage.IsFull = true
		}

		// Update the package number of items.
		shippingPackage.NumberOfItems++

		// Store the updated shipping package back in the map.
		shippingPackages[nearestPackageSize] = shippingPackage

		// Check if additional packages of the same type are really required to prevent waste.
		currentPackageCount := shippingPackage.NumberOfItems
		if currentPackageCount > 1 {
			nextPackageSizeIndex := nearestPackageSizeIndex + 1

			if nextPackageSizeIndex < len(packageSizes) {
				nextPackageSize := packageSizes[nextPackageSizeIndex]
				preventWastingPackages(nearestPackageSize, nextPackageSize, shippingPackage.NumberOfItems, shippingPackages)
			}
		}

		// Recalculate remainingItems
		remainingItems -= nearestPackageSize
	}
	return shippingPackages
}

// Prevent wasting packages by checking if the current package size times the additional packages equals the next package size.
func preventWastingPackages(currentPackageSize int, nextPackageSize int, currentNumberOfAdditionalPackages int, shippingPackages map[int]models.ShippingPackage) {
	// Calculate the expected size of the next shipping package.
	expectedNextPackageSize := currentPackageSize * currentNumberOfAdditionalPackages

	if expectedNextPackageSize == nextPackageSize {
		// Increment the count of the next shipping package.
		nextShippingPackage := shippingPackages[nextPackageSize]
		nextShippingPackage.NumberOfItems++
		shippingPackages[nextPackageSize] = nextShippingPackage

		// Remove the current package size from the map since it's no longer needed.
		delete(shippingPackages, currentPackageSize)
	}
}

// findNearestWithIndex finds the nearest element in an integer slice based on the target value.
func findNearestWithIndex(arr []int, target int) (int, int) {
	if len(arr) == 0 {
		return -1, -1 // No elements in the array.
	}

	nearest := arr[0]
	nearestIndex := 0
	minDiff := math.Abs(float64(target - nearest))

	for i, num := range arr {
		diff := math.Abs(float64(target - num))

		if diff < minDiff {
			minDiff = diff
			nearest = num
			nearestIndex = i
		}
	}

	return nearest, nearestIndex
}

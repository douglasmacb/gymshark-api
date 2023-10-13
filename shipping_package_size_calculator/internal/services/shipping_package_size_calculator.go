package services

import (
	"context"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
	"math"
)

type Repository interface {
	ShippingPackagesSizes(ctx context.Context) ([]int, error)
}

type ShippingPackageSizeCalculator struct {
	logger     logging.Logger
	repository Repository
}

func New(log logging.Logger, repo Repository) ShippingPackageSizeCalculator {
	return ShippingPackageSizeCalculator{
		logger:     log,
		repository: repo,
	}
}

func (s ShippingPackageSizeCalculator) ShippingPackageSizeCalculator(e models.ShippingPackageSizeCalculator) ([]models.ShippingPackage, error) {
	s.logger.Info("Serving ShippingPackageSizeCalculator event", logging.Int("numberOfItemsOrdered", e.NumberOfItemsOrdered))

	packages, err := s.calculate(e)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (s ShippingPackageSizeCalculator) loadPackagesSizes() ([]int, error) {

	shippingPackagesSizes, err := s.repository.ShippingPackagesSizes(context.Background())
	if err != nil {
		return nil, err
	}

	return shippingPackagesSizes, nil
}

func (s ShippingPackageSizeCalculator) calculate(e models.ShippingPackageSizeCalculator) ([]models.ShippingPackage, error) {
	// Load package sizes from environment variables or database (e.g., DynamoDB).
	packageSizes, err := s.loadPackagesSizes()
	if err != nil {
		return nil, err
	}

	shippingPackages := calculateShippingPackages(e.NumberOfItemsOrdered, packageSizes)
	calculatedPackagesReadyForShipping := make([]models.ShippingPackage, 0, len(shippingPackages))

	for _, shippingPackage := range shippingPackages {
		// Only whole packs can be sent.
		if shippingPackage.IsFull {
			calculatedPackagesReadyForShipping = append(calculatedPackagesReadyForShipping, shippingPackage)
		}
	}

	return calculatedPackagesReadyForShipping, nil
}

// calculateShippingPackages calculates the number of shipping packages needed for a given number of items and package sizes.
// Time complexity is O(numberOfItemsOrdered * log(n))
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

		shippingPackage.Size = nearestPackageSize

		// Update the package number of items.
		shippingPackage.Quantity++

		// Store the updated shipping package back in the map.
		shippingPackages[nearestPackageSize] = shippingPackage

		// Check if additional packages of the same type are really required to prevent waste.
		currentPackageCount := shippingPackage.Quantity
		if currentPackageCount > 1 {
			nextPackageSizeIndex := nearestPackageSizeIndex + 1

			if nextPackageSizeIndex < len(packageSizes) {
				nextPackageSize := packageSizes[nextPackageSizeIndex]
				preventWastingPackages(nearestPackageSize, nextPackageSize, shippingPackage.Quantity, shippingPackages)
			}
		}

		// Recalculate remainingItems
		remainingItems -= nearestPackageSize
	}
	return shippingPackages
}

// Prevent wasting packages by checking if the current package size times the additional packages equals the next package size.
// Time complexity is constant, O(1)
func preventWastingPackages(currentPackageSize int, nextPackageSize int, currentNumberOfAdditionalPackages int, shippingPackages map[int]models.ShippingPackage) {
	// Calculate the expected size of the next shipping package.
	expectedNextPackageSize := currentPackageSize * currentNumberOfAdditionalPackages

	if expectedNextPackageSize == nextPackageSize {
		// Increment the count of the next shipping package.
		nextShippingPackage := shippingPackages[nextPackageSize]
		nextShippingPackage.Quantity++
		shippingPackages[nextPackageSize] = nextShippingPackage

		// Remove the current package size from the map since it's no longer needed.
		delete(shippingPackages, currentPackageSize)
	}
}

// findNearestWithIndex finds the nearest element in an integer slice based on the target value. Time complexity is linear, O(n)
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

package services

import (
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/logging"
	"github.com/douglasmacb/gymshark-api/shipping_package_size_calculator/internal/models"
	"os"
	"reflect"
	"testing"
)

const packageSizesEnvPropertyName = "PACKAGE_SIZES"

var packagesSizes = []int{250, 500, 1000, 2000, 5000}
var packagesSizesString = "[250, 500, 1000, 2000, 5000]"

func TestNew(t *testing.T) {
	l, _ := logging.New()

	want := New(l)

	if got := New(l); !reflect.DeepEqual(got, want) {
		t.Errorf("New() = %v, want %v", got, want)
	}
}

func Test_calculateShippingPackages(t *testing.T) {

	type args struct {
		numberOfItemsOrdered int
		packageSizes         []int
	}
	tests := []struct {
		name string
		args args
		want map[int]models.ShippingPackage
	}{
		{
			name: "calculate shipping packages, with 1 item",
			args: args{
				numberOfItemsOrdered: 1,
				packageSizes:         packagesSizes,
			},
			want: map[int]models.ShippingPackage{
				packagesSizes[0]: {
					NumberOfItems: 1,
					IsFull:        false,
				},
			},
		},
		{
			name: "calculate shipping packages, with 250 items",
			args: args{
				numberOfItemsOrdered: 250,
				packageSizes:         packagesSizes,
			},
			want: map[int]models.ShippingPackage{
				packagesSizes[0]: {
					NumberOfItems: 1,
					IsFull:        true,
				},
			},
		},
		{
			name: "calculate shipping packages, with 251 items",
			args: args{
				numberOfItemsOrdered: 251,
				packageSizes:         packagesSizes,
			},
			want: map[int]models.ShippingPackage{
				packagesSizes[1]: {
					NumberOfItems: 1,
					IsFull:        false,
				},
			},
		},
		{
			name: "calculate shipping packages, with 501 items",
			args: args{
				numberOfItemsOrdered: 501,
				packageSizes:         packagesSizes,
			},
			want: map[int]models.ShippingPackage{
				packagesSizes[0]: {
					NumberOfItems: 1,
					IsFull:        false,
				},
				packagesSizes[1]: {
					NumberOfItems: 1,
					IsFull:        true,
				},
			},
		},
		{
			name: "calculate shipping packages, with 12001 items",
			args: args{
				numberOfItemsOrdered: 12001,
				packageSizes:         packagesSizes,
			},
			want: map[int]models.ShippingPackage{
				packagesSizes[0]: {
					NumberOfItems: 1,
					IsFull:        false,
				},
				packagesSizes[3]: {
					NumberOfItems: 1,
					IsFull:        true,
				},
				packagesSizes[4]: {
					NumberOfItems: 2,
					IsFull:        true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := calculateShippingPackages(tt.args.numberOfItemsOrdered, tt.args.packageSizes)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateShippingPackages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preventWastingPackages(t *testing.T) {
	shippingPackages := make(map[int]models.ShippingPackage)

	type args struct {
		currentPackageSize                int
		nextPackageSize                   int
		currentNumberOfAdditionalPackages int
		shippingPackages                  map[int]models.ShippingPackage
	}
	tests := []struct {
		name       string
		args       args
		want       map[int]models.ShippingPackage
		beforeTest func()
	}{
		{
			name: "prevent wasting packages, ok",
			args: args{
				currentPackageSize:                packagesSizes[0],
				nextPackageSize:                   packagesSizes[1],
				currentNumberOfAdditionalPackages: 2,
				shippingPackages:                  shippingPackages,
			},
			want: map[int]models.ShippingPackage{
				packagesSizes[1]: {
					NumberOfItems: 1,
					IsFull:        false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preventWastingPackages(tt.args.currentPackageSize, tt.args.nextPackageSize, tt.args.currentNumberOfAdditionalPackages, tt.args.shippingPackages)
			if !reflect.DeepEqual(tt.args.shippingPackages, tt.want) {
				t.Errorf("preventWastingPackages() got = %v, want %v", tt.args.shippingPackages, tt.want)
			}
		})
	}
}

func Test_findNearestWithIndex(t *testing.T) {

	array := []int{300, 500, 700, 1000}

	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantIndex int
	}{
		{
			name: "find nearest with index when target is 250",
			args: args{
				arr:    array,
				target: 250,
			},
			want:      300,
			wantIndex: 0,
		},
		{
			name: "find nearest with index when target is 550",
			args: args{
				arr:    array,
				target: 550,
			},
			want:      500,
			wantIndex: 1,
		},
		{
			name: "find nearest with index when target is 950",
			args: args{
				arr:    array,
				target: 950,
			},
			want:      1000,
			wantIndex: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findNearestWithIndex(tt.args.arr, tt.args.target)
			if got != tt.want {
				t.Errorf("findNearestWithIndex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantIndex {
				t.Errorf("findNearestWithIndex() got1 = %v, want %v", got1, tt.wantIndex)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		e models.ShippingPackageSizeCalculator
	}
	tests := []struct {
		name       string
		args       args
		want       []string
		wantErr    bool
		beforeTest func()
	}{
		{
			name: "calculate with 1 item",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 1,
				},
			},
			want:    []string{},
			wantErr: false,
			beforeTest: func() {
				os.Setenv(packageSizesEnvPropertyName, packagesSizesString)
			},
		},
		{
			name: "calculate with 250 items",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 250,
				},
			},
			want:    []string{"1 x 250"},
			wantErr: false,
			beforeTest: func() {
				os.Setenv(packageSizesEnvPropertyName, packagesSizesString)
			},
		},
		{
			name: "calculate with 251 items",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 251,
				},
			},
			want:    []string{},
			wantErr: false,
			beforeTest: func() {
				os.Setenv(packageSizesEnvPropertyName, packagesSizesString)
			},
		},
		{
			name: "calculate with 501 items",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 501,
				},
			},
			want:    []string{"1 x 500"},
			wantErr: false,
			beforeTest: func() {
				os.Setenv(packageSizesEnvPropertyName, packagesSizesString)
			},
		},
		{
			name: "calculate with 12001 items",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 12001,
				},
			},
			want:    []string{"2 x 5000", "1 x 2000"},
			wantErr: false,
			beforeTest: func() {
				os.Setenv(packageSizesEnvPropertyName, packagesSizesString)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest()

			got, err := calculate(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingPackageSizeCalculator_ShippingPackageSizeCalculator(t *testing.T) {

	log, _ := logging.New()

	type args struct {
		e models.ShippingPackageSizeCalculator
	}
	tests := []struct {
		name       string
		args       args
		want       []string
		wantErr    bool
		beforeTest func()
	}{
		{
			name: "shipping package size calculator, should return an error if calculate throws an error",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 5,
				},
			},
			want:       nil,
			wantErr:    true,
			beforeTest: func() {},
		},
		{
			name: "shipping package size calculator, should return an error if there are no packages returned",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 1,
				},
			},
			want:    nil,
			wantErr: true,
			beforeTest: func() {
				os.Setenv(packageSizesEnvPropertyName, packagesSizesString)
			},
		},
		{
			name: "shipping package size calculator, should return packages with success",
			args: args{
				e: models.ShippingPackageSizeCalculator{
					NumberOfItemsOrdered: 12001,
				},
			},
			want:    []string{"2 x 5000", "1 x 2000"},
			wantErr: false,
			beforeTest: func() {
				os.Setenv(packageSizesEnvPropertyName, packagesSizesString)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShippingPackageSizeCalculator{
				logger: log,
			}

			tt.beforeTest()

			got, err := s.ShippingPackageSizeCalculator(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShippingPackageSizeCalculator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShippingPackageSizeCalculator() got = %v, want %v", got, tt.want)
			}
		})
	}
}

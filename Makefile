run-terraform:
	terraform -chdir=infra/terraform init
	terraform -chdir=infra/terraform destroy --auto-approve
	rm -fr infra/terraform/tf_generated
	terraform -chdir=infra/terraform apply --auto-approve

test:
	cd shipping_package_size_calculator && go test ./... -v

upgrade-dependencies:
	go get -t -v -u -d ./...

lint-shipping-package-size-calculator:
	cd shipping_package_size_calculator && golangci-lint run ./...

vulncheck:
	govulncheck ./...
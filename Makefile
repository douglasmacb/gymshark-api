run-terraform:
	terraform -chdir=infra/terraform init
	terraform -chdir=infra/terraform destroy --auto-approve
	rm -fr infra/terraform/tf_generated
	terraform -chdir=infra/terraform apply --auto-approve
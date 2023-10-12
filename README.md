# gymshark-api

Gymshark API

shipping-package-size-calculator lambda - Responsible for calculating the number of packs gymshark needs to ship to the customer.

## Dependencies

### Required
* Go 1.21+
* Terraform v1.6+
* AWS CLI

### Optional
* golangci-lint
* govulncheck

## Test
At the root path of the gymshark-api, run the following command to execute the unit tests
```bash
make test
```

## Configuration

The PACKAGE_SIZES environment variable is a required configuration that contains a list of package sizes. While an initial value is provided during Terraform provisioning, this approach offers flexibility by allowing modifications to package sizes without altering the implementation. As a next step, consider transitioning to a database table for managing package sizes.

## Deploy
For deploying the lambda function to AWS cloud provider, you have to create a aws profile called gymshark-aws.

```bash
aws configure --profile gymshark-aws
```

After configuring your AWS profile, you can execute the following command to deploy the function to the cloud provider, making it ready for use.

```bash
make run-terraform
```

After deployment, Terraform will provide the Lambda function endpoint:

#### Example:
lambda_function_url = "https://lyqy2yvqi6rak3edlus5ne47fi0vigqk.lambda-url.us-east-1.on.aws/"



## Usage

```bash
curl -X POST -H "Content-Type: application/json" -d '{"numberOfItemsOrdered": 12001}' https://lyqy2yvqi6rak3edlus5ne47fi0vigqk.lambda-url.us-east-1.on.aws/

RESPONSE BODY:

{
    "data": [
        "2 x 5000",
        "1 x 2000"
    ],
    "success": true,
    "status": 200
}

curl -X POST -H "Content-Type: application/json" -d '{"numberOfItemsOrdered": 2}' https://lyqy2yvqi6rak3edlus5ne47fi0vigqk.lambda-url.us-east-1.on.aws/

RESPONSE BODY:

{
    "success": false,
    "message": "no complete packages found for the given number of items ordered",
    "status": 400
}
```

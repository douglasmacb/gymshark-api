resource "aws_dynamodb_table" "shipping-table" {
  name           = local.dynamodb_table_name
  billing_mode   = "PROVISIONED"
  read_capacity  = local.dynamodb_read_capacity
  write_capacity = local.dynamodb_write_capacity
  hash_key       = local.dynamodb_hash_key
  range_key      = local.dynamodb_range_key

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "SK"
    type = "N"
  }

  tags = {
    Name        = local.dynamodb_table_name
    Environment = var.environment
  }
}
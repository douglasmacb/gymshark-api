terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    archive = {
      source = "hashicorp/archive"
    }
    null = {
      source = "hashicorp/null"
    }
  }

  required_version = ">= 1.3.7"
}

provider "aws" {
  region  = "us-east-1"
  profile = "gymshark-aws"

  default_tags {
    tags = {
      app = "gymshark-api"
    }
  }
}
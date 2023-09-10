terraform {

  cloud {
    organization = "ut080"

    workspaces {
      name = "bcs-portal"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  required_version = ">= 1.5.7"
}

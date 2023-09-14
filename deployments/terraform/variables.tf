variable "region" {
  description = "AWS region"
  default     = "us-west-2"
}

variable "terraform_cloud_role_arn" {
  description = "The ARN to use when granting permissions to the Terraform Cloud role"
}

variable "aws_principal_arn" {
  description = "Value to use as the principal for IAM users for KMS keys"
  default = ""
}

variable "amplify_repository" {
  description = "Git repository used with AWS Amplify"
  default = "https://github.com/ut080/bcs-portal"
}

variable "amplify_repository_access_token" {
  description = "The repository access token that Amplify will use to deploy the website"
  default = ""
}

variable "bcs_portal_domain" {
  description = "The domain name for the BCS portal"
  default = ""
}
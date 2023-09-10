variable "region" {
  description = "AWS region"
  default     = "us-west-2"
}

variable "amplify_repository" {
  description = "Git repository used with AWS Amplify"
  default = "https://github.com/ut080/bcs-portal"
}

variable "bcs_portal_domain" {
  description = "The domain name for the BCS portal"
  default = "blackhawkcadetsquadron.org"
}
terraform {
  required_version = ">= 0.12.0"

  backend "s3" {
    bucket         = "go-contacts-terraform"
    region         = "us-east-1"
    dynamodb_table = "go-contacts-terraformStateLock"
  }
}

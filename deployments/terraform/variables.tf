provider "aws" {
  version = "2.28.1"
  region  = var.region
}

provider "aws" {
  alias  = "dns"
  region = var.region

  assume_role {
    role_arn = "arn:aws:iam::275098837840:role/HylandIOModifyRole"
  }
}

variable "aws_profile" {
  description = "Which AWS profile is should be used? Defaults to \"default\""
  default     = "default"
}

variable "region" {
  default = "us-east-1"
}

variable "pe_dev_public_ingress_name" {
  default = "ingress-public.pe-dev.platform.hyland.io"
}
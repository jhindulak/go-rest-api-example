data "aws_route53_zone" "hylandio" {
  name     = "hyland.io."
  provider = "aws.dns"
}

data "aws_vpc" "pe-dev" {
  default = false
  state   = "available"

  tags = {
    Name = "platform-engineering-dev"
  }
}

data "aws_subnet_ids" "pe-dev-public" {
  vpc_id = data.aws_vpc.pe-dev.id

  tags = {
    Name = "*public*"
  }
}

data "aws_security_group" "pe_dev_eks_worker" {
  vpc_id = data.aws_vpc.pe-dev.id

  tags = {
    Name = "pe-dev-eks_worker_sg"
  }
}
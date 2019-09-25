// Create Route53 alias to environment
resource "aws_route53_record" "go-contacts-url" {
  name    = var.application_url
  type    = "CNAME"
  zone_id = data.aws_route53_zone.hylandio.id
  records = [var.pe_dev_public_ingress_name]
  ttl     = "300"

  provider = "aws.dns"
}

// Create database
resource "aws_security_group" "postgres-to-eks-pe-dev" {
  name = "postgres_to_eks_pe_dev"
  description = "Allows the pe-dev EKS cluster to connect to databases on port 5432"
  vpc_id = data.aws_vpc.pe-dev.id

  ingress {
    from_port = 5432
    protocol = "TCP"
    to_port = 5432
    security_groups = [data.aws_security_group.pe_dev_eks_worker.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

module "postgres_database" {
  source  = "terraform-aws-modules/rds/aws"
  version = "2.5.0"

  db_subnet_group_name = "default-${data.aws_vpc.pe-dev.id}" // Fix this, no data provider for db_subnet_groups
  parameter_group_name = "default.postgres10"
  family = "postgres10"
  allocated_storage = 5
  backup_window = "03:00-06:00"
  backup_retention_period = 0
  engine = "postgres"
  engine_version = "10.6"
  name = "gocontacts"
  username = var.database_username
  password = var.database_password
  identifier = "go-contacts-${var.env}"
  instance_class = "db.t2.micro"
  maintenance_window = "Sun:00:00-Sun:03:00"
  port = "5432"
  publicly_accessible = "false"
  vpc_security_group_ids = [aws_security_group.postgres-to-eks-pe-dev.id]

  tags = {
    Terraform = "true"
    Environment = var.env
  }
}
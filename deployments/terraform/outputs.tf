output "database_port" {
  value = module.postgres_database.this_db_instance_port
}

output "database_name" {
  value = module.postgres_database.this_db_instance_name
}

output "database_username" {
  value = module.postgres_database.this_db_instance_username
}

output "database_password" {
  value = module.postgres_database.this_db_instance_password
  sensitive = true
}

output "route53_record" {
  value = aws_route53_record.go-contacts-url.fqdn
}

output "database_address" {
  value = module.postgres_database.this_db_instance_address
}
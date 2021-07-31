output "vpc_id" {
  value = module.vpc.vpc_id
  description = "The vpc id created"
}
output "instance_private_ip_addr" {
  value = aws_instance.alation_aws_instance.private_ip
  description = "The private IP address of the EC2 server instance."
}

output "instance_public_ip_addr" {
  value = aws_instance.alation_aws_instance.public_ip
  description = "The public IP address of the EC2 server instance."
}

output "instance_name" {
  value = aws_instance.alation_aws_instance.name
  description = "The Name of the EC2 server instance."
}

output "complete_details" {
  value = aws_instance
  description = "AWS Instance complete details"
}

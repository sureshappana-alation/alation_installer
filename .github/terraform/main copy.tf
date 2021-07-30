# terraform {
#   required_providers {
#     aws = {
#       source  = "hashicorp/aws"
#       version = "~> 3.0"
#     }
#   }
# }

# provider "aws" {
#   region  = "us-east-2"
#   profile = "engineering"
# }

# # module "vpc" {
# #   source = "terraform-aws-modules/vpc/aws"

# #   name = "my-vpc"
# #   cidr = "10.36.0.0/16"

# #   azs             = ["us-east-2a"]
# #   private_subnets = ["10.36.0.0/20"]
# #   # public_subnets  = ["10.36.128.0/21", "10.36.136.0/21", "10.36.144.0/21"]

# #   # One NAT Gateway per availability zone
# #   enable_nat_gateway = true
# #   single_nat_gateway = true
# #   one_nat_gateway_per_az = false

# #   # Enable VPC Flow Logs
# #   # Cloudwatch log group and IAM role will be created
# #   # enable_flow_log                      = false
# #   # create_flow_log_cloudwatch_log_group = false
# #   # create_flow_log_cloudwatch_iam_role  = false
# #   # flow_log_max_aggregation_interval    = 60
# #   # vpc_flow_log_tags = {
# #   #   Name = "vpc-flow-logs-infra-us-east-2"
# #   # }

# #   # Enable DHCP Options
# #   enable_dhcp_options = true
# #   enable_dns_support = true
# #   enable_dns_hostnames = true

# #   # Default Tags
# #   tags = {
# #     Terraform = "true"
# #     Environment = "Infrastructure"
# #   }
# # }


# resource "aws_vpc" "my_vpc" {
#   cidr_block = "172.16.0.0/16"

#   tags = {
#     Name = "tf-example"
#   }
# }

# resource "aws_subnet" "my_subnet" {
#   vpc_id                  = aws_vpc.my_vpc.id
#   cidr_block              = "172.16.10.0/24"
#   availability_zone       = "us-east-2a"
#   map_public_ip_on_launch = "true"

#   tags = {
#     Name = "tf-example"
#   }
# }


# resource "aws_internet_gateway" "my_internet_gateway" {
#   vpc_id = aws_vpc.my_vpc.id
#   tags = {
#     Name = "tf-example"
#   }
# }

# resource "aws_route_table" "my_route_table" {
#   vpc_id = aws_vpc.my_vpc.id

#   route {
#     //associated subnet can reach everywhere
#     cidr_block = "0.0.0.0/0"
#     //CRT uses this IGW to reach internet
#     gateway_id = aws_internet_gateway.my_internet_gateway.id
#   }

#   tags = {
#     Name = "tf-example"
#   }
# }

# resource "aws_route_table_association" "prod-crta-public-subnet-1" {
#   subnet_id      = aws_subnet.my_subnet.id
#   route_table_id = aws_route_table.my_route_table.id
# }

# resource "aws_security_group" "ssh-allowed" {
#   vpc_id = aws_vpc.my_vpc.id

#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = -1
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   ingress {
#     from_port = 22
#     to_port   = 22
#     protocol  = "tcp"
#     // This means, all ip address are allowed to ssh ! 
#     // Do not do it in the production. 
#     // Put your office or home address in it!
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   //If you do not add this rule, you can not reach the NGIX  
#   # ingress {
#   #     from_port = 80
#   #     to_port = 80
#   #     protocol = "tcp"
#   #     cidr_blocks = ["0.0.0.0/0"]
#   # }
#   tags = {
#     Name = "ssh-allowed"
#   }
# }

# resource "aws_instance" "foo" {
#   ami           = var.image_id
#   instance_type = "t3.2xlarge"
#   # VPC
#   subnet_id = aws_subnet.my_subnet.id
#   # Security Group
#   vpc_security_group_ids = ["${aws_security_group.ssh-allowed.id}"]
#   # the Public SSH key
#   key_name = var.keyName
#   # # nginx installation
#   # provisioner "file" {
#   #     source = "nginx.sh"
#   #     destination = "/tmp/nginx.sh"
#   # }
#   # data drive
#   ebs_block_device {
#     device_name           = "/dev/xvda"
#     volume_size           = 100
#     volume_type           = "gp2"
#     delete_on_termination = true
#     encrypted             = true
#   }
#   tags = {
#     Name = var.instanceName
#   }
#   provisioner "remote-exec" {
#     inline = [
#       "aws configure set aws_access_key_id ${var.aws_access_key_id}",
#       "aws configure set aws_secret_access_key ${var.aws_secret_access_key}",
#       "aws s3 cp s3://unified-installer-build-pipeline-release/${var.alation_version}.tar.gz .",
#       "pwd",
#       "ls -al",
#       "tar xvzf ./${var.alation_version}.tar.gz",
#       "cd ./${var.alation_version}",
#       "chmod +x ./installer",
#       "./installer"
#     ]
#     connection {
#       user        = "ec2-user"
#       private_key = file(var.keyPath)
#       host        = self.public_ip
#     }
#   }
# }

# # resource "aws_network_interface" "foo" {
# #   subnet_id   = aws_subnet.my_subnet.id
# #   private_ips = ["172.16.10.100"]

# #   tags = {
# #     Name = "primary_network_interface"
# #   }
# # }

# # resource "aws_instance" "foo" {
# #   ami           = var.image_id #"ami-0443305dabd4be2bc" # us-east-2
# #   instance_type = "t3.2xlarge"
# #   key_name        = var.keyName

# #   network_interface {
# #     network_interface_id = aws_network_interface.foo.id
# #     device_index         = 0
# #   }
# #   tags = {
# #     Name = var.instanceName
# #   }

# #   provisioner "remote-exec" {
# #     inline = [
# #       "aws s3 cp s3://unified-installer-build-pipeline-release/${var.alation_version}.tar.gz .",
# #       "chmod +x ${var.alation_version}.tar.gz",
# #       "tar xvzf ${var.alation_version}",
# #       "cd ${var.alation_version}/installer"
# #     ]
# #     connection {
# #         type        = "ssh"
# #         user        = "ec2-user"
# #         password    = ""
# #         private_key = file(var.keyPath)
# #         host        = self.public_ip
# #     }
# #   }
# # }


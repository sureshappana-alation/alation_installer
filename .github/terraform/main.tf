terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }

  backend "s3" {
    bucket = "gh-action-tf-state"
    region = "us-east-2"
  }
}

provider "aws" {
  region = "us-east-2"
}


module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = var.instanceName
  cidr = "10.36.0.0/16"

  azs             = ["us-east-2a"]
  private_subnets = ["10.36.0.0/20"]
  public_subnets  = ["10.36.128.0/21"]

  # Internet Gateway
  create_igw = true

  # Enable DHCP Options
  enable_dhcp_options  = true
  enable_dns_support   = true
  enable_dns_hostnames = true

  # Default Tags
  tags = {
    Terraform   = "true"
    Environment = "Cloud Infrastructure"
    Owner       = var.instanceName
  }
}

# Create Security Group
resource "aws_security_group" "ssh-allowed" {
  vpc_id = module.vpc.vpc_id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = var.instanceName
  }
}

# Finally, create AWS instance
resource "aws_instance" "alation_aws_instance" {
  ami           = var.image_id
  instance_type = "t3.2xlarge"
  # VPC
  subnet_id = module.vpc.public_subnets[0]
  #   # Security Group
  vpc_security_group_ids = ["${aws_security_group.ssh-allowed.id}"]
  # the Public SSH key
  key_name = var.keyName

  # data drive
  ebs_block_device {
    device_name           = "/dev/xvda"
    volume_size           = 100
    volume_type           = "gp2"
    delete_on_termination = true
    encrypted             = true
  }

  provisioner "remote-exec" {
    inline = [
        "aws configure set aws_access_key_id ${var.aws_access_key_id}",
        "aws configure set aws_secret_access_key ${var.aws_secret_access_key}",
        "aws s3 cp s3://unified-installer-build-pipeline-release/${var.alation_version}.tar.gz .",
        "tar xvzf ./${var.alation_version}.tar.gz",
        "cd ./${var.alation_version}",
        "chmod +x ./installer",
        "./installer"
    ]
    connection {
      user        = "ec2-user"
      private_key = file(var.keyPath)
      host        = self.public_ip
    }
  }
  tags = {
    Name = var.instanceName
  }
}

variable "image_id" {
  type = string
}
variable "alation_version" {
  type = string,
  default="alation-k8s-master-20210729.30"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "us-east-2"
  profile = "engineering"
}

resource "aws_vpc" "my_vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "tf-example"
  }
}

resource "aws_subnet" "my_subnet" {
  vpc_id            = aws_vpc.my_vpc.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = "us-east-2a"

  tags = {
    Name = "tf-example"
  }
}

resource "aws_network_interface" "foo" {
  subnet_id   = aws_subnet.my_subnet.id
  private_ips = ["172.16.10.100"]

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_instance" "foo" {
  ami           = var.image_id #"ami-0443305dabd4be2bc" # us-east-2
  instance_type = "t3.2xlarge"

  network_interface {
    network_interface_id = aws_network_interface.foo.id
    device_index         = 0
  }
  tags = {
    Name = "Foo"
  }

  provisioner "remote-exec" {
    inline = [
      "aws s3 cp s3://unified-installer-build-pipeline-release/${var.alation_version}.tar.gz .",
      "chmod +x ${var.alation_version}.tar.gz",
      "tar xvzf ${var.alation_version}",
      "cd ${var.alation_version}/installer"
    ]
  }
}


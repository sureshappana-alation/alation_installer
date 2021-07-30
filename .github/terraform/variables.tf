variable "image_id" {
  type = string
}

variable "alation_version" {
  type = string
}

variable "keyName" {
  default = "gh-actions-terraform"
}

variable "keyPath" {
  default = "~/.ssh/gh-actions-terraform.pem"
}

variable "instanceName" {
  default = "my new instance"
}

variable "aws_access_key_id" {
  default = "<PUT IN YOUR INSTANCE NAME>"
}

variable "aws_secret_access_key" {
  default = "<PUT IN YOUR INSTANCE NAME>"
}

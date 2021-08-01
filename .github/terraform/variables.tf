variable "image_id" {
  type = string
}

variable "alation_version" {
  type = string
}

variable "keyName" {
  type = string  
  default = "gh-actions-terraform"
}

variable "keyPath" {
  type = string  
  default = "~/.ssh/gh-actions-terraform.pem"
}

variable "instanceName" {
  type = string  
  default = "my new instance"
}

variable "aws_access_key_id" {
  type = string  
  default = "<PUT IN YOUR INSTANCE NAME>"
}

variable "aws_secret_access_key" {
  type    = string
  default = "<PUT IN YOUR INSTANCE NAME>"
}

variable "build_download_url" {
  type    = string
  default = "download url"
}

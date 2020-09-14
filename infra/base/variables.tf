variable "public_key" {}
variable "private_key" {}
variable "ssh_fingerprint" {}

variable "do_token" {
  description = "This is digitalocean apikey that will be given in runtime with terraform apply -var \"do_token=abc..\""
}

variable "image" {
  description = "Image of instance(droplet)"
  default     = "ubuntu-18-04-x64"
}

variable "size" {
  description = "Size of instance(droplet)"
  default     = "s-1vcpu-1gb"
}


variable "region" {
  description = "Region of instance(droplet)"
  default     = "AMS3"
}

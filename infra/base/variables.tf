
variable "do_token" {
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

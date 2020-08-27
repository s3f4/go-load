resource "digitalocean_droplet" "workers" {
  image  = var.os
  name   = "Worker-22"
  region = "AMS3"
  size   = "1GB"

  ssh_keys = [
    var.ssh_fingerprint
  ]
}
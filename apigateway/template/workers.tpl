resource "digitalocean_droplet" "workers" {
  image  = var.os
  name   = "Worker-{{.index}}"
  region = "{{.region}}"
  size   = "{{.size}}"

  ssh_keys = [
    var.ssh_fingerprint
  ]
}
resource "digitalocean_droplet" "workers" {
  image  = var.os
  name   = "Worker-{{.Index}}"
  region = {{.Region}}
  size   = {{.Size}}

  ssh_keys = [
    var.ssh_fingerprint
  ]
}
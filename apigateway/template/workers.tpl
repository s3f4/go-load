resource "digitalocean_droplet" "workers" {
  count  = {{.Count}}
  image  = var.os
  name   = "Worker-${count.index + 1}"

  region = "{{.Region}}"
  size   = "{{.Size}}"

  ssh_keys = [
    var.ssh_fingerprint
  ]
}
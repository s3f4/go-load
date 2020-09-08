resource "digitalocean_droplet" "workers" {
  count  = {{.Count}}
  image  = "{{.Image}}"
  name   = "Worker-${count.index + 1}"

  region = "{{.Region}}"
  size   = "{{.Size}}"

  ssh_keys = [
    "abc"
  ]
}
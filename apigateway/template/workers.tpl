data "digitalocean_ssh_key" "for_master" {
  name = "id_rsa_for_master"
}

resource "digitalocean_droplet" "workers" {
  count  = {{.Count}}
  image  = "{{.Image}}"
  name   = "Worker-${count.index + 1}"

  region = "{{.Region}}"
  size   = "{{.Size}}"

  ssh_keys = [
    data.digitalocean_ssh_key.for_master.id
  ]
}

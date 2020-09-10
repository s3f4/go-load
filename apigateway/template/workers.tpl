# Create a new SSH key
resource "digitalocean_ssh_key" "master" {
  name       = "Master instance ssh key"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "digitalocean_droplet" "workers" {
  count  = {{.Count}}
  image  = "{{.Image}}"
  name   = "Worker-${count.index + 1}"

  region = "{{.Region}}"
  size   = "{{.Size}}"

  ssh_keys = [
    digitalocean_ssh_key.master.fingerprint
  ]
}
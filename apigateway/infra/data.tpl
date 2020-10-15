data "digitalocean_regions" "available" {
  filter {
    key    = "available"
    values = ["true"]
  }
}

data "digitalocean_account" "account" {
}

{{ if ne .env "development" }}
data "digitalocean_ssh_key" "for_master" {
	name = "id_rsa_for_master"
}
{{ else }}
resource "digitalocean_ssh_key" "for_master" {
	name       = "id_rsa_for_master"
	public_key = file("~/.ssh/id_rsa.pub")
}
data "digitalocean_ssh_key" "for_master" {
  depends_on = [digitalocean_ssh_key.for_master]
  name       = "id_rsa_for_master"
}
{{ end }}
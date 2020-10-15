data "digitalocean_regions" "available" {
  filter {
    key    = "available"
    values = ["true"]
  }
}

data "digitalocean_ssh_key" "for_master" {
  name = "id_rsa_for_master"
}

data "digitalocean_account" "account" {
}

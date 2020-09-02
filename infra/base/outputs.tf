output "master_ipv4_address" {
  value = "${digitalocean_droplet.master.ipv4_address}"
}

output "data_ipv4_address"{
  value = "${digitalocean_droplet.data.ipv4_address}"
}
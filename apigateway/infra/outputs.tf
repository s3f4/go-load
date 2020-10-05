output "regions" {
  value = data.digitalocean_regions.available
}

output "account"{
  value = data.digitalocean_account.account
}

output "workers" {
  value = "${digitalocean_droplet.workers}"
}

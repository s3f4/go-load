output "regions" {
  value = data.digitalocean_regions.available
}

output "workers" {
  value = "${digitalocean_droplet.workers}"
}

output "workers" {
  value = "${digitalocean_droplet.workers}"
}

output "regions" {
  value = data.digitalocean_regions.available
}

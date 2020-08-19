# Master instance for docker swarm
resource "digitalocean_droplet" "master"{
    image = var.os
    name = "Master"
    region = var.region
    size = var.size
}
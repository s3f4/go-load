data "digitalocean_regions" "available" {
  filter {
    key    = "available"
    values = ["true"]
  }
} 

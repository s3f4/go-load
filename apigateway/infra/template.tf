resource "local_file" "inventory" {
  content = templatefile("inventory.tmpl", {
    workers = join("\n", digitalocean_droplet.workers.*.ipv4_address)
  })
  filename = "./ansible/inventory.txt"
}

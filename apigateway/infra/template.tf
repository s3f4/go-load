resource "local_file" "inventory" {
  content = templatefile("./ansible/inventory.tmpl", {
    workers = join("\n", digitalocean_droplet.workers.*.ipv4_address)
  })
  filename = "./ansible/inventory.txt"
}

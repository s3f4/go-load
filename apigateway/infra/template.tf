resource "local_file" "inventory" {
  content = templatefile("ansible/inventory.tmpl", {
    workers   = "${digitalocean_droplet.workers.ipv4_address}"
  })
  filename = "inventory.txt"
}

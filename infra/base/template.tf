resource "local_file" "Inventory" {
  content = templatefile("ansible/inventory.tmpl", {
    master = "${digitalocean_droplet.master.ipv4_address}"
    data   = "${digitalocean_droplet.data.ipv4_address}"
  })
  filename = "inventory.txt"
}

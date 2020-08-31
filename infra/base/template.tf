resource "local_file" "inventory" {
  content = templatefile("ansible/inventory.tmpl", {
    master = "${digitalocean_droplet.master.ipv4_address}"
    data   = "${digitalocean_droplet.data.ipv4_address}"
  })
  filename = "inventory.txt"
}

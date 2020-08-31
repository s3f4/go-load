# Master instance for docker swarm
resource "digitalocean_droplet" "master" {
  image     = var.os
  name      = "Master"
  region    = var.region
  size      = var.size

  ssh_keys = [
    var.ssh_fingerprint
  ]

  connection {
    host        = self.ipv4_address
    user        = "root"
    type        = "ssh"
    private_key = file(var.private_key)
    timeout     = "2m"
  }

  ## Install ansible to master instance
  provisioner "file" {
    source      = "../scripts/install-ansible.sh"
    destination = "/tmp/install-ansible.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/install-ansible.sh",
      "/tmp/install-ansible.sh",
    ]
  }

  ## Upload inventory file to master instance
  # provisioner "file" {
  #   source      = file(local_file.inventory.filename)
  #   destination = "/etc/inventory.txt"
  # }

  ## Install terraform to master instance
  provisioner "file" {
    source      = "../scripts/install-terraform.sh"
    destination = "/tmp/install-terraform.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/install-terraform.sh",
      "/tmp/install-terraform.sh",
    ]
  }
}

# Data instance for docker swarm
resource "digitalocean_droplet" "data" {
  image  = var.os
  name   = "Data"
  region = var.region
  size   = var.size
  ssh_keys = [
    var.ssh_fingerprint
  ]
}

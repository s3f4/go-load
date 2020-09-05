# Master instance for docker swarm

# Create a new SSH key
resource "digitalocean_ssh_key" "for_master" {
  name       = "id_rsa_for_master"
  public_key = file("~/.ssh/id_rsa_for_master.pub")
}

resource "digitalocean_droplet" "master" {
  image  = var.os
  name   = "Master"
  region = var.region
  size   = var.size

  ssh_keys = [
    var.ssh_fingerprint,
    digitalocean_ssh_key.for_master.fingerprint
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

  // upload playbook
  provisioner "file" {
    source      = "ansible/docker-playbook.yml"
    destination = "/etc/ansible/docker-playbook.yml"
  }

  ## Upload ssh private key for ansible master 
  provisioner "file" {
    source      = "~/.ssh/id_rsa_for_master"
    destination = "~/.ssh/id_rsa"
  }

  ## Upload ssh public key for ansible master 
  provisioner "file" {
    source      = "~/.ssh/id_rsa_for_master.pub"
    destination = "~/.ssh/id_rsa.pub"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod 600 ~/.ssh/id_rsa",
      "chmod 644 ~/.ssh/id_rsa.pub",
    ]
  }
}



# Create a new Droplet using the SSH key
# Data instance for docker swarm
resource "digitalocean_droplet" "data" {
  image  = var.os
  name   = "Data"
  region = var.region
  size   = var.size
  ssh_keys = [
    var.ssh_fingerprint,
    digitalocean_ssh_key.for_master.fingerprint
  ]
}

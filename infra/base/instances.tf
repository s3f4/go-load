# Master instance for docker swarm

resource "digitalocean_ssh_key" "ssh" {
  name       = "id_rsa"
  public_key = file("~/.ssh/id_rsa.pub")
}

# Create a new SSH key
resource "digitalocean_ssh_key" "for_master" {
  name       = "id_rsa_for_master"
  public_key = file("~/.ssh/id_rsa_for_master.pub")
}

resource "digitalocean_droplet" "master" {
  image  = var.image
  name   = "Master"
  region = var.region
  size   = var.size

  ssh_keys = [
    digitalocean_ssh_key.ssh.fingerprint,
    digitalocean_ssh_key.for_master.fingerprint
  ]

  connection {
    host        = self.ipv4_address
    user        = "root"
    type        = "ssh"
    private_key = file("~/.ssh/id_rsa")
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

  // upload playbook
  provisioner "file" {
    source      = "ansible/hosts.yml"
    destination = "/etc/ansible/hosts.yml"
  }

  provisioner "file" {
    source      = "ansible/known_hosts.yml"
    destination = "/etc/ansible/known_hosts.yml"
  }

   provisioner "file" {
    source      = "ansible/label.yml"
    destination = "/etc/ansible/label.yml"
  }

  provisioner "file" {
    source      = "ansible/cert.yml"
    destination = "/etc/ansible/cert.yml"
  }

  provisioner "file" {
    source      = "ansible/docker-playbook.yml"
    destination = "/etc/ansible/docker-playbook.yml"
  }

  provisioner "file" {
    source      = "ansible/swarm-init-deploy.yml"
    destination = "/etc/ansible/swarm-init-deploy.yml"
  }

  provisioner "file" {
    source      = "ansible/swarm-join.yml"
    destination = "/etc/ansible/swarm-join.yml"
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
      "mkdir ~/app"
    ]
  }

  provisioner "file" {
    source      = abspath("../../apigateway")
    destination = "~/app/apigateway"
  }

  provisioner "file" {
    source      = abspath("../../eventhandler")
    destination = "~/app/eventhandler"
  }


  provisioner "file" {
    source      = abspath("../../web")
    destination = "~/app/web"
  }


  provisioner "file" {
    source      = abspath("../../worker")
    destination = "~/app/worker"
  }

  provisioner "file" {
    source      = abspath("../../swarm-prod.yml")
    destination = "~/app/swarm-prod.yml"
  }

  provisioner "file" {
    source      = abspath("../../build.sh")
    destination = "~/app/build.sh"
  }
}



# Create a new Droplet using the SSH key
# Data instance for docker swarm
resource "digitalocean_droplet" "data" {
  image  = var.image
  name   = "Data"
  region = var.region
  size   = var.size
  ssh_keys = [
    digitalocean_ssh_key.ssh.fingerprint,
    digitalocean_ssh_key.for_master.fingerprint
  ]
}

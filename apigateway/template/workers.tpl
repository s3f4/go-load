variable regions {
  default = [
    {{- range $index,$instance := .Instances }}
    {{ $instance -}}
    {{- end }}
  ]
}

locals {
  regions = { for r in var.regions :
    r.index => r
  }
}

resource "digitalocean_droplet" "workers" {
  for_each = local.regions
  name     = "worker-${each.value.reg}-${each.value.instance_number}"
  region   = each.value.reg
  size     = "s-1vcpu-1gb"
  image    = "ubuntu-18-04-x64"

  ssh_keys = [
	  {{ if ne .env "development" }}data.{{ end }}digitalocean_ssh_key.for_master.id
  ]

  connection {
    host        = self.ipv4_address
    user        = "root"
    type        = "ssh"
    private_key = file("~/.ssh/id_rsa")
    timeout     = "2m"
  }

  provisioner "remote-exec" {
    inline = [
      "mkdir -p ~/app/worker/log",
    ]
  }
}
variable regions {
  default = [
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
	  data.digitalocean_ssh_key.for_master.id
  ]
}
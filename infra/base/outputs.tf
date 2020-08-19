output "master_ipv4_address"{
   value = "${digitalocean_droplet.master.ipv4_address}" 
}

output "master_ipv4_name"{
   value = "${digitalocean_droplet.master.name}" 
}
#!/bin/bash

#set -euo pipefail

apt update
apt install unzip
wget https://releases.hashicorp.com/terraform/0.13.1/terraform_0.13.1_linux_amd64.zip
unzip terraform_0.13.1_linux_amd64.zip
mv terraform /usr/bin/

terraform --version
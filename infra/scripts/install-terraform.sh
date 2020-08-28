#!/bin/bash

#set -euo pipefail

sudo apt update
sudo apt-get install unzip
wget https://releases.hashicorp.com/terraform/0.13.1/terraform_0.13.1_linux_amd64.zip
unzip terraform_0.13.1_linux_amd64.zip
sudo mv terraform /usr/local/bin/

terraform --version
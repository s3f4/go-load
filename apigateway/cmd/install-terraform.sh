#!/bin/sh

apk update
apk add unzip
wget https://releases.hashicorp.com/terraform/0.14.4/terraform_0.14.4_linux_amd64.zip
unzip terraform_0.14.4_linux_amd64.zip
mv terraform /usr/bin/

terraform --version
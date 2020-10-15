#!/bin/sh

#set -euo pipefail

apk update
#apk add --yes software-properties-common
#apt-add-repository --yes --update ppa:ansible/ansible
#apt install --yes ansible
apk add ansible
apk add python3
ansible --version
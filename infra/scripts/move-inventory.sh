#!/bin/bash

cd ../base
master=$(terraform output master_ipv4_address)
scp inventory.txt root@$master:/etc/ansible/inventory.txt
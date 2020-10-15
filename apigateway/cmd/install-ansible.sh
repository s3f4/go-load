#!/bin/sh

apk update
apk add ansible
apk add python3
ansible --version
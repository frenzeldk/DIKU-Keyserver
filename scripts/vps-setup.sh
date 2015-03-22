#!/bin/bash
# User creation is not documented in this shell script.
# The set-up of the ssh-daemon and password-less elevation is not documented either.

# Installation of golang and sqlite
sudo apt-get -y update
sudo apt-get -y upgrade
sudo apt-get -y install golang sqlite


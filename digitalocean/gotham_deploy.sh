#!/usr/bin/env bash

# This script builds, deploys and installs the gotham API onto DigitalOcean.

set -e

pushd /vagrant
./gotham_build.sh
./gotham_package.sh
popd

pushd /vagrant/inst
ssh -i /ssh_keys/digital_ocean root@162.243.69.92 'rm -rf ~/gotham_pkg && mkdir ~/gotham_pkg'
scp -Cri /ssh_keys/digital_ocean ./ root@162.243.69.92:~/gotham_pkg
ssh -i /ssh_keys/digital_ocean root@162.243.69.92 'cd ~/gotham_pkg/ && sudo ./gotham_install.sh && sudo start gotham'
popd

#!/usr/bin/env bash

# This script builds, deploys and installs the gotham API onto DigitalOcean.

pushd /vagrant/src/gotham
go install
popd

scp -Cri /ssh_keys/digital_ocean /vagrant/scripts/gotham_config /vagrant/scripts/gotham_install.sh /vagrant/bin/gotham root@162.243.69.92:~/
ssh -i /ssh_keys/digital_ocean root@162.243.69.92 '~/gotham_install.sh'

#!/usr/bin/env bash

# This script builds, deploys and installs the gotham API onto DigitalOcean.

pushd /vagrant/src/gotham
go install
popd

scp -C -i /ssh_keys/digital_ocean /vagrant/scripts/install.sh /vagrant/bin/gotham root@162.243.69.92:~/
ssh -i /ssh_keys/digital_ocean root@162.243.69.92 '~/install.sh'
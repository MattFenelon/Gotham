#!/usr/bin/env bash

wget --no-verbose https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.2.linux-amd64.tar.gz

# Add Go to path
sudo sh -c "echo 'export PATH=$PATH:/usr/local/go/bin' > /etc/profile.d/gopath.sh"
# Configure GOPATH for our shared folder
sudo sh -c "echo 'export GOPATH=/vagrant' >> /etc/profile.d/gopath.sh"
sudo chmod a+x /etc/profile.d/gopath.sh
source /etc/profile.d/gopath.sh
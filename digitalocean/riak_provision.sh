#!/usr/bin/env bash

scp -Cri /ssh_keys/digital_ocean /vagrant/env/riak root@162.243.69.92:~/
ssh -i /ssh_keys/digital_ocean root@162.243.69.92 'cd ~/riak && ./riak_install.sh ./config/'
#!/usr/bin/env bash

scp -Cri /ssh_keys/digital_ocean /vagrant/scripts/riak_install.sh /vagrant/scripts/riak_config root@162.243.69.92:~/
ssh -i /ssh_keys/digital_ocean root@162.243.69.92 '~/riak_install.sh ./riak_config/'
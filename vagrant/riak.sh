#!/usr/bin/env bash

wget --no-verbose http://s3.amazonaws.com/downloads.basho.com/riak/1.4/1.4.2/ubuntu/precise/riak_1.4.2-1_amd64.deb
sudo dpkg -i riak_1.4.2-1_amd64.deb

rsync -rv /vagrant/vagrant/fs/ /
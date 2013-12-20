#!/usr/bin/env bash

set -e

adduser --disabled-password --gecos "" mtf

mkdir /home/mtf/.ssh -p
chmod 777 /home/mtf/.ssh
cat mtf.pub > /home/mtf/.ssh/authorized_keys
chmod 660 /home/mtf/.ssh
chmod 600 /home/mtf/.ssh/authorized_keys
chown -R mtf:mtf /home/mtf/.ssh
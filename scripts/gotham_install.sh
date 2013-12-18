#!/usr/bin/env bash

# This script installs the gotham API on the remote host.
# It is expected that the compiled API is contained in the same
# location as this script.

pkill gotham
rm -Rf /usr/local/gotham

mkdir /usr/local/gotham
cp gotham /usr/local/gotham
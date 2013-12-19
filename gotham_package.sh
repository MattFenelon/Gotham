#!/usr/bin/env bash

# Cause any non-0 exit code to fail the script immediately
set -e

outdir=./inst
gothamOut=$outdir/usr/local/gotham

rm -rf $outdir
mkdir $gothamOut -p

# Copy Gotham API executable
cp ./bin/gotham $gothamOut -v

# Copy extra installation files
rsync -rv ./etc/gotham/ $outdir

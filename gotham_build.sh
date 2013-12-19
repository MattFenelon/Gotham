#!/usr/bin/env bash

# Cause any non-0 exit code to fail the script immediately
set -e

go test ./src/domain_tests ./src/persistence/filestore_tests ./src/persistence/riak_tests ./src/http_tests
go install ./src/gotham
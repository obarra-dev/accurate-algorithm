#!/usr/bin/env bash

set -e

ls -l

cd /go/alggo
echo "Running tests in alggo directory"
ls -l

echo "The go version is:"
go version

go test

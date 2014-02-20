#!/bin/bash
#
# This script builds the application from source.
set -e

# If we're building on Windows, specify an extension
EXTENSION=""
if [ "$(go env GOOS)" = "windows" ]; then
    EXTENSION=".exe"
fi

# Install dependencies
echo "--> Installing dependencies to speed up builds..."
go get -u ./...

# Build!
echo "--> Building percherond"
go build \
    -v \
    -o dist/percherond${EXTENSION} \
    percherond/*.go

# Build!
echo "--> Building perchauthd"
go build \
    -v \
    -o dist/perchauthd${EXTENSION} \
    perchauthd/*.go

# Build!
echo "--> Building perchcli"
go build \
    -v \
    -o dist/perchcli${EXTENSION} \
    perchcli/*.go

# Build!
echo "--> Building perchlookupd"
go build \
    -v \
    -o dist/perchlookupd${EXTENSION} \
    perchlookupd/*.go

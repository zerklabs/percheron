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
go get ./...

# Build!
echo "--> Building percherond"
go build \
    -v \
    -o dist/percherond${EXTENSION} \
    $(percheron/*.go percherond/*.go)

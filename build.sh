#!/bin/bash
#
# This script builds the application from source.
# If we're building on Windows, specify an extension
EXTENSION=""
DISTPATH="dist/"
if [ "$(go env GOOS)" = "windows" ]; then
    EXTENSION=".exe"
    DISTPATH="dist\\"
fi

# Build!
echo "--> Building percherond"
go build -v -o ${DISTPATH}percherond${EXTENSION} percheron/percherond

# Build!
echo "--> Building perchauthd"
go build -v -o ${DISTPATH}perchauthd${EXTENSION} percheron/perchauthd

# Build!
echo "--> Building perch-cli"
go build -v -o ${DISTPATH}perchcli${EXTENSION} percheron/perch-cli

# Build!
echo "--> Building perchlookupd"
go build -v -o ${DISTPATH}perchlookupd${EXTENSION} percheron/perchlookupd

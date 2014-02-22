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


echo "--> Building percheron"
go build -v

# Build!
echo "--> Building percherond"
go build -v -o dist/percherond${EXTENSION} github.com/zerklabs/percheron/percherond

# Build!
echo "--> Building perchauthd"
go build -v -o dist/perchauthd${EXTENSION} github.com/zerklabs/percheron/perchauthd

# Build!
echo "--> Building perch-cli"
go build -v -o dist/perchcli${EXTENSION} github.com/zerklabs/percheron/perch-cli

# Build!
echo "--> Building perchlookupd"
go build -v -o dist/perchlookupd${EXTENSION} github.com/zerklabs/percheron/perchlookupd

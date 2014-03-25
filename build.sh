#!/bin/bash

EXTENSION=""
ARCH="$(go env GOHOSTARCH)"
DISTPATH="bin/linux-${ARCH}"
if [ "$(go env GOOS)" = "windows" ]; then
    EXTENSION=".exe"
    DISTPATH="bin\\windows-${ARCH}"
fi

echo "--> Building percheron"
go build -v

echo "--> Building percherond"
go build -v -o $DISTPATH/percherond${EXTENSION} github.com/zerklabs/percheron/percherond

echo "--> Building perchauthd"
go build -v -o $DISTPATH/perchauthd${EXTENSION} github.com/zerklabs/percheron/perchauthd

echo "--> Building perchlookupd"
go build -v -o $DISTPATH/perchlookupd${EXTENSION} github.com/zerklabs/percheron/perchlookupd

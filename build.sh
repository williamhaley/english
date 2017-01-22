#!/usr/bin/env bash

export GOPATH=`pwd`
export GOBIN=$GOPATH/bin

# download and install packages and dependencies
# go get

# Download, then install packages.
go get -d
go install

go fmt


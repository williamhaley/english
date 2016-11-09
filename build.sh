#!/usr/bin/env bash

export GOPATH=`pwd`
export GOBIN=$GOPATH/bin

go get

go fmt


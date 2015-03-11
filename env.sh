#!/usr/bin/env bash
GOROOT=`go env GOROOT`
export GOROOT
unset GOPATH
GOPATH=`pwd`
export GOPATH
GOBIN=$GOPATH/bin
export GOBIN
PATH=$GOPATH/bin:$PATH
export PATH

#/bin/bash

set -eux

export GOPATH=$(pwd)

gopath_is_defined=$GOPATH

go install -v all

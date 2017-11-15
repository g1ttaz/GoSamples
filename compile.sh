#!/bin/sh

PACKAGE_PREFIX="github.com/g1ttaz/GoSamples"
SUBPACKAGES=`ls -d exercise-*`

for pkg in ${SUBPACKAGES}; do
    go build -o ${GOPATH}/bin/$pkg ${PACKAGE_PREFIX}/$pkg
done

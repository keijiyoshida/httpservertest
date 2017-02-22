#!/bin/bash

VERSION=$1

if [ -e ./out ]; then
    rm -rf ./out
fi

mkdir out

docker run --rm -it -v "$PWD":/usr/local/go/src/github.com/keijiyoshida/httpservertest -w /usr/local/go/src/github.com/keijiyoshida/httpservertest golang:1.8 ./cross_compile_internal.sh $VERSION

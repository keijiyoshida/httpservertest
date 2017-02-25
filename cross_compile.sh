#!/bin/bash

VERSION=$1

if [ -e ./out ]; then
    rm -rf ./out
fi

mkdir out

for GOOS in darwin linux; do
    for GOARCH in 386 amd64; do
        mkdir out/httpservertest-$VERSION-$GOOS-$GOARCH
        docker run --rm -it -v "$PWD":/usr/local/go/src/github.com/keijiyoshida/httpservertest -w /usr/local/go/src/github.com/keijiyoshida/httpservertest golang:1.8 /bin/bash -c "GOOS=$GOOS GOARCH=$GOARC go build -o out/httpservertest-$VERSION-$GOOS-$GOARCH/httpservertest cmd/httpservertest/main.go"
        tar -zcf out/httpservertest-$VERSION-$GOOS-$GOARCH.tar.gz -C out httpservertest-$VERSION-$GOOS-$GOARCH
        rm -rf out/httpservertest-$VERSION-$GOOS-$GOARCH
    done
done

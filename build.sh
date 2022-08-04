#!/bin/bash

version="0.0.3"

GOOS=linux GOARCH=amd64 go build -o ww
tar -zcvf workwork_${version}_linux_x86_64.tar.gz ww
rm ww

GOOS=windows GOARCH=amd64 go build -o ww
tar -zcvf workwork_${version}_windows.tar.gz ww
rm ww

GOOS=darwin GOARCH=amd64 go build -o ww
tar -zcvf workwork_${version}_darwin.tar.gz ww

shasum -a 256 workwork_${version}_darwin.tar.gz


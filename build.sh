#!/bin/bash

# Build
[ -d ../bin ] || mkdir ../bin
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o ../bin/workspace-linux-amd64 .
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -v -o ../bin/workspace-darwin-amd64 .

# Compress
cd ../bin
cp workspace-linux-amd64 workspace-linux-amd64-upx
cp workspace-darwin-amd64 workspace-darwin-amd64-upx

upx -9 workspace-linux-amd64-upx
upx -9 workspace-darwin-amd64-upx

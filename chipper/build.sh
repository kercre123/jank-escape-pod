#!/bin/bash
echo "Building chipper..."
if [[ $(arch) == "aarch64" ]]; then
   ARCH=arm64
elif [[ $(arch) == "x86_64" ]]; then
   ARCH=amd64
fi
CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} /usr/local/go/bin/go build \
-ldflags "-w -s -extldflags "-static"" \
-trimpath \
-o chipper cmd/main.go
echo "Built chipper!"

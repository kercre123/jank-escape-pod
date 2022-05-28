#!/bin/bash
if [[ $(arch) == "armv7l" ]]; then
	mkdir -p build
	echo "Building vic-cloud (direct because host arch is armv7l)..."
  	/usr/local/go/bin/go build  \
	-tags nolibopusfile,vicos \
	--trimpath \
	-ldflags '-w -s -linkmode internal -extldflags "-static" -r /anki/lib' \
	-o build/vic-cloud \
	cloud/main.go \
	cloud/cert.go
else
	echo "Building vic-cloud (docker)..."
	docker build -t armbuilder docker-builder/.
	echo `/usr/local/go/bin/go version` && cd $(PWD) && /usr/local/go/bin/go mod download
	mkdir -p build
	docker container run  \
	-v "$(PWD)":/go/src/digital-dream-labs/vector-cloud \
	-v $(GOPATH)/pkg/mod:/go/pkg/mod \
	-w /go/src/digital-dream-labs/vector-cloud \
	--user $(UID):$(GID) \
	armbuilder \
	go build  \
	-tags nolibopusfile,vicos \
	--trimpath \
	-ldflags '-w -s -linkmode internal -extldflags "-static" -r /anki/lib' \
	-o build/vic-cloud \
	cloud/main.go \
	cloud/cert.go

	docker container run \
	-v "$(PWD)":/go/src/digital-dream-labs/vector-cloud \
	-v $(GOPATH)/pkg/mod:/go/pkg/mod \
	-w /go/src/digital-dream-labs/vector-cloud \
	--user $(UID):$(GID) \
	armbuilder \
	upx build/vic-cloud
fi

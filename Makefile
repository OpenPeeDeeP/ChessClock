CC_VERSION=v0.1.0
CCD_VERSION=v0.1.0

build: dep proto build_cc build_ccd

build_cc:
	@go build -ldflags "-X main.version=${CC_VERSION}" -o build/cc ./cmd/cc

build_ccd:
	@go build -ldflags "-X main.version=${CCD_VERSION}" -o build/ccd ./cmd/ccd

build_clean:
	@rm -r build

dep:
	@dep ensure

proto:
	@protoc -I chessclock/ chessclock/chessclock.proto --go_out=plugins=grpc:chessclock

proto_clean:
	@rm chessclock/chessclock.pb.go

clean: build_clean

.PHONY: proto proto_clean build_cc build_ccd build build_clean clean dep install

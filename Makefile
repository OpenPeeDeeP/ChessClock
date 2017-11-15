build: proto build_cc build_ccd

build_cc:
	@go build -o build/cc ./cmd/cc

build_ccd:
	@go build -o build/ccd ./cmd/ccd

build_clean:
	@rm -r build

proto:
	@protoc -I chessclock/ chessclock/chessclock.proto --go_out=plugins=grpc:chessclock

proto_clean:
	@rm chessclock/chessclock.pb.go

clean: build_clean

.PHONY: proto proto_clean build_cc build_ccd build build_clean clean

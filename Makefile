
GO_MOD="github.com/sourcenetwork/source-zanzibar"

.PHONY: build
build:
	go build -gcflags="-e" -v -o build/auth-cli cmd/auth-cli/main.go

.PHONY: proto
proto:
	docker run --rm -it --workdir /app/proto -v ${PWD}:/app ghcr.io/cosmos/proto-builder buf generate


.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w $(fd '\.go$')

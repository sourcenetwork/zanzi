
GO_MOD="github.com/sourcenetwork/zanzi"

.PHONY: build
build:
	go build -gcflags="-e" -v -o build/auth-cli cmd/auth-cli/main.go

.PHONY: proto
proto:
	docker image build --file proto/Dockerfile --tag zanzi-proto-builder:latest proto/
	docker run --rm -it --workdir /app/proto -v $(PWD):/app zanzi-proto-builder:latest buf generate --verbose
	go mod tidy



.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w $(fd '\.go$')

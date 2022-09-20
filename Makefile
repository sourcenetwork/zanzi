
GO_MOD="github.com/sourcenetwork/source-zanzibar"

.PHONY: build
build:
	go build -o build/auth-cli cmd/auth-cli/main.go

.PHONY: proto
proto: proto/*/*.proto
	protoc --go_out=. --go_opt=module=${GO_MOD} proto/*/*.proto

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w $(fd '\.go$')

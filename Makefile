
GO_MOD="github.com/sourcenetwork/source-zanzibar"

.PHONY: build
build:
	go build -o build/auth-cli cmd/auth-cli/main.go

.PHONY: proto
proto:
	protoc --go_out=. -Iproto --go_opt=module=${GO_MOD} $$(find . -iname '*.proto' | sort)

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w $(fd '\.go$')

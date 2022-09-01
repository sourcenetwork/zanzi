
GO_MOD="github.com/sourcenetwork/source-zanzibar"

.PHONY: proto
proto: proto/*/*.proto
	protoc --go_out=. --go_opt=module=${GO_MOD} proto/*/*.proto

.PHONY: test
test:
	go test ./...

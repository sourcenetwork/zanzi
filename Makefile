
GO_MOD="github.com/sourcenetwork/zanzi"

.PHONY: build
build: zanzi-cli zanzid

.PHONY: zanzi-cli
zanzi-cli:
	go build -gcflags="-e" -v -o build/zanzi-cli cmd/zanzi-cli/main.go

.PHONY: zanzid
zanzid:
	go build -o build/zanzid cmd/zanzid/main.go

.PHONY: proto
proto:
	./proto/build-docker.sh
	docker run --rm -it --workdir /app/proto --user $(UID):$(GID) -v $(PWD):/app zanzi-proto-builder:latest buf generate --verbose
	#sudo mv zanzi/api/* pkg/api/
	#sudo mv zanzi/core/* pkg/core/
	#sudo rmdir zanzi/api
	#sudo rmdir zanzi/core
	#sudo rmdir zanzi
	#go mod tidy

.PHONY: fmt-proto
fmt-proto:
	docker image build --file proto/Dockerfile --tag zanzi-proto-builder:latest proto/
	docker run --rm -it --workdir /app/proto --user $(UID):$(GID) -v $(PWD):/app zanzi-proto-builder:latest buf format


.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: clean
clean:
	rm build/*

.PHONY: swagger
swagger:
	docker image build --file openapi/nginx.dockerfile -t zanzi-nginx openapi
	echo Nginx listening at port 7777
	echo Swagger accessible at endpoints /relation_graph/ and /policy/ 
	docker run --rm --publish 7777:80 zanzi-nginx

.PHONY: example
example:
	go build -o build/example example/embedded/main.go

.PHONY: docs
docs:
	docker image build --file devtools/docs.dockerfile -t zanzi-docs .
	echo "starting pkgsite in local port 7778"
	docker run --rm --publish 7778:8080 --volume .:/app zanzi-docs

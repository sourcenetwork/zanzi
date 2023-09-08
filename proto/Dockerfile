## To test locally:
# docker build --pull --rm -f "contrib/devtools/Dockerfile" -t cosmossdk-proto:latest "contrib/devtools"
# docker run --rm -v $(pwd):/workspace --workdir /workspace cosmossdk-proto sh ./scripts/protocgen.sh

# Source: https://github.com/cosmos/cosmos-sdk/blob/v0.46.7/contrib/devtools/Dockerfile

FROM golang:1.18-alpine

ARG PROTOBUF_VERSION=1.28.0 
ARG GRPC_GATEWAY_VERSION=2.15.0
ARG GEN_COBRA_VERSION=1.1.0
ARG GEN_GO_GRPC_VERSION=latest
ARG BUF_VERSION=1.25.1
ARG VALIDATE_VERSION=1.0.2

RUN apk update &&\
    apk add --no-cache git curl &&\
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOBUF_VERSION} &&\
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${GEN_GO_GRPC_VERSION} &&\
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v${GRPC_GATEWAY_VERSION} &&\
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v${GRPC_GATEWAY_VERSION} &&\
    go install github.com/NathanBaulch/protoc-gen-cobra@v${GEN_COBRA_VERSION} &&\
    go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@v${VALIDATE_VERSION} &&\
    curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(uname -s)-$(uname -m)" -o "/usr/local/bin/buf" &&\
    chmod +x "/usr/local/bin/buf"

ENV PATH="/usr/local/bin:$PATH"

FROM FROM bufbuild/buf:1.17.0

LABEL description="image with buf and google's Go generation plugin"


ENV PATH=/root/go/bin:$PATH

RUN apk update &&\
    apk add --no-cache go &&\
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

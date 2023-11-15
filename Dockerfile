# syntax=docker/dockerfile:1

# Build stage
FROM docker.io/golang:1.18 AS BUILD

WORKDIR /app

COPY go.mod go.sum Makefile .git ./
COPY . .
RUN make build

# Run stage
FROM gcr.io/distroless/base-debian11
WORKDIR /

COPY --from=BUILD /app/build/zanzid /zanzid

EXPOSE 8080

CMD ["/zanzid"]

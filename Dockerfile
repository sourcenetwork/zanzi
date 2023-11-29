# syntax=docker/dockerfile:1

# Build stage
FROM docker.io/golang:1.18 AS BUILD

WORKDIR /app

COPY go.mod go.sum raccoondb ./ 
RUN go mod download
COPY . .
RUN go build -o build/zanzid cmd/zanzid/main.go

# Run stage
FROM gcr.io/distroless/base-debian11
WORKDIR /

COPY --from=BUILD /app/build/zanzid /zanzid

EXPOSE 8080
EXPOSE 8090

CMD ["/zanzid"]

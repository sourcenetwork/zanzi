FROM golang:1.20-alpine

WORKDIR /app

RUN go install golang.org/x/pkgsite/cmd/pkgsite@latest

COPY . /app

EXPOSE 8080

CMD ["pkgsite", "-http", "0.0.0.0:8080"]

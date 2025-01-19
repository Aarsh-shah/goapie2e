FROM golang:1.23.1-bullseye as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && \
    go mod verify

COPY . ./

RUN go build -o . ./...

RUN echo here

FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install ca-certificates procps -y && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=build /app/teste2e ./

ENTRYPOINT ["./teste2e"]

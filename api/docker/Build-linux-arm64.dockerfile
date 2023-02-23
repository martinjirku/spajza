## Build
FROM golang:1.20.1 as build

WORKDIR /api

ADD * ../

RUN env GOOS=darwin GOARCH=linux go build -o zasobar-api ./cmd/server/main.go

## Deploy
FROM alpine:3.17.2

WORKDIR /

COPY --from=build /api/zasobar-api  /usr/bin

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["docker-gs-ping"]
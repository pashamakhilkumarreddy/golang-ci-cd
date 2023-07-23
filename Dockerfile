FROM golang:1.20.6-alpine as build

ARG VERSION=dev

WORKDIR /go/src/app
COPY . .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go

FROM debian:buster-slim
COPY --from=build /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD [ "main" ]
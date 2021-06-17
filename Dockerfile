# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/

FROM golang:1.16-alpine as builder
RUN apk add --no-cache git
RUN apk add --no-cache git

ENV GO111MODULE on
ENV GOPROXY "https://proxy.golang.org/"

RUN mkdir -p /src/apple-wwdc21-demo
WORKDIR /src/apple-wwdc21-demo

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

ADD . .
RUN go build -o /bin/app

######## Dockerize
FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .
COPY ./assets /bin/assets
COPY ./templates /bin/templates

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production

EXPOSE 9000

ENTRYPOINT /bin/app

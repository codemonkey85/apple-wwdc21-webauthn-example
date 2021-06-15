# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM golang:1.16-alpine as builder
RUN apk add --no-cache git
RUN apk add --no-cache git

ENV GO111MODULE on
ENV GOPROXY "https://proxy.golang.org/"
ENV GOPRIVATE gitlab.com/hanko
ARG GITLAB_ACCESS_TOKEN
ARG GITLAB_USER

# Acces to private repos
RUN git config --global url."https://${GITLAB_USER}:${GITLAB_ACCESS_TOKEN}@gitlab.com/hanko".insteadOf "https://gitlab.com/hanko"


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
COPY ./assets /bin/asstes

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production

EXPOSE 9000

ENTRYPOINT /bin/app

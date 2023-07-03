# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.19 AS base
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src
COPY ./ ./
RUN go mod download

# build the application
FROM base AS build
RUN go build -ldflags="-w -s" -o /app/main ./internal/main/main.go

# Import the binary from build stage
FROM gcr.io/distroless/static:nonroot as prd
COPY --from=build /app/main /

# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/main"]

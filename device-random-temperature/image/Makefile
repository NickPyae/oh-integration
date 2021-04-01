# Copyright 2021 Dell Inc, or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

.PHONY: build test clean prepare update docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-random-temperature

.PHONY: $(MICROSERVICES)

DOCKERS=device_random_temperature_go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X device-random-temperature.Version=$(VERSION)"

build: $(MICROSERVICES)
        $(GO) build ./...

cmd/device-random-temperature:
        $(GO) build $(GOFLAGS) -o $@ ./cmd

test:
        $(GO) test ./... -coverprofile=coverage.out
        $(GO) vet ./...
        gofmt -l .
        [ "`gofmt -l .`" = "" ]

clean:
        rm -f $(MICROSERVICES)

docker: $(DOCKERS)

device_random_temperature_go:
        docker build \
                --label "git_sha=$(GIT_SHA)" \
                -t device-random-temperature-go:$(GIT_SHA) \
                -t device-random-temperature-go:$(VERSION)-dev \
                .

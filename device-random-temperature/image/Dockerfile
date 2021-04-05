# Copyright 2021 Dell Inc, or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

ARG BASE=golang:1.15-alpine3.12
FROM ${BASE} AS builder
ARG ALPINE_PKG_BASE="build-base git openssh-client"
ARG ALPINE_PKG_EXTRA=""

# Replicate the APK repository override.
# If it is no longer necessary to avoid the CDN mirros we should consider dropping this as it is brittle.
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
# Install our build time packages.
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR $GOPATH/src/HelloSally/device-random-temperature/image

COPY . .

# To run tests in the build container:
#   docker build --build-arg 'MAKE=build test' .
# This is handy of you do your Docker business on a Mac
ARG MAKE='make build'
RUN $MAKE

FROM alpine

ENV APP_PORT=49988
EXPOSE $APP_PORT

COPY --from=builder /go/src/HelloSally/device-random-temperature/image/cmd /
# COPY --from=builder /go/src/github.com/edgexfoundry/device-random/LICENSE /
# COPY --from=builder /go/src/github.com/edgexfoundry/device-random/Attribution.txt /

ENTRYPOINT ["/device-random-temperature"]
CMD ["--cp=consul://edgex-core-consul:8500", "--confdir=/res", "--registry"]

# Copyright 2021 Dell Inc, or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.15.7-alpine3.13

WORKDIR /app

COPY . /app/

RUN go mod tidy

EXPOSE 49989

ENTRYPOINT ["go","run","main.go"]
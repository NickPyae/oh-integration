FROM golang:alpine AS base

WORKDIR /email-alert

COPY go.mod go.sum main.go ./

RUN go mod tidy
RUN go build .

FROM alpine:latest

COPY --from=base /email-alert/email-alert /

ENTRYPOINT ["/email-alert"]
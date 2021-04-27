# Naive way of building image
FROM golang:1.15.7-alpine3.13

WORKDIR /app

COPY . /app/

RUN go mod tidy

EXPOSE 49990

ENTRYPOINT ["go","run","main.go"]

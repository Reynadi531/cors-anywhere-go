FROM golang:1.17-alpine

RUN apk update && \
  apk add musl-dev openssl git bash --no-cache

WORKDIR /app

COPY . . 

RUN go mod tidy
RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]

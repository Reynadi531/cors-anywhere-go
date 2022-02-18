FROM golang:1.17-alpine

RUN apk update && \
  apk add msul-dev openssl ca-certificate bash --no-cache

WORKDIR /app

COPY ../ .

RUN go mod tidy
RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]

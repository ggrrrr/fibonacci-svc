# ARG SERVICE
# FROM golang:1.21 AS builder
# FROM golang:1.21-alpine AS builder
FROM golang:1.21-alpine AS builder

# RUN apt update && apt install telnet curl
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY  . ./

RUN go build -o app main.go

FROM alpine:latest AS runtime
RUN apk add --no-cache curl
WORKDIR /app
COPY --from=builder /app/app /app/app

CMD ["/app/app"]

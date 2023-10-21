# Path: Dockerfile
# This is a generated file. Do not modify.
FROM golang:1.21 as builder

ARG SHA

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app cmd/main.go

FROM alpine:latest as release
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app/bin/app .
CMD ["./app"]

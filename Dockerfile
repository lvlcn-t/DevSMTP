FROM golang:1.21-alpine3.18 AS builder

RUN apk add --no-cache git

COPY . /tmp/server

WORKDIR /tmp/server

RUN go mod download

RUN go build -o bin/server -ldflags="-s -w" ./cmd/server/main.go

FROM alpine:3.18

ARG user_id=1000
RUN adduser -S $user_id -G root -u $user_id

COPY --from=builder --chown=$user_id:root /tmp/server/bin/server /app/devsmtp

EXPOSE 8080

USER $user_id

CMD ["/app/devsmtp"]

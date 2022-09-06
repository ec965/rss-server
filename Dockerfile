# syntax=docker/dockerfile:1
FROM golang:1.19-alpine as builder
WORKDIR /opt/rss-server/
RUN apk add build-base
COPY . .
RUN go build
RUN go build -o migrate cmd/migrate/main.go

FROM alpine:latest
WORKDIR /opt
COPY --from=builder /opt/rss-server/rss-server ./
COPY --from=builder /opt/rss-server/migrate ./
COPY --from=builder /opt/rss-server/db/migrations ./db/migrations
CMD ["./rss-server"]

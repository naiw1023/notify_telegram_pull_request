FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -v -o notify-telegram .

FROM alpine:latest

COPY --from=builder /app/notify_telegram_pull_request /notify_telegram_pull_request

ENTRYPOINT ["/notify_telegram_pull_request"]

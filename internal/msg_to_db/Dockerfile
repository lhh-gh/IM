FROM imbase AS builder

WORKDIR /build/internal/msg_to_db/

RUN go build -o main main.go

FROM golang:alpine

ARG SERVICE_DIR=/build/internal/msg_to_db

COPY --from=builder $SERVICE_DIR/etc/config.yaml /app/etc/config.yaml
COPY --from=builder $SERVICE_DIR/main          /app/main

WORKDIR /app

ENTRYPOINT ["/app/main"]
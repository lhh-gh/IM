FROM imbase AS builder

WORKDIR /build/internal/im_server/

RUN go build -o server server.go

FROM golang:alpine

ARG SERVICE_DIR=/build/internal/im_server

COPY --from=builder $SERVICE_DIR/etc/server.yaml /app/etc/server.yaml
COPY --from=builder $SERVICE_DIR/server          /app/server

WORKDIR /app

EXPOSE 8082

ENTRYPOINT ["/app/server"]
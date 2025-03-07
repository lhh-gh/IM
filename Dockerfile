FROM golang:alpine AS base

LABEL stage=gobuilder
LABEL authors="sillycat"

ENV CGO_ENABLED 0

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .
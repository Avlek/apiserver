FROM golang:1.11-alpine AS build

RUN go build cmd/apiserver/main.go

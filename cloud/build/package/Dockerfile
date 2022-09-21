# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ../../
COPY go.sum ../../
RUN go mod download

COPY ../../* ./

RUN go build -o ./cmd/cloud_service/cloud_service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /cloud_service /cloud_service

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/cloud_service"]
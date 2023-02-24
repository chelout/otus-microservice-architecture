FROM golang:1.19-alpine AS build

WORKDIR /app

COPY ./ ./

RUN go mod download && \
    CGO_ENABLED=0 go build -o bin/healthcheck -ldflags="-w -s" ./cmd/api

FROM alpine

WORKDIR /

COPY --from=build /app/bin/healthcheck /healthcheck

EXPOSE 8000

ENTRYPOINT ["/healthcheck"]
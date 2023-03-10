FROM golang:1.20-alpine AS build

WORKDIR /app

COPY ./ ./

RUN go mod download && \
    CGO_ENABLED=0 go build -o bin/api -ldflags="-w -s" ./cmd/api && \
    CGO_ENABLED=0 go build -o bin/migrate -ldflags="-w -s" ./cmd/console

FROM alpine

WORKDIR /

COPY --from=build /app/bin/api /api
COPY --from=build /app/bin/migrate /migrate
COPY --from=build /app/deployments/db/migrations /deployments/db/migrations

EXPOSE 8000

ENTRYPOINT ["/api"]

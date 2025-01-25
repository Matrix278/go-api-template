FROM golang:1.23.0-bookworm AS builder

ARG APP_PORT

WORKDIR /app

COPY . .

RUN go mod download \
    && CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./go-backend -a -ldflags "-s -w" -installsuffix cgo .

FROM alpine:latest AS application

COPY --from=builder /app/go-backend ./go-backend
COPY --from=builder /app/configuration/ ./configuration/
COPY --from=builder /app/docs/ ./docs/

EXPOSE ${APP_PORT}

ENTRYPOINT ["./go-backend"]

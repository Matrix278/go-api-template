FROM golang:1.24.2-bookworm AS builder

ARG APP_PORT

WORKDIR /app

COPY . .

RUN go mod download \
    && CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./go-backend -a -ldflags "-s -w" -installsuffix cgo .

FROM alpine:latest AS application

RUN apk add --no-cache git go \
    && go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app

# Copy compiled backend
COPY --from=builder /app/go-backend ./go-backend
COPY --from=builder /app/configuration/ ./configuration/
COPY --from=builder /app/docs/ ./docs/
COPY --from=builder /app/migrations/ ./migrations/

EXPOSE ${APP_PORT}

# Ensure goose is available in PATH
ENV PATH="/root/go/bin:$PATH"

# Run migrations before starting the app
CMD ["sh", "-c", "goose -dir ./migrations postgres \"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable\" up && ./go-backend"]

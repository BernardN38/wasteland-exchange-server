# --- build stage ---
FROM golang:alpine AS builder

WORKDIR /app

# Install CA certs + git if you have private deps (safe even if you don't)
RUN apk add --no-cache ca-certificates git

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./...

# --- runtime stage ---
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/server /app/server

# Your app listens on 8080 by default in the example
EXPOSE 8080

# Optional: run as non-root (recommended)
RUN adduser -D -H appuser && chown -R appuser:appuser /app
USER appuser

ENTRYPOINT ["/app/server"]

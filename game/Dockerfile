FROM golang:alpine AS builder
WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /server .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /server /server
COPY  /common ./common
COPY  /saves ./saves

EXPOSE 8080
ENTRYPOINT ["/server"]

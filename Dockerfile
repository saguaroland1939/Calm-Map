FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o calm-map ./cmd/server

FROM alpine:3.20
COPY --from=builder /app/calm-map /usr/local/bin/calm-map
EXPOSE 8080
CMD ["calm-map"]

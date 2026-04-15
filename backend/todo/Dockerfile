# ── Build stage ────────────────────────────────────────────
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download deps first — cached unless go.mod/go.sum change
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ── Run stage ──────────────────────────────────────────────
FROM alpine:3.19

WORKDIR /app

# Certificates needed for HTTPS outbound calls
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./server"]
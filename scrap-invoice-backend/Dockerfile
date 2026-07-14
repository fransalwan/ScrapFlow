# === STAGE 1: Builder ===
FROM golang:1.24.4-alpine AS builder

# Tambahkan baris ini buat fix network/proxy
ENV GOPROXY=https://proxy.golang.org,direct
ENV CGO_ENABLED=0

# Install dependensi build
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod dan go.sum dulu buat cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi Go (CGO disabled biar static binary, bisa jalan di Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# === STAGE 2: Runner ===
FROM alpine:latest

# Install CA certificates (buat HTTPS request) dan timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /app/main .

# Expose port yang dipakai aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
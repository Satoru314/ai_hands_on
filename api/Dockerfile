# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

# Copy source code
COPY ./api .

# デバッグ: コピーされたファイル構造を確認
RUN echo "=== Copied files ===" && ls -la
RUN echo "=== cmd directory ===" && ls -la cmd/
RUN echo "=== server directory ===" && ls -la cmd/server/

# Build the application
RUN go build -o server cmd/server/main.go

# デバッグ: ビルド結果を確認
RUN echo "=== Build result ===" && ls -la server
RUN echo "=== Server permissions ===" && ls -la server

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/server .

# デバッグ: runtime環境確認
RUN echo "=== Runtime files ===" && ls -la
RUN echo "=== Server in runtime ===" && ls -la server

# Expose port
EXPOSE 8080

# Run the binary
CMD ["/root/server"]
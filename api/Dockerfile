# Development stage with Air
FROM golang:1.23-alpine AS dev

WORKDIR /app

# Install Air for hot reloading (compatible with Go 1.23)
RUN go install github.com/cosmtrek/air@v1.40.4

# Copy go mod files
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

# Copy source code
COPY ./api .

# Expose port
EXPOSE 8080

# Use Air for hot reloading
CMD ["air", "-c", ".air.toml"]
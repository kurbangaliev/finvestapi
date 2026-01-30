# Stage 1: Build the Go binary
FROM golang:1.25-bookworm AS builder

LABEL authors="denislam"
LABEL maintainer="denislam_"

WORKDIR /app

# Copy Go modules and download dependencies to leverage Docker's build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application with CGO_ENABLED=0 to produce a static binary
# and optional flags to strip debug information for a smaller size
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/backend-server ./cmd/backend-server.go
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/frontend-server ./cmd/frontend-server.go

# Stage 2: Create a minimal production image
FROM debian:bookworm-slim
# Install necessary certificates if the app makes outgoing HTTPS requests
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/* /app/
COPY web/ /app/web/
COPY scripts/start.sh .
RUN chmod +x start.sh

# Optional: Run as a non-root user for better security
RUN groupadd -r appgroup && useradd -r -g appgroup appuser
USER appuser

# Expose the port your application listens on
EXPOSE 8443
EXPOSE 8081

# Command to run the executable when the container starts
CMD ["./start.sh"]
# Dockerfile (Corrected with templ generate and multi-binary build)

# --- Stage 1: Builder ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install the templ CLI tool.
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy dependency files and download them to leverage Docker layer caching.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Run templ generate to create the _templ.go files.
RUN /go/bin/templ generate

# --- CHANGE: Build all commands in the /cmd directory ---
RUN for d in cmd/*; do go build -o /app/$(basename $d) ./$d; done

# --- Stage 2: Final Image ---
# This stage creates the small, clean final image.
FROM alpine:latest

WORKDIR /app

# Copy necessary assets and the final compiled binaries from the builder stage.
COPY --from=builder /app/web ./web
COPY --from=builder /app/server .
COPY --from=builder /app/server .
COPY --from=builder /app/cli .

# Note: We do not need to copy the migrations/ folder because the .sql files
# are now embedded directly inside the /app/cli binary.

EXPOSE 8080

# The default command is to run the web server. This will be overridden
# in our deploy-compose.yml for production startup.
CMD ["/app/server"]
